[TOC]

## 前言

对于开发的小伙伴们来说，代码包已经是开发过程中的家常便饭了。比如，PHP有composer，Java有maven，前端有npm，yarn，Mac程序安装有brew，Linux程序安装有yum。使用这些包，可以便于我们很好的管理代码引入的外部代码组件，有助于我们提高开发效率，同时也有助于我们代码管理更加优雅。下文主要就是分享个人自定义一个composer代码包，代码仅仅是作为演示示例，没有实际效果。同时也会不断更新，大家可以关注关注。

## 具体实现

### 创建一个远程git代码仓库。

创建代码仓库，我是选择的GitHub作为远程仓库。创建好后，我们直接拉取到本地。后面的所有代码都是在该仓库下进行的。具体创建仓库和拉取代码就不做演示。

### 创建源码目录

我们在仓库(为了方便描述，后文便用项目一词来替代该词)下面，创建一个src的目录，用于存储我们实际的代码。
![](https://oscimg.oschina.net/oscnet/up-b024457e4d6df4b6730bb582cf967b0e3d7.png)

### 编写实际代码

下面这个文件，是后面实际需要调用的类。具体的演示代码，可以通过[演示代码仓库](https://github.com/bruceqiq/composer_test "演示代码")进行查看。主要的作用就是利用工厂模式，Cache去调用实际的缓存类。
```php
<?php
// composer演示代码
declare(strict_types=1);

// 这里的命名空间根据自己情况填写，后面生成的composer.json文件需要使用该命名空间。
namespace Bruce;

use Bruce\Client\Redis;

class Cache
{
    public $redisHandle = '';

    public function __construct()
    {
        $this->redisHandle = new Redis();
    }
}
```
创建好代码仓库中的所有文件，这时候我们就可以开始生成composer.json文件了。
![](https://oscimg.oschina.net/oscnet/up-e15d9676f2c3f1deacd636c50a3b81d3f9f.png)
### 生成composer.json配置文件

下面就是具体生成的过程了，记住一定要在项目的根目录下面使用composer init命令。
```shell
composer init
Welcome to the Composer config generator
This command will guide you through creating your composer.json config.

#项目命名空间
Package name (<vendor>/<name>) [bruce_redis/redis]: bruce_redis/redis
#项目描述
Description []: composer test
#作者信息
Author [卡二条 <2665274677@qq.com>, n to skip]: 卡二条 <2665274677@qq.com>
#输入最低稳定版本
Minimum Stability []: dev
#项目类型
Package Type (e.g. library, project, metapackage, composer-plugin) []: library
#授权类型
License []: 
Define your dependencies.

#依赖信息
Would you like to define your dependencies (require) interactively [yes]? yes
#如果需要依赖，则输入要安装的依赖
Search for a package: php
Enter the version constraint to require (or leave blank to use the latest version): >=7.0
Search for a package: 
Would you like to define your dev dependencies (require-dev) interactively [yes]? yes
Search for a package: php
Enter the version constraint to require (or leave blank to use the latest version): >=7.0
Search for a package: 
{
  "name": "bruce_redis/redis",
  "description": "composer test",
  "type": "library",
  "require": {
    "php": ">=7.0"
  },
  "require-dev": {
    "php": ">=7.0"
  }

#确认构建项目，生成composer.json
Do you confirm generation [yes]? yes
Would you like the vendor directory added to your .gitignore [yes]? yes
Would you like to install dependencies now [yes]? yes
Loading composer repositories with package information
Updating dependencies (including require-dev)
Nothing to install or update
Writing lock file
Generating autoload files
```
> 这里需要注意一下<vendor>/<name>配置项

根据[composer中文网](https://docs.phpcomposer.com/04-schema.html#package-name "官方文档")翻译的是供应商名称和项目名称。实际指的就是，你在其他项目中安装该composer包时，在你其他项目的vendor目录下面名称。例如，我将示例的代码的该配置项，设置为bruce_redis/redis。我再其他项目使用该包，展现的形式就是下图的效果。
![](https://oscimg.oschina.net/oscnet/up-ae1c8986128319b0b9937577aa14c96938f.png)
### 配置自动加载

在生成composer.json文件之后，我们的代码被安装到其他项目，还不能直接使用，因为不能为composer自动加载，我们还需要做如下配置，在composer.json文件中添加如下代码。
```php
"autoload": {
    "psr-4": {
      "Bruce\\": "src/"
    }
}
```
该配置项主要的目是为了，告知composer命名空间为Bruce的类，源码目录都在src下面,autoload_psr4.php文件再做文件加载时，就是根据这个目录去下载。这里的Bruce就是编写实际代码中提到的命令空间。如果你不是这个名字，改为自己的即可。例如，你写的是Alibab\Composer,自动加载中的配置就是下面的配置:
```php
"autoload": {
    "psr-4": {
      "Alibab\Composer\\": "src/"
    }
}
```

### 发布代码

做好上面的配置之后，我们就可以将代码发布到[packagist](https://packagist.org/ "packagist")。首先我们需要将代码更新到GitHub仓库。登录packagist，然后点击右上角的submit按钮，在输入框中输入GitHub仓库地址即可。
![](https://oscimg.oschina.net/oscnet/up-a3b395ccc5f4706d8a196ca1a3afbfd11c1.png)
## 效果演示

### 安装代码
发布好代码之后，就可以直接在项目中引入该包。
> 如果你是使用的composer阿里源，或许不能立即使用，因为阿里的源还没完全同步过来。可切换为官方源。

```shell
composer require bruce_redis/redis dev-master
```
使用该命令之后，出现下图的结果，则证明安装成功了。你就可以按照GitHub仓库里面的示例代码进行操作了。
![](https://oscimg.oschina.net/oscnet/up-c093220abe5830e404e7260efc6ba283ade.png)
### 自动加载

在上面提到composer会自动加载，通过下图，我们可以发现composer根据配置自动去加载了文件。
![](https://oscimg.oschina.net/oscnet/up-eac2a79a8e39f55523eaaa516b70cb49fa8.png)
> 这里的/bruce_redis/redis/src值就是,composer.json里面的name(<vendor>/<name>)值。