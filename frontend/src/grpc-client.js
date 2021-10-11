import { GatewayServiceClient } from './proto/internal/pkg/delivery/grpc/gateway-service/proto/scheme_grpc_web_pb.js'
import { User } from './proto/internal/pkg/delivery/grpc/auth-service/proto/scheme_pb.js'
import { Hotel, Room } from './proto/internal/pkg/delivery/grpc/hotel-service/proto/scheme_pb.js'
import { Reservation } from './proto/internal/pkg/delivery/grpc/reservation-service/proto/scheme_pb.js'
import { Token, Empty, UUID } from './proto/internal/pkg/delivery/grpc/commonProto/common_pb.js'
import { gatewayLocation }  from './consts/config.js'

export default class GatewayClient {
    constructor() {
        this.client = new GatewayServiceClient(gatewayLocation, null, null)
    }

    /**
     *
     * @param {{login: String, password: String}} payload
     * @param {function} callback
     */
    login(payload, callback) {
        console.log('Login payload: ', payload)
        const req = new User()
        req.setLogin(payload.login)
        req.setPassword(payload.password)
        this.client.login(req, {}, callback)
    }

    checkAuth(token, callback) {
        const req = new Token()
        req.setValue(token)
        this.client.checkAuth(req, {}, callback)
    }

    /**
     *
     * @param {{login: String, password: String}} payload
     * @param {function} callback
     */
    signup(payload, callback) {
        console.log('signup payload: ', payload)
        const req = new User()
        req.setLogin(payload.login)
        req.setPassword(payload.password)
        this.client.addUser(req, {}, callback)
    }

    getHotels(callback) {
        const req = new Empty()
        this.client.getHotels(req, {}, callback)
    }

    getHotel(uuid, callback) {
        console.log('getHotel uuid is: ', uuid)
        const req = new UUID()
        req.setValue(uuid)
        this.client.getHotel(req, {}, callback)
    }

    getRoom(uuid, callback) {
        console.log('getHotel uuid is: ', uuid)
        const req = new UUID()
        req.setValue(uuid)
        this.client.getRoom(req, {}, callback)
    }

    /**
     *
     * @param {{userUuid, roomUuid, date, token}} payload
     * @param callback
     */
    reserveRoom(payload, callback) {
        console.log('reserveRoom payload: ', payload)
        const userUuid = new UUID()
        userUuid.setValue(payload.userUuid)

        const roomUuid = new UUID()
        roomUuid.setValue(payload.roomUuid)

        const req = new Reservation()
        req.setUseruuid(userUuid)
        req.setRoomuuid(roomUuid)
        req.setDate(payload.date)

        this.client.addReservation(req, {"authorization": payload.token}, callback)
    }

    /**
     *
     * @param {{userUuid, token}} payload
     * @param callback
     */
    getReservations(payload, callback) {
        console.log('getReservations payload: ', payload)
        const req = new UUID()
        req.setValue(payload.userUuid)

        this.client.getReservationsByUser(req, {"authorization": payload.token}, callback)
    }

    /**
     *
     * @param {{paymentUuid, token}} payload
     * @param callback
     */
    getPayment(payload, callback) {
        console.log('getPayment payload: ', payload)
        const req = new UUID()
        req.setValue(payload.paymentUuid)

        this.client.getPayment(req, {"authorization": payload.token}, callback)
    }

    /**
     *
     * @param {{userUuid, token}} payload
     * @param callback
     */
    getLoyalty(payload, callback) {
        console.log('getLoyalty payload: ', payload)
        const req = new UUID()
        req.setValue(payload.userUuid)

        this.client.getDiscount(req, {"authorization": payload.token}, callback)
    }

    /**
     *
     * @param {{reservationUuid, token}} payload
     * @param callback
     */
    createPayment(payload, callback) {
        console.log('createPayment payload: ', payload)
        const req = new UUID()
        req.setValue(payload.reservationUuid)

        this.client.createPayment(req, {"authorization": payload.token}, callback)
    }

    /**
     *
     * @param {{reservationUuid, token}} payload
     * @param callback
     */
    cancelReservation(payload, callback) {
        console.log('createPayment payload: ', payload)
        const req = new UUID()
        req.setValue(payload.reservationUuid)

        this.client.cancelReservation(req, {"authorization": payload.token}, callback)
    }

    /**
     *
     * @param {{hotel: {name, description, country, city, address}, token}} payload
     * @param callback
     */
    createHotel(payload, callback) {
        console.log('createHotel payload: ', payload)
        const req = new Hotel()
        req.setName(payload.hotel.name)
        req.setDescription(payload.hotel.description)
        req.setCountry(payload.hotel.country)
        req.setCity(payload.hotel.city)
        req.setAddress(payload.hotel.address)

        this.client.addHotel(req, {"authorization": payload.token}, callback)
    }

    /**
     *
     * @param {{room: {hotelUuid, roomType, amount, beds, offers, nightPrice}, token}} payload
     * @param callback
     */
    createRoom(payload, callback) {
        console.log('createRoom payload: ', payload)
        const req = new Room()
        req.setHoteluuid(payload.room.hotelUuid)
        req.setRoomtype(payload.room.roomType)
        req.setAmount(payload.room.amount)
        req.setBeds(payload.room.beds)
        req.setOffersList(payload.room.offers)
        req.setNightprice(payload.room.nightPrice)

        this.client.addRoom(req, {"authorization": payload.token}, callback)
    }

    /**
     *
     * @param {{room: {hotelUuid, roomUuid, roomType, amount, beds, offers, nightPrice}, token}} payload
     * @param callback
     */
    patchRoom(payload, callback) {
        console.log('patchRoom payload: ', payload)
        const req = new Room()
        req.setHoteluuid(payload.room.hotelUuid)
        req.setRoomuuid(payload.room.roomUuid)
        req.setRoomtype(payload.room.roomType)
        req.setAmount(payload.room.amount)
        req.setBeds(payload.room.beds)
        req.setOffersList(payload.room.offers)
        req.setNightprice(payload.room.nightPrice)

        this.client.patchRoom(req, {"authorization": payload.token}, callback)
    }
}