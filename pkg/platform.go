package pkg

import (
	"fmt"

	"github.com/frankbraun/codechain/util"
)

// computed by `go tool dist list`
var validPlatforms = []string{
	"aix_ppc64",
	"android_386",
	"android_amd64",
	"android_arm",
	"android_arm64",
	"darwin_386",
	"darwin_amd64",
	"darwin_arm",
	"darwin_arm64",
	"dragonfly_amd64",
	"freebsd_386",
	"freebsd_amd64",
	"freebsd_arm",
	"js_wasm",
	"linux_386",
	"linux_amd64",
	"linux_arm",
	"linux_arm64",
	"linux_mips",
	"linux_mips64",
	"linux_mips64le",
	"linux_mipsle",
	"linux_ppc64",
	"linux_ppc64le",
	"linux_s390x",
	"nacl_386",
	"nacl_amd64p32",
	"nacl_arm",
	"netbsd_386",
	"netbsd_amd64",
	"netbsd_arm",
	"openbsd_386",
	"openbsd_amd64",
	"openbsd_arm",
	"plan9_386",
	"plan9_amd64",
	"plan9_arm",
	"solaris_amd64",
	"windows_386",
	"windows_amd64",
	"windows_arm",
}

// platformIsValid makes sure the given platform is valid.
func platformIsValid(platform string) error {
	if !util.ContainsString(validPlatforms, platform) {
		return fmt.Errorf("'%s' is not a valid platform", platform)
	}
	return nil
}
