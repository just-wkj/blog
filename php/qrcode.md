# php二维码单文件

```php
<?php
/**
 * @author: justwkj
 * @date: 2020/9/3 23:17
 * @email: justwkj@gmail.com
 * @desc:
 */

namespace common\utils;


class Qrcode {

    /**
     *  生成二维码
     * @param $url
     * @return string
     * @author: justwkj
     * @date: 2020/9/3 23:17
     */
    public static function createOld($url) {
        return 'http://qr.topscan.com/api.php?text=' . $url;
        //return 'https://chart.apis.google.com/chart?cht=qr&chs=300x300&chl='.$url;
    }

    /**
     *  类库生成图片
     * @param $url
     * @return string
     * @author: justwkj
     * @date: 2020/9/8 21:24
     */
    public static function create($url='') {
        //生成二维码图片
        $path = __DIR__.'/../../frontend/web/qrcode/';
        if(!is_dir($path)){
            mkdir($path);
        }
        $filename =  md5($url) . '.png';//生成二维码图片
        $QR = $path . $filename;            //已经生成的原始二维码图

        //已经生成的图片不再生成,直接读取即可
        if(file_exists($QR)){
            return "http://".$_SERVER['HTTP_HOST']."/qrcode/".$filename;
        }


        require_once __DIR__."/../extend/phpqrcode.php";

        $value = $url;                  //二维码内容
        $errorCorrectionLevel = 'H';    //容错级别
        $matrixPointSize = 10;           //生成图片大小


        \common\extend\QRcode::png($value, $path . $filename, $errorCorrectionLevel, $matrixPointSize, 2);

        $QR = imagecreatefromstring(file_get_contents($QR));        //目标图象连接资源。

        imagepng($QR,$path.$filename );//带Logo二维码
        return "http://".$_SERVER['HTTP_HOST']."/qrcode/".$filename;
    }
}


```