export default class Advertisng {
  constructor(w, time) {
    this.width = w ? w: 375
    this.time = time ? time: 60
    this.ad = null
    this.left = 0
    this.top = 0
    this.id = ''
  }

  // banner广告
  initBanner(id, left, top) {
    this.id = id
    this.left = left
    this.top = top
    this.ad = window.wx.createBannerAd({
      adUnitId: id,
      adIntervals: this.time,
      style: {
        left: left,
        top: top,
        width: this.width,
      }
    })
    this.onLoad()
    this.onError()
  }

  // 激励视频
  initVideo(id) {
    this.ad = wx.createRewardedVideoAd({
      adUnitId: id
    })
    this.onLoad()
    this.onError()
    this.onClose()
  }

  // 插屏广告hy
  initInterstitial(id) {
    this.ad = wx.createInterstitialAd({
      adUnitId: id
    })
    this.onLoad()
    this.onError()
    this.onClose()
  }

  // 格子广告
  initGrid(id, left, top) {
    this.ad = wx.createGridAd({
      adUnitId: id,
      adTheme: 'white',
      gridCount: 5,
      style: {
        left: left,
        top: top,
        width: this.width,
        opacity: 0.8
      }
    })
    this.onLoad()
    this.onError()
  }

  // banner广告重新创建
  reLoad() {
    if (this.id == '') {
      console.log('目前只支持已创建的banner广告重置')
      return
    }
    this.destory()
    this.ad = window.wx.createBannerAd({
      adUnitId: this.id,
      adIntervals: this.time,
      style: {
        left: this.left,
        top: this.top,
        width: this.width,
      }
    })
    this.onLoad()
    this.onError()
  }

  show() {
    this.ad.show().then(() => console.log('banner 广告显示')).catch(err => console.log(err))
  }
  hide() {
    this.ad.hide()
  }
  destory() {
    this.ad.destory()
    this.ad = null
  }
  onLoad() {
    this.ad.onLoad(() => {
      console.log('banner 广告加载成功')
    })
  }
  onError() {
    this.ad.onError(err => {
      console.log(err)
    })
  }
  onClose() {
    this.ad.onClose(res => {
      console.log('广告关闭')
      if (res && res.isEnded || res === undefined) {
        // 正常播放结束，可以下发游戏奖励
      } else {
        // 播放中途退出，不下发游戏奖励
      }
    })
  }
}
