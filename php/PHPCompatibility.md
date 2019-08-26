# php不同版本语法检查

不同项目的的php版本不同,导致某些语法解析报错  
比如 一个项目用php5.2 一个项目用php7.2, 5.2的项目可能已经不继续开发了,但是可能会
新增一部分小功能. 可能不经意的就是定义数组为 [], 或者使用类常量等都会导致报错.

可以使用phpcs来处理这类问题,需要安装一个 [PHPCompatibility](https://www.sitepoint.com/quick-intro-phpcompatibility-standard-for-phpcs-are-you-php7-ready/) 的拓展

简单的列一下安装和使用步骤:
语法检查可以安装在某个项目下,也可以全局安装,本人选择全局安装,当做工具来使用
步骤如下:

全局安装phpcs  
```
#全局安装phpcs
composer global require "squizlabs/php_codesniffer=2.*" 

# 进入目录
cd ~/.composer/vendor/squizlabs/php_codesniffer/CodeSniffer/Standards

# clone PHPCompatibility 到文件夹
git clone https://github.com/wimg/PHPCompatibility.git

# 设计
~/.composer/vendor/bin/phpcs --config-set installed_paths ~/.composer/vendor/squizlabs/php_codesniffer/CodeSniffer/Standards/PHPCompatibility

```

安装配置之后就可以使用了

```bash
# 只检测php文件, 是否满足5.2的语法
~/.composer/vendor/bin/phpcs --standard=PHPCompatibility --extensions=php --runtime-set testVersion 5.2 ./path
```
执行结果如下 ![](http://img.justwkj.com/20190826174733.png)

如果把结果保存在文件中可以添加参数 `--report-full=filename`

```bash
# 只检测php文件, 是否满足5.2的语法 检查结果保存在php.txt中
~/.composer/vendor/bin/phpcs --standard=PHPCompatibility --extensions=php --report-full=php.txt --runtime-set testVersion 5.2 ./path
```
