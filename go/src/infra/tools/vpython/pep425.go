// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package vpython

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/vpython/api/vpython"
	"go.chromium.org/luci/vpython/cipd"
)

// pep425MacPlatform is a parsed PEP425 Mac platform string.
//
// The string is formatted:
// macosx_<maj>_<min>_<cpu-arch>
//
// For example:
//	- macosx_10_6_intel
//	- macosx_10_0_fat
//	- macosx_10_2_x86_64
type pep425MacPlatform struct {
	major int
	minor int
	arch  string
}

// parsePEP425MacPlatform parses a pep425MacPlatform from the supplied
// platform string. If the string does not contain a recognizable Mac
// platform, this function returns nil.
func parsePEP425MacPlatform(v string) *pep425MacPlatform {
	parts := strings.SplitN(v, "_", 4)
	if len(parts) != 4 {
		return nil
	}
	if parts[0] != "macosx" {
		return nil
	}

	var ma pep425MacPlatform
	var err error
	if ma.major, err = strconv.Atoi(parts[1]); err != nil {
		return nil
	}
	if ma.minor, err = strconv.Atoi(parts[2]); err != nil {
		return nil
	}

	ma.arch = parts[3]
	return &ma
}

// less returns true if "ma" represents a Mac version before "other".
func (ma *pep425MacPlatform) less(other *pep425MacPlatform) bool {
	switch {
	case ma.major < other.major:
		return true
	case ma.major > other.major:
		return false
	case ma.minor < other.minor:
		return true
	default:
		return false
	}
}

// pep425IsBetterMacPlatform processes two PEP425 platform strings and
// returns true if "candidate" is a superior PEP425 tag candidate than "cur".
//
// This function favors, in order:
//	- Mac platforms over non-Mac platforms,
//	- "intel" package builds over non-"intel"
//	- Older Mac versions over newer ones
func pep425IsBetterMacPlatform(cur, candidate string) bool {
	// Parse a Mac platform string
	curPlatform := parsePEP425MacPlatform(cur)
	candidatePLatform := parsePEP425MacPlatform(candidate)
	switch {
	case curPlatform == nil:
		return candidatePLatform != nil
	case candidatePLatform == nil:
		return false
	case curPlatform.arch != "intel" && candidatePLatform.arch == "intel":
		// Prefer "intel" architecture over others, since it's more modern and
		// generic.
		return true
	case curPlatform.arch == "intel" && candidatePLatform.arch != "intel":
		return false
	case candidatePLatform.less(curPlatform):
		// We prefer the lowest Mac architecture available.
		return true
	default:
		return false
	}
}

// Determies if the specified platform is a Linux platform and, if so, if it
// is a "manylinux1_" Linux platform.
func isLinuxPlatform(plat string) (is bool, many bool) {
	switch {
	case strings.HasPrefix(plat, "linux_"):
		is = true
	case strings.HasPrefix(plat, "manylinux1_"):
		is, many = true, true
	}
	return
}

// pep425IsBetterLinuxPlatform processes two PEP425 platform strings and
// returns true if "candidate" is a superior PEP425 tag candidate than "cur".
//
// This function favors, in order:
//	- Linux platforms over non-Linux platforms.
//	- "manylinux1_" over non-"manylinux1_".
//
// Examples of expected Linux platform strings are:
//	- linux1_x86_64
//	- linux1_i686
//	- manylinux1_i686
func pep425IsBetterLinuxPlatform(cur, candidate string) bool {
	// We prefer "manylinux1_" platforms over "linux_" platforms.
	curIs, curMany := isLinuxPlatform(cur)
	candidateIs, candidateMany := isLinuxPlatform(candidate)
	switch {
	case !curIs:
		return candidateIs
	case !candidateIs:
		return false
	case curMany:
		return false
	default:
		return candidateMany
	}
}

// preferredPlatformFuncForTagSet examines a tag set and returns a function
// that compares two "platform" tags.
//
// The comparison function is chosen based on the operating system represented
// by the tag set. This choice is made with the assumption that the tag set
// represents a realistic platform (e.g., no mixed Mac and Linux tags).
func preferredPlatformFuncForTagSet(tags []*vpython.PEP425Tag) func(cur, candidate string) bool {
	// Identify the operating system from the tag set. Iterate through tags until
	// we see an indicator.
	for _, tag := range tags {
		// Linux?
		if is, _ := isLinuxPlatform(tag.Platform); is {
			return pep425IsBetterLinuxPlatform
		}

		// Mac
		if plat := parsePEP425MacPlatform(tag.Platform); plat != nil {
			return pep425IsBetterMacPlatform
		}
	}

	// No opinion.
	return func(cur, candidate string) bool { return false }
}

// isNewerPy3Abi returns true if the candidate string identifies a new, unstable
// ABI that should be preferred over the long-term stable "abi3", which we don't
// build wheels against.
func isNewerPy3Abi(cur, candidate string) bool {
	// We don't bother finding the latest ABI (e.g. preferring "cp39" over
	// "cp38"). Each release only has one supported unstable ABI, so we should
	// never encounter more than one anyway.
	return (cur == "abi3" || cur == "none") && strings.HasPrefix(candidate, "cp3")
}

// Prefer specific Python (e.g., cp27) over generic (e.g., py27).
func isSpecificImplAbi(python string) bool {
	return !strings.HasPrefix(python, "py")
}

// pep425TagSelector chooses the "best" PEP425 tag from a set of potential tags.
// This "best" tag will be used to resolve our CIPD templates and allow for
// Python implementation-specific CIPD template parameters.
func pep425TagSelector(tags []*vpython.PEP425Tag) *vpython.PEP425Tag {
	var best *vpython.PEP425Tag

	// isPreferredOSPlatform is an OS-specific platform preference function.
	isPreferredOSPlatform := preferredPlatformFuncForTagSet(tags)

	isBetter := func(t *vpython.PEP425Tag) bool {
		switch {
		case best == nil:
			return true
		case t.Count() > best.Count():
			// More populated fields is more specificity.
			return true
		case best.AnyPlatform() && !t.AnyPlatform():
			// More specific platform is preferred.
			return true
		case !best.HasABI() && t.HasABI():
			// More specific ABI is preferred.
			return true
		case isNewerPy3Abi(best.Abi, t.Abi):
			// Prefer the newest supported ABI tag. In theory this can break if
			// we have wheels built against a long-term stable ABI like abi3, as
			// we'll only look for packages built against the newest, unstable
			// ABI. But in practice that doesn't happen, as dockerbuild
			// produces packages tagged with the unstable ABIs.
			return true
		case isPreferredOSPlatform(best.Platform, t.Platform) && (isSpecificImplAbi(t.Python) || !isSpecificImplAbi(best.Python)):
			// Prefer a better platform, but not if it means moving
			// to a less-specific ABI.
			return true
		case isSpecificImplAbi(t.Python) && !isSpecificImplAbi(best.Python):
			return true

		default:
			return false
		}
	}

	for _, t := range tags {
		tag := proto.Clone(t).(*vpython.PEP425Tag)

		// pip has a bug, fixed in 19.2.3 where it adds the "m" ABI tag on 3.8
		// for some platforms, despite this flag being removed in 3.8. We work
		// around it here by stripping it from the ABI string.
		//
		// TODO: Remove this workaround when we've updated to pip >= 19.2.3.
		if strings.HasPrefix(tag.Python, "cp3") {
			tag.Abi = strings.TrimSuffix(tag.Abi, "m")
		}
		if isBetter(tag) {
			best = tag
		}
	}
	return best
}

// getPEP425CIPDTemplates returns the set of CIPD template strings for a
// given PEP425 tag.
//
// Template parameters are derived from the most representative PEP425 tag.
// Any missing tag parameters will result in their associated template
// parameters not getting exported.
//
// The full set of exported tag parameters is:
// - py_python: The PEP425 "python" tag value (e.g., "cp27").
// - py_abi: The PEP425 Python ABI (e.g., "cp27mu").
// - py_platform: The PEP425 Python platform (e.g., "manylinux1_x86_64").
// - py_tag: The full PEP425 tag (e.g., "cp27-cp27mu-manylinux1_x86_64").
//
// This function also backports the Python platform into the CIPD "platform"
// field, ensuring that regardless of the host platform, the Python CIPD
// wheel is chosen based solely on that host's Python interpreter.
//
// Infra CIPD packages tend to use "${platform}" (generic) combined with
// "${py_abi}" and "${py_platform}" to identify its packages.
func getPEP425CIPDTemplateForTag(tag *vpython.PEP425Tag) (map[string]string, error) {
	if tag == nil {
		return nil, errors.New("no PEP425 tag")
	}

	template := make(map[string]string, 4)
	if tag.Python != "" {
		template["py_python"] = tag.Python

		// TODO(dnj): Remove this once everything has been updated to "py_python".
		template["py_version"] = tag.Python
	}
	if tag.Abi != "" {
		template["py_abi"] = tag.Abi
	}
	if tag.Platform != "" {
		template["py_platform"] = tag.Platform

		// TODO(dnj): Remove this once everything has been updated to "py_platform".
		template["py_arch"] = tag.Platform
	}
	if tag.Python != "" && tag.Abi != "" && tag.Platform != "" {
		template["py_tag"] = tag.TagString()
	}

	// Override the CIPD "platform" based on the PEP425 tag. This allows selection
	// of Python wheels based on the platform of the Python executable rather
	// than the platform of the underlying operating system.
	//
	// For example, a 64-bit Windows version can run 32-bit Python, and we'll
	// want to use 32-bit Python wheels.
	platform := cipd.PlatformForPEP425Tag(tag)
	if platform == "" {
		return nil, errors.Reason("failed to infer CIPD platform for tag [%s]", tag).Err()
	}
	template["platform"] = platform

	// Build the sum tag, "vpython_platform",
	// "${platform}_${py_python}_${py_abi}"
	if platform != "" && tag.Python != "" && tag.Abi != "" {
		template["vpython_platform"] = fmt.Sprintf("%s_%s_%s", platform, tag.Python, tag.Abi)
	}

	return template, nil
}
