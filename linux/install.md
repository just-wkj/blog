# 安装centos7

公司闲置了几台windows电脑,打算装centos7,后面作为服务器使用

安装步骤大致如下:

###1. 下载安装镜像
去[官网](http://isoredirect.centos.org/centos/7/isos/x86_64/CentOS-7-x86_64-Minimal-1810.iso)找到合适的镜像,
最好直接使用阿里云的镜像,下载比较快
###2. 设置Upan启动盘
使用[UltraISO](https://github.com/chenhaoxiang/Java/blob/master/UltraISO/UltraISO.zip)工具来设置启动盘,
打开UltraISO, 选择刚才下载的镜像,双击打开,之后点击菜单的 启动-写入硬盘镜像,
之后选择硬盘驱动器,就是选择你的U盘,点击写入,执行完毕, 点击返回关闭.
![](http://img.justwkj.com/20190830182604.png)
![](http://img.justwkj.com/20190830182548.png)
![](http://img.justwkj.com/20190830182719.png)

###3. 开始安装
网上好多教程随便搜一个就可以,bios设置U盘启动正确就可以了

###4. 问题处理

安装过程页面出现
![](http://img.justwkj.com/20190830183048.png)选择直接安装,
可能执行不成功,可以根据提示 编辑安装的设置
出现如下图所示
![](http://img.justwkj.com/20190830183157.png),需要修改镜像文件所在的位置
由于我们不知道镜像到底在什么位置,可以如下修改
![](http://img.justwkj.com/20190830183327.png)
后面部分修改为 `initrd=initrd.img linux dd quiet`
重启之后发现如下展示
![](http://img.justwkj.com/20190830183436.png)
可以看到Upan是sdc4
再次修改之前的部分 `inst.stage2=hd:/dev/sdc4 quiet`
重启就可以了

###5. 联网设置
编辑
```bash
TYPE=Ethernet
BOOTPROTO=static
DEFROUTE=yes
PEERDNS=yes
PEERROUTES=yes
IPV4_FAILURE_FATAL=yes
IPV6INIT=yes
IPV6_AUTOCONF=yes
IPV6_DEFROUTE=yes
IPV6_PEERDNS=yes
IPV6_PEERROUTES=yes
IPV6_FAILURE_FATAL=no
NAME=ens34
UUID=a2216a1d-8b95-4b92-9c37-2c60588bf599


ONBOOT=yes              #开启自动启用网络连接
IPADDR0=192.168.120.120  #设置IP地址
PREFIXO0=24             #设置子网掩码
GATEWAY0=192.168.108.252  #设置网关
DNS1=114.114.114.114    #这个是国内的DNS地址，是固定的；
```
之后可以 
重启网络服务 `service network resatrt`  
测试ping `ping www.baidu.com`


