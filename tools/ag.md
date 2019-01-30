# ag
 ag号称比grep更快的一个搜索工具
 官网地址 [github](https://github.com/ggreer/the_silver_searcher)

1. mac 安装
   - `brew install the_silver_searcher`
2. centos 安装
    > `yum install -y pcre-devel`      
      `yum install -y xz-devel`  
      `cd /usr/local/src`  
      ` git clone https://github.com/ggreer/the_silver_searcher.git`  
      `cd the_silver_searcher`  
      `./build.sh && make install`  
3. 部分参数介绍
    - -A --after [LINES] Print lines after match (Default: 2)
    - -B --before [LINES] Print lines before match (Default: 2)
    - -C --context [LINES] Print lines before and after matches (Default: 2)
    - -g --filename-pattern PATTERN
    - -l --files-with-matches Only print filenames that contain matches
      (don't print the matching lines)
    - -L --files-without-matches
      Only print filenames that don't contain matches
    - -i --ignore-case Match case insensitively
    - -a --all-types Search all files (doesn't include hidden files
    or patterns from ignore files)
    - --depth NUM Search up to NUM directories deep (Default: 25)
    - -v --invert-match
4. 参考博客 [简书](https://www.jianshu.com/p/a6a373636894)
