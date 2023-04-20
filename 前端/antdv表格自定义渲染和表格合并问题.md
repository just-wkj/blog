# antdv表格自定义渲染和表格合并问题

### 问题描述: 
> 使用antdv的表格组件, 需要自定义渲染某一列, 并且需要合并单元格, 但是合并单元格的时候, 会导致自定义渲染的内容不显示.  
 同时也看到一个人提出了同样的问题,下面的解决方案亲测可用  
[segmentfault同问题解决方案](https://segmentfault.com/q/1010000023578435)

### 解决方案:
> 使用`customCell`配合`scopedSlots`来实现自定义渲染和合并单元格

```javascript
const columns = [
    {
        width: 120,
        align: 'center',
        title: '区域',
        dataIndex: 'quyu',
        scopedSlots: {
            customRender: 'quyu'
        },
        customCell: (record, index) => {
            const obj = {
                style: {},
                attrs: {
                    rowSpan: 0
                }
            }
            if (parseInt(record.data_source) === 1) {
                obj.attrs.rowSpan = 2
            } else {
                obj.style.display = 'none'
            }
            return obj
        },
        fixed: 'left'
    },
    {
        width: 100,
        align: 'center',
        title: '数据来源',
        dataIndex: 'shujulaiyuan',
        scopedSlots: { customRender: 'shujulaiyuan' },
        fixed: 'left'
    },
   {
        title: '操作',
        align: 'center',
        dataIndex: 'action',
        width: 80,
        fixed: 'right',
        scopedSlots: { customRender: 'operat' },
        customCell: (record, index) => {
          const obj = {
            style: {},
            attrs: {
              rowSpan: 0
            }
          }
          if (parseInt(record.data_source) === 1) {
            obj.attrs.rowSpan = 2
          } else {
            obj.style.display = 'none'
          }
          return obj
        }
      }
```
