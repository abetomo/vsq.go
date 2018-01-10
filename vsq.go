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

func (vsq *VerySimpleQueue) load(filePath string) (VsqData, error) {
	var err error
	vsq.Data, err = load(filePath)
	if err != nil {
		return VsqData{}, err
	}

	vsq.FilePath = filePath
	return vsq.Data, nil
}

func (vsq VerySimpleQueue) size() int {
	return len(vsq.Data.Value)
}

func (vsq VerySimpleQueue) writeDbFile() {
	bytes, _ := json.Marshal(vsq.Data)
	ioutil.WriteFile(vsq.FilePath, bytes, 0644)
}

func (vsq *VerySimpleQueue) shift() (string, error) {
	if vsq.size() == 0 {
		return "", errors.New("size is zero")
	}
	value := vsq.Data.Value[0]
	vsq.Data.Value = vsq.Data.Value[1:]
	vsq.writeDbFile()
	return value, nil
}

func (vsq *VerySimpleQueue) unshift(data string) int {
	value := &vsq.Data.Value
	// https://github.com/golang/go/wiki/SliceTricks#unshift
	(*value) = append([]string{data}, (*value)...)
	vsq.writeDbFile()
	return vsq.size()
}

// TODO: Other functions
