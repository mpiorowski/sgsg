// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: auth.proto

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

type Role int32

const (
	Role_ROLE_UNSET Role = 0
	Role_ROLE_USER  Role = 1
	Role_ROLE_ADMIN Role = 2
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "ROLE_UNSET",
		1: "ROLE_USER",
		2: "ROLE_ADMIN",
	}
	Role_value = map[string]int32{
		"ROLE_UNSET": 0,
		"ROLE_USER":  1,
		"ROLE_ADMIN": 2,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_auth_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Created            string `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated            string `protobuf:"bytes,3,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted            string `protobuf:"bytes,4,opt,name=deleted,proto3" json:"deleted,omitempty"`
	Email              string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Sub                string `protobuf:"bytes,6,opt,name=sub,proto3" json:"sub,omitempty"`
	Role               Role   `protobuf:"varint,7,opt,name=role,proto3,enum=proto.Role" json:"role,omitempty"`
	Avatar             string `protobuf:"bytes,8,opt,name=avatar,proto3" json:"avatar,omitempty"`
	SubscriptionId     string `protobuf:"bytes,9,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	SubscriptionEnd    string `protobuf:"bytes,10,opt,name=subscription_end,json=subscriptionEnd,proto3" json:"subscription_end,omitempty"`
	SubscriptionCheck  string `protobuf:"bytes,11,opt,name=subscription_check,json=subscriptionCheck,proto3" json:"subscription_check,omitempty"`
	SubscriptionActive bool   `protobuf:"varint,12,opt,name=subscription_active,json=subscriptionActive,proto3" json:"subscription_active,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *User) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

func (x *User) GetDeleted() string {
	if x != nil {
		return x.Deleted
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

func (x *User) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_ROLE_UNSET
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

func (x *User) GetSubscriptionEnd() string {
	if x != nil {
		return x.SubscriptionEnd
	}
	return ""
}

func (x *User) GetSubscriptionCheck() string {
	if x != nil {
		return x.SubscriptionCheck
	}
	return ""
}

func (x *User) GetSubscriptionActive() bool {
	if x != nil {
		return x.SubscriptionActive
	}
	return false
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x75, 0x62, 0x12, 0x1f, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x12,
	0x2d, 0x0a, 0x12, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x2f,
	0x0a, 0x13, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2a,
	0x35, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x4f, 0x4c, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x4f, 0x4c, 0x45, 0x5f,
	0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x41,
	0x44, 0x4d, 0x49, 0x4e, 0x10, 0x02, 0x42, 0x0c, 0x5a, 0x0a, 0x73, 0x67, 0x73, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_auth_proto_goTypes = []interface{}{
	(Role)(0),    // 0: proto.Role
	(*User)(nil), // 1: proto.User
}
var file_auth_proto_depIdxs = []int32{
	0, // 0: proto.User.role:type_name -> proto.Role
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		EnumInfos:         file_auth_proto_enumTypes,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}