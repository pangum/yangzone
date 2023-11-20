package get

import (
	"github.com/pangum/grpc"
	"github.com/pangum/pangu"
	"gitlab.com/ruijc/app/core"
)

type Link struct {
	pangu.Get

	Client *grpc.Client
	App    core.Id
}
