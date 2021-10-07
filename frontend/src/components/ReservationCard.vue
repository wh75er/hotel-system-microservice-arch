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
          v-if='reservations[scope.$index].Payment.Status === "New"'
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
  props: [
      'user',
      'reservations',
  ],
  data() {
    return {
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
