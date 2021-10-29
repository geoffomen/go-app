package cronjob

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/geoffomen/go-app/pkg/cronjob/robfigimp"
)

func TestCron(t *testing.T) {
	ins, _ := robfigimp.New(log.Default())

	ins.AddScheduleFunc("*/3 * * * * *", func() {
		fmt.Println("hello world.")
	})
	ins.Start()

	time.Sleep(5 * time.Second)
	ins.Stop()
	time.Sleep(5 * time.Second)
	fmt.Println("bye")
}
