# xdebug调试

1. 查看PHP版本 `<?php phpinfo();`  根据版本安装xdebug,windows用户可以去官网查找对应版本<https://windows.php.net/downloads/pecl/releases/>
2. php.ini 结尾处添加xdebug配置
    > [xdebug]  
      ; 引入xdebug   
      zend_extension="C:\bin\php_xdebug-2.5.4-5.6-vc11-nts.dll"  
      ; 允许远程调试  
      xdebug.remote_enable = 1  
      ; 远程调试自动开启  
      xdebug.remote_autostart = 1  
      ; 远程调试端口,默认9000 可以自行修改  
      xdebug.remote_port =  9001  
3. 重启php,查看php配置,如图所示,则已经生效![](http://img.justwkj.com/20190529091750.png)
4. 打开本地项目,添加断点 ![](http://img.justwkj.com/20190529091802.png)![](http://img.justwkj.com/20190529091816.png)
    浏览器直接访问即可,可以看到phpstorm已经有调试信息![](http://img.justwkj.com/20190529091826.png)
