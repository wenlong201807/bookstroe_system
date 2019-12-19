<template>
  <div>
    <h2>我的订单</h2>
    <Table class="page_top" border :columns="initColumns" :data="initData"></Table>
    <Modal v-model="modalState" title="修改物流状态" @on-ok="okAmendState" @on-cancel="cancelAmendState">
      <p>订单号：{{orderId}}</p>
      <p>发货状态</p>
    </Modal>
    <Modal v-model="modelCheckOrderDetail" title="订单详情" @on-cancel="cancelCheckOrderDetail" footer-hide>
      <p>图书分类ID：{{OrderDetail.OrderItemID}}</p>
      <p>图书ID：{{OrderDetail.Count}}</p>
      <p>书名：{{OrderDetail.Title}}</p>
      <p>作者：{{OrderDetail.Author}}</p>
      <p>单价：{{OrderDetail.Price}}</p>
      <p>总价：{{OrderDetail.Amount}}</p>
      <p>图片路径：{{OrderDetail.ImgPath}}</p>
      <p>订单编号：{{OrderDetail.OrderID}}</p>

    </Modal>
  </div>
</template>

<script>
import PostHeader from '@/utils/postCommon.js'
export default {
  data() {
    return {
      modelCheckOrderDetail: false,
      OrderDetail: {
        OrderItemID: 6,
        Count: 1,
        Amount: 22.2,
        Title: '给孩子的诗',
        Author: '北岛',
        Price: 22.2,
        ImgPath: 'static/img/default.jpg',
        OrderID: 'f8dd1418-fcf5-41d1-7d7c-5f759a7d6ee9'
      },
      modalState: false,
      orderId: '',
      current_id: 0,
      initColumns: [
        {
          title: '用户ID',
          key: 'UserID'
          // width: 100
        },
        {
          title: '订单ID',
          key: 'OrderID',
          width: 290
        },
        {
          title: '购买数量',
          key: 'TotalCount'
          // width: 160
        },
        {
          title: '消费总金额',
          key: 'TotalAmount'
          // width: 130
        },
        {
          title: '物流状态',
          key: 'State',
          width: 170,
          render: (h, params) => {
            // let result = '' // 待发货，已发货，已收货
            return h('div', {}, this.stateList[params.row.State].conente)
          }
        },
        {
          title: '操作',
          key: 'action',
          fixed: 'right',
          width: 290,
          align: 'center',
          render: (h, params) => {
            return h('div', [
              h(
                'Button',
                {
                  props: {
                    type: 'primary',
                    size: 'small',
                    disabled: params.row.State !== 1 // 0 未发货  1 已发货 2 交易完成
                  },
                  style: {
                    marginRight: '5px'
                  },
                  on: {
                    click: () => {
                      console.log(params.row)
                      this.modalState = true
                      this.orderId = params.row.OrderID
                    }
                  }
                },
                '收货'
              ),
              h(
                'Button',
                {
                  props: {
                    type: 'error',
                    size: 'small',
                    disabled: true
                  },
                  on: {
                    click: () => {
                      console.log(params.row.Id)
                      // this.current_id = params.row.Id
                      // this.remove()
                    }
                  }
                },
                '删除'
              ),
              h(
                'Button',
                {
                  props: {
                    type: 'success',
                    size: 'small'
                  },
                  on: {
                    click: () => {
                      console.log(params.row)
                      this.modelCheckOrderDetail = true
                      this.orderId = params.row.OrderID
                      this.getOrderItemsInfo()
                      // this.OrderDetail.OrderID = params.row.OrderID
                      // this.OrderDetail.CreateTime = params.row.CreateTime
                      // this.OrderDetail.TotalCount = params.row.TotalCount
                      // this.OrderDetail.TotalAmount = params.row.TotalAmount
                      // this.OrderDetail.State = params.row.State
                      // this.OrderDetail.UserID = params.row.UserID
                    }
                  }
                },
                '订单详情'
              )
            ])
          }
        }
      ],
      initData: [],
      pageSizeOps: [3, 5, 9],
      totalCount: 0,
      currentPage: 1,
      currentRows: 10,
      stateList: [
        // 0 未发货  1 已发货 2 交易完成
        { state: 0, conente: '未发货' },
        { state: 1, conente: '已发货' },
        { state: 2, conente: '交易完成' }
      ]
    }
  },
  created() {
    this.getMyOrders()
  },
  methods: {
    okAmendState() {
      let postData = {
        orderId: this.orderId
      }
      this.$axios.put('/home/receiveOrder', postData, PostHeader).then(res => {
        console.log('receiveOrder', res)
        this.getMyOrders()
        this.modalState = false
        this.orderId = ''
      })
    },
    cancelAmendState() {
      this.modalState = false
    },
    cancelCheckOrderDetail() {
      this.modelCheckOrderDetail = false
    },
    getOrderItemsInfo() {
      this.$axios
        .get('/home/getOrderItemsInfo?orderId=' + this.orderId)
        .then(res => {
          console.log('getOrderItemsInfo', res)

          this.OrderDetail.OrderItemID = res.data[0].OrderItemID
          this.OrderDetail.Count = res.data[0].Count
          this.OrderDetail.Amount = res.data[0].Amount
          this.OrderDetail.Title = res.data[0].Title
          this.OrderDetail.Author = res.data[0].Author
          this.OrderDetail.Price = res.data[0].Price
          this.OrderDetail.ImgPath = res.data[0].ImgPath
          this.OrderDetail.OrderID = res.data[0].OrderID
        })
    },
    getMyOrders() {
      this.$axios.get('/home/getMyOrders').then(res => {
        console.log('getMyOrders', res)
        this.initData = res.data
      })
    }
  }
}
</script>

<style>
</style>
