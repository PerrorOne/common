package utils

func If(ok bool, trueValue, falseValue interface{}) interface{} {
	if ok {
		return trueValue
	}
	return falseValue
}
