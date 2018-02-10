package struct2json

import (
	"reflect"

	"fmt"
)

/*

{
	order: 0
	type: Struct
    members: [
		{
			order: 0
			name: "hogehoge"
			type: string
			value: "hogeee"
		},
		{
			order: 1
			name: "hogeStruct"
			type: struct
			value: { type: "struct" ....}

		}
	]
}


 */
type Struct2json struct {
	Order     int                    `json:"order"`
	Type      string                 `json:"type"`
	PrimValue string                 `json:"value,omitempty"`
	StMembers map[string]Struct2json `json:"members,omitempty"`
}



func reflectStruct(v reflect.Value) (st Struct2json) {
	fmt.Printf("\nin reflectStruct: %s, %s, %s\n", v.Kind().String(),v.Type().String(),v.String())

	structName := ""
	var currentField reflect.Value
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("in reflectStruct Recover! err=%s\n, StructName:%s ,FieldName:%s, \n v= %v" ,
				err, structName, currentField.Kind().String(), v)
			return
		}
	}()

	kind := v.Kind()
	structName = v.Type().String()
	st.Type = kind.String()

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		field := v.Field(i)
		currentField = field

		var member Struct2json

		switch field.Kind() {
		case reflect.String:
			member = Create(field.Interface())
		case reflect.Ptr:
			member = Create(reflect.Indirect(field).Interface())
		default:
			if field.CanInterface() {
				member = Create(field.Interface())
			} else {
				member = Create(fmt.Sprintf("%s", "unknown"))
			}
		}
		member.Order = i

		if st.StMembers == nil {
			st.StMembers = make(map[string]Struct2json)
		}
		st.StMembers[f.Name] = member
	}
	return st
}
func Create(obj interface{}) (st Struct2json) {
	v := reflect.ValueOf(obj)
	kind := v.Kind()
	fmt.Printf("\nin Create: %s, %s, %s\n", v.Kind().String(),v.Type().String(),v.String())

	//defer func() {
	//	err := recover()
	//	if err != nil {
	//		fmt.Println("Create Recover!:", err)
	//		fmt.Printf("in Recover!: %#v\n\n", kind.String())
	//	}
	//	return
	//}()


	st.Type = kind.String()


	switch kind {
	case reflect.Array, reflect.Slice:
		//fmt.Println("in Array")
		for i := 0; i < v.Len(); i++ {

			member := Create(v.Index(i).Interface())
			member.Order = i
			if st.StMembers == nil {
				st.StMembers = make(map[string]Struct2json)
			}
			st.StMembers[fmt.Sprintf("%d", i)] = member
		}
	case reflect.Ptr:
		fmt.Printf("in ptr %s \n",v.Kind().String())
		if st.StMembers == nil {
			st.StMembers = make(map[string]Struct2json)
		}

		if v.CanInterface() {
			member := Create(reflect.Indirect(v).Interface())
			st.StMembers["0"] = member
		} else {
			fmt.Printf("invalid: %#v\n", v)
		}
	case reflect.Struct:
		return reflectStruct(v)
	default:
		st.PrimValue = fmt.Sprintf("%v", v)
	}

	return st
}
