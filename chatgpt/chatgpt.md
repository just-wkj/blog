# chatgpt

本地直接调用openai接口经常会出现超时,并且自己写的页面也比较丑.
所以在github找到了一个开源项目 [chatgpt-web](https://github.com/Chanzhaoyu/chatgpt-web) ,然后自己改了一下,放到自己的服务器上.

项目地址: <http://ai.justwkj.com:3002/#/chat/1002>

```shell
docker run --name chatgpt-web -d -p 127.0.0.1:3032:3002 --env OPENAI_API_KEY=sk-1xxxxx chatgpt-web
```
