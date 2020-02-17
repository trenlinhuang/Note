## **前言**

前些日子了解到mqtt这样一个协议，可以在web上达到即时通讯的效果，但网上并不能很方便地找到一篇**目前版本**的在node下正确实现这个协议的博客。

自己捣鼓了一段时间，理解不深刻，但也算是基本能够达到使用目的。

本文目的为对MQTT有需求的学习者提供一定程度上的便利，省去了查阅官方文档及源码的功夫，但尚未对离线消息的接收顺序进行处理。

## **代码**

**服务端: server.js**

    //服务端引入中间件mosca
    let mosca = require('mosca')
    let settings = {
      port: 5112
    }
    let server = new mosca.Server(settings)
    server.on('ready', function(){
        console.log('Mosca server is up and running at port 5112');  
    })
    server.on('published', function(packet, client) {
      console.log('Published', packet.payload)
    })
    
    server.on('clientDisconnected', function(client){
      console.log('disconnected: ', client.id)
    })
**推送端: pub.js**

    //客户端引入mqtt
    let mqtt = require('mqtt');
    
    let client = mqtt.connect('mqtt://localhost', {
      port: 5112,
      clientId: 'cli_pub',
    })
    
    let num = 0;
    setInterval(function (){
      client.publish('test', 
      'Hello mqtt ' + (++num),
      {qos:1},
      () => console.log(num));
    }, 1000)
  
  **订阅端: sub.js**

    let mqtt = require('mqtt')
    
    let client = mqtt.connect('mqtt://localhost', {
      port: 5112,
      clientId: 'cli_sub',
    })
    
    client.subscribe('test',{qos:1})
    
    client.on('message', function (topic, message) {
      console.log('received message: ', message.toString())
    })
---
server运行后，先启动sub，再启动pub，即可在sub中接收到推送过来的消息序列
**至此实现了简单的即时推送**

## **离线推送相关配置及简要介绍**

离线配置-服务端：
---

要实现消息的离线推送，必然需要一个存储临时数据的部件
此处用到的是mongo，当然可以根据需要选择其他的存储工具

**server.js中的settings需更改为:**

    let settings = {
      port: 5112,
      persistence:{    //增加了此项
        factory: mosca.persistence.Mongo,
        url: "mongodb://localhost:27017/mosca"
      }
    }
factory: 引入mosca对特定存储工具的一些处理方法
url: 其中的 27017 为mongo所监听的端口号，mosca为存储相关数据的数据库

值得一提的是：配置好mongo的环境后，不需要提前在mongo中手动创建，若数据库不存在会自动生成，而且mosca会为你作好其他一切基本事项 **(即：若只想临时体验下效果，甚至可以暂时把mongo放一边 )**

在mongo中，可以看到自动新添了db: mosca
及其下的collection(相当于关系型数据库中的表/关系)

![clipboard.png](https://image-static.segmentfault.com/174/846/1748466158-5c487e8e29a75_articlex)

离线配置-客户端：
---

**pub.js和sub.js中的client中都可以改为：**

    let client = mqtt.connect('mqtt://localhost', {
      port: 5112,
      clientId: 'cli_**',
      clean: false//增加了此项
    })
clientId: 区分客户端的识别码
clean: 此处决定了客户端在服务端的session是否会被清除，默认为true，为实现离线推送，我们需要将其保留
**clean及上文中的persistence为实现离线推送的关键配置**

mqtt.connect()会返回一个mqttClient对象，包含了：reconnect(), subscribe(), publish()等一系列方法。

本文中发送端接收端被分为了pub.js和sub.js两个独立文件，仅仅为了方便在不同控制台中观察效果
一个client可以既为推送端，又为订阅端

---
**至此，所有代码已完成**

其他介绍：
---

**client.subscribe():**
为本客户端订阅一个话题，所有订阅此话题的用户都会收到在此话题下推送的信息

    //client.subscribe(topic,opts)
    client.subscribe('test',{qos:1})
opts中的qos为通信机制，控制发送端与接收端的互锁程度
上文中的其中一个collection: subscriptions即记录各用户话题订阅情况

用户cli_sub及cli2_sub订阅了话题test:
![clipboard.png](https://segmentfault.com/img/bVbnHio?w=563&h=134)
(新增一个cli2_pub，下文有用)

*注：*
*重复执行脚本sub.js实际上对topic进行了重复订阅*
*实际编码时，应避免topic的重复订阅，即使重复订阅并不影响实现效果*

**client.publish():**

向指定topic发送数据
*message为Buffer或String格式，可以通过序列化或转json实现对复杂数据对象的传送*

    //client.publish(topic, message, opts, callback)
    let num = 0;
    setInterval(function (){
      client.publish('test', 
      'Hello mqtt ' + (++num),
      {qos:1},
      () => console.log(num));
    }, 1000)
参数不再赘述
此处用一个定时器定时在 topic: test 下发送'Hello mqtt 1,2,3..'

用回调函数实时打印一下发送的num：
![clipboard.png](https://segmentfault.com/img/bVbnHol?w=268&h=48)

当订阅者处于离线状态时，可以在collection: packets中查看到临时数据的存储情况：
![clipboard.png](https://segmentfault.com/img/bVbnHm5?w=992&h=171)
mosca把每一条推送消息为所有订阅用户都生成了独立的记录，用同一个messageId进行关联

当其中一个用户(cli2_sub)上线时，获取到其对应的数据，
![clipboard.png](https://segmentfault.com/img/bVbnHok?w=269&h=52)
而后数据库中相应记录便会被删除
![clipboard.png](https://segmentfault.com/img/bVbnHnS?w=997&h=107)
此时仅有cli_sub用户的数据
当cli2_sub上线接收消息后，packets中记录将被清空

**client.on():**
即在client上触发的事件，此处只列举消息接收事件

    //client.on(event, callback)
    client.on('message', function (topic, message) {
      console.log('received message: ', message.toString())
    })
处理为简单地打印到控制台

## **附**

>mosca.js文档：
https://www.npmjs.com/package/mosca

>mqtt.js文档：
https://www.npmjs.com/package/mqtt

>windows环境下mongo的配置：
https://jingyan.baidu.com/article/6b97984dbeef881ca2b0bf3e.html

>及一位前辈的文章：
https://www.jianshu.com/p/8315acec4e6b

转载请注明出处 ; )