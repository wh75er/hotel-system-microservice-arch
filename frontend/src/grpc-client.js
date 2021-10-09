import { GatewayServiceClient } from './proto/internal/pkg/delivery/grpc/gateway-service/proto/scheme_grpc_web_pb.js'
import { User } from './proto/internal/pkg/delivery/grpc/auth-service/proto/scheme_pb.js'
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
    async login(payload, callback) {
        console.log('Login payload: ', payload)
        const req = new User()
        req.setLogin(payload.login)
        req.setPassword(payload.password)
        this.client.login(req, {}, callback)
    }
}