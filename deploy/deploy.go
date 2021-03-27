package main

import (
	"flag"
	"fmt"
	"log"
	"gserver.v2/admin/dbs"
	"gserver.v2/deploy/application"
)

var (
	service *string
	host *string
	port *string
	pid *string
	websocketport *string
	workdir *string
	container_id *string

) 


// 删除应用               deploy -service=kill -containerid=查询数据库或者docker ps -l
// 部署 g-web-restapi    deploy -service=g-web-restapi -wd=/home/qiongwei/mycode/goprj/gserver.v3/ -port=8090
// 
func main() {

	workdir = flag.String("wd", ".", "it's user send workdir[help message]")

	service = flag.String("service", "no service", "it's user send message[help message]")
	
	host = flag.String("host", "default host", "it's user send host[help message]")

	port = flag.String("port", "default port", "it's user send port[help message]")

	pid = flag.String("pid", "default pid", "it's user send pid[help message]")

	websocketport = flag.String("websocketport", "default websocketport", "it's user send websocketport[help message]")

	container_id = flag.String("containerid", "", "it's user send pid[help message]")

	flag.Parse();

	fmt.Println(*service)

	var err error


	// 发布应用类型为 g-web-restapi
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



	// 构建程序，并生成应用文件
	_, err = application.ExecCICmd(*workdir, "/bin/sh", "./build.sh")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 拷贝应用到k8s目录
	_,err = application.ExecCICmd(*workdir, "cp", "./gserverv2", "/home/qiongwei/myapp/k8s/app")
	if err != nil {
		fmt.Println(err)
	}

	// 拷贝环境变量到k8s目录
	_, err = application.ExecCICmd(*workdir, "cp", "./bin/conf/env.json", "/home/qiongwei/myapp/k8s/app")
	if err != nil {
		fmt.Println(err)
	}

	// 发布应用类型为 g-web-restapi
	if *service == "g-web-restapi" {
		// 部署 g-web-restapi 
		containerId, serviceId, err := application.DeployingGWebRestApi(*workdir, *service, *port)
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
