// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/events.proto

package voltha // import "github.com/opencord/voltha-protos/go/voltha"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/opencord/voltha-protos/go/common"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ConfigEventType_ConfigEventType int32

const (
	ConfigEventType_add    ConfigEventType_ConfigEventType = 0
	ConfigEventType_remove ConfigEventType_ConfigEventType = 1
	ConfigEventType_update ConfigEventType_ConfigEventType = 2
)

var ConfigEventType_ConfigEventType_name = map[int32]string{
	0: "add",
	1: "remove",
	2: "update",
}
var ConfigEventType_ConfigEventType_value = map[string]int32{
	"add":    0,
	"remove": 1,
	"update": 2,
}

func (x ConfigEventType_ConfigEventType) String() string {
	return proto.EnumName(ConfigEventType_ConfigEventType_name, int32(x))
}
func (ConfigEventType_ConfigEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_events_26477559e5dbfb38, []int{0, 0}
}

type KpiEventType_KpiEventType int32

const (
	KpiEventType_slice KpiEventType_KpiEventType = 0
	KpiEventType_ts    KpiEventType_KpiEventType = 1
)

var KpiEventType_KpiEventType_name = map[int32]string{
	0: "slice",
	1: "ts",
}
var KpiEventType_KpiEventType_value = map[string]int32{
	"slice": 0,
	"ts":    1,
}

func (x KpiEventType_KpiEventType) String() string {
	return proto.EnumName(KpiEventType_KpiEventType_name, int32(x))
}
func (KpiEventType_KpiEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_events_26477559e5dbfb38, []int{2, 0}
}

type AlarmEventType_AlarmEventType int32

const (
	AlarmEventType_COMMUNICATION AlarmEventType_AlarmEventType = 0
	AlarmEventType_ENVIRONMENT   AlarmEventType_AlarmEventType = 1
	AlarmEventType_EQUIPMENT     AlarmEventType_AlarmEventType = 2
	AlarmEventType_SERVICE       AlarmEventType_AlarmEventType = 3
	AlarmEventType_PROCESSING    AlarmEventType_AlarmEventType = 4
	AlarmEventType_SECURITY      AlarmEventType_AlarmEventType = 5
)

var AlarmEventType_AlarmEventType_name = map[int32]string{
	0: "COMMUNICATION",
	1: "ENVIRONMENT",
	2: "EQUIPMENT",
	3: "SERVICE",
	4: "PROCESSING",
	5: "SECURITY",
}
var AlarmEventType_AlarmEventType_value = map[string]int32{
	"COMMUNICATION": 0,
	"ENVIRONMENT":   1,
	"EQUIPMENT":     2,
	"SERVICE":       3,
	"PROCESSING":    4,
	"SECURITY":      5,
}

func (x AlarmEventType_AlarmEventType) String() string {
	return proto.EnumName(AlarmEventType_AlarmEventType_name, int32(x))
}
func (AlarmEventType_AlarmEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_events_26477559e5dbfb38, []int{8, 0}
}

type AlarmEventCategory_AlarmEventCategory int32

const (
	AlarmEventCategory_PON AlarmEventCategory_AlarmEventCategory = 0
	AlarmEventCategory_OLT AlarmEventCategory_AlarmEventCategory = 1
	AlarmEventCategory_ONT AlarmEventCategory_AlarmEventCategory = 2
	AlarmEventCategory_ONU AlarmEventCategory_AlarmEventCategory = 3
	AlarmEventCategory_NNI AlarmEventCategory_AlarmEventCategory = 4
)

var AlarmEventCategory_AlarmEventCategory_name = map[int32]string{
	0: "PON",
	1: "OLT",
	2: "ONT",
	3: "ONU",
	4: "NNI",
}
var AlarmEventCategory_AlarmEventCategory_value = map[string]int32{
	"PON": 0,
	"OLT": 1,
	"ONT": 2,
	"ONU": 3,
	"NNI": 4,
}

func (x AlarmEventCategory_AlarmEventCategory) String() string {
	return proto.EnumName(AlarmEventCategory_AlarmEventCategory_name, int32(x))
}
func (AlarmEventCategory_AlarmEventCategory) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_events_26477559e5dbfb38, []int{9, 0}
}

type AlarmEventState_AlarmEventState int32

const (
	AlarmEventState_RAISED  AlarmEventState_AlarmEventState = 0
	AlarmEventState_CLEARED AlarmEventState_AlarmEventState = 1
)

var AlarmEventState_AlarmEventState_name = map[int32]string{
	0: "RAISED",
	1: "CLEARED",
}
var AlarmEventState_AlarmEventState_value = map[string]int32{
	"RAISED":  0,
	"CLEARED": 1,
}

func (x AlarmEventState_AlarmEventState) String() string {
	return proto.EnumName(AlarmEventState_AlarmEventState_name, int32(x))
}
func (AlarmEventState_AlarmEventState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_events_26477559e5dbfb38, []int{10, 0}
}

type AlarmEventSeverity_AlarmEventSeverity int32

const (
	AlarmEventSeverity_INDETERMINATE AlarmEventSeverity_AlarmEventSeverity = 0
	AlarmEventSeverity_WARNING       AlarmEventSeverity_AlarmEventSeverity = 1
	AlarmEventSeverity_MINOR         AlarmEventSeverity_AlarmEventSeverity = 2
	AlarmEventSeverity_MAJOR         AlarmEventSeverity_AlarmEventSeverity = 3
	AlarmEventSeverity_CRITICAL      AlarmEventSeverity_AlarmEventSeverity = 4
)

var AlarmEventSeverity_AlarmEventSeverity_name = map[int32]string{
	0: "INDETERMINATE",
	1: "WARNING",
	2: "MINOR",
	3: "MAJOR",
	4: "CRITICAL",
}
var AlarmEventSeverity_AlarmEventSeverity_value = map[string]int32{
	"INDETERMINATE": 0,
	"WARNING":       1,
	"MINOR":         2,
	"MAJOR":         3,
	"CRITICAL":      4,
}

func (x AlarmEventSeverity_AlarmEventSeverity) String() string {
	return proto.EnumName(AlarmEventSeverity_AlarmEventSeverity_name, int32(x))
}
func (AlarmEventSeverity_AlarmEventSeverity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_events_26477559e5dbfb38, []int{11, 0}
}

type ConfigEventType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigEventType) Reset()         { *m = ConfigEventType{} }
func (m *ConfigEventType) String() string { return proto.CompactTextString(m) }
func (*ConfigEventType) ProtoMessage()    {}
func (*ConfigEventType) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_26477559e5dbfb38, []int{0}
}
func (m *ConfigEventType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigEventType.Unmarshal(m, b)
}
func (m *ConfigEventType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigEventType.Marshal(b, m, deterministic)
}
func (dst *ConfigEventType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigEventType.Merge(dst, src)
}
func (m *ConfigEventType) XXX_Size() int {
	return xxx_messageInfo_ConfigEventType.Size(m)
}
func (m *ConfigEventType) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigEventType.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigEventType proto.InternalMessageInfo

type ConfigEvent struct {
	Type                 ConfigEventType_ConfigEventType `protobuf:"varint,1,opt,name=type,proto3,enum=voltha.ConfigEventType_ConfigEventType" json:"type,omitempty"`
	Hash                 string                          `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Data                 string                          `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ConfigEvent) Reset()         { *m = ConfigEvent{} }
func (m *ConfigEvent) String() string { return proto.CompactTextString(m) }
func (*ConfigEvent) ProtoMessage()    {}
func (*ConfigEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_26477559e5dbfb38, []int{1}
}
func (m *ConfigEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigEvent.Unmarshal(m, b)
}
func (m *ConfigEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigEvent.Marshal(b, m, deterministic)
}
func (dst *ConfigEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigEvent.Merge(dst, src)
}
func (m *ConfigEvent) XXX_Size() int {
	return xxx_messageInfo_ConfigEvent.Size(m)
}
func (m *ConfigEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigEvent proto.InternalMessageInfo

func (m *ConfigEvent) GetType() ConfigEventType_ConfigEventType {
	if m != nil {
		return m.Type
	}
	return ConfigEventType_add
}

func (m *ConfigEvent) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *ConfigEvent) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type KpiEventType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KpiEventType) Reset()         { *m = KpiEventType{} }
func (m *KpiEventType) String() string { return proto.CompactTextString(m) }
func (*KpiEventType) ProtoMessage()    {}
func (*KpiEventType) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_26477559e5dbfb38, []int{2}
}
func (m *KpiEventType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KpiEventType.Unmarshal(m, b)
}
func (m *KpiEventType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KpiEventType.Marshal(b, m, deterministic)
}
func (dst *KpiEventType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KpiEventType.Merge(dst, src)
}
func (m *KpiEventType) XXX_Size() int {
	return xxx_messageInfo_KpiEventType.Size(m)
}
func (m *KpiEventType) XXX_DiscardUnknown() {
	xxx_messageInfo_KpiEventType.DiscardUnknown(m)
}

var xxx_messageInfo_KpiEventType proto.InternalMessageInfo

//
// Struct to convey a dictionary of metric metadata.
type MetricMetaData struct {
	Title           string  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Ts              float64 `protobuf:"fixed64,2,opt,name=ts,proto3" json:"ts,omitempty"`
	LogicalDeviceId string  `protobuf:"bytes,3,opt,name=logical_device_id,json=logicalDeviceId,proto3" json:"logical_device_id,omitempty"`
	// (equivalent to the DPID that ONOS has
	// for the VOLTHA device without the
	//  'of:' prefix
	SerialNo             string            `protobuf:"bytes,4,opt,name=serial_no,json=serialNo,proto3" json:"serial_no,omitempty"`
	DeviceId             string            `protobuf:"bytes,5,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Context              map[string]string `protobuf:"bytes,6,rep,name=context,proto3" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MetricMetaData) Reset()         { *m = MetricMetaData{} }
func (m *MetricMetaData) String() string { return proto.CompactTextString(m) }
func (*MetricMetaData) ProtoMessage()    {}
func (*MetricMetaData) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_26477559e5dbfb38, []int{3}
}
func (m *MetricMetaData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricMetaData.Unmarshal(m, b)
}
func (m *MetricMetaData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricMetaData.Marshal(b, m, deterministic)
}
func (dst *MetricMetaData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricMetaData.Merge(dst, src)
}
func (m *MetricMetaData) XXX_Size() int {
	return xxx_messageInfo_MetricMetaData.Size(m)
}
func (m *MetricMetaData) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricMetaData.DiscardUnknown(m)
}

var xxx_messageInfo_MetricMetaData proto.InternalMessageInfo

func (m *MetricMetaData) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MetricMetaData) GetTs() float64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *MetricMetaData) GetLogicalDeviceId() string {
	if m != nil {
		return m.LogicalDeviceId
	}
	return ""
}

func (m *MetricMetaData) GetSerialNo() string {
	if m != nil {
		return m.SerialNo
	}
	return ""
}

func (m *MetricMetaData) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *MetricMetaData) GetContext() map[string]string {
	if m != nil {
		return m.Context
	}
	return nil
}

//
// Struct to convey a dictionary of metric->value pairs. Typically used in
// pure shared-timestamp or shared-timestamp + shared object prefix situations.
type MetricValuePairs struct {
	// Metric / value pairs.
	Metrics              map[string]float32 `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MetricValuePairs) Reset()         { *m = MetricValuePairs{} }
func (m *MetricValuePairs) String() string { return proto.CompactTextString(m) }
func (*MetricValuePairs) ProtoMessage()    {}
func (*MetricValuePairs) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_26477559e5dbfb38, []int{4}
}
func (m *MetricValuePairs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricValuePairs.Unmarshal(m, b)
}
func (m *MetricValuePairs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricValuePairs.Marshal(b, m, deterministic)
}
func (dst *MetricValuePairs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricValuePairs.Merge(dst, src)
}
func (m *MetricValuePairs) XXX_Size() int {
	return xxx_messageInfo_MetricValuePairs.Size(m)
}
func (m *MetricValuePairs) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricValuePairs.DiscardUnknown(m)
}

var xxx_messageInfo_MetricValuePairs proto.InternalMessageInfo

func (m *MetricValuePairs) GetMetrics() map[string]float32 {
	if m != nil {
		return m.Metrics
	}
	return nil
}

//
// Struct to group metadata for a metric (or group of metrics) with the key-value
// pairs of collected metrics
type MetricInformation struct {
	Metadata             *MetricMetaData    `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Metrics              map[string]float32 `protobuf:"bytes,2,rep,name=metrics,proto3" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MetricInformation) Reset()         { *m = MetricInformation{} }
func (m *MetricInformation) String() string { return proto.CompactTextString(m) }
func (*MetricInformation) ProtoMessage()    {}
func (*MetricInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_26477559e5dbfb38, []int{5}
}
func (m *MetricInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricInformation.Unmarshal(m, b)
}
func (m *MetricInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricInformation.Marshal(b, m, deterministic)
}
func (dst *MetricInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricInformation.Merge(dst, src)
}
func (m *MetricInformation) XXX_Size() int {
	return xxx_messageInfo_MetricInformation.Size(m)
}
func (m *MetricInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricInformation.DiscardUnknown(m)
}

var xxx_messageInfo_MetricInformation proto.InternalMessageInfo

func (m *MetricInformation) GetMetadata() *MetricMetaData {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *MetricInformation) GetMetrics() map[string]float32 {
	if m != nil {
		return m.Metrics
	}
	return nil
}

//
// Legacy KPI Event structured.  In mid-August, the KPI event format was updated
//                               to a more easily parsable format. See VOL-1140
//                               for more information.
type KpiEvent struct {
	Type                 KpiEventType_KpiEventType    `protobuf:"varint,1,opt,name=type,proto3,enum=voltha.KpiEventType_KpiEventType" json:"type,omitempty"`
	Ts                   float32                      `protobuf:"fixed32,2,opt,name=ts,proto3" json:"ts,omitempty"`
	Prefixes             map[string]*MetricValuePairs `protobuf:"bytes,3,rep,name=prefixes,proto3" json:"prefixes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *KpiEvent) Reset()         { *m = KpiEvent{} }
func (m *KpiEvent) String() string { return proto.CompactTextString(m) }
func (*KpiEvent) ProtoMessage()    {}
func (*KpiEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_26477559e5dbfb38, []int{6}
}
func (m *KpiEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KpiEvent.Unmarshal(m, b)
}
func (m *KpiEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KpiEvent.Marshal(b, m, deterministic)
}
func (dst *KpiEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KpiEvent.Merge(dst, src)
}
func (m *KpiEvent) XXX_Size() int {
	return xxx_messageInfo_KpiEvent.Size(m)
}
func (m *KpiEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_KpiEvent.DiscardUnknown(m)
}

var xxx_messageInfo_KpiEvent proto.InternalMessageInfo

func (m *KpiEvent) GetType() KpiEventType_KpiEventType {
	if m != nil {
		return m.Type
	}
	return KpiEventType_slice
}

func (m *KpiEvent) GetTs() float32 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *KpiEvent) GetPrefixes() map[string]*MetricValuePairs {
	if m != nil {
		return m.Prefixes
	}
	return nil
}

type KpiEvent2 struct {
	// Type of KPI Event
	Type KpiEventType_KpiEventType `protobuf:"varint,1,opt,name=type,proto3,enum=voltha.KpiEventType_KpiEventType" json:"type,omitempty"`
	// Fields used when for slice:
	Ts                   float64              `protobuf:"fixed64,2,opt,name=ts,proto3" json:"ts,omitempty"`
	SliceData            []*MetricInformation `protobuf:"bytes,3,rep,name=slice_data,json=sliceData,proto3" json:"slice_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *KpiEvent2) Reset()         { *m = KpiEvent2{} }
func (m *KpiEvent2) String() string { return proto.CompactTextString(m) }
func (*KpiEvent2) ProtoMessage()    {}
func (*KpiEvent2) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_26477559e5dbfb38, []int{7}
}
func (m *KpiEvent2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KpiEvent2.Unmarshal(m, b)
}
func (m *KpiEvent2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KpiEvent2.Marshal(b, m, deterministic)
}
func (dst *KpiEvent2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KpiEvent2.Merge(dst, src)
}
func (m *KpiEvent2) XXX_Size() int {
	return xxx_messageInfo_KpiEvent2.Size(m)
}
func (m *KpiEvent2) XXX_DiscardUnknown() {
	xxx_messageInfo_KpiEvent2.DiscardUnknown(m)
}

var xxx_messageInfo_KpiEvent2 proto.InternalMessageInfo

func (m *KpiEvent2) GetType() KpiEventType_KpiEventType {
	if m != nil {
		return m.Type
	}
	return KpiEventType_slice
}

func (m *KpiEvent2) GetTs() float64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *KpiEvent2) GetSliceData() []*MetricInformation {
	if m != nil {
		return m.SliceData
	}
	return nil
}

//
// Identify to the area of the system impacted by the alarm
type AlarmEventType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmEventType) Reset()         { *m = AlarmEventType{} }
func (m *AlarmEventType) String() string { return proto.CompactTextString(m) }
func (*AlarmEventType) ProtoMessage()    {}
func (*AlarmEventType) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_26477559e5dbfb38, []int{8}
}
func (m *AlarmEventType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmEventType.Unmarshal(m, b)
}
func (m *AlarmEventType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmEventType.Marshal(b, m, deterministic)
}
func (dst *AlarmEventType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmEventType.Merge(dst, src)
}
func (m *AlarmEventType) XXX_Size() int {
	return xxx_messageInfo_AlarmEventType.Size(m)
}
func (m *AlarmEventType) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmEventType.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmEventType proto.InternalMessageInfo

//
// Identify to the functional category originating the alarm
type AlarmEventCategory struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmEventCategory) Reset()         { *m = AlarmEventCategory{} }
func (m *AlarmEventCategory) String() string { return proto.CompactTextString(m) }
func (*AlarmEventCategory) ProtoMessage()    {}
func (*AlarmEventCategory) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_26477559e5dbfb38, []int{9}
}
func (m *AlarmEventCategory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmEventCategory.Unmarshal(m, b)
}
func (m *AlarmEventCategory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmEventCategory.Marshal(b, m, deterministic)
}
func (dst *AlarmEventCategory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmEventCategory.Merge(dst, src)
}
func (m *AlarmEventCategory) XXX_Size() int {
	return xxx_messageInfo_AlarmEventCategory.Size(m)
}
func (m *AlarmEventCategory) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmEventCategory.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmEventCategory proto.InternalMessageInfo

//
// Active state of the alarm
type AlarmEventState struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmEventState) Reset()         { *m = AlarmEventState{} }
func (m *AlarmEventState) String() string { return proto.CompactTextString(m) }
func (*AlarmEventState) ProtoMessage()    {}
func (*AlarmEventState) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_26477559e5dbfb38, []int{10}
}
func (m *AlarmEventState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmEventState.Unmarshal(m, b)
}
func (m *AlarmEventState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmEventState.Marshal(b, m, deterministic)
}
func (dst *AlarmEventState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmEventState.Merge(dst, src)
}
func (m *AlarmEventState) XXX_Size() int {
	return xxx_messageInfo_AlarmEventState.Size(m)
}
func (m *AlarmEventState) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmEventState.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmEventState proto.InternalMessageInfo

//
// Identify the overall impact of the alarm on the system
type AlarmEventSeverity struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmEventSeverity) Reset()         { *m = AlarmEventSeverity{} }
func (m *AlarmEventSeverity) String() string { return proto.CompactTextString(m) }
func (*AlarmEventSeverity) ProtoMessage()    {}
func (*AlarmEventSeverity) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_26477559e5dbfb38, []int{11}
}
func (m *AlarmEventSeverity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmEventSeverity.Unmarshal(m, b)
}
func (m *AlarmEventSeverity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmEventSeverity.Marshal(b, m, deterministic)
}
func (dst *AlarmEventSeverity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmEventSeverity.Merge(dst, src)
}
func (m *AlarmEventSeverity) XXX_Size() int {
	return xxx_messageInfo_AlarmEventSeverity.Size(m)
}
func (m *AlarmEventSeverity) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmEventSeverity.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmEventSeverity proto.InternalMessageInfo

//
//
type AlarmEvent struct {
	// Unique ID for this alarm.  e.g. voltha.some_olt.1234
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Refers to the area of the system impacted by the alarm
	Type AlarmEventType_AlarmEventType `protobuf:"varint,2,opt,name=type,proto3,enum=voltha.AlarmEventType_AlarmEventType" json:"type,omitempty"`
	// Refers to functional category of the alarm
	Category AlarmEventCategory_AlarmEventCategory `protobuf:"varint,3,opt,name=category,proto3,enum=voltha.AlarmEventCategory_AlarmEventCategory" json:"category,omitempty"`
	// Current active state of the alarm
	State AlarmEventState_AlarmEventState `protobuf:"varint,4,opt,name=state,proto3,enum=voltha.AlarmEventState_AlarmEventState" json:"state,omitempty"`
	// Overall impact of the alarm on the system
	Severity AlarmEventSeverity_AlarmEventSeverity `protobuf:"varint,5,opt,name=severity,proto3,enum=voltha.AlarmEventSeverity_AlarmEventSeverity" json:"severity,omitempty"`
	// Timestamp at which the alarm was first raised
	RaisedTs float32 `protobuf:"fixed32,6,opt,name=raised_ts,json=raisedTs,proto3" json:"raised_ts,omitempty"`
	// Timestamp at which the alarm was reported
	ReportedTs float32 `protobuf:"fixed32,7,opt,name=reported_ts,json=reportedTs,proto3" json:"reported_ts,omitempty"`
	// Timestamp at which the alarm has changed since it was raised
	ChangedTs float32 `protobuf:"fixed32,8,opt,name=changed_ts,json=changedTs,proto3" json:"changed_ts,omitempty"`
	// Identifier of the originating resource of the alarm
	ResourceId string `protobuf:"bytes,9,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// Textual explanation of the alarm
	Description string `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`
	// Key/Value storage for extra information that may give context to the alarm
	Context map[string]string `protobuf:"bytes,11,rep,name=context,proto3" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// logical device id
	LogicalDeviceId string `protobuf:"bytes,12,opt,name=logical_device_id,json=logicalDeviceId,proto3" json:"logical_device_id,omitempty"`
	// alarm_type  name indicates clearly the name of the alarm
	AlarmTypeName        string   `protobuf:"bytes,13,opt,name=alarm_type_name,json=alarmTypeName,proto3" json:"alarm_type_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmEvent) Reset()         { *m = AlarmEvent{} }
func (m *AlarmEvent) String() string { return proto.CompactTextString(m) }
func (*AlarmEvent) ProtoMessage()    {}
func (*AlarmEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_26477559e5dbfb38, []int{12}
}
func (m *AlarmEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmEvent.Unmarshal(m, b)
}
func (m *AlarmEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmEvent.Marshal(b, m, deterministic)
}
func (dst *AlarmEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmEvent.Merge(dst, src)
}
func (m *AlarmEvent) XXX_Size() int {
	return xxx_messageInfo_AlarmEvent.Size(m)
}
func (m *AlarmEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmEvent proto.InternalMessageInfo

func (m *AlarmEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AlarmEvent) GetType() AlarmEventType_AlarmEventType {
	if m != nil {
		return m.Type
	}
	return AlarmEventType_COMMUNICATION
}

func (m *AlarmEvent) GetCategory() AlarmEventCategory_AlarmEventCategory {
	if m != nil {
		return m.Category
	}
	return AlarmEventCategory_PON
}

func (m *AlarmEvent) GetState() AlarmEventState_AlarmEventState {
	if m != nil {
		return m.State
	}
	return AlarmEventState_RAISED
}

func (m *AlarmEvent) GetSeverity() AlarmEventSeverity_AlarmEventSeverity {
	if m != nil {
		return m.Severity
	}
	return AlarmEventSeverity_INDETERMINATE
}

func (m *AlarmEvent) GetRaisedTs() float32 {
	if m != nil {
		return m.RaisedTs
	}
	return 0
}

func (m *AlarmEvent) GetReportedTs() float32 {
	if m != nil {
		return m.ReportedTs
	}
	return 0
}

func (m *AlarmEvent) GetChangedTs() float32 {
	if m != nil {
		return m.ChangedTs
	}
	return 0
}

func (m *AlarmEvent) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *AlarmEvent) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AlarmEvent) GetContext() map[string]string {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *AlarmEvent) GetLogicalDeviceId() string {
	if m != nil {
		return m.LogicalDeviceId
	}
	return ""
}

func (m *AlarmEvent) GetAlarmTypeName() string {
	if m != nil {
		return m.AlarmTypeName
	}
	return ""
}

func init() {
	proto.RegisterType((*ConfigEventType)(nil), "voltha.ConfigEventType")
	proto.RegisterType((*ConfigEvent)(nil), "voltha.ConfigEvent")
	proto.RegisterType((*KpiEventType)(nil), "voltha.KpiEventType")
	proto.RegisterType((*MetricMetaData)(nil), "voltha.MetricMetaData")
	proto.RegisterMapType((map[string]string)(nil), "voltha.MetricMetaData.ContextEntry")
	proto.RegisterType((*MetricValuePairs)(nil), "voltha.MetricValuePairs")
	proto.RegisterMapType((map[string]float32)(nil), "voltha.MetricValuePairs.MetricsEntry")
	proto.RegisterType((*MetricInformation)(nil), "voltha.MetricInformation")
	proto.RegisterMapType((map[string]float32)(nil), "voltha.MetricInformation.MetricsEntry")
	proto.RegisterType((*KpiEvent)(nil), "voltha.KpiEvent")
	proto.RegisterMapType((map[string]*MetricValuePairs)(nil), "voltha.KpiEvent.PrefixesEntry")
	proto.RegisterType((*KpiEvent2)(nil), "voltha.KpiEvent2")
	proto.RegisterType((*AlarmEventType)(nil), "voltha.AlarmEventType")
	proto.RegisterType((*AlarmEventCategory)(nil), "voltha.AlarmEventCategory")
	proto.RegisterType((*AlarmEventState)(nil), "voltha.AlarmEventState")
	proto.RegisterType((*AlarmEventSeverity)(nil), "voltha.AlarmEventSeverity")
	proto.RegisterType((*AlarmEvent)(nil), "voltha.AlarmEvent")
	proto.RegisterMapType((map[string]string)(nil), "voltha.AlarmEvent.ContextEntry")
	proto.RegisterEnum("voltha.ConfigEventType_ConfigEventType", ConfigEventType_ConfigEventType_name, ConfigEventType_ConfigEventType_value)
	proto.RegisterEnum("voltha.KpiEventType_KpiEventType", KpiEventType_KpiEventType_name, KpiEventType_KpiEventType_value)
	proto.RegisterEnum("voltha.AlarmEventType_AlarmEventType", AlarmEventType_AlarmEventType_name, AlarmEventType_AlarmEventType_value)
	proto.RegisterEnum("voltha.AlarmEventCategory_AlarmEventCategory", AlarmEventCategory_AlarmEventCategory_name, AlarmEventCategory_AlarmEventCategory_value)
	proto.RegisterEnum("voltha.AlarmEventState_AlarmEventState", AlarmEventState_AlarmEventState_name, AlarmEventState_AlarmEventState_value)
	proto.RegisterEnum("voltha.AlarmEventSeverity_AlarmEventSeverity", AlarmEventSeverity_AlarmEventSeverity_name, AlarmEventSeverity_AlarmEventSeverity_value)
}

func init() { proto.RegisterFile("voltha_protos/events.proto", fileDescriptor_events_26477559e5dbfb38) }

var fileDescriptor_events_26477559e5dbfb38 = []byte{
	// 1033 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xae, 0x9d, 0xff, 0x93, 0x36, 0xf5, 0x8e, 0x10, 0x32, 0xe5, 0x67, 0xbb, 0x46, 0x2c, 0xd5,
	0xa2, 0xa6, 0x22, 0x08, 0x69, 0xb7, 0xa8, 0x82, 0xe0, 0x5a, 0xc8, 0xd0, 0x38, 0x65, 0x92, 0x76,
	0x81, 0x9b, 0x68, 0xd6, 0x9e, 0xa6, 0x16, 0x89, 0xc7, 0xb2, 0xa7, 0xd1, 0xe6, 0x8e, 0x07, 0x40,
	0xe2, 0x72, 0x9f, 0x88, 0x37, 0xe1, 0x41, 0xd0, 0xcc, 0xd8, 0xb5, 0x93, 0x7a, 0xc5, 0xc5, 0x8a,
	0xbb, 0x39, 0xdf, 0xf9, 0xf1, 0x77, 0xce, 0x9c, 0x39, 0xc7, 0x70, 0xb0, 0x62, 0x0b, 0x7e, 0x4b,
	0x66, 0x71, 0xc2, 0x38, 0x4b, 0x4f, 0xe8, 0x8a, 0x46, 0x3c, 0xed, 0x4b, 0x09, 0x35, 0x95, 0xee,
	0xc0, 0xdc, 0xb4, 0x59, 0x52, 0x4e, 0x94, 0xc5, 0xc1, 0x47, 0x73, 0xc6, 0xe6, 0x0b, 0x7a, 0x42,
	0xe2, 0xf0, 0x84, 0x44, 0x11, 0xe3, 0x84, 0x87, 0x2c, 0xca, 0xfc, 0x2d, 0x07, 0xf6, 0x6d, 0x16,
	0xdd, 0x84, 0x73, 0x47, 0x44, 0x9d, 0xae, 0x63, 0x6a, 0x0d, 0x1e, 0x40, 0xa8, 0x05, 0x35, 0x12,
	0x04, 0xc6, 0x0e, 0x02, 0x68, 0x26, 0x74, 0xc9, 0x56, 0xd4, 0xd0, 0xc4, 0xf9, 0x2e, 0x0e, 0x08,
	0xa7, 0x86, 0x6e, 0x25, 0xd0, 0x2d, 0xf9, 0xa0, 0x6f, 0xa0, 0xce, 0xd7, 0x31, 0x35, 0xb5, 0x43,
	0xed, 0xa8, 0x37, 0xf8, 0xbc, 0xaf, 0xc8, 0xf5, 0xb7, 0xc2, 0x6e, 0xcb, 0x58, 0x3a, 0x21, 0x04,
	0xf5, 0x5b, 0x92, 0xde, 0x9a, 0xfa, 0xa1, 0x76, 0xd4, 0xc1, 0xf2, 0x2c, 0xb0, 0x80, 0x70, 0x62,
	0xd6, 0x14, 0x26, 0xce, 0xd6, 0x97, 0xb0, 0xfb, 0x53, 0x1c, 0x16, 0xbc, 0x9f, 0x6c, 0xca, 0xa8,
	0x03, 0x8d, 0x74, 0x11, 0xfa, 0xd4, 0xd8, 0x41, 0x4d, 0xd0, 0x79, 0x6a, 0x68, 0xd6, 0x1b, 0x1d,
	0x7a, 0x23, 0xca, 0x93, 0xd0, 0x1f, 0x51, 0x4e, 0xce, 0x09, 0x27, 0xe8, 0x3d, 0x68, 0xf0, 0x90,
	0x2f, 0x14, 0xd7, 0x0e, 0x56, 0x02, 0xea, 0x09, 0x07, 0xc9, 0x40, 0xc3, 0x3a, 0x4f, 0xd1, 0x33,
	0x78, 0xb4, 0x60, 0xf3, 0xd0, 0x27, 0x8b, 0x59, 0x40, 0x57, 0xa1, 0x4f, 0x67, 0x61, 0x90, 0x91,
	0xd9, 0xcf, 0x14, 0xe7, 0x12, 0x77, 0x03, 0xf4, 0x21, 0x74, 0x52, 0x9a, 0x84, 0x64, 0x31, 0x8b,
	0x98, 0x59, 0x97, 0x36, 0x6d, 0x05, 0x78, 0x4c, 0x28, 0x8b, 0x00, 0x0d, 0xa5, 0x0c, 0x72, 0xcf,
	0x33, 0x68, 0xf9, 0x2c, 0xe2, 0xf4, 0x35, 0x37, 0x9b, 0x87, 0xb5, 0xa3, 0xee, 0xe0, 0xd3, 0xbc,
	0x72, 0x9b, 0xa4, 0x45, 0xe1, 0x84, 0x95, 0x13, 0xf1, 0x64, 0x8d, 0x73, 0x9f, 0x83, 0x53, 0xd8,
	0x2d, 0x2b, 0x90, 0x01, 0xb5, 0xdf, 0xe9, 0x3a, 0x4b, 0x4c, 0x1c, 0x45, 0xb2, 0x2b, 0xb2, 0xb8,
	0xa3, 0x59, 0x6d, 0x95, 0x70, 0xaa, 0x3f, 0xd7, 0xac, 0xbf, 0x34, 0x30, 0xd4, 0x47, 0xae, 0x05,
	0x76, 0x49, 0xc2, 0x24, 0x45, 0xdf, 0x42, 0x6b, 0x29, 0xb1, 0xd4, 0xd4, 0x24, 0x9f, 0xcf, 0x36,
	0xf9, 0x14, 0xa6, 0x19, 0x90, 0x66, 0x8c, 0x32, 0x2f, 0xc1, 0xa8, 0xac, 0xf8, 0x2f, 0x46, 0x7a,
	0x99, 0xd1, 0xdf, 0x1a, 0x3c, 0x52, 0xce, 0x6e, 0x74, 0xc3, 0x92, 0xa5, 0x6c, 0x5b, 0x34, 0x80,
	0xb6, 0xe8, 0x6d, 0xd9, 0x0c, 0x22, 0x4c, 0x77, 0xf0, 0x7e, 0x75, 0x8d, 0xf0, 0xbd, 0x1d, 0xfa,
	0xae, 0x48, 0x43, 0x97, 0x69, 0x3c, 0xdd, 0x74, 0x29, 0xc5, 0xff, 0x1f, 0xf2, 0xf8, 0x47, 0x83,
	0x76, 0xde, 0x97, 0xe8, 0xeb, 0x8d, 0x87, 0xf1, 0x24, 0xe7, 0x51, 0xee, 0xdb, 0x0d, 0x21, 0x7b,
	0x12, 0x45, 0x3b, 0xea, 0xb2, 0x1d, 0x4f, 0xa1, 0x1d, 0x27, 0xf4, 0x26, 0x7c, 0x4d, 0x53, 0xb3,
	0x26, 0x53, 0xfa, 0x64, 0x3b, 0x54, 0xff, 0x32, 0x33, 0x50, 0xa9, 0xdc, 0xdb, 0x1f, 0x5c, 0xc1,
	0xde, 0x86, 0xaa, 0x22, 0x99, 0x7e, 0x39, 0x99, 0xee, 0xc0, 0x7c, 0xdb, 0xad, 0x97, 0xd3, 0xfc,
	0x53, 0x83, 0x4e, 0xfe, 0xed, 0xc1, 0xbb, 0xe7, 0xa9, 0x9e, 0xdd, 0x73, 0x00, 0xf9, 0x84, 0x67,
	0xd9, 0xe3, 0x17, 0x99, 0x7e, 0xf0, 0xd6, 0xcb, 0xc3, 0x1d, 0x69, 0x2c, 0x6e, 0xdf, 0xfa, 0x43,
	0x83, 0xde, 0x70, 0x41, 0x92, 0x65, 0x31, 0x1f, 0xa2, 0x6d, 0x04, 0x3d, 0x82, 0x3d, 0x7b, 0x3c,
	0x1a, 0x5d, 0x79, 0xae, 0x3d, 0x9c, 0xba, 0x63, 0xcf, 0xd8, 0x41, 0xfb, 0xd0, 0x75, 0xbc, 0x6b,
	0x17, 0x8f, 0xbd, 0x91, 0xe3, 0x4d, 0x0d, 0x0d, 0xed, 0x41, 0xc7, 0xf9, 0xf9, 0xca, 0xbd, 0x94,
	0xa2, 0x8e, 0xba, 0xd0, 0x9a, 0x38, 0xf8, 0xda, 0xb5, 0x1d, 0xa3, 0x86, 0x7a, 0x00, 0x97, 0x78,
	0x6c, 0x3b, 0x93, 0x89, 0xeb, 0xfd, 0x60, 0xd4, 0xd1, 0x2e, 0xb4, 0x27, 0x8e, 0x7d, 0x85, 0xdd,
	0xe9, 0xaf, 0x46, 0xc3, 0x7a, 0x09, 0xa8, 0xf8, 0x9e, 0x4d, 0x38, 0x9d, 0xb3, 0x64, 0x6d, 0x0d,
	0xab, 0x50, 0x31, 0x60, 0x2f, 0xe5, 0xf7, 0x5b, 0x50, 0x1b, 0x5f, 0x88, 0xef, 0x8a, 0x83, 0xfc,
	0xa2, 0x3c, 0x5c, 0x19, 0x35, 0x71, 0xf0, 0x3c, 0xd7, 0xa8, 0x5b, 0x67, 0xb0, 0x5f, 0x84, 0x98,
	0x70, 0xc2, 0xa9, 0xf5, 0xec, 0x01, 0x24, 0xc6, 0x33, 0x1e, 0xba, 0x13, 0xe7, 0xdc, 0xd8, 0x11,
	0xac, 0xed, 0x0b, 0x67, 0x88, 0x9d, 0x73, 0x43, 0xb3, 0xa2, 0x32, 0x83, 0x09, 0x5d, 0xd1, 0x24,
	0xe4, 0x6b, 0xeb, 0x97, 0x2a, 0x54, 0x54, 0xc8, 0xf5, 0xce, 0x9d, 0xa9, 0x83, 0x47, 0xae, 0x37,
	0x9c, 0x3a, 0x2a, 0xd6, 0xcb, 0x21, 0xf6, 0x44, 0xc6, 0x9a, 0x98, 0xb1, 0x23, 0xd7, 0x1b, 0x63,
	0x43, 0x97, 0xc7, 0xe1, 0x8f, 0x63, 0x6c, 0xd4, 0x44, 0x1d, 0x6c, 0xec, 0x4e, 0x5d, 0x7b, 0x78,
	0x61, 0xd4, 0xad, 0x37, 0x0d, 0x80, 0x22, 0xb4, 0xb8, 0xe3, 0x30, 0xc8, 0xba, 0x4d, 0x0f, 0x03,
	0xf4, 0x22, 0x6b, 0x15, 0x5d, 0xb6, 0xca, 0xfd, 0x84, 0xd9, 0xbc, 0xaa, 0x2d, 0x31, 0x6b, 0x17,
	0x17, 0xda, 0x7e, 0x56, 0x41, 0x39, 0x8c, 0x7b, 0x83, 0xe3, 0x87, 0xee, 0x79, 0x8d, 0x2b, 0x20,
	0x7c, 0xef, 0x8e, 0xce, 0xa0, 0x91, 0x8a, 0xb2, 0xc9, 0x81, 0x5d, 0x5a, 0x59, 0x5b, 0x55, 0xdd,
	0x96, 0xb1, 0xf2, 0x12, 0x4c, 0xd2, 0xac, 0x66, 0x72, 0xaa, 0x57, 0x32, 0xc9, 0xab, 0x5a, 0x01,
	0xe1, 0x7b, 0x77, 0xb1, 0x21, 0x12, 0x12, 0xa6, 0x34, 0x98, 0xf1, 0xd4, 0x6c, 0xca, 0x27, 0xdf,
	0x56, 0xc0, 0x34, 0x45, 0x8f, 0xa1, 0x9b, 0xd0, 0x98, 0x25, 0x5c, 0xa9, 0x5b, 0x52, 0x0d, 0x39,
	0x34, 0x4d, 0xd1, 0xc7, 0x00, 0xfe, 0x2d, 0x89, 0xe6, 0x4a, 0xdf, 0x96, 0xfa, 0x4e, 0x86, 0xe4,
	0xfe, 0x29, 0xbb, 0x4b, 0xd4, 0x02, 0xea, 0xc8, 0x5b, 0x80, 0x1c, 0x72, 0x03, 0x74, 0x08, 0xdd,
	0x80, 0xa6, 0x7e, 0x12, 0xc6, 0xe2, 0x45, 0x99, 0x20, 0x0d, 0xca, 0x10, 0x7a, 0x51, 0x2c, 0xa9,
	0xae, 0x7c, 0x90, 0x8f, 0x1f, 0x66, 0x5a, 0xbd, 0xa0, 0xaa, 0xb7, 0xe8, 0x6e, 0xf5, 0x16, 0x7d,
	0x0a, 0xfb, 0x44, 0xc4, 0x9b, 0x89, 0x9b, 0x9e, 0x45, 0x64, 0x49, 0xcd, 0x3d, 0x69, 0xb9, 0x27,
	0x61, 0xd1, 0x05, 0x1e, 0x59, 0xd2, 0x77, 0x59, 0x7a, 0xdf, 0x1f, 0xff, 0xf6, 0xc5, 0x3c, 0xe4,
	0xb7, 0x77, 0xaf, 0xfa, 0x3e, 0x5b, 0x9e, 0xb0, 0x98, 0x46, 0x3e, 0x4b, 0x82, 0x13, 0x95, 0xce,
	0x71, 0xf6, 0x2b, 0x35, 0x67, 0x19, 0xf0, 0xaa, 0x29, 0x91, 0xaf, 0xfe, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xba, 0xd4, 0xc8, 0xc8, 0x90, 0x09, 0x00, 0x00,
}