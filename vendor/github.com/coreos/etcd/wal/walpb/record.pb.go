// Code generated by protoc-gen-gogo.
// source: record.proto
// DO NOT EDIT!

/*
	Package walpb is a generated protocol buffer package.

	It is generated from these files:
		record.proto

	It has these top-level messages:
		Record
		Snapshot
*/
package walpb

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Record struct {
	Type             int64  `protobuf:"varint,1,opt,name=type" json:"type"`
	Crc              uint32 `protobuf:"varint,2,opt,name=crc" json:"crc"`
	Data             []byte `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptorRecord, []int{0} }

type Snapshot struct {
	Index            uint64 `protobuf:"varint,1,opt,name=index" json:"index"`
	Term             uint64 `protobuf:"varint,2,opt,name=term" json:"term"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Snapshot) Reset()                    { *m = Snapshot{} }
func (m *Snapshot) String() string            { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()               {}
func (*Snapshot) Descriptor() ([]byte, []int) { return fileDescriptorRecord, []int{1} }

func init() {
	proto.RegisterType((*Record)(nil), "walpb.Record")
	proto.RegisterType((*Snapshot)(nil), "walpb.Snapshot")
}
func (m *Record) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Record) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintRecord(data, i, uint64(m.Type))
	data[i] = 0x10
	i++
	i = encodeVarintRecord(data, i, uint64(m.Crc))
	if m.Data != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintRecord(data, i, uint64(len(m.Data)))
		i += copy(data[i:], m.Data)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Snapshot) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Snapshot) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintRecord(data, i, uint64(m.Index))
	data[i] = 0x10
	i++
	i = encodeVarintRecord(data, i, uint64(m.Term))
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Record(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Record(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRecord(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Record) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovRecord(uint64(m.Type))
	n += 1 + sovRecord(uint64(m.Crc))
	if m.Data != nil {
		l = len(m.Data)
		n += 1 + l + sovRecord(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Snapshot) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovRecord(uint64(m.Index))
	n += 1 + sovRecord(uint64(m.Term))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRecord(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRecord(x uint64) (n int) {
	return sovRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Record) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecord
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Record: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Record: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Type |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Crc", wireType)
			}
			m.Crc = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Crc |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], data[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRecord(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRecord
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Snapshot) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecord
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Snapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Snapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Index |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Term |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRecord(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRecord
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRecord(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRecord
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRecord
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRecord
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipRecord(data[start:])
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
	ErrInvalidLengthRecord = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRecord   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorRecord = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4a, 0x4d, 0xce,
	0x2f, 0x4a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0x4f, 0xcc, 0x29, 0x48, 0x92,
	0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x8b, 0xe8, 0x83, 0x58, 0x10, 0x49, 0x25, 0x3f, 0x2e, 0xb6,
	0x20, 0xb0, 0x62, 0x21, 0x09, 0x2e, 0x96, 0x92, 0xca, 0x82, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0x66, 0x27, 0x96, 0x13, 0xf7, 0xe4, 0x19, 0x82, 0xc0, 0x22, 0x42, 0x62, 0x5c, 0xcc, 0xc9, 0x45,
	0xc9, 0x12, 0x4c, 0x40, 0x09, 0x5e, 0xa8, 0x04, 0x48, 0x40, 0x48, 0x88, 0x8b, 0x25, 0x25, 0xb1,
	0x24, 0x51, 0x82, 0x19, 0x28, 0xc1, 0x13, 0x04, 0x66, 0x2b, 0x39, 0x70, 0x71, 0x04, 0xe7, 0x25,
	0x16, 0x14, 0x67, 0xe4, 0x97, 0x08, 0x49, 0x71, 0xb1, 0x66, 0xe6, 0xa5, 0xa4, 0x56, 0x80, 0x8d,
	0x64, 0x81, 0xea, 0x84, 0x08, 0x81, 0x6d, 0x4b, 0x2d, 0xca, 0x05, 0x1b, 0xca, 0x02, 0xb7, 0x0d,
	0x28, 0xe2, 0x24, 0x72, 0xe2, 0xa1, 0x1c, 0xc3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x80, 0xf8, 0x01,
	0x10, 0xcf, 0x78, 0x2c, 0xc7, 0x00, 0x08, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x5e, 0x5c, 0x46, 0xd3,
	0x00, 0x00, 0x00,
}
