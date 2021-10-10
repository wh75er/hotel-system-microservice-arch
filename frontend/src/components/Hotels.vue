<template>
  <div class="hotels">
    <div v-if="hotels.length === 0">There's nothing here yet</div>
    <HotelCard class="hotels__item"
        :to="{ name: 'hotel', params: { id: hotel.HotelUuid}}"
        v-for="hotel in hotels"
        :key="hotel"
        :hotel="hotel"
        @click="onHotelClick(hotel.HotelUuid)"
    />
  </div>
</template>

<script>
import HotelCard from './HotelCard.vue'
import {ElNotification} from "element-plus";
import { ref } from 'vue';

export default {
  name: "Header",
  components: {
    HotelCard
  },
  setup() {
    let hotels = ref([])
    return {
      hotels
    }
  },
  mounted() {
    this.gatewayClient.getHotels(function(error, response) {
      if (error) {
        ElNotification({
          title: 'Error',
          message: error.message,
          type: 'error',
        })
      } else {
        const hotels = response.getHotelsList()
        this.hotels = []
        for (const ph of hotels) {
          console.log(ph)
          console.log(ph.getName())
          this.hotels.push({
            HotelUuid: ph.getHoteluuid(),
            Name: ph.getName(),
            Description: ph.getDescription(),
            Country: ph.getCountry(),
            Address: ph.getAddress(),
            City: ph.getCity(),
          })
        }
      }
    }.bind(this))
  },
  // data() {
  //   return {
  //     hotels: [
  //       {
  //         HotelUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //         Name: 'HOTEL TEST NAME',
  //         Description: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Commodi, dignissimos dolore, enim exercitationem facilis hic inventore magni maiores maxime molestias natus omnis optio perspiciatis porro provident repudiandae ullam ut voluptate?',
  //         Country: 'Russia',
  //         Address: 'Pushkina, 24',
  //         City: 'Moscow',
  //       },
  //       {
  //         HotelUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //         Name: 'HOTEL 2 TEST NAME',
  //         Description: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Commodi, dignissimos dolore, enim exercitationem facilis hic inventore magni maiores maxime molestias natus omnis optio perspiciatis porro provident repudiandae ullam ut voluptate?',
  //         City: 'Sydney',
  //         Country: 'Australia',
  //         Address: 'Tuntero, 10',
  //       },
  //       {
  //         HotelUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //         Name: 'HOTEL 3 TEST NAME',
  //         Description: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Commodi, dignissimos dolore, enim exercitationem facilis hic inventore magni maiores maxime molestias natus omnis optio perspiciatis porro provident repudiandae ullam ut voluptate?',
  //         City: 'Sydney',
  //         Country: 'Australia',
  //         Address: 'Tuntero, 10',
  //       },
  //       {
  //         HotelUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //         Name: 'HOTEL 4 TEST NAME',
  //         Description: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Commodi, dignissimos dolore, enim exercitationem facilis hic inventore magni maiores maxime molestias natus omnis optio perspiciatis porro provident repudiandae ullam ut voluptate?',
  //         City: 'Sydney',
  //         Country: 'Australia',
  //         Address: 'Tuntero, 10',
  //       }
  //     ],
  //   }
  // },
  methods: {
    onHotelClick(hotelUuid) {
      console.log('HERE', hotelUuid)
      console.log('router: ', this.$router)
      this.$router.push({name: 'hotel', params: {id: hotelUuid}})
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