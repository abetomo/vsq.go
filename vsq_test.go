// WIP
package vsq

import (
	"errors"
	"reflect"
	"testing"
)

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

// TODO: Other functions
