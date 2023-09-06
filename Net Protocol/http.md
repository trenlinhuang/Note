## HTTP协议
HTTP（Hypertext Transfer Protocol）定义了客户端和服务器进行报文交换的方式，以TCP作为支撑运输层协议。HTTP不保存关于客户的任何信息，是一个无状态协议。HTTP定义的是通信协议，如何解释web页面是浏览器的职责。

- 非持续连接：一个TCP连接只负责一个请求和一个响应的传输，结束后即关闭。每次请求响应，连接建立+数据传输总共需要两个RTT
- 持续连接：一个完整页面或多个完整页面在同一个持续TCP连接上进行，经过一定时间间隔未使用后会关闭连接

### HTTP报文格式
请求报文格式：
```
// 1. 请求行：方法，url，版本
GET /somedir/page.html HTTP/1.1 
// 2. 首部行：字段：值
Host: www.someschool.edu
Connection: close  // 不使用持续链接
Uesr-agent: Mozilla/5.0
Accept-language: en
// 空行
// 3. 实体体：
(data data data ...)
```

响应报文格式：
```
// 1. 状态行：版本，状态码，短语
HTTP/1.1 200 OK
// 2. 首部行：字段：值
Connection: close
Date: Tue, 18 Aug 2015 15:44:04 GMT
Server: Apache/2.2.3 (centos)
Content-Length: 6812
Content-Type: text/html
// 空行
// 3. 实体体：
(data data data ...)
```
