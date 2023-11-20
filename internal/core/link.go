package core

import (
	"context"

	"github.com/goexl/gox"
	"github.com/pangum/grpc"
	"github.com/pangum/yangzone/internal/internal/constant"
	"github.com/yangzone/core/link"
	"github.com/yangzone/protocol"
	"gitlab.com/ruijc/app/core"
)

// Link 链接
type Link struct {
	protocol.LinkClient

	app core.Id
	_   gox.CannotCopy
}

func NewLink(client *grpc.Client, app core.Id) *Link {
	return &Link{
		LinkClient: protocol.NewLinkClient(client.Connection(constant.LabelYangzone)),
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
