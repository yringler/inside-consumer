package main

import (
	"fmt"

	wordpressxml "github.com/grokify/wordpress-xml-go"
)

func main() {
	wp := wordpressxml.NewWordpressXml()
	err := wp.ReadXml("archive.xml")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(wp.Channel.Items))
	fmt.Println(wp.Channel.Items[0].Content)
	fmt.Println(wp.Channel.Items[0].Categories[0].DisplayName)
}
