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



func refrectStruct(v reflect.Value) Struct2json {

	fieldName := ""
	structName := ""
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("Recover! %s, FieldName:%s, StructName:%s \n" ,err, fieldName, structName)
		}
	}()
	st := Struct2json{
		0,
		"",
		"",
		nil,
	}

	kind := v.Kind()
	structName = v.Type().String()
	st.Type = kind.String()

	t := v.Type()
	fmt.Printf("in Struct\n %#v\n\n", v)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fieldName = f.Name
		field := v.Field(i)


		var member Struct2json

		if  v.Field(i).CanInterface() {
			member = Create(field.Interface())
		} else {
			member = Create(fmt.Sprintf("%s", field))
		}

		member.Order = i

		if st.StMembers == nil {
			st.StMembers = make(map[string]Struct2json)
		}
		st.StMembers[f.Name] = member
	}
	return st
}
func Create(obj interface{}) Struct2json {
	//defer func() {
	//	err := recover()
	//	if err != nil {
	//		fmt.Println("Recover!:", err)
	//	}
	//}()

	st := Struct2json{
		0,
		"",
		"",
		nil,
	}

	v := reflect.ValueOf(obj)
	kind := v.Kind()
	st.Type = kind.String()

	switch kind {
	case reflect.Struct:
		return refrectStruct(v)
	case reflect.Array, reflect.Slice:
		//fmt.Println("in Array")
		for i := 0; i < v.Len(); i++ {

			//fmt.Printf("%#v\n", Create(v.Index(i).Interface()))
			member := Create(v.Index(i).Interface())
			member.Order = i
			if st.StMembers == nil {
				st.StMembers = make(map[string]Struct2json)
			}
			st.StMembers[fmt.Sprintf("%d", i)] = member
		}
	case reflect.Ptr:
		if st.StMembers == nil {
			st.StMembers = make(map[string]Struct2json)
		}

		v = reflect.Indirect(v)
		if v.CanInterface() {
			fmt.Printf("%#v\n", v.Interface())
			member := Create(v.Interface())
			st.StMembers["0"] = member
		} else {
			fmt.Printf("invalid: %#v\n", v)
		}

	default:
		//fmt.Println("in default")
		st.PrimValue = fmt.Sprintf("%v", v)
	}

	return st
}
