<template>
  <div>

    <div class="storeContain">
      <Card v-for="item in initData" :key="item.ID" style="width:245px">
        <div style="text-align:center">
          <img :src="item.ImgPath">
          <h3>作者：{{item.Author}}</h3>
          <h3>书名：{{item.Title}}</h3>
          <h3>价格：{{item.Price}}</h3>
          <h3>销量：{{item.Sales}}</h3>
          <h3>库存：{{item.Stock}}</h3>
          <Button type="primary" @click="addBook2Cart(item.ID)">加入购物车</Button>
        </div>
      </Card>
    </div>

    <Page class="page_top" :total="totalCount" :page-size="currentRows" @on-change="changePage" show-elevator show-total />
  </div>
</template>

<script>
import PostHeader from '@/utils/postCommon.js'
export default {
  data() {
    return {
      current_id: 0,
      // purchaseId:0,
      initData: [],
      pageSizeOps: [3, 5, 9],
      totalCount: 0,
      currentPage: 1,
      currentRows: 10
    }
  },
  created() {
    this.rangeBook()
  },
  methods: {
    addBook2Cart(ID) {
      console.log('购买的ID：', ID)
      let postData = {
        bookId: ID
      }
      this.$axios.post('/home/addBook2Cart', postData, PostHeader).then(res => {
        console.log('addBook2Cart', res)
      })
    },
    changePage(page) {
      this.currentPage = page
      this.rangeBook()
    },
    rangeBook() {
      let postData = {
        pageNo: this.currentPage + '' // 第几页
        // pageNo: '2' // 第几页
      }
      this.$axios.get('/home/getPageBooks', { params: postData }).then(res => {
        console.log('getPageBooksByPrice', res)
        console.log(res.data.Book)
        this.initData = res.data.Book
        this.totalCount = res.data.TotalRecord
        console.log(this.initData)
        // if (res.data.Msg === '登录成功') {
        //   this.$Message.success('success!')
        //   this.$router.push({ path: '/home/item' })
        // }
      })
    }
  }
}
</script>

<style>
.storeContain {
  display: flex;
  justify-content: flex-start;
}
</style>
