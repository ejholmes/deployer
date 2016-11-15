// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/tracing/trace.proto
// DO NOT EDIT!

/*
Package google_tracing_v1 is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/genproto/googleapis/tracing/trace.proto

It has these top-level messages:
	TraceId
	Module
	StackTrace
	LabelValue
	Span
	Trace
*/
package google_tracing_v1 // import "google.golang.org/genproto/googleapis/tracing"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"
import google_rpc "google.golang.org/genproto/googleapis/rpc/status"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The type of the network event. SENT or RECV event.
type Span_TimeEvent_NetworkEvent_Type int32

const (
	Span_TimeEvent_NetworkEvent_UNSPECIFIED Span_TimeEvent_NetworkEvent_Type = 0
	Span_TimeEvent_NetworkEvent_SENT        Span_TimeEvent_NetworkEvent_Type = 1
	Span_TimeEvent_NetworkEvent_RECV        Span_TimeEvent_NetworkEvent_Type = 2
)

var Span_TimeEvent_NetworkEvent_Type_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "SENT",
	2: "RECV",
}
var Span_TimeEvent_NetworkEvent_Type_value = map[string]int32{
	"UNSPECIFIED": 0,
	"SENT":        1,
	"RECV":        2,
}

func (x Span_TimeEvent_NetworkEvent_Type) String() string {
	return proto.EnumName(Span_TimeEvent_NetworkEvent_Type_name, int32(x))
}
func (Span_TimeEvent_NetworkEvent_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0, 1, 0}
}

// The type of the link.
type Span_Link_Type int32

const (
	Span_Link_UNSPECIFIED Span_Link_Type = 0
	Span_Link_CHILD       Span_Link_Type = 1
	Span_Link_PARENT      Span_Link_Type = 2
)

var Span_Link_Type_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "CHILD",
	2: "PARENT",
}
var Span_Link_Type_value = map[string]int32{
	"UNSPECIFIED": 0,
	"CHILD":       1,
	"PARENT":      2,
}

func (x Span_Link_Type) String() string {
	return proto.EnumName(Span_Link_Type_name, int32(x))
}
func (Span_Link_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 1, 0} }

// A TraceId uniquely identifies a Trace. It is conceptually a 128-bit value,
// represented as a string, containing the hex-encoded value.
type TraceId struct {
	// Trace ID specified as a hex-encoded string. *Must* be 32 bytes long.
	HexEncoded string `protobuf:"bytes,1,opt,name=hex_encoded,json=hexEncoded" json:"hex_encoded,omitempty"`
}

func (m *TraceId) Reset()                    { *m = TraceId{} }
func (m *TraceId) String() string            { return proto.CompactTextString(m) }
func (*TraceId) ProtoMessage()               {}
func (*TraceId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Module struct {
	// Binary module.
	// E.g. main binary, kernel modules, and dynamic libraries
	// such as libc.so, sharedlib.so
	Module string `protobuf:"bytes,1,opt,name=module" json:"module,omitempty"`
	// Build_id is a unique identifier for the module,
	// probably a hash of its contents
	BuildId string `protobuf:"bytes,2,opt,name=build_id,json=buildId" json:"build_id,omitempty"`
}

func (m *Module) Reset()                    { *m = Module{} }
func (m *Module) String() string            { return proto.CompactTextString(m) }
func (*Module) ProtoMessage()               {}
func (*Module) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type StackTrace struct {
	// Stack frames of this stack trace.
	StackFrame []*StackTrace_StackFrame `protobuf:"bytes,1,rep,name=stack_frame,json=stackFrame" json:"stack_frame,omitempty"`
	// User can choose to use his own hash function to hash large labels to save
	// network bandwidth and storage.
	// Typical usage is to pass both initially to inform the storage of the
	// mapping. And in subsequent calls, pass in stack_trace_hash_id only.
	// User shall verify the hash value is successfully stored.
	StackTraceHashId uint64 `protobuf:"varint,2,opt,name=stack_trace_hash_id,json=stackTraceHashId" json:"stack_trace_hash_id,omitempty"`
}

func (m *StackTrace) Reset()                    { *m = StackTrace{} }
func (m *StackTrace) String() string            { return proto.CompactTextString(m) }
func (*StackTrace) ProtoMessage()               {}
func (*StackTrace) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StackTrace) GetStackFrame() []*StackTrace_StackFrame {
	if m != nil {
		return m.StackFrame
	}
	return nil
}

// Presents a single stack frame in a stack trace.
type StackTrace_StackFrame struct {
	// Fully qualified names which uniquely identify function/method/etc.
	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName" json:"function_name,omitempty"`
	// Used when function name is ‘mangled’. Not guaranteed to be fully
	// qualified but usually it is.
	OrigFunctionName string `protobuf:"bytes,2,opt,name=orig_function_name,json=origFunctionName" json:"orig_function_name,omitempty"`
	// File name of the frame.
	FileName string `protobuf:"bytes,3,opt,name=file_name,json=fileName" json:"file_name,omitempty"`
	// Line number of the frame.
	LineNumber int64 `protobuf:"varint,4,opt,name=line_number,json=lineNumber" json:"line_number,omitempty"`
	// Column number is important in JavaScript(anonymous functions),
	// Might not be available in some languages.
	ColumnNumber int64 `protobuf:"varint,5,opt,name=column_number,json=columnNumber" json:"column_number,omitempty"`
	// Binary module the code is loaded from.
	LoadModule *Module `protobuf:"bytes,6,opt,name=load_module,json=loadModule" json:"load_module,omitempty"`
	// source_version is deployment specific. It might be
	// better to be stored in deployment metadata.
	// However, in distributed tracing, it’s hard to keep track of
	// source/binary versions at one place for all spans.
	SourceVersion string `protobuf:"bytes,7,opt,name=source_version,json=sourceVersion" json:"source_version,omitempty"`
}

func (m *StackTrace_StackFrame) Reset()                    { *m = StackTrace_StackFrame{} }
func (m *StackTrace_StackFrame) String() string            { return proto.CompactTextString(m) }
func (*StackTrace_StackFrame) ProtoMessage()               {}
func (*StackTrace_StackFrame) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *StackTrace_StackFrame) GetLoadModule() *Module {
	if m != nil {
		return m.LoadModule
	}
	return nil
}

// Allowed label values.
type LabelValue struct {
	// The value of the label.
	//
	// Types that are valid to be assigned to Value:
	//	*LabelValue_StringValue
	//	*LabelValue_IntValue
	//	*LabelValue_BoolValue
	Value isLabelValue_Value `protobuf_oneof:"value"`
}

func (m *LabelValue) Reset()                    { *m = LabelValue{} }
func (m *LabelValue) String() string            { return proto.CompactTextString(m) }
func (*LabelValue) ProtoMessage()               {}
func (*LabelValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isLabelValue_Value interface {
	isLabelValue_Value()
}

type LabelValue_StringValue struct {
	StringValue string `protobuf:"bytes,1,opt,name=string_value,json=stringValue,oneof"`
}
type LabelValue_IntValue struct {
	IntValue int64 `protobuf:"varint,2,opt,name=int_value,json=intValue,oneof"`
}
type LabelValue_BoolValue struct {
	BoolValue bool `protobuf:"varint,3,opt,name=bool_value,json=boolValue,oneof"`
}

func (*LabelValue_StringValue) isLabelValue_Value() {}
func (*LabelValue_IntValue) isLabelValue_Value()    {}
func (*LabelValue_BoolValue) isLabelValue_Value()   {}

func (m *LabelValue) GetValue() isLabelValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *LabelValue) GetStringValue() string {
	if x, ok := m.GetValue().(*LabelValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *LabelValue) GetIntValue() int64 {
	if x, ok := m.GetValue().(*LabelValue_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *LabelValue) GetBoolValue() bool {
	if x, ok := m.GetValue().(*LabelValue_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LabelValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LabelValue_OneofMarshaler, _LabelValue_OneofUnmarshaler, _LabelValue_OneofSizer, []interface{}{
		(*LabelValue_StringValue)(nil),
		(*LabelValue_IntValue)(nil),
		(*LabelValue_BoolValue)(nil),
	}
}

func _LabelValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LabelValue)
	// value
	switch x := m.Value.(type) {
	case *LabelValue_StringValue:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case *LabelValue_IntValue:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntValue))
	case *LabelValue_BoolValue:
		t := uint64(0)
		if x.BoolValue {
			t = 1
		}
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("LabelValue.Value has unexpected type %T", x)
	}
	return nil
}

func _LabelValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LabelValue)
	switch tag {
	case 1: // value.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &LabelValue_StringValue{x}
		return true, err
	case 2: // value.int_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &LabelValue_IntValue{int64(x)}
		return true, err
	case 3: // value.bool_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &LabelValue_BoolValue{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _LabelValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LabelValue)
	// value
	switch x := m.Value.(type) {
	case *LabelValue_StringValue:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *LabelValue_IntValue:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.IntValue))
	case *LabelValue_BoolValue:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A span represents a single operation within a trace. Spans can be nested
// and form a trace tree. Often, a trace contains a root span that describes the
// end-to-end latency and, optionally, one or more subspans for
// its sub-operations. Spans do not need to be contiguous. There may be gaps
// between spans in a trace.
type Span struct {
	// Identifier for the span. Must be a 64-bit integer other than 0 and
	// unique within a trace.
	Id uint64 `protobuf:"fixed64,1,opt,name=id" json:"id,omitempty"`
	// Name of the span. The span name is sanitized and displayed in the
	// Stackdriver Trace tool in the {% dynamic print site_values.console_name %}.
	// The name may be a method name or some other per-call site name.
	// For the same executable and the same call point, a best practice is
	// to use a consistent name, which makes it easier to correlate
	// cross-trace spans.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// ID of parent span. 0 or missing if this is a root span.
	ParentId uint64 `protobuf:"fixed64,3,opt,name=parent_id,json=parentId" json:"parent_id,omitempty"`
	// Local machine clock in nanoseconds from the UNIX epoch,
	// at which span execution started.
	// On the server side these are the times when the server application
	// handler starts running.
	LocalStartTime *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=local_start_time,json=localStartTime" json:"local_start_time,omitempty"`
	// Local machine clock in nanoseconds from the UNIX epoch,
	// at which span execution ended.
	// On the server side these are the times when the server application
	// handler finishes running.
	LocalEndTime *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=local_end_time,json=localEndTime" json:"local_end_time,omitempty"`
	// Properties of a span. Labels at the span level.
	// E.g.
	// "/instance_id": "my-instance"
	// "/zone": "us-central1-a"
	// "/grpc/peer_address": "ip:port" (dns, etc.)
	// "/grpc/deadline": "Duration"
	// "/http/user_agent"
	// "/http/request_bytes": 300
	// "/http/response_bytes": 1200
	// "/http/url": google.com/apis
	// "/pid"
	// "abc.com/mylabel": "my label value"
	Labels map[string]*LabelValue `protobuf:"bytes,6,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Stack trace captured at the start of the span. This is optional.
	StackTrace *StackTrace `protobuf:"bytes,7,opt,name=stack_trace,json=stackTrace" json:"stack_trace,omitempty"`
	// A collection of time-stamped events.
	TimeEvents []*Span_TimeEvent `protobuf:"bytes,8,rep,name=time_events,json=timeEvents" json:"time_events,omitempty"`
	// A collection of links.
	Links []*Span_Link `protobuf:"bytes,9,rep,name=links" json:"links,omitempty"`
	// The final status of the Span. This is optional.
	Status *google_rpc.Status `protobuf:"bytes,10,opt,name=status" json:"status,omitempty"`
	// True if this Span has a remote parent (is an RPC server Span).
	HasRemoteParent bool `protobuf:"varint,11,opt,name=has_remote_parent,json=hasRemoteParent" json:"has_remote_parent,omitempty"`
}

func (m *Span) Reset()                    { *m = Span{} }
func (m *Span) String() string            { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()               {}
func (*Span) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Span) GetLocalStartTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.LocalStartTime
	}
	return nil
}

func (m *Span) GetLocalEndTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.LocalEndTime
	}
	return nil
}

func (m *Span) GetLabels() map[string]*LabelValue {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Span) GetStackTrace() *StackTrace {
	if m != nil {
		return m.StackTrace
	}
	return nil
}

func (m *Span) GetTimeEvents() []*Span_TimeEvent {
	if m != nil {
		return m.TimeEvents
	}
	return nil
}

func (m *Span) GetLinks() []*Span_Link {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *Span) GetStatus() *google_rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// A time-stamped annotation in the Span.
type Span_TimeEvent struct {
	// The local machine absolute timestamp when this event happened.
	LocalTime *google_protobuf1.Timestamp `protobuf:"bytes,1,opt,name=local_time,json=localTime" json:"local_time,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Span_TimeEvent_Annotation_
	//	*Span_TimeEvent_NetworkEvent_
	Value isSpan_TimeEvent_Value `protobuf_oneof:"value"`
}

func (m *Span_TimeEvent) Reset()                    { *m = Span_TimeEvent{} }
func (m *Span_TimeEvent) String() string            { return proto.CompactTextString(m) }
func (*Span_TimeEvent) ProtoMessage()               {}
func (*Span_TimeEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type isSpan_TimeEvent_Value interface {
	isSpan_TimeEvent_Value()
}

type Span_TimeEvent_Annotation_ struct {
	Annotation *Span_TimeEvent_Annotation `protobuf:"bytes,2,opt,name=annotation,oneof"`
}
type Span_TimeEvent_NetworkEvent_ struct {
	NetworkEvent *Span_TimeEvent_NetworkEvent `protobuf:"bytes,3,opt,name=network_event,json=networkEvent,oneof"`
}

func (*Span_TimeEvent_Annotation_) isSpan_TimeEvent_Value()   {}
func (*Span_TimeEvent_NetworkEvent_) isSpan_TimeEvent_Value() {}

func (m *Span_TimeEvent) GetValue() isSpan_TimeEvent_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Span_TimeEvent) GetLocalTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.LocalTime
	}
	return nil
}

func (m *Span_TimeEvent) GetAnnotation() *Span_TimeEvent_Annotation {
	if x, ok := m.GetValue().(*Span_TimeEvent_Annotation_); ok {
		return x.Annotation
	}
	return nil
}

func (m *Span_TimeEvent) GetNetworkEvent() *Span_TimeEvent_NetworkEvent {
	if x, ok := m.GetValue().(*Span_TimeEvent_NetworkEvent_); ok {
		return x.NetworkEvent
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Span_TimeEvent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Span_TimeEvent_OneofMarshaler, _Span_TimeEvent_OneofUnmarshaler, _Span_TimeEvent_OneofSizer, []interface{}{
		(*Span_TimeEvent_Annotation_)(nil),
		(*Span_TimeEvent_NetworkEvent_)(nil),
	}
}

func _Span_TimeEvent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Span_TimeEvent)
	// value
	switch x := m.Value.(type) {
	case *Span_TimeEvent_Annotation_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Annotation); err != nil {
			return err
		}
	case *Span_TimeEvent_NetworkEvent_:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NetworkEvent); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Span_TimeEvent.Value has unexpected type %T", x)
	}
	return nil
}

func _Span_TimeEvent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Span_TimeEvent)
	switch tag {
	case 2: // value.annotation
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Span_TimeEvent_Annotation)
		err := b.DecodeMessage(msg)
		m.Value = &Span_TimeEvent_Annotation_{msg}
		return true, err
	case 3: // value.network_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Span_TimeEvent_NetworkEvent)
		err := b.DecodeMessage(msg)
		m.Value = &Span_TimeEvent_NetworkEvent_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Span_TimeEvent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Span_TimeEvent)
	// value
	switch x := m.Value.(type) {
	case *Span_TimeEvent_Annotation_:
		s := proto.Size(x.Annotation)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Span_TimeEvent_NetworkEvent_:
		s := proto.Size(x.NetworkEvent)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Text annotation with a set of labels.
type Span_TimeEvent_Annotation struct {
	// A user-supplied message describing the event.
	Description string `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	// A set of labels on the annotation.
	Labels map[string]*LabelValue `protobuf:"bytes,2,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Span_TimeEvent_Annotation) Reset()                    { *m = Span_TimeEvent_Annotation{} }
func (m *Span_TimeEvent_Annotation) String() string            { return proto.CompactTextString(m) }
func (*Span_TimeEvent_Annotation) ProtoMessage()               {}
func (*Span_TimeEvent_Annotation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0, 0} }

func (m *Span_TimeEvent_Annotation) GetLabels() map[string]*LabelValue {
	if m != nil {
		return m.Labels
	}
	return nil
}

// An event describing an RPC message sent/received on the network.
type Span_TimeEvent_NetworkEvent struct {
	// If available, this is the kernel time:
	// For sent messages, this is the time at which the first bit was sent.
	// For received messages, this is the time at which the last bit was
	// received.
	KernelTime *google_protobuf1.Timestamp      `protobuf:"bytes,1,opt,name=kernel_time,json=kernelTime" json:"kernel_time,omitempty"`
	Type       Span_TimeEvent_NetworkEvent_Type `protobuf:"varint,2,opt,name=type,enum=google.tracing.v1.Span_TimeEvent_NetworkEvent_Type" json:"type,omitempty"`
	// Every message has an identifier, that must be different from all the
	// network messages in this span.
	// This is very important when the request/response are streamed.
	MessageId uint64 `protobuf:"varint,3,opt,name=message_id,json=messageId" json:"message_id,omitempty"`
	// Number of bytes send/receive.
	MessageSize uint64 `protobuf:"varint,4,opt,name=message_size,json=messageSize" json:"message_size,omitempty"`
}

func (m *Span_TimeEvent_NetworkEvent) Reset()         { *m = Span_TimeEvent_NetworkEvent{} }
func (m *Span_TimeEvent_NetworkEvent) String() string { return proto.CompactTextString(m) }
func (*Span_TimeEvent_NetworkEvent) ProtoMessage()    {}
func (*Span_TimeEvent_NetworkEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0, 1}
}

func (m *Span_TimeEvent_NetworkEvent) GetKernelTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.KernelTime
	}
	return nil
}

// Link one span with another which may be in a different Trace. Used (for
// example) in batching operations, where a single batch handler processes
// multiple requests from different traces.
type Span_Link struct {
	// The trace and span identifier of the linked span.
	TraceId *TraceId       `protobuf:"bytes,1,opt,name=trace_id,json=traceId" json:"trace_id,omitempty"`
	SpanId  uint64         `protobuf:"fixed64,2,opt,name=span_id,json=spanId" json:"span_id,omitempty"`
	Type    Span_Link_Type `protobuf:"varint,3,opt,name=type,enum=google.tracing.v1.Span_Link_Type" json:"type,omitempty"`
}

func (m *Span_Link) Reset()                    { *m = Span_Link{} }
func (m *Span_Link) String() string            { return proto.CompactTextString(m) }
func (*Span_Link) ProtoMessage()               {}
func (*Span_Link) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 1} }

func (m *Span_Link) GetTraceId() *TraceId {
	if m != nil {
		return m.TraceId
	}
	return nil
}

// A trace describes how long it takes for an application to perform some
// operations. It consists of a tree of spans, each of which contains details
// about an operation with time information and operation details.
type Trace struct {
	// Globally unique identifier for the trace. Common to all the spans.
	TraceId *TraceId `protobuf:"bytes,1,opt,name=trace_id,json=traceId" json:"trace_id,omitempty"`
	// Collection of spans in the trace. The root span has parent_id == 0.
	Spans []*Span `protobuf:"bytes,2,rep,name=spans" json:"spans,omitempty"`
}

func (m *Trace) Reset()                    { *m = Trace{} }
func (m *Trace) String() string            { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()               {}
func (*Trace) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Trace) GetTraceId() *TraceId {
	if m != nil {
		return m.TraceId
	}
	return nil
}

func (m *Trace) GetSpans() []*Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

func init() {
	proto.RegisterType((*TraceId)(nil), "google.tracing.v1.TraceId")
	proto.RegisterType((*Module)(nil), "google.tracing.v1.Module")
	proto.RegisterType((*StackTrace)(nil), "google.tracing.v1.StackTrace")
	proto.RegisterType((*StackTrace_StackFrame)(nil), "google.tracing.v1.StackTrace.StackFrame")
	proto.RegisterType((*LabelValue)(nil), "google.tracing.v1.LabelValue")
	proto.RegisterType((*Span)(nil), "google.tracing.v1.Span")
	proto.RegisterType((*Span_TimeEvent)(nil), "google.tracing.v1.Span.TimeEvent")
	proto.RegisterType((*Span_TimeEvent_Annotation)(nil), "google.tracing.v1.Span.TimeEvent.Annotation")
	proto.RegisterType((*Span_TimeEvent_NetworkEvent)(nil), "google.tracing.v1.Span.TimeEvent.NetworkEvent")
	proto.RegisterType((*Span_Link)(nil), "google.tracing.v1.Span.Link")
	proto.RegisterType((*Trace)(nil), "google.tracing.v1.Trace")
	proto.RegisterEnum("google.tracing.v1.Span_TimeEvent_NetworkEvent_Type", Span_TimeEvent_NetworkEvent_Type_name, Span_TimeEvent_NetworkEvent_Type_value)
	proto.RegisterEnum("google.tracing.v1.Span_Link_Type", Span_Link_Type_name, Span_Link_Type_value)
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/tracing/trace.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 1122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x55, 0xdb, 0x6e, 0x1b, 0x37,
	0x13, 0xd6, 0xf9, 0x30, 0xab, 0x38, 0x0a, 0x7f, 0xfc, 0x8d, 0xa2, 0xd6, 0xb0, 0xad, 0x20, 0x80,
	0xe1, 0x26, 0x2b, 0x54, 0x46, 0x80, 0x24, 0x46, 0x8b, 0xc6, 0xb6, 0x5c, 0x09, 0x48, 0x05, 0x81,
	0x72, 0x8c, 0xde, 0x2d, 0xa8, 0x5d, 0x7a, 0x45, 0x68, 0xc5, 0x5d, 0x2c, 0x29, 0x37, 0xf6, 0x6d,
	0xdf, 0xa0, 0xef, 0xd0, 0xdb, 0xbe, 0x41, 0x1f, 0xa4, 0x7d, 0x9a, 0x82, 0x87, 0x95, 0x55, 0xc4,
	0x8e, 0x9d, 0x02, 0xbd, 0x5a, 0xf2, 0x9b, 0x6f, 0x66, 0x87, 0xdf, 0xcc, 0x90, 0xf0, 0x3a, 0x8c,
	0xe3, 0x30, 0xa2, 0x6e, 0x18, 0x47, 0x84, 0x87, 0x6e, 0x9c, 0x86, 0xdd, 0x90, 0xf2, 0x24, 0x8d,
	0x65, 0xdc, 0x35, 0x26, 0x92, 0x30, 0xd1, 0x95, 0x29, 0xf1, 0x19, 0x0f, 0xf5, 0x97, 0xba, 0xda,
	0x8c, 0x1e, 0x59, 0x57, 0x6b, 0x73, 0x2f, 0xbe, 0x69, 0x0f, 0xef, 0x17, 0x8d, 0x24, 0xac, 0x2b,
	0x68, 0x7a, 0xc1, 0x7c, 0xea, 0xc7, 0xfc, 0x9c, 0x85, 0x5d, 0xc2, 0x79, 0x2c, 0x89, 0x64, 0x31,
	0x17, 0x26, 0x7a, 0xfb, 0x20, 0x64, 0x72, 0xb6, 0x9c, 0xba, 0x7e, 0xbc, 0xe8, 0x9a, 0x70, 0x5d,
	0x6d, 0x98, 0x2e, 0xcf, 0xbb, 0x89, 0xbc, 0x4c, 0xa8, 0xe8, 0x4a, 0xb6, 0xa0, 0x42, 0x92, 0x45,
	0x72, 0xbd, 0xb2, 0xce, 0xdf, 0xde, 0x2f, 0x8f, 0x34, 0xf1, 0xbb, 0x42, 0x12, 0xb9, 0x14, 0xf6,
	0x63, 0xdc, 0x3b, 0x7b, 0x50, 0x3d, 0x55, 0x07, 0x1d, 0x06, 0x68, 0x0b, 0x9c, 0x19, 0xfd, 0xe0,
	0x51, 0xee, 0xc7, 0x01, 0x0d, 0x5a, 0xf9, 0xed, 0xfc, 0x6e, 0x1d, 0xc3, 0x8c, 0x7e, 0xe8, 0x1b,
	0xa4, 0x73, 0x00, 0x95, 0x1f, 0xe3, 0x60, 0x19, 0x51, 0xf4, 0x05, 0x54, 0x16, 0x7a, 0x65, 0x59,
	0x76, 0x87, 0x9e, 0x40, 0x6d, 0xba, 0x64, 0x51, 0xe0, 0xb1, 0xa0, 0x55, 0xd0, 0x96, 0xaa, 0xde,
	0x0f, 0x83, 0xce, 0xef, 0x45, 0x80, 0x89, 0x24, 0xfe, 0x5c, 0xff, 0x0e, 0x0d, 0xc1, 0x11, 0x6a,
	0xe7, 0x9d, 0xa7, 0x64, 0xa1, 0xc2, 0x14, 0x77, 0x9d, 0xde, 0xae, 0xfb, 0x91, 0xce, 0xee, 0xb5,
	0x8f, 0x59, 0x9e, 0x28, 0x3e, 0x06, 0xb1, 0x5a, 0xa3, 0x17, 0xf0, 0x3f, 0x13, 0x4a, 0x57, 0xcc,
	0x9b, 0x11, 0x31, 0xcb, 0xfe, 0x5f, 0xc2, 0x4d, 0xb1, 0xf2, 0x1f, 0x10, 0x31, 0x1b, 0x06, 0xed,
	0xdf, 0x0a, 0x36, 0x11, 0xe3, 0xfd, 0x14, 0x1e, 0x9c, 0x2f, 0xb9, 0xaf, 0xea, 0xe1, 0x71, 0x93,
	0x8a, 0xca, 0xbb, 0x91, 0x81, 0x23, 0x45, 0x7a, 0x0e, 0x28, 0x4e, 0x59, 0xe8, 0xfd, 0x93, 0x69,
	0x4e, 0xd8, 0x54, 0x96, 0x93, 0x75, 0xf6, 0x97, 0x50, 0x3f, 0x67, 0x11, 0x35, 0xa4, 0xa2, 0x26,
	0xd5, 0x14, 0xa0, 0x8d, 0x5b, 0xe0, 0x44, 0x8c, 0x53, 0x8f, 0x2f, 0x17, 0x53, 0x9a, 0xb6, 0x4a,
	0xdb, 0xf9, 0xdd, 0x22, 0x06, 0x05, 0x8d, 0x34, 0xa2, 0x12, 0xf2, 0xe3, 0x68, 0xb9, 0xe0, 0x19,
	0xa5, 0xac, 0x29, 0x0d, 0x03, 0x5a, 0xd2, 0x1b, 0x70, 0xa2, 0x98, 0x04, 0x9e, 0xad, 0x42, 0x65,
	0x3b, 0xbf, 0xeb, 0xf4, 0x9e, 0xdc, 0x20, 0x9f, 0x29, 0x18, 0x06, 0xc5, 0xb6, 0xc5, 0x7b, 0x06,
	0x1b, 0x22, 0x5e, 0xa6, 0x3e, 0xf5, 0x2e, 0x68, 0x2a, 0x58, 0xcc, 0x5b, 0x55, 0x9d, 0xe3, 0x03,
	0x83, 0x9e, 0x19, 0xb0, 0x73, 0x05, 0xf0, 0x8e, 0x4c, 0x69, 0x74, 0x46, 0xa2, 0xa5, 0x92, 0xa9,
	0x21, 0x64, 0xca, 0x78, 0xe8, 0x5d, 0xa8, 0xbd, 0x51, 0x69, 0x90, 0xc3, 0x8e, 0x41, 0x0d, 0x69,
	0x13, 0xea, 0x8c, 0x4b, 0xcb, 0x50, 0xea, 0x14, 0x07, 0x39, 0x5c, 0x63, 0x5c, 0x1a, 0xf3, 0x16,
	0xc0, 0x34, 0x8e, 0x23, 0x6b, 0x57, 0xc2, 0xd4, 0x06, 0x39, 0x5c, 0x57, 0x98, 0x26, 0x1c, 0x56,
	0xa1, 0xac, 0x6d, 0x9d, 0x5f, 0x1b, 0x50, 0x9a, 0x24, 0x84, 0xa3, 0x0d, 0x28, 0x30, 0xd3, 0x8a,
	0x15, 0x5c, 0x60, 0x01, 0x42, 0x50, 0x5a, 0x93, 0x5e, 0xaf, 0x95, 0xdc, 0x09, 0x49, 0x29, 0x97,
	0xaa, 0xea, 0x45, 0x4d, 0xad, 0x19, 0x60, 0x18, 0xa0, 0x63, 0x68, 0x46, 0xb1, 0x4f, 0x22, 0x4f,
	0x48, 0x92, 0x4a, 0x4f, 0x4d, 0x8f, 0xd6, 0xdc, 0xe9, 0xb5, 0x33, 0xb5, 0xb2, 0x59, 0x73, 0x4f,
	0xb3, 0xd1, 0xc2, 0x1b, 0xda, 0x67, 0xa2, 0x5c, 0x14, 0x88, 0xbe, 0x07, 0x83, 0x78, 0x94, 0x07,
	0x26, 0x46, 0xf9, 0xce, 0x18, 0x0d, 0xed, 0xd1, 0xe7, 0x81, 0x8e, 0x70, 0x00, 0x95, 0x48, 0xa9,
	0x29, 0x5a, 0x15, 0xdd, 0xea, 0x4f, 0x6f, 0x6a, 0xf5, 0x84, 0x70, 0x57, 0x6b, 0x2e, 0xfa, 0x5c,
	0xa6, 0x97, 0xd8, 0xba, 0xa0, 0xef, 0xb2, 0x61, 0xd1, 0x1d, 0xae, 0xcb, 0xe5, 0xf4, 0x36, 0x3f,
	0x39, 0x2c, 0x76, 0x42, 0xcc, 0xb0, 0x1d, 0x82, 0xa3, 0x92, 0xf6, 0xe8, 0x05, 0xe5, 0x52, 0xb4,
	0x6a, 0x3a, 0x83, 0x9d, 0xdb, 0x32, 0x50, 0xf9, 0xf6, 0x15, 0x13, 0x83, 0xcc, 0x96, 0x02, 0xf5,
	0xa0, 0x1c, 0x31, 0x3e, 0x17, 0xad, 0xba, 0xf6, 0xfe, 0xea, 0xd6, 0xfc, 0x19, 0x9f, 0x63, 0x43,
	0x45, 0x7b, 0x50, 0x31, 0x97, 0x4d, 0x0b, 0x74, 0xca, 0x28, 0x73, 0x4a, 0x13, 0x5f, 0xe5, 0x2a,
	0x97, 0x02, 0x5b, 0x06, 0xda, 0x83, 0x47, 0x33, 0x22, 0xbc, 0x94, 0x2e, 0x62, 0x49, 0x3d, 0x53,
	0xbf, 0x96, 0xa3, 0x7a, 0x04, 0x3f, 0x9c, 0x11, 0x81, 0x35, 0x3e, 0xd6, 0x70, 0xfb, 0xcf, 0x32,
	0xd4, 0x57, 0x59, 0xa2, 0xd7, 0x00, 0xa6, 0x38, 0xba, 0x30, 0xf9, 0x3b, 0x0b, 0x53, 0xd7, 0x6c,
	0x5d, 0x95, 0x11, 0xc0, 0xf5, 0x75, 0xac, 0x9b, 0xca, 0xe9, 0x3d, 0xbf, 0x53, 0x17, 0xf7, 0xed,
	0xca, 0x67, 0x90, 0xc3, 0x6b, 0x11, 0xd0, 0x7b, 0x78, 0xc0, 0xa9, 0xfc, 0x39, 0x4e, 0xe7, 0x46,
	0x6b, 0xdd, 0x8e, 0x4e, 0xcf, 0xbd, 0x3b, 0xe4, 0xc8, 0xb8, 0xe9, 0xcd, 0x20, 0x87, 0x1b, 0x7c,
	0x6d, 0xdf, 0xfe, 0x2b, 0x0f, 0x70, 0xfd, 0x4f, 0xb4, 0x0d, 0x4e, 0x40, 0x85, 0x9f, 0xb2, 0x44,
	0xa7, 0x6d, 0x2e, 0xac, 0x75, 0x08, 0x8d, 0x57, 0xdd, 0x56, 0xd0, 0xd5, 0x7a, 0xf5, 0x39, 0x67,
	0xba, 0xa9, 0x05, 0xdb, 0x3f, 0x81, 0xb3, 0x06, 0xa3, 0x26, 0x14, 0xe7, 0xf4, 0xd2, 0xfe, 0x5a,
	0x2d, 0xd1, 0xbe, 0x9d, 0x5d, 0xab, 0xe2, 0x4d, 0xdd, 0x79, 0x7d, 0x9d, 0x60, 0xc3, 0x7d, 0x53,
	0x78, 0x95, 0x6f, 0xff, 0x52, 0x80, 0xc6, 0xfa, 0xe9, 0xd1, 0x01, 0x38, 0x73, 0x9a, 0x72, 0x7a,
	0xef, 0x82, 0x82, 0xa1, 0xeb, 0x8a, 0xfe, 0x00, 0x25, 0xf5, 0x62, 0xea, 0x2c, 0x36, 0x7a, 0xfb,
	0x9f, 0x27, 0xbc, 0x7b, 0x7a, 0x99, 0x50, 0xac, 0x03, 0xa0, 0x4d, 0x80, 0x05, 0x15, 0x82, 0x84,
	0x34, 0xbb, 0x56, 0x4a, 0xb8, 0x6e, 0x91, 0x61, 0x80, 0x76, 0xa0, 0x91, 0x99, 0x05, 0xbb, 0x32,
	0x77, 0x4a, 0x09, 0x3b, 0x16, 0x9b, 0xb0, 0x2b, 0xda, 0xf9, 0x1a, 0x4a, 0x2a, 0x1e, 0x7a, 0x08,
	0xce, 0xfb, 0xd1, 0x64, 0xdc, 0x3f, 0x1a, 0x9e, 0x0c, 0xfb, 0xc7, 0xcd, 0x1c, 0xaa, 0x41, 0x69,
	0xd2, 0x1f, 0x9d, 0x36, 0xf3, 0x6a, 0x85, 0xfb, 0x47, 0x67, 0xcd, 0xc2, 0xea, 0xea, 0x6b, 0xff,
	0x91, 0x87, 0x92, 0x9a, 0x21, 0xf4, 0x12, 0x6a, 0xe6, 0x41, 0xb3, 0x17, 0xe0, 0x9a, 0x06, 0x6b,
	0xa7, 0xb1, 0x8f, 0x37, 0xae, 0x4a, 0xfb, 0x8a, 0x3f, 0x86, 0xaa, 0x48, 0x08, 0xcf, 0x5e, 0xc0,
	0x0a, 0xae, 0xa8, 0xed, 0x30, 0x40, 0x2f, 0xad, 0x32, 0x45, 0xad, 0xcc, 0xce, 0xa7, 0xe6, 0x77,
	0x4d, 0x87, 0x8e, 0x7b, 0xdb, 0x29, 0xea, 0x50, 0x3e, 0x1a, 0x0c, 0xdf, 0x1d, 0x37, 0xf3, 0x08,
	0xa0, 0x32, 0x7e, 0x8b, 0xd5, 0x91, 0x0a, 0xff, 0x5d, 0xa3, 0x74, 0x16, 0x50, 0x36, 0xd7, 0xd9,
	0xbf, 0x54, 0xe6, 0x05, 0x94, 0x95, 0x14, 0xd9, 0x4c, 0x3c, 0xbe, 0x45, 0x01, 0x6c, 0x58, 0x87,
	0xcf, 0xe0, 0xff, 0x7e, 0xbc, 0xf8, 0x98, 0x74, 0x08, 0x3a, 0xf2, 0x58, 0xf5, 0xe1, 0x38, 0x3f,
	0xad, 0xe8, 0x86, 0xdc, 0xff, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x32, 0x14, 0xf4, 0x2b, 0x5e, 0x0a,
	0x00, 0x00,
}
