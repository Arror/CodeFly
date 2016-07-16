package global

// GenerateCommandInfo 命令信息结构
type GenerateCommandInfo struct {
	Lang   string
	Input  string
	Output string
}

// GenCmdInfo 命令信息对象
var GenCmdInfo = &GenerateCommandInfo{}

// CheckValidity 检查输入命令合法性
func (gci *GenerateCommandInfo) CheckValidity() error {
	return nil
}
