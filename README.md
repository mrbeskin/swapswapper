# update-swap-id

This is a simple Go library for replacing a Linux system swap UUID in `/etc/fstab`. 

It makes some assumptions:
- you have an `/etc/fstab` file
- you have a swap in `/etc/fstab` denoted by a UUID. Something like this:
```
# <device>                                <dir> <type> <options> <dump> <fsck> 
UUID=52dab1c1-0c6a-4f19-bdb7-d6e6617b26cf none  swap   defaults  0      0 
```


## Usage 

### As a Go Library

There is one function. 

example: 

```
package main

import "github.com/mrbeskin/swapswapper"

func main() {
	swapswapper.ReplaceSwapUUID("<your new UUID>")
}
```

If you did this, and meet the above assumptions, your `/etc/fstab` will have the value of `<your new UUID>` as the UUID value for the first swap partition listed. That's it.  
