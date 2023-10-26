package plugin

import (
	"github.com/pangum/yangzone/internal/core"
	"github.com/pangum/yangzone/internal/get"
)

type Creator struct {
	// 解决命名空间问题
}

func (c *Creator) NewLink(get get.Link) *core.Link {
	return core.NewLink(get.Client, get.App)
}
