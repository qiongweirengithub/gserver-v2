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
package battleroomsvctest

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/liangdas/armyant/task"
	"github.com/liangdas/armyant/work"
)

func NewWork(manager *Manager) *Work {
	this := new(Work)
	this.manager = manager
	//opts := this.GetDefaultOptions("tcp://127.0.0.1:3563")
	opts := this.GetDefaultOptions("ws://127.0.0.1:3653")
	opts.SetConnectionLostHandler(func(client MQTT.Client, err error) {
		fmt.Println("ConnectionLost", err.Error())
	})
	opts.SetOnConnectHandler(func(client MQTT.Client) {
		fmt.Println("OnConnectHandler")
	})
	// load root ca
	// 需要一个证书，这里使用的这个网站提供的证书https://curl.haxx.se/docs/caextract.html
	err := this.Connect(opts)
	if err != nil {
		fmt.Println(err.Error())
	}

	this.On("/room_table/new_event", func(client MQTT.Client, msg MQTT.Message) {
		//服务端主动下发玩家加入事件
		fmt.Println(msg.Topic(), string(msg.Payload()))
	})


	this.On("room_table/event_conform", func(client MQTT.Client, msg MQTT.Message) {
		//服务端主动下发玩家加入事件
		fmt.Println(msg.Topic(), string(msg.Payload()))
	})

	return this
}

/**
Work 代表一个协程内具体执行任务工作者
*/
type Work struct {
	work.MqttWork
	manager *Manager
}

func (this *Work) UnmarshalResult(payload []byte) map[string]interface{} {
	rmsg := map[string]interface{}{}
	json.Unmarshal(payload, &rmsg)
	return rmsg["Result"].(map[string]interface{})
}


func (this *Work) RunWorker(t task.Task) {
	// msg, err := this.Request("helloworld@helloworld001/HD_say", []byte(`i want test module id`))
	fmt.Println("sending request")

	msg, err := this.Request("battleroomsvc@12352/HD_create_table", []byte(`{"table_id":"testtable01"}`))

	msg, err = this.Request("battleroomsvc@12352/HD_join_table", []byte(`{"table_id":"testtable01"}`))

	// msg, err := this.Request("battleroomsvc/table_create", []byte(`{"tableId":"testtable01"}`))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(msg.Topic(), string(msg.Payload()))
}
func (this *Work) Init(t task.Task) {

}
func (this *Work) Close(t task.Task) {
	this.GetClient().Disconnect(0)
}
