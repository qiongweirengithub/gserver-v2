#go build -o improving-server improving-server.go
/home/qiongwei/mycode/goprj/gserver.v2/gserverv2 -pid helloworld001 -conf /home/qiongwei/mycode/goprj/gserver.v2/bin/conf/helloworld-release/server-helloworld1.json -wd /home/qiongwei/mycode/goprj/gserver.v2 &
/home/qiongwei/mycode/goprj/gserver.v2/gserverv2 -pid helloworld002 -conf /home/qiongwei/mycode/goprj/gserver.v2/bin/conf/helloworld-release/server-helloworld2.json -wd /home/qiongwei/mycode/goprj/gserver.v2 &
/home/qiongwei/mycode/goprj/gserver.v2/gserverv2 -pid helloworld003 -conf /home/qiongwei/mycode/goprj/gserver.v2/bin/conf/helloworld-release/server-helloworld3.json -wd /home/qiongwei/mycode/goprj/gserver.v2 &


# pgrep -lf "helloworld0" | awk '{print $1}'| xargs kill -15