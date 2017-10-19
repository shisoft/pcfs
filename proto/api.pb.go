// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/api.proto

/*
Package client is a generated protocol buffer package.

It is generated from these files:
	proto/api.proto

It has these top-level messages:
	BlockData
	Block
	FileMeta
	Directory
	Volume
	HostStash
	OpenRequest
	GetBlockRequest
	GetBlockResponse
	AppendToBlockRequest
	DeleteBlockRequest
	CreateBlockRequest
	Node
	BlockStashSuggestion
	WriteResult
	NewDirectoryContract
	AcquireFileWriteLockContract
	ReleaseFileWriteLockContract
	TouchFileContract
	ConfirmBlockContract
	ConfirmBlockCreationContract
	FileWriteLock
	Nothing
*/
package client

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BlockData struct {
	Group uint64 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	Index uint64 `protobuf:"varint,2,opt,name=index" json:"index,omitempty"`
	Size  uint32 `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	File  []byte `protobuf:"bytes,4,opt,name=file,proto3" json:"file,omitempty"`
	Data  []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *BlockData) Reset()                    { *m = BlockData{} }
func (m *BlockData) String() string            { return proto.CompactTextString(m) }
func (*BlockData) ProtoMessage()               {}
func (*BlockData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BlockData) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *BlockData) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *BlockData) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *BlockData) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *BlockData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Block struct {
	Hosts []uint64 `protobuf:"varint,2,rep,packed,name=hosts" json:"hosts,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Block) GetHosts() []uint64 {
	if m != nil {
		return m.Hosts
	}
	return nil
}

type FileMeta struct {
	Name         string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Size         uint64   `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	LastModified uint64   `protobuf:"varint,3,opt,name=last_modified,json=lastModified" json:"last_modified,omitempty"`
	CreatedAt    uint64   `protobuf:"varint,4,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	BlockSize    uint32   `protobuf:"varint,6,opt,name=block_size,json=blockSize" json:"block_size,omitempty"`
	Key          []byte   `protobuf:"bytes,7,opt,name=key,proto3" json:"key,omitempty"`
	Blocks       []*Block `protobuf:"bytes,8,rep,name=blocks" json:"blocks,omitempty"`
}

func (m *FileMeta) Reset()                    { *m = FileMeta{} }
func (m *FileMeta) String() string            { return proto.CompactTextString(m) }
func (*FileMeta) ProtoMessage()               {}
func (*FileMeta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FileMeta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FileMeta) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *FileMeta) GetLastModified() uint64 {
	if m != nil {
		return m.LastModified
	}
	return 0
}

func (m *FileMeta) GetCreatedAt() uint64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *FileMeta) GetBlockSize() uint32 {
	if m != nil {
		return m.BlockSize
	}
	return 0
}

func (m *FileMeta) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *FileMeta) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

type Directory struct {
	Name  string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Key   []byte   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Files [][]byte `protobuf:"bytes,3,rep,name=files,proto3" json:"files,omitempty"`
}

func (m *Directory) Reset()                    { *m = Directory{} }
func (m *Directory) String() string            { return proto.CompactTextString(m) }
func (*Directory) ProtoMessage()               {}
func (*Directory) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Directory) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Directory) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Directory) GetFiles() [][]byte {
	if m != nil {
		return m.Files
	}
	return nil
}

type Volume struct {
	Name         string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Key          []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Replications uint32 `protobuf:"varint,3,opt,name=replications" json:"replications,omitempty"`
	BlockSize    uint32 `protobuf:"varint,4,opt,name=block_size,json=blockSize" json:"block_size,omitempty"`
	RootDir      []byte `protobuf:"bytes,5,opt,name=root_dir,json=rootDir,proto3" json:"root_dir,omitempty"`
}

func (m *Volume) Reset()                    { *m = Volume{} }
func (m *Volume) String() string            { return proto.CompactTextString(m) }
func (*Volume) ProtoMessage()               {}
func (*Volume) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Volume) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Volume) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Volume) GetReplications() uint32 {
	if m != nil {
		return m.Replications
	}
	return 0
}

func (m *Volume) GetBlockSize() uint32 {
	if m != nil {
		return m.BlockSize
	}
	return 0
}

func (m *Volume) GetRootDir() []byte {
	if m != nil {
		return m.RootDir
	}
	return nil
}

type HostStash struct {
	HostId   uint64 `protobuf:"varint,1,opt,name=host_id,json=hostId" json:"host_id,omitempty"`
	Capacity uint64 `protobuf:"varint,2,opt,name=capacity" json:"capacity,omitempty"`
	Used     uint64 `protobuf:"varint,3,opt,name=used" json:"used,omitempty"`
	Owner    uint64 `protobuf:"varint,4,opt,name=owner" json:"owner,omitempty"`
}

func (m *HostStash) Reset()                    { *m = HostStash{} }
func (m *HostStash) String() string            { return proto.CompactTextString(m) }
func (*HostStash) ProtoMessage()               {}
func (*HostStash) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *HostStash) GetHostId() uint64 {
	if m != nil {
		return m.HostId
	}
	return 0
}

func (m *HostStash) GetCapacity() uint64 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *HostStash) GetUsed() uint64 {
	if m != nil {
		return m.Used
	}
	return 0
}

func (m *HostStash) GetOwner() uint64 {
	if m != nil {
		return m.Owner
	}
	return 0
}

type OpenRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Key  []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *OpenRequest) Reset()                    { *m = OpenRequest{} }
func (m *OpenRequest) String() string            { return proto.CompactTextString(m) }
func (*OpenRequest) ProtoMessage()               {}
func (*OpenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *OpenRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OpenRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type GetBlockRequest struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *GetBlockRequest) Reset()                    { *m = GetBlockRequest{} }
func (m *GetBlockRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBlockRequest) ProtoMessage()               {}
func (*GetBlockRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetBlockRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type GetBlockResponse struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *GetBlockResponse) Reset()                    { *m = GetBlockResponse{} }
func (m *GetBlockResponse) String() string            { return proto.CompactTextString(m) }
func (*GetBlockResponse) ProtoMessage()               {}
func (*GetBlockResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetBlockResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type AppendToBlockRequest struct {
	Group  uint64 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	Index  uint64 `protobuf:"varint,2,opt,name=index" json:"index,omitempty"`
	Offset uint32 `protobuf:"varint,3,opt,name=offset" json:"offset,omitempty"`
	File   []byte `protobuf:"bytes,4,opt,name=file,proto3" json:"file,omitempty"`
	Data   []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *AppendToBlockRequest) Reset()                    { *m = AppendToBlockRequest{} }
func (m *AppendToBlockRequest) String() string            { return proto.CompactTextString(m) }
func (*AppendToBlockRequest) ProtoMessage()               {}
func (*AppendToBlockRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AppendToBlockRequest) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *AppendToBlockRequest) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *AppendToBlockRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *AppendToBlockRequest) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *AppendToBlockRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeleteBlockRequest struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *DeleteBlockRequest) Reset()                    { *m = DeleteBlockRequest{} }
func (m *DeleteBlockRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteBlockRequest) ProtoMessage()               {}
func (*DeleteBlockRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DeleteBlockRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type CreateBlockRequest struct {
	Group     uint64 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	Index     uint64 `protobuf:"varint,2,opt,name=index" json:"index,omitempty"`
	File      []byte `protobuf:"bytes,4,opt,name=file,proto3" json:"file,omitempty"`
	Signature []byte `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *CreateBlockRequest) Reset()                    { *m = CreateBlockRequest{} }
func (m *CreateBlockRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateBlockRequest) ProtoMessage()               {}
func (*CreateBlockRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *CreateBlockRequest) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *CreateBlockRequest) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *CreateBlockRequest) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *CreateBlockRequest) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Node struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	NodeId  uint64 `protobuf:"varint,2,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Remains uint64 `protobuf:"varint,3,opt,name=remains" json:"remains,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Node) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Node) GetNodeId() uint64 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *Node) GetRemains() uint64 {
	if m != nil {
		return m.Remains
	}
	return 0
}

type BlockStashSuggestion struct {
	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *BlockStashSuggestion) Reset()                    { *m = BlockStashSuggestion{} }
func (m *BlockStashSuggestion) String() string            { return proto.CompactTextString(m) }
func (*BlockStashSuggestion) ProtoMessage()               {}
func (*BlockStashSuggestion) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *BlockStashSuggestion) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type WriteResult struct {
	Succeed bool   `protobuf:"varint,1,opt,name=succeed" json:"succeed,omitempty"`
	Remains uint64 `protobuf:"varint,2,opt,name=remains" json:"remains,omitempty"`
	Key     []byte `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *WriteResult) Reset()                    { *m = WriteResult{} }
func (m *WriteResult) String() string            { return proto.CompactTextString(m) }
func (*WriteResult) ProtoMessage()               {}
func (*WriteResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *WriteResult) GetSucceed() bool {
	if m != nil {
		return m.Succeed
	}
	return false
}

func (m *WriteResult) GetRemains() uint64 {
	if m != nil {
		return m.Remains
	}
	return 0
}

func (m *WriteResult) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type NewDirectoryContract struct {
	ParentDir []byte     `protobuf:"bytes,1,opt,name=parent_dir,json=parentDir,proto3" json:"parent_dir,omitempty"`
	Dir       *Directory `protobuf:"bytes,3,opt,name=dir" json:"dir,omitempty"`
}

func (m *NewDirectoryContract) Reset()                    { *m = NewDirectoryContract{} }
func (m *NewDirectoryContract) String() string            { return proto.CompactTextString(m) }
func (*NewDirectoryContract) ProtoMessage()               {}
func (*NewDirectoryContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *NewDirectoryContract) GetParentDir() []byte {
	if m != nil {
		return m.ParentDir
	}
	return nil
}

func (m *NewDirectoryContract) GetDir() *Directory {
	if m != nil {
		return m.Dir
	}
	return nil
}

type AcquireFileWriteLockContract struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *AcquireFileWriteLockContract) Reset()                    { *m = AcquireFileWriteLockContract{} }
func (m *AcquireFileWriteLockContract) String() string            { return proto.CompactTextString(m) }
func (*AcquireFileWriteLockContract) ProtoMessage()               {}
func (*AcquireFileWriteLockContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *AcquireFileWriteLockContract) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type ReleaseFileWriteLockContract struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *ReleaseFileWriteLockContract) Reset()                    { *m = ReleaseFileWriteLockContract{} }
func (m *ReleaseFileWriteLockContract) String() string            { return proto.CompactTextString(m) }
func (*ReleaseFileWriteLockContract) ProtoMessage()               {}
func (*ReleaseFileWriteLockContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *ReleaseFileWriteLockContract) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type TouchFileContract struct {
	ClientTime uint64 `protobuf:"varint,1,opt,name=client_time,json=clientTime" json:"client_time,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Dir        []byte `protobuf:"bytes,3,opt,name=dir,proto3" json:"dir,omitempty"`
	Volume     []byte `protobuf:"bytes,4,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (m *TouchFileContract) Reset()                    { *m = TouchFileContract{} }
func (m *TouchFileContract) String() string            { return proto.CompactTextString(m) }
func (*TouchFileContract) ProtoMessage()               {}
func (*TouchFileContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *TouchFileContract) GetClientTime() uint64 {
	if m != nil {
		return m.ClientTime
	}
	return 0
}

func (m *TouchFileContract) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TouchFileContract) GetDir() []byte {
	if m != nil {
		return m.Dir
	}
	return nil
}

func (m *TouchFileContract) GetVolume() []byte {
	if m != nil {
		return m.Volume
	}
	return nil
}

type ConfirmBlockContract struct {
	NodeId uint64              `protobuf:"varint,1,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Index  uint64              `protobuf:"varint,2,opt,name=index" json:"index,omitempty"`
	File   []byte              `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
	Req    *CreateBlockRequest `protobuf:"bytes,4,opt,name=req" json:"req,omitempty"`
}

func (m *ConfirmBlockContract) Reset()                    { *m = ConfirmBlockContract{} }
func (m *ConfirmBlockContract) String() string            { return proto.CompactTextString(m) }
func (*ConfirmBlockContract) ProtoMessage()               {}
func (*ConfirmBlockContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *ConfirmBlockContract) GetNodeId() uint64 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *ConfirmBlockContract) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ConfirmBlockContract) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *ConfirmBlockContract) GetReq() *CreateBlockRequest {
	if m != nil {
		return m.Req
	}
	return nil
}

type ConfirmBlockCreationContract struct {
	Index      uint64   `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	ClientTime uint64   `protobuf:"varint,2,opt,name=client_time,json=clientTime" json:"client_time,omitempty"`
	NodeIds    []uint64 `protobuf:"varint,3,rep,packed,name=node_ids,json=nodeIds" json:"node_ids,omitempty"`
	File       []byte   `protobuf:"bytes,4,opt,name=file,proto3" json:"file,omitempty"`
}

func (m *ConfirmBlockCreationContract) Reset()                    { *m = ConfirmBlockCreationContract{} }
func (m *ConfirmBlockCreationContract) String() string            { return proto.CompactTextString(m) }
func (*ConfirmBlockCreationContract) ProtoMessage()               {}
func (*ConfirmBlockCreationContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *ConfirmBlockCreationContract) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ConfirmBlockCreationContract) GetClientTime() uint64 {
	if m != nil {
		return m.ClientTime
	}
	return 0
}

func (m *ConfirmBlockCreationContract) GetNodeIds() []uint64 {
	if m != nil {
		return m.NodeIds
	}
	return nil
}

func (m *ConfirmBlockCreationContract) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

type FileWriteLock struct {
	Group uint64 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	Owner uint64 `protobuf:"varint,2,opt,name=owner" json:"owner,omitempty"`
	Key   []byte `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *FileWriteLock) Reset()                    { *m = FileWriteLock{} }
func (m *FileWriteLock) String() string            { return proto.CompactTextString(m) }
func (*FileWriteLock) ProtoMessage()               {}
func (*FileWriteLock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *FileWriteLock) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *FileWriteLock) GetOwner() uint64 {
	if m != nil {
		return m.Owner
	}
	return 0
}

func (m *FileWriteLock) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type Nothing struct {
}

func (m *Nothing) Reset()                    { *m = Nothing{} }
func (m *Nothing) String() string            { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()               {}
func (*Nothing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func init() {
	proto.RegisterType((*BlockData)(nil), "client.BlockData")
	proto.RegisterType((*Block)(nil), "client.Block")
	proto.RegisterType((*FileMeta)(nil), "client.FileMeta")
	proto.RegisterType((*Directory)(nil), "client.Directory")
	proto.RegisterType((*Volume)(nil), "client.Volume")
	proto.RegisterType((*HostStash)(nil), "client.HostStash")
	proto.RegisterType((*OpenRequest)(nil), "client.OpenRequest")
	proto.RegisterType((*GetBlockRequest)(nil), "client.GetBlockRequest")
	proto.RegisterType((*GetBlockResponse)(nil), "client.GetBlockResponse")
	proto.RegisterType((*AppendToBlockRequest)(nil), "client.AppendToBlockRequest")
	proto.RegisterType((*DeleteBlockRequest)(nil), "client.DeleteBlockRequest")
	proto.RegisterType((*CreateBlockRequest)(nil), "client.CreateBlockRequest")
	proto.RegisterType((*Node)(nil), "client.Node")
	proto.RegisterType((*BlockStashSuggestion)(nil), "client.BlockStashSuggestion")
	proto.RegisterType((*WriteResult)(nil), "client.WriteResult")
	proto.RegisterType((*NewDirectoryContract)(nil), "client.NewDirectoryContract")
	proto.RegisterType((*AcquireFileWriteLockContract)(nil), "client.AcquireFileWriteLockContract")
	proto.RegisterType((*ReleaseFileWriteLockContract)(nil), "client.ReleaseFileWriteLockContract")
	proto.RegisterType((*TouchFileContract)(nil), "client.TouchFileContract")
	proto.RegisterType((*ConfirmBlockContract)(nil), "client.ConfirmBlockContract")
	proto.RegisterType((*ConfirmBlockCreationContract)(nil), "client.ConfirmBlockCreationContract")
	proto.RegisterType((*FileWriteLock)(nil), "client.FileWriteLock")
	proto.RegisterType((*Nothing)(nil), "client.Nothing")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PCFS service

type PCFSClient interface {
	GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*GetBlockResponse, error)
	AppendToBlock(ctx context.Context, in *AppendToBlockRequest, opts ...grpc.CallOption) (*WriteResult, error)
	CreateBlock(ctx context.Context, in *CreateBlockRequest, opts ...grpc.CallOption) (*WriteResult, error)
	DeleteBlock(ctx context.Context, in *DeleteBlockRequest, opts ...grpc.CallOption) (*WriteResult, error)
	SuggestBlockStash(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (*BlockStashSuggestion, error)
}

type pCFSClient struct {
	cc *grpc.ClientConn
}

func NewPCFSClient(cc *grpc.ClientConn) PCFSClient {
	return &pCFSClient{cc}
}

func (c *pCFSClient) GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*GetBlockResponse, error) {
	out := new(GetBlockResponse)
	err := grpc.Invoke(ctx, "/client.PCFS/GetBlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pCFSClient) AppendToBlock(ctx context.Context, in *AppendToBlockRequest, opts ...grpc.CallOption) (*WriteResult, error) {
	out := new(WriteResult)
	err := grpc.Invoke(ctx, "/client.PCFS/AppendToBlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pCFSClient) CreateBlock(ctx context.Context, in *CreateBlockRequest, opts ...grpc.CallOption) (*WriteResult, error) {
	out := new(WriteResult)
	err := grpc.Invoke(ctx, "/client.PCFS/CreateBlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pCFSClient) DeleteBlock(ctx context.Context, in *DeleteBlockRequest, opts ...grpc.CallOption) (*WriteResult, error) {
	out := new(WriteResult)
	err := grpc.Invoke(ctx, "/client.PCFS/DeleteBlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pCFSClient) SuggestBlockStash(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (*BlockStashSuggestion, error) {
	out := new(BlockStashSuggestion)
	err := grpc.Invoke(ctx, "/client.PCFS/SuggestBlockStash", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PCFS service

type PCFSServer interface {
	GetBlock(context.Context, *GetBlockRequest) (*GetBlockResponse, error)
	AppendToBlock(context.Context, *AppendToBlockRequest) (*WriteResult, error)
	CreateBlock(context.Context, *CreateBlockRequest) (*WriteResult, error)
	DeleteBlock(context.Context, *DeleteBlockRequest) (*WriteResult, error)
	SuggestBlockStash(context.Context, *Nothing) (*BlockStashSuggestion, error)
}

func RegisterPCFSServer(s *grpc.Server, srv PCFSServer) {
	s.RegisterService(&_PCFS_serviceDesc, srv)
}

func _PCFS_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PCFSServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.PCFS/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PCFSServer).GetBlock(ctx, req.(*GetBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PCFS_AppendToBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendToBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PCFSServer).AppendToBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.PCFS/AppendToBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PCFSServer).AppendToBlock(ctx, req.(*AppendToBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PCFS_CreateBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PCFSServer).CreateBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.PCFS/CreateBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PCFSServer).CreateBlock(ctx, req.(*CreateBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PCFS_DeleteBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PCFSServer).DeleteBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.PCFS/DeleteBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PCFSServer).DeleteBlock(ctx, req.(*DeleteBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PCFS_SuggestBlockStash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nothing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PCFSServer).SuggestBlockStash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.PCFS/SuggestBlockStash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PCFSServer).SuggestBlockStash(ctx, req.(*Nothing))
	}
	return interceptor(ctx, in, info, handler)
}

var _PCFS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "client.PCFS",
	HandlerType: (*PCFSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlock",
			Handler:    _PCFS_GetBlock_Handler,
		},
		{
			MethodName: "AppendToBlock",
			Handler:    _PCFS_AppendToBlock_Handler,
		},
		{
			MethodName: "CreateBlock",
			Handler:    _PCFS_CreateBlock_Handler,
		},
		{
			MethodName: "DeleteBlock",
			Handler:    _PCFS_DeleteBlock_Handler,
		},
		{
			MethodName: "SuggestBlockStash",
			Handler:    _PCFS_SuggestBlockStash_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api.proto",
}

func init() { proto.RegisterFile("proto/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 946 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x4b, 0x6f, 0x23, 0x45,
	0x17, 0x4d, 0xbb, 0x3b, 0x7e, 0x5c, 0xdb, 0xca, 0xa4, 0x3e, 0xeb, 0x9b, 0x1e, 0xcb, 0x23, 0xac,
	0x8a, 0x18, 0x79, 0x81, 0xc2, 0x28, 0xb3, 0x63, 0x03, 0x21, 0xd6, 0x0c, 0x23, 0x91, 0x80, 0xda,
	0x11, 0x48, 0x6c, 0xac, 0x9a, 0xee, 0xb2, 0x53, 0xa4, 0xd3, 0xd5, 0xa9, 0xaa, 0x66, 0x08, 0x4b,
	0x84, 0xc4, 0x86, 0xdf, 0xc6, 0x96, 0xbf, 0x83, 0xea, 0xd1, 0x0f, 0xdb, 0x99, 0x10, 0xd8, 0xd5,
	0xbd, 0x55, 0x7d, 0x5f, 0xe7, 0xdc, 0x63, 0xc3, 0x41, 0x2e, 0xb8, 0xe2, 0x9f, 0x92, 0x9c, 0x1d,
	0x9b, 0x13, 0x6a, 0xc7, 0x29, 0xa3, 0x99, 0xc2, 0x12, 0x7a, 0x5f, 0xa6, 0x3c, 0xbe, 0x9e, 0x13,
	0x45, 0xd0, 0x08, 0xf6, 0xd7, 0x82, 0x17, 0x79, 0xe8, 0x4d, 0xbd, 0x59, 0x10, 0x59, 0x43, 0x7b,
	0x59, 0x96, 0xd0, 0x9f, 0xc3, 0x96, 0xf5, 0x1a, 0x03, 0x21, 0x08, 0x24, 0xfb, 0x85, 0x86, 0xfe,
	0xd4, 0x9b, 0x0d, 0x23, 0x73, 0xd6, 0xbe, 0x15, 0x4b, 0x69, 0x18, 0x4c, 0xbd, 0xd9, 0x20, 0x32,
	0x67, 0xed, 0x4b, 0x88, 0x22, 0xe1, 0xbe, 0xf5, 0xe9, 0x33, 0x7e, 0x0e, 0xfb, 0x26, 0xa9, 0x0e,
	0x7d, 0xc5, 0xa5, 0x92, 0x61, 0x6b, 0xea, 0xeb, 0xd0, 0xc6, 0xc0, 0x7f, 0x7a, 0xd0, 0x7d, 0xcd,
	0x52, 0x7a, 0x4e, 0x15, 0xd1, 0xdf, 0x67, 0xe4, 0x86, 0x9a, 0x92, 0x7a, 0x91, 0x39, 0x57, 0xb9,
	0x6d, 0x41, 0x36, 0xf7, 0x11, 0x0c, 0x53, 0x22, 0xd5, 0xf2, 0x86, 0x27, 0x6c, 0xc5, 0x68, 0x62,
	0x0a, 0x0b, 0xa2, 0x81, 0x76, 0x9e, 0x3b, 0x1f, 0x7a, 0x0e, 0x10, 0x0b, 0x4a, 0x14, 0x4d, 0x96,
	0x44, 0x99, 0x32, 0x83, 0xa8, 0xe7, 0x3c, 0xa7, 0x4a, 0x5f, 0xbf, 0xd3, 0x75, 0x2d, 0x4d, 0xf4,
	0xb6, 0xe9, 0xac, 0x67, 0x3c, 0x0b, 0x9d, 0xe2, 0x09, 0xf8, 0xd7, 0xf4, 0x2e, 0xec, 0x98, 0x4e,
	0xf4, 0x11, 0x7d, 0x0c, 0x6d, 0x73, 0x2d, 0xc3, 0xee, 0xd4, 0x9f, 0xf5, 0x4f, 0x86, 0xc7, 0x76,
	0xac, 0xc7, 0xa6, 0xbd, 0xc8, 0x5d, 0xe2, 0x37, 0xd0, 0x9b, 0x33, 0x41, 0x63, 0xc5, 0xc5, 0xdd,
	0xbd, 0x0d, 0xb9, 0xc8, 0xad, 0x3a, 0xf2, 0x08, 0xf6, 0xf5, 0xf8, 0x64, 0xe8, 0x4f, 0xfd, 0xd9,
	0x20, 0xb2, 0x06, 0xfe, 0xc3, 0x83, 0xf6, 0x77, 0x3c, 0x2d, 0xec, 0x0c, 0x1e, 0x11, 0x06, 0xc3,
	0x40, 0xd0, 0x3c, 0x65, 0x31, 0x51, 0x8c, 0x67, 0xd2, 0xa1, 0xb5, 0xe1, 0xdb, 0xea, 0x3a, 0xd8,
	0xee, 0xfa, 0x19, 0x74, 0x05, 0xe7, 0x6a, 0x99, 0x30, 0xe1, 0x40, 0xec, 0x68, 0x7b, 0xce, 0x04,
	0xfe, 0x11, 0x7a, 0x5f, 0x71, 0xa9, 0x16, 0x8a, 0xc8, 0x2b, 0xf4, 0x14, 0x3a, 0x1a, 0xbe, 0x25,
	0x4b, 0x1c, 0x7d, 0xda, 0xda, 0x7c, 0x9b, 0xa0, 0x31, 0x74, 0x63, 0x92, 0x93, 0x98, 0xa9, 0x3b,
	0x87, 0x58, 0x65, 0xeb, 0x2e, 0x0a, 0x59, 0x81, 0x65, 0xce, 0xba, 0x75, 0xfe, 0x3e, 0xa3, 0xc2,
	0xe1, 0x63, 0x0d, 0xfc, 0x0a, 0xfa, 0xdf, 0xe4, 0x34, 0x8b, 0xe8, 0x6d, 0x41, 0xa5, 0x7a, 0x5c,
	0xfb, 0xf8, 0x08, 0x0e, 0xde, 0x50, 0x65, 0xc1, 0x70, 0x1f, 0xba, 0x47, 0x5e, 0xfd, 0xe8, 0x05,
	0x3c, 0xa9, 0x1f, 0xc9, 0x9c, 0x67, 0xb2, 0x66, 0xad, 0xd7, 0x60, 0xed, 0xaf, 0x1e, 0x8c, 0x4e,
	0xf3, 0x9c, 0x66, 0xc9, 0x25, 0xdf, 0x08, 0xf9, 0x6f, 0xd6, 0xe6, 0xff, 0xd0, 0xe6, 0xab, 0x95,
	0xa4, 0xca, 0x41, 0xe1, 0xac, 0x47, 0xaf, 0xce, 0x0b, 0x40, 0x73, 0x9a, 0x52, 0x45, 0xff, 0xa1,
	0x29, 0x01, 0xe8, 0xcc, 0xf0, 0xfa, 0x3f, 0x57, 0x7a, 0x5f, 0x45, 0x13, 0xe8, 0x49, 0xb6, 0xce,
	0x88, 0x2a, 0x04, 0x75, 0x65, 0xd5, 0x0e, 0xbc, 0x80, 0xe0, 0x82, 0x27, 0x14, 0x85, 0xd0, 0x21,
	0x49, 0x22, 0xa8, 0x94, 0x0e, 0x9e, 0xd2, 0xd4, 0x1c, 0xc9, 0x78, 0x42, 0x35, 0x47, 0x6c, 0xae,
	0xb6, 0x36, 0xdf, 0x26, 0xfa, 0x13, 0x41, 0x6f, 0x08, 0x73, 0x14, 0x0d, 0xa2, 0xd2, 0xc4, 0x9f,
	0xc1, 0xc8, 0xb4, 0x60, 0x48, 0xb6, 0x28, 0xd6, 0x6b, 0x2a, 0x35, 0x6d, 0x11, 0x86, 0x7d, 0xfd,
	0xad, 0x4e, 0xa1, 0x37, 0x6f, 0x50, 0x6e, 0x9e, 0xae, 0x20, 0xb2, 0x57, 0x78, 0x01, 0xfd, 0xef,
	0x05, 0x53, 0x34, 0xa2, 0xb2, 0x48, 0x95, 0x4e, 0x22, 0x8b, 0x38, 0xa6, 0xd4, 0x32, 0xb4, 0x1b,
	0x95, 0x66, 0x33, 0x7d, 0x6b, 0x23, 0x7d, 0x39, 0x59, 0xbf, 0x9e, 0xec, 0x0f, 0x30, 0xba, 0xa0,
	0xef, 0xab, 0x7d, 0x3e, 0xe3, 0x99, 0x12, 0x24, 0x36, 0xe2, 0x91, 0x13, 0x41, 0x33, 0xbb, 0x29,
	0x16, 0x8a, 0x9e, 0xf5, 0xcc, 0x99, 0x40, 0x47, 0xe0, 0x6b, 0xbf, 0x0e, 0xd4, 0x3f, 0x39, 0x2c,
	0xab, 0xad, 0xc2, 0x44, 0xfa, 0x16, 0xbf, 0x84, 0xc9, 0x69, 0x7c, 0x5b, 0x30, 0x41, 0xb5, 0xfe,
	0x99, 0xda, 0xbf, 0xe6, 0xf1, 0x75, 0x95, 0x63, 0x17, 0xe7, 0x97, 0x30, 0x89, 0x68, 0x4a, 0x89,
	0x7c, 0xf4, 0x17, 0x02, 0x0e, 0x2f, 0x79, 0x11, 0x5f, 0xe9, 0xf7, 0xd5, 0xb3, 0x8f, 0xa0, 0x6f,
	0x2b, 0x5a, 0x2a, 0xe6, 0xb6, 0x2a, 0x88, 0xc0, 0xba, 0x2e, 0x59, 0x43, 0x6e, 0x5a, 0x9b, 0xfb,
	0x56, 0xb6, 0x34, 0x30, 0xf5, 0x6b, 0x76, 0xff, 0x64, 0xe4, 0xc9, 0xb1, 0xc6, 0x59, 0xf8, 0x77,
	0x0f, 0x46, 0x67, 0x3c, 0x5b, 0x31, 0x71, 0x63, 0xc0, 0xac, 0xf2, 0x36, 0x08, 0xe1, 0x6d, 0x10,
	0xe2, 0x61, 0x4e, 0xfa, 0x0d, 0x4e, 0x7e, 0x02, 0xbe, 0xa0, 0xb7, 0x26, 0x61, 0xff, 0x64, 0x5c,
	0x0e, 0x76, 0x97, 0xfc, 0x91, 0x7e, 0x86, 0x7f, 0xf3, 0x60, 0xb2, 0x51, 0x89, 0x7e, 0xc7, 0x78,
	0x56, 0x55, 0x54, 0x25, 0xf6, 0x9a, 0x89, 0xb7, 0xe6, 0xd3, 0xda, 0x99, 0xcf, 0x33, 0xe8, 0xba,
	0x46, 0xac, 0x64, 0x07, 0x51, 0xc7, 0x76, 0x22, 0xef, 0x5b, 0x24, 0x7c, 0x0e, 0xc3, 0x0d, 0xbc,
	0x3e, 0xbc, 0x99, 0x56, 0x0a, 0x5b, 0x0d, 0x29, 0x2c, 0x31, 0x0d, 0x6a, 0x4c, 0x7b, 0xd0, 0xb9,
	0xe0, 0xea, 0x8a, 0x65, 0xeb, 0x93, 0xbf, 0x5a, 0x10, 0x7c, 0x7b, 0xf6, 0x7a, 0x81, 0x3e, 0x87,
	0x6e, 0x29, 0x6b, 0xe8, 0x69, 0x39, 0x96, 0x2d, 0x35, 0x1c, 0x87, 0xbb, 0x17, 0x56, 0x01, 0xf1,
	0x1e, 0x9a, 0xc3, 0x70, 0x43, 0xee, 0xd0, 0xa4, 0x7c, 0x7c, 0x9f, 0x0a, 0x8e, 0xff, 0x57, 0xde,
	0x36, 0x56, 0x0e, 0xef, 0xa1, 0x2f, 0xa0, 0xdf, 0xc0, 0x02, 0x3d, 0x00, 0xd0, 0x03, 0x11, 0x1a,
	0x92, 0x57, 0x47, 0xd8, 0xd5, 0xc1, 0x0f, 0x45, 0x98, 0xc3, 0xa1, 0x53, 0x8e, 0x5a, 0x4a, 0xd0,
	0x41, 0xad, 0x18, 0x66, 0x72, 0xe3, 0xc9, 0xc6, 0x8f, 0xf7, 0x96, 0xde, 0xe0, 0xbd, 0x77, 0x6d,
	0xf3, 0xcf, 0xe9, 0xd5, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x51, 0x99, 0x02, 0x3b, 0x4c, 0x09,
	0x00, 0x00,
}
