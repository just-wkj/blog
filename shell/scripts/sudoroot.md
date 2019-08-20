# shell脚本切换root

执行某些脚本的时候需要切换到root,才可以执行
切换root用户需要输入密码,可以采用如下形式

`echo YOUR_PWD| sudo -S command`

示例如下: 重启本地web服务

~/server/etc/root_key 里存放root密码
```bash
ROOT_KEY=~/server/etc/root_key
echo "kill nginx ..."
cat $ROOT_KEY |sudo -S pkill nginx
echo "start nginx..."
cat $ROOT_KEY |sudo -S nginx
echo "kill php-fpm..."
pkill php-fpm
echo "start php-fpm...."
echo "ok"
exit 0
```
