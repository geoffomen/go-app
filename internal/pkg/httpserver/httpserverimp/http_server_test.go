package httpserverimp

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mylogger struct{}

func (ml *mylogger) Infof(format string, args ...interface{}) {
	log.Printf(format, args...)
}
func (ml *mylogger) Errorf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (ml *mylogger) Fatalf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func TestServer(t *testing.T) {

	type ReqStruct struct {
		Key1 string
		Key2 int
		Key3 struct {
			Key5 string
		}
		Key4 []int
	}
	arg := ReqStruct{
		Key1: "val1",
		Key2: 100000,
		Key3: struct{ Key5 string }{
			Key5: "val5",
		},
		Key4: []int{1, 2, 3},
	}
	b, _ := json.Marshal(arg)
	req, _ := http.NewRequest(http.MethodPost, "http://127.0.0.1:50000", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	srv := New(&Options{
		Port:   50000,
		Logger: &mylogger{},
	})
	srv.AddRouter(map[string]interface{}{
		"POST /": func(arg ReqStruct, req *http.Request, w http.ResponseWriter) (interface{}, error) {
			assert.Equal(t, "val1", arg.Key1)
			assert.Equal(t, 100000, arg.Key2)
			assert.Equal(t, "val5", arg.Key3.Key5)
			assert.Equal(t, []int{1, 2, 3}, arg.Key4)
			return nil, nil
		},
	})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		// defer wg.Done()  // this scenario not neccessary
		srv.Listen()
	}()
	time.Sleep(3 * time.Second)

	client := http.Client{}
	_, _ = client.Do(req)
	wg.Wait()
}
