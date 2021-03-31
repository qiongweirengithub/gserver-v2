package application

import (
	"fmt"
)


func DeployingGAuthServuce(wd string, service string)  (string, string, error) {

	// check param
	// 1. port 是否被占用
	fmt.Println("deploying auth svc")

	_, err := ExecCICmd(wd, "cp", "./bin/conf/auth-release/server.json", "./")
	if err != nil {
		fmt.Print(err)
		return "", "", err
	}

	_, err = ExecCICmd(wd, "cp", "./bin/conf/auth-release/auth.Dockerfile", wd)
	if err != nil {
		fmt.Print(err)
		return "","", err
	}

	_, err = ExecCDCmd(wd, "docker", "build", "-t", service, "-f", wd + "/auth.Dockerfile", ".")
	if err != nil {
		fmt.Print(err)
		return "","", err
	}
	fmt.Println("auth svc docker build succ")

	appId := "-" + RandString(6)
	containerName := service + appId
	containerPidFileName := service+appId+".pid"

	_, err = ExecCDCmd(wd, "docker","run","-d","--name", containerName, "--cidfile" , containerPidFileName, service)

	if err != nil {
		fmt.Print(err)
		return "","", err
	}

	fmt.Println("auth svc running ")

	pdiFile := wd + "/" + containerPidFileName
	pid := GetPid(pdiFile)
	fmt.Println("pid file:", pdiFile, ", pid: ", pid)

	// record to db
	return pid, containerName, nil

}