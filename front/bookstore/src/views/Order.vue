<template>
  <div>
    <h3>order</h3>
    <Table class="page_top" border :columns="initColumns" :data="initData"></Table>
  </div>
</template>

<script>
export default {
  data() {
    return {
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
          width: 170
        },
        {
          title: '操作',
          key: 'action',
          fixed: 'right',
          width: 130,
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
                      console.log(params.row)
                      // this.$router.push({
                      //   name: 'itemAmend',
                      //   params: params.row
                      // })
                      // this.show(params.index)/home/itemAmend
                    }
                  }
                },
                '修改'
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
                      console.log(params.row.Id)
                      // this.current_id = params.row.Id
                      this.remove()
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
      currentRows: 10
    }
  },
  created() {
    this.initGet()
  },
  methods: {
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
@import '../styles/common.less';
</style>
