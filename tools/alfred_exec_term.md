# alfred快读执行终端命令



###  [参考网址: https://www.jianshu.com/p/ec4390117f24](https://www.jianshu.com/p/ec4390117f24)



### 1. 执行pyhton脚本

- preference中点击workflow,按照如图选择keyword to script

  ​	![](http://img.justwkj.com/20210923112056.png)

- 输入关键字和描述，保存

  ![](http://img.justwkj.com/20210923112209.png)

- 右键创建script
  - ![](http://img.justwkj.com/20210923112248.png)

- 输入执行的脚本并保存

  ```py
  # -*- coding:utf-8 -*-
  
  import json
  import urllib
  import urllib2
  
  
  def login():
      url = 'http://p.nju.edu.cn/portal_io/login'
      username = 'xxxx'  # 可将密码等保存至文件
      password = 'xxxxx'
      data = {'username': username, 'password': password}
      postdata = urllib.urlencode(data).encode('utf-8')
      try:
          request = urllib2.Request(url, postdata)
          response = urllib2.urlopen(request)
          res = json.loads(response.read().decode('utf-8'))
          # print res["reply_code"]
      except Exception as e:
          print(e)
  
  
  if __name__ == '__main__':
      login()
  ```

- 快捷键打开alfred输入框，输入你的关键字，回车



### 2. 执行终端命令

步骤和上个例子一样，不过第四步选择terminal command

![](http://img.justwkj.com/20210923112456.png)

​						![](http://img.justwkj.com/20210923112605.png)

