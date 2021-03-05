# mac系统搭建python+uiautomator2环境

介绍：这个主要用来弄安卓自动化，底层调用adb命令来驱动安卓设备，下面介绍在mac上安装环境

一:首先电脑安装adb环境

当然在mac系统下安装很方便，一个命令即可

终端输入：brew cask install android-platform-tools

如果跟我看到的一致 那么就是安装成功了

![](http://img.justwkj.com/20210305094645.png)


当我们安装完成之后，需要检测下看看可以连接上你的真机

操作步骤：

1.通过数据线连接,然后在手机中打开USB调试



2.终端运行 adb devices 可以看到相对应的设备就是代表连接上了



 

二：安装uiautomator2模块

1.终端输入：pip3 install  uiautomator2  -i https://pypi.douban.com/simple



 

三：编写代码

import uiautomator2

device = uiautomator2.connect()
print(device.device_info)
运行此代码，你的手机会提示需要安装二个文件，根据提示安装即可





如果你运行上方代码成功输出以下信息，则代表你的环境基本安装好了



 

四：安装weditor(这个主要是用于元素定位的，可以将手机投屏到电脑，直接点击区域定位到元素)

pip3 install weditor -i https://pypi.douban.com/simple



安装完之后，运行weditor

终端输入 python3 -m webditor



当你终端输入之后，会自动打开您的浏览器，你就可以看到如下界面(必须USB连接手机，并且打开实时按钮)，这个时候你就可以定位你的APP元素了




————————————————
版权声明：本文为CSDN博主「Tester_xjp」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/tester_xjp/article/details/104870045
参考链接：https://www.cnblogs.com/fnng/p/8486863.html
