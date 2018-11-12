package fakeserverconf_test

import (
	. "github.com/khurlbut/fakeserverconf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Configurator Tests", func() {

	It("should return a default Root page config", func() {
		c := DefaultConfig()
		Ω(c[0].Path).Should(Equal("/"))
		Ω(c[0].Status).Should(Equal(200))
		Ω(c[0].Body).Should(Equal("Welcome to fakeserver. This is the Default configuration for Root path (/)"))
	})

	It("should read a single config from a JSON string", func() {
		c := ReadJSON(`[{"Path":"/","Status": 200,"Body":"This is from a JSON string"}]`)
		Ω(c[0].Path).Should(Equal("/"))
		Ω(c[0].Status).Should(Equal(200))
		Ω(c[0].Body).Should(Equal("This is from a JSON string"))
	})

	It("should read multiple configs from a JSON string array", func() {
		c := ReadJSON(`[{"Path":"/","Status": 200,"Body":"Body 1"},{"Path":"/","Status": 200,"Body":"Body 2"}]`)
		Ω(c[0].Path).Should(Equal("/"))
		Ω(c[0].Status).Should(Equal(200))
		Ω(c[0].Body).Should(Equal("Body 1"))
		Ω(c[1].Path).Should(Equal("/"))
		Ω(c[1].Status).Should(Equal(200))
		Ω(c[1].Body).Should(Equal("Body 2"))
	})

})
