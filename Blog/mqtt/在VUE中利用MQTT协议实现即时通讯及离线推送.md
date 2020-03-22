## **前言**

建议先阅读：
[在Node.js下运用MQTT协议实现即时通讯及离线推送](https://segmentfault.com/a/1190000018002561)

以前尝试在vue中用上mqtt，了解到mqtt实质上是基于websocket进行数据通信，所以上文中在node下实现的服务端此时不能满足需求

## **代码**

**服务端: server.js**

    let http     = require('http')
    , httpServer = http.createServer()
    , mosca = require('mosca')

    let settings = {
      port: 5112,
      persistence:{
        factory: mosca.persistence.Mongo,
        url: "mongodb://localhost:27017/mosca"
      }
    }
    let server = new mosca.Server(settings)
    
    server.attachHttpServer(httpServer)
    server.on('published', function(packet, client) {
      console.log('Published',  packet.payload.toString());
    })
    httpServer.listen(3003)
    server.on('ready', function(){
      console.log('server is running at port 3003');  
    })
服务端mosca的实例化相较于在Node中的实现并没有改动
而是将其为websocket形式进行适配

**客户端: mqtt.js**

    let mqtt = require('mqtt')
    let client = {}
    export default {
      launch(id, callback) {
        client = mqtt('mqtt://ip', {
          port: 3003,
          clientId: id,
          clean: false
        })
        client.on('message', (topic, message) => {
          callback(topic, message)
        })
      },
      end() {
        client.end()
      },
      subscribe(topic) {
        client.subscribe(topic, {qos: 1})
        console.log('subscribe:', topic)
      },
      publish(topic, message) {
        client.publish(topic, JSON.stringify(message), {qos: 1})
      }
    }
独立地对mqtt进行简单地封装，方便调用
值得注意的是此时的协议头仍为mqtt，
但mqtt源码会以ws形式进行通信

**main.js:**
再把mqtt绑到vue原型链上

    import mqtt from './api/mqtt'
    Vue.prototype.$mqtt = mqtt

现在便可在vue环境中对mqtt进行调用

    this.$mqtt.launch(this.user._id, (topic, source) => {
        console.log('message: ', JSON.parse(source.toString()))
    })
转载请注明出处 ; )