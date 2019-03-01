# git多个远程库提交 公钥生成 tag global全局信息设置

git提交使用秘钥的比较简单
首先生成秘钥
`ssh-keygen -t rsa -C "注释内容，一般为邮件地址"` 主要是用来标识用户的
之后吧你的id_rsa.pub的内容填写在相应的地方

`git remote add origin git@gixxxxx.git`

`git remote set-url --add origin git@github.com:xxxxX.git`

git push 即可
若设置的名称不同 则可用 `git push -u origin master`提交

添加tag
` git tag -a v0.2 -m "0.2版本 2017-xx上线版本"`  
删除本地tag  
`git tag -d v0.1`  
删除远程tag  
`git push origin :refs/tags/v0.2`  

设置查看git 全局信息  
`git config -l`

`git config --global user.name "Your Name"`  
`git config --global user.email you@email.com`  
