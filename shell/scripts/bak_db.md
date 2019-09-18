# 自动备份数据库

### 1.备份单个数据库
比如把某个数据库备份到 '/opt/bak',创建bak_db.sh 如下
```bash
cd /opt/bak &&\
/bin/mysqldump -uroot -pYOURPWD -B $1 > "$1"_$(date "+%F" -d -1day).sql
```
比如备份 test库, `sh bak_db.sh test`


### 2.备份多个数据库
bak_db_all.sh
```bash
cd /opt/bak &&\
mysql -uroot -pYOURPWD "-BNe show databases;"|grep -v "information_schema"|grep -v "performance_schema" >all_db_name.txt
ALL_DB_NAME=`cat all_db_name.txt`
for db in $ALL_DB_NAME
do
	/bin/mysqldump -uroot -pYOURPWD -B $db > "$db"_$(date "+%F" -d -1day).sql
done
```
直接执行(`sh bak_db_all.sh`)或者设置定时任务,如crontab每天0点执行  
脚本自己需要按照自己实际修改

> 00 00 * * * /bin/sh /server/jobs/bak_db_all.sh >/dev/null 2>&1
