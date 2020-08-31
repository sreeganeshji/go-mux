package main

import (
	"os"
	"fmt"
)

func main(){
	a := App{}
	a.Initialize(os.Getenv("POSTGRES_USER"),os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	a.Run(":8010")


	//for a, ele:= range(os.Environ()){
	//	fmt.Println(a,ele)
	//}


}