# gitbook自动更新

```bash
# author: wkj
# date: 2019-01-24
# desc: github同步数据,编译成静态文件,并推送到gitee仓库中

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
   git push gitee master
}
```

