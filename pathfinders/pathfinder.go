package pathfinders

type pathfinder interface {
	Eval(data map[string]interface{}, path string) interface{}
}

type Pathfinder struct {
	Path string
	pathfinder
}

func (pf *Pathfinder) Eval(data map[string]interface{}, path string) interface{} {

	if result, err := EvalPointer(data, path); err == nil {
		return result
	}

	if ar := ConformsArraySelector(pf.Path); ar != nil {
		return ar.Eval(data, path)
	}

	return ""
}
