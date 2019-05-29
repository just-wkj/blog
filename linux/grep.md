# grep

grep 快速查询和过滤数据

## 参数介绍

## 常见场景
1. 查询某目录以及子目录下 某些字符出现的次数  
    ```bash
    #查询当前目录下,所有log文件中包含test的出现次数
    grep -Rc --include="*.log"  test
    ```
    结果如下:  
    ![](http://img.justwkj.com/20190529091503.png)
