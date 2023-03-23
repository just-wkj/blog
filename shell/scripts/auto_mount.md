# 自动挂载脚本

> 由于nas服务器的录音文件是存放在nas服务器上的，所以需要将nas服务器上的录音文件挂载到本地服务器上，这样才能够访问到录音文件。
> 但是每次重启服务器后，都需要手动挂载，所以需要写一个脚本，每次重启服务器后，自动挂载nas服务器上的录音文件。

```shell
#!/bin/bash
#*******************************************
#*Author: wangkeji
#*Date: 2023-03-23
#*email: justwkj@gmail.com
#*Description: 检测是否挂载nas目录
#*******************************************

if mount|grep '/home/wwwroot/dfy_record/Sounds' > /dev/null; then
  echo "远程目录已经挂载"
else
  # 如果没有挂载，则尝试挂载该目录
  echo "正在尝试挂载远程目录"
  mount -t nfs 192.168.0.62:/volume3/录音 /home/wwwroot/dfy_record/Sounds
  if [[ $? -eq 0 ]]; then
    echo "远程目录挂载成功"
    cd /home/wwwroot/dfy_record/ || exit
    chmod -R 777 ./Sounds/*
    chmod -R 777 ./Sounds
    echo "修改权限成功"
  else
    echo "远程目录挂载失败"
  fi
fi

```
