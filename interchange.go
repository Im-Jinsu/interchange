package interchange

import "reflect"

type Gate struct {
	Type         string
	TargetArray  []interface{}
	IntArray     []int
	Int8Array    []int8
	Int16Array   []int16
	Int32Array   []int32
	Int64Array   []int64
	UintArray    []uint
	Uint8Array   []uint8
	Uint16Array  []uint16
	Uit32Array   []uint32
	Uit64Array   []uint64
	StringArray  []string
	BoolArray    []bool
	Float32Array []float32
	Float64Array []float64
}

func (gate *Gate) Do() {
	if len(gate.TargetArray) > 0 {
		switch reflect.TypeOf(gate.TargetArray[0]).String() {
		case "string":
			for _, r := range gate.TargetArray {
				gate.StringArray = append(gate.StringArray, r.(string))
			}
		case "int":
			for _, r := range gate.TargetArray {
				gate.IntArray = append(gate.IntArray, r.(int))
			}
		}
	}
}
