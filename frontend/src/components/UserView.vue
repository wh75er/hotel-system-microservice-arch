<template>
  <div class="user-view">
    <UserInfo/>
    <ReservationCard/>
  </div>
</template>

<script>
import UserInfo from './UserInfo.vue'
import ReservationCard from './ReservationCard.vue'

export default {
  name: "HotelView",
  components: {
    UserInfo,
    ReservationCard,
  }
}
</script>

<style scoped>
@import '../../public/main.css';
.user-view {
}

</style>