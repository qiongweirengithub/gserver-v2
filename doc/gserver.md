### 环境

1. docker
    1. 安装
        https://blog.csdn.net/huyaowei789/article/details/106329204
    2. 加速
        https://www.jianshu.com/p/293b4737c938?utm_campaign=maleskine&utm_content=note&utm_medium=seo_notes&utm_source=recommendation

2. go
    1. 安装
        https://www.jianshu.com/p/2aa806b0e4b6
        
    2. 加速
        https://blog.csdn.net/sinat_28371057/article/details/113716714
        https://blog.csdn.net/qq_42409788/article/details/104416422
        
3. consul
    1. 安装
        https://learn.hashicorp.com/tutorials/consul/get-started-agent?in=consul/getting-started
    2. 启动
        consul agent -dev -client 0.0.0.0

4. nats
    1. 安装
        https://docs.nats.io/nats-server/installation
    2. 启动
        1. docker 见安装文档
        2. 源码安装，直接 nats-server 即可启动，建议用 docker 方式

5. mysql
    1. 安装
    2. 值：
        aliyun机器

6. github
    1. 加速 
        推荐使用https://github.zhlh6.cn/

7. vscode
    1. 配置golang开发环境
    2. 配置 remote ssh



### 部署方法

1. 确定启动了 consul
2. 确定启动了 nats

3. 项目代码 
    github  git@github.com:qiongweirengithub/gserver-v2.git
    branch  main-v3，0

4. 切换deploy目录
    go install deploy.go    ->   将部署脚本安装到机器上

5. 执行部署指令
    部署指令见  deploy.go 说明文档
    // 删除应用               		deploy -service=kill -containerid=查询数据库或者docker ps -l
    // 部署 g-web-restapi    		deploy -service=g-web-restapi -port=8090
    // 部署 g-gate-connectionsvc    deploy -service=g-gate-connectionsvc -websocketport=3653
    // 部署 g-authservice           deploy -service=g-authservice
    // 部署 g-battleroomsvc         deploy -service=g-battleroomsvc -roomid=12345

### 部署脚本


### 业务架构


### 常见问题
    阿里云问题见  aliyun.md



