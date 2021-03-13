## 分析与定义用户指令文档

|client|server|交互对象|相关信息|
|----|----|----|----|
|发送匹配请求|返回roomid|match包|详见match包
|加入房间|房间信息更新|game包|详见game/roomsvc.go
|执行动作|事件同步|game包|详见game/roomsvc.go

