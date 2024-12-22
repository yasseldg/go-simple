package repos

import (
	"sync"
	"testing"
	"time"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repos/rMongo"
)

func TestRunConcurr(t *testing.T) {
	mongo := rMongo.New()
	RunConcurr(mongo)
}

func TestFindFilter(t *testing.T) {
	filter := NewFilter().Uuid("uuid_6")
	doc, err := findFilter(filter, 1)
	if err != nil {
		t.Errorf("findFilter() error = %v", err)
	}
	if doc == nil {
		t.Errorf("Expected document, got nil")
	}
}

func TestConcurrentFindFilter(t *testing.T) {
	mongo := rMongo.New()
	err := config(mongo)
	if err != nil {
		t.Errorf("config() error = %v", err)
		return
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		filter := NewFilter().Uuid("uuid_6")
		doc, err := findFilter(filter, 1)
		if err != nil {
			sLog.Error("find(): %s", err)
		}
		if doc != nil {
			doc.Log()
		}
	}()

	time.Sleep(1 * time.Second)

	wg.Add(1)
	go func() {
		defer wg.Done()
		filter := NewFilter().Uuid("uuid_4")
		doc, err := findFilter(filter, 1)
		if err != nil {
			sLog.Error("find(): %s", err)
		}
		if doc != nil {
			doc.Log()
		}
	}()

	wg.Wait()
}
