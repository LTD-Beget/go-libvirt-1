// Code generated by "stringer -type=TypedParamTypes -output=consts_string.go"; DO NOT EDIT.

package libvirt

import "fmt"

const _TypedParamTypes_name = "TypedParamTypeINTTypedParamTypeUINTTypedParamTypeLLONGTypedParamTypeULLONGTypedParamTypeDOUBLETypedParamTypeBOOLEANTypedParamTypeSTRINGTypedParamTypeLAST"

var _TypedParamTypes_index = [...]uint8{0, 17, 35, 54, 74, 94, 115, 135, 153}

func (i TypedParamTypes) String() string {
	i -= 1
	if i >= TypedParamTypes(len(_TypedParamTypes_index)-1) {
		return fmt.Sprintf("TypedParamTypes(%d)", i+1)
	}
	return _TypedParamTypes_name[_TypedParamTypes_index[i]:_TypedParamTypes_index[i+1]]
}
