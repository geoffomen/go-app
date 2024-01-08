package echoargsdm

import "ibingli.com/internal/pkg/myDatabase"

type SrvIface interface {
	NewWithDb(db myDatabase.Iface) SrvIface
}
