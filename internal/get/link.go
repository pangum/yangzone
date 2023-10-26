package get

import (
	"gitea.com/ruijc/app"
	"github.com/pangum/grpc"
	"github.com/pangum/pangu"
)

type Link struct {
	pangu.Get

	Client *grpc.Client
	App    app.Id
}
