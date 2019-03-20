// +build go1.12,!go1.13

package stdlib

// Code generated by 'goexports math/rand'. DO NOT EDIT.

import (
	"math/rand"
	"reflect"
)

func init() {
	Value["math/rand"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ExpFloat64":  reflect.ValueOf(rand.ExpFloat64),
		"Float32":     reflect.ValueOf(rand.Float32),
		"Float64":     reflect.ValueOf(rand.Float64),
		"Int":         reflect.ValueOf(rand.Int),
		"Int31":       reflect.ValueOf(rand.Int31),
		"Int31n":      reflect.ValueOf(rand.Int31n),
		"Int63":       reflect.ValueOf(rand.Int63),
		"Int63n":      reflect.ValueOf(rand.Int63n),
		"Intn":        reflect.ValueOf(rand.Intn),
		"New":         reflect.ValueOf(rand.New),
		"NewSource":   reflect.ValueOf(rand.NewSource),
		"NewZipf":     reflect.ValueOf(rand.NewZipf),
		"NormFloat64": reflect.ValueOf(rand.NormFloat64),
		"Perm":        reflect.ValueOf(rand.Perm),
		"Read":        reflect.ValueOf(rand.Read),
		"Seed":        reflect.ValueOf(rand.Seed),
		"Shuffle":     reflect.ValueOf(rand.Shuffle),
		"Uint32":      reflect.ValueOf(rand.Uint32),
		"Uint64":      reflect.ValueOf(rand.Uint64),

		// type definitions
		"Rand":     reflect.ValueOf((*rand.Rand)(nil)),
		"Source":   reflect.ValueOf((*rand.Source)(nil)),
		"Source64": reflect.ValueOf((*rand.Source64)(nil)),
		"Zipf":     reflect.ValueOf((*rand.Zipf)(nil)),
	}
}