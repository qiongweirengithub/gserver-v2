/**
一定要记得在confin.json配置这个模块的参数,否则无法使用
*/
package auth

import (
	"fmt"
	"time"
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/module/base"
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/server"
)

var AService = func() module.Module {
	this := new(Auth)
	return this
}

type Auth struct {
	basemodule.BaseModule
}

func (self *Auth) GetType() string {
	//很关键,需要与配置文件中的Module配置对应
	return "auth-service"
}
func (self *Auth) Version() string {
	//可以在监控时了解代码版本
	return "1.0.0"
}
func (self *Auth) OnAppConfigurationLoaded(app module.App) {
	//当App初始化时调用，这个接口不管这个模块是否在这个进程运行都会调用
	self.BaseModule.OnAppConfigurationLoaded(app)
}
func (self *Auth) OnInit(app module.App, settings *conf.ModuleSettings) {
	

	var nodeId string = settings.ProcessID

	self.BaseModule.OnInit(self, app, settings,
		server.RegisterInterval(15*time.Second),
		server.RegisterTTL(30*time.Second),
		// 注册到 consul 的服务id， 这个是用来做服务路由的，因此一定要
		server.Id(nodeId),
	)
	
	self.GetServer().Options().Metadata["state"] = "alive"
	self.GetServer().RegisterGO("/register", self.register)
	self.GetServer().RegisterGO("/login", self.login) 
	self.GetServer().RegisterGO("/logout", self.logout)
	self.GetServer().RegisterGO("HD_hello", self.gatesay) 

	log.Info("%v模块初始化完成...", self.GetType())
}

func (self *Auth) Run(closeSig chan bool) {
	log.Info("%v模块运行中...", self.GetType())
	<-closeSig
	log.Info("%v模块已停止...", self.GetType())
}

func (self *Auth) OnDestroy() {
	//一定继承
	self.BaseModule.OnDestroy()
	log.Info("%v模块已回收...", self.GetType())
}
func (self *Auth) register(name string, pwd string) (r string, err error) {
	return fmt.Sprintf("hi %v register", name), nil
}

func (self *Auth) login(name string, pwd string) (r string, err error) {
	return fmt.Sprintf("hi %v login", name), nil
}

func (self *Auth) logout(token string) (r string, err error) {
	return fmt.Sprintf("hi %v logout", token), nil
}


func (self *Auth) gatesay(session gate.Session, msg map[string]interface{}) (r string, err error) {
	session.Send("/gate/send/test", []byte(fmt.Sprintf("send hi to %v", msg["name"])))
	return fmt.Sprintf("hi %v 你在网关 %v, session:%v", msg["name"], session.GetServerId(), session.GetSessionId()), nil
}


func (self *Auth) gatesay2(session gate.Session, msg map[string]interface{}) (r string, err error) {
	session.Send("/gate/send/test", []byte(fmt.Sprintf("send hi to %v", msg["name"])))
	return fmt.Sprintf("hi %v 你在网关 %v, session:%v", msg["name"], session.GetServerId(), session.GetSessionId()), nil
}



