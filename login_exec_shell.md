# 登录服务器并执行脚本

`ssh renhetest "cd /home/wwwroot/renhetest ;  php think clear"`

双引号，必须有，如果不加双引号，第二个ls命令在本地执行，同时两个命令之间用分号隔开


有些远程执行的命令内容较多，单一命令无法完成，考虑脚本方式实现：
```shell
#!/bin/bash
ssh user@remoteNode > /dev/null 2>&1 << eeooff
cd /home
touch abcdefg.txt
exit
eeooff
echo done!

```
远程执行的内容在“<< eeooff ” 至“ eeooff ”之间，在远程机器上的操作就位于其中，注意的点：

<< eeooff，ssh后直到遇到eeooff这样的内容结束，eeooff可以随便修改成其他形式。
重定向目的在于不显示远程的输出了。在结束前，加exit退出远程节点。
