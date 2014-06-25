// Code generated by protoc-gen-gogo.
// source: state.proto
// DO NOT EDIT!

package mesosproto

import testing14 "testing"
import math_rand14 "math/rand"
import time14 "time"
import code_google_com_p_gogoprotobuf_proto8 "github.com/yifan-gu/go-mesos/3rdparty/code.google.com/p/gogoprotobuf/proto"
import testing15 "testing"
import math_rand15 "math/rand"
import time15 "time"
import encoding_json2 "encoding/json"
import testing16 "testing"
import math_rand16 "math/rand"
import time16 "time"
import code_google_com_p_gogoprotobuf_proto9 "github.com/yifan-gu/go-mesos/3rdparty/code.google.com/p/gogoprotobuf/proto"
import math_rand17 "math/rand"
import time17 "time"
import testing17 "testing"
import fmt4 "fmt"
import math_rand18 "math/rand"
import time18 "time"
import testing18 "testing"
import code_google_com_p_gogoprotobuf_proto10 "github.com/yifan-gu/go-mesos/3rdparty/code.google.com/p/gogoprotobuf/proto"
import math_rand19 "math/rand"
import time19 "time"
import testing19 "testing"
import fmt5 "fmt"
import go_parser2 "go/parser"
import math_rand20 "math/rand"
import time20 "time"
import testing20 "testing"
import code_google_com_p_gogoprotobuf_proto11 "github.com/yifan-gu/go-mesos/3rdparty/code.google.com/p/gogoprotobuf/proto"

func TestEntryProto(t *testing14.T) {
	popr := math_rand14.New(math_rand14.NewSource(time14.Now().UnixNano()))
	p := NewPopulatedEntry(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto8.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &Entry{}
	if err := code_google_com_p_gogoprotobuf_proto8.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestEntryMarshalTo(t *testing14.T) {
	popr := math_rand14.New(math_rand14.NewSource(time14.Now().UnixNano()))
	p := NewPopulatedEntry(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &Entry{}
	if err := code_google_com_p_gogoprotobuf_proto8.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkEntryProtoMarshal(b *testing14.B) {
	popr := math_rand14.New(math_rand14.NewSource(616))
	total := 0
	pops := make([]*Entry, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedEntry(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto8.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkEntryProtoUnmarshal(b *testing14.B) {
	popr := math_rand14.New(math_rand14.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto8.Marshal(NewPopulatedEntry(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &Entry{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := code_google_com_p_gogoprotobuf_proto8.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestOperationProto(t *testing14.T) {
	popr := math_rand14.New(math_rand14.NewSource(time14.Now().UnixNano()))
	p := NewPopulatedOperation(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto8.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &Operation{}
	if err := code_google_com_p_gogoprotobuf_proto8.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestOperationMarshalTo(t *testing14.T) {
	popr := math_rand14.New(math_rand14.NewSource(time14.Now().UnixNano()))
	p := NewPopulatedOperation(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &Operation{}
	if err := code_google_com_p_gogoprotobuf_proto8.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkOperationProtoMarshal(b *testing14.B) {
	popr := math_rand14.New(math_rand14.NewSource(616))
	total := 0
	pops := make([]*Operation, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedOperation(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto8.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkOperationProtoUnmarshal(b *testing14.B) {
	popr := math_rand14.New(math_rand14.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto8.Marshal(NewPopulatedOperation(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &Operation{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := code_google_com_p_gogoprotobuf_proto8.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestOperation_SnapshotProto(t *testing14.T) {
	popr := math_rand14.New(math_rand14.NewSource(time14.Now().UnixNano()))
	p := NewPopulatedOperation_Snapshot(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto8.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &Operation_Snapshot{}
	if err := code_google_com_p_gogoprotobuf_proto8.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestOperation_SnapshotMarshalTo(t *testing14.T) {
	popr := math_rand14.New(math_rand14.NewSource(time14.Now().UnixNano()))
	p := NewPopulatedOperation_Snapshot(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &Operation_Snapshot{}
	if err := code_google_com_p_gogoprotobuf_proto8.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkOperation_SnapshotProtoMarshal(b *testing14.B) {
	popr := math_rand14.New(math_rand14.NewSource(616))
	total := 0
	pops := make([]*Operation_Snapshot, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedOperation_Snapshot(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto8.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkOperation_SnapshotProtoUnmarshal(b *testing14.B) {
	popr := math_rand14.New(math_rand14.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto8.Marshal(NewPopulatedOperation_Snapshot(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &Operation_Snapshot{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := code_google_com_p_gogoprotobuf_proto8.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestOperation_ExpungeProto(t *testing14.T) {
	popr := math_rand14.New(math_rand14.NewSource(time14.Now().UnixNano()))
	p := NewPopulatedOperation_Expunge(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto8.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &Operation_Expunge{}
	if err := code_google_com_p_gogoprotobuf_proto8.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestOperation_ExpungeMarshalTo(t *testing14.T) {
	popr := math_rand14.New(math_rand14.NewSource(time14.Now().UnixNano()))
	p := NewPopulatedOperation_Expunge(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &Operation_Expunge{}
	if err := code_google_com_p_gogoprotobuf_proto8.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkOperation_ExpungeProtoMarshal(b *testing14.B) {
	popr := math_rand14.New(math_rand14.NewSource(616))
	total := 0
	pops := make([]*Operation_Expunge, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedOperation_Expunge(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto8.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkOperation_ExpungeProtoUnmarshal(b *testing14.B) {
	popr := math_rand14.New(math_rand14.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto8.Marshal(NewPopulatedOperation_Expunge(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &Operation_Expunge{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := code_google_com_p_gogoprotobuf_proto8.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestEntryJSON(t *testing15.T) {
	popr := math_rand15.New(math_rand15.NewSource(time15.Now().UnixNano()))
	p := NewPopulatedEntry(popr, true)
	jsondata, err := encoding_json2.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &Entry{}
	err = encoding_json2.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestOperationJSON(t *testing15.T) {
	popr := math_rand15.New(math_rand15.NewSource(time15.Now().UnixNano()))
	p := NewPopulatedOperation(popr, true)
	jsondata, err := encoding_json2.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &Operation{}
	err = encoding_json2.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestOperation_SnapshotJSON(t *testing15.T) {
	popr := math_rand15.New(math_rand15.NewSource(time15.Now().UnixNano()))
	p := NewPopulatedOperation_Snapshot(popr, true)
	jsondata, err := encoding_json2.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &Operation_Snapshot{}
	err = encoding_json2.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestOperation_ExpungeJSON(t *testing15.T) {
	popr := math_rand15.New(math_rand15.NewSource(time15.Now().UnixNano()))
	p := NewPopulatedOperation_Expunge(popr, true)
	jsondata, err := encoding_json2.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &Operation_Expunge{}
	err = encoding_json2.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestEntryProtoText(t *testing16.T) {
	popr := math_rand16.New(math_rand16.NewSource(time16.Now().UnixNano()))
	p := NewPopulatedEntry(popr, true)
	data := code_google_com_p_gogoprotobuf_proto9.MarshalTextString(p)
	msg := &Entry{}
	if err := code_google_com_p_gogoprotobuf_proto9.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestEntryProtoCompactText(t *testing16.T) {
	popr := math_rand16.New(math_rand16.NewSource(time16.Now().UnixNano()))
	p := NewPopulatedEntry(popr, true)
	data := code_google_com_p_gogoprotobuf_proto9.CompactTextString(p)
	msg := &Entry{}
	if err := code_google_com_p_gogoprotobuf_proto9.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestOperationProtoText(t *testing16.T) {
	popr := math_rand16.New(math_rand16.NewSource(time16.Now().UnixNano()))
	p := NewPopulatedOperation(popr, true)
	data := code_google_com_p_gogoprotobuf_proto9.MarshalTextString(p)
	msg := &Operation{}
	if err := code_google_com_p_gogoprotobuf_proto9.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestOperationProtoCompactText(t *testing16.T) {
	popr := math_rand16.New(math_rand16.NewSource(time16.Now().UnixNano()))
	p := NewPopulatedOperation(popr, true)
	data := code_google_com_p_gogoprotobuf_proto9.CompactTextString(p)
	msg := &Operation{}
	if err := code_google_com_p_gogoprotobuf_proto9.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestOperation_SnapshotProtoText(t *testing16.T) {
	popr := math_rand16.New(math_rand16.NewSource(time16.Now().UnixNano()))
	p := NewPopulatedOperation_Snapshot(popr, true)
	data := code_google_com_p_gogoprotobuf_proto9.MarshalTextString(p)
	msg := &Operation_Snapshot{}
	if err := code_google_com_p_gogoprotobuf_proto9.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestOperation_SnapshotProtoCompactText(t *testing16.T) {
	popr := math_rand16.New(math_rand16.NewSource(time16.Now().UnixNano()))
	p := NewPopulatedOperation_Snapshot(popr, true)
	data := code_google_com_p_gogoprotobuf_proto9.CompactTextString(p)
	msg := &Operation_Snapshot{}
	if err := code_google_com_p_gogoprotobuf_proto9.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestOperation_ExpungeProtoText(t *testing16.T) {
	popr := math_rand16.New(math_rand16.NewSource(time16.Now().UnixNano()))
	p := NewPopulatedOperation_Expunge(popr, true)
	data := code_google_com_p_gogoprotobuf_proto9.MarshalTextString(p)
	msg := &Operation_Expunge{}
	if err := code_google_com_p_gogoprotobuf_proto9.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestOperation_ExpungeProtoCompactText(t *testing16.T) {
	popr := math_rand16.New(math_rand16.NewSource(time16.Now().UnixNano()))
	p := NewPopulatedOperation_Expunge(popr, true)
	data := code_google_com_p_gogoprotobuf_proto9.CompactTextString(p)
	msg := &Operation_Expunge{}
	if err := code_google_com_p_gogoprotobuf_proto9.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestEntryStringer(t *testing17.T) {
	popr := math_rand17.New(math_rand17.NewSource(time17.Now().UnixNano()))
	p := NewPopulatedEntry(popr, false)
	s1 := p.String()
	s2 := fmt4.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestOperationStringer(t *testing17.T) {
	popr := math_rand17.New(math_rand17.NewSource(time17.Now().UnixNano()))
	p := NewPopulatedOperation(popr, false)
	s1 := p.String()
	s2 := fmt4.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestOperation_SnapshotStringer(t *testing17.T) {
	popr := math_rand17.New(math_rand17.NewSource(time17.Now().UnixNano()))
	p := NewPopulatedOperation_Snapshot(popr, false)
	s1 := p.String()
	s2 := fmt4.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestOperation_ExpungeStringer(t *testing17.T) {
	popr := math_rand17.New(math_rand17.NewSource(time17.Now().UnixNano()))
	p := NewPopulatedOperation_Expunge(popr, false)
	s1 := p.String()
	s2 := fmt4.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestEntrySize(t *testing18.T) {
	popr := math_rand18.New(math_rand18.NewSource(time18.Now().UnixNano()))
	p := NewPopulatedEntry(popr, true)
	size2 := code_google_com_p_gogoprotobuf_proto10.Size(p)
	data, err := code_google_com_p_gogoprotobuf_proto10.Marshal(p)
	if err != nil {
		panic(err)
	}
	size := p.Size()
	if len(data) != size {
		t.Fatalf("size %v != marshalled size %v", size, len(data))
	}
	if size2 != size {
		t.Fatalf("size %v != before marshal proto.Size %v", size, size2)
	}
	size3 := code_google_com_p_gogoprotobuf_proto10.Size(p)
	if size3 != size {
		t.Fatalf("size %v != after marshal proto.Size %v", size, size3)
	}
}

func BenchmarkEntrySize(b *testing18.B) {
	popr := math_rand18.New(math_rand18.NewSource(616))
	total := 0
	pops := make([]*Entry, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedEntry(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestOperationSize(t *testing18.T) {
	popr := math_rand18.New(math_rand18.NewSource(time18.Now().UnixNano()))
	p := NewPopulatedOperation(popr, true)
	size2 := code_google_com_p_gogoprotobuf_proto10.Size(p)
	data, err := code_google_com_p_gogoprotobuf_proto10.Marshal(p)
	if err != nil {
		panic(err)
	}
	size := p.Size()
	if len(data) != size {
		t.Fatalf("size %v != marshalled size %v", size, len(data))
	}
	if size2 != size {
		t.Fatalf("size %v != before marshal proto.Size %v", size, size2)
	}
	size3 := code_google_com_p_gogoprotobuf_proto10.Size(p)
	if size3 != size {
		t.Fatalf("size %v != after marshal proto.Size %v", size, size3)
	}
}

func BenchmarkOperationSize(b *testing18.B) {
	popr := math_rand18.New(math_rand18.NewSource(616))
	total := 0
	pops := make([]*Operation, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedOperation(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestOperation_SnapshotSize(t *testing18.T) {
	popr := math_rand18.New(math_rand18.NewSource(time18.Now().UnixNano()))
	p := NewPopulatedOperation_Snapshot(popr, true)
	size2 := code_google_com_p_gogoprotobuf_proto10.Size(p)
	data, err := code_google_com_p_gogoprotobuf_proto10.Marshal(p)
	if err != nil {
		panic(err)
	}
	size := p.Size()
	if len(data) != size {
		t.Fatalf("size %v != marshalled size %v", size, len(data))
	}
	if size2 != size {
		t.Fatalf("size %v != before marshal proto.Size %v", size, size2)
	}
	size3 := code_google_com_p_gogoprotobuf_proto10.Size(p)
	if size3 != size {
		t.Fatalf("size %v != after marshal proto.Size %v", size, size3)
	}
}

func BenchmarkOperation_SnapshotSize(b *testing18.B) {
	popr := math_rand18.New(math_rand18.NewSource(616))
	total := 0
	pops := make([]*Operation_Snapshot, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedOperation_Snapshot(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestOperation_ExpungeSize(t *testing18.T) {
	popr := math_rand18.New(math_rand18.NewSource(time18.Now().UnixNano()))
	p := NewPopulatedOperation_Expunge(popr, true)
	size2 := code_google_com_p_gogoprotobuf_proto10.Size(p)
	data, err := code_google_com_p_gogoprotobuf_proto10.Marshal(p)
	if err != nil {
		panic(err)
	}
	size := p.Size()
	if len(data) != size {
		t.Fatalf("size %v != marshalled size %v", size, len(data))
	}
	if size2 != size {
		t.Fatalf("size %v != before marshal proto.Size %v", size, size2)
	}
	size3 := code_google_com_p_gogoprotobuf_proto10.Size(p)
	if size3 != size {
		t.Fatalf("size %v != after marshal proto.Size %v", size, size3)
	}
}

func BenchmarkOperation_ExpungeSize(b *testing18.B) {
	popr := math_rand18.New(math_rand18.NewSource(616))
	total := 0
	pops := make([]*Operation_Expunge, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedOperation_Expunge(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestEntryGoString(t *testing19.T) {
	popr := math_rand19.New(math_rand19.NewSource(time19.Now().UnixNano()))
	p := NewPopulatedEntry(popr, false)
	s1 := p.GoString()
	s2 := fmt5.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser2.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func TestOperationGoString(t *testing19.T) {
	popr := math_rand19.New(math_rand19.NewSource(time19.Now().UnixNano()))
	p := NewPopulatedOperation(popr, false)
	s1 := p.GoString()
	s2 := fmt5.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser2.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func TestOperation_SnapshotGoString(t *testing19.T) {
	popr := math_rand19.New(math_rand19.NewSource(time19.Now().UnixNano()))
	p := NewPopulatedOperation_Snapshot(popr, false)
	s1 := p.GoString()
	s2 := fmt5.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser2.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func TestOperation_ExpungeGoString(t *testing19.T) {
	popr := math_rand19.New(math_rand19.NewSource(time19.Now().UnixNano()))
	p := NewPopulatedOperation_Expunge(popr, false)
	s1 := p.GoString()
	s2 := fmt5.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser2.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func TestEntryVerboseEqual(t *testing20.T) {
	popr := math_rand20.New(math_rand20.NewSource(time20.Now().UnixNano()))
	p := NewPopulatedEntry(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto11.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &Entry{}
	if err := code_google_com_p_gogoprotobuf_proto11.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestOperationVerboseEqual(t *testing20.T) {
	popr := math_rand20.New(math_rand20.NewSource(time20.Now().UnixNano()))
	p := NewPopulatedOperation(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto11.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &Operation{}
	if err := code_google_com_p_gogoprotobuf_proto11.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestOperation_SnapshotVerboseEqual(t *testing20.T) {
	popr := math_rand20.New(math_rand20.NewSource(time20.Now().UnixNano()))
	p := NewPopulatedOperation_Snapshot(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto11.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &Operation_Snapshot{}
	if err := code_google_com_p_gogoprotobuf_proto11.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestOperation_ExpungeVerboseEqual(t *testing20.T) {
	popr := math_rand20.New(math_rand20.NewSource(time20.Now().UnixNano()))
	p := NewPopulatedOperation_Expunge(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto11.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &Operation_Expunge{}
	if err := code_google_com_p_gogoprotobuf_proto11.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}

//These tests are generated by github.com/yifan-gu/go-mesos/3rdparty/code.google.com/p/gogoprotobuf/plugin/testgen
