package main

import (
	"bytes"
	"flag"
	"fmt"
	"os/exec"
	"math/rand"
	"time"
)

var (
	service *string
	host *string
	port *string
	pid *string
	websocketport *string
	workdir *string
	r *rand.Rand
) 


func init() {
    r = rand.New(rand.NewSource(time.Now().Unix()))
}

func RandString(len int) string {
    bytes := make([]byte, len)
    for i := 0; i < len; i++ {
        b := r.Intn(26) + 65
        bytes[i] = byte(b)
    }
    return string(bytes)
}

func execCDCmd(name string, arg ...string) (_ string, err error) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(name, arg...)
	cmd.Dir = "/home/qiongwei/myapp/k8s/app/" 

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	errs:=cmd.Run()
	if errs != nil {
		fmt.Print(errs)
	}
	a:= out.Bytes();
	fmt.Print(string(a))

	return string(a), nil
}


func execCICmd(name string, arg ...string) (_ string, err error) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(name, arg...)
	cmd.Dir = *workdir 

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	errs:=cmd.Run()
	if errs != nil {
		fmt.Print(errs)
	}
	a:= out.Bytes();
	fmt.Print(string(a))

	return string(a), nil
}

func main() {

	workdir = flag.String("wd", "default message", "it's user send workdir[help message]")

	service = flag.String("service", "default message", "it's user send message[help message]")
	
	host = flag.String("host", "default host", "it's user send host[help message]")

	port = flag.String("port", "default port", "it's user send port[help message]")

	pid = flag.String("pid", "default pid", "it's user send pid[help message]")

	websocketport = flag.String("websocketport", "default websocketport", "it's user send websocketport[help message]")

	flag.Parse();

	fmt.Println(*service)

	var err error

	_, err = execCICmd("pwd")


	if err != nil {
		fmt.Print(err)
		return
	}
	_, err = execCICmd("pwd")
	if err != nil {
		fmt.Print(err)
	}

	_, err = execCICmd("ls", "-lh")

	
	_, err = execCICmd("/bin/sh", "./build.sh")
	if err != nil {
		fmt.Print(err)
		return
	}


	_,err = execCICmd("cp", "./gserverv2", "/home/qiongwei/myapp/k8s/app")
	if err != nil {
		fmt.Print(err)
	}

	_, err = execCICmd("cp", "./bin/conf/env.json", "/home/qiongwei/myapp/k8s/app")
	if err != nil {
		fmt.Print(err)
	}


	if *service == "g-web-restapi" {

		// check param
		// 1. port 是否被占用

		fmt.Println("deploying web")


		_, err = execCICmd("cp", "./bin/conf/web-release/server-port.json", "/home/qiongwei/myapp/k8s/app/server-web.json")
		if err != nil {
			fmt.Print(err)
			return
		}

		_, err = execCICmd("sed", "-i","s/PORT/" + *port + "/g", "/home/qiongwei/myapp/k8s/app/server-web.json")
		if err != nil {
			fmt.Print(err)
			return
		}

		fmt.Println("web app config json succ")
		_, err = execCICmd("cp", "./bin/conf/web-release/web.Dockerfile", "/home/qiongwei/myapp/k8s/app")
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Println("web docker file succ")

		_, err = execCDCmd("docker", "build", "-t", "g-web-restapi", "-f", "/home/qiongwei/myapp/k8s/app/web.Dockerfile", ".")
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Println("web docker build succ")


		fmt.Println("runing with ports: ", *port)

		appId := "-" + RandString(6)
		containerName := *service + appId
		containerPidFileName := *service+appId+".pid"

		_, err = execCDCmd("docker","run","-d","--name", containerName, "--cidfile" , containerPidFileName,"-p"+*port+":"+*port, "g-web-restapi")
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Println("recording web app: ", containerName)

		// record to db

		return
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
