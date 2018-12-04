/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/vendor/k8s.io/api/scheduling/v1beta1/generated.proto
// DO NOT EDIT!

/*
	Package v1beta1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/vendor/k8s.io/api/scheduling/v1beta1/generated.proto

	It has these top-level messages:
		PriorityClass
		PriorityClassList
*/
package v1beta1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *PriorityClass) Reset()                    { *m = PriorityClass{} }
func (*PriorityClass) ProtoMessage()               {}
func (*PriorityClass) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *PriorityClassList) Reset()                    { *m = PriorityClassList{} }
func (*PriorityClassList) ProtoMessage()               {}
func (*PriorityClassList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func init() {
	proto.RegisterType((*PriorityClass)(nil), "k8s.io.api.scheduling.v1beta1.PriorityClass")
	proto.RegisterType((*PriorityClassList)(nil), "k8s.io.api.scheduling.v1beta1.PriorityClassList")
}
func (m *PriorityClass) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriorityClass) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x10
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Value))
	dAtA[i] = 0x18
	i++
	if m.GlobalDefault {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	dAtA[i] = 0x22
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Description)))
	i += copy(dAtA[i:], m.Description)
	return i, nil
}

func (m *PriorityClassList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriorityClassList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ListMeta.Size()))
	n2, err := m.ListMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PriorityClass) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	n += 1 + sovGenerated(uint64(m.Value))
	n += 2
	l = len(m.Description)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *PriorityClassList) Size() (n int) {
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PriorityClass) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PriorityClass{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`GlobalDefault:` + fmt.Sprintf("%v", this.GlobalDefault) + `,`,
		`Description:` + fmt.Sprintf("%v", this.Description) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PriorityClassList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PriorityClassList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "PriorityClass", "PriorityClass", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PriorityClass) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: PriorityClass: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriorityClass: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalDefault", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
			m.GlobalDefault = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *PriorityClassList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: PriorityClassList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriorityClassList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, PriorityClass{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
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
				next, err := skipGenerated(dAtA[start:])
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
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/api/scheduling/v1beta1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0x8e, 0xd3, 0x4c,
	0x14, 0xc5, 0x3d, 0xd9, 0x2f, 0xfa, 0x82, 0xa3, 0x48, 0x60, 0x84, 0x64, 0x45, 0xc2, 0x6b, 0x2d,
	0x8d, 0x0b, 0x76, 0x86, 0x2c, 0x7f, 0x84, 0x44, 0x67, 0x56, 0x20, 0x24, 0x10, 0xe0, 0x82, 0x02,
	0x51, 0x30, 0xb6, 0xef, 0x3a, 0x43, 0x6c, 0x8f, 0x35, 0x73, 0x6d, 0x69, 0x3b, 0x6a, 0x2a, 0x1e,
	0x8a, 0x22, 0xe5, 0x96, 0x5b, 0xad, 0x88, 0x79, 0x11, 0x64, 0xc7, 0xac, 0x13, 0xa2, 0x05, 0x3a,
	0xcf, 0xb9, 0xe7, 0x77, 0x66, 0xee, 0x91, 0xcd, 0x67, 0x8b, 0xc7, 0x9a, 0x0a, 0xc9, 0x16, 0x65,
	0x08, 0x2a, 0x07, 0x04, 0xcd, 0x2a, 0xc8, 0x63, 0xa9, 0x58, 0x37, 0xe0, 0x85, 0x60, 0x3a, 0x9a,
	0x43, 0x5c, 0xa6, 0x22, 0x4f, 0x58, 0x35, 0x0b, 0x01, 0xf9, 0x8c, 0x25, 0x90, 0x83, 0xe2, 0x08,
	0x31, 0x2d, 0x94, 0x44, 0x69, 0xdd, 0x5e, 0xdb, 0x29, 0x2f, 0x04, 0xed, 0xed, 0xb4, 0xb3, 0x4f,
	0x0f, 0x13, 0x81, 0xf3, 0x32, 0xa4, 0x91, 0xcc, 0x58, 0x22, 0x13, 0xc9, 0x5a, 0x2a, 0x2c, 0x4f,
	0xda, 0x53, 0x7b, 0x68, 0xbf, 0xd6, 0x69, 0xd3, 0x07, 0xfd, 0xe5, 0x19, 0x8f, 0xe6, 0x22, 0x07,
	0x75, 0xca, 0x8a, 0x45, 0xd2, 0x08, 0x9a, 0x65, 0x80, 0x9c, 0x55, 0x3b, 0x6f, 0x98, 0xb2, 0xab,
	0x28, 0x55, 0xe6, 0x28, 0x32, 0xd8, 0x01, 0x1e, 0xfd, 0x0d, 0x68, 0x36, 0xc9, 0xf8, 0x0e, 0x77,
	0xff, 0x2a, 0xae, 0x44, 0x91, 0x32, 0x91, 0xa3, 0x46, 0xf5, 0x3b, 0x74, 0xf0, 0x65, 0x60, 0x4e,
	0xde, 0x28, 0x21, 0x95, 0xc0, 0xd3, 0xa7, 0x29, 0xd7, 0xda, 0xfa, 0x68, 0x8e, 0x9a, 0x55, 0x62,
	0x8e, 0xdc, 0x26, 0x2e, 0xf1, 0xc6, 0x47, 0xf7, 0x68, 0x5f, 0xe3, 0x65, 0x32, 0x2d, 0x16, 0x49,
	0x23, 0x68, 0xda, 0xb8, 0x69, 0x35, 0xa3, 0xaf, 0xc3, 0x4f, 0x10, 0xe1, 0x2b, 0x40, 0xee, 0x5b,
	0xcb, 0x8b, 0x7d, 0xa3, 0xbe, 0xd8, 0x37, 0x7b, 0x2d, 0xb8, 0x4c, 0xb5, 0xee, 0x98, 0xc3, 0x8a,
	0xa7, 0x25, 0xd8, 0x03, 0x97, 0x78, 0x43, 0x7f, 0xd2, 0x99, 0x87, 0xef, 0x1a, 0x31, 0x58, 0xcf,
	0xac, 0x27, 0xe6, 0x24, 0x49, 0x65, 0xc8, 0xd3, 0x63, 0x38, 0xe1, 0x65, 0x8a, 0xf6, 0x9e, 0x4b,
	0xbc, 0x91, 0x7f, 0xab, 0x33, 0x4f, 0x9e, 0x6f, 0x0e, 0x83, 0x6d, 0xaf, 0xf5, 0xd0, 0x1c, 0xc7,
	0xa0, 0x23, 0x25, 0x0a, 0x14, 0x32, 0xb7, 0xff, 0x73, 0x89, 0x77, 0xcd, 0xbf, 0xd9, 0xa1, 0xe3,
	0xe3, 0x7e, 0x14, 0x6c, 0xfa, 0x0e, 0xbe, 0x11, 0xf3, 0xc6, 0x56, 0x19, 0x2f, 0x85, 0x46, 0xeb,
	0xc3, 0x4e, 0x21, 0xf4, 0xdf, 0x0a, 0x69, 0xe8, 0xb6, 0x8e, 0xeb, 0xdd, 0xcd, 0xa3, 0x5f, 0xca,
	0x46, 0x19, 0x6f, 0xcd, 0xa1, 0x40, 0xc8, 0xb4, 0x3d, 0x70, 0xf7, 0xbc, 0xf1, 0xd1, 0x5d, 0xfa,
	0xc7, 0x5f, 0x96, 0x6e, 0x3d, 0xaf, 0xaf, 0xee, 0x45, 0x13, 0x11, 0xac, 0x93, 0xfc, 0xc3, 0xe5,
	0xca, 0x31, 0xce, 0x56, 0x8e, 0x71, 0xbe, 0x72, 0x8c, 0xcf, 0xb5, 0x43, 0x96, 0xb5, 0x43, 0xce,
	0x6a, 0x87, 0x9c, 0xd7, 0x0e, 0xf9, 0x5e, 0x3b, 0xe4, 0xeb, 0x0f, 0xc7, 0x78, 0xff, 0x7f, 0x17,
	0xf9, 0x33, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xb3, 0xc6, 0x7a, 0x6d, 0x03, 0x00, 0x00,
}
