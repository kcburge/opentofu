package tofu

//go:generate stringer -type=DependenciesModeType dependencies_mode_type.go

// DependenciesModeType is an enum
// Set to indicate how dependencies should be handled.
//   include: include all specified targets and implicit dependencies (standard behavior).
//   fail: fail if any implicit dependecies are targeted
//   exclude: only target explicit targets, ignoring dependencies

type DependenciesModeType int

const (
	DependenciesModeTypeInvalid DependenciesModeType = 0
	DependenciesModeTypeInclude DependenciesModeType = iota
	DependenciesModeTypeFail
	DependenciesModeTypeExclude
)

// DependenciesModeTypeMap is a mapping of human-readable string to
// DependenciesModeType. This is useful to use as the mechanism for
// human input for configurable graph types.
var DependenciesModeTypeMap = map[string]DependenciesModeType{
	"invalid": DependenciesModeTypeInvalid,
	"include": DependenciesModeTypeInclude,
	"fail":    DependenciesModeTypeFail,
	"exclude": DependenciesModeTypeExclude,
}
