package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gock "gopkg.in/h2non/gock.v1"
)

var _ = Describe("Get Root Password", func() {
	Describe("Get Root Password", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getRootPasswordResponse>
							<requestId>074fa0db-a289-4272-8ed5-9f94a0ae5173</requestId>
		  					<returnCode>0</returnCode>
							<returnMessage>success</returnMessage>
							<rootPassword>H53yn-TE+%25ee</rootPassword>
						</getRootPasswordResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get root password", func() {
			reqParams := new(RequestGetRootPassword)
			reqParams.ServerInstanceNo = "354246"
			reqParams.PrivateKey = "-----BEGIN RSA PRIVATE KEY-----\n" +
				"MIIEogIBAAKCAQEAjtfNZ9dtOUkVYCOmn9/ICKPOF3359qZ1hOeOvmkfINwxNlMs\n" +
				"6bp7U7UOyJZwjouL85jhcHlgR8CTOKlSR9SbQ0gGMBUJfCksv3VuzIpTL0S/8dn6\n" +
				"rcMKhPLBfrVA6R1r4ixJrvls4U4BBDqVU1bF10CDAqnrLWRZHMsSWC2jNx+av8cR\n" +
				"S74TX8tKGD3fLAl+3PcgY9rk5PYghe5DOEXRAtiBOFWhwr4SINFT6exKfg/lba1Q\n" +
				"2kEkeZ4t0Zxu1lUcsULxbyek0rm/STaCQDrU/rsITDVIdlJ1x5UgE7StU/EB2P1F\n" +
				"MM1GZlFtONzArOkQ10HXLpuYnr0LAyhgQQDLzwIDAQABAoIBADR5LDjIV5Rit5qA\n" +
				"njY+09HwssXIfmnQTn2LPtHN9iRxHxbwyR/3ZFSkv5CKS0DiHIBZGf/iZq8VHEAF\n" +
				"W3eNr5ewvq0tHTXiPlD7B7/V1KcZn1VuGPt6GcQrFU/Br92y9C26zTqc4BWIPTFb\n" +
				"f/2Ec9sWFMqGtyEOqvm707pW0Y1u3tetE46PytT/E7I0z6EGzXycvR1dvHvYKkr9\n" +
				"KCF7cbD6SaNqWQ+DfOdU6M+S8ifY7Bfkb4K88rly0DkdFADWSo6p0iC8+uRxHkzY\n" +
				"zWzQUNXfunF+iLpbw2CjHyEYDkP9Id1FiUpH6MzaSDvWsOg2FrofiCb3kPYqRygm\n" +
				"zo2QlCECgYEAzlzArqDbYyUTGaBQ6fK9mEN7HDWAtpPNQ5/Cf4YAHCIrJgM3mTdI\n" +
				"Hvyd2MvSwPwOWC+zPLeN9FJ5oiHdsWZQ3Id0bq80c8QN9hL2qX7kieMJAFcx6527\n" +
				"3roGozaGcrd8qJfLCxFFd9HLpKskGvs1SQMwuQByv0vce1mq8WBQGNECgYEAsTOy\n" +
				"Rb+YgV68ZI8EG9B083YrHAJtWA7HX8NfohEQMn8gUYYJtTOMXFGTcjkO3+rXHUNq\n" +
				"Yn3S2Xx1sWIANXVEMijYfLQ9zLyqdaTjpCqUNzNLsQa//KLHH7Wqejl/hIuC3WcZ\n" +
				"ovXNgh4eTvrwKA+mEyHTkZGuT1jcfDLEndIDwp8CgYAtTtYLQUwJz5kK9lTtmJsh\n" +
				"9CbygZi5/WeC4m+8ufoKM0JP/ULWw/l0vDgX4mBqeKzZldd/jDmBnB4Sh+b5zSIR\n" +
				"2TeqGZOlmzUJgw5fFnotAAj45ywRtDcMZsGQUidgxoJ7LaCp1GfxL8HLbLvnpKIF\n" +
				"Xsryw3NuJsTI3F4Y0vGjEQKBgE9XdlsEjqReZrwfzRcnHzZjkqnX7F2hQGQAdrF6\n" +
				"MmZpxW67NV2M1yv+2qYqJh7cZ7USmm+/Jsl6kLf7yVwmRYKwBz7O9VzWnPxP+B1p\n" +
				"czPEI///pbbYq4/nJ43ScOZIETBRGYl9xaadFbUHZsCJQTMP0ks7j3C3uQFjOeXw\n" +
				"Qv7BAoGAPZyrtfgpWNr6vlY17yTQii6M9t1eMhPpLyLnuQg5ENOHA2qZynrekFI+\n" +
				"c+TsJ5aFShvIc818YMZlg/yVO76z9FS0oPCQ2tutzX3881EJlstU6JFjk5j3SDHP\n" +
				"cHFcbmSZLpJXzJ0vJp2CH5l1uiMrF1vsEBf7uNjHAL1fx7zf424=\n" +
				"-----END RSA PRIVATE KEY-----"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetRootPassword(reqParams)

			Expect(err).To(BeNil())
			Expect(result.RequestID).To(Equal("074fa0db-a289-4272-8ed5-9f94a0ae5173"))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(result.RootPassword).To(Equal("H53yn-TE+%25ee"))
		})

	})

	Describe("Invalid Private Key", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
					<returnCode>24001</returnCode>
					<returnMessage>
					Invalid authentication key. Please check authentication key or input type and try again.
					</returnMessage>
					</responseError>`)
		})

		AfterEach(func() {
			gock.Off()
		})

		It("should fail if Server instance number is valid but, Private key is invalid", func() {
			reqParams := new(RequestGetRootPassword)
			reqParams.ServerInstanceNo = "354246"
			reqParams.PrivateKey = "-----BEGIN RSA PRIVATE KEY-----\n" +
				"fake private key\n" +
				"-----END RSA PRIVATE KEY-----"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetRootPassword(reqParams)

			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
			Expect(result.ReturnCode).To(Equal(24001))
			Expect(result.ReturnMessage).To(Equal("Invalid authentication key. Please check authentication key or input type and try again."))
		})
	})

	Describe("Invalid Server Instance Number", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
					<returnCode>1300</returnCode>
					<returnMessage>
					Please try your call again later. Temporarily out of service. If error continue, Please contact our customer service center.
					</returnMessage>
					</responseError>`)
		})

		AfterEach(func() {
			gock.Off()
		})

		It("should fail if Server instance number is invalid", func() {
			reqParams := new(RequestGetRootPassword)
			reqParams.ServerInstanceNo = "354246"
			reqParams.PrivateKey = "-----BEGIN RSA PRIVATE KEY-----\n" +
				"fake private key\n" +
				"-----END RSA PRIVATE KEY-----"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetRootPassword(reqParams)

			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
			Expect(result.ReturnCode).To(Equal(1300))
			Expect(result.ReturnMessage).To(Equal("Please try your call again later. Temporarily out of service. If error continue, Please contact our customer service center."))
		})
	})
})
