package fakeserverconf_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFakeserverconf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fakeserverconf Suite")
}
