// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package tlslib provides the canonical implementation of a common TLS server.
package tlslib

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"path"
	"strings"

	"go.chromium.org/chromiumos/config/go/api/test/tls"
	"golang.org/x/crypto/ssh"
)

type provisionState struct {
	s                 *Server
	c                 *ssh.Client
	dutName           string
	imagePath         string
	targetBuilderPath string
}

func newProvisionState(s *Server, req *tls.ProvisionDutRequest) (*provisionState, error) {
	p := &provisionState{
		s:       s,
		dutName: req.Name,
	}

	// Verify the incoming request path oneof is valid.
	switch t := req.GetTargetBuild().GetPathOneof().(type) {
	// Requests with gs_path_prefix should be in the format:
	// gs://chromeos-image-archive/eve-release/R86-13388.0.0
	case *tls.ChromeOsImage_GsPathPrefix:
		p.imagePath = req.GetTargetBuild().GetGsPathPrefix()
	default:
		return nil, fmt.Errorf("newProvisionState: unsupported ImagePathOneof in ProvisionDutRequest, %T", t)
	}

	u, uErr := url.Parse(p.imagePath)
	if uErr != nil {
		return nil, fmt.Errorf("setPaths: failed to parse path %s, %s", p.imagePath, uErr)
	}

	d, version := path.Split(u.Path)
	p.targetBuilderPath = path.Join(path.Base(d), version)

	return p, nil
}

func (p *provisionState) connect(ctx context.Context, addr string) (func(), error) {
	c, err := p.s.clientPool.GetContext(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("connect: DUT unreachable, %s", err)
	}

	p.c = c
	disconnect := func(c *ssh.Client) func() {
		return func() {
			p.s.clientPool.Put(addr, c)
		}
	}(p.c)
	return disconnect, nil
}

func (p *provisionState) shouldProvisionOS() bool {
	// Get the current builder path.
	// If the builder path is missing or fails to be retrieved, continue to provision.
	builderPath, err := getBuilderPath(p.c)
	if err != nil {
		log.Printf("provision: failed to get pre-provision builder path, %s", err)
		return true
	}
	// Only provision the OS if any of the following are true:
	//  - the DUT is not on the requested OS.
	//  - the force marker exists on the DUT.
	shouldProvision := false
	if builderPath != p.targetBuilderPath {
		log.Printf("Going to provision DUT from %s to %s", builderPath, p.targetBuilderPath)
		shouldProvision = true
	}
	if shouldForceProvision(p.c) {
		if !shouldProvision {
			log.Printf("Going to force provision to %s", p.targetBuilderPath)
			shouldProvision = true
		} else {
			log.Printf("Ignoring force provision as already provisioning")
		}
	}
	return shouldProvision
}

// provisionOS will provision the OS, but on failure it will set op.Result to longrunning.Operation_Error
// and return the same error message
func (p *provisionState) provisionOS(ctx context.Context) error {
	r, err := getRootDev(p.c)
	if err != nil {
		return fmt.Errorf("provisionOS: failed to get root device from DUT, %s", err)
	}

	stopSystemDaemons(p.c)

	// Only clear the inactive verified DLC marks if the DLCs exist.
	dlcsExist, err := pathExists(p.c, dlcLibDir)
	if err != nil {
		return fmt.Errorf("provisionOS: failed to check if DLC is enabled, %s", err)
	}
	if dlcsExist {
		if err := clearInactiveDLCVerifiedMarks(p.c, r); err != nil {
			return fmt.Errorf("provisionOS: failed to clear inactive verified DLC marks, %s", err)
		}
	}

	pi := getPartitionInfo(r)
	if err := p.installPartitions(ctx, pi); err != nil {
		return fmt.Errorf("provisionOS: failed to provision the OS, %s", err)
	}
	if err := p.postInstall(pi); err != nil {
		p.revertProvisionOS(pi)
		return fmt.Errorf("provisionOS: failed to set next kernel, %s", err)
	}
	if err := clearTPM(p.c); err != nil {
		p.revertProvisionOS(pi)
		return fmt.Errorf("provisionOS: failed to clear TPM owner, %s", err)
	}
	return nil
}

func (p *provisionState) wipeStateful(ctx context.Context) error {
	if err := runCmd(p.c, "echo 'fast keepimg' > /mnt/stateful_partition/factory_install_reset"); err != nil {
		return fmt.Errorf("wipeStateful: Failed to to write to factory reset file, %s", err)
	}

	runLabMachineAutoReboot(p.c)
	if err := rebootDUT(ctx, p.c); err != nil {
		return fmt.Errorf("wipeStateful: failed to reboot DUT, %s", err)
	}
	return nil
}

func (p *provisionState) provisionStateful(ctx context.Context) error {
	stopSystemDaemons(p.c)

	if err := p.installStateful(ctx); err != nil {
		p.revertStatefulInstall()
		return fmt.Errorf("provisionStateful: falied to install stateful partition, %s", err)
	}

	runLabMachineAutoReboot(p.c)
	if err := rebootDUT(ctx, p.c); err != nil {
		return fmt.Errorf("provisionStateful: failed to reboot DUT, %s", err)
	}
	return nil
}

func (p *provisionState) verifyOSProvision() error {
	actualBuilderPath, err := getBuilderPath(p.c)
	if err != nil {
		return fmt.Errorf("verify OS provision: failed to get builder path, %s", err)
	}
	if actualBuilderPath != p.targetBuilderPath {
		return fmt.Errorf("verify OS provision: DUT is on builder path %s when expected builder path is %s",
			actualBuilderPath, p.targetBuilderPath)
	}
	return nil
}

const (
	fetchUngzipConvertCmd = `curl -S -s -v -# -C - --retry 3 --retry-delay 60 %[1]s | gzip -d | dd of=%[2]s obs=2M
pipestatus=("${PIPESTATUS[@]}")
if [[ "${pipestatus[0]}" -ne 0 ]]; then
  echo "$(date --rfc-3339=seconds) ERROR: Fetching %[1]s failed." >&2
  exit 1
elif [[ "${pipestatus[1]}" -ne 0 ]]; then
  echo "$(date --rfc-3339=seconds) ERROR: Decompressing %[1]s failed." >&2
  exit 1
elif [[ "${pipestatus[2]}" -ne 0 ]]; then
  echo "$(date --rfc-3339=seconds) ERROR: Writing to %[2]s failed." >&2
  exit 1
fi`
)

// installKernel updates kernelPartition on disk.
func (p *provisionState) installKernel(ctx context.Context, pi partitionInfo) error {
	url, err := p.s.cacheForDut(ctx, path.Join(p.imagePath, "full_dev_part_KERN.bin.gz"), p.dutName)
	if err != nil {
		return fmt.Errorf("install kernel: failed to get GS Cache URL, %s", err)
	}
	return runCmd(p.c, fmt.Sprintf(fetchUngzipConvertCmd, url, pi.inactiveKernel))
}

// installRoot updates rootPartition on disk.
func (p *provisionState) installRoot(ctx context.Context, pi partitionInfo) error {
	url, err := p.s.cacheForDut(ctx, path.Join(p.imagePath, "full_dev_part_ROOT.bin.gz"), p.dutName)
	if err != nil {
		return fmt.Errorf("install root: failed to get GS Cache URL, %s", err)
	}
	return runCmd(p.c, fmt.Sprintf(fetchUngzipConvertCmd, url, pi.inactiveRoot))
}

const (
	statefulPath           = "/mnt/stateful_partition"
	updateStatefulFilePath = statefulPath + "/.update_available"
)

// installStateful updates the stateful partition on disk (finalized after a reboot).
func (p *provisionState) installStateful(ctx context.Context) error {
	url, err := p.s.cacheForDut(ctx, path.Join(p.imagePath, "stateful.tgz"), p.dutName)
	if err != nil {
		return fmt.Errorf("install stateful: failed to get GS Cache URL, %s", err)
	}
	return runCmd(p.c, strings.Join([]string{
		fmt.Sprintf("rm -rf %[1]s %[2]s/var_new %[2]s/dev_image_new", updateStatefulFilePath, statefulPath),
		fmt.Sprintf("curl -v -# -C - --retry 3 --retry-delay 60 %s | tar --ignore-command-error --overwrite --directory=%s -xzf -", url, statefulPath),
		fmt.Sprintf("echo -n clobber > %s", updateStatefulFilePath),
	}, " && "))
}

func (p *provisionState) installPartitions(ctx context.Context, pi partitionInfo) error {
	if err := p.installKernel(ctx, pi); err != nil {
		log.Printf("installPartitions: failed to install kernel.")
		return err
	}
	if err := p.installRoot(ctx, pi); err != nil {
		log.Printf("installPartitions: failed to install rootfs.")
		return err
	}
	return nil
}

func (p *provisionState) postInstall(pi partitionInfo) error {
	return runCmd(p.c, strings.Join([]string{
		"tmpmnt=$(mktemp -d)",
		fmt.Sprintf("mount -o ro %s ${tmpmnt}", pi.inactiveRoot),
		fmt.Sprintf("${tmpmnt}/postinst %s", pi.inactiveRoot),
		"{ umount ${tmpmnt} || true; }",
		"{ rmdir ${tmpmnt} || true; }",
	}, " && "))
}

func (p *provisionState) revertProvisionOS(pi partitionInfo) {
	p.revertStatefulInstall()
	p.revertPostInstall(pi)
}

func (p *provisionState) revertStatefulInstall() {
	const (
		varNew      = "var_new"
		devImageNew = "dev_image_new"
	)
	varNewPath := path.Join(statefulPath, varNew)
	devImageNewPath := path.Join(statefulPath, devImageNew)
	err := runCmd(p.c, fmt.Sprintf("rm -rf %s %s %s", varNewPath, devImageNewPath, updateStatefulFilePath))
	if err != nil {
		log.Printf("revert stateful install: failed to revert stateful installation, %s", err)
	}
}

func (p *provisionState) revertPostInstall(pi partitionInfo) {
	if err := runCmd(p.c, fmt.Sprintf("/postinst %s 2>&1", pi.activeRoot)); err != nil {
		log.Printf("revert post install: failed to revert postinst, %s", err)
	}
}

func (p *provisionState) provisionDLCs(ctx context.Context, specs []*tls.ProvisionDutRequest_DLCSpec) error {
	if len(specs) == 0 {
		return nil
	}

	// Stop dlcservice daemon in order to not interfere with provisioning DLCs.
	if err := runCmd(p.c, "stop dlcservice"); err != nil {
		log.Printf("provision DLCs: failed to stop dlcservice daemon, %s", err)
	}
	defer func() {
		if err := runCmd(p.c, "start dlcservice"); err != nil {
			log.Printf("provision DLCs: failed to start dlcservice daemon, %s", err)
		}
	}()

	var err error
	r, err := getRootDev(p.c)
	if err != nil {
		return fmt.Errorf("provision DLCs: failed to get root device from DUT, %s", err)
	}

	// TODO(kimjae): Can parallelize, once outputs can be sorted.
	for _, spec := range specs {
		dlcID := spec.GetId()
		dlcOutputDir := path.Join(dlcCacheDir, dlcID, dlcPackage)
		if err := p.installDLC(ctx, spec, dlcOutputDir, getActiveDLCSlot(r)); err != nil {
			return fmt.Errorf("provision DLCs: failed to install the following DLC %s (%s)", dlcID, err)
		}
	}

	return nil
}

func (p *provisionState) installDLC(ctx context.Context, spec *tls.ProvisionDutRequest_DLCSpec, dlcOutputDir string, slot dlcSlot) error {
	verified, err := isDLCVerified(p.c, spec, slot)
	if err != nil {
		return fmt.Errorf("install DLC: failed is DLC verified check, %s", err)
	}

	dlcID := spec.GetId()
	// Skip installing the DLC if already verified.
	if verified {
		log.Printf("Provision DLC %s skipped as already verified", dlcID)
		return nil
	}

	dlcURL := path.Join(p.imagePath, "dlc", dlcID, dlcPackage, dlcImage)
	url, err := p.s.cacheForDut(ctx, dlcURL, p.dutName)
	if err != nil {
		return fmt.Errorf("install DLC: failed to get GS Cache server, %s", err)
	}

	dlcOutputSlotDir := path.Join(dlcOutputDir, string(slot))
	dlcOutputImage := path.Join(dlcOutputSlotDir, dlcImage)
	dlcCmd := fmt.Sprintf(`mkdir -p %[1]s && curl -v -# -C - --retry 3 --retry-delay 60 --output %[2]s %[3]s`,
		dlcOutputSlotDir, dlcOutputImage, url)
	if err := runCmd(p.c, dlcCmd); err != nil {
		return fmt.Errorf("provision DLC: failed to provision DLC %s, %s", dlcID, err)
	}
	return nil
}
