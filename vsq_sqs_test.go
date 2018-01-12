// WIP
package vsq

import (
	"errors"
	"io/ioutil"
	"reflect"
	"testing"
)

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
	actual, err := vsq.load("testdata/data_file_like_sqs.json")

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
	if _, err := vsq.load("testdata/data_file_like_sqs.json"); err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if expected := 3; vsq.size() != 3 {
		t.Fatalf("got %#v\nwant %#v", vsq.size(), expected)
	}
}

func TestSizeNoValue_LikeSQS(t *testing.T) {
	var vsq VerySimpleQueueLikeSQS
	if _, err := vsq.load("not_exist"); err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if expected := 0; vsq.size() != 0 {
		t.Fatalf("got %#v\nwant %#v", vsq.size(), expected)
	}
}

func TestWriteDbFile_LikeSQS(t *testing.T) {
	removeTestFile()

	var vsq VerySimpleQueueLikeSQS
	if _, err := vsq.load(testFilePath); err != nil {
		t.Fatalf("failed test %#v", err)
	}
	vsq.writeDbFile()

	bytes, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if expected := []byte(`{"Name":"VerySimpleQueueLikeSQS","Value":{}}`); !reflect.DeepEqual(bytes, expected) {
		t.Fatalf("got %#v\nwant %#v", bytes, expected)
	}
}
