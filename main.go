package main

import(
	"fmt"
)

type Company struct {
	name string
}

func main(){
	c := Company{"erkan"}
	fmt.Println("main-age = ",c.name)
	fmt.Printf("com: %p\n", &c)
	metod(c)
	fmt.Println("main-age 22= ",c.name)
	/*age:=41
	fmt.Println("main-age = ",age)
	//fmt.Println("main-address of age",&age)
	call(age)
	fmt.Println("main-age = ",age)
	//callWithPointer(&age)
	 */
	//fmt.Println("main-age = ",age)
}


func metod(c Company){

	fmt.Printf("com AFTER: %p\n", &c)

	fmt.Println("VALUE main-address of age", c.name)
	c.name = "ABUZAETTIN"
	fmt.Println("call-value = ", c.name)
	//fmt.Println("call-address of value",&value)
}


func call(value int){
	fmt.Println("VALUE main-address of age",value)
	value+=1
	fmt.Println("call-value = ",value)
	//fmt.Println("call-address of value",&value)
}

func callWithPointer(value *int){
	*value+=1
	fmt.Println("callWithPointer-value = ",*value)
	fmt.Println("callWithPointer-address of value",value)
}