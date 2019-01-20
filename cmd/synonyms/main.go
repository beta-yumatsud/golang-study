package main

import (
	"../thesaurus"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	apiKey := os.Getenv("BHT_APIKEY")
	thesaurusAPI := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurusAPI.Synonyms(word)
		if err != nil {
			log.Fatalf("%qの類似語検索に失敗しました: %v\n", word, err)
		}
		if len(syns) == 0 {
			log.Fatalf("%qに類似語はありませんでした\n", word)
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
