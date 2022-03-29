// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package config

import (
	"log"

	"google.golang.org/protobuf/types/known/durationpb"
)

func crosDeployPlan() *Plan {
	return &Plan{
		CriticalActions: []string{
			"Set needs_deploy state",
			"Clean up",
			"Servo has USB-key with require image",
			"Device is pingable before deploy",
			"DUT has expected OS version",
			"DUT has expected dev firmware",
			"Switch to secure-mode and reboot",
			"Collect DUT labels",
			"Deployment checks",
			"DUT verify",
		},
		Actions: crosDeployAndRepairActions(),
	}
}

func deployActions() map[string]*Action {
	return map[string]*Action{
		"DUT is in dev-mode and allowed to boot from USB-key": {
			Docs:        []string{"Verify that device is set to boot in DEV mode and enabled to boot from USB-drive."},
			ExecTimeout: &durationpb.Duration{Seconds: 2000},
			ExecName:    "cros_read_gbb_by_servo",
			ExecExtraArgs: []string{
				"validate_in_dev_mode:true",
				"validate_usb_boot_enabled:true",
				"remove_file:false",
			},
			RecoveryActions: []string{"Set GBB flags to 0x18 by servo"},
		},
		"Device is pingable before deploy": {
			Docs: []string{
				"Verify that device is present in setup.",
				"All devices is pingable by default even they have prod images on them.",
				"If device is not pingable then device is off on not connected",
			},
			ExecName:    "cros_ping",
			ExecTimeout: &durationpb.Duration{Seconds: 15},
			RecoveryActions: []string{
				"Cold reset DUT by servo and wait to boot",
				"Power cycle DUT by RPM and wait",
				"Set GBB flags to 0x18 by servo",
				"Install OS in DEV mode",
			},
		},
		"DUT has expected OS version": {
			Docs: []string{"Verify that device has stable version OS on it and version is match."},
			Dependencies: []string{
				"Device is pingable before deploy",
				"has_stable_version_cros_image",
				"Device NOT booted from USB-drive",
			},
			ExecName: "cros_is_on_stable_version",
			RecoveryActions: []string{
				"Quick provision OS",
				"Install OS in DEV mode",
			},
		},
		"DUT has expected dev firmware": {
			Docs:         []string{"Verify that FW on the DUT has dev keys."},
			Dependencies: []string{"cros_ssh"},
			ExecName:     "cros_has_dev_signed_firmware",
			ExecTimeout:  &durationpb.Duration{Seconds: 600},
			RecoveryActions: []string{
				"Update DUT firmware with factory mode and restart by servo",
				"Update DUT firmware with factory mode and restart by host",
			},
		},
		"Update DUT firmware with factory mode and restart by servo": {
			Docs: []string{
				"Force update FW on the DUT by factory mode.",
				"Reboot device by servo",
			},
			Conditions: []string{"servo_state_is_working"},
			Dependencies: []string{
				"cros_ssh",
				"Disable software-controlled write-protect for 'host'",
				"Disable software-controlled write-protect for 'ec'",
			},
			ExecTimeout: &durationpb.Duration{Seconds: 900},
			ExecName:    "cros_run_firmware_update",
			ExecExtraArgs: []string{
				"mode:factory",
				"force:true",
				"reboot:by_servo",
				"updater_timeout:600",
			},
		},
		"Update DUT firmware with factory mode and restart by host": {
			Docs: []string{
				"Force update FW on the DUT by factory mode.",
				"Reboot device by host",
			},
			Conditions: []string{"servo_state_is_not_working"},
			Dependencies: []string{
				"cros_ssh",
				"Disable software-controlled write-protect for 'host'",
				"Disable software-controlled write-protect for 'ec'",
			},
			ExecTimeout: &durationpb.Duration{Seconds: 900},
			ExecName:    "cros_run_firmware_update",
			ExecExtraArgs: []string{
				"mode:factory",
				"force:true",
				"reboot:by_host",
				"updater_timeout:600",
			},
		},
		"Deployment checks": {
			Docs: []string{"Run some specif checks as part of deployment."},
			Dependencies: []string{
				"Verify battery charging level",
				"Verify RPM config (without battery)",
				"Verify boot in recovery mode",
			},
			ExecName: "sample_pass",
		},
		"Verify battery charging level": {
			Docs: []string{
				"Battery will be checked that it can be charged to the 80% as if device cannot then probably device is not fully prepared for deployment.",
				"If battery is not charged, then we will re-check every 15 minutes for 8 time to allows to charge the battery.",
				"Dues overheat battery in audio boxes mostly it deployed ",
			},
			Conditions: []string{
				"Is not in audio box",
				"Battery is expected on device",
				"Battery is present on device",
			},
			ExecName: "cros_battery_changable_to_expected_level",
			ExecExtraArgs: []string{
				"charge_retry_count:8",
				"charge_retry_interval:900",
			},
			ExecTimeout: &durationpb.Duration{Seconds: 9000},
		},
		"Verify boot in recovery mode": {
			Docs: []string{
				"Devices deployed with servo in the pools required secure mode need to be able to be boot in recovery mode.",
			},
			Conditions: []string{
				"Pools required to be in Secure mode",
			},
			Dependencies: []string{
				"dut_servo_host_present",
				"servo_state_is_working",
			},
			ExecName: "cros_verify_boot_in_recovery_mode",
			ExecExtraArgs: []string{
				"boot_timeout:480",
				"boot_interval:10",
				"halt_timeout:120",
				"ignore_reboot_failure:false",
			},
			ExecTimeout: &durationpb.Duration{Seconds: 900},
			RecoveryActions: []string{
				// The only reason why it can fail on good DUT is that USB-key has not good image.
				"Download stable image to USB-key",
			},
		},
		"Verify RPM config (without battery)": {
			Docs: []string{
				"Verify RPM configs and set RPM state",
				"Not applicable for cr50 servos based on b/205728276",
				"Action is not critical as it updates own state.",
			},
			Conditions: []string{
				"dut_servo_host_present",
				"servo_state_is_working",
				"is_servo_main_ccd_cr50",
				"has_rpm_info",
				"No Battery is present on device",
			},
			ExecName:               "rpm_audit",
			ExecTimeout:            &durationpb.Duration{Seconds: 600},
			AllowFailAfterRecovery: true,
		},
		"DUT verify": {
			Docs:         []string{"Run all repair critcal actions."},
			Dependencies: crosRepairPlan().GetCriticalActions(),
			ExecName:     "sample_pass",
		},
		"Install OS in DEV mode": {
			Docs: []string{"Install OS on the device from USB-key when device is in DEV-mode."},
			Dependencies: []string{
				"Set GBB flags to 0x18 by servo",
				"Boot DUT from USB in DEV mode",
				"Run install after boot from USB-drive",
				"Cold reset DUT by servo and wait to boot",
				"Wait DUT to be SSHable after reset",
			},
			ExecName: "sample_pass",
		},
		"Clean up": {
			Docs:         []string{"Verify that device is set to boot in DEV mode and enabled to boot from USB-drive."},
			Conditions:   []string{"dut_servo_host_present"},
			Dependencies: []string{"cros_remove_default_ap_file_servo_host"},
			ExecName:     "sample_pass",
		},
		"Collect DUT labels": {
			Docs: []string{"Updating device info in inventory."},
			Dependencies: []string{
				"cros_ssh",
				"cros_update_hwid_to_inventory",
				"cros_update_serial_number_inventory",
				"device_sku",
				"servo_type_label",
			},
			ExecName: "sample_pass",
		},
		"servo_type_label": {
			Docs:                   []string{"Update the servo type label for the DUT info."},
			ExecName:               "servo_update_servo_type_label",
			AllowFailAfterRecovery: true,
		},
	}
}

func crosDeployAndRepairActions() map[string]*Action {
	combo := deployActions()
	for name, action := range crosRepairActions() {
		if _, ok := combo[name]; ok {
			log.Fatalf("duplicate name in crosDeploy and crosRepair plan actions: %s", name)
		}
		combo[name] = action
	}
	return combo
}
