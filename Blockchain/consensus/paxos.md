# Paxos
## 背景
在常见的分布式系统中，总会发生诸如机器宕机或网络异常（包括消息的延迟、丢失、重复、乱序，还有网络分区）等情况。Paxos算法需要解决的问题就是如何在一个可能发生上述异常的分布式系统中，快速且正确地在集群内部对某个数据的值达成一致，并且保证不论发生以上任何异常，都不会破坏整个系统的一致性。<sup>[1]</sup>

注：这里某个数据的值并不只是狭义上的某个数，它可以是一条日志，也可以是一条命令（command）。。。根据应用场景不同，某个数据的值有不同的含义。

## 角色
一个节点可以扮演多个角色
* Client: 系统外部角色，请求发起者，不直接参与一致性算法
* Proposer: 接收Client请求，替民众提出议案(propose)
* Acceptor(Voter): 提议投票和接收者，只有在形成一定人数时(Quorum多数派)，提议才会被接受，类似于国会
* Learner: 提议接收者，对集群一致性无影响，作备份的记录员

### Basic Paxos
#### 阶段
##### Phase 1a: Prepare
proposer提出一个议案N，此N大于此proposer之前提出的提案编号。请求acceptors的Quorum接受
##### Phase 1b: Promise
如果N大于此acceptor之前接受的任何提案编号则接受，否则拒绝
##### Phase 2a: Accept
如果达到了多数派(一半以上)，proposer会发出accept请求，此请求包含提案编号N及其内容
##### Phase 2b: Accepted
如果此acceptor在此期间没有收到任何编号大于N的提案，则接受此提案内容，否则忽略

![basic paxos](/Blockchain/img/basic-paxos.png)

#### 存在的问题
##### 活锁(liveness)或dueling问题
两个Proposer不断交替提出更大的N时，系统无法正常运行

![basic paxos weakness](/Blockchain/img/basic-paxos-weakness.png)

解决方案：失败时用一个随机的timeout
##### 难实现、效率低（2轮RPC Remote Procedure Call）

### Multi Paxos
#### 角色
* Client
* Server: 一段固定时间内，其中某个Server为Proposer

竞选Proposer需要一轮RPC，确定Proposer后只需要一轮RPC

![basic paxos weakness](/Blockchain/img/multi-paxos.png)

简化后：

![basic paxos weakness](/Blockchain/img/multi-paxos-smiplify.png)


## 引用
* [1] [分布式系列文章——Paxos算法原理与推导](https://www.cnblogs.com/linbingdong/p/6253479.html)
* [2] [一致性算法（Paxos、Raft、ZAB）](https://www.bilibili.com/video/BV1TW411M7Fx?from=search&seid=5528910020621807956&spm_id_from=333.337.0.0)