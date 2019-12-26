# B站视频下载脚本

```bash
#!/bin/bash
# author: wkj
# mail: justwkj@gmail.com
# desc: B站视频快速下载

BILIBILI_URL='https://www.bilibili.com/video/av'
VIDEO=0
P=0

usage(){
	echo "Usage: 
        -v number  视频文件编号, eg: -v 57721377
        -p number 下载某个章节  eg: -p 1
        -n number 所有视频章节 eg: -n 20

        demo: ${0} -v 57721377 -n 20 下载 av57721377 1-20个章节的视频
        demo: ${0} -v 57721377 -p 2 下载 av57721377  第2个章节的视频
    " 
	exit
}


if [ "$1" != "-v" ];then
usage 
fi

if [ "$2" -gt 0 ] 2>/dev/null ;then 
 VIDEO=$2
else 
 echo "-v $2 is not number "
 exit 
fi 

if [ "$4" -gt 0 ] 2>/dev/null ;then 
P=$4
else 
 echo "-n|-p $4 is not number "
 exit 
fi 

donwnloadOne(){
    echo "you-get  https://www.bilibili.com/video/av${VIDEO}?p=${1}"
    you-get  https://www.bilibili.com/video/av${VIDEO}?p=${1}
}

downloadSome(){
    for((i=1;i<=${1};i++));
    do
    echo "you-get  https://www.bilibili.com/video/av${VIDEO}?p=${i}"
    you-get  https://www.bilibili.com/video/av${VIDEO}?p=${i}
    done
}

case "$3" in
	-n)
		downloadSome ${N}
	;;
	-p)
    echo ${P}
     donwnloadOne ${P}
		;;
	*)
		usage
	;;
esac
```
