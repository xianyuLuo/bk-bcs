// Code generated by protoc-gen-go.
// source: sidecar.proto
// DO NOT EDIT!

/*
Package sidecar is a generated protocol buffer package.

It is generated from these files:
	sidecar.proto

It has these top-level messages:
	ModInfo
	ConfigsMetadata
	PingReq
	PingResp
	InjectReq
	InjectResp
	WatchReloadReq
	WatchReloadResp
	ReportReloadReq
	ReportReloadResp
*/
package sidecar

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "bk-bscp/internal/protocol/common"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

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

type ModInfo struct {
	BusinessName string            `protobuf:"bytes,1,opt,name=businessName" json:"businessName,omitempty"`
	AppName      string            `protobuf:"bytes,2,opt,name=appName" json:"appName,omitempty"`
	ClusterName  string            `protobuf:"bytes,3,opt,name=clusterName" json:"clusterName,omitempty"`
	ZoneName     string            `protobuf:"bytes,4,opt,name=zoneName" json:"zoneName,omitempty"`
	Dc           string            `protobuf:"bytes,5,opt,name=dc" json:"dc,omitempty"`
	IP           string            `protobuf:"bytes,6,opt,name=IP" json:"IP,omitempty"`
	Labels       map[string]string `protobuf:"bytes,7,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Path         string            `protobuf:"bytes,8,opt,name=path" json:"path,omitempty"`
	IsReady      bool              `protobuf:"varint,9,opt,name=isReady" json:"isReady,omitempty"`
}

func (m *ModInfo) Reset()                    { *m = ModInfo{} }
func (m *ModInfo) String() string            { return proto.CompactTextString(m) }
func (*ModInfo) ProtoMessage()               {}
func (*ModInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ModInfo) GetBusinessName() string {
	if m != nil {
		return m.BusinessName
	}
	return ""
}

func (m *ModInfo) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *ModInfo) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ModInfo) GetZoneName() string {
	if m != nil {
		return m.ZoneName
	}
	return ""
}

func (m *ModInfo) GetDc() string {
	if m != nil {
		return m.Dc
	}
	return ""
}

func (m *ModInfo) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *ModInfo) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ModInfo) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ModInfo) GetIsReady() bool {
	if m != nil {
		return m.IsReady
	}
	return false
}

type ConfigsMetadata struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Fpath string `protobuf:"bytes,2,opt,name=fpath" json:"fpath,omitempty"`
}

func (m *ConfigsMetadata) Reset()                    { *m = ConfigsMetadata{} }
func (m *ConfigsMetadata) String() string            { return proto.CompactTextString(m) }
func (*ConfigsMetadata) ProtoMessage()               {}
func (*ConfigsMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ConfigsMetadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ConfigsMetadata) GetFpath() string {
	if m != nil {
		return m.Fpath
	}
	return ""
}

type PingReq struct {
	Seq uint64 `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
}

func (m *PingReq) Reset()                    { *m = PingReq{} }
func (m *PingReq) String() string            { return proto.CompactTextString(m) }
func (*PingReq) ProtoMessage()               {}
func (*PingReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PingReq) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

type PingResp struct {
	Seq     uint64         `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	ErrCode common.ErrCode `protobuf:"varint,2,opt,name=errCode,enum=common.ErrCode" json:"errCode,omitempty"`
	ErrMsg  string         `protobuf:"bytes,3,opt,name=errMsg" json:"errMsg,omitempty"`
	Mods    []*ModInfo     `protobuf:"bytes,4,rep,name=mods" json:"mods,omitempty"`
}

func (m *PingResp) Reset()                    { *m = PingResp{} }
func (m *PingResp) String() string            { return proto.CompactTextString(m) }
func (*PingResp) ProtoMessage()               {}
func (*PingResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PingResp) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *PingResp) GetErrCode() common.ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return common.ErrCode_E_OK
}

func (m *PingResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *PingResp) GetMods() []*ModInfo {
	if m != nil {
		return m.Mods
	}
	return nil
}

type InjectReq struct {
	Seq          uint64            `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	BusinessName string            `protobuf:"bytes,2,opt,name=businessName" json:"businessName,omitempty"`
	AppName      string            `protobuf:"bytes,3,opt,name=appName" json:"appName,omitempty"`
	Labels       map[string]string `protobuf:"bytes,4,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *InjectReq) Reset()                    { *m = InjectReq{} }
func (m *InjectReq) String() string            { return proto.CompactTextString(m) }
func (*InjectReq) ProtoMessage()               {}
func (*InjectReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *InjectReq) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *InjectReq) GetBusinessName() string {
	if m != nil {
		return m.BusinessName
	}
	return ""
}

func (m *InjectReq) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *InjectReq) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type InjectResp struct {
	Seq     uint64         `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	ErrCode common.ErrCode `protobuf:"varint,2,opt,name=errCode,enum=common.ErrCode" json:"errCode,omitempty"`
	ErrMsg  string         `protobuf:"bytes,3,opt,name=errMsg" json:"errMsg,omitempty"`
}

func (m *InjectResp) Reset()                    { *m = InjectResp{} }
func (m *InjectResp) String() string            { return proto.CompactTextString(m) }
func (*InjectResp) ProtoMessage()               {}
func (*InjectResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *InjectResp) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *InjectResp) GetErrCode() common.ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return common.ErrCode_E_OK
}

func (m *InjectResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type WatchReloadReq struct {
	Seq          uint64 `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	BusinessName string `protobuf:"bytes,2,opt,name=businessName" json:"businessName,omitempty"`
	AppName      string `protobuf:"bytes,3,opt,name=appName" json:"appName,omitempty"`
}

func (m *WatchReloadReq) Reset()                    { *m = WatchReloadReq{} }
func (m *WatchReloadReq) String() string            { return proto.CompactTextString(m) }
func (*WatchReloadReq) ProtoMessage()               {}
func (*WatchReloadReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *WatchReloadReq) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *WatchReloadReq) GetBusinessName() string {
	if m != nil {
		return m.BusinessName
	}
	return ""
}

func (m *WatchReloadReq) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

type WatchReloadResp struct {
	Seq            uint64             `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	ErrCode        common.ErrCode     `protobuf:"varint,2,opt,name=errCode,enum=common.ErrCode" json:"errCode,omitempty"`
	ErrMsg         string             `protobuf:"bytes,3,opt,name=errMsg" json:"errMsg,omitempty"`
	Releaseid      string             `protobuf:"bytes,4,opt,name=releaseid" json:"releaseid,omitempty"`
	MultiReleaseid string             `protobuf:"bytes,5,opt,name=multiReleaseid" json:"multiReleaseid,omitempty"`
	ReleaseName    string             `protobuf:"bytes,6,opt,name=releaseName" json:"releaseName,omitempty"`
	ReloadType     int32              `protobuf:"varint,7,opt,name=reloadType" json:"reloadType,omitempty"`
	RootPath       string             `protobuf:"bytes,8,opt,name=rootPath" json:"rootPath,omitempty"`
	Metadatas      []*ConfigsMetadata `protobuf:"bytes,9,rep,name=metadatas" json:"metadatas,omitempty"`
}

func (m *WatchReloadResp) Reset()                    { *m = WatchReloadResp{} }
func (m *WatchReloadResp) String() string            { return proto.CompactTextString(m) }
func (*WatchReloadResp) ProtoMessage()               {}
func (*WatchReloadResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *WatchReloadResp) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *WatchReloadResp) GetErrCode() common.ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return common.ErrCode_E_OK
}

func (m *WatchReloadResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *WatchReloadResp) GetReleaseid() string {
	if m != nil {
		return m.Releaseid
	}
	return ""
}

func (m *WatchReloadResp) GetMultiReleaseid() string {
	if m != nil {
		return m.MultiReleaseid
	}
	return ""
}

func (m *WatchReloadResp) GetReleaseName() string {
	if m != nil {
		return m.ReleaseName
	}
	return ""
}

func (m *WatchReloadResp) GetReloadType() int32 {
	if m != nil {
		return m.ReloadType
	}
	return 0
}

func (m *WatchReloadResp) GetRootPath() string {
	if m != nil {
		return m.RootPath
	}
	return ""
}

func (m *WatchReloadResp) GetMetadatas() []*ConfigsMetadata {
	if m != nil {
		return m.Metadatas
	}
	return nil
}

type ReportReloadReq struct {
	Seq            uint64 `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	BusinessName   string `protobuf:"bytes,2,opt,name=businessName" json:"businessName,omitempty"`
	AppName        string `protobuf:"bytes,3,opt,name=appName" json:"appName,omitempty"`
	Releaseid      string `protobuf:"bytes,4,opt,name=releaseid" json:"releaseid,omitempty"`
	MultiReleaseid string `protobuf:"bytes,5,opt,name=multiReleaseid" json:"multiReleaseid,omitempty"`
	ReloadTime     string `protobuf:"bytes,6,opt,name=reloadTime" json:"reloadTime,omitempty"`
	ReloadCode     int32  `protobuf:"varint,7,opt,name=reloadCode" json:"reloadCode,omitempty"`
	ReloadMsg      string `protobuf:"bytes,8,opt,name=reloadMsg" json:"reloadMsg,omitempty"`
}

func (m *ReportReloadReq) Reset()                    { *m = ReportReloadReq{} }
func (m *ReportReloadReq) String() string            { return proto.CompactTextString(m) }
func (*ReportReloadReq) ProtoMessage()               {}
func (*ReportReloadReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ReportReloadReq) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *ReportReloadReq) GetBusinessName() string {
	if m != nil {
		return m.BusinessName
	}
	return ""
}

func (m *ReportReloadReq) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *ReportReloadReq) GetReleaseid() string {
	if m != nil {
		return m.Releaseid
	}
	return ""
}

func (m *ReportReloadReq) GetMultiReleaseid() string {
	if m != nil {
		return m.MultiReleaseid
	}
	return ""
}

func (m *ReportReloadReq) GetReloadTime() string {
	if m != nil {
		return m.ReloadTime
	}
	return ""
}

func (m *ReportReloadReq) GetReloadCode() int32 {
	if m != nil {
		return m.ReloadCode
	}
	return 0
}

func (m *ReportReloadReq) GetReloadMsg() string {
	if m != nil {
		return m.ReloadMsg
	}
	return ""
}

type ReportReloadResp struct {
	Seq     uint64         `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	ErrCode common.ErrCode `protobuf:"varint,2,opt,name=errCode,enum=common.ErrCode" json:"errCode,omitempty"`
	ErrMsg  string         `protobuf:"bytes,3,opt,name=errMsg" json:"errMsg,omitempty"`
}

func (m *ReportReloadResp) Reset()                    { *m = ReportReloadResp{} }
func (m *ReportReloadResp) String() string            { return proto.CompactTextString(m) }
func (*ReportReloadResp) ProtoMessage()               {}
func (*ReportReloadResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ReportReloadResp) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *ReportReloadResp) GetErrCode() common.ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return common.ErrCode_E_OK
}

func (m *ReportReloadResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*ModInfo)(nil), "sidecar.ModInfo")
	proto.RegisterType((*ConfigsMetadata)(nil), "sidecar.ConfigsMetadata")
	proto.RegisterType((*PingReq)(nil), "sidecar.PingReq")
	proto.RegisterType((*PingResp)(nil), "sidecar.PingResp")
	proto.RegisterType((*InjectReq)(nil), "sidecar.InjectReq")
	proto.RegisterType((*InjectResp)(nil), "sidecar.InjectResp")
	proto.RegisterType((*WatchReloadReq)(nil), "sidecar.WatchReloadReq")
	proto.RegisterType((*WatchReloadResp)(nil), "sidecar.WatchReloadResp")
	proto.RegisterType((*ReportReloadReq)(nil), "sidecar.ReportReloadReq")
	proto.RegisterType((*ReportReloadResp)(nil), "sidecar.ReportReloadResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Instance service

type InstanceClient interface {
	Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingResp, error)
	Inject(ctx context.Context, in *InjectReq, opts ...grpc.CallOption) (*InjectResp, error)
	WatchReload(ctx context.Context, in *WatchReloadReq, opts ...grpc.CallOption) (Instance_WatchReloadClient, error)
	ReportReload(ctx context.Context, in *ReportReloadReq, opts ...grpc.CallOption) (*ReportReloadResp, error)
}

type instanceClient struct {
	cc *grpc.ClientConn
}

func NewInstanceClient(cc *grpc.ClientConn) InstanceClient {
	return &instanceClient{cc}
}

func (c *instanceClient) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingResp, error) {
	out := new(PingResp)
	err := grpc.Invoke(ctx, "/sidecar.Instance/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) Inject(ctx context.Context, in *InjectReq, opts ...grpc.CallOption) (*InjectResp, error) {
	out := new(InjectResp)
	err := grpc.Invoke(ctx, "/sidecar.Instance/Inject", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) WatchReload(ctx context.Context, in *WatchReloadReq, opts ...grpc.CallOption) (Instance_WatchReloadClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Instance_serviceDesc.Streams[0], c.cc, "/sidecar.Instance/WatchReload", opts...)
	if err != nil {
		return nil, err
	}
	x := &instanceWatchReloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Instance_WatchReloadClient interface {
	Recv() (*WatchReloadResp, error)
	grpc.ClientStream
}

type instanceWatchReloadClient struct {
	grpc.ClientStream
}

func (x *instanceWatchReloadClient) Recv() (*WatchReloadResp, error) {
	m := new(WatchReloadResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *instanceClient) ReportReload(ctx context.Context, in *ReportReloadReq, opts ...grpc.CallOption) (*ReportReloadResp, error) {
	out := new(ReportReloadResp)
	err := grpc.Invoke(ctx, "/sidecar.Instance/ReportReload", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Instance service

type InstanceServer interface {
	Ping(context.Context, *PingReq) (*PingResp, error)
	Inject(context.Context, *InjectReq) (*InjectResp, error)
	WatchReload(*WatchReloadReq, Instance_WatchReloadServer) error
	ReportReload(context.Context, *ReportReloadReq) (*ReportReloadResp, error)
}

func RegisterInstanceServer(s *grpc.Server, srv InstanceServer) {
	s.RegisterService(&_Instance_serviceDesc, srv)
}

func _Instance_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sidecar.Instance/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).Ping(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_Inject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).Inject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sidecar.Instance/Inject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).Inject(ctx, req.(*InjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_WatchReload_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchReloadReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InstanceServer).WatchReload(m, &instanceWatchReloadServer{stream})
}

type Instance_WatchReloadServer interface {
	Send(*WatchReloadResp) error
	grpc.ServerStream
}

type instanceWatchReloadServer struct {
	grpc.ServerStream
}

func (x *instanceWatchReloadServer) Send(m *WatchReloadResp) error {
	return x.ServerStream.SendMsg(m)
}

func _Instance_ReportReload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportReloadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).ReportReload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sidecar.Instance/ReportReload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).ReportReload(ctx, req.(*ReportReloadReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Instance_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sidecar.Instance",
	HandlerType: (*InstanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Instance_Ping_Handler,
		},
		{
			MethodName: "Inject",
			Handler:    _Instance_Inject_Handler,
		},
		{
			MethodName: "ReportReload",
			Handler:    _Instance_ReportReload_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchReload",
			Handler:       _Instance_WatchReload_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sidecar.proto",
}

func init() { proto.RegisterFile("sidecar.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 2575 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe4, 0x59, 0xeb, 0x6f, 0x1b, 0xc7,
	0x11, 0xc7, 0xd1, 0xb2, 0x1e, 0x2b, 0xc7, 0xb2, 0xd7, 0x79, 0xd0, 0x8c, 0xeb, 0x32, 0x57, 0x14,
	0x91, 0x2f, 0xa6, 0x28, 0x5e, 0xe2, 0x48, 0x62, 0x9a, 0xd4, 0x27, 0xd7, 0x45, 0x95, 0xd4, 0xb1,
	0x70, 0x4e, 0xd0, 0x02, 0x41, 0x91, 0x9e, 0xc8, 0xb3, 0xc2, 0x98, 0x22, 0x19, 0x1e, 0xe5, 0x47,
	0x0d, 0x01, 0x52, 0xac, 0xf7, 0x8b, 0xec, 0x59, 0x96, 0x65, 0x4a, 0xb1, 0xad, 0x54, 0xb6, 0x15,
	0x27, 0x16, 0xe5, 0xc4, 0xb1, 0xa8, 0x87, 0x25, 0xa0, 0x1f, 0xfb, 0x51, 0xfd, 0x0f, 0x8a, 0xdb,
	0x25, 0x05, 0xb4, 0x75, 0xd0, 0x2f, 0x29, 0x0a, 0x14, 0x28, 0x6e, 0xf7, 0x48, 0x1e, 0x49, 0x39,
	0x4e, 0x9a, 0xb4, 0x01, 0x92, 0x4f, 0xe2, 0xcd, 0xce, 0xec, 0xfe, 0x76, 0x66, 0xf6, 0xb7, 0x33,
	0x2b, 0xf0, 0x88, 0xe2, 0xf3, 0xca, 0x1e, 0x29, 0x5c, 0x13, 0x0a, 0x07, 0x23, 0x41, 0x58, 0x66,
	0x7c, 0xda, 0x1c, 0x2d, 0x27, 0x1c, 0x2d, 0x8a, 0x27, 0xe4, 0xf4, 0x05, 0x22, 0x72, 0x38, 0x20,
	0xf9, 0x9d, 0x44, 0xc1, 0x13, 0xf4, 0x3b, 0x3d, 0xc1, 0xb6, 0xb6, 0x60, 0xc0, 0xf8, 0x43, 0xed,
	0x6c, 0x7b, 0x5a, 0x83, 0xc1, 0x56, 0xbf, 0xec, 0x94, 0x42, 0x3e, 0xa7, 0x14, 0x08, 0x04, 0x23,
	0x52, 0xc4, 0x17, 0x0c, 0x28, 0xc6, 0xe8, 0x7e, 0x6a, 0xeb, 0x68, 0x95, 0x03, 0x0e, 0xe5, 0x94,
	0xd4, 0xda, 0x2a, 0x87, 0x9d, 0xc1, 0x10, 0xd1, 0x28, 0xd6, 0x66, 0x2f, 0x96, 0x82, 0xb2, 0x23,
	0x41, 0x6f, 0x53, 0xe0, 0x78, 0x10, 0xbe, 0x08, 0xb6, 0xb5, 0xb4, 0x2b, 0xbe, 0x80, 0xac, 0x28,
	0xaf, 0x4a, 0x6d, 0xb2, 0x95, 0xb1, 0x33, 0xd5, 0x15, 0x8d, 0xbb, 0x55, 0xe1, 0x71, 0x0e, 0xe2,
	0x81, 0x4e, 0x74, 0xfb, 0xb2, 0x96, 0x8c, 0xa3, 0xc1, 0x69, 0x14, 0x1b, 0x49, 0xdd, 0x98, 0x3f,
	0x7d, 0x50, 0xcc, 0x53, 0x87, 0xcf, 0x82, 0x32, 0x29, 0x14, 0x22, 0x96, 0x96, 0x22, 0x4b, 0xb4,
	0x34, 0x96, 0x1a, 0x9b, 0xcd, 0x5a, 0x66, 0x34, 0xe1, 0x0b, 0xa0, 0xd2, 0xe3, 0x6f, 0x57, 0x22,
	0x72, 0x98, 0x18, 0x6e, 0x29, 0x32, 0xdc, 0x98, 0xec, 0x4d, 0xad, 0xce, 0x64, 0x0d, 0xcd, 0xda,
	0xf0, 0x00, 0x28, 0xff, 0x5d, 0x30, 0x20, 0x13, 0xcb, 0x92, 0xe2, 0x25, 0x67, 0x6e, 0xa0, 0xe1,
	0xa5, 0xac, 0x65, 0x56, 0x15, 0xee, 0x03, 0x16, 0xaf, 0xc7, 0xba, 0xd5, 0x64, 0x90, 0x1a, 0xf8,
	0x20, 0x15, 0xeb, 0xc5, 0x53, 0x4b, 0xb8, 0x7f, 0x0d, 0x5f, 0xed, 0x4b, 0x27, 0x7a, 0x4f, 0x1f,
	0x14, 0x2d, 0x5e, 0x0f, 0xdc, 0x0d, 0x2c, 0x4d, 0xcd, 0xd6, 0x52, 0xa2, 0x5a, 0xa1, 0x0a, 0xa5,
	0x9c, 0xa5, 0xa9, 0xf9, 0xb4, 0x5d, 0xb4, 0x34, 0x35, 0xc3, 0x66, 0x50, 0xea, 0x97, 0x5a, 0x64,
	0xbf, 0x62, 0x2d, 0xb3, 0x6f, 0xa9, 0xae, 0xe4, 0xf7, 0xd4, 0x64, 0xa2, 0x6b, 0xf8, 0xb3, 0xe6,
	0x97, 0x64, 0xf8, 0x70, 0x20, 0x12, 0x3e, 0xd3, 0x68, 0x53, 0x85, 0x27, 0x38, 0x68, 0x78, 0x61,
	0xee, 0xb2, 0xb6, 0x3a, 0x44, 0x8d, 0x4f, 0x77, 0x96, 0x8b, 0xc6, 0x3c, 0x50, 0x00, 0x25, 0x21,
	0x29, 0xf2, 0x96, 0xb5, 0x9c, 0x2c, 0xe7, 0x50, 0x05, 0x8e, 0xfb, 0xb1, 0x61, 0x31, 0x35, 0x6b,
	0x4c, 0x9e, 0x8a, 0x77, 0x6f, 0xf4, 0x8c, 0xa4, 0xee, 0xcd, 0xa5, 0x56, 0xce, 0xa3, 0x5b, 0x13,
	0xe9, 0x85, 0x04, 0x5a, 0xed, 0x3e, 0xdd, 0xc9, 0x88, 0xc4, 0x14, 0xfe, 0x89, 0x01, 0x65, 0x3e,
	0x45, 0x94, 0x25, 0xef, 0x19, 0x6b, 0x85, 0x9d, 0xa9, 0x2e, 0x6f, 0x4c, 0x30, 0xaa, 0xf0, 0x21,
	0xc3, 0x3d, 0x99, 0x4e, 0x5c, 0x13, 0x42, 0x21, 0x3c, 0x3b, 0x8d, 0xde, 0xbb, 0x88, 0x27, 0x12,
	0x28, 0x76, 0x1d, 0x2d, 0x7c, 0x8c, 0xe6, 0x6f, 0xa7, 0x96, 0x6f, 0xf2, 0x43, 0x0c, 0x9a, 0x9a,
	0xdd, 0xe8, 0x8c, 0xa3, 0x64, 0x12, 0x9f, 0xeb, 0x41, 0xbd, 0x9f, 0x6a, 0xc9, 0xa1, 0xfb, 0x2b,
	0xc3, 0x1b, 0x53, 0x9d, 0xe9, 0xeb, 0x5d, 0x34, 0xe6, 0x78, 0x7c, 0x11, 0xcd, 0x0d, 0xe3, 0xfe,
	0x18, 0xfe, 0x64, 0x16, 0xf5, 0x5c, 0x43, 0xb1, 0x51, 0x7d, 0xb6, 0x81, 0x11, 0x6d, 0x25, 0x4e,
	0x27, 0xc1, 0x43, 0x03, 0x28, 0x3a, 0x4e, 0xb1, 0xa1, 0xde, 0x1e, 0x34, 0xb7, 0x78, 0x7f, 0x65,
	0x38, 0x6b, 0x6b, 0x4c, 0xb5, 0xf6, 0x5e, 0x3a, 0xd1, 0x65, 0xcc, 0x40, 0x66, 0xd3, 0xee, 0xad,
	0xa1, 0xd5, 0x8b, 0xa9, 0xc1, 0xbb, 0xb8, 0xb3, 0x4b, 0x4b, 0x2e, 0x45, 0xc2, 0xed, 0xb2, 0x98,
	0xd9, 0x80, 0xad, 0x01, 0x54, 0x9a, 0x5c, 0x08, 0x77, 0x80, 0x2d, 0x27, 0xe4, 0x33, 0x34, 0x2b,
	0x45, 0xfd, 0x27, 0x7c, 0x14, 0x6c, 0x3d, 0x29, 0xf9, 0xdb, 0x8d, 0x7c, 0x13, 0xe9, 0x87, 0xdb,
	0x52, 0xcf, 0xb8, 0x05, 0x55, 0x78, 0x09, 0xfc, 0x84, 0xcb, 0xa4, 0x36, 0xef, 0x3a, 0x46, 0xdd,
	0x47, 0x5d, 0xaf, 0x25, 0x6f, 0xe1, 0x81, 0xb5, 0xf4, 0xbd, 0x7b, 0x68, 0xa6, 0x3b, 0x15, 0xeb,
	0x4d, 0xc5, 0xbb, 0xa9, 0x9b, 0xa9, 0x77, 0xb4, 0xb5, 0x69, 0xdc, 0x95, 0x60, 0x67, 0x18, 0x50,
	0x75, 0x28, 0x18, 0x38, 0xee, 0x6b, 0x55, 0x8e, 0xc8, 0x11, 0xc9, 0x2b, 0x45, 0x24, 0xf8, 0x34,
	0x28, 0x09, 0xe4, 0x4e, 0xc6, 0x2e, 0x55, 0xd8, 0xc1, 0x6d, 0x33, 0xb6, 0x9c, 0x49, 0x33, 0xa2,
	0x00, 0x1d, 0x60, 0xeb, 0x71, 0x12, 0x4b, 0x7a, 0x12, 0x9e, 0x50, 0x85, 0x47, 0xb9, 0x2a, 0x43,
	0xf3, 0x56, 0x2c, 0x13, 0xb5, 0x72, 0x91, 0x6a, 0xb9, 0x1b, 0x54, 0xe1, 0x79, 0xf0, 0x1c, 0x57,
	0xb8, 0x1e, 0xff, 0x14, 0xf5, 0x5b, 0x58, 0xf6, 0x07, 0x25, 0x6f, 0x36, 0xf4, 0x14, 0x20, 0x8e,
	0x46, 0xd3, 0x6b, 0xf3, 0xec, 0x87, 0x0c, 0x28, 0x6b, 0xf6, 0x05, 0x5a, 0x45, 0xf9, 0x1d, 0xd8,
	0x06, 0xb6, 0x28, 0xf2, 0x3b, 0x04, 0x5d, 0x49, 0xe3, 0x1b, 0xaa, 0xf0, 0x6b, 0xae, 0x2a, 0x9d,
	0x58, 0xc0, 0xb7, 0xdf, 0x45, 0x4b, 0x51, 0xd4, 0x7f, 0x11, 0x45, 0x17, 0xf8, 0xc3, 0xe9, 0xf9,
	0x73, 0xfa, 0x4e, 0xc7, 0x17, 0x53, 0xf1, 0xee, 0x82, 0xc1, 0xfb, 0x2b, 0xc3, 0xa9, 0xb1, 0x59,
	0x6d, 0x69, 0x14, 0x0f, 0xf7, 0xa1, 0xb9, 0x38, 0xbe, 0x78, 0x77, 0xe3, 0xe2, 0x1d, 0xd4, 0xdb,
	0x43, 0x4f, 0x04, 0xbe, 0x12, 0xc3, 0x1f, 0x4d, 0x53, 0x13, 0x51, 0x5f, 0xc7, 0xed, 0x52, 0x85,
	0x1a, 0xf0, 0x24, 0x97, 0x59, 0x9e, 0x07, 0xfa, 0x0f, 0xaa, 0xb0, 0xce, 0xe8, 0x1a, 0x7c, 0x15,
	0x7c, 0xe4, 0xac, 0x9d, 0x55, 0xe4, 0x77, 0x58, 0xb7, 0x9d, 0x75, 0xb1, 0x1d, 0xec, 0xfb, 0x5b,
	0x41, 0x39, 0x55, 0x57, 0x42, 0xb0, 0x8f, 0x31, 0xe3, 0x6d, 0x57, 0x85, 0x30, 0x57, 0x85, 0xce,
	0x8f, 0xa0, 0xa5, 0xb1, 0x1c, 0xde, 0x37, 0x0b, 0x04, 0xfb, 0xed, 0x5a, 0x72, 0xf4, 0x0b, 0xf6,
	0xa0, 0x25, 0x3b, 0xd3, 0x7d, 0x77, 0xbe, 0xe2, 0x4e, 0x60, 0x03, 0x28, 0x93, 0xc3, 0xe1, 0x43,
	0x41, 0x2f, 0x4d, 0xa5, 0xed, 0x7c, 0x55, 0x8d, 0xc1, 0xb8, 0x87, 0xa9, 0xb8, 0x71, 0xbb, 0x2a,
	0x54, 0x72, 0x15, 0x1b, 0x63, 0x97, 0xd2, 0x89, 0x44, 0xea, 0x6a, 0x97, 0x98, 0xd1, 0x87, 0xfb,
	0x40, 0xa9, 0x1c, 0x0e, 0x1f, 0x51, 0x5a, 0x0d, 0xee, 0xda, 0xa9, 0x0a, 0xdb, 0xb9, 0x6d, 0x54,
	0x91, 0x86, 0x48, 0x34, 0x14, 0xe0, 0x59, 0x50, 0xd2, 0x16, 0xf4, 0x2a, 0xd6, 0x12, 0xc2, 0x17,
	0x3b, 0x0a, 0xf9, 0xa2, 0xf1, 0x75, 0x55, 0x10, 0xb9, 0x6d, 0x46, 0x8a, 0x92, 0x18, 0xf3, 0x8d,
	0xb8, 0x3f, 0x86, 0x06, 0xaf, 0xa0, 0xd8, 0x68, 0x7a, 0x6d, 0x0c, 0x4d, 0x5e, 0x36, 0x0f, 0x56,
	0xa3, 0x91, 0x0b, 0x99, 0xb4, 0x8e, 0x26, 0xb4, 0xe5, 0x6b, 0xe6, 0x9c, 0x46, 0x33, 0x71, 0x2d,
	0x79, 0x53, 0x08, 0x85, 0xf6, 0x89, 0x64, 0x51, 0xf7, 0xac, 0x45, 0x15, 0x66, 0x2c, 0xc0, 0xc9,
	0x65, 0xfd, 0x4f, 0xe3, 0x45, 0xfd, 0x4b, 0xe3, 0xb5, 0xce, 0x64, 0xb6, 0xb5, 0xce, 0x18, 0xa0,
	0xf9, 0xcf, 0x18, 0xf8, 0x37, 0xc6, 0x1c, 0xc4, 0xfd, 0x76, 0xd6, 0x50, 0x62, 0xdd, 0xf6, 0x5a,
	0xfa, 0x75, 0x44, 0x69, 0xd5, 0x07, 0x8f, 0xbe, 0xa2, 0x8f, 0xea, 0xcb, 0xb1, 0x6e, 0xfb, 0x1b,
	0x67, 0x59, 0xf3, 0xf5, 0xa0, 0x8f, 0x67, 0xbe, 0x75, 0x2d, 0xe3, 0x02, 0xd0, 0xc5, 0x52, 0x28,
	0xa4, 0x4b, 0x4c, 0xcc, 0xae, 0x4b, 0x8d, 0x4f, 0x7d, 0x44, 0x67, 0x6e, 0x5d, 0x44, 0xfe, 0xee,
	0x67, 0xbd, 0x1e, 0xfd, 0xc3, 0xeb, 0xa9, 0x25, 0x60, 0x9a, 0x9a, 0x09, 0x2e, 0xbe, 0xae, 0xa6,
	0xb6, 0xa6, 0xb6, 0x86, 0x88, 0x28, 0x9f, 0xb2, 0x6e, 0xfb, 0x59, 0xf6, 0x84, 0x4b, 0x1f, 0x3d,
	0x49, 0xc4, 0x27, 0x78, 0xf2, 0x9b, 0x67, 0x3b, 0xf6, 0xdb, 0x59, 0xfd, 0xdc, 0xe9, 0x9f, 0x4e,
	0xfd, 0xa0, 0x39, 0x33, 0x17, 0x30, 0xdb, 0xf1, 0x9b, 0x0e, 0xf6, 0x2f, 0x25, 0xa0, 0xa2, 0x29,
	0xf0, 0xb6, 0xec, 0x89, 0xfc, 0xff, 0x0f, 0x15, 0xac, 0x2b, 0xb8, 0x84, 0x2d, 0x26, 0xaa, 0xf9,
	0xc2, 0xeb, 0xd7, 0x91, 0xbb, 0x7e, 0xb7, 0x98, 0x6c, 0x1e, 0x74, 0xf1, 0x9e, 0x63, 0xb2, 0xf7,
	0x17, 0xcd, 0xc7, 0xbd, 0xd9, 0x7c, 0xcc, 0xee, 0x3d, 0xef, 0x06, 0x7b, 0x45, 0x15, 0x7e, 0xc1,
	0x3d, 0x85, 0xa7, 0xae, 0xe0, 0xa9, 0x49, 0x4a, 0xe1, 0x3a, 0x6f, 0x92, 0x7c, 0xc4, 0x57, 0xfb,
	0x52, 0xb7, 0x56, 0x37, 0x26, 0x7b, 0x51, 0xac, 0x9f, 0xff, 0xd1, 0xcb, 0xc7, 0x8e, 0xbe, 0x6a,
	0x6c, 0x7f, 0x78, 0x1c, 0x77, 0x2e, 0x6f, 0x5c, 0x58, 0x43, 0x4b, 0xd7, 0x37, 0x2e, 0xc5, 0x50,
	0xff, 0x5d, 0x2d, 0xb9, 0xe4, 0xaa, 0xe5, 0x9f, 0xcb, 0x5c, 0x79, 0x5f, 0x87, 0xe2, 0x57, 0x18,
	0x55, 0x48, 0x32, 0xe0, 0x18, 0x97, 0x8b, 0x15, 0xbf, 0xd7, 0x8c, 0xc7, 0xa0, 0x49, 0x02, 0xd5,
	0xcc, 0x4a, 0xeb, 0x4c, 0x9e, 0xe3, 0xd6, 0x99, 0x8c, 0x4f, 0xd6, 0x19, 0x03, 0x16, 0x2f, 0xc1,
	0x37, 0x0b, 0x52, 0xfe, 0xab, 0xa5, 0xf2, 0x43, 0x13, 0xb0, 0x83, 0xfd, 0xdc, 0x02, 0x40, 0x06,
	0xf8, 0xf7, 0x8f, 0x0b, 0xdd, 0x67, 0x54, 0xe1, 0x24, 0x38, 0xcc, 0x99, 0x5c, 0xf0, 0xe0, 0xe0,
	0x3d, 0x8c, 0xa2, 0x1c, 0xf0, 0x99, 0x2f, 0x4f, 0x50, 0x1d, 0xec, 0x3f, 0x2d, 0x60, 0xfb, 0xaf,
	0xa4, 0x88, 0xe7, 0x2d, 0x91, 0x5c, 0xab, 0xdf, 0xdd, 0x33, 0xee, 0x1e, 0x65, 0x54, 0x61, 0x90,
	0x01, 0x4d, 0x5c, 0xc1, 0x7e, 0x79, 0x6b, 0x7a, 0xee, 0xfd, 0x8d, 0x89, 0x1e, 0x2a, 0xd0, 0x96,
	0x86, 0xb4, 0xe5, 0xbb, 0x5f, 0xea, 0x84, 0xf0, 0x2f, 0xc2, 0x17, 0xfe, 0xfb, 0x83, 0xd1, 0xc1,
	0xde, 0xad, 0x04, 0x55, 0x79, 0x50, 0xbe, 0x87, 0x55, 0xc0, 0x79, 0x06, 0x54, 0x84, 0x65, 0xbf,
	0x2c, 0x29, 0xb2, 0xcf, 0x6b, 0xb4, 0x2d, 0x64, 0xd7, 0x3b, 0xf1, 0xf8, 0x3c, 0x8a, 0xfe, 0x1e,
	0x25, 0xcf, 0xa5, 0x06, 0xfa, 0xf1, 0xd4, 0x47, 0x4d, 0x3f, 0xe3, 0xdf, 0xc8, 0x2a, 0x22, 0x75,
	0xb8, 0xad, 0xdd, 0x1f, 0xf1, 0x89, 0x19, 0x81, 0xb6, 0xa4, 0xe2, 0xf1, 0x6b, 0xe8, 0xd6, 0x04,
	0x9a, 0x9a, 0xd5, 0x6b, 0x75, 0x52, 0xbd, 0xa3, 0x85, 0xb9, 0x8d, 0xbe, 0x28, 0x2d, 0x17, 0xf1,
	0xd4, 0x12, 0x21, 0xdc, 0x21, 0x14, 0xbd, 0x49, 0xd5, 0xf2, 0x27, 0x38, 0x7d, 0x50, 0xcc, 0xe1,
	0x80, 0x1f, 0x30, 0x60, 0x7b, 0xfe, 0xb8, 0xd1, 0x20, 0x75, 0x32, 0xaa, 0xd0, 0xc1, 0xed, 0x29,
	0xc2, 0x56, 0x9d, 0x5a, 0xee, 0xd6, 0xeb, 0xe5, 0x95, 0xe8, 0xbe, 0xff, 0x35, 0xcc, 0x02, 0x60,
	0xf0, 0x25, 0x50, 0x69, 0x2c, 0x48, 0x4e, 0x02, 0xed, 0xce, 0xf6, 0xa8, 0xc2, 0x6e, 0xee, 0xb1,
	0x02, 0x98, 0xb9, 0xb6, 0xd1, 0x64, 0x00, 0xc3, 0x00, 0xd0, 0xb5, 0x5f, 0x3b, 0x13, 0x92, 0xad,
	0x65, 0x76, 0xa6, 0x7a, 0x6b, 0xa3, 0xa8, 0x0a, 0x47, 0xb9, 0x6d, 0x46, 0xa1, 0x7d, 0x7b, 0x19,
	0x5d, 0x1e, 0xe2, 0x7f, 0x5a, 0x7b, 0x7f, 0x25, 0x4e, 0x67, 0xc3, 0x93, 0x77, 0xf0, 0xf8, 0xbc,
	0xdd, 0xee, 0xd2, 0x25, 0x93, 0x97, 0xf1, 0x72, 0xdc, 0x6e, 0xd7, 0x92, 0xcb, 0xe9, 0xeb, 0x5d,
	0x78, 0x22, 0x61, 0xe4, 0x18, 0x89, 0xb2, 0x71, 0x96, 0xce, 0x0d, 0xe3, 0x4b, 0x09, 0xd1, 0xb4,
	0x8a, 0x1e, 0xf5, 0xf2, 0x70, 0x30, 0x18, 0x69, 0xce, 0x35, 0x78, 0xa7, 0x54, 0x21, 0xc2, 0x55,
	0xd1, 0x3e, 0x0e, 0x5f, 0x5d, 0x4c, 0x4d, 0xce, 0xa1, 0x7b, 0x17, 0x78, 0x09, 0x45, 0x6f, 0xe2,
	0x89, 0x84, 0x92, 0xe9, 0x57, 0x16, 0xd1, 0xa5, 0x59, 0xd4, 0xdb, 0x93, 0x8a, 0x77, 0x9b, 0x3b,
	0x3e, 0xea, 0xca, 0xf4, 0x5a, 0xdf, 0xc9, 0xa0, 0xbf, 0xbd, 0x4d, 0x41, 0xb1, 0x51, 0x34, 0x35,
	0x6b, 0x50, 0x07, 0x31, 0xd1, 0x92, 0xb7, 0xb4, 0x7b, 0xa4, 0xdd, 0x58, 0xbc, 0xab, 0x25, 0x47,
	0xb4, 0x64, 0x27, 0x9a, 0x8b, 0xa7, 0x26, 0x93, 0x28, 0x36, 0x2c, 0x66, 0x81, 0xc0, 0xdb, 0x0c,
	0xa8, 0x68, 0x33, 0x9a, 0x0d, 0xc5, 0x5a, 0x41, 0xea, 0x00, 0x6b, 0xb6, 0x0e, 0x28, 0xe8, 0x46,
	0x1a, 0x75, 0xce, 0xe6, 0xf6, 0xd2, 0x96, 0x04, 0x75, 0xc5, 0x37, 0xed, 0x4a, 0xf8, 0xd7, 0xf0,
	0xd5, 0x45, 0x3c, 0x32, 0x97, 0x4e, 0x5c, 0xa3, 0x02, 0xd4, 0x3f, 0x9f, 0x87, 0x6b, 0xf5, 0x46,
	0x2a, 0xde, 0x4d, 0x51, 0xd3, 0x1d, 0xeb, 0xcd, 0xd8, 0xd0, 0x1a, 0x8a, 0x8e, 0xd3, 0x53, 0x4a,
	0x27, 0xc3, 0x17, 0xe6, 0xf1, 0xc8, 0x1c, 0xea, 0x8a, 0x53, 0xcf, 0x8b, 0x39, 0x94, 0xee, 0x7f,
	0x58, 0x54, 0xe1, 0xef, 0x16, 0xf0, 0x32, 0x57, 0xc8, 0x24, 0xfc, 0xde, 0x62, 0x56, 0xc3, 0x9f,
	0x76, 0xa1, 0x95, 0xe8, 0xc3, 0x2e, 0x90, 0x2b, 0x16, 0x38, 0x65, 0xf9, 0x4a, 0x35, 0x6e, 0x36,
	0xe1, 0x75, 0xd1, 0xf1, 0xfa, 0xba, 0x06, 0xa9, 0xa5, 0xfe, 0xb8, 0xc3, 0x53, 0x27, 0x1f, 0x70,
	0xb8, 0x5c, 0x72, 0x83, 0xa3, 0xc1, 0xeb, 0x39, 0xe0, 0x38, 0xc0, 0x1f, 0x78, 0xae, 0xb6, 0xf6,
	0x78, 0x43, 0x03, 0x5f, 0x57, 0x4f, 0x0a, 0xe3, 0xbc, 0x34, 0xce, 0x5a, 0xf2, 0xf5, 0xfc, 0xc3,
	0x2c, 0x4d, 0xd9, 0xcb, 0xba, 0xd9, 0x80, 0x7c, 0xca, 0x6e, 0x48, 0x8c, 0x41, 0x23, 0xcb, 0x0c,
	0xb0, 0x99, 0xf0, 0x16, 0x95, 0xb7, 0x04, 0x44, 0xc6, 0x91, 0xb4, 0x44, 0x37, 0xa6, 0x54, 0xe4,
	0xf0, 0x49, 0x39, 0x5c, 0x73, 0x46, 0x6a, 0xf3, 0xeb, 0x4a, 0x3f, 0xa7, 0xd5, 0x31, 0xeb, 0x94,
	0x23, 0x1e, 0x52, 0x12, 0x5f, 0xaa, 0x00, 0x55, 0xa2, 0x1c, 0x0a, 0x86, 0x23, 0xdf, 0xf5, 0x4b,
	0x13, 0xce, 0x6c, 0xc2, 0xcf, 0x3d, 0x8c, 0x2a, 0x74, 0x31, 0x9c, 0x1d, 0x2d, 0x7c, 0x8c, 0x07,
	0x6e, 0xa4, 0xa7, 0x87, 0x53, 0x63, 0x57, 0xf0, 0x85, 0x7e, 0xbd, 0x00, 0xfe, 0xb6, 0x08, 0x7b,
	0xf1, 0x41, 0x84, 0x1d, 0x63, 0x54, 0x61, 0x98, 0xe1, 0x9e, 0x79, 0x18, 0xd6, 0x6f, 0x91, 0xc0,
	0x7f, 0x9b, 0x25, 0x60, 0x5f, 0x96, 0xbf, 0x0f, 0xaa, 0xc2, 0x8b, 0x1c, 0xa4, 0xe0, 0x68, 0x8e,
	0xd1, 0x8c, 0xe0, 0x9f, 0xa6, 0xb2, 0xb0, 0x49, 0xb6, 0xdf, 0xce, 0xd7, 0xba, 0x1a, 0x1c, 0xb5,
	0xf5, 0x0e, 0xbe, 0xc1, 0xee, 0xaa, 0x73, 0xbb, 0xea, 0xdd, 0x3c, 0x2f, 0x9a, 0xe6, 0x84, 0xee,
	0xcc, 0x0a, 0xe4, 0x36, 0xa7, 0x14, 0x4f, 0x9e, 0xe0, 0x1e, 0x33, 0xe6, 0x26, 0x9e, 0xc9, 0x5d,
	0xe4, 0x26, 0x6d, 0xf8, 0x02, 0x89, 0x7f, 0x50, 0xf2, 0xea, 0xd7, 0x39, 0xa5, 0xea, 0x1f, 0xa8,
	0x82, 0x8d, 0xb3, 0x16, 0x9b, 0x52, 0x8a, 0x13, 0x73, 0xfa, 0xee, 0xcf, 0x2c, 0xaa, 0xf0, 0x57,
	0x0b, 0x68, 0xe3, 0x0a, 0x8f, 0x0b, 0x9f, 0x67, 0xae, 0x25, 0x07, 0xf1, 0xe0, 0x97, 0xed, 0x4a,
	0x4c, 0x1b, 0x5b, 0x67, 0x76, 0xe5, 0x90, 0x66, 0x57, 0xe5, 0xff, 0xc5, 0xc0, 0xcf, 0x99, 0xaf,
	0xd5, 0xb1, 0x7c, 0x43, 0x84, 0x56, 0xdb, 0xf2, 0x7c, 0x9d, 0x54, 0xd7, 0x50, 0xef, 0x90, 0x1b,
	0x3c, 0x2e, 0x6a, 0x59, 0xef, 0xe1, 0x9f, 0xdd, 0x8c, 0xd0, 0x8c, 0x1d, 0xe9, 0x56, 0x9b, 0x44,
	0x33, 0xa7, 0x64, 0x62, 0xe1, 0xec, 0x86, 0x59, 0x37, 0x7b, 0xec, 0xf5, 0x43, 0x87, 0x0e, 0x1f,
	0x3b, 0xc6, 0x76, 0xb0, 0xff, 0xb6, 0x80, 0x1d, 0xf9, 0xce, 0xfe, 0xfe, 0xf6, 0x53, 0x45, 0x8e,
	0xd8, 0x2c, 0xed, 0xbe, 0xd9, 0x7e, 0x8a, 0xff, 0x63, 0x09, 0x28, 0x6f, 0x0a, 0x28, 0x11, 0x29,
	0xe0, 0x91, 0xe1, 0xbb, 0x0c, 0x28, 0x69, 0xf6, 0x05, 0x5a, 0x61, 0xee, 0x79, 0xcb, 0x78, 0x1e,
	0xb4, 0xed, 0x2c, 0x90, 0x28, 0x21, 0xf6, 0x75, 0x55, 0x78, 0x1e, 0xee, 0x4e, 0x8f, 0x2e, 0xa0,
	0xe8, 0xb8, 0x92, 0xf7, 0x42, 0x4b, 0x0e, 0x96, 0xed, 0xc1, 0x43, 0xef, 0xfe, 0x59, 0x3b, 0x6f,
	0x79, 0x9c, 0xdd, 0xe9, 0x3c, 0xe9, 0x72, 0xfa, 0x8c, 0xc5, 0x9d, 0x21, 0x5f, 0xa0, 0xd5, 0xcd,
	0x70, 0x3a, 0x88, 0x52, 0xda, 0x58, 0x42, 0x58, 0xfc, 0xaa, 0x61, 0xdb, 0x55, 0x24, 0x53, 0x42,
	0xec, 0x51, 0x55, 0xa8, 0x81, 0x30, 0xf3, 0x36, 0x9d, 0x6b, 0x44, 0x6d, 0xd6, 0x62, 0x99, 0x09,
	0x82, 0x95, 0xdd, 0x95, 0x07, 0xc1, 0x47, 0x66, 0xd4, 0x41, 0x9c, 0x67, 0x40, 0xa5, 0xa9, 0x42,
	0x81, 0x4f, 0x64, 0x57, 0xcd, 0x6f, 0xc6, 0x6c, 0xd6, 0xcd, 0x07, 0x28, 0xa6, 0x6a, 0x08, 0x69,
	0x55, 0x13, 0x36, 0x55, 0x35, 0xb6, 0x4d, 0x64, 0x04, 0xcd, 0x5e, 0x76, 0x77, 0x1e, 0x9a, 0x53,
	0xfa, 0x84, 0x0e, 0xaa, 0xe6, 0x66, 0xb8, 0x5a, 0x06, 0x5e, 0x60, 0xc0, 0x36, 0x73, 0x8e, 0xc0,
	0xdc, 0xea, 0x05, 0x84, 0x65, 0xdb, 0xfd, 0x80, 0x11, 0x25, 0xc4, 0x36, 0x13, 0x60, 0x34, 0x9b,
	0x8c, 0x72, 0x8b, 0x06, 0x6c, 0x13, 0x19, 0x01, 0xf6, 0x43, 0xd6, 0x96, 0x07, 0x2c, 0x4c, 0x66,
	0xcc, 0x21, 0x6b, 0x6c, 0x53, 0x85, 0xb7, 0xa1, 0x00, 0xea, 0xd3, 0xe7, 0xdf, 0xdb, 0xf8, 0x38,
	0x89, 0xa7, 0x46, 0xd0, 0xe0, 0xb4, 0xf1, 0x5e, 0xbe, 0xf8, 0x09, 0x8a, 0xce, 0x1b, 0x0f, 0x9d,
	0xf6, 0x4c, 0xa2, 0xd9, 0x85, 0xe6, 0xa6, 0x74, 0xe2, 0x0e, 0x9e, 0x18, 0xc5, 0xe3, 0x7d, 0x78,
	0xfa, 0x0f, 0xfc, 0x56, 0xf2, 0xfa, 0xc7, 0x31, 0x0c, 0xbf, 0x43, 0x0a, 0x85, 0xfc, 0x3e, 0x0f,
	0xf9, 0x37, 0x97, 0xf3, 0x6d, 0x25, 0x18, 0x70, 0x17, 0x49, 0x5a, 0x4a, 0xc9, 0x7f, 0xbf, 0x9e,
	0xfd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x44, 0xe6, 0x1b, 0x92, 0x1b, 0x00, 0x00,
}