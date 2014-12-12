package diego_smoke_tests_test

import (
	"time"

	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/cloudfoundry-incubator/cf-test-helpers/generator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Staging", func() {
	var appName string

	BeforeEach(func() {
		appName = generator.RandomName()
	})

	AfterEach(func() {
		Eventually(cf.Cf("logs", appName, "--recent")).Should(gexec.Exit())
		Eventually(cf.Cf("delete", "-r", "-f", appName)).Should(gexec.Exit(0))
	})

	It("works", func() {
		Eventually(cf.Cf("push", appName, "-p", "dora", "--no-start")).Should(gexec.Exit(0))
		Eventually(cf.Cf("set-env", appName, "DIEGO_STAGE_BETA", "true")).Should(gexec.Exit(0))
		Eventually(cf.Cf("start", appName), 5*time.Minute).Should(gexec.Exit(0))
	})
})
