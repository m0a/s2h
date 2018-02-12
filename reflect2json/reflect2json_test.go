package reflect2json

import (
	"testing"
	"fmt"
	"reflect"
	"encoding/json"
)


func TestStructFieldisNil(t *testing.T) {

	type SturctFieledNil struct{
		xxx *SturctFieledNil
		A string
	}

	testValue2 := SturctFieledNil{
		A: "sampleB",
		xxx: nil,
	}

	testValue := SturctFieledNil{
		A: "sampleA",
		xxx: &testValue2,
	}
	v := Create(reflect.ValueOf(testValue))
	v = v.Fields["0"]

	if len(v.Fields) != 0 {
		t.Logf("%#v\n", v)
		t.Errorf("expect: v.stMembers not nil, but: %v", v.Fields)
	}
}


type HogeHoge struct {
	Hell  string
	World int
	A     interface{}
}



func TestNestMap(t *testing.T) {

	myMap := map[string]interface{}{
		"abc":1,
		"def":map[string]int{"hij": 2, "lmn": 3},
	}
	v := Create(reflect.ValueOf(myMap))


	if v.Fields["def"].Fields["lmn"].Value != "3" {
		json, _ := json.Marshal(v)
		t.Logf("%v",string(json))
		t.Errorf("expect: 3 but %s", v.Fields["def"].Fields["lmn"].Value)

	}

}

func TestMap(t *testing.T) {

	myMap := map[string]int{
		"abc":1,
		"def":2,
	}
	v := Create(reflect.ValueOf(myMap))

	//t.Logf("%#v", v)
	//json, _ := json.Marshal(v)
	//t.Logf("%v",string(json))

	if v.Fields["def"].Value != "2" {
		t.Errorf("expect: 2 but %s", v.Fields["def"].Value)
	}

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

	v := Create(reflect.ValueOf(myValue))
	v = v.Fields["0"]
	if len(v.Fields) != 3 {
		t.Errorf("expect: 3, but: %d", len(v.Fields))
	}
	if v.Type != "reflect2json.HogeHoge" {
		t.Errorf("expect: struct but %s", v.Type)
	}
	if v.Fields["Hell"].Value != "hell?" {
		t.Errorf("expect: hell? but %s", v.Fields["Hell"].Value)
	}
}

func TestPtrValue(t *testing.T) {
	myValue := &HogeHoge {
		"hell?",
		12,
		"c",
	}

	v := Create(reflect.ValueOf(myValue))
	v = v.Fields["0"]
	if len(v.Fields) != 3 {
		t.Errorf("expect: 3, but: %d", len(v.Fields))
	}
	if v.Type != "reflect2json.HogeHoge" {
		t.Errorf("expect: struct but %s", v.Type)
	}
	if v.Fields["Hell"].Value != "hell?" {
		t.Errorf("expect: hell? but %s", v.Fields["Hell"].Value)
	}
}



func TestStructValue(t *testing.T) {
	myValue := HogeHoge {
		"hell?",
		12,
		"c",
	}

	v := Create(reflect.ValueOf(myValue))
	//t.Logf("%#v", v)
	//json, _ := json.Marshal(v)
	//t.Logf("%v", string(json))
	if len(v.Fields) != 3 {
		t.Errorf("expect: 3, but: %d", len(v.Fields))
	}
	if v.Type != "reflect2json.HogeHoge" {
		t.Logf("%#v\n", v)
		t.Errorf("expect: reflect2json.HogeHoge but %s", v.Type)
	}
	if v.Fields["Hell"].Value != "hell?" {
		t.Errorf("expect: hell? but %s", v.Fields["Hell"].Fields)
	}
}


func TestArrayValue(t *testing.T) {
	myValue := [...]string{
		"a","b","c",
	}

	v := Create(reflect.ValueOf(myValue))
	if len(v.Fields) != 3 {
		t.Errorf("expect: 3, but: %d", len(v.Fields))
	}
	if v.Type != "[3]string" {
		t.Logf("%#v\n", v)
		t.Errorf("expect: array but %s", v.Type)
	}
	if v.Fields["0"].Value != "a" {
		t.Errorf("%#v", v.Fields)
		t.Errorf("expect: a but %s", v.Fields["0"].Value)

	}
}


func TestStringValue(t *testing.T) {
	myValue := "something"
	v := Create(reflect.ValueOf(myValue))
	expect :=`{0  string something map[]}`
	result := fmt.Sprintf("%v", v)
	if result != expect {
		t.Errorf ("expect: \n\t%v\n but: \n\t%s\n", expect, result)
	}
}


func TestIntValue(t *testing.T) {
	myValue := 2
	v := Create(reflect.ValueOf(myValue))
	expect :=`{0  int 2 map[]}`
	result := fmt.Sprintf("%v", v)
	if result != expect {
		t.Errorf ("expect: \n\t%v\n but: \n\t%s\n", expect, result)
	}
}

