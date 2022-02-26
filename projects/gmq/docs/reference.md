# Reference

- https://github.com/nsqio/go-nsq
- https://github.com/nsqio/nsq

- https://code.aliyun.com/middlewarerace2017/open-messaging-demo

阅读Open-Messaging的接口代码(本工程内除了demo目录外的其他代码)，了解Topic，Queue的基本概念，并实现一个进程内消息引擎。

提示：Topic类似于水坝（蓄积功能，消峰填谷之利器），Queue类似于水渠；每当新建一个Queue的时候，可以选择绑定到几个Topic，类似于水渠从水坝引水； 
每个Topic可以被任意多个Queue绑定，这点与现实生活不太一样，因为数据可以多次拷贝； 
在发送的时候，可以选择发送到Topic，也可以选择直接发送到Queue；直接发送到Queue的数据只能被对应Queue消费，不能被其他Queue读取到； 
在消费的时候，除了要读取绑定的Topic的数据，还要去取直接发送到该Queue的数据；

## OpenMessaging 规范
