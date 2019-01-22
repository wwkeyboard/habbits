package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/wwkeyboard/habbits/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
