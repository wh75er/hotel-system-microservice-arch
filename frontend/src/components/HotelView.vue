<template>
  <div class="hotel-view">
    <HotelInfo/>
    <ReservationCard/>
  </div>
</template>

<script>
import HotelInfo from './HotelInfo.vue'
// import RoomCard from './RoomCard.vue'
import ReservationCard from './ReservationCard.vue'

export default {
  name: "HotelView",
  components: {
    HotelInfo,
    // RoomCard,
    ReservationCard,
  }
}
</script>

<style scoped>
@import '../../public/main.css';
.hotel-view {
}

</style>