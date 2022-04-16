## 文章说明

一直知道opcache可以提高PHP性能，但没有具体的关注，更多的利用其他的组件来提升系统的性能。一次无意开启了opcache之后，并随意设置了一些配置。结果导致后面在使用一个项目时，发现`项目总是不会读取到最新的代码，而是隔一段时间才会执行到最新代码`。排查了很久才想起来开启了opcache，于是对opcache做了一个简单的学习与总结。
> 发现这个优化小技巧之后，后面也会对稍微底层进行探索学习，欢迎大家持续关注该文。

## 什么是opcache

OPcache 通过将 PHP 脚本预编译的字节码存储到`共享内存`中来提升 PHP 的性能， 存储预编译字节码的好处就是 省去了每次加载和解析 PHP 脚本的开销。

## opcache运行原理

### 不使用opcache

在使用opcache之前，我们事先看一个request，PHP的一个大致处理流程是如何的。大致的示意图如下：
![Snipaste_2021-09-24_23-18-32](https://gitee.com/bruce_qiq/picture/raw/master/2021-9-24/1632496909208-Snipaste_2021-09-24_23-18-32.png)

1. 首先会去模块初始化一次，也就是加载我们php.ini当中的一些配置信息，这里需要根据配置信息初始化一次。

2. 初始化完php.ini的配置信息之后，第二步就是针对当前请求的信息做一次初始化。例如我的一些get、post以及$_SEVER等相关的信息。

3. 得到上面1和2中的信息之后，则时候就会去真正执行我们的php脚本文件内容了，也就是我们写的代码。是怎么去实现的呢？如下图：
![1128628-20180504142714761-711951956](https://gitee.com/bruce_qiq/picture/raw/master/2021-9-25/1632543136562-1128628-20180504142714761-711951956.png)
Zend引擎读取.php文件-->扫描其词典和表达式 -->解析文件-->创建要执行的计算机代码(称为Opcode)-->最后执行Opcode--> response 返回。

4. 执行完php脚本文件内容之后，这时候会针对1和2中的一些初始化信息，进行销毁。

### 使用opcache

当使用opcache之后，当一个请求来了之后，依然的会去执行上面提到的1和2，进行模块和请求的初始化。接着就会去编译php脚本文件内容，opcache也是在这一个阶段才会产生作用。

通过上面的第3步，我们可以看到每一次请求都会去解析php文件内容，不管是php文件的内容是否发生变化，都会执行这样的一个重复流程来生成opcode。

opcache的作用就是减少每次请求都会去编译php脚本文件，第一次将编译好的脚本文件内容缓存起来，下一次请求就不需要去重复编译了，而是直接冲内存中取就行了。减少了CPU和内存的消耗。

![Snipaste_2021-09-24_23-18-16](https://gitee.com/bruce_qiq/picture/raw/master/2021-9-24/1632496899073-Snipaste_2021-09-24_23-18-16.png)

1. 首先会去模块初始化一次，也就是加载我们php.ini当中的一些配置信息，这里需要根据配置信息初始化一次。

2. 初始化完php.ini的配置信息之后，第二步就是针对当前请求的信息做一次初始化。例如我的一些get、post以及$_SEVER等相关的信息。

3. 此时去解析php脚本文件，首先会去判断opcode是否存在，如果不存在就执行一个编译流程并缓存到共享内存中。当存在opcode时，则直接使用共享内存中的opcode，不会再进行一次编译的过程。
![1128628-20180504142702126-1584014725](https://gitee.com/bruce_qiq/picture/raw/master/2021-9-25/1632543119211-1128628-20180504142702126-1584014725.png)

4. 执行完php脚本文件内容之后，这时候会针对1和2中的一些初始化信息，进行销毁。

### 使用总结

1. 通过上面的对比，很容易看得出来opcache执行的时段在于编译php脚本文件，减少了编译的过程。 

2. 对于模块初始化、请求初始化等这样的一个重复流程，该如何优化。这里可以去了解一下swoole。

3. 可能会存在这样一个疑问，opcode给缓存起来了，如果我们更新了代码，这时候还是会加载旧的opcode还是重新编译一次opcode并缓存起来呢？后面我们会单独总结。

## opcache配置说明
```php
[opcache]
; 是否快开启opcache缓存。
;opcache.enable=1

; 是否在cli模式下开启opcache。
;opcache.enable_cli=1

; opcache共享内存的大小(单位是M)。
;opcache.memory_consumption=128

; 预留字符串的的内存大小(单位是M)。
;opcache.interned_strings_buffer=8

; 在hash表中存储的最大脚本文件数量，范围是200到1000000之间。实际的情况是在{ 223, 463, 983, 1979, 3907, 7963, 16229, 32531, 65407, 130987 }中找到第一个大于等于设置值的质数。最小范围是200。
;opcache.max_accelerated_files=10000

; 浪费内存的上线，如果超过这个上线，opcache将重新启动。
;opcache.max_wasted_percentage=5

; 如果启用，opcache将会在hash表的脚本键后面增加一个文件目录，避免吃同名的脚本产生冲突。禁用的话可以提高性能，但是也容易导致应用不可用。
;opcache.use_cwd=1

; 如果启用(1)，opcache会每隔设置的值时间来判断脚本是否更新。如果禁用(0)，则不会自动检测脚本更新，必须通过重启PHP服务，或者使用opcache_reset()、opcache_invalidate()函数来刷新缓存。
;opcache.validate_timestamps=1

; opcache检查脚本是否更新的时间周期(单位是秒)，如果设置为0则会针对每一个请求进行检查更新，如果validate_timestamps=0，该值不会生效。
;opcache.revalidate_freq=60

; 如果禁用，在统一include_path下面已经缓存的文件将被重用，因此无法找到该路径下的同名文件。
;opcache.revalidate_path=0

; 是否保存PHP脚本中的注释内容。禁用，则不会缓存PHP代码中的注释，可以减少文件中的体积，但是一些依赖注释或者注解将无法使用。
;opcache.save_comments=1

; 如果启用，则会使用快速停止续发事件。 所谓快速停止续发事件是指依赖 Zend 引擎的内存管理模块 一次释放全部请求变量的内存，而不是依次释放每一个已分配的内存块。
; 在php7.2.0开始，被移除，这类说的事件将会在PHP中自动处理。
;opcache.fast_shutdown=1

; 如果启用，在调用file_exists()、is_file()和is_readable()函数时，不管文件是否被缓存，都会检测操作码。如果禁用，可能读取的内容是一些旧数据。
;opcache.enable_file_override=0

; 控制优化级别，是一个二进制的位的掩码。
;opcache.optimization_level=0xffffffff

; 不进行编译优化的配置文件路径。该文件中配置具体哪些不被编译的文件。如果文中每行的开头是";"开头，则会被视为注释。黑名单中的文件名，可以是通配符，也可以使用前缀。
; 例如配置文件的路径是"/home/blacklist.txt"，则该配置的值就是该路径。
; 配置的内容可以是如下格式

; 这是一段注释，在解析的时候因为开头是;，则会被视为注释
;/var/www/a.php
;/var/www/a/b.php

;opcache.blacklist_filename=

; 以字节为单位的缓存的文件大小上限。设置为 0 表示缓存全部文件。
;opcache.max_file_size=0

; 每个N次请求会检查缓存校验和，0是不检查。该项对性能有较大影响，尽量在调试环境中使用。
;opcache.consistency_checks=0

; 如果缓存处于非激活状态，等待多少秒之后计划重启。 如果超出了设定时间，则 OPcache 模块将杀除持有缓存锁的进程， 并进行重启。
;opcache.force_restart_timeout=180

; 错误日志文件位置，不填写将默认输出到服务器的错误日志文件中。
;opcache.error_log=

; 错误日志文件等级。
; 默认情况下，仅有致命级别（0）及错误级别（1）的日志会被记录。 其他可用的级别有：警告（2），信息（3）和调试（4）。
; 如何设置的是1以上，在进行force_restart_timeout选项时，会将错误日志中插入一条警告信息。
;opcache.log_verbosity_level=1

; opcache首选的内存模块，不配置则自动选择。可以选择的值有mmap，shm, posix 以及 win32。
;opcache.preferred_memory_model=

; 保护共享内存，以避免执行脚本时发生非预期的写入。 仅用于内部调试。
;opcache.protect_memory=0

; 只允许指定字符串开头的PHP脚本调用opcache api函数，默认不做限制。
;opcache.restrict_api=

; 在 Windows 平台上共享内存段的基地址。 所有的 PHP 进程都将共享内存映射到同样的地址空间。 使用此配置指令避免“无法重新附加到基地址”的错误。
;opcache.mmap_base=

; 配置二级缓存目录并启用二级缓存。 启用二级缓存可以在 SHM 内存满了、服务器重启或者重置 SHM 的时候提高性能。 默认值为空字符串 ""，表示禁用基于文件的缓存。
;opcache.file_cache=

; 启用或禁用在共享内存中的 opcode 缓存。
;opcache.file_cache_only=0

; 当从文件缓存中加载脚本的时候，是否对文件的校验和进行验证。
;opcache.file_cache_consistency_checks=1

; 在 Windows 平台上，当一个进程无法附加到共享内存的时候， 使用基于文件的缓存。需要开启opcache.file_cache_only选项。建议开启此选项，否则可能导致进程无法启动。
;opcache.file_cache_fallback=1

; 启用或者禁用将 PHP 代码（文本段）拷贝到 HUGE PAGES 中。 此项配置指令可以提高性能，但是需要在 OS 层面进行对应的配置。
;opcache.huge_code_pages=1

; 针对当前用户，验证缓存文件的访问权限。
;opcache.validate_permission=0

; 在 chroot 的环境中避免命名冲突。 为了防止进程访问到 chroot 环境之外的文件，应该在 chroot 的情况下启用这个选项。
;opcache.validate_root=0
```

## 配置示例

下面这一段代码是PHP官方给的一个示例配置，推荐使用该配置项进行配置，也可以根据自己实际的情况进行单独配置。
```php
;opcache.memory_consumption=128
;opcache.interned_strings_buffer=8
;opcache.max_accelerated_files=4000
;opcache.revalidate_freq=60
;opcache.fast_shutdown=1
;opcache.enable_cli=1
```

## 问题总结

1. 如何更新opcode?

编译好的opcode会添加到共享内存中，如果我们更新了代码就需要去更新opcode，否则得到的代码还是旧的opcode。就会发生文章开头说到的情况。要解决这个问题，我们有几种方式。
```php
; 方法一
直接重启我们的php进程，但这样会导致服务中断，是一种不推荐的方式。

; 方法二
根据官方给出的函数，进行设置。在代码中使用opcache_reset()或者使用opcache_invalidate()函数进行充值opcode。直接通过一个特殊的链接去执行这个函数即可。

; 方法三
使用php.ini中的配置项实现自动充值opcode。
opcache.validate_timestamps = 1
opcache.revalidate_freq  = 60
```

## 效果演示

![Snipaste_2021-09-25_12-42-05](https://gitee.com/bruce_qiq/picture/raw/master/2021-9-25/1632544942250-Snipaste_2021-09-25_12-42-05.png)

![Snipaste_2021-09-25_12-41-15](https://gitee.com/bruce_qiq/picture/raw/master/2021-9-25/1632544950016-Snipaste_2021-09-25_12-41-15.png)

1. 上面的两张图，第一张是未开启opcache的一个压测，第二个是开启opcache的一个压测。

2. 从截图上来看，开启opcache开启之后，有一些小幅度的提升。也并没有网上说的翻倍的提升。

3. 这里的提升不能说opcache的提升效果不明显，这需要根据综合因素决定，这里的演示使用Mac操作本身就会降低很多。