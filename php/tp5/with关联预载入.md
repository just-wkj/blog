# tp5 with关联预载入

以前觉得tp with使用限制很大,后来发现是研究不够

举个栗子: 获取关联数据的时候需要排序

```php
    //imgUrl获取图片需要进行排序操作,直接使用无法实现
    self:with(['imgs.imgUrl', 'properties'])->find();
    
    ...转化一下
    self:with(['imgs.imgUrl'])->with(['properties'])->find();
    
    ...再次转化
    self:with([
        'imgs' => function($query){
            $querey->with(['imgUrl'])->order('sort', 'desc');
        }
    ])->with(['properties'])->find();
```
