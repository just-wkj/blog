# mac 下环境搭建

近期开始学习《UNIX高级环境编程》,看到第一张,代码里发现有apue.h文件的引入,编译不通过

```C
#include "apue.h"
#include <dirent.h>

int
main(int argc, char *argv[])
{
    DIR             *dp;
    struct dirent   *dirp;

    if (argc != 2)
        err_quit("usage: ls directory_name");

    if ((dp = opendir(argv[1])) == NULL)
        err_sys("can′t open %s", argv[1]);
    while ((dirp = readdir(dp)) != NULL)
        printf("%s\n", dirp->d_name);

    closedir(dp);
    exit(0);
}
```

其中apue.h头文件诗作者自定义的,需要去找到这个源码才可以.

打开网站<http://www.apuebook.com](https://link.jianshu.com/?t=http://www.apuebook.com>,找到版本对应的代码下载,

或者直接下载 <http://www.apuebook.com/src.3e.tar.gz>

下载完成后开始后续操作:

经Google后发现，下载的源码需要编译，并且把*aupe.h*和*error.h*两个文件复制到头文件目录下(我放到了*/usr/local/include/*)

### 编译源码

下载好源码之后,解压,进入apue.3e目录,执行编译操作

```shell
cd apue.3e && make
```



### 复制头文件到 `/usr/local/include`

```shell
cp ./include/apue.h /usr/local/include/
cp ./lib/error.c /usr/local/include/
# 编辑apue.h添加在最后一行**#end if之前**插入 #include "error.c"
vim /usr/local/include/apue.h

```

再次自行demo.编译

`cc myls.c` 或者 `clang myls.c` 编译成功

`./a.out /dev` 可以列出dev下的文件列表, 完美!

