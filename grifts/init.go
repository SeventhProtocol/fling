package grifts

import (
	"github.com/SeventhProtocol/fling/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
