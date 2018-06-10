// Code generated by protoc-gen-go. DO NOT EDIT.
// source: facebook/about_you.proto

/*
Package grain_facebook is a generated protocol buffer package.

It is generated from these files:
	facebook/about_you.proto
	facebook/ads.proto
	facebook/apps.proto
	facebook/comments.proto
	facebook/events.proto
	facebook/following_and_followers.proto
	facebook/friends.proto
	facebook/groups.proto

It has these top-level messages:
	AddressBook
	Ads
	InstalledApps
	PostsFromApps
	YourApps
	Comment
	PhotoMetadata
	MediaMetadata
	Media
	AttachmentData
	Attachment
	Post
	YourPosts
	Coordinate
	Event
	EventInvitations
	EventResponses
	YourEvents
	FollowedPages
	Friend
	FriendsAdded
	RejectedRequests
	RemovedFriends
	SentRequests
	GroupsYouManage
	GroupMembershipActivity
	GroupPostsComments
*/
package grain_facebook

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddressBook struct {
	AddressBook *AddressBook_AddressBookEntry `protobuf:"bytes,1,opt,name=address_book,json=addressBook" json:"address_book,omitempty"`
}

func (m *AddressBook) Reset()                    { *m = AddressBook{} }
func (m *AddressBook) String() string            { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()               {}
func (*AddressBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddressBook) GetAddressBook() *AddressBook_AddressBookEntry {
	if m != nil {
		return m.AddressBook
	}
	return nil
}

type AddressBook_Details struct {
	ContactPoint string `protobuf:"bytes,1,opt,name=contact_point,json=contactPoint" json:"contact_point,omitempty"`
}

func (m *AddressBook_Details) Reset()                    { *m = AddressBook_Details{} }
func (m *AddressBook_Details) String() string            { return proto.CompactTextString(m) }
func (*AddressBook_Details) ProtoMessage()               {}
func (*AddressBook_Details) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *AddressBook_Details) GetContactPoint() string {
	if m != nil {
		return m.ContactPoint
	}
	return ""
}

type AddressBook_Contact struct {
	Name    string                 `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Details []*AddressBook_Details `protobuf:"bytes,2,rep,name=details" json:"details,omitempty"`
}

func (m *AddressBook_Contact) Reset()                    { *m = AddressBook_Contact{} }
func (m *AddressBook_Contact) String() string            { return proto.CompactTextString(m) }
func (*AddressBook_Contact) ProtoMessage()               {}
func (*AddressBook_Contact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *AddressBook_Contact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddressBook_Contact) GetDetails() []*AddressBook_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

type AddressBook_AddressBookEntry struct {
	AddressBook []*AddressBook_Contact `protobuf:"bytes,1,rep,name=address_book,json=addressBook" json:"address_book,omitempty"`
}

func (m *AddressBook_AddressBookEntry) Reset()                    { *m = AddressBook_AddressBookEntry{} }
func (m *AddressBook_AddressBookEntry) String() string            { return proto.CompactTextString(m) }
func (*AddressBook_AddressBookEntry) ProtoMessage()               {}
func (*AddressBook_AddressBookEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

func (m *AddressBook_AddressBookEntry) GetAddressBook() []*AddressBook_Contact {
	if m != nil {
		return m.AddressBook
	}
	return nil
}

func init() {
	proto.RegisterType((*AddressBook)(nil), "grain.facebook.AddressBook")
	proto.RegisterType((*AddressBook_Details)(nil), "grain.facebook.AddressBook.Details")
	proto.RegisterType((*AddressBook_Contact)(nil), "grain.facebook.AddressBook.Contact")
	proto.RegisterType((*AddressBook_AddressBookEntry)(nil), "grain.facebook.AddressBook.AddressBookEntry")
}

func init() { proto.RegisterFile("facebook/about_you.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0x4b, 0x4c, 0x4e,
	0x4d, 0xca, 0xcf, 0xcf, 0xd6, 0x4f, 0x4c, 0xca, 0x2f, 0x2d, 0x89, 0xaf, 0xcc, 0x2f, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4b, 0x2f, 0x4a, 0xcc, 0xcc, 0xd3, 0x83, 0xc9, 0x2b, 0x9d,
	0x60, 0xe2, 0xe2, 0x76, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x76, 0xca, 0xcf, 0xcf, 0x16, 0xf2,
	0xe7, 0xe2, 0x49, 0x84, 0x70, 0xe3, 0x41, 0xf2, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x3a,
	0x7a, 0xa8, 0xda, 0xf4, 0x90, 0xb4, 0x20, 0xb3, 0x5d, 0xf3, 0x4a, 0x8a, 0x2a, 0x83, 0xb8, 0x13,
	0x11, 0x22, 0x52, 0x7a, 0x5c, 0xec, 0x2e, 0xa9, 0x25, 0x89, 0x99, 0x39, 0xc5, 0x42, 0xca, 0x5c,
	0xbc, 0xc9, 0xf9, 0x79, 0x25, 0x89, 0xc9, 0x25, 0xf1, 0x05, 0xf9, 0x99, 0x79, 0x25, 0x60, 0xc3,
	0x39, 0x83, 0x78, 0xa0, 0x82, 0x01, 0x20, 0x31, 0xa9, 0x18, 0x2e, 0x76, 0x67, 0x08, 0x5f, 0x48,
	0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x15, 0xaa, 0x0c, 0xcc, 0x16, 0xb2, 0xe5, 0x62, 0x4f, 0x81,
	0x18, 0x27, 0xc1, 0xa4, 0xc0, 0xac, 0xc1, 0x6d, 0xa4, 0x8c, 0xcf, 0x69, 0x50, 0x9b, 0x83, 0x60,
	0x7a, 0xa4, 0xa2, 0xb8, 0x04, 0xd0, 0x9d, 0x2b, 0xe4, 0x86, 0xe1, 0x65, 0x82, 0xe6, 0x42, 0x5d,
	0x88, 0xe2, 0xd3, 0x24, 0x36, 0x70, 0x08, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xae, 0x0f,
	0xda, 0x15, 0x7d, 0x01, 0x00, 0x00,
}