> 专注于PHP、MySQL、Linux和前端开发，感兴趣的感谢点个关注哟！！！文章整理在[GitHub](https://github.com/bruceqiq/code_study),主要包含的技术有PHP、Redis、MySQL、JavaScript、HTML&CSS、Linux、Java、Golang、Linux和工具资源等相关理论知识、面试题和实战内容。

目前，PHP是用于Web开发的最流行的脚本语言。你可以在互联网上随手找到关于PHP大量资料，包括文档、教程、工具等等。PHP不仅是一种功能丰富的语言，它还能帮助开发人员轻松地创建更好的网络环境。
该文将总结几款PHP非常实用的类库。

## PhpFastCache

phpFastCache是一个开源的PHP缓存库，只提供一个简单的PHP文件，可方便集成到已有项目，支持多种缓存方法，包括：apc, memcache, memcached, wincache, files, pdo and mpdo。可通过简单的API来定义缓存的有效时间。
![WX20210623-202016@2x](https://gitee.com/bruce_qiq/picture/raw/master/2021-6-23/1624450836011-WX20210623-202016@2x.png)

官方地址：https://www.phpfastcache.com/


## pChart

pChart是一个基于GD library（图形处理函数库）开发的PHP图表制作开源项目。支持多种图表类型包括：
1. Line chart
2. Cubic curve chart
3. Plot chart
4. Bar chart
5. Filled line chart
6. Filled cubic curve chart
7. Pie chart
8. Radars chart
9. Limits chart

![截屏2021-06-23 下午8.23.52](https://gitee.com/bruce_qiq/picture/raw/master/2021-6-23/1624451040940-%E6%88%AA%E5%B1%8F2021-06-23%20%E4%B8%8B%E5%8D%888.23.52.png)

官网地址：http://www.pchart.net/features

## Munee

Munee是一个集图片尺寸调整、CSS-JS合并/压缩、缓存等功能于一身的PHP库。可以在服务器端和客户端缓存资源。它集成了PHP图片操作库Imagine来实现图片尺寸调整和剪切，之后进行缓存。
Munee可以自动编译LESS、SCSS和CoffeeScript，并且可以把CSS+JS文件合并成一个单一的请求，此外，还可以对这些CSS-JS文件进行精缩，让其拥有更好的性能表现。该库还可以轻易地与任何代码集成。

![Snipaste_2021-06-24_16-53-17](https://gitee.com/bruce_qiq/picture/raw/master/2021-6-24/1624524814129-Snipaste_2021-06-24_16-53-17.png)

官网地址：http://mun.ee/

## gantti

gantti, 一个简单的PHP甘特图类 Gantti一个简单的PHP甘特图类特性生成有效的HTML5使用SASS样式表定制非常容易在包括 IE7，IE8和IE9在内的所有主流浏览器中工作不需要 javascript
![v2-1dac14575e5c22357f297f763a4e8148_r](https://gitee.com/bruce_qiq/picture/raw/master/2021-6-24/1624524996845-v2-1dac14575e5c22357f297f763a4e8148_r.jpg)

官网地址：https://github.com/bastianallgeier/gantti

## whoops

是PHP的错误处理器框架。开箱即用，它提供了一个相当错误的界面，可以帮助你调试您的 Web 项目，但在n内核它是一个简单而强大的堆叠错误处理系统。
1.灵活、基于堆栈的错误处理
2.独立图书馆（目前）无需依赖
3.用于处理异常、跟踪帧及其数据的简单 API
4.包括一个漂亮的rad错误页面为您的webapp项目
5.包括直接在编辑器和 IDE 中打开引用文件的能力
6.包括不同响应格式的处理程序（杰森、XML、SOAP）
7.易于扩展和集成现有库
8.清洁、结构良好和经过测试的代码基础

![687474703a2f2f692e696d6775722e636f6d2f305651706539362e706e67](https://gitee.com/bruce_qiq/picture/raw/master/2021-6-24/1624525150131-687474703a2f2f692e696d6775722e636f6d2f305651706539362e706e67.png)

官网地址：https://github.com/filp/whoops

## php-image-cache

图像缓存是一个微小的PHP类，接受.png、.jpg或.gif图像，然后压缩、移动和缓存用户浏览器中的图像。然后，它将返回图像的新源，以打印成图像标签。

通过压缩和缓存图像，页面加载时间可以显著缩短。页面加载时间是用户保留的最大因素之一，我们都看到了关于加载时间如何影响公司底线的研究。但是，当尝试计算这些因素时，我很难找到一个简单而直接的 PHP 类来缓存和加载图像。所以我做了一个。

下面是本类可以产生差异的示例。下面，左侧是本地测试环境中在 2.19 秒内加载的大型.png文件的屏幕截图。右边是脚本运行后拍摄的屏幕截图，图像被压缩和缓存，在 23 毫秒内加载。由于脚本，图像加载时间缩短了2167毫秒-这是资源加载时间的98.95%减少！

![ss1-full](https://gitee.com/bruce_qiq/picture/raw/master/2021-6-24/1624525252338-ss1-full.jpg)

官网地址：https://nielse63.github.io/php-image-cache/

## Imagine

Imagine是一个面向对象的PHP类库，用于图片操作。这个类库能够处理一些常用到的操作如：调整大小、裁剪、应用过滤器等。其Color类库可用于对任意对定的颜色生成RGB值。并且还提供一些方法来绘制图形如：圆弧，椭圆，线，片等。此外，还可以利用一个灵活的字体类来加载任意字体文件，然后将文字插入到图片中。
![d9dd65ed7f76cca1775c18db4bcb6cb3](https://gitee.com/bruce_qiq/picture/raw/master/2021-6-24/1624525430617-d9dd65ed7f76cca1775c18db4bcb6cb3.jpg)

官网地址：https://github.com/avalanche123/Imagine

