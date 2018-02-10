package reflect2json

import (
	"testing"
	"fmt"
	"reflect"
)


func TestHowtoAccess(t *testing.T) {

	type Some struct{
		some string
		A string
		b int
		c bool
	}

	S := Some{
		A:"someA",
		b:12,
		c: true,
		some: "something",
	}

	valueOfS := reflect.ValueOf(S)
	typeOfS := reflect.TypeOf(S)

	fmt.Printf("Type:%s,Kind:%s\n",typeOfS.Name() , typeOfS.Kind().String())
	for i:=0; i< valueOfS.NumField(); i++ {
		fieldName := typeOfS.Field(i).Name
		value := valueOfS.Field(i).Interface()
		fmt.Printf("%s: %s\n",fieldName, value)
	}

}


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
	v = v.Members["0"]

	if len(v.Members) != 0 {
		t.Logf("%#v\n", v)
		t.Errorf("expect: v.stMembers not nil, but: %v", v.Members)
	}
	//if v.Type != "struct" {
	//	t.Errorf("expect: struct but %s", v.Type)
	//}
	//if v.StMembers["Hell"].PrimValue != "hell?" {
	//	t.Errorf("expect: hell? but %s", v.StMembers["Hell"].PrimValue)
	//}
}


type HogeHoge struct {
	Hell  string
	World int
	A     interface{}
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

	if v.Members["def"].Value != "2" {
		t.Errorf("expect: 2 but %s", v.Members["def"].Value)
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
	v = v.Members["0"]
	if len(v.Members) != 3 {
		t.Errorf("expect: 3, but: %d", len(v.Members))
	}
	if v.Type != "reflect2json.HogeHoge" {
		t.Errorf("expect: struct but %s", v.Type)
	}
	if v.Members["Hell"].Value != "hell?" {
		t.Errorf("expect: hell? but %s", v.Members["Hell"].Value)
	}
}

func TestPtrValue(t *testing.T) {
	myValue := &HogeHoge {
		"hell?",
		12,
		"c",
	}

	v := Create(reflect.ValueOf(myValue))
	v = v.Members["0"]
	if len(v.Members) != 3 {
		t.Errorf("expect: 3, but: %d", len(v.Members))
	}
	if v.Type != "reflect2json.HogeHoge" {
		t.Errorf("expect: struct but %s", v.Type)
	}
	if v.Members["Hell"].Value != "hell?" {
		t.Errorf("expect: hell? but %s", v.Members["Hell"].Value)
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
	if len(v.Members) != 3 {
		t.Errorf("expect: 3, but: %d", len(v.Members))
	}
	if v.Type != "reflect2json.HogeHoge" {
		t.Logf("%#v\n", v)
		t.Errorf("expect: reflect2json.HogeHoge but %s", v.Type)
	}
	if v.Members["Hell"].Value != "hell?" {
		t.Errorf("expect: hell? but %s", v.Members["Hell"].Members)
	}
}


func TestArrayValue(t *testing.T) {
	myValue := [...]string{
		"a","b","c",
	}

	v := Create(reflect.ValueOf(myValue))
	if len(v.Members) != 3 {
		t.Errorf("expect: 3, but: %d", len(v.Members))
	}
	if v.Type != "[3]string" {
		t.Logf("%#v\n", v)
		t.Errorf("expect: array but %s", v.Type)
	}
	if v.Members["0"].Value != "a" {
		t.Errorf("%#v", v.Members)
		t.Errorf("expect: a but %s", v.Members["0"].Value)

	}
}


func TestStringValue(t *testing.T) {
	myValue := "something"
	v := Create(reflect.ValueOf(myValue))
	expect :=`reflect2json.ReflectJSON{Order:0, Type:"string", Value:"something", Members:map[string]reflect2json.ReflectJSON(nil)}`
	result := fmt.Sprintf("%#v", v)
	if result != expect {
		t.Errorf ("expect: \n\t%v\n but: \n\t%s\n", expect, result)
	}
}


func TestIntValue(t *testing.T) {
	myValue := 2
	v := Create(reflect.ValueOf(myValue))
	expect :=`reflect2json.ReflectJSON{Order:0, Type:"int", Value:"2", Members:map[string]reflect2json.ReflectJSON(nil)}`
	result := fmt.Sprintf("%#v", v)
	if result != expect {
		t.Errorf ("expect: \n\t%v\n but: \n\t%s\n", expect, result)
	}
}

