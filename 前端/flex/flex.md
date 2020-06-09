# flex

作为一个后端开发,还是很怕前端页面排版的,有时候可能比较顺,有时候怎么也调整不好.
以前听说过 flex 不过没怎么研究过. 今天无意中看到一个关于flex的视频 [flex简介](https://www.youtube.com/watch?v=CB-s9s1-r7Q)  |
[flex响应式布局](https://www.youtube.com/watch?v=94EPx1NMMjU&t=5s) ,讲的比较浅显,
比较适合后端开发者看. 


下面是简单的demo:

```html

<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
<style>
    nav {
        display: flex;
    }

    nav .logo {
        /*没有弹性固定宽度*/
        flex: none;
        width: 100px;

        background-color: red;
    }

    nav .menu {
        /*尽量把容器的空间张满*/
        flex: auto;
        
        background-color: blue;
    }

    nav .user {
        flex: none;
        width: 100px;
        background-color: green;
    }


    main {
        display: flex;
        flex-wrap: wrap; /* 需要响应式布局的时候设置 */
        justify-content: center; /* flex子元素居中 flex-start center flex-end */
        align-items: center; /*flex子元素 垂直居中对齐  flex-start center flex-end*/
    }

    main > .item {
        flex: none; /* flex:none; 代表固定不可以动态变化,必须设置固定的宽度 */
        width: 270px;
        margin: 10px;
        background-color: beige;
    }


    /*规划:*/
    /*     >=1200px 4栏位 固定尺寸*/
    /*     500px~1200px 2栏位 弹性宽度*/
    /*     <500px 1栏位 弹性尺寸*/

    @media (max-width: 1200px) {
        main > .item {
            width: 45%;
        }
    }

    @media (max-width: 500px) {
        main > .item {
            width: 95%;
        }
    }

</style>


<nav>
    <div class="logo">logo</div>
    <div class="menu">menu</div>
    <div class="user">user</div>
</nav>

<!--
    规划:
        >=1200px 4栏位 固定尺寸
        500px~1200px 2栏位 弹性宽度
        <500px 1栏位 弹性尺寸
-->
<main>
    <div class="item">new1 <br/ > today is good day</div>
    <div class="item">new2</div>
    <div class="item">new3</div>
    <div class="item">new4</div>
</main>

</body>
</html>

```
