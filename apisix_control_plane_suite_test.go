package apisix_control_plane_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestApisixControlPlane(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ApisixControlPlane Suite")
}
