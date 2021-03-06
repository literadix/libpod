// +build !linux

package libpod

import (
	"os"
	"os/exec"
)

func (r *OCIRuntime) moveConmonToCgroup(ctr *Container, cgroupParent string, cmd *exec.Cmd) error {
	return ErrOSNotSupported
}

func newPipe() (parent *os.File, child *os.File, err error) {
	return nil, nil, ErrNotImplemented
}

func (r *OCIRuntime) createContainer(ctr *Container, cgroupParent string, restoreOptions *ContainerCheckpointOptions) (err error) {
	return ErrNotImplemented
}

func (r *OCIRuntime) pathPackage() string {
	return ""
}

func (r *OCIRuntime) conmonPackage() string {
	return ""
}

func (r *OCIRuntime) createOCIContainer(ctr *Container, cgroupParent string, restoreOptions *ContainerCheckpointOptions) (err error) {
	return ErrOSNotSupported
}

func (r *OCIRuntime) execStopContainer(ctr *Container, timeout uint) error {
	return ErrOSNotSupported
}

func (r *OCIRuntime) stopContainer(ctr *Container, timeout uint) error {
	return ErrOSNotSupported
}
