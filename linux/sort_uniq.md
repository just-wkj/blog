# 重复统计 sort uniq

统计日志中重复数据有多少条,这样的需求很常见
实际场景: 接口日志被刷,查看日志,统计用户分别请求多少次

### `uniq`  简介
`uniq` 是去重操作,简单的列举几个参数如下:  
![](http://img.justwkj.com/20190529085316.png)  
文档中的 -u 参数表述并不准确,并非把所有的行直接去重,实际效果是:发现有连续重复的行,则把连续的数据全部删除掉,比如  
![](http://img.justwkj.com/20190529085401.png)  

-d 参数是 和-u有点类似,只展示有连续重复的行,有的话则展示一条,示例如下:  
![](http://img.justwkj.com/20190529090400.png)  

所以单独的使用uniq可能会有各种问题,并不能得到想要的数据,比如只想要有重复的数据,或者没有重复的数据  

### `sort`  简介
`sort` 是排序操作,默认是升序  
-r 参数是反转,也就是倒序排列  
![](http://img.justwkj.com/20190529090344.png)  



### 组合使用

场景1: 获取xx文件中没有重复的行 `cat xx|sort|uniq -u`  
![](http://img.justwkj.com/20190529090441.png)  
场景2: 获取xx文件中重复的行 `cat xx|sort|uniq -d`  
![](http://img.justwkj.com/20190529090451.png)  
场景3: 获取xx文件中重复的行并统计重复行的次数 `cat xx|sort|uniq -dc`  
![](http://img.justwkj.com/20190529090502.png)  

某日志文件,查询不同接口请求次数  
日志文件如下:
![](http://img.justwkj.com/20190529090520.png)  
执行操作 ` awk  '{print $4}' res.log | sort|uniq -c | sort -r  > out.txt`  
可以看到 前面是请求数量,后面是请求接口地址,可以排查问题
![](http://img.justwkj.com/20190529090536.png)
