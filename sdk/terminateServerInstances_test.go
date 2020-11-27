package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Terminate Server Instances", func() {
	Describe("Terminate Server Instances", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusOK).BodyString(`<terminateServerInstancesResponse>
					<requestId>8c5572d4-8895-4fad-afca-2a2762fb8b70</requestId>
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
								<code>NSTOP</code>
								<codeName>Server normal stopped state</codeName>
							</serverInstanceStatus>
							<serverInstanceOperation>
								<code>TERMT</code>
								<codeName>Server TERMINATE OP</codeName>
							</serverInstanceOperation>
							<serverInstanceStatusName>terminating</serverInstanceStatusName>
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
				</terminateServerInstancesResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})

		It("should stop server instances", func() {
			reqParams := new(RequestTerminateServerInstances)
			reqParams.ServerInstanceNoList = []string{"551218"}

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.TerminateServerInstances(reqParams)

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
			Expect(serverinstance.ServerInstanceStatusName).To(Equal("terminating"))
		})
	})

	Describe("Invalid server instance number", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
					<returnCode>10713</returnCode>
					<returnMessage>Not found contract information. Please check your input parameter.</returnMessage>
				</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should fail if server instance number is invalid", func() {
			reqParams := new(RequestTerminateServerInstances)
			reqParams.ServerInstanceNoList = []string{"551230"}

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.TerminateServerInstances(reqParams)

			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
			Expect(result.ReturnCode).To(Equal(10713))
			Expect(result.ReturnMessage).To(Equal("Not found contract information. Please check your input parameter."))
		})
	})

	Describe("Unable to terminate", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
					<returnCode>25022</returnCode>
					<returnMessage>Unable to return the server since (other) user is either operating target server or an error in the target server. Please confirm the status of server/ storage.</returnMessage>
				</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should fail if server instance status is not stopped", func() {
			reqParams := new(RequestTerminateServerInstances)
			reqParams.ServerInstanceNoList = []string{"551233"}

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.TerminateServerInstances(reqParams)

			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
			Expect(result.ReturnCode).To(Equal(25022))
			Expect(result.ReturnMessage).To(Equal("Unable to return the server since (other) user is either operating target server or an error in the target server. Please confirm the status of server/ storage."))
		})
	})

	Describe("The input parameter server instance number is required", func() {
		It("should be fail with no parameters", func() {
			reqParams := new(RequestTerminateServerInstances)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.TerminateServerInstances(reqParams)

			Expect(err.Error()).To(ContainSubstring("serverInstanceNoList is required"))
		})
	})
})
