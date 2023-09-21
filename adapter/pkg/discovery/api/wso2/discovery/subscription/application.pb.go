//
//  Copyright (c) 2023, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: wso2/discovery/subscription/application.proto

package subscription

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

// Application data model
type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId    string            `protobuf:"bytes,1,opt,name=eventId,proto3" json:"eventId,omitempty"`
	Uuid       string            `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name       string            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Owner      string            `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Attributes map[string]string `protobuf:"bytes,5,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_subscription_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_subscription_application_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_subscription_application_proto_rawDescGZIP(), []int{0}
}

func (x *Application) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *Application) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Application) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Application) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Application) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type SecuritySchemes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OAuth2 *OAuth2 `protobuf:"bytes,1,opt,name=oAuth2,proto3" json:"oAuth2,omitempty"`
}

func (x *SecuritySchemes) Reset() {
	*x = SecuritySchemes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_subscription_application_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecuritySchemes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecuritySchemes) ProtoMessage() {}

func (x *SecuritySchemes) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_subscription_application_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecuritySchemes.ProtoReflect.Descriptor instead.
func (*SecuritySchemes) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_subscription_application_proto_rawDescGZIP(), []int{1}
}

func (x *SecuritySchemes) GetOAuth2() *OAuth2 {
	if x != nil {
		return x.OAuth2
	}
	return nil
}

type OAuth2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Environments []*Environment `protobuf:"bytes,1,rep,name=environments,proto3" json:"environments,omitempty"`
}

func (x *OAuth2) Reset() {
	*x = OAuth2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_subscription_application_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuth2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuth2) ProtoMessage() {}

func (x *OAuth2) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_subscription_application_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuth2.ProtoReflect.Descriptor instead.
func (*OAuth2) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_subscription_application_proto_rawDescGZIP(), []int{2}
}

func (x *OAuth2) GetEnvironments() []*Environment {
	if x != nil {
		return x.Environments
	}
	return nil
}

type Environment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvID                 string `protobuf:"bytes,1,opt,name=envID,proto3" json:"envID,omitempty"`
	ApplicationIdentifier string `protobuf:"bytes,2,opt,name=applicationIdentifier,proto3" json:"applicationIdentifier,omitempty"`
	KeyType               string `protobuf:"bytes,3,opt,name=keyType,proto3" json:"keyType,omitempty"`
}

func (x *Environment) Reset() {
	*x = Environment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_subscription_application_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Environment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Environment) ProtoMessage() {}

func (x *Environment) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_subscription_application_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Environment.ProtoReflect.Descriptor instead.
func (*Environment) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_subscription_application_proto_rawDescGZIP(), []int{3}
}

func (x *Environment) GetEnvID() string {
	if x != nil {
		return x.EnvID
	}
	return ""
}

func (x *Environment) GetApplicationIdentifier() string {
	if x != nil {
		return x.ApplicationIdentifier
	}
	return ""
}

func (x *Environment) GetKeyType() string {
	if x != nil {
		return x.KeyType
	}
	return ""
}

var File_wso2_discovery_subscription_application_proto protoreflect.FileDescriptor

var file_wso2_discovery_subscription_application_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1b, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xfe, 0x01, 0x0a,
	0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x58, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x3d,
	0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4e, 0x0a,
	0x0f, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x73,
	0x12, 0x3b, 0x0a, 0x06, 0x6f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f,
	0x41, 0x75, 0x74, 0x68, 0x32, 0x52, 0x06, 0x6f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x22, 0x56, 0x0a,
	0x06, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x12, 0x4c, 0x0a, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x73, 0x0a, 0x0b, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6e, 0x76, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6e, 0x76, 0x49, 0x44, 0x12, 0x34, 0x0a, 0x15, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x42, 0x93, 0x01, 0x0a, 0x2c, 0x6f,
	0x72, 0x67, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x61, 0x70, 0x6b, 0x2e, 0x65, 0x6e, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x10, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x3b, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wso2_discovery_subscription_application_proto_rawDescOnce sync.Once
	file_wso2_discovery_subscription_application_proto_rawDescData = file_wso2_discovery_subscription_application_proto_rawDesc
)

func file_wso2_discovery_subscription_application_proto_rawDescGZIP() []byte {
	file_wso2_discovery_subscription_application_proto_rawDescOnce.Do(func() {
		file_wso2_discovery_subscription_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_wso2_discovery_subscription_application_proto_rawDescData)
	})
	return file_wso2_discovery_subscription_application_proto_rawDescData
}

var file_wso2_discovery_subscription_application_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_wso2_discovery_subscription_application_proto_goTypes = []interface{}{
	(*Application)(nil),     // 0: wso2.discovery.subscription.Application
	(*SecuritySchemes)(nil), // 1: wso2.discovery.subscription.SecuritySchemes
	(*OAuth2)(nil),          // 2: wso2.discovery.subscription.OAuth2
	(*Environment)(nil),     // 3: wso2.discovery.subscription.Environment
	nil,                     // 4: wso2.discovery.subscription.Application.AttributesEntry
}
var file_wso2_discovery_subscription_application_proto_depIdxs = []int32{
	4, // 0: wso2.discovery.subscription.Application.attributes:type_name -> wso2.discovery.subscription.Application.AttributesEntry
	2, // 1: wso2.discovery.subscription.SecuritySchemes.oAuth2:type_name -> wso2.discovery.subscription.OAuth2
	3, // 2: wso2.discovery.subscription.OAuth2.environments:type_name -> wso2.discovery.subscription.Environment
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_wso2_discovery_subscription_application_proto_init() }
func file_wso2_discovery_subscription_application_proto_init() {
	if File_wso2_discovery_subscription_application_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wso2_discovery_subscription_application_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_wso2_discovery_subscription_application_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecuritySchemes); i {
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
		file_wso2_discovery_subscription_application_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuth2); i {
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
		file_wso2_discovery_subscription_application_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Environment); i {
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
			RawDescriptor: file_wso2_discovery_subscription_application_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wso2_discovery_subscription_application_proto_goTypes,
		DependencyIndexes: file_wso2_discovery_subscription_application_proto_depIdxs,
		MessageInfos:      file_wso2_discovery_subscription_application_proto_msgTypes,
	}.Build()
	File_wso2_discovery_subscription_application_proto = out.File
	file_wso2_discovery_subscription_application_proto_rawDesc = nil
	file_wso2_discovery_subscription_application_proto_goTypes = nil
	file_wso2_discovery_subscription_application_proto_depIdxs = nil
}
