# github webhook部署

平时使用自动部署,大多是直接git或者jenkins加git来实现,具体就是定时的去pull代码,发现有代码就去执行部署的一系列操作,webhook很少用,觉得有些麻烦

此处简单的记录一下github的webhook使用,选择php脚本来接口处理

下面是具体步骤:

1. github创建仓库,在webhook页面设置你的回调地址,并且设置你的秘钥
   ![](http://img.justwkj.com/20190529091345.png)

2. 设置好之后,我们来写php的回调处理代码

   ```php
   	<?php
   	$secret = "第一步骤设置的秘钥";
   	$signature = $_SERVER['HTTP_X_HUB_SIGNATURE'];
   	if ($signature) {
   	    $hash = "sha1=" . hash_hmac('sha1', file_get_contents('php://input'), $secret);
   	    if (strcmp($signature, $hash) == 0) {
   	    	 //如果php没有限制 exec,shell_exec等函数话,在这里就可以直接执行你的脚本了
   	    	 //因为我限制了执行函数,曲线救国,生成一个临时文件并写入当前的时间戳
   			 file_put_contents('./blog_justwkj_com.txt', time());
   			 echo 'ok';
   	        exit();
   	    }
   	}
   	http_response_code(404);
   ```

3. 验证是否成功,点击 Redeliver ![](http://img.justwkj.com/20190529091405.png),发现响应是200, 而且生成了临时文件 blog\_justwkj\_com.txt,此时则说明webhook执行成功了.如果你的php没有限制系统命令执行函数如:shell_exec,则可以直接在签名匹配成功之后,执行我们相应的脚本即可.

4. 为了安全起见,我在服务器上限制了php执行系统的命令,所以根据临时文件的内容来判断是否有推送.如果有推送,则文件内容为时间戳,我们执行部署脚本,再把文件清空,等待再次收到推送,写了一个简单的shell脚本,以守护进程的方式执行,即可完成部署的操作

   ```bash
   #! /bin/sh
   
   while [ true ]
   do
       [ -s blog_justwkj_com.txt ] && {
   
               echo 'updating...'
               # 执行部署脚本, 更新代码,环境设置等等..
               
               # 清空临时文件
               cat  /dev/null > blog_justwkj_com.txt
   
       }
       sleep 5
   done
   ```

5. 总结一下吧,对于限制了系统函数执行的php脚本而言,webhook显得很鸡肋了.需要一个守护进程的shell脚本配合.不如直接写一个脚本,每分钟或者自定义的时间定时的`git pull`一次代码,发现有数据更新,则执行部署的相关操作. 我目前基本都是直接使用脚本定时抓取git内容来部署,操作最简单,webhook用的少.
