package pkg

// ObjectFrom ...
func ObjectFrom(str string) *Object {
	if str == "" {
		return &Object{Value: &Value{Kind: Value_EMPTY}}
	}

	return &Object{
		Value: &Value{
			Kind: Value_is_str,
			Str:  str,
		},
	}
}
