// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: proto/te.proto

package pb

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

type Te int32

const (
	Te_Gu    Te = 0
	Te_Choki Te = 1
	Te_Pa    Te = 2
	Te_None  Te = 3
)

// Enum value maps for Te.
var (
	Te_name = map[int32]string{
		0: "Gu",
		1: "Choki",
		2: "Pa",
		3: "None",
	}
	Te_value = map[string]int32{
		"Gu":    0,
		"Choki": 1,
		"Pa":    2,
		"None":  3,
	}
)

func (x Te) Enum() *Te {
	p := new(Te)
	*p = x
	return p
}

func (x Te) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Te) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_te_proto_enumTypes[0].Descriptor()
}

func (Te) Type() protoreflect.EnumType {
	return &file_proto_te_proto_enumTypes[0]
}

func (x Te) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Te.Descriptor instead.
func (Te) EnumDescriptor() ([]byte, []int) {
	return file_proto_te_proto_rawDescGZIP(), []int{0}
}

var File_proto_te_proto protoreflect.FileDescriptor

var file_proto_te_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x2a, 0x29, 0x0a, 0x02, 0x54, 0x65, 0x12, 0x06, 0x0a, 0x02,
	0x47, 0x75, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x68, 0x6f, 0x6b, 0x69, 0x10, 0x01, 0x12,
	0x06, 0x0a, 0x02, 0x50, 0x61, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10,
	0x03, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x78, 0x78, 0x61, 0x72, 0x75, 0x70, 0x61, 0x6b, 0x61, 0x78, 0x78, 0x2f, 0x7a, 0x79, 0x61, 0x6e,
	0x6b, 0x65, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_te_proto_rawDescOnce sync.Once
	file_proto_te_proto_rawDescData = file_proto_te_proto_rawDesc
)

func file_proto_te_proto_rawDescGZIP() []byte {
	file_proto_te_proto_rawDescOnce.Do(func() {
		file_proto_te_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_te_proto_rawDescData)
	})
	return file_proto_te_proto_rawDescData
}

var file_proto_te_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_te_proto_goTypes = []interface{}{
	(Te)(0), // 0: game.Te
}
var file_proto_te_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_te_proto_init() }
func file_proto_te_proto_init() {
	if File_proto_te_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_te_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_te_proto_goTypes,
		DependencyIndexes: file_proto_te_proto_depIdxs,
		EnumInfos:         file_proto_te_proto_enumTypes,
	}.Build()
	File_proto_te_proto = out.File
	file_proto_te_proto_rawDesc = nil
	file_proto_te_proto_goTypes = nil
	file_proto_te_proto_depIdxs = nil
}
