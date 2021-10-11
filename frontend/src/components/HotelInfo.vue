<template>
  <el-card class="hotel-info text-wrapping">
    <template #header>
      <div class="card-header">
        <span class="header-text card-header__label">{{hotel.Name}}</span>
      </div>
    </template>
    <div class="box-card__text box-card__item">
      <span class="small-title">Description:</span>
      <div class="box-card__caption caption-text">{{hotel.Description}}</div>
    </div>
    <div class="box-card__text box-card__item">
      <span class="small-title">Location:</span>
      <div class="box-card__caption caption-text">{{hotel.Country}}, {{hotel.City}}</div>
    </div>
    <div class="box-card__text box-card__item">
      <span class="small-title">Address:</span>
      <div class="box-card__caption caption-text">{{hotel.Address}}</div>
    </div>
    <div class="box-card__item">
      <DatePicker/>
    </div>
    <div class="box-card__item box-card__admin-buttons">
      <el-button @click="onAddHotelClick" round>Add Room</el-button>
    </div>
  </el-card>
</template>

<style>
@import '../../public/main.css';
.hotel-info {
  width: 100%;
}
.box-card__admin-buttons {
  display: flex;
  width: 100%;
  justify-content: right;
}
</style>

<script>
import DatePicker from "./DatePicker.vue";
export default {
  name: "Header",
  components: {
    DatePicker,
  },
  props: [
      'hotel',
  ],
  // data() {
  //   return {
  //     hotel: {
  //       Name: 'HOTEL TEST NAME',
  //       Description: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Commodi, dignissimos dolore, enim exercitationem facilis hic inventore magni maiores maxime molestias natus omnis optio perspiciatis porro provident repudiandae ullam ut voluptate?',
  //       City: 'Moscow',
  //       Country: 'Russia',
  //       Address: 'Pushkina, 24',
  //     },
  //   }
  // },
  methods: {
    onAddHotelClick() {
      this.$router.push({name: 'roomCreate', params: { id: this.hotel.HotelUuid }})
    }
  },
}
</script>

