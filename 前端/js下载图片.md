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


    function aa(){
        miniUrl = '/piao.png' //同一域名的图片
        let fileName =  miniUrl.substring(miniUrl.lastIndexOf('/') + 1)
        downFile(miniUrl, fileName)
    }
</script>



```