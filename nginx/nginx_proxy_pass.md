# nginx proxy_pass使用


### 1. proxy_pass 绝对路径和相对路径的使用区分
在nginx中配置proxy_pass代理转发时，如果在proxy_pass后面的url加/，表示绝对根路径；如果没有/，表示相对路径，把匹配的路径部分也给代理走。

假设下面四种情况分别用 http://192.168.1.1/proxy/test.html 进行访问。

```nginx
#第一种：
location /proxy/ {
    proxy_pass http://127.0.0.1/;
}
#代理到URL：http://127.0.0.1/test.html


#第二种（相对于第一种，最后少一个 / ）
location /proxy/ {
    proxy_pass http://127.0.0.1;
}
#代理到URL：http://127.0.0.1/proxy/test.html


#第三种：
location /proxy/ {
    proxy_pass http://127.0.0.1/aaa/;
}
#代理到URL：http://127.0.0.1/aaa/test.html


#第四种（相对于第三种，最后少一个 / ）
location /proxy/ {
    proxy_pass http://127.0.0.1/aaa;
}
#代理到URL：http://127.0.0.1/aaatest.html

```
### 2. proxy 使用
之前自己搭建一个git服务器, 代理本地3000端口
``` nginx
server {
        listen       80;
        server_name  git.moyixi.cn;
        #location / {
         #   proxy_pass   http://127.0.0.1:3000;
        #}
        rewrite ^/(.*) https://git.moyixi.cn/$1 last;
}

server {
        listen 443;
        server_name git.moyixi.cn; #填写绑定证书的域名
        ssl on;
        ssl_certificate cert/1_git.moyixi.cn_bundle.crt;
        ssl_certificate_key cert/2_git.moyixi.cn.key;
        ssl_session_timeout 5m;
        ssl_protocols TLSv1 TLSv1.1 TLSv1.2; #按照这个协议配置
        ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:HIGH:!aNULL:!MD5:!RC4:!DHE;#按照这个套件配置
        ssl_prefer_server_ciphers on;
        location / {
            proxy_pass http://127.0.0.1:3000;
        }
    }
```
---
参考博客 
- <https://blog.csdn.net/zhongzh86/article/details/70173174>
- <https://www.jianshu.com/p/b113bd14f584>
