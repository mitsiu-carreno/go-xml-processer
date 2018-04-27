package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
)

type Node struct{
	XMLName xml.Name
	Attrs 	[]xml.Attr 	`xml:"-"`
	Content []byte		`xml:",innerxml"`
	Nodes	[]Node		`xml:",any"`
}

func (n *Node) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	n.Attrs = start.Attr
	type node Node

	return d.DecodeElement((*node)(n), &start)
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

	buf := bytes.NewBuffer(xmlFile)

	dec := xml.NewDecoder(buf)

	var n Node
	err = dec.Decode(&n)
	check(err)
	/*
	walk([]Node{n}, func(n Node) bool {
				fmt.Println(string(n.XMLName))
				fmt.Println(string(n.Attrs))
				fmt.Println(string(n.Content))
				return true
			}
		)
		*/
}

func walk(nodes []Node, f func(Node) bool){
	for _, n := range nodes {
		if(f(n)){
			walk(n.Nodes, f)
		}
	}
}