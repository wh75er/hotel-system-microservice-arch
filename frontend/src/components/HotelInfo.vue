<template>
  <el-card class="hotel-info">
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
  </el-card>
</template>

<style>
@import '../../public/main.css';
.hotel-info {
  width: 100%;
}
</style>

<script>
import { ref } from 'vue';
import DatePicker from "./DatePicker.vue";
import Events from "../consts/events.js";
export default {
  name: "Header",
  components: {
    DatePicker,
  },
  props: {
  },
  setup() {
    const pickedDate = ref(0)
    return {
      pickedDate
    }
  },
  data() {
    return {
      hotel: {
        Name: 'HOTEL TEST NAME',
        Description: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Commodi, dignissimos dolore, enim exercitationem facilis hic inventore magni maiores maxime molestias natus omnis optio perspiciatis porro provident repudiandae ullam ut voluptate?',
        City: 'Moscow',
        Country: 'Russia',
        Address: 'Pushkina, 24',
      },
    }
  },
  methods: {
  },
  mounted() {
    this.emitter.on(Events.reservationDateChanged, (newDate) => {
      console.log('ROOM CARD RECEIVER GOT NEW DATE: ', newDate)
      this.pickedDate = newDate;
      console.log('Value in reference: ', this.pickedDate)
    })
  }
}
</script>

