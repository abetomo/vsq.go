package vsq

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"os"
	"sort"
	"time"
)

type VsqDataLikeSQS struct {
	Name  string
	Value map[string]string
}

type VsqDataLikeSQSValue struct {
	Id   string
	Body string
}

type VerySimpleQueueLikeSQS struct {
	Data     VsqDataLikeSQS
	FilePath string
}

func UniqId() string {
	timestamp := time.Now().Unix()
	return fmt.Sprintf("%d-%s", timestamp/10, uuid.Must(uuid.NewV4(), nil))
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

func (vsq *VerySimpleQueueLikeSQS) Load(filePath string) (VsqDataLikeSQS, error) {
	var err error
	vsq.Data, err = loadLikeSQS(filePath)
	if err != nil {
		return VsqDataLikeSQS{}, err
	}

	vsq.FilePath = filePath
	return vsq.Data, nil
}

func (vsq VerySimpleQueueLikeSQS) Size() int {
	return len(vsq.Data.Value)
}

func (vsq VerySimpleQueueLikeSQS) writeDbFile() {
	bytes, _ := json.Marshal(vsq.Data)
	ioutil.WriteFile(vsq.FilePath, bytes, 0644)
}

func (vsq *VerySimpleQueueLikeSQS) Send(data string, idFunc func() string) string {
	id := idFunc()
	if vsq.Data.Value == nil {
		vsq.Data.Value = map[string]string{}
	}
	vsq.Data.Value[id] = data
	defer vsq.writeDbFile()
	return id
}

func (vsq VerySimpleQueueLikeSQS) keys() []string {
	ks := []string{}
	for k, _ := range vsq.Data.Value {
		ks = append(ks, k)
	}
	return ks
}

func (vsq VerySimpleQueueLikeSQS) Receive() (VsqDataLikeSQSValue, error) {
	if vsq.Size() == 0 {
		return VsqDataLikeSQSValue{}, errors.New("size is zero")
	}
	keys := vsq.keys()
	sort.Strings(keys)
	id := keys[0]
	return VsqDataLikeSQSValue{id, vsq.Data.Value[id]}, nil
}

func (vsq *VerySimpleQueueLikeSQS) Delete(id string) bool {
	if _, ok := vsq.Data.Value[id]; ok == false {
		return false
	}
	delete(vsq.Data.Value, id)
	defer vsq.writeDbFile()
	return true
}
