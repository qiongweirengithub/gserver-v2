package battle

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/liangdas/mqant-modules/room"
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
)

type GTable struct {
	room.QTable
	module  module.RPCModule
	players map[string]room.BasePlayer
}

func (this *GTable) GetSeats() map[string]room.BasePlayer {
	return this.players
}
func (this *GTable) GetModule() module.RPCModule {
	return this.module
}

func (this *GTable) OnCreate() {
	//可以加载数据
	log.Info("GTable OnCreate")
	//一定要调用QTable.OnCreate()
	this.QTable.OnCreate()
}

/**
每帧都会调用
*/
func (this *GTable) Update(ds time.Duration) {}

/**
 * 处理每帧消息
 * 
 */
func (this *GTable) Receive(msg *room.QueueMsg, index int) {
	log.Info("帧同步消息:", msg)
	action := msg.Func
	var playermsg map[string]interface{} = msg.Params[1].(map[string]interface{})
	session := msg.Params[0].(gate.Session)	
	player := this.GetSeats()[session.GetSessionId()];

	if "join" == action {
		if len(this.GetSeats()) >= 2 {
			return
		}
		if _, ok := this.GetSeats()[session.GetSessionId()]; ok {
			return
		}
		log.Info("new player join:", playermsg)
		player = &room.BasePlayerImp{}
		player.Bind(session)
		this.GetSeats()[session.GetSessionId()] = player
	} else if "exit" == action {
		delete(this.GetSeats(), session.GetSessionId())
	} else {
		log.Info("unsupport action:", action)
	}

	if player == nil {
		log.Error("player is nul", session.GetSessionId())
		return
	}
	
	// 更新session
    player.OnRequest(session)

	// 通知所有人
	data, err := json.Marshal(playermsg)
	if err != nil {
		log.Error("msg serilized fail", playermsg)
		return
	}

	for _,pl := range this.GetSeats() {
		log.Info("syn to all player event:", playermsg)
		pl.Session().Send("/room_table/new_event", []byte(string(data)))		
	}

}


func NewTable(module module.RPCModule, opts ...room.Option) *GTable {
	this := &GTable{
		module:  module,
		players: map[string]room.BasePlayer{},
	}
	opts = append(opts, room.TimeOut(60))
	opts = append(opts, room.Update(this.Update))
	opts = append(opts, room.NoFound(func(msg *room.QueueMsg) (value reflect.Value, e error) {
		//return reflect.ValueOf(this.doSay), nil
		return reflect.Zero(reflect.ValueOf("").Type()), errors.New("no found handler")
	}))
	opts = append(opts, room.SetRecoverHandle(func(msg *room.QueueMsg, err error) {
		log.Error("Recover %v Error: %v", msg.Func, err.Error())
	}))

	frameUpdateReceiver := this
	//设置每帧消息处理器
	this.QueueTable.SetReceive(frameUpdateReceiver)

	opts = append(opts, room.SetErrorHandle(func(msg *room.QueueMsg, err error) {
		log.Error("Error %v Error: %v", msg.Func, err.Error())
	}))
	this.OnInit(this, opts...)
	return this
}

func (this *GTable) doSay(session gate.Session, msg map[string]interface{}) (err error) {
	player := this.FindPlayer(session)
	if player == nil {
		return errors.New("no join")
	}
	player.OnRequest(session)
	_ = this.NotifyCallBackMsg("/room/say", []byte(fmt.Sprintf("say hi from %v", msg["name"])))
	return nil
}



/**
* 用户新的动作
*/
func (this *GTable) doAction(session gate.Session, msg map[string]interface{}) (err error) {
	player := this.FindPlayer(session)
	player.OnRequest(session)
	name := msg["name"]
	action := msg["action"]
	_ = this.NotifyCallBackMsg("/room_table/event_conform", []byte(fmt.Sprintf("welcome to %v:%v", name, action)))
	return nil
}
