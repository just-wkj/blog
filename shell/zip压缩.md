#zip压缩

zip压缩排除指定目录

-x 参数 后面需要使用引号
````shell
zip -r advert_test.zip ./advert_test/ -x "advert_test/basic/images/attachment/*"
````