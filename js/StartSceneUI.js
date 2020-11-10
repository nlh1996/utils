cc.Class({
  extends: cc.Component,

  properties: {
    m_ok_btn:{
      default:null,
      type:cc.Button
    },
    // 头像显示
    sprite: cc.Sprite,
    setting:null,
    decideFlag:null,
  },

  // // LIFE-CYCLE CALLBACKS:
  onLoad () {
    // 远程 url 不带图片后缀名，此时必须指定远程图片文件的类型
    // cc.loader.load({url: '', type: 'png'}, function (err,tex) {
    //   if(err){
    //     cc.error(err);
    //     return;
    //   }
    //   if(cc.isValid(tex)==false){
    //       return;
    //   }
    //   if(cc.isValid(self,true)==false||cc.isValid(self.sprite,true)==false){
    //       return;
    //   }
    //   self.sprite.spriteFrame = new cc.SpriteFrame(tex)
    // });
  },

  auth() {
    let sysInfo = wx.getSystemInfoSync();
    console.log("微信信息",sysInfo);
    var top = sysInfo.statusBarHeight;
    if(top > 20){
      console.log("顶部距离 top：",top)
    }
    //获取微信界面大小
    let width = sysInfo.screenWidth;
    let height = sysInfo.screenHeight;
    cc.sys.localStorage.setItem('width', width)
    cc.sys.localStorage.setItem('height', height)
    let self = this

    wx.getSetting({
      success (res) {
        console.log(res.authSetting);
        if (res.authSetting["scope.userInfo"]) {
          console.log("用户已授权");
          wx.getUserInfo({
            success(res){
              console.log(res);
              self.userInfo = res.userInfo;
              //此时可进行登录操作
              self.login()
            }
          });
        }else {
          console.log("用户未授权");
          let button = wx.createUserInfoButton({
            type: 'image',
            image: "decide.png",
            // text: '',
            style: {
              left: width/2-75,
              top: height*0.75,
              width: 150,
              height: 45,
            }
          });
          button.onTap((res) => {
            if (res.userInfo) {
              console.log("用户授权:", res);
              self.userInfo = res.userInfo;
              //此时可进行登录操作
              button.destroy();
              self.login()
            }else {
              console.log("用户拒绝授权:", res);
            }
          });
        }
      }
    })
  },

  login() {
    console.log("login: ");
    wx.login({ 
      success (res) {
        if(res.code) {
          wx.request({
            url: 'https://yinghuo2018.com:10200/v1/onload',
            data: {
              code: res.code,
              ChannelName : null,
            },
            success (res) {
              // 成功响应
              console.log("login: success", res)
            },
            fail (err) {
              console.log("login: fail", err);
            }
          })
        }else {
          console.log('登录失败！' + res.errMsg)
        }
      }
    })
  },

  authHw(){
    let self = this;
    console.log("华为游戏登陆")
    hbs.gameLogin({
      forceLogin: 1, //强制登录，未登录时会弹出登录对话框
      appid: "102475943", //appid需要与华为开发者联盟后台配置一致
      success: function (res) {
        console.log("华为游戏game login: success res",res);
        let url = "https://yinghuo2018.com:10200/v1/hwload"//www.
        let playerId = self.intToStr(res.playerId);
        let displayName = self.intToStr(res.displayName);
        let playerLevel = self.intToStr(res.playerLevel);
        let isAuth = self.intToStr(res.isAuth);
        let ts = self.intToStr(res.ts);
        let gameAuthSign = self.intToStr(res.gameAuthSign);

        let data = {
          playerId: playerId,
          displayName: displayName,
          playerLevel: playerLevel,
          isAuth: isAuth,
          ts: ts,
          gameAuthSign: gameAuthSign,

          ChannelName : window.dataMgr.getChannelData().CHANNEL.HW,
        };
        var callback = function(tdata){
          console.log("华为游戏game callback",tdata);
          var json = self.strToJson(tdata);
          self.closeSelfClick();
          // //进入分区
          cc.sys.localStorage.setItem("account",json.account);
          cc.sys.localStorage.setItem("token",json.token);
          var call = function(){
            let startGameUI = window.uiMgr.getUI("startGameUI");
            startGameUI.getComponent("StartGameUI").setData(json);
          }
          window.uiMgr.openUI("startGameUI",GameEnum.LAYER.NORMAL,call);
        }
        window.dataMgr.getHttpData().Post(url,data,callback)
      },
      fail(data, code) {
        console.log("华为游戏on gameLogin fail: " + data + "," + code);
      },
      complete() {
        console.log("华为游戏on gameLogin: complete");
      }
    });
  },

  start () {
  },

});
