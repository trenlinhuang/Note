# 容错问题和Paxos算法

## 2.1 客户端/服务器
**定理2.7：** 系统有多个服务器、多个客户端时，各个服务器接收到的命令顺序可能不同，这将导致不一致的状态

**证明：** 假定有两个客户端`u1`, `u2`，两个服务器端`s1`,`s2`，两个服务器共同维护一个初始值为0的变量`x`。
`u1`发送命令 x=x+1，`u2`发送命令 x=2*x。

若`s1`先收到`u1`的命令，`s2`先收到`u2`的命令，两个服务器将产生不同的计算结果，导致不一致的状态。

### 状态复制（State Replication）
**定义 2.8：** 对于一组节点，如果所有节点均以相同的顺序执行一个命令序列`c1`, `c2`, `c3`, ...，则这组节点实现了状态复制。
* 状态复制时分布式系统的基本性质
* 对于金融科技的从业人员来说，状态复制经常等同于区块链

### 串行化器（Serializer）
单个服务器可以天然实现状态复制，可将单个服务器视为一个串行化器。通过让串行化器来分发命令，自动对请求排序并获得状态复制。

**算法 2.9：** 借助串行化器实现状态复制
* 所有客户端的请求都发送到串行化器，一次只发送一条
* 串行化器逐条转发命令到所有服务器
* 对于某条命令，当串行化器接收到所有服务器的响应时，再响应客户端

也被称为主从复制（Master-Slave Replication）不够去中心化*

### 两阶段协议（Tow-Phase Protocal）
**算法 2.10：** 两阶段协议
```js
// 阶段1
客户端向所有的服务器请求锁；
// 阶段2
if(该客户端获得了所有服务器的锁) {
  该客户端向每个服务器发送命令，随即释放锁；
}
else {
  该客户端释放已经获得的锁；
  等待一段时间，再重新进入阶段1；
}
```
* 不同的领域中有不同的名称，比如两段锁协议（Two-Phase-Locking）
  * ？：数据库中有两段锁协议，但似乎和两段锁协议有冲突，两阶段协议更像：一次封锁法

## 2.2 Paxos
### 票（Ticket）
一张票是一个弱化形式的锁
* 可重新发布：一个服务器可以随时发布新的票，哪怕前面发布的票还没有被释放。
* 票可以过期：当客户端使用一张票 t 来给服务器发送消息时，如果不是最新的票服务器不会接收。

另：
* 解决了宕机问题：若一个客户端在得到票后宕机，服务器可以发布新的票为其他客户端提供服务
* 票可以用计数器来实现：每当服务器收到一个（发布票）的请求时，将计数器+1，当某客户端尝试使用某个票时，服务器可判定该票是否已经过期
  * ？：新请求是否会使旧请求失效
  * ？：服务器可判定该票是否已经过期，应该是服务器可决定该票是否已经过期
      > the server can determine if the ticket is expired.

**算法 2.12：** 朴素的基于票的协议
```js
// 阶段1
客户端向所有服务器请求票；
// 阶段2
if(收到了绝大多数服务器的响应) {
  客户端将票和命令一起发送给每一个服务器；
  当票有效时，服务器存储命令并给该客户端一个正反馈响应；
}
else {
  客户端等待，然后重新进入阶段1；
}
// 阶段3
if(客户端收到了绝大多数服务器的正反馈响应) {
  客户端告诉所有服务器执行之前存储的命令
}
else {
  客户端等待，然后重新进入阶段1；
}
```
* 算法存在的问题：如果`u1`在各个服务器上存储了命令`c1`，但还没来得及通知所有服务器执行命令时，`u2`经过阶段2将部分服务器的命令更改为了`c2`，那么当`u1`通知服务器执行命令时，有的执行了`c1`，有的执行了`c2`。
* 一个想法：当服务器不仅能发布票，还能发布当前存储的命令时，当`u2`得知`u1`已经存储命令，`u2`就可以不让服务器存储`c2`，从而避免问题。
* 由于不同的服务器存储的命令可能不同，当没有支持数过半的命令时，客户端们可以支持任何命令，当存在支持数过半的命令时，所有的客户端必须支持这个过半的命令。

### Paxos算法

**算法 2.13：** Paxos

#### 初始化阶段

客户端（提案者）
```js
c; // 待执行命令
t = 0; // 当前尝试票号
```
服务端（接收者）
```js
Tmax = 0; // 已发布的最大票号
C; // 当前存储命令
Tstore = 0; // 当前存储的命令C对应的票
```
#### 阶段1
客户端（提案者）
```js
t = t + 1;
向所有服务器发消息，请求编号为t的票;
```
服务端（接收者）
```js
// 更新Tmax并响应当前存储的命令C及其票Tstore
if(t > Tmax) {
  Tmax = t;
  回复：ok(Tstore, C);
}
```
#### 阶段2
客户端（提案者）？
```js
if(过半数服务器回复ok) {
  选择Tstore最大的(Tstore, C); 
  if(Tstore > 0) {
    c = C;
  }
  向回复了ok的服务器发送消息：propose(t, c);
}
```
服务端（接收者）
```js
// 存储最大票的提案
if(t == Tmax) {
  C = c;
  Tstore = t;
  回复：success;
}
```
#### 阶段3
客户端（提案者）
```js
if(过半数服务器回复success) {
  向每个服务器发送消息：execute(c);
}
```
### Basic Paxos
#### 角色
* Client: 系统外部角色，请求发起者，不直接参与一致性算法
* Proposer: 接收Client请求，替民众提出议案(propose)
* Acceptor(Voter): 提议投票和接收者，只有在形成一定人数时(Quorum多数派)，提议才会被接受，类似于国会
* Learner: 提议接收者，对集群一致性无影响，作备份的记录员
#### 阶段
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

## Reference
[1] [Paxos 分布式共识算法的介绍和 Go 实现，论文笔记](https://www.bilibili.com/video/BV1C5411L7qT?from=search&seid=3598025483296716677&spm_id_from=333.337.0.0)