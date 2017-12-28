package main

import (
	"github.com/VitaliiMichailovich/exercism-go/go/flatten-array"
	"fmt"
)

func main() {
	fmt.Println(flatten.Flatten([]interface{}{interface{}(nil), []interface{}{[]interface{}{[]interface{}{interface{}(nil)}}}, interface{}(nil), interface{}(nil), []interface{}{[]interface{}{interface{}(nil), interface{}(nil)}, interface{}(nil)}, interface{}(nil)}))
}
