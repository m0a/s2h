package struct2json

import (
	"testing"
	"fmt"
)


func TestStructFieldisNil(t *testing.T) {

	type SturctFieledNil struct{
		ptr *SturctFieledNil
		A string
	}

	testValue2 := SturctFieledNil{
		A: "sampleB",
		ptr: nil,
	}

	testValue := SturctFieledNil{
		A: "sampleA",
		ptr: &testValue2,
	}
	v := Create(testValue)
	v = v.StMembers["0"]
	if len(v.StMembers) != 3 {
		t.Errorf("expect: 3, but: %d", len(v.StMembers))
	}
	if v.Type != "struct" {
		t.Errorf("expect: struct but %s", v.Type)
	}
	if v.StMembers["Hell"].PrimValue != "hell?" {
		t.Errorf("expect: hell? but %s", v.StMembers["Hell"].PrimValue)
	}
}


type HogeHoge struct {
	Hell  string
	World int
	A     interface{}
}


func TestNestStruct(t *testing.T) {

	c:=	HogeHoge {
		Hell: "hellll",
		World: 2,
	}
	myValue := &HogeHoge {
		Hell: "hell?",
		World: 12,
		A: &c,
	}

	v := Create(myValue)
	v = v.StMembers["0"]
	if len(v.StMembers) != 3 {
		t.Errorf("expect: 3, but: %d", len(v.StMembers))
	}
	if v.Type != "struct" {
		t.Errorf("expect: struct but %s", v.Type)
	}
	if v.StMembers["Hell"].PrimValue != "hell?" {
		t.Errorf("expect: hell? but %s", v.StMembers["Hell"].PrimValue)
	}
}

func TestPtrValue(t *testing.T) {
	myValue := &HogeHoge {
		"hell?",
		12,
		"c",
	}

	v := Create(myValue)
	v = v.StMembers["0"]
	if len(v.StMembers) != 3 {
		t.Errorf("expect: 3, but: %d", len(v.StMembers))
	}
	if v.Type != "struct" {
		t.Errorf("expect: struct but %s", v.Type)
	}
	if v.StMembers["Hell"].PrimValue != "hell?" {
		t.Errorf("expect: hell? but %s", v.StMembers["Hell"].PrimValue)
	}
}



func TestStructValue(t *testing.T) {
	myValue := HogeHoge {
		"hell?",
		12,
		"c",
	}

	v := Create(myValue)
	//t.Logf("%#v", v)
	//json, _ := json.Marshal(v)
	//t.Logf("%v", string(json))
	if len(v.StMembers) != 3 {
		t.Errorf("expect: 3, but: %d", len(v.StMembers))
	}
	if v.Type != "struct" {
		t.Errorf("expect: struct but %s", v.Type)
	}
	if v.StMembers["Hell"].PrimValue != "hell?" {
		t.Errorf("expect: hell? but %s", v.StMembers["Hell"].PrimValue)
	}
}


func TestArrayValue(t *testing.T) {
	myValue := [...]string{
		"a","b","c",
	}

	v := Create(myValue)
	if len(v.StMembers) != 3 {
		t.Errorf("expect: 3, but: %d", len(v.StMembers))
	}
	if v.Type != "array" {
		t.Errorf("expect: array but %s", v.Type)
	}
	if v.StMembers["0"].PrimValue != "a" {
		t.Errorf("%#v", v.StMembers)
		t.Errorf("expect: a but %s", v.StMembers["0"].PrimValue)

	}
}


func TestStringValue(t *testing.T) {
	myValue := "something"
	v := Create(myValue)
	expect :=`struct2json.Struct2json{Order:0, Type:"string", PrimValue:"something", StMembers:map[string]struct2json.Struct2json(nil)}`
	result := fmt.Sprintf("%#v", v)
	if result != expect {
		t.Errorf ("expect: \n\t%v\n but: \n\t%s\n", expect, result)
	}
}


func TestIntValue(t *testing.T) {
	myValue := 2
	v := Create(myValue)
	expect :=`struct2json.Struct2json{Order:0, Type:"int", PrimValue:"2", StMembers:map[string]struct2json.Struct2json(nil)}`
	result := fmt.Sprintf("%#v", v)
	if result != expect {
		t.Errorf ("expect: \n\t%v\n but: \n\t%s\n", expect, result)
	}
}

