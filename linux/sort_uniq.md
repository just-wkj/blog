# 重复统计

统计日志中重复数据有多少条,这样的需求很常见
实际场景: 接口日志被刷,查看日志,统计用户分别请求多少次

### `uniq`  简介
`uniq` 是去重操作,简单的列举几个参数如下:  
![](https://ws1.sinaimg.cn/large/0063sFGSgy1g3h4dnk7mhj30qi0bq74o.jpg)  
文档中的 -u 参数表述并不准确,并非把所有的行直接去重,实际效果是:发现有连续重复的行,则把连续的数据全部删除掉,比如  
![](https://ws1.sinaimg.cn/large/0063sFGSgy1g3h4iryaxlj30g607pdfu.jpg)  

-d 参数是 和-u有点类似,只展示有连续重复的行,有的话则展示一条,示例如下:  
![](https://ws1.sinaimg.cn/large/0063sFGSgy1g3h4n5128sj30gm08ljrf.jpg)  

所以单独的使用uniq可能会有各种问题,并不能得到想要的数据,比如只想要有重复的数据,或者没有重复的数据  

### `sort`  简介
`sort` 是排序操作,默认是升序  
-r 参数是反转,也就是倒序排列  
![](https://ws1.sinaimg.cn/large/0063sFGSgy1g3h4q9be2ij30gg0d9aa6.jpg)  



### 组合使用

场景1: 获取xx文件中没有重复的行 `cat xx|sort|uniq -u`  
![](https://ws1.sinaimg.cn/large/0063sFGSgy1g3h4sffx9xj30ht09jaa4.jpg)  
场景2: 获取xx文件中重复的行 `cat xx|sort|uniq -d`  
![](https://ws1.sinaimg.cn/large/0063sFGSgy1g3h4tk6yw7j30js09uaa5.jpg)  
场景3: 获取xx文件中重复的行并统计重复行的次数 `cat xx|sort|uniq -dc`  
![](https://ws1.sinaimg.cn/large/0063sFGSgy1g3h4v2wg5yj30iz07zmx7.jpg)  

某日志文件,查询不同接口请求次数  
日志文件如下:
![](https://ws1.sinaimg.cn/large/0063sFGSgy1g3h4y9iganj31230ft0xa.jpg)  
执行操作 ` awk  '{print $4}' res.log | sort|uniq -c | sort -r  > out.txt`  
可以看到 前面是请求数量,后面是请求接口地址,可以排查问题
![](https://ws1.sinaimg.cn/large/0063sFGSgy1g3h502kg8wj30fm0cegm3.jpg)
