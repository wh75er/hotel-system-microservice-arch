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
proto.proto.ReservationServiceClient =
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
proto.proto.ReservationServicePromiseClient =
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
 *   !proto.common.Credentials,
 *   !proto.common.Token>}
 */
const methodDescriptor_ReservationService_GetToken = new grpc.web.MethodDescriptor(
  '/proto.ReservationService/GetToken',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.Credentials,
  internal_pkg_delivery_grpc_commonProto_common_pb.Token,
  /**
   * @param {!proto.common.Credentials} request
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
 *   !proto.common.Credentials,
 *   !proto.common.Token>}
 */
const methodInfo_ReservationService_GetToken = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_commonProto_common_pb.Token,
  /**
   * @param {!proto.common.Credentials} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.Token.deserializeBinary
);


/**
 * @param {!proto.common.Credentials} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.common.Token)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.common.Token>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.ReservationServiceClient.prototype.getToken =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.ReservationService/GetToken',
      request,
      metadata || {},
      methodDescriptor_ReservationService_GetToken,
      callback);
};


/**
 * @param {!proto.common.Credentials} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.common.Token>}
 *     Promise that resolves to the response
 */
proto.proto.ReservationServicePromiseClient.prototype.getToken =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.ReservationService/GetToken',
      request,
      metadata || {},
      methodDescriptor_ReservationService_GetToken);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.Reservation,
 *   !proto.common.UUID>}
 */
const methodDescriptor_ReservationService_AddReservation = new grpc.web.MethodDescriptor(
  '/proto.ReservationService/AddReservation',
  grpc.web.MethodType.UNARY,
  proto.proto.Reservation,
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
const methodInfo_ReservationService_AddReservation = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.proto.ReservationServiceClient.prototype.addReservation =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.ReservationService/AddReservation',
      request,
      metadata || {},
      methodDescriptor_ReservationService_AddReservation,
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
proto.proto.ReservationServicePromiseClient.prototype.addReservation =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.ReservationService/AddReservation',
      request,
      metadata || {},
      methodDescriptor_ReservationService_AddReservation);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.common.Empty>}
 */
const methodDescriptor_ReservationService_CancelReservation = new grpc.web.MethodDescriptor(
  '/proto.ReservationService/CancelReservation',
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
const methodInfo_ReservationService_CancelReservation = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.proto.ReservationServiceClient.prototype.cancelReservation =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.ReservationService/CancelReservation',
      request,
      metadata || {},
      methodDescriptor_ReservationService_CancelReservation,
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
proto.proto.ReservationServicePromiseClient.prototype.cancelReservation =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.ReservationService/CancelReservation',
      request,
      metadata || {},
      methodDescriptor_ReservationService_CancelReservation);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.proto.Reservation>}
 */
const methodDescriptor_ReservationService_GetReservation = new grpc.web.MethodDescriptor(
  '/proto.ReservationService/GetReservation',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  proto.proto.Reservation,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Reservation.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.UUID,
 *   !proto.proto.Reservation>}
 */
const methodInfo_ReservationService_GetReservation = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.Reservation,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Reservation.deserializeBinary
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
proto.proto.ReservationServiceClient.prototype.getReservation =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.ReservationService/GetReservation',
      request,
      metadata || {},
      methodDescriptor_ReservationService_GetReservation,
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
proto.proto.ReservationServicePromiseClient.prototype.getReservation =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.ReservationService/GetReservation',
      request,
      metadata || {},
      methodDescriptor_ReservationService_GetReservation);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.proto.Reservations>}
 */
const methodDescriptor_ReservationService_GetReservationsByUser = new grpc.web.MethodDescriptor(
  '/proto.ReservationService/GetReservationsByUser',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  proto.proto.Reservations,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Reservations.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.UUID,
 *   !proto.proto.Reservations>}
 */
const methodInfo_ReservationService_GetReservationsByUser = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.Reservations,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Reservations.deserializeBinary
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
proto.proto.ReservationServiceClient.prototype.getReservationsByUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.ReservationService/GetReservationsByUser',
      request,
      metadata || {},
      methodDescriptor_ReservationService_GetReservationsByUser,
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
proto.proto.ReservationServicePromiseClient.prototype.getReservationsByUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.ReservationService/GetReservationsByUser',
      request,
      metadata || {},
      methodDescriptor_ReservationService_GetReservationsByUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.common.UUID>}
 */
const methodDescriptor_ReservationService_CreatePayment = new grpc.web.MethodDescriptor(
  '/proto.ReservationService/CreatePayment',
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
const methodInfo_ReservationService_CreatePayment = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.proto.ReservationServiceClient.prototype.createPayment =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.ReservationService/CreatePayment',
      request,
      metadata || {},
      methodDescriptor_ReservationService_CreatePayment,
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
proto.proto.ReservationServicePromiseClient.prototype.createPayment =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.ReservationService/CreatePayment',
      request,
      metadata || {},
      methodDescriptor_ReservationService_CreatePayment);
};


module.exports = proto.proto;

