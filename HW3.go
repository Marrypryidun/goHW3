package main

import (
	"fmt"
	"math"
)
type person struct {
	name string
	age int
	height float64
	weight float64
	job *worker
}
type worker struct {
	position string
	experience float64
}
func (p person) String() string {
	return fmt.Sprintf("%v (%v years): %v (%v years of experience)", p.name, p.age,p.job.position,p.job.experience)
}
func (receiver *person) happyBirthday (){
	fmt.Println("Happy birthday",receiver.name)
}
func (p person) bodyMassIndex() {
	var index float64 = p.weight/math.Pow(p.height,2)
	fmt.Println("BMI:",index)

}
func main() {
	first := &person{"Maria", 18, 1.68,47,&worker{"QA tester", 0.5}}
	//second := &person{}
	fmt.Println(first/*,second panic*/)
	first.happyBirthday();
	first.bodyMassIndex()
	
}
