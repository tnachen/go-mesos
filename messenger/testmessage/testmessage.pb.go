// Code generated by protoc-gen-gogo.
// source: testmessage.proto
// DO NOT EDIT!

/*
Package testmessage is a generated protocol buffer package.

It is generated from these files:
	testmessage.proto

It has these top-level messages:
	SmallMessage
	MediumMessage
	BigMessage
	LargeMessage
*/
package testmessage

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"

import io "io"
import code_google_com_p_gogoprotobuf_proto "code.google.com/p/gogoprotobuf/proto"

import fmt "fmt"
import strings "strings"
import reflect "reflect"

import fmt1 "fmt"
import strings1 "strings"
import code_google_com_p_gogoprotobuf_proto1 "code.google.com/p/gogoprotobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect1 "reflect"

import fmt2 "fmt"
import bytes "bytes"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type SmallMessage struct {
	Values           []string `protobuf:"bytes,1,rep" json:"Values,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SmallMessage) Reset()      { *m = SmallMessage{} }
func (*SmallMessage) ProtoMessage() {}

func (m *SmallMessage) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type MediumMessage struct {
	Values           []string `protobuf:"bytes,1,rep" json:"Values,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *MediumMessage) Reset()      { *m = MediumMessage{} }
func (*MediumMessage) ProtoMessage() {}

func (m *MediumMessage) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type BigMessage struct {
	Values           []string `protobuf:"bytes,1,rep" json:"Values,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *BigMessage) Reset()      { *m = BigMessage{} }
func (*BigMessage) ProtoMessage() {}

func (m *BigMessage) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type LargeMessage struct {
	Values           []string `protobuf:"bytes,1,rep" json:"Values,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *LargeMessage) Reset()      { *m = LargeMessage{} }
func (*LargeMessage) ProtoMessage() {}

func (m *LargeMessage) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
}
func (m *SmallMessage) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, string(data[index:postIndex]))
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *MediumMessage) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, string(data[index:postIndex]))
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *BigMessage) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, string(data[index:postIndex]))
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *LargeMessage) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, string(data[index:postIndex]))
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (this *SmallMessage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SmallMessage{`,
		`Values:` + fmt.Sprintf("%v", this.Values) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *MediumMessage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MediumMessage{`,
		`Values:` + fmt.Sprintf("%v", this.Values) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *BigMessage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&BigMessage{`,
		`Values:` + fmt.Sprintf("%v", this.Values) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *LargeMessage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LargeMessage{`,
		`Values:` + fmt.Sprintf("%v", this.Values) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTestmessage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SmallMessage) Size() (n int) {
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			l = len(s)
			n += 1 + l + sovTestmessage(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *MediumMessage) Size() (n int) {
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			l = len(s)
			n += 1 + l + sovTestmessage(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *BigMessage) Size() (n int) {
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			l = len(s)
			n += 1 + l + sovTestmessage(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *LargeMessage) Size() (n int) {
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			l = len(s)
			n += 1 + l + sovTestmessage(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTestmessage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTestmessage(x uint64) (n int) {
	return sovTestmessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func NewPopulatedSmallMessage(r randyTestmessage, easy bool) *SmallMessage {
	this := &SmallMessage{}
	if r.Intn(10) != 0 {
		v1 := r.Intn(10)
		this.Values = make([]string, v1)
		for i := 0; i < v1; i++ {
			this.Values[i] = randStringTestmessage(r)
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTestmessage(r, 2)
	}
	return this
}

func NewPopulatedMediumMessage(r randyTestmessage, easy bool) *MediumMessage {
	this := &MediumMessage{}
	if r.Intn(10) != 0 {
		v2 := r.Intn(10)
		this.Values = make([]string, v2)
		for i := 0; i < v2; i++ {
			this.Values[i] = randStringTestmessage(r)
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTestmessage(r, 2)
	}
	return this
}

func NewPopulatedBigMessage(r randyTestmessage, easy bool) *BigMessage {
	this := &BigMessage{}
	if r.Intn(10) != 0 {
		v3 := r.Intn(10)
		this.Values = make([]string, v3)
		for i := 0; i < v3; i++ {
			this.Values[i] = randStringTestmessage(r)
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTestmessage(r, 2)
	}
	return this
}

func NewPopulatedLargeMessage(r randyTestmessage, easy bool) *LargeMessage {
	this := &LargeMessage{}
	if r.Intn(10) != 0 {
		v4 := r.Intn(10)
		this.Values = make([]string, v4)
		for i := 0; i < v4; i++ {
			this.Values[i] = randStringTestmessage(r)
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTestmessage(r, 2)
	}
	return this
}

type randyTestmessage interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTestmessage(r randyTestmessage) rune {
	res := rune(r.Uint32() % 1112064)
	if 55296 <= res {
		res += 2047
	}
	return res
}
func randStringTestmessage(r randyTestmessage) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneTestmessage(r)
	}
	return string(tmps)
}
func randUnrecognizedTestmessage(r randyTestmessage, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldTestmessage(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldTestmessage(data []byte, r randyTestmessage, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateTestmessage(data, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		data = encodeVarintPopulateTestmessage(data, uint64(v6))
	case 1:
		data = encodeVarintPopulateTestmessage(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateTestmessage(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateTestmessage(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateTestmessage(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateTestmessage(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *SmallMessage) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *SmallMessage) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			data[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *MediumMessage) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *MediumMessage) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			data[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *BigMessage) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *BigMessage) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			data[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *LargeMessage) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *LargeMessage) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			data[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func encodeFixed64Testmessage(data []byte, offset int, v uint64) int {
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
func encodeFixed32Testmessage(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTestmessage(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (this *SmallMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings1.Join([]string{`&testmessage.SmallMessage{` + `Values:` + fmt1.Sprintf("%#v", this.Values), `XXX_unrecognized:` + fmt1.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *MediumMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings1.Join([]string{`&testmessage.MediumMessage{` + `Values:` + fmt1.Sprintf("%#v", this.Values), `XXX_unrecognized:` + fmt1.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *BigMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings1.Join([]string{`&testmessage.BigMessage{` + `Values:` + fmt1.Sprintf("%#v", this.Values), `XXX_unrecognized:` + fmt1.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *LargeMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings1.Join([]string{`&testmessage.LargeMessage{` + `Values:` + fmt1.Sprintf("%#v", this.Values), `XXX_unrecognized:` + fmt1.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func valueToGoStringTestmessage(v interface{}, typ string) string {
	rv := reflect1.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect1.Indirect(rv).Interface()
	return fmt1.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringTestmessage(e map[int32]code_google_com_p_gogoprotobuf_proto1.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings1.Join(ss, ",") + "}"
	return s
}
func (this *SmallMessage) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt2.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*SmallMessage)
	if !ok {
		return fmt2.Errorf("that is not of type *SmallMessage")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt2.Errorf("that is type *SmallMessage but is nil && this != nil")
	} else if this == nil {
		return fmt2.Errorf("that is type *SmallMessagebut is not nil && this == nil")
	}
	if len(this.Values) != len(that1.Values) {
		return fmt2.Errorf("Values this(%v) Not Equal that(%v)", len(this.Values), len(that1.Values))
	}
	for i := range this.Values {
		if this.Values[i] != that1.Values[i] {
			return fmt2.Errorf("Values this[%v](%v) Not Equal that[%v](%v)", i, this.Values[i], i, that1.Values[i])
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt2.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *SmallMessage) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*SmallMessage)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Values) != len(that1.Values) {
		return false
	}
	for i := range this.Values {
		if this.Values[i] != that1.Values[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MediumMessage) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt2.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*MediumMessage)
	if !ok {
		return fmt2.Errorf("that is not of type *MediumMessage")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt2.Errorf("that is type *MediumMessage but is nil && this != nil")
	} else if this == nil {
		return fmt2.Errorf("that is type *MediumMessagebut is not nil && this == nil")
	}
	if len(this.Values) != len(that1.Values) {
		return fmt2.Errorf("Values this(%v) Not Equal that(%v)", len(this.Values), len(that1.Values))
	}
	for i := range this.Values {
		if this.Values[i] != that1.Values[i] {
			return fmt2.Errorf("Values this[%v](%v) Not Equal that[%v](%v)", i, this.Values[i], i, that1.Values[i])
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt2.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *MediumMessage) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*MediumMessage)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Values) != len(that1.Values) {
		return false
	}
	for i := range this.Values {
		if this.Values[i] != that1.Values[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *BigMessage) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt2.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*BigMessage)
	if !ok {
		return fmt2.Errorf("that is not of type *BigMessage")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt2.Errorf("that is type *BigMessage but is nil && this != nil")
	} else if this == nil {
		return fmt2.Errorf("that is type *BigMessagebut is not nil && this == nil")
	}
	if len(this.Values) != len(that1.Values) {
		return fmt2.Errorf("Values this(%v) Not Equal that(%v)", len(this.Values), len(that1.Values))
	}
	for i := range this.Values {
		if this.Values[i] != that1.Values[i] {
			return fmt2.Errorf("Values this[%v](%v) Not Equal that[%v](%v)", i, this.Values[i], i, that1.Values[i])
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt2.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *BigMessage) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*BigMessage)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Values) != len(that1.Values) {
		return false
	}
	for i := range this.Values {
		if this.Values[i] != that1.Values[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LargeMessage) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt2.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*LargeMessage)
	if !ok {
		return fmt2.Errorf("that is not of type *LargeMessage")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt2.Errorf("that is type *LargeMessage but is nil && this != nil")
	} else if this == nil {
		return fmt2.Errorf("that is type *LargeMessagebut is not nil && this == nil")
	}
	if len(this.Values) != len(that1.Values) {
		return fmt2.Errorf("Values this(%v) Not Equal that(%v)", len(this.Values), len(that1.Values))
	}
	for i := range this.Values {
		if this.Values[i] != that1.Values[i] {
			return fmt2.Errorf("Values this[%v](%v) Not Equal that[%v](%v)", i, this.Values[i], i, that1.Values[i])
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt2.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *LargeMessage) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*LargeMessage)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Values) != len(that1.Values) {
		return false
	}
	for i := range this.Values {
		if this.Values[i] != that1.Values[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
