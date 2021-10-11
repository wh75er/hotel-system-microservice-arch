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
proto.proto.LoyaltyServiceClient =
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
proto.proto.LoyaltyServicePromiseClient =
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
const methodDescriptor_LoyaltyService_GetToken = new grpc.web.MethodDescriptor(
  '/proto.LoyaltyService/GetToken',
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
const methodInfo_LoyaltyService_GetToken = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.proto.LoyaltyServiceClient.prototype.getToken =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.LoyaltyService/GetToken',
      request,
      metadata || {},
      methodDescriptor_LoyaltyService_GetToken,
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
proto.proto.LoyaltyServicePromiseClient.prototype.getToken =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.LoyaltyService/GetToken',
      request,
      metadata || {},
      methodDescriptor_LoyaltyService_GetToken);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.proto.Loyalty>}
 */
const methodDescriptor_LoyaltyService_GetDiscount = new grpc.web.MethodDescriptor(
  '/proto.LoyaltyService/GetDiscount',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  proto.proto.Loyalty,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Loyalty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.UUID,
 *   !proto.proto.Loyalty>}
 */
const methodInfo_LoyaltyService_GetDiscount = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.Loyalty,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Loyalty.deserializeBinary
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
proto.proto.LoyaltyServiceClient.prototype.getDiscount =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.LoyaltyService/GetDiscount',
      request,
      metadata || {},
      methodDescriptor_LoyaltyService_GetDiscount,
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
proto.proto.LoyaltyServicePromiseClient.prototype.getDiscount =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.LoyaltyService/GetDiscount',
      request,
      metadata || {},
      methodDescriptor_LoyaltyService_GetDiscount);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.common.Empty>}
 */
const methodDescriptor_LoyaltyService_AddUser = new grpc.web.MethodDescriptor(
  '/proto.LoyaltyService/AddUser',
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
const methodInfo_LoyaltyService_AddUser = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.proto.LoyaltyServiceClient.prototype.addUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.LoyaltyService/AddUser',
      request,
      metadata || {},
      methodDescriptor_LoyaltyService_AddUser,
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
proto.proto.LoyaltyServicePromiseClient.prototype.addUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.LoyaltyService/AddUser',
      request,
      metadata || {},
      methodDescriptor_LoyaltyService_AddUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.UpdateDiscountRequest,
 *   !proto.common.Empty>}
 */
const methodDescriptor_LoyaltyService_UpdateDiscount = new grpc.web.MethodDescriptor(
  '/proto.LoyaltyService/UpdateDiscount',
  grpc.web.MethodType.UNARY,
  proto.proto.UpdateDiscountRequest,
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  /**
   * @param {!proto.proto.UpdateDiscountRequest} request
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
 *   !proto.proto.UpdateDiscountRequest,
 *   !proto.common.Empty>}
 */
const methodInfo_LoyaltyService_UpdateDiscount = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  /**
   * @param {!proto.proto.UpdateDiscountRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.proto.UpdateDiscountRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.common.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.common.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.LoyaltyServiceClient.prototype.updateDiscount =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.LoyaltyService/UpdateDiscount',
      request,
      metadata || {},
      methodDescriptor_LoyaltyService_UpdateDiscount,
      callback);
};


/**
 * @param {!proto.proto.UpdateDiscountRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.common.Empty>}
 *     Promise that resolves to the response
 */
proto.proto.LoyaltyServicePromiseClient.prototype.updateDiscount =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.LoyaltyService/UpdateDiscount',
      request,
      metadata || {},
      methodDescriptor_LoyaltyService_UpdateDiscount);
};


module.exports = proto.proto;

