# nginx rewrite 重写

1. 路由重写,定向跳转,比如域名跳转或者http需要跳转到https经常会使用到
    ```nginx
        server { 
          listen 80;
          server_name www.justwkj.com;
          # 域名直接跳转到新的域名下带上参数 可以重写去除index之类的
          rewrite ^/(.*) http://blog.moyixi.cn/$1 permanent;
        }
    ```
2. 重写indx
    ```nginx
    ##return 示例 禁止IE和FireFox 
    if ($http_user_agent ~* "spider|MSIE"){
        return 403;
    }
    # 文件不存在重写index
    location / {
       if (!-e $request_filename){
            rewrite ^/(.*)$ /index.php/$1 last;
       }
    }
    ``` 
3. flag标志位
    - last : 相当于Apache的[L]标记，表示完成rewrite  
    - break : 停止执行当前虚拟主机的后续rewrite指令集  
    - redirect : 返回302临时重定向，地址栏会显示跳转后的地址  
    - permanent : 返回301永久重定向，地址栏会显示跳转后的地址  
    > 因为301和302不能简单的只返回状态码，还必须有重定向的URL，直接用return只是返回了状态码,需要根据情况选择return还是rewrite的redirect还是permanent  
     last 和 break 区别：
     last一般写在server和if中，而break一般使用在location中
     last不终止重写后的url匹配，即新的url会再从server走一遍匹配流程，而break终止重写后的匹配
      break和last都能组织继续执行后面的rewrite指令  
      下面是示例
      
      ```nginx
        server {
            listen 80 default_server;
            server_name test.com;
            root www;
        
            location /break/ {
                rewrite ^/break/(.*) /test/$1 break;
                echo "break page";
            } 
        
            location /last/ {
                 rewrite ^/last/(.*) /test/$1 last;
                 echo "last page";
            }    
        
            location /test/ {
               echo "test page";
            }
        }
        # 请求 http://test.com/break/*** 输出 break page
        # 请求 http://test.com/last/***  输出 test  page
      ```
4. 表达式
    - 直接比较变量和内容时，使用=或!=
    - ~正则表达式匹配，~*不区分大小写的匹配，!~区分大小写的不匹配
    - -f和!-f用来判断是否存在文件
    - -d和!-d用来判断是否存在目录
    - -e和!-e用来判断是否存在文件或目录
    - -x和!-x用来判断文件是否可执行
    
5. apache重写
> 很少使用apache了,把之前用过的部分配置贴出来
    
```apacheconfig
RewriteEngine on
RewriteCond %{REQUEST_FILENAME} !-d
RewriteCond %{REQUEST_FILENAME} !-f
RewriteRule ^(.*)$ index.php?/$1 [L]

#获取参数
RewriteEngine on
RewriteCond %{REQUEST_FILENAME} !-d
RewriteCond %{REQUEST_FILENAME} !-f
RewriteRule ^(.*)$ /index.php?%{QUERY_STRING}
```
