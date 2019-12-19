<template>
  <div class="homeContainer">
    <div class="LeftMenu">
      <Menu @on-select="onSelect" width="210" theme="dark" :active-name="currentPath" :open-names="['use']">
        <Submenu name="use">
          <template slot="title">
            <Icon type="ios-paper" />
            第一版效果篇
          </template>
          <MenuGroup title="购买者">
            <MenuItem name="/Home/Buyer/Store">书城</MenuItem>
            <MenuItem name="/Home/Buyer/MyCart">我的购物车</MenuItem>
            <MenuItem name="/Home/Buyer/MyOrder">我的订单</MenuItem>
          </MenuGroup>
          <MenuGroup title="管理员">
            <MenuItem name="/Home/Admin/Book">图书管理</MenuItem>
            <MenuItem name="/Home/Admin/Cart">购物车管理</MenuItem>
            <MenuItem name="/Home/Admin/Order">订单管理</MenuItem>
            <MenuItem name="/Home/Admin/User">用户管理</MenuItem>
          </MenuGroup>
          <MenuGroup title="登录与注册">
            <MenuItem name="/Login">登录</MenuItem>
            <MenuItem name="/Regist">注册</MenuItem>

          </MenuGroup>

          <!-- <MenuItem v-for="item in MenuList" :key="item.path" :name="item.path">{{ item.title }}</MenuItem> -->
          <!-- <MenuItem to="/Home/Book" name="/Home/Book">书城管理</MenuItem>:to="item.path" -->
        </Submenu>
        <Submenu name="1">
          <template slot="title">
            <Icon type="ios-paper" />
            内容管理
          </template>
          <MenuItem name="1-1">文章管理</MenuItem>
          <MenuItem name="1-2">评论管理</MenuItem>
          <MenuItem name="1-3">举报管理</MenuItem>
        </Submenu>
        <Submenu name="2">
          <template slot="title">
            <Icon type="ios-people" />
            用户管理
          </template>
          <MenuItem name="2-1">新增用户</MenuItem>
          <MenuItem name="2-2">活跃用户</MenuItem>
        </Submenu>
        <Submenu name="3">
          <template slot="title">
            <Icon type="ios-stats" />
            统计分析
          </template>
          <MenuGroup title="使用">
            <MenuItem name="3-1">新增和启动</MenuItem>
            <MenuItem name="3-2">活跃分析</MenuItem>
            <MenuItem name="3-3">时段分析</MenuItem>
          </MenuGroup>
          <MenuGroup title="留存">
            <MenuItem name="3-4">用户留存</MenuItem>
            <MenuItem name="3-5">流失用户</MenuItem>
          </MenuGroup>
        </Submenu>
      </Menu>
    </div>
    <div class="RightView">
      <div class="overall">
        <h2 class="usernameCla">{{username}}</h2>
        <div class="time">{{timer}}</div>
        <div class="loginoutCla" @click="bookstroeLogout">退出</div>
      </div>
      <router-view />
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      currentPath: '',
      username: localStorage.username || '未登录',
      timer: this.CurentTime(),
      MenuList: [
        { id: 1, title: '书城购物', path: '/Home/Store' },
        { id: 2, title: '书城管理', path: '/Home/Book' },
        { id: 3, title: '订单管理', path: '/Home/Order' },
        { id: 4, title: '购物车管理', path: '/Home/Cart' },
        { id: 5, title: '登录', path: '/Login' },
        { id: 6, title: '注册', path: '/Regist' }
      ]
    }
  },
  created() {
    this.currentPath = this.$route.path
    // console.log(this.$route.path)
    // let intervalId = setInterval(function() {
    //   console.log(4)
    //   this.CurentTime()
    //   // clearInterval(intervalId);
    // }, 1000)
  },
  methods: {
    bookstroeLogout() {
      this.$axios.post('/home/logout').then(res => {
        console.log('退出登录', res)
        if (res.data.Msg === '退出成功') {
          localStorage.clear()
          this.$router.push('/Login')
        }
      })
    },
    CurentTime() {
      // https://blog.csdn.net/qq_33242126/article/details/79279322
      var now = new Date()

      var year = now.getFullYear() // 年
      var month = now.getMonth() + 1 // 月
      var day = now.getDate() // 日
      var hh = now.getHours() // 时
      var mm = now.getMinutes() // 分
      // var m = now.getSeconds() // 秒
      var clock = year + '-'
      if (month < 10) {
        clock += '0'
      }
      clock += month + '-'
      if (day < 10) {
        clock += '0'
      }
      clock += day + ' '
      if (hh < 10) {
        clock += '0'
      }
      clock += hh + ':'
      if (mm < 10) clock += '0'
      clock += mm
      // if (m < 10) clock += '0'
      // clock += m
      return clock
    },

    onSelect(pageName) {
      console.log(pageName)
      this.$router.push(pageName)
    }
  }
}
</script>

<style lang="less" scoped>
.homeContainer {
  width: 100wh;
  height: 100vh;
  min-width: 1000px;
  min-height: 800px;
  // background: pink;
  display: flex;
  .LeftMenu {
    // background: yellowgreen;
    width: 210px;
  }
  .RightView {
    background: bisque;
    width: calc(100% - 210px);
    .overall {
      height: 60px;
      line-height: 60px;
      background: burlywood;
      display: flex;
      justify-content: flex-end; /*右对齐*/
      .usernameCla {
        font-size: 18px;
        color: red;
        margin-right: 20px;
      }
      .time {
        margin-right: 20px;
      }
      .loginoutCla {
        margin-right: 20px;
        cursor: pointer;
      }
    }
  }
}
</style>
