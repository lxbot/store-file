package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/lxbot/lxlib/v2"
	"github.com/lxbot/lxlib/v2/common"
	"github.com/lxbot/lxlib/v2/lxtypes"
)

type M = map[string]interface{}

const fileName = "store.json"

var filePath string
var wg sync.WaitGroup

func main() {
	store, getCh, setCh := lxlib.NewStore()

	var dir string
	if fp, err := filepath.Abs(filepath.Join("data", "store")); err == nil {
		dir = fp
	} else {
		common.FatalLog(err)
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		common.FatalLog(err)
	}
	if fp, err := filepath.Abs(filepath.Join(dir, fileName)); err == nil {
		filePath = fp
	} else {
		common.FatalLog(err)
	}
	common.InfoLog("store file path:", filePath)

	for {
		select {
		case event := <-*getCh:
			onGet(event, store)
		case event := <-*setCh:
			onSet(event)
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

func onGet(event *lxtypes.StoreEvent, store *lxlib.Store) {
	m := M{}
	if b, err := os.ReadFile(filePath); err == nil {
		_ = json.Unmarshal(b, &m)
	}

	if value, ok := m[event.Key]; ok {
		event.Value = value
	}

	store.SendGetResult(event)
}

func onSet(event *lxtypes.StoreEvent) {
	m := M{}
	if b, err := os.ReadFile(filePath); err == nil {
		_ = json.Unmarshal(b, &m)
	}

	m[event.Key] = event.Value

	b, err := json.Marshal(m)
	if err != nil {
		common.DebugLog(err)
	}
	_ = os.WriteFile(filePath, b, 0600)
}
