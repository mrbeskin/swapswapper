package swapswapper

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// ReplaceSwapUUID reads /etc/fstab and replaces the first swap id
// found there with the one passed in
func ReplaceSwapUUID(newUUID string) error {
	fstabName := "/etc/fstab"
	fstab, err := ioutil.ReadFile(fstabName)
	if err != nil {
		return err
	}
	newFstab, err := replaceSwapUUID(newUUID, string(fstab))
	if err != nil {
		return fmt.Errorf("modifying file: %v", err)
	}
	ioutil.WriteFile(fstabName, []byte(newFstab), 0644)
	return nil
}

func replaceSwapUUID(newUUID, fstab string) (string, error) {
	lines := strings.Split(fstab, "\n")
	for i, line := range lines {
		line = strings.TrimRight(line, " ")
		if strings.HasPrefix(line, "UUID=") {
			if strings.Contains(line, "swap") {
				toks := strings.Split(line, " ")
				newPrefix := "UUID=" + newUUID
				toks[0] = newPrefix
				newLine := strings.Join(toks, " ")
				lines[i] = newLine
				newFstab := strings.Join(lines, "\n")
				return newFstab, nil
			}
		}
	}
	return fstab, fmt.Errorf("no swap found")
}
