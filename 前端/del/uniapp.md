# uniapp
```js

import request from '@/util/request';
import apiMap from '@/util/apimap';
import NavList from "@/components/nav-list";


let pageNum = 0;
let totalCount = 0;

export default {
  components: { NavList },
  data () {
    return {
      stateList: [
        {
          label: '可使用',
          value: 1,
        },
        {
          label: '已使用',
          value: 2,
        },
        {
          label: '已失效',
          value: 3,
        }
      ],
      state: 1,
      dataList: [{}, {}],
    }
  },
  onLoad () {
    pageNum = 0;
    totalCount = 0;
    this.queryList();
  },
  onPullDownRefresh: function (e) {
    this.refresh();
  },
  onReachBottom: function (e) {
    this.loadMore();
  },
  methods: {
    // 刷新
    refresh () {
      pageNum = 0;
      this.queryList();
    },

    // 加载更多
    loadMore () {
      if (this.dataList.length >= totalCount) return;
      pageNum += 1;
      this.queryList(true);
    },

    // 查询列表接口
    queryList (loadMore = false) {
      const params = {
        page:pageNum,
        status: this.state,
      }
      uni.showLoading({
        title: '加载中...',
        mask: true,
      })
      request.fetch(apiMap.userCouponList, params)
        .then((data) => {
          uni.hideLoading();
          uni.stopPullDownRefresh();
          console.log(data)
          totalCount = data.count;
          this.dataList = loadMore ? this.dataList.concat(data.lists) : data.lists;
        }).catch(() => {
          uni.hideLoading();
          uni.stopPullDownRefresh();
        })
    },

    navChange (v) {
      this.state = v
      console.log('nav change')
      console.log(this.state)
      this.queryList();
    },
    _detail (id) {
      uni.navigateTo({
        url: `/pages/order/detail/detail?id=${id}`
      });
    }
  }
}

```