[TOC]

## 简介

本文对使用hyperf框架的表单验证中遇到的两个小细节做一个分享。具体的两点如下：

1. 自定义验证异常数据返回格式。
`该问题主要在下面的第3点体现。`
2. 自定义验证规则。
`该问题主要在下面的第6点体现。`
## 自定义验证异常格式

1. 首选根据官方文档进行操作，安装验证组件。

```php
composer require hyperf/validation
php bin/hyperf.php vendor:publish hyperf/translation
php bin/hyperf.php vendor:publish hyperf/validation
```

1. 接着在配置文件`config/autoload/middlewares.php`，中添加验证异常中间件。这里的异常中间件为框架自带的异常处理中间件。

```php
<?php

declare(strict_types=1);
/**
 * This file is part of api.
 *
 * @link     https://www.qqdeveloper.io
 * @document https://www.qqdeveloper.wiki
 * @contact  2665274677@qq.com
 * @license  Apache2.0
 */
use Hyperf\Validation\Middleware\ValidationMiddleware;

return [
    'http' => [
        ValidationMiddleware::class,
    ],
];
```

1. 自定义一个验证异常处理器。`这一步是最重要的异步`，官方文档有提及到使用框架自带的异常处理器，如果你没有特别的需求，可以直接按照官方文档操作即可。由于我们的异常接口返回数据格式要返回一个`json`的格式，而不是默认的一个`文本`格式。

```php
<?php

declare(strict_types=1);
/**
 * This file is part of api.
 *
 * @link     https://www.qqdeveloper.io
 * @document https://www.qqdeveloper.wiki
 * @contact  2665274677@qq.com
 * @license  Apache2.0
 */
namespace App\Exception\Handler;

use Hyperf\ExceptionHandler\ExceptionHandler;
use Hyperf\HttpMessage\Stream\SwooleStream;
use Hyperf\Validation\ValidationException;
use Psr\Http\Message\ResponseInterface;
use Throwable;

/**
 * 自定义表单验证异常处理器.
 *
 * Class FromValidateExceptionHandler
 */
class FromValidateExceptionHandler extends ExceptionHandler
{
    public function handle(Throwable $throwable, ResponseInterface $response)
    {
        if ($throwable instanceof ValidationException) {
            // 格式化异常数据格式
            $data = json_encode([
                'code' => $throwable->getCode(),
                // 获取异常信息
                'message' => $throwable->validator->errors()->first(),
                'data' => [],
            ]);
            $this->stopPropagation();
            return $response->withStatus(422)->withBody(new SwooleStream($data));
        }

        return $response;
    }
    // 异常处理器处理该异常
    public function isValid(Throwable $throwable): bool
    {
        return true;
    }
}
```

1. 编写完验证异常处理器之后，将该异常添加到异常配置文件`config/autoload/exceptions.php`中。由于hyperf中异常处理器的配置顺序会影响到异常的处理顺序，这里可以随机顺序配置。

```php
<?php

declare(strict_types=1);
/**
 * This file is part of api.
 *
 * @link     https://www.qqdeveloper.io
 * @document https://www.qqdeveloper.wiki
 * @contact  2665274677@qq.com
 * @license  Apache2.0
 */
use App\Exception\Handler\FromValidateExceptionHandler;

return [
    'handler' => [
        'http' => [
            Hyperf\HttpServer\Exception\Handler\HttpExceptionHandler::class,
            App\Exception\Handler\AppExceptionHandler::class,
            // 自定义的验证异常处理器
            FromValidateExceptionHandler::class,
        ],
    ],
];
```

5. 剩下的代码就按照文档操作，编写一个`独立的验证类`文件，在对应的`控制器中`的`方法`采用`依赖注入`的方式调用即可。输出的结果，格式就和下面的一样了。
![Snipaste_2021-06-30_18-38-48](https://gitee.com/bruce_qiq/picture/raw/master/2021-6-30/1625049543644-Snipaste_2021-06-30_18-38-48.png)

## 自定义验证规则

为什么有自定义验证规则呢？无非就是官网提供的验证规则属于常见的，可能你会根据项目的需要，自定义一些规则，这时候就需要你单独定义一个规则了。`我们这里创建一个money的验证规则，验证金额是否合法。`

1. 创建一个监听器。
```php
<?php

declare(strict_types=1);
/**
 * This file is part of api.
 *
 * @link     https://www.qqdeveloper.io
 * @document https://www.qqdeveloper.wiki
 * @contact  2665274677@qq.com
 * @license  Apache2.0
 */
namespace App\Listener;

use Hyperf\Event\Contract\ListenerInterface;
use Hyperf\Validation\Contract\ValidatorFactoryInterface;
use Hyperf\Validation\Event\ValidatorFactoryResolved;

/**
 * 验证器监听器.
 *
 * Class ValidatorFactoryResolvedListener
 */
class ValidatorFactoryResolvedListener implements ListenerInterface
{
    public function listen(): array
    {
        return [
            ValidatorFactoryResolved::class,
        ];
    }

    public function process(object $event)
    {
        /** @var ValidatorFactoryInterface $validatorFactory */
        $validatorFactory = $event->validatorFactory;
        // 注册了 money 验证器
        $validatorFactory->extend('money', function ($attribute, $value, $parameters, $validator) {
            var_dump(time());
            $pregResult = preg_match('/^[0-9]{1,20}(\.[0-9]{1,2})?$/', $value);
            // true则返回错误信息;false则不返回错误，表示验证通过
            return empty($pregResult) ? true : false;
        });
        $validatorFactory->replacer('money', function ($message, $attribute, $rule, $parameters) {
            return str_replace(':money', $attribute, $message);
        });
    }
}
```

1. 注册监听器到`config/autoload/listeners`配置文件中。

```php
<?php

declare(strict_types=1);
/**
 * This file is part of api.
 *
 * @link     https://www.qqdeveloper.io
 * @document https://www.qqdeveloper.wiki
 * @contact  2665274677@qq.com
 * @license  Apache2.0
 */
use App\Listener\ValidatorFactoryResolvedListener;

return [
    ValidatorFactoryResolvedListener::class,
];
```

1. 自定义一个独立验证类文件。

```php
<?php

declare(strict_types=1);
/**
 * This file is part of api.
 *
 * @link     https://www.qqdeveloper.io
 * @document https://www.qqdeveloper.wiki
 * @contact  2665274677@qq.com
 * @license  Apache2.0
 */
namespace App\Request;

use Hyperf\Validation\Request\FormRequest;

class FooRequest extends FormRequest
{
    public function authorize(): bool
    {
        return true;
    }

    public function rules(): array
    {
        return [
            'money' => 'money',
        ];
    }
}
```

1. 自定义验证字段信息。找到`storage/languages/zh_CN/validation.php`文件。在下面添加如下两行代码，关于en文件下的验证字段配置信息，可以添加也可以不添加，根据实际需要添加即可。

```php
'money' => ':attribute格式错误',
'attributes' => [
    'money' => '金额',
],
```
5. 在对应的控制器中使用`依赖注入`的方式对独立的验证类文件进行注访问。这样我们的一个`独立验证规则`就可以配置好了。效果如下：
![Snipaste_2021-06-30_18-38-48](https://gitee.com/bruce_qiq/picture/raw/master/2021-6-30/1625049543644-Snipaste_2021-06-30_18-38-48.png)

6. 或许这么定义之后，发现自定义规则没有起作用，这种情况，获取是你没有传递该参数名导致的。只有你传递了参数名，该验证规则才会生效。
