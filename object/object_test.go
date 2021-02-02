package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("Strings with same content should have the same hash key")
	}
	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("Strings with same content should have the same hash key")
	}
	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("Strings with different content should have different hash key")
	}
}

func TestIntegerHashKey(t *testing.T) {
	hello1 := &Integer{Value: 1}
	hello2 := &Integer{Value: 1}
	diff1 := &Integer{Value: 2}
	diff2 := &Integer{Value: 2}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("Integers with same content should have the same hash key")
	}
	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("Integer with same content should have the same hash key")
	}
	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("Integer with different content should have different hash key")
	}
}

func TestBooleanHashKey(t *testing.T) {
	hello1 := &Boolean{Value: true}
	hello2 := &Boolean{Value: true}
	diff1 := &Boolean{Value: false}
	diff2 := &Boolean{Value: false}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("Booleans with same content should have the same hash key")
	}
	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("Booleans with same content should have the same hash key")
	}
	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("Booleans with different content should have different hash key")
	}
}
