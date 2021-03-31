// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: eth/v1/beacon_state.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/protoc-gen-go/descriptor"
	github_com_prysmaticlabs_eth2_types "github.com/prysmaticlabs/eth2-types"
	_ "github.com/prysmaticlabs/ethereumapis/eth/ext"
	github_com_prysmaticlabs_go_bitfield "github.com/prysmaticlabs/go-bitfield"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type BeaconState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GenesisTime                 uint64                                          `protobuf:"varint,1,opt,name=genesis_time,json=genesisTime,proto3" json:"genesis_time,omitempty"`
	GenesisValidatorsRoot       []byte                                          `protobuf:"bytes,2,opt,name=genesis_validators_root,json=genesisValidatorsRoot,proto3" json:"genesis_validators_root,omitempty" ssz-size:"32"`
	Slot                        github_com_prysmaticlabs_eth2_types.Slot        `protobuf:"varint,3,opt,name=slot,proto3" json:"slot,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.Slot"`
	Fork                        *Fork                                           `protobuf:"bytes,4,opt,name=fork,proto3" json:"fork,omitempty"`
	LatestBlockHeader           *BeaconBlockHeader                              `protobuf:"bytes,5,opt,name=latest_block_header,json=latestBlockHeader,proto3" json:"latest_block_header,omitempty"`
	BlockRoots                  [][]byte                                        `protobuf:"bytes,6,rep,name=block_roots,json=blockRoots,proto3" json:"block_roots,omitempty" ssz-size:"8192"`
	StateRoots                  [][]byte                                        `protobuf:"bytes,7,rep,name=state_roots,json=stateRoots,proto3" json:"state_roots,omitempty" ssz-size:"8192"`
	HistoricalRoots             [][]byte                                        `protobuf:"bytes,8,rep,name=historical_roots,json=historicalRoots,proto3" json:"historical_roots,omitempty" ssz-max:"16777216" ssz-size:"32"`
	Eth1Data                    *Eth1Data                                       `protobuf:"bytes,9,opt,name=eth1_data,json=eth1Data,proto3" json:"eth1_data,omitempty"`
	Eth1DataVotes               []*Eth1Data                                     `protobuf:"bytes,10,rep,name=eth1_data_votes,json=eth1DataVotes,proto3" json:"eth1_data_votes,omitempty" ssz-max:"1024"`
	Eth1DepositIndex            uint64                                          `protobuf:"varint,11,opt,name=eth1_deposit_index,json=eth1DepositIndex,proto3" json:"eth1_deposit_index,omitempty"`
	Validators                  []*Validator                                    `protobuf:"bytes,12,rep,name=validators,proto3" json:"validators,omitempty" ssz-max:"1099511627776"`
	Balances                    []uint64                                        `protobuf:"varint,13,rep,packed,name=balances,proto3" json:"balances,omitempty" ssz-max:"1099511627776"`
	RandaoMixes                 [][]byte                                        `protobuf:"bytes,14,rep,name=randao_mixes,json=randaoMixes,proto3" json:"randao_mixes,omitempty" ssz-size:"65536"`
	Slashings                   []uint64                                        `protobuf:"varint,15,rep,packed,name=slashings,proto3" json:"slashings,omitempty" ssz-size:"65536"`
	PreviousEpochAttestations   []*PendingAttestation                           `protobuf:"bytes,16,rep,name=previous_epoch_attestations,json=previousEpochAttestations,proto3" json:"previous_epoch_attestations,omitempty" ssz-max:"4096"`
	CurrentEpochAttestations    []*PendingAttestation                           `protobuf:"bytes,17,rep,name=current_epoch_attestations,json=currentEpochAttestations,proto3" json:"current_epoch_attestations,omitempty" ssz-max:"4096"`
	JustificationBits           github_com_prysmaticlabs_go_bitfield.Bitvector4 `protobuf:"bytes,18,opt,name=justification_bits,json=justificationBits,proto3" json:"justification_bits,omitempty" cast-type:"github.com/prysmaticlabs/go-bitfield.Bitvector4" ssz-size:"1"`
	PreviousJustifiedCheckpoint *Checkpoint                                     `protobuf:"bytes,19,opt,name=previous_justified_checkpoint,json=previousJustifiedCheckpoint,proto3" json:"previous_justified_checkpoint,omitempty"`
	CurrentJustifiedCheckpoint  *Checkpoint                                     `protobuf:"bytes,20,opt,name=current_justified_checkpoint,json=currentJustifiedCheckpoint,proto3" json:"current_justified_checkpoint,omitempty"`
	FinalizedCheckpoint         *Checkpoint                                     `protobuf:"bytes,21,opt,name=finalized_checkpoint,json=finalizedCheckpoint,proto3" json:"finalized_checkpoint,omitempty"`
}

func (x *BeaconState) Reset() {
	*x = BeaconState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_beacon_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeaconState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeaconState) ProtoMessage() {}

func (x *BeaconState) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_beacon_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeaconState.ProtoReflect.Descriptor instead.
func (*BeaconState) Descriptor() ([]byte, []int) {
	return file_eth_v1_beacon_state_proto_rawDescGZIP(), []int{0}
}

func (x *BeaconState) GetGenesisTime() uint64 {
	if x != nil {
		return x.GenesisTime
	}
	return 0
}

func (x *BeaconState) GetGenesisValidatorsRoot() []byte {
	if x != nil {
		return x.GenesisValidatorsRoot
	}
	return nil
}

func (x *BeaconState) GetSlot() github_com_prysmaticlabs_eth2_types.Slot {
	if x != nil {
		return x.Slot
	}
	return github_com_prysmaticlabs_eth2_types.Slot(0)
}

func (x *BeaconState) GetFork() *Fork {
	if x != nil {
		return x.Fork
	}
	return nil
}

func (x *BeaconState) GetLatestBlockHeader() *BeaconBlockHeader {
	if x != nil {
		return x.LatestBlockHeader
	}
	return nil
}

func (x *BeaconState) GetBlockRoots() [][]byte {
	if x != nil {
		return x.BlockRoots
	}
	return nil
}

func (x *BeaconState) GetStateRoots() [][]byte {
	if x != nil {
		return x.StateRoots
	}
	return nil
}

func (x *BeaconState) GetHistoricalRoots() [][]byte {
	if x != nil {
		return x.HistoricalRoots
	}
	return nil
}

func (x *BeaconState) GetEth1Data() *Eth1Data {
	if x != nil {
		return x.Eth1Data
	}
	return nil
}

func (x *BeaconState) GetEth1DataVotes() []*Eth1Data {
	if x != nil {
		return x.Eth1DataVotes
	}
	return nil
}

func (x *BeaconState) GetEth1DepositIndex() uint64 {
	if x != nil {
		return x.Eth1DepositIndex
	}
	return 0
}

func (x *BeaconState) GetValidators() []*Validator {
	if x != nil {
		return x.Validators
	}
	return nil
}

func (x *BeaconState) GetBalances() []uint64 {
	if x != nil {
		return x.Balances
	}
	return nil
}

func (x *BeaconState) GetRandaoMixes() [][]byte {
	if x != nil {
		return x.RandaoMixes
	}
	return nil
}

func (x *BeaconState) GetSlashings() []uint64 {
	if x != nil {
		return x.Slashings
	}
	return nil
}

func (x *BeaconState) GetPreviousEpochAttestations() []*PendingAttestation {
	if x != nil {
		return x.PreviousEpochAttestations
	}
	return nil
}

func (x *BeaconState) GetCurrentEpochAttestations() []*PendingAttestation {
	if x != nil {
		return x.CurrentEpochAttestations
	}
	return nil
}

func (x *BeaconState) GetJustificationBits() github_com_prysmaticlabs_go_bitfield.Bitvector4 {
	if x != nil {
		return x.JustificationBits
	}
	return github_com_prysmaticlabs_go_bitfield.Bitvector4(nil)
}

func (x *BeaconState) GetPreviousJustifiedCheckpoint() *Checkpoint {
	if x != nil {
		return x.PreviousJustifiedCheckpoint
	}
	return nil
}

func (x *BeaconState) GetCurrentJustifiedCheckpoint() *Checkpoint {
	if x != nil {
		return x.CurrentJustifiedCheckpoint
	}
	return nil
}

func (x *BeaconState) GetFinalizedCheckpoint() *Checkpoint {
	if x != nil {
		return x.FinalizedCheckpoint
	}
	return nil
}

type PendingAttestation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AggregationBits github_com_prysmaticlabs_go_bitfield.Bitlist       `protobuf:"bytes,1,opt,name=aggregation_bits,json=aggregationBits,proto3" json:"aggregation_bits,omitempty" cast-type:"github.com/prysmaticlabs/go-bitfield.Bitlist" ssz-max:"2048"`
	Data            *AttestationData                                   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	InclusionDelay  github_com_prysmaticlabs_eth2_types.Slot           `protobuf:"varint,3,opt,name=inclusion_delay,json=inclusionDelay,proto3" json:"inclusion_delay,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.Slot"`
	ProposerIndex   github_com_prysmaticlabs_eth2_types.ValidatorIndex `protobuf:"varint,4,opt,name=proposer_index,json=proposerIndex,proto3" json:"proposer_index,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.ValidatorIndex"`
}

func (x *PendingAttestation) Reset() {
	*x = PendingAttestation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_beacon_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PendingAttestation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PendingAttestation) ProtoMessage() {}

func (x *PendingAttestation) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_beacon_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PendingAttestation.ProtoReflect.Descriptor instead.
func (*PendingAttestation) Descriptor() ([]byte, []int) {
	return file_eth_v1_beacon_state_proto_rawDescGZIP(), []int{1}
}

func (x *PendingAttestation) GetAggregationBits() github_com_prysmaticlabs_go_bitfield.Bitlist {
	if x != nil {
		return x.AggregationBits
	}
	return github_com_prysmaticlabs_go_bitfield.Bitlist(nil)
}

func (x *PendingAttestation) GetData() *AttestationData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PendingAttestation) GetInclusionDelay() github_com_prysmaticlabs_eth2_types.Slot {
	if x != nil {
		return x.InclusionDelay
	}
	return github_com_prysmaticlabs_eth2_types.Slot(0)
}

func (x *PendingAttestation) GetProposerIndex() github_com_prysmaticlabs_eth2_types.ValidatorIndex {
	if x != nil {
		return x.ProposerIndex
	}
	return github_com_prysmaticlabs_eth2_types.ValidatorIndex(0)
}

type Committee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index      uint64                                   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Slot       github_com_prysmaticlabs_eth2_types.Slot `protobuf:"varint,2,opt,name=slot,proto3" json:"slot,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.Slot"`
	Validators []uint64                                 `protobuf:"varint,3,rep,packed,name=validators,proto3" json:"validators,omitempty" ssz-max:"1099511627776"`
}

func (x *Committee) Reset() {
	*x = Committee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_beacon_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Committee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Committee) ProtoMessage() {}

func (x *Committee) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_beacon_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Committee.ProtoReflect.Descriptor instead.
func (*Committee) Descriptor() ([]byte, []int) {
	return file_eth_v1_beacon_state_proto_rawDescGZIP(), []int{2}
}

func (x *Committee) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Committee) GetSlot() github_com_prysmaticlabs_eth2_types.Slot {
	if x != nil {
		return x.Slot
	}
	return github_com_prysmaticlabs_eth2_types.Slot(0)
}

func (x *Committee) GetValidators() []uint64 {
	if x != nil {
		return x.Validators
	}
	return nil
}

type Fork struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreviousVersion []byte                                    `protobuf:"bytes,1,opt,name=previous_version,json=previousVersion,proto3" json:"previous_version,omitempty" ssz-size:"4"`
	CurrentVersion  []byte                                    `protobuf:"bytes,2,opt,name=current_version,json=currentVersion,proto3" json:"current_version,omitempty" ssz-size:"4"`
	Epoch           github_com_prysmaticlabs_eth2_types.Epoch `protobuf:"varint,3,opt,name=epoch,proto3" json:"epoch,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.Epoch"`
}

func (x *Fork) Reset() {
	*x = Fork{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_beacon_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fork) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fork) ProtoMessage() {}

func (x *Fork) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_beacon_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fork.ProtoReflect.Descriptor instead.
func (*Fork) Descriptor() ([]byte, []int) {
	return file_eth_v1_beacon_state_proto_rawDescGZIP(), []int{3}
}

func (x *Fork) GetPreviousVersion() []byte {
	if x != nil {
		return x.PreviousVersion
	}
	return nil
}

func (x *Fork) GetCurrentVersion() []byte {
	if x != nil {
		return x.CurrentVersion
	}
	return nil
}

func (x *Fork) GetEpoch() github_com_prysmaticlabs_eth2_types.Epoch {
	if x != nil {
		return x.Epoch
	}
	return github_com_prysmaticlabs_eth2_types.Epoch(0)
}

var File_eth_v1_beacon_state_proto protoreflect.FileDescriptor

var file_eth_v1_beacon_state_proto_rawDesc = []byte{
	0x0a, 0x19, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x65, 0x74, 0x68, 0x2f, 0x65, 0x78, 0x74, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x74,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x65, 0x74, 0x68, 0x2f,
	0x76, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa3, 0x0b, 0x0a, 0x0b, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69,
	0x73, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x17, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73,
	0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x5f, 0x72, 0x6f, 0x6f, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x15,
	0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x40, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x2c, 0x82, 0xb5, 0x18, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62,
	0x73, 0x2f, 0x65, 0x74, 0x68, 0x32, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x6f,
	0x74, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x66, 0x6f, 0x72, 0x6b, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x6b, 0x52, 0x04, 0x66, 0x6f,
	0x72, 0x6b, 0x12, 0x52, 0x0a, 0x13, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x11, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x72, 0x6f, 0x6f, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0c, 0x42, 0x08, 0x8a, 0xb5, 0x18,
	0x04, 0x38, 0x31, 0x39, 0x32, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x6f, 0x6f, 0x74,
	0x73, 0x12, 0x29, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0c, 0x42, 0x08, 0x8a, 0xb5, 0x18, 0x04, 0x38, 0x31, 0x39, 0x32,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x73, 0x12, 0x3d, 0x0a, 0x10,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x0c, 0x42, 0x12, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x92, 0xb5,
	0x18, 0x08, 0x31, 0x36, 0x37, 0x37, 0x37, 0x32, 0x31, 0x36, 0x52, 0x0f, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x6f, 0x6f, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x09, 0x65,
	0x74, 0x68, 0x31, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x74, 0x68, 0x31, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x65, 0x74, 0x68, 0x31, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x4b, 0x0a, 0x0f, 0x65, 0x74, 0x68, 0x31, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65,
	0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x74, 0x68, 0x31, 0x44, 0x61, 0x74, 0x61, 0x42, 0x08, 0x92, 0xb5, 0x18, 0x04, 0x31, 0x30, 0x32,
	0x34, 0x52, 0x0d, 0x65, 0x74, 0x68, 0x31, 0x44, 0x61, 0x74, 0x61, 0x56, 0x6f, 0x74, 0x65, 0x73,
	0x12, 0x2c, 0x0a, 0x12, 0x65, 0x74, 0x68, 0x31, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x65, 0x74,
	0x68, 0x31, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x4d,
	0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x0c, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x11,
	0x92, 0xb5, 0x18, 0x0d, 0x31, 0x30, 0x39, 0x39, 0x35, 0x31, 0x31, 0x36, 0x32, 0x37, 0x37, 0x37,
	0x36, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x2d, 0x0a,
	0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x04, 0x42,
	0x11, 0x92, 0xb5, 0x18, 0x0d, 0x31, 0x30, 0x39, 0x39, 0x35, 0x31, 0x31, 0x36, 0x32, 0x37, 0x37,
	0x37, 0x36, 0x52, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x0c,
	0x72, 0x61, 0x6e, 0x64, 0x61, 0x6f, 0x5f, 0x6d, 0x69, 0x78, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x03,
	0x28, 0x0c, 0x42, 0x09, 0x8a, 0xb5, 0x18, 0x05, 0x36, 0x35, 0x35, 0x33, 0x36, 0x52, 0x0b, 0x72,
	0x61, 0x6e, 0x64, 0x61, 0x6f, 0x4d, 0x69, 0x78, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x09, 0x73, 0x6c,
	0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x04, 0x42, 0x09, 0x8a,
	0xb5, 0x18, 0x05, 0x36, 0x35, 0x35, 0x33, 0x36, 0x52, 0x09, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x6d, 0x0a, 0x1b, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f,
	0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0x92,
	0xb5, 0x18, 0x04, 0x34, 0x30, 0x39, 0x36, 0x52, 0x19, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75,
	0x73, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x6b, 0x0a, 0x1a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x70,
	0x6f, 0x63, 0x68, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75,
	0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0x92, 0xb5, 0x18,
	0x04, 0x34, 0x30, 0x39, 0x36, 0x52, 0x18, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x45, 0x70,
	0x6f, 0x63, 0x68, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x67, 0x0a, 0x12, 0x6a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x62, 0x69, 0x74, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x38, 0x8a, 0xb5, 0x18,
	0x01, 0x31, 0x82, 0xb5, 0x18, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x67,
	0x6f, 0x2d, 0x62, 0x69, 0x74, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x42, 0x69, 0x74, 0x76, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x34, 0x52, 0x11, 0x6a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x74, 0x73, 0x12, 0x5f, 0x0a, 0x1d, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x6f, 0x75, 0x73, 0x5f, 0x6a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x1b, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x4a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x5d, 0x0a, 0x1c, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x1a, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x4a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x4e, 0x0a, 0x14, 0x66, 0x69, 0x6e, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75,
	0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x13, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0xe5, 0x02, 0x0a, 0x12, 0x50, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x63, 0x0a, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62,
	0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x38, 0x92, 0xb5, 0x18, 0x04, 0x32,
	0x30, 0x34, 0x38, 0x82, 0xb5, 0x18, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f,
	0x67, 0x6f, 0x2d, 0x62, 0x69, 0x74, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x42, 0x69, 0x74, 0x6c,
	0x69, 0x73, 0x74, 0x52, 0x0f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x69, 0x74, 0x73, 0x12, 0x34, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x55, 0x0a, 0x0f, 0x69, 0x6e,
	0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x2c, 0x82, 0xb5, 0x18, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62,
	0x73, 0x2f, 0x65, 0x74, 0x68, 0x32, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x6f,
	0x74, 0x52, 0x0e, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x61,
	0x79, 0x12, 0x5d, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x42, 0x36, 0x82, 0xb5, 0x18, 0x32, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61,
	0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x32, 0x2d, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x22, 0x83, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x40, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x2c, 0x82, 0xb5, 0x18, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73,
	0x2f, 0x65, 0x74, 0x68, 0x32, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x6f, 0x74,
	0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x22, 0xad, 0x01, 0x0a, 0x04, 0x46, 0x6f, 0x72, 0x6b, 0x12,
	0x30, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x05, 0x8a, 0xb5, 0x18, 0x01, 0x34,
	0x52, 0x0f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x2e, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x05, 0x8a, 0xb5, 0x18, 0x01,
	0x34, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x43, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x2d, 0x82, 0xb5, 0x18, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65,
	0x74, 0x68, 0x32, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x52,
	0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x42, 0x7b, 0x0a, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x42,
	0x65, 0x61, 0x63, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72,
	0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0xaa,
	0x02, 0x0f, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x45, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0xca, 0x02, 0x0f, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5c, 0x45, 0x74, 0x68,
	0x5c, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eth_v1_beacon_state_proto_rawDescOnce sync.Once
	file_eth_v1_beacon_state_proto_rawDescData = file_eth_v1_beacon_state_proto_rawDesc
)

func file_eth_v1_beacon_state_proto_rawDescGZIP() []byte {
	file_eth_v1_beacon_state_proto_rawDescOnce.Do(func() {
		file_eth_v1_beacon_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_eth_v1_beacon_state_proto_rawDescData)
	})
	return file_eth_v1_beacon_state_proto_rawDescData
}

var file_eth_v1_beacon_state_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eth_v1_beacon_state_proto_goTypes = []interface{}{
	(*BeaconState)(nil),        // 0: ethereum.eth.v1.BeaconState
	(*PendingAttestation)(nil), // 1: ethereum.eth.v1.PendingAttestation
	(*Committee)(nil),          // 2: ethereum.eth.v1.Committee
	(*Fork)(nil),               // 3: ethereum.eth.v1.Fork
	(*BeaconBlockHeader)(nil),  // 4: ethereum.eth.v1.BeaconBlockHeader
	(*Eth1Data)(nil),           // 5: ethereum.eth.v1.Eth1Data
	(*Validator)(nil),          // 6: ethereum.eth.v1.Validator
	(*Checkpoint)(nil),         // 7: ethereum.eth.v1.Checkpoint
	(*AttestationData)(nil),    // 8: ethereum.eth.v1.AttestationData
}
var file_eth_v1_beacon_state_proto_depIdxs = []int32{
	3,  // 0: ethereum.eth.v1.BeaconState.fork:type_name -> ethereum.eth.v1.Fork
	4,  // 1: ethereum.eth.v1.BeaconState.latest_block_header:type_name -> ethereum.eth.v1.BeaconBlockHeader
	5,  // 2: ethereum.eth.v1.BeaconState.eth1_data:type_name -> ethereum.eth.v1.Eth1Data
	5,  // 3: ethereum.eth.v1.BeaconState.eth1_data_votes:type_name -> ethereum.eth.v1.Eth1Data
	6,  // 4: ethereum.eth.v1.BeaconState.validators:type_name -> ethereum.eth.v1.Validator
	1,  // 5: ethereum.eth.v1.BeaconState.previous_epoch_attestations:type_name -> ethereum.eth.v1.PendingAttestation
	1,  // 6: ethereum.eth.v1.BeaconState.current_epoch_attestations:type_name -> ethereum.eth.v1.PendingAttestation
	7,  // 7: ethereum.eth.v1.BeaconState.previous_justified_checkpoint:type_name -> ethereum.eth.v1.Checkpoint
	7,  // 8: ethereum.eth.v1.BeaconState.current_justified_checkpoint:type_name -> ethereum.eth.v1.Checkpoint
	7,  // 9: ethereum.eth.v1.BeaconState.finalized_checkpoint:type_name -> ethereum.eth.v1.Checkpoint
	8,  // 10: ethereum.eth.v1.PendingAttestation.data:type_name -> ethereum.eth.v1.AttestationData
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_eth_v1_beacon_state_proto_init() }
func file_eth_v1_beacon_state_proto_init() {
	if File_eth_v1_beacon_state_proto != nil {
		return
	}
	file_eth_v1_attestation_proto_init()
	file_eth_v1_beacon_block_proto_init()
	file_eth_v1_validator_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eth_v1_beacon_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeaconState); i {
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
		file_eth_v1_beacon_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PendingAttestation); i {
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
		file_eth_v1_beacon_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Committee); i {
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
		file_eth_v1_beacon_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fork); i {
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
			RawDescriptor: file_eth_v1_beacon_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eth_v1_beacon_state_proto_goTypes,
		DependencyIndexes: file_eth_v1_beacon_state_proto_depIdxs,
		MessageInfos:      file_eth_v1_beacon_state_proto_msgTypes,
	}.Build()
	File_eth_v1_beacon_state_proto = out.File
	file_eth_v1_beacon_state_proto_rawDesc = nil
	file_eth_v1_beacon_state_proto_goTypes = nil
	file_eth_v1_beacon_state_proto_depIdxs = nil
}
