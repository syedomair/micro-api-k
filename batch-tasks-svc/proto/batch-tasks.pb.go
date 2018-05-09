// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/batch-tasks.proto

/*
Package batch_tasks is a generated protocol buffer package.

It is generated from these files:
	proto/batch-tasks.proto

It has these top-level messages:
	BatchTask
	ResponseBatchTask
	Response
	ResponseList
	UserShorten
	UserList
	RequestQuery
	BatchTaskMessage
*/
package batch_tasks

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BatchTask struct {
	Id          string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	ClientId    string `protobuf:"bytes,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	ApiName     string `protobuf:"bytes,3,opt,name=api_name,json=apiName" json:"api_name,omitempty"`
	Status      string `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	CreatedAt   string `protobuf:"bytes,6,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	CompletedAt string `protobuf:"bytes,7,opt,name=completed_at,json=completedAt" json:"completed_at,omitempty"`
}

func (m *BatchTask) Reset()                    { *m = BatchTask{} }
func (m *BatchTask) String() string            { return proto.CompactTextString(m) }
func (*BatchTask) ProtoMessage()               {}
func (*BatchTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BatchTask) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BatchTask) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *BatchTask) GetApiName() string {
	if m != nil {
		return m.ApiName
	}
	return ""
}

func (m *BatchTask) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *BatchTask) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *BatchTask) GetCompletedAt() string {
	if m != nil {
		return m.CompletedAt
	}
	return ""
}

type ResponseBatchTask struct {
	Result string            `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	Error  map[string]string `protobuf:"bytes,2,rep,name=error" json:"error,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Data   *BatchTask        `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *ResponseBatchTask) Reset()                    { *m = ResponseBatchTask{} }
func (m *ResponseBatchTask) String() string            { return proto.CompactTextString(m) }
func (*ResponseBatchTask) ProtoMessage()               {}
func (*ResponseBatchTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ResponseBatchTask) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *ResponseBatchTask) GetError() map[string]string {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ResponseBatchTask) GetData() *BatchTask {
	if m != nil {
		return m.Data
	}
	return nil
}

type Response struct {
	Result string            `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	Error  map[string]string `protobuf:"bytes,2,rep,name=error" json:"error,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Data   map[string]string `protobuf:"bytes,3,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *Response) GetError() map[string]string {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type ResponseList struct {
	Result string            `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	Error  map[string]string `protobuf:"bytes,2,rep,name=error" json:"error,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Data   *UserList         `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *ResponseList) Reset()                    { *m = ResponseList{} }
func (m *ResponseList) String() string            { return proto.CompactTextString(m) }
func (*ResponseList) ProtoMessage()               {}
func (*ResponseList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ResponseList) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *ResponseList) GetError() map[string]string {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ResponseList) GetData() *UserList {
	if m != nil {
		return m.Data
	}
	return nil
}

type UserShorten struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	ClientId  string `protobuf:"bytes,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,5,opt,name=email" json:"email,omitempty"`
	IsAdmin   string `protobuf:"bytes,6,opt,name=is_admin,json=isAdmin" json:"is_admin,omitempty"`
	CreatedAt string `protobuf:"bytes,7,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *UserShorten) Reset()                    { *m = UserShorten{} }
func (m *UserShorten) String() string            { return proto.CompactTextString(m) }
func (*UserShorten) ProtoMessage()               {}
func (*UserShorten) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UserShorten) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserShorten) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *UserShorten) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserShorten) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UserShorten) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserShorten) GetIsAdmin() string {
	if m != nil {
		return m.IsAdmin
	}
	return ""
}

func (m *UserShorten) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *UserShorten) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type UserList struct {
	Offset string         `protobuf:"bytes,1,opt,name=offset" json:"offset,omitempty"`
	Limit  string         `protobuf:"bytes,2,opt,name=limit" json:"limit,omitempty"`
	Count  string         `protobuf:"bytes,3,opt,name=count" json:"count,omitempty"`
	List   []*UserShorten `protobuf:"bytes,4,rep,name=list" json:"list,omitempty"`
}

func (m *UserList) Reset()                    { *m = UserList{} }
func (m *UserList) String() string            { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()               {}
func (*UserList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UserList) GetOffset() string {
	if m != nil {
		return m.Offset
	}
	return ""
}

func (m *UserList) GetLimit() string {
	if m != nil {
		return m.Limit
	}
	return ""
}

func (m *UserList) GetCount() string {
	if m != nil {
		return m.Count
	}
	return ""
}

func (m *UserList) GetList() []*UserShorten {
	if m != nil {
		return m.List
	}
	return nil
}

type RequestQuery struct {
	Offset  string `protobuf:"bytes,1,opt,name=offset" json:"offset,omitempty"`
	Limit   string `protobuf:"bytes,2,opt,name=limit" json:"limit,omitempty"`
	Orderby string `protobuf:"bytes,3,opt,name=orderby" json:"orderby,omitempty"`
	Sort    string `protobuf:"bytes,4,opt,name=sort" json:"sort,omitempty"`
	Filter  string `protobuf:"bytes,5,opt,name=filter" json:"filter,omitempty"`
	Search  string `protobuf:"bytes,6,opt,name=search" json:"search,omitempty"`
}

func (m *RequestQuery) Reset()                    { *m = RequestQuery{} }
func (m *RequestQuery) String() string            { return proto.CompactTextString(m) }
func (*RequestQuery) ProtoMessage()               {}
func (*RequestQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RequestQuery) GetOffset() string {
	if m != nil {
		return m.Offset
	}
	return ""
}

func (m *RequestQuery) GetLimit() string {
	if m != nil {
		return m.Limit
	}
	return ""
}

func (m *RequestQuery) GetOrderby() string {
	if m != nil {
		return m.Orderby
	}
	return ""
}

func (m *RequestQuery) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *RequestQuery) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *RequestQuery) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

type BatchTaskMessage struct {
	BatchTaskId string `protobuf:"bytes,1,opt,name=batch_task_id,json=batchTaskId" json:"batch_task_id,omitempty"`
	ClientId    string `protobuf:"bytes,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	Limit       string `protobuf:"bytes,3,opt,name=limit" json:"limit,omitempty"`
	Offset      string `protobuf:"bytes,4,opt,name=offset" json:"offset,omitempty"`
	Orderby     string `protobuf:"bytes,5,opt,name=orderby" json:"orderby,omitempty"`
	Sort        string `protobuf:"bytes,6,opt,name=sort" json:"sort,omitempty"`
	Filter      string `protobuf:"bytes,7,opt,name=filter" json:"filter,omitempty"`
	Search      string `protobuf:"bytes,8,opt,name=search" json:"search,omitempty"`
}

func (m *BatchTaskMessage) Reset()                    { *m = BatchTaskMessage{} }
func (m *BatchTaskMessage) String() string            { return proto.CompactTextString(m) }
func (*BatchTaskMessage) ProtoMessage()               {}
func (*BatchTaskMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *BatchTaskMessage) GetBatchTaskId() string {
	if m != nil {
		return m.BatchTaskId
	}
	return ""
}

func (m *BatchTaskMessage) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *BatchTaskMessage) GetLimit() string {
	if m != nil {
		return m.Limit
	}
	return ""
}

func (m *BatchTaskMessage) GetOffset() string {
	if m != nil {
		return m.Offset
	}
	return ""
}

func (m *BatchTaskMessage) GetOrderby() string {
	if m != nil {
		return m.Orderby
	}
	return ""
}

func (m *BatchTaskMessage) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *BatchTaskMessage) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *BatchTaskMessage) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

func init() {
	proto.RegisterType((*BatchTask)(nil), "batch_tasks.BatchTask")
	proto.RegisterType((*ResponseBatchTask)(nil), "batch_tasks.ResponseBatchTask")
	proto.RegisterType((*Response)(nil), "batch_tasks.Response")
	proto.RegisterType((*ResponseList)(nil), "batch_tasks.ResponseList")
	proto.RegisterType((*UserShorten)(nil), "batch_tasks.UserShorten")
	proto.RegisterType((*UserList)(nil), "batch_tasks.UserList")
	proto.RegisterType((*RequestQuery)(nil), "batch_tasks.RequestQuery")
	proto.RegisterType((*BatchTaskMessage)(nil), "batch_tasks.BatchTaskMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BatchTasksService service

type BatchTasksServiceClient interface {
	GetAllUser(ctx context.Context, in *RequestQuery, opts ...grpc.CallOption) (*ResponseBatchTask, error)
	GetAllUserStatus(ctx context.Context, in *BatchTask, opts ...grpc.CallOption) (*ResponseBatchTask, error)
	GetAllUserOutput(ctx context.Context, in *BatchTask, opts ...grpc.CallOption) (*ResponseList, error)
}

type batchTasksServiceClient struct {
	cc *grpc.ClientConn
}

func NewBatchTasksServiceClient(cc *grpc.ClientConn) BatchTasksServiceClient {
	return &batchTasksServiceClient{cc}
}

func (c *batchTasksServiceClient) GetAllUser(ctx context.Context, in *RequestQuery, opts ...grpc.CallOption) (*ResponseBatchTask, error) {
	out := new(ResponseBatchTask)
	err := grpc.Invoke(ctx, "/batch_tasks.BatchTasksService/GetAllUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *batchTasksServiceClient) GetAllUserStatus(ctx context.Context, in *BatchTask, opts ...grpc.CallOption) (*ResponseBatchTask, error) {
	out := new(ResponseBatchTask)
	err := grpc.Invoke(ctx, "/batch_tasks.BatchTasksService/GetAllUserStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *batchTasksServiceClient) GetAllUserOutput(ctx context.Context, in *BatchTask, opts ...grpc.CallOption) (*ResponseList, error) {
	out := new(ResponseList)
	err := grpc.Invoke(ctx, "/batch_tasks.BatchTasksService/GetAllUserOutput", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BatchTasksService service

type BatchTasksServiceServer interface {
	GetAllUser(context.Context, *RequestQuery) (*ResponseBatchTask, error)
	GetAllUserStatus(context.Context, *BatchTask) (*ResponseBatchTask, error)
	GetAllUserOutput(context.Context, *BatchTask) (*ResponseList, error)
}

func RegisterBatchTasksServiceServer(s *grpc.Server, srv BatchTasksServiceServer) {
	s.RegisterService(&_BatchTasksService_serviceDesc, srv)
}

func _BatchTasksService_GetAllUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchTasksServiceServer).GetAllUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/batch_tasks.BatchTasksService/GetAllUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchTasksServiceServer).GetAllUser(ctx, req.(*RequestQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _BatchTasksService_GetAllUserStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchTask)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchTasksServiceServer).GetAllUserStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/batch_tasks.BatchTasksService/GetAllUserStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchTasksServiceServer).GetAllUserStatus(ctx, req.(*BatchTask))
	}
	return interceptor(ctx, in, info, handler)
}

func _BatchTasksService_GetAllUserOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchTask)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchTasksServiceServer).GetAllUserOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/batch_tasks.BatchTasksService/GetAllUserOutput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchTasksServiceServer).GetAllUserOutput(ctx, req.(*BatchTask))
	}
	return interceptor(ctx, in, info, handler)
}

var _BatchTasksService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "batch_tasks.BatchTasksService",
	HandlerType: (*BatchTasksServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllUser",
			Handler:    _BatchTasksService_GetAllUser_Handler,
		},
		{
			MethodName: "GetAllUserStatus",
			Handler:    _BatchTasksService_GetAllUserStatus_Handler,
		},
		{
			MethodName: "GetAllUserOutput",
			Handler:    _BatchTasksService_GetAllUserOutput_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/batch-tasks.proto",
}

func init() { proto.RegisterFile("proto/batch-tasks.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 755 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6e, 0x13, 0x49,
	0x10, 0xd6, 0x8c, 0xff, 0xcb, 0xd9, 0xdd, 0xa4, 0x95, 0x4d, 0x26, 0xce, 0x7a, 0x37, 0x3b, 0xbb,
	0x87, 0x04, 0x81, 0x2d, 0x12, 0x09, 0xa2, 0x5c, 0x90, 0x11, 0x11, 0x8a, 0xc4, 0x8f, 0x70, 0xe0,
	0x6c, 0x3a, 0x9e, 0x76, 0xd2, 0x64, 0x3c, 0x3d, 0x74, 0xd7, 0x44, 0xb2, 0x80, 0x4b, 0x5e, 0x81,
	0x1b, 0xef, 0xc0, 0x73, 0x70, 0x0f, 0x0f, 0xc0, 0x05, 0xde, 0x03, 0x75, 0x4f, 0x7b, 0xfc, 0x13,
	0x9b, 0x28, 0xca, 0x6d, 0xea, 0xbf, 0xbe, 0xaf, 0xa6, 0xaa, 0x61, 0x35, 0x96, 0x02, 0x45, 0xf3,
	0x88, 0x62, 0xf7, 0xe4, 0x0e, 0x52, 0x75, 0xaa, 0x1a, 0x46, 0x43, 0xaa, 0x46, 0xd5, 0x31, 0xaa,
	0xda, 0x5f, 0xc7, 0x42, 0x1c, 0x87, 0xac, 0x49, 0x63, 0xde, 0xa4, 0x51, 0x24, 0x90, 0x22, 0x17,
	0x91, 0x75, 0xf5, 0x3f, 0x3b, 0x50, 0x79, 0xa8, 0xbd, 0x5f, 0x52, 0x75, 0x4a, 0x7e, 0x07, 0x97,
	0x07, 0x9e, 0xb3, 0xe1, 0x6c, 0x56, 0xda, 0x2e, 0x0f, 0xc8, 0x3a, 0x54, 0xba, 0x21, 0x67, 0x11,
	0x76, 0x78, 0xe0, 0xb9, 0x46, 0x5d, 0x4e, 0x15, 0x07, 0x01, 0x59, 0x83, 0x32, 0x8d, 0x79, 0x27,
	0xa2, 0x7d, 0xe6, 0xe5, 0x8c, 0xad, 0x44, 0x63, 0xfe, 0x8c, 0xf6, 0x19, 0x59, 0x81, 0xa2, 0x42,
	0x8a, 0x89, 0xf2, 0xf2, 0xc6, 0x60, 0x25, 0x52, 0x07, 0xe8, 0x4a, 0x46, 0x91, 0x05, 0x1d, 0x8a,
	0x5e, 0xd1, 0xd8, 0x2a, 0x56, 0xd3, 0x42, 0xf2, 0x2f, 0x2c, 0x74, 0x45, 0x3f, 0x0e, 0x99, 0x75,
	0x28, 0x19, 0x87, 0x6a, 0xa6, 0x6b, 0xa1, 0x7f, 0xe1, 0xc0, 0x52, 0x9b, 0xa9, 0x58, 0x44, 0x8a,
	0x8d, 0xfa, 0x5e, 0x81, 0xa2, 0x64, 0x2a, 0x09, 0xd1, 0xf6, 0x6e, 0x25, 0xf2, 0x00, 0x0a, 0x4c,
	0x4a, 0x21, 0x3d, 0x77, 0x23, 0xb7, 0x59, 0xdd, 0xde, 0x6a, 0x8c, 0x11, 0xd3, 0xb8, 0x94, 0xa6,
	0xb1, 0xaf, 0x7d, 0xf7, 0x23, 0x94, 0x83, 0x76, 0x1a, 0x47, 0x6e, 0x41, 0x3e, 0xa0, 0x48, 0x0d,
	0xbe, 0xea, 0xf6, 0xca, 0x44, 0x7c, 0x16, 0xd7, 0x36, 0x3e, 0xb5, 0x5d, 0x80, 0x51, 0x02, 0xb2,
	0x08, 0xb9, 0x53, 0x36, 0xb0, 0xfd, 0xe8, 0x4f, 0xb2, 0x0c, 0x85, 0x33, 0x1a, 0x26, 0xcc, 0x12,
	0x99, 0x0a, 0x7b, 0xee, 0xae, 0xe3, 0x9f, 0xbb, 0x50, 0x1e, 0x76, 0x33, 0x17, 0xcb, 0xbd, 0x49,
	0x2c, 0x1b, 0x33, 0xb1, 0xcc, 0x80, 0xb0, 0x93, 0x41, 0xd0, 0x61, 0xff, 0xcc, 0x0e, 0x7b, 0x44,
	0x91, 0xa6, 0x51, 0x37, 0xc4, 0x52, 0xbb, 0x0f, 0x95, 0x2c, 0xd9, 0xb5, 0x48, 0xf8, 0xe2, 0xc0,
	0xc2, 0xb0, 0x9f, 0x27, 0x5c, 0xe1, 0x5c, 0x22, 0xf6, 0x26, 0x89, 0xf8, 0x7f, 0x26, 0x22, 0x9d,
	0x61, 0x06, 0x19, 0x5b, 0x13, 0xf3, 0xfc, 0x73, 0x22, 0xf4, 0x95, 0x62, 0x52, 0x87, 0xdd, 0x78,
	0x9c, 0x3f, 0x1c, 0xa8, 0xea, 0x64, 0x87, 0x27, 0x42, 0x22, 0x8b, 0xae, 0xb7, 0x55, 0x75, 0x80,
	0x1e, 0x97, 0x0a, 0xc7, 0xf7, 0xaa, 0x62, 0x34, 0x66, 0xb3, 0xd6, 0xa1, 0x12, 0xd2, 0xa1, 0x35,
	0x5d, 0xae, 0xb2, 0x56, 0x18, 0xe3, 0x32, 0x14, 0x58, 0x9f, 0xf2, 0xd0, 0x2b, 0xa4, 0x2d, 0x19,
	0x41, 0xef, 0x29, 0x57, 0x1d, 0x1a, 0xf4, 0x79, 0x64, 0x57, 0xae, 0xc4, 0x55, 0x4b, 0x8b, 0x53,
	0xfb, 0x58, 0x9a, 0xde, 0xc7, 0x3a, 0x40, 0x12, 0x07, 0x43, 0x73, 0x39, 0x35, 0x5b, 0x4d, 0x0b,
	0xfd, 0xf7, 0x50, 0x1e, 0x72, 0xa6, 0x87, 0x25, 0x7a, 0x3d, 0xc5, 0xb2, 0x61, 0xa5, 0x92, 0x6e,
	0x29, 0xe4, 0x7d, 0x8e, 0x43, 0x96, 0x8c, 0xa0, 0xb5, 0x5d, 0x91, 0x44, 0x68, 0xf1, 0xa5, 0x02,
	0xb9, 0x0d, 0xf9, 0x90, 0x2b, 0xf4, 0xf2, 0x66, 0xae, 0xde, 0xa5, 0xe1, 0x58, 0x3e, 0xdb, 0xc6,
	0xcb, 0xff, 0x64, 0xfe, 0x97, 0xb7, 0x09, 0x53, 0xf8, 0x22, 0x61, 0x72, 0x70, 0xcd, 0x16, 0x3c,
	0x28, 0x09, 0x19, 0x30, 0x79, 0x34, 0x18, 0x1e, 0x2f, 0x2b, 0x12, 0x02, 0x79, 0x25, 0x24, 0x5a,
	0x76, 0xcd, 0xb7, 0xce, 0xdd, 0xe3, 0x21, 0x32, 0x69, 0xa9, 0xb5, 0x92, 0x39, 0x74, 0x8c, 0xca,
	0xee, 0x89, 0x65, 0xd6, 0x4a, 0xfe, 0x37, 0x07, 0x16, 0xb3, 0xfb, 0xf0, 0x94, 0x29, 0x45, 0x8f,
	0x19, 0xf1, 0xe1, 0xb7, 0x11, 0xa4, 0x4e, 0xf6, 0x4b, 0xa4, 0xd7, 0x5a, 0x3b, 0x1e, 0x5c, 0xf1,
	0x6f, 0x64, 0x48, 0x72, 0xe3, 0x48, 0x46, 0xb8, 0xf3, 0x13, 0xb8, 0xc7, 0x10, 0x16, 0x66, 0x23,
	0x2c, 0xce, 0x44, 0x58, 0x9a, 0x83, 0xb0, 0x3c, 0x8e, 0x70, 0xfb, 0xc2, 0x85, 0xa5, 0x0c, 0xa1,
	0x3a, 0x64, 0xf2, 0x8c, 0x77, 0x19, 0x79, 0x0d, 0xf0, 0x98, 0x61, 0x2b, 0x0c, 0xf5, 0xbc, 0xc8,
	0xda, 0xd4, 0x6a, 0x8e, 0x86, 0x55, 0xfb, 0xfb, 0xd7, 0xa7, 0xd8, 0x5f, 0x3d, 0xff, 0xfa, 0xfd,
	0xa3, 0xbb, 0x44, 0xfe, 0x68, 0x9e, 0xdd, 0x4d, 0x5f, 0xb8, 0x66, 0xa2, 0x98, 0x54, 0x44, 0xc0,
	0xe2, 0xa8, 0xc2, 0x61, 0xfa, 0xac, 0xcc, 0xb9, 0xcb, 0x57, 0x16, 0xf9, 0xcf, 0x14, 0xa9, 0x93,
	0xf5, 0xa9, 0x22, 0xcd, 0xf4, 0xb9, 0x6a, 0xbe, 0xe3, 0xc1, 0x07, 0xf2, 0x66, 0xbc, 0xe0, 0xf3,
	0x04, 0xe3, 0x04, 0xe7, 0x16, 0x5c, 0x9b, 0x7b, 0x8b, 0xe6, 0xd7, 0x12, 0x26, 0xa5, 0xa9, 0x75,
	0x54, 0x34, 0x8f, 0xf2, 0xce, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x77, 0x4a, 0x58, 0x13, 0xda,
	0x07, 0x00, 0x00,
}
