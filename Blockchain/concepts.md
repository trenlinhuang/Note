#### 共识（consensus）
系统中有n个节点，其中最多有f个节点可能崩溃。节点i从一个输入值vi开始，所有节点最终必须要从全部输入值中最终选择一个决策值。
* 一致性（Agreement）：所有好节点的决策值必定相同
* 可终止性（Termination）：所有好节点在有限的时间内结束决策过程
* 有效性（Validity）：选择出的决策值必须是某个节点的输入值

#### 原子广播 (atomic broadcast)
原子广播是一个分布式原语，其保证各节点收到并相同次序的消息并处理或者没有副作用的中止操作。即，保证消息在分布式系统中的原子性。在容忍崩溃的系统中，原子广播等同与共识.
* 有效性：正常节点发出的消息，会被其他所有节点接收到
* 一致性：一个正常的节点收到消息，则其他正常节点最终也会收到该消息
* 完整性：一个消息对于每个节点来说，最多接收到一次
* 顺序性：各节点收到的消息顺序和消息发出的顺序是一致的
#### CAP定理
加州大学的计算机科学家 Eric Brewer 提出，分布式系统有三个指标：Consistency（一致性）、 Availability（可用性）、Partition tolerance（分区容错性），即一个分布式系统不可能同时具有以上三个特征。
##### Partition tolerance
大多数分布式系统都分布在多个子网络。每个子网络就叫做一个区（partition）。分区容错的意思是，区间通信可能失败。比如，一台服务器放在中国，另一台服务器放在美国，这就是两个区，它们之间可能无法通信。

一般来说，分区容错无法避免，所有的分布式系统应考虑分区容错性。因此可以认为 CAP 的 P 总是成立。CAP 定理告诉我们，剩下的 C 和 A 无法同时做到。
##### Consistency
Consistency 中文叫做"一致性"。意思是，对某一节点执行写操作之后，所有节点上的读操作必须返回该值。即在客户端看来所有的服务器节点应具有一致的状态。
##### Availability
Availability 中文叫做"可用性"，意思是只要收到用户的请求，服务器就必须给出回应。

用户可以选择向服务器 G1 或 G2 发起读操作。不管是哪台服务器，只要收到请求，就必须告诉用户，到底是 v0 还是 v1，否则就不满足可用性。

#### Safety & Liveness
* Safety: The replicated service satisfies linearizability
* Liveness: Clients eventually receive replies to their requests
#### 线性一致性 (Linearizability)
操作多节点系统就像操作单节点系统一样
> It behaves like a centralized implementation that executesoperations atomically one at a time

#### 网络假设
##### 同步网络
网络中的消息能够在一个已知的时间 Δ 内到达，这是一种非常理想的假设，实际的工程实践中很难保证这一假设成立。
##### 半同步网络
在一个 GST（global stabilization time）事件之后，消息在 Δ 时间内到达。随着网络技术的发展，这种假设已经能符合我们生活中遇到多绝大多数网络情况。
##### 异步网络
只保证消息最终能到达，并没有到达时间的限制，这几乎涵盖所有网络情况。由此可见，异步 BFT 协议的鲁棒性最强，即使在极端网络条件下协议也不会丧失活性。