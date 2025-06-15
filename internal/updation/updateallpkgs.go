package updation

func UpdateAllPkgs(pkgs []string) bool {
	// success, err := common.ExecutePrevScript("choco upgrade all", "-y")
	// if err != nil {
	// 	return false
	// }
	// return success
	var success bool
	i := 0
	for i < len(pkgs) {
		success = UpdatePkg(pkgs[i])
		i++
	}
	return success
}
