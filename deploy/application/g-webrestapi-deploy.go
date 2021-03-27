package application

import (
	"fmt"
)


func DeployingGWebRestApi(wd string, service string, port string)  (string, string, error) {

	// check param
	// 1. port 是否被占用

	fmt.Println("deploying web")

	_, err := ExecCICmd(wd, "cp", "./bin/conf/web-release/server-port.json", "/home/qiongwei/myapp/k8s/app/server-web.json")
	if err != nil {
		fmt.Print(err)
		return "", "", err
	}

	_, err = ExecCICmd(wd, "sed", "-i","s/PORT/" + port + "/g", "/home/qiongwei/myapp/k8s/app/server-web.json")
	if err != nil {
		fmt.Print(err)
		return "","", err
	}

	fmt.Println("web app config json succ")
	_, err = ExecCICmd(wd, "cp", "./bin/conf/web-release/web.Dockerfile", "/home/qiongwei/myapp/k8s/app")
	if err != nil {
		fmt.Print(err)
		return "","", err
	}
	fmt.Println("web docker file succ")

	_, err = ExecCDCmd("docker", "build", "-t", "g-web-restapi", "-f", "/home/qiongwei/myapp/k8s/app/web.Dockerfile", ".")
	if err != nil {
		fmt.Print(err)
		return "","", err
	}
	fmt.Println("web docker build succ")

	appId := "-" + RandString(6)
	containerName := service + appId
	containerPidFileName := service+appId+".pid"

	_, err = ExecCDCmd("docker","run","-d","--name", containerName, "--cidfile" , containerPidFileName,"-p"+port+":"+port, "g-web-restapi")
	if err != nil {
		fmt.Print(err)
		return "","", err
	}

	fmt.Println("runing with ports: ", port)

	pdiFile := CD_Dir + containerPidFileName
	pid := GetPid(pdiFile)
	fmt.Println("pid file:", pdiFile, ", pid: ", pid)

	// record to db
	return pid, containerName, nil

}