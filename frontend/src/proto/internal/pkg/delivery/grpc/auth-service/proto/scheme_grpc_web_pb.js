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
proto.proto.AuthServiceClient =
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
proto.proto.AuthServicePromiseClient =
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
const methodDescriptor_AuthService_GetToken = new grpc.web.MethodDescriptor(
  '/proto.AuthService/GetToken',
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
const methodInfo_AuthService_GetToken = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.proto.AuthServiceClient.prototype.getToken =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.AuthService/GetToken',
      request,
      metadata || {},
      methodDescriptor_AuthService_GetToken,
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
proto.proto.AuthServicePromiseClient.prototype.getToken =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.AuthService/GetToken',
      request,
      metadata || {},
      methodDescriptor_AuthService_GetToken);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.User,
 *   !proto.common.Empty>}
 */
const methodDescriptor_AuthService_AddUser = new grpc.web.MethodDescriptor(
  '/proto.AuthService/AddUser',
  grpc.web.MethodType.UNARY,
  proto.proto.User,
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
const methodInfo_AuthService_AddUser = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.proto.AuthServiceClient.prototype.addUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.AuthService/AddUser',
      request,
      metadata || {},
      methodDescriptor_AuthService_AddUser,
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
proto.proto.AuthServicePromiseClient.prototype.addUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.AuthService/AddUser',
      request,
      metadata || {},
      methodDescriptor_AuthService_AddUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.proto.User>}
 */
const methodDescriptor_AuthService_GetUser = new grpc.web.MethodDescriptor(
  '/proto.AuthService/GetUser',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  proto.proto.User,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.User.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.UUID,
 *   !proto.proto.User>}
 */
const methodInfo_AuthService_GetUser = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.User,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.User.deserializeBinary
);


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.User)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.User>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.AuthServiceClient.prototype.getUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.AuthService/GetUser',
      request,
      metadata || {},
      methodDescriptor_AuthService_GetUser,
      callback);
};


/**
 * @param {!proto.common.UUID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.User>}
 *     Promise that resolves to the response
 */
proto.proto.AuthServicePromiseClient.prototype.getUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.AuthService/GetUser',
      request,
      metadata || {},
      methodDescriptor_AuthService_GetUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.User,
 *   !proto.common.Token>}
 */
const methodDescriptor_AuthService_Login = new grpc.web.MethodDescriptor(
  '/proto.AuthService/Login',
  grpc.web.MethodType.UNARY,
  proto.proto.User,
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
const methodInfo_AuthService_Login = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.proto.AuthServiceClient.prototype.login =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.AuthService/Login',
      request,
      metadata || {},
      methodDescriptor_AuthService_Login,
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
proto.proto.AuthServicePromiseClient.prototype.login =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.AuthService/Login',
      request,
      metadata || {},
      methodDescriptor_AuthService_Login);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.Token,
 *   !proto.proto.Role>}
 */
const methodDescriptor_AuthService_CheckAuth = new grpc.web.MethodDescriptor(
  '/proto.AuthService/CheckAuth',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.Token,
  proto.proto.Role,
  /**
   * @param {!proto.common.Token} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Role.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.Token,
 *   !proto.proto.Role>}
 */
const methodInfo_AuthService_CheckAuth = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.Role,
  /**
   * @param {!proto.common.Token} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Role.deserializeBinary
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
proto.proto.AuthServiceClient.prototype.checkAuth =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.AuthService/CheckAuth',
      request,
      metadata || {},
      methodDescriptor_AuthService_CheckAuth,
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
proto.proto.AuthServicePromiseClient.prototype.checkAuth =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.AuthService/CheckAuth',
      request,
      metadata || {},
      methodDescriptor_AuthService_CheckAuth);
};


module.exports = proto.proto;

