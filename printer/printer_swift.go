package printer

import (
	"fmt"

	"CodeFly/global"
	"CodeFly/reader"
)

// Print 输出文件
func Print(info *global.GenerateCommandInfo) error {
	if info.Lang != global.Swift {
		return fmt.Errorf("不支持的语言: %s", info.Lang)
	}

	fmt.Println("Hallo!!!")

	return nil
}

func printWith(str *reader.SwiftThriftReader) error {
	return nil
}
