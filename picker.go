// Package logrus
// 模块名: 选择器
// 功能描述: 描述
// 作者:  yr  2025/12/1 00:09
// 最后更新:  yr  2025/12/1 00:09
package logrus

import (
	"io"
)

type IPicker interface {
	Pick(entry *Entry) io.Writer
}

type Picker struct {
}

func (p *Picker) Pick(entry *Entry) io.Writer {
	return entry.Logger.Out
}
