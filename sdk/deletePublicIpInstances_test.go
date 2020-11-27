package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Delete Public IP Instances", func() {
	Describe("Delete Public IP Instances", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusOK).BodyString(`<deletePublicIpInstancesResponse>
					<requestId>f6c7908e-936c-42ad-9e3b-73e4b0765f59</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<publicIpInstanceList>
					<publicIpInstance>
					<publicIpInstanceNo>388561</publicIpInstanceNo>
					<publicIp>192.168.120.175</publicIp>
					<publicIpDescription/>
					<createDate>2017-12-01T22:44:53+0900</createDate>
					<internetLineType>
					<code>PUBLC</code>
					<codeName>PUBLC</codeName>
					</internetLineType>
					<publicIpInstanceStatusName>terminated</publicIpInstanceStatusName>
					<publicIpInstanceStatus>
					<code>TERMT</code>
					<codeName>NET TERMINATED State</codeName>
					</publicIpInstanceStatus>
					<publicIpInstanceOperation>
					<code>NULL</code>
					<codeName>NET NULL OP</codeName>
					</publicIpInstanceOperation>
					<publicIpKindType>
					<code>GEN</code>
					<codeName>General</codeName>
					</publicIpKindType>
					<serverInstanceAssociatedWithPublicIp/>
					</publicIpInstance>
					</publicIpInstanceList>
					</deletePublicIpInstancesResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})

		It("should Delete Public IP instances", func() {
			reqParams := new(RequestDeletePublicIPInstances)
			reqParams.PublicIPInstanceNoList = []string{"388561"}

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.DeletePublicIPInstances(reqParams)

			Expect(err).To(BeNil())
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.PublicIPInstanceList)).To(Equal(1))

			publicIPInstance := result.PublicIPInstanceList[0]
			Expect(publicIPInstance.PublicIPInstanceNo).To(Equal("388561"))
			Expect(publicIPInstance.PublicIP).To(Equal("192.168.120.175"))
			Expect(publicIPInstance.PublicIPInstanceStatusName).To(Equal("terminated"))
		})
	})

	Describe("Invalid Public IP number", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
					<returnCode>10713</returnCode>
					<returnMessage>
					Not found contract information. Please check your input parameter.
					</returnMessage>
					</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should fail if Public IP instance number is invalid", func() {
			reqParams := new(RequestDeletePublicIPInstances)
			reqParams.PublicIPInstanceNoList = []string{"aaaa", "bbb"}

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.DeletePublicIPInstances(reqParams)

			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
			Expect(result.ReturnCode).To(Equal(10713))
			Expect(result.ReturnMessage).To(Equal("Not found contract information. Please check your input parameter."))
		})
	})

	Describe("Unable to delete", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
					<returnCode>24073</returnCode>
					<returnMessage>
					Unable to destroy the server since a public IP is associated with the server. First, please disassociate a public IP from the server. publicIpInstanceNo = 322396
					</returnMessage>
					</responseError>`)

		})
		AfterEach(func() {
			gock.Off()
		})
		It("should fail if public IP instance number is invalid", func() {
			reqParams := new(RequestDeletePublicIPInstances)
			reqParams.PublicIPInstanceNoList = []string{"388979"}

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.DeletePublicIPInstances(reqParams)

			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
			Expect(result.ReturnCode).To(Equal(24073))
			Expect(result.ReturnMessage).To(Equal("Unable to destroy the server since a public IP is associated with the server. First, please disassociate a public IP from the server. publicIpInstanceNo = 322396"))
		})
	})

	Describe("The input parameter public IP instance number is required", func() {
		It("should be fail with no parameters", func() {
			reqParams := new(RequestDeletePublicIPInstances)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.DeletePublicIPInstances(reqParams)

			Expect(err.Error()).To(ContainSubstring("Required field is not specified"))
		})
	})
})
