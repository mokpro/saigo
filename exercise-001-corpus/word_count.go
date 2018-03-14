package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mokpro/saigo/exercise-001-corpus/corpus"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("File name missing!")
	}
	content, err := ioutil.ReadFile(args[1])
	if err != nil {
		panic(err)
	}

	wordFreqs := corpus.Analyze(string(content))
	for _, word := range wordFreqs {
		fmt.Println(word.Count, " ", word.Word)
	}
}
