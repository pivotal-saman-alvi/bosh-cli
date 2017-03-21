package cmd

import (
	"os"

	goflags "github.com/jessevdk/go-flags"
	"path/filepath"
	"fmt"

	boshreldir "github.com/cloudfoundry/bosh-cli/releasedir"
	boshui "github.com/cloudfoundry/bosh-cli/ui"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
)

type AddBlobCmd struct {
	blobsDir boshreldir.BlobsDir
	fs       boshsys.FileSystem
	ui       boshui.UI
}

func NewAddBlobCmd(blobsDir boshreldir.BlobsDir, fs boshsys.FileSystem, ui boshui.UI) AddBlobCmd {
	return AddBlobCmd{blobsDir: blobsDir, fs: fs, ui: ui}
}

func (c AddBlobCmd) Run(opts AddBlobOpts) error {
	file, err := c.fs.OpenFile(opts.Args.Path, os.O_RDONLY, 0)
	if err != nil {
		return bosherr.WrapErrorf(err, "Opening blob")
	}

	defer file.Close()

	blob, err := c.blobsDir.TrackBlob(opts.Args.BlobsPath, file)
	if err != nil {
		return bosherr.WrapErrorf(err, "Tracking blob")
	}

	c.ui.PrintLinef("Added blob '%s'", blob.Path)

	return nil
}

func (c AddBlobCmd) Complete(match string) []goflags.Completion {
	fmt.Sprintf("=========== hallelujah")
	matchedFiles, _ := filepath.Glob(match + "*")
	ret := make([]goflags.Completion, len(matchedFiles))

	for i, v := range matchedFiles {
		ret[i].Item = v
	}

	return ret
}


