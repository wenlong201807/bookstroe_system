<template>
  <div class="containerLogin">
    <div class="loginCla">
      <Form ref="formInline" :model="formInline" :rules="ruleInline" inline>
        <FormItem prop="username">
          <i-input type="text" v-model="formInline.username" placeholder="用户名">
            <Icon type="ios-person-outline" slot="prepend"></Icon>
          </i-input>
        </FormItem>
        <FormItem prop="password">
          <i-input type="password" v-model="formInline.password" placeholder="密码">
            <Icon type="ios-lock-outline" slot="prepend"></Icon>
          </i-input>
        </FormItem>
        <FormItem>
          <Button type="primary" @click="login('formInline')">登录</Button>
          <Button type="warning" @click="goRegist">去注册</Button>
        </FormItem>
      </Form>
    </div>
  </div>
</template>

<script>
import PostHeader from '@/utils/postCommon.js'
export default {
  name: 'login',
  data() {
    return {
      formInline: {
        username: '',
        password: '123456'
      },
      ruleInline: {
        username: [
          {
            required: true,
            message: 'Please fill in the user name',
            trigger: 'blur'
          }
        ],
        password: [
          {
            required: true,
            message: 'Please fill in the password.',
            trigger: 'blur'
          },
          {
            type: 'string',
            min: 6,
            message: 'The password length cannot be less than 6 bits',
            trigger: 'blur'
          }
        ]
      }
    }
  },
  methods: {
    goRegist() {
      this.$router.push({
        path: '/Regist'
        // params: {
        //   page: '1',
        //   code: '8989',
        //   username: this.formInline.username
        // }
        // this.$route.params.page
      })
    },
    login(name) {
      console.log(this.formInline.username, this.formInline.password)
      this.$refs[name].validate(valid => {
        if (valid) {
          let postData = {
            username: this.formInline.username,
            password: this.formInline.password
          }
          this.$axios
            .post('/home/login', postData, PostHeader)
            .then(res => {
              console.log('login', res)
              if (res.data.Msg === '登录成功') {
                // this.$Message.success('success!')
                // 检验成功 设置登录状态并且跳转到/
                localStorage.setItem('iview_login', true)
                localStorage.setItem('username', this.formInline.username)
                let path = ''
                if (this.formInline.username !== 'admin') {
                  path = '/Home/Buyer/Store'
                } else {
                  path = '/Home/Admin/Book'
                }
                this.$router.push({
                  path
                  // params: {
                  //   page: '1',
                  //   code: '8989',
                  //   username: this.formInline.username
                  // }
                  // this.$route.params.page
                })
              }
            })
            .catch(e => {
              console.log(e)
            })
        } else {
          this.$Message.error('Fail!')
        }
      })
    }
  }
}
</script>

<style lang="less" scoped>
.containerLogin {
  height: 100vh;
  width: 100vw;
  // background: url()
}
.loginCla {
  height: 200px;
  width: 800px;
  padding: 60px 0px 0px 157px;
  margin: 50px auto 0px;
  border: 1px solid #ccc;
  border-radius: 5px;
  box-shadow: 2px 2px 2px pink;
}
</style>
