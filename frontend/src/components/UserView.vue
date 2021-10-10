<template>
  <div class="user-view">
    <UserInfo :user="user"/>
    <ReservationCard :user="user"/>
  </div>
</template>

<script>
import UserInfo from './UserInfo.vue'
import ReservationCard from './ReservationCard.vue'
import Events from '../consts/events.js'
import {ref} from "vue";

export default {
  name: "UserView",
  components: {
    UserInfo,
    ReservationCard,
  },
  setup() {
    const user = ref({
      Login: '',
      UserUuid: '',
      Role: '',
      LoyaltyStatus: '',
    })
    return {
      user,
    }
  },
  beforeRouteEnter(to, from, next) {
    next(vm => {
      console.log(vm)
      if (!vm.userSingletone.claims.userUuid) {
        vm.emitter.emit(Events.unauthorizedRedirect)
      }
    })
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

    this.gatewayClient.getLoyalty(
        {userUuid: claims.userUuid, token: this.userSingletone.token},
        function (error, response) {
          if (error) {
            console.log('Failed to retrieve discount information: ', error)
          } else {
            this.user.LoyaltyStatus = response.getStatus()
          }
        }.bind(this)
    )
  },
  // data() {
  //   return {
  //     data: {
  //       user: {
  //         Login: 'cool login',
  //         UserUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //         Role: 'user',
  //         LoyaltyStatus: 'Bronze',
  //       },
  //       reservations: [
  //         {
  //           ReservationUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //           RoomUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //           UserUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //           PaymentUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //           Status: 'active',
  //           Date: '1633593545',
  //           Hotel: {
  //             HotelUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //             Name: 'HOTEL TEST NAME',
  //             Description: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Commodi, dignissimos dolore, enim exercitationem facilis hic inventore magni maiores maxime molestias natus omnis optio perspiciatis porro provident repudiandae ullam ut voluptate?',
  //             City: 'Moscow',
  //             Country: 'Russia',
  //             Address: 'Pushkina, 24',
  //           },
  //           Room: {
  //             RoomUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //             RoomType: 'Family',
  //             Beds: 4,
  //             Offers: [
  //               'food included',
  //               'pool included',
  //             ],
  //             Price: 400,
  //           },
  //           Payment: {
  //             PaymentUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //             UserUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //             Status: 'New',
  //             Price: 200,
  //           },
  //         },
  //         {
  //           ReservationUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //           RoomUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //           UserUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //           PaymentUUid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //           Status: 'active',
  //           Date: '1633593545',
  //           Hotel: {
  //             HotelUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //             Name: 'Best Hotel',
  //             Description: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Commodi, dignissimos dolore, enim exercitationem facilis hic inventore magni maiores maxime molestias natus omnis optio perspiciatis porro provident repudiandae ullam ut voluptate?',
  //             City: 'Moscow',
  //             Country: 'Russia',
  //             Address: 'Pushkina, 24',
  //           },
  //           Room: {
  //             RoomUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //             RoomType: 'Family',
  //             Beds: 4,
  //             Offers: [
  //               'food included',
  //               'pool included',
  //             ],
  //             Price: 400,
  //           },
  //           Payment: {
  //             PaymentUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //             UserUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //             Status: 'New',
  //             Price: 200,
  //           },
  //         },
  //         {
  //           ReservationUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //           RoomUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //           UserUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //           PaymentUUid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //           Status: 'active',
  //           Date: '1633593545',
  //           Hotel: {
  //             HotelUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //             Name: 'Paid hotel',
  //             Description: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Commodi, dignissimos dolore, enim exercitationem facilis hic inventore magni maiores maxime molestias natus omnis optio perspiciatis porro provident repudiandae ullam ut voluptate?',
  //             City: 'Moscow',
  //             Country: 'Russia',
  //             Address: 'Pushkina, 24',
  //           },
  //           Room: {
  //             RoomUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //             RoomType: 'Family',
  //             Beds: 4,
  //             Offers: [
  //               'food included',
  //               'pool included',
  //             ],
  //             Price: 400,
  //           },
  //           Payment: {
  //             PaymentUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //             UserUuid: '34ee2e3d-83c6-4ed8-b884-af42bdb6e53a',
  //             Status: 'Paid',
  //             Price: 200,
  //           },
  //         },
  //       ],
  //     },
  //   }
  // }
}
</script>

<style scoped>
@import '../../public/main.css';
.user-view {
}

</style>