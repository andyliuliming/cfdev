package bosh

import (
	"os"

	"runtime"

	"code.cloudfoundry.org/cfdev/bosh"
	"code.cloudfoundry.org/cfdev/config"
	"code.cloudfoundry.org/cfdev/errors"
	"code.cloudfoundry.org/cfdev/shell"
	"github.com/spf13/cobra"
)

type UI interface {
	Say(message string, args ...interface{})
}

type Provisioner interface {
	FetchBOSHConfig() (bosh.Config, error)
}

type Bosh struct {
	Exit        chan struct{}
	UI          UI
	Config      config.Config
	Provisioner Provisioner
}

func (b *Bosh) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "bosh",
		Run: func(cmd *cobra.Command, args []string) {
			if runtime.GOOS != "windows" {
				b.UI.Say(`Usage: eval $(cf dev bosh env)`)
			} else {
				b.UI.Say(`Usage: cf dev bosh env | Invoke-Expression`)
			}
		},
	}
	envCmd := &cobra.Command{
		Use:  "env",
		RunE: b.RunE,
	}
	cmd.AddCommand(envCmd)
	return cmd
}

func (b *Bosh) RunE(cmd *cobra.Command, args []string) error {
	go func() {
		<-b.Exit
		os.Exit(128)
	}()

	config, err := b.Provisioner.FetchBOSHConfig()
	if err != nil {
		return errors.SafeWrap(err, "failed to fetch bosh configuration")
	}

	env := shell.Environment{StateDir: b.Config.StateDir}
	shellScript, err := env.Prepare(config)
	if err != nil {
		return errors.SafeWrap(err, "failed to prepare bosh configuration")
	}

	b.UI.Say(shellScript)
	return nil
}
