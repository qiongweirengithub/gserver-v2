package admin

import (

	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	basemodule "github.com/liangdas/mqant/module/base"

	"io"

	"net/http"
)

var RestApi = func() module.Module {
	this := new(GameApi)
	return this
}

type GameApi struct {
	basemodule.BaseModule
	Port string
}

func (self *GameApi) GetType() string {
	//很关键,需要与配置文件中的Module配置对应
	return "admin-restapi"
}
func (self *GameApi) Version() string {
	//可以在监控时了解代码版本
	return "1.0.0"
}
func (self *GameApi) OnInit(app module.App, settings *conf.ModuleSettings) {
	self.BaseModule.OnInit(self, app, settings)
	self.Port = ":" + (string)(self.GetModuleSettings().Settings["Port"].(string))
}

func (self *GameApi)startHttpServer(port string) *http.Server {

	srv := &http.Server{Addr: port}

	// register api
	http.HandleFunc("/service/list", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		rstr := r.Form.Get("account")
		log.Info("admin request param %v,", rstr, )
		_, _ = io.WriteString(w, rstr)
	})

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			log.Info("Httpserver: ListenAndServe() error: %s", err)
		}
	}()

	// 获取机器列表

	// 获取机器端口列表

	// 获取service列表

	// 发布service到服务器m:port

	return srv
}

func (self *GameApi) Run(closeSig chan bool) {

	log.Info("web: starting HTTP server : %s", self.Port)

	srv := self.startHttpServer(self.Port)
	<-closeSig
	log.Info("web: stopping HTTP server")
	// now close the server gracefully ("shutdown")
	// timeout could be given instead of nil as a https://golang.org/pkg/context/
	if err := srv.Shutdown(nil); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
	log.Info("web: done. exiting")
}

func (self *GameApi) OnDestroy() {
	//一定别忘了继承
	self.BaseModule.OnDestroy()
}
