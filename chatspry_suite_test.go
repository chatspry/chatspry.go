package chatspry

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestChatspryGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "chatspry.go Suite")
}
