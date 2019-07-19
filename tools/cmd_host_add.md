# cmd host快速添加

测试机开发时,项目需要被产品运营或者测试访问查看,经常需要添加host来操作,
为了方便快速添加,可以写一个简单的bat脚本

```bat
attrib -R C:\WINDOWS\system32\drivers\etc\hosts 
@echo.>>C:\WINDOWS\system32\drivers\etc\hosts 
@echo 192.168.104.104 www.test.com >>C:\WINDOWS\system32\drivers\etc\hosts 
```
