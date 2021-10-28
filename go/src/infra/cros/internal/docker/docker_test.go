package docker_test

import (
	"context"
	"errors"
	"testing"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/strslice"

	"infra/cros/internal/cmd"
	"infra/cros/internal/docker"
)

func TestRunContainer(t *testing.T) {
	ctx := context.Background()

	containerConfig := &container.Config{
		Cmd:   strslice.StrSlice{"ls", "-l"},
		User:  "testuser",
		Image: "testimage",
	}

	hostConfig := &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: "/tmp/hostdir",
				Target: "/usr/local/containerdir",
			},
			{
				Type:     mount.TypeBind,
				Source:   "/othersource",
				Target:   "/othertarget",
				ReadOnly: true,
			},
		},
		NetworkMode: "host",
	}

	cmdRunner := cmd.FakeCommandRunner{}
	cmdRunner.ExpectedCmd = []string{
		"docker", "run",
		"--user", "testuser",
		"--network", "host",
		"--mount=source=/tmp/hostdir,target=/usr/local/containerdir,type=bind",
		"--mount=source=/othersource,target=/othertarget,type=bind,readonly",
		"testimage",
		"ls", "-l",
	}

	err := docker.RunContainer(ctx, cmdRunner, containerConfig, hostConfig)
	if err != nil {
		t.Fatalf("RunContainer failed: %s", err)
	}
}

func TestRunContainer_CmdError(t *testing.T) {
	ctx := context.Background()

	containerConfig := &container.Config{
		Cmd: strslice.StrSlice{"ls", "-l"},
	}

	hostConfig := &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: "/tmp/hostdir",
				Target: "/usr/local/containerdir",
			},
		},
	}

	cmdRunner := cmd.FakeCommandRunner{}
	cmdRunner.FailCommand = true
	cmdRunner.FailError = errors.New("docker cmd failed.")

	err := docker.RunContainer(ctx, cmdRunner, containerConfig, hostConfig)
	if err == nil {
		t.Errorf("RunContainer expected to fail")
	}
}
