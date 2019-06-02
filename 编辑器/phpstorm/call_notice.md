# phpstorm 为静态方法添加提示

php开发的时候,会遇到魔术方法的使用,但是直接使用都是没有提示的,如何才能添加代码提示呢

比如:

![](http://img.justwkj.com/20190602145039.png)
直接使用静态方法调用不存在方法的时候,代码是没有提示的,
此时只需要在,当前的class添加 PHPDoc 的注释即可

![](http://img.justwkj.com/20190602145325.png)
```php
/**
 * Class ShopConf
 * @package app\common
 *
 * 添加IDE提示,如新增加参数,需要增加提示
 *
 * @method static SHOP_ID() 获取APP_ID
 * @method static SHOP_DOMAIN() 获取域名
 * @method static APP_ID() 获取APP_ID
   @method static someClass xxMethod(int $id) 这是描述文字
 *
 */
```
> 
- @method 是PHPDoc标记
- static 标志静态方法
- someClass或$ this – 返回类型
- xxMethod 方法名称
- (int $id) 参数提示

此时再次调用,就可以看到提示了
![](http://img.justwkj.com/20190602145759.png)
