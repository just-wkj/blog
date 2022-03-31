# windows微信多开

windows下使用bat进行微信多开

1. 右键微信快捷方式-属性
![](http://img.justwkj.com/20220331110315.png)
2. 快捷方式-目标 复制
![](http://img.justwkj.com/20220331110436.png)
得到如下 "E:\Program Files\Tencent\WeChat\WeChat.exe"
3. 新建 微信多开.bat,内容如下 (演示双开,如果需要打开更多,多复制几行即可)
```bash
start "1" "E:\Program Files\Tencent\WeChat\WeChat.exe"
start "2" "E:\Program Files\Tencent\WeChat\WeChat.exe"
```
4. 退出微信,点击 微信多开.bat 即可实现多开
![](http://img.justwkj.com/20220331110713.png)
