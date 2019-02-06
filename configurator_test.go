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
		c := ReadJSON(`{"Port": "8000", "Pages":[]}`)
		Ω(c.Port).Should(Equal("8000"))
	})

	It("should read an IP address string from the configuration", func() {
		c := ReadJSON(`{"IPAddress": "127.0.0.1", "Pages":[]}`)
		Ω(c.IPAddress).Should(Equal("127.0.0.1"))
	})

	It("should read an array of page headers from the configuration", func() {
		c := ReadJSON(`{"Pages":[{"Path":"/","Status": 200,"Body":"Body", "Headers":["key1:val1","key2:val2"]}]}`)
		Ω(c.Pages[0].Headers[0]).Should(Equal("key1:val1"))
		Ω(c.Pages[0].Headers[1]).Should(Equal("key2:val2"))
	})

	It("should read an array of cookes from the configuration", func() {
		c := ReadJSON(`{"Pages":[{"Path":"/","Status": 200,"Body":"Body", "Headers":["key1:val1","key2:val2"], "Cookies":["cook1:val1","cook2:val2"]}]}`)
		Ω(c.Pages[0].Cookies[0]).Should(Equal("cook1:val1"))
		Ω(c.Pages[0].Cookies[1]).Should(Equal("cook2:val2"))
	})

	It("should read an array of injection keys from the configuration", func() {
		c := ReadJSON(`{"Pages":[{"Path":"/","Status": 200,"Body":"Body", "InjectionKeys": ["path", "uri"], "Headers":["key1:val1","key2:val2"], "Cookies":["cook1:val1","cook2:val2"]}]}`)
		Ω(c.Pages[0].InjectionKeys[0]).Should(Equal("path"))
		Ω(c.Pages[0].InjectionKeys[1]).Should(Equal("uri"))
	})

	It("should read an array of service endpoints from the configuration", func() {
		c := ReadJSON(`{"Pages":[{"Path":"/","Status": 200,"Body":"Body", "InjectionKeys": ["path", "uri"], "Headers":["key1:val1","key2:val2"], "Cookies":["cook1:val1","cook2:val2"], "ServiceEndpoints": ["uri1","uri2","uri3"]}]}`)
		Ω(c.Pages[0].ServiceEndpoints[0]).Should(Equal("uri1"))
		Ω(c.Pages[0].ServiceEndpoints[1]).Should(Equal("uri2"))
		Ω(c.Pages[0].ServiceEndpoints[2]).Should(Equal("uri3"))
	})

})
