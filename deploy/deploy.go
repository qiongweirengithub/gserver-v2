package main

import (
	"flag"
	"fmt"
	"log"
	"gserver.v2/admin/dbs"
	"gserver.v2/deploy/application"
	"io/ioutil"
	"os/exec"
	"os"
	"path/filepath"	
	"strings"
	"errors"
)

var (
	service *string
	host *string
	port *string
	pid *string
	websocketport *string
	ci_dir string
	container_id *string
	roomid *string

	// project_url string = "git@github.com:qiongweirengithub/gserver-v2.git"
	project_url string = "/home/qiongwei/mycode/goprj/gserver.v3/"

) 



// 检查当前目录是否可用
func checkDirAvaliable(dir string) bool {
	dirFile, _ := ioutil.ReadDir(dir)
    if len(dirFile) == 0 {
		return true;
    } else {
        fmt.Println(dir + " is not empty dir!")
		return false
    }
}


func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}


// 删除应用               		deploy -service=kill -containerid=查询数据库或者docker ps -l
// 部署 g-web-restapi    		deploy -service=g-web-restapi -port=8090
// 部署 g-gate-connectionsvc    deploy -service=g-gate-connectionsvc -websocketport=3653
// 部署 g-authservice           deploy -service=g-authservice
// 部署 g-battleroomsvc         deploy -service=g-battleroomsvc -roomid=12345

// 
func main() {

	service = flag.String("service", "no service", "it's user send message[help message]")
	
	host = flag.String("host", "default host", "it's user send host[help message]")

	port = flag.String("port", "default port", "it's user send port[help message]")

	roomid = flag.String("roomid", "", "it's user send pid[help message]")

	pid = flag.String("pid", "default pid", "it's user send pid[help message]")

	websocketport = flag.String("websocketport", "default websocketport", "it's user send websocketport[help message]")

	container_id = flag.String("containerid", "", "it's user send pid[help message]")


	flag.Parse();


	var err error

	// 下线应用
	if *service == "kill" || *service == "rm" {

		fmt.Println(application.YellowBg, "start killing svc: ", *service, application.Reset)

		if *container_id == "" || len(*container_id) <20 {
			fmt.Println(application.Red, "wrong container id: ", *container_id, application.Reset)
			return
		}
		fmt.Println(application.Green, "killing container : ", *container_id, application.Reset)
		_, err = application.ExecCmd("docker", "container", "rm", "-f", *container_id)
		if err != nil {
			fmt.Println(application.Red, "kill container fail: ", *container_id, err, application.Reset)
		} else {
			fmt.Println(application.Green, "container:", *container_id, " kill success", application.Reset)
		}
		dbs.DeleteService(*container_id)		
		return
	}

	// ================ 发布应用 ====================

	fmt.Println(application.YellowBg, "start deploy svc: ", *service, application.Reset)

	// 创建临时文件夹
	
	tmpDir := "/tmp/gserver-" + application.RandString(20)
	fmt.Println(application.MagentaBg, "create tmp dir ", tmpDir, application.Reset)
	application.ExecCmd("mkdir", tmpDir)

	// 清理tmp数据
	// defer application.ExecCmd("rm", "-fr",  tmpDir)

	ci_dir = tmpDir

	// 拉取项目
	fmt.Println(application.MagentaBg, "loading project", project_url, application.Reset)
	application.ExecCICmd(ci_dir, "git", "clone", project_url)
	
	// 切换到git分支
	// ci_dir = ci_dir + "/gserver-v2"
	ci_dir = ci_dir + "/gserver.v3"
	_, err = application.ExecCICmd(ci_dir, "git", "checkout", "main-v3.0")
	if err != nil {
		fmt.Println(application.Red, "branch checkout fail: ", *container_id, err, application.Reset)
		return
	}
	
	// 构建程序，并生成应用文件
	fmt.Println(application.MagentaBg, "building project", project_url, application.Reset)
	_, err = application.ExecCICmd(ci_dir, "/bin/sh", "./build.sh")
	if err != nil {
		fmt.Println(application.Red, "project build fail: ", *container_id, err, application.Reset)
		return
	}

	// 拷贝环境变量到构建目录
	fmt.Println(application.MagentaBg, "copying project env", project_url, application.Reset)
	_, err = application.ExecCICmd(ci_dir, "cp", "./bin/conf/env.json", "./")
	if err != nil {
		fmt.Println(application.Red, "copying project env fail", project_url, err, application.Reset)
		return
	}

	// 发布应用类型为 g-web-restapi
	if *service == "g-web-restapi" {
		// 部署 g-web-restapi 
		//  TODO 检查端口
		fmt.Println(application.MagentaBg, "deploy service", *service, application.Reset)
		containerId, serviceId, err := application.DeployingGWebRestApi(ci_dir, *service, *port)
		if err != nil {
			fmt.Println(application.Red, "deploy service ", *service, " fail", err, application.Reset)
			return
		}
		_, err = dbs.CreateGservice(serviceId, 1, "127.0.0.1", *port, "image", containerId, "g_web");

		if err != nil {
			log.Fatalln(err)
			fmt.Println(application.Red, "deploy service ", *service, " fail", err, application.Reset)
			return
		}
	}

	if *service == "g-authservice" {
		// 部署 g-authservice
		//  TODO 检查端口
		containerId, serviceId, err := application.DeployingGAuthServuce(ci_dir, *service)
		if err != nil {
			log.Fatalln(err)
			return
		}

		fmt.Println("record g-auth service: ", containerId, "-", serviceId)
		_, err = dbs.CreateGservice(serviceId, 1, "127.0.0.1", *port, "image", containerId, "g_auth");

		if err != nil {
			log.Fatalln(err)
			return
		}
		return
	}


	if *service == "g-gate-connectionsvc" {
		// 部署 g-gate-connectionsvc 
		//  TODO 检查端口
		containerId, serviceId, err := application.DeployingGGateconnectionsvc(ci_dir, *service, *websocketport)
		if err != nil {
			log.Fatalln(err)
			return
		}

		fmt.Println("record gate connection svc : ", containerId, "-", serviceId)
		_, err = dbs.CreateGservice(serviceId, 1, "127.0.0.1", *websocketport, "image", containerId, "g-gate-connectionsvc");

		if err != nil {
			log.Fatalln(err)
			return
		}
		return
	}

	if *service == "g-battleroomsvc" {

		fmt.Println(application.MagentaBg, *service, "is ready to deploy", *service,  application.Reset)

		containerId, serviceId, err := application.DeployingGBattleRoomsvc(ci_dir, *service, *roomid)
		if err != nil {
			fmt.Println(application.Red, "deploying svc", *service, " error", application.Reset)
			return
		}
		_, err = dbs.CreateGservice(serviceId, 1, "127.0.0.1", *websocketport, "image", containerId, "g-gate-connectionsvc");

		if err != nil {
			fmt.Println(application.Red, "recording svc", *service, " error", application.Reset)
			return
		}
		return
	}
}
