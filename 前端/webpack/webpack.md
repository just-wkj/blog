# webpack
###### 日期:2019-03-24


最近对前端兴趣比较浓厚,打算研究研究前端的技术,为了学习一下react,打算先简单的研究webpack  
以前对webpack也是是简单的理解为打包工具,最近看了下webpack的东西打算记录下来

## 如何手动创建一个webpack项目
1. 创建文件夹`mkdir webpack-test && cd webpack-test`, 项目快速初始化 `npm init -y`
2. 创建目录 `mkdir {src,dist}` 
3. 创建文件 `touch src/{index.html,index.js}`
4. 安装webpck `npm i webpack -D` 或者使用cnpm, 解释一下**-D是开发依赖,相当于--save-dev**
5. 安装webpck-cli `npm i webpack-cli -D` 或者使用cnpm, 可以直接合并成  `npm i webpack webpack-cli -D`
6. 执行  `webpack` webpack4.0 约定大约配置 入口文件是 src/index.js 输出是dist/main.js
7. 安装webpck-dev-server `cnpm i webpack-dev-server -D`,可以在package.json的scrpts中添加自定义命令比如dev
`"dev": "webpack-dev-server --open --port 3001 --hot --progress --compress --host 127.0.0.1"`,此处参数简单的解释一下,
> --open 打开浏览器,不传参数打开默认浏览器,可以指定eg:--open fixfox  
> --port 指定端口  
> --hot 热更新  
> --progress 打包进程  
> --compress 压缩打包的文件  
> --host 指定域名,默认是localhost