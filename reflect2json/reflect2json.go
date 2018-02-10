package reflect2json

import (
	"reflect"
	"fmt"
)

type ReflectJSON struct {
	Order   int                    `json:"order"`
	Type    string                 `json:"type,omitempty"`
	Kind    string                 `json:"kind"`
	Value   string                 `json:"value,omitempty"`
	Members map[string]ReflectJSON `json:"members,omitempty"`
}

func makeMembers(members map[string]ReflectJSON) map[string]ReflectJSON {
	if members == nil {
		return make(map[string]ReflectJSON)
	}
	return members
}

func Create(value reflect.Value) (rj ReflectJSON) {
	defer func(){
		err := recover()
		if err != nil {
			fmt.Printf("in Create  err=%s\n, value=%v",err, value)
		}
	}()
	kind := value.Kind()
	rj.Kind = kind.String()

	if kind == reflect.Interface {
		value = reflect.ValueOf(value.Interface())
		kind = value.Kind()
	}

	switch kind {
	case reflect.Array, reflect.Slice:
		typeOfV := value.Type()
		rj.Type = typeOfV.String()
		for i := 0; i < value.Len(); i++ {
			member := Create(value.Index(i))
			member.Order = i
			rj.Members = makeMembers(rj.Members)
			rj.Members[fmt.Sprintf("%d", i)] = member
		}
	case reflect.Ptr:
		typeOfV := value.Type()
		rj.Type = typeOfV.String()
		rj.Members = makeMembers(rj.Members)
		member := Create(reflect.Indirect(value))
		rj.Members["0"] = member
	case reflect.Struct:
		return reflectStruct(value)
	case reflect.Map:
		return reflectMap(value)
	case reflect.String:
		rj.Value = value.String()
	case reflect.Bool:
		rj.Value = fmt.Sprintf("%b",value.Bool())
	case reflect.Invalid:
		rj.Type = "nil"
		rj.Value = "nil"
	case reflect.Interface:
		panic(value)
	default:
		fmt.Printf("default: kind = %s\n", kind.String())
		rj.Value = fmt.Sprintf("%v", value)
	}

	return rj
}

func reflectStruct(v reflect.Value) (rj ReflectJSON) {
	defer func(){
		err := recover()
		if err != nil {
			fmt.Printf("in reflectStruct  err=%s\n, value=%v",err, v)
		}
	}()
	rj.Members = makeMembers(rj.Members)
	t := v.Type()
	rj.Type = t.String()
	rj.Kind = v.Kind().String()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		field := v.Field(i)
		var member ReflectJSON

		switch field.Kind() {
		case reflect.Ptr:
			if !field.IsNil() {
				field = reflect.Indirect(field)
			}

			member = Create(field)
		default:
			member = Create(field)
		}
		member.Order = i

		rj.Members[f.Name] = member
	}
	return rj
}

func reflectMap(v reflect.Value) (rj ReflectJSON) {
	defer func(){
		err := recover()
		if err != nil {
			fmt.Printf("in reflectMap  err=%s\n, value=%v",err, v)
		}
	}()

	t := v.Type()
	rj.Type = t.String()
	rj.Kind = v.Kind().String()
	rj.Members = makeMembers(rj.Members)

	for i,key := range v.MapKeys() {
		var member ReflectJSON
		//field := v.MapIndex(key)
		field := reflect.ValueOf(v.MapIndex(key).Interface())

		switch field.Kind() {
		//case reflect.Ptr:
		//	if !field.IsNil() {
		//		field = reflect.Indirect(field)
		//	}
		//
		//	member = Create(field)
		default:
			member = Create(field)
		}

		member.Order = i
		rj.Members[key.String()] = member
	}
	return rj
}