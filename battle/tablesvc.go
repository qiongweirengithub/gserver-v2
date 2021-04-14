package battle

import (
	"errors"
	"fmt"
	"github.com/liangdas/mqant-modules/room"
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"reflect"
	"time"
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
*   处理每帧消息
*/
func (this *GTable) Receive(msg *room.QueueMsg, index int) {
	log.Info("帧同步消息:", msg)
	action := msg.Func
	playermsg := msg.Params[2].(string)
	if "join" == action {
		log.Info("new player join:", playermsg)
		session := msg.Params[1].(gate.Session)	
		player := &room.BasePlayerImp{}
		player.Bind(session)
		player.OnRequest(session)
		this.players[session.GetSessionId()] = player
	} else {
		log.Info("unsupport action:", action)
	}

	for _,pl := range this.GetSeats() {
		log.Info("syn to all player join:", playermsg)
		pl.Session().Send("/gate/send/test", []byte(fmt.Sprintf("send hi to %v", action)))		
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
