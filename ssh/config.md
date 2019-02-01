# ssh config

服务器比较多的情况下快速,如何才能快速登录各种服务器.  
可以选择一些软件,我比较常用的是 `Secure CRT` 这个软件,一次设置之后,以后新增服务器只需要复制一次就可以了.
有些情况下还是喜欢在git bash 或者各种编辑器的终端中直接登录服务器,一旦服务器多了就需要采取一些措施

1. 使用alias
    ```bash
    alias sshxx="ssh root@HOST -p PORT"
    alias hui="ssh root@HOST -p PORT '/bin/sh /jobs/xx.sh'"
    ```
    使用了alias之后 直接在终端执行 `sshxx` 即可登录服务器
2. 配置ssh config
    ```
    Host testserver1
        HostName YOUR_SERVER_IP
        # 用户名
        User root
        # 端口
        Port 22
        # 本地的id_rsa私钥 特别注意 IdentityFile这个单词不要拼写错了
        IdentityFile /c/Users/365/.ssh/id_rsa
    
    Host testserver2
        HostName YOUR_SERVER_IP
        User root
        Port 2233
        IdentityFile /c/Users/365/.ssh/id_rsa
    ```
    配置之后 使用 `ssh testserver2` 即可登录服务器  
    配置config也可以方便的使用scp来上传文件,比如: 把当前目录下的a.txt上传到服务器的/tmp目录下 `scp a.txt testserver2:/tmp`
