# gogs搭建Git服务平台

[gogs git服务][1]
### 二进制包没有app.ini 去github上找一个
[app.ini 地址][2]

###安装git
    yum install -y git
    git --version
    git version 1.7.1

###安装mysql-server
    yum install -y mysql-server
    mysql --version
    mysql  Ver 14.14 Distrib 5.1.73, for redhat-linux-gnu (x86_64) using readline 5.1

###启动数据库
    service mysqld start
    chkconfig mysqld on

###创建gogs数据库
    cd /home/git/gogs/scripts
### mysql -u root -p < mysql.sql
    mysql -u root -p
###（输入密码，无密码直接跳过）
    set global storage_engine = 'InnoDB';
    create database gogs character set utf8 collate utf8_bin;
    create user 'gogs'@'localhost' identified by 'gogs';
    grant all privileges on gogs.* to 'gogs'@'localhost';
    flush privileges;
    exit;


  [1]: https://gogs.io/
  [2]: https://github.com/gogits/gogs/blob/master/conf/app.ini
