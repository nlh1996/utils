class CusHttp {
  constructor () {
    this._http = cc.loader.getXMLHttpRequest()
    this._callback
  }
  /**
   * Get 请求
   * @param {*} Url 
   * @param {*} cb 
   */
  Get (Url, cb) {
    this._http.open("GET", Url, true)
    this._callback = cb
    this._http.onreadystatechange = this.result.bind(this)
    this._http.send()
  }
  Post (Url, data, cb) {
    data = JSON.stringify(data)
    this._http.open("POST", Url, true)
    this._callback = cb
    this._http.onreadystatechange = this.result.bind(this)
    this._http.send(data)
  }
  result () {
    if (this._http.readyState == 4) {
      if (this._http.status == 200) {
        if (this._callback) {
          this._callback(JSON.parse(this._http.response))
        }
      }else{
        console.log(this._http.response)
      }
    }
  }
}

const http = new CusHttp()
http._http.setRequestHeader("Content-Type", "application/json;charset=UTF-8")
http._http.timeout = 10000

export default http
