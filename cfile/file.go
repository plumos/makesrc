package cfile
import (
	"io/ioutil"
)

func WriteToFile(content string, path string) {

	err := ioutil.WriteFile(path, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}
