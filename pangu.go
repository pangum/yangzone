package yangzone

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependency(newLink)
}
