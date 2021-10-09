/**
 * @fileoverview gRPC-Web generated client stub for proto
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var internal_pkg_delivery_grpc_commonProto_common_pb = require('../../../../../../internal/pkg/delivery/grpc/commonProto/common_pb.js')

var internal_pkg_delivery_grpc_auth$service_proto_scheme_pb = require('../../../../../../internal/pkg/delivery/grpc/auth-service/proto/scheme_pb.js')

var internal_pkg_delivery_grpc_hotel$service_proto_scheme_pb = require('../../../../../../internal/pkg/delivery/grpc/hotel-service/proto/scheme_pb.js')

var internal_pkg_delivery_grpc_loyalty$service_proto_scheme_pb = require('../../../../../../internal/pkg/delivery/grpc/loyalty-service/proto/scheme_pb.js')

var internal_pkg_delivery_grpc_payment$service_proto_scheme_pb = require('../../../../../../internal/pkg/delivery/grpc/payment-service/proto/scheme_pb.js')

var internal_pkg_delivery_grpc_reservation$service_proto_scheme_pb = require('../../../../../../internal/pkg/delivery/grpc/reservation-service/proto/scheme_pb.js')
const proto = {};
proto.proto = require('./scheme_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.GatewayServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.GatewayServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.Reservation,
 *   !proto.common.UUID>}
 */
const methodDescriptor_GatewayService_AddReservation = new grpc.web.MethodDescriptor(
  '/proto.GatewayService/AddReservation',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_reservation$service_proto_scheme_pb.Reservation,
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  /**
   * @param {!proto.proto.Reservation} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.Reservation,
 *   !proto.common.UUID>}
 */
const methodInfo_GatewayService_AddReservation = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  /**
   * @param {!proto.proto.Reservation} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID.deserializeBinary
);


/**
 * @param {!proto.proto.Reservation} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.common.UUID)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.common.UUID>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.GatewayServiceClient.prototype.addReservation =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.GatewayService/AddReservation',
      request,
      metadata || {},
      methodDescriptor_GatewayService_AddReservation,
      callback);
};


/**
 * @param {!proto.proto.Reservation} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.common.UUID>}
 *     Promise that resolves to the response
 */
proto.proto.GatewayServicePromiseClient.prototype.addReservation =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.GatewayService/AddReservation',
      request,
      metadata || {},
      methodDescriptor_GatewayService_AddReservation);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.common.Empty>}
 */
const methodDescriptor_GatewayService_CancelReservation = new grpc.web.MethodDescriptor(
  '/proto.GatewayService/CancelReservation',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.UUID,
 *   !proto.common.Empty>}
 */
const methodInfo_GatewayService_CancelReservation = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.common.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.common.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.GatewayServiceClient.prototype.cancelReservation =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.GatewayService/CancelReservation',
      request,
      metadata || {},
      methodDescriptor_GatewayService_CancelReservation,
      callback);
};


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.common.Empty>}
 *     Promise that resolves to the response
 */
proto.proto.GatewayServicePromiseClient.prototype.cancelReservation =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.GatewayService/CancelReservation',
      request,
      metadata || {},
      methodDescriptor_GatewayService_CancelReservation);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.proto.Reservation>}
 */
const methodDescriptor_GatewayService_GetReservation = new grpc.web.MethodDescriptor(
  '/proto.GatewayService/GetReservation',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  internal_pkg_delivery_grpc_reservation$service_proto_scheme_pb.Reservation,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_reservation$service_proto_scheme_pb.Reservation.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.UUID,
 *   !proto.proto.Reservation>}
 */
const methodInfo_GatewayService_GetReservation = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_reservation$service_proto_scheme_pb.Reservation,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_reservation$service_proto_scheme_pb.Reservation.deserializeBinary
);


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.Reservation)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.Reservation>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.GatewayServiceClient.prototype.getReservation =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.GatewayService/GetReservation',
      request,
      metadata || {},
      methodDescriptor_GatewayService_GetReservation,
      callback);
};


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.Reservation>}
 *     Promise that resolves to the response
 */
proto.proto.GatewayServicePromiseClient.prototype.getReservation =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.GatewayService/GetReservation',
      request,
      metadata || {},
      methodDescriptor_GatewayService_GetReservation);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.proto.Reservations>}
 */
const methodDescriptor_GatewayService_GetReservationsByUser = new grpc.web.MethodDescriptor(
  '/proto.GatewayService/GetReservationsByUser',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  internal_pkg_delivery_grpc_reservation$service_proto_scheme_pb.Reservations,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_reservation$service_proto_scheme_pb.Reservations.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.UUID,
 *   !proto.proto.Reservations>}
 */
const methodInfo_GatewayService_GetReservationsByUser = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_reservation$service_proto_scheme_pb.Reservations,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_reservation$service_proto_scheme_pb.Reservations.deserializeBinary
);


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.Reservations)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.Reservations>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.GatewayServiceClient.prototype.getReservationsByUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.GatewayService/GetReservationsByUser',
      request,
      metadata || {},
      methodDescriptor_GatewayService_GetReservationsByUser,
      callback);
};


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.Reservations>}
 *     Promise that resolves to the response
 */
proto.proto.GatewayServicePromiseClient.prototype.getReservationsByUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.GatewayService/GetReservationsByUser',
      request,
      metadata || {},
      methodDescriptor_GatewayService_GetReservationsByUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.common.UUID>}
 */
const methodDescriptor_GatewayService_CreatePayment = new grpc.web.MethodDescriptor(
  '/proto.GatewayService/CreatePayment',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.UUID,
 *   !proto.common.UUID>}
 */
const methodInfo_GatewayService_CreatePayment = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID.deserializeBinary
);


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.common.UUID)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.common.UUID>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.GatewayServiceClient.prototype.createPayment =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.GatewayService/CreatePayment',
      request,
      metadata || {},
      methodDescriptor_GatewayService_CreatePayment,
      callback);
};


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.common.UUID>}
 *     Promise that resolves to the response
 */
proto.proto.GatewayServicePromiseClient.prototype.createPayment =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.GatewayService/CreatePayment',
      request,
      metadata || {},
      methodDescriptor_GatewayService_CreatePayment);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.User,
 *   !proto.common.Empty>}
 */
const methodDescriptor_GatewayService_AddUser = new grpc.web.MethodDescriptor(
  '/proto.GatewayService/AddUser',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_auth$service_proto_scheme_pb.User,
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  /**
   * @param {!proto.proto.User} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.User,
 *   !proto.common.Empty>}
 */
const methodInfo_GatewayService_AddUser = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  /**
   * @param {!proto.proto.User} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.proto.User} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.common.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.common.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.GatewayServiceClient.prototype.addUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.GatewayService/AddUser',
      request,
      metadata || {},
      methodDescriptor_GatewayService_AddUser,
      callback);
};


/**
 * @param {!proto.proto.User} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.common.Empty>}
 *     Promise that resolves to the response
 */
proto.proto.GatewayServicePromiseClient.prototype.addUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.GatewayService/AddUser',
      request,
      metadata || {},
      methodDescriptor_GatewayService_AddUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.User,
 *   !proto.common.Token>}
 */
const methodDescriptor_GatewayService_Login = new grpc.web.MethodDescriptor(
  '/proto.GatewayService/Login',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_auth$service_proto_scheme_pb.User,
  internal_pkg_delivery_grpc_commonProto_common_pb.Token,
  /**
   * @param {!proto.proto.User} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.Token.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.User,
 *   !proto.common.Token>}
 */
const methodInfo_GatewayService_Login = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_commonProto_common_pb.Token,
  /**
   * @param {!proto.proto.User} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.Token.deserializeBinary
);


/**
 * @param {!proto.proto.User} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.common.Token)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.common.Token>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.GatewayServiceClient.prototype.login =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.GatewayService/Login',
      request,
      metadata || {},
      methodDescriptor_GatewayService_Login,
      callback);
};


/**
 * @param {!proto.proto.User} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.common.Token>}
 *     Promise that resolves to the response
 */
proto.proto.GatewayServicePromiseClient.prototype.login =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.GatewayService/Login',
      request,
      metadata || {},
      methodDescriptor_GatewayService_Login);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.Token,
 *   !proto.proto.Role>}
 */
const methodDescriptor_GatewayService_CheckAuth = new grpc.web.MethodDescriptor(
  '/proto.GatewayService/CheckAuth',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.Token,
  internal_pkg_delivery_grpc_auth$service_proto_scheme_pb.Role,
  /**
   * @param {!proto.common.Token} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_auth$service_proto_scheme_pb.Role.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.Token,
 *   !proto.proto.Role>}
 */
const methodInfo_GatewayService_CheckAuth = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_auth$service_proto_scheme_pb.Role,
  /**
   * @param {!proto.common.Token} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_auth$service_proto_scheme_pb.Role.deserializeBinary
);


/**
 * @param {!proto.common.Token} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.Role)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.Role>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.GatewayServiceClient.prototype.checkAuth =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.GatewayService/CheckAuth',
      request,
      metadata || {},
      methodDescriptor_GatewayService_CheckAuth,
      callback);
};


/**
 * @param {!proto.common.Token} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.Role>}
 *     Promise that resolves to the response
 */
proto.proto.GatewayServicePromiseClient.prototype.checkAuth =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.GatewayService/CheckAuth',
      request,
      metadata || {},
      methodDescriptor_GatewayService_CheckAuth);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.Hotel,
 *   !proto.common.Empty>}
 */
const methodDescriptor_GatewayService_AddHotel = new grpc.web.MethodDescriptor(
  '/proto.GatewayService/AddHotel',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_hotel$service_proto_scheme_pb.Hotel,
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  /**
   * @param {!proto.proto.Hotel} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.Hotel,
 *   !proto.common.Empty>}
 */
const methodInfo_GatewayService_AddHotel = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  /**
   * @param {!proto.proto.Hotel} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.proto.Hotel} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.common.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.common.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.GatewayServiceClient.prototype.addHotel =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.GatewayService/AddHotel',
      request,
      metadata || {},
      methodDescriptor_GatewayService_AddHotel,
      callback);
};


/**
 * @param {!proto.proto.Hotel} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.common.Empty>}
 *     Promise that resolves to the response
 */
proto.proto.GatewayServicePromiseClient.prototype.addHotel =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.GatewayService/AddHotel',
      request,
      metadata || {},
      methodDescriptor_GatewayService_AddHotel);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.proto.Hotel>}
 */
const methodDescriptor_GatewayService_GetHotel = new grpc.web.MethodDescriptor(
  '/proto.GatewayService/GetHotel',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  internal_pkg_delivery_grpc_hotel$service_proto_scheme_pb.Hotel,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_hotel$service_proto_scheme_pb.Hotel.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.UUID,
 *   !proto.proto.Hotel>}
 */
const methodInfo_GatewayService_GetHotel = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_hotel$service_proto_scheme_pb.Hotel,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_hotel$service_proto_scheme_pb.Hotel.deserializeBinary
);


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.Hotel)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.Hotel>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.GatewayServiceClient.prototype.getHotel =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.GatewayService/GetHotel',
      request,
      metadata || {},
      methodDescriptor_GatewayService_GetHotel,
      callback);
};


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.Hotel>}
 *     Promise that resolves to the response
 */
proto.proto.GatewayServicePromiseClient.prototype.getHotel =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.GatewayService/GetHotel',
      request,
      metadata || {},
      methodDescriptor_GatewayService_GetHotel);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.Empty,
 *   !proto.proto.HotelsResponse>}
 */
const methodDescriptor_GatewayService_GetHotels = new grpc.web.MethodDescriptor(
  '/proto.GatewayService/GetHotels',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  internal_pkg_delivery_grpc_hotel$service_proto_scheme_pb.HotelsResponse,
  /**
   * @param {!proto.common.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_hotel$service_proto_scheme_pb.HotelsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.Empty,
 *   !proto.proto.HotelsResponse>}
 */
const methodInfo_GatewayService_GetHotels = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_hotel$service_proto_scheme_pb.HotelsResponse,
  /**
   * @param {!proto.common.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_hotel$service_proto_scheme_pb.HotelsResponse.deserializeBinary
);


/**
 * @param {!proto.common.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.HotelsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.HotelsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.GatewayServiceClient.prototype.getHotels =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.GatewayService/GetHotels',
      request,
      metadata || {},
      methodDescriptor_GatewayService_GetHotels,
      callback);
};


/**
 * @param {!proto.common.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.HotelsResponse>}
 *     Promise that resolves to the response
 */
proto.proto.GatewayServicePromiseClient.prototype.getHotels =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.GatewayService/GetHotels',
      request,
      metadata || {},
      methodDescriptor_GatewayService_GetHotels);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.Hotel,
 *   !proto.common.Empty>}
 */
const methodDescriptor_GatewayService_PatchHotel = new grpc.web.MethodDescriptor(
  '/proto.GatewayService/PatchHotel',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_hotel$service_proto_scheme_pb.Hotel,
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  /**
   * @param {!proto.proto.Hotel} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.Hotel,
 *   !proto.common.Empty>}
 */
const methodInfo_GatewayService_PatchHotel = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  /**
   * @param {!proto.proto.Hotel} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.proto.Hotel} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.common.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.common.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.GatewayServiceClient.prototype.patchHotel =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.GatewayService/PatchHotel',
      request,
      metadata || {},
      methodDescriptor_GatewayService_PatchHotel,
      callback);
};


/**
 * @param {!proto.proto.Hotel} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.common.Empty>}
 *     Promise that resolves to the response
 */
proto.proto.GatewayServicePromiseClient.prototype.patchHotel =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.GatewayService/PatchHotel',
      request,
      metadata || {},
      methodDescriptor_GatewayService_PatchHotel);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.common.Empty>}
 */
const methodDescriptor_GatewayService_DeleteHotel = new grpc.web.MethodDescriptor(
  '/proto.GatewayService/DeleteHotel',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.UUID,
 *   !proto.common.Empty>}
 */
const methodInfo_GatewayService_DeleteHotel = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.common.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.common.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.GatewayServiceClient.prototype.deleteHotel =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.GatewayService/DeleteHotel',
      request,
      metadata || {},
      methodDescriptor_GatewayService_DeleteHotel,
      callback);
};


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.common.Empty>}
 *     Promise that resolves to the response
 */
proto.proto.GatewayServicePromiseClient.prototype.deleteHotel =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.GatewayService/DeleteHotel',
      request,
      metadata || {},
      methodDescriptor_GatewayService_DeleteHotel);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.Room,
 *   !proto.common.Empty>}
 */
const methodDescriptor_GatewayService_AddRoom = new grpc.web.MethodDescriptor(
  '/proto.GatewayService/AddRoom',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_hotel$service_proto_scheme_pb.Room,
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  /**
   * @param {!proto.proto.Room} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.Room,
 *   !proto.common.Empty>}
 */
const methodInfo_GatewayService_AddRoom = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  /**
   * @param {!proto.proto.Room} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.proto.Room} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.common.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.common.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.GatewayServiceClient.prototype.addRoom =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.GatewayService/AddRoom',
      request,
      metadata || {},
      methodDescriptor_GatewayService_AddRoom,
      callback);
};


/**
 * @param {!proto.proto.Room} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.common.Empty>}
 *     Promise that resolves to the response
 */
proto.proto.GatewayServicePromiseClient.prototype.addRoom =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.GatewayService/AddRoom',
      request,
      metadata || {},
      methodDescriptor_GatewayService_AddRoom);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.proto.RoomsResponse>}
 */
const methodDescriptor_GatewayService_GetRooms = new grpc.web.MethodDescriptor(
  '/proto.GatewayService/GetRooms',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  internal_pkg_delivery_grpc_hotel$service_proto_scheme_pb.RoomsResponse,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_hotel$service_proto_scheme_pb.RoomsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.UUID,
 *   !proto.proto.RoomsResponse>}
 */
const methodInfo_GatewayService_GetRooms = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_hotel$service_proto_scheme_pb.RoomsResponse,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_hotel$service_proto_scheme_pb.RoomsResponse.deserializeBinary
);


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.RoomsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.RoomsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.GatewayServiceClient.prototype.getRooms =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.GatewayService/GetRooms',
      request,
      metadata || {},
      methodDescriptor_GatewayService_GetRooms,
      callback);
};


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.RoomsResponse>}
 *     Promise that resolves to the response
 */
proto.proto.GatewayServicePromiseClient.prototype.getRooms =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.GatewayService/GetRooms',
      request,
      metadata || {},
      methodDescriptor_GatewayService_GetRooms);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.proto.Room>}
 */
const methodDescriptor_GatewayService_GetRoom = new grpc.web.MethodDescriptor(
  '/proto.GatewayService/GetRoom',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  internal_pkg_delivery_grpc_hotel$service_proto_scheme_pb.Room,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_hotel$service_proto_scheme_pb.Room.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.UUID,
 *   !proto.proto.Room>}
 */
const methodInfo_GatewayService_GetRoom = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_hotel$service_proto_scheme_pb.Room,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_hotel$service_proto_scheme_pb.Room.deserializeBinary
);


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.Room)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.Room>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.GatewayServiceClient.prototype.getRoom =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.GatewayService/GetRoom',
      request,
      metadata || {},
      methodDescriptor_GatewayService_GetRoom,
      callback);
};


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.Room>}
 *     Promise that resolves to the response
 */
proto.proto.GatewayServicePromiseClient.prototype.getRoom =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.GatewayService/GetRoom',
      request,
      metadata || {},
      methodDescriptor_GatewayService_GetRoom);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.Room,
 *   !proto.common.Empty>}
 */
const methodDescriptor_GatewayService_PatchRoom = new grpc.web.MethodDescriptor(
  '/proto.GatewayService/PatchRoom',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_hotel$service_proto_scheme_pb.Room,
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  /**
   * @param {!proto.proto.Room} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.Room,
 *   !proto.common.Empty>}
 */
const methodInfo_GatewayService_PatchRoom = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  /**
   * @param {!proto.proto.Room} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.proto.Room} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.common.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.common.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.GatewayServiceClient.prototype.patchRoom =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.GatewayService/PatchRoom',
      request,
      metadata || {},
      methodDescriptor_GatewayService_PatchRoom,
      callback);
};


/**
 * @param {!proto.proto.Room} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.common.Empty>}
 *     Promise that resolves to the response
 */
proto.proto.GatewayServicePromiseClient.prototype.patchRoom =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.GatewayService/PatchRoom',
      request,
      metadata || {},
      methodDescriptor_GatewayService_PatchRoom);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.common.Empty>}
 */
const methodDescriptor_GatewayService_DeleteRoom = new grpc.web.MethodDescriptor(
  '/proto.GatewayService/DeleteRoom',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.UUID,
 *   !proto.common.Empty>}
 */
const methodInfo_GatewayService_DeleteRoom = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.common.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.common.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.GatewayServiceClient.prototype.deleteRoom =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.GatewayService/DeleteRoom',
      request,
      metadata || {},
      methodDescriptor_GatewayService_DeleteRoom,
      callback);
};


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.common.Empty>}
 *     Promise that resolves to the response
 */
proto.proto.GatewayServicePromiseClient.prototype.deleteRoom =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.GatewayService/DeleteRoom',
      request,
      metadata || {},
      methodDescriptor_GatewayService_DeleteRoom);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.proto.Loyalty>}
 */
const methodDescriptor_GatewayService_GetDiscount = new grpc.web.MethodDescriptor(
  '/proto.GatewayService/GetDiscount',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  internal_pkg_delivery_grpc_loyalty$service_proto_scheme_pb.Loyalty,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_loyalty$service_proto_scheme_pb.Loyalty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.UUID,
 *   !proto.proto.Loyalty>}
 */
const methodInfo_GatewayService_GetDiscount = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_loyalty$service_proto_scheme_pb.Loyalty,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_loyalty$service_proto_scheme_pb.Loyalty.deserializeBinary
);


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.Loyalty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.Loyalty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.GatewayServiceClient.prototype.getDiscount =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.GatewayService/GetDiscount',
      request,
      metadata || {},
      methodDescriptor_GatewayService_GetDiscount,
      callback);
};


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.Loyalty>}
 *     Promise that resolves to the response
 */
proto.proto.GatewayServicePromiseClient.prototype.getDiscount =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.GatewayService/GetDiscount',
      request,
      metadata || {},
      methodDescriptor_GatewayService_GetDiscount);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.proto.Payment>}
 */
const methodDescriptor_GatewayService_GetPayment = new grpc.web.MethodDescriptor(
  '/proto.GatewayService/GetPayment',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  internal_pkg_delivery_grpc_payment$service_proto_scheme_pb.Payment,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_payment$service_proto_scheme_pb.Payment.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.UUID,
 *   !proto.proto.Payment>}
 */
const methodInfo_GatewayService_GetPayment = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_payment$service_proto_scheme_pb.Payment,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_payment$service_proto_scheme_pb.Payment.deserializeBinary
);


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.Payment)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.Payment>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.GatewayServiceClient.prototype.getPayment =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.GatewayService/GetPayment',
      request,
      metadata || {},
      methodDescriptor_GatewayService_GetPayment,
      callback);
};


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.Payment>}
 *     Promise that resolves to the response
 */
proto.proto.GatewayServicePromiseClient.prototype.getPayment =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.GatewayService/GetPayment',
      request,
      metadata || {},
      methodDescriptor_GatewayService_GetPayment);
};


module.exports = proto.proto;

