# phpstorm 简单设置

自从phpstorm升级更新了之后**打开速度从5分钟到5秒钟**决定重新使用phpstorm来编码.以前都是editplus,以后会是两者结合了,主要看情况了.

基本快捷键自定义快捷键就不说了,来设置下文件头部和函数注释
- 文件头设置
    > settings->editor->file and code templates->include->PHP file Header
        系统的改成自己需要的  
        /**  
        * author:wkj  
        * createTime:${DATE} ${TIME}  
        * description:  
        */  
![请输入图片描述][1]

- 函数注释 写函数注释的时候希望加上日期,但是phpstorm并不支持,只能间接的完成了
>settings->editor->file and code templates->include->PHP Function Doc Comment
还是修改成自己所需的样式  
    /**  
    *  
    *@author:wkj  
    *@date    
    ${PARAM_DOC}  
    \#if (${TYPE_HINT} != "void") * @return  
    \#end
    */ 
![请输入图片描述][2]
这里的date 我个人喜欢在后面添加日期来标注这个函数是哪天写的,但是这样写是无效的,必须手动把日期写在后面!!这点太费事,因为比较懒所以想到另一个方式快捷的处理
如下:
settings->editor->file and code templates->include->Live Templates->PHP

![请输入图片描述][3]
![请输入图片描述][4]
![请输入图片描述][5]

###这些设置之后在date后面按下tab键即可输入当前时间
总体的效果图如下
![请输入图片描述][6]


  [1]: http://ww1.sinaimg.cn/large/0063sFGSgy1fds23pkqjmj30p60bb0ui.jpg
  [2]: http://ww1.sinaimg.cn/large/0063sFGSgy1fds24d7kmbj30u50fl40w.jpg
  [3]: http://ww1.sinaimg.cn/large/0063sFGSgy1fds1yhps5aj30uy0k3juz.jpg
  [4]: http://ww1.sinaimg.cn/large/0063sFGSgy1fds1zdk6e4j30o905jmxd.jpg
  [5]: http://ww1.sinaimg.cn/large/0063sFGSgy1fds210zfa3j30k80apjsb.jpg
  [6]: http://ww1.sinaimg.cn/large/0063sFGSgy1fds29hjlufj30fq0en0tg.jpg
