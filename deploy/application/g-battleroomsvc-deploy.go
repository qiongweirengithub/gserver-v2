package application

import (
	"fmt"
)


func DeployingGBattleRoomsvc(wd string, service string, roomid string)  (string, string, error) {

	// check param
	// 1. port 是否被占用
	fmt.Println(MagentaBg, service, "preparing docker env", Reset)

	_, err := ExecCICmd(wd, "cp", "./bin/conf/battle-release/server.json", "./")
	if err != nil {
		fmt.Println(Red, "copying server.json error for battle room svc ", err, Reset)
		return "", "", err
	}

	_, err = ExecCICmd(wd, "sed", "-i","s/ROOMID/" + roomid + "/g", "./server.json")
	if err != nil {
		fmt.Println(Red, "replace roomid for battle room svc errr ", err, Reset)
		return "","", err
	}


	_, err = ExecCICmd(wd, "cp", "./bin/conf/battle-release/battle.Dockerfile", wd)
	if err != nil {
		fmt.Println(Red, "copying battle.Dockerfile error for battle room svc ", err, Reset)
		return "","", err
	}


	fmt.Println(MagentaBg, service, "image is builing", Reset)
	_, err = ExecCDCmd(wd, "docker", "build", "-t", service, "-f", wd + "/battle.Dockerfile", ".")
	if err != nil {
		fmt.Println(Red, "building image error for battle room svc ", err, Reset)
		return "","", err
	}

	fmt.Println(MagentaBg, service, "starting in docker", Reset)
	appId := "-" + RandString(6)
	containerName := service + appId
	containerPidFileName := service+appId+".pid"
	_, err = ExecCDCmd(wd, "docker","run","-d","--name", containerName, "--cidfile" , containerPidFileName, service)
	if err != nil {
		fmt.Println(Red, "running docker error for battle room svc ", err, Reset)
		return "","", err
	}

	fmt.Println(MagentaBg, service, "started success in docker", Reset)

	pidFile := wd + "/" + containerPidFileName
	pid := GetPid(pidFile)
	fmt.Println(Yellow, "--->   roomid: ", roomid, Reset)
	fmt.Println(Yellow, "--->   pid:    ", pid,  Reset)
	fmt.Println(Yellow, "--->   pidfile:", pidFile,  Reset)

	return pid, containerName, nil

}