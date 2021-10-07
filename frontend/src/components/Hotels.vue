<template>
  <div class="hotels">
    <HotelCard class="hotels__item" v-for="hotel in hotels" :key="hotel" :hotel="hotel"/>
  </div>
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
  data() {
    return {
      hotels: [
        {
          Name: 'HOTEL TEST NAME',
          Description: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Commodi, dignissimos dolore, enim exercitationem facilis hic inventore magni maiores maxime molestias natus omnis optio perspiciatis porro provident repudiandae ullam ut voluptate?',
          Country: 'Russia',
          Address: 'Pushkina, 24',
          City: 'Moscow',
        },
        {
          Name: 'HOTEL 2 TEST NAME',
          Description: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Commodi, dignissimos dolore, enim exercitationem facilis hic inventore magni maiores maxime molestias natus omnis optio perspiciatis porro provident repudiandae ullam ut voluptate?',
          City: 'Sydney',
          Country: 'Australia',
          Address: 'Tuntero, 10',
        },
        {
          Name: 'HOTEL 3 TEST NAME',
          Description: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Commodi, dignissimos dolore, enim exercitationem facilis hic inventore magni maiores maxime molestias natus omnis optio perspiciatis porro provident repudiandae ullam ut voluptate?',
          City: 'Sydney',
          Country: 'Australia',
          Address: 'Tuntero, 10',
        },
        {
          Name: 'HOTEL 4 TEST NAME',
          Description: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Commodi, dignissimos dolore, enim exercitationem facilis hic inventore magni maiores maxime molestias natus omnis optio perspiciatis porro provident repudiandae ullam ut voluptate?',
          City: 'Sydney',
          Country: 'Australia',
          Address: 'Tuntero, 10',
        }
      ],
    }
  }
}
</script>

<style scoped>
@import '../../public/main.css';
  .hotels {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
  }

  .hotels__item {
    margin: 0 0 10px 0;
  }
</style>