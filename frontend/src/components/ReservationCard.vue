<template>
  <el-table class="reservation-card" :data="reservations" border style="width: 100%">
    <el-table-column prop="Hotel.Name" label="Hotel" width="180" >
      <template #default="scope">
        <router-link :to="{ name: 'hotel', params: { id: reservations[scope.$index].Hotel.HotelUuid}}">
          {{reservations[scope.$index].Hotel.Name}}
        </router-link>
      </template>
    </el-table-column>
    <el-table-column prop="Room.RoomType" label="Room Type" width="180" />
    <el-table-column prop="Payment.PaymentUuid" label="PaymentUuid" />
    <el-table-column prop="Payment.Price" label="Price" width="180" />
    <el-table-column prop="Payment.Status" label="Status" width="180" />
    <el-table-column width="75" >
      <template #default="scope">
        <el-button
          type="text"
          size="small"
          @click.prevent="clicked(scope.$index, reservations)"
        >
          Pay
        </el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  data() {
    return {
      reservations: [
        {
          ReservationUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
          RoomUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
          UserUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
          PaymentUUid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
          Status: 'active',
          Date: '1633593545',
          Hotel: {
            HotelUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
            Name: 'HOTEL TEST NAME',
            Description: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Commodi, dignissimos dolore, enim exercitationem facilis hic inventore magni maiores maxime molestias natus omnis optio perspiciatis porro provident repudiandae ullam ut voluptate?',
            City: 'Moscow',
            Country: 'Russia',
            Address: 'Pushkina, 24',
          },
          Room: {
            RoomUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
            RoomType: 'Family',
            Beds: 4,
            Offers: [
              'food included',
              'pool included',
            ],
            Price: 400,
          },
          Payment: {
            PaymentUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
            UserUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
            Status: 'New',
            Price: 200,
          },
        },
        {
          ReservationUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
          RoomUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
          UserUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
          PaymentUUid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
          Status: 'active',
          Date: '1633593545',
          Hotel: {
            HotelUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
            Name: 'Best Hotel',
            Description: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Commodi, dignissimos dolore, enim exercitationem facilis hic inventore magni maiores maxime molestias natus omnis optio perspiciatis porro provident repudiandae ullam ut voluptate?',
            City: 'Moscow',
            Country: 'Russia',
            Address: 'Pushkina, 24',
          },
          Room: {
            RoomUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
            RoomType: 'Family',
            Beds: 4,
            Offers: [
              'food included',
              'pool included',
            ],
            Price: 400,
          },
          Payment: {
            PaymentUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
            UserUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
            Status: 'New',
            Price: 200,
          },
        },
      ],
    }
  },
  methods: {
    clicked(index, rows) {
      console.log(index, rows)
    }
  }
}
</script>

<style>
.reservation-card {
  box-shadow: var(--el-box-shadow-light);
}
</style>
