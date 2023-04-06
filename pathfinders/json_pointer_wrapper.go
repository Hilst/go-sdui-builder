package pathfinders

import (
	"github.com/qri-io/jsonpointer"
)

func EvalPointer(data interface{}, path string) (result interface{}, err error) {
	if ptr, err := jsonpointer.Parse(path); err == nil {
		return ptr.Eval(data)
	}
	return path, err
}
