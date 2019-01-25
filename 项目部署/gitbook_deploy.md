# gitbook 自动部署

此博客也同步更新在gitbook上,不过gitbook经常抽风,或者显示不全,自定义了域名之后也会有文档展示不出来,故打算用gitbook编译生成的静态文件来部署博客

如下:
1. 在服务器上随便建立一个房博客的目录,如 /home/wwwroot/blog.justwkj.com,并把github的仓库clone下来
2. 添加一个自动部署的脚本,如 blog.deploy.sh
    ```bash
    #!/bin/bash
    GIT_DIR="/home/wwwroot/blog.justwkj.com"
    GIT_TMP_LOG="./git_update.log"
    cd $GIT_DIR
    git pull   > $GIT_TMP_LOG 2>&1
    GIT_LOG=`sed -n "1p" ${GIT_TMP_LOG}`
    OK_LOG='Already up-to-date.' 
    [ "$GIT_LOG" != "$OK_LOG" ] && {
        echo '更新中...'
        gitbook init
        gitbook build
        rm -rf $GIT_TMP_LOG
    }
    ```
3. 把blog.deploy.sh放在/server/jobs 目录下(自己随便定义的)
4. 添加定时脚本没分钟执行一次 `* * * * * /server/jobs/gitbook.blog.sh >/dev/null 2>&1`
5. 代码已经自动更新好了,下面就申请一下域名,配置nginx
    > 申请域名配置ssl, 域名指向设置 脚本中的GIT_DIR目录下的_book目录 
    `root /home/wwwroot/blog.justwkj.com/_book`;   
6. 大功告成,此时即可访问你的博客了,如 <https://blog.justwkj.com/>
