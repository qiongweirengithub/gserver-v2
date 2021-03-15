package main

import (
	"fmt"
	"net/http"

	"github.com/liangdas/mqant"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/registry"
	"github.com/liangdas/mqant/registry/consul"
	"github.com/nats-io/nats.go"

	"gserver.v2/web"
	"gserver.v2/gate"
	"gserver.v2/battle"
	"gserver.v2/auth"
	"gserver.v2/env"
)


func main() {

	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()

	var consul_addr string = env.GetRsAddr()
	var nats_addr string = env.GetNcAddr()

	fmt.Println("staring")

	// rs := consul.NewRegistry(func(options *registry.Options) {
	// 	options.Addrs = []string{"172.17.0.1:8500"}
	// })
	// nc, err := nats.Connect("nats://172.17.0.1:4222", nats.MaxReconnects(10000))
	
	rs := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{consul_addr}
	})



	nc, err := nats.Connect(nats_addr, nats.MaxReconnects(10000))
	if err != nil {
		log.Error(err.Error())
		return
	}


	app := mqant.CreateApp(
		module.Debug(true),  //只有是在调试模式下才会在控制台打印日志, 非调试模式下只在日志文件中输出日志
		module.Parse(true),
		module.Nats(nc),     //指定nats rpc, 此处设置为空， 在配置加载完成后进行设置
		module.Registry(rs), //指定服务发现， 此处设置为空， 在配置加载完成后进行设置
	)

	_ = app.OnConfigurationLoaded(func(app module.App) {

	 	consul_addr = app.GetSettings().Settings["consul_addr"].(string)
    	nats_addr = app.GetSettings().Settings["nats_addr"].(string)
		fmt.Println("consul: " + consul_addr + ", nats: " + nats_addr)

	})

	err = app.Run(

		// 启动长链接业务
		gate.ConnectionSvc(),

		// 网管信息查询
		gate.GateService(),

		// 授权相关服务
		auth.AService(),

		// 对局服务(对战列表，对战详情)
		battle.BService(),

		// 对上述服务进行 rest 汇总
		web.RestApi(),
		
	)
	if err != nil {
		log.Error(err.Error())
	}

}
