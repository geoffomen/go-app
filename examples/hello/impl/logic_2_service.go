package impl

import (
	"fmt"

	"github.com/storm-5/go-app/pkg/myerr"
)

func (srv *Service) Error() (string, error) {
	err := func() error {
		err := myerr.New(fmt.Errorf("first")).AddMsgf("second")
		return err
	}()
	myerr.New(err).AddMsgf("third").AddMsgf("%s", "fourth").SetCode(500)

	return "", err
}
