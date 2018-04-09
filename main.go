package main

import (
	"io/ioutil"
	"encoding/xml"
	"fmt"
	"os"
)

type Factura struct { 
	
	Receptor []struct {
		Nombre string `xml:"nombre,attr"`
	} `xml:"Receptor"`
	
}

type Receptor struct {
	Receptor []struct {
		Nombre string `xml:"nombre,attr"`
	} `xml:"Receptor"`
}


func check(err error){
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
}

func main(){
	xmlFile, err := os.Open("./assets/XML/3596656.xml")
	//xmlFile, err := os.Open("./assets/XML/test.xml")
	check(err)
	defer xmlFile.Close()

	b, err := ioutil.ReadAll(xmlFile)

	var q Factura
	xml.Unmarshal(b, &q)

	fmt.Println(q)
}