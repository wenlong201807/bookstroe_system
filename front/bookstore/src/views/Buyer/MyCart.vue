<template>
  <div>
    <Button type="primary" @click="checkout">结算</Button>
    <Button type="primary" @click="clearCartModel">清空购物车</Button>
    <Table class="page_top" border :columns="initColumns" :data="initData"></Table>
    <div>
      购买总数量：<span class="money">{{TotalCount}}</span>
      购买总金额<span class="money">{{TotalAmount}}</span>
      购买者：<span class="money">{{UserName}}</span>
      购物车号码： <span>{{CartID}}</span>

    </div>
    <Modal v-model="modelCheckDetail" title="购买的图书详情" @on-cancel="cancelCheckDetail" footer-hide>
      <p>图书ID：{{detailData.ID}}</p>
      <p>书名：{{detailData.Title}}</p>
      <p>作者：{{detailData.Author}}</p>
      <p>单价：{{detailData.Price}}</p>
      <p>销量：{{detailData.Sales}}</p>
      <p>库存：{{detailData.Stock}}</p>
      <p>图片路径：{{detailData.ImgPath}}</p>
    </Modal>
    <Modal v-model="modelDeleteCartItem" title="删除购物项确认框" @on-ok="okDeleteCartItem" @on-cancel="cancelDeleteCartItem">
      <p>确定要删除这个宝贝吗？</p>
    </Modal>
    <Modal v-model="modelDeleteCart" title="删除购物车确认框" @on-ok="okDeleteCart" @on-cancel="cancelDeleteCart">
      <p>确定要删除整个购物车吗？</p>
    </Modal>
  </div>
</template>

<script>
export default {
  data() {
    return {
      modelDeleteCart: false,
      modelDeleteCartItem: false,
      cartItemId: 0, // 删除购物项使用
      modelCheckDetail: false,
      detailData: {
        ID: 0,
        Title: '',
        Author: '',
        Price: 0,
        Sales: 0,
        Stock: 0,
        ImgPath: ''
      },
      initColumns: [
        {
          title: '购物项ID',
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
                      // this.detailData = params.row.Book
                      this.detailData.ID = params.row.Book.ID
                      this.detailData.Title = params.row.Book.Title
                      this.detailData.Author = params.row.Book.Author
                      this.detailData.Price = params.row.Book.Price
                      this.detailData.Sales = params.row.Book.Sales
                      this.detailData.Stock = params.row.Book.Stock
                      this.detailData.ImgPath = params.row.Book.ImgPath
                      this.modelCheckDetail = true
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
                      this.modelDeleteCartItem = true
                      this.cartItemId = params.row.CartItemID
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
    clearCartModel() {
      this.modelDeleteCart = true
    },
    okDeleteCart() {
      let cartId = this.CartID
      this.$axios.delete('/home/deleteCart?cartId=' + cartId).then(res => {
        setTimeout(() => {
          this.$Message.success('删除购物车成功')
          this.$router.push('/Home/Buyer/Store')
        }, 366)
        this.modelDeleteCart = false
        console.log('deleteCart', res)
      })
    },
    cancelDeleteCart() {
      this.modelDeleteCart = false
      this.$Message.success('取消删除购物车')
    },
    okDeleteCartItem() {
      let cartItemId = this.cartItemId
      this.$axios
        .delete('/home/deleteCartItem?cartItemId=' + cartItemId)
        .then(res => {
          this.modelDeleteCartItem = false
          this.$Message.success('删除购物项成功')
          this.getCartInfo()
          console.log('deleteCartItem', res)
        })
    },
    cancelDeleteCartItem() {
      this.modelDeleteCartItem = false
      this.$Message.success('取消删除购物项')
    },
    checkout() {
      this.$axios.post('/home/checkout').then(res => {
        console.log('checkout', res)
        setTimeout(() => {
          this.$Message.success('生成订单：' + res.data)
          this.$router.push('/Home/Buyer/MyOrder')
        }, 366)
      })
    },
    cancelCheckDetail() {
      this.detailData.ID = ''
      this.detailData.Title = ''
      this.detailData.Author = ''
      this.detailData.Price = ''
      this.detailData.Sales = ''
      this.detailData.Stock = ''
      this.detailData.ImgPath = ''
      this.modelCheckDetail = false
    },
    getCartInfo() {
      this.$axios
        .get('/home/getCartInfo')
        .then(res => {
          console.log('getCartInfo', res)
          this.initData = res.data.Data.CartItems
          this.TotalCount = res.data.Data.TotalCount
          this.TotalAmount = res.data.Data.TotalAmount
          this.UserName = res.data.Data.UserName
          this.CartID = res.data.Data.CartID // 清空购物车使用
        })
        .catch(e => {
          console.log(e)
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
