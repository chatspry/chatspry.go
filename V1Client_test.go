package chatspry

import (
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

const (
	SessionID  = "ABC-123-DEFG-4567"
	UserName   = "Foo Bar"
	UserHandle = "foobar"
)

func strPtr(str string) *string { return &str }

func mockLogin(w http.ResponseWriter, r *http.Request) {
	resp := map[string]interface{}{
		"session": map[string]string{
			"id": SessionID,
		},
		"user": map[string]string{
			"id":     SessionID,
			"name":   UserName,
			"handle": UserHandle,
		},
	}
	json, _ := json.Marshal(&resp)
	w.Write(json)
}

var _ = Describe("V1Client", func() {
	It("should accept a base URL", func() {
		url := "foo.bar.baz/api"
		client := NewV1Client(url)

		Expect(client.BaseURL).To(BeEquivalentTo(url))
	})

	It("should be able to login", func() {
		http.Handle("/v1/session", http.HandlerFunc(mockLogin))
		server := httptest.NewServer(nil)
		defer server.Close()

		client := NewV1Client(server.URL)
		err := client.Login("foo", "bar")

		Expect(err).To(BeNil())
		Expect(client.SessionID).To(BeEquivalentTo(SessionID))
		Expect(client.User).NotTo(BeNil())
		Expect(client.User).To(BeEquivalentTo(&User{
			ID:     strPtr(SessionID),
			Name:   strPtr(UserName),
			Handle: strPtr(UserHandle),
		}))
	})
})
