# Diego Smoke Tests

This test suite is a variant of the [CF Smoke Tests][smoke-tests] that exercises
[Diego](https://github.com/cloudfoundry-incubator/diego-release).

## Running the tests

### Test Setup

Prerequisites:

- You should have user access to an org and a space in the CF instance you intend to test.
- The [CF cli](https://github.com/cloudfoundry/cli) and the [Ginkgo CLI](https://github.com/onsi/ginkgo) should both be on your PATH.
- This repository and its Go dependencies should be present in your GOPATH.
- The environment variable `SMOKE_TESTS_APPS_DOMAIN` should be exported to the
  name of a domain associated with the space above. For testing on a local BOSH-lite,
  "10.244.0.34.xip.io" is recommended. 

Before running the smoke tests, log into your CF instance with the CLI and
target the org and the space in which you intend to run the smoke tests. The
smoke tests will create an app which will be deleted under normal operation.
In the case of failure, though, it may remain present.

To run the tests, change your working directory to the root of this repo and run
`ginkgo`.

[smoke-tests]: https://github.com/cloudfoundry/cf-smoke-tests
