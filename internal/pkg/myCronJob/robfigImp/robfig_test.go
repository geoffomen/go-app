package robfigImp

import (
	"fmt"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	ins, _ := New(newLogger())

	ins.AddScheduleFunc("*/3 * * * * *", func() {
		fmt.Println("hello world.")
	})
	go ins.Start()

	time.Sleep(5 * time.Second)
	ins.Stop()
	time.Sleep(5 * time.Second)
	fmt.Println("bye")
}
