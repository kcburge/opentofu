// Code generated by "stringer -type=DependenciesModeType dependencies_mode_type.go"; DO NOT EDIT.

package tofu

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DependenciesModeTypeInvalid-0]
	_ = x[DependenciesModeTypeInclude-1]
	_ = x[DependenciesModeTypeFail-2]
	_ = x[DependenciesModeTypeExclude-3]
}

const _DependenciesModeType_name = "DependenciesModeTypeInvalidDependenciesModeTypeIncludeDependenciesModeTypeFailDependenciesModeTypeExclude"

var _DependenciesModeType_index = [...]uint8{0, 27, 54, 78, 105}

func (i DependenciesModeType) String() string {
	if i < 0 || i >= DependenciesModeType(len(_DependenciesModeType_index)-1) {
		return "DependenciesModeType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DependenciesModeType_name[_DependenciesModeType_index[i]:_DependenciesModeType_index[i+1]]
}
