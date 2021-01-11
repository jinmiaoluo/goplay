package main

import (
	"bytes"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("small.txt")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("content is: %s\n", content)

	// 只替换一次的情况
	//bytes.Replace(content, []byte("你好"), []byte("你好呀"), 1)

	// 全部替换
	result := bytes.ReplaceAll(content, []byte("你好"), []byte("你好呀"))
	if err := ioutil.WriteFile("result.txt", result, 0644); err != nil {
		log.Fatal(err)
	}
}