package hponcfg

import (
	"bytes"
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
func (data []byte) SendXMLToHPOnCfg() ([]byte, error) {

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
