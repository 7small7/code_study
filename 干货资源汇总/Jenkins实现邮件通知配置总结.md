[TOC]
## 文章简介

本文分享Jenkins实现邮件发送，安装Jenkins可参考[jenkins安装](http://www.qqdeveloper.com/2019/10/06/jenkins-install/)。写这篇文章，是在实际使用Jenkins过程中遇到这样一个问题，当每次Jenkins构建成功或者失败后，需要个人登录Jenkins查看构建结果，同时在构建前做了数据备份，也需要手动的拷贝一份备份文件到本地。为了解决这个问题，便想到了Jenkins的邮件功能。在个人实践中，在每次master分支自动构建前，需要将数据库和代码打包、备份，在Jenkins构建结束之后，将备份的文件发送给对应的负责人。

> 文章部分细节的地方可能省略待过，因此需要对Jenkins有一定使用的用户比较合适，如果不熟悉的，可以参考上面提交的文章。文章涉及到不准确的信息还望反馈。

## 大致逻辑

![](https://oscimg.oschina.net/oscnet/up-6c23bf7022e10814f44f0d376e8381ca02f.png)

接下来整个流程，也都按照该流程进行演示如何配置。

### 安装插件

首先点击管理，然后点击插件管理，跳转到插件中心。
![](https://oscimg.oschina.net/oscnet/up-c6926409d2ab2548829de388ebee93a850c.png)
![](https://oscimg.oschina.net/oscnet/up-84b952a2f57047200f96ae1ee7d62289278.png)
如果没有安装过email对应的插件，点击可选插件按钮，然后输入email关键词，进行搜索即可。这里我已经安装过了，为了演示选择已安装菜单。
![](https://oscimg.oschina.net/oscnet/up-17a7175156143690e7e64877b7b944ef498.png)
> 在安装插件的过程中，需要注意一个版本号。2.73版本的插件有一个bug，就是配置正确却不能发送邮件。建议避开这个版本号。

## 系统配置

安装好插件之后，接下来就需要进入系统配置。进入系统配置，主要配置两个地方，一个是管理员的邮箱地址，一个是插件的配置信息。
![](https://oscimg.oschina.net/oscnet/up-b2869ef4e3fd658286392373fc9dd66a477.png)
![](https://oscimg.oschina.net/oscnet/up-190e26896b41a685bdeae71da2433bf5004.png)
![](https://oscimg.oschina.net/oscnet/up-3b0adc0017f6f1cc657872a3f9c1f98d0dc.png)
![](https://oscimg.oschina.net/oscnet/up-5e8f0e815813df6f83b01b7cff3a71f1e75.png)
> 记住插件发送邮件的账号一定的和管理员的邮箱账号一致，否者会出现下面错误信息。

错误信息:
```shell
com.sun.mail.smtp.SMTPSenderFailedException: 553 Mail from must equal authorized user
at com.sun.mail.smtp.SMTPTransport.mailFrom(SMTPTransport.java:1587)
Caused: com.sun.mail.smtp.SMTPSendFailedException: 553 Mail from must equal authorized user;
  nested exception is:
    com.sun.mail.smtp.SMTPSenderFailedException: 553 Mail from must equal authorized user
    at com.sun.mail.smtp.SMTPTransport.issueSendCommand(SMTPTransport.java:2057)
    at com.sun.mail.smtp.SMTPTransport.mailFrom(SMTPTransport.java:1580)
    at com.sun.mail.smtp.SMTPTransport.sendMessage(SMTPTransport.java:1097)
    at javax.mail.Transport.send0(Transport.java:195)
    at javax.mail.Transport.send(Transport.java:124)
    at hudson.tasks.Mailer$DescriptorImpl.doSendTestMail(Mailer.java:581)
    at java.lang.invoke.MethodHandle.invokeWithArguments(MethodHandle.java:627)
```
错误信息的大致意思就是说，邮箱的授权用户(插件配置的账号)和from(邮件发送者)账号不一致。
> 这里其实有点小困惑，邮件配置授权账户，为什么管理员的邮件账号还必须保持一致。按理来说，既然配置了授权用户，就采用配置的授权邮箱进行发送呗。可能是系统使用的管理员邮件发送。这里的配置授权仅仅是为了授权第三方邮箱账户吧。
![](https://oscimg.oschina.net/oscnet/up-84fc1dec6d8c55679bc4d50bc7345121baf.png)

## 项目配置

接下来，我们创建一个任务测试发送邮件。至于具体的配置这里省略待过，直接记录配置邮件的地方。
![](https://oscimg.oschina.net/oscnet/up-4c73b6965dc8e334fcd1aad7efc0e53c76a.png)
![](https://oscimg.oschina.net/oscnet/up-e073d94bde8efece4b98ac3e7dacd954943.png)
这里面就是针对该项目的具体配置信息，上面我们提交到系统配置，属于全局配置。如果这里没有配置，则默认走全局配置。大致的配置信息和全局配置都是相同的作用，可以根据不同的任务，自行配置即可。
![](https://oscimg.oschina.net/oscnet/up-ec6d8e67198087ffea504373fa2eb0b0b73.png)
![](https://oscimg.oschina.net/oscnet/up-3829348338a5f6aaab5c8aa75a8389025cb.png)![](https://oscimg.oschina.net/oscnet/up-c18d04d80cf59964078ce64659461710eb3.png)

## 邮件测试

最后配置完毕，当提交代码待仓库后，使用webhooks自动触发构建，登录邮箱就可以查看到下面的一些构建基础信息了。如果邮件需要更多的配置信息，直接到任务中去配置即可。
![](https://oscimg.oschina.net/oscnet/up-1af70a392b7556981a9263584c1bbd73b53.png)
​