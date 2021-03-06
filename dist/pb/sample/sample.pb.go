// Code generated by furo-proto-gen. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: sample/sample.proto

package samplepb

import (
	furo "github.com/theNorstroem/FuroBaseSpecs/dist/pb/furo"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A sample type
type Sample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ULID as required string.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Readonly display name
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// True when a draft exists
	InEdit bool `protobuf:"varint,3,opt,name=in_edit,json=inEdit,proto3" json:"in_edit,omitempty"`
	// Repeated string
	Tags []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	// Oneof example
	Is string `protobuf:"bytes,6,opt,name=is,proto3" json:"is,omitempty"`
	// Oneof example
	Val int32 `protobuf:"varint,7,opt,name=val,proto3" json:"val,omitempty"`
	// Just another demo field
	Newfield string `protobuf:"bytes,8,opt,name=newfield,proto3" json:"newfield,omitempty"`
	// Types that are assignable to Oneofname:
	//	*Sample_This
	Oneofname isSample_Oneofname `protobuf_oneof:"oneofname"`
}

func (x *Sample) Reset() {
	*x = Sample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_sample_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sample) ProtoMessage() {}

func (x *Sample) ProtoReflect() protoreflect.Message {
	mi := &file_sample_sample_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sample.ProtoReflect.Descriptor instead.
func (*Sample) Descriptor() ([]byte, []int) {
	return file_sample_sample_proto_rawDescGZIP(), []int{0}
}

func (x *Sample) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Sample) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Sample) GetInEdit() bool {
	if x != nil {
		return x.InEdit
	}
	return false
}

func (x *Sample) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Sample) GetIs() string {
	if x != nil {
		return x.Is
	}
	return ""
}

func (x *Sample) GetVal() int32 {
	if x != nil {
		return x.Val
	}
	return 0
}

func (x *Sample) GetNewfield() string {
	if x != nil {
		return x.Newfield
	}
	return ""
}

func (m *Sample) GetOneofname() isSample_Oneofname {
	if m != nil {
		return m.Oneofname
	}
	return nil
}

func (x *Sample) GetThis() string {
	if x, ok := x.GetOneofname().(*Sample_This); ok {
		return x.This
	}
	return ""
}

type isSample_Oneofname interface {
	isSample_Oneofname()
}

type Sample_This struct {
	// Oneof example
	This string `protobuf:"bytes,5,opt,name=this,proto3,oneof"`
}

func (*Sample_This) isSample_Oneofname() {}

// Collectioncontainer which holds a sample.Sample
type SampleCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the data contains a sample.Sample
	Entities []*SampleEntity `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"`
	// the Hateoas links
	Links []*furo.Link `protobuf:"bytes,2,rep,name=links,proto3" json:"links,omitempty"`
}

func (x *SampleCollection) Reset() {
	*x = SampleCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_sample_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleCollection) ProtoMessage() {}

func (x *SampleCollection) ProtoReflect() protoreflect.Message {
	mi := &file_sample_sample_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleCollection.ProtoReflect.Descriptor instead.
func (*SampleCollection) Descriptor() ([]byte, []int) {
	return file_sample_sample_proto_rawDescGZIP(), []int{1}
}

func (x *SampleCollection) GetEntities() []*SampleEntity {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *SampleCollection) GetLinks() []*furo.Link {
	if x != nil {
		return x.Links
	}
	return nil
}

// SimpleEntity
type SampleEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the data contains a sample.Sample
	Data *Sample `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// the Hateoas links
	Links []*furo.Link `protobuf:"bytes,2,rep,name=links,proto3" json:"links,omitempty"`
	// Meta for the response
	Meta *furo.Meta `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *SampleEntity) Reset() {
	*x = SampleEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_sample_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleEntity) ProtoMessage() {}

func (x *SampleEntity) ProtoReflect() protoreflect.Message {
	mi := &file_sample_sample_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleEntity.ProtoReflect.Descriptor instead.
func (*SampleEntity) Descriptor() ([]byte, []int) {
	return file_sample_sample_proto_rawDescGZIP(), []int{2}
}

func (x *SampleEntity) GetData() *Sample {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SampleEntity) GetLinks() []*furo.Link {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *SampleEntity) GetMeta() *furo.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

var File_sample_sample_proto protoreflect.FileDescriptor

var file_sample_sample_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x0f, 0x66,
	0x75, 0x72, 0x6f, 0x2f, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9,
	0x01, 0x0a, 0x06, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x69, 0x6e, 0x5f, 0x65, 0x64, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69,
	0x6e, 0x45, 0x64, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6e,
	0x65, 0x77, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x65, 0x77, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x04, 0x74, 0x68, 0x69, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x74, 0x68, 0x69, 0x73, 0x42, 0x0b, 0x0a,
	0x09, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x66, 0x0a, 0x10, 0x53, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30,
	0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x20, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e,
	0x6b, 0x73, 0x22, 0x74, 0x0a, 0x0c, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x4c, 0x69, 0x6e,
	0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x42, 0x65, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c,
	0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x0b, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x73, 0x73, 0x65, 0x31, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x70, 0x62, 0x2f,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3b, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sample_sample_proto_rawDescOnce sync.Once
	file_sample_sample_proto_rawDescData = file_sample_sample_proto_rawDesc
)

func file_sample_sample_proto_rawDescGZIP() []byte {
	file_sample_sample_proto_rawDescOnce.Do(func() {
		file_sample_sample_proto_rawDescData = protoimpl.X.CompressGZIP(file_sample_sample_proto_rawDescData)
	})
	return file_sample_sample_proto_rawDescData
}

var file_sample_sample_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sample_sample_proto_goTypes = []interface{}{
	(*Sample)(nil),           // 0: sample.Sample
	(*SampleCollection)(nil), // 1: sample.SampleCollection
	(*SampleEntity)(nil),     // 2: sample.SampleEntity
	(*furo.Link)(nil),        // 3: furo.Link
	(*furo.Meta)(nil),        // 4: furo.Meta
}
var file_sample_sample_proto_depIdxs = []int32{
	2, // 0: sample.SampleCollection.entities:type_name -> sample.SampleEntity
	3, // 1: sample.SampleCollection.links:type_name -> furo.Link
	0, // 2: sample.SampleEntity.data:type_name -> sample.Sample
	3, // 3: sample.SampleEntity.links:type_name -> furo.Link
	4, // 4: sample.SampleEntity.meta:type_name -> furo.Meta
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_sample_sample_proto_init() }
func file_sample_sample_proto_init() {
	if File_sample_sample_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sample_sample_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sample); i {
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
		file_sample_sample_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleCollection); i {
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
		file_sample_sample_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleEntity); i {
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
	file_sample_sample_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Sample_This)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sample_sample_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sample_sample_proto_goTypes,
		DependencyIndexes: file_sample_sample_proto_depIdxs,
		MessageInfos:      file_sample_sample_proto_msgTypes,
	}.Build()
	File_sample_sample_proto = out.File
	file_sample_sample_proto_rawDesc = nil
	file_sample_sample_proto_goTypes = nil
	file_sample_sample_proto_depIdxs = nil
}
