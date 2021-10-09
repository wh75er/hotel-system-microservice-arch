<template>
<el-container>
  <el-header><Header :login="login"/></el-header>
  <el-main><router-view/></el-main>
</el-container>
</template>

<script>
import Header from './Header.vue'
import Events from "@/consts/events";
import { ref } from 'vue';

export default {
  name: 'Layout',
  components: {
    Header,
  },
  setup() {
    const login = ref('')
    return {
      login,
    }
  },
  mounted() {
    this.emitter.on(Events.userLoggedIn, (token) => {
      this.userSingletone.login(token)
      this.$router.push({name: 'home'})
      this.login = this.userSingletone.claims.login
    })
    this.emitter.on(Events.userLoggedOut, () => {
      this.userSingletone.logout()
      this.$router.push({name: 'home'})
      this.login = ''
    })
    if (this.userSingletone.token) {
      console.log('Found token: checking authorization')
      this.gatewayClient.checkAuth(this.userSingletone.token, function(error) {
        if (error) {
          this.emitter.emit(Events.userLoggedOut)
        } else {
          this.login = this.userSingletone.claims.login
        }
      }.bind(this))
    }
  },
}
</script>

<style>
@import '../../public/main.css';
  .el-header, .el-footer {
    background-color: #B3C0D1;
    padding: 0 0;
    color: #333;
    text-align: center;
    line-height: 60px;
  }
  
  .el-aside {
    background-color: #D3DCE6;
    color: #333;
    text-align: center;
    line-height: 200px;
  }
  
  .el-main {
    background-color: #E9EEF3;
    color: #333;
  }

  body {
    background-color: #E9EEF3;
  }
  
  body > .el-container {
    margin-bottom: 40px;
  }
  
  .el-container:nth-child(5) .el-aside,
  .el-container:nth-child(6) .el-aside {
    line-height: 260px;
  }
  
  .el-container:nth-child(7) .el-aside {
    line-height: 320px;
  }
</style>

