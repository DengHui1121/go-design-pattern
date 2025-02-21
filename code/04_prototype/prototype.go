package _4_prototype

import (
	"encoding/json"
	"time"
)

var Keywords map[string]Keyword

type Keyword struct {
	word    string
	visit   int
	updated *time.Time
}

// 深拷贝
func (k *Keyword) Clone() *Keyword {
	var newKeyword Keyword
	b, _ := json.Marshal(k)
	json.Unmarshal(b, &newKeyword)
	return &newKeyword
}

// 浅拷贝
