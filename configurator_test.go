package fakeserverconf_test

import (
	. "github.com/khurlbut/fakeserverconf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Configurator Tests", func() {
	c := Configurator()

	It("should return a default Root page config", func() {
		Ω(c[0].Path).Should(Equal("/"))
		Ω(c[0].Status).Should(Equal(200))
		Ω(c[0].Body).Should(Equal("Welcome to fakeserver. This is the Default configuration for Root path (/)"))
	})

})
