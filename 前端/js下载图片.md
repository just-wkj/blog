#js下载图片

```js


// <div onclick="aa()"> 下载图片</div>

<script>
    function downFile (res, name){
        var aLink = document.createElement('a');
        aLink.href = res;
        aLink.download = name|| 'qrcode.jpg';
        aLink.style.display = 'none';
        document.body.appendChild(aLink);
        aLink.click();
        document.body.removeChild(aLink);
    };


    // 只能下载同一个域名的图片,二级域名不一样不可以
    function aa(){
        miniUrl = '/piao.png' //同一域名的图片
        let fileName =  miniUrl.substring(miniUrl.lastIndexOf('/') + 1)
        downFile(miniUrl, fileName)
    }


    //可以下载不同域名的数据
    funciton downloadImgByBase64(url){
    let img = new Image()

    img.onload = function() {
    let canvas = document.createElement('canvas')
    canvas.width = img.width
    canvas.height = img.height
    let ctx = canvas.getContext('2d')
    // 将img中的内容画到画布上
    ctx.drawImage(img, 0, 0, canvas.width, canvas.height)
    // 将画布内容转换为base64
    let base64 = canvas.toDataURL()
    // 创建a链接
    let a = document.createElement('a')
    a.href = base64
    a.download = ''
    // 触发a链接点击事件，浏览器开始下载文件
    a.click()
}

    img.src = url
    // 必须设置，否则canvas中的内容无法转换为base64
    img.setAttribute('crossOrigin', 'Anonymous')
}



</script>



```
如果是用第二种还是报跨域的错误
可以在nginx添加header
```nginx
    add_header Access-Control-Allow-Origin *;
	add_header Access-Control-Allow-Headers X-Requested-With;
	add_header Access-Control-Allow-Methods GET,POST,OPTIONS;



server {
    listen       8082;
   
	server_name localhost;
	
	add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;
	add_header Access-Control-Allow-Origin *;
	add_header Access-Control-Allow-Headers X-Requested-With;
	add_header Access-Control-Allow-Methods GET,POST,OPTIONS;
}

```