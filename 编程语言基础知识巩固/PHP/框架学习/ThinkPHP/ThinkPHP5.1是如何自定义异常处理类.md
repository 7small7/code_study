[TOC]

ThinkPHP5 之后封装了系统的异常类操作，同时我们也可以在 config 目录下面的 app.php 配置文件中设置我们自定义的异常处理操作.配置项如下

```php
// 异常处理handle类 留空使用 \think\exception\Handle
'exception_handle' => '',
```

要实现自定义，其实实现原理很简单，我们可以把系统默认的异常类当做父类，我们自身创建的类当做子类，子类去集成父类并重写父类的方法，这样就可以实现自定义异常类了。通过查看系统异常类，可以发现只要是通过如下方法实现(下面的方法为\think\exception\Handle 类的 render 方法)。

```php
 public function render(Exception $e)
    {
        if ($this->render && $this->render instanceof \Closure) {
            $result = call_user_func_array($this->render, [$e]);

            if ($result) {
                return $result;
            }
        }

        if ($e instanceof HttpException) {
            return $this->renderHttpException($e);
        } else {
            return $this->convertExceptionToResponse($e);
        }
    }
```

当我们框架报告错误时，也是通过该方法实现错误渲染的。我们只要对下面这个方法进行重写，报错内容的格式按照我们自身的需求来写，这样就可以实现自定义了。

```php
namespace app\common\exception;

use Exception;
use think\exception\Handle;

class CommonException extends Handle
{
    /**
     * 公共异常处理函数
     * @param Exception $e
     * @return \think\Response|void
     */
    public function render(Exception $e)
    {
        // error方法为项目中创建的公共方法，在使用中需要修改为你自身的方法
        return error('在' . $e->getFile() . '的第' . $e->getLine() . '行,发生错误为' . $e->getMessage(), []);
    }
}
```

修改配置文件(config/app.php)

```php
// 异常处理handle类 留空使用 \think\exception\Handle
'exception_handle' => 'app\common\exception\CommonException',
```

错误报告对比
图一为系统默认异常界面，图二为自定义异常界面。
![屏幕快照 2019-06-23 16.40.48](http://qiniucloud.qqdeveloper.com/mweb/%E5%B1%8F%E5%B9%95%E5%BF%AB%E7%85%A7%202019-06-23%2016.40.48.png)

![屏幕快照 2019-06-23 16.40.30](http://qiniucloud.qqdeveloper.com/mweb/%E5%B1%8F%E5%B9%95%E5%BF%AB%E7%85%A7%202019-06-23%2016.40.30.png)