package pathfinders

import (
	"fmt"
	"regexp"
	"strings"
)

const ORIGIN = "origin"
const ARRAY = "array"
const SELECTOR = "selector"
const MATCHER = "matcher"

type ArraySlector struct {
	Path  string
	paths map[string]string
}

func ConformsArraySelector(path string) *ArraySlector {
	as := &ArraySlector{
		Path:  path,
		paths: make(map[string]string),
	}
	regex := regexp.MustCompile(`(?P<origin>.*?)(?P<array>\[(?P<selector>.*?)==(?P<matcher>.*?)\])`)
	match := regex.FindStringSubmatch(as.Path)
	for i, name := range regex.SubexpNames() {
		if i > 0 && i < len(match) {
			as.paths[name] = strings.Trim(match[i], " ")
		}
	}

	if len(as.paths) == 4 {
		return as
	}

	return nil
}

func (as *ArraySlector) Eval(data map[string]interface{}, path string) interface{} {

	ptrResult, err := EvalPointer(data, as.paths[ORIGIN])
	if err != nil {
		return ""
	}
	array := ptrResult.([]interface{})
	for i, d := range array {
		array[i] = d.(map[string]interface{})
	}

	var index string = ""
	for i, b := range array {
		subject, _ := EvalPointer(b, as.paths[SELECTOR])

		if subject.(string) == as.paths[MATCHER] {
			index = fmt.Sprint(i)
			break
		}
	}

	if index == "" {
		return index
	}

	as.Path = strings.Replace(as.Path, as.paths[ARRAY], "/"+index, 1)

	if result, err := EvalPointer(data, as.Path); err == nil {
		return result
	}

	return ""
}
