package reflect2json

import (
	"reflect"
	"fmt"
)

type ReflectJSON struct {
	Order  int                    `json:"order"`
	Type   string                 `json:"type,omitempty"`
	Kind   string                 `json:"kind"`
	Value  string                 `json:"value,omitempty"`
	Fields map[string]ReflectJSON `json:"fields,omitempty"`
}

func makeFields(members map[string]ReflectJSON) map[string]ReflectJSON {
	if members == nil {
		return make(map[string]ReflectJSON)
	}
	return members
}

var ptrList map[reflect.Value]bool = make(map[reflect.Value]bool)

func panicRecover(rj *ReflectJSON)  {
	if err := recover(); err != nil {
		fmt.Errorf("\npanicRecover err=%s\n", err)
		rj.Kind = reflect.Invalid.String()
		switch err.(type) {
		case string:
			rj.Value = err.(string)
		case error:
			rj.Value = err.(error).Error()
		default:
			rj.Value = fmt.Sprintf("%v", err)

		}
	} else {
		// fmt.Printf("\npanicRecover throuth\n")
	}
}

func Create(value reflect.Value) (rj ReflectJSON) {

	defer panicRecover(&rj)

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
			rj.Fields = makeFields(rj.Fields)
			rj.Fields[fmt.Sprintf("%d", i)] = member
		}
	case reflect.Ptr:
		typeOfV := value.Type()
		rj.Type = typeOfV.String()
		rj.Fields = makeFields(rj.Fields)
		var member ReflectJSON
		if check, ok := ptrList[value]; ok && check {
			member = Create(reflect.ValueOf("cycle loop:" + fmt.Sprintf("%x", value.Pointer())))
		} else {
			ptrList[value] = true

			member = Create(reflect.Indirect(value))
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

func reflectStruct(v reflect.Value) (rj ReflectJSON) {
	defer panicRecover(&rj)

	rj.Fields = makeFields(rj.Fields)
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

				if check, ok := ptrList[field]; ok && check {
					field = reflect.ValueOf("cycle loop")
				} else {
					ptrList[field] = true
					field = reflect.Indirect(field)
				}
			}

			member = Create(field)
		default:
			member = Create(field)
		}
		member.Order = i

		rj.Fields[f.Name] = member
	}
	return rj
}

func reflectMap(v reflect.Value) (rj ReflectJSON) {
	defer panicRecover(&rj)


	t := v.Type()
	rj.Type = t.String()
	rj.Kind = v.Kind().String()
	rj.Fields = makeFields(rj.Fields)

	for i, key := range v.MapKeys() {
		var member ReflectJSON
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
		//	member = Create(field)
		default:
			member = Create(field)
		}

		member.Order = i
		rj.Fields[key.String()] = member
	}
	return rj
}
