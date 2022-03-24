# vue引入原生高德地图

最近项目需要引入高德地图组件,于是查看文档很快写了一个html+js的demo. 
但是由于项目是vue开发的,所以需要改造一下. 之后google搜索了很久也没搜到好的方案,好多推荐vue-map,
但是看到vue-map文档都显示不全,觉得使用起来堪忧.

最终在github上找到一个demo,根绝这个demo进行改在
[github vue-gaode-js](https://github.com/fenglingcong/vue-gaode-js)

下面是改造方案

1. index.html页面直接引入高德js,本来是想和上面参考的一样使用在加载,但是由于使用到AMapUI会报错,所以干脆全都引入

```html
<!DOCTYPE html>
<html lang="zh-cmn-Hans">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width,initial-scale=1.0">
    <title>系统</title>
    <link rel="icon" href="<%= BASE_URL %>logo.png">
    <script src="<%= BASE_URL %>cdn/babel-polyfill/polyfill_7_2_5.js"></script>
    <script type="text/javascript">
        window._AMapSecurityConfig = {
            securityJsCode: '你的高德秘钥',
        }
    </script>
    <script type="text/javascript" src="//webapi.amap.com/maps?v=2.0&key=你的高德key&plugin=AMap.ToolBar"></script>
    <!-- UI组件库 1.0 -->
    <script src="//webapi.amap.com/ui/1.1/main.js?v=1.1.1"></script>
    <!-- 全局配置 -->
    <script>
        window._CONFIG = {};
        /*  window._CONFIG = ['domianURL']='http://127.0.0.1:8080/jeecg-boot';*/
    </script>
</head>
<body>
<div id="app"></div>
</body>
</html>
```

2. 创建组件components/amap/amap.vue
```html
<template>
  <div class='app-inner'>
    <a-row>
      <a-col span="4">
        <div id="pickerBox">
          <a-input id="pickerInput" placeholder="输入关键字选取地点" />
          <div id="poiInfo"></div>
        </div>
      </a-col>
      <a-col span="20">
        <div>
          经纬: {{lnglat && lnglat.lng || '--'}}
          维度: {{lnglat && lnglat.lat || '--'}}
          地址：{{address || '--'}}
        </div>
      </a-col>
    </a-row>
    <div id="container"></div>


    <div v-if="false">
      <hr>
      <h3>生成静态地图图片</h3>
      <img class="map-img" :src="mapUrl" alt="map">
    </div>
  </div>
</template>

<script>


export default {
  name: 'Map',
  data() {
    return {
      map: null,
      marker: null,
      geocoder: null,
      autoComplete: null,
      address: '',
      lnglat: null,
      showSelect: false,
      resultList: [],
    };
  },
  computed: {
    mapUrl() {
      return  '';
      // const lnglatStr = this.lnglat ? `${this.lnglat.lng},${this.lnglat.lat}` : '116.407526,39.90403';
      // const url = `https://restapi.amap.com/v3/staticmap?location=${lnglatStr}&zoom=13&size=750*300&markers=mid,,A:${lnglatStr}&key=8461bc5da657e97a65d0065888a5bbcb`;
      // return url;
    },
  },
  created() {},
  mounted () {
    this.initMap();
  },
  methods: {
    initMap(){
      const that = this
      AMapUI.loadUI(['misc/PositionPicker', 'misc/PoiPicker'], function(PositionPicker, PoiPicker) {
        var poiPicker = new PoiPicker({
          //city:'北京',
          input: 'pickerInput'
        });

        window.poiPicker = poiPicker;

        var marker = new AMap.Marker();
        var infoWindow = new AMap.InfoWindow({
          offset: new AMap.Pixel(0, -20)
        });

        //选取了某个POI
        poiPicker.on('poiPicked', function (poiResult) {
          marker.setPosition(poiResult.item.location);
          map.setCenter(marker.getPosition());
        });
        //初始化poiPicker
        // poiPickerReady(poiPicker);
        var map = new AMap.Map('container', {
          zoom: 16,
          scrollWheel: true
        })
        var positionPicker = new PositionPicker({
          mode: 'dragMap',
          // mode: 'dragMarker',
          map: map
        });

        positionPicker.on('success', function(positionResult) {
          var _position = positionResult.position
          that.lnglat = {
            lng:_position.lng,
            lat:_position.lat
          }
          that.address = positionResult.address;
        });
        positionPicker.on('fail', function(positionResult) {

        });
        positionPicker.start();
        map.panBy(0, 1);

        map.addControl(new AMap.ToolBar({
          liteStyle: true
        }))
      });
    }
  },
};
</script>

<style lang='less' scoped>
.app-inner {
  text-align: center;
}

#pickerInput2{
  width: 220px;
  border: 1px solid #ccc;
}
#container {
  margin: 30px auto;
  width: 100%;
  height: 400px;;
}

.address_info{
  position: absolute;
  top: 100px;
  right: 20%;
}

.input {
  &-wrap {
    margin: 20px 0;
  }

  &-box {
    position: relative;
    display: inline-block;
  }

  &-text {
    width: 300px;
    height: 30px;
  }

  &-btn {
    padding: 0 10px;
    height: 36px;
  }
  &-result {

    &__list {
      padding-bottom: 10px;
      position: absolute;
      top: 30px;
      left: 0;
      z-index: 99;
      width: 100%;
      background-color: #fff;
      box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    }

    &__item {
      padding: 0 10px;
      width: 100%;
      line-height: 30px;
      text-align: left;
      border-bottom: 1px solid #eee;

      &:last-child {
        border-bottom: 0;
      }
    }
  }
}

.map-img {
  margin: 0 0 30px;
}
</style>
```

3. 页面使用Map组件

```html
<template>
  <a-modal v-model="visible" title="地理位置" width="80%" height="80%" @ok="handleOk">
    <Map />
  </a-modal>
</template>

<script>
  import Vue from 'vue'
  import Map from '@/components/amap/Map.vue';

  export default {
    name: "MapSelectModal",
    components: {
      Map
    },
    data() {
      return {
        visible: false,
        }
      }
    },
  
    methods: {
      showModal() {
        this.visible = true;
      },
      handleOk(e) {
        this.visible = false;
        this.$emit('ok');
      },
    }
  }
</script>

<style scoped>
.amap-wrapper {
  width: 500px;
  height: 500px;
}

#container {
  width: 100%;
  height: 100%;
  margin: 0px;
  font-size: 13px;
}

#pickerBox {
  position: absolute;
  z-index: 9999;
  top: 50px;
  right: 30px;
  width: 300px;
}

#pickerInput {
  width: 200px;
  padding: 5px 5px;
}

#poiInfo {
  background: #fff;
}

.amap_lib_placeSearch .poibox.highlight {
  background-color: #CAE1FF;
}

.amap_lib_placeSearch .poi-more {
  display: none !important;
}
</style>

```
这样就可以使用了
![](http://img.justwkj.com/20220324150213.png)
