<template>
  <el-card class="statistics-card text-wrapping">
    <template #header>
      <div class="card-header">
        <span class="header-text card-header__label">Statistics Data</span>
      </div>
    </template>
    <div class="box-card__text box-card__item">
      <span class="small-title">Reservations Amount:</span>
      <div class="box-card__caption caption-text">{{stats.reservationsAmount}}</div>
    </div>
    <div class="box-card__text box-card__item">
      <span class="small-title">Free Rooms Amount:</span>
      <div class="box-card__caption caption-text">{{stats.roomsAmount}}</div>
    </div>
  </el-card>
</template>

<style>
@import '../../public/main.css';
.statistics-card {
  width: 100%;
}
</style>

<script>
import {ElNotification} from "element-plus";
import {ref} from 'vue';

export default {
  name: "StatisticsCard",
  setup() {
    let stats = ref ({
      roomsAmount: '',
      reservationsAmount: '',
    })
    return {
      stats
    }
  },
  mounted() {
    this.gatewayClient.getStats(this.userSingletone.token, function(error, response) {
      if (error) {
        console.log('Error in get stats: ', error)
        ElNotification({
          title: 'Error',
          message: error.message,
          type: 'error',
        })
      } else {
        this.stats.roomsAmount = response.getRoomsamount()
        this.stats.reservationsAmount = response.getReservationsamount()
        console.log(this.stats)
      }
    }.bind(this))
  },
}
</script>