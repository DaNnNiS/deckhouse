// Copyright 2024 Flant JSC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: dhctl.proto

package dhctl

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_dhctl_proto protoreflect.FileDescriptor

var file_dhctl_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x64,
	0x68, 0x63, 0x74, 0x6c, 0x1a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x87, 0x01, 0x0a, 0x05, 0x44, 0x48, 0x43, 0x54, 0x4c, 0x12, 0x38, 0x0a, 0x05,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x13, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x64, 0x68, 0x63,
	0x74, 0x6c, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x44, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74,
	0x72, 0x61, 0x70, 0x12, 0x17, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x42, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x64,
	0x68, 0x63, 0x74, 0x6c, 0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0a, 0x5a, 0x08,
	0x70, 0x62, 0x2f, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_dhctl_proto_goTypes = []interface{}{
	(*CheckRequest)(nil),      // 0: dhctl.CheckRequest
	(*BootstrapRequest)(nil),  // 1: dhctl.BootstrapRequest
	(*CheckResponse)(nil),     // 2: dhctl.CheckResponse
	(*BootstrapResponse)(nil), // 3: dhctl.BootstrapResponse
}
var file_dhctl_proto_depIdxs = []int32{
	0, // 0: dhctl.DHCTL.Check:input_type -> dhctl.CheckRequest
	1, // 1: dhctl.DHCTL.Bootstrap:input_type -> dhctl.BootstrapRequest
	2, // 2: dhctl.DHCTL.Check:output_type -> dhctl.CheckResponse
	3, // 3: dhctl.DHCTL.Bootstrap:output_type -> dhctl.BootstrapResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dhctl_proto_init() }
func file_dhctl_proto_init() {
	if File_dhctl_proto != nil {
		return
	}
	file_check_proto_init()
	file_bootstrap_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dhctl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dhctl_proto_goTypes,
		DependencyIndexes: file_dhctl_proto_depIdxs,
	}.Build()
	File_dhctl_proto = out.File
	file_dhctl_proto_rawDesc = nil
	file_dhctl_proto_goTypes = nil
	file_dhctl_proto_depIdxs = nil
}
