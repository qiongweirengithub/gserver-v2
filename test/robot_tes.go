// Copyright 2014 hey Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"fmt"
	"github.com/liangdas/armyant/task"
	// "gserver.v2/test/authservicetest"
	"gserver.v2/test/battleroomsvctest"
	"os"
	"os/signal"
)

func main() {

	// task_authservice := task.LoopTask{
	// 	C: 10, //并发数
	// }
	// manager := authservicetest.NewManager(task_authservice) //房间模型的demo

	fmt.Println("开始压测请等待")
	// task_authservice.Run(manager)



	task_battleroomsvctest := task.LoopTask{
		C: 1, //并发数
	}
	manager_authservicetest := battleroomsvctest.NewManager(task_battleroomsvctest) //房间模型的demo

	task_battleroomsvctest.Run(manager_authservicetest)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	// task_authservice.Stop()
	task_battleroomsvctest.Stop()

	os.Exit(1)
}
