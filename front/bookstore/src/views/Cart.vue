<template>
  <div>
    <h3>cart</h3>
    <Table class="page_top" border :columns="initColumns" :data="initData"></Table>
    <div>
      购买总数量：<span class="money">{{TotalCount}}</span>
      购买总金额<span class="money">{{TotalAmount}}</span>
      购买者：<span class="money">{{UserName}}</span>
      购物车号码： <span>{{CartID}}</span>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      initColumns: [
        {
          title: '用户ID',
          key: 'CartItemID',
          width: 100
        },
        {
          title: '购买数量',
          key: 'Count'
          // width: 290
        },
        {
          title: '购买金额小计',
          key: 'Amount',
          width: 160
        },
        {
          title: '购物车ID',
          key: 'CartID',
          width: 270
        },
        {
          title: '操作',
          key: 'action',
          fixed: 'right',
          width: 180,
          align: 'center',
          render: (h, params) => {
            return h('div', [
              h(
                'Button',
                {
                  props: {
                    type: 'primary',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px'
                  },
                  on: {
                    click: () => {
                      console.log(params.row.Book)
                      // this.$router.push({
                      //   name: 'itemAmend',
                      //   params: params.row
                      // })
                      // this.show(params.index)/home/itemAmend
                    }
                  }
                },
                '查看详情'
              ),
              h(
                'Button',
                {
                  props: {
                    type: 'error',
                    size: 'small'
                  },
                  on: {
                    click: () => {
                      console.log(params.row)
                      // this.current_id = params.row.Id
                      // this.remove()
                    }
                  }
                },
                '删除'
              )
            ])
          }
        }
      ],
      initData: [],
      TotalCount: 0,
      TotalAmount: 0,
      UserName: '',
      CartID: ''
    }
  },
  created() {
    this.getCartInfo()
  },
  methods: {
    getCartInfo() {
      this.$axios.get('/home/getCartInfo').then(res => {
        console.log('getCartInfo', res)
        this.initData = res.data.Data.CartItems
        this.TotalCount = res.data.Data.TotalCount
        this.TotalAmount = res.data.Data.TotalAmount
        this.UserName = res.data.Data.UserName
        this.CartID = res.data.Data.CartID
      })
    }
  }
}
</script>

<style lang="less" scoped>
.money {
  color: red;
  font-size: 20px;
  margin-right: 15px;
}
</style>
