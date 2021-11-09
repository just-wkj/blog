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

![](http://img.justwkj.com/20211109101343.png)
```HTML
<STYLE>
    .DEMO0 {
        MARGIN: 0 AUTO;
        WIDTH: 700PX;
        HEIGHT: 30%;
        BACKGROUND: CADETBLUE;
        DISPLAY: FLEX;
        /*主轴默认*/
        FLEX-DIRECTION: ROW;
        JUSTIFY-CONTENT: CENTER;
        ALIGN-ITEMS: CENTER;
    }

    .DEMO0 SPAN {
        WIDTH: 100PX;
        HEIGHT: 100PX;
        BACKGROUND: PINK;
        MARGIN-LEFT: 10PX;
    }

    .DEMO1 {
        MARGIN: 0 AUTO;
        WIDTH: 700PX;
        HEIGHT: 30%;
        BACKGROUND: CADETBLUE;
        DISPLAY: FLEX;
        /*主轴默认*/
        FLEX-DIRECTION: ROW;
        /*FLEX-DIRECTION: ROW-REVERSE;*/
        /*FLEX-DIRECTION: COLUMN;*/
        /*FLEX-DIRECTION: COLUMN-REVERSE;*/

        JUSTIFY-CONTENT: FLEX-START;
        /*//FLEX 默认不换行,如果放不下则缩小子元素宽度*/
        FLEX-WRAP: WRAP;
        /*ALIGN-ITEMS: CENTER;*/

        /*有了换行需要使用ALIGN-CONTENT*/
        ALIGN-CONTENT: SPACE-AROUND;
    }

    .DEMO1 SPAN {
        WIDTH: 200PX;
        HEIGHT: 100PX;
        BACKGROUND: PINK;
        MARGIN-LEFT: 10PX;
    }

    .DEMO2 {
        WIDTH: 80%;
        HEIGHT: 150PX;
        MARGIN: 0 AUTO;
        BACKGROUND: #CCC;
        DISPLAY: FLEX;
    }

    .DEMO2 SPAN:NTH-CHILD(1) {
        BACKGROUND: RED;
        WIDTH: 150PX;
    }

    .DEMO2 SPAN:NTH-CHILD(2) {
        BACKGROUND: YELLOWGREEN;
        /*//剩余空间分配*/
        FLEX: 1;
    }

    .DEMO2 SPAN:NTH-CHILD(3) {
        BACKGROUND: GREEN;
        WIDTH: 150PX;
    }

    .DEMO3 {
        WIDTH: 80%;
        HEIGHT: 150PX;
        MARGIN: 0 AUTO;
        BACKGROUND: #CCC;
        DISPLAY: FLEX;
    }

    .DEMO3 SPAN {
        /*子元素不设置宽度,剩余空间是全部,都设置FLEX1,则表示平均分配*/
        FLEX: 1;
        BORDER: 1PX SOLID RED;
    }

    .DEMO4 {
        WIDTH: 80%;
        HEIGHT: 150PX;
        MARGIN: 0 AUTO;
        BACKGROUND: #CCC;
        DISPLAY: FLEX;
    }

    .DEMO4 SPAN:NTH-CHILD(1) {
        FLEX: 1;
        BORDER: 1PX SOLID RED;
    }

    .DEMO4 SPAN:NTH-CHILD(2) {
        FLEX: 2;
        BORDER: 1PX SOLID RED;
    }

    .DEMO4 SPAN:NTH-CHILD(3) {
        FLEX: 4;
        BORDER: 1PX SOLID RED;
    }

    .DEMO4 SPAN:NTH-CHILD(4) {
        WIDTH: 100PX;
        BORDER: 1PX SOLID RED;
    }

    .DEMO5 {
        WIDTH: 80%;
        HEIGHT: 300PX;
        MARGIN: 0 AUTO;
        BACKGROUND: #CCC;
        DISPLAY: FLEX;
        JUSTIFY-CONTENT: CENTER;
        ALIGN-ITEMS: FLEX-END;
    }

    .DEMO5 SPAN {
        WIDTH: 100PX;
        HEIGHT: 100PX;
        BORDER: 1PX SOLID RED;
    }

    .DEMO5 SPAN:NTH-CHILD(2) {
        ALIGN-SELF: FLEX-START;
    }

    .DEMO5 SPAN:NTH-CHILD(4) {
        /*排序默认0,越小越靠左*/
        ORDER: -1;
    }
</STYLE>
<P>1. 简单DEMO对齐</P>
<DIV CLASS="DEMO0">
    <SPAN>1</SPAN>
    <SPAN>2</SPAN>
    <SPAN>3</SPAN>
</DIV>

<P>1. 简单DEMO</P>
<DIV CLASS="DEMO1">
    <SPAN>1</SPAN>
    <SPAN>2</SPAN>
    <SPAN>3</SPAN>
    <SPAN>4</SPAN>
    <SPAN>5</SPAN>
</DIV>

<P>2.两侧固定,中间占满</P>
<P CLASS="DEMO2">
    <SPAN>1</SPAN>
    <SPAN>2</SPAN>
    <SPAN>3</SPAN>
</P>

<P>3.平均三等分</P>
<P CLASS="DEMO3">
    <SPAN>1</SPAN>
    <SPAN>2</SPAN>
    <SPAN>3</SPAN>
</P>
<P>4.自定义大小切分</P>
<P CLASS="DEMO4">
    <SPAN>1</SPAN>
    <SPAN>2</SPAN>
    <SPAN>3</SPAN>
    <SPAN>4</SPAN>
</P>

<P>5.子元素单独排列</P>
<P CLASS="DEMO5">
    <SPAN>1</SPAN>
    <SPAN>2</SPAN>
    <SPAN>3</SPAN>
    <SPAN>4</SPAN>
</P>


```
