package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Access Control Rule List", func() {
	Describe("Get Access Control Rule List", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getAccessControlRuleListResponse>
					<requestId>ca53a246-3af4-4d25-9327-7cde849cbabf</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>6</totalRows>
					<accessControlRuleList>
					  <accessControlRule>
						<accessControlRuleConfigurationNo>25363</accessControlRuleConfigurationNo>
						<protocolType>
						  <code>TCP</code>
						  <codeName>tcp</codeName>
						</protocolType>
						<sourceIp>0.0.0.0/0</sourceIp>
						<destinationPort>22</destinationPort>
						<accessControlRuleDescription></accessControlRuleDescription>
					  </accessControlRule>
					  <accessControlRule>
						<accessControlRuleConfigurationNo>25364</accessControlRuleConfigurationNo>
						<protocolType>
						  <code>TCP</code>
						  <codeName>tcp</codeName>
						</protocolType>
						<sourceIp>0.0.0.0/0</sourceIp>
						<destinationPort>3389</destinationPort>
						<accessControlRuleDescription></accessControlRuleDescription>
					  </accessControlRule>
					  <accessControlRule>
						<accessControlRuleConfigurationNo>25365</accessControlRuleConfigurationNo>
						<protocolType>
						  <code>TCP</code>
						  <codeName>tcp</codeName>
						</protocolType>
						<sourceIp></sourceIp>
						<sourceAccessControlRuleConfigurationNo>4964</sourceAccessControlRuleConfigurationNo>
						<sourceAccessControlRuleName>ncloud-default-acg</sourceAccessControlRuleName>
						<destinationPort>1-65535</destinationPort>
						<accessControlRuleDescription></accessControlRuleDescription>
					  </accessControlRule>
					  <accessControlRule>
						<accessControlRuleConfigurationNo>25366</accessControlRuleConfigurationNo>
						<protocolType>
						  <code>UDP</code>
						  <codeName>udp</codeName>
						</protocolType>
						<sourceIp></sourceIp>
						<sourceAccessControlRuleConfigurationNo>4964</sourceAccessControlRuleConfigurationNo>
						<sourceAccessControlRuleName>ncloud-default-acg</sourceAccessControlRuleName>
						<destinationPort>1-65535</destinationPort>
						<accessControlRuleDescription></accessControlRuleDescription>
					  </accessControlRule>
					  <accessControlRule>
						<accessControlRuleConfigurationNo>25367</accessControlRuleConfigurationNo>
						<protocolType>
						  <code>ICMP</code>
						  <codeName>icmp</codeName>
						</protocolType>
						<sourceIp></sourceIp>
						<sourceAccessControlRuleConfigurationNo>4964</sourceAccessControlRuleConfigurationNo>
						<sourceAccessControlRuleName>ncloud-default-acg</sourceAccessControlRuleName>
						<accessControlRuleDescription></accessControlRuleDescription>
					  </accessControlRule>
					  <accessControlRule>
						<accessControlRuleConfigurationNo>99006</accessControlRuleConfigurationNo>
						<protocolType>
						  <code>TCP</code>
						  <codeName>tcp</codeName>
						</protocolType>
						<sourceIp>0.0.0.0/0</sourceIp>
						<destinationPort>5985-5986</destinationPort>
						<accessControlRuleDescription>winrm 접속 포트 허용</accessControlRuleDescription>
					  </accessControlRule>
					</accessControlRuleList>
				  </getAccessControlRuleListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get Acess Control Rule list", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetAccessControlRuleList("4964")

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(6))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.AccessControlRuleList)).To(Equal(6))

			accessControlRule := result.AccessControlRuleList[0]
			Expect(accessControlRule.AccessControlRuleConfigurationNo).To(Equal("25363"))
			Expect(accessControlRule.ProtocolType.Code).To(Equal("TCP"))
			Expect(accessControlRule.ProtocolType.CodeName).To(Equal("tcp"))
			Expect(accessControlRule.SourceIP).To(Equal("0.0.0.0/0"))
			Expect(accessControlRule.DestinationPort).To(Equal("22"))

			accessControlRule = result.AccessControlRuleList[5]
			Expect(accessControlRule.AccessControlRuleConfigurationNo).To(Equal("99006"))
			Expect(accessControlRule.ProtocolType.Code).To(Equal("TCP"))
			Expect(accessControlRule.ProtocolType.CodeName).To(Equal("tcp"))
			Expect(accessControlRule.SourceIP).To(Equal("0.0.0.0/0"))
			Expect(accessControlRule.DestinationPort).To(Equal("5985-5986"))
		})
	})

	Describe("There is no Access Control Rule list", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
					<returnCode>24076</returnCode>
					<returnMessage>It is ACG(access control group) you do not own. Serial number - 49645</returnMessage>
				</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get No Access Control Rule list", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetAccessControlRuleList("49645")

			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
			Expect(result.ReturnCode).To(Equal(24076))
			Expect(result.ReturnMessage).To(ContainSubstring("It is ACG(access control group) you do not own. Serial number -"))
		})
	})
})
