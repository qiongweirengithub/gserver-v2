/**
*  一定要记得在confin.json配置这个模块的参数,否则无法使用
*  唯一作用是 处理游戏内 逻辑， 不提供其他服务
*  以单一服务的形式运行在机器上，负责管理一批 table 的生命周期
*  单一机器可以运行多个此服务实例
 */
package battle

import (
	"errors"
	"fmt"
	"time"

	"github.com/liangdas/mqant-modules/room"
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/module/base"
	"github.com/liangdas/mqant/server"

)

var Room = func() module.Module {
	this := new(RoomSvc)
	return this
}

type RoomSvc struct {
	basemodule.BaseModule
	room *room.Room
}

func (self *RoomSvc) GetType() string {
	//很关键,需要与配置文件中的Module配置对应
	return "game-roomsvc"
}
func (self *RoomSvc) Version() string {
	//可以在监控时了解代码版本
	return "1.0.0"
}
func (self *RoomSvc) OnAppConfigurationLoaded(app module.App) {
	//当App初始化时调用，这个接口不管这个模块是否在这个进程运行都会调用
	self.BaseModule.OnAppConfigurationLoaded(app)
}
func (self *RoomSvc) OnInit(app module.App, settings *conf.ModuleSettings) {
	self.BaseModule.OnInit(self, app, settings)
	self.GetServer().Options().Metadata["state"] = "alive"

	// 启动后自动把自己的信息同步到数据库

	// 固定id,可定向访问
	var nodeId string = settings.ProcessID
	self.BaseModule.OnInit(self, app, settings,
		server.RegisterInterval(15*time.Second),
		server.RegisterTTL(30*time.Second),
		// 注册到 consul 的服务id， 这个是用来做服务路由的，因此一定要
		server.Id(nodeId),
	)


	self.GetServer().RegisterGO("/table/create", self.createTable)
	log.Info("%v模块初始化完成...", self.GetType())
}

func (self *RoomSvc) Run(closeSig chan bool) {
	log.Info("%v模块运行中...", self.GetType())
	<-closeSig
	log.Info("%v模块已停止...", self.GetType())
}

func (self *RoomSvc) OnDestroy() {
	//一定继承
	self.BaseModule.OnDestroy()
	log.Info("%v模块已回收...", self.GetType())
}

func (self *RoomSvc) createTable(module module.RPCModule, tableId string) (room.BaseTable, error) {
	table := NewTable(
		module,
		room.TableId(tableId),
		room.Router(func(TableId string) string {
			return fmt.Sprintf("%v://%v/%v", self.GetType(), self.GetServerId(), tableId)
		}),
		room.DestroyCallbacks(func(table room.BaseTable) error {
			log.Info("回收了房间: %v", table.TableId())
			_ = self.room.DestroyTable(table.TableId())
			return nil
		}),
	)

	// 更新此room的状态到数据库 

	return table, nil
}

func (self *RoomSvc) gatesay(session gate.Session, msg map[string]interface{}) (r string, err error) {
	table_id := msg["table_id"].(string)
	action := msg["action"].(string)
	table := self.room.GetTable(table_id)
	if table == nil {
		return "操作失败", errors.New("房间不存在")
	}
	erro := table.PutQueue(action, session, msg)
	if erro != nil {
		return "", erro
	}
	return "success", nil
}

