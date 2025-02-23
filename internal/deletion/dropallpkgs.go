package deletion

func DropAllPkgs(pkgNames []string) bool {
	var success bool
	for _, pkg := range pkgNames {
		success = DeletePkg(pkg)
	}
	return success
}
