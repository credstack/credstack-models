// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: proto/user/user.proto

package user

import (
	header "github.com/credstack/credstack-models/proto/header"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// User - Represents a user registered with CredStack. Describes all standard claims for the OpenID Connect spec
type User struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// header - Shared data values used across all objects
	Header *header.Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty" bson:"header"` // @gotags: bson:"header"
	// sub - A UUID v5 based on the users email address, used for identifying the user. This is the same as the identifier in the header
	Sub string `protobuf:"bytes,2,opt,name=sub,proto3" json:"sub,omitempty" bson:"sub"` // @gotags: bson:"sub"
	// username - The username of the user. Does not need to be unique as primary lookup for the user is done via userId
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty" bson:"username"` // @gotags: bson:"username"
	// email - The primary email address for the user. Must be unique
	Email string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty" bson:"email"` // @gotags: bson:"email"
	// email_verified - A boolean value describing if the user has validated there email address
	EmailVerified bool `protobuf:"varint,5,opt,name=email_verified,json=emailVerified,proto3" json:"email_verified,omitempty" bson:"email_verified"` // @gotags: bson:"email_verified"
	// given_name - The users first name. Can be null, and does not need to be unique
	GivenName string `protobuf:"bytes,6,opt,name=given_name,json=givenName,proto3" json:"given_name,omitempty" bson:"given_name"` // @gotags: bson:"given_name"
	// middle_name - The users middle name. Can be null, and does not need to be unique
	MiddleName string `protobuf:"bytes,7,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty" bson:"middle_name"` // @gotags: bson:"middle_name"
	// family_name - The users last name. Can be null, and does not need to be unique
	FamilyName string `protobuf:"bytes,8,opt,name=family_name,json=familyName,proto3" json:"family_name,omitempty" bson:"family_name"` // @gotags: bson:"family_name"
	// gender - The users gender. Can be null
	Gender string `protobuf:"bytes,9,opt,name=gender,proto3" json:"gender,omitempty" bson:"gender"` // @gotags: bson:"gender"
	// birth_date - The users birthdate represented as YYYY-MM-DD
	BirthDate string `protobuf:"bytes,10,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty" bson:"birth_date"` // @gotags: bson:"birth_date"
	// zone_info - The timezone for the user that should be respected by timestamps. Defaults to America/New_York
	ZoneInfo string `protobuf:"bytes,11,opt,name=zone_info,json=zoneInfo,proto3" json:"zone_info,omitempty" bson:"zone_info"` // @gotags: bson:"zone_info"
	// phone_number - The users phone number in the following format +1800-555-555
	PhoneNumber string `protobuf:"bytes,12,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty" bson:"phone_number"` // @gotags: bson:"phone_number"
	// phone_number_verified - A boolean value describing if the user has validated there email address
	PhoneNumberVerified bool `protobuf:"varint,13,opt,name=phone_number_verified,json=phoneNumberVerified,proto3" json:"phone_number_verified,omitempty" bson:"phone_number_verified"` // @gotags: bson:"phone_number_verified"
	// address - The users physical mailing address. Can be null
	Address string `protobuf:"bytes,14,opt,name=address,proto3" json:"address,omitempty" bson:"address"` // @gotags: bson:"address"
	// credential - A message representing the users hashed password. Will be left null when exporting unless explicitly declared
	Credential *UserCredential `protobuf:"bytes,15,opt,name=credential,proto3" json:"credential,omitempty" bson:"credential"` // @gotags: bson:"credential"
	// scopes - A list of scope ID's that were directly assigned to the user
	Scopes []string `protobuf:"bytes,16,rep,name=scopes,proto3" json:"scopes,omitempty" bson:"scopes"` // @gotags: bson:"scopes"
	// roles - A list of role ID's that were directly assigned to the user
	Roles         []string `protobuf:"bytes,17,rep,name=roles,proto3" json:"roles,omitempty" bson:"roles"` // @gotags: bson:"roles"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_proto_user_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[0]
	if x != nil {
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
	return file_proto_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetHeader() *header.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *User) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetEmailVerified() bool {
	if x != nil {
		return x.EmailVerified
	}
	return false
}

func (x *User) GetGivenName() string {
	if x != nil {
		return x.GivenName
	}
	return ""
}

func (x *User) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *User) GetFamilyName() string {
	if x != nil {
		return x.FamilyName
	}
	return ""
}

func (x *User) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *User) GetBirthDate() string {
	if x != nil {
		return x.BirthDate
	}
	return ""
}

func (x *User) GetZoneInfo() string {
	if x != nil {
		return x.ZoneInfo
	}
	return ""
}

func (x *User) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *User) GetPhoneNumberVerified() bool {
	if x != nil {
		return x.PhoneNumberVerified
	}
	return false
}

func (x *User) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *User) GetCredential() *UserCredential {
	if x != nil {
		return x.Credential
	}
	return nil
}

func (x *User) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *User) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

var File_proto_user_user_proto protoreflect.FileDescriptor

const file_proto_user_user_proto_rawDesc = "" +
	"\n" +
	"\x15proto/user/user.proto\x12\x04user\x1a\x1bproto/user/credential.proto\x1a\x19proto/header/header.proto\"\xa3\x04\n" +
	"\x04User\x12&\n" +
	"\x06header\x18\x01 \x01(\v2\x0e.header.HeaderR\x06header\x12\x10\n" +
	"\x03sub\x18\x02 \x01(\tR\x03sub\x12\x1a\n" +
	"\busername\x18\x03 \x01(\tR\busername\x12\x14\n" +
	"\x05email\x18\x04 \x01(\tR\x05email\x12%\n" +
	"\x0eemail_verified\x18\x05 \x01(\bR\remailVerified\x12\x1d\n" +
	"\n" +
	"given_name\x18\x06 \x01(\tR\tgivenName\x12\x1f\n" +
	"\vmiddle_name\x18\a \x01(\tR\n" +
	"middleName\x12\x1f\n" +
	"\vfamily_name\x18\b \x01(\tR\n" +
	"familyName\x12\x16\n" +
	"\x06gender\x18\t \x01(\tR\x06gender\x12\x1d\n" +
	"\n" +
	"birth_date\x18\n" +
	" \x01(\tR\tbirthDate\x12\x1b\n" +
	"\tzone_info\x18\v \x01(\tR\bzoneInfo\x12!\n" +
	"\fphone_number\x18\f \x01(\tR\vphoneNumber\x122\n" +
	"\x15phone_number_verified\x18\r \x01(\bR\x13phoneNumberVerified\x12\x18\n" +
	"\aaddress\x18\x0e \x01(\tR\aaddress\x124\n" +
	"\n" +
	"credential\x18\x0f \x01(\v2\x14.user.UserCredentialR\n" +
	"credential\x12\x16\n" +
	"\x06scopes\x18\x10 \x03(\tR\x06scopes\x12\x14\n" +
	"\x05roles\x18\x11 \x03(\tR\x05rolesB2Z0github.com/credstack/credstack-models/proto/userb\x06proto3"

var (
	file_proto_user_user_proto_rawDescOnce sync.Once
	file_proto_user_user_proto_rawDescData []byte
)

func file_proto_user_user_proto_rawDescGZIP() []byte {
	file_proto_user_user_proto_rawDescOnce.Do(func() {
		file_proto_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_user_user_proto_rawDesc), len(file_proto_user_user_proto_rawDesc)))
	})
	return file_proto_user_user_proto_rawDescData
}

var file_proto_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_user_user_proto_goTypes = []any{
	(*User)(nil),           // 0: user.User
	(*header.Header)(nil),  // 1: header.Header
	(*UserCredential)(nil), // 2: user.UserCredential
}
var file_proto_user_user_proto_depIdxs = []int32{
	1, // 0: user.User.header:type_name -> header.Header
	2, // 1: user.User.credential:type_name -> user.UserCredential
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_user_user_proto_init() }
func file_proto_user_user_proto_init() {
	if File_proto_user_user_proto != nil {
		return
	}
	file_proto_user_credential_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_user_user_proto_rawDesc), len(file_proto_user_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_user_user_proto_goTypes,
		DependencyIndexes: file_proto_user_user_proto_depIdxs,
		MessageInfos:      file_proto_user_user_proto_msgTypes,
	}.Build()
	File_proto_user_user_proto = out.File
	file_proto_user_user_proto_goTypes = nil
	file_proto_user_user_proto_depIdxs = nil
}
