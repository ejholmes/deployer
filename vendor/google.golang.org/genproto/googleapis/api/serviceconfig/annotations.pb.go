// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/api/serviceconfig/annotations.proto
// DO NOT EDIT!

/*
Package google_api is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/genproto/googleapis/api/serviceconfig/annotations.proto
	google.golang.org/genproto/googleapis/api/serviceconfig/auth.proto
	google.golang.org/genproto/googleapis/api/serviceconfig/backend.proto
	google.golang.org/genproto/googleapis/api/serviceconfig/billing.proto
	google.golang.org/genproto/googleapis/api/serviceconfig/consumer.proto
	google.golang.org/genproto/googleapis/api/serviceconfig/context.proto
	google.golang.org/genproto/googleapis/api/serviceconfig/control.proto
	google.golang.org/genproto/googleapis/api/serviceconfig/documentation.proto
	google.golang.org/genproto/googleapis/api/serviceconfig/endpoint.proto
	google.golang.org/genproto/googleapis/api/serviceconfig/http.proto
	google.golang.org/genproto/googleapis/api/serviceconfig/logging.proto
	google.golang.org/genproto/googleapis/api/serviceconfig/log.proto
	google.golang.org/genproto/googleapis/api/serviceconfig/monitoring.proto
	google.golang.org/genproto/googleapis/api/serviceconfig/service.proto
	google.golang.org/genproto/googleapis/api/serviceconfig/system_parameter.proto
	google.golang.org/genproto/googleapis/api/serviceconfig/usage.proto

It has these top-level messages:
	Authentication
	AuthenticationRule
	AuthProvider
	OAuthRequirements
	AuthRequirement
	Backend
	BackendRule
	Billing
	BillingStatusRule
	ProjectProperties
	Property
	Context
	ContextRule
	Control
	Documentation
	DocumentationRule
	Page
	Endpoint
	Http
	HttpRule
	CustomHttpPattern
	Logging
	LogDescriptor
	Monitoring
	Service
	SystemParameters
	SystemParameterRule
	SystemParameter
	Usage
	UsageRule
*/
package google_api // import "google.golang.org/genproto/googleapis/api/serviceconfig"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "google.golang.org/genproto/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

var E_Http = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*HttpRule)(nil),
	Field:         72295728,
	Name:          "google.api.http",
	Tag:           "bytes,72295728,opt,name=http",
}

func init() {
	proto.RegisterExtension(E_Http)
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/api/serviceconfig/annotations.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xf2, 0x4c, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4b, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0xcb, 0x2f, 0x4a, 0xd7, 0x4f, 0x4f,
	0xcd, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x87, 0x48, 0x25, 0x16, 0x64, 0x16, 0xeb, 0x27, 0x16,
	0x64, 0xea, 0x17, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x26, 0xe7, 0xe7, 0xa5, 0x65, 0xa6, 0xeb,
	0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15, 0xeb, 0x81, 0x95, 0x0b, 0x71,
	0x41, 0x8d, 0x4a, 0x2c, 0xc8, 0x94, 0x72, 0x22, 0xd7, 0xd8, 0x8c, 0x92, 0x92, 0x02, 0x88, 0x79,
	0x52, 0x26, 0x78, 0xcc, 0x00, 0x93, 0x49, 0xa5, 0x69, 0xfa, 0x29, 0xa9, 0xc5, 0xc9, 0x45, 0x99,
	0x05, 0x25, 0xf9, 0x45, 0x10, 0x5d, 0x56, 0xde, 0x5c, 0x2c, 0x20, 0x33, 0x84, 0xe4, 0xf4, 0xa0,
	0xda, 0x61, 0x4a, 0xf5, 0x7c, 0x53, 0x4b, 0x32, 0xf2, 0x53, 0xfc, 0x0b, 0xc0, 0x6e, 0x96, 0xd8,
	0x70, 0x6a, 0x8f, 0x92, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x88, 0x1e, 0xc2, 0xdd, 0x7a, 0x1e, 0x25,
	0x25, 0x05, 0x41, 0xa5, 0x39, 0xa9, 0x41, 0x60, 0x43, 0x9c, 0xb4, 0xb9, 0xf8, 0x92, 0xf3, 0x73,
	0x91, 0x14, 0x38, 0x09, 0x38, 0x22, 0xfc, 0x1d, 0x00, 0x32, 0x39, 0x80, 0x71, 0x11, 0x13, 0x8b,
	0xbb, 0x63, 0x80, 0x67, 0x12, 0x1b, 0xd8, 0x26, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x01,
	0xd8, 0x8e, 0xc1, 0x53, 0x01, 0x00, 0x00,
}
