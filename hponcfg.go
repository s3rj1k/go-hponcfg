package hponcfg

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

// GenericHPOnCfgWrapper - wrapper for HPOnCfg RO functions to circumvent stuck BMC
func GenericHPOnCfgWrapper(f func() ([]byte, error), retries int) ([]byte, error) {

	// try multiple times, (BMC bug?)
	for i := 0; i < retries; i++ {

		// cooldown
		time.Sleep(3 * time.Second)

		// get data from hponconfig
		data, err := f()
		if err == nil {
			return data, nil
		}

	}

	return []byte{}, fmt.Errorf("failed to get data from HPOnCfg multiple times")
}

// GetHPOnCfgHealthXML - run hponcfg and get GET_EMBEDDED_HEALTH data from STDOUT
func GetHPOnCfgHealthXML() ([]byte, error) {

	// command to run
	cmd := exec.Command("hponcfg", "-i")

	// stdin for hponcfg, GET_EMBEDDED_HEALTH
	cmd.Stdin = strings.NewReader(`
		<RIBCL VERSION="2.22">
			<LOGIN USER_LOGIN="Administrator" PASSWORD="">
				<SERVER_INFO MODE="read">
					<GET_EMBEDDED_HEALTH/>
				</SERVER_INFO>
			</LOGIN>
		</RIBCL>
	`)

	// syscall magic
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid:   true,
		Pdeathsig: syscall.SIGKILL,
	}

	// run command and capture output
	out, err := cmd.CombinedOutput()
	if err != nil {
		return []byte{}, err
	}

	return out, nil
}

// GetHPOnCfgNetworkXML - run hponcfg and get GET_NETWORK_SETTINGS data from STDOUT
func GetHPOnCfgNetworkXML() ([]byte, error) {

	// command to run
	cmd := exec.Command("hponcfg", "-i")

	// stdin for hponcfg, GET_NETWORK_SETTINGS
	cmd.Stdin = strings.NewReader(`
		<RIBCL VERSION="2.22">
			<LOGIN USER_LOGIN="Administrator" PASSWORD="">
				<RIB_INFO mode="read">
					<GET_NETWORK_SETTINGS/>
				</RIB_INFO>
			</LOGIN>
		</RIBCL>
	`)

	// syscall magic
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid:   true,
		Pdeathsig: syscall.SIGKILL,
	}

	// run command and capture output
	out, err := cmd.CombinedOutput()
	if err != nil {
		return []byte{}, err
	}

	return out, nil
}

// GetHPOnCfgBMCXML - run hponcfg and get GET_FW_VERSION data from STDOUT (only BMC related data)
func GetHPOnCfgBMCXML() ([]byte, error) {

	// command to run
	cmd := exec.Command("hponcfg", "-i")

	// stdin for hponcfg, GET_FW_VERSION
	cmd.Stdin = strings.NewReader(`
		<RIBCL VERSION="2.22">
			<LOGIN USER_LOGIN="Administrator" PASSWORD="">
				<RIB_INFO MODE="read">
					<GET_FW_VERSION/>
				</RIB_INFO>
			</LOGIN>
		</RIBCL>
	`)

	// syscall magic
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid:   true,
		Pdeathsig: syscall.SIGKILL,
	}

	// run command and capture output
	out, err := cmd.CombinedOutput()
	if err != nil {
		return []byte{}, err
	}

	return out, nil
}

// FactoryDefaults - resets BMC to factory default config using hponcfg
func FactoryDefaults() ([]byte, error) {

	// command to run
	cmd := exec.Command("hponcfg", "-i")

	// stdin for hponcfg, FACTORY_DEFAULTS
	cmd.Stdin = strings.NewReader(`
		<RIBCL VERSION="2.0">
			<LOGIN USER_LOGIN="Administrator" PASSWORD="">
				<RIB_INFO MODE="write">
					<FACTORY_DEFAULTS/>
				</RIB_INFO>
			</LOGIN>
		</RIBCL>
	`)

	// syscall magic
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid:   true,
		Pdeathsig: syscall.SIGKILL,
	}

	// run command and capture output
	out, err := cmd.CombinedOutput()
	if err != nil {
		return []byte{}, err
	}

	return out, nil
}

// SendXMLToHPOnCfg - resets BMC to factory default config using hponcfg
func SendXMLToHPOnCfg(data []byte) ([]byte, error) {

	// command to run
	cmd := exec.Command("hponcfg", "-i")

	// stdin for hponcfg
	cmd.Stdin = bytes.NewReader(data)

	// syscall magic
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid:   true,
		Pdeathsig: syscall.SIGKILL,
	}

	// run command and capture output
	out, err := cmd.CombinedOutput()
	if err != nil {
		return []byte{}, err
	}

	return out, nil
}

// FlashFirmware - flashes firmware using hponcfg
func FlashFirmware(isILO bool, firmware string) error {

	xml := `
		<RIBCL VERSION="2.0">
			<LOGIN USER_LOGIN="Administrator" PASSWORD="">
				<RIB_INFO MODE="write">
					<TPM_ENABLED VALUE="Yes"/>
					### REPLACE ###
				</RIB_INFO>
			</LOGIN>
		</RIBCL>
	`

	// check if firmware file exists
	if _, err := os.Stat(firmware); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("firmware file not found in: %s", firmware)
		}
	}

	if isILO {
		xml = strings.Replace(
			xml,
			"### REPLACE ###",
			fmt.Sprintf("<UPDATE_RIB_FIRMWARE IMAGE_LOCATION=\"%s\"/>", firmware),
			1,
		)
	} else {
		xml = strings.Replace(
			xml,
			"### REPLACE ###",
			fmt.Sprintf("<UPDATE_FIRMWARE IMAGE_LOCATION=\"%s\"/>", firmware),
			1,
		)
	}

	// command to run
	cmd := exec.Command("hponcfg", "-i")

	// stdin for hponcfg
	cmd.Stdin = strings.NewReader(xml)

	// syscall magic
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid:   true,
		Pdeathsig: syscall.SIGKILL,
	}

	// run command
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}
