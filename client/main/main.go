package main

import (
	"EasyLinks/client"
)

func main() {
	cln := client.NewClient("http://localhost:8081")
	resp, err := cln.CreateLink()
	if err != nil {
		panic(err)
	}
	println(resp)

	link, err := cln.GetLink()
	if err != nil {
		panic(err)
	}
	println(link)

}
