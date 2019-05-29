# composer TransportException报错

1. 错误展示![](http://img.justwkj.com/20190529091619.png)
2. 解决方案
- 可以先更新composer `composer self-update`,再次尝试
- 镜像设置中国 ` composer config -g repo.packagist composer https://packagist.phpcomposer.com` 
- secure-http 改成false `composer config -g secure-http false`
- 修改composer.json 文件 
```json
{
    "repositories": {
        "packagist": {
            "type": "composer",
            "url": "https://packagist.phpcomposer.com"
        }
    }
}
```
