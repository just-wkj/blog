# mysql 数据备份与恢复
`mysqldump -u 用户名 -p [密码] 数据库 > 备份文件`  
`mysqldump -u username -p dbname table1 table2 ...> BackupName.sql`  
-B 和 --databases 意思一样
`mysqldump -u username -p -B dbname > BackupName.sql`  
如 `mysqldump -u root -p test > /opt/test_20170326.sql`  

两个参数介绍 -t 只备份数据 -d 只备份表结构  

备份多数据库  
`mysqldump -u username -p --databases dbname2 dbname2 > Backup.sql`  
备份所有数据库  
`mysqldump -u username -p -all-databases > BackupName.sql`  
示例:  
`mysqldump -u -root -p -all-databases > D:\all.sql`  

还原  
`mysql -uroot -p test < /opt/test_20170326.sql `  
如果用-B 不需要选择数据库了  
`mysql -uroot -p < /opt/test_20170326.sql `  
查看下数据库列表  
`mysql -e "show databases;"`  

文件可能很大需要压缩一下  
`mysqldump -u root -p -B baktest1 |gzip> /opt/test.sql.gz`  

mysql 导入数据  
mysql命令行下 `source /data/xxx.sql`  

---

一篇参数介绍的文章 <https://segmentfault.com/a/1190000000621104>
