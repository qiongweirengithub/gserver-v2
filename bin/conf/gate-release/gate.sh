#go build -o improving-server improving-server.go
/home/qiongwei/mycode/goprj/gserver.v2/gserverv2 -pid gate001 -conf /home/qiongwei/mycode/goprj/gserver.v2/bin/conf/gate-release/server-gate001.json -wd /home/qiongwei/mycode/goprj/gserver.v2 &
/home/qiongwei/mycode/goprj/gserver.v2/gserverv2 -pid gate002 -conf /home/qiongwei/mycode/goprj/gserver.v2/bin/conf/gate-release/server-gate002.json -wd /home/qiongwei/mycode/goprj/gserver.v2 &
/home/qiongwei/mycode/goprj/gserver.v2/gserverv2 -pid gate003 -conf /home/qiongwei/mycode/goprj/gserver.v2/bin/conf/gate-release/server-gate003.json -wd /home/qiongwei/mycode/goprj/gserver.v2 &

# pgrep -lf "gate0" | awk '{print $1}'| xargs kill -15