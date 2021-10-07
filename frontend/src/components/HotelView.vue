<template>
  <div class="hotel-view">
    <HotelInfo/>
    <RoomCard/>
  </div>
</template>

<script>
import HotelInfo from './HotelInfo.vue'
import RoomCard from './RoomCard.vue'

export default {
  name: "HotelView",
  components: {
    HotelInfo,
    RoomCard,
  }
}
</script>

<style scoped>
@import '../../public/main.css';
.hotel-view {
}

</style>