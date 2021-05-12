package exapmle_test

import (
	"reflect"
	"testing"

	"github.com/im-jinsu/interchange"
)

func TestConvert(t *testing.T) {
	var tmp = []interface{}{
		"apple",
		"grape",
		"orange",
	}
	interchange := new(interchange.Gate)
	interchange.TargetArray = tmp
	interchange.Do()
	for _, v := range interchange.StringArray {
		println(reflect.TypeOf(v).String(), v)
	}
}
