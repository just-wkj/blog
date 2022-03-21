# 打包构建报错

## - vue项目打包报错
项目安装第三方组件,某些时候可能会无法构建项目,运行 `npm run build`会报错
但是运行 `npm run serve`正常.

报错如下
>  vue-antd-jeecg@3.0.0 build
 vue-cli-service build
⠋  Building for production... ERROR  TypeError: Class extends value undefined is not a constructor or null
TypeError: Class extends value undefined is not a constructor or null
at Object.<anonymous> (/Users/wangkeji/code/dfy/dfy_gitrepo/dfy_crm/ant-design-vue-jeecg/node_modules/mini-css-extract-plugin/dist/CssDependency.js:12:46)
at Module._compile (node:internal/modules/cjs/loader:1101:14)
at Object.Module._extensions..js (node:internal/modules/cjs/loader:1153:10)
at Module.load (node:internal/modules/cjs/loader:981:32)
at Function.Module._load (node:internal/modules/cjs/loader:822:12)
at Module.require (node:internal/modules/cjs/loader:1005:19)
at require (node:internal/modules/cjs/helpers:102:18)
at Object.<anonymous> (/Users/wangkeji/code/dfy/dfy_gitrepo/dfy_crm/ant-design-vue-jeecg/node_modules/mini-css-extract-plugin/dist/index.js:12:45)
at Module._compile (node:internal/modules/cjs/loader:1101:14)
at Object.Module._extensions..js (node:internal/modules/cjs/loader:1153:10)
at Module.load (node:internal/modules/cjs/loader:981:32)
at Function.Module._load (node:internal/modules/cjs/loader:822:12)
at Module.require (node:internal/modules/cjs/loader:1005:19)
at require (node:internal/modules/cjs/helpers:102:18)
at Object.<anonymous> (/Users/wangkeji/code/dfy/dfy_gitrepo/dfy_crm/ant-design-vue-jeecg/node_modules/mini-css-extract-plugin/dist/cjs.js:3:18)
at Module._compile (node:internal/modules/cjs/loader:1101:14)  

![](http://img.justwkj.com/20220321155858.png)

## - 解决方案
1. 删除 node_modules 和 lock文件
2. 执行 `npm install` 和 `npm run build` 
