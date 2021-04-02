package application

import (
	"fmt"
)


func DeployingGBattleRoomsvc(wd string, service string, roomid string)  (string, string, error) {

	// check param
	// 1. port 是否被占用

	fmt.Println("deploying battle room service")

	_, err := ExecCICmd(wd, "cp", "./bin/conf/battle-release/server.json", "./")
	if err != nil {
		fmt.Print(err)
		return "", "", err
	}

	_, err = ExecCICmd(wd, "sed", "-i","s/ROOMID/" + roomid + "/g", "./server.json")
	if err != nil {
		fmt.Print(err)
		return "","", err
	}


	_, err = ExecCICmd(wd, "cp", "./bin/conf/battle-release/battle.Dockerfile", wd)
	if err != nil {
		fmt.Print(err)
		return "","", err
	}

	_, err = ExecCDCmd(wd, "docker", "build", "-t", service, "-f", wd + "/battle.Dockerfile", ".")
	if err != nil {
		fmt.Print(err)
		return "","", err
	}
	fmt.Println("battle docker build succ")

	appId := "-" + RandString(6)
	containerName := service + appId
	containerPidFileName := service+appId+".pid"

	_, err = ExecCDCmd(wd, "docker","run","-d","--name", containerName, "--cidfile" , containerPidFileName, service)
	if err != nil {
		fmt.Print(err)
		return "","", err
	}

	fmt.Println("new battle room service: ", roomid)

	pdiFile := wd + "/" + containerPidFileName
	pid := GetPid(pdiFile)
	fmt.Println("pid file:", pdiFile, ", pid: ", pid)

	// record to db
	return pid, containerName, nil

}