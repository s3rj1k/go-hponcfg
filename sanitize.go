package hponcfg

import (
	"bytes"
	"errors"
	"regexp"
)

// SanitizeHPOnCfgXML - sanitizes output of hponcfg XML data (assume that XML has no inner value only attributes)
func SanitizeHPOnCfgXML(data []byte) ([]byte, error) {

	// fix XML tags in firmware section for proper unmarshalling
	data = regexp.MustCompile("(?i)<INDEX_[0-9]+>").ReplaceAllLiteral(data, []byte("<INDEX>"))
	data = regexp.MustCompile("(?i)</INDEX_[0-9]+>").ReplaceAllLiteral(data, []byte("</INDEX>"))

	// replace non-ASCII characters with space
	data = regexp.MustCompile("[[:^ascii:]]").ReplaceAllLiteral(data, []byte(" "))

	// replace control characters with space
	data = regexp.MustCompile("[[:cntrl:]]").ReplaceAllLiteral(data, []byte(" "))

	// find all strings inside <> using non-gready regexp
	clean := regexp.MustCompile("<.+?>").FindAll(data, -1)
	if clean == nil {
		return []byte{}, errors.New("failed to find XML elements")
	}

	return bytes.Join(clean, []byte("")), nil
}
