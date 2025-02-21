package listing

import (
	"depoty/internal/util/common"
	"regexp"
)

func ListPkgs() []string {
	text := common.ExecuteScript("choco list", "")

	var textSlice []string = make([]string, 0, 100)

	// Match the regular expression
	re := regexp.MustCompile(`(\S+)\s+([\d.]+)`)
	match := re.FindAllStringSubmatch(text, -1)

	for _, item := range match {
		textSlice = append(textSlice, item...)

	}
	return textSlice
}
