package pkg

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/copier"
)

// ObjectFrom ...
func ObjectFrom(data []byte) (o *Object) {
	if len(data) == 0 {
		return &Object{Value: &Value{Kind: Value_EMPTY}}
	}

	var v interface{}
	if err := json.Unmarshal(data, &v); err == nil {
		o = objectFromInterface(v)
	} else {
		o = objectFromInterface(string(data))
	}

	return o
}

// CloneObject ...
func CloneObject(i *Object) *Object {
	var o Object
	// deep copy
	copier.Copy(&o, i)

	// Split data. Get more data today by doing calls to Print!
	if o.Value != nil {
		for i := range o.Value.List {
			if o.Value.List[i].Value != nil {
				for j := range o.Value.List[i].Value.List {
					o.Value.List[i].Value.List[j].Value = nil
				}
			}
		}
		for k := range o.Value.Hashmap {
			if o.Value.Hashmap[k].Value != nil {
				for l := range o.Value.Hashmap[k].Value.Hashmap {
					o.Value.Hashmap[k].Value.Hashmap[l].Value = nil
				}
			}
		}
	}
	return &o
}

func number(vv float64) *Object {
	return Put(&Object{
		Ptr: ptr(vv),
		Value: &Value{
			Kind:   Value_is_number,
			Number: vv,
		},
	})
}

func objectFromInterface(v interface{}) *Object {
	switch vv := v.(type) {
	case nil:
		return Put(&Object{
			Ptr:   ptr(vv),
			Value: &Value{Kind: Value_is_null},
		})

	case bool:
		return Put(&Object{
			Ptr: ptr(vv),
			Value: &Value{
				Kind:    Value_is_bool,
				Boolean: vv,
			},
		})

	case string:
		return Put(&Object{
			Ptr: ptr(vv),
			Value: &Value{
				Kind: Value_is_str,
				Str:  vv,
			},
		})

	case float32:
		return number(float64(vv))
	case float64:
		return number(float64(vv))
	case int:
		return number(float64(vv))
	case int8:
		return number(float64(vv))
	case int16:
		return number(float64(vv))
	case int32:
		return number(float64(vv))
	case int64:
		return number(float64(vv))
	case uint:
		return number(float64(vv))
	case uint8:
		return number(float64(vv))
	case uint16:
		return number(float64(vv))
	case uint32:
		return number(float64(vv))
	case uint64:
		return number(float64(vv))

	case []interface{}:
		xs := make([]*Object, 0, len(vv))
		for _, x := range vv {
			y := objectFromInterface(x)
			xs = append(xs, y)
		}
		return Put(&Object{
			Ptr: ptr(vv), // ptr(vv) = ptr([42 , [...]])
			Value: &Value{
				Kind: Value_is_list,
				List: xs,
			},
		})

	case map[string]interface{}:
		xs := make(map[string]*Object, len(vv))
		for k, x := range vv {
			y := objectFromInterface(x)
			xs[k] = y
		}
		return Put(&Object{
			Ptr: ptr(vv),
			Value: &Value{
				Kind:    Value_is_hashmap,
				Hashmap: xs,
			},
		})

	default:
		panic(fmt.Sprintf("unhandled %T %v", vv, vv))
	}
}
