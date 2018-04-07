package main

import (
	"io/ioutil"
	"encoding/xml"
	"fmt"
)

type Factura struct{

}

func check(err error){
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
}

func main(){
	xmlFile, err := os.Open("assets/3596656.xml")
	check(err)
	defer xmlFile.Close()

	b, err := ioutil.ReadAll(xmlFile)

	var q Query
	xml.Unmarshal(b, &q)

	fmt.Println()
}