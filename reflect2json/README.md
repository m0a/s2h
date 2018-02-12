# reflect2json

Go's Reflection Info to json

type ReflectJSON struct {
Order int `json:"order"` // sort order
Type string `json:"type,omitempty"` // reflect.Type.String()
Kind string `json:"kind"` // reflect.Type.Kind().String()
Value string `json:"value,omitempty"` // reflect.Value.String()
Fields map[string]ReflectJSON `json:"fields,omitempty"` // reflect.Value.Fields
}

TypeScript interface example

```
export type GoReflectKind = "ptr" | "struct" | "string" | string; // TODO: all kind rewrite!

export default interface GoReflect {
  kind: GoReflectKind;
  type?: string;
  value?: string;
  order: number;
  fields: { [key: string]: GoReflect };
};
```
