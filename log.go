package log

import (
	"errors"
	"os"

	"github.com/fatih/color"
)

func Error(err_message error) {
	err := errors.New(err_message.Error())
	ErrorNotice := color.New(color.FgHiRed).Add(color.Bold).Add(color.BgRed).PrintfFunc()
	ErrorText := color.New(color.FgWhite).Add(color.Bold).PrintfFunc()

	ErrorNotice(" ERROR ")
	ErrorText(" %s\n", err)
	os.Exit(1)
}

func Info(info_message string) {
	InfoNotice := color.New(color.FgHiBlue).Add(color.Bold).Add(color.BgBlue).PrintfFunc()
	InfoText := color.New(color.FgWhite).Add(color.Bold).PrintfFunc()

	InfoNotice(" INFO ")
	InfoText(" %s\n", info_message)
}

func Warn(warn_message string) {
	WarnNotice := color.New(color.FgHiYellow).Add(color.Bold).Add(color.BgYellow).PrintfFunc()
	WarnText := color.New(color.FgWhite).Add(color.Bold).PrintfFunc()

	WarnNotice(" WARN ")
	WarnText(" %s\n", warn_message)

}

func Debug(debug_message string) {
	DebugNotice := color.New(color.FgHiMagenta).Add(color.Bold).Add(color.BgMagenta).PrintfFunc()
	DebugText := color.New(color.FgWhite).Add(color.Bold).PrintfFunc()

	DebugNotice(" DEBUG ")
	DebugText(" %s\n", debug_message)
}

func Success(success_message string) {
	SuccessNotice := color.New(color.FgHiGreen).Add(color.Bold).Add(color.BgGreen).PrintfFunc()
	SuccessText := color.New(color.FgWhite).Add(color.Bold).PrintfFunc()

	SuccessNotice(" SUCCESS ")
	SuccessText(" %s\n", success_message)
}
