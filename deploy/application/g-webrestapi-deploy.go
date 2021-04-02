package application

import (
	"fmt"
)


func DeployingGWebRestApi(wd string, service string, port string)  (string, string, error) {

	// check param
	// 1. port 是否被占用

	fmt.Println("deploying web")

	_, err := ExecCICmd(wd, "cp", "./bin/conf/web-release/server.json", "./")
	if err != nil {
		fmt.Print(err)
		return "", "", err
	}

	_, err = ExecCICmd(wd, "sed", "-i","s/PORT/" + port + "/g", "./server.json")
	if err != nil {
		fmt.Print(err)
		return "","", err
	}


	_, err = ExecCICmd(wd, "cp", "./bin/conf/web-release/web.Dockerfile", wd)
	if err != nil {
		fmt.Print(err)
		return "","", err
	}

	_, err = ExecCDCmd(wd, "docker", "build", "-t", service, "-f", wd + "/web.Dockerfile", ".")
	if err != nil {
		fmt.Print(err)
		return "","", err
	}
	fmt.Println("web docker build succ")

	appId := "-" + RandString(6)
	containerName := service + appId
	containerPidFileName := service+appId+".pid"

	_, err = ExecCDCmd(wd, "docker","run","-d","--name", containerName, "--cidfile" , containerPidFileName,"-p"+port+":"+port, service)
	if err != nil {
		fmt.Print(err)
		return "","", err
	}

	fmt.Println("runing with ports: ", port)

	pdiFile := wd + "/" + containerPidFileName
	pid := GetPid(pdiFile)
	fmt.Println("pid file:", pdiFile, ", pid: ", pid)

	// record to db
	return pid, containerName, nil

}