package yangzone

import (
	"context"

	"gitea.com/ruijc/app"
	"github.com/goexl/gox"
	"github.com/pangum/yangzone/internal/get"
	"github.com/yangzone/core/link"
	"github.com/yangzone/protocol"
)

// Link 链接
type Link struct {
	protocol.LinkClient

	app app.Id
	_   gox.CannotCopy
}

func newLink(get get.Link) *Link {
	return &Link{
		LinkClient: protocol.NewLinkClient(get.Client.Connection("yangzone")),
		app:        get.App,
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
