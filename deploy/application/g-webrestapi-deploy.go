package application

import (
	"fmt"
)


func DeployingGWebRestApi(wd string, service string, port string)  (string, string, error) {

	// check param
	// 1. port 是否被占用

	fmt.Println(MagentaBg, service, "preparing docker env", Reset)

	_, err := ExecCICmd(wd, "cp", "./bin/conf/web-release/server.json", "./")
	if err != nil {
		fmt.Println(Red, "copying server.json error for web rest api svc ", err, Reset)
		return "", "", err
	}

	_, err = ExecCICmd(wd, "sed", "-i","s/PORT/" + port + "/g", "./server.json")
	if err != nil {
		fmt.Println(Red, "replace roomid for web rest api errr ", err, Reset)
		return "","", err
	}


	_, err = ExecCICmd(wd, "cp", "./bin/conf/web-release/web.Dockerfile", wd)
	if err != nil {
		fmt.Println(Red, "copying battle.Dockerfile error for web rest api ", err, Reset)
		return "","", err
	}

	_, err = ExecCDCmd(wd, "docker", "build", "-t", service, "-f", wd + "/web.Dockerfile", ".")
	if err != nil {
		fmt.Println(Red, "building image error for web rest api ", err, Reset)
		return "","", err
	}

	fmt.Println(MagentaBg, service, "starting in docker", Reset)

	appId := "-" + RandString(6)
	containerName := service + appId
	containerPidFileName := service+appId+".pid"

	_, err = ExecCDCmd(wd, "docker","run","-d","--name", containerName, "--cidfile" , containerPidFileName,"-p"+port+":"+port, service)
	if err != nil {
		fmt.Println(Red, "running docker error for web rest api ", err, Reset)
		return "","", err
	}

	fmt.Println(MagentaBg, service, "started success in docker", Reset)

	pidFile := wd + "/" + containerPidFileName
	pid := GetPid(pidFile)
	fmt.Println(Yellow, "--->   pid:    ", pid,  Reset)
	fmt.Println(Yellow, "--->   pidfile:", pidFile,  Reset)

	// record to db
	return pid, containerName, nil

}