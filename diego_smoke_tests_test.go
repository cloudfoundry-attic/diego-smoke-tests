package diego_smoke_tests_test

import (
	"os"
	"time"

	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/cloudfoundry-incubator/cf-test-helpers/generator"
	"github.com/cloudfoundry-incubator/cf-test-helpers/runner"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Staging and running", func() {
	var (
		appName    string
		appsDomain string
		appRoute   string
	)

	BeforeEach(func() {
		appName = generator.RandomName()
		appsDomain = os.Getenv("SMOKE_TESTS_APPS_DOMAIN")
		Ω(appsDomain).ShouldNot(BeEmpty(), "must set $SMOKE_TESTS_APPS_DOMAIN")
		appRoute = "http://" + appName + "." + appsDomain + "/"
	})

	AfterEach(func() {
		Eventually(cf.Cf("logs", appName, "--recent")).Should(gexec.Exit())
		Eventually(cf.Cf("delete", "-r", "-f", appName)).Should(gexec.Exit(0))
	})

	It("works", func() {
		Eventually(cf.Cf("push", appName, "-p", "dora", "--no-start")).Should(gexec.Exit(0))
		enableDiego(appName)
		Eventually(cf.Cf("start", appName), 5*time.Minute).Should(gexec.Exit(0))

		Eventually(func() *gexec.Session {
			curl := runner.Curl(appRoute)
			curl.Wait()
			return curl
		}).Should(gbytes.Say("Hi, I'm Dora!"))
	})
})
