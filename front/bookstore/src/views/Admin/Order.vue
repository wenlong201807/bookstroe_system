<template>
  <div>
    <h3>order</h3>
    <Table class="page_top" border :columns="initColumns" :data="initData"></Table>
    <Modal v-model="modalState" title="修改物流状态" @on-ok="okAmendState" @on-cancel="cancelAmendState">
      <p>订单号：{{orderId}}</p>
      <p>发货状态</p>
    </Modal>
    <Modal v-model="modalRemoveOrder" title="修改物流状态" @on-ok="okRemoveOrder" @on-cancel="cancelRemoveOrder">
      <p>订单号：{{orderId}}</p>
      <p>确定移除交易完成的订单吗？</p>
    </Modal>
  </div>
</template>

<script>
import PostHeader from '@/utils/postCommon.js'
export default {
  data() {
    return {
      modalRemoveOrder: false,
      modalState: false,
      orderId: '',
      current_id: 0,
      initColumns: [
        {
          title: '用户ID',
          key: 'UserID',
          width: 100
        },
        {
          title: '订单ID',
          key: 'OrderID'
          // width: 290
        },
        {
          title: '购买数量',
          key: 'TotalCount',
          width: 160
        },
        {
          title: '消费总金额',
          key: 'TotalAmount',
          width: 130
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
          width: 230,
          align: 'center',
          render: (h, params) => {
            return h('div', [
              h(
                'Button',
                {
                  props: {
                    type: 'primary',
                    size: 'small',
                    disabled: params.row.State !== 0 // 0 未发货  1 已发货 2 交易完成
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
                '发货'
              ),
              h(
                'Button',
                {
                  props: {
                    type: 'error',
                    size: 'small',
                    disabled: params.row.State !== 2 // 0 未发货  1 已发货 2 交易完成
                  },
                  on: {
                    click: () => {
                      console.log(params.row.Id)
                      this.orderId = params.row.OrderID
                      this.modalRemoveOrder = true
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
    this.initGet()
  },
  methods: {
    okRemoveOrder() {
      console.log('想要删除我？', this.orderId)
      // let postData = {
      //   orderId: this.orderId
      // }
      // this.$axios.put('/home/sendOrder', postData, PostHeader).then(res => {
      //   console.log('sendOrder', res)
      //   this.initGet()
      //   this.modalState = false
      //   this.orderId = ''
      // })
    },
    cancelRemoveOrder() {
      this.orderId = ''
      this.modalRemoveOrder = false
    },
    okAmendState() {
      let postData = {
        orderId: this.orderId
      }
      this.$axios.put('/home/sendOrder', postData, PostHeader).then(res => {
        console.log('sendOrder', res)
        this.initGet()
        this.modalState = false
        this.orderId = ''
      })
    },
    cancelAmendState() {
      this.modalState = false
    },
    remove() {
      // this.$axios
      //   .delete('/company/deleteContent?ids=' + this.current_id)
      //   .then(res => {
      //     console.log(res)
      //     this.initGet()
      //   })
    },
    initGet() {
      // let params = {
      //   page: this.currentPage,
      //   rows: this.currentRows
      // }
      this.$axios.get('/home/getorders').then(res => {
        // this.$axios.get('/api/showContent', { params }).then(res => {
        console.log('getorders', res)
        this.initData = res.data.Orders
        // this.totalCount = res.data.total
      })
    }
  }
}
</script>

<style scoped lang="less">
@import '../../styles/common.less';
</style>
