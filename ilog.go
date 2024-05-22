// Package logrus
// @Title  接口
// @Description  接口
// @Author  yr  2024/5/22 上午10:24
// @Update  yr  2024/5/22 上午10:24
package logrus

import "io"

type ILogger interface {
	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	TraceFn(fn LogFunction)
	DebugFn(fn LogFunction)
	InfoFn(fn LogFunction)
	PrintFn(fn LogFunction)
	WarnFn(fn LogFunction)
	WarningFn(fn LogFunction)
	ErrorFn(fn LogFunction)
	FatalFn(fn LogFunction)
	PanicFn(fn LogFunction)

	Traceln(args ...interface{})
	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Println(args ...interface{})
	Warnln(args ...interface{})
	Warningln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Panicln(args ...interface{})

	Exit(code int)
	Writer() *io.PipeWriter
	WriterLevel(level Level) *io.PipeWriter
	GetOutput() io.Writer
	SetNoLock()
	Release(writer io.Writer) error
}
