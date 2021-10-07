<template>
  <div class="hotels"><HotelCard/></div>
</template>

<script>
import HotelCard from './HotelCard.vue'

export default {
  name: "Header",
  components: {
    HotelCard
  },
  props: {
  },
}
</script>

<style scoped>
@import '../../public/main.css';
  .hotels {
    display: flex;
    justify-content: center;
    background-color: red;
  }
</style>