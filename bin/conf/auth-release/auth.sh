#go build -o improving-server improving-server.go
/home/qiongwei/mycode/goprj/gserver.v2/gserverv2 -pid auth001 -conf /home/qiongwei/mycode/goprj/gserver.v2/bin/conf/auth-release/server-auth001.json -wd /home/qiongwei/mycode/goprj/gserver.v2 &
/home/qiongwei/mycode/goprj/gserver.v2/gserverv2 -pid auth002 -conf /home/qiongwei/mycode/goprj/gserver.v2/bin/conf/auth-release/server-auth002.json -wd /home/qiongwei/mycode/goprj/gserver.v2 &
/home/qiongwei/mycode/goprj/gserver.v2/gserverv2 -pid auth003 -conf /home/qiongwei/mycode/goprj/gserver.v2/bin/conf/auth-release/server-auth003.json -wd /home/qiongwei/mycode/goprj/gserver.v2 &

# pgrep -lf "auth0" | awk '{print $1}'| xargs kill -15