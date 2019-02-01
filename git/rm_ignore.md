# git ignore 去除已经加入仓库的文件

如果项目中不小心把某些不要加入仓库的东西,比如缓存,上传文件等加入了git仓库,可以修改.gitignore之后使用如下命令执行,重新提交即可

`git rm -r --cached .`

简单的介绍下 rm 和git rm区别:
- rm 是bash的命令和shell无关
- git rm 是git下的删除命令, 删除git追踪文件，同时删除文件
- git rm –cached 也是是git下的删除命令, 删除git追踪文件，但是不删除文件
