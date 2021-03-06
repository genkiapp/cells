/*
 * Copyright (c) 2018. Abstrium SAS <team (at) pydio.com>
 * This file is part of Pydio Cells.
 *
 * Pydio Cells is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio Cells is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio Cells.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

/*
Package chat is a generated protocol buffer package.

It is generated from these files:
	chat.proto

It has these top-level messages:
	ChatRoom
	ChatMessage
	PutRoomRequest
	PutRoomResponse
	PostMessageRequest
	PostMessageResponse
	DeleteMessageRequest
	DeleteMessageResponse
	ListMessagesRequest
	ListMessagesResponse
	ListRoomsRequest
	ListRoomsResponse
	DeleteRoomRequest
	DeleteRoomResponse
	ChatEvent
	WebSocketMessage
*/
package chat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import activity "github.com/pydio/cells/common/proto/activity"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RoomType int32

const (
	RoomType_GLOBAL    RoomType = 0
	RoomType_WORKSPACE RoomType = 1
	RoomType_USER      RoomType = 2
	RoomType_NODE      RoomType = 3
)

var RoomType_name = map[int32]string{
	0: "GLOBAL",
	1: "WORKSPACE",
	2: "USER",
	3: "NODE",
}
var RoomType_value = map[string]int32{
	"GLOBAL":    0,
	"WORKSPACE": 1,
	"USER":      2,
	"NODE":      3,
}

func (x RoomType) String() string {
	return proto.EnumName(RoomType_name, int32(x))
}
func (RoomType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type WsMessageType int32

const (
	WsMessageType_JOIN        WsMessageType = 0
	WsMessageType_LEAVE       WsMessageType = 1
	WsMessageType_POST        WsMessageType = 2
	WsMessageType_ROOM_UPDATE WsMessageType = 3
	WsMessageType_HISTORY     WsMessageType = 4
	WsMessageType_DELETE_MSG  WsMessageType = 5
	WsMessageType_DELETE_ROOM WsMessageType = 6
)

var WsMessageType_name = map[int32]string{
	0: "JOIN",
	1: "LEAVE",
	2: "POST",
	3: "ROOM_UPDATE",
	4: "HISTORY",
	5: "DELETE_MSG",
	6: "DELETE_ROOM",
}
var WsMessageType_value = map[string]int32{
	"JOIN":        0,
	"LEAVE":       1,
	"POST":        2,
	"ROOM_UPDATE": 3,
	"HISTORY":     4,
	"DELETE_MSG":  5,
	"DELETE_ROOM": 6,
}

func (x WsMessageType) String() string {
	return proto.EnumName(WsMessageType_name, int32(x))
}
func (WsMessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ChatRoom struct {
	Uuid           string   `protobuf:"bytes,1,opt,name=Uuid" json:"Uuid,omitempty"`
	Type           RoomType `protobuf:"varint,2,opt,name=Type,enum=chat.RoomType" json:"Type,omitempty"`
	RoomTypeObject string   `protobuf:"bytes,3,opt,name=RoomTypeObject" json:"RoomTypeObject,omitempty"`
	RoomLabel      string   `protobuf:"bytes,4,opt,name=RoomLabel" json:"RoomLabel,omitempty"`
	Users          []string `protobuf:"bytes,5,rep,name=Users" json:"Users,omitempty"`
	LastUpdated    int32    `protobuf:"varint,6,opt,name=LastUpdated" json:"LastUpdated,omitempty"`
}

func (m *ChatRoom) Reset()                    { *m = ChatRoom{} }
func (m *ChatRoom) String() string            { return proto.CompactTextString(m) }
func (*ChatRoom) ProtoMessage()               {}
func (*ChatRoom) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ChatRoom) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ChatRoom) GetType() RoomType {
	if m != nil {
		return m.Type
	}
	return RoomType_GLOBAL
}

func (m *ChatRoom) GetRoomTypeObject() string {
	if m != nil {
		return m.RoomTypeObject
	}
	return ""
}

func (m *ChatRoom) GetRoomLabel() string {
	if m != nil {
		return m.RoomLabel
	}
	return ""
}

func (m *ChatRoom) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *ChatRoom) GetLastUpdated() int32 {
	if m != nil {
		return m.LastUpdated
	}
	return 0
}

type ChatMessage struct {
	Uuid      string           `protobuf:"bytes,1,opt,name=Uuid" json:"Uuid,omitempty"`
	RoomUuid  string           `protobuf:"bytes,2,opt,name=RoomUuid" json:"RoomUuid,omitempty"`
	Message   string           `protobuf:"bytes,3,opt,name=Message" json:"Message,omitempty"`
	Author    string           `protobuf:"bytes,4,opt,name=Author" json:"Author,omitempty"`
	Timestamp int64            `protobuf:"varint,5,opt,name=Timestamp" json:"Timestamp,omitempty"`
	Activity  *activity.Object `protobuf:"bytes,6,opt,name=Activity" json:"Activity,omitempty"`
}

func (m *ChatMessage) Reset()                    { *m = ChatMessage{} }
func (m *ChatMessage) String() string            { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()               {}
func (*ChatMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ChatMessage) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ChatMessage) GetRoomUuid() string {
	if m != nil {
		return m.RoomUuid
	}
	return ""
}

func (m *ChatMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ChatMessage) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *ChatMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ChatMessage) GetActivity() *activity.Object {
	if m != nil {
		return m.Activity
	}
	return nil
}

type PutRoomRequest struct {
	Room *ChatRoom `protobuf:"bytes,1,opt,name=Room" json:"Room,omitempty"`
}

func (m *PutRoomRequest) Reset()                    { *m = PutRoomRequest{} }
func (m *PutRoomRequest) String() string            { return proto.CompactTextString(m) }
func (*PutRoomRequest) ProtoMessage()               {}
func (*PutRoomRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PutRoomRequest) GetRoom() *ChatRoom {
	if m != nil {
		return m.Room
	}
	return nil
}

type PutRoomResponse struct {
	Room *ChatRoom `protobuf:"bytes,1,opt,name=Room" json:"Room,omitempty"`
}

func (m *PutRoomResponse) Reset()                    { *m = PutRoomResponse{} }
func (m *PutRoomResponse) String() string            { return proto.CompactTextString(m) }
func (*PutRoomResponse) ProtoMessage()               {}
func (*PutRoomResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PutRoomResponse) GetRoom() *ChatRoom {
	if m != nil {
		return m.Room
	}
	return nil
}

type PostMessageRequest struct {
	Messages []*ChatMessage `protobuf:"bytes,1,rep,name=Messages" json:"Messages,omitempty"`
}

func (m *PostMessageRequest) Reset()                    { *m = PostMessageRequest{} }
func (m *PostMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*PostMessageRequest) ProtoMessage()               {}
func (*PostMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PostMessageRequest) GetMessages() []*ChatMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

type PostMessageResponse struct {
	Success  bool           `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
	Messages []*ChatMessage `protobuf:"bytes,2,rep,name=Messages" json:"Messages,omitempty"`
}

func (m *PostMessageResponse) Reset()                    { *m = PostMessageResponse{} }
func (m *PostMessageResponse) String() string            { return proto.CompactTextString(m) }
func (*PostMessageResponse) ProtoMessage()               {}
func (*PostMessageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PostMessageResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PostMessageResponse) GetMessages() []*ChatMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

type DeleteMessageRequest struct {
	Messages []*ChatMessage `protobuf:"bytes,1,rep,name=Messages" json:"Messages,omitempty"`
}

func (m *DeleteMessageRequest) Reset()                    { *m = DeleteMessageRequest{} }
func (m *DeleteMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteMessageRequest) ProtoMessage()               {}
func (*DeleteMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteMessageRequest) GetMessages() []*ChatMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

type DeleteMessageResponse struct {
	Success bool `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
}

func (m *DeleteMessageResponse) Reset()                    { *m = DeleteMessageResponse{} }
func (m *DeleteMessageResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteMessageResponse) ProtoMessage()               {}
func (*DeleteMessageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DeleteMessageResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ListMessagesRequest struct {
	RoomUuid string `protobuf:"bytes,1,opt,name=RoomUuid" json:"RoomUuid,omitempty"`
	// List starting at a given message ID
	LastMessage string `protobuf:"bytes,2,opt,name=LastMessage" json:"LastMessage,omitempty"`
	Offset      int64  `protobuf:"varint,3,opt,name=Offset" json:"Offset,omitempty"`
	Limit       int64  `protobuf:"varint,4,opt,name=Limit" json:"Limit,omitempty"`
}

func (m *ListMessagesRequest) Reset()                    { *m = ListMessagesRequest{} }
func (m *ListMessagesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListMessagesRequest) ProtoMessage()               {}
func (*ListMessagesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ListMessagesRequest) GetRoomUuid() string {
	if m != nil {
		return m.RoomUuid
	}
	return ""
}

func (m *ListMessagesRequest) GetLastMessage() string {
	if m != nil {
		return m.LastMessage
	}
	return ""
}

func (m *ListMessagesRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListMessagesRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ListMessagesResponse struct {
	Message *ChatMessage `protobuf:"bytes,1,opt,name=Message" json:"Message,omitempty"`
}

func (m *ListMessagesResponse) Reset()                    { *m = ListMessagesResponse{} }
func (m *ListMessagesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListMessagesResponse) ProtoMessage()               {}
func (*ListMessagesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ListMessagesResponse) GetMessage() *ChatMessage {
	if m != nil {
		return m.Message
	}
	return nil
}

type ListRoomsRequest struct {
	ByType     RoomType `protobuf:"varint,1,opt,name=ByType,enum=chat.RoomType" json:"ByType,omitempty"`
	TypeObject string   `protobuf:"bytes,2,opt,name=TypeObject" json:"TypeObject,omitempty"`
}

func (m *ListRoomsRequest) Reset()                    { *m = ListRoomsRequest{} }
func (m *ListRoomsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRoomsRequest) ProtoMessage()               {}
func (*ListRoomsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ListRoomsRequest) GetByType() RoomType {
	if m != nil {
		return m.ByType
	}
	return RoomType_GLOBAL
}

func (m *ListRoomsRequest) GetTypeObject() string {
	if m != nil {
		return m.TypeObject
	}
	return ""
}

type ListRoomsResponse struct {
	Room *ChatRoom `protobuf:"bytes,1,opt,name=Room" json:"Room,omitempty"`
}

func (m *ListRoomsResponse) Reset()                    { *m = ListRoomsResponse{} }
func (m *ListRoomsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListRoomsResponse) ProtoMessage()               {}
func (*ListRoomsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ListRoomsResponse) GetRoom() *ChatRoom {
	if m != nil {
		return m.Room
	}
	return nil
}

type DeleteRoomRequest struct {
	Room *ChatRoom `protobuf:"bytes,1,opt,name=Room" json:"Room,omitempty"`
}

func (m *DeleteRoomRequest) Reset()                    { *m = DeleteRoomRequest{} }
func (m *DeleteRoomRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRoomRequest) ProtoMessage()               {}
func (*DeleteRoomRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *DeleteRoomRequest) GetRoom() *ChatRoom {
	if m != nil {
		return m.Room
	}
	return nil
}

type DeleteRoomResponse struct {
	Success bool `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
}

func (m *DeleteRoomResponse) Reset()                    { *m = DeleteRoomResponse{} }
func (m *DeleteRoomResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteRoomResponse) ProtoMessage()               {}
func (*DeleteRoomResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *DeleteRoomResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ChatEvent struct {
	Message *ChatMessage `protobuf:"bytes,1,opt,name=Message" json:"Message,omitempty"`
	Room    *ChatRoom    `protobuf:"bytes,2,opt,name=Room" json:"Room,omitempty"`
	Details string       `protobuf:"bytes,3,opt,name=Details" json:"Details,omitempty"`
}

func (m *ChatEvent) Reset()                    { *m = ChatEvent{} }
func (m *ChatEvent) String() string            { return proto.CompactTextString(m) }
func (*ChatEvent) ProtoMessage()               {}
func (*ChatEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *ChatEvent) GetMessage() *ChatMessage {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *ChatEvent) GetRoom() *ChatRoom {
	if m != nil {
		return m.Room
	}
	return nil
}

func (m *ChatEvent) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

type WebSocketMessage struct {
	Type    WsMessageType `protobuf:"varint,1,opt,name=Type,json=@type,enum=chat.WsMessageType" json:"Type,omitempty"`
	Room    *ChatRoom     `protobuf:"bytes,2,opt,name=Room" json:"Room,omitempty"`
	Message *ChatMessage  `protobuf:"bytes,3,opt,name=Message" json:"Message,omitempty"`
}

func (m *WebSocketMessage) Reset()                    { *m = WebSocketMessage{} }
func (m *WebSocketMessage) String() string            { return proto.CompactTextString(m) }
func (*WebSocketMessage) ProtoMessage()               {}
func (*WebSocketMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *WebSocketMessage) GetType() WsMessageType {
	if m != nil {
		return m.Type
	}
	return WsMessageType_JOIN
}

func (m *WebSocketMessage) GetRoom() *ChatRoom {
	if m != nil {
		return m.Room
	}
	return nil
}

func (m *WebSocketMessage) GetMessage() *ChatMessage {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*ChatRoom)(nil), "chat.ChatRoom")
	proto.RegisterType((*ChatMessage)(nil), "chat.ChatMessage")
	proto.RegisterType((*PutRoomRequest)(nil), "chat.PutRoomRequest")
	proto.RegisterType((*PutRoomResponse)(nil), "chat.PutRoomResponse")
	proto.RegisterType((*PostMessageRequest)(nil), "chat.PostMessageRequest")
	proto.RegisterType((*PostMessageResponse)(nil), "chat.PostMessageResponse")
	proto.RegisterType((*DeleteMessageRequest)(nil), "chat.DeleteMessageRequest")
	proto.RegisterType((*DeleteMessageResponse)(nil), "chat.DeleteMessageResponse")
	proto.RegisterType((*ListMessagesRequest)(nil), "chat.ListMessagesRequest")
	proto.RegisterType((*ListMessagesResponse)(nil), "chat.ListMessagesResponse")
	proto.RegisterType((*ListRoomsRequest)(nil), "chat.ListRoomsRequest")
	proto.RegisterType((*ListRoomsResponse)(nil), "chat.ListRoomsResponse")
	proto.RegisterType((*DeleteRoomRequest)(nil), "chat.DeleteRoomRequest")
	proto.RegisterType((*DeleteRoomResponse)(nil), "chat.DeleteRoomResponse")
	proto.RegisterType((*ChatEvent)(nil), "chat.ChatEvent")
	proto.RegisterType((*WebSocketMessage)(nil), "chat.WebSocketMessage")
	proto.RegisterEnum("chat.RoomType", RoomType_name, RoomType_value)
	proto.RegisterEnum("chat.WsMessageType", WsMessageType_name, WsMessageType_value)
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 847 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0xae, 0xe3, 0x24, 0x9b, 0x1c, 0xb3, 0xa9, 0x77, 0x76, 0xdb, 0xba, 0x06, 0xa1, 0xc8, 0x17,
	0x55, 0xd4, 0x42, 0x02, 0xe1, 0xa7, 0xe2, 0x06, 0xc8, 0x26, 0xd6, 0xb6, 0xe0, 0xad, 0xa3, 0x71,
	0xc2, 0x0a, 0x2e, 0xa8, 0x1c, 0x67, 0xda, 0x18, 0xd6, 0xeb, 0x90, 0x99, 0x44, 0xca, 0x25, 0x0f,
	0xc0, 0x23, 0xf0, 0x24, 0xdc, 0xf0, 0x68, 0x68, 0x7e, 0xec, 0x38, 0x3f, 0xb0, 0x5d, 0x71, 0xe7,
	0xf3, 0xf7, 0xcd, 0x77, 0xce, 0x7c, 0x73, 0x12, 0x80, 0x68, 0x16, 0xb2, 0xf6, 0x7c, 0x91, 0xb2,
	0x14, 0x95, 0xf9, 0xb7, 0x3d, 0x78, 0x1b, 0xb3, 0xd9, 0x72, 0xd2, 0x8e, 0xd2, 0xa4, 0x33, 0x5f,
	0x4f, 0xe3, 0xb4, 0x43, 0xc9, 0x62, 0x15, 0x47, 0x84, 0x76, 0xa2, 0x34, 0x49, 0xd2, 0x9b, 0x8e,
	0xc8, 0xee, 0x84, 0x11, 0x8b, 0x57, 0x31, 0x5b, 0xe7, 0x1f, 0x94, 0x2d, 0x48, 0x98, 0x48, 0x2c,
	0xe7, 0x6f, 0x0d, 0x6a, 0xfd, 0x59, 0xc8, 0x70, 0x9a, 0x26, 0x08, 0x41, 0x79, 0xbc, 0x8c, 0xa7,
	0x96, 0xd6, 0xd4, 0x5a, 0x75, 0x2c, 0xbe, 0x91, 0x03, 0xe5, 0xd1, 0x7a, 0x4e, 0xac, 0x52, 0x53,
	0x6b, 0x35, 0xba, 0x8d, 0xb6, 0xe0, 0xc1, 0xb3, 0xb9, 0x17, 0x8b, 0x18, 0x7a, 0x02, 0x8d, 0xcc,
	0xe3, 0x4f, 0x7e, 0x21, 0x11, 0xb3, 0x74, 0x81, 0xb0, 0xe3, 0x45, 0x1f, 0x40, 0x9d, 0x7b, 0xbc,
	0x70, 0x42, 0xae, 0xad, 0xb2, 0x48, 0xd9, 0x38, 0xd0, 0x19, 0x54, 0xc6, 0x94, 0x2c, 0xa8, 0x55,
	0x69, 0xea, 0xad, 0x3a, 0x96, 0x06, 0x6a, 0x82, 0xe1, 0x85, 0x94, 0x8d, 0xe7, 0xd3, 0x90, 0x91,
	0xa9, 0x55, 0x6d, 0x6a, 0xad, 0x0a, 0x2e, 0xba, 0x9c, 0xbf, 0x34, 0x30, 0x78, 0x0b, 0x97, 0x84,
	0xd2, 0xf0, 0x2d, 0x39, 0xd8, 0x85, 0x0d, 0x35, 0x7e, 0x90, 0xf0, 0x97, 0x84, 0x3f, 0xb7, 0x91,
	0x05, 0x47, 0xaa, 0x54, 0xd1, 0xce, 0x4c, 0xf4, 0x10, 0xaa, 0xbd, 0x25, 0x9b, 0xa5, 0x0b, 0x45,
	0x56, 0x59, 0xbc, 0x8f, 0x51, 0x9c, 0x10, 0xca, 0xc2, 0x64, 0x6e, 0x55, 0x9a, 0x5a, 0x4b, 0xc7,
	0x1b, 0x07, 0xfa, 0x08, 0x6a, 0x3d, 0x35, 0x6a, 0x41, 0xd7, 0xe8, 0x9a, 0xed, 0x6c, 0xf6, 0x6d,
	0x39, 0x09, 0x9c, 0x67, 0x38, 0x9f, 0x43, 0x63, 0xb8, 0x14, 0xe3, 0xc7, 0xe4, 0xb7, 0x25, 0xa1,
	0x8c, 0x4f, 0x9c, 0x9b, 0x82, 0xbf, 0x91, 0x4d, 0x3c, 0xbb, 0x23, 0x2c, 0x62, 0xce, 0x17, 0x70,
	0x3f, 0xaf, 0xa2, 0xf3, 0xf4, 0x86, 0x92, 0x77, 0x2a, 0xeb, 0x03, 0x1a, 0xa6, 0x34, 0x9b, 0x54,
	0x76, 0xe0, 0xc7, 0x50, 0x53, 0x1e, 0x6a, 0x69, 0x4d, 0xbd, 0x65, 0x74, 0x4f, 0x36, 0xd5, 0x59,
	0x6e, 0x9e, 0xe2, 0xfc, 0x0c, 0xa7, 0x5b, 0x20, 0xea, 0x7c, 0x0b, 0x8e, 0x82, 0x65, 0x14, 0x11,
	0x4a, 0x05, 0x85, 0x1a, 0xce, 0xcc, 0x2d, 0xfc, 0xd2, 0xed, 0xf8, 0x2e, 0x9c, 0x0d, 0xc8, 0x35,
	0x61, 0xe4, 0xff, 0xd1, 0xfc, 0x14, 0x1e, 0xec, 0xc0, 0xdc, 0x46, 0xd4, 0xf9, 0x5d, 0x83, 0x53,
	0x2f, 0xce, 0x5b, 0xa3, 0xd9, 0xc9, 0x45, 0xf5, 0x68, 0x3b, 0xea, 0x51, 0xfa, 0xcc, 0x14, 0x24,
	0xc5, 0x55, 0x74, 0x71, 0x15, 0xf9, 0x6f, 0xde, 0x50, 0x22, 0x5f, 0x85, 0x8e, 0x95, 0xc5, 0xf5,
	0xee, 0xc5, 0x49, 0xcc, 0x84, 0xb8, 0x74, 0x2c, 0x0d, 0xa7, 0x0f, 0x67, 0xdb, 0x14, 0x14, 0xeb,
	0x67, 0x1b, 0x95, 0xca, 0x1b, 0x3e, 0xd0, 0x7c, 0x96, 0xe1, 0xfc, 0x04, 0x26, 0x07, 0xe1, 0x24,
	0xf3, 0x26, 0x9e, 0x40, 0xf5, 0x7c, 0x2d, 0x9e, 0xb2, 0x76, 0xf0, 0x29, 0xab, 0x28, 0xfa, 0x10,
	0xa0, 0xf0, 0x90, 0x65, 0x3f, 0x05, 0x8f, 0xf3, 0x1c, 0x4e, 0x0a, 0xd8, 0x77, 0x10, 0xdf, 0x73,
	0x38, 0x91, 0x17, 0x72, 0x57, 0xb1, 0xb7, 0x01, 0x15, 0x0b, 0x6f, 0xbd, 0xc6, 0x15, 0xd4, 0x39,
	0x82, 0xbb, 0x22, 0x37, 0xec, 0x4e, 0x73, 0xcb, 0xd9, 0x94, 0xfe, 0x9d, 0x0d, 0x3f, 0x77, 0x40,
	0x58, 0x18, 0x5f, 0xd3, 0x6c, 0x5d, 0x28, 0xd3, 0xf9, 0x43, 0x03, 0xf3, 0x8a, 0x4c, 0x82, 0x34,
	0xfa, 0x95, 0xe4, 0xb7, 0xdf, 0x52, 0xfb, 0x53, 0x0e, 0xfd, 0x54, 0x42, 0x5e, 0x51, 0x15, 0x16,
	0x93, 0xaf, 0x7c, 0xcb, 0xf8, 0xe0, 0xdf, 0xe5, 0xf0, 0x67, 0xdb, 0xbb, 0xea, 0x3f, 0xbb, 0x79,
	0xfa, 0x95, 0x94, 0xad, 0xb8, 0x55, 0x80, 0xea, 0x85, 0xe7, 0x9f, 0xf7, 0x3c, 0xf3, 0x1e, 0x3a,
	0x86, 0xfa, 0x95, 0x8f, 0xbf, 0x0f, 0x86, 0xbd, 0xbe, 0x6b, 0x6a, 0xa8, 0x06, 0xe5, 0x71, 0xe0,
	0x62, 0xb3, 0xc4, 0xbf, 0x5e, 0xf9, 0x03, 0xd7, 0xd4, 0x9f, 0x26, 0x70, 0xbc, 0xc5, 0x91, 0x87,
	0xbe, 0xf3, 0x5f, 0xbe, 0x32, 0xef, 0xa1, 0x3a, 0x54, 0x3c, 0xb7, 0xf7, 0x83, 0xaa, 0x1c, 0xfa,
	0xc1, 0xc8, 0x2c, 0xa1, 0xfb, 0x60, 0x60, 0xdf, 0xbf, 0x7c, 0x3d, 0x1e, 0x0e, 0x7a, 0x23, 0xd7,
	0xd4, 0x91, 0x01, 0x47, 0x2f, 0x5e, 0x06, 0x23, 0x1f, 0xff, 0x68, 0x96, 0x51, 0x03, 0x60, 0xe0,
	0x7a, 0xee, 0xc8, 0x7d, 0x7d, 0x19, 0x5c, 0x98, 0x15, 0x9e, 0xad, 0x6c, 0x5e, 0x64, 0x56, 0xbb,
	0x7f, 0xea, 0x72, 0x85, 0x07, 0xf2, 0x17, 0x0c, 0x7d, 0x09, 0x47, 0x6a, 0xbd, 0xa1, 0x33, 0xd9,
	0xe0, 0xf6, 0x8e, 0xb4, 0x1f, 0xec, 0x78, 0x95, 0x26, 0xbe, 0x01, 0xd8, 0x28, 0x05, 0x3d, 0x92,
	0x49, 0x7b, 0xa2, 0xb3, 0xad, 0xfd, 0x80, 0x02, 0xf8, 0x1a, 0xea, 0xb9, 0xb8, 0xd1, 0x43, 0x99,
	0xb6, 0xfb, 0x92, 0xec, 0x47, 0x7b, 0x7e, 0x59, 0xfd, 0x89, 0x86, 0x2e, 0xe0, 0xbd, 0xe2, 0xeb,
	0x45, 0x8f, 0x37, 0xa9, 0x3b, 0x4b, 0xc5, 0xb6, 0x0f, 0x85, 0x72, 0xa0, 0x73, 0x30, 0x0a, 0x4b,
	0x16, 0x29, 0xc6, 0xfb, 0xcb, 0xdb, 0x7e, 0x7c, 0x20, 0xa2, 0x9a, 0x79, 0x01, 0xc7, 0x5b, 0x1b,
	0x10, 0xd9, 0xc5, 0xbe, 0x77, 0x70, 0xde, 0x3f, 0x18, 0x93, 0x48, 0x93, 0xaa, 0xf8, 0xb3, 0xf0,
	0xd9, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x96, 0x6a, 0xe1, 0x85, 0x86, 0x08, 0x00, 0x00,
}
