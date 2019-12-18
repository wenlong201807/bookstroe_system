<template>
  <div class="container">
    <Button type="primary" @click="addBook">添加图书</Button>
    <Table class="page_top" border :columns="initColumns" :data="initData"></Table>
    <Page class="page_top" :total="totalCount" :page-size="currentRows" @on-change="changePage" show-elevator show-total />

    <Modal v-model="modalAddBook" title="添加图书" footer-hide @on-cancel="cancelAddBook('fromAdd')">
      <Form ref="fromAdd" :model="fromAddAmend" :rules="rulesAddAmend" :label-width="80">
        <FormItem label="书名" prop="Title">
          <Input v-model.trim="fromAddAmend.Title" placeholder="Enter your Title"></Input>
        </FormItem>
        <FormItem label="作者" prop="Author">
          <Input v-model.trim="fromAddAmend.Author" placeholder="Enter your Author"></Input>
        </FormItem>
        <FormItem label="单价" prop="Price">
          <Input v-model.trim="fromAddAmend.Price" placeholder="Enter your Price"></Input>
        </FormItem>
        <FormItem label="销量" prop="Sales">
          <Input v-model.trim="fromAddAmend.Sales" placeholder="Enter your Sales"></Input>
        </FormItem>
        <FormItem label="库存" prop="Stock">
          <Input v-model.trim="fromAddAmend.Stock" placeholder="Enter your Stock"></Input>
        </FormItem>
        <FormItem label="图片路径" prop="ImgPath">
          <Input v-model.trim="fromAddAmend.ImgPath" placeholder="Enter your ImgPath"></Input>
        </FormItem>

        <FormItem>
          <Button type="primary" @click="okAddBook('fromAdd')">提交</Button>
          <Button @click="cancelAddBook('fromAdd')" style="margin-left: 8px">取消</Button>
        </FormItem>
      </Form>
      <!-- <div slot="footer">
        <Button type="error" size="large" long :loading="modal_loading" @click="del">Delete</Button>
      </div> -->
    </Modal>

    <Modal v-model="modalAmendBook" title="修改图书" footer-hide @on-cancel="cancelAmendBook('fromAmend')">
      <Form ref="fromAmend" :model="fromAddAmend" :rules="rulesAddAmend" :label-width="80">
        <FormItem label="书名" prop="Title">
          <Input v-model.trim="fromAddAmend.Title" placeholder="Enter your Title"></Input>
        </FormItem>
        <FormItem label="作者" prop="Author">
          <Input v-model.trim="fromAddAmend.Author" placeholder="Enter your Author"></Input>
        </FormItem>
        <FormItem label="单价" prop="Price">
          <Input v-model.trim="fromAddAmend.Price" placeholder="Enter your Price"></Input>
        </FormItem>
        <FormItem label="销量" prop="Sales">
          <Input v-model.trim="fromAddAmend.Sales" placeholder="Enter your Sales"></Input>
        </FormItem>
        <FormItem label="库存" prop="Stock">
          <Input v-model.trim="fromAddAmend.Stock" placeholder="Enter your Stock"></Input>
        </FormItem>
        <FormItem label="图片路径" prop="ImgPath">
          <Input v-model.trim="fromAddAmend.ImgPath" placeholder="Enter your ImgPath"></Input>
        </FormItem>

        <FormItem>
          <Button type="primary" @click="okAmendBook('fromAmend')">提交</Button>
          <Button @click="cancelAmendBook('fromAmend')" style="margin-left: 8px">取消</Button>
        </FormItem>
      </Form>
      <!-- <div slot="footer">
        <Button type="error" size="large" long :loading="modal_loading" @click="del">Delete</Button>
      </div> -->
    </Modal>

    <Modal v-model="modalRemoveBook" title="删除图书提醒框" @on-ok="okRemoveBook" @on-cancel="cancelRemoveBook">
      <p>确定要删除吗？</p>
    </Modal>
  </div>
</template>

<script>
import PostHeader from '@/utils/postCommon.js'
export default {
  data() {
    return {
      modalRemoveBook: false,
      modalAmendBook: false,
      modalAddBook: false,
      fromAddAmend: {
        Title: '',
        Author: '',
        Price: '',
        Sales: '',
        Stock: '',
        ImgPath: ''
      },
      rulesAddAmend: {
        Title: {
          required: true,
          message: 'Title cannot be empty',
          trigger: 'blur'
        },
        Author: {
          required: true,
          message: 'Author cannot be empty',
          trigger: 'blur'
        },
        Price: [
          {
            required: true,
            message: 'Price cannot be empty',
            trigger: 'blur'
          },
          {
            message: '只能是数字',
            trigger: 'change',
            pattern: /^(([1-9]\d{0,3})|0)(\.\d{0,2})?$/
          }
        ],
        Sales: [
          {
            required: true,
            message: 'Sales cannot be empty',
            trigger: 'blur'
          },
          {
            message: '只能是数字',
            trigger: 'change',
            pattern: /^(([1-9]\d{0,3})|0)(\.\d{0,2})?$/
          }
        ],
        Stock: [
          {
            required: true,
            message: 'Stock cannot be empty',
            trigger: 'blur'
          },
          {
            message: '只能是数字',
            trigger: 'change',
            pattern: /^(([1-9]\d{0,3})|0)(\.\d{0,2})?$/
          }
        ],
        ImgPath: {
          required: true,
          message: 'ImgPath cannot be empty',
          trigger: 'blur'
        }
      },
      currentId: 0,
      // purchaseId:0,
      initColumns: [
        {
          title: '图书ID',
          key: 'ID',
          width: 120
        },
        {
          title: '标题',
          key: 'Title',
          width: 130
        },
        {
          title: '作者',
          key: 'Author',
          width: 130
        },
        {
          title: '单价',
          key: 'Price',
          width: 130
        },
        {
          title: '销量',
          key: 'Sales'
          // width: 130
        },
        {
          title: '库存',
          key: 'Stock'
          // width: 130
        },
        {
          title: '图片路径',
          key: 'ImgPath'
          // width: 130
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
                      // console.log(params.row)
                      this.amendBook(params.row)
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
                      console.log(params.row)
                      this.modalRemoveBook = true
                      this.currentId = params.row.ID
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
      totalCount: 0,
      currentPage: 1,
      currentRows: 4
    }
  },
  created() {
    this.rangeBook()
  },
  methods: {
    okAddBook(name) {
      this.$refs[name].validate(valid => {
        if (valid) {
          let postData = {
            title: this.fromAddAmend.Title,
            author: this.fromAddAmend.Author,
            price: this.fromAddAmend.Price,
            sales: this.fromAddAmend.Sales,
            stock: this.fromAddAmend.Stock
          }
          this.$axios
            .post('/home/addBook', postData, PostHeader)
            .then(res => {
              console.log('addBook', res)

              if (res.data.Msg === '添加成功') {
                this.rangeBook()
                this.$Message.success('添加成功')
                this.$refs[name].resetFields()
                this.modalAddBook = false
                //   this.$router.push({ path: '/home/item' })
              }
            })
            .catch(e => {
              this.$Message.success('添加失败')
              this.modalAddBook = true
            })
        }
      })
    },
    cancelAddBook(name) {
      this.$refs[name].resetFields()
      this.modalAddBook = false
    },
    okAmendBook(name) {
      this.$refs[name].validate(valid => {
        if (valid) {
          let postData = {
            bookId: this.currentId,
            title: this.fromAddAmend.Title,
            author: this.fromAddAmend.Author,
            price: this.fromAddAmend.Price,
            sales: this.fromAddAmend.Sales,
            stock: this.fromAddAmend.Stock
          }
          this.$axios
            .post('/home/upduateBookPage', postData, PostHeader)
            .then(res => {
              console.log('upduateBookPage', res)
              if (res.data.Msg === '修改成功') {
                this.rangeBook()
                this.$Message.success('修改成功')
                this.$refs[name].resetFields()
                this.modalAmendBook = false
                // this.$router.push({ path: '/home/item' })
              }
            })
        }
      })
    },
    cancelAmendBook(name) {
      this.$refs[name].resetFields()
      this.modalAmendBook = false
    },
    addBook() {
      this.modalAddBook = true
    },
    amendBook(row) {
      this.modalAmendBook = true
      console.log('amendBook**row:', row)
      this.currentId = row.ID
      this.fromAddAmend = row
      console.log('修改前获取参数', this.fromAddAmend)
    },
    okRemoveBook() {
      console.log('removeBook')
      this.$axios
        .delete('/home/deleteBook?bookId=' + this.currentId)
        .then(res => {
          console.log(res)
          if (res.data.Msg === '删除成功') {
            this.rangeBook()
            this.$Message.success('删除成功')
            this.modalRemoveBook = false
            // this.$router.push({ path: '/home/item' })
          }
        })
    },
    cancelRemoveBook() {
      this.modalRemoveBook = false
      this.currentId = 0
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
        this.currentRows = res.data.PageSize
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

<style scoped lang="less">
@import '../styles/common.less';
</style>
