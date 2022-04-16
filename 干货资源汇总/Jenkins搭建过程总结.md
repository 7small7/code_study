[TOC]
## 文章简介
本文将介绍在Linux、Mac和Windows三种系统上安装jenkins，分别采用源码、docker以及war文件三种方式进行安装。每种安装方式都可以换到其他的系统上面去安装。

## windows安装

首先windows需要配置Java环境，这里不介绍了。
**a.** 下载war包，这个包只是在Java源码的基础打包后的一个包。
http://mirrors.jenkins.io/war-stable/latest/jenkins.war，该包是最新的包。

**b.** 运行war包。将下载的war包放在你自定义的目录。运行如下的命令。
```shell
java --jar jenkins.war
```
![windows-run-jenkins](http://qiniucloud.qqdeveloper.com/mweb/windows-run-jenkins.png)

**c.** 此时浏览器访问http://127.0.0.1:8080根据界面操作即可。
![chrome-inviate-jenkins](http://qiniucloud.qqdeveloper.com/mweb/chrome-inviate-jenkins.png)


## Mac安装
Mac安装我们采用的是docker安装，直接下载安装包点击下一步即可，这里不单独介绍。
**a.** 检验docker是否安装成功，使用如下命令，如出现下图的界面则表示安装成功。
```shell
docker -h
```
![docker-command](http://qiniucloud.qqdeveloper.com/mweb/docker-command.png)


**b.** docker安装，需要使用镜像，我们可以自己编写一个镜像也可以找开源的镜像。这里推荐jenkinsci/blueocean镜像，一直在使用，没什么问题。这里是将本地的8080端口映射到容器的8080端口，因为jenkins默认的端口是8080。第一个8080是主机的端口号，第二个8080是容器的端口号，也就是jenkins容器的端口号。我们可以根据实际情况将第一个端口号改为主机上面没有使用的端口号。
```shell
docker run -p 8080:8080 jenkinsci/blueocean
```
> enkinsci/blueocean每次发布Blue Ocean新版本时，都会发布新镜像。您可以在标签 page页上看到以前发布的镜像版本列表 。
您还可以使用其他Jenkins Docker镜像（在Docker Hub上可通过jenkins/jenkins获取）。 但是，这些不会随Blue Ocean的发布而提供，需要通过 Jenkins中的Manage Jenkins > Manage Plugins页面进行安装。 在Blue Ocean入门中了解更多信息。

输入上面的安装命令之后，docker会自动去拉取镜像，我们不需要单独操作，稍等一会就好了。
![docker-install-image](http://qiniucloud.qqdeveloper.com/mweb/docker-install-image.png)


## Linux安装

### 源码安装

Linux使用的是centos7的系统，我们采用源码的方式进行安装，同样的也需要安装Java环境，建议安装的版本不低于1.8。安装的具体操作不介绍，可以参考http://www.qqdeveloper.com/2019/08/29/java-centos-path/
**a.** 安装好Java之后，下载jenkins源码包。分别执行下面4行命令。
```shell
wget -O /etc/yum.repos.d/jenkins.repo http://pkg.jenkins-ci.org/redhat/jenkins.repo
yum -y update nss
rpm --import https://jenkins-ci.org/redhat/jenkins-ci.org.key
yum install -y jenkins
```
**b.** 如何管理jenkins服务。根据上面a中安装的操作，我们已经将jenkins安装到系统服务中了。便可以使用系统服务管理的命令进行管理。
```shell
systemctl start jenkins// 启动服务
systemctl stop jenkins// 停止服务
systemctl restart jenkins// 重启服务
systemctl status jenkins// 查看服务状态
```
> centos7的system改为了systemctl

### 依赖于Tomcat进行安装

第一步：下载Tomcat到/opt/目录下面，并解压。

第二步：到Jenkins官网下载jenkins.war文件包，放在/opt/apache-tomcat-8.5.40/webapps

第三步：启动Tomcat,<kbd>/opt/apache-tomcat-8.5.40/bin/startup.sh</kbd>

执行上诉操作之后，浏览器访问ip:8080/jenkins,根据提示操作即可。

### 直接使用war包安装

将war下载到指定目录，然后执行如下命令。该命令的含义是启动一个端口为80的Jenkins服务，并以守护进程的方式运行。
```shell
nohup java -jar jenkins.war  --ajp13Port=-1 --httpPort=80 >/dev/null 2>&1 &
```