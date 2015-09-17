package main


import ("fmt")

type Test struct{
	Abc []string
	Test Test2
}

type Test2 struct {
	Names []string
	
}

func main(){
	var t Test
	t.Dis="asd"
	setValues(&t)
	fmt.Println(t)
}


func setValues( t *Test){
	jarjar:= []string{"asd","baasdfasdfrk"}
	binks:= []string{"the ","force"}
	
	t.Abc=jarjar



	//fmt.Println(dog)
}