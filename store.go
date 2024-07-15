package main

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/lxbot/lxlib/v2"
	"github.com/lxbot/lxlib/v2/common"
	"github.com/lxbot/lxlib/v2/lxtypes"
)

type M = map[string]interface{}

const fileName = "store.json"

var wg sync.WaitGroup

func main() {
	store, getCh, setCh := lxlib.NewStore()

	for {
		select {
		case event := <-*getCh:
			onGet(event, store)
		case event := <-*setCh:
			onSet(event)
		}
	}
}

func onGet(event *lxtypes.StoreEvent, store *lxlib.Store) {
	m := M{}
	if b, err := os.ReadFile(fileName); err == nil {
		_ = json.Unmarshal(b, &m)
	}

	if value, ok := m[event.Key]; ok {
		event.Value = value
	}

	store.SendGetResult(event)
}

func onSet(event *lxtypes.StoreEvent) {
	m := M{}
	if b, err := os.ReadFile(fileName); err == nil {
		_ = json.Unmarshal(b, &m)
	}

	m[event.Key] = event.Value

	b, err := json.Marshal(m)
	if err != nil {
		common.DebugLog(err)
	}
	_ = os.WriteFile(fileName, b, 0600)
}
