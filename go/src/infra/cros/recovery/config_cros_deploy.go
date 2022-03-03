// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package recovery

// List of critical actions for deployment of the ChromeOS.
var crosDeployPlanCriticalActionList = []string{
	"Clean up",
	"Servo has USB-key with require image",
	"Device is pingable before deploy",
	"DUT has expected OS version",
	"DUT has expected test firmware",
	"Collect DUT labels",
	"Deployment checks",
	"DUT verify",
}

// List of actions configs for deployment of the ChromeOS.
var crosDeployPlanActions = `
"DUT is in dev-mode and allowed to boot from USB-key":{
	"docs":[
		"Verify that device is set to boot in DEV mode and enabled to boot from USB-drive."
	],
	"exec_timeout": {
		"seconds":2000
	},
	"exec_name":"cros_read_gbb_by_servo",
	"exec_extra_args":[
		"validate_in_dev_mode:true",
		"validate_usb_boot_enabled:true",
		"remove_file:false"
	],
	"recovery_actions":[
		"Set GBB flags to 0x18 by servo"
	]
},
"Device is pingable before deploy":{
	"docs":[
		"Verify that device is present in setup.",
		"All devices is pingable by default even they have prod images on them.",
		"If device is not pingable then device is off on not connected"
	],
	"exec_name":"cros_ping",
	"exec_timeout": {
		"seconds":15
	},
	"recovery_actions":[
		"Power cycle DUT by RPM and wait",
		"Set GBB flags to 0x18 by servo",
		"Install OS in DEV mode"
	]
},
"Cold reset DUT by servo and wait to boot":{
	"docs":[
		"Verify that device has stable version OS on it and version is match."
	],
	"dependencies":[
		"dut_servo_host_present",
		"servo_state_is_working",
		"Cold reset DUT by servo",
		"Wait DUT to be pingable after reset"
	],
	"exec_name":"sample_pass",
	"run_control": 1
},
"Cold reset DUT by servo":{
	"docs":[
		"Verify that device has stable version OS on it and version is match."
	],
	"dependencies":[
		"dut_servo_host_present",
		"servo_state_is_working"
	],
	"exec_name":"servo_set",
	"exec_extra_args":[
		"command:power_state",
		"string_value:reset"
	],
	"run_control": 1
},
"DUT has expected OS version":{
	"docs":[
		"Verify that device has stable version OS on it and version is match."
	],
	"dependencies":[
		"Device is pingable before deploy",
		"has_stable_version_cros_image",
		"has_test_cros_image"
	],
	"exec_name":"cros_is_on_stable_version",
	"recovery_actions":[
		"Install OS in DEV mode"
	]
},
"DUT has expected test firmware":{
	"docs":[
		"Verify that FW on the DUT has dev keys."
	],
	"dependencies":[
		"cros_ssh"
	],
	"exec_name":"cros_has_dev_signed_firmware",
	"exec_timeout": {
		"seconds":600
	},
	"recovery_actions":[
		"Update DUT firmware with factory mode and restart by servo",
		"Update DUT firmware with factory mode and restart by host"
	]
},
"Update DUT firmware with factory mode and restart by servo":{
	"docs":[
		"Force update FW on the DUT by factory mode.",
		"Reboot device by servo"
	],
	"conditions":[
		"servo_state_is_working"
	],
	"dependencies":[
		"cros_ssh"
	],
	"exec_name":"cros_run_firmware_update",
	"exec_extra_args":[
		"mode:factory",
		"force:true",
		"reboot:by_servo"
	]
},
"Update DUT firmware with factory mode and restart by host":{
	"docs":[
		"Force update FW on the DUT by factory mode.",
		"Reboot device by host"
	],
	"conditions":[
		"servo_state_is_not_working"
	],
	"dependencies":[
		"cros_ssh"
	],
	"exec_name":"cros_run_firmware_update",
	"exec_extra_args":[
		"mode:factory",
		"force:true",
		"reboot:by_host"
	]
},
"Deployment checks":{
	"docs":[
		"Run some specif checks as part of deployment.",
		"TODO: Need finished analysis to create final list"
	],
	"dependencies":[
		"Verify boot in recovery mode",
		"Verify RPM config (not critical)"
	],
	"exec_name":"sample_fail"
},
"Verify boot in recovery mode":{
	"docs":[
		"TODO: Not implemented yet!"
	],
	"exec_name":"sample_fail"
},
"Verify RPM config (not critical)":{
	"docs":[
		"TODO: Not implemented yet!"
	],
	"exec_name":"sample_fail"
},
"DUT verify":{
	"docs":[
		"Run all repair critcal actions."
	],
	"dependencies":[` + joinCriticalList(crosRepairPlanCriticalActionList) + `],
	"exec_name":"sample_pass"
},
"has_test_cros_image":{
	"docs":[
		"Verify that device has test OS image on it.",
		"TODO: Need implement"
	],
	"exec_name":"sample_fail"
},
"Install OS in DEV mode":{
	"docs":[
		"Install OS on the device from USB-key when device is in DEV-mode.",
		"TODO: Need implement"
	],
	"dependencies":[
		"Set GBB flags to 0x18 by servo",
		"Boot DUT from USB in DEV mode",
		"Run install after boot from USB-drive",
		"Cold reset DUT by servo and wait to boot",
		"wait_device_to_boot_after_reset"
	],
	"exec_name":"sample_pass"
},
"Boot DUT from USB in DEV mode":{
	"docs":[
		"Restart and try to boot from USB-drive"
	],
	"exec_name":"cros_dev_mode_boot_from_servo_usb_drive",
	"exec_extra_args":[
		"boot_timeout:180",
		"retry_interval:2"
	],
	"exec_timeout": {
		"seconds":300
	}
},
"Run install after boot from USB-drive":{
	"docs":[
		"Perform install process"
	],
	"exec_name":"cros_run_chromeos_install_command_after_boot_usbdrive",
	"exec_timeout": {
		"seconds":1200
	}
},
"Clean up":{
	"docs":[
		"Verify that device is set to boot in DEV mode and enabled to boot from USB-drive."
	],
	"dependencies":[
		"cros_remove_default_ap_file_servo_host"
	],
	"exec_name":"sample_pass"
},
"Collect DUT labels":{
	"dependencies":[
		"device_sku",
		"servo_type_label"
	],
	"exec_name":"sample_pass"
},
"servo_type_label":{
	"docs":[
		"Update the servo type label for the DUT info."
	],
	"exec_name":"servo_update_servo_type_label",
	"allow_fail_after_recovery": true
},
"device_sku":{
	"docs":[
		"Update the device_sku label from the device if not present in inventory data."
	],
	"conditions":[
		"dut_does_not_have_device_sku"
	],
	"exec_name":"cros_update_device_sku",
	"allow_fail_after_recovery": true
},
"dut_does_not_have_device_sku":{
	"docs":[
		"Confirm that the DUT itself does not have device_sku label."
	],
	"conditions":[
		"dut_has_device_sku"
	],
	"exec_name":"sample_fail"
},
`

// Represents the Chrome OS deployment plan for DUT.
var crosDeployPlanBody = `"critical_actions": [` + joinCriticalList(crosDeployPlanCriticalActionList) + `],
"actions": {` + crosDeployPlanActions + crosRepairPlanActions + `}`
