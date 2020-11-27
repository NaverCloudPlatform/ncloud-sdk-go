package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reboot Server Instances", func() {
	Describe("Reboot Server Instances", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusOK).BodyString(`<rebootServerInstancesResponse>
					<requestId>3c8797b6-f48c-4c1e-bb7f-d323f187bf14</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<serverInstanceList>
						<serverInstance>
							<serverInstanceNo>551218</serverInstanceNo>
							<serverName>wwef</serverName>
							<serverDescription></serverDescription>
							<cpuCount>2</cpuCount>
							<memorySize>4294967296</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<loginKeyName>packer-1512025059</loginKeyName>
							<isFeeChargingMonitoring>false</isFeeChargingMonitoring>
							<publicIp></publicIp>
							<privateIp>10.33.1.201</privateIp>
							<serverImageName>centos-7.3-64</serverImageName>
							<serverInstanceStatus>
								<code>RUN</code>
								<codeName>Server run state</codeName>
							</serverInstanceStatus>
							<serverInstanceOperation>
								<code>SHTDN</code>
								<codeName>Server SHUTDOWN OP</codeName>
							</serverInstanceOperation>
							<serverInstanceStatusName>shutting down</serverInstanceStatusName>
							<createDate>2017-12-01T20:08:40+0900</createDate>
							<uptime>2017-12-01T20:11:34+0900</uptime>
							<serverImageProductCode>SPSW0LINUX000046</serverImageProductCode>
							<serverProductCode>SPSVRSSD00000003</serverProductCode>
							<isProtectServerTermination>false</isProtectServerTermination>
							<portForwardingPublicIp>220.230.115.116</portForwardingPublicIp>
							<zone>
								<zoneNo>2</zoneNo>
								<zoneName>KR-1</zoneName>
								<zoneDescription>가산 NANG zone</zoneDescription>
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
							<userData></userData>
							<accessControlGroupList>
								<accessControlGroup>
									<accessControlGroupConfigurationNo>4964</accessControlGroupConfigurationNo>
									<accessControlGroupName>ncloud-default-acg</accessControlGroupName>
									<accessControlGroupDescription>Default AccessControlGroup</accessControlGroupDescription>
									<isDefault>true</isDefault>
									<createDate>2017-02-23T10:25:39+0900</createDate>
								</accessControlGroup>
							</accessControlGroupList>
						</serverInstance>
					</serverInstanceList>
				</rebootServerInstancesResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})

		It("should reboot server instances", func() {
			reqParams := new(RequestRebootServerInstances)
			reqParams.ServerInstanceNoList = []string{"551227"}

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.RebootServerInstances(reqParams)

			Expect(err).To(BeNil())
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.ServerInstanceList)).To(Equal(1))

			serverinstance := result.ServerInstanceList[0]
			Expect(serverinstance.ServerInstanceNo).To(Equal("551218"))
			Expect(serverinstance.ServerName).To(Equal("wwef"))
			Expect(serverinstance.Region.RegionCode).To(Equal("KR"))
			Expect(serverinstance.Region.RegionName).To(Equal("Korea"))
			Expect(serverinstance.Zone.ZoneName).To(Equal("KR-1"))
			Expect(serverinstance.PlatformType.Code).To(Equal("LNX64"))
			Expect(serverinstance.ServerInstanceStatusName).To(Equal("shutting down"))
		})
	})

	Describe("Invalid server instance number", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
					<returnCode>24120</returnCode>
					<returnMessage>The input parameter server instance number is invalid. </returnMessage>
				</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should fail if server instance number is invalid", func() {
			reqParams := new(RequestRebootServerInstances)
			reqParams.ServerInstanceNoList = []string{"aaaa", "bbb"}

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.RebootServerInstances(reqParams)

			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
			Expect(result.ReturnCode).To(Equal(24120))
			Expect(result.ReturnMessage).To(Equal("The input parameter server instance number is invalid."))
		})
	})

	Describe("Unable to reboot", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
					<returnCode>25041</returnCode>
					<returnMessage>Unable to perform operation command since (other) user is either manipulating the target server or a problem in target server. Please confirm server status once again. </returnMessage>
				</responseError>`)

		})
		AfterEach(func() {
			gock.Off()
		})
		It("should fail if server instance number is invalid", func() {
			reqParams := new(RequestRebootServerInstances)
			reqParams.ServerInstanceNoList = []string{"551218", "551227"}

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.RebootServerInstances(reqParams)

			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
			Expect(result.ReturnCode).To(Equal(25041))
			Expect(result.ReturnMessage).To(Equal("Unable to perform operation command since (other) user is either manipulating the target server or a problem in target server. Please confirm server status once again."))
		})
	})

	Describe("The input parameter server instance number is required", func() {
		It("should be fail with no parameters", func() {
			reqParams := new(RequestRebootServerInstances)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.RebootServerInstances(reqParams)

			Expect(err.Error()).To(ContainSubstring("serverInstanceNoList is required"))
		})
	})
})
