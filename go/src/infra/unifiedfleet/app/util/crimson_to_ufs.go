// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package util

import (
	"fmt"
	"strings"

	fleet "infra/unifiedfleet/api/v1/proto"

	crimsoncommon "go.chromium.org/luci/machine-db/api/common/v1"
	crimsonconfig "go.chromium.org/luci/machine-db/api/config/v1"
	"go.chromium.org/luci/machine-db/api/crimson/v1"
)

// ToChromeMachines converts crimson machines to UFS format.
func ToChromeMachines(old []*crimson.Machine, machineToNics map[string][]string, machineToDracs map[string]string) []*fleet.Machine {
	newObjects := make([]*fleet.Machine, len(old))
	for i, o := range old {
		newObjects[i] = &fleet.Machine{
			// Temporarily use existing display name as browser machine's name instead of serial number/assettag
			Name:     o.Name,
			Location: toLocation(o.Rack, o.Datacenter),
			Device: &fleet.Machine_ChromeBrowserMachine{
				ChromeBrowserMachine: &fleet.ChromeBrowserMachine{
					// RpmInterface is not available for browser machine.
					// KvmInterface is currently attached to rack.
					// NetworkDeviceInterface is attached to the nics.
					DisplayName:      o.Name,
					ChromePlatform:   FormatResourceName(o.Platform),
					Nics:             machineToNics[o.Name],
					Drac:             machineToDracs[o.Name],
					DeploymentTicket: o.DeploymentTicket,
				},
			},
			Realm: BrowserLabAdminRealm,
		}
	}
	return newObjects
}

func toLocation(rack, datacenter string) *fleet.Location {
	return &fleet.Location{
		Rack: rack,
		Lab:  ToLab(strings.ToLower(datacenter)),
	}
}

// ToChromePlatforms converts platforms in static file to UFS format.
func ToChromePlatforms(oldP *crimsonconfig.Platforms) []*fleet.ChromePlatform {
	ps := oldP.GetPlatform()
	newP := make([]*fleet.ChromePlatform, len(ps))
	for i, p := range ps {
		newP[i] = &fleet.ChromePlatform{
			Name:         FormatResourceName(p.GetName()),
			Manufacturer: p.GetManufacturer(),
			Description:  p.GetDescription(),
		}
	}
	return newP
}

// ToOses converts the os versions to UFS format.
func ToOses(old []*crimson.OS) []*fleet.OSVersion {
	newOSes := make([]*fleet.OSVersion, len(old))
	for i, p := range old {
		newOSes[i] = &fleet.OSVersion{
			Value:       FormatResourceName(p.GetName()),
			Description: p.GetDescription(),
		}
	}
	return newOSes
}

// ProcessDatacenters converts datacenters to several UFS objects
func ProcessDatacenters(dc *crimsonconfig.Datacenter) ([]*fleet.Rack, []*fleet.RackLSE, []*fleet.KVM, []*fleet.Switch, []*fleet.DHCPConfig) {
	dcName := dc.GetName()
	switches := make([]*fleet.Switch, 0)
	racks := make([]*fleet.Rack, 0)
	rackLSEs := make([]*fleet.RackLSE, 0)
	rackToKvms := make(map[string][]string, 0)
	kvms := make([]*fleet.KVM, 0)
	dhcps := make([]*fleet.DHCPConfig, 0)
	for _, oldKVM := range dc.GetKvm() {
		name := oldKVM.GetName()
		k := &fleet.KVM{
			Name:           name,
			MacAddress:     oldKVM.GetMacAddress(),
			ChromePlatform: FormatResourceName(oldKVM.GetPlatform()),
		}
		kvms = append(kvms, k)
		rackName := oldKVM.GetRack()
		rackToKvms[rackName] = append(rackToKvms[rackName], name)
		dhcps = append(dhcps, &fleet.DHCPConfig{
			MacAddress: oldKVM.GetMacAddress(),
			Hostname:   name,
			Ip:         oldKVM.GetIpv4(),
		})
	}
	for _, old := range dc.GetRack() {
		rackName := old.GetName()
		switchNames := make([]string, 0)
		for _, crimsonSwitch := range old.GetSwitch() {
			s := &fleet.Switch{
				Name:         crimsonSwitch.GetName(),
				CapacityPort: crimsonSwitch.GetPorts(),
				Description:  crimsonSwitch.GetDescription(),
			}
			switches = append(switches, s)
			switchNames = append(switchNames, s.GetName())
		}
		// Also add the kvms which is attached to the rack in rack definitation
		found := false
		for _, rack := range rackToKvms[rackName] {
			if rack == old.GetKvm() {
				found = true
				break
			}
		}
		if !found {
			rackToKvms[rackName] = append(rackToKvms[rackName], old.GetKvm())
		}
		r := &fleet.Rack{
			Name:     rackName,
			Location: toLocation(rackName, dcName),
			Rack: &fleet.Rack_ChromeBrowserRack{
				ChromeBrowserRack: &fleet.ChromeBrowserRack{
					Kvms:     rackToKvms[rackName],
					Switches: switchNames,
				},
			},
		}
		rlse := &fleet.RackLSE{
			Name:             GetRackHostname(rackName),
			RackLsePrototype: "browser-lab:normal",
			Lse: &fleet.RackLSE_ChromeBrowserRackLse{
				ChromeBrowserRackLse: &fleet.ChromeBrowserRackLSE{
					// Still keep them as they are potential hostnames
					Kvms:     rackToKvms[rackName],
					Switches: switchNames,
				},
			},
			Racks: []string{rackName},
		}
		racks = append(racks, r)
		rackLSEs = append(rackLSEs, rlse)
	}
	return racks, rackLSEs, kvms, switches, dhcps
}

// ProcessNetworkInterfaces converts nics and dracs to several UFS formats for further importing
func ProcessNetworkInterfaces(nics []*crimson.NIC, dracs []*crimson.DRAC, machines []*crimson.Machine) ([]*fleet.Nic, []*fleet.Drac, []*fleet.DHCPConfig, map[string][]string, map[string]string) {
	machineToNics := make(map[string][]string, 0)
	machineToDracs := make(map[string]string, 0)
	machineMap := make(map[string]*crimson.Machine, len(machines))
	newNics := make([]*fleet.Nic, 0)
	newDracs := make([]*fleet.Drac, 0)
	dhcps := make([]*fleet.DHCPConfig, 0)
	for _, machine := range machines {
		machineMap[machine.GetName()] = machine
	}
	for _, nic := range nics {
		name := GetNicName(nic.GetName(), nic.GetMachine())
		switch nic.GetName() {
		case "drac":
			// Use ListDrac() as the source of truth for drac
			continue
		default:
			// lab and rack are for indexing nic table
			var rack string
			var lab string
			machine, ok := machineMap[nic.GetMachine()]
			if ok {
				rack = machine.GetRack()
				lab = ToLab(strings.ToLower(machine.GetDatacenter())).String()
			}
			// Multiple nic names, e.g. eth0, eth1, bmc
			newNic := &fleet.Nic{
				Name:       name,
				MacAddress: nic.GetMacAddress(),
				SwitchInterface: &fleet.SwitchInterface{
					Switch: nic.GetSwitch(),
					Port:   nic.GetSwitchport(),
				},
				Rack: rack,
				Lab:  lab,
			}
			newNics = append(newNics, newNic)
			machineToNics[nic.GetMachine()] = append(machineToNics[nic.GetMachine()], name)
		}
		if ip := nic.GetIpv4(); ip != "" {
			dhcps = append(dhcps, &fleet.DHCPConfig{
				MacAddress: nic.GetMacAddress(),
				Hostname:   name,
				Ip:         nic.GetIpv4(),
			})
		}
	}
	for _, drac := range dracs {
		// lab and rack are for indexing drac table
		var rack string
		var lab string
		machine, ok := machineMap[drac.GetMachine()]
		if ok {
			rack = machine.GetRack()
			lab = ToLab(strings.ToLower(machine.GetDatacenter())).String()
		}
		hostname := FormatResourceName(drac.GetName())
		d := &fleet.Drac{
			Name: hostname,
			// Inject machine name to display name
			DisplayName: GetNicName("drac", drac.GetMachine()),
			MacAddress:  drac.GetMacAddress(),
			SwitchInterface: &fleet.SwitchInterface{
				Switch: drac.GetSwitch(),
				Port:   drac.GetSwitchport(),
			},
			Rack: rack,
			Lab:  lab,
		}
		newDracs = append(newDracs, d)
		machineToDracs[drac.GetMachine()] = hostname
		if ip := drac.GetIpv4(); ip != "" {
			dhcps = append(dhcps, &fleet.DHCPConfig{
				MacAddress: drac.GetMacAddress(),
				Hostname:   hostname,
				Ip:         drac.GetIpv4(),
			})
		}
	}
	return newNics, newDracs, dhcps, machineToNics, machineToDracs
}

// ToMachineLSEs converts crimson data to UFS LSEs.
func ToMachineLSEs(hosts []*crimson.PhysicalHost, vms []*crimson.VM) ([]*fleet.MachineLSE, []*fleet.IP, []*fleet.DHCPConfig) {
	hostToVMs := make(map[string][]*fleet.VM, 0)
	ips := make([]*fleet.IP, 0)
	dhcps := make([]*fleet.DHCPConfig, 0)
	for _, vm := range vms {
		name := vm.GetName()
		v := &fleet.VM{
			Name: name,
			OsVersion: &fleet.OSVersion{
				Value: vm.GetOs(),
			},
			Hostname: name,
		}
		hostToVMs[vm.GetHost()] = append(hostToVMs[vm.GetHost()], v)
		ip := FormatIP(vm.GetVlan(), vm.GetIpv4(), true)
		if ip != nil {
			ips = append(ips, ip)
		}
		dhcps = append(dhcps, &fleet.DHCPConfig{
			Hostname: v.GetHostname(),
			Ip:       vm.GetIpv4(),
			// No mac address found
		})
	}
	lses := make([]*fleet.MachineLSE, 0)
	var lsePrototype string
	for _, h := range hosts {
		name := h.GetName()
		vms := hostToVMs[name]
		if len(vms) > 0 {
			lsePrototype = "browser-lab:vm"
		} else {
			lsePrototype = "browser-lab:no-vm"
		}
		lse := &fleet.MachineLSE{
			Name:                name,
			MachineLsePrototype: lsePrototype,
			Hostname:            name,
			Machines:            []string{h.GetMachine()},
			Lse: &fleet.MachineLSE_ChromeBrowserMachineLse{
				ChromeBrowserMachineLse: &fleet.ChromeBrowserMachineLSE{
					Vms:        vms,
					VmCapacity: h.GetVmSlots(),
					OsVersion: &fleet.OSVersion{
						Value: h.GetOs(),
					},
				},
			},
		}
		lses = append(lses, lse)
		ip := FormatIP(h.GetVlan(), h.GetIpv4(), true)
		if ip != nil {
			ips = append(ips, ip)
		}
		dhcps = append(dhcps, &fleet.DHCPConfig{
			Hostname:   h.GetName(),
			Ip:         h.GetIpv4(),
			MacAddress: h.GetMacAddress(),
		})
	}
	return lses, ips, dhcps
}

// ToState converts crimson state to UFS state.
func ToState(state crimsoncommon.State) fleet.State {
	switch state {
	case crimsoncommon.State_SERVING:
		return fleet.State_STATE_SERVING
	case crimsoncommon.State_DECOMMISSIONED:
		return fleet.State_STATE_DECOMMISSIONED
	case crimsoncommon.State_REPAIR:
		return fleet.State_STATE_NEEDS_REPAIR
	case crimsoncommon.State_TEST:
		return fleet.State_STATE_DEPLOYED_TESTING
	case crimsoncommon.State_PRERELEASE:
		return fleet.State_STATE_DEPLOYED_PRE_SERVING
	case crimsoncommon.State_FREE:
		return fleet.State_STATE_REGISTERED
	}
	return fleet.State_STATE_UNSPECIFIED
}

// ToLab converts the crimson lab string to UFS lab.
func ToLab(datacenter string) fleet.Lab {
	switch strings.ToLower(datacenter) {
	case "atl97":
		return fleet.Lab_LAB_DATACENTER_ATL97
	case "iad97":
		return fleet.Lab_LAB_DATACENTER_IAD97
	case "mtv96":
		return fleet.Lab_LAB_DATACENTER_MTV96
	case "mtv97":
		return fleet.Lab_LAB_DATACENTER_MTV97
	case "lab01":
		return fleet.Lab_LAB_DATACENTER_FUCHSIA
	default:
		return fleet.Lab_LAB_UNSPECIFIED
	}
}

// GetNicName formats a nic name with its attached machine
func GetNicName(nicName, machineName string) string {
	return fmt.Sprintf("%s-%s", machineName, nicName)
}
