<template>
  <div class="header">
    <div class="start-block">
      <div class="start-block__label">
        <span class="header-text label-text">Booktle</span>
      </div>
    </div>
    <div class="end-block">
      <div v-if="login" class="user-data end-block__item">
        <span class="user-data__login">{{login}}</span>
        <el-button type="danger">Log out</el-button>
      </div>
      <div v-else class="signup-buttons end-block__item">
        <el-button type="primary">Log in</el-button>
        <el-button type="warning">Sign up</el-button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Header",
  props: {
    login: String
  }
}
</script>

<style scoped>
@import '../../public/main.css';
  .header {
    display: flex;
    margin: 0 12px 0 55px;
    justify-content: space-between;
  }
  .start-block {
    display: flex;
  }
  .label-text {
    font-size: var(--el-font-size-extra-large);
  }
  .end-block {
    display: flex;
  }
  .end-block__item {
    margin-left: 15px
  }
  .user-data__login {
    font-size: 20px;
    padding: 0 10px;
  }
</style>