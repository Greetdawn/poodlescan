package logger

import (
	"fmt"
	pb "poodle/pkg/mygrpc"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SRV *pb.Kernel_SendOrderServer

// 是否打印内核日志
var IsPrintLogInfo bool = true

const (
	LOG_TERMINAL LOG_OUTPUT_MODE = iota
	LOG_FILE
	LOG_TERMINAL_FILE
)

var SugarLogger *zap.SugaredLogger

type LOG_OUTPUT_MODE byte

// 自定义时间显示格式：控制台和文件相同
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("[%s]", t.Format("2006-01-02 15:04:05")))
}

// 自定义时间显示格式：控制台和文件相同
func CustomLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var lev string
	switch l {
	case zapcore.DebugLevel:
		lev = fmt.Sprintf("%-8s", "[DEBUG]")
	case zapcore.InfoLevel:
		lev = fmt.Sprintf("%-8s", "[INFO]")
	case zapcore.WarnLevel:
		lev = fmt.Sprintf("%-8s", "[WARN]")
	case zapcore.ErrorLevel:
		lev = fmt.Sprintf("%-8s", "[ERROR]")
	case zapcore.DPanicLevel:
		lev = fmt.Sprintf("%-8s", "[DPANIC]")
	case zapcore.PanicLevel:
		lev = fmt.Sprintf("%-8s", "[PANIC]")
	case zapcore.FatalLevel:
		lev = fmt.Sprintf("%-8s", "[FATAL]")
	}
	enc.AppendString(lev)
}

// 自定义显示调用者
func FullCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(caller.String())
}

// 自定义显示调用者
func NoCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(" ")
	enc = nil
}

// 自定义Zap日志编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = CustomTimeEncoder
	// zap库默认
	//encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeLevel = CustomLevelEncoder
	// 关闭显示调用者信息
	encoderConfig.EncodeCaller = NoCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func init() {
	initLogger()
}

func initLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.FatalLevel)
	// 不用打印调用者 logger := zap.New(core, zap.AddCaller())
	logger := zap.New(core)
	SugarLogger = logger.Sugar()
	defer SugarLogger.Sync()
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "test.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// 记录INFO日志
func LogInfo(log string, mode LOG_OUTPUT_MODE) {
	if SRV != nil {
		_ = (*SRV).Send(&pb.SendOrderReply{
			Info: log,
		})
	}
	if !IsPrintLogInfo {
		return
	}

	if mode == LOG_FILE {
		SugarLogger.Info(log)
	} else if mode == LOG_TERMINAL {
		logTeriminal("["+FgGreen("INFO")+"]", log)
	} else if mode == LOG_TERMINAL_FILE {
		SugarLogger.Info(log)
		logTeriminal("["+FgGreen("INFO")+"]", log)
	}
}

// 记录Warn日志
func LogWarn(log string, mode LOG_OUTPUT_MODE) {
	if SRV != nil {
		_ = (*SRV).Send(&pb.SendOrderReply{
			Info: log,
		})
	}

	if !IsPrintLogInfo {
		return
	}
	if mode == LOG_FILE {
		SugarLogger.Warn(log)
	} else if mode == LOG_TERMINAL {
		logTeriminal("["+FgYellow("WARN")+"]", log)
	} else if mode == LOG_TERMINAL_FILE {
		SugarLogger.Warn(log)
		logTeriminal("["+FgYellow("WARN")+"]", log)
	}
}

// 记录Error日志
func LogError(log string, mode LOG_OUTPUT_MODE) {
	if SRV != nil {
		_ = (*SRV).Send(&pb.SendOrderReply{
			Info: log,
		})
	}
	if !IsPrintLogInfo {
		return
	}
	if mode == LOG_FILE {
		SugarLogger.Error(log)
	} else if mode == LOG_TERMINAL {
		logTeriminal("["+BgRedFgWhite("ERROR")+"]", log)
	} else if mode == LOG_TERMINAL_FILE {
		SugarLogger.Error(log)
		logTeriminal("["+BgRedFgWhite("ERROR")+"]", log)
	}
}

func LogNoFormat(log string, mode LOG_OUTPUT_MODE) {
	if !IsPrintLogInfo {
		return
	}
	if mode == LOG_FILE {
		SugarLogger.Info(log)
	} else if mode == LOG_TERMINAL {
		fmt.Print(log)
	} else if mode == LOG_TERMINAL_FILE {
		SugarLogger.Info(log)
		fmt.Print(log)
	}
}

func getShortCallerInfo() string {
	_, f, line, _ := runtime.Caller(3)
	names := strings.Split(f, "/")
	size := len(names)
	return names[size-2] + "/" + names[size-1] + ":" + strconv.Itoa(line)
}

func logTeriminal(level, log string) {
	// 不显示调用者
	fmt.Printf("[%s] %s %s\n", FgGreen(time.Now().Format("2006-01-02 15:04:05")), level, log)
	// fmt.Printf(fmt.Sprintf("[%s] %s %s %s\n", FgGreen(time.Now().Format("2006-01-02 15:04:05")), level, getShortCallerInfo(), log))
}
