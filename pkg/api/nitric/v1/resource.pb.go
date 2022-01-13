// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: resource/v1/resource.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type ResourceType int32

const (
	ResourceType_Api          ResourceType = 0
	ResourceType_Function     ResourceType = 1
	ResourceType_Bucket       ResourceType = 2
	ResourceType_Queue        ResourceType = 3
	ResourceType_Topic        ResourceType = 4
	ResourceType_Schedule     ResourceType = 5
	ResourceType_Subscription ResourceType = 6
	ResourceType_Collection   ResourceType = 7
	ResourceType_Policy       ResourceType = 8
)

// Enum value maps for ResourceType.
var (
	ResourceType_name = map[int32]string{
		0: "Api",
		1: "Function",
		2: "Bucket",
		3: "Queue",
		4: "Topic",
		5: "Schedule",
		6: "Subscription",
		7: "Collection",
		8: "Policy",
	}
	ResourceType_value = map[string]int32{
		"Api":          0,
		"Function":     1,
		"Bucket":       2,
		"Queue":        3,
		"Topic":        4,
		"Schedule":     5,
		"Subscription": 6,
		"Collection":   7,
		"Policy":       8,
	}
)

func (x ResourceType) Enum() *ResourceType {
	p := new(ResourceType)
	*p = x
	return p
}

func (x ResourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_resource_v1_resource_proto_enumTypes[0].Descriptor()
}

func (ResourceType) Type() protoreflect.EnumType {
	return &file_resource_v1_resource_proto_enumTypes[0]
}

func (x ResourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceType.Descriptor instead.
func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return file_resource_v1_resource_proto_rawDescGZIP(), []int{0}
}

type Action int32

const (
	// Bucket Permissions: 0XX
	Action_BucketFileList   Action = 0
	Action_BucketFileGet    Action = 1
	Action_BucketFilePut    Action = 2
	Action_BucketFileDelete Action = 3
	// Topic Permissions: 2XX
	Action_TopicList         Action = 200
	Action_TopicDetail       Action = 201
	Action_TopicEventPublish Action = 202
	// Queue Permissions: 3XX
	Action_QueueSend    Action = 300
	Action_QueueReceive Action = 301
	Action_QueueList    Action = 302
	Action_QueueDetail  Action = 303
	// Collection Permissions: 4XX
	Action_CollectionDocumentRead   Action = 400
	Action_CollectionDocumentWrite  Action = 401
	Action_CollectionDocumentDelete Action = 402
	Action_CollectionQuery          Action = 403
	Action_CollectionList           Action = 404
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0:   "BucketFileList",
		1:   "BucketFileGet",
		2:   "BucketFilePut",
		3:   "BucketFileDelete",
		200: "TopicList",
		201: "TopicDetail",
		202: "TopicEventPublish",
		300: "QueueSend",
		301: "QueueReceive",
		302: "QueueList",
		303: "QueueDetail",
		400: "CollectionDocumentRead",
		401: "CollectionDocumentWrite",
		402: "CollectionDocumentDelete",
		403: "CollectionQuery",
		404: "CollectionList",
	}
	Action_value = map[string]int32{
		"BucketFileList":           0,
		"BucketFileGet":            1,
		"BucketFilePut":            2,
		"BucketFileDelete":         3,
		"TopicList":                200,
		"TopicDetail":              201,
		"TopicEventPublish":        202,
		"QueueSend":                300,
		"QueueReceive":             301,
		"QueueList":                302,
		"QueueDetail":              303,
		"CollectionDocumentRead":   400,
		"CollectionDocumentWrite":  401,
		"CollectionDocumentDelete": 402,
		"CollectionQuery":          403,
		"CollectionList":           404,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_resource_v1_resource_proto_enumTypes[1].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_resource_v1_resource_proto_enumTypes[1]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_resource_v1_resource_proto_rawDescGZIP(), []int{1}
}

type PolicyResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Principals []*Resource `protobuf:"bytes,1,rep,name=principals,proto3" json:"principals,omitempty"`
	Actions    []Action    `protobuf:"varint,2,rep,packed,name=actions,proto3,enum=nitric.resource.v1.Action" json:"actions,omitempty"`
	Resources  []*Resource `protobuf:"bytes,3,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *PolicyResource) Reset() {
	*x = PolicyResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_v1_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyResource) ProtoMessage() {}

func (x *PolicyResource) ProtoReflect() protoreflect.Message {
	mi := &file_resource_v1_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyResource.ProtoReflect.Descriptor instead.
func (*PolicyResource) Descriptor() ([]byte, []int) {
	return file_resource_v1_resource_proto_rawDescGZIP(), []int{0}
}

func (x *PolicyResource) GetPrincipals() []*Resource {
	if x != nil {
		return x.Principals
	}
	return nil
}

func (x *PolicyResource) GetActions() []Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *PolicyResource) GetResources() []*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ResourceType `protobuf:"varint,1,opt,name=type,proto3,enum=nitric.resource.v1.ResourceType" json:"type,omitempty"`
	Name string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_v1_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_resource_v1_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_resource_v1_resource_proto_rawDescGZIP(), []int{1}
}

func (x *Resource) GetType() ResourceType {
	if x != nil {
		return x.Type
	}
	return ResourceType_Api
}

func (x *Resource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ResourceDeclareRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// Types that are assignable to Config:
	//	*ResourceDeclareRequest_Policy
	//	*ResourceDeclareRequest_Bucket
	//	*ResourceDeclareRequest_Queue
	//	*ResourceDeclareRequest_Topic
	//	*ResourceDeclareRequest_Collection
	Config isResourceDeclareRequest_Config `protobuf_oneof:"config"`
}

func (x *ResourceDeclareRequest) Reset() {
	*x = ResourceDeclareRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_v1_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceDeclareRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceDeclareRequest) ProtoMessage() {}

func (x *ResourceDeclareRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resource_v1_resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceDeclareRequest.ProtoReflect.Descriptor instead.
func (*ResourceDeclareRequest) Descriptor() ([]byte, []int) {
	return file_resource_v1_resource_proto_rawDescGZIP(), []int{2}
}

func (x *ResourceDeclareRequest) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (m *ResourceDeclareRequest) GetConfig() isResourceDeclareRequest_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (x *ResourceDeclareRequest) GetPolicy() *PolicyResource {
	if x, ok := x.GetConfig().(*ResourceDeclareRequest_Policy); ok {
		return x.Policy
	}
	return nil
}

func (x *ResourceDeclareRequest) GetBucket() *BucketResource {
	if x, ok := x.GetConfig().(*ResourceDeclareRequest_Bucket); ok {
		return x.Bucket
	}
	return nil
}

func (x *ResourceDeclareRequest) GetQueue() *QueueResource {
	if x, ok := x.GetConfig().(*ResourceDeclareRequest_Queue); ok {
		return x.Queue
	}
	return nil
}

func (x *ResourceDeclareRequest) GetTopic() *TopicResource {
	if x, ok := x.GetConfig().(*ResourceDeclareRequest_Topic); ok {
		return x.Topic
	}
	return nil
}

func (x *ResourceDeclareRequest) GetCollection() *CollectionResource {
	if x, ok := x.GetConfig().(*ResourceDeclareRequest_Collection); ok {
		return x.Collection
	}
	return nil
}

type isResourceDeclareRequest_Config interface {
	isResourceDeclareRequest_Config()
}

type ResourceDeclareRequest_Policy struct {
	Policy *PolicyResource `protobuf:"bytes,10,opt,name=policy,proto3,oneof"`
}

type ResourceDeclareRequest_Bucket struct {
	Bucket *BucketResource `protobuf:"bytes,11,opt,name=bucket,proto3,oneof"`
}

type ResourceDeclareRequest_Queue struct {
	Queue *QueueResource `protobuf:"bytes,12,opt,name=queue,proto3,oneof"`
}

type ResourceDeclareRequest_Topic struct {
	Topic *TopicResource `protobuf:"bytes,13,opt,name=topic,proto3,oneof"`
}

type ResourceDeclareRequest_Collection struct {
	Collection *CollectionResource `protobuf:"bytes,14,opt,name=collection,proto3,oneof"`
}

func (*ResourceDeclareRequest_Policy) isResourceDeclareRequest_Config() {}

func (*ResourceDeclareRequest_Bucket) isResourceDeclareRequest_Config() {}

func (*ResourceDeclareRequest_Queue) isResourceDeclareRequest_Config() {}

func (*ResourceDeclareRequest_Topic) isResourceDeclareRequest_Config() {}

func (*ResourceDeclareRequest_Collection) isResourceDeclareRequest_Config() {}

type BucketResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BucketResource) Reset() {
	*x = BucketResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_v1_resource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BucketResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BucketResource) ProtoMessage() {}

func (x *BucketResource) ProtoReflect() protoreflect.Message {
	mi := &file_resource_v1_resource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BucketResource.ProtoReflect.Descriptor instead.
func (*BucketResource) Descriptor() ([]byte, []int) {
	return file_resource_v1_resource_proto_rawDescGZIP(), []int{3}
}

type QueueResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueueResource) Reset() {
	*x = QueueResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_v1_resource_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueResource) ProtoMessage() {}

func (x *QueueResource) ProtoReflect() protoreflect.Message {
	mi := &file_resource_v1_resource_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueResource.ProtoReflect.Descriptor instead.
func (*QueueResource) Descriptor() ([]byte, []int) {
	return file_resource_v1_resource_proto_rawDescGZIP(), []int{4}
}

type TopicResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TopicResource) Reset() {
	*x = TopicResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_v1_resource_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicResource) ProtoMessage() {}

func (x *TopicResource) ProtoReflect() protoreflect.Message {
	mi := &file_resource_v1_resource_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicResource.ProtoReflect.Descriptor instead.
func (*TopicResource) Descriptor() ([]byte, []int) {
	return file_resource_v1_resource_proto_rawDescGZIP(), []int{5}
}

type CollectionResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CollectionResource) Reset() {
	*x = CollectionResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_v1_resource_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionResource) ProtoMessage() {}

func (x *CollectionResource) ProtoReflect() protoreflect.Message {
	mi := &file_resource_v1_resource_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionResource.ProtoReflect.Descriptor instead.
func (*CollectionResource) Descriptor() ([]byte, []int) {
	return file_resource_v1_resource_proto_rawDescGZIP(), []int{6}
}

type ResourceDeclareResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResourceDeclareResponse) Reset() {
	*x = ResourceDeclareResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_v1_resource_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceDeclareResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceDeclareResponse) ProtoMessage() {}

func (x *ResourceDeclareResponse) ProtoReflect() protoreflect.Message {
	mi := &file_resource_v1_resource_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceDeclareResponse.ProtoReflect.Descriptor instead.
func (*ResourceDeclareResponse) Descriptor() ([]byte, []int) {
	return file_resource_v1_resource_proto_rawDescGZIP(), []int{7}
}

var File_resource_v1_resource_proto protoreflect.FileDescriptor

var file_resource_v1_resource_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6e, 0x69,
	0x74, 0x72, 0x69, 0x63, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a, 0x0e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0a,
	0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0a,
	0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6e, 0x69,
	0x74, 0x72, 0x69, 0x63, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x3a, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x54, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x98, 0x03, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44,
	0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x06, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x3c, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x06, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x39, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x39,
	0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x48, 0x00, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x48, 0x0a, 0x0a, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x10, 0x0a,
	0x0e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22,
	0x0f, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x22, 0x0f, 0x0a, 0x0d, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x19, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x44, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2a, 0x83, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x70, 0x69, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x75, 0x65, 0x10,
	0x03, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x10, 0x08, 0x2a, 0xdc, 0x02, 0x0a, 0x06, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x75, 0x74, 0x10, 0x02, 0x12, 0x14, 0x0a,
	0x10, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x09, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74,
	0x10, 0xc8, 0x01, 0x12, 0x10, 0x0a, 0x0b, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x10, 0xc9, 0x01, 0x12, 0x16, 0x0a, 0x11, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x10, 0xca, 0x01, 0x12, 0x0e, 0x0a,
	0x09, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x10, 0xac, 0x02, 0x12, 0x11, 0x0a,
	0x0c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x10, 0xad, 0x02,
	0x12, 0x0e, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x10, 0xae, 0x02,
	0x12, 0x10, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x10,
	0xaf, 0x02, 0x12, 0x1b, 0x0a, 0x16, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x61, 0x64, 0x10, 0x90, 0x03, 0x12,
	0x1c, 0x0a, 0x17, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x10, 0x91, 0x03, 0x12, 0x1d, 0x0a,
	0x18, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x92, 0x03, 0x12, 0x14, 0x0a, 0x0f,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x10,
	0x93, 0x03, 0x12, 0x13, 0x0a, 0x0e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x10, 0x94, 0x03, 0x32, 0x75, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x07, 0x44, 0x65,
	0x63, 0x6c, 0x61, 0x72, 0x65, 0x12, 0x2a, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x44, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2b, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44,
	0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6e,
	0x0a, 0x1b, 0x69, 0x6f, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x0c, 0x6e, 0x69, 0x74, 0x72,
	0x69, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0xaa, 0x02, 0x18, 0x4e, 0x69, 0x74, 0x72, 0x69,
	0x63, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0xca, 0x02, 0x18, 0x4e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x5c, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resource_v1_resource_proto_rawDescOnce sync.Once
	file_resource_v1_resource_proto_rawDescData = file_resource_v1_resource_proto_rawDesc
)

func file_resource_v1_resource_proto_rawDescGZIP() []byte {
	file_resource_v1_resource_proto_rawDescOnce.Do(func() {
		file_resource_v1_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_resource_v1_resource_proto_rawDescData)
	})
	return file_resource_v1_resource_proto_rawDescData
}

var file_resource_v1_resource_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_resource_v1_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_resource_v1_resource_proto_goTypes = []interface{}{
	(ResourceType)(0),               // 0: nitric.resource.v1.ResourceType
	(Action)(0),                     // 1: nitric.resource.v1.Action
	(*PolicyResource)(nil),          // 2: nitric.resource.v1.PolicyResource
	(*Resource)(nil),                // 3: nitric.resource.v1.Resource
	(*ResourceDeclareRequest)(nil),  // 4: nitric.resource.v1.ResourceDeclareRequest
	(*BucketResource)(nil),          // 5: nitric.resource.v1.BucketResource
	(*QueueResource)(nil),           // 6: nitric.resource.v1.QueueResource
	(*TopicResource)(nil),           // 7: nitric.resource.v1.TopicResource
	(*CollectionResource)(nil),      // 8: nitric.resource.v1.CollectionResource
	(*ResourceDeclareResponse)(nil), // 9: nitric.resource.v1.ResourceDeclareResponse
}
var file_resource_v1_resource_proto_depIdxs = []int32{
	3,  // 0: nitric.resource.v1.PolicyResource.principals:type_name -> nitric.resource.v1.Resource
	1,  // 1: nitric.resource.v1.PolicyResource.actions:type_name -> nitric.resource.v1.Action
	3,  // 2: nitric.resource.v1.PolicyResource.resources:type_name -> nitric.resource.v1.Resource
	0,  // 3: nitric.resource.v1.Resource.type:type_name -> nitric.resource.v1.ResourceType
	3,  // 4: nitric.resource.v1.ResourceDeclareRequest.resource:type_name -> nitric.resource.v1.Resource
	2,  // 5: nitric.resource.v1.ResourceDeclareRequest.policy:type_name -> nitric.resource.v1.PolicyResource
	5,  // 6: nitric.resource.v1.ResourceDeclareRequest.bucket:type_name -> nitric.resource.v1.BucketResource
	6,  // 7: nitric.resource.v1.ResourceDeclareRequest.queue:type_name -> nitric.resource.v1.QueueResource
	7,  // 8: nitric.resource.v1.ResourceDeclareRequest.topic:type_name -> nitric.resource.v1.TopicResource
	8,  // 9: nitric.resource.v1.ResourceDeclareRequest.collection:type_name -> nitric.resource.v1.CollectionResource
	4,  // 10: nitric.resource.v1.ResourceService.Declare:input_type -> nitric.resource.v1.ResourceDeclareRequest
	9,  // 11: nitric.resource.v1.ResourceService.Declare:output_type -> nitric.resource.v1.ResourceDeclareResponse
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_resource_v1_resource_proto_init() }
func file_resource_v1_resource_proto_init() {
	if File_resource_v1_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resource_v1_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyResource); i {
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
		file_resource_v1_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_resource_v1_resource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceDeclareRequest); i {
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
		file_resource_v1_resource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BucketResource); i {
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
		file_resource_v1_resource_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueResource); i {
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
		file_resource_v1_resource_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicResource); i {
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
		file_resource_v1_resource_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionResource); i {
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
		file_resource_v1_resource_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceDeclareResponse); i {
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
	file_resource_v1_resource_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ResourceDeclareRequest_Policy)(nil),
		(*ResourceDeclareRequest_Bucket)(nil),
		(*ResourceDeclareRequest_Queue)(nil),
		(*ResourceDeclareRequest_Topic)(nil),
		(*ResourceDeclareRequest_Collection)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resource_v1_resource_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_resource_v1_resource_proto_goTypes,
		DependencyIndexes: file_resource_v1_resource_proto_depIdxs,
		EnumInfos:         file_resource_v1_resource_proto_enumTypes,
		MessageInfos:      file_resource_v1_resource_proto_msgTypes,
	}.Build()
	File_resource_v1_resource_proto = out.File
	file_resource_v1_resource_proto_rawDesc = nil
	file_resource_v1_resource_proto_goTypes = nil
	file_resource_v1_resource_proto_depIdxs = nil
}
