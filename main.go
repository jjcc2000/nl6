package main

import (
	"fmt"
	"math"
)	
type Person struct{
	first string
	last string 
	age int
}

func main(){

	x:=pp(XX(1))
	fmt.Println(x)
		
}
//TODO://Exercise Number 1	
func boo()int{
	return 1
}
func bar() (string ,int){
	x:= "Johan"
	return x,20
}
//TODO://Exercise Number 2
func ej2(x...int) int{
	total:=0
	for _, v:=range x{
		total+= v
	}
	return total
}
func ej2_1(x []int) int {
	total:= 0
	for _, val := range x{
		total+=val
	}
	return total
}
//TODO:// Exercise Number 3////////
func ej3(){
//FIXME:// The defer keyword is use to put the function at the bottom
	defer fmt.Println("This function is going last in the main")
	fmt.Println("This is going first")
}
//TODO:Exercise Number 4//////////

func speak(){
	fmt.Println("This is the person speaking")

}
func (p Person) speak(){
	fmt.Println("Just conect the Speak method to the struct "+
	"call Person therefore the method can be call just with a Type person")

	fmt.Println("The name is ", p.first," the second name is ",p.last," the age ",p.age)
}
////////////////////////////////////TODO:// Exercise Number 5 ///////////////
type square struct {
	lenght float64
}
type circle struct{
	radio float64
}
//Create a type shape which defines an interface as anything that has the area method
type shape	interface{
	area() float64
}
func (c circle) area() float64 {
	return math.Pi * c.radio *c.radio
}
func (s square) area() float64{
	return s.lenght * s.lenght
}
func info(s shape) {
	fmt.Println(s.area())
}

////////////////////////////////TODO://Exercise Number 6 ///////////////////////
///////////////////FIXME:////// TO create an uninamous function just add () after the closing brace to executed///////////// 



////////////////////////////////TODO://EXERCISE NUMBER 7 //////////////////////
///////////////FIXME:// x:= func (){ .......} to asign the func to a variable///

//////////////////////////////TODO://EXERCISE NUmber 8/////////////////////////////
func y(i int) func() int {
	return func() int{
		return i+1
	}
}
////////////////////////////////TODO://Exercise Number 9/////////////////////////
////////////////////////FIXME: This is callbacks FIXME://///////////////////////////
func XX(i int) int{
	return i+1
}
func pp(i int) int{
	return i*2
}
//////////////////////////////TODO://Exercise Number10///////////////////////////
