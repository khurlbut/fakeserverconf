package fakeserverconf_test

import (
	. "github.com/khurlbut/fakeserverconf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Configurator Tests", func() {

	It("should return a default Root page config", func() {
		c := DefaultConfig()
		Ω(c.Pages[0].Path).Should(Equal("/"))
		Ω(c.Pages[0].Status).Should(Equal(200))
		Ω(c.Pages[0].Body).Should(Equal("Welcome to fakeserver. This is the Default configuration for Root path (/)"))
	})

	It("should read a single config from a JSON string", func() {
		c := ReadJSON(`{"Pages":[{"Path":"/","Status": 200,"Body":"This is from a JSON string"}]}`)
		Ω(c.Pages[0].Path).Should(Equal("/"))
		Ω(c.Pages[0].Status).Should(Equal(200))
		Ω(c.Pages[0].Body).Should(Equal("This is from a JSON string"))
	})

	It("should read multiple configs from a JSON string array", func() {
		c := ReadJSON(`{"Pages":[{"Path":"/","Status": 200,"Body":"Body 1"},{"Path":"/","Status": 200,"Body":"Body 2"}]}`)
		Ω(c.Pages[0].Path).Should(Equal("/"))
		Ω(c.Pages[0].Status).Should(Equal(200))
		Ω(c.Pages[0].Body).Should(Equal("Body 1"))
		Ω(c.Pages[1].Path).Should(Equal("/"))
		Ω(c.Pages[1].Status).Should(Equal(200))
		Ω(c.Pages[1].Body).Should(Equal("Body 2"))
	})

	It("should read a port number from the configuration", func() {
		c := ReadJSON(`{"Port": 8000, "Pages":[]}`)
		Ω(c.Port).Should(Equal(8000))
	})

	It("should read an IP address string from the configuration", func() {
		c := ReadJSON(`{"IPAddress": "host", "Pages":[]}`)
		Ω(c.IPAddress).Should(Equal("host"))
	})

})
