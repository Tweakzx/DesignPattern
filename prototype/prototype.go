package prototype

import (
	"encoding/json"
	"time"
)

type Keyword struct {
	word      string
	visit     int
	UpdatedAt *time.Time
}

//使用序列化和反序列化实现深拷贝
func (k *Keyword) Clone() *Keyword {
	var newKeyword Keyword
	b, _ := json.Marshal(k)
	json.Unmarshal(b, &newKeyword)
	return &newKeyword
}


type Keywords map[string]*Keyword

func (words Keywords) Clone(updateWords []*Keyword) Keywords {
	newKeywords := Keywords{}
	for k, v := range words {
		//浅拷贝
		newKeywords[k] = v
	}
	for _, word := range updateWords {
		//深拷贝
		newKeywords[word.word] = word.Clone()
	}
	return newKeywords
}