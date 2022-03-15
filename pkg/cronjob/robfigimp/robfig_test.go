package robfigimp

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	ins, _ := New(log.Default())

	ins.AddScheduleFunc("*/3 * * * * *", func() {
		fmt.Println("hello world.")
	})
	ins.Start()

	time.Sleep(5 * time.Second)
	ins.Stop()
	time.Sleep(5 * time.Second)
	fmt.Println("bye")
}
