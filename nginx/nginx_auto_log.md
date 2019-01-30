# nginx 日志轮询切割

原理:
    定时复制日志文件到指定位置,然后平滑重启nginx

操作如下:
1. 创建一个shell脚本 如: /server/jobs/cut_blog_nginx_log.sh
2. 编辑脚本内容如下
    ```bash
     cd /home/wwwlogs &&\
     /bin/mv /home/wwwlogs/blog.moyixi.cn.log  blog.moyixi.cn_$(date "+%F" -d -1day).log 
     /usr/local/nginx/sbin/nginx -s reload
    ```
3. 每天零点执行日志切割,添加定时任务  
    `00 00 * * * /bin/sh /server/jobs/cut_blog_nginx_log.sh >/dev/null 2>&1`
