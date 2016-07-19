package printer

import (
	"fmt"

	"CodeFly/reader"
)

// Print 输出文件
func Print(r *reader.ThriftReader) error {

	str := reader.SwiftReader
	str.InitSwiftThrift(r)

	return fmt.Errorf("Error")
}

func printWith(str *reader.SwiftThriftReader) error {
	return fmt.Errorf("Error")
}
