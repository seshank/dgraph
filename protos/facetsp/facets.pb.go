// Code generated by protoc-gen-gogo.
// source: protos/facetsp/facets.proto
// DO NOT EDIT!

/*
	Package facetsp is a generated protocol buffer package.

	It is generated from these files:
		protos/facetsp/facets.proto

	It has these top-level messages:
		Facet
		Param
		Facets
		List
		Function
		FilterTree
*/
package facetsp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Facet_ValType int32

const (
	Facet_STRING   Facet_ValType = 0
	Facet_INT      Facet_ValType = 1
	Facet_FLOAT    Facet_ValType = 2
	Facet_BOOL     Facet_ValType = 3
	Facet_DATETIME Facet_ValType = 4
)

var Facet_ValType_name = map[int32]string{
	0: "STRING",
	1: "INT",
	2: "FLOAT",
	3: "BOOL",
	4: "DATETIME",
}
var Facet_ValType_value = map[string]int32{
	"STRING":   0,
	"INT":      1,
	"FLOAT":    2,
	"BOOL":     3,
	"DATETIME": 4,
}

func (x Facet_ValType) String() string {
	return proto.EnumName(Facet_ValType_name, int32(x))
}
func (Facet_ValType) EnumDescriptor() ([]byte, []int) { return fileDescriptorFacets, []int{0, 0} }

type Facet struct {
	Key     string        `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value   []byte        `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ValType Facet_ValType `protobuf:"varint,3,opt,name=val_type,json=valType,proto3,enum=facetsp.Facet_ValType" json:"val_type,omitempty"`
	Tokens  []string      `protobuf:"bytes,4,rep,name=tokens" json:"tokens,omitempty"`
}

func (m *Facet) Reset()                    { *m = Facet{} }
func (m *Facet) String() string            { return proto.CompactTextString(m) }
func (*Facet) ProtoMessage()               {}
func (*Facet) Descriptor() ([]byte, []int) { return fileDescriptorFacets, []int{0} }

func (m *Facet) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Facet) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Facet) GetValType() Facet_ValType {
	if m != nil {
		return m.ValType
	}
	return Facet_STRING
}

func (m *Facet) GetTokens() []string {
	if m != nil {
		return m.Tokens
	}
	return nil
}

type Param struct {
	AllKeys bool     `protobuf:"varint,1,opt,name=all_keys,json=allKeys,proto3" json:"all_keys,omitempty"`
	Keys    []string `protobuf:"bytes,2,rep,name=keys" json:"keys,omitempty"`
}

func (m *Param) Reset()                    { *m = Param{} }
func (m *Param) String() string            { return proto.CompactTextString(m) }
func (*Param) ProtoMessage()               {}
func (*Param) Descriptor() ([]byte, []int) { return fileDescriptorFacets, []int{1} }

func (m *Param) GetAllKeys() bool {
	if m != nil {
		return m.AllKeys
	}
	return false
}

func (m *Param) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type Facets struct {
	Facets []*Facet `protobuf:"bytes,1,rep,name=facets" json:"facets,omitempty"`
}

func (m *Facets) Reset()                    { *m = Facets{} }
func (m *Facets) String() string            { return proto.CompactTextString(m) }
func (*Facets) ProtoMessage()               {}
func (*Facets) Descriptor() ([]byte, []int) { return fileDescriptorFacets, []int{2} }

func (m *Facets) GetFacets() []*Facet {
	if m != nil {
		return m.Facets
	}
	return nil
}

type List struct {
	FacetsList []*Facets `protobuf:"bytes,1,rep,name=facets_list,json=facetsList" json:"facets_list,omitempty"`
}

func (m *List) Reset()                    { *m = List{} }
func (m *List) String() string            { return proto.CompactTextString(m) }
func (*List) ProtoMessage()               {}
func (*List) Descriptor() ([]byte, []int) { return fileDescriptorFacets, []int{3} }

func (m *List) GetFacetsList() []*Facets {
	if m != nil {
		return m.FacetsList
	}
	return nil
}

type Function struct {
	Name string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Key  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Args []string `protobuf:"bytes,3,rep,name=args" json:"args,omitempty"`
}

func (m *Function) Reset()                    { *m = Function{} }
func (m *Function) String() string            { return proto.CompactTextString(m) }
func (*Function) ProtoMessage()               {}
func (*Function) Descriptor() ([]byte, []int) { return fileDescriptorFacets, []int{4} }

func (m *Function) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Function) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Function) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

// Op and Children are internal nodes and Func on leaves.
type FilterTree struct {
	Op       string        `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	Children []*FilterTree `protobuf:"bytes,2,rep,name=children" json:"children,omitempty"`
	Func     *Function     `protobuf:"bytes,3,opt,name=func" json:"func,omitempty"`
}

func (m *FilterTree) Reset()                    { *m = FilterTree{} }
func (m *FilterTree) String() string            { return proto.CompactTextString(m) }
func (*FilterTree) ProtoMessage()               {}
func (*FilterTree) Descriptor() ([]byte, []int) { return fileDescriptorFacets, []int{5} }

func (m *FilterTree) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *FilterTree) GetChildren() []*FilterTree {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *FilterTree) GetFunc() *Function {
	if m != nil {
		return m.Func
	}
	return nil
}

func init() {
	proto.RegisterType((*Facet)(nil), "facetsp.Facet")
	proto.RegisterType((*Param)(nil), "facetsp.Param")
	proto.RegisterType((*Facets)(nil), "facetsp.Facets")
	proto.RegisterType((*List)(nil), "facetsp.List")
	proto.RegisterType((*Function)(nil), "facetsp.Function")
	proto.RegisterType((*FilterTree)(nil), "facetsp.FilterTree")
	proto.RegisterEnum("facetsp.Facet_ValType", Facet_ValType_name, Facet_ValType_value)
}
func (m *Facet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Facet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFacets(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFacets(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	if m.ValType != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintFacets(dAtA, i, uint64(m.ValType))
	}
	if len(m.Tokens) > 0 {
		for _, s := range m.Tokens {
			dAtA[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *Param) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Param) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AllKeys {
		dAtA[i] = 0x8
		i++
		if m.AllKeys {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Keys) > 0 {
		for _, s := range m.Keys {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *Facets) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Facets) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Facets) > 0 {
		for _, msg := range m.Facets {
			dAtA[i] = 0xa
			i++
			i = encodeVarintFacets(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *List) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *List) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.FacetsList) > 0 {
		for _, msg := range m.FacetsList {
			dAtA[i] = 0xa
			i++
			i = encodeVarintFacets(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Function) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Function) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFacets(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Key) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFacets(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.Args) > 0 {
		for _, s := range m.Args {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *FilterTree) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FilterTree) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Op) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFacets(dAtA, i, uint64(len(m.Op)))
		i += copy(dAtA[i:], m.Op)
	}
	if len(m.Children) > 0 {
		for _, msg := range m.Children {
			dAtA[i] = 0x12
			i++
			i = encodeVarintFacets(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Func != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintFacets(dAtA, i, uint64(m.Func.Size()))
		n1, err := m.Func.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeFixed64Facets(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Facets(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintFacets(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Facet) Size() (n int) {
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovFacets(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovFacets(uint64(l))
	}
	if m.ValType != 0 {
		n += 1 + sovFacets(uint64(m.ValType))
	}
	if len(m.Tokens) > 0 {
		for _, s := range m.Tokens {
			l = len(s)
			n += 1 + l + sovFacets(uint64(l))
		}
	}
	return n
}

func (m *Param) Size() (n int) {
	var l int
	_ = l
	if m.AllKeys {
		n += 2
	}
	if len(m.Keys) > 0 {
		for _, s := range m.Keys {
			l = len(s)
			n += 1 + l + sovFacets(uint64(l))
		}
	}
	return n
}

func (m *Facets) Size() (n int) {
	var l int
	_ = l
	if len(m.Facets) > 0 {
		for _, e := range m.Facets {
			l = e.Size()
			n += 1 + l + sovFacets(uint64(l))
		}
	}
	return n
}

func (m *List) Size() (n int) {
	var l int
	_ = l
	if len(m.FacetsList) > 0 {
		for _, e := range m.FacetsList {
			l = e.Size()
			n += 1 + l + sovFacets(uint64(l))
		}
	}
	return n
}

func (m *Function) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovFacets(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovFacets(uint64(l))
	}
	if len(m.Args) > 0 {
		for _, s := range m.Args {
			l = len(s)
			n += 1 + l + sovFacets(uint64(l))
		}
	}
	return n
}

func (m *FilterTree) Size() (n int) {
	var l int
	_ = l
	l = len(m.Op)
	if l > 0 {
		n += 1 + l + sovFacets(uint64(l))
	}
	if len(m.Children) > 0 {
		for _, e := range m.Children {
			l = e.Size()
			n += 1 + l + sovFacets(uint64(l))
		}
	}
	if m.Func != nil {
		l = m.Func.Size()
		n += 1 + l + sovFacets(uint64(l))
	}
	return n
}

func sovFacets(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFacets(x uint64) (n int) {
	return sovFacets(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Facet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFacets
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Facet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Facet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFacets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFacets
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFacets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFacets
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValType", wireType)
			}
			m.ValType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFacets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValType |= (Facet_ValType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFacets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFacets
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tokens = append(m.Tokens, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFacets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFacets
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Param) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFacets
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Param: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Param: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllKeys", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFacets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AllKeys = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFacets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFacets
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Keys = append(m.Keys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFacets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFacets
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Facets) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFacets
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Facets: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Facets: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Facets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFacets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFacets
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Facets = append(m.Facets, &Facet{})
			if err := m.Facets[len(m.Facets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFacets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFacets
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *List) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFacets
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: List: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: List: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FacetsList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFacets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFacets
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FacetsList = append(m.FacetsList, &Facets{})
			if err := m.FacetsList[len(m.FacetsList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFacets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFacets
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Function) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFacets
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Function: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Function: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFacets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFacets
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFacets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFacets
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Args", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFacets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFacets
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Args = append(m.Args, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFacets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFacets
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FilterTree) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFacets
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FilterTree: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FilterTree: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFacets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFacets
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Op = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Children", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFacets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFacets
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Children = append(m.Children, &FilterTree{})
			if err := m.Children[len(m.Children)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Func", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFacets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFacets
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Func == nil {
				m.Func = &Function{}
			}
			if err := m.Func.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFacets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFacets
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFacets(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFacets
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFacets
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFacets
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthFacets
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFacets
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipFacets(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthFacets = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFacets   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("protos/facetsp/facets.proto", fileDescriptorFacets) }

var fileDescriptorFacets = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x52, 0x51, 0x8b, 0xd3, 0x40,
	0x10, 0xbe, 0x4d, 0xd2, 0x24, 0x9d, 0x1e, 0x35, 0x8e, 0x72, 0x44, 0x84, 0x12, 0x02, 0x4a, 0x9e,
	0x7a, 0x67, 0x04, 0xf1, 0xb5, 0xc7, 0x5d, 0xa4, 0x58, 0xaf, 0xb2, 0x06, 0x5f, 0xcb, 0x1a, 0xf7,
	0x34, 0x74, 0x2f, 0x09, 0xd9, 0x6d, 0x21, 0xff, 0xc4, 0x3f, 0xe3, 0xbb, 0x8f, 0xfe, 0x04, 0xa9,
	0x7f, 0x44, 0xb2, 0x59, 0xaf, 0xdc, 0xd3, 0x7e, 0xf3, 0xcd, 0x7c, 0xb3, 0xdf, 0xcc, 0x2e, 0x3c,
	0x6f, 0xda, 0x5a, 0xd5, 0xf2, 0xfc, 0x96, 0x15, 0x5c, 0xc9, 0xc6, 0x9c, 0x73, 0xcd, 0xa2, 0x67,
	0xd8, 0xf8, 0x27, 0x81, 0x51, 0xd6, 0x63, 0x0c, 0xc0, 0xde, 0xf2, 0x2e, 0x24, 0x11, 0x49, 0xc6,
	0xb4, 0x87, 0xf8, 0x14, 0x46, 0x7b, 0x26, 0x76, 0x3c, 0xb4, 0x22, 0x92, 0x9c, 0xd2, 0x21, 0xc0,
	0x57, 0xe0, 0xef, 0x99, 0xd8, 0xa8, 0xae, 0xe1, 0xa1, 0x1d, 0x91, 0x64, 0x9a, 0x9e, 0xcd, 0x4d,
	0xb7, 0xb9, 0xee, 0x34, 0xff, 0xcc, 0x44, 0xde, 0x35, 0x9c, 0x7a, 0xfb, 0x01, 0xe0, 0x19, 0xb8,
	0xaa, 0xde, 0xf2, 0x4a, 0x86, 0x4e, 0x64, 0x27, 0x63, 0x6a, 0xa2, 0x78, 0x01, 0x9e, 0xa9, 0x45,
	0x00, 0xf7, 0x53, 0x4e, 0x97, 0x37, 0xef, 0x82, 0x13, 0xf4, 0xc0, 0x5e, 0xde, 0xe4, 0x01, 0xc1,
	0x31, 0x8c, 0xb2, 0xd5, 0x7a, 0x91, 0x07, 0x16, 0xfa, 0xe0, 0x5c, 0xae, 0xd7, 0xab, 0xc0, 0xc6,
	0x53, 0xf0, 0xaf, 0x16, 0xf9, 0x75, 0xbe, 0xfc, 0x70, 0x1d, 0x38, 0xf1, 0x1b, 0x18, 0x7d, 0x64,
	0x2d, 0xbb, 0xc3, 0x67, 0xe0, 0x33, 0x21, 0x36, 0x5b, 0xde, 0x49, 0x3d, 0x83, 0x4f, 0x3d, 0x26,
	0xc4, 0x7b, 0xde, 0x49, 0x44, 0x70, 0x34, 0x6d, 0xe9, 0xcb, 0x35, 0x8e, 0x2f, 0xc0, 0xd5, 0x66,
	0x25, 0xbe, 0x04, 0x77, 0xb0, 0x1f, 0x92, 0xc8, 0x4e, 0x26, 0xe9, 0xf4, 0xe1, 0x34, 0xd4, 0x64,
	0xe3, 0xb7, 0xe0, 0xac, 0x4a, 0xa9, 0xf0, 0x02, 0x26, 0x03, 0xb3, 0x11, 0xa5, 0x54, 0x46, 0xf4,
	0xe8, 0xa1, 0x48, 0x52, 0x18, 0xe2, 0x5e, 0x11, 0x5f, 0x81, 0x9f, 0xed, 0xaa, 0x42, 0x95, 0x75,
	0xd5, 0x7b, 0xa9, 0xd8, 0x1d, 0x37, 0x6b, 0xd6, 0xf8, 0xff, 0xe6, 0xad, 0xe3, 0xe6, 0x11, 0x1c,
	0xd6, 0x7e, 0x93, 0xa1, 0x3d, 0x38, 0xee, 0x71, 0xac, 0x00, 0xb2, 0x52, 0x28, 0xde, 0xe6, 0x2d,
	0xe7, 0x38, 0x05, 0xab, 0x6e, 0x4c, 0x17, 0xab, 0x6e, 0xf0, 0x1c, 0xfc, 0xe2, 0x7b, 0x29, 0xbe,
	0xb6, 0xbc, 0xd2, 0x73, 0x4e, 0xd2, 0x27, 0x47, 0x4b, 0xf7, 0x32, 0x7a, 0x5f, 0x84, 0x2f, 0xc0,
	0xb9, 0xdd, 0x55, 0x85, 0x7e, 0xc2, 0x49, 0xfa, 0xf8, 0x58, 0x6c, 0x9c, 0x52, 0x9d, 0xbe, 0x0c,
	0x7e, 0x1d, 0x66, 0xe4, 0xf7, 0x61, 0x46, 0xfe, 0x1c, 0x66, 0xe4, 0xc7, 0xdf, 0xd9, 0xc9, 0x17,
	0x57, 0xff, 0xa0, 0xd7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x17, 0xcc, 0xe7, 0x19, 0x60, 0x02,
	0x00, 0x00,
}
