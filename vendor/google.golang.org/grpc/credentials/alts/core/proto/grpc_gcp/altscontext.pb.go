// Code generated by protoc-gen-go. DO NOT EDIT.
// source: altscontext.proto

/*
Package grpc_gcp is a generated protocol buffer package.

It is generated from these files:
	altscontext.proto
	handshaker.proto
	transport_security_common.proto

It has these top-level messages:
	AltsContext
	Endpoint
	Identity
	StartClientHandshakeReq
	ServerHandshakeParameters
	StartServerHandshakeReq
	NextHandshakeMessageReq
	HandshakerReq
	HandshakerResult
	HandshakerStatus
	HandshakerResp
	RpcProtocolVersions
*/
package grpc_gcp

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

type AltsContext struct {
	// The application protocol negotiated for this connection.
	ApplicationProtocol string `protobuf:"bytes,1,opt,name=application_protocol,json=applicationProtocol" json:"application_protocol,omitempty"`
	// The record protocol negotiated for this connection.
	RecordProtocol string `protobuf:"bytes,2,opt,name=record_protocol,json=recordProtocol" json:"record_protocol,omitempty"`
	// The security level of the created secure channel.
	SecurityLevel SecurityLevel `protobuf:"varint,3,opt,name=security_level,json=securityLevel,enum=grpc.gcp.SecurityLevel" json:"security_level,omitempty"`
	// The peer service account.
	PeerServiceAccount string `protobuf:"bytes,4,opt,name=peer_service_account,json=peerServiceAccount" json:"peer_service_account,omitempty"`
	// The local service account.
	LocalServiceAccount string `protobuf:"bytes,5,opt,name=local_service_account,json=localServiceAccount" json:"local_service_account,omitempty"`
	// The RPC protocol versions supported by the peer.
	PeerRpcVersions *RpcProtocolVersions `protobuf:"bytes,6,opt,name=peer_rpc_versions,json=peerRpcVersions" json:"peer_rpc_versions,omitempty"`
}

func (m *AltsContext) Reset()                    { *m = AltsContext{} }
func (m *AltsContext) String() string            { return proto.CompactTextString(m) }
func (*AltsContext) ProtoMessage()               {}
func (*AltsContext) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AltsContext) GetApplicationProtocol() string {
	if m != nil {
		return m.ApplicationProtocol
	}
	return ""
}

func (m *AltsContext) GetRecordProtocol() string {
	if m != nil {
		return m.RecordProtocol
	}
	return ""
}

func (m *AltsContext) GetSecurityLevel() SecurityLevel {
	if m != nil {
		return m.SecurityLevel
	}
	return SecurityLevel_SECURITY_NONE
}

func (m *AltsContext) GetPeerServiceAccount() string {
	if m != nil {
		return m.PeerServiceAccount
	}
	return ""
}

func (m *AltsContext) GetLocalServiceAccount() string {
	if m != nil {
		return m.LocalServiceAccount
	}
	return ""
}

func (m *AltsContext) GetPeerRpcVersions() *RpcProtocolVersions {
	if m != nil {
		return m.PeerRpcVersions
	}
	return nil
}

func init() {
	proto.RegisterType((*AltsContext)(nil), "grpc.gcp.AltsContext")
}

func init() { proto.RegisterFile("altscontext.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xd9, 0xaa, 0x45, 0x53, 0x6c, 0x69, 0x6c, 0xe9, 0x22, 0x88, 0xc5, 0x8b, 0x3d, 0x2d,
	0xba, 0xde, 0x85, 0xea, 0x49, 0xf0, 0x20, 0x5b, 0xf0, 0x1a, 0xe2, 0x18, 0x4a, 0x20, 0xcd, 0x84,
	0x49, 0xba, 0xe8, 0xab, 0xfa, 0x34, 0xb2, 0xc9, 0x6e, 0x5b, 0xf4, 0x98, 0xf9, 0xbe, 0x3f, 0x33,
	0x93, 0xb0, 0xb1, 0x34, 0xc1, 0x03, 0xda, 0xa0, 0xbe, 0x42, 0xe1, 0x08, 0x03, 0xf2, 0xd3, 0x35,
	0x39, 0x28, 0xd6, 0xe0, 0x2e, 0xaf, 0x03, 0x49, 0xeb, 0x1d, 0x52, 0x10, 0x5e, 0xc1, 0x96, 0x74,
	0xf8, 0x16, 0x80, 0x9b, 0x0d, 0xda, 0xa4, 0xde, 0xfc, 0xf4, 0xd8, 0x60, 0x69, 0x82, 0x7f, 0x4e,
	0x17, 0xf0, 0x7b, 0x36, 0x91, 0xce, 0x19, 0x0d, 0x32, 0x68, 0xb4, 0x22, 0x4a, 0x80, 0x26, 0xcf,
	0xe6, 0xd9, 0xe2, 0xac, 0xba, 0x38, 0x60, 0x6f, 0x2d, 0xe2, 0xb7, 0x6c, 0x44, 0x0a, 0x90, 0x3e,
	0xf7, 0x76, 0x2f, 0xda, 0xc3, 0x54, 0xde, 0x89, 0x8f, 0x6c, 0xb8, 0x1b, 0xc2, 0xa8, 0x5a, 0x99,
	0xfc, 0x68, 0x9e, 0x2d, 0x86, 0xe5, 0xac, 0xe8, 0xe6, 0x2d, 0x56, 0x2d, 0x7f, 0x6d, 0x70, 0x75,
	0xee, 0x0f, 0x8f, 0xfc, 0x8e, 0x4d, 0x9c, 0x52, 0x24, 0xbc, 0xa2, 0x5a, 0x83, 0x12, 0x12, 0x00,
	0xb7, 0x36, 0xe4, 0xc7, 0xb1, 0x1b, 0x6f, 0xd8, 0x2a, 0xa1, 0x65, 0x22, 0xbc, 0x64, 0x53, 0x83,
	0x20, 0xcd, 0xbf, 0xc8, 0x49, 0x5a, 0x27, 0xc2, 0x3f, 0x99, 0x17, 0x36, 0x8e, 0x5d, 0xc8, 0x81,
	0xa8, 0x15, 0x79, 0x8d, 0xd6, 0xe7, 0xfd, 0x79, 0xb6, 0x18, 0x94, 0x57, 0xfb, 0x41, 0x2b, 0x07,
	0xdd, 0x5e, 0xef, 0xad, 0x54, 0x8d, 0x9a, 0x5c, 0xe5, 0xa0, 0x2b, 0x3c, 0xcd, 0xd8, 0x54, 0x63,
	0xca, 0x34, 0x9f, 0x54, 0x68, 0x1b, 0x14, 0x59, 0x69, 0x3e, 0xfa, 0xf1, 0xa5, 0x1e, 0x7e, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x36, 0xd0, 0xe1, 0x63, 0xbc, 0x01, 0x00, 0x00,
}
