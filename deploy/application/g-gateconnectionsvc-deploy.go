package application

import (
	"fmt"
)


func DeployingGGateconnectionsvc(wd string, service string, wsaddr string)  (string, string, error) {

	// check param
	// 1. port 是否被占用

	fmt.Println("deploying gate connection svc")

	_, err := ExecCICmd(wd, "cp", "./bin/conf/gate-release/server.json", "./")
	if err != nil {
		fmt.Print(err)
		return "", "", err
	}

	_, err = ExecCICmd(wd, "sed", "-i","s/WSADDR/" + wsaddr + "/g", "./server.json")
	if err != nil {
		fmt.Print(err)
		return "","", err
	}


	_, err = ExecCICmd(wd, "cp", "./bin/conf/gate-release/gate.Dockerfile", wd)
	if err != nil {
		fmt.Print(err)
		return "","", err
	}

	_, err = ExecCDCmd(wd, "docker", "build", "-t", service, "-f", wd + "/gate.Dockerfile", ".")
	if err != nil {
		fmt.Print(err)
		return "","", err
	}
	fmt.Println("gate connection svc docker build succ")

	appId := "-" + RandString(6)
	containerName := service + appId
	containerPidFileName := service+appId+".pid"

	_, err = ExecCDCmd(wd, "docker","run","-d","--name", containerName, "--cidfile" , containerPidFileName,"-p"+wsaddr+":"+wsaddr, service)

	if err != nil {
		fmt.Print(err)
		return "","", err
	}

	fmt.Println("gate connection svc running with ports: ", wsaddr)

	pdiFile := wd + "/" + containerPidFileName
	pid := GetPid(pdiFile)
	fmt.Println("pid file:", pdiFile, ", pid: ", pid)

	// record to db
	return pid, containerName, nil

}