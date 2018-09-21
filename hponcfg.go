package hponcfg

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

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
func FlashFirmware(printToStdout, isILO bool, firmware string) error {

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

	// command to run
	cmd := exec.Command("hponcfg", "-i")

	if isILO {
		xml = strings.Replace(
			xml,
			"### REPLACE ###",
			fmt.Sprint("<UPDATE_RIB_FIRMWARE IMAGE_LOCATION=\"%s\"/>", firmware),
			1,
		)
	} else {
		xml = strings.Replace(
			xml,
			"### REPLACE ###",
			fmt.Sprint("<UPDATE_FIRMWARE IMAGE_LOCATION=\"%s\"/>", firmware),
			1,
		)
	}

	if printToStdout {

		// set stdout pipe
		cmdReader, err := cmd.StdoutPipe()
		if err != nil {
			return fmt.Errorf("failed to set stdout pipe")
		}

		// set scanner for cmd output
		scanner := bufio.NewScanner(cmdReader)

		// print lines to stdout
		go func() {
			for scanner.Scan() {
				fmt.Printf(scanner.Text())
			}
		}()

	}

	// syscall magic
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid:   true,
		Pdeathsig: syscall.SIGKILL,
	}

	// run command
	cmd.Start()

	// wait for command to complete execution
	err := cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}
