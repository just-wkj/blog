# react动态添加input

```js

channels = [''] // channels格式

// onInputChange.bind(this,index) 需要使用 xx.bind(this,params)才能触发绑定事件
{channels.map((item, index) => (
    <Row style={{marginTop:10}}>
        <Col span={4} style={{lineHeight:'30px'}}>渠道名称:</Col>
        <Col span={20}>
            <Input placeholder="请输入发送二维码的渠道名称,最多十个字" value={item}
                   onChange={props.onInputChange.bind(this,index)}
                   maxLength={10}
            />
        </Col>
    </Row>
))}

// input修改
onInputChange = {(index, e) => {
    dispatch('shop:changeChannel', {index, value:e.target.value});
}}

// input修改绑定
export const changeChannel = action("shop:changeChannel", (store, {index, value}) => {
    store.setState((state) => {
        state.channels[index] = value;
        let newChannel =  state.channels;
        state.channels = newChannel;
    })
});

```
相关参考: 
- [react 动态添加input框后获取值的问题](https://segmentfault.com/q/1010000017986817)  
- [react动态添加数据](https://blog.csdn.net/wgf5845201314/article/details/88675338)  
