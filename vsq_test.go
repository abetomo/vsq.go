// WIP
package vsq

import (
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

const testFilePath = "/tmp/vsq_test.json"

func removeTestFile() {
	if _, err := os.Stat(testFilePath); err == nil {
		os.Remove(testFilePath)
	}
}

func TestMain(m *testing.M) {
	exit := m.Run()
	removeTestFile()
	os.Exit(exit)
}

func Test_loadSuccess(t *testing.T) {
	actual, err := load("testdata/data_file.json")

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if expected := (VsqData{"VerySimpleQueue", []string{"1", "2", "3"}}); !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %#v\nwant %#v", actual, expected)
	}
}

func Test_loadSucessFileNotExist(t *testing.T) {
	actual, err := load("not_exist")

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if expected := (VsqData{"VerySimpleQueue", []string{}}); !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %#v\nwant %#v", actual, expected)
	}
}

func Test_loadFailedNameInvalid(t *testing.T) {
	actual, err := load("testdata/name_invalid_file.json")

	if err == nil {
		t.Fatalf("Succeeded with failed test")
	}

	if expected := errors.New("not a data file of VerySimpleQueue"); !reflect.DeepEqual(err, expected) {
		t.Fatalf("got %#v\nwant %#v", err, expected)
	}

	if expected := (VsqData{}); !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %#v\nwant %#v", actual, expected)
	}
}

func Test_loadFailedValueInvalid(t *testing.T) {
	actual, err := load("testdata/value_invalid_file.json")

	if err == nil {
		t.Fatalf("Succeeded with failed test")
	}

	if expected := errors.New("not a data file of VerySimpleQueue"); !reflect.DeepEqual(err, expected) {
		t.Fatalf("got %#v\nwant %#v", err, expected)
	}

	if expected := (VsqData{}); !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %#v\nwant %#v", actual, expected)
	}
}

func TestLoadSuccess(t *testing.T) {
	var vsq VerySimpleQueue
	actual, err := vsq.load("testdata/data_file.json")

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	dataExpected := (VsqData{"VerySimpleQueue", []string{"1", "2", "3"}})
	if !reflect.DeepEqual(actual, dataExpected) {
		t.Fatalf("got %#v\nwant %#v (return value)", actual, dataExpected)
	}
	if !reflect.DeepEqual(vsq.Data, dataExpected) {
		t.Fatalf("got %#v\nwant %#v (vsq.Data)", vsq.Data, dataExpected)
	}

	if expected := "testdata/data_file.json"; vsq.FilePath != expected {
		t.Fatalf("got %#v\nwant %#v (vsq.FilePath)", vsq.FilePath, expected)
	}
}

func TestSizeWithValue(t *testing.T) {
	var vsq VerySimpleQueue
	if _, err := vsq.load("testdata/data_file.json"); err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if expected := 3; vsq.size() != 3 {
		t.Fatalf("got %#v\nwant %#v", vsq.size(), expected)
	}
}

func TestSizeNoValue(t *testing.T) {
	var vsq VerySimpleQueue
	if _, err := vsq.load("not_exist"); err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if expected := 0; vsq.size() != 0 {
		t.Fatalf("got %#v\nwant %#v", vsq.size(), expected)
	}
}

func TestWriteDbFile(t *testing.T) {
	var vsq VerySimpleQueue
	if _, err := vsq.load(testFilePath); err != nil {
		t.Fatalf("failed test %#v", err)
	}
	vsq.writeDbFile()

	bytes, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if expected := []byte(`{"Name":"VerySimpleQueue","Value":[]}`); !reflect.DeepEqual(bytes, expected) {
		t.Fatalf("got %#v\nwant %#v", bytes, expected)
	}
}

func TestUnshiftSize1(t *testing.T) {
	removeTestFile()

	var vsq VerySimpleQueue
	if _, err := vsq.load(testFilePath); err != nil {
		t.Fatalf("failed test %#v", err)
	}

	size := vsq.unshift("hoge")

	if expected := 1; size != expected {
		t.Fatalf("got %#v\nwant %#v", size, expected)
	}

	if expected := (VsqData{"VerySimpleQueue", []string{"hoge"}}); !reflect.DeepEqual(vsq.Data, expected) {
		t.Fatalf("got %#v\nwant %#v", vsq.Data, expected)
	}

	bytes, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if expected := []byte(`{"Name":"VerySimpleQueue","Value":["hoge"]}`); !reflect.DeepEqual(bytes, expected) {
		t.Fatalf("got %#v\nwant %#v", bytes, expected)
	}
}

func TestUnshiftSize3(t *testing.T) {
	removeTestFile()

	var vsq VerySimpleQueue
	if _, err := vsq.load(testFilePath); err != nil {
		t.Fatalf("failed test %#v", err)
	}

	vsq.unshift("hoge")
	vsq.unshift("fuga")
	size := vsq.unshift("piyo")

	if expected := 3; size != expected {
		t.Fatalf("got %#v\nwant %#v", size, expected)
	}

	if expected := (VsqData{"VerySimpleQueue", []string{"piyo", "fuga", "hoge"}}); !reflect.DeepEqual(vsq.Data, expected) {
		t.Fatalf("got %#v\nwant %#v", vsq.Data, expected)
	}

	bytes, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if expected := []byte(`{"Name":"VerySimpleQueue","Value":["piyo","fuga","hoge"]}`); !reflect.DeepEqual(bytes, expected) {
		t.Fatalf("got %#v\nwant %#v", bytes, expected)
	}
}

func TestShiftSuccess(t *testing.T) {
	removeTestFile()

	var vsq VerySimpleQueue
	if _, err := vsq.load(testFilePath); err != nil {
		t.Fatalf("failed test %#v", err)
	}
	vsq.unshift("hoge")

	value, err := vsq.shift()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if expected := "hoge"; value != expected {
		t.Fatalf("got %#v\nwant %#v", value, expected)
	}
}

func TestShiftFailed(t *testing.T) {
	removeTestFile()

	var vsq VerySimpleQueue
	if _, err := vsq.load(testFilePath); err != nil {
		t.Fatalf("failed test %#v", err)
	}

	value, err := vsq.shift()
	if err == nil {
		t.Fatalf("Succeeded with failed test")
	}

	if expected := ""; value != expected {
		t.Fatalf("got %#v\nwant %#v", value, expected)
	}
}

// TODO: Other functions
