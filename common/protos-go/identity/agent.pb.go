// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: identity/agent.proto

package identity

import (
	common "github.com/hyperledger-labs/weaver-dlt-interoperability/common/protos-go/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Unique identifier for a unit of a network that runs an IIN agent
type SecurityDomainMemberIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecurityDomain string `protobuf:"bytes,1,opt,name=security_domain,json=securityDomain,proto3" json:"security_domain,omitempty"`
	MemberId       string `protobuf:"bytes,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
}

func (x *SecurityDomainMemberIdentity) Reset() {
	*x = SecurityDomainMemberIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityDomainMemberIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityDomainMemberIdentity) ProtoMessage() {}

func (x *SecurityDomainMemberIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_identity_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityDomainMemberIdentity.ProtoReflect.Descriptor instead.
func (*SecurityDomainMemberIdentity) Descriptor() ([]byte, []int) {
	return file_identity_agent_proto_rawDescGZIP(), []int{0}
}

func (x *SecurityDomainMemberIdentity) GetSecurityDomain() string {
	if x != nil {
		return x.SecurityDomain
	}
	return ""
}

func (x *SecurityDomainMemberIdentity) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

type SecurityDomainMemberIdentityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceNetwork     *SecurityDomainMemberIdentity `protobuf:"bytes,1,opt,name=source_network,json=sourceNetwork,proto3" json:"source_network,omitempty"`
	RequestingNetwork *SecurityDomainMemberIdentity `protobuf:"bytes,2,opt,name=requesting_network,json=requestingNetwork,proto3" json:"requesting_network,omitempty"`
	Nonce             string                        `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *SecurityDomainMemberIdentityRequest) Reset() {
	*x = SecurityDomainMemberIdentityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityDomainMemberIdentityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityDomainMemberIdentityRequest) ProtoMessage() {}

func (x *SecurityDomainMemberIdentityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_identity_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityDomainMemberIdentityRequest.ProtoReflect.Descriptor instead.
func (*SecurityDomainMemberIdentityRequest) Descriptor() ([]byte, []int) {
	return file_identity_agent_proto_rawDescGZIP(), []int{1}
}

func (x *SecurityDomainMemberIdentityRequest) GetSourceNetwork() *SecurityDomainMemberIdentity {
	if x != nil {
		return x.SourceNetwork
	}
	return nil
}

func (x *SecurityDomainMemberIdentityRequest) GetRequestingNetwork() *SecurityDomainMemberIdentity {
	if x != nil {
		return x.RequestingNetwork
	}
	return nil
}

func (x *SecurityDomainMemberIdentityRequest) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

// Association of signature (over arbitrary data) and signer identity
type Attestation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnitIdentity *SecurityDomainMemberIdentity `protobuf:"bytes,1,opt,name=unit_identity,json=unitIdentity,proto3" json:"unit_identity,omitempty"`
	Certificate  string                        `protobuf:"bytes,2,opt,name=certificate,proto3" json:"certificate,omitempty"`
	Signature    string                        `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Attestation) Reset() {
	*x = Attestation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attestation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attestation) ProtoMessage() {}

func (x *Attestation) ProtoReflect() protoreflect.Message {
	mi := &file_identity_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attestation.ProtoReflect.Descriptor instead.
func (*Attestation) Descriptor() ([]byte, []int) {
	return file_identity_agent_proto_rawDescGZIP(), []int{2}
}

func (x *Attestation) GetUnitIdentity() *SecurityDomainMemberIdentity {
	if x != nil {
		return x.UnitIdentity
	}
	return nil
}

func (x *Attestation) GetCertificate() string {
	if x != nil {
		return x.Certificate
	}
	return ""
}

func (x *Attestation) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

// Attested security domain membership by a single member
type AttestedMembership struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Membership  *common.Membership `protobuf:"bytes,1,opt,name=membership,proto3" json:"membership,omitempty"`
	Attestation *Attestation       `protobuf:"bytes,2,opt,name=attestation,proto3" json:"attestation,omitempty"`
}

func (x *AttestedMembership) Reset() {
	*x = AttestedMembership{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttestedMembership) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestedMembership) ProtoMessage() {}

func (x *AttestedMembership) ProtoReflect() protoreflect.Message {
	mi := &file_identity_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestedMembership.ProtoReflect.Descriptor instead.
func (*AttestedMembership) Descriptor() ([]byte, []int) {
	return file_identity_agent_proto_rawDescGZIP(), []int{3}
}

func (x *AttestedMembership) GetMembership() *common.Membership {
	if x != nil {
		return x.Membership
	}
	return nil
}

func (x *AttestedMembership) GetAttestation() *Attestation {
	if x != nil {
		return x.Attestation
	}
	return nil
}

// Counter attestation over security domain membership attested by its participants
type AttestedSecurityDomain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecurityDomain *AttestedSecurityDomain_AttestedMembershipSet `protobuf:"bytes,1,opt,name=security_domain,json=securityDomain,proto3" json:"security_domain,omitempty"`
	Attestations   []*Attestation                                `protobuf:"bytes,2,rep,name=attestations,proto3" json:"attestations,omitempty"`
}

func (x *AttestedSecurityDomain) Reset() {
	*x = AttestedSecurityDomain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_agent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttestedSecurityDomain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestedSecurityDomain) ProtoMessage() {}

func (x *AttestedSecurityDomain) ProtoReflect() protoreflect.Message {
	mi := &file_identity_agent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestedSecurityDomain.ProtoReflect.Descriptor instead.
func (*AttestedSecurityDomain) Descriptor() ([]byte, []int) {
	return file_identity_agent_proto_rawDescGZIP(), []int{4}
}

func (x *AttestedSecurityDomain) GetSecurityDomain() *AttestedSecurityDomain_AttestedMembershipSet {
	if x != nil {
		return x.SecurityDomain
	}
	return nil
}

func (x *AttestedSecurityDomain) GetAttestations() []*Attestation {
	if x != nil {
		return x.Attestations
	}
	return nil
}

type AttestedSecurityDomain_AttestedMembershipSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Membership   *common.Membership `protobuf:"bytes,1,opt,name=membership,proto3" json:"membership,omitempty"`
	Attestations []*Attestation     `protobuf:"bytes,2,rep,name=attestations,proto3" json:"attestations,omitempty"`
}

func (x *AttestedSecurityDomain_AttestedMembershipSet) Reset() {
	*x = AttestedSecurityDomain_AttestedMembershipSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_agent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttestedSecurityDomain_AttestedMembershipSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestedSecurityDomain_AttestedMembershipSet) ProtoMessage() {}

func (x *AttestedSecurityDomain_AttestedMembershipSet) ProtoReflect() protoreflect.Message {
	mi := &file_identity_agent_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestedSecurityDomain_AttestedMembershipSet.ProtoReflect.Descriptor instead.
func (*AttestedSecurityDomain_AttestedMembershipSet) Descriptor() ([]byte, []int) {
	return file_identity_agent_proto_rawDescGZIP(), []int{4, 0}
}

func (x *AttestedSecurityDomain_AttestedMembershipSet) GetMembership() *common.Membership {
	if x != nil {
		return x.Membership
	}
	return nil
}

func (x *AttestedSecurityDomain_AttestedMembershipSet) GetAttestations() []*Attestation {
	if x != nil {
		return x.Attestations
	}
	return nil
}

var File_identity_agent_proto protoreflect.FileDescriptor

var file_identity_agent_proto_rawDesc = []byte{
	0x0a, 0x14, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x1a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61,
	0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x64, 0x0a, 0x1c, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x22, 0xed, 0x01, 0x0a, 0x23, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x53, 0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x12, 0x5b, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x11,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x0d, 0x75, 0x6e, 0x69, 0x74, 0x5f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0c, 0x75, 0x6e,
	0x69, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x92, 0x01, 0x0a, 0x12, 0x41,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x12, 0x3d, 0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x12, 0x3d, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0xda, 0x02, 0x0a, 0x16, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x65, 0x0a, 0x0f, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x53, 0x65,
	0x74, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x3f, 0x0a, 0x0c, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x1a, 0x97, 0x01, 0x0a, 0x15, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x53, 0x65, 0x74, 0x12, 0x3d, 0x0a, 0x0a,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52,
	0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x3f, 0x0a, 0x0c, 0x61,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xb4, 0x03, 0x0a,
	0x08, 0x49, 0x49, 0x4e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x54, 0x0a, 0x11, 0x53, 0x79, 0x6e,
	0x63, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2c,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x63, 0x6b, 0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12,
	0x5f, 0x0a, 0x1c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2c, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x0f, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x63, 0x6b, 0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00,
	0x12, 0x52, 0x0a, 0x19, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x41,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x63, 0x6b, 0x2e, 0x41,
	0x63, 0x6b, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x63, 0x6b, 0x2e,
	0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x63, 0x6b, 0x2e, 0x41, 0x63,
	0x6b, 0x22, 0x00, 0x42, 0x75, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2d,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x65, 0x72, 0x2d, 0x64, 0x6c, 0x74, 0x2d,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67,
	0x6f, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_identity_agent_proto_rawDescOnce sync.Once
	file_identity_agent_proto_rawDescData = file_identity_agent_proto_rawDesc
)

func file_identity_agent_proto_rawDescGZIP() []byte {
	file_identity_agent_proto_rawDescOnce.Do(func() {
		file_identity_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_identity_agent_proto_rawDescData)
	})
	return file_identity_agent_proto_rawDescData
}

var file_identity_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_identity_agent_proto_goTypes = []interface{}{
	(*SecurityDomainMemberIdentity)(nil),                 // 0: identity.agent.SecurityDomainMemberIdentity
	(*SecurityDomainMemberIdentityRequest)(nil),          // 1: identity.agent.SecurityDomainMemberIdentityRequest
	(*Attestation)(nil),                                  // 2: identity.agent.Attestation
	(*AttestedMembership)(nil),                           // 3: identity.agent.AttestedMembership
	(*AttestedSecurityDomain)(nil),                       // 4: identity.agent.AttestedSecurityDomain
	(*AttestedSecurityDomain_AttestedMembershipSet)(nil), // 5: identity.agent.AttestedSecurityDomain.AttestedMembershipSet
	(*common.Membership)(nil),                            // 6: common.membership.Membership
	(*common.Ack)(nil),                                   // 7: common.ack.Ack
}
var file_identity_agent_proto_depIdxs = []int32{
	0,  // 0: identity.agent.SecurityDomainMemberIdentityRequest.source_network:type_name -> identity.agent.SecurityDomainMemberIdentity
	0,  // 1: identity.agent.SecurityDomainMemberIdentityRequest.requesting_network:type_name -> identity.agent.SecurityDomainMemberIdentity
	0,  // 2: identity.agent.Attestation.unit_identity:type_name -> identity.agent.SecurityDomainMemberIdentity
	6,  // 3: identity.agent.AttestedMembership.membership:type_name -> common.membership.Membership
	2,  // 4: identity.agent.AttestedMembership.attestation:type_name -> identity.agent.Attestation
	5,  // 5: identity.agent.AttestedSecurityDomain.security_domain:type_name -> identity.agent.AttestedSecurityDomain.AttestedMembershipSet
	2,  // 6: identity.agent.AttestedSecurityDomain.attestations:type_name -> identity.agent.Attestation
	6,  // 7: identity.agent.AttestedSecurityDomain.AttestedMembershipSet.membership:type_name -> common.membership.Membership
	2,  // 8: identity.agent.AttestedSecurityDomain.AttestedMembershipSet.attestations:type_name -> identity.agent.Attestation
	0,  // 9: identity.agent.IINAgent.SyncExternalState:input_type -> identity.agent.SecurityDomainMemberIdentity
	0,  // 10: identity.agent.IINAgent.RequestIdentityConfiguration:input_type -> identity.agent.SecurityDomainMemberIdentity
	3,  // 11: identity.agent.IINAgent.SendIdentityConfiguration:input_type -> identity.agent.AttestedMembership
	4,  // 12: identity.agent.IINAgent.RequestAttestation:input_type -> identity.agent.AttestedSecurityDomain
	4,  // 13: identity.agent.IINAgent.SendAttestation:input_type -> identity.agent.AttestedSecurityDomain
	7,  // 14: identity.agent.IINAgent.SyncExternalState:output_type -> common.ack.Ack
	7,  // 15: identity.agent.IINAgent.RequestIdentityConfiguration:output_type -> common.ack.Ack
	7,  // 16: identity.agent.IINAgent.SendIdentityConfiguration:output_type -> common.ack.Ack
	7,  // 17: identity.agent.IINAgent.RequestAttestation:output_type -> common.ack.Ack
	7,  // 18: identity.agent.IINAgent.SendAttestation:output_type -> common.ack.Ack
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_identity_agent_proto_init() }
func file_identity_agent_proto_init() {
	if File_identity_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_identity_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityDomainMemberIdentity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_identity_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityDomainMemberIdentityRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_identity_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attestation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_identity_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestedMembership); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_identity_agent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestedSecurityDomain); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_identity_agent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestedSecurityDomain_AttestedMembershipSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_identity_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_identity_agent_proto_goTypes,
		DependencyIndexes: file_identity_agent_proto_depIdxs,
		MessageInfos:      file_identity_agent_proto_msgTypes,
	}.Build()
	File_identity_agent_proto = out.File
	file_identity_agent_proto_rawDesc = nil
	file_identity_agent_proto_goTypes = nil
	file_identity_agent_proto_depIdxs = nil
}
