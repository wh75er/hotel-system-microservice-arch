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
proto.proto.StatServiceClient =
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
proto.proto.StatServicePromiseClient =
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
const methodDescriptor_StatService_GetToken = new grpc.web.MethodDescriptor(
  '/proto.StatService/GetToken',
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
const methodInfo_StatService_GetToken = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.proto.StatServiceClient.prototype.getToken =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.StatService/GetToken',
      request,
      metadata || {},
      methodDescriptor_StatService_GetToken,
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
proto.proto.StatServicePromiseClient.prototype.getToken =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.StatService/GetToken',
      request,
      metadata || {},
      methodDescriptor_StatService_GetToken);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.Empty,
 *   !proto.proto.Stat>}
 */
const methodDescriptor_StatService_GetStat = new grpc.web.MethodDescriptor(
  '/proto.StatService/GetStat',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  proto.proto.Stat,
  /**
   * @param {!proto.common.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Stat.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.Empty,
 *   !proto.proto.Stat>}
 */
const methodInfo_StatService_GetStat = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.Stat,
  /**
   * @param {!proto.common.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Stat.deserializeBinary
);


/**
 * @param {!proto.common.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.Stat)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.Stat>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.StatServiceClient.prototype.getStat =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.StatService/GetStat',
      request,
      metadata || {},
      methodDescriptor_StatService_GetStat,
      callback);
};


/**
 * @param {!proto.common.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.Stat>}
 *     Promise that resolves to the response
 */
proto.proto.StatServicePromiseClient.prototype.getStat =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.StatService/GetStat',
      request,
      metadata || {},
      methodDescriptor_StatService_GetStat);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.Delta,
 *   !proto.common.Empty>}
 */
const methodDescriptor_StatService_UpdateRoomsAmount = new grpc.web.MethodDescriptor(
  '/proto.StatService/UpdateRoomsAmount',
  grpc.web.MethodType.UNARY,
  proto.proto.Delta,
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  /**
   * @param {!proto.proto.Delta} request
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
 *   !proto.proto.Delta,
 *   !proto.common.Empty>}
 */
const methodInfo_StatService_UpdateRoomsAmount = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  /**
   * @param {!proto.proto.Delta} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.proto.Delta} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.common.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.common.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.StatServiceClient.prototype.updateRoomsAmount =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.StatService/UpdateRoomsAmount',
      request,
      metadata || {},
      methodDescriptor_StatService_UpdateRoomsAmount,
      callback);
};


/**
 * @param {!proto.proto.Delta} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.common.Empty>}
 *     Promise that resolves to the response
 */
proto.proto.StatServicePromiseClient.prototype.updateRoomsAmount =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.StatService/UpdateRoomsAmount',
      request,
      metadata || {},
      methodDescriptor_StatService_UpdateRoomsAmount);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.Delta,
 *   !proto.common.Empty>}
 */
const methodDescriptor_StatService_UpdateReservationsAmount = new grpc.web.MethodDescriptor(
  '/proto.StatService/UpdateReservationsAmount',
  grpc.web.MethodType.UNARY,
  proto.proto.Delta,
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  /**
   * @param {!proto.proto.Delta} request
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
 *   !proto.proto.Delta,
 *   !proto.common.Empty>}
 */
const methodInfo_StatService_UpdateReservationsAmount = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  /**
   * @param {!proto.proto.Delta} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.proto.Delta} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.common.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.common.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.StatServiceClient.prototype.updateReservationsAmount =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.StatService/UpdateReservationsAmount',
      request,
      metadata || {},
      methodDescriptor_StatService_UpdateReservationsAmount,
      callback);
};


/**
 * @param {!proto.proto.Delta} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.common.Empty>}
 *     Promise that resolves to the response
 */
proto.proto.StatServicePromiseClient.prototype.updateReservationsAmount =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.StatService/UpdateReservationsAmount',
      request,
      metadata || {},
      methodDescriptor_StatService_UpdateReservationsAmount);
};


module.exports = proto.proto;

