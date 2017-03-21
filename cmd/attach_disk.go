package cmd

import goflags "github.com/jessevdk/go-flags"
import boshdir "github.com/cloudfoundry/bosh-cli/director"
import boshui "github.com/cloudfoundry/bosh-cli/ui"
import "fmt"

type AttachDiskCmd struct {
	deployment boshdir.Deployment
	ui boshui.UI
}

func NewAttachDiskCmd(deployment boshdir.Deployment, ui boshui.UI) AttachDiskCmd {
	return AttachDiskCmd {
		deployment: deployment,
		ui: ui,
	}
}

func (c AttachDiskCmd) Run(opts AttachDiskOpts) error {
	return c.deployment.AttachDisk(opts.Args.Slug, opts.Args.DiskCID)
}

func (c AttachDiskCmd) Complete(match string) []goflags.Completion {
	fmt.Println("==== WHY HELLO BOB IM ON MY MERRY WAY NOW")
	c.ui.PrintLinef("attaching a disk!", c.deployment.Name())
	ret := make([]goflags.Completion, 1)
	ret[0] = goflags.Completion{Item: c.deployment.Name()}

	return ret
}