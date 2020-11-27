package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Delete Login Key", func() {
	Describe("Delete Login Key", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusOK).BodyString(`<deleteLoginKeyResponse>
					<requestId>aec19696-46c3-4841-967b-ce5399d151cf</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					</deleteLoginKeyResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should delete login key", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.DeleteLoginKey("packer-1512019276")

			Expect(err).To(BeNil())
			Expect(result.RequestID).To(Equal("aec19696-46c3-4841-967b-ce5399d151cf"))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
		})
	})

	Describe("Not exist login key name", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
					<returnCode>10406</returnCode>
					<returnMessage>The authentication key does not exist.</returnMessage>
					</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should fail if key name is not exist", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.DeleteLoginKey("test01912")

			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
			Expect(result.ReturnCode).To(Equal(10406))
			Expect(result.ReturnMessage).To(Equal("The authentication key does not exist."))
		})
	})

	Describe("keyName length", func() {
		It("should be larger than 3", func() {
			conn := NewConnection(accessKey, secretKey)
			_, err := conn.DeleteLoginKey("12")

			Expect(err.Error()).To(ContainSubstring("Length of KeyName should be min 3 or max 30"))
		})

		It("should be smaller than 30", func() {
			conn := NewConnection(accessKey, secretKey)
			_, err := conn.DeleteLoginKey(String(31))

			Expect(err.Error()).To(ContainSubstring("Length of KeyName should be min 3 or max 30"))
		})
	})
})
