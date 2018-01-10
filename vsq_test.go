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

// TODO: Other functions
