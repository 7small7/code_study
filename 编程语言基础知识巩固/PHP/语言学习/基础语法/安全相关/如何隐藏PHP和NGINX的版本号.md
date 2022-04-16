[TOC]

## 文章简介

一般的网站，在发送请求信息时，在HTTP的response头信息中都会包含一个web服务器的信息。如下示例:
```shell
HTTP/1.1 200 OK
Server: nginx/0.8.31
Date: Aug, 13 Jan 2010 08:17:30 GMT
Content-Type: text/html
Content-Length: 2341
Last-Modified: Mon, 11 Jan 2010 08:45:11 GMT
Connection: keep-alive
Keep-Alive: timeout=15
Accept-Ranges: bytes
X-Powered-By: PHP/5.2.4
```
为了网站安全，我们一般建议将NGINX和PHP的版本信息给隐藏了。那该如何做呢？

## 隐藏版本信息

### 修改php配置

找到php.ini文件，添加或者修改下面的配置项
```php
// 将该项配置值On改为Off即可，如果不存在添加该项。
server_tokens off;
```
### 修改php-fpm配置

找到php-fpm配置文件。
将下面的配置项进行修改
```php
// 修改前
fastcgi_param SERVER_SOFTWARE nginx/$nginx_version;
// 修改后
fastcgi_param SERVER_SOFTWARE nginx;
```

## 隐藏PHP版本信息

隐藏PHP的版本信息，直接修改php.ini配置文件即可。
```php
// 将该项配置值On改为Off即可，如果不存在添加该项。
expose_php Off
```

## 效果演示

配置好之后，我们重启NGINX和PHP服务。对之前的uri进行重新访问，就会发现PHP版本信息不存在了。
```javascript
HTTP/1.1 200 OK
Server: nginx/0.8.31
Date: Aug, 13 Jan 2010 08:17:30 GMT
Content-Type: text/html
Content-Length: 2341
Last-Modified: Mon, 11 Jan 2010 08:45:11 GMT
Connection: keep-alive
Keep-Alive: timeout=15
Accept-Ranges: bytes
```