// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: directory/model.proto

package directory

import (
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

type DirectoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *DirectoryRequest) Reset() {
	*x = DirectoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directory_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectoryRequest) ProtoMessage() {}

func (x *DirectoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_directory_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectoryRequest.ProtoReflect.Descriptor instead.
func (*DirectoryRequest) Descriptor() ([]byte, []int) {
	return file_directory_model_proto_rawDescGZIP(), []int{0}
}

func (x *DirectoryRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type FileCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size string `protobuf:"bytes,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *FileCollection) Reset() {
	*x = FileCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directory_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileCollection) ProtoMessage() {}

func (x *FileCollection) ProtoReflect() protoreflect.Message {
	mi := &file_directory_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileCollection.ProtoReflect.Descriptor instead.
func (*FileCollection) Descriptor() ([]byte, []int) {
	return file_directory_model_proto_rawDescGZIP(), []int{1}
}

func (x *FileCollection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileCollection) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

type DirectoryCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size string `protobuf:"bytes,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *DirectoryCollection) Reset() {
	*x = DirectoryCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directory_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectoryCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectoryCollection) ProtoMessage() {}

func (x *DirectoryCollection) ProtoReflect() protoreflect.Message {
	mi := &file_directory_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectoryCollection.ProtoReflect.Descriptor instead.
func (*DirectoryCollection) Descriptor() ([]byte, []int) {
	return file_directory_model_proto_rawDescGZIP(), []int{2}
}

func (x *DirectoryCollection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DirectoryCollection) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

type DirectoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files       []*FileCollection      `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	Directories []*DirectoryCollection `protobuf:"bytes,2,rep,name=directories,proto3" json:"directories,omitempty"`
}

func (x *DirectoryResponse) Reset() {
	*x = DirectoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directory_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectoryResponse) ProtoMessage() {}

func (x *DirectoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_directory_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectoryResponse.ProtoReflect.Descriptor instead.
func (*DirectoryResponse) Descriptor() ([]byte, []int) {
	return file_directory_model_proto_rawDescGZIP(), []int{3}
}

func (x *DirectoryResponse) GetFiles() []*FileCollection {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *DirectoryResponse) GetDirectories() []*DirectoryCollection {
	if x != nil {
		return x.Directories
	}
	return nil
}

var File_directory_model_proto protoreflect.FileDescriptor

var file_directory_model_proto_rawDesc = []byte{
	0x0a, 0x15, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x26, 0x0a, 0x10, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x22, 0x38, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x3d, 0x0a, 0x13, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x92, 0x01, 0x0a, 0x11, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x35, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x46, 0x0a, 0x0b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x42, 0x15,
	0x5a, 0x13, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_directory_model_proto_rawDescOnce sync.Once
	file_directory_model_proto_rawDescData = file_directory_model_proto_rawDesc
)

func file_directory_model_proto_rawDescGZIP() []byte {
	file_directory_model_proto_rawDescOnce.Do(func() {
		file_directory_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_directory_model_proto_rawDescData)
	})
	return file_directory_model_proto_rawDescData
}

var file_directory_model_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_directory_model_proto_goTypes = []interface{}{
	(*DirectoryRequest)(nil),    // 0: directory.model.DirectoryRequest
	(*FileCollection)(nil),      // 1: directory.model.FileCollection
	(*DirectoryCollection)(nil), // 2: directory.model.DirectoryCollection
	(*DirectoryResponse)(nil),   // 3: directory.model.DirectoryResponse
}
var file_directory_model_proto_depIdxs = []int32{
	1, // 0: directory.model.DirectoryResponse.files:type_name -> directory.model.FileCollection
	2, // 1: directory.model.DirectoryResponse.directories:type_name -> directory.model.DirectoryCollection
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_directory_model_proto_init() }
func file_directory_model_proto_init() {
	if File_directory_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_directory_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectoryRequest); i {
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
		file_directory_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileCollection); i {
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
		file_directory_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectoryCollection); i {
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
		file_directory_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectoryResponse); i {
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
			RawDescriptor: file_directory_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_directory_model_proto_goTypes,
		DependencyIndexes: file_directory_model_proto_depIdxs,
		MessageInfos:      file_directory_model_proto_msgTypes,
	}.Build()
	File_directory_model_proto = out.File
	file_directory_model_proto_rawDesc = nil
	file_directory_model_proto_goTypes = nil
	file_directory_model_proto_depIdxs = nil
}