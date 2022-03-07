package main

import "fmt"

type school interface {
	building()
	classrooms()
	students()
}
type vshs struct {
	code   int
	branch string
}

func (v vshs) building() {

}
func (v vshs) classrooms() {

}
func (v vshs) students() {

}
func isRecognised(s school) {
	fmt.Println(s, "is recognised by schools")
}
func main() {
	vshs1 := vshs{10, "jyothi nagar"}
	isRecognised(vshs1)
}
