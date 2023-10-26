package yangzone

import (
	"github.com/pangum/pangu"
	"github.com/pangum/yangzone/internal/plugin"
)

func init() {
	creator := new(plugin.Creator)
	pangu.New().Get().Dependency().Put(
		creator.NewLink,
	).Build().Build().Apply()
}
