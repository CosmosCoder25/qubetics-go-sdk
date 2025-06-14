package v2ray

import (
	"path/filepath"
)

// execFile returns the name of the executable file.
func (c *Client) execFile(name string) string {
	return ".\\" + filepath.Join("V2Ray", name+".exe")
}
