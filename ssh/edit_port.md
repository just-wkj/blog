# ssh 端口修改

突然收到邮件,![请输入图片描述][1]
立刻上服务器,修改ssh端口,root密码

`vim  /etc/ssh/sshd_config`  
> Port 新的端口号  

`service sshd restart`

修改root密码
`passwd root`

查看定时任务crontab是否有异常

  [1]: http://ww1.sinaimg.cn/large/0063sFGSgy1fe8n0o1zpyj30lx079jry.jpg
