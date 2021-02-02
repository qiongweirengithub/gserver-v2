package gate

import (
    "github.com/liangdas/mqant/conf"
    "github.com/liangdas/mqant/gate"
    "github.com/liangdas/mqant/gate/base"
	"github.com/liangdas/mqant/module"
    "github.com/liangdas/mqant/log"
)

var Module = func() module.Module {
    gate := new(Gate)
    return gate
}

type Gate struct {
    basegate.Gate
    ws_addr string
    tcp_addr string

}

func (this *Gate) GetType() string {
    //很关键,需要与配置文件中的Module配置对应
    return "Gate"
}
func (this *Gate) Version() string {
    //可以在监控时了解代码版本
    return "1.0.0"
}

func (this *Gate) OnInit(app module.App, settings *conf.ModuleSettings) {
    //注意这里一定要用 gate.Gate 而不是 module.BaseModule

    log.Info("gate: starting %s ", settings)

    this.ws_addr = ":" + (string) (settings.Settings["wsaddr"].(string))
    this.tcp_addr = ":" + (string) (settings.Settings["tcpaddr"].(string))

    log.Info("gate: starting ws port:%s, tcp port:%s", this.ws_addr, this.tcp_addr)

    this.Gate.OnInit(this, app, settings,
        gate.WsAddr(this.ws_addr),
        gate.TcpAddr(this.tcp_addr),
	)
	log.Info("%v模块运行中...", "gate")
	
}
