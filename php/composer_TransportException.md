# composer TransportException报错

1. 错误展示![](https://ws1.sinaimg.cn/large/0063sFGSgy1g1f7eikef5j31a904x0t0.jpg)
2. 解决方案
- 可以先更新composer `composer self-update`,再次尝试
- 镜像设置中国 ` composer config -g repo.packagist composer https://packagist.phpcomposer.com` 
- secure-http 改成false `composer config -g secure-http false`
