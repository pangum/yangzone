package plugin

import (
	"github.com/pangum/yangzone/internal/core"
	"github.com/pangum/yangzone/internal/put"
)

type Constructor struct {
	// 构造方法
}

func (*Constructor) NewLink(link put.Link) *core.Link {
	return core.NewLink(link.Client, link.App)
}
