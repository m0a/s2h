package reflect2json

import (
	"reflect"
	"fmt"
	"encoding/json"
	"os"
)

type reflectJSON struct {
	Order  int                    `json:"order"`
	Type   string                 `json:"type,omitempty"`
	Kind   string                 `json:"kind"`
	Value  string                 `json:"value,omitempty"`
	Fields map[string]reflectJSON `json:"fields,omitempty"`
}

// main routines
func Reflect2JSON(value interface{}) string {
	rj := walk(reflect.ValueOf(value))
	bytes, err := json.Marshal(rj)
	if err != nil {
		return  ""
	}
	return  string(bytes)
}


func makeFields(members map[string]reflectJSON) map[string]reflectJSON {
	if members == nil {
		return make(map[string]reflectJSON)
	}
	return members
}

var ptrList map[reflect.Value]bool = make(map[reflect.Value]bool)

func panicRecover(rj *reflectJSON)  {
	if err := recover(); err != nil {
		fmt.Fprintf(os.Stderr,"\npanicRecover err=%s\n", err)
		rj.Kind = reflect.Invalid.String()
		switch err.(type) {
		case string:
			rj.Value = err.(string)
		case error:
			rj.Value = err.(error).Error()
		default:
			rj.Value = fmt.Sprintf("%v", err)
		}
	}
}



func walk(value reflect.Value) (rj reflectJSON) {

	defer panicRecover(&rj)

	kind := value.Kind()
	rj.Kind = kind.String()
	rj.Fields = makeFields(rj.Fields)

	if kind == reflect.Interface {
		value = reflect.ValueOf(value.Interface())
		kind = value.Kind()
	}

	switch kind {
	case reflect.Array, reflect.Slice:
		typeOfV := value.Type()
		rj.Type = typeOfV.String()
		for i := 0; i < value.Len(); i++ {
			member := walk(value.Index(i))
			member.Order = i
			rj.Fields[fmt.Sprintf("%d", i)] = member
		}
	case reflect.Ptr:
		typeOfV := value.Type()
		rj.Type = typeOfV.String()
		var member reflectJSON
		if check, ok := ptrList[value]; ok && check {
			member = walk(reflect.ValueOf("cycle loop:" + fmt.Sprintf("%x", value.Pointer())))
		} else {
			ptrList[value] = true

			member = walk(reflect.Indirect(value))
			rj.Value = fmt.Sprintf("%x", value.Pointer())
		}

		rj.Fields["0"] = member
	case reflect.Struct:
		return reflectStruct(value)
	case reflect.Map:
		return reflectMap(value)
	case reflect.String:
		rj.Value = value.String()
	case reflect.Invalid:
		rj.Type = "nil"
		rj.Value = "nil"
	case reflect.Interface:
		panic(value)
	default:
		rj.Value = fmt.Sprintf("%v", value)
	}

	return rj
}

func reflectStruct(v reflect.Value) (rj reflectJSON) {
	defer panicRecover(&rj)

	rj.Fields = makeFields(rj.Fields)
	t := v.Type()
	rj.Type = t.String()
	rj.Kind = v.Kind().String()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		field := v.Field(i)
		var member reflectJSON

		switch field.Kind() {
		case reflect.Ptr:
			if !field.IsNil() {

				if check, ok := ptrList[field]; ok && check {
					field = reflect.ValueOf("cycle loop")
				} else {
					ptrList[field] = true
					field = reflect.Indirect(field)
				}
			}

			member = walk(field)
		default:
			member = walk(field)
		}
		member.Order = i

		rj.Fields[f.Name] = member
	}
	return rj
}

func reflectMap(v reflect.Value) (rj reflectJSON) {
	defer panicRecover(&rj)
	rj.Fields = makeFields(rj.Fields)


	t := v.Type()
	rj.Type = t.String()
	rj.Kind = v.Kind().String()

	for i, key := range v.MapKeys() {
		var member reflectJSON
		field := v.MapIndex(key)
		if field.Kind() == reflect.Interface {
			field = reflect.ValueOf(field.Interface())
		}

		switch field.Kind() {
		//case reflect.Ptr:
		//	if !field.IsNil() {
		//		field = reflect.Indirect(field)
		//	}
		//
		//	member = walk(field)
		default:
			member = walk(field)
		}

		member.Order = i
		rj.Fields[key.String()] = member
	}
	return rj
}
