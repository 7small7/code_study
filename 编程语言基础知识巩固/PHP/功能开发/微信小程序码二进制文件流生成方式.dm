## 文章说明

在一次项目中，需要使用到微信小程序二维码生成的接口，微信Api返回的是一个二进制文件流的格式。在向前端返回时，可以输出一个图片的格式，也可以输出二进制文件流的格式。
输出二进制文件流的好处在于不需要服务端存储图片，如果是输出图片，就需要服务端事先写入一个图片，然后返回给前端一个图片地址，这样增加了服务器的处理逻辑与存储。
最终的解决方案是输出一个二进制文件流。

## 服务端

这里使用PHP代码生成。
```php
<?php
// 微信小程序公钥和秘钥
define('id', '');
define('secret', '');

function curlPost($url, $postData)
{
    $header = array (
        'Accept: application/json',
    );
    $curl   = curl_init();
    curl_setopt($curl, CURLOPT_URL, $url);
    curl_setopt($curl, CURLOPT_HEADER, 0);
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1);
    curl_setopt($curl, CURLOPT_TIMEOUT, 10);

    curl_setopt($curl, CURLOPT_HTTPHEADER, $header);
    curl_setopt($curl, CURLOPT_SSL_VERIFYPEER, FALSE);
    curl_setopt($curl, CURLOPT_SSL_VERIFYHOST, FALSE);
    curl_setopt($curl, CURLOPT_POST, 1);
    curl_setopt($curl, CURLOPT_POSTFIELDS, $postData);
    $data = curl_exec($curl);
    curl_close($curl);
    return $data;
}

function curlGet($url)
{
    $header = array (
        'Accept: application/json',
    );
    $curl   = curl_init();
    curl_setopt($curl, CURLOPT_URL, $url);
    curl_setopt($curl, CURLOPT_HEADER, 0);
    curl_setopt($curl, CURLOPT_TIMEOUT, 1);
    curl_setopt($curl, CURLOPT_HTTPHEADER, $header);
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1);
    curl_setopt($curl, CURLOPT_SSL_VERIFYPEER, false);
    curl_setopt($curl, CURLOPT_SSL_VERIFYHOST, false);
    $data = curl_exec($curl);
    curl_close($curl);

    return json_decode($data, true);
}


function getAccessToken()
{
    $appid  = id;
    $secret = secret;
    $url    = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=$appid&secret=$secret";
    return curlGet($url)['access_token'];
}

function getImage()
{
    $token = getAccessToken();
    $url   = "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=$token";
    return curlPost($url, json_encode(['path' => 'pages/index/index'], JSON_UNESCAPED_UNICODE));
}

// 返回二进制文件流
header("Content-type:image/gif,image/png,image/jpg,image/jpeg");
echo getImage();
```

## 小程序端

小程序端首选需要你创建一个wxml的文件和wxjs的文件。
js文件内容:
```javascript
// index.js
// 获取应用实例
const app = getApp()

Page({
  data: {
    imageUrl: '/images/7acb0a46f21fbe09f646230a4473a8358744adaa.jpeg',// 用来默认展示的图片
    imageBuffer: '',// 用来后端存储的二进制流文件。
  },
  onLoad() { this.getImageUrl() },
  getImageUrl() {
    let th = this
    wx.request({
      url: 'http://localhost/image.php',// 这里是后端服务接口
      responseType: 'arraybuffer',
      success(res) {
        th.setData({
          imageBuffer: res.data,//组件库2.4之后，不能使用微信够逼玩意的函数。"data:image/png;base64," + wx.arrayBufferToBase64(res.data)
        })
      }
    })
  },
  saveImage() {
    let that = this
    var fileManager = wx.getFileSystemManager()
    var fielName = wx.env.USER_DATA_PATH + '/' + Math.ceil(Math.random() * 1000) + '.png'// 随机生成一个图片名称
    fileManager.writeFile({
      filePath: fielName,
      data: that.data.imageBuffer,//如果是一个base64的数据的话，就使用这个分割一下。that.data.imageUrl.slice(22),
      encoding: 'base64',
      success(res) {
        wx.saveImageToPhotosAlbum({
          filePath: fielName,
        })
      }
    })
  }
})
```
wxml文件内容:
```xml
<view>
  <image src="{{imageUrl}}"></image>
  <button bindtap="saveImage">保存图片</button>
</view>
```