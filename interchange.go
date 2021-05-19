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
	Uint32Array  []uint32
	Uint64Array  []uint64
	StringArray  []string
	BoolArray    []bool
	Float32Array []float32
	Float64Array []float64
}

func (gate *Gate) Do() {
	if len(gate.TargetArray) > 0 {
		switch reflect.TypeOf(gate.TargetArray[0]).String() {
		case "int":
			for _, r := range gate.TargetArray {
				gate.IntArray = append(gate.IntArray, r.(int))
			}
		case "int8":
			for _, r := range gate.TargetArray {
				gate.Int8Array = append(gate.Int8Array, r.(int8))
			}
		case "int16":
			for _, r := range gate.TargetArray {
				gate.Int16Array = append(gate.Int16Array, r.(int16))
			}
		case "int32":
			for _, r := range gate.TargetArray {
				gate.Int32Array = append(gate.Int32Array, r.(int32))
			}
		case "int64":
			for _, r := range gate.TargetArray {
				gate.Int64Array = append(gate.Int64Array, r.(int64))
			}
		case "uint":
			for _, r := range gate.TargetArray {
				gate.UintArray = append(gate.UintArray, r.(uint))
			}
		case "uint8":
			for _, r := range gate.TargetArray {
				gate.Uint8Array = append(gate.Uint8Array, r.(uint8))
			}
		case "uint16":
			for _, r := range gate.TargetArray {
				gate.Uint16Array = append(gate.Uint16Array, r.(uint16))
			}
		case "uint32":
			for _, r := range gate.TargetArray {
				gate.Uint32Array = append(gate.Uint32Array, r.(uint32))
			}
		case "uint64":
			for _, r := range gate.TargetArray {
				gate.Uint64Array = append(gate.Uint64Array, r.(uint64))
			}
		case "string":
			for _, r := range gate.TargetArray {
				gate.StringArray = append(gate.StringArray, r.(string))
			}
		case "bool":
			for _, r := range gate.TargetArray {
				gate.BoolArray = append(gate.BoolArray, r.(bool))
			}
		case "float32":
			for _, r := range gate.TargetArray {
				gate.Float32Array = append(gate.Float32Array, r.(float32))
			}
		case "float64":
			for _, r := range gate.TargetArray {
				gate.Float64Array = append(gate.Float64Array, r.(float64))
			}
		}
	}
}
