# phpstorm terminal 设置git bash

jetbrains 系列的IDE编辑器功能都很多,今天来设置一下内置终端工具  

##起因
为什么要写这篇博客,以前只是简单的设置了git bash作为phpstorm的终端工具,发现经常会存在一些问题:
比如: npm经常无效,自定义的alias无效,自己创建的shell脚本和安装的工具有时候都不能用. 今天写代码的时候,调试接口,调用jq格式化json,发现无效!! 忍不了!!一定要把这个问题解决掉

### 编辑器terminal位置 git bash展示
如图所示编辑器左下角有,Terminal选项,点击会有terminal终端,windows下默认的是cmd,cmd极其的难用
![](http://img.justwkj.com/20190529092146.png)
考虑到我们已经安装过git,此处就用git的终端工具来替代


打开git bash,发现如下 显示的 `MINGW64`
![](http://img.justwkj.com/20190529092156.png)
### phpstorm设置
开始设置phpstorm或者webstorm
![](http://img.justwkj.com/20190529092205.png)
一定要选择安装目录下面的bin\bash.exe 或者bin\sh.exe,这样选择的终端是64位的,
如果选择下面的终端就是32位的
![](http://img.justwkj.com/20190529092215.png),正常来说没有什么问题,如果你自定义了一些自己的脚本,或者安装了一些软件,在32位的终端里面是看不到的
比如我本地安装了jq工具,下面分别是MINGW32 MINGW64查询的结果  
![](http://img.justwkj.com/20190529092223.png)  
![](http://img.justwkj.com/20190529092234.png)  
此外.比如安装了npm,yarn等工具可能都会存在问题,所以推荐使用上面的

### 颜色设置
假设git bash地址是 C:\Program Files\Git\bin\bash.exe
如果直接设置 C:\Program Files\Git\bin\bash.exe 终端工具不会有高亮展示
设置成 "C:\Program Files\Git\bin\bash.exe" --login -i 文件目录会高亮展示,注意双引号不要丢了

