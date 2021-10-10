import {ElNotification} from "element-plus";

export default function getHeavyReservations(payload, gatewayClient, reservations, updateCallback) {
    gatewayClient.getReservations(payload, function(error, response) {
        if (error) {
            console.log('Failed to retrieve reservations: ', error)
            ElNotification({
                title: 'Error',
                message: `Failed to get reservations list: ${error.message}`,
                type: 'error',
            })
        } else {
            console.log('response from server: ', response)
            const prs = response.getValueList()
            for (const pr of prs) {
                let currentReservation = {
                    ReservationUuid: pr.getReservationuuid().getValue(),
                    RoomUuid: pr.getRoomuuid().getValue(),
                    UserUuid: pr.getUseruuid().getValue(),
                    PaymentUuid: pr.getPaymentuuid().getValue(),
                    Status: pr.getStatus(),
                    Date: pr.getDate() * 1000,
                }

                getRoomAndHotel(gatewayClient, currentReservation, updateCallback)
                getPayment(gatewayClient, currentReservation, payload.token, updateCallback)

                reservations.push(
                    currentReservation
                )
            }
        }
    })
}

function getRoomAndHotel(gatewayClient, currentReservation, updateCallback) {
    if (!currentReservation) { return }
    gatewayClient.getRoom(currentReservation.RoomUuid, function(error, response) {
        if (error) {
            console.log('Failed to retrieve room: ', error)
        } else {
            const room = {
                RoomUuid: response.getRoomuuid(),
                HotelUuid: response.getHoteluuid(),
                RoomType: response.getRoomtype(),
                Beds: response.getBeds(),
                Offers: response.getOffersList(),
                Price: response.getNightprice(),
            }
            Object.assign(currentReservation, {Room: room})

            getHotel(gatewayClient, currentReservation, updateCallback)
        }
    })
}

function getHotel(gatewayClient, currentReservation, updateCallback) {
    if (!(currentReservation && currentReservation.Room && currentReservation.Room.HotelUuid)) { return }
    gatewayClient.getHotel(currentReservation.Room.HotelUuid, function(error, response) {
        if (error) {
            console.log('Failed to retrieve hotel: ', error)
        } else {
            const hotel = {
                HotelUuid: response.getHoteluuid(),
                Name: response.getName(),
                Description: response.getDescription(),
                Country: response.getCountry(),
                Address: response.getAddress(),
                City: response.getCity(),
            }
            Object.assign(currentReservation, {Hotel: hotel})
            updateCallback()
        }
    })
}

export function getPayment(gatewayClient, currentReservation, token, updateCallback) {
    if (!(currentReservation && currentReservation.PaymentUuid)) { return }
    gatewayClient.getPayment({ paymentUuid: currentReservation.PaymentUuid, token: token},
        function(error, response) {
        if (error) {
            console.log('Failed to retrieve payment: ', error)
        } else {
            const payment = {
                PaymentUuid: response.getPaymentuuid().getValue(),
                UserUuid: response.getUseruuid().getValue(),
                Status: response.getStatus(),
                Price: response.getPrice(),
            }
            Object.assign(currentReservation, {Payment: payment})
            updateCallback()
        }
    })
}