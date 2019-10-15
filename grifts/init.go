package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/splice/myawesomeapp/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
