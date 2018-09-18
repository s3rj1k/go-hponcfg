package hponcfg

import (
	"bufio"
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

	// convert output to scanner
	scanner := bufio.NewScanner(
		bytes.NewReader(out),
	)

	// filtered raw XML
	xml := make([]byte, 0)

	// loop-over lines
	for scanner.Scan() {
		// get current line
		line := bytes.TrimSpace(scanner.Bytes())
		// check if line begins with XML
		if bytes.HasPrefix(line, []byte("<")) {
			xml = append(xml, line...)
		}
	}

	// if scanner fails, return error
	if err := scanner.Err(); err != nil {
		return []byte{}, err
	}

	return xml, nil
}
