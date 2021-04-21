# php安装oracle

1. [下载两个压缩包](https://www.oracle.com/database/technologies/instant-client/macos-intel-x86-downloads.html)
    - Basic Package
    - SDK Package  
    下载后将2个压缩包进行解压并将SDK压缩包里的sdk文件夹复制到Basic文件夹
2. 链接  
   链接的目的是让驱动能被其他开发工具或软件能找到驱动目录
    假设上面的解压后的文件夹为 `/extensions`，则执行如下命令
   `ln -s /extensions/libclntsh.dylib /usr/local/lib/`
3. 安装oci8
    `
 # 根据版本安装,可以直接通过下载链接查看
`pecl install oci8-2.2.0
`
  
提示文字如下   
`
running: phpize  
Configuring for:  
PHP Api Version:         20180731  
Zend Module Api No:      20180731  
Zend Extension Api No:     
Please provide the path to the ORACLE_HOME directory. Use 'instantclient,/path/to/instant/client/lib' if you're compiling with Oracle Instant Client [autodetect] :
`
则输入上文中设置的目录 instantclient,/extensions，确保SDK文件夹也在此目录，稍等片刻后即可编译成功且自动加入php.ini

执行 php -m 可查看扩展是否载入成功

## 源码安装
`./configure  --with-oci8=instantclient,/Users/wangkeji/soft/oracle_ext --with-php-config=/usr/local/Cellar/php/7.4.6_1/bin/php-config`
php.ini添加so文件,重启即可

