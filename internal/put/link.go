package put

import (
	"github.com/pangum/grpc"
	"github.com/pangum/pangu"
	"gitlab.com/ruijc/app/core"
)

type Link struct {
	pangu.Put

	Client *grpc.Client
	App    core.Id
}
