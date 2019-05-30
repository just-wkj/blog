# 依赖注入

### DI 和IoC是什么
DI  全称: Dependency Injection  依赖注入  
IoC 全称: Inversion of Control  控制反转

控制反转是一种设计模式,依赖注入是该设计模式的一种实现.

IoC描述大致如下:
> 高层模块不应该依赖底层模块，他们都应该依赖抽象。抽象不应该依赖于细节，细节应该依赖于抽象。


真实的dependency injection container会提供更多的特性，如
- 自动绑定（Autowiring）或 自动解析（Automatic Resolution）
- 注释解析器（Annotations）
- 延迟注入（Lazy injection）

下面是简单的自动绑定的容易实现代码:
```php
<?php

class A{
    public function doSomeThing(){
        echo __METHOD__ . "|";
    }
}

class B{
    public $a;

    public function __construct(A $a){
        $this->a = $a;
    }

    public function doSomeThing(){
        $this->a->doSomeThing();
        echo __METHOD__ . "|";
    }
}

class C{
    public $b;

    public function __construct(B $b){
        $this->b = $b;
    }

    public function doSomeThing(){
        $this->b->doSomeThing();
        echo __METHOD__ . "|";
    }
}

class Container{
    public $registry = [];

    public function __set($name, $value){
        $this->registry[$name] = $value;
    }

    public function __get($name){
        //return $this->registry[$name]($this);
        if (!isset($this->registry[$name])) {
            throw new Exception('请先设置容器参数 ' . $name);
        }

        return $this->build($this->registry[$name]);
    }

    private function build($className){
        //如果是闭包,直接调用即可
        if ($className instanceof Closure) {
            return $className($this);
        }

        //反射class
        $reflect = new ReflectionClass($className);

        //检查是否能被实例化,排除abstract 和interface
        if (!$reflect->isInstantiable()) {
            throw new Exception($className . " 不能被实例化");
        }
        //获取构造函数
        $constructor = $reflect->getConstructor();

        if (is_null($constructor)) {//无构造函数
            return new $className;
            //或
            //return $reflect->newInstance();
        }

        //获取构造函数的参数
        $params = $constructor->getParameters();
        //获取参数所需的依赖
        $dependencies = $this->getDependencies($params);

        //实例化class,并添加构造函数参数
        return $reflect->newInstanceArgs($dependencies);
    }

    //获取构造函数参数依赖
    private function getDependencies(array $params){
        $dependencies = [];
        foreach ($params as $param) {
            $dependency = $param->getClass();
            if (is_null($dependency)) {//非class 参数
                $dependencies[] = $this->resolveNonClass($param);
            } else {//构造函数依赖是class, 递归的构建class的依赖
                $dependencies[] = $this->build($dependency->name);
            }
        }

        return $dependencies;
    }

    private function resolveNonClass($param){
        // 有默认值则返回默认值
        if ($param->isDefaultValueAvailable()) {
            return $param->getDefaultValue();
        }

        throw new Exception('缺少构造函数参数');
    }
}

//常规调用class C 中的 doSomeThing 方法
echo PHP_EOL . '===old==' . PHP_EOL;
$old = new C(new B(new A()));
$old->doSomeThing();

//依靠依赖注入 调用 class C 中的 doSomeThing 方法
echo PHP_EOL . '===new==' . PHP_EOL;
$di = new Container();//实例化容器
$di->c = C::class;//class C 放在容器中
$di->c->doSomething();//调用  C 中的 doSomeThing 方法, 此时C需要实例化B,B需要实例化A,这些操作都在容器中实现

//依靠依赖注入 调用 class B 中的 doSomeThing 方法
echo PHP_EOL . '===new2==' . PHP_EOL;
$di = new Container();
$di->b = B::class;
$di->b->doSomething();

```

//TODO 有时间看下TP还有laravel是如何实现的


参考文章:
- [PHP程序员如何理解依赖注入容器(dependency injection container)](https://segmentfault.com/a/1190000002424023)  
- [聊一聊PHP的依赖注入(DI) 和 控制反转(IoC)](https://segmentfault.com/a/1190000007209266)
