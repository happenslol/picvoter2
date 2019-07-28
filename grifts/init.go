package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/happenslol/picvoter2/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
