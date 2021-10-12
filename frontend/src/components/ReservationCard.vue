<template>
  <el-table class="reservation-card" :data="reservations" border style="width: 100%">
    <el-table-column prop="Hotel.Name" label="Hotel" width="180" >
      <template #default="scope">
        <router-link v-if="reservations[scope.$index] && reservations[scope.$index].Hotel"
                     :to="{ name: 'hotel', params: { id: reservations[scope.$index].Hotel.HotelUuid}}">
          {{reservations[scope.$index].Hotel.Name}}
        </router-link>
      </template>
    </el-table-column>
    <el-table-column prop="Room.RoomType" label="Room Type" width="180" />
    <el-table-column prop="Status" label="Status" width="80" />
    <el-table-column prop="Payment.PaymentUuid" label="PaymentUuid" />
    <el-table-column prop="Payment.Price" label="Price" width="80" />
    <el-table-column prop="Payment.Status" label="Payment Status" width="80" />
    <el-table-column prop="Date" label="Date" width="100" />
    <el-table-column width="75" >
      <template #default="scope">
        <iframe v-if="reservations[scope.$index] && reservations[scope.$index].Payment && reservations[scope.$index].Payment.Status && reservations[scope.$index].Payment.Status === 'New' && reservations[scope.$index].Status && reservations[scope.$index].Status !== 'canceled'"
            :src="'https://yoomoney.ru/quickpay/button-widget?targets=Pay%20for%20reservation&default-sum=' + reservations[scope.$index].Payment.Price +'&label=' + reservations[scope.$index].Payment.PaymentUuid + '&button-text=11&any-card-payment-type=on&button-size=s&button-color=orange&successURL=http%3A%2F%2Fbookinghotelservice.ru&quickpay=small&account=4100117213941944'"
            width="127"
            height="25"
            frameborder="0"
            allowtransparency="true"
            scrolling="no">
        </iframe>
        <el-button
            type="text"
            size="small"
            v-else-if="reservations[scope.$index] && reservations[scope.$index].Status && reservations[scope.$index].Status !== 'canceled' && reservations[scope.$index].Payment && reservations[scope.$index].Payment.Status && reservations[scope.$index].Payment.Status !== 'Paid'"
            @click.prevent="createPaymentClicked(scope.$index, reservations)"
        >
            create<br>paycheck
        </el-button>
        <el-button
            type="text"
            size="small"
            v-if="reservations[scope.$index].Status && reservations[scope.$index].Status !== 'canceled'"
            @click.prevent="cancelClicked(scope.$index, reservations)"
        >
          Cancel
        </el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import Events from "@/consts/events";
import getHeavyReservations from "@/helpers/heavyReservations";
import { ref } from 'vue';
import {ElNotification} from "element-plus";
import {getPayment} from '../helpers/heavyReservations.js';

export default {
  props: [
      'user',
  ],
  setup() {
    let reservations = ref([])
    return {
      reservations,
    }
  },
  mounted() {
    if (!this.userSingletone.claims.userUuid) {
      this.emitter.emit(Events.unauthorizedRedirect)
    }

    const claims = this.userSingletone.claims

    console.log('user claims: ', claims)
    this.user.Login = claims.login
    this.user.UserUuid = claims.userUuid
    this.user.Role = claims.role

    console.log(this)
    getHeavyReservations(
        { userUuid: this.userSingletone.claims.userUuid, token: this.userSingletone.token },
        this.gatewayClient,
        this.reservations,
        () => {
          const tmp = this.reservations
          this.reservations = []
          this.reservations = tmp
        }
    )
    console.log('reservations: ', this.reservations)
  },
  data() {
    return {
    }
  },
  methods: {
    payClicked(index, rows) {
      console.log(index, rows)
    },
    createPaymentClicked(index, rows) {
      console.log('Payment creation button clicked')
      console.log(index, rows)
      console.log(rows[index].ReservationUuid)
      this.gatewayClient.createPayment(
          {reservationUuid: rows[index].ReservationUuid, token: this.userSingletone.token},
          function (error, response) {
            if (error) {
              console.log('Failed to create payment for reservation: ', error)
              ElNotification({
                title: 'Error',
                message: 'Failed to create payment for that reservation. Try later',
                type: 'error',
              })
            } else {
              rows[index].PaymentUuid = response.getValue()
              getPayment(this.gatewayClient, rows[index], this.userSingletone.token, () => {
                this.reservations = rows
              })

            }
          }.bind(this)
      )
    },
    cancelClicked(index, rows) {
      if (!(rows && rows[index] && rows[index].ReservationUuid && rows[index].Status)) { return }
      this.gatewayClient.cancelReservation(
          { reservationUuid: rows[index].ReservationUuid, token: this.userSingletone.token },
          function (error) {
            if (error) {
              console.log('Failed to cancel reservation: ', error)
              ElNotification({
                title: 'Error',
                message: 'Failed to cancel reservation. Try later',
                type: 'error',
              })
            } else {
              rows[index].Status = 'canceled'
            }
          }
      )
    }
  }
}
</script>

<style>
.reservation-card {
  box-shadow: var(--el-box-shadow-light);
}
</style>
