// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.5
// source: grpc-service/notifications-service.proto

package grpc_service

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type OrderDirection int32

const (
	OrderDirection_ASC  OrderDirection = 0
	OrderDirection_DESC OrderDirection = 1
)

// Enum value maps for OrderDirection.
var (
	OrderDirection_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	OrderDirection_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x OrderDirection) Enum() *OrderDirection {
	p := new(OrderDirection)
	*p = x
	return p
}

func (x OrderDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_service_notifications_service_proto_enumTypes[0].Descriptor()
}

func (OrderDirection) Type() protoreflect.EnumType {
	return &file_grpc_service_notifications_service_proto_enumTypes[0]
}

func (x OrderDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderDirection.Descriptor instead.
func (OrderDirection) EnumDescriptor() ([]byte, []int) {
	return file_grpc_service_notifications_service_proto_rawDescGZIP(), []int{0}
}

type SearchTemplatesRequest_OrderBy int32

const (
	SearchTemplatesRequest_NAME SearchTemplatesRequest_OrderBy = 0
)

// Enum value maps for SearchTemplatesRequest_OrderBy.
var (
	SearchTemplatesRequest_OrderBy_name = map[int32]string{
		0: "NAME",
	}
	SearchTemplatesRequest_OrderBy_value = map[string]int32{
		"NAME": 0,
	}
)

func (x SearchTemplatesRequest_OrderBy) Enum() *SearchTemplatesRequest_OrderBy {
	p := new(SearchTemplatesRequest_OrderBy)
	*p = x
	return p
}

func (x SearchTemplatesRequest_OrderBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SearchTemplatesRequest_OrderBy) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_service_notifications_service_proto_enumTypes[1].Descriptor()
}

func (SearchTemplatesRequest_OrderBy) Type() protoreflect.EnumType {
	return &file_grpc_service_notifications_service_proto_enumTypes[1]
}

func (x SearchTemplatesRequest_OrderBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SearchTemplatesRequest_OrderBy.Descriptor instead.
func (SearchTemplatesRequest_OrderBy) EnumDescriptor() ([]byte, []int) {
	return file_grpc_service_notifications_service_proto_rawDescGZIP(), []int{4, 0}
}

type SearchNotificationsRequest_OrderBy int32

const (
	SearchNotificationsRequest_CREATED_AT SearchNotificationsRequest_OrderBy = 0
)

// Enum value maps for SearchNotificationsRequest_OrderBy.
var (
	SearchNotificationsRequest_OrderBy_name = map[int32]string{
		0: "CREATED_AT",
	}
	SearchNotificationsRequest_OrderBy_value = map[string]int32{
		"CREATED_AT": 0,
	}
)

func (x SearchNotificationsRequest_OrderBy) Enum() *SearchNotificationsRequest_OrderBy {
	p := new(SearchNotificationsRequest_OrderBy)
	*p = x
	return p
}

func (x SearchNotificationsRequest_OrderBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SearchNotificationsRequest_OrderBy) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_service_notifications_service_proto_enumTypes[2].Descriptor()
}

func (SearchNotificationsRequest_OrderBy) Type() protoreflect.EnumType {
	return &file_grpc_service_notifications_service_proto_enumTypes[2]
}

func (x SearchNotificationsRequest_OrderBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SearchNotificationsRequest_OrderBy.Descriptor instead.
func (SearchNotificationsRequest_OrderBy) EnumDescriptor() ([]byte, []int) {
	return file_grpc_service_notifications_service_proto_rawDescGZIP(), []int{7, 0}
}

type Notification_Type int32

const (
	Notification_EMAIL Notification_Type = 0
)

// Enum value maps for Notification_Type.
var (
	Notification_Type_name = map[int32]string{
		0: "EMAIL",
	}
	Notification_Type_value = map[string]int32{
		"EMAIL": 0,
	}
)

func (x Notification_Type) Enum() *Notification_Type {
	p := new(Notification_Type)
	*p = x
	return p
}

func (x Notification_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Notification_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_service_notifications_service_proto_enumTypes[3].Descriptor()
}

func (Notification_Type) Type() protoreflect.EnumType {
	return &file_grpc_service_notifications_service_proto_enumTypes[3]
}

func (x Notification_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Notification_Type.Descriptor instead.
func (Notification_Type) EnumDescriptor() ([]byte, []int) {
	return file_grpc_service_notifications_service_proto_rawDescGZIP(), []int{9, 0}
}

type SendEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email      string            `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	TemplateId string            `protobuf:"bytes,2,opt,name=templateId,proto3" json:"templateId,omitempty"`
	Context    map[string]string `protobuf:"bytes,3,rep,name=context,proto3" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SendEmailRequest) Reset() {
	*x = SendEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_service_notifications_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailRequest) ProtoMessage() {}

func (x *SendEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_notifications_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailRequest.ProtoReflect.Descriptor instead.
func (*SendEmailRequest) Descriptor() ([]byte, []int) {
	return file_grpc_service_notifications_service_proto_rawDescGZIP(), []int{0}
}

func (x *SendEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SendEmailRequest) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *SendEmailRequest) GetContext() map[string]string {
	if x != nil {
		return x.Context
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_service_notifications_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_notifications_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_grpc_service_notifications_service_proto_rawDescGZIP(), []int{1}
}

type CreateTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Template string `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *CreateTemplateRequest) Reset() {
	*x = CreateTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_service_notifications_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTemplateRequest) ProtoMessage() {}

func (x *CreateTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_notifications_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTemplateRequest.ProtoReflect.Descriptor instead.
func (*CreateTemplateRequest) Descriptor() ([]byte, []int) {
	return file_grpc_service_notifications_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTemplateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTemplateRequest) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

type CreateTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateTemplateResponse) Reset() {
	*x = CreateTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_service_notifications_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTemplateResponse) ProtoMessage() {}

func (x *CreateTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_notifications_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTemplateResponse.ProtoReflect.Descriptor instead.
func (*CreateTemplateResponse) Descriptor() ([]byte, []int) {
	return file_grpc_service_notifications_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTemplateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SearchTemplatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string                         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Page           int32                          `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size           int32                          `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	OrderBy        SearchTemplatesRequest_OrderBy `protobuf:"varint,4,opt,name=orderBy,proto3,enum=SearchTemplatesRequest_OrderBy" json:"orderBy,omitempty"`
	OrderDirection OrderDirection                 `protobuf:"varint,5,opt,name=orderDirection,proto3,enum=OrderDirection" json:"orderDirection,omitempty"`
}

func (x *SearchTemplatesRequest) Reset() {
	*x = SearchTemplatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_service_notifications_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchTemplatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchTemplatesRequest) ProtoMessage() {}

func (x *SearchTemplatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_notifications_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchTemplatesRequest.ProtoReflect.Descriptor instead.
func (*SearchTemplatesRequest) Descriptor() ([]byte, []int) {
	return file_grpc_service_notifications_service_proto_rawDescGZIP(), []int{4}
}

func (x *SearchTemplatesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SearchTemplatesRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchTemplatesRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SearchTemplatesRequest) GetOrderBy() SearchTemplatesRequest_OrderBy {
	if x != nil {
		return x.OrderBy
	}
	return SearchTemplatesRequest_NAME
}

func (x *SearchTemplatesRequest) GetOrderDirection() OrderDirection {
	if x != nil {
		return x.OrderDirection
	}
	return OrderDirection_ASC
}

type SearchTemplatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*Template `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Count int32       `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *SearchTemplatesResponse) Reset() {
	*x = SearchTemplatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_service_notifications_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchTemplatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchTemplatesResponse) ProtoMessage() {}

func (x *SearchTemplatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_notifications_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchTemplatesResponse.ProtoReflect.Descriptor instead.
func (*SearchTemplatesResponse) Descriptor() ([]byte, []int) {
	return file_grpc_service_notifications_service_proto_rawDescGZIP(), []int{5}
}

func (x *SearchTemplatesResponse) GetList() []*Template {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *SearchTemplatesResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Template struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Template string `protobuf:"bytes,3,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *Template) Reset() {
	*x = Template{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_service_notifications_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Template) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Template) ProtoMessage() {}

func (x *Template) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_notifications_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Template.ProtoReflect.Descriptor instead.
func (*Template) Descriptor() ([]byte, []int) {
	return file_grpc_service_notifications_service_proto_rawDescGZIP(), []int{6}
}

func (x *Template) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Template) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Template) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

type SearchNotificationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page           int32                              `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size           int32                              `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	OrderBy        SearchNotificationsRequest_OrderBy `protobuf:"varint,3,opt,name=orderBy,proto3,enum=SearchNotificationsRequest_OrderBy" json:"orderBy,omitempty"`
	OrderDirection OrderDirection                     `protobuf:"varint,4,opt,name=orderDirection,proto3,enum=OrderDirection" json:"orderDirection,omitempty"`
}

func (x *SearchNotificationsRequest) Reset() {
	*x = SearchNotificationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_service_notifications_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchNotificationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchNotificationsRequest) ProtoMessage() {}

func (x *SearchNotificationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_notifications_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchNotificationsRequest.ProtoReflect.Descriptor instead.
func (*SearchNotificationsRequest) Descriptor() ([]byte, []int) {
	return file_grpc_service_notifications_service_proto_rawDescGZIP(), []int{7}
}

func (x *SearchNotificationsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchNotificationsRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SearchNotificationsRequest) GetOrderBy() SearchNotificationsRequest_OrderBy {
	if x != nil {
		return x.OrderBy
	}
	return SearchNotificationsRequest_CREATED_AT
}

func (x *SearchNotificationsRequest) GetOrderDirection() OrderDirection {
	if x != nil {
		return x.OrderDirection
	}
	return OrderDirection_ASC
}

type SearchNotificationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*Notification `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Count int32           `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *SearchNotificationsResponse) Reset() {
	*x = SearchNotificationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_service_notifications_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchNotificationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchNotificationsResponse) ProtoMessage() {}

func (x *SearchNotificationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_notifications_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchNotificationsResponse.ProtoReflect.Descriptor instead.
func (*SearchNotificationsResponse) Descriptor() ([]byte, []int) {
	return file_grpc_service_notifications_service_proto_rawDescGZIP(), []int{8}
}

func (x *SearchNotificationsResponse) GetList() []*Notification {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *SearchNotificationsResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Addressee string            `protobuf:"bytes,2,opt,name=addressee,proto3" json:"addressee,omitempty"`
	Message   string            `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Type      Notification_Type `protobuf:"varint,4,opt,name=type,proto3,enum=Notification_Type" json:"type,omitempty"`
	CreatedAt string            `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_service_notifications_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_notifications_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_grpc_service_notifications_service_proto_rawDescGZIP(), []int{9}
}

func (x *Notification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Notification) GetAddressee() string {
	if x != nil {
		return x.Addressee
	}
	return ""
}

func (x *Notification) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Notification) GetType() Notification_Type {
	if x != nil {
		return x.Type
	}
	return Notification_EMAIL
}

func (x *Notification) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

var File_grpc_service_notifications_service_proto protoreflect.FileDescriptor

var file_grpc_service_notifications_service_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x10, 0x53,
	0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x1a,
	0x3a, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x07, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x47, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x28, 0x0a,
	0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xdd, 0x01, 0x0a, 0x16, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x39,
	0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1f, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x37, 0x0a, 0x0e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x13, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x00, 0x22, 0x4e, 0x0a, 0x17, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x22, 0xd7, 0x01, 0x0a, 0x1a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x37, 0x0a, 0x0e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x19, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x0e, 0x0a,
	0x0a, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x10, 0x00, 0x22, 0x56, 0x0a,
	0x1b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xaf, 0x01, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x65, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x26,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x11, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05,
	0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x00, 0x2a, 0x23, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x01, 0x32, 0xa1, 0x02, 0x0a,
	0x14, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x11, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x43, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x17, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x13,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x61, 0x79, 0x65, 0x72, 0x6b, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_service_notifications_service_proto_rawDescOnce sync.Once
	file_grpc_service_notifications_service_proto_rawDescData = file_grpc_service_notifications_service_proto_rawDesc
)

func file_grpc_service_notifications_service_proto_rawDescGZIP() []byte {
	file_grpc_service_notifications_service_proto_rawDescOnce.Do(func() {
		file_grpc_service_notifications_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_service_notifications_service_proto_rawDescData)
	})
	return file_grpc_service_notifications_service_proto_rawDescData
}

var file_grpc_service_notifications_service_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_grpc_service_notifications_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_grpc_service_notifications_service_proto_goTypes = []interface{}{
	(OrderDirection)(0),                     // 0: OrderDirection
	(SearchTemplatesRequest_OrderBy)(0),     // 1: SearchTemplatesRequest.OrderBy
	(SearchNotificationsRequest_OrderBy)(0), // 2: SearchNotificationsRequest.OrderBy
	(Notification_Type)(0),                  // 3: Notification.Type
	(*SendEmailRequest)(nil),                // 4: SendEmailRequest
	(*Empty)(nil),                           // 5: Empty
	(*CreateTemplateRequest)(nil),           // 6: CreateTemplateRequest
	(*CreateTemplateResponse)(nil),          // 7: CreateTemplateResponse
	(*SearchTemplatesRequest)(nil),          // 8: SearchTemplatesRequest
	(*SearchTemplatesResponse)(nil),         // 9: SearchTemplatesResponse
	(*Template)(nil),                        // 10: Template
	(*SearchNotificationsRequest)(nil),      // 11: SearchNotificationsRequest
	(*SearchNotificationsResponse)(nil),     // 12: SearchNotificationsResponse
	(*Notification)(nil),                    // 13: Notification
	nil,                                     // 14: SendEmailRequest.ContextEntry
}
var file_grpc_service_notifications_service_proto_depIdxs = []int32{
	14, // 0: SendEmailRequest.context:type_name -> SendEmailRequest.ContextEntry
	1,  // 1: SearchTemplatesRequest.orderBy:type_name -> SearchTemplatesRequest.OrderBy
	0,  // 2: SearchTemplatesRequest.orderDirection:type_name -> OrderDirection
	10, // 3: SearchTemplatesResponse.list:type_name -> Template
	2,  // 4: SearchNotificationsRequest.orderBy:type_name -> SearchNotificationsRequest.OrderBy
	0,  // 5: SearchNotificationsRequest.orderDirection:type_name -> OrderDirection
	13, // 6: SearchNotificationsResponse.list:type_name -> Notification
	3,  // 7: Notification.type:type_name -> Notification.Type
	4,  // 8: NotificationsService.SendEmail:input_type -> SendEmailRequest
	6,  // 9: NotificationsService.CreateTemplate:input_type -> CreateTemplateRequest
	8,  // 10: NotificationsService.SearchTemplates:input_type -> SearchTemplatesRequest
	11, // 11: NotificationsService.SearchNotifications:input_type -> SearchNotificationsRequest
	5,  // 12: NotificationsService.SendEmail:output_type -> Empty
	7,  // 13: NotificationsService.CreateTemplate:output_type -> CreateTemplateResponse
	9,  // 14: NotificationsService.SearchTemplates:output_type -> SearchTemplatesResponse
	12, // 15: NotificationsService.SearchNotifications:output_type -> SearchNotificationsResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_grpc_service_notifications_service_proto_init() }
func file_grpc_service_notifications_service_proto_init() {
	if File_grpc_service_notifications_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_service_notifications_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailRequest); i {
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
		file_grpc_service_notifications_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_grpc_service_notifications_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTemplateRequest); i {
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
		file_grpc_service_notifications_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTemplateResponse); i {
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
		file_grpc_service_notifications_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchTemplatesRequest); i {
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
		file_grpc_service_notifications_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchTemplatesResponse); i {
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
		file_grpc_service_notifications_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Template); i {
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
		file_grpc_service_notifications_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchNotificationsRequest); i {
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
		file_grpc_service_notifications_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchNotificationsResponse); i {
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
		file_grpc_service_notifications_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
			RawDescriptor: file_grpc_service_notifications_service_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_service_notifications_service_proto_goTypes,
		DependencyIndexes: file_grpc_service_notifications_service_proto_depIdxs,
		EnumInfos:         file_grpc_service_notifications_service_proto_enumTypes,
		MessageInfos:      file_grpc_service_notifications_service_proto_msgTypes,
	}.Build()
	File_grpc_service_notifications_service_proto = out.File
	file_grpc_service_notifications_service_proto_rawDesc = nil
	file_grpc_service_notifications_service_proto_goTypes = nil
	file_grpc_service_notifications_service_proto_depIdxs = nil
}
