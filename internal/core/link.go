package core

import (
	"context"

	"gitea.com/ruijc/app"
	"github.com/goexl/gox"
	"github.com/pangum/grpc"
	"github.com/yangzone/core/link"
	"github.com/yangzone/protocol"
)

// Link 链接
type Link struct {
	protocol.LinkClient

	app app.Id
	_   gox.CannotCopy
}

func NewLink(client *grpc.Client, app app.Id) *Link {
	return &Link{
		LinkClient: protocol.NewLinkClient(client.Connection("yangzone")),
		app:        app,
	}
}

func (l *Link) Add(ctx context.Context, req *link.AddReq) (rsp *link.AddRsp, err error) {
	req.App = int64(l.app)
	rsp, err = l.LinkClient.Add(ctx, req)

	return
}

func (l *Link) Paging(ctx context.Context, req *link.PagingReq) (rsp *link.PagingRsp, err error) {
	req.App = int64(l.app)
	rsp, err = l.LinkClient.Paging(ctx, req)

	return
}
