#go build -o improving-server improving-server.go
/home/qiongwei/mycode/goprj/gserver.v2/gserverv2 -pid gateapi001 -conf /home/qiongwei/mycode/goprj/gserver.v2/bin/conf/gateapi-release/server-gateapi001.json -wd /home/qiongwei/mycode/goprj/gserver.v2 &
/home/qiongwei/mycode/goprj/gserver.v2/gserverv2 -pid gateapi002 -conf /home/qiongwei/mycode/goprj/gserver.v2/bin/conf/gateapi-release/server-gateapi002.json -wd /home/qiongwei/mycode/goprj/gserver.v2 &
/home/qiongwei/mycode/goprj/gserver.v2/gserverv2 -pid gateapi003 -conf /home/qiongwei/mycode/goprj/gserver.v2/bin/conf/gateapi-release/server-gateapi003.json -wd /home/qiongwei/mycode/goprj/gserver.v2 &

# pgrep -lf "gateapi0" | awk '{print $1}'| xargs kill -15