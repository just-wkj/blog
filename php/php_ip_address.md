# php 根据ip获取省份信息



```php
<?php  
namespace app\common\util;  
class IpLocationUtil {  
  
    const IP_URL = 'https://www.ipshudi.com/%s.htm';  
  
    /**  
     *  ip 查询省市区  
     * @param $ip  
     * @return array  
     * @author: justwkj  
     * @date: 2023/12/3 10:35  
     *  //根据https://www.ip138.com/ 点击查询获取  
     */  
    public static function getLocation($ip) {  
        $html = Curl::get(sprintf(self::IP_URL, $ip));  
        $pattern = '/<span>中国\s(.*?)<a.*?>(.*?)<\/a>(.*?)<\/span>/';  
        preg_match($pattern, $html, $matches);  
        $province = $city = '';  
        if ($matches && count($matches) == 4) {  
            $province = $matches[1];  
            $city = $matches[2];  
            if (empty($province)) { //比如上海市,只返回城市没有省份  
                $province = $city;  
            }  
            $province = trim($province);  
            $city = trim($city);  
        }  
        return [  
            'province' => $province,  
            'city' => $city,  
            'address' =>implode('', array_unique([$province, $city]))  
        ];  
    }  
}
```
