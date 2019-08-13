# justsplit 快速操作一列数据

```bash
#!/bin/bash
#auth: wkj
#date: 2019-08-13
#mail: justwkj@gmail.com
#func: 对一列数据进行简单处理,方便处理数据,直接转换成php的数据,或者用来拼接sql in 语句所需的子元素

usage(){
        echo "Usage: ${0} {-d, -dx, -l, -lx, -a, -as}"
        echo "-d 逗号分隔,带引号"
        echo "-dx 逗号分隔"
        echo "-l 数据合并成1行,逗号分隔,带引号"
        echo "-lx 数据合并成1行,逗号分隔"
        echo "-a 转化为php数组"
        echo "-as 转化为php数组 k=>v"
        exit
}

[ "$#" != 1 ] &&  usage

douhao(){
    awk '{printf "\"%s\",\n", $1}'
}
douhaox(){
    awk '{printf "%s,\n", $1}'
}
line(){
    awk '{printf "\"%s\",", $1}'
}
linex(){
    awk '{printf "%s,", $1}'
}
array(){
 awk 'BEGIN{print "["}{printf "\"%s\",\n", $1,$NF}END{print"]"}'
}
array_with_index(){
	awk 'BEGIN{print "["}{printf "\"%s\" => \"%s\",\n", $1,$NF}END{print"]"}'
}
case "$1" in
        -d)
               douhao
        ;;
        -dx)
               douhaox
        ;;
        -l)
               line
        ;;
        -lx)
               linex
        ;;
        -a)
               array
        ;;
        -as)
               array_with_index
   		;;
        *)
               usage
        ;;
esac

```
