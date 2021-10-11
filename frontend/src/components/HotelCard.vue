<template>
  <el-card class="hotel-card" shadow="hover">
    <template #header>
      <div class="card-header">
        <span class="header-text card-header__label">{{hotel.Name}}</span>
      </div>
    </template>
    <div class="box-card__text box-card__item">
      <span class="small-title">Description:</span>
      <div class="box-card__caption caption-text text-overflow">{{hotel.Description}}</div>
    </div>
    <div class="box-card__text box-card__item">
      <span class="small-title">Address:</span>
      <div class="box-card__caption caption-text">{{hotel.City}}</div>
    </div>
  </el-card>
</template>

<style>
@import '../../public/main.css';
.hotel-card {
  width: 480px;
  cursor: pointer;
}
</style>

<script>
export default {
  name: "Header",
  props: [
    'hotel',
  ],
  // data() {
  //   return {
  //     hotel: {
  //       Name: 'HOTEL TEST NAME',
  //       Description: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Commodi, dignissimos dolore, enim exercitationem facilis hic inventore magni maiores maxime molestias natus omnis optio perspiciatis porro provident repudiandae ullam ut voluptate?',
  //       City: 'Moscow',
  //     }
  //   }
  // }
}
</script>

