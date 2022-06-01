// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: events/events.proto

package events

import (
	_ "github.com/clubo-app/protobuf/party"
	profile "github.com/clubo-app/protobuf/profile"
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

type Registered struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Firstname string `protobuf:"bytes,4,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,5,opt,name=lastname,proto3" json:"lastname,omitempty"`
}

func (x *Registered) Reset() {
	*x = Registered{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registered) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registered) ProtoMessage() {}

func (x *Registered) ProtoReflect() protoreflect.Message {
	mi := &file_events_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registered.ProtoReflect.Descriptor instead.
func (*Registered) Descriptor() ([]byte, []int) {
	return file_events_events_proto_rawDescGZIP(), []int{0}
}

func (x *Registered) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Registered) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Registered) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Registered) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *Registered) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

// id must always be defined but for the rest just define the properties that got updated
type ProfileUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *profile.Profile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *ProfileUpdated) Reset() {
	*x = ProfileUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileUpdated) ProtoMessage() {}

func (x *ProfileUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_events_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileUpdated.ProtoReflect.Descriptor instead.
func (*ProfileUpdated) Descriptor() ([]byte, []int) {
	return file_events_events_proto_rawDescGZIP(), []int{1}
}

func (x *ProfileUpdated) GetProfile() *profile.Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

type PartyCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title         string  `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	IsPublic      bool    `protobuf:"varint,4,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
	Lat           float32 `protobuf:"fixed32,5,opt,name=lat,proto3" json:"lat,omitempty"`
	Long          float32 `protobuf:"fixed32,6,opt,name=long,proto3" json:"long,omitempty"`
	Geohash       string  `protobuf:"bytes,7,opt,name=geohash,proto3" json:"geohash,omitempty"`
	StreetAddress string  `protobuf:"bytes,8,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	PostalCode    string  `protobuf:"bytes,9,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	State         string  `protobuf:"bytes,10,opt,name=state,proto3" json:"state,omitempty"`
	Country       string  `protobuf:"bytes,11,opt,name=country,proto3" json:"country,omitempty"`
	StartDate     string  `protobuf:"bytes,12,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	CreatedAt     string  `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *PartyCreated) Reset() {
	*x = PartyCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyCreated) ProtoMessage() {}

func (x *PartyCreated) ProtoReflect() protoreflect.Message {
	mi := &file_events_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyCreated.ProtoReflect.Descriptor instead.
func (*PartyCreated) Descriptor() ([]byte, []int) {
	return file_events_events_proto_rawDescGZIP(), []int{2}
}

func (x *PartyCreated) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PartyCreated) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PartyCreated) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PartyCreated) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

func (x *PartyCreated) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *PartyCreated) GetLong() float32 {
	if x != nil {
		return x.Long
	}
	return 0
}

func (x *PartyCreated) GetGeohash() string {
	if x != nil {
		return x.Geohash
	}
	return ""
}

func (x *PartyCreated) GetStreetAddress() string {
	if x != nil {
		return x.StreetAddress
	}
	return ""
}

func (x *PartyCreated) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *PartyCreated) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *PartyCreated) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *PartyCreated) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *PartyCreated) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type PartyUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title         string  `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	IsPublic      bool    `protobuf:"varint,4,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
	Lat           float32 `protobuf:"fixed32,5,opt,name=lat,proto3" json:"lat,omitempty"`
	Long          float32 `protobuf:"fixed32,6,opt,name=long,proto3" json:"long,omitempty"`
	StreetAddress string  `protobuf:"bytes,7,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	PostalCode    string  `protobuf:"bytes,8,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	State         string  `protobuf:"bytes,9,opt,name=state,proto3" json:"state,omitempty"`
	Country       string  `protobuf:"bytes,10,opt,name=country,proto3" json:"country,omitempty"`
	StartDate     string  `protobuf:"bytes,11,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	CreatedAt     string  `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *PartyUpdated) Reset() {
	*x = PartyUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyUpdated) ProtoMessage() {}

func (x *PartyUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_events_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyUpdated.ProtoReflect.Descriptor instead.
func (*PartyUpdated) Descriptor() ([]byte, []int) {
	return file_events_events_proto_rawDescGZIP(), []int{3}
}

func (x *PartyUpdated) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PartyUpdated) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PartyUpdated) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PartyUpdated) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

func (x *PartyUpdated) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *PartyUpdated) GetLong() float32 {
	if x != nil {
		return x.Long
	}
	return 0
}

func (x *PartyUpdated) GetStreetAddress() string {
	if x != nil {
		return x.StreetAddress
	}
	return ""
}

func (x *PartyUpdated) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *PartyUpdated) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *PartyUpdated) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *PartyUpdated) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *PartyUpdated) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type FriendRemoved struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FriendId string `protobuf:"bytes,2,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
}

func (x *FriendRemoved) Reset() {
	*x = FriendRemoved{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendRemoved) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendRemoved) ProtoMessage() {}

func (x *FriendRemoved) ProtoReflect() protoreflect.Message {
	mi := &file_events_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendRemoved.ProtoReflect.Descriptor instead.
func (*FriendRemoved) Descriptor() ([]byte, []int) {
	return file_events_events_proto_rawDescGZIP(), []int{4}
}

func (x *FriendRemoved) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *FriendRemoved) GetFriendId() string {
	if x != nil {
		return x.FriendId
	}
	return ""
}

type FriendRequested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FriendId string `protobuf:"bytes,2,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
}

func (x *FriendRequested) Reset() {
	*x = FriendRequested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_events_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendRequested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendRequested) ProtoMessage() {}

func (x *FriendRequested) ProtoReflect() protoreflect.Message {
	mi := &file_events_events_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendRequested.ProtoReflect.Descriptor instead.
func (*FriendRequested) Descriptor() ([]byte, []int) {
	return file_events_events_proto_rawDescGZIP(), []int{5}
}

func (x *FriendRequested) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *FriendRequested) GetFriendId() string {
	if x != nil {
		return x.FriendId
	}
	return ""
}

type FriendAccepted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FriendId string `protobuf:"bytes,2,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
}

func (x *FriendAccepted) Reset() {
	*x = FriendAccepted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_events_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendAccepted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendAccepted) ProtoMessage() {}

func (x *FriendAccepted) ProtoReflect() protoreflect.Message {
	mi := &file_events_events_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendAccepted.ProtoReflect.Descriptor instead.
func (*FriendAccepted) Descriptor() ([]byte, []int) {
	return file_events_events_proto_rawDescGZIP(), []int{6}
}

func (x *FriendAccepted) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *FriendAccepted) GetFriendId() string {
	if x != nil {
		return x.FriendId
	}
	return ""
}

var File_events_events_proto protoreflect.FileDescriptor

var file_events_events_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x11, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x22, 0xe0, 0x02, 0x0a, 0x0c, 0x50, 0x61, 0x72, 0x74, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x10, 0x0a,
	0x03, 0x6c, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6c,
	0x6f, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x65, 0x6f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x65, 0x6f, 0x68, 0x61, 0x73, 0x68, 0x12, 0x25, 0x0a,
	0x0e, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61,
	0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xc6, 0x02, 0x0a, 0x0c, 0x50, 0x61, 0x72, 0x74, 0x79, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c,
	0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x45, 0x0a, 0x0d,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x0f, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x0e,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x49, 0x64, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x75, 0x62, 0x6f, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_events_events_proto_rawDescOnce sync.Once
	file_events_events_proto_rawDescData = file_events_events_proto_rawDesc
)

func file_events_events_proto_rawDescGZIP() []byte {
	file_events_events_proto_rawDescOnce.Do(func() {
		file_events_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_events_events_proto_rawDescData)
	})
	return file_events_events_proto_rawDescData
}

var file_events_events_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_events_events_proto_goTypes = []interface{}{
	(*Registered)(nil),      // 0: events.Registered
	(*ProfileUpdated)(nil),  // 1: events.ProfileUpdated
	(*PartyCreated)(nil),    // 2: events.PartyCreated
	(*PartyUpdated)(nil),    // 3: events.PartyUpdated
	(*FriendRemoved)(nil),   // 4: events.FriendRemoved
	(*FriendRequested)(nil), // 5: events.FriendRequested
	(*FriendAccepted)(nil),  // 6: events.FriendAccepted
	(*profile.Profile)(nil), // 7: profile.Profile
}
var file_events_events_proto_depIdxs = []int32{
	7, // 0: events.ProfileUpdated.profile:type_name -> profile.Profile
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_events_events_proto_init() }
func file_events_events_proto_init() {
	if File_events_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_events_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registered); i {
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
		file_events_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileUpdated); i {
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
		file_events_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyCreated); i {
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
		file_events_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyUpdated); i {
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
		file_events_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendRemoved); i {
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
		file_events_events_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendRequested); i {
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
		file_events_events_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendAccepted); i {
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
			RawDescriptor: file_events_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_events_events_proto_goTypes,
		DependencyIndexes: file_events_events_proto_depIdxs,
		MessageInfos:      file_events_events_proto_msgTypes,
	}.Build()
	File_events_events_proto = out.File
	file_events_events_proto_rawDesc = nil
	file_events_events_proto_goTypes = nil
	file_events_events_proto_depIdxs = nil
}
