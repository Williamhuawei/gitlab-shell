// Code generated by protoc-gen-go. DO NOT EDIT.
// source: diff.proto

package gitaly

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CommitDiffRequest struct {
	Repository             *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	LeftCommitId           string      `protobuf:"bytes,2,opt,name=left_commit_id,json=leftCommitId" json:"left_commit_id,omitempty"`
	RightCommitId          string      `protobuf:"bytes,3,opt,name=right_commit_id,json=rightCommitId" json:"right_commit_id,omitempty"`
	IgnoreWhitespaceChange bool        `protobuf:"varint,4,opt,name=ignore_whitespace_change,json=ignoreWhitespaceChange" json:"ignore_whitespace_change,omitempty"`
	Paths                  [][]byte    `protobuf:"bytes,5,rep,name=paths,proto3" json:"paths,omitempty"`
	CollapseDiffs          bool        `protobuf:"varint,6,opt,name=collapse_diffs,json=collapseDiffs" json:"collapse_diffs,omitempty"`
	EnforceLimits          bool        `protobuf:"varint,7,opt,name=enforce_limits,json=enforceLimits" json:"enforce_limits,omitempty"`
	MaxFiles               int32       `protobuf:"varint,8,opt,name=max_files,json=maxFiles" json:"max_files,omitempty"`
	MaxLines               int32       `protobuf:"varint,9,opt,name=max_lines,json=maxLines" json:"max_lines,omitempty"`
	MaxBytes               int32       `protobuf:"varint,10,opt,name=max_bytes,json=maxBytes" json:"max_bytes,omitempty"`
	SafeMaxFiles           int32       `protobuf:"varint,11,opt,name=safe_max_files,json=safeMaxFiles" json:"safe_max_files,omitempty"`
	SafeMaxLines           int32       `protobuf:"varint,12,opt,name=safe_max_lines,json=safeMaxLines" json:"safe_max_lines,omitempty"`
	SafeMaxBytes           int32       `protobuf:"varint,13,opt,name=safe_max_bytes,json=safeMaxBytes" json:"safe_max_bytes,omitempty"`
}

func (m *CommitDiffRequest) Reset()                    { *m = CommitDiffRequest{} }
func (m *CommitDiffRequest) String() string            { return proto.CompactTextString(m) }
func (*CommitDiffRequest) ProtoMessage()               {}
func (*CommitDiffRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *CommitDiffRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *CommitDiffRequest) GetLeftCommitId() string {
	if m != nil {
		return m.LeftCommitId
	}
	return ""
}

func (m *CommitDiffRequest) GetRightCommitId() string {
	if m != nil {
		return m.RightCommitId
	}
	return ""
}

func (m *CommitDiffRequest) GetIgnoreWhitespaceChange() bool {
	if m != nil {
		return m.IgnoreWhitespaceChange
	}
	return false
}

func (m *CommitDiffRequest) GetPaths() [][]byte {
	if m != nil {
		return m.Paths
	}
	return nil
}

func (m *CommitDiffRequest) GetCollapseDiffs() bool {
	if m != nil {
		return m.CollapseDiffs
	}
	return false
}

func (m *CommitDiffRequest) GetEnforceLimits() bool {
	if m != nil {
		return m.EnforceLimits
	}
	return false
}

func (m *CommitDiffRequest) GetMaxFiles() int32 {
	if m != nil {
		return m.MaxFiles
	}
	return 0
}

func (m *CommitDiffRequest) GetMaxLines() int32 {
	if m != nil {
		return m.MaxLines
	}
	return 0
}

func (m *CommitDiffRequest) GetMaxBytes() int32 {
	if m != nil {
		return m.MaxBytes
	}
	return 0
}

func (m *CommitDiffRequest) GetSafeMaxFiles() int32 {
	if m != nil {
		return m.SafeMaxFiles
	}
	return 0
}

func (m *CommitDiffRequest) GetSafeMaxLines() int32 {
	if m != nil {
		return m.SafeMaxLines
	}
	return 0
}

func (m *CommitDiffRequest) GetSafeMaxBytes() int32 {
	if m != nil {
		return m.SafeMaxBytes
	}
	return 0
}

// A CommitDiffResponse corresponds to a single changed file in a commit.
type CommitDiffResponse struct {
	FromPath []byte `protobuf:"bytes,1,opt,name=from_path,json=fromPath,proto3" json:"from_path,omitempty"`
	ToPath   []byte `protobuf:"bytes,2,opt,name=to_path,json=toPath,proto3" json:"to_path,omitempty"`
	// Blob ID as returned via `git diff --full-index`
	FromId       string `protobuf:"bytes,3,opt,name=from_id,json=fromId" json:"from_id,omitempty"`
	ToId         string `protobuf:"bytes,4,opt,name=to_id,json=toId" json:"to_id,omitempty"`
	OldMode      int32  `protobuf:"varint,5,opt,name=old_mode,json=oldMode" json:"old_mode,omitempty"`
	NewMode      int32  `protobuf:"varint,6,opt,name=new_mode,json=newMode" json:"new_mode,omitempty"`
	Binary       bool   `protobuf:"varint,7,opt,name=binary" json:"binary,omitempty"`
	RawPatchData []byte `protobuf:"bytes,9,opt,name=raw_patch_data,json=rawPatchData,proto3" json:"raw_patch_data,omitempty"`
	EndOfPatch   bool   `protobuf:"varint,10,opt,name=end_of_patch,json=endOfPatch" json:"end_of_patch,omitempty"`
	// Indicates the diff file at which we overflow according to the limitations sent,
	// in which case only this attribute will be set.
	OverflowMarker bool `protobuf:"varint,11,opt,name=overflow_marker,json=overflowMarker" json:"overflow_marker,omitempty"`
	Collapsed      bool `protobuf:"varint,12,opt,name=collapsed" json:"collapsed,omitempty"`
}

func (m *CommitDiffResponse) Reset()                    { *m = CommitDiffResponse{} }
func (m *CommitDiffResponse) String() string            { return proto.CompactTextString(m) }
func (*CommitDiffResponse) ProtoMessage()               {}
func (*CommitDiffResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *CommitDiffResponse) GetFromPath() []byte {
	if m != nil {
		return m.FromPath
	}
	return nil
}

func (m *CommitDiffResponse) GetToPath() []byte {
	if m != nil {
		return m.ToPath
	}
	return nil
}

func (m *CommitDiffResponse) GetFromId() string {
	if m != nil {
		return m.FromId
	}
	return ""
}

func (m *CommitDiffResponse) GetToId() string {
	if m != nil {
		return m.ToId
	}
	return ""
}

func (m *CommitDiffResponse) GetOldMode() int32 {
	if m != nil {
		return m.OldMode
	}
	return 0
}

func (m *CommitDiffResponse) GetNewMode() int32 {
	if m != nil {
		return m.NewMode
	}
	return 0
}

func (m *CommitDiffResponse) GetBinary() bool {
	if m != nil {
		return m.Binary
	}
	return false
}

func (m *CommitDiffResponse) GetRawPatchData() []byte {
	if m != nil {
		return m.RawPatchData
	}
	return nil
}

func (m *CommitDiffResponse) GetEndOfPatch() bool {
	if m != nil {
		return m.EndOfPatch
	}
	return false
}

func (m *CommitDiffResponse) GetOverflowMarker() bool {
	if m != nil {
		return m.OverflowMarker
	}
	return false
}

func (m *CommitDiffResponse) GetCollapsed() bool {
	if m != nil {
		return m.Collapsed
	}
	return false
}

type CommitDeltaRequest struct {
	Repository    *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	LeftCommitId  string      `protobuf:"bytes,2,opt,name=left_commit_id,json=leftCommitId" json:"left_commit_id,omitempty"`
	RightCommitId string      `protobuf:"bytes,3,opt,name=right_commit_id,json=rightCommitId" json:"right_commit_id,omitempty"`
	Paths         [][]byte    `protobuf:"bytes,4,rep,name=paths,proto3" json:"paths,omitempty"`
}

func (m *CommitDeltaRequest) Reset()                    { *m = CommitDeltaRequest{} }
func (m *CommitDeltaRequest) String() string            { return proto.CompactTextString(m) }
func (*CommitDeltaRequest) ProtoMessage()               {}
func (*CommitDeltaRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *CommitDeltaRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *CommitDeltaRequest) GetLeftCommitId() string {
	if m != nil {
		return m.LeftCommitId
	}
	return ""
}

func (m *CommitDeltaRequest) GetRightCommitId() string {
	if m != nil {
		return m.RightCommitId
	}
	return ""
}

func (m *CommitDeltaRequest) GetPaths() [][]byte {
	if m != nil {
		return m.Paths
	}
	return nil
}

type CommitDelta struct {
	FromPath []byte `protobuf:"bytes,1,opt,name=from_path,json=fromPath,proto3" json:"from_path,omitempty"`
	ToPath   []byte `protobuf:"bytes,2,opt,name=to_path,json=toPath,proto3" json:"to_path,omitempty"`
	// Blob ID as returned via `git diff --full-index`
	FromId  string `protobuf:"bytes,3,opt,name=from_id,json=fromId" json:"from_id,omitempty"`
	ToId    string `protobuf:"bytes,4,opt,name=to_id,json=toId" json:"to_id,omitempty"`
	OldMode int32  `protobuf:"varint,5,opt,name=old_mode,json=oldMode" json:"old_mode,omitempty"`
	NewMode int32  `protobuf:"varint,6,opt,name=new_mode,json=newMode" json:"new_mode,omitempty"`
}

func (m *CommitDelta) Reset()                    { *m = CommitDelta{} }
func (m *CommitDelta) String() string            { return proto.CompactTextString(m) }
func (*CommitDelta) ProtoMessage()               {}
func (*CommitDelta) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *CommitDelta) GetFromPath() []byte {
	if m != nil {
		return m.FromPath
	}
	return nil
}

func (m *CommitDelta) GetToPath() []byte {
	if m != nil {
		return m.ToPath
	}
	return nil
}

func (m *CommitDelta) GetFromId() string {
	if m != nil {
		return m.FromId
	}
	return ""
}

func (m *CommitDelta) GetToId() string {
	if m != nil {
		return m.ToId
	}
	return ""
}

func (m *CommitDelta) GetOldMode() int32 {
	if m != nil {
		return m.OldMode
	}
	return 0
}

func (m *CommitDelta) GetNewMode() int32 {
	if m != nil {
		return m.NewMode
	}
	return 0
}

type CommitDeltaResponse struct {
	Deltas []*CommitDelta `protobuf:"bytes,1,rep,name=deltas" json:"deltas,omitempty"`
}

func (m *CommitDeltaResponse) Reset()                    { *m = CommitDeltaResponse{} }
func (m *CommitDeltaResponse) String() string            { return proto.CompactTextString(m) }
func (*CommitDeltaResponse) ProtoMessage()               {}
func (*CommitDeltaResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *CommitDeltaResponse) GetDeltas() []*CommitDelta {
	if m != nil {
		return m.Deltas
	}
	return nil
}

type CommitPatchRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	Revision   []byte      `protobuf:"bytes,2,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (m *CommitPatchRequest) Reset()                    { *m = CommitPatchRequest{} }
func (m *CommitPatchRequest) String() string            { return proto.CompactTextString(m) }
func (*CommitPatchRequest) ProtoMessage()               {}
func (*CommitPatchRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *CommitPatchRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *CommitPatchRequest) GetRevision() []byte {
	if m != nil {
		return m.Revision
	}
	return nil
}

type CommitPatchResponse struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *CommitPatchResponse) Reset()                    { *m = CommitPatchResponse{} }
func (m *CommitPatchResponse) String() string            { return proto.CompactTextString(m) }
func (*CommitPatchResponse) ProtoMessage()               {}
func (*CommitPatchResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *CommitPatchResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type RawDiffRequest struct {
	Repository    *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	LeftCommitId  string      `protobuf:"bytes,2,opt,name=left_commit_id,json=leftCommitId" json:"left_commit_id,omitempty"`
	RightCommitId string      `protobuf:"bytes,3,opt,name=right_commit_id,json=rightCommitId" json:"right_commit_id,omitempty"`
}

func (m *RawDiffRequest) Reset()                    { *m = RawDiffRequest{} }
func (m *RawDiffRequest) String() string            { return proto.CompactTextString(m) }
func (*RawDiffRequest) ProtoMessage()               {}
func (*RawDiffRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *RawDiffRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *RawDiffRequest) GetLeftCommitId() string {
	if m != nil {
		return m.LeftCommitId
	}
	return ""
}

func (m *RawDiffRequest) GetRightCommitId() string {
	if m != nil {
		return m.RightCommitId
	}
	return ""
}

type RawDiffResponse struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *RawDiffResponse) Reset()                    { *m = RawDiffResponse{} }
func (m *RawDiffResponse) String() string            { return proto.CompactTextString(m) }
func (*RawDiffResponse) ProtoMessage()               {}
func (*RawDiffResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *RawDiffResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type RawPatchRequest struct {
	Repository    *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	LeftCommitId  string      `protobuf:"bytes,2,opt,name=left_commit_id,json=leftCommitId" json:"left_commit_id,omitempty"`
	RightCommitId string      `protobuf:"bytes,3,opt,name=right_commit_id,json=rightCommitId" json:"right_commit_id,omitempty"`
}

func (m *RawPatchRequest) Reset()                    { *m = RawPatchRequest{} }
func (m *RawPatchRequest) String() string            { return proto.CompactTextString(m) }
func (*RawPatchRequest) ProtoMessage()               {}
func (*RawPatchRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *RawPatchRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *RawPatchRequest) GetLeftCommitId() string {
	if m != nil {
		return m.LeftCommitId
	}
	return ""
}

func (m *RawPatchRequest) GetRightCommitId() string {
	if m != nil {
		return m.RightCommitId
	}
	return ""
}

type RawPatchResponse struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *RawPatchResponse) Reset()                    { *m = RawPatchResponse{} }
func (m *RawPatchResponse) String() string            { return proto.CompactTextString(m) }
func (*RawPatchResponse) ProtoMessage()               {}
func (*RawPatchResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

func (m *RawPatchResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CommitDiffRequest)(nil), "gitaly.CommitDiffRequest")
	proto.RegisterType((*CommitDiffResponse)(nil), "gitaly.CommitDiffResponse")
	proto.RegisterType((*CommitDeltaRequest)(nil), "gitaly.CommitDeltaRequest")
	proto.RegisterType((*CommitDelta)(nil), "gitaly.CommitDelta")
	proto.RegisterType((*CommitDeltaResponse)(nil), "gitaly.CommitDeltaResponse")
	proto.RegisterType((*CommitPatchRequest)(nil), "gitaly.CommitPatchRequest")
	proto.RegisterType((*CommitPatchResponse)(nil), "gitaly.CommitPatchResponse")
	proto.RegisterType((*RawDiffRequest)(nil), "gitaly.RawDiffRequest")
	proto.RegisterType((*RawDiffResponse)(nil), "gitaly.RawDiffResponse")
	proto.RegisterType((*RawPatchRequest)(nil), "gitaly.RawPatchRequest")
	proto.RegisterType((*RawPatchResponse)(nil), "gitaly.RawPatchResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DiffService service

type DiffServiceClient interface {
	// Returns stream of CommitDiffResponse with patches chunked over messages
	CommitDiff(ctx context.Context, in *CommitDiffRequest, opts ...grpc.CallOption) (DiffService_CommitDiffClient, error)
	// Return a stream so we can divide the response in chunks of deltas
	CommitDelta(ctx context.Context, in *CommitDeltaRequest, opts ...grpc.CallOption) (DiffService_CommitDeltaClient, error)
	CommitPatch(ctx context.Context, in *CommitPatchRequest, opts ...grpc.CallOption) (DiffService_CommitPatchClient, error)
	RawDiff(ctx context.Context, in *RawDiffRequest, opts ...grpc.CallOption) (DiffService_RawDiffClient, error)
	RawPatch(ctx context.Context, in *RawPatchRequest, opts ...grpc.CallOption) (DiffService_RawPatchClient, error)
}

type diffServiceClient struct {
	cc *grpc.ClientConn
}

func NewDiffServiceClient(cc *grpc.ClientConn) DiffServiceClient {
	return &diffServiceClient{cc}
}

func (c *diffServiceClient) CommitDiff(ctx context.Context, in *CommitDiffRequest, opts ...grpc.CallOption) (DiffService_CommitDiffClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DiffService_serviceDesc.Streams[0], c.cc, "/gitaly.DiffService/CommitDiff", opts...)
	if err != nil {
		return nil, err
	}
	x := &diffServiceCommitDiffClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DiffService_CommitDiffClient interface {
	Recv() (*CommitDiffResponse, error)
	grpc.ClientStream
}

type diffServiceCommitDiffClient struct {
	grpc.ClientStream
}

func (x *diffServiceCommitDiffClient) Recv() (*CommitDiffResponse, error) {
	m := new(CommitDiffResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *diffServiceClient) CommitDelta(ctx context.Context, in *CommitDeltaRequest, opts ...grpc.CallOption) (DiffService_CommitDeltaClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DiffService_serviceDesc.Streams[1], c.cc, "/gitaly.DiffService/CommitDelta", opts...)
	if err != nil {
		return nil, err
	}
	x := &diffServiceCommitDeltaClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DiffService_CommitDeltaClient interface {
	Recv() (*CommitDeltaResponse, error)
	grpc.ClientStream
}

type diffServiceCommitDeltaClient struct {
	grpc.ClientStream
}

func (x *diffServiceCommitDeltaClient) Recv() (*CommitDeltaResponse, error) {
	m := new(CommitDeltaResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *diffServiceClient) CommitPatch(ctx context.Context, in *CommitPatchRequest, opts ...grpc.CallOption) (DiffService_CommitPatchClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DiffService_serviceDesc.Streams[2], c.cc, "/gitaly.DiffService/CommitPatch", opts...)
	if err != nil {
		return nil, err
	}
	x := &diffServiceCommitPatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DiffService_CommitPatchClient interface {
	Recv() (*CommitPatchResponse, error)
	grpc.ClientStream
}

type diffServiceCommitPatchClient struct {
	grpc.ClientStream
}

func (x *diffServiceCommitPatchClient) Recv() (*CommitPatchResponse, error) {
	m := new(CommitPatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *diffServiceClient) RawDiff(ctx context.Context, in *RawDiffRequest, opts ...grpc.CallOption) (DiffService_RawDiffClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DiffService_serviceDesc.Streams[3], c.cc, "/gitaly.DiffService/RawDiff", opts...)
	if err != nil {
		return nil, err
	}
	x := &diffServiceRawDiffClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DiffService_RawDiffClient interface {
	Recv() (*RawDiffResponse, error)
	grpc.ClientStream
}

type diffServiceRawDiffClient struct {
	grpc.ClientStream
}

func (x *diffServiceRawDiffClient) Recv() (*RawDiffResponse, error) {
	m := new(RawDiffResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *diffServiceClient) RawPatch(ctx context.Context, in *RawPatchRequest, opts ...grpc.CallOption) (DiffService_RawPatchClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DiffService_serviceDesc.Streams[4], c.cc, "/gitaly.DiffService/RawPatch", opts...)
	if err != nil {
		return nil, err
	}
	x := &diffServiceRawPatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DiffService_RawPatchClient interface {
	Recv() (*RawPatchResponse, error)
	grpc.ClientStream
}

type diffServiceRawPatchClient struct {
	grpc.ClientStream
}

func (x *diffServiceRawPatchClient) Recv() (*RawPatchResponse, error) {
	m := new(RawPatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for DiffService service

type DiffServiceServer interface {
	// Returns stream of CommitDiffResponse with patches chunked over messages
	CommitDiff(*CommitDiffRequest, DiffService_CommitDiffServer) error
	// Return a stream so we can divide the response in chunks of deltas
	CommitDelta(*CommitDeltaRequest, DiffService_CommitDeltaServer) error
	CommitPatch(*CommitPatchRequest, DiffService_CommitPatchServer) error
	RawDiff(*RawDiffRequest, DiffService_RawDiffServer) error
	RawPatch(*RawPatchRequest, DiffService_RawPatchServer) error
}

func RegisterDiffServiceServer(s *grpc.Server, srv DiffServiceServer) {
	s.RegisterService(&_DiffService_serviceDesc, srv)
}

func _DiffService_CommitDiff_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CommitDiffRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DiffServiceServer).CommitDiff(m, &diffServiceCommitDiffServer{stream})
}

type DiffService_CommitDiffServer interface {
	Send(*CommitDiffResponse) error
	grpc.ServerStream
}

type diffServiceCommitDiffServer struct {
	grpc.ServerStream
}

func (x *diffServiceCommitDiffServer) Send(m *CommitDiffResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DiffService_CommitDelta_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CommitDeltaRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DiffServiceServer).CommitDelta(m, &diffServiceCommitDeltaServer{stream})
}

type DiffService_CommitDeltaServer interface {
	Send(*CommitDeltaResponse) error
	grpc.ServerStream
}

type diffServiceCommitDeltaServer struct {
	grpc.ServerStream
}

func (x *diffServiceCommitDeltaServer) Send(m *CommitDeltaResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DiffService_CommitPatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CommitPatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DiffServiceServer).CommitPatch(m, &diffServiceCommitPatchServer{stream})
}

type DiffService_CommitPatchServer interface {
	Send(*CommitPatchResponse) error
	grpc.ServerStream
}

type diffServiceCommitPatchServer struct {
	grpc.ServerStream
}

func (x *diffServiceCommitPatchServer) Send(m *CommitPatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DiffService_RawDiff_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RawDiffRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DiffServiceServer).RawDiff(m, &diffServiceRawDiffServer{stream})
}

type DiffService_RawDiffServer interface {
	Send(*RawDiffResponse) error
	grpc.ServerStream
}

type diffServiceRawDiffServer struct {
	grpc.ServerStream
}

func (x *diffServiceRawDiffServer) Send(m *RawDiffResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DiffService_RawPatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RawPatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DiffServiceServer).RawPatch(m, &diffServiceRawPatchServer{stream})
}

type DiffService_RawPatchServer interface {
	Send(*RawPatchResponse) error
	grpc.ServerStream
}

type diffServiceRawPatchServer struct {
	grpc.ServerStream
}

func (x *diffServiceRawPatchServer) Send(m *RawPatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _DiffService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.DiffService",
	HandlerType: (*DiffServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CommitDiff",
			Handler:       _DiffService_CommitDiff_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CommitDelta",
			Handler:       _DiffService_CommitDelta_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CommitPatch",
			Handler:       _DiffService_CommitPatch_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RawDiff",
			Handler:       _DiffService_RawDiff_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RawPatch",
			Handler:       _DiffService_RawPatch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "diff.proto",
}

func init() { proto.RegisterFile("diff.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 753 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0x4d, 0x6f, 0xdb, 0x46,
	0x10, 0x2d, 0xf5, 0x41, 0x51, 0x23, 0x5a, 0x76, 0xd7, 0x85, 0x4d, 0xcb, 0x3d, 0x08, 0x44, 0xed,
	0xaa, 0x28, 0x60, 0x14, 0xea, 0xa5, 0xa7, 0x02, 0xb5, 0x8d, 0x16, 0x36, 0x6c, 0xd4, 0x60, 0x0e,
	0x39, 0x12, 0x6b, 0xed, 0x52, 0x5a, 0x84, 0xe4, 0x2a, 0xbb, 0x1b, 0xcb, 0xfa, 0x1b, 0xc9, 0x8f,
	0xc8, 0x25, 0xff, 0x28, 0xbf, 0x22, 0xf7, 0x1c, 0x82, 0xdd, 0x25, 0x29, 0xca, 0x56, 0x72, 0x71,
	0x0e, 0xbe, 0x69, 0xdf, 0x7b, 0x9c, 0x19, 0xbe, 0x37, 0x4b, 0x08, 0x80, 0xb0, 0x24, 0x39, 0x99,
	0x0b, 0xae, 0x38, 0x72, 0xa7, 0x4c, 0xe1, 0x74, 0x39, 0xf0, 0xe5, 0x0c, 0x0b, 0x4a, 0x2c, 0x1a,
	0x7e, 0x6e, 0xc2, 0x8f, 0x67, 0x3c, 0xcb, 0x98, 0x3a, 0x67, 0x49, 0x12, 0xd1, 0xd7, 0x6f, 0xa8,
	0x54, 0x68, 0x0c, 0x20, 0xe8, 0x9c, 0x4b, 0xa6, 0xb8, 0x58, 0x06, 0xce, 0xd0, 0x19, 0xf5, 0xc6,
	0xe8, 0xc4, 0x16, 0x38, 0x89, 0x2a, 0x26, 0xaa, 0xa9, 0xd0, 0x2f, 0xd0, 0x4f, 0x69, 0xa2, 0xe2,
	0x89, 0xa9, 0x16, 0x33, 0x12, 0x34, 0x86, 0xce, 0xa8, 0x1b, 0xf9, 0x1a, 0xb5, 0x2d, 0x2e, 0x08,
	0x3a, 0x86, 0x6d, 0xc1, 0xa6, 0xb3, 0xba, 0xac, 0x69, 0x64, 0x5b, 0x06, 0xae, 0x74, 0x7f, 0x41,
	0xc0, 0xa6, 0x39, 0x17, 0x34, 0x5e, 0xcc, 0x98, 0xa2, 0x72, 0x8e, 0x27, 0x34, 0x9e, 0xcc, 0x70,
	0x3e, 0xa5, 0x41, 0x6b, 0xe8, 0x8c, 0xbc, 0x68, 0xcf, 0xf2, 0x2f, 0x2b, 0xfa, 0xcc, 0xb0, 0xe8,
	0x27, 0x68, 0xcf, 0xb1, 0x9a, 0xc9, 0xa0, 0x3d, 0x6c, 0x8e, 0xfc, 0xc8, 0x1e, 0xd0, 0x11, 0xf4,
	0x27, 0x3c, 0x4d, 0xf1, 0x5c, 0xd2, 0x58, 0x9b, 0x22, 0x03, 0xd7, 0x54, 0xd9, 0x2a, 0x51, 0xfd,
	0xfa, 0x46, 0x46, 0xf3, 0x84, 0x8b, 0x09, 0x8d, 0x53, 0x96, 0x31, 0x25, 0x83, 0x8e, 0x95, 0x15,
	0xe8, 0x95, 0x01, 0xd1, 0x21, 0x74, 0x33, 0x7c, 0x1f, 0x27, 0x2c, 0xa5, 0x32, 0xf0, 0x86, 0xce,
	0xa8, 0x1d, 0x79, 0x19, 0xbe, 0xff, 0x57, 0x9f, 0x4b, 0x32, 0x65, 0x39, 0x95, 0x41, 0xb7, 0x22,
	0xaf, 0xf4, 0xb9, 0x24, 0x6f, 0x97, 0x8a, 0xca, 0x00, 0x2a, 0xf2, 0x54, 0x9f, 0xb5, 0x85, 0x12,
	0x27, 0x34, 0x5e, 0xd5, 0xee, 0x19, 0x85, 0xaf, 0xd1, 0xeb, 0xb2, 0x7e, 0x5d, 0x65, 0x9b, 0xf8,
	0x6b, 0x2a, 0xdb, 0xa8, 0xae, 0xb2, 0xdd, 0xb6, 0xd6, 0x54, 0xa6, 0x63, 0xf8, 0xb1, 0x01, 0xa8,
	0x1e, 0xbf, 0x9c, 0xf3, 0x5c, 0x52, 0x3d, 0x65, 0x22, 0x78, 0x16, 0x6b, 0xef, 0x4c, 0xfc, 0x7e,
	0xe4, 0x69, 0xe0, 0x06, 0xab, 0x19, 0xda, 0x87, 0x8e, 0xe2, 0x96, 0x6a, 0x18, 0xca, 0x55, 0xbc,
	0x24, 0xcc, 0x53, 0x55, 0xa6, 0xae, 0x3e, 0x5e, 0x10, 0xb4, 0x0b, 0x6d, 0xc5, 0x35, 0xdc, 0x32,
	0x70, 0x4b, 0xf1, 0x0b, 0x82, 0x0e, 0xc0, 0xe3, 0x29, 0x89, 0x33, 0x4e, 0x68, 0xd0, 0x36, 0xa3,
	0x75, 0x78, 0x4a, 0xae, 0x39, 0xa1, 0x9a, 0xca, 0xe9, 0xc2, 0x52, 0xae, 0xa5, 0x72, 0xba, 0x30,
	0xd4, 0x1e, 0xb8, 0xb7, 0x2c, 0xc7, 0x62, 0x59, 0x04, 0x53, 0x9c, 0xf4, 0xeb, 0x0a, 0xbc, 0xd0,
	0x53, 0x4d, 0x66, 0x31, 0xc1, 0x0a, 0x1b, 0xe7, 0xfd, 0xc8, 0x17, 0x78, 0x71, 0xa3, 0xc1, 0x73,
	0xac, 0x30, 0x1a, 0x82, 0x4f, 0x73, 0x12, 0xf3, 0xc4, 0x0a, 0x4d, 0x00, 0x5e, 0x04, 0x34, 0x27,
	0xff, 0x27, 0x46, 0x85, 0x7e, 0x85, 0x6d, 0x7e, 0x47, 0x45, 0x92, 0xf2, 0x45, 0x9c, 0x61, 0xf1,
	0x8a, 0x0a, 0x93, 0x81, 0x17, 0xf5, 0x4b, 0xf8, 0xda, 0xa0, 0xe8, 0x67, 0xe8, 0x96, 0xab, 0x43,
	0x4c, 0x00, 0x5e, 0xb4, 0x02, 0x2e, 0x5b, 0x9e, 0xb7, 0xd3, 0x0d, 0x3f, 0x38, 0x95, 0xbb, 0x34,
	0x55, 0xf8, 0xf9, 0xdc, 0xae, 0xea, 0x8e, 0xb4, 0x6a, 0x77, 0x24, 0x7c, 0xef, 0x40, 0xaf, 0x36,
	0xee, 0xf3, 0xdd, 0x82, 0xf0, 0x14, 0x76, 0xd7, 0x7c, 0x2d, 0xd6, 0xf6, 0x77, 0x70, 0x89, 0x06,
	0x64, 0xe0, 0x0c, 0x9b, 0xa3, 0xde, 0x78, 0xb7, 0x34, 0xb5, 0x2e, 0x2e, 0x24, 0x21, 0x29, 0xb3,
	0x31, 0xc1, 0x3f, 0x25, 0x9b, 0x01, 0x78, 0x82, 0xde, 0x31, 0xc9, 0x78, 0x5e, 0x78, 0x51, 0x9d,
	0xc3, 0xdf, 0xca, 0x49, 0x8b, 0x2e, 0xc5, 0xa4, 0x08, 0x5a, 0x66, 0x49, 0xad, 0xab, 0xe6, 0x77,
	0xf8, 0xd6, 0x81, 0x7e, 0x84, 0x17, 0xcf, 0xea, 0x3b, 0x1c, 0x1e, 0xc1, 0x76, 0x35, 0xd3, 0x37,
	0x66, 0x7f, 0xe7, 0x18, 0xdd, 0x93, 0xad, 0xfc, 0xbe, 0xc3, 0x1f, 0xc3, 0xce, 0x6a, 0xa8, 0xaf,
	0x4f, 0x3f, 0xfe, 0xd4, 0x80, 0x9e, 0x7e, 0xc5, 0x17, 0x54, 0xdc, 0xb1, 0x09, 0x45, 0xff, 0x01,
	0xac, 0x3e, 0x8a, 0xe8, 0xe0, 0xc1, 0x16, 0xad, 0xf2, 0x19, 0x0c, 0x36, 0x51, 0xb6, 0x51, 0xf8,
	0xc3, 0x1f, 0x0e, 0xba, 0x5c, 0xbf, 0x50, 0x83, 0x4d, 0xfb, 0x58, 0x94, 0x3a, 0xdc, 0xc8, 0x6d,
	0xaa, 0x65, 0x3f, 0x54, 0x0f, 0x6a, 0xd5, 0x9d, 0x7f, 0x58, 0x6b, 0xcd, 0x00, 0x53, 0xeb, 0x6f,
	0xe8, 0x14, 0xa9, 0xa2, 0xbd, 0x2a, 0x91, 0xb5, 0xd5, 0x1b, 0xec, 0x3f, 0xc2, 0x6b, 0xcf, 0xff,
	0x03, 0x5e, 0x69, 0x2c, 0xaa, 0x0b, 0xd7, 0xa6, 0x08, 0x1e, 0x13, 0xab, 0x12, 0xb7, 0xae, 0xf9,
	0xff, 0xf1, 0xe7, 0x97, 0x00, 0x00, 0x00, 0xff, 0xff, 0x55, 0x7c, 0x0d, 0x4f, 0xa3, 0x08, 0x00,
	0x00,
}
