#go build -o improving-server improving-server.go
/home/qiongwei/mycode/goprj/gserver.v2/gserverv2 -pid web001 -conf /home/qiongwei/mycode/goprj/gserver.v2/bin/conf/web-release/server-web01.json -wd /home/qiongwei/mycode/goprj/gserver.v2 &
/home/qiongwei/mycode/goprj/gserver.v2/gserverv2 -pid web002 -conf /home/qiongwei/mycode/goprj/gserver.v2/bin/conf/web-release/server-web02.json -wd /home/qiongwei/mycode/goprj/gserver.v2 &


# pgrep -lf "web0" | awk '{print $1}'| xargs kill -15