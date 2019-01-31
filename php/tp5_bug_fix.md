# tp5漏洞修复
# thinkphp 漏洞修复

1.  报错信息关闭
      > 配置文件config.php中  
      > 1、show_error_msg 设置为  false  
      > 2、app_debug 设置为 false   
2. 控制器漏洞修复
    1. composer更新 
        > 如果你使用composer安装，并且一直保持最新版本使用的话，执行 composer update topthink/framework
    2. 手动修复
        - thinkphp5.0 版本  
            a. 需要修改 think\App类的module方法  553行左右 ![](https://ws1.sinaimg.cn/large/0063sFGSgy1fzpitawzz3j30pt0560t4.jpg)
             修复代码如下: 
            ```php
                // 获取控制器名
                   $controller = strip_tags($result[1] ?: $config['default_controller']);
                   //修复漏洞
                   if (!preg_match('/^[A-Za-z](\w|\.)*$/', $controller)) {
                       throw new HttpException(404, 'controller not exists:' . $controller);
                   }
            ```
        - thinkphp5.1 版本  
            a. 需要修改 think\route\dispatch\Url类的parseUrl方法 61行左右
            ![](https://ws1.sinaimg.cn/large/0063sFGSgy1fzpj5rntirj30o807caal.jpg)
            修复代码如下:

            ```php
              if ($this->param['auto_search']) {
                    $controller = $this->autoFindController($module, $path);
                } else {
                    // 解析控制器
                    $controller = !empty($path) ? array_shift($path) : null;
                }
                //修复漏洞
                if ($controller && !preg_match('/^[A-Za-z][\w|\.]*$/', $controller)) {
                    throw new HttpException(404, 'controller not exists:' . $controller);
                }
            ```
3. method漏洞修复
    > 需要修改 think\Request类的method方法  508行左右
    > ![](https://ws1.sinaimg.cn/large/0063sFGSgy1fzpjjtbuqpj30zs0f075p.jpg)
    > 修改为:
    > ![](https://ws1.sinaimg.cn/large/0063sFGSgy1fzpjkqtxqej31040ermyu.jpg)
    > 修复代码如下:   
    ```php
      //$this->method = strtoupper($_POST[Config::get('var_method')]);
       //$this->{$this->method}($_POST);
       //修复漏洞
       $method = strtoupper($_POST[Config::get('var_method')]);
       if (in_array($method, ['GET', 'POST', 'DELETE', 'PUT', 'PATCH'])) {
           $this->method = $method;
           $this->{$this->method}($_POST);
       } else {
           $this->method = 'POST';
       }
       unset($_POST[Config::get('var_method')]);
    ```
4. 漏洞修复参考地址
    - <https://blog.thinkphp.cn/869075>
    - <https://yq.aliyun.com/articles/686397>
