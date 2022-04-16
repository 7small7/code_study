[TOC]

## 总结

1. $this表示对象的引用，写在类中的非静态方法中；

2. self和this不同，它指向类本身，不指向任何实例化对象；

3. static一般用来访问类中的静态变量和静态方法。

## $this使用

$this 表示对象的引用 $this写在类中的非静态方法中， 实例化该类，谁调用该方法（一般是对象调用）$this则表示该对象的引用。
```php
class Person {

    public $name;

    public function getName() {
    	var_dump($this);
        echo $this->name;
    }

}

$p = new Person();
$p2 = new Person();

$p->name = "小红";
$p2->name = "小明";

$p->getName();  // 小红
$p2->getName();  // 小明
/**
object(Person)#1 (1) {
  ["name"]=>
  string(6) "小红"
}
object(Person)#2 (1) {
  ["name"]=>
  string(6) "小明"
}
**/
```

## self使用

self 和 this 不同，它指向类本身，不指向任何实例化对象，一般用来访问类中的静态变量和静态方法,也是写在类中的方法。self写在哪个类中则表示该类的引用。
```php
class Person {

    public static $name = "小红";

    public static function getName() {
        echo self::$name;
    }

}

$p = new Person();

$p2 = new Person();

$p::getName();  // 小红

$p2::getName();  // 小红

$p::$name = "小明";

$p::getName();  // 小明

$p2::getName();  // 小明
```

## static使用

static关键字来定义静态方法和属性，也可用于定义静态变量以及后期静态绑定。也是写在类中的方法，也是那个类调用该方法static就表示那个类（绑定那个类）。
```php
class A {

    public function say() {
        echo "Hello";
    }

    public function saySelf() {
       //static 和 self 可以调用非静态方法  不能调用非静态属性  静态方法中不能有$this
      // static 和 self 可以写在非静态的方法中 ，可以使用对象调用
        self::say();

    }

    public function sayStatic() {
        static::say();// 在调用时，static指的就是B class
    }

}

class B extends A {

    public function say() {
        echo "World";
    }

}

$b = new B();

$b->say();  // World

$b->saySelf();  // Hello

$b->sayStatic();  // World
```