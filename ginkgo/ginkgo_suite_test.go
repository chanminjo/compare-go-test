package ginkgo

import (
        "testing"

        . "github.com/onsi/ginkgo/v2"
        . "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
        RegisterFailHandler(Fail)
        RunSpecs(t, "Ginkgo Suite")
}

var _ = Describe("Books", func() {
        It("1+1=2", func() {

                Expect(Add(1, 3)).To(Equal(4))
        })
})
