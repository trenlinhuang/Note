## TCP报文段结构
```
=======32比特=======
源端口号 | 目的端口号
       序号
      确认号
首部长|保留|flag|接收窗口
校验和   |紧急数据指针
       选项
       数据
=======32比特=======
```

- 32比特序号和确认号字段用于实现可靠数据传输
- 4比特首部长度字段，以32比特为单位。由于TCP选项字段的原因，TCP首部长度是可变的（典型长度20字节）
- flag为标志位，包括：
  1. CWR
  2. ECE
  3. URG
  4. ACK
  5. PSH
  6. RST: 用于连接建立和拆除
  7. SYN: 用于连接建立和拆除
  8. FIN: 用于连接建立和拆除
- 16比特接收窗口用于流量控制

### 序号和确认号
TCP把数据看成一个无结构的、有序的字节流。一个报文段的序号是该报文段首字节的字节流编号。

如果限制每个报文段最多包含1000个字节，那么一个500,000个字节的文件将被拆成500个报文段，序号依次为0, 1000, 2000, ..., 490000

主机A向主机B发送数据的同时，主机B也可能在向主机A发送数据，确认号用于由主机A向主机B指明希望收到的下一个字节。

## TCP连接管理
### 连接建立
- 第一步：（SYN报文段）客户端TCP发送一个不包含应用层数据的报文段，SYN被置为1。TCP报文段序号被置为一个随机数client_isn
- 第二步：（SYNACK报文段）服务器为该TCP连接分配缓存和变量，向客户端发送允许连接的报文段，确认号字段被设为client_isn+1，序号字段被设为随机数server_isn，SYN比特被置为1
- 第三步：客户端收到SYNACK报文段后，客户也要给该连接分配缓存和变量。客户主机则向服务器发送另外一个报文段，确认号字段被设为server_isn+1，因为连接已经建立，SYN比特被置为0。该报文段可以在报文负载中携带客户到服务器的数据。

![tcp%20connection%20establish](./tcp/tcp%20connection%20establish.png)

### 连接释放
1. 客户都安打算关闭连接时，客户TCP会向服务器进程发送一个特殊的TCP报文段。这个特殊的报文段让其首部中的一个标志位FIN比特被置为1。服务器收到该报文段后，服务器发送一个确认报文段。
2. 随后服务器发送它自己的中止报文段，其FIN比特被置为1。最后，客户端再对这个中止报文段进行确认。

![tcp%20connection%20close](./tcp/tcp%20connection%20close.png)

即四次挥手