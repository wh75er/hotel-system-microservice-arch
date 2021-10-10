import {ElNotification} from "element-plus";

export default function updateReservationPayment(gatewayClient, paymentUuid, token, updateCallback) {
    gatewayClient.getPayment(
        {paymentUuid: paymentUuid, token: token},
        function (error, response) {
            if (error) {
                console.log('Failed to get payment: ', error)
                ElNotification({
                    title: 'Error',
                    message: 'Failed to get created payment. Try later',
                    type: 'error',
                })
            } else {

            }
        }
    )
}