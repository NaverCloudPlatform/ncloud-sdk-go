package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Disassociate Public IP From Server Instance", func() {
	Describe("Disassociate Public IP From Server Instance", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusOK).BodyString(`<disassociatePublicIpFromServerInstanceResponse>
					<requestId>bd83cf83-0361-4c74-aa2e-3e5243a58f1d</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<publicIpInstanceList>
					<publicIpInstance>
					<publicIpInstanceNo>388979</publicIpInstanceNo>
					<publicIp>192.168.120.117</publicIp>
					<publicIpDescription>packer</publicIpDescription>
					<createDate>2017-12-02T14:50:48+0900</createDate>
					<internetLineType>
					<code>PUBLC</code>
					<codeName>PUBLC</codeName>
					</internetLineType>
					<publicIpInstanceStatusName>disusing</publicIpInstanceStatusName>
					<publicIpInstanceStatus>
					<code>USED</code>
					<codeName>NET USED state</codeName>
					</publicIpInstanceStatus>
					<publicIpInstanceOperation>
					<code>DISUS</code>
					<codeName>NET DISUSE OP</codeName>
					</publicIpInstanceOperation>
					<publicIpKindType>
					<code>GEN</code>
					<codeName>General</codeName>
					</publicIpKindType>
					<serverInstanceAssociatedWithPublicIp>
					<serverInstanceNo>387235</serverInstanceNo>
					<serverName>s-541onrn9sndwwb004</serverName>
					<serverDescription>Auto scaling group 'g0927' 소속의 자동 생성된 서버</serverDescription>
					<cpuCount>2</cpuCount>
					<memorySize>4294967296</memorySize>
					<baseBlockStorageSize>53687091200</baseBlockStorageSize>
					<platformType>
					<code>LNX64</code>
					<codeName>Linux 64 Bit</codeName>
					</platformType>
					<loginKeyName>safdsfa</loginKeyName>
					<isFeeChargingMonitoring>false</isFeeChargingMonitoring>
					<publicIp>192.168.120.117</publicIp>
					<privateIp>10.113.245.85</privateIp>
					<serverImageName>centos-6.3-64</serverImageName>
					<serverInstanceStatus>
					<code>RUN</code>
					<codeName>Server RUN State</codeName>
					</serverInstanceStatus>
					<serverInstanceOperation>
					<code>NULL</code>
					<codeName>Server NULL OP</codeName>
					</serverInstanceOperation>
					<serverInstanceStatusName>running</serverInstanceStatusName>
					<createDate>2017-11-29T15:05:21+0900</createDate>
					<uptime>2017-11-29T15:09:53+0900</uptime>
					<serverImageProductCode>SPSW0LINUX000031</serverImageProductCode>
					<serverProductCode>SPSVRSSD00000003</serverProductCode>
					<isProtectServerTermination>false</isProtectServerTermination>
					<portForwardingPublicIp>192.168.120.111</portForwardingPublicIp>
					<zone>
					<zoneNo>2</zoneNo>
					<zoneName>KR-1</zoneName>
					<zoneDescription>KR-1 zone</zoneDescription>
					</zone>
					<region>
					<regionNo>1</regionNo>
					<regionCode>KR</regionCode>
					<regionName>Korea</regionName>
					</region>
					<baseBlockStorageDiskType>
					<code>NET</code>
					<codeName>Network Storage</codeName>
					</baseBlockStorageDiskType>
					<baseBlockStroageDiskDetailType>
					<code>SSD</code>
					<codeName>SSD</codeName>
					</baseBlockStroageDiskDetailType>
					<internetLineType>
					<code>PUBLC</code>
					<codeName>PUBLC</codeName>
					</internetLineType>
					<userData/>
					<accessControlGroupList>
					<accessControlGroup>
					<accessControlGroupConfigurationNo>4054</accessControlGroupConfigurationNo>
					<accessControlGroupName>brick-test-acg11</accessControlGroupName>
					<accessControlGroupDescription>bbbbbbb</accessControlGroupDescription>
					<isDefault>false</isDefault>
					<createDate>2017-03-29T23:03:39+0900</createDate>
					</accessControlGroup>
					</accessControlGroupList>
					</serverInstanceAssociatedWithPublicIp>
					</publicIpInstance>
					</publicIpInstanceList>
					</disassociatePublicIpFromServerInstanceResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should disassociate public IP", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.DisassociatePublicIP("388979")

			Expect(err).To(BeNil())
			Expect(result.RequestID).To(Equal("bd83cf83-0361-4c74-aa2e-3e5243a58f1d"))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
		})
	})

	Describe("The input parameter public IP instance number is required", func() {
		It("should be fail with no parameters", func() {
			conn := NewConnection(accessKey, secretKey)
			_, err := conn.DisassociatePublicIP("")

			Expect(err.Error()).To(ContainSubstring("Required field is not specified"))
		})
	})

	Describe("Unable to disassociate", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
						<returnCode>28102</returnCode>
						<returnMessage>
						The public IP is already disassociated from the server.
						</returnMessage>
						</responseError>`)

		})
		AfterEach(func() {
			gock.Off()
		})
		It("should fail if public IP instance number is invalid", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.DisassociatePublicIP("388979")

			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
			Expect(result.ReturnCode).To(Equal(28102))
			Expect(result.ReturnMessage).To(Equal("The public IP is already disassociated from the server."))
		})
	})
})
