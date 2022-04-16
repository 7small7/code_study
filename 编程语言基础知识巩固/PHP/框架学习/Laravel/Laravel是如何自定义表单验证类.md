[TOC]

## 问题背景

最近在公司的项目开发中使用到了 laravel 框架，采用的是前后端开发的模式。接触过前后端开发模式的小伙伴应该都知道，后端返回的数据格式需要尽可能搞得保证一致性，这样前端在处理时也方便处理。我们先通过观看下面的两张接口返回的效果图吧，这样或许会更加的直观一些。
![](http://qiniucloud.qqdeveloper.com/laravel-from-2.png)

<center>laravel默认的输出格式(图一)</center>

![](http://qiniucloud.qqdeveloper.com/laravel-from-1.png)

<center>修改后的输出格式(图二)</center>
或许通过上面两张图，你还是未看出有什么区别的话。这里我用文字描述一下吧。
这种情况是发生在laravel做表单验证的情况下发生的。前端向我后端接口发送一个POST请求时，发送了一个title和body的字段。我后端需要对两个字段做一些非空验证。按照框架手册来进行的话，输出的格式就是图一的格式。然后后端统一的输出格式是图二中的格式，如果按照图一的格式输出肯定是不行，这样就需要我们做一个特殊处理。

## 问题排查

首先我们可以通过文档参看到如下信息。下面划线的部分，提到的返回信息是将<font color='red'>所有未验证通过</font>的数据都返回给前端，就如图一中的数据格式。
![](http://qiniucloud.qqdeveloper.com/laravel-from-3.png)

<center>laravel默认的输出格式(图三)</center>

## 预期效果

通过图三我们知道了 laravel 默认的是返回一个带 422 的 http 状态码并且将所有的验证错误信息都返回。
然而我们需要的只是如图二的格式，单个的输出错误信息。大致的解决思路就是在输出的时候，我们去默认显示第一个未通过的验证信息，当通过之后，之前第二个未严重通过的就变成了第一个，这样依次循环下去，我们的每个数据就得到了验证。验证的地方我们选择框架异常统一处理的地方，这样每次验证都自动的进行处理。

## 解决方案

> 该框架是 laravel5.8 的情况下进行编写,如果版本不同,或许还需要特殊的处理,不过处理的思路可以参考下面的。

1.创建一个表单验证器。执行下面的命令之后，我们在`php app/Http/Requests`目录下面就可以看到该类文件了。

```shell
php artisan make:request ProjectValidate
```

2.定义验证规则。rules 方法是定义验证规则,而 messages 方法则是定义返回的错误信息，该方法也可以省略掉，这样提示的信息就是英文而不是图一或图二看到的中文了。

```php
namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class ProjectValidate extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     *
     * @return bool
     */
    public function authorize()
    {
        return true;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array
     */
    public function rules()
    {
        return [
            'title' => 'bail|required',
            'body' => 'required',
        ];
    }

    /**
     * define the validation message
     *
     * @return array
     */
    public function messages()
    {
        return [
            'title.required' => '文章标题必填',
            'body.required' => '文章内容必填',
        ];
    }
}
```

3.使用验证器。这里我定义了一个控制器，在 addData 方法中，使用依赖注入的方式去实现数据的验证。<font color='red'>记住，该方法体在未通过数据验证的情况下是不会去执行的。</font>

```php
namespace App\Http\Controllers\Backend\Project;

use App\Http\Requests\ProjectValidate;
use App\Http\Controllers\Backend\UCenter;

class Index extends UCenter
{

    public function index()
    {
        return success();
    }

    public function addData(ProjectValidate $request)
    {
        $validated = $request->validated();
        return success($validated);
    }
}
```

4.统一处理数据格式。找到`php App\Exceptions\Handler.php`文件，找到下面的方法，修改为如下内容。这时候在做表单验证就会显示图二的格式信息了。

```php
public function render($request, Exception $exception)
    {
        if ($exception instanceof ValidationException) {
            // 方式一
            $message = $exception->validator->getMessageBag()->first()
            // 方式二
            // 只读取错误中的第一个错误信息
            $errors  = $exception->errors();
            $message = '';
            // 框架返回的是二维数组，因此需要去循环读取第一个数据
            foreach ($errors as $key => $val) {
                $keys    = array_key_first($val);
                $message = $val[$keys];
                break;
            }
            return response()->json(['code' => 1001, 'message' => $message, 'data' => []], 422);
        }
        return parent::render($request, $exception);
    }
```

## 总结

1.优势
输出固定的格式，前端在处理数据的时候，不需要做特别的格式上面调整。

2.劣势
这样的方式验证，每验证一次，就会向后端发送一个 http 请求。