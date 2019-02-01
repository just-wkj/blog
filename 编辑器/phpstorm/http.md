# http调试

接口调试的工具有很多,我比较常用是postman,很不错的  
去年看到phpstorm更新可以使用http接口调试了,但是一直没怎么去研究,前段时间使用了一阵子,觉得还可以

phpstorm下面有两种方式调试

1. 图形话的调试 Tools->Http Clinet->Test RestFul Web Service,打开的界面如图中下半部分
 ![](https://ws1.sinaimg.cn/large/0063sFGSgy1fzqr5hjejyj30wz0pa40f.jpg)
    可以方便直接设置get,post数据等,不过这种界面调试的后面应该会被慢慢废弃掉
    
2. 创建一个.http的文件如 test.http
    ![](https://ws1.sinaimg.cn/large/0063sFGSgy1fzqrahfnh5j310o0gl0tr.jpg)
    文件创建之后,上面编辑器默认生成的一些简单的示例,右侧分别是查看请求历史和请求案例类型  
   如果想设置不同的环境或者默认的参数值,可以添加一个`http-client.env.json`的文件,比如:
   ```json
    {
       "test":{
         "host":"127.0.0.1",
         "city":"nj"
       },
       "线上":{
         "host":"127.0.0.1:7788",
         "city":"njxx"
       }
     }
   ```
   比如 `GET http://{{host}}/api/item?id={{city}}` 这里面的host和city都是在上面的json配置文件中设置的,运行的时候选择线上,效果如下
   ![](https://ws1.sinaimg.cn/large/0063sFGSgy1fzqrex8ctaj30rt06b74h.jpg)
    点击右上角的 show http history可以查看请求的日志,日志里面也有返回的值都会保存为文件,也可以在历史记录中再次请求
    ![](https://ws1.sinaimg.cn/large/0063sFGSgy1fzqrhe6521j30d605b748.jpg)
