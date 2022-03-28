# node版本管理

不同的前端项目对node版本需求不一样,一些老旧的项目依赖的node版本个很低.

今天遇到一个 [node-sass](https://github.com/sass/node-sass/releases) 版本报错的

由于本地的node版本是16,node-sass版本至少是V6.0.0,但是项目依赖的版本是v4.14.1
尝试把项目依赖改为v6.0.0依旧不行,包语法错误.所以只能切换node版本了

## 切换node版本
1. 安装node版本管理模块n  
`sudo npm install n -g`
   
2. 安装稳定版
`sudo n stable`

3. 安装最新版
`sudo n latest`
4. 安装指定版本,升级/降级  
`sudo n 版本号//例如：sudo n 14`


所以执行上述1和4步骤即可
全局安装模块n    
`sudo n 14` 之后就可以运营低版本的项目了
