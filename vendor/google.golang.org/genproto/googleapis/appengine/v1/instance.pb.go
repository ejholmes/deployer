// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/appengine/v1/instance.proto
// DO NOT EDIT!

package google_appengine_v1 // import "google.golang.org/genproto/googleapis/appengine/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_protobuf2 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Availability of the instance.
type Instance_Availability int32

const (
	Instance_UNSPECIFIED Instance_Availability = 0
	Instance_RESIDENT    Instance_Availability = 1
	Instance_DYNAMIC     Instance_Availability = 2
)

var Instance_Availability_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "RESIDENT",
	2: "DYNAMIC",
}
var Instance_Availability_value = map[string]int32{
	"UNSPECIFIED": 0,
	"RESIDENT":    1,
	"DYNAMIC":     2,
}

func (x Instance_Availability) String() string {
	return proto.EnumName(Instance_Availability_name, int32(x))
}
func (Instance_Availability) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 0} }

// An Instance resource is the computing unit that App Engine uses to
// automatically scale an application.
type Instance struct {
	// Full path to the Instance resource in the API.
	// Example: `apps/myapp/services/default/versions/v1/instances/instance-1`.
	//
	// @OutputOnly
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Relative name of the instance within the version.
	// Example: `instance-1`.
	//
	// @OutputOnly
	Id string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	// App Engine release this instance is running on.
	//
	// @OutputOnly
	AppEngineRelease string `protobuf:"bytes,3,opt,name=app_engine_release,json=appEngineRelease" json:"app_engine_release,omitempty"`
	// Availability of the instance.
	//
	// @OutputOnly
	Availability Instance_Availability `protobuf:"varint,4,opt,name=availability,enum=google.appengine.v1.Instance_Availability" json:"availability,omitempty"`
	// Name of the virtual machine where this instance lives. Only applicable
	// for instances in App Engine flexible environment.
	//
	// @OutputOnly
	VmName string `protobuf:"bytes,5,opt,name=vm_name,json=vmName" json:"vm_name,omitempty"`
	// Zone where the virtual machine is located. Only applicable for instances
	// in App Engine flexible environment.
	//
	// @OutputOnly
	VmZoneName string `protobuf:"bytes,6,opt,name=vm_zone_name,json=vmZoneName" json:"vm_zone_name,omitempty"`
	// Virtual machine ID of this instance. Only applicable for instances in
	// App Engine flexible environment.
	//
	// @OutputOnly
	VmId string `protobuf:"bytes,7,opt,name=vm_id,json=vmId" json:"vm_id,omitempty"`
	// Time that this instance was started.
	//
	// @OutputOnly
	StartTime *google_protobuf2.Timestamp `protobuf:"bytes,8,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// Number of requests since this instance was started.
	//
	// @OutputOnly
	Requests int32 `protobuf:"varint,9,opt,name=requests" json:"requests,omitempty"`
	// Number of errors since this instance was started.
	//
	// @OutputOnly
	Errors int32 `protobuf:"varint,10,opt,name=errors" json:"errors,omitempty"`
	// Average queries per second (QPS) over the last minute.
	//
	// @OutputOnly
	Qps float32 `protobuf:"fixed32,11,opt,name=qps" json:"qps,omitempty"`
	// Average latency (ms) over the last minute.
	//
	// @OutputOnly
	AverageLatency int32 `protobuf:"varint,12,opt,name=average_latency,json=averageLatency" json:"average_latency,omitempty"`
	// Total memory in use (bytes).
	//
	// @OutputOnly
	MemoryUsage int64 `protobuf:"varint,13,opt,name=memory_usage,json=memoryUsage" json:"memory_usage,omitempty"`
	// Status of the virtual machine where this instance lives. Only applicable
	// for instances in App Engine flexible environment.
	//
	// @OutputOnly
	VmStatus string `protobuf:"bytes,14,opt,name=vm_status,json=vmStatus" json:"vm_status,omitempty"`
	// Whether this instance is in debug mode. Only applicable for instances in
	// App Engine flexible environment.
	//
	// @OutputOnly
	VmDebugEnabled bool `protobuf:"varint,15,opt,name=vm_debug_enabled,json=vmDebugEnabled" json:"vm_debug_enabled,omitempty"`
}

func (m *Instance) Reset()                    { *m = Instance{} }
func (m *Instance) String() string            { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()               {}
func (*Instance) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Instance) GetStartTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func init() {
	proto.RegisterType((*Instance)(nil), "google.appengine.v1.Instance")
	proto.RegisterEnum("google.appengine.v1.Instance_Availability", Instance_Availability_name, Instance_Availability_value)
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/appengine/v1/instance.proto", fileDescriptor4)
}

var fileDescriptor4 = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x51, 0x4d, 0x4f, 0xdb, 0x40,
	0x10, 0xad, 0x03, 0x84, 0x64, 0x12, 0x42, 0xb4, 0x48, 0x65, 0x95, 0x5e, 0x5c, 0x2e, 0xb5, 0x50,
	0x65, 0x0b, 0x7a, 0x6a, 0x7b, 0x02, 0x92, 0x4a, 0x96, 0xda, 0x08, 0x39, 0x70, 0x68, 0x2f, 0xd6,
	0x26, 0x19, 0xdc, 0x95, 0xbc, 0x1f, 0x78, 0xd7, 0x2b, 0xa5, 0xbf, 0xb1, 0x3f, 0xaa, 0xf2, 0x3a,
	0xa1, 0x54, 0xe2, 0xd0, 0xdb, 0xce, 0x7b, 0x33, 0x6f, 0xdf, 0xbc, 0x81, 0xab, 0x42, 0xa9, 0xa2,
	0xc4, 0xb8, 0x50, 0x25, 0x93, 0x45, 0xac, 0xaa, 0x22, 0x29, 0x50, 0xea, 0x4a, 0x59, 0x95, 0xb4,
	0x14, 0xd3, 0xdc, 0x24, 0x4c, 0x6b, 0x94, 0x05, 0x97, 0x98, 0xb8, 0x8b, 0x84, 0x4b, 0x63, 0x99,
	0x5c, 0x61, 0xec, 0xdb, 0xc8, 0xc9, 0x56, 0xe2, 0xa9, 0x27, 0x76, 0x17, 0x93, 0xf4, 0x7f, 0x75,
	0x79, 0x62, 0xb0, 0x72, 0x7c, 0x85, 0x2b, 0x25, 0x1f, 0x78, 0x91, 0x30, 0x29, 0x95, 0x65, 0x96,
	0x2b, 0x69, 0x5a, 0xfd, 0xc9, 0xe7, 0x82, 0xdb, 0x9f, 0xf5, 0x32, 0x5e, 0x29, 0x91, 0xb4, 0x72,
	0x89, 0x27, 0x96, 0xf5, 0x43, 0xa2, 0xed, 0x46, 0xa3, 0x49, 0x2c, 0x17, 0x68, 0x2c, 0x13, 0xfa,
	0xef, 0xab, 0x1d, 0x3e, 0xfb, 0xbd, 0x0f, 0xbd, 0x74, 0xeb, 0x97, 0x10, 0xd8, 0x97, 0x4c, 0x20,
	0x0d, 0xc2, 0x20, 0xea, 0x67, 0xfe, 0x4d, 0x46, 0xd0, 0xe1, 0x6b, 0xda, 0xf1, 0x48, 0x87, 0xaf,
	0xc9, 0x7b, 0x20, 0x4c, 0xeb, 0xbc, 0xdd, 0x24, 0xaf, 0xb0, 0x44, 0x66, 0x90, 0xee, 0x79, 0x7e,
	0xcc, 0xb4, 0x9e, 0x79, 0x22, 0x6b, 0x71, 0x32, 0x87, 0x21, 0x73, 0x8c, 0x97, 0x6c, 0xc9, 0x4b,
	0x6e, 0x37, 0x74, 0x3f, 0x0c, 0xa2, 0xd1, 0xe5, 0x79, 0xfc, 0x42, 0x24, 0xf1, 0xce, 0x46, 0x7c,
	0xf5, 0x6c, 0x22, 0xfb, 0x67, 0x9e, 0x9c, 0xc2, 0xa1, 0x13, 0xb9, 0x37, 0x79, 0xe0, 0xbf, 0xec,
	0x3a, 0x31, 0x6f, 0x6c, 0x86, 0x30, 0x74, 0x22, 0xff, 0xa5, 0x24, 0xb6, 0x6c, 0xd7, 0xb3, 0xe0,
	0xc4, 0x0f, 0x25, 0xd1, 0x77, 0x9c, 0xc0, 0x81, 0x13, 0x39, 0x5f, 0xd3, 0xc3, 0x76, 0x3b, 0x27,
	0xd2, 0x35, 0xf9, 0x08, 0x60, 0x2c, 0xab, 0x6c, 0xde, 0xe4, 0x42, 0x7b, 0x61, 0x10, 0x0d, 0x2e,
	0x27, 0x3b, 0x77, 0xbb, 0x14, 0xe3, 0xbb, 0x5d, 0x68, 0x59, 0xdf, 0x77, 0x37, 0x35, 0x99, 0x40,
	0xaf, 0xc2, 0xc7, 0x1a, 0x8d, 0x35, 0xb4, 0x1f, 0x06, 0xd1, 0x41, 0xf6, 0x54, 0x93, 0xd7, 0xd0,
	0xc5, 0xaa, 0x52, 0x95, 0xa1, 0xe0, 0x99, 0x6d, 0x45, 0xc6, 0xb0, 0xf7, 0xa8, 0x0d, 0x1d, 0x84,
	0x41, 0xd4, 0xc9, 0x9a, 0x27, 0x79, 0x07, 0xc7, 0xcc, 0x61, 0xc5, 0x0a, 0xcc, 0x4b, 0x66, 0x51,
	0xae, 0x36, 0x74, 0xe8, 0x47, 0x46, 0x5b, 0xf8, 0x6b, 0x8b, 0x92, 0xb7, 0x30, 0x14, 0x28, 0x54,
	0xb5, 0xc9, 0x6b, 0xc3, 0x0a, 0xa4, 0x47, 0x61, 0x10, 0xed, 0x65, 0x83, 0x16, 0xbb, 0x6f, 0x20,
	0xf2, 0x06, 0xfa, 0x4e, 0xe4, 0xc6, 0x32, 0x5b, 0x1b, 0x3a, 0xf2, 0x5b, 0xf6, 0x9c, 0x58, 0xf8,
	0x9a, 0x44, 0x30, 0x76, 0x22, 0x5f, 0xe3, 0xb2, 0x2e, 0x72, 0x94, 0x6c, 0x59, 0xe2, 0x9a, 0x1e,
	0x87, 0x41, 0xd4, 0xcb, 0x46, 0x4e, 0x4c, 0x1b, 0x78, 0xd6, 0xa2, 0x67, 0x9f, 0x60, 0xf8, 0xfc,
	0x02, 0xe4, 0x18, 0x06, 0xf7, 0xf3, 0xc5, 0xed, 0xec, 0x26, 0xfd, 0x92, 0xce, 0xa6, 0xe3, 0x57,
	0x64, 0x08, 0xbd, 0x6c, 0xb6, 0x48, 0xa7, 0xb3, 0xf9, 0xdd, 0x38, 0x20, 0x03, 0x38, 0x9c, 0x7e,
	0x9f, 0x5f, 0x7d, 0x4b, 0x6f, 0xc6, 0x9d, 0xeb, 0x73, 0x38, 0x5d, 0x29, 0xf1, 0xd2, 0x79, 0xaf,
	0x8f, 0x76, 0xf7, 0xbd, 0x6d, 0x62, 0xbd, 0x0d, 0x96, 0x5d, 0x9f, 0xef, 0x87, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x9e, 0x6c, 0x5b, 0xd2, 0x63, 0x03, 0x00, 0x00,
}
