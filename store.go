package main

import (
	"encoding/json"
	"io/ioutil"
)

type M = map[string]interface{}

const fileName = "store.json"
var ch *chan M

func Boot(c *chan M) {
	ch = c
}

func Set(key string, value interface{}) {
	m := map[string]interface{}{}
	if b, err := ioutil.ReadFile(fileName); err == nil {
		_ = json.Unmarshal(b, &m)
	}

	m[key] = value

	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	_ = ioutil.WriteFile(fileName, b, 0600)
}

func Get(key string) interface{} {
	m := map[string]interface{}{}
	if b, err := ioutil.ReadFile(fileName); err == nil {
		_ = json.Unmarshal(b, &m)
	}

	return m[key]
}
