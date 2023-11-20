package plugin

import (
	"github.com/pangum/yangzone/internal/core"
	"github.com/pangum/yangzone/internal/get"
)

type Constructor struct {
	// 构造方法
}

func (*Constructor) NewLink(get get.Link) *core.Link {
	return core.NewLink(get.Client, get.App)
}
