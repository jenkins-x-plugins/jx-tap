// +build !windows

package app

import (
	"github.com/jenkins-x-plugins/jx-tap/pkg/cmd"
)

// Run runs the command, if args are not nil they will be set on the command
func Run(args []string) error {
	c := cmd.Main()
	if args != nil {
		args = args[1:]
		c.SetArgs(args)
	}
	return c.Execute()
}
