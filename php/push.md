# 企业微信钉钉等消息推送

由于经常需要用到推送提醒的功能,之前就用了server酱的推送. 由于server酱后期免费条数过少,
所以寻找替代品.  
替代品找到 微信菌 vercel [github地址](https://github.com/zhheo/work-weixin-msg-sever-api)
试用了一下发现就是调用企业微信接口,延迟10来秒太慢了,自己写一个

上代码

```php
<?php


Notify::vercel('短信提醒vercel');
Notify::weixin('短信提醒wecha');
Notify::dingTalk('钉钉');

class NotifyConfig {
    // 企业微信配置 https://work.weixin.qq.com/
    const WX_ID = 'wwe4eb****企业微信ID';
    const WX_SECRET = '企业微信秘钥';
    const WX_AGENT_ID = 1000002;//应用的id根据实际情况填写

    //钉钉配置 钉钉需要设置机器人
    const DING_TALK_ACCESS_TOKEN = '钉钉的access_token';
}

class Notify {
    public static function vercel($msg) {
        $url = vsprintf('https://work-weixin-msg-sever-api.vercel.app/api?id=%s&secert=%s&agentId=%s&msg=%s', [
            NotifyConfig::WX_ID,
            NotifyConfig::WX_SECRET,
            NotifyConfig::WX_AGENT_ID,
            $msg
        ]);
        return self::curlGet($url);
    }

    public static function weixin($msg = '') {
        $id = NotifyConfig::WX_ID;
        $secret = NotifyConfig::WX_SECRET;
        $agentId = NotifyConfig::WX_AGENT_ID;
        $url = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" . $id . "&corpsecret=" . $secret;

        $res = json_decode(self::curlPostJson($url), true);
        if (!empty($res['access_token'])) {
            $sendUrl = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" . $res['access_token'];
            $res = self::curlPostJson($sendUrl, [
                "safe" => 0,
                "touser" => '@all',
                "msgtype" => 'text',
                "agentid" => $agentId,
                "text" => [
                    "content" => $msg
                ]
            ]);
            return $res;
        }
        return 0;
    }

    public static function dingTalk($text = '', $phoneArr = []) {
        $webhook = "https://oapi.dingtalk.com/robot/send?access_token=" . NotifyConfig::DING_TALK_ACCESS_TOKEN;
        $message = "监控报警:" . $text;
        $data = [
            'msgtype' => 'text',
            'text' => ['content' => $message],
            'at' => [
                "atMobiles" => $phoneArr, //["18252022430"],
                "isAtAll" => false, // 不@用户
            ],

        ];
        $res = self::curlPostJson($webhook, $data);
        return $res;
    }

    /**
     *  post json
     * @param $url
     * @param array $post
     * @param array $headers
     * @param int $timeOut
     * @return bool|string
     * @author: justwkj
     * @date: 2020/3/31 09:03
     */
    public static function curlPostJson($url, array $post = array(), $headers = array('Content-Type' => 'application/json;charset=utf-8'), $timeOut = 10) {
        $defaults = array(
            CURLOPT_POST => 1,
            CURLOPT_HEADER => 0,
            CURLOPT_URL => $url,
            CURLOPT_FRESH_CONNECT => 1,
            CURLOPT_RETURNTRANSFER => 1,
            CURLOPT_FORBID_REUSE => 1,
            CURLOPT_TIMEOUT => $timeOut,
            CURLOPT_POSTFIELDS => json_encode($post)
        );

        $ch = curl_init();
        curl_setopt_array($ch, $defaults);
        if ($headers) {
            $_header = array();
            foreach ($headers as $key => $vo) {
                $_header[] = $key . ':' . $vo;
            }
            curl_setopt($ch, CURLOPT_HTTPHEADER, $_header);
        }
        if (stripos($url, "https://") !== false) {
            curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);
            curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, false);
            curl_setopt($ch, CURLOPT_SSLVERSION, 1);
        }
        if (!$result = curl_exec($ch)) {
            trigger_error(curl_error($ch));
        }
        curl_close($ch);

        return $result;
    }

    /**
     *  curl get
     * @param $url
     * @param array $params
     * @param array $headers
     * @param int $timeOut
     * @param bool $isFollow
     * @return bool|string
     * @author: justwkj
     * @date: 2019-09-16 13:48
     */
    public static function curlGet($url, $params = [], $headers = [], $timeOut = 10, $isFollow = false) {
        $paramsStr = '';
        if ($params) {
            if (strpos($url, '?') === false) {
                $paramsStr .= '?';
            } else {
                $paramsStr .= '&';
            }
            $paramsStr .= http_build_query($params);
        }

        $url = rtrim($url . $paramsStr, '&');
        $ch = curl_init();
        if (stripos($url, "https://") !== false) {
            curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);
            curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, false);
            curl_setopt($ch, CURLOPT_SSLVERSION, 1);
        }
        curl_setopt($ch, CURLOPT_URL, $url);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
        curl_setopt($ch, CURLOPT_TIMEOUT, $timeOut);
        if ($isFollow) {
            curl_setopt($ch, CURLOPT_FOLLOWLOCATION, true);
        }
        if ($headers) {
            $_header = [];
            foreach ($headers as $key => $vo) {
                $_header[] = $key . ':' . $vo;
            }
            curl_setopt($ch, CURLOPT_HTTPHEADER, $_header);
        }
        $sContent = curl_exec($ch);
        $aStatus = curl_getinfo($ch);
        curl_close($ch);

        if (intval($aStatus["http_code"]) == 200) {
            return $sContent;
        } else {
            return false;
        }
    }

}

```