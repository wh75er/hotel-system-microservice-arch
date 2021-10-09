import { GatewayServiceClient } from './proto/internal/pkg/delivery/grpc/gateway-service/proto/scheme_grpc_web_pb.js'
import { User } from './proto/internal/pkg/delivery/grpc/auth-service/proto/scheme_pb.js'
import { Token } from './proto/internal/pkg/delivery/grpc/commonProto/common_pb.js'
import { gatewayLocation }  from './consts/config.js'

export default class GatewayClient {
    constructor() {
        this.client = new GatewayServiceClient(gatewayLocation, null, null)
    }

    /**
     *
     * @param {{login: String, password: String}} payload
     * @param {function} callback
     * @returns {Promise<void>}
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
     * @returns {Promise<void>}
     */
    signup(payload, callback) {
        console.log('signup payload: ', payload)
        const req = new User()
        req.setLogin(payload.login)
        req.setPassword(payload.password)
        this.client.addUser(req, {}, callback)
    }
}