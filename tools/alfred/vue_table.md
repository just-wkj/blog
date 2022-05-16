# alfred 自动生成vue table column

![](http://img.justwkj.com/20220516135218.png)
```php

<?php
/**
 * @author: justwkj
 * @date: 2022/5/16 11:18
 * @email: justwkj@gmail.com
 * @desc: 根据数据库生成vue表字段
 */

function outputJson($data) {
    $itemArr = [];
    foreach ($data as $datum) {
        $tmp = [
            'title' => $datum['title'],
            'arg' => isset($datum['arg']) ? $datum['arg'] : '',
        ];
        if (isset($datum['quicklookurl'])) {
            $tmp['quicklookurl'] = $datum['quicklookurl'];
        }
        $itemArr[] = $tmp;
    }
    $data = ['items' => $itemArr];
    echo json_encode($data);
    die;
}

$input = shell_exec('pbpaste');


$align = 'center';//空字符串,或者center,left,right
$action = true;//

preg_match_all("/\s+`(.*)?`.*?COMMENT\s+'(.*)?',/", $input, $matches);

$vueColumn = [];
if (count($matches) != 3 || empty($matches[1]) || empty($matches[2])) {
    outputJson([[
        "title" => '请复制建表语句(包含注释)',
        "arg" => '',
    ]]);

}
foreach ($matches[1] as $index => $column) {
    $tmp = [
        'title' => $matches[2][$index],
        'dataIndex' => $column,
    ];
    if ($align) {
        $tmp['align'] = $align;
    }

    $vueColumn[] = $tmp;
}

if ($action) {
    $vueColumn[] = [
        'title' => '操作',
        'dataIndex' => 'action',
        'width' => '180px',
        'scopedSlots' => [
            'customRender' => 'action',
        ],
    ];
}

$str = json_encode($vueColumn, JSON_PRETTY_PRINT|JSON_UNESCAPED_UNICODE);
outputJson([[
    "title" => '格式化完毕:' . $str,
    "arg" => $str,
]]);

```
