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
	project_url string = "git@github.com:qiongweirengithub/gserver-v2.git"
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


// 删除应用               deploy -service=kill -containerid=查询数据库或者docker ps -l
// 部署 g-web-restapi    deploy -service=g-web-restapi -wd=/home/qiongwei/mycode/goprj/gserver.v3/ -port=8090
// 
func main() {

	service = flag.String("service", "no service", "it's user send message[help message]")
	
	host = flag.String("host", "default host", "it's user send host[help message]")

	port = flag.String("port", "default port", "it's user send port[help message]")

	pid = flag.String("pid", "default pid", "it's user send pid[help message]")

	websocketport = flag.String("websocketport", "default websocketport", "it's user send websocketport[help message]")

	container_id = flag.String("containerid", "", "it's user send pid[help message]")

	flag.Parse();

	fmt.Println(*service)

	var err error

	// 下线应用
	if *service == "kill" || *service == "rm" {
		if *container_id == "" || len(*container_id) <20 {
			fmt.Println("wrong container id: ", *container_id)
			return
		}
		_, err = application.ExecCmd("docker", "container", "rm", "-f", *container_id)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("container:", *container_id, " kill success")
		}
		dbs.DeleteService(*container_id)		
		return
	}

	// ================ 发布应用 ====================

	// 创建临时文件夹
	tmpDir := "/tmp/gserver-" + application.RandString(20)
	fmt.Println("creating tmp dir : ",  tmpDir)
	application.ExecCmd("mkdir", tmpDir)

	// 清理tmp数据
	// defer application.ExecCmd("rm", "-fr",  tmpDir)

	ci_dir = tmpDir

	// 拉取项目
	application.ExecCICmd(ci_dir, "git", "clone", project_url)
	
	// 切换到git目录
	ci_dir = ci_dir + "/gserver-v2"
	application.ExecCICmd(ci_dir, "ls", "-lh")
	application.ExecCICmd(ci_dir, "git", "checkout", "main-v3.0")
	application.ExecCICmd(ci_dir, "git", "branch")

	
	// 构建程序，并生成应用文件
	_, err = application.ExecCICmd(ci_dir, "/bin/sh", "./build.sh")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 拷贝环境变量到构建目录
	_, err = application.ExecCICmd(ci_dir, "cp", "./bin/conf/env.json", "./")
	if err != nil {
		fmt.Println(err)
	}

	// 发布应用类型为 g-web-restapi
	if *service == "g-web-restapi" {
		// 部署 g-web-restapi 
		containerId, serviceId, err := application.DeployingGWebRestApi(ci_dir, *service, *port)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("record web app: ", containerId, "-", serviceId)
		_, err = dbs.CreateGservice(serviceId, 1, "127.0.0.1", *port, "image", containerId, "g_web");

		if err != nil {
			log.Fatalln(err)
		}
	}

	if *service == "auth" {
		return
	}


	if *service == "gate" {
		
		return
	}

	if *service == "web" {

		return
	}

	if *service == "web" {

		return
	}

	fmt.Println("deploying service: ", *service)
	fmt.Println("deploying host: ", *host)
}
