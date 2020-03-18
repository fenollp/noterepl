package pkg

import (
	"encoding/json"
	"fmt"
)

// ObjectFrom ...
func ObjectFrom(data []byte) *Object {
	if len(data) == 0 {
		return &Object{Value: &Value{Kind: Value_EMPTY}}
	}

	var v interface{}
	if err := json.Unmarshal(data, &v); err == nil {
		return objectFromInterface(v)
	}

	return objectFromInterface(string(data))
}

func number(vv float64) *Object {
	return &Object{
		Ptr: ptr(vv),
		Value: &Value{
			Kind:   Value_is_number,
			Number: vv,
		},
	}
}

func objectFromInterface(v interface{}) *Object {
	switch vv := v.(type) {
	case nil:
		return &Object{
			Ptr:   ptr(vv),
			Value: &Value{Kind: Value_is_null},
		}

	case bool:
		return &Object{
			Ptr: ptr(vv),
			Value: &Value{
				Kind:    Value_is_bool,
				Boolean: vv,
			},
		}

	case string:
		return &Object{
			Ptr: ptr(vv),
			Value: &Value{
				Kind: Value_is_str,
				Str:  vv,
			},
		}

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
		return &Object{
			Ptr: ptr(vv),
			Value: &Value{
				Kind: Value_is_list,
				List: xs,
			},
		}

	default:
		panic(fmt.Sprintf("unhandled %T %v", vv, vv))
	}
}
