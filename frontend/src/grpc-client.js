import { GatewayServiceClient } from './proto/internal/pkg/delivery/grpc/gateway-service/proto/scheme_grpc_web_pb.js'
import { User } from './proto/internal/pkg/delivery/grpc/auth-service/proto/scheme_pb.js'
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

    /**
     *
     * @param {{userUuid, roomUuid, date, token}} payload
     * @param callback
     */
    reserveRoom(payload, callback) {
        console.log('reserveRoom payload: ', payload)
        const userUuid = new UUID
        userUuid.setValue(payload.userUuid)

        const roomUuid = new UUID
        roomUuid.setValue(payload.roomUuid)

        const req = new Reservation()
        req.setUseruuid(userUuid)
        req.setRoomuuid(roomUuid)
        req.setDate(payload.date)

        this.client.addReservation(req, {"authorization": payload.token}, callback)
    }
}