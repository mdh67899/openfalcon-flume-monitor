open-falcon flume监控脚本
================================

系统需求
--------------------------------
操作系统：Linux

Python >= 2.6

python-requests


主要逻辑
--------------------------------
通过flume进程运行时暴露的http接口获取flume组件信息，通过python解析组件信息，将得到的json数据格式的结果输出到标准输出

汇报字段
--------------------------------

flume metrics页面内的数据分为三部分:

1. SOURCE：flume的数据源组件，所有收集日志的第一个到达的地方

2. CHANNEL：flume的通道组件，对数据具有缓存的作用

3. SINK：数据离开flume前的最后一个组件，负责从channel中取走数据，然后发送到缓存系统或者持久化数据库

除去StartTime、StopTime、Type以及ChannelCapacity(通道容量)之外，所有数据均汇报，上报的数据metric会添加相应的name前缀，并且在tag中上传flume进程的监听端口号：

| metric |  tag | type | note |
|-----|------|------|------|
|OpenConnectionCount|port=3000|GAUGE|打开的连接数|
|AppendBatchReceivedCount|port=3000|COUNTER|source端刚刚追加的批数量|
|AppendBatchAcceptedCount|port=3000|COUNTER|追加到channel中的批数量|
|AppendReceivedCount|port=3000|COUNTER|source追加目前收到的数量|
|AppendAcceptedCount|port=3000|COUNTER|放入channel的event数量|
|EventReceivedCount|port=3000|COUNTER|source端成功收到的event数量|
|EventAcceptedCount|port=3000|COUNTER|成功放入channel的event数量|
|ChannelFillPercentage|port=3000|GAUGE|通道使用比例|
|ChannelSize|port=3000|GAUGE|目前在channel中的event数量|
|EventPutSuccessCount|port=3000|COUNTER|成功放入channel的event数量|
|EventPutAttemptCount|port=3000|COUNTER|尝试放入将event放入channel的次数|
|EventTakeSuccessCount|port=3000|COUNTER|从channel中成功取走的event数量|
|EventTakeAttemptCount|port=3000|COUNTER|尝试从channel中取走event的次数|
|BatchCompleteCount|port=3000|COUNTER|完成的批数量|
|ConnectionFailedCount|port=3000|COUNTER|连接失败数|
|EventDrainAttemptCount|port=3000|COUNTER|尝试提交的event数量|
|ConnectionCreatedCount|port=3000|COUNTER|创建连接数|
|BatchEmptyCount|port=3000|COUNTER|批量取空的数量|
|ConnectionClosedCount|port=3000|COUNTER|关闭连接数量|
|EventDrainSuccessCount|port=3000|COUNTER|成功发送event的数量|
|BatchUnderflowCount|port=3000|COUNTER|正处于批量处理的batch数|


监控告警项设置建议
--------------------------------

说明: 请根据实际情部署情况以及使用方式，自行调整监控项触发条件，以下报警条件只是基础监控，详细报警条件请自行调整

| 监控项 | 告警触发条件 | 备注 |
|-----|-----|-----|
|net.port.listen/port=3000|all(3)==0|最近三分钟flume进程监听的metrics http端口挂掉了，检查flume进程是否有问题|
|ChannelFillPercentage|all(3)>90|最近三分钟flume channel空间使用率超过90%，检查sink的发送速率是否有问题|
|EventReceivedCount|all(3)==0|最近三分钟flume source组件最近三分钟收到event为0，检查flume进程是否有问题|
|EventAcceptedCount|all(3)==0|最近三分钟flume成功放入channel组件的event数量为0，检查flume进程是否有问题|

参考资料
--------------------------------
flume官网：http://flume.apache.org/FlumeUserGuide.html#json-reporting
