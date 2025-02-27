// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: remote/kv.proto

package remote

import (
	context "context"
	types "github.com/ledgerwatch/erigon-lib/gointerfaces/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// KVClient is the client API for KV service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KVClient interface {
	// Version returns the service version number
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*types.VersionReply, error)
	// Tx exposes read-only transactions for the key-value store
	//
	// When tx open, client must receive 1 message from server with txID
	// When cursor open, client must receive 1 message from server with cursorID
	// Then only client can initiate messages from server
	Tx(ctx context.Context, opts ...grpc.CallOption) (KV_TxClient, error)
	StateChanges(ctx context.Context, in *StateChangeRequest, opts ...grpc.CallOption) (KV_StateChangesClient, error)
	// Snapshots returns list of current snapshot files. Then client can just open all of them.
	Snapshots(ctx context.Context, in *SnapshotsRequest, opts ...grpc.CallOption) (*SnapshotsReply, error)
	// Temporal methods
	DomainGet(ctx context.Context, in *DomainGetReq, opts ...grpc.CallOption) (*DomainGetReply, error)
	HistoryGet(ctx context.Context, in *HistoryGetReq, opts ...grpc.CallOption) (*HistoryGetReply, error)
	IndexRange(ctx context.Context, in *IndexRangeReq, opts ...grpc.CallOption) (KV_IndexRangeClient, error)
	Range(ctx context.Context, in *RangeReq, opts ...grpc.CallOption) (KV_RangeClient, error)
}

type kVClient struct {
	cc grpc.ClientConnInterface
}

func NewKVClient(cc grpc.ClientConnInterface) KVClient {
	return &kVClient{cc}
}

func (c *kVClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*types.VersionReply, error) {
	out := new(types.VersionReply)
	err := c.cc.Invoke(ctx, "/remote.KV/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) Tx(ctx context.Context, opts ...grpc.CallOption) (KV_TxClient, error) {
	stream, err := c.cc.NewStream(ctx, &KV_ServiceDesc.Streams[0], "/remote.KV/Tx", opts...)
	if err != nil {
		return nil, err
	}
	x := &kVTxClient{stream}
	return x, nil
}

type KV_TxClient interface {
	Send(*Cursor) error
	Recv() (*Pair, error)
	grpc.ClientStream
}

type kVTxClient struct {
	grpc.ClientStream
}

func (x *kVTxClient) Send(m *Cursor) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kVTxClient) Recv() (*Pair, error) {
	m := new(Pair)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kVClient) StateChanges(ctx context.Context, in *StateChangeRequest, opts ...grpc.CallOption) (KV_StateChangesClient, error) {
	stream, err := c.cc.NewStream(ctx, &KV_ServiceDesc.Streams[1], "/remote.KV/StateChanges", opts...)
	if err != nil {
		return nil, err
	}
	x := &kVStateChangesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KV_StateChangesClient interface {
	Recv() (*StateChangeBatch, error)
	grpc.ClientStream
}

type kVStateChangesClient struct {
	grpc.ClientStream
}

func (x *kVStateChangesClient) Recv() (*StateChangeBatch, error) {
	m := new(StateChangeBatch)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kVClient) Snapshots(ctx context.Context, in *SnapshotsRequest, opts ...grpc.CallOption) (*SnapshotsReply, error) {
	out := new(SnapshotsReply)
	err := c.cc.Invoke(ctx, "/remote.KV/Snapshots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) DomainGet(ctx context.Context, in *DomainGetReq, opts ...grpc.CallOption) (*DomainGetReply, error) {
	out := new(DomainGetReply)
	err := c.cc.Invoke(ctx, "/remote.KV/DomainGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) HistoryGet(ctx context.Context, in *HistoryGetReq, opts ...grpc.CallOption) (*HistoryGetReply, error) {
	out := new(HistoryGetReply)
	err := c.cc.Invoke(ctx, "/remote.KV/HistoryGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) IndexRange(ctx context.Context, in *IndexRangeReq, opts ...grpc.CallOption) (KV_IndexRangeClient, error) {
	stream, err := c.cc.NewStream(ctx, &KV_ServiceDesc.Streams[2], "/remote.KV/IndexRange", opts...)
	if err != nil {
		return nil, err
	}
	x := &kVIndexRangeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KV_IndexRangeClient interface {
	Recv() (*IndexRangeReply, error)
	grpc.ClientStream
}

type kVIndexRangeClient struct {
	grpc.ClientStream
}

func (x *kVIndexRangeClient) Recv() (*IndexRangeReply, error) {
	m := new(IndexRangeReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kVClient) Range(ctx context.Context, in *RangeReq, opts ...grpc.CallOption) (KV_RangeClient, error) {
	stream, err := c.cc.NewStream(ctx, &KV_ServiceDesc.Streams[3], "/remote.KV/Range", opts...)
	if err != nil {
		return nil, err
	}
	x := &kVRangeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KV_RangeClient interface {
	Recv() (*Pairs, error)
	grpc.ClientStream
}

type kVRangeClient struct {
	grpc.ClientStream
}

func (x *kVRangeClient) Recv() (*Pairs, error) {
	m := new(Pairs)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KVServer is the server API for KV service.
// All implementations must embed UnimplementedKVServer
// for forward compatibility
type KVServer interface {
	// Version returns the service version number
	Version(context.Context, *emptypb.Empty) (*types.VersionReply, error)
	// Tx exposes read-only transactions for the key-value store
	//
	// When tx open, client must receive 1 message from server with txID
	// When cursor open, client must receive 1 message from server with cursorID
	// Then only client can initiate messages from server
	Tx(KV_TxServer) error
	StateChanges(*StateChangeRequest, KV_StateChangesServer) error
	// Snapshots returns list of current snapshot files. Then client can just open all of them.
	Snapshots(context.Context, *SnapshotsRequest) (*SnapshotsReply, error)
	// Temporal methods
	DomainGet(context.Context, *DomainGetReq) (*DomainGetReply, error)
	HistoryGet(context.Context, *HistoryGetReq) (*HistoryGetReply, error)
	IndexRange(*IndexRangeReq, KV_IndexRangeServer) error
	Range(*RangeReq, KV_RangeServer) error
	mustEmbedUnimplementedKVServer()
}

// UnimplementedKVServer must be embedded to have forward compatible implementations.
type UnimplementedKVServer struct {
}

func (UnimplementedKVServer) Version(context.Context, *emptypb.Empty) (*types.VersionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedKVServer) Tx(KV_TxServer) error {
	return status.Errorf(codes.Unimplemented, "method Tx not implemented")
}
func (UnimplementedKVServer) StateChanges(*StateChangeRequest, KV_StateChangesServer) error {
	return status.Errorf(codes.Unimplemented, "method StateChanges not implemented")
}
func (UnimplementedKVServer) Snapshots(context.Context, *SnapshotsRequest) (*SnapshotsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Snapshots not implemented")
}
func (UnimplementedKVServer) DomainGet(context.Context, *DomainGetReq) (*DomainGetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DomainGet not implemented")
}
func (UnimplementedKVServer) HistoryGet(context.Context, *HistoryGetReq) (*HistoryGetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HistoryGet not implemented")
}
func (UnimplementedKVServer) IndexRange(*IndexRangeReq, KV_IndexRangeServer) error {
	return status.Errorf(codes.Unimplemented, "method IndexRange not implemented")
}
func (UnimplementedKVServer) Range(*RangeReq, KV_RangeServer) error {
	return status.Errorf(codes.Unimplemented, "method Range not implemented")
}
func (UnimplementedKVServer) mustEmbedUnimplementedKVServer() {}

// UnsafeKVServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KVServer will
// result in compilation errors.
type UnsafeKVServer interface {
	mustEmbedUnimplementedKVServer()
}

func RegisterKVServer(s grpc.ServiceRegistrar, srv KVServer) {
	s.RegisterService(&KV_ServiceDesc, srv)
}

func _KV_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.KV/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_Tx_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KVServer).Tx(&kVTxServer{stream})
}

type KV_TxServer interface {
	Send(*Pair) error
	Recv() (*Cursor, error)
	grpc.ServerStream
}

type kVTxServer struct {
	grpc.ServerStream
}

func (x *kVTxServer) Send(m *Pair) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kVTxServer) Recv() (*Cursor, error) {
	m := new(Cursor)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _KV_StateChanges_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StateChangeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KVServer).StateChanges(m, &kVStateChangesServer{stream})
}

type KV_StateChangesServer interface {
	Send(*StateChangeBatch) error
	grpc.ServerStream
}

type kVStateChangesServer struct {
	grpc.ServerStream
}

func (x *kVStateChangesServer) Send(m *StateChangeBatch) error {
	return x.ServerStream.SendMsg(m)
}

func _KV_Snapshots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnapshotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Snapshots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.KV/Snapshots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Snapshots(ctx, req.(*SnapshotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_DomainGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DomainGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).DomainGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.KV/DomainGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).DomainGet(ctx, req.(*DomainGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_HistoryGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HistoryGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).HistoryGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.KV/HistoryGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).HistoryGet(ctx, req.(*HistoryGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_IndexRange_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(IndexRangeReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KVServer).IndexRange(m, &kVIndexRangeServer{stream})
}

type KV_IndexRangeServer interface {
	Send(*IndexRangeReply) error
	grpc.ServerStream
}

type kVIndexRangeServer struct {
	grpc.ServerStream
}

func (x *kVIndexRangeServer) Send(m *IndexRangeReply) error {
	return x.ServerStream.SendMsg(m)
}

func _KV_Range_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RangeReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KVServer).Range(m, &kVRangeServer{stream})
}

type KV_RangeServer interface {
	Send(*Pairs) error
	grpc.ServerStream
}

type kVRangeServer struct {
	grpc.ServerStream
}

func (x *kVRangeServer) Send(m *Pairs) error {
	return x.ServerStream.SendMsg(m)
}

// KV_ServiceDesc is the grpc.ServiceDesc for KV service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KV_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "remote.KV",
	HandlerType: (*KVServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _KV_Version_Handler,
		},
		{
			MethodName: "Snapshots",
			Handler:    _KV_Snapshots_Handler,
		},
		{
			MethodName: "DomainGet",
			Handler:    _KV_DomainGet_Handler,
		},
		{
			MethodName: "HistoryGet",
			Handler:    _KV_HistoryGet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tx",
			Handler:       _KV_Tx_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StateChanges",
			Handler:       _KV_StateChanges_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "IndexRange",
			Handler:       _KV_IndexRange_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Range",
			Handler:       _KV_Range_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "remote/kv.proto",
}
