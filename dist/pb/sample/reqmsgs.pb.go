// Code generated by furo-proto-gen. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: sample/reqmsgs.proto

package samplepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// request message for CreateSample
type CreateSampleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body with sample.Sample
	Body *Sample `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *CreateSampleRequest) Reset() {
	*x = CreateSampleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_reqmsgs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSampleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSampleRequest) ProtoMessage() {}

func (x *CreateSampleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sample_reqmsgs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSampleRequest.ProtoReflect.Descriptor instead.
func (*CreateSampleRequest) Descriptor() ([]byte, []int) {
	return file_sample_reqmsgs_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSampleRequest) GetBody() *Sample {
	if x != nil {
		return x.Body
	}
	return nil
}

// request message for DeleteSample
type DeleteSampleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body with google.protobuf.Empty
	Body *emptypb.Empty `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// The query param smp stands for the id of a  Sample.
	Smp string `protobuf:"bytes,2,opt,name=smp,proto3" json:"smp,omitempty"`
}

func (x *DeleteSampleRequest) Reset() {
	*x = DeleteSampleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_reqmsgs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSampleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSampleRequest) ProtoMessage() {}

func (x *DeleteSampleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sample_reqmsgs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSampleRequest.ProtoReflect.Descriptor instead.
func (*DeleteSampleRequest) Descriptor() ([]byte, []int) {
	return file_sample_reqmsgs_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteSampleRequest) GetBody() *emptypb.Empty {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *DeleteSampleRequest) GetSmp() string {
	if x != nil {
		return x.Smp
	}
	return ""
}

// request message for GetSample
type GetSampleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body with google.protobuf.Empty
	Body *emptypb.Empty `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// The query param smp stands for the id of a Sample.
	Smp string `protobuf:"bytes,2,opt,name=smp,proto3" json:"smp,omitempty"`
}

func (x *GetSampleRequest) Reset() {
	*x = GetSampleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_reqmsgs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSampleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSampleRequest) ProtoMessage() {}

func (x *GetSampleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sample_reqmsgs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSampleRequest.ProtoReflect.Descriptor instead.
func (*GetSampleRequest) Descriptor() ([]byte, []int) {
	return file_sample_reqmsgs_proto_rawDescGZIP(), []int{2}
}

func (x *GetSampleRequest) GetBody() *emptypb.Empty {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *GetSampleRequest) GetSmp() string {
	if x != nil {
		return x.Smp
	}
	return ""
}

// request message for ListSamples
type ListSamplesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body with google.protobuf.Empty
	Body *emptypb.Empty `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Search query param.
	Q string `protobuf:"bytes,2,opt,name=q,proto3" json:"q,omitempty"`
	// Filter query param
	Filter string `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	// Pagination query param.
	Page string `protobuf:"bytes,4,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListSamplesRequest) Reset() {
	*x = ListSamplesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_reqmsgs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSamplesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSamplesRequest) ProtoMessage() {}

func (x *ListSamplesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sample_reqmsgs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSamplesRequest.ProtoReflect.Descriptor instead.
func (*ListSamplesRequest) Descriptor() ([]byte, []int) {
	return file_sample_reqmsgs_proto_rawDescGZIP(), []int{3}
}

func (x *ListSamplesRequest) GetBody() *emptypb.Empty {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *ListSamplesRequest) GetQ() string {
	if x != nil {
		return x.Q
	}
	return ""
}

func (x *ListSamplesRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListSamplesRequest) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

// request message for UpdateSample
type UpdateSampleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body with sample.Sample
	Body *Sample `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// The query param smp stands for the id of a  Sample.
	Smp string `protobuf:"bytes,2,opt,name=smp,proto3" json:"smp,omitempty"`
	// Needed to patch a record
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateSampleRequest) Reset() {
	*x = UpdateSampleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_reqmsgs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSampleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSampleRequest) ProtoMessage() {}

func (x *UpdateSampleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sample_reqmsgs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSampleRequest.ProtoReflect.Descriptor instead.
func (*UpdateSampleRequest) Descriptor() ([]byte, []int) {
	return file_sample_reqmsgs_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateSampleRequest) GetBody() *Sample {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *UpdateSampleRequest) GetSmp() string {
	if x != nil {
		return x.Smp
	}
	return ""
}

func (x *UpdateSampleRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_sample_reqmsgs_proto protoreflect.FileDescriptor

var file_sample_reqmsgs_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x72, 0x65, 0x71, 0x6d, 0x73, 0x67, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x39, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x53, 0x0a,
	0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x6d, 0x70, 0x22, 0x50, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x73, 0x6d, 0x70, 0x22, 0x7a, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x01, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x22, 0x88, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x6d, 0x70, 0x12, 0x3b,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x66, 0x0a, 0x1b, 0x63,
	0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72,
	0x69, 0x61, 0x6c, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x0c, 0x52, 0x65, 0x71, 0x6d,
	0x73, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x73, 0x73, 0x65, 0x31, 0x2f, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x74,
	0x2f, 0x70, 0x62, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3b, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sample_reqmsgs_proto_rawDescOnce sync.Once
	file_sample_reqmsgs_proto_rawDescData = file_sample_reqmsgs_proto_rawDesc
)

func file_sample_reqmsgs_proto_rawDescGZIP() []byte {
	file_sample_reqmsgs_proto_rawDescOnce.Do(func() {
		file_sample_reqmsgs_proto_rawDescData = protoimpl.X.CompressGZIP(file_sample_reqmsgs_proto_rawDescData)
	})
	return file_sample_reqmsgs_proto_rawDescData
}

var file_sample_reqmsgs_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sample_reqmsgs_proto_goTypes = []interface{}{
	(*CreateSampleRequest)(nil),   // 0: sample.CreateSampleRequest
	(*DeleteSampleRequest)(nil),   // 1: sample.DeleteSampleRequest
	(*GetSampleRequest)(nil),      // 2: sample.GetSampleRequest
	(*ListSamplesRequest)(nil),    // 3: sample.ListSamplesRequest
	(*UpdateSampleRequest)(nil),   // 4: sample.UpdateSampleRequest
	(*Sample)(nil),                // 5: sample.Sample
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
	(*fieldmaskpb.FieldMask)(nil), // 7: google.protobuf.FieldMask
}
var file_sample_reqmsgs_proto_depIdxs = []int32{
	5, // 0: sample.CreateSampleRequest.body:type_name -> sample.Sample
	6, // 1: sample.DeleteSampleRequest.body:type_name -> google.protobuf.Empty
	6, // 2: sample.GetSampleRequest.body:type_name -> google.protobuf.Empty
	6, // 3: sample.ListSamplesRequest.body:type_name -> google.protobuf.Empty
	5, // 4: sample.UpdateSampleRequest.body:type_name -> sample.Sample
	7, // 5: sample.UpdateSampleRequest.update_mask:type_name -> google.protobuf.FieldMask
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_sample_reqmsgs_proto_init() }
func file_sample_reqmsgs_proto_init() {
	if File_sample_reqmsgs_proto != nil {
		return
	}
	file_sample_sample_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sample_reqmsgs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSampleRequest); i {
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
		file_sample_reqmsgs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSampleRequest); i {
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
		file_sample_reqmsgs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSampleRequest); i {
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
		file_sample_reqmsgs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSamplesRequest); i {
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
		file_sample_reqmsgs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSampleRequest); i {
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
			RawDescriptor: file_sample_reqmsgs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sample_reqmsgs_proto_goTypes,
		DependencyIndexes: file_sample_reqmsgs_proto_depIdxs,
		MessageInfos:      file_sample_reqmsgs_proto_msgTypes,
	}.Build()
	File_sample_reqmsgs_proto = out.File
	file_sample_reqmsgs_proto_rawDesc = nil
	file_sample_reqmsgs_proto_goTypes = nil
	file_sample_reqmsgs_proto_depIdxs = nil
}
