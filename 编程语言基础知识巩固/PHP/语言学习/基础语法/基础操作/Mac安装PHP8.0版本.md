[TOC]

## 下载源码

```shell
wget https://www.php.net/distributions/php-8.0.0.tar.gz
tar zxvf php-8.0.0.tar.gz
cd php-8.0.0
```
## 安装 PHP
```shell
# 生成 configure 文件
./buildconf --force
# 配置构建流程（最小化安装）
./configure --prefix=/usr/local/php80 \
--with-config-file-path=/usr/local/php80 \
--enable-cli \
--without-iconv
# 构建 && 安装
make && sudo make install
# 设置配置文件
sudo cp php.ini-development /usr/local/php80/php.ini
```
## 解决问题

### 报错信息
在执行<kbd>make</kbd>命令时，可能会出现如下错误信息。
```shell
configure: error: The pkg-config script could not be found or is too old.  Make sure it
is in your PATH or set the PKG_CONFIG environment variable to the full
path to pkg-config.

Alternatively, you may set the environment variables LIBXML_CFLAGS
and LIBXML_LIBS to avoid the need to call pkg-config.
See the pkg-config man page for more details.
```
### 解决方式
使用下面的命令，安装好之后，重新执行<kbd>make</kbd>命令即可。
```shell
brew install pkg-config
```

## 设置环境变量

由于 Mac 下默认自带 PHP 环境，这里修改默认的 PHP 版本。
```shell
sudo vim ~/.zshrc
```
在文件底部添加如下配置信息。
```shell
alias php="/usr/local/php80/bin/php"
```