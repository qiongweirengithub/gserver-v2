package web

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/module/base"
	"github.com/liangdas/mqant/rpc"
	"github.com/liangdas/mqant/rpc/pb"

	"io"

	"gserver.v2/rpctest"

	"net/http"
	"time"
)

var Module = func() module.Module {
    this := new(Web)
    return this
}

type Web struct {
    basemodule.BaseModule
    Port string
}

func (self *Web) GetType() string {
    //很关键,需要与配置文件中的Module配置对应
    return "Web"
}
func (self *Web) Version() string {
    //可以在监控时了解代码版本
    return "1.0.0"
}
func (self *Web) OnInit(app module.App, settings *conf.ModuleSettings) {
    self.BaseModule.OnInit(self, app, settings)
    self.Port = ":" + (string) (self.GetModuleSettings().Settings["Port"].(string))
}

func (self *Web)startHttpServer(port string) *http.Server {

	srv := &http.Server{Addr: port}
	rpctest.Module()
	
	// hello world api
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        _ = r.ParseForm()
        rstr,err:=mqrpc.String(
            self.Call(
            context.Background(),
            "helloworld",
            "/say/hi",
            mqrpc.Param(r.Form.Get("name")),
            ),
        )
        log.Info("RpcCall %v , err %v",rstr,err)
        _, _ = io.WriteString(w, rstr)
	})
	
	http.HandleFunc("/test/proto", func(w http.ResponseWriter, r *http.Request) {
        _ = r.ParseForm()
        ctx, _ := context.WithTimeout(context.TODO(), time.Second*3)
        protobean := new(rpcpb.ResultInfo)
        err:=mqrpc.Proto(protobean,func() (reply interface{}, errstr interface{}) {
            return self.RpcCall(
                ctx,
                "rpctest",     //要访问的moduleType
                "/test/proto", //访问模块中handler路径
                mqrpc.Param(&rpcpb.ResultInfo{Error: *proto.String(r.Form.Get("message"))}),
            )
        })
        log.Info("RpcCall %v , err %v",protobean,err)
        if err!=nil{
            _, _ = io.WriteString(w, err.Error())
        }
        _, _ = io.WriteString(w, protobean.Error)
    })


    go func() {
        if err := srv.ListenAndServe(); err != nil {
            // cannot panic, because this probably is an intentional close
            log.Info("Httpserver: ListenAndServe() error: %s", err)
        }
    }()
    // returning reference so caller can call Shutdown()
    return srv
}

func (self *Web) Run(closeSig chan bool) {

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

func (self *Web) OnDestroy() {
    //一定别忘了继承
    self.BaseModule.OnDestroy()
}
