package deletion

func DropAllPkgs(pkg []string) bool {
	var success bool
	i := 0
	for i < len(pkg) {
		success = DeletePkg(pkg)
		i++
	}
	return success
}
