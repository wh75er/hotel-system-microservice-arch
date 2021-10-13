// source: internal/pkg/delivery/grpc/hotel-service/proto/scheme.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var internal_pkg_delivery_grpc_commonProto_common_pb = require('../../../../../../internal/pkg/delivery/grpc/commonProto/common_pb.js');
goog.object.extend(proto, internal_pkg_delivery_grpc_commonProto_common_pb);
goog.exportSymbol('proto.proto.Hotel', null, global);
goog.exportSymbol('proto.proto.HotelsResponse', null, global);
goog.exportSymbol('proto.proto.Room', null, global);
goog.exportSymbol('proto.proto.RoomsResponse', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.proto.Room = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.proto.Room.repeatedFields_, null);
};
goog.inherits(proto.proto.Room, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.proto.Room.displayName = 'proto.proto.Room';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.proto.Hotel = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.proto.Hotel.repeatedFields_, null);
};
goog.inherits(proto.proto.Hotel, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.proto.Hotel.displayName = 'proto.proto.Hotel';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.proto.HotelsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.proto.HotelsResponse.repeatedFields_, null);
};
goog.inherits(proto.proto.HotelsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.proto.HotelsResponse.displayName = 'proto.proto.HotelsResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.proto.RoomsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.proto.RoomsResponse.repeatedFields_, null);
};
goog.inherits(proto.proto.RoomsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.proto.RoomsResponse.displayName = 'proto.proto.RoomsResponse';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.proto.Room.repeatedFields_ = [7];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.proto.Room.prototype.toObject = function(opt_includeInstance) {
  return proto.proto.Room.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.proto.Room} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.Room.toObject = function(includeInstance, msg) {
  var f, obj = {
    roomtype: jspb.Message.getFieldWithDefault(msg, 1, ""),
    amount: jspb.Message.getFieldWithDefault(msg, 2, 0),
    beds: jspb.Message.getFieldWithDefault(msg, 3, 0),
    hoteluuid: jspb.Message.getFieldWithDefault(msg, 4, ""),
    roomuuid: jspb.Message.getFieldWithDefault(msg, 5, ""),
    creationdate: jspb.Message.getFieldWithDefault(msg, 6, 0),
    offersList: (f = jspb.Message.getRepeatedField(msg, 7)) == null ? undefined : f,
    nightprice: jspb.Message.getFloatingPointFieldWithDefault(msg, 8, 0.0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.proto.Room}
 */
proto.proto.Room.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.proto.Room;
  return proto.proto.Room.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.Room} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.Room}
 */
proto.proto.Room.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setRoomtype(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setAmount(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setBeds(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setHoteluuid(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setRoomuuid(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setCreationdate(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.addOffers(value);
      break;
    case 8:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setNightprice(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.proto.Room.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.proto.Room.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.Room} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.Room.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRoomtype();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAmount();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
  f = message.getBeds();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
  f = message.getHoteluuid();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getRoomuuid();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getCreationdate();
  if (f !== 0) {
    writer.writeInt64(
      6,
      f
    );
  }
  f = message.getOffersList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      7,
      f
    );
  }
  f = message.getNightprice();
  if (f !== 0.0) {
    writer.writeFloat(
      8,
      f
    );
  }
};


/**
 * optional string RoomType = 1;
 * @return {string}
 */
proto.proto.Room.prototype.getRoomtype = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.Room} returns this
 */
proto.proto.Room.prototype.setRoomtype = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional int64 Amount = 2;
 * @return {number}
 */
proto.proto.Room.prototype.getAmount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.proto.Room} returns this
 */
proto.proto.Room.prototype.setAmount = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional int64 Beds = 3;
 * @return {number}
 */
proto.proto.Room.prototype.getBeds = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.proto.Room} returns this
 */
proto.proto.Room.prototype.setBeds = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional string hotelUuid = 4;
 * @return {string}
 */
proto.proto.Room.prototype.getHoteluuid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.Room} returns this
 */
proto.proto.Room.prototype.setHoteluuid = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string roomUuid = 5;
 * @return {string}
 */
proto.proto.Room.prototype.getRoomuuid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.Room} returns this
 */
proto.proto.Room.prototype.setRoomuuid = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional int64 CreationDate = 6;
 * @return {number}
 */
proto.proto.Room.prototype.getCreationdate = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/**
 * @param {number} value
 * @return {!proto.proto.Room} returns this
 */
proto.proto.Room.prototype.setCreationdate = function(value) {
  return jspb.Message.setProto3IntField(this, 6, value);
};


/**
 * repeated string Offers = 7;
 * @return {!Array<string>}
 */
proto.proto.Room.prototype.getOffersList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 7));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.proto.Room} returns this
 */
proto.proto.Room.prototype.setOffersList = function(value) {
  return jspb.Message.setField(this, 7, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.proto.Room} returns this
 */
proto.proto.Room.prototype.addOffers = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 7, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.proto.Room} returns this
 */
proto.proto.Room.prototype.clearOffersList = function() {
  return this.setOffersList([]);
};


/**
 * optional float NightPrice = 8;
 * @return {number}
 */
proto.proto.Room.prototype.getNightprice = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 8, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.proto.Room} returns this
 */
proto.proto.Room.prototype.setNightprice = function(value) {
  return jspb.Message.setProto3FloatField(this, 8, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.proto.Hotel.repeatedFields_ = [10];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.proto.Hotel.prototype.toObject = function(opt_includeInstance) {
  return proto.proto.Hotel.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.proto.Hotel} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.Hotel.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, ""),
    hoteluuid: jspb.Message.getFieldWithDefault(msg, 2, ""),
    description: jspb.Message.getFieldWithDefault(msg, 4, ""),
    country: jspb.Message.getFieldWithDefault(msg, 5, ""),
    city: jspb.Message.getFieldWithDefault(msg, 6, ""),
    address: jspb.Message.getFieldWithDefault(msg, 7, ""),
    isready: jspb.Message.getBooleanFieldWithDefault(msg, 8, false),
    creationdate: jspb.Message.getFieldWithDefault(msg, 9, 0),
    roomsList: jspb.Message.toObjectList(msg.getRoomsList(),
    proto.proto.Room.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.proto.Hotel}
 */
proto.proto.Hotel.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.proto.Hotel;
  return proto.proto.Hotel.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.Hotel} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.Hotel}
 */
proto.proto.Hotel.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setHoteluuid(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setDescription(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setCountry(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setCity(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setAddress(value);
      break;
    case 8:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIsready(value);
      break;
    case 9:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setCreationdate(value);
      break;
    case 10:
      var value = new proto.proto.Room;
      reader.readMessage(value,proto.proto.Room.deserializeBinaryFromReader);
      msg.addRooms(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.proto.Hotel.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.proto.Hotel.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.Hotel} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.Hotel.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getHoteluuid();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getDescription();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getCountry();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getCity();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getAddress();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getIsready();
  if (f) {
    writer.writeBool(
      8,
      f
    );
  }
  f = message.getCreationdate();
  if (f !== 0) {
    writer.writeInt64(
      9,
      f
    );
  }
  f = message.getRoomsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      10,
      f,
      proto.proto.Room.serializeBinaryToWriter
    );
  }
};


/**
 * optional string Name = 1;
 * @return {string}
 */
proto.proto.Hotel.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.Hotel} returns this
 */
proto.proto.Hotel.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string HotelUuid = 2;
 * @return {string}
 */
proto.proto.Hotel.prototype.getHoteluuid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.Hotel} returns this
 */
proto.proto.Hotel.prototype.setHoteluuid = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string Description = 4;
 * @return {string}
 */
proto.proto.Hotel.prototype.getDescription = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.Hotel} returns this
 */
proto.proto.Hotel.prototype.setDescription = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string Country = 5;
 * @return {string}
 */
proto.proto.Hotel.prototype.getCountry = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.Hotel} returns this
 */
proto.proto.Hotel.prototype.setCountry = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string City = 6;
 * @return {string}
 */
proto.proto.Hotel.prototype.getCity = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.Hotel} returns this
 */
proto.proto.Hotel.prototype.setCity = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional string Address = 7;
 * @return {string}
 */
proto.proto.Hotel.prototype.getAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.Hotel} returns this
 */
proto.proto.Hotel.prototype.setAddress = function(value) {
  return jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional bool IsReady = 8;
 * @return {boolean}
 */
proto.proto.Hotel.prototype.getIsready = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 8, false));
};


/**
 * @param {boolean} value
 * @return {!proto.proto.Hotel} returns this
 */
proto.proto.Hotel.prototype.setIsready = function(value) {
  return jspb.Message.setProto3BooleanField(this, 8, value);
};


/**
 * optional int64 CreationDate = 9;
 * @return {number}
 */
proto.proto.Hotel.prototype.getCreationdate = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 9, 0));
};


/**
 * @param {number} value
 * @return {!proto.proto.Hotel} returns this
 */
proto.proto.Hotel.prototype.setCreationdate = function(value) {
  return jspb.Message.setProto3IntField(this, 9, value);
};


/**
 * repeated Room rooms = 10;
 * @return {!Array<!proto.proto.Room>}
 */
proto.proto.Hotel.prototype.getRoomsList = function() {
  return /** @type{!Array<!proto.proto.Room>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.proto.Room, 10));
};


/**
 * @param {!Array<!proto.proto.Room>} value
 * @return {!proto.proto.Hotel} returns this
*/
proto.proto.Hotel.prototype.setRoomsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 10, value);
};


/**
 * @param {!proto.proto.Room=} opt_value
 * @param {number=} opt_index
 * @return {!proto.proto.Room}
 */
proto.proto.Hotel.prototype.addRooms = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 10, opt_value, proto.proto.Room, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.proto.Hotel} returns this
 */
proto.proto.Hotel.prototype.clearRoomsList = function() {
  return this.setRoomsList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.proto.HotelsResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.proto.HotelsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.proto.HotelsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.proto.HotelsResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.HotelsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    hotelsList: jspb.Message.toObjectList(msg.getHotelsList(),
    proto.proto.Hotel.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.proto.HotelsResponse}
 */
proto.proto.HotelsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.proto.HotelsResponse;
  return proto.proto.HotelsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.HotelsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.HotelsResponse}
 */
proto.proto.HotelsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.proto.Hotel;
      reader.readMessage(value,proto.proto.Hotel.deserializeBinaryFromReader);
      msg.addHotels(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.proto.HotelsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.proto.HotelsResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.HotelsResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.HotelsResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHotelsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.proto.Hotel.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Hotel Hotels = 1;
 * @return {!Array<!proto.proto.Hotel>}
 */
proto.proto.HotelsResponse.prototype.getHotelsList = function() {
  return /** @type{!Array<!proto.proto.Hotel>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.proto.Hotel, 1));
};


/**
 * @param {!Array<!proto.proto.Hotel>} value
 * @return {!proto.proto.HotelsResponse} returns this
*/
proto.proto.HotelsResponse.prototype.setHotelsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.proto.Hotel=} opt_value
 * @param {number=} opt_index
 * @return {!proto.proto.Hotel}
 */
proto.proto.HotelsResponse.prototype.addHotels = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.proto.Hotel, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.proto.HotelsResponse} returns this
 */
proto.proto.HotelsResponse.prototype.clearHotelsList = function() {
  return this.setHotelsList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.proto.RoomsResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.proto.RoomsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.proto.RoomsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.proto.RoomsResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.RoomsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    roomsList: jspb.Message.toObjectList(msg.getRoomsList(),
    proto.proto.Room.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.proto.RoomsResponse}
 */
proto.proto.RoomsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.proto.RoomsResponse;
  return proto.proto.RoomsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.RoomsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.RoomsResponse}
 */
proto.proto.RoomsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.proto.Room;
      reader.readMessage(value,proto.proto.Room.deserializeBinaryFromReader);
      msg.addRooms(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.proto.RoomsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.proto.RoomsResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.RoomsResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.RoomsResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRoomsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.proto.Room.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Room Rooms = 1;
 * @return {!Array<!proto.proto.Room>}
 */
proto.proto.RoomsResponse.prototype.getRoomsList = function() {
  return /** @type{!Array<!proto.proto.Room>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.proto.Room, 1));
};


/**
 * @param {!Array<!proto.proto.Room>} value
 * @return {!proto.proto.RoomsResponse} returns this
*/
proto.proto.RoomsResponse.prototype.setRoomsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.proto.Room=} opt_value
 * @param {number=} opt_index
 * @return {!proto.proto.Room}
 */
proto.proto.RoomsResponse.prototype.addRooms = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.proto.Room, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.proto.RoomsResponse} returns this
 */
proto.proto.RoomsResponse.prototype.clearRoomsList = function() {
  return this.setRoomsList([]);
};


goog.object.extend(exports, proto.proto);
