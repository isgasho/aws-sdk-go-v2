// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

type ComplexNestedErrorData struct {
	Foo *string
}

type EmptyStruct struct {
}

type KitchenSink struct {
	Blob                   []byte
	Boolean                *bool
	Double                 *float64
	EmptyStruct            *EmptyStruct
	Float                  *float32
	HttpdateTimestamp      *time.Time
	Integer                *int32
	Iso8601Timestamp       *time.Time
	JsonValue              *string
	ListOfLists            [][]*string
	ListOfMapsOfStrings    []map[string]*string
	ListOfStrings          []*string
	ListOfStructs          []*SimpleStruct
	Long                   *int64
	MapOfListsOfStrings    map[string][]*string
	MapOfMaps              map[string]map[string]*string
	MapOfStrings           map[string]*string
	MapOfStructs           map[string]*SimpleStruct
	RecursiveList          []*KitchenSink
	RecursiveMap           map[string]*KitchenSink
	RecursiveStruct        *KitchenSink
	SimpleStruct           *SimpleStruct
	String_                *string
	StructWithLocationName *StructWithLocationName
	Timestamp              *time.Time
	UnixTimestamp          *time.Time
}

// A union with a representative set of types for members.
type MyUnion interface {
	isMyUnion()
}

type MyUnionMemberStringValue struct {
	Value string
}

func (*MyUnionMemberStringValue) isMyUnion() {}

type MyUnionMemberBooleanValue struct {
	Value bool
}

func (*MyUnionMemberBooleanValue) isMyUnion() {}

type MyUnionMemberNumberValue struct {
	Value int32
}

func (*MyUnionMemberNumberValue) isMyUnion() {}

type MyUnionMemberBlobValue struct {
	Value []byte
}

func (*MyUnionMemberBlobValue) isMyUnion() {}

type MyUnionMemberTimestampValue struct {
	Value time.Time
}

func (*MyUnionMemberTimestampValue) isMyUnion() {}

type MyUnionMemberEnumValue struct {
	Value FooEnum
}

func (*MyUnionMemberEnumValue) isMyUnion() {}

type MyUnionMemberListValue struct {
	Value []*string
}

func (*MyUnionMemberListValue) isMyUnion() {}

type MyUnionMemberMapValue struct {
	Value map[string]*string
}

func (*MyUnionMemberMapValue) isMyUnion() {}

type MyUnionMemberStructureValue struct {
	Value *GreetingStruct
}

func (*MyUnionMemberStructureValue) isMyUnion() {}

type SimpleStruct struct {
	Value *string
}

type StructWithLocationName struct {
	Value *string
}

type GreetingStruct struct {
	Hi *string
}

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte
}

func (*UnknownUnionMember) isMyUnion() {}
