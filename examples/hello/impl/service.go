package impl

import (
	"fmt"

	"github.com/storm-5/go-app/pkg/database"
	"github.com/storm-5/go-app/pkg/myerr"
	"github.com/storm-5/go-app/pkg/webfw"

	"github.com/storm-5/go-app/examples/hello"
)

type Service struct {
	db *database.Client
}

var instance *Service = &Service{}

// New ...
func New(db *database.Client,
) *Service {
	*instance = Service{
		db: db,
	}
	return instance
}

// GetInstance ..
func GetInstance() *Service {
	return instance
}

// NewWithDBClient ..
func (srv *Service) NewWithDb(db *database.Client) hello.Iface {
	newSrv := *instance
	newSrv.db = db
	return &newSrv
}

func (srv *Service) SayHello(sessInfo webfw.SessionInfo) (interface{}, error) {
	return fmt.Sprintf("HELLO, %s!", sessInfo.Name), nil
}

func (srv *Service) Echo(args hello.EchoReqDto) (hello.EchoRspDto, error) {
	rsp := hello.EchoRspDto{}
	rsp.IntVal = args.IntVal
	rsp.IntPtrVal = *args.IntPtrVal
	rsp.StrVal = args.StrVal
	rsp.StructVal.Id = args.StructVal.Id
	rsp.SliceVal = args.SliceVal
	return rsp, nil
}

func (srv *Service) Error() (string, error) {
	err := func() error {
		err := myerr.New(fmt.Errorf("first")).AddMsgf("second")
		return err
	}()
	myerr.New(err).AddMsgf("third").AddMsgf("%s", "fourth").SetCode(500)

	return "", err
}
