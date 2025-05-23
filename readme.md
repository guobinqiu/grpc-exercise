RPC
---
rpc到底是个啥？rpc是remote procedure call的缩写，中文名叫远程过程调用。procedure称为过程，过程和函数的本质区别是函数是有返回值的，过程是没有返回值的。rpc可以简单理解为跨进程的过程调用，同一主机上两个不同进程间的调用可以叫rpc，不同主机间跨网络的调用也可以叫rpc。

我们把rpc中的两个角色分别称作producer和consumer。producer作为服务的提供者既要有过程声明（也叫存根），又要有过程实现；consumer作为服务的消费者只需要过程声明，用来告诉producer我想要调用你哪一个过程。

producer和consumer两个进程间的通信需要建立在一个点对点的tcp连接之上，出于性能考虑这个连接最好是一个长连接。consumer需要把想要调用的过程名及过程参数通过编码转换成字节数组发送给producer。producer收到数据后先进行解码，读出其中的过程名和过程参数，然后在反射中代入该过程名及过程参数，通过反射调用自身的过程实现得到执行结果，然后把该结果编码后发回给consumer。consumer收到数据后再对其进行解码，这就完成了一次rpc调用。

在一个应用里，可以同时拥有rpc的producer和consumer。

这中间需要构建的模块至少应有基于socket的网络通信模块，数据的编解码器，服务提供者的反射处理器，以及双方对于过程声明的约定。

使用rpc需要注意的是由于消费端的代码是同步执行的，需要等待服务端返回才能继续往下走，所以服务端的性能直接制约着客户端性能，所以对于耗时长的操作请使用异步的，基于消息队列的架构来处理。

gRPC
---
A high-performance, open-source universal RPC framework

![gRPC](https://grpc.io/img/landing-2.svg)

### Examples

- [go as serverside](go-as-serverside/readme.md)

- [gRPC的4种模式](https://github.com/guobinqiu/grpc-f4)

### Tutorial

> https://grpc.io/docs/tutorials/

