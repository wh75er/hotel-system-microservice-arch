<template>
  <div class="hotel-view">
    <HotelInfo :hotel="hotel"/>
    <RoomCard :rooms="rooms"/>
  </div>
</template>

<script>
import HotelInfo from './HotelInfo.vue'
import RoomCard from './RoomCard.vue'
import { ref } from 'vue'
import {ElNotification} from "element-plus";

export default {
  name: "HotelView",
  setup() {
    const hotel = ref({})
    const rooms = ref([])
    return {
      hotel,
      rooms,
    }
  },
  mounted() {
    console.log('UUID OF HOTEL IS: ', this.$route.params.id)
    this.gatewayClient.getHotel(this.$route.params.id, function(error, ph) {
          if (error) {
            console.log('getHotel error: ', error)
            ElNotification({
              title: 'Error',
              message: error.message,
              type: 'error',
            })
          } else {
            this.hotel = {
              HotelUuid: ph.getHoteluuid(),
              Name: ph.getName(),
              Description: ph.getDescription(),
              Country: ph.getCountry(),
              Address: ph.getAddress(),
              City: ph.getCity(),
            }
            const rooms = ph.getRoomsList()
            console.log('Rooms response: ', rooms)
            this.rooms = []
            for (const rh of rooms) {
              console.log(rh)
              console.log(rh.getRoomuuid())
              this.rooms.push({
                RoomUuid: rh.getRoomuuid(),
                HotelUuid: rh.getHoteluuid(),
                RoomType: rh.getRoomtype(),
                Beds: rh.getBeds(),
                Offers: rh.getOffersList(),
                Price: rh.getNightprice(),
              })
            }
            console.log('result of rooms: ', this.rooms)
          }
    }.bind(this))
  },
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