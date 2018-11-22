package main

import (
	"testing"
)

var testSwpFstab = `
# <device>                                <dir> <type> <options> <dump> <fsck>
UUID=CBB6-24F2                            /boot vfat   defaults  0      2
UUID=0a3407de-014b-458b-b5c1-848e92a327a3 /     ext4   defaults  0      1
UUID=b411dc99-f0a0-4c87-9e05-184977be8539 /home ext4   defaults  0      2
UUID=f9fe0b69-a280-415d-a03a-a32752370dee none  swap   defaults  0      0
`

var testNoSwpFstab = `
# <device>                                <dir> <type> <options> <dump> <fsck>
UUID=CBB6-24F2                            /boot vfat   defaults  0      2
UUID=0a3407de-014b-458b-b5c1-848e92a327a3 /     ext4   defaults  0      1
UUID=b411dc99-f0a0-4c87-9e05-184977be8539 /home ext4   defaults  0      2
`

var testSwpFstabExpected = `
# <device>                                <dir> <type> <options> <dump> <fsck>
UUID=CBB6-24F2                            /boot vfat   defaults  0      2
UUID=0a3407de-014b-458b-b5c1-848e92a327a3 /     ext4   defaults  0      1
UUID=b411dc99-f0a0-4c87-9e05-184977be8539 /home ext4   defaults  0      2
UUID=52dab1c1-0c6a-4f19-bdb7-d6e6617b26cf none  swap   defaults  0      0
`

// success
func TestReplaceSwapUUID(t *testing.T) {
	newFstab, err := replaceSwapUUID("52dab1c1-0c6a-4f19-bdb7-d6e6617b26cf", testSwpFstab)
	if err != nil {
		t.Error(err)
	}
	if newFstab != testSwpFstabExpected {
		t.Errorf("expected fstab output:\n%v\nreturned fstab:\n%v\n", testSwpFstabExpected, newFstab)
	}
}

// Test fstab with no swp value
func TestReplaceSwapUUIDErr(t *testing.T) {
	_, err := replaceSwapUUID("52dab1c1-0c6a-4f19-bdb7-d6e6617b26cf", testNoSwpFstab)
	if err == nil {
		t.Errorf("expected error, but got nil")
	}
}
