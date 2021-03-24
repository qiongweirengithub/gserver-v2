/**
一定要记得在confin.json配置这个模块的参数,否则无法使用
*/
package gate

import (
	"fmt"
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/module/base"
)

var GateService = func() module.Module {
	this := new(Service)
	return this
}

type Service struct {
	basemodule.BaseModule
}

func (self *Service) GetType() string {
	//很关键,需要与配置文件中的Module配置对应
	return "gate-service"
}
func (self *Service) Version() string {
	//可以在监控时了解代码版本
	return "1.0.0"
}
func (self *Service) OnAppConfigurationLoaded(app module.App) {
	//当App初始化时调用，这个接口不管这个模块是否在这个进程运行都会调用
	self.BaseModule.OnAppConfigurationLoaded(app)
}
func (self *Service) OnInit(app module.App, settings *conf.ModuleSettings) {
	self.BaseModule.OnInit(self, app, settings)
	self.GetServer().Options().Metadata["state"] = "alive"
	self.GetServer().RegisterGO("/gateinfo/load", self.loadGate)
	log.Info("%v模块初始化完成...", self.GetType())
}

func (self *Service) Run(closeSig chan bool) {
	log.Info("%v模块运行中...", self.GetType())
	<-closeSig
	log.Info("%v模块已停止...", self.GetType())
}

func (self *Service) OnDestroy() {
	//一定继承
	self.BaseModule.OnDestroy()
	log.Info("%v模块已回收...", self.GetType())
}
func (self *Service) loadGate(token string) (r string, err error) {
	return fmt.Sprintf("hi %v load gate", token), nil
}

