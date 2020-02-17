package main

import (
	"github.com/grokify/wordpress-xml-go"
)

func main() {
	wp := wordpressxml.NewWordpressXml()
	err := wp.ReadXml("insidechassidus.xml")
	if err != nil {
		panic(err)
	}
	wp.WriteMetaCsv("articles.csv")
}