package glog

import (
	"context"
	"fmt"

	"github.com/nutscloud/go-util/log"

	"github.com/golang/glog"
)

type LevelLog int

func (log *LevelLog) Info(args ...interface{}) {
	glog.V(glog.Level(*log)).Info(args)
}

func (log *LevelLog) Infof(format string, args ...interface{}) {
	glog.V(glog.Level(*log)).Infof(format, args)
}

type Glog struct {
	ctx context.Context
}

func NewGlog() *Glog {
	return &Glog{}
}

func (log *Glog) WithContext(ctx context.Context) log.Logger {
	log.ctx = ctx
	return log
}

func (log *Glog) Info(args ...interface{}) {
	glog.InfoDepth(1, args)
}

func (log *Glog) Infof(format string, args ...interface{}) {
	glog.InfoDepth(1, fmt.Sprintf(format, args))
}

func (log *Glog) InfoDepth(depth int, args ...interface{}) {
	glog.InfoDepth(depth+1, args)
}

func (log *Glog) V(level int) log.LevelLogger {
	levelLog := LevelLog(level)
	return &levelLog
}

func (log *Glog) SetLevel(l int) error {
	var level glog.Level
	if err := level.Set(fmt.Sprintf("%d", l)); err != nil {
		return fmt.Errorf("failed set glog.logging.verbosity %s: %v\n", l, err)
	}
	return nil
}

func (log *Glog) Warning(args ...interface{}) {
	glog.WarningDepth(1, args)
}

func (log *Glog) Warningf(format string, args ...interface{}) {
	glog.WarningDepth(1, fmt.Sprintf(format, args))
}

func (log *Glog) WarningDepth(depth int, args ...interface{}) {
	glog.WarningDepth(depth+1, args)
}

func (log *Glog) Error(args ...interface{}) {
	glog.ErrorDepth(1, args)
}

func (log *Glog) Errorf(format string, args ...interface{}) {
	glog.ErrorDepth(1, fmt.Sprintf(format, args))
}

func (log *Glog) ErrorDepth(depth int, args ...interface{}) {
	glog.ErrorDepth(depth+1, args)
}

func (log *Glog) Fatal(args ...interface{}) {
	glog.FatalDepth(1, args)
}

func (log *Glog) Fatalf(format string, args ...interface{}) {
	glog.FatalDepth(1, fmt.Sprintf(format, args))
}

func (log *Glog) FatalDepth(depth int, args ...interface{}) {
	glog.FatalDepth(depth+1, args)
}

func (log *Glog) Exit(args ...interface{}) {
	glog.ExitDepth(1, args)
}

func (log *Glog) Exitf(format string, args ...interface{}) {
	glog.ExitDepth(1, fmt.Sprintf(format, args))
}

func (log *Glog) ExitDepth(depth int, args ...interface{}) {
	glog.ExitDepth(depth+1, args)
}

func (log *Glog) Flush() {
	glog.Flush()
}
