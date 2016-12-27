package filteredError

import "strings"

type FilteredError struct {
	err        error
	ReplaceMap map[string]string
}

func NewFilteredError(err error, replacers map[string]string) FilteredError {
	return FilteredError{ReplaceMap: replacers, err: err}
}

func (fe FilteredError) AddReplacer(object, replacer string) {
	if fe.ReplaceMap == nil {
		fe.ReplaceMap = map[string]string{}
	}
	fe.ReplaceMap[object] = replacer
}

func (fe FilteredError) RemoveReplacer(object string) {
	if fe.ReplaceMap == nil {
		return
	}
	delete(fe.ReplaceMap, object)
}

func (fe FilteredError) Error() (returnMessage string) {
	if fe.err == nil {
		return ""
	}
	returnMessage = fe.err.Error()
	if fe.ReplaceMap == nil {
		return
	}
	for k, v := range fe.ReplaceMap {
		returnMessage = strings.Replace(returnMessage, k, v, -1)
	}

	return
}
