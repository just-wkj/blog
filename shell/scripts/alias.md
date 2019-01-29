# 快速操作alias

---
有时候会经常使用alias每次都要去修改文本,比较费事,于是写一个脚本来方便编辑操作

```bash
# author: wkj
# mail: justwkj@gmail.com
# desc: 快速添加alias和删除alias

#!/bin/bash
VIMRC=~/.bashrc
usage(){
    echo "Usage ${0} [-a/-d] aliasName realCommand"
    exit
}

case $1 in
    -a)
        [ $# -ne 3 ] && usage
        # 删除旧的alias
        sed -i -e '/^alias '"${2}"'/d' $VIMRC
        echo alias ${2}=\"${3}\" >> $VIMRC
        ;;
    -d)
        [ $# -ne 2 ] && usage
        sed -i -e '/^alias '"${2}"'/d' $VIMRC
        ;;
    *)
       usage
        ;;
esac


```
