[TOC]

## Composer介绍

Composer 是 PHP 的一个<font color='red'>包依赖</font>管理工具。我们可以在项目中声明所依赖的外部工具库，Composer 会帮你安装这些依赖的库文件，有了它，我们就可以很轻松的使用一个命令将其他人的优秀代码引用到我们的项目中来。Composer 默认情况下不是全局安装，而是基于指定的项目的某个目录中(例如 vendor)进行安装。Composer 需要 PHP 5.3.2+ 以上版本，且需要开启 openssl。Composer 可运行在 Windows 、 Linux 以及 OSX 平台上。

## Composer安装

### Windows 安装
Wondows 平台上，我们只需要下载 Composer-Setup.exe 后，一步步安装即可。需要注意的是你需要开启 openssl 配置，我们打开 php 目录下的 php.ini，将 extension=php_openssl.dll 前面的分号去掉就可以了。
![windows安装](http://qiniucloud.qqdeveloper.com/How-to-Install-Composer-on-Windows-Specify-PHP-File-Location.png)
安装完成之后，检测是否安装成功。可以使用

```shell
composer --version
```

命令查看,如下图:
![检测是否安装成功](http://qiniucloud.qqdeveloper.com/Test-Whether-The-Composer-is-Successfully-Installed.png)
### Linux 安装

```shell
// 下载composer文件
php -r "copy('https://install.phpcomposer.com/installer', 'composer-setup.php');"
// 使用PHP解释器安装composer
php composer-setup.php
// 移动到系统可执行文件目录，方便我们后期直接使用composer命令进行全局调用
mv composer.phar /usr/local/bin/composer
```

### Mac Os 安装

```php
------直接安装
// 下载并安装
curl -sS https://getcomposer.org/installer | php
// 移动到可执行文件目录，便于全局调用
sudo mv composer.phar /usr/local/bin/composer
------使用Mac上面的brew包管理工具安装
brew install composer
// 检测是否安装成功
composer --version
```

### 如何切换 composer 镜像源
现在阿里处理自己的 composer 镜像源，并且能够做到与 Packagist 官网实时同步，推荐使用阿里的 composer 镜像源.

```shell
// 切换镜像源
composer config -g repo.packagist composer https://mirrors.aliyun.com/composer/
// 取消配置
composer config -g --unset repos.packagist
// 只针对当前项目切换镜像源(非全局切换)
composer config repo.packagist composer https://mirrors.aliyun.com/composer/
// 取消当前项目的镜像源
composer config --unset repos.packagist
```

### composer 更新
composer 的更新可以使用自身的命令来更新

```shell
composer selfupdate
```

## Composer 使用

Composer 的使用，我们常见的几个操作如下：

### composer install
当我们的 composer.json 文件中增加了项目的依赖关系,如下：

```shell
{
    "require": {
        "topthink/framework": "5.0.5",
    }
}
```

此时我们使用 composer install 时，会自动根据包中的依赖关系，来安装相对应的包。

### composer update
该命令会根据 composer.json 配置文件中包依赖以及相应的版本，更新包的版本，但是该命令会将所有的包都更新到最新版本，在实际的项目中需要谨慎使用，尤其是在生产环境上面。

### composer require
第 1 点中，我们讲到了如何去根据 composer.json 配置文件中的包依赖配置，安装对应的包。该命令可以不需要根据配置文件，而是去手动安装一个包。安装完之后，对应的依赖配置会自动添加在 composer.json 配置文件中。

### composer remove
该命令可以移除指定的包依赖，例如我们不需要依赖某个包直接使用该命令+包名

```shell
composer remove "topthink/framework": "5.0.5"
```

同样的，执行该命令之后，composer.json 配置文件中的包依赖会自动删除,无需我们手动操作。

### composer show
该命令主要是为了查看当前项目已经安装的包相关信息。

```shell
// 查看所有已经安装的包的信息
composer show
// 指定查看已经安装的包的信息
composer show topthink/framework
```

## Composer.json 与 Composer.lock 的区别是什么

我们在使用 composer 的过程中会发现，当我们执行 composer update 等命令，会发现无意中多了一个 composer.lock 文件。那这个文件主要是干什么的呢？该文件主要是管理包版本使用的，当我们在使用 composer update 命令时，composer 会自动根据 composer.json 的包版本依赖，生成对应的 composer.lock 文件，当我们下次在执行 composer 命令的时候，首先也会去读取 composer.lock 文件的内容。

## Composer 版本约束

在我们使用 composer 安装包时，不得不考虑的就是一个版本问题，因为不同的版本，存在兼容性问题，因此我们在使用该工具安装包时需要特别的注意包版本，如果使用不当很容易导致项目因为包版本问题瘫痪。常见的几种如下： 1.精准版本
明确要安装到那个版本,如需要安装包的版本是 1.2.3

```shell
"topthink/think-captcha": "v1.2.3",
```

2.通配符
既满足指定范围即可,如下范围在 5.0 到 5.1 之间

```shell
"topthink/framework": "5.0.*",
```

3.范围
范围常用的操作符有>，>=，<，<=，!=。你可以定义多个范围，使用空格或者逗号 , 表示逻辑上的与，使用双竖线 || 表示逻辑上的或。其中与的优先级会大于或。

```shell
// 表示大于等于0.90并且小于3.0的版本
"ruflin/elastica": ">=0.90 <3.0",
```

4.波浪符 ~
该操作符限制最小版本号。

```php
允许表达式中的最后一位版本号达到最大值
```

如~1.2 与>=1.2 <2.0 相等,~1.5.6 与>=1.5.6 < 1.6.0 相等。也就是主版本号与次版本号保持不变，修复版本号可以达到最大值。 5.折音符 ^
该操作符约束锁定最大版本号。

```php
锁定表达不变的是第一位主版本号，允许升级版本到安全的版本号
```

如^1.2 就等于>=1.2 <2.0，^1.2.3 就等于>=1.2.3 < 2.0.0。

## 语义化

什么是语义化呢？说的简单一点即是版本号管理。我们的包一般分文如下格式组成:
`php 主版本号+次版本号+修复版本`
如上面的例子中，我们都提到了一些包的版本号是 x.x.x。第一位就是主版本号，第二位就是次版本号，第三位就是针对一些 bug 修复来的修复版本号。具体的可以参考https://semver.org/lang/zh-CN/

## Composer 使用优化

1.composer 加载类型
composer 加载类型包括 classmap,psr-0,psr-4,file.psr-0 逐渐的被抛弃了，由于一些老项目还在使用该规则，因此部分项目仍在使用。大多数的是使用 psr-4。classmap 是包文件的映射处理，下面有讲。file 主要加载一些 helper 的操作。

1.composer dump-autoload -o
该命令会根据包的命令空间和路径生成文件映射,当去加载包的时候，会根据映射去加载包文件。这样会加快我们的包文件访问速度。当我们执行了该命令，可以查看如下如的界面。被圈出来的就是类映射配置。
![](http://qiniucloud.qqdeveloper.com/%E5%B1%8F%E5%B9%95%E5%BF%AB%E7%85%A7%202019-09-08%2016.53.11.png)
composer 具体怎么去处理这个加载顺序逻辑，我们可以通过查看 composer 加载类的处理顺序。下图中圈出的方法，首先就是去加载 classmap，没找到在去加载 psr-4。
![](http://qiniucloud.qqdeveloper.com/%E5%B1%8F%E5%B9%95%E5%BF%AB%E7%85%A7%202019-09-08%2016.58.41.png)

2.composer dump-autoload -a
该命令主要的是功能是，当在我们 1 中执行了命令，会生成映射文件。如果当去加载映射文件没有找到时，则提示包文件不存在。