// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pulsar.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pulsar.proto

It has these top-level messages:
	Status
	KeyValuePair
	VectorRowRecord
	AttrRecord
	VectorRecord
	VectorParam
	FieldValue
	Entities
	InsertParam
	EntityIds
	DeleteByIDParam
	SearchParam
	QueryResult
*/
package pb

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

type ErrorCode int32

const (
	ErrorCode_SUCCESS                 ErrorCode = 0
	ErrorCode_UNEXPECTED_ERROR        ErrorCode = 1
	ErrorCode_CONNECT_FAILED          ErrorCode = 2
	ErrorCode_PERMISSION_DENIED       ErrorCode = 3
	ErrorCode_COLLECTION_NOT_EXISTS   ErrorCode = 4
	ErrorCode_ILLEGAL_ARGUMENT        ErrorCode = 5
	ErrorCode_ILLEGAL_DIMENSION       ErrorCode = 7
	ErrorCode_ILLEGAL_INDEX_TYPE      ErrorCode = 8
	ErrorCode_ILLEGAL_COLLECTION_NAME ErrorCode = 9
	ErrorCode_ILLEGAL_TOPK            ErrorCode = 10
	ErrorCode_ILLEGAL_ROWRECORD       ErrorCode = 11
	ErrorCode_ILLEGAL_VECTOR_ID       ErrorCode = 12
	ErrorCode_ILLEGAL_SEARCH_RESULT   ErrorCode = 13
	ErrorCode_FILE_NOT_FOUND          ErrorCode = 14
	ErrorCode_META_FAILED             ErrorCode = 15
	ErrorCode_CACHE_FAILED            ErrorCode = 16
	ErrorCode_CANNOT_CREATE_FOLDER    ErrorCode = 17
	ErrorCode_CANNOT_CREATE_FILE      ErrorCode = 18
	ErrorCode_CANNOT_DELETE_FOLDER    ErrorCode = 19
	ErrorCode_CANNOT_DELETE_FILE      ErrorCode = 20
	ErrorCode_BUILD_INDEX_ERROR       ErrorCode = 21
	ErrorCode_ILLEGAL_NLIST           ErrorCode = 22
	ErrorCode_ILLEGAL_METRIC_TYPE     ErrorCode = 23
	ErrorCode_OUT_OF_MEMORY           ErrorCode = 24
)

var ErrorCode_name = map[int32]string{
	0:  "SUCCESS",
	1:  "UNEXPECTED_ERROR",
	2:  "CONNECT_FAILED",
	3:  "PERMISSION_DENIED",
	4:  "COLLECTION_NOT_EXISTS",
	5:  "ILLEGAL_ARGUMENT",
	7:  "ILLEGAL_DIMENSION",
	8:  "ILLEGAL_INDEX_TYPE",
	9:  "ILLEGAL_COLLECTION_NAME",
	10: "ILLEGAL_TOPK",
	11: "ILLEGAL_ROWRECORD",
	12: "ILLEGAL_VECTOR_ID",
	13: "ILLEGAL_SEARCH_RESULT",
	14: "FILE_NOT_FOUND",
	15: "META_FAILED",
	16: "CACHE_FAILED",
	17: "CANNOT_CREATE_FOLDER",
	18: "CANNOT_CREATE_FILE",
	19: "CANNOT_DELETE_FOLDER",
	20: "CANNOT_DELETE_FILE",
	21: "BUILD_INDEX_ERROR",
	22: "ILLEGAL_NLIST",
	23: "ILLEGAL_METRIC_TYPE",
	24: "OUT_OF_MEMORY",
}
var ErrorCode_value = map[string]int32{
	"SUCCESS":                 0,
	"UNEXPECTED_ERROR":        1,
	"CONNECT_FAILED":          2,
	"PERMISSION_DENIED":       3,
	"COLLECTION_NOT_EXISTS":   4,
	"ILLEGAL_ARGUMENT":        5,
	"ILLEGAL_DIMENSION":       7,
	"ILLEGAL_INDEX_TYPE":      8,
	"ILLEGAL_COLLECTION_NAME": 9,
	"ILLEGAL_TOPK":            10,
	"ILLEGAL_ROWRECORD":       11,
	"ILLEGAL_VECTOR_ID":       12,
	"ILLEGAL_SEARCH_RESULT":   13,
	"FILE_NOT_FOUND":          14,
	"META_FAILED":             15,
	"CACHE_FAILED":            16,
	"CANNOT_CREATE_FOLDER":    17,
	"CANNOT_CREATE_FILE":      18,
	"CANNOT_DELETE_FOLDER":    19,
	"CANNOT_DELETE_FILE":      20,
	"BUILD_INDEX_ERROR":       21,
	"ILLEGAL_NLIST":           22,
	"ILLEGAL_METRIC_TYPE":     23,
	"OUT_OF_MEMORY":           24,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DataType int32

const (
	DataType_NONE          DataType = 0
	DataType_BOOL          DataType = 1
	DataType_INT8          DataType = 2
	DataType_INT16         DataType = 3
	DataType_INT32         DataType = 4
	DataType_INT64         DataType = 5
	DataType_FLOAT         DataType = 10
	DataType_DOUBLE        DataType = 11
	DataType_STRING        DataType = 20
	DataType_VECTOR_BINARY DataType = 100
	DataType_VECTOR_FLOAT  DataType = 101
)

var DataType_name = map[int32]string{
	0:   "NONE",
	1:   "BOOL",
	2:   "INT8",
	3:   "INT16",
	4:   "INT32",
	5:   "INT64",
	10:  "FLOAT",
	11:  "DOUBLE",
	20:  "STRING",
	100: "VECTOR_BINARY",
	101: "VECTOR_FLOAT",
}
var DataType_value = map[string]int32{
	"NONE":          0,
	"BOOL":          1,
	"INT8":          2,
	"INT16":         3,
	"INT32":         4,
	"INT64":         5,
	"FLOAT":         10,
	"DOUBLE":        11,
	"STRING":        20,
	"VECTOR_BINARY": 100,
	"VECTOR_FLOAT":  101,
}

func (x DataType) String() string {
	return proto.EnumName(DataType_name, int32(x))
}
func (DataType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Status struct {
	ErrorCode ErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,enum=pb.ErrorCode" json:"error_code,omitempty"`
	Reason    string    `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Status) GetErrorCode() ErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return ErrorCode_SUCCESS
}

func (m *Status) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type KeyValuePair struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *KeyValuePair) Reset()                    { *m = KeyValuePair{} }
func (m *KeyValuePair) String() string            { return proto.CompactTextString(m) }
func (*KeyValuePair) ProtoMessage()               {}
func (*KeyValuePair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *KeyValuePair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValuePair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type VectorRowRecord struct {
	FloatData  []float32 `protobuf:"fixed32,1,rep,packed,name=float_data,json=floatData" json:"float_data,omitempty"`
	BinaryData []byte    `protobuf:"bytes,2,opt,name=binary_data,json=binaryData,proto3" json:"binary_data,omitempty"`
}

func (m *VectorRowRecord) Reset()                    { *m = VectorRowRecord{} }
func (m *VectorRowRecord) String() string            { return proto.CompactTextString(m) }
func (*VectorRowRecord) ProtoMessage()               {}
func (*VectorRowRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *VectorRowRecord) GetFloatData() []float32 {
	if m != nil {
		return m.FloatData
	}
	return nil
}

func (m *VectorRowRecord) GetBinaryData() []byte {
	if m != nil {
		return m.BinaryData
	}
	return nil
}

type AttrRecord struct {
	Int32Value  []int32   `protobuf:"varint,1,rep,packed,name=int32_value,json=int32Value" json:"int32_value,omitempty"`
	Int64Value  []int64   `protobuf:"varint,2,rep,packed,name=int64_value,json=int64Value" json:"int64_value,omitempty"`
	FloatValue  []float32 `protobuf:"fixed32,3,rep,packed,name=float_value,json=floatValue" json:"float_value,omitempty"`
	DoubleValue []float64 `protobuf:"fixed64,4,rep,packed,name=double_value,json=doubleValue" json:"double_value,omitempty"`
}

func (m *AttrRecord) Reset()                    { *m = AttrRecord{} }
func (m *AttrRecord) String() string            { return proto.CompactTextString(m) }
func (*AttrRecord) ProtoMessage()               {}
func (*AttrRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AttrRecord) GetInt32Value() []int32 {
	if m != nil {
		return m.Int32Value
	}
	return nil
}

func (m *AttrRecord) GetInt64Value() []int64 {
	if m != nil {
		return m.Int64Value
	}
	return nil
}

func (m *AttrRecord) GetFloatValue() []float32 {
	if m != nil {
		return m.FloatValue
	}
	return nil
}

func (m *AttrRecord) GetDoubleValue() []float64 {
	if m != nil {
		return m.DoubleValue
	}
	return nil
}

type VectorRecord struct {
	Records []*VectorRowRecord `protobuf:"bytes,1,rep,name=records" json:"records,omitempty"`
}

func (m *VectorRecord) Reset()                    { *m = VectorRecord{} }
func (m *VectorRecord) String() string            { return proto.CompactTextString(m) }
func (*VectorRecord) ProtoMessage()               {}
func (*VectorRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *VectorRecord) GetRecords() []*VectorRowRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

type VectorParam struct {
	Json      string        `protobuf:"bytes,1,opt,name=json" json:"json,omitempty"`
	RowRecord *VectorRecord `protobuf:"bytes,2,opt,name=row_record,json=rowRecord" json:"row_record,omitempty"`
}

func (m *VectorParam) Reset()                    { *m = VectorParam{} }
func (m *VectorParam) String() string            { return proto.CompactTextString(m) }
func (*VectorParam) ProtoMessage()               {}
func (*VectorParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *VectorParam) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

func (m *VectorParam) GetRowRecord() *VectorRecord {
	if m != nil {
		return m.RowRecord
	}
	return nil
}

type FieldValue struct {
	FieldName    string        `protobuf:"bytes,1,opt,name=field_name,json=fieldName" json:"field_name,omitempty"`
	Type         DataType      `protobuf:"varint,2,opt,name=type,enum=pb.DataType" json:"type,omitempty"`
	AttrRecord   *AttrRecord   `protobuf:"bytes,3,opt,name=attr_record,json=attrRecord" json:"attr_record,omitempty"`
	VectorRecord *VectorRecord `protobuf:"bytes,4,opt,name=vector_record,json=vectorRecord" json:"vector_record,omitempty"`
}

func (m *FieldValue) Reset()                    { *m = FieldValue{} }
func (m *FieldValue) String() string            { return proto.CompactTextString(m) }
func (*FieldValue) ProtoMessage()               {}
func (*FieldValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *FieldValue) GetFieldName() string {
	if m != nil {
		return m.FieldName
	}
	return ""
}

func (m *FieldValue) GetType() DataType {
	if m != nil {
		return m.Type
	}
	return DataType_NONE
}

func (m *FieldValue) GetAttrRecord() *AttrRecord {
	if m != nil {
		return m.AttrRecord
	}
	return nil
}

func (m *FieldValue) GetVectorRecord() *VectorRecord {
	if m != nil {
		return m.VectorRecord
	}
	return nil
}

type Entities struct {
	Status   *Status       `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Ids      []int64       `protobuf:"varint,2,rep,packed,name=ids" json:"ids,omitempty"`
	ValidRow []bool        `protobuf:"varint,3,rep,packed,name=valid_row,json=validRow" json:"valid_row,omitempty"`
	Fields   []*FieldValue `protobuf:"bytes,4,rep,name=fields" json:"fields,omitempty"`
}

func (m *Entities) Reset()                    { *m = Entities{} }
func (m *Entities) String() string            { return proto.CompactTextString(m) }
func (*Entities) ProtoMessage()               {}
func (*Entities) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Entities) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Entities) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *Entities) GetValidRow() []bool {
	if m != nil {
		return m.ValidRow
	}
	return nil
}

func (m *Entities) GetFields() []*FieldValue {
	if m != nil {
		return m.Fields
	}
	return nil
}

type InsertParam struct {
	CollectionName     string          `protobuf:"bytes,1,opt,name=collection_name,json=collectionName" json:"collection_name,omitempty"`
	Fields             []*FieldValue   `protobuf:"bytes,2,rep,name=fields" json:"fields,omitempty"`
	EntityIdArray      []int64         `protobuf:"varint,3,rep,packed,name=entity_id_array,json=entityIdArray" json:"entity_id_array,omitempty"`
	PartitionTag       string          `protobuf:"bytes,4,opt,name=partition_tag,json=partitionTag" json:"partition_tag,omitempty"`
	Timestamp          []int64         `protobuf:"varint,5,rep,packed,name=timestamp" json:"timestamp,omitempty"`
	GrpcServerClientId int64           `protobuf:"varint,6,opt,name=grpc_server_client_id,json=grpcServerClientId" json:"grpc_server_client_id,omitempty"`
	ExtraParams        []*KeyValuePair `protobuf:"bytes,7,rep,name=extra_params,json=extraParams" json:"extra_params,omitempty"`
}

func (m *InsertParam) Reset()                    { *m = InsertParam{} }
func (m *InsertParam) String() string            { return proto.CompactTextString(m) }
func (*InsertParam) ProtoMessage()               {}
func (*InsertParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *InsertParam) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *InsertParam) GetFields() []*FieldValue {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *InsertParam) GetEntityIdArray() []int64 {
	if m != nil {
		return m.EntityIdArray
	}
	return nil
}

func (m *InsertParam) GetPartitionTag() string {
	if m != nil {
		return m.PartitionTag
	}
	return ""
}

func (m *InsertParam) GetTimestamp() []int64 {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *InsertParam) GetGrpcServerClientId() int64 {
	if m != nil {
		return m.GrpcServerClientId
	}
	return 0
}

func (m *InsertParam) GetExtraParams() []*KeyValuePair {
	if m != nil {
		return m.ExtraParams
	}
	return nil
}

type EntityIds struct {
	Status        *Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	EntityIdArray []int64 `protobuf:"varint,2,rep,packed,name=entity_id_array,json=entityIdArray" json:"entity_id_array,omitempty"`
}

func (m *EntityIds) Reset()                    { *m = EntityIds{} }
func (m *EntityIds) String() string            { return proto.CompactTextString(m) }
func (*EntityIds) ProtoMessage()               {}
func (*EntityIds) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *EntityIds) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *EntityIds) GetEntityIdArray() []int64 {
	if m != nil {
		return m.EntityIdArray
	}
	return nil
}

type DeleteByIDParam struct {
	CollectionName     string  `protobuf:"bytes,1,opt,name=collection_name,json=collectionName" json:"collection_name,omitempty"`
	IdArray            []int64 `protobuf:"varint,2,rep,packed,name=id_array,json=idArray" json:"id_array,omitempty"`
	Timestamp          []int64 `protobuf:"varint,3,rep,packed,name=timestamp" json:"timestamp,omitempty"`
	GrpcServerClientId int64   `protobuf:"varint,4,opt,name=grpc_server_client_id,json=grpcServerClientId" json:"grpc_server_client_id,omitempty"`
}

func (m *DeleteByIDParam) Reset()                    { *m = DeleteByIDParam{} }
func (m *DeleteByIDParam) String() string            { return proto.CompactTextString(m) }
func (*DeleteByIDParam) ProtoMessage()               {}
func (*DeleteByIDParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DeleteByIDParam) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *DeleteByIDParam) GetIdArray() []int64 {
	if m != nil {
		return m.IdArray
	}
	return nil
}

func (m *DeleteByIDParam) GetTimestamp() []int64 {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *DeleteByIDParam) GetGrpcServerClientId() int64 {
	if m != nil {
		return m.GrpcServerClientId
	}
	return 0
}

type SearchParam struct {
	CollectionName     string          `protobuf:"bytes,1,opt,name=collection_name,json=collectionName" json:"collection_name,omitempty"`
	PartitionTagArray  []string        `protobuf:"bytes,2,rep,name=partition_tag_array,json=partitionTagArray" json:"partition_tag_array,omitempty"`
	VectorParam        []*VectorParam  `protobuf:"bytes,3,rep,name=vector_param,json=vectorParam" json:"vector_param,omitempty"`
	Dsl                string          `protobuf:"bytes,4,opt,name=dsl" json:"dsl,omitempty"`
	Timestamp          []int64         `protobuf:"varint,5,rep,packed,name=timestamp" json:"timestamp,omitempty"`
	GrpcServerClientId int64           `protobuf:"varint,6,opt,name=grpc_server_client_id,json=grpcServerClientId" json:"grpc_server_client_id,omitempty"`
	ExtraParams        []*KeyValuePair `protobuf:"bytes,7,rep,name=extra_params,json=extraParams" json:"extra_params,omitempty"`
}

func (m *SearchParam) Reset()                    { *m = SearchParam{} }
func (m *SearchParam) String() string            { return proto.CompactTextString(m) }
func (*SearchParam) ProtoMessage()               {}
func (*SearchParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *SearchParam) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *SearchParam) GetPartitionTagArray() []string {
	if m != nil {
		return m.PartitionTagArray
	}
	return nil
}

func (m *SearchParam) GetVectorParam() []*VectorParam {
	if m != nil {
		return m.VectorParam
	}
	return nil
}

func (m *SearchParam) GetDsl() string {
	if m != nil {
		return m.Dsl
	}
	return ""
}

func (m *SearchParam) GetTimestamp() []int64 {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *SearchParam) GetGrpcServerClientId() int64 {
	if m != nil {
		return m.GrpcServerClientId
	}
	return 0
}

func (m *SearchParam) GetExtraParams() []*KeyValuePair {
	if m != nil {
		return m.ExtraParams
	}
	return nil
}

type QueryResult struct {
	Status      *Status         `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Entities    *Entities       `protobuf:"bytes,2,opt,name=entities" json:"entities,omitempty"`
	RowNum      int64           `protobuf:"varint,3,opt,name=row_num,json=rowNum" json:"row_num,omitempty"`
	Scores      []float32       `protobuf:"fixed32,4,rep,packed,name=scores" json:"scores,omitempty"`
	Distances   []float32       `protobuf:"fixed32,5,rep,packed,name=distances" json:"distances,omitempty"`
	ExtraParams []*KeyValuePair `protobuf:"bytes,6,rep,name=extra_params,json=extraParams" json:"extra_params,omitempty"`
}

func (m *QueryResult) Reset()                    { *m = QueryResult{} }
func (m *QueryResult) String() string            { return proto.CompactTextString(m) }
func (*QueryResult) ProtoMessage()               {}
func (*QueryResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *QueryResult) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *QueryResult) GetEntities() *Entities {
	if m != nil {
		return m.Entities
	}
	return nil
}

func (m *QueryResult) GetRowNum() int64 {
	if m != nil {
		return m.RowNum
	}
	return 0
}

func (m *QueryResult) GetScores() []float32 {
	if m != nil {
		return m.Scores
	}
	return nil
}

func (m *QueryResult) GetDistances() []float32 {
	if m != nil {
		return m.Distances
	}
	return nil
}

func (m *QueryResult) GetExtraParams() []*KeyValuePair {
	if m != nil {
		return m.ExtraParams
	}
	return nil
}

func init() {
	proto.RegisterType((*Status)(nil), "pb.Status")
	proto.RegisterType((*KeyValuePair)(nil), "pb.KeyValuePair")
	proto.RegisterType((*VectorRowRecord)(nil), "pb.VectorRowRecord")
	proto.RegisterType((*AttrRecord)(nil), "pb.AttrRecord")
	proto.RegisterType((*VectorRecord)(nil), "pb.VectorRecord")
	proto.RegisterType((*VectorParam)(nil), "pb.VectorParam")
	proto.RegisterType((*FieldValue)(nil), "pb.FieldValue")
	proto.RegisterType((*Entities)(nil), "pb.Entities")
	proto.RegisterType((*InsertParam)(nil), "pb.InsertParam")
	proto.RegisterType((*EntityIds)(nil), "pb.EntityIds")
	proto.RegisterType((*DeleteByIDParam)(nil), "pb.DeleteByIDParam")
	proto.RegisterType((*SearchParam)(nil), "pb.SearchParam")
	proto.RegisterType((*QueryResult)(nil), "pb.QueryResult")
	proto.RegisterEnum("pb.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterEnum("pb.DataType", DataType_name, DataType_value)
}

func init() { proto.RegisterFile("pulsar.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xdd, 0x72, 0xdb, 0x44,
	0x14, 0xae, 0x2c, 0xc7, 0xb1, 0x8e, 0x9c, 0x64, 0xb3, 0x49, 0x1a, 0x77, 0x0a, 0xd3, 0x60, 0x66,
	0x8a, 0xa7, 0x03, 0xe9, 0xd4, 0x29, 0x19, 0x6e, 0xb8, 0x70, 0xa4, 0x75, 0xab, 0xa9, 0x2c, 0xa5,
	0x2b, 0xb9, 0x3f, 0x57, 0x9a, 0x8d, 0xb5, 0x14, 0x81, 0x6d, 0x79, 0x56, 0x72, 0x82, 0x1f, 0x80,
	0x07, 0x80, 0x47, 0xe0, 0x82, 0x7b, 0x98, 0xe1, 0x79, 0x78, 0x15, 0x66, 0x57, 0x52, 0xec, 0x96,
	0x96, 0x49, 0xef, 0xb8, 0x3b, 0xfb, 0x9d, 0x9f, 0xfd, 0xce, 0x77, 0xd6, 0x47, 0x86, 0xd6, 0x7c,
	0x31, 0xc9, 0x98, 0x38, 0x9e, 0x8b, 0x34, 0x4f, 0x71, 0x6d, 0x7e, 0xd1, 0xf1, 0xa0, 0x11, 0xe4,
	0x2c, 0x5f, 0x64, 0xf8, 0x4b, 0x00, 0x2e, 0x44, 0x2a, 0xa2, 0x71, 0x1a, 0xf3, 0xb6, 0x76, 0xa4,
	0x75, 0xb7, 0x7b, 0x5b, 0xc7, 0xf3, 0x8b, 0x63, 0x22, 0x51, 0x2b, 0x8d, 0x39, 0x35, 0x78, 0x65,
	0xe2, 0xdb, 0xd0, 0x10, 0x9c, 0x65, 0xe9, 0xac, 0x5d, 0x3b, 0xd2, 0xba, 0x06, 0x2d, 0x4f, 0x9d,
	0x53, 0x68, 0x3d, 0xe3, 0xcb, 0x17, 0x6c, 0xb2, 0xe0, 0xe7, 0x2c, 0x11, 0x18, 0x81, 0xfe, 0x23,
	0x5f, 0xaa, 0x72, 0x06, 0x95, 0x26, 0xde, 0x87, 0x8d, 0x4b, 0xe9, 0x2e, 0x13, 0x8b, 0x43, 0xe7,
	0x39, 0xec, 0xbc, 0xe0, 0xe3, 0x3c, 0x15, 0x34, 0xbd, 0xa2, 0x7c, 0x9c, 0x8a, 0x18, 0x7f, 0x0a,
	0xf0, 0xdd, 0x24, 0x65, 0x79, 0x14, 0xb3, 0x9c, 0xb5, 0xb5, 0x23, 0xbd, 0x5b, 0xa3, 0x86, 0x42,
	0x6c, 0x96, 0x33, 0x7c, 0x0f, 0xcc, 0x8b, 0x64, 0xc6, 0xc4, 0xb2, 0xf0, 0xcb, 0x6a, 0x2d, 0x0a,
	0x05, 0x24, 0x03, 0x3a, 0xbf, 0x6a, 0x00, 0xfd, 0x3c, 0x17, 0x65, 0xb9, 0x7b, 0x60, 0x26, 0xb3,
	0xfc, 0xa4, 0x17, 0x15, 0xb7, 0xcb, 0x7a, 0x1b, 0x14, 0x14, 0xa4, 0xe8, 0x96, 0x01, 0xa7, 0x8f,
	0xa3, 0x8a, 0x9e, 0xde, 0xd5, 0x55, 0xc0, 0xe9, 0xe3, 0xeb, 0x80, 0x82, 0x50, 0x11, 0xa0, 0x2b,
	0x46, 0x05, 0xc7, 0x22, 0xe0, 0x33, 0x68, 0xc5, 0xe9, 0xe2, 0x62, 0xc2, 0xcb, 0x88, 0xfa, 0x91,
	0xde, 0xd5, 0xa8, 0x59, 0x60, 0x2a, 0xa4, 0xf3, 0x2d, 0xb4, 0xca, 0x3e, 0x0b, 0x56, 0x5f, 0xc1,
	0xa6, 0x50, 0x56, 0xa6, 0x18, 0x99, 0xbd, 0x3d, 0x29, 0xf9, 0x3b, 0x52, 0xd0, 0x2a, 0xa6, 0x43,
	0xc1, 0x2c, 0x7c, 0xe7, 0x4c, 0xb0, 0x29, 0xc6, 0x50, 0xff, 0x41, 0xce, 0xa0, 0x90, 0x57, 0xd9,
	0xf8, 0x21, 0x80, 0x48, 0xaf, 0xa2, 0x22, 0x43, 0xc9, 0x62, 0xf6, 0xd0, 0x5a, 0xd1, 0xa2, 0xa2,
	0x21, 0xaa, 0xe2, 0x9d, 0xbf, 0x34, 0x80, 0x41, 0xc2, 0x27, 0x71, 0xd1, 0x84, 0x94, 0x5d, 0x9e,
	0xa2, 0x19, 0x9b, 0xf2, 0xb2, 0xb2, 0xa1, 0x10, 0x8f, 0x4d, 0x39, 0x3e, 0x82, 0x7a, 0xbe, 0x9c,
	0x17, 0xd3, 0xdb, 0xee, 0xb5, 0x64, 0x61, 0xa9, 0x76, 0xb8, 0x9c, 0x73, 0xaa, 0x3c, 0xf8, 0x21,
	0x98, 0x2c, 0xcf, 0x45, 0xc5, 0x40, 0x57, 0x0c, 0xb6, 0x65, 0xe0, 0x6a, 0x1a, 0x14, 0xd8, 0x6a,
	0x32, 0x5f, 0xc3, 0xd6, 0xa5, 0xe2, 0x56, 0xa5, 0xd4, 0x3f, 0x40, 0xba, 0x75, 0xb9, 0x76, 0xea,
	0xfc, 0xac, 0x41, 0x93, 0xcc, 0xf2, 0x24, 0x4f, 0x78, 0x86, 0x3b, 0xd0, 0xc8, 0xd4, 0x3b, 0x56,
	0x8c, 0xcd, 0x1e, 0xc8, 0xe4, 0xe2, 0x65, 0xd3, 0xd2, 0x23, 0xdf, 0x62, 0x12, 0x67, 0xe5, 0x60,
	0xa5, 0x89, 0xef, 0x82, 0x71, 0xc9, 0x26, 0x49, 0x1c, 0x89, 0xf4, 0x4a, 0xcd, 0xb3, 0x49, 0x9b,
	0x0a, 0xa0, 0xe9, 0x15, 0xbe, 0x0f, 0x0d, 0xd5, 0x76, 0xa6, 0xe6, 0x58, 0xb6, 0xb0, 0x12, 0x8a,
	0x96, 0xde, 0xce, 0x1f, 0x35, 0x30, 0x9d, 0x59, 0xc6, 0x45, 0x5e, 0x0c, 0xe5, 0x0b, 0xd8, 0x19,
	0xa7, 0x93, 0x09, 0x1f, 0xe7, 0x49, 0x3a, 0x5b, 0x57, 0x71, 0x7b, 0x05, 0x2b, 0x29, 0x57, 0x17,
	0xd4, 0xfe, 0xeb, 0x02, 0x7c, 0x1f, 0x76, 0xb8, 0xec, 0x73, 0x19, 0x25, 0x71, 0xc4, 0x84, 0x60,
	0x4b, 0xc5, 0x55, 0xa7, 0x5b, 0x05, 0xec, 0xc4, 0x7d, 0x09, 0xe2, 0xcf, 0x61, 0x6b, 0xce, 0x84,
	0x14, 0x24, 0x9d, 0x45, 0x39, 0x7b, 0xa3, 0x74, 0x34, 0x68, 0xeb, 0x1a, 0x0c, 0xd9, 0x1b, 0xfc,
	0x09, 0x18, 0x79, 0x32, 0xe5, 0x59, 0xce, 0xa6, 0xf3, 0xf6, 0x86, 0x2a, 0xb3, 0x02, 0xf0, 0x23,
	0x38, 0x78, 0x23, 0xe6, 0xe3, 0x28, 0xe3, 0xe2, 0x92, 0x8b, 0x68, 0x3c, 0x49, 0xf8, 0x2c, 0x8f,
	0x92, 0xb8, 0xdd, 0x38, 0xd2, 0xba, 0x3a, 0xc5, 0xd2, 0x19, 0x28, 0x9f, 0xa5, 0x5c, 0x4e, 0x8c,
	0x4f, 0xa0, 0xc5, 0x7f, 0xca, 0x05, 0x8b, 0xe6, 0xb2, 0xfb, 0xac, 0xbd, 0xa9, 0x7a, 0x51, 0xc3,
	0x5b, 0xdf, 0x04, 0xd4, 0x54, 0x51, 0x4a, 0xa2, 0xac, 0xf3, 0x12, 0x0c, 0x52, 0x72, 0xbf, 0xd9,
	0xec, 0xde, 0xa3, 0x41, 0xed, 0x3d, 0x1a, 0x74, 0x7e, 0xd7, 0x60, 0xc7, 0xe6, 0x13, 0x9e, 0xf3,
	0xb3, 0xa5, 0x63, 0x7f, 0xe4, 0x40, 0xee, 0x40, 0xf3, 0x9d, 0xea, 0x9b, 0x49, 0xa9, 0xed, 0x5b,
	0xb2, 0xe9, 0x37, 0x96, 0xad, 0xfe, 0x21, 0xd9, 0x3a, 0x7f, 0xd6, 0xc0, 0x0c, 0x38, 0x13, 0xe3,
	0xef, 0x3f, 0x92, 0xe4, 0x31, 0xec, 0xbd, 0x35, 0xe5, 0x35, 0xbe, 0x06, 0xdd, 0x5d, 0x9f, 0x75,
	0xc1, 0xbc, 0x07, 0xe5, 0xcf, 0xa6, 0x18, 0x90, 0x22, 0x6f, 0xf6, 0x76, 0x56, 0x3f, 0x2e, 0x75,
	0x3f, 0x35, 0x2f, 0xd7, 0xf6, 0x0a, 0x02, 0x3d, 0xce, 0x26, 0xe5, 0xfb, 0x91, 0xe6, 0xff, 0xe4,
	0xd9, 0xfc, 0xad, 0x81, 0xf9, 0x7c, 0xc1, 0xc5, 0x92, 0xf2, 0x6c, 0x31, 0xc9, 0x6f, 0xf4, 0x72,
	0xba, 0xd0, 0xe4, 0xe5, 0x96, 0x28, 0xb7, 0xa1, 0x5a, 0x5a, 0xd5, 0xe6, 0xa0, 0xd7, 0x5e, 0x7c,
	0x08, 0x9b, 0x72, 0x73, 0xce, 0x16, 0x53, 0xb5, 0xb4, 0x74, 0xda, 0x10, 0xe9, 0x95, 0xb7, 0x98,
	0xca, 0x8f, 0x5d, 0x36, 0x4e, 0x05, 0x2f, 0x36, 0x41, 0x8d, 0x96, 0x27, 0x29, 0x4a, 0x9c, 0x64,
	0x39, 0x9b, 0x8d, 0x79, 0xa6, 0x44, 0xa9, 0xd1, 0x15, 0xf0, 0xaf, 0x0e, 0x1b, 0x37, 0xe8, 0xf0,
	0xc1, 0x6f, 0x75, 0x30, 0xae, 0x3f, 0xb8, 0xd8, 0x84, 0xcd, 0x60, 0x64, 0x59, 0x24, 0x08, 0xd0,
	0x2d, 0xbc, 0x0f, 0x68, 0xe4, 0x91, 0x57, 0xe7, 0xc4, 0x0a, 0x89, 0x1d, 0x11, 0x4a, 0x7d, 0x8a,
	0x34, 0x8c, 0x61, 0xdb, 0xf2, 0x3d, 0x8f, 0x58, 0x61, 0x34, 0xe8, 0x3b, 0x2e, 0xb1, 0x51, 0x0d,
	0x1f, 0xc0, 0xee, 0x39, 0xa1, 0x43, 0x27, 0x08, 0x1c, 0xdf, 0x8b, 0x6c, 0xe2, 0x39, 0xc4, 0x46,
	0x3a, 0xbe, 0x03, 0x07, 0x96, 0xef, 0xba, 0xc4, 0x0a, 0x25, 0xec, 0xf9, 0x61, 0x44, 0x5e, 0x39,
	0x41, 0x18, 0xa0, 0xba, 0xac, 0xed, 0xb8, 0x2e, 0x79, 0xd2, 0x77, 0xa3, 0x3e, 0x7d, 0x32, 0x1a,
	0x12, 0x2f, 0x44, 0x1b, 0xb2, 0x4e, 0x85, 0xda, 0xce, 0x90, 0x78, 0xb2, 0x1c, 0xda, 0xc4, 0xb7,
	0x01, 0x57, 0xb0, 0xe3, 0xd9, 0xe4, 0x55, 0x14, 0xbe, 0x3e, 0x27, 0xa8, 0x89, 0xef, 0xc2, 0x61,
	0x85, 0xaf, 0xdf, 0xd3, 0x1f, 0x12, 0x64, 0x60, 0x04, 0xad, 0xca, 0x19, 0xfa, 0xe7, 0xcf, 0x10,
	0xac, 0x57, 0xa7, 0xfe, 0x4b, 0x4a, 0x2c, 0x9f, 0xda, 0xc8, 0x5c, 0x87, 0x5f, 0x10, 0x2b, 0xf4,
	0x69, 0xe4, 0xd8, 0xa8, 0x25, 0xc9, 0x57, 0x70, 0x40, 0xfa, 0xd4, 0x7a, 0x1a, 0x51, 0x12, 0x8c,
	0xdc, 0x10, 0x6d, 0x49, 0x09, 0x06, 0x8e, 0x4b, 0x54, 0x47, 0x03, 0x7f, 0xe4, 0xd9, 0x68, 0x1b,
	0xef, 0x80, 0x39, 0x24, 0x61, 0xbf, 0xd2, 0x64, 0x47, 0xde, 0x6f, 0xf5, 0xad, 0xa7, 0xa4, 0x42,
	0x10, 0x6e, 0xc3, 0xbe, 0xd5, 0xf7, 0x64, 0x92, 0x45, 0x49, 0x3f, 0x24, 0xd1, 0xc0, 0x77, 0x6d,
	0x42, 0xd1, 0xae, 0x6c, 0xf0, 0x1d, 0x8f, 0xe3, 0x12, 0x84, 0xd7, 0x32, 0x6c, 0xe2, 0x92, 0x55,
	0xc6, 0xde, 0x5a, 0x46, 0xe5, 0x91, 0x19, 0xfb, 0xb2, 0x99, 0xb3, 0x91, 0xe3, 0xda, 0xa5, 0x50,
	0xc5, 0xd0, 0x0e, 0xf0, 0x2e, 0x6c, 0x55, 0xcd, 0x78, 0xae, 0x13, 0x84, 0xe8, 0x36, 0x3e, 0x84,
	0xbd, 0x0a, 0x1a, 0x92, 0x90, 0x3a, 0x56, 0xa1, 0xea, 0xa1, 0x8c, 0xf5, 0x47, 0x61, 0xe4, 0x0f,
	0xa2, 0x21, 0x19, 0xfa, 0xf4, 0x35, 0x6a, 0x3f, 0xf8, 0x45, 0x83, 0x66, 0xf5, 0xd1, 0xc5, 0x4d,
	0xa8, 0x7b, 0xbe, 0x47, 0xd0, 0x2d, 0x69, 0x9d, 0xf9, 0xbe, 0x8b, 0x34, 0x69, 0x39, 0x5e, 0xf8,
	0x0d, 0xaa, 0x61, 0x03, 0x36, 0x1c, 0x2f, 0x7c, 0x74, 0x8a, 0xf4, 0xd2, 0x3c, 0xe9, 0xa1, 0x7a,
	0x69, 0x9e, 0x3e, 0x46, 0x1b, 0xd2, 0x1c, 0xb8, 0x7e, 0x3f, 0x44, 0x80, 0x01, 0x1a, 0xb6, 0x3f,
	0x3a, 0x73, 0x09, 0x32, 0xa5, 0x1d, 0x84, 0xd4, 0xf1, 0x9e, 0xa0, 0x7d, 0xc9, 0xa0, 0x9c, 0xc4,
	0x99, 0xe3, 0xf5, 0xe9, 0x6b, 0x14, 0x4b, 0x35, 0x4b, 0xa8, 0x48, 0xe6, 0x17, 0x0d, 0xf5, 0x9f,
	0xf2, 0xe4, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x17, 0xca, 0xe4, 0x63, 0x0a, 0x00, 0x00,
}