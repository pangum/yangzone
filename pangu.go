package yangzone

import (
	"github.com/pangum/pangu"
	"github.com/pangum/yangzone/internal/plugin"
)

func init() {
	ctor := new(plugin.Constructor)
	pangu.New().Get().Dependency().Put(
		ctor.NewLink,
	).Build().Build().Apply()
}
