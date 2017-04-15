package hostkeys

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var matchKeyBlock = regexp.MustCompile(`(?s)-----BEGIN SSH HOST KEY KEYS-----\r\n(.*)\r\n-----END SSH HOST KEY KEYS-----`)

// ErrNoStartHostKeysBlock indicates that the required block of text containing
// host keys was not found in the CloudInit output.
var ErrNoStartHostKeysBlock = errors.New("No BEGIN SSH HOST KEY KEYS block found in CloudInit output")

// Parse extracts the SSH host public keys from CloudInit log output
// such as is obtained from the Get-Console-Output command in AWS, or perhaps via
// alternative means for other cloud providers.
func Parse(cloudInitOutput string) ([]string, error) {
	blocks := matchKeyBlock.FindAllStringSubmatch(cloudInitOutput, -1)

	if len(blocks) == 0 {
		return nil, ErrNoStartHostKeysBlock
	}

	var keys []string
	for _, block := range blocks {
		keyLines := strings.SplitAfter(block[1], "\r\n")
		for _, keyLine := range keyLines {
			keyParts := strings.SplitN(keyLine, " ", 3)
			if len(keyParts) < 2 || keyParts[0] == "\r\n" || keyParts[1] == "\r\n" {
				return nil, fmt.Errorf("invalid host key format: %s", strings.TrimSpace(keyLine))
			}

			keyNoComment := fmt.Sprintf("%s %s", keyParts[0], keyParts[1])
			keys = append(keys, keyNoComment)
		}
	}

	return keys, nil
}
