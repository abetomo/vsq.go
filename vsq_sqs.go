// WIP
package vsq

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type VsqDataLikeSQS struct {
	Name  string
	Value map[string]string
}

type VerySimpleQueueLikeSQS struct {
	Data     VsqDataLikeSQS
	FilePath string
}

func loadLikeSQS(filePath string) (VsqDataLikeSQS, error) {
	if _, err := os.Stat(filePath); err != nil {
		defaultVsqData := VsqDataLikeSQS{"VerySimpleQueueLikeSQS", map[string]string{}}
		return defaultVsqData, nil
	}

	var vsqData VsqDataLikeSQS
	bytes, _ := ioutil.ReadFile(filePath)
	if err := json.Unmarshal(bytes, &vsqData); err != nil {
		return VsqDataLikeSQS{}, errors.New("not a data file of VerySimpleQueueLikeSQS")
	}

	if vsqData.Name != "VerySimpleQueueLikeSQS" {
		return VsqDataLikeSQS{}, errors.New("not a data file of VerySimpleQueueLikeSQS")
	}
	return vsqData, nil
}

func (vsq *VerySimpleQueueLikeSQS) load(filePath string) (VsqDataLikeSQS, error) {
	var err error
	vsq.Data, err = loadLikeSQS(filePath)
	if err != nil {
		return VsqDataLikeSQS{}, err
	}

	vsq.FilePath = filePath
	return vsq.Data, nil
}

func (vsq VerySimpleQueueLikeSQS) size() int {
	return len(vsq.Data.Value)
}

func (vsq VerySimpleQueueLikeSQS) writeDbFile() {
	bytes, _ := json.Marshal(vsq.Data)
	ioutil.WriteFile(vsq.FilePath, bytes, 0644)
}
