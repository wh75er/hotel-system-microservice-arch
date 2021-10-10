<template>
  <el-table class="room-card" :data="rooms" border style="width: 100%">
    <el-table-column prop="RoomType" label="Room Type" width="180" />
    <el-table-column prop="Beds" label="Beds" width="180" />
    <el-table-column prop="Offers" label="Offers" />
    <el-table-column prop="Price" label="Price" width="180" />
    <el-table-column width="75" >
      <template #default="scope">
        <el-button
          type="text"
          size="small"
          @click.prevent="clicked(scope.$index, rooms, reservationDate)"
        >
          Reserve
        </el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import {ref} from "vue";
import Events from "@/consts/events";
import {ElNotification} from "element-plus";

export default {
  setup() {
    const reservationDate = ref(0)
    return {
      reservationDate
    }
  },
  props: [
      'rooms'
  ],
  // data() {
  //   return {
  //     rooms: [
  //       {
  //         RoomType: 'Family',
  //         Beds: 4,
  //         Offers: [
  //             "food included",
  //             "pool included",
  //         ],
  //         Price: 400,
  //       },
  //       {
  //         RoomType: 'Lone wolf',
  //         Beds: 1,
  //         Offers: [
  //           "food included",
  //         ],
  //         Price: 150,
  //       },
  //     ],
  //   }
  // },
  methods: {
    clicked(index, rows, reservationDate) {
      if (!reservationDate) {
        ElNotification({
          title: 'Error',
          message: 'Please pick a reservation date',
          type: 'error',
        })
        return
      }
      console.log(index, rows)
      console.log(rows[0])
      console.log(reservationDate)
      this.gatewayClient.reserveRoom({
        userUuid: this.userSingletone.claims.userUuid,
        roomUuid: rows[index].RoomUuid,
        date: reservationDate * 0.001,
        token: this.userSingletone.token,
      }, function (error, response) {
        if (error) {
          console.log('reservation error: ', error)
          ElNotification({
            title: 'Error',
            message: error.message,
            type: 'error',
          })
        } else {
          console.log('response is: ', response)
          ElNotification({
            title: 'Success',
            message: `Successfully reserved room with uuid:`, //${response.getValue()}`,
            type: 'success',
          })
        }
      }.bind(this))
    }
  },
  mounted() {
    this.emitter.on(Events.reservationDateChanged, (newDate) => {
      console.log('ROOM CARD RECEIVER GOT NEW DATE: ', newDate)
      this.reservationDate = newDate;
      console.log('Value in reference: ', this.reservationDate)
    })
  }
}
</script>

<style>
.room-card {
  box-shadow: var(--el-box-shadow-light);
}
</style>
