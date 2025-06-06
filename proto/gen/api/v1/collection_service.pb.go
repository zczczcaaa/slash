// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/v1/collection_service.proto

package apiv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Collection struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatorId     int32                  `protobuf:"varint,2,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	CreatedTime   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	UpdatedTime   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	Name          string                 `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Title         string                 `protobuf:"bytes,7,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	ShortcutIds   []int32                `protobuf:"varint,9,rep,packed,name=shortcut_ids,json=shortcutIds,proto3" json:"shortcut_ids,omitempty"`
	Visibility    Visibility             `protobuf:"varint,10,opt,name=visibility,proto3,enum=slash.api.v1.Visibility" json:"visibility,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Collection) Reset() {
	*x = Collection{}
	mi := &file_api_v1_collection_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Collection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Collection) ProtoMessage() {}

func (x *Collection) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_collection_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Collection.ProtoReflect.Descriptor instead.
func (*Collection) Descriptor() ([]byte, []int) {
	return file_api_v1_collection_service_proto_rawDescGZIP(), []int{0}
}

func (x *Collection) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Collection) GetCreatorId() int32 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *Collection) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *Collection) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *Collection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Collection) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Collection) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Collection) GetShortcutIds() []int32 {
	if x != nil {
		return x.ShortcutIds
	}
	return nil
}

func (x *Collection) GetVisibility() Visibility {
	if x != nil {
		return x.Visibility
	}
	return Visibility_VISIBILITY_UNSPECIFIED
}

type ListCollectionsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCollectionsRequest) Reset() {
	*x = ListCollectionsRequest{}
	mi := &file_api_v1_collection_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCollectionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCollectionsRequest) ProtoMessage() {}

func (x *ListCollectionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_collection_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCollectionsRequest.ProtoReflect.Descriptor instead.
func (*ListCollectionsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_collection_service_proto_rawDescGZIP(), []int{1}
}

type ListCollectionsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Collections   []*Collection          `protobuf:"bytes,1,rep,name=collections,proto3" json:"collections,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCollectionsResponse) Reset() {
	*x = ListCollectionsResponse{}
	mi := &file_api_v1_collection_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCollectionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCollectionsResponse) ProtoMessage() {}

func (x *ListCollectionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_collection_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCollectionsResponse.ProtoReflect.Descriptor instead.
func (*ListCollectionsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_collection_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListCollectionsResponse) GetCollections() []*Collection {
	if x != nil {
		return x.Collections
	}
	return nil
}

type GetCollectionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCollectionRequest) Reset() {
	*x = GetCollectionRequest{}
	mi := &file_api_v1_collection_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCollectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionRequest) ProtoMessage() {}

func (x *GetCollectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_collection_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionRequest.ProtoReflect.Descriptor instead.
func (*GetCollectionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_collection_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetCollectionRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCollectionByNameRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCollectionByNameRequest) Reset() {
	*x = GetCollectionByNameRequest{}
	mi := &file_api_v1_collection_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCollectionByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionByNameRequest) ProtoMessage() {}

func (x *GetCollectionByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_collection_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionByNameRequest.ProtoReflect.Descriptor instead.
func (*GetCollectionByNameRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_collection_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetCollectionByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateCollectionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Collection    *Collection            `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCollectionRequest) Reset() {
	*x = CreateCollectionRequest{}
	mi := &file_api_v1_collection_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCollectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCollectionRequest) ProtoMessage() {}

func (x *CreateCollectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_collection_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCollectionRequest.ProtoReflect.Descriptor instead.
func (*CreateCollectionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_collection_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateCollectionRequest) GetCollection() *Collection {
	if x != nil {
		return x.Collection
	}
	return nil
}

type UpdateCollectionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Collection    *Collection            `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	UpdateMask    *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCollectionRequest) Reset() {
	*x = UpdateCollectionRequest{}
	mi := &file_api_v1_collection_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCollectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCollectionRequest) ProtoMessage() {}

func (x *UpdateCollectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_collection_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCollectionRequest.ProtoReflect.Descriptor instead.
func (*UpdateCollectionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_collection_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCollectionRequest) GetCollection() *Collection {
	if x != nil {
		return x.Collection
	}
	return nil
}

func (x *UpdateCollectionRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type DeleteCollectionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCollectionRequest) Reset() {
	*x = DeleteCollectionRequest{}
	mi := &file_api_v1_collection_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCollectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCollectionRequest) ProtoMessage() {}

func (x *DeleteCollectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_collection_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCollectionRequest.ProtoReflect.Descriptor instead.
func (*DeleteCollectionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_collection_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteCollectionRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_api_v1_collection_service_proto protoreflect.FileDescriptor

const file_api_v1_collection_service_proto_rawDesc = "" +
	"\n" +
	"\x1fapi/v1/collection_service.proto\x12\fslash.api.v1\x1a\x13api/v1/common.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x17google/api/client.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a google/protobuf/field_mask.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xe2\x02\n" +
	"\n" +
	"Collection\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x1d\n" +
	"\n" +
	"creator_id\x18\x02 \x01(\x05R\tcreatorId\x12=\n" +
	"\fcreated_time\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\vcreatedTime\x12=\n" +
	"\fupdated_time\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\vupdatedTime\x12\x12\n" +
	"\x04name\x18\x06 \x01(\tR\x04name\x12\x14\n" +
	"\x05title\x18\a \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\b \x01(\tR\vdescription\x12!\n" +
	"\fshortcut_ids\x18\t \x03(\x05R\vshortcutIds\x128\n" +
	"\n" +
	"visibility\x18\n" +
	" \x01(\x0e2\x18.slash.api.v1.VisibilityR\n" +
	"visibility\"\x18\n" +
	"\x16ListCollectionsRequest\"U\n" +
	"\x17ListCollectionsResponse\x12:\n" +
	"\vcollections\x18\x01 \x03(\v2\x18.slash.api.v1.CollectionR\vcollections\"&\n" +
	"\x14GetCollectionRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"0\n" +
	"\x1aGetCollectionByNameRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"S\n" +
	"\x17CreateCollectionRequest\x128\n" +
	"\n" +
	"collection\x18\x01 \x01(\v2\x18.slash.api.v1.CollectionR\n" +
	"collection\"\x90\x01\n" +
	"\x17UpdateCollectionRequest\x128\n" +
	"\n" +
	"collection\x18\x01 \x01(\v2\x18.slash.api.v1.CollectionR\n" +
	"collection\x12;\n" +
	"\vupdate_mask\x18\x02 \x01(\v2\x1a.google.protobuf.FieldMaskR\n" +
	"updateMask\")\n" +
	"\x17DeleteCollectionRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id2\x83\x06\n" +
	"\x11CollectionService\x12{\n" +
	"\x0fListCollections\x12$.slash.api.v1.ListCollectionsRequest\x1a%.slash.api.v1.ListCollectionsResponse\"\x1b\x82\xd3\xe4\x93\x02\x15\x12\x13/api/v1/collections\x12t\n" +
	"\rGetCollection\x12\".slash.api.v1.GetCollectionRequest\x1a\x18.slash.api.v1.Collection\"%\xdaA\x02id\x82\xd3\xe4\x93\x02\x1a\x12\x18/api/v1/collections/{id}\x12[\n" +
	"\x13GetCollectionByName\x12(.slash.api.v1.GetCollectionByNameRequest\x1a\x18.slash.api.v1.Collection\"\x00\x12|\n" +
	"\x10CreateCollection\x12%.slash.api.v1.CreateCollectionRequest\x1a\x18.slash.api.v1.Collection\"'\x82\xd3\xe4\x93\x02!:\n" +
	"collection\"\x13/api/v1/collections\x12\xa5\x01\n" +
	"\x10UpdateCollection\x12%.slash.api.v1.UpdateCollectionRequest\x1a\x18.slash.api.v1.Collection\"P\xdaA\x16collection,update_mask\x82\xd3\xe4\x93\x021:\n" +
	"collection\x1a#/api/v1/collections/{collection.id}\x12x\n" +
	"\x10DeleteCollection\x12%.slash.api.v1.DeleteCollectionRequest\x1a\x16.google.protobuf.Empty\"%\xdaA\x02id\x82\xd3\xe4\x93\x02\x1a*\x18/api/v1/collections/{id}B\xb4\x01\n" +
	"\x10com.slash.api.v1B\x16CollectionServiceProtoP\x01Z6github.com/yourselfhosted/slash/proto/gen/api/v1;apiv1\xa2\x02\x03SAX\xaa\x02\fSlash.Api.V1\xca\x02\fSlash\\Api\\V1\xe2\x02\x18Slash\\Api\\V1\\GPBMetadata\xea\x02\x0eSlash::Api::V1b\x06proto3"

var (
	file_api_v1_collection_service_proto_rawDescOnce sync.Once
	file_api_v1_collection_service_proto_rawDescData []byte
)

func file_api_v1_collection_service_proto_rawDescGZIP() []byte {
	file_api_v1_collection_service_proto_rawDescOnce.Do(func() {
		file_api_v1_collection_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1_collection_service_proto_rawDesc), len(file_api_v1_collection_service_proto_rawDesc)))
	})
	return file_api_v1_collection_service_proto_rawDescData
}

var file_api_v1_collection_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_v1_collection_service_proto_goTypes = []any{
	(*Collection)(nil),                 // 0: slash.api.v1.Collection
	(*ListCollectionsRequest)(nil),     // 1: slash.api.v1.ListCollectionsRequest
	(*ListCollectionsResponse)(nil),    // 2: slash.api.v1.ListCollectionsResponse
	(*GetCollectionRequest)(nil),       // 3: slash.api.v1.GetCollectionRequest
	(*GetCollectionByNameRequest)(nil), // 4: slash.api.v1.GetCollectionByNameRequest
	(*CreateCollectionRequest)(nil),    // 5: slash.api.v1.CreateCollectionRequest
	(*UpdateCollectionRequest)(nil),    // 6: slash.api.v1.UpdateCollectionRequest
	(*DeleteCollectionRequest)(nil),    // 7: slash.api.v1.DeleteCollectionRequest
	(*timestamppb.Timestamp)(nil),      // 8: google.protobuf.Timestamp
	(Visibility)(0),                    // 9: slash.api.v1.Visibility
	(*fieldmaskpb.FieldMask)(nil),      // 10: google.protobuf.FieldMask
	(*emptypb.Empty)(nil),              // 11: google.protobuf.Empty
}
var file_api_v1_collection_service_proto_depIdxs = []int32{
	8,  // 0: slash.api.v1.Collection.created_time:type_name -> google.protobuf.Timestamp
	8,  // 1: slash.api.v1.Collection.updated_time:type_name -> google.protobuf.Timestamp
	9,  // 2: slash.api.v1.Collection.visibility:type_name -> slash.api.v1.Visibility
	0,  // 3: slash.api.v1.ListCollectionsResponse.collections:type_name -> slash.api.v1.Collection
	0,  // 4: slash.api.v1.CreateCollectionRequest.collection:type_name -> slash.api.v1.Collection
	0,  // 5: slash.api.v1.UpdateCollectionRequest.collection:type_name -> slash.api.v1.Collection
	10, // 6: slash.api.v1.UpdateCollectionRequest.update_mask:type_name -> google.protobuf.FieldMask
	1,  // 7: slash.api.v1.CollectionService.ListCollections:input_type -> slash.api.v1.ListCollectionsRequest
	3,  // 8: slash.api.v1.CollectionService.GetCollection:input_type -> slash.api.v1.GetCollectionRequest
	4,  // 9: slash.api.v1.CollectionService.GetCollectionByName:input_type -> slash.api.v1.GetCollectionByNameRequest
	5,  // 10: slash.api.v1.CollectionService.CreateCollection:input_type -> slash.api.v1.CreateCollectionRequest
	6,  // 11: slash.api.v1.CollectionService.UpdateCollection:input_type -> slash.api.v1.UpdateCollectionRequest
	7,  // 12: slash.api.v1.CollectionService.DeleteCollection:input_type -> slash.api.v1.DeleteCollectionRequest
	2,  // 13: slash.api.v1.CollectionService.ListCollections:output_type -> slash.api.v1.ListCollectionsResponse
	0,  // 14: slash.api.v1.CollectionService.GetCollection:output_type -> slash.api.v1.Collection
	0,  // 15: slash.api.v1.CollectionService.GetCollectionByName:output_type -> slash.api.v1.Collection
	0,  // 16: slash.api.v1.CollectionService.CreateCollection:output_type -> slash.api.v1.Collection
	0,  // 17: slash.api.v1.CollectionService.UpdateCollection:output_type -> slash.api.v1.Collection
	11, // 18: slash.api.v1.CollectionService.DeleteCollection:output_type -> google.protobuf.Empty
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_v1_collection_service_proto_init() }
func file_api_v1_collection_service_proto_init() {
	if File_api_v1_collection_service_proto != nil {
		return
	}
	file_api_v1_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1_collection_service_proto_rawDesc), len(file_api_v1_collection_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_collection_service_proto_goTypes,
		DependencyIndexes: file_api_v1_collection_service_proto_depIdxs,
		MessageInfos:      file_api_v1_collection_service_proto_msgTypes,
	}.Build()
	File_api_v1_collection_service_proto = out.File
	file_api_v1_collection_service_proto_goTypes = nil
	file_api_v1_collection_service_proto_depIdxs = nil
}
