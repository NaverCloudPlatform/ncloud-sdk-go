package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create Public IP", func() {
	Describe("Create Public IP", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusOK).BodyString(`<createPublicIpInstanceResponse>
					<requestId>f2807cbe-f876-4fe5-b42e-2d0f96b68a8b</requestId>
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
					<publicIpInstanceStatusName>using</publicIpInstanceStatusName>
					<publicIpInstanceStatus>
					<code>CREAT</code>
					<codeName>NET CREAT State</codeName>
					</publicIpInstanceStatus>
					<publicIpInstanceOperation>
					<code>USE</code>
					<codeName>NET USE OP</codeName>
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
					<zone>
          <zoneNo>3</zoneNo>
          <zoneName>KR-2</zoneName>
          <zoneCode>KR-2</zoneCode>
          <zoneDescription>평촌 zone</zoneDescription>
          <regionNo>1</regionNo>
          </zone>
					</publicIpInstance>
					</publicIpInstanceList>
					</createPublicIpInstanceResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should create Public IP", func() {
			reqParams := new(RequestCreatePublicIPInstance)
			reqParams.ServerInstanceNo = "551236"
			reqParams.PublicIPDescription = "public ip desc by SDK go"
			reqParams.InternetLineTypeCode = "PUBLC"
			reqParams.ZoneNo = "3"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.CreatePublicIPInstance(reqParams)

			Expect(err).To(BeNil())
			publicIPInstance := result.PublicIPInstanceList[0]

			Expect(result.RequestID).To(Equal("f2807cbe-f876-4fe5-b42e-2d0f96b68a8b"))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.PublicIPInstanceList)).To(Equal(1))
			Expect(publicIPInstance.PublicIP).To(Equal("192.168.120.117"))
			Expect(publicIPInstance.PublicIPInstanceStatusName).To(Equal("using"))
			Expect(publicIPInstance.Zone.ZoneNo).To(Equal("3"))
			Expect(publicIPInstance.Zone.ZoneName).To(Equal("KR-2"))
			Expect(publicIPInstance.Zone.ZoneCode).To(Equal("KR-2"))
			Expect(publicIPInstance.Zone.ZoneDescription).To(Equal("평촌 zone"))
			Expect(publicIPInstance.Zone.RegionNo).To(Equal("1"))
		})
	})

	Describe("Unable to create", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
					<returnCode>24120</returnCode>
					<returnMessage>
					The input parameter server instance number is invalid.
					</returnMessage>
					</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should fail if The input parameter server instance number is invalid", func() {
			reqParams := new(RequestCreatePublicIPInstance)
			reqParams.ServerInstanceNo = "551236"
			reqParams.PublicIPDescription = "public ip desc by SDK go"
			reqParams.InternetLineTypeCode = "PUBLC"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.CreatePublicIPInstance(reqParams)

			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
			Expect(result.ReturnCode).To(Equal(24120))
			Expect(result.ReturnMessage).To(Equal("The input parameter server instance number is invalid."))
		})
	})
})
