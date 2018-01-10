// WIP
package vsq

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type VsqData struct {
	Name  string
	Value []string
}

type VerySimpleQueue struct {
	Data     VsqData
	FilePath string
}

func load(filePath string) (VsqData, error) {
	if _, err := os.Stat(filePath); err != nil {
		defaultVsqData := VsqData{"VerySimpleQueue", []string{}}
		return defaultVsqData, nil
	}

	var vsqData VsqData
	bytes, _ := ioutil.ReadFile(filePath)
	if err := json.Unmarshal(bytes, &vsqData); err != nil {
		return VsqData{}, errors.New("not a data file of VerySimpleQueue")
	}

	if vsqData.Name != "VerySimpleQueue" {
		return VsqData{}, errors.New("not a data file of VerySimpleQueue")
	}
	return vsqData, nil
}

// TODO: Other functions
