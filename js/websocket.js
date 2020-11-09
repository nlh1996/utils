export default {
  sock: null,
  handlerFuc: new Map(),

  on_open: function () {
    console.log("ws connected...")
  },
  
  on_message: function (event) {
    let data = JSON.parse(event.data)
    switch(data.status) {
      // ping 消息
      case 100:
        break
      // 返回消息成功
      case 200:
        this.handlerFuc.get(data.id)(data.data)
        this.handlerFuc.delete(data.id)
        break
      // 服务器主动推送消息,接受消息需要注册对应的方法
      case 300:
        this.handlerFuc.get(data.id)(data.data)
        break
      // 消息出错
      case 400:
        console.log(data)
        this.handlerFuc.delete(data.id)
        break
      default:
        console.log("Unexcepted error !")
    }
  },

  on_close: function () {
    this.close()
  },

  on_error: function () {
    this.close()
  },
  
  close: function () {
    if(this.sock) {
      this.sock.close()
      this.sock = null
    }
  },

  connect: function (url) {
    this.sock = new WebSocket(url)
    this.sock.binaryType = "arraybuffer"
    this.sock.onopen = this.on_open.bind(this)
    this.sock.onmessage = this.on_message.bind(this)
    this.sock.onclose = this.on_close.bind(this)
    this.sock.onerror = this.on_error.bind(this)
  },

  send: function (data, cb) {
    if(this.handlerFuc.has(data.id)) {
      console.log("当前网络不稳定,请稍后发送！！！")
      return
    }
    this.sock.send(JSON.stringify(data))
    this.handlerFuc.set(data.id, cb)
  },

  register: function (id, cb) {
    if(this.handlerFuc.has(id)) {
      console.log(id + "该消息已经注册过")
      return
    }else {
      this.handlerFuc.set(id, cb)
    }
  }
}