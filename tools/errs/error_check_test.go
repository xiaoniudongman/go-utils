package errs_test

import (
	"fmt"
	"testing"

	"github.com/getsentry/raven-go"

	"github.com/xiaoniudongman/go-utils/tools/errs"

	"github.com/xiaoniudongman/go-utils/config"
	"github.com/xiaoniudongman/go-utils/tools/logs"
)

const (
	logFile    = "../../config/test.yaml"
	logxmlFile = "../../config/seelog.xml"
)

func TestErrLog1(b *testing.T) {
	appConfig := &config.ConfigEngine{}
	logs.LoggerSetup(logxmlFile)
	errs.CheckFatalErr(appConfig.Load(logFile))
	errs.CheckFatalErr(appConfig.SentryRavenInit("Sentry_dsn"))
	//errs.CheckErrSendEmail(fmt.Errorf("asdfasdfsafdsdfs"))
	raven.CaptureError(fmt.Errorf("asdfasdfsafdsdfs"), nil)
}

func TestErrLog(b *testing.T) {
	appConfig := &config.ConfigEngine{}
	logs.LoggerSetup(logxmlFile)
	errs.CheckFatalErr(appConfig.Load(logFile))
	errs.CheckFatalErr(appConfig.SentryRavenInit("Sentry_dsn"))
	errs.CheckErrSendEmail(fmt.Errorf("asdfasdfsafdsdfs"))
}

func TestRusErrLog1(b *testing.T) {
	appConfig := &config.ConfigEngine{}
	errs.CheckFatalErr(appConfig.Load(logFile))
	errs.CheckFatalErr(appConfig.SentryRavenInit("Sentry_dsn"))
	errs.CheckLogrusCaptureError(fmt.Errorf("asdf"), nil, "yyy", "ppp")
}
