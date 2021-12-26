// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: src/proto/render_server.proto

package proto

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

type RenderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Objects []*Object `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *RenderRequest) Reset() {
	*x = RenderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_render_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenderRequest) ProtoMessage() {}

func (x *RenderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_render_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenderRequest.ProtoReflect.Descriptor instead.
func (*RenderRequest) Descriptor() ([]byte, []int) {
	return file_src_proto_render_server_proto_rawDescGZIP(), []int{0}
}

func (x *RenderRequest) GetObjects() []*Object {
	if x != nil {
		return x.Objects
	}
	return nil
}

type RenderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageBytes []byte `protobuf:"bytes,1,opt,name=imageBytes,proto3" json:"imageBytes,omitempty"`
}

func (x *RenderResponse) Reset() {
	*x = RenderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_render_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenderResponse) ProtoMessage() {}

func (x *RenderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_render_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenderResponse.ProtoReflect.Descriptor instead.
func (*RenderResponse) Descriptor() ([]byte, []int) {
	return file_src_proto_render_server_proto_rawDescGZIP(), []int{1}
}

func (x *RenderResponse) GetImageBytes() []byte {
	if x != nil {
		return x.ImageBytes
	}
	return nil
}

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shape    string  `protobuf:"bytes,1,opt,name=shape,proto3" json:"shape,omitempty"`
	Position *Vector `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	Scale    *Vector `protobuf:"bytes,3,opt,name=scale,proto3" json:"scale,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_render_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_render_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_src_proto_render_server_proto_rawDescGZIP(), []int{2}
}

func (x *Object) GetShape() string {
	if x != nil {
		return x.Shape
	}
	return ""
}

func (x *Object) GetPosition() *Vector {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *Object) GetScale() *Vector {
	if x != nil {
		return x.Scale
	}
	return nil
}

type Vector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float64 `protobuf:"fixed64,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *Vector) Reset() {
	*x = Vector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_render_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vector) ProtoMessage() {}

func (x *Vector) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_render_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vector.ProtoReflect.Descriptor instead.
func (*Vector) Descriptor() ([]byte, []int) {
	return file_src_proto_render_server_proto_rawDescGZIP(), []int{3}
}

func (x *Vector) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vector) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Vector) GetZ() float64 {
	if x != nil {
		return x.Z
	}
	return 0
}

var File_src_proto_render_server_proto protoreflect.FileDescriptor

var file_src_proto_render_server_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x40,
	0x0a, 0x0d, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2f, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x22, 0x30, 0x0a, 0x0e, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x22, 0x7e, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x68, 0x61, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x61,
	0x70, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x22, 0x32, 0x0a, 0x06, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x0c, 0x0a, 0x01,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x01, 0x7a, 0x32, 0x5a, 0x0a, 0x08, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x65, 0x72, 0x12, 0x4e, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_proto_render_server_proto_rawDescOnce sync.Once
	file_src_proto_render_server_proto_rawDescData = file_src_proto_render_server_proto_rawDesc
)

func file_src_proto_render_server_proto_rawDescGZIP() []byte {
	file_src_proto_render_server_proto_rawDescOnce.Do(func() {
		file_src_proto_render_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_proto_render_server_proto_rawDescData)
	})
	return file_src_proto_render_server_proto_rawDescData
}

var file_src_proto_render_server_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_src_proto_render_server_proto_goTypes = []interface{}{
	(*RenderRequest)(nil),  // 0: render_server.RenderRequest
	(*RenderResponse)(nil), // 1: render_server.RenderResponse
	(*Object)(nil),         // 2: render_server.Object
	(*Vector)(nil),         // 3: render_server.Vector
}
var file_src_proto_render_server_proto_depIdxs = []int32{
	2, // 0: render_server.RenderRequest.objects:type_name -> render_server.Object
	3, // 1: render_server.Object.position:type_name -> render_server.Vector
	3, // 2: render_server.Object.scale:type_name -> render_server.Vector
	0, // 3: render_server.Renderer.RequestRender:input_type -> render_server.RenderRequest
	1, // 4: render_server.Renderer.RequestRender:output_type -> render_server.RenderResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_src_proto_render_server_proto_init() }
func file_src_proto_render_server_proto_init() {
	if File_src_proto_render_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_proto_render_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenderRequest); i {
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
		file_src_proto_render_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenderResponse); i {
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
		file_src_proto_render_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
		file_src_proto_render_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vector); i {
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
			RawDescriptor: file_src_proto_render_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_proto_render_server_proto_goTypes,
		DependencyIndexes: file_src_proto_render_server_proto_depIdxs,
		MessageInfos:      file_src_proto_render_server_proto_msgTypes,
	}.Build()
	File_src_proto_render_server_proto = out.File
	file_src_proto_render_server_proto_rawDesc = nil
	file_src_proto_render_server_proto_goTypes = nil
	file_src_proto_render_server_proto_depIdxs = nil
}
