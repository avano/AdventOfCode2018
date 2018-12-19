package util

import (
	"io/ioutil"
	"path"
	"runtime"
	"strings"
)

// GetInputString loads the input as string
func GetInputString() string {
	_, file, _, _ := runtime.Caller(1)

	day := file[strings.Index(file, ".go")-3 : strings.Index(file, ".go")-1]

	data, err := ioutil.ReadFile(path.Join(path.Dir(file), "../../assets/day"+day))
	if err != nil {
		data, err2 := ioutil.ReadFile(path.Join(path.Dir(file), "../../assets/day"+day+"_example"))
		if err2 != nil {
			panic(err2)
		}
		return strings.TrimSpace(string(data))
	}

	return strings.TrimSpace(string(data))
}
