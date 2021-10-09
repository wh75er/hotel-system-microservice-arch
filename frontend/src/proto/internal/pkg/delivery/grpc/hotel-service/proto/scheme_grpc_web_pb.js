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
proto.proto.HotelServiceClient =
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
proto.proto.HotelServicePromiseClient =
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
const methodDescriptor_HotelService_GetToken = new grpc.web.MethodDescriptor(
  '/proto.HotelService/GetToken',
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
const methodInfo_HotelService_GetToken = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.proto.HotelServiceClient.prototype.getToken =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.HotelService/GetToken',
      request,
      metadata || {},
      methodDescriptor_HotelService_GetToken,
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
proto.proto.HotelServicePromiseClient.prototype.getToken =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.HotelService/GetToken',
      request,
      metadata || {},
      methodDescriptor_HotelService_GetToken);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.Hotel,
 *   !proto.common.Empty>}
 */
const methodDescriptor_HotelService_AddHotel = new grpc.web.MethodDescriptor(
  '/proto.HotelService/AddHotel',
  grpc.web.MethodType.UNARY,
  proto.proto.Hotel,
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
const methodInfo_HotelService_AddHotel = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.proto.HotelServiceClient.prototype.addHotel =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.HotelService/AddHotel',
      request,
      metadata || {},
      methodDescriptor_HotelService_AddHotel,
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
proto.proto.HotelServicePromiseClient.prototype.addHotel =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.HotelService/AddHotel',
      request,
      metadata || {},
      methodDescriptor_HotelService_AddHotel);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.proto.Hotel>}
 */
const methodDescriptor_HotelService_GetHotel = new grpc.web.MethodDescriptor(
  '/proto.HotelService/GetHotel',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  proto.proto.Hotel,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Hotel.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.UUID,
 *   !proto.proto.Hotel>}
 */
const methodInfo_HotelService_GetHotel = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.Hotel,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Hotel.deserializeBinary
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
proto.proto.HotelServiceClient.prototype.getHotel =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.HotelService/GetHotel',
      request,
      metadata || {},
      methodDescriptor_HotelService_GetHotel,
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
proto.proto.HotelServicePromiseClient.prototype.getHotel =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.HotelService/GetHotel',
      request,
      metadata || {},
      methodDescriptor_HotelService_GetHotel);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.Empty,
 *   !proto.proto.HotelsResponse>}
 */
const methodDescriptor_HotelService_GetHotels = new grpc.web.MethodDescriptor(
  '/proto.HotelService/GetHotels',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.Empty,
  proto.proto.HotelsResponse,
  /**
   * @param {!proto.common.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.HotelsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.Empty,
 *   !proto.proto.HotelsResponse>}
 */
const methodInfo_HotelService_GetHotels = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.HotelsResponse,
  /**
   * @param {!proto.common.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.HotelsResponse.deserializeBinary
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
proto.proto.HotelServiceClient.prototype.getHotels =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.HotelService/GetHotels',
      request,
      metadata || {},
      methodDescriptor_HotelService_GetHotels,
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
proto.proto.HotelServicePromiseClient.prototype.getHotels =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.HotelService/GetHotels',
      request,
      metadata || {},
      methodDescriptor_HotelService_GetHotels);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.Hotel,
 *   !proto.common.Empty>}
 */
const methodDescriptor_HotelService_PatchHotel = new grpc.web.MethodDescriptor(
  '/proto.HotelService/PatchHotel',
  grpc.web.MethodType.UNARY,
  proto.proto.Hotel,
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
const methodInfo_HotelService_PatchHotel = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.proto.HotelServiceClient.prototype.patchHotel =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.HotelService/PatchHotel',
      request,
      metadata || {},
      methodDescriptor_HotelService_PatchHotel,
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
proto.proto.HotelServicePromiseClient.prototype.patchHotel =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.HotelService/PatchHotel',
      request,
      metadata || {},
      methodDescriptor_HotelService_PatchHotel);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.common.Empty>}
 */
const methodDescriptor_HotelService_DeleteHotel = new grpc.web.MethodDescriptor(
  '/proto.HotelService/DeleteHotel',
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
const methodInfo_HotelService_DeleteHotel = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.proto.HotelServiceClient.prototype.deleteHotel =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.HotelService/DeleteHotel',
      request,
      metadata || {},
      methodDescriptor_HotelService_DeleteHotel,
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
proto.proto.HotelServicePromiseClient.prototype.deleteHotel =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.HotelService/DeleteHotel',
      request,
      metadata || {},
      methodDescriptor_HotelService_DeleteHotel);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.Room,
 *   !proto.common.Empty>}
 */
const methodDescriptor_HotelService_AddRoom = new grpc.web.MethodDescriptor(
  '/proto.HotelService/AddRoom',
  grpc.web.MethodType.UNARY,
  proto.proto.Room,
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
const methodInfo_HotelService_AddRoom = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.proto.HotelServiceClient.prototype.addRoom =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.HotelService/AddRoom',
      request,
      metadata || {},
      methodDescriptor_HotelService_AddRoom,
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
proto.proto.HotelServicePromiseClient.prototype.addRoom =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.HotelService/AddRoom',
      request,
      metadata || {},
      methodDescriptor_HotelService_AddRoom);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.proto.RoomsResponse>}
 */
const methodDescriptor_HotelService_GetRooms = new grpc.web.MethodDescriptor(
  '/proto.HotelService/GetRooms',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  proto.proto.RoomsResponse,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.RoomsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.UUID,
 *   !proto.proto.RoomsResponse>}
 */
const methodInfo_HotelService_GetRooms = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.RoomsResponse,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.RoomsResponse.deserializeBinary
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
proto.proto.HotelServiceClient.prototype.getRooms =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.HotelService/GetRooms',
      request,
      metadata || {},
      methodDescriptor_HotelService_GetRooms,
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
proto.proto.HotelServicePromiseClient.prototype.getRooms =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.HotelService/GetRooms',
      request,
      metadata || {},
      methodDescriptor_HotelService_GetRooms);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.proto.Room>}
 */
const methodDescriptor_HotelService_GetRoom = new grpc.web.MethodDescriptor(
  '/proto.HotelService/GetRoom',
  grpc.web.MethodType.UNARY,
  internal_pkg_delivery_grpc_commonProto_common_pb.UUID,
  proto.proto.Room,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Room.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.common.UUID,
 *   !proto.proto.Room>}
 */
const methodInfo_HotelService_GetRoom = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.Room,
  /**
   * @param {!proto.common.UUID} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Room.deserializeBinary
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
proto.proto.HotelServiceClient.prototype.getRoom =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.HotelService/GetRoom',
      request,
      metadata || {},
      methodDescriptor_HotelService_GetRoom,
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
proto.proto.HotelServicePromiseClient.prototype.getRoom =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.HotelService/GetRoom',
      request,
      metadata || {},
      methodDescriptor_HotelService_GetRoom);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.common.Empty>}
 */
const methodDescriptor_HotelService_TakeRoom = new grpc.web.MethodDescriptor(
  '/proto.HotelService/TakeRoom',
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
const methodInfo_HotelService_TakeRoom = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.proto.HotelServiceClient.prototype.takeRoom =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.HotelService/TakeRoom',
      request,
      metadata || {},
      methodDescriptor_HotelService_TakeRoom,
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
proto.proto.HotelServicePromiseClient.prototype.takeRoom =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.HotelService/TakeRoom',
      request,
      metadata || {},
      methodDescriptor_HotelService_TakeRoom);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.common.Empty>}
 */
const methodDescriptor_HotelService_DismissRoom = new grpc.web.MethodDescriptor(
  '/proto.HotelService/DismissRoom',
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
const methodInfo_HotelService_DismissRoom = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.proto.HotelServiceClient.prototype.dismissRoom =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.HotelService/DismissRoom',
      request,
      metadata || {},
      methodDescriptor_HotelService_DismissRoom,
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
proto.proto.HotelServicePromiseClient.prototype.dismissRoom =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.HotelService/DismissRoom',
      request,
      metadata || {},
      methodDescriptor_HotelService_DismissRoom);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.Room,
 *   !proto.common.Empty>}
 */
const methodDescriptor_HotelService_PatchRoom = new grpc.web.MethodDescriptor(
  '/proto.HotelService/PatchRoom',
  grpc.web.MethodType.UNARY,
  proto.proto.Room,
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
const methodInfo_HotelService_PatchRoom = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.proto.HotelServiceClient.prototype.patchRoom =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.HotelService/PatchRoom',
      request,
      metadata || {},
      methodDescriptor_HotelService_PatchRoom,
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
proto.proto.HotelServicePromiseClient.prototype.patchRoom =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.HotelService/PatchRoom',
      request,
      metadata || {},
      methodDescriptor_HotelService_PatchRoom);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.common.UUID,
 *   !proto.common.Empty>}
 */
const methodDescriptor_HotelService_DeleteRoom = new grpc.web.MethodDescriptor(
  '/proto.HotelService/DeleteRoom',
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
const methodInfo_HotelService_DeleteRoom = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.proto.HotelServiceClient.prototype.deleteRoom =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.HotelService/DeleteRoom',
      request,
      metadata || {},
      methodDescriptor_HotelService_DeleteRoom,
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
proto.proto.HotelServicePromiseClient.prototype.deleteRoom =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.HotelService/DeleteRoom',
      request,
      metadata || {},
      methodDescriptor_HotelService_DeleteRoom);
};


module.exports = proto.proto;

