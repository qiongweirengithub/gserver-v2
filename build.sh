#go build -o improving-server improving-server.go



 export GOPATH=$GOPATH:$PWD

 echo $GOPATH

 echo "go build mqant-example"

#  go build -o gserverv2 main.go

#  静态编译
CGO_ENABLED=0 go build -a -o gserverv2 -ldflags '-extldflags "-static"' .

 echo "编译完成"

 cp gserverv2 ~/myapp/k8s/app/
 cp ./bin/conf/env.json ~/myapp/k8s/app/
 cp ./bin/conf/server.json ~/myapp/k8s/app/
