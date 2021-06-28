package main

type PointerTest struct {
	number uint
	str string
}

func definePointer() (ptrTest *PointerTest) {
	// The new built-in function allocates memory. The first argument is a type,
	// not a value, and the value returned is a pointer to a newly
	// allocated zero value of that type.
	ptrTest = new(PointerTest)
	//var ptrTest *PointerTest	// nil create nil pointer exception because ptrTest is pointer type of PointerTest
	ptrTest.number = 1
	ptrTest.str = "this is a test"
	return
}
