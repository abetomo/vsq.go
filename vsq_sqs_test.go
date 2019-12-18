package vsq

import (
	"errors"
	"io/ioutil"
	"reflect"
	"regexp"
	"testing"
)

func mockUniqId() string {
	return "test-id"
}

func TestUniqId(t *testing.T) {
	actual := UniqId()
	if r := regexp.MustCompile(`\d{9}-[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`); !r.MatchString(actual) {
		t.Fatalf("%v not match %v", actual, r)
	}
}

func Test_loadLikeSQSSuccess(t *testing.T) {
	actual, err := loadLikeSQS("testdata/data_file_like_sqs.json")

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if expected := (VsqDataLikeSQS{"VerySimpleQueueLikeSQS", map[string]string{"id1": "1", "id2": "2", "id3": "3"}}); !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %#v\nwant %#v", actual, expected)
	}
}

func Test_loadLikeSQSSuccessFileNotExist(t *testing.T) {
	actual, err := loadLikeSQS("not_exist")

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if expected := (VsqDataLikeSQS{"VerySimpleQueueLikeSQS", map[string]string{}}); !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %#v\nwant %#v", actual, expected)
	}
}

func Test_loadLikeSQSFailedNameInvalid(t *testing.T) {
	actual, err := loadLikeSQS("testdata/name_invalid_file.json")

	if err == nil {
		t.Fatalf("Succeeded with failed test")
	}

	if expected := errors.New("not a data file of VerySimpleQueueLikeSQS"); !reflect.DeepEqual(err, expected) {
		t.Fatalf("got %#v\nwant %#v", err, expected)
	}

	if expected := (VsqDataLikeSQS{}); !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %#v\nwant %#v", actual, expected)
	}
}

func Test_loadLikeSQSFailedValueInvalid(t *testing.T) {
	actual, err := loadLikeSQS("testdata/value_invalid_file_like_sqs.json")

	if err == nil {
		t.Fatalf("Succeeded with failed test")
	}

	if expected := errors.New("not a data file of VerySimpleQueueLikeSQS"); !reflect.DeepEqual(err, expected) {
		t.Fatalf("got %#v\nwant %#v", err, expected)
	}

	if expected := (VsqDataLikeSQS{}); !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %#v\nwant %#v", actual, expected)
	}
}

func TestLoadSuccess_LikeSQS(t *testing.T) {
	var vsq VerySimpleQueueLikeSQS
	actual, err := vsq.Load("testdata/data_file_like_sqs.json")

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	dataExpected := (VsqDataLikeSQS{"VerySimpleQueueLikeSQS", map[string]string{"id1": "1", "id2": "2", "id3": "3"}})
	if !reflect.DeepEqual(actual, dataExpected) {
		t.Fatalf("got %#v\nwant %#v (return value)", actual, dataExpected)
	}
	if !reflect.DeepEqual(vsq.Data, dataExpected) {
		t.Fatalf("got %#v\nwant %#v (vsq.Data)", vsq.Data, dataExpected)
	}

	if expected := "testdata/data_file_like_sqs.json"; vsq.FilePath != expected {
		t.Fatalf("got %#v\nwant %#v (vsq.FilePath)", vsq.FilePath, expected)
	}
}

func TestSizeWithValue_LikeSQS(t *testing.T) {
	var vsq VerySimpleQueueLikeSQS
	if _, err := vsq.Load("testdata/data_file_like_sqs.json"); err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if expected := 3; vsq.Size() != expected {
		t.Fatalf("got %#v\nwant %#v", vsq.Size(), expected)
	}
}

func TestSizeNoValue_LikeSQS(t *testing.T) {
	var vsq VerySimpleQueueLikeSQS
	if _, err := vsq.Load("not_exist"); err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if expected := 0; vsq.Size() != expected {
		t.Fatalf("got %#v\nwant %#v", vsq.Size(), expected)
	}
}

func TestWriteDbFile_LikeSQS(t *testing.T) {
	removeTestFile()

	var vsq VerySimpleQueueLikeSQS
	if _, err := vsq.Load(testFilePath()); err != nil {
		t.Fatalf("failed test %#v", err)
	}
	vsq.writeDbFile()

	bytes, err := ioutil.ReadFile(testFilePath())
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if expected := []byte(`{"Name":"VerySimpleQueueLikeSQS","Value":{}}`); !reflect.DeepEqual(bytes, expected) {
		t.Fatalf("got %#v\nwant %#v", bytes, expected)
	}
}

func TestSend(t *testing.T) {
	removeTestFile()

	var vsq VerySimpleQueueLikeSQS
	if _, err := vsq.Load(testFilePath()); err != nil {
		t.Fatalf("failed test %#v", err)
	}

	id := vsq.Send("hoge", mockUniqId)

	if expected := "test-id"; id != expected {
		t.Fatalf("got %#v\nwant %#v", id, expected)
	}

	if expected := (VsqDataLikeSQS{"VerySimpleQueueLikeSQS", map[string]string{"test-id": "hoge"}}); !reflect.DeepEqual(vsq.Data, expected) {
		t.Fatalf("got %#v\nwant %#v", vsq.Data, expected)
	}

	bytes, err := ioutil.ReadFile(testFilePath())
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if expected := []byte(`{"Name":"VerySimpleQueueLikeSQS","Value":{"test-id":"hoge"}}`); !reflect.DeepEqual(bytes, expected) {
		t.Fatalf("got %#v\nwant %#v", bytes, expected)
	}
}

func TestKeys(t *testing.T) {
	removeTestFile()

	var vsq VerySimpleQueueLikeSQS
	var keys []string
	if _, err := vsq.Load(testFilePath()); err != nil {
		t.Fatalf("failed test %#v", err)
	}

	keys = vsq.keys()
	if expected := []string{}; !reflect.DeepEqual(keys, expected) {
		t.Fatalf("got %#v\nwant %#v", keys, expected)
	}

	vsq.Send("hoge", mockUniqId)
	keys = vsq.keys()
	if expected := []string{"test-id"}; !reflect.DeepEqual(keys, expected) {
		t.Fatalf("got %#v\nwant %#v", keys, expected)
	}
}

func TestReceiveSuccess(t *testing.T) {
	removeTestFile()

	var vsq VerySimpleQueueLikeSQS
	if _, err := vsq.Load(testFilePath()); err != nil {
		t.Fatalf("failed test %#v", err)
	}
	vsq.Send("hoge", mockUniqId)

	value, err := vsq.Receive()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if expected := (VsqDataLikeSQSValue{"test-id", "hoge"}); !reflect.DeepEqual(value, expected) {
		t.Fatalf("got %#v\nwant %#v", value, expected)
	}
}

func TestReceiveFailed(t *testing.T) {
	removeTestFile()

	var vsq VerySimpleQueueLikeSQS
	if _, err := vsq.Load(testFilePath()); err != nil {
		t.Fatalf("failed test %#v", err)
	}

	value, err := vsq.Receive()
	if err == nil {
		t.Fatalf("Succeeded with failed test")
	}

	if expected := (VsqDataLikeSQSValue{}); !reflect.DeepEqual(value, expected) {
		t.Fatalf("got %#v\nwant %#v", value, expected)
	}
}

func TestDeleteTrue(t *testing.T) {
	removeTestFile()

	var vsq VerySimpleQueueLikeSQS
	if _, err := vsq.Load(testFilePath()); err != nil {
		t.Fatalf("failed test %#v", err)
	}

	vsq.Send("hoge", mockUniqId)
	if expected := 1; vsq.Size() != expected {
		t.Fatalf("got %#v\nwant %#v", vsq.Size(), expected)
	}

	ret := vsq.Delete("test-id")
	if expected := true; ret != expected {
		t.Fatalf("got %#v\nwant %#v", ret, expected)
	}

	if expected := 0; vsq.Size() != expected {
		t.Fatalf("got %#v\nwant %#v", vsq.Size(), expected)
	}

	bytes, err := ioutil.ReadFile(testFilePath())
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if expected := []byte(`{"Name":"VerySimpleQueueLikeSQS","Value":{}}`); !reflect.DeepEqual(bytes, expected) {
		t.Fatalf("got %#v\nwant %#v", bytes, expected)
	}
}

func TestDeleteFalse(t *testing.T) {
	removeTestFile()

	var vsq VerySimpleQueueLikeSQS
	if _, err := vsq.Load(testFilePath()); err != nil {
		t.Fatalf("failed test %#v", err)
	}

	ret := vsq.Delete("hoge")

	if expected := false; ret != expected {
		t.Fatalf("got %#v\nwant %#v", ret, expected)
	}
}
