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

	"gserver.v2/helloworld"
	"gserver.v2/rpctest"
	"gserver.v2/web"
	"gserver.v2/gate"
	"gserver.v2/auth"
)

func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()
	time := helloworld.H1{Time: "xx"}
	fmt.Println(time)
	

	rs := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})

	nc, err := nats.Connect("nats://127.0.0.1:4222", nats.MaxReconnects(10000))
	app := mqant.CreateApp(
		module.Debug(true),  //只有是在调试模式下才会在控制台打印日志, 非调试模式下只在日志文件中输出日志
		module.Parse(true),
		module.Nats(nc),     //指定nats rpc
		module.Registry(rs), //指定服务发现
	)

	err = app.Run(
		helloworld.Module(),
		web.Module(),
		rpctest.Module(),
		mgate.Module(),

		// 业务相关
		auth.Module(),
		web.ApiModule(),
	)
	if err != nil {
		log.Error(err.Error())
	}

}
