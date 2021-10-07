<template>
  <el-card class="user-info">
    <template #header>
      <div class="card-header">
        <span class="header-text card-header__label">{{user.Login}}</span>
      </div>
    </template>
    <div class="box-card__text box-card__item">
      <span class="small-title">Role:</span>
      <div class="box-card__caption caption-text">{{user.Role}}</div>
    </div>
    <div class="box-card__text box-card__item">
      <span class="small-title">Discount status:</span>
      <div class="box-card__caption caption-text">{{user.LoyaltyStatus}}</div>
    </div>
  </el-card>
</template>

<style>
@import '../../public/main.css';
.user-info {
  width: 100%;
}
</style>

<script>
export default {
  name: "Header",
  components: {
  },
  props: [
    'user'
  ],
}
</script>

