// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.27.3
// source: pkg/proto/configuration/cas/cas.proto

package cas

import (
	eviction "github.com/buildbarn/bb-storage/pkg/proto/configuration/eviction"
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

type CachingDirectoryFetcherConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaximumCount           int64                           `protobuf:"varint,1,opt,name=maximum_count,json=maximumCount,proto3" json:"maximum_count,omitempty"`
	MaximumSizeBytes       int64                           `protobuf:"varint,2,opt,name=maximum_size_bytes,json=maximumSizeBytes,proto3" json:"maximum_size_bytes,omitempty"`
	CacheReplacementPolicy eviction.CacheReplacementPolicy `protobuf:"varint,3,opt,name=cache_replacement_policy,json=cacheReplacementPolicy,proto3,enum=buildbarn.configuration.eviction.CacheReplacementPolicy" json:"cache_replacement_policy,omitempty"`
}

func (x *CachingDirectoryFetcherConfiguration) Reset() {
	*x = CachingDirectoryFetcherConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_cas_cas_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CachingDirectoryFetcherConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CachingDirectoryFetcherConfiguration) ProtoMessage() {}

func (x *CachingDirectoryFetcherConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_cas_cas_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CachingDirectoryFetcherConfiguration.ProtoReflect.Descriptor instead.
func (*CachingDirectoryFetcherConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_cas_cas_proto_rawDescGZIP(), []int{0}
}

func (x *CachingDirectoryFetcherConfiguration) GetMaximumCount() int64 {
	if x != nil {
		return x.MaximumCount
	}
	return 0
}

func (x *CachingDirectoryFetcherConfiguration) GetMaximumSizeBytes() int64 {
	if x != nil {
		return x.MaximumSizeBytes
	}
	return 0
}

func (x *CachingDirectoryFetcherConfiguration) GetCacheReplacementPolicy() eviction.CacheReplacementPolicy {
	if x != nil {
		return x.CacheReplacementPolicy
	}
	return eviction.CacheReplacementPolicy(0)
}

var File_pkg_proto_configuration_cas_cas_proto protoreflect.FileDescriptor

var file_pkg_proto_configuration_cas_cas_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x61, 0x73, 0x2f, 0x63, 0x61,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61,
	0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x63, 0x61, 0x73, 0x1a, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x65, 0x76,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x01, 0x0a, 0x24, 0x43, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x46, 0x65, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23,
	0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x10, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x72, 0x0a, 0x18, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x65, 0x76,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x16, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2f, 0x62, 0x62,
	0x2d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x61, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_configuration_cas_cas_proto_rawDescOnce sync.Once
	file_pkg_proto_configuration_cas_cas_proto_rawDescData = file_pkg_proto_configuration_cas_cas_proto_rawDesc
)

func file_pkg_proto_configuration_cas_cas_proto_rawDescGZIP() []byte {
	file_pkg_proto_configuration_cas_cas_proto_rawDescOnce.Do(func() {
		file_pkg_proto_configuration_cas_cas_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_configuration_cas_cas_proto_rawDescData)
	})
	return file_pkg_proto_configuration_cas_cas_proto_rawDescData
}

var file_pkg_proto_configuration_cas_cas_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_proto_configuration_cas_cas_proto_goTypes = []interface{}{
	(*CachingDirectoryFetcherConfiguration)(nil), // 0: buildbarn.configuration.cas.CachingDirectoryFetcherConfiguration
	(eviction.CacheReplacementPolicy)(0),         // 1: buildbarn.configuration.eviction.CacheReplacementPolicy
}
var file_pkg_proto_configuration_cas_cas_proto_depIdxs = []int32{
	1, // 0: buildbarn.configuration.cas.CachingDirectoryFetcherConfiguration.cache_replacement_policy:type_name -> buildbarn.configuration.eviction.CacheReplacementPolicy
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_proto_configuration_cas_cas_proto_init() }
func file_pkg_proto_configuration_cas_cas_proto_init() {
	if File_pkg_proto_configuration_cas_cas_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_configuration_cas_cas_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CachingDirectoryFetcherConfiguration); i {
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
			RawDescriptor: file_pkg_proto_configuration_cas_cas_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_configuration_cas_cas_proto_goTypes,
		DependencyIndexes: file_pkg_proto_configuration_cas_cas_proto_depIdxs,
		MessageInfos:      file_pkg_proto_configuration_cas_cas_proto_msgTypes,
	}.Build()
	File_pkg_proto_configuration_cas_cas_proto = out.File
	file_pkg_proto_configuration_cas_cas_proto_rawDesc = nil
	file_pkg_proto_configuration_cas_cas_proto_goTypes = nil
	file_pkg_proto_configuration_cas_cas_proto_depIdxs = nil
}
