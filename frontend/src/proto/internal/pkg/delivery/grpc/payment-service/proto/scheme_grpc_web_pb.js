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
proto.proto.PaymentServiceClient =
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
proto.proto.PaymentServicePromiseClient =
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
const methodDescriptor_PaymentService_GetToken = new grpc.web.MethodDescriptor(
  '/proto.PaymentService/GetToken',
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
const methodInfo_PaymentService_GetToken = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.proto.PaymentServiceClient.prototype.getToken =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.PaymentService/GetToken',
      request,
      metadata || {},
      methodDescriptor_PaymentService_GetToken,
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
proto.proto.PaymentServicePromiseClient.prototype.getToken =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.PaymentService/GetToken',
      request,
      metadata || {},
      methodDescriptor_PaymentService_GetToken);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.CreatePaymentRequest,
 *   !proto.common.UUID>}
 */
const methodDescriptor_PaymentService_CreatePayment = new grpc.web.MethodDescriptor(
  '/proto.PaymentService/CreatePayment',
  grpc.web.MethodType.UNARY,
  proto.proto.CreatePaymentRequest,
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  /**
   * @param {!proto.proto.CreatePaymentRequest} request
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
 *   !proto.proto.CreatePaymentRequest,
 *   !proto.common.UUID>}
 */
const methodInfo_PaymentService_CreatePayment = new grpc.web.AbstractClientBase.MethodInfo(
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  /**
   * @param {!proto.proto.CreatePaymentRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID.deserializeBinary
);


/**
 * @param {!proto.proto.CreatePaymentRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.common.UUID)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.common.UUID>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.PaymentServiceClient.prototype.createPayment =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.PaymentService/CreatePayment',
      request,
      metadata || {},
      methodDescriptor_PaymentService_CreatePayment,
      callback);
};


/**
 * @param {!proto.proto.CreatePaymentRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.common.UUID>}
 *     Promise that resolves to the response
 */
proto.proto.PaymentServicePromiseClient.prototype.createPayment =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.PaymentService/CreatePayment',
      request,
      metadata || {},
      methodDescriptor_PaymentService_CreatePayment);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.proto.Payment>}
 */
const methodDescriptor_PaymentService_GetPayment = new grpc.web.MethodDescriptor(
  '/proto.PaymentService/GetPayment',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  proto.proto.Payment,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Payment.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.UUID,
 *   !proto.proto.Payment>}
 */
const methodInfo_PaymentService_GetPayment = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.Payment,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Payment.deserializeBinary
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
proto.proto.PaymentServiceClient.prototype.getPayment =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.PaymentService/GetPayment',
      request,
      metadata || {},
      methodDescriptor_PaymentService_GetPayment,
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
proto.proto.PaymentServicePromiseClient.prototype.getPayment =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.PaymentService/GetPayment',
      request,
      metadata || {},
      methodDescriptor_PaymentService_GetPayment);
};


module.exports = proto.proto;

