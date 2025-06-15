package deletion

func DropAllPkgs(pkgs []string) bool {
	var success bool
	i := 0
	for i < len(pkgs) {
		success = DeletePkg(pkgs[i])
		i++
	}
	return success
}
