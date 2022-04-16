## 命令参数
命名参数可以让函数或者方法的调用更加清晰直观，对于如下的函数定义:
```php
function foo(string $a, string $b, ?string $c = null, ?string $d = null) {
}
```
PHP8中你可以通过下面的方式传入参数进行调用。
```php
foo(
    b: 'value b', 
    a: 'value a', 
    d: 'value d',
) {}
```