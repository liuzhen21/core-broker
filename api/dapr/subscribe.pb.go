// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: api/dapr/subscribe.proto

package dapr

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ListTopicSubscriptionsResponse is the message including the list of the subscribing topics.
type ListTopicSubscriptionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of topics.
	Subscriptions []*TopicSubscription `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
}

func (x *ListTopicSubscriptionsResponse) Reset() {
	*x = ListTopicSubscriptionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dapr_subscribe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTopicSubscriptionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTopicSubscriptionsResponse) ProtoMessage() {}

func (x *ListTopicSubscriptionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_dapr_subscribe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTopicSubscriptionsResponse.ProtoReflect.Descriptor instead.
func (*ListTopicSubscriptionsResponse) Descriptor() ([]byte, []int) {
	return file_api_dapr_subscribe_proto_rawDescGZIP(), []int{0}
}

func (x *ListTopicSubscriptionsResponse) GetSubscriptions() []*TopicSubscription {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

// TopicSubscription represents topic and metadata.
type TopicSubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the pubsub containing the topic below to subscribe to.
	PubsubName string `protobuf:"bytes,1,opt,name=pubsub_name,json=pubsubName,proto3" json:"pubsub_name,omitempty"`
	// Required. The name of topic which will be subscribed
	Topic string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	// The optional properties used for this topic's subscription e.g. session id
	Metadata map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The optional routing rules to match against. In the gRPC interface, OnTopicEvent
	// is still invoked but the matching path is sent in the TopicEventRequest.
	Route string `protobuf:"bytes,5,opt,name=route,proto3" json:"route,omitempty"`
}

func (x *TopicSubscription) Reset() {
	*x = TopicSubscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dapr_subscribe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicSubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicSubscription) ProtoMessage() {}

func (x *TopicSubscription) ProtoReflect() protoreflect.Message {
	mi := &file_api_dapr_subscribe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicSubscription.ProtoReflect.Descriptor instead.
func (*TopicSubscription) Descriptor() ([]byte, []int) {
	return file_api_dapr_subscribe_proto_rawDescGZIP(), []int{1}
}

func (x *TopicSubscription) GetPubsubName() string {
	if x != nil {
		return x.PubsubName
	}
	return ""
}

func (x *TopicSubscription) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *TopicSubscription) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *TopicSubscription) GetRoute() string {
	if x != nil {
		return x.Route
	}
	return ""
}

var File_api_dapr_subscribe_proto protoreflect.FileDescriptor

var file_api_dapr_subscribe_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x70, 0x72, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x70, 0x69, 0x2e,
	0x64, 0x61, 0x70, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x63, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x41, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x61, 0x70, 0x72, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0xe4, 0x01, 0x0a, 0x11, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x75,
	0x62, 0x73, 0x75, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x12, 0x45, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x1a, 0x3b,
	0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x71, 0x0a, 0x09, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x64, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x3b,
	0x0a, 0x08, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6b, 0x65, 0x65, 0x6c, 0x2d, 0x69,
	0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x64, 0x61, 0x70, 0x72, 0x3b, 0x64, 0x61, 0x70, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_dapr_subscribe_proto_rawDescOnce sync.Once
	file_api_dapr_subscribe_proto_rawDescData = file_api_dapr_subscribe_proto_rawDesc
)

func file_api_dapr_subscribe_proto_rawDescGZIP() []byte {
	file_api_dapr_subscribe_proto_rawDescOnce.Do(func() {
		file_api_dapr_subscribe_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_dapr_subscribe_proto_rawDescData)
	})
	return file_api_dapr_subscribe_proto_rawDescData
}

var file_api_dapr_subscribe_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_dapr_subscribe_proto_goTypes = []interface{}{
	(*ListTopicSubscriptionsResponse)(nil), // 0: api.dapr.ListTopicSubscriptionsResponse
	(*TopicSubscription)(nil),              // 1: api.dapr.TopicSubscription
	nil,                                    // 2: api.dapr.TopicSubscription.MetadataEntry
	(*emptypb.Empty)(nil),                  // 3: google.protobuf.Empty
}
var file_api_dapr_subscribe_proto_depIdxs = []int32{
	1, // 0: api.dapr.ListTopicSubscriptionsResponse.subscriptions:type_name -> api.dapr.TopicSubscription
	2, // 1: api.dapr.TopicSubscription.metadata:type_name -> api.dapr.TopicSubscription.MetadataEntry
	3, // 2: api.dapr.Subscribe.GetSubscribe:input_type -> google.protobuf.Empty
	0, // 3: api.dapr.Subscribe.GetSubscribe:output_type -> api.dapr.ListTopicSubscriptionsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_dapr_subscribe_proto_init() }
func file_api_dapr_subscribe_proto_init() {
	if File_api_dapr_subscribe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_dapr_subscribe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTopicSubscriptionsResponse); i {
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
		file_api_dapr_subscribe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicSubscription); i {
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
			RawDescriptor: file_api_dapr_subscribe_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_dapr_subscribe_proto_goTypes,
		DependencyIndexes: file_api_dapr_subscribe_proto_depIdxs,
		MessageInfos:      file_api_dapr_subscribe_proto_msgTypes,
	}.Build()
	File_api_dapr_subscribe_proto = out.File
	file_api_dapr_subscribe_proto_rawDesc = nil
	file_api_dapr_subscribe_proto_goTypes = nil
	file_api_dapr_subscribe_proto_depIdxs = nil
}
