package main_test

import (
	main "github/repeating-string-testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CountRepeatedWords", func() {
	var (
		input    string
		expected int
	)

	BeforeEach(func() {
		input = ""
		expected = 0
	})

	When("input is empty", func() {
		It("should return 0", func() {
			Expect(main.CountRepeatedWords(input)).To(Equal(expected))
		})
	})

	When("input is only one word", func() {
		BeforeEach(func() {
			input = "hello"
			expected = 1
		})

		It("should return 1", func() {
			Expect(main.CountRepeatedWords(input)).To(Equal(expected))
		})
	})

	When("input has multiple words", func() {
		Context("and all words are the same", func() {
			BeforeEach(func() {
				input = "hello hello hello hello hello"
				expected = 5
			})

			It("should count the number of repeated words", func() {
				Expect(main.CountRepeatedWords(input)).To(Equal(expected))
			})
		})

		Context("and only the first word is repeated", func() {
			BeforeEach(func() {
				input = "hello world world world"
				expected = 1
			})

			It("should count only the first word", func() {
				Expect(main.CountRepeatedWords(input)).To(Equal(expected))
			})
		})
	})
})
