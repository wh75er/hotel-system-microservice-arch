// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: internal/pkg/delivery/grpc/loyalty-service/proto/scheme.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	commonProto "hotel-booking-system/internal/pkg/delivery/grpc/commonProto"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpdateDiscountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUid      *commonProto.UUID `protobuf:"bytes,1,opt,name=userUid,proto3" json:"userUid,omitempty"`
	Contribution float32           `protobuf:"fixed32,2,opt,name=contribution,proto3" json:"contribution,omitempty"`
}

func (x *UpdateDiscountRequest) Reset() {
	*x = UpdateDiscountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDiscountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDiscountRequest) ProtoMessage() {}

func (x *UpdateDiscountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDiscountRequest.ProtoReflect.Descriptor instead.
func (*UpdateDiscountRequest) Descriptor() ([]byte, []int) {
	return file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateDiscountRequest) GetUserUid() *commonProto.UUID {
	if x != nil {
		return x.UserUid
	}
	return nil
}

func (x *UpdateDiscountRequest) GetContribution() float32 {
	if x != nil {
		return x.Contribution
	}
	return 0
}

type Loyalty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid           *commonProto.UUID `protobuf:"bytes,1,opt,name=UserUuid,proto3" json:"UserUuid,omitempty"`
	Status             string            `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	Discount           int64             `protobuf:"varint,3,opt,name=Discount,proto3" json:"Discount,omitempty"`
	ContributionAmount float32           `protobuf:"fixed32,4,opt,name=ContributionAmount,proto3" json:"ContributionAmount,omitempty"`
}

func (x *Loyalty) Reset() {
	*x = Loyalty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Loyalty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Loyalty) ProtoMessage() {}

func (x *Loyalty) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Loyalty.ProtoReflect.Descriptor instead.
func (*Loyalty) Descriptor() ([]byte, []int) {
	return file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_rawDescGZIP(), []int{1}
}

func (x *Loyalty) GetUserUuid() *commonProto.UUID {
	if x != nil {
		return x.UserUuid
	}
	return nil
}

func (x *Loyalty) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Loyalty) GetDiscount() int64 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *Loyalty) GetContributionAmount() float32 {
	if x != nil {
		return x.ContributionAmount
	}
	return 0
}

var File_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto protoreflect.FileDescriptor

var file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6c, 0x6f, 0x79,
	0x61, 0x6c, 0x74, 0x79, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x15, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x55, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55,
	0x55, 0x49, 0x44, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x97, 0x01, 0x0a, 0x07, 0x4c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x12, 0x28, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x12, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xdc, 0x01, 0x0a, 0x0e, 0x4c,
	0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x1a, 0x0d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x12,
	0x2d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0c,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x1a, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x22, 0x00, 0x12, 0x28,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x1a, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x47, 0x5a, 0x45, 0x68, 0x6f, 0x74,
	0x65, 0x6c, 0x2d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6c, 0x6f, 0x79,
	0x61, 0x6c, 0x74, 0x79, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_rawDescOnce sync.Once
	file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_rawDescData = file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_rawDesc
)

func file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_rawDescGZIP() []byte {
	file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_rawDescOnce.Do(func() {
		file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_rawDescData)
	})
	return file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_rawDescData
}

var file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_goTypes = []interface{}{
	(*UpdateDiscountRequest)(nil),   // 0: proto.UpdateDiscountRequest
	(*Loyalty)(nil),                 // 1: proto.Loyalty
	(*commonProto.UUID)(nil),        // 2: common.UUID
	(*commonProto.Credentials)(nil), // 3: common.Credentials
	(*commonProto.Token)(nil),       // 4: common.Token
	(*commonProto.Empty)(nil),       // 5: common.Empty
}
var file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_depIdxs = []int32{
	2, // 0: proto.UpdateDiscountRequest.userUid:type_name -> common.UUID
	2, // 1: proto.Loyalty.UserUuid:type_name -> common.UUID
	3, // 2: proto.LoyaltyService.GetToken:input_type -> common.Credentials
	2, // 3: proto.LoyaltyService.GetDiscount:input_type -> common.UUID
	2, // 4: proto.LoyaltyService.AddUser:input_type -> common.UUID
	0, // 5: proto.LoyaltyService.UpdateDiscount:input_type -> proto.UpdateDiscountRequest
	4, // 6: proto.LoyaltyService.GetToken:output_type -> common.Token
	1, // 7: proto.LoyaltyService.GetDiscount:output_type -> proto.Loyalty
	5, // 8: proto.LoyaltyService.AddUser:output_type -> common.Empty
	5, // 9: proto.LoyaltyService.UpdateDiscount:output_type -> common.Empty
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_init() }
func file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_init() {
	if File_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDiscountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Loyalty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_goTypes,
		DependencyIndexes: file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_depIdxs,
		MessageInfos:      file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_msgTypes,
	}.Build()
	File_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto = out.File
	file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_rawDesc = nil
	file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_goTypes = nil
	file_internal_pkg_delivery_grpc_loyalty_service_proto_scheme_proto_depIdxs = nil
}
