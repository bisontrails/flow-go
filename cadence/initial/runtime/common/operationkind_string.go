// Code generated by "stringer -type=OperationKind"; DO NOT EDIT.

package common

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OperationKindUnknown-0]
	_ = x[OperationKindUnary-1]
	_ = x[OperationKindBinary-2]
	_ = x[OperationKindTernary-3]
}

const _OperationKind_name = "OperationKindUnknownOperationKindUnaryOperationKindBinaryOperationKindTernary"

var _OperationKind_index = [...]uint8{0, 20, 38, 57, 77}

func (i OperationKind) String() string {
	if i >= OperationKind(len(_OperationKind_index)-1) {
		return "OperationKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OperationKind_name[_OperationKind_index[i]:_OperationKind_index[i+1]]
}