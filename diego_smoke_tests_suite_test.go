package diego_smoke_tests_test

import (
	"strings"
	"testing"
	"time"

	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

func TestDiegoSmokeTests(t *testing.T) {
	SetDefaultEventuallyTimeout(time.Minute)
	SetDefaultEventuallyPollingInterval(time.Second)

	RegisterFailHandler(Fail)
	RunSpecs(t, "DiegoSmokeTests Suite")
}

func guidForAppName(appName string) string {
	cfApp := cf.Cf("app", appName, "--guid")
	Expect(cfApp.Wait()).To(gexec.Exit(0))

	appGuid := strings.TrimSpace(string(cfApp.Out.Contents()))
	Expect(appGuid).NotTo(Equal(""))
	return appGuid
}

func enableDiego(appName string) {
	guid := guidForAppName(appName)
	Eventually(cf.Cf("curl", "/v2/apps/"+guid, "-X", "PUT", "-d", `{"diego": true}`)).Should(gexec.Exit(0))
}
