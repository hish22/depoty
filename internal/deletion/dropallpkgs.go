package deletion

import "runtime"

func DropAllPkgs(pkgs []string) bool {

	switch runtime.GOOS {
	case "windows":
		var success bool
		i := 0
		for i < len(pkgs) {
			success = DeletePkg(pkgs[i])
			i++
		}
		return success
	case "linux":
		return false
	default:
		return false
	}

}
