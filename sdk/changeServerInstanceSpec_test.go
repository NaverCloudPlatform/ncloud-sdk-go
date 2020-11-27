package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Change Server Instance Spec", func() {
	Describe("Change Server Instance Spec", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusOK).BodyString(`<changeServerInstanceSpecResponse>
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
							<serverInstanceStatusName>init</serverInstanceStatusName>
							<createDate>2017-12-01T20:08:40+0900</createDate>
							<uptime>2017-12-01T20:11:34+0900</uptime>
							<serverImageProductCode>SPSW0LINUX000046</serverImageProductCode>
							<serverProductCode>SPSVRSTAND000004</serverProductCode>
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
				</changeServerInstanceSpecResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})

		It("should change server instance spec", func() {
			reqParams := new(RequestChangeServerInstanceSpec)
			reqParams.ServerInstanceNo = "551227"
			reqParams.ServerProductCode = "SPSVRSTAND000004"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.ChangeServerInstanceSpec(reqParams)

			Expect(err).To(BeNil())
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.ServerInstanceList)).To(Equal(1))

			serverInstance := result.ServerInstanceList[0]
			Expect(serverInstance.ServerInstanceNo).To(Equal("551218"))
			Expect(serverInstance.ServerProductCode).To(Equal("SPSVRSTAND000004"))
			Expect(serverInstance.ServerName).To(Equal("wwef"))
			Expect(serverInstance.Region.RegionCode).To(Equal("KR"))
			Expect(serverInstance.Region.RegionName).To(Equal("Korea"))
			Expect(serverInstance.Zone.ZoneName).To(Equal("KR-1"))
			Expect(serverInstance.PlatformType.Code).To(Equal("LNX64"))
			Expect(serverInstance.ServerInstanceStatusName).To(Equal("init"))
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
			reqParams := new(RequestChangeServerInstanceSpec)
			reqParams.ServerInstanceNo = "aaaa"
			reqParams.ServerProductCode = "aaaa"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.ChangeServerInstanceSpec(reqParams)

			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
			Expect(result.ReturnCode).To(Equal(24120))
			Expect(result.ReturnMessage).To(Equal("The input parameter server instance number is invalid."))
		})
	})

	Describe("Invalid server product code", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
					<returnCode>10712</returnCode>
  					<returnMessage>Not found product. Please check your input product.  </returnMessage>
				</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})

		It("should fail if server instance number is invalid", func() {
			reqParams := new(RequestChangeServerInstanceSpec)
			reqParams.ServerInstanceNo = "123"
			reqParams.ServerProductCode = "invalid"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.ChangeServerInstanceSpec(reqParams)

			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
			Expect(result.ReturnCode).To(Equal(10712))
			Expect(result.ReturnMessage).To(Equal("Not found product. Please check your input product."))
		})
	})

	Describe("Unable to change", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
  					<returnCode>10503</returnCode>
  					<returnMessage>You may not change to spec of same level.</returnMessage>
				</responseError>`)

		})
		AfterEach(func() {
			gock.Off()
		})
		It("should fail if server product code is same level", func() {
			reqParams := new(RequestChangeServerInstanceSpec)
			reqParams.ServerInstanceNo = "123"
			reqParams.ServerProductCode = "SPSVRSTAND000004"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.ChangeServerInstanceSpec(reqParams)

			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
			Expect(result.ReturnCode).To(Equal(10503))
			Expect(result.ReturnMessage).To(Equal("You may not change to spec of same level."))
		})
	})

	Describe("The input parameter server instance number and server product code are required", func() {
		It("should be fail with no parameters", func() {
			reqParams := new(RequestChangeServerInstanceSpec)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.ChangeServerInstanceSpec(reqParams)

			Expect(err.Error()).To(ContainSubstring("ServerInstanceNo field is required"))
		})

		It("should be fail with no server product code parameter", func() {
			reqParams := new(RequestChangeServerInstanceSpec)
			reqParams.ServerInstanceNo = "123"
			conn := NewConnection(accessKey, secretKey)
			_, err := conn.ChangeServerInstanceSpec(reqParams)

			Expect(err.Error()).To(ContainSubstring("ServerProductCode field is required"))
		})
	})

})
