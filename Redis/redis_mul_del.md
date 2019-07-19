# redis批量删除key
redis 只有del单个删除键的命令,如果需要批量删除,可以使用xargs来操作

` redis-cli -a PASSWD keys "[45678]*"| xargs redis-cli -a PASSWD del`

如果要删除指定数据库的数据,可以用-n 选择数据库
`redis-cli -a PASSWD -n 0  keys "4*" | xargs redis-cli -a PASSWD -n 0  del`
