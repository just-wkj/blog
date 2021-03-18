#BurpSuite安装使用

-  [参考文章](https://www.sqlsec.com/2019/11/macbp.html)


平台	下载链接
JAR	https://portswigger.net/burp/releases/download?product=pro&version=2020.11.3&type=Jar
Linux (64-bit)	https://portswigger.net/burp/releases/download?product=pro&version=2020.11.3&type=Linux
Mac OSX	https://portswigger.net/burp/releases/download?product=pro&version=2020.11.3&type=MacOsx
Windows (64-bit)	https://portswigger.net/burp/releases/download?product=pro&version=2020.11.3&type=WindowsX64



首先下载官网封装的 Mac OSX 的版本：burpsuite_pro_macos_v2020_11_3.dmg 正常安装一下。

此时会在应用程序文件夹创建 Burp Suite Professional 的应用程序，国光分析了一下这个官方封装的 BP 里面也打包好了对应的 java 版本，这样哪怕本机是低版本的 JAVA 也是可以正常运行这个 BP 的，美滋滋呀。

下载拷贝注册机
接着下载好国光打包的注册机和汉化包：

文件名	下载地址
macOS Burpsuite.zip （文件大小：995.9 K）	蓝奏云
解压注册机和汉化包，在终端中打开 BP 程序内置的 JAVA 文件路径：

Bash
cd /Applications/Burp\ Suite\ Professional.app/Contents/java/app && open .

自定义启动参数
复制成功Bash
open /Applications/Burp\ Suite\ Professional.app/Contents/vmoptions.txt
下面给了两种示例：

汉化脚本加载启动
Properties
-noverify
-javaagent:BurpSuiteLoader.jar
-javaagent:BurpSuiteCn.jar
-Dfile.encoding=utf-8
-XX:MaxRAMPercentage=50
-include-options user.vmoptions
英文原版启动
Properties
-noverify
-javaagent:BurpSuiteLoader.jar
-XX:MaxRAMPercentage=50
-include-options user.vmoptions
选择自己想要的启动方式保存并退出，国光这里选择了喜闻乐见的汉化启动方式：


激活 Burpsuite
如果之前你没有激活过 Burpsuite 的话，那么第一次可能需要提示输入许可证秘钥，运行如下命令打开注册机：

复制成功Bash
chmod +x /Applications/Burp\ Suite\ Professional.app/Contents/java/app/license_key.sh && sh /Applications/Burp\ Suite\ Professional.app/Contents/java/app/license_key.sh
下面国光啰嗦一下激活方式，理论上这里应该直接去百度的，但是这篇文章流量实在太多了，国光我妥协了。

首先将注册机的「License」复制粘贴到 BP 的许可证秘钥框中，然后点击「下一步」：


接着点击「手动激活」：


首先点击注册机里面的「辅助请求」按钮，将信息粘贴到注册机中的 「Activation Request」中，此时注册机会在「Activation Response」生成返回信息，然后将这个返回信息粘贴到 BP 中，点击「粘贴响应」按钮即可，最后点击「下一个」：


最后如果顺利的话就激活成功了，祝你使用愉快：


授人以鱼
BurpSuite Pro 2020.9.2

资源下载
因为蓝奏云免费账户最高支持单个文件 100MB 大小，而分享的这个 BP 大小为 500 多MB（内置了浏览器），所以国光使用分卷压缩了，一共分了 6 卷。另外这个版本也是自带高版本 JDK 的，原理是使用了相对路径的 JAVA 来调用执行 BurpSuite。

总的来说国光还是比较推荐大家安装的，而且这个版本已经过卡巴斯基检测了，理论上不会存在后门：

文件名	下载链接
BurpSuite2020.9.2.7z.001	蓝奏云
BurpSuite2020.9.2.7z.002	蓝奏云
BurpSuite2020.9.2.7z.003	蓝奏云
BurpSuite2020.9.2.7z.004	蓝奏云
BurpSuite2020.9.2.7z.005	蓝奏云
BurpSuite2020.9.2.7z.006	蓝奏云
又因为蓝奏云不支持001、002 等这种奇怪的后缀上传，所以国光都在文件名后面手动添加了iso后缀。

所以大家下载完这两个文件的时候，把这 6 个文件放在同一级的目录下，然后手动删掉 iso 后缀，然后解压 BurpSuite2020.9.2.7z.001即可。

因为老是有粗心的网友评论里面留言说不能解压，所以国光我下面再啰嗦一遍：

下载两个文件到同一级目录
手动删掉这 6 个文件的 iso 后缀
解压 BurpSuite2020.9.2.7z.001
基本使用
打包的 BP 下面一共 3 个 sh 脚本，下面简单概括一下 3 个脚本的作用：

Bash
# 在 BP 的目录下给脚本执行权限
chmod +x *.sh

# 给 jar 执行权限
chmod +x *.jar

# 启动英文版 BP
start_EN.sh

# 启动中文版 BP
start_CN.sh

# 如果第一次 BP 启动需要 licensekey 的话 那么就执行此脚本重新激活一下 后面就不用再激活了
license_key.sh
优化体验
接下来的任务就是如何将脚本创建快捷方式到 「启动台」了，可以参考下面的 「授人以渔」部分，自学一下问题不大。

授人以渔
资源下载
Burp Suite Pro 2.1.05 破解版
Oracle Java 版本为 1.8 的版本是可以成功运行的，感兴趣的朋友可以自行去安装配置这个 JAVA 版本


文件名	下载地址1	下载地址2
burpsuite_pro_v2.1.05.jar	OneDriver	买不起会员
burp-loader-keygen-2_1_05.jar	OneDriver	蓝奏云
BurpSuite Pro 2.1.06 破解版
Oracle Java 版本为 1.8 的版本是可以成功运行的，感兴趣的朋友可以自行去安装配置这个 JAVA 版本


文件名	下载地址1	下载地址2
burpsuite_pro_v2.1.06.jar	OneDriver	买不起会员
burp-loader-keygen-2_1_06.jar	旁边蓝奏云速度比较快	蓝奏云
Burp Suite Pro 2.x 系列汉化包
文件名	下载地址1	下载地址2
BurpSuiteCn.jar	OneDriver	蓝奏云
Burp Suite icns 图标
国光从官方正版的 BP 里面提取出的正版的图标文件

文件名	下载地址	备用链接
BurpSuite.icns	Onedriver	蓝奏云
Burpsuite激活
这里不赘述了，大家都是做安全的，我这里再啰嗦的话 就显得很多余了。如果 BurpSuite 没有激活的话，那么让我帮你百度一下：

让我帮你百度一下 Burpsuite如何激活 2333

BP 不激活的话，下面的命令行封装的 APP 每次启动都是要重新激活的，这样的话就很难受了。

命令行模拟启动
启动示例
当前目录结构如下：

Bash
$ pwd
/Users/sec/Documents/Sec/BurpSuite/2.1.06

$ tree
.
├── BurpSuite.icns
├── BurpSuiteCn.jar
├── burp-loader-keygen-2_1_06.jar
└── burpsuite_pro_v2.1.06.jar

0 directories, 4 files
知道路径的情况下，下面尝试命令行直接启动试试看：

Bash
$ pwd
/Users/sec/Documents/Sec/BurpSuite/2.1.06

$ java -Xdock:icon=BurpSuite.icns -javaagent:BurpSuiteCn.jar -Xbootclasspath/p:burp-loader-keygen.jar -jar burpsuite_pro_v2.0.11beta.jar
-Xdock:icon=BurpSuite.icns 用于在 macOS 的 Dock 栏显示图标，该方法由评论区的网友 Nick Yan 提供

启动测试没有问题，下面使用&&将命令合为一个命令，用于下方制作APP使用，不同版本的启动命令分别如下，大家用的话修改成对应自己的目录信息即可。

2.1.05 启动命令
Bash
cd /Users/sec/Documents/Sec/BurpSuite/2.1.05 && java -noverify -Xdock:icon=BurpSuite.icns -javaagent:BurpSuiteCn.jar -Xbootclasspath/p:burp-loader-keygen-2_1_05.jar -jar burpsuite_pro_v2.1.05.jar
2.1.06 启动命令
Bash
cd /Users/sec/Documents/Sec/BurpSuite/2.1.06 && java -noverify -Xdock:icon=BurpSuite.icns -javaagent:BurpSuiteCn.jar -Xbootclasspath/p:burp-loader-keygen-2_1_06.jar -jar burpsuite_pro_v2.1.06.jar
制作应用程序
新建工作流程
借助 macOS 自带的自动操作 automator.app 可以轻松地将命令行封装成一个应用程序出来：


首先打开自动操作，选择左下角的新建文稿:


选择文稿类型为「应用程序」，点击「选取」：


左侧列表中找到运行Shell脚本然后拖入进去：


把上文中的命令行模拟启动对应的命令粘贴进来，点击右上角运行测试一下看看能不能成功启动：


OK 测试应该是没有问题的，然后将 BP 关掉，command+s保存一下，保存的文件格式选择应用程序，自定义命名一下，然后保存到应用程序中：


然后就可以在启动台LaunchPad中找到这个孤零零的应用程序了，点击一下就可以成功启动BP了：


美化BP图标
当然这么丑的图标，国光我是看不下去的，下面来使用上文中下载的 BP 图标文件。

在「应用程序」文件夹中找到我们刚刚制作的「BurpSuite」应用程序，「右键」点击「显示简介」，然后将下载好的icns图标文件拖入到简介的左上角：


最后的效果如下：


OK 美化成功，现在就可以在 mac 下优雅地启动 BP 了~~

其他配置
Firefox Dev
经国光本人一个个测试，最终发现使用 Firefox 的 Dev 版本中的 56.0b9 版本可以完美地安装老版本的Hackbar 插件，安装过后，手工注入的速度嗖嗖的上涨，并且 Firefox 可以完美的导入 BP 的 SSL 证书，很方便我们后期使用 BurpSuite 抓取 HTTPS 网站。

平台	下载地址（官方）	备用下载（Onedriver）
Windows64位	Firefox Setup 56.0b9.exe	Firefox Setup 56.0b9.exe
Windows32位	Firefox Setup 56.0b9.exe	32位 国光比较懒 不提供备用下载
Linux x86_64	firefox-56.0b9.tar.bz2	firefox-56.0b9.tar.bz2
macOS	Firefox 56.0b9.dmg	Firefox 56.0b9.dmg
关闭自动更新
首先下载完启动Firefox Developer Edition一定要断网！断网！断网！因为火狐的策略问题，默认会自动更新，联网安装的话，刚刚安装完的时候就会被强制更新最新版本了，所以断网安装大法保平安。安装完成后，设置里面勾选不检查更新。然后就可以恢复联网了：


允许第三方插件
火狐安全策略的原因是无法顺利安装一些第三方插件的，得手动去高级选项里面关闭一下，浏览器打开：

about:config

点击我了解此风险：



来源: 国光
文章作者: 国光
文章链接: https://www.sqlsec.com/2019/11/macbp.html
咳咳又想白嫖文章？本文章著作权归作者所有，任何形式的转载都请注明出处。