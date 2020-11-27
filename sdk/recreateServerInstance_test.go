package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Recreate Server Instance", func() {
	Describe("Recreate linux server instance", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusOK).BodyString(`<recreateServerInstanceResponse>
					<requestId>77868086-533e-4b31-8722-b6e84dfdcddd</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<serverInstanceList>
						<serverInstance>
							<serverInstanceNo>550601</serverInstanceNo>
							<serverName>svr-9bb4ded6ecfc049</serverName>
							<serverDescription></serverDescription>
							<cpuCount>1</cpuCount>
							<memorySize>1073741824</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<loginKeyName>packer-1512025059</loginKeyName>
							<isFeeChargingMonitoring>false</isFeeChargingMonitoring>
							<publicIp></publicIp>
							<privateIp>10.33.25.170</privateIp>
							<serverImageName>centos-7.3-64</serverImageName>
							<serverInstanceStatus>
								<code>INIT</code>
								<codeName>Server init state</codeName>
							</serverInstanceStatus>
							<serverInstanceOperation>
								<code>NULL</code>
								<codeName>Server NULL OP</codeName>
							</serverInstanceOperation>
							<serverInstanceStatusName>init</serverInstanceStatusName>
							<createDate>2017-12-01T13:50:46+0900</createDate>
							<uptime>2017-12-01T13:50:46+0900</uptime>
							<serverImageProductCode>SPSW0LINUX000046</serverImageProductCode>
							<serverProductCode>SPSVRSTAND000056</serverProductCode>
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
							<baseBlockStroageDiskDetailType>
								<code>HDD</code>
								<codeName>HDD</codeName>
							</baseBlockStroageDiskDetailType>
							<userData></userData>
							<accessControlGroupList/>
						</serverInstance>
					</serverInstanceList>
				</recreateServerInstanceResponse>`)

		})
		AfterEach(func() {
			gock.Off()
		})
		It("should recreate linux server instance", func() {
			reqParams := new(RequestRecreateServerInstance)
			reqParams.ServerInstanceNo = "550601"
			reqParams.ChangeServerImageProductCode = "SPSW0LINUX000046"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.RecreateServerInstance(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.ServerInstanceList)).To(Equal(1))

			serverInstance := result.ServerInstanceList[0]
			Expect(serverInstance.ServerInstanceNo).To(Equal("550601"))
			Expect(serverInstance.ServerName).To(Equal("svr-9bb4ded6ecfc049"))
			Expect(serverInstance.Region.RegionCode).To(Equal("KR"))
			Expect(serverInstance.Region.RegionName).To(Equal("Korea"))
			Expect(serverInstance.Zone.ZoneName).To(Equal("KR-1"))
			Expect(serverInstance.PlatformType.Code).To(Equal("LNX64"))
			Expect(serverInstance.ServerInstanceStatusName).To(Equal("init"))
			Expect(serverInstance.PortForwardingPublicIP).To(Equal("220.230.115.116"))
			Expect(serverInstance.Uptime).To(Equal("2017-12-01T13:50:46+0900"))
			Expect(serverInstance.BaseBlockStorageDiskType.Code).To(Equal("NET"))
		})
	})

	Describe("Authorize fail", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusUnauthorized).BodyString(`<responseError>
					<returnCode>800</returnCode>
					<returnMessage>Expired url.</returnMessage>
					</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should be failed by authorization fail", func() {
			reqParams := new(RequestRecreateServerInstance)
			reqParams.ServerInstanceNo = "550601"
			reqParams.ChangeServerImageProductCode = "SPSW0LINUX000046"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.RecreateServerInstance(reqParams)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Expired url."))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})

	Describe("Check Arguments", func() {
		It("should be error : 'ServerInstanceNo field is required'", func() {
			reqParams := new(RequestRecreateServerInstance)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.RecreateServerInstance(reqParams)

			Expect(err.Error()).To(ContainSubstring("ServerInstanceNo field is required"))
		})

		It("should be error : 'Length of ServerInstanceName should be min 3 or max 30'", func() {
			reqParams := new(RequestRecreateServerInstance)
			reqParams.ServerInstanceNo = "550601"
			reqParams.ServerInstanceName = String(31)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.RecreateServerInstance(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of ServerInstanceName to be in the range (3 - 30)"))

			reqParams.ServerInstanceName = "Se"

			_, err = conn.RecreateServerInstance(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of ServerInstanceName to be in the range (3 - 30)"))
		})

		It("should be error : 'Length of ChangeServerImageProductCode should be min 1 or max 20'", func() {
			reqParams := new(RequestRecreateServerInstance)
			reqParams.ServerInstanceNo = "550601"
			reqParams.ChangeServerImageProductCode = String(21)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.RecreateServerInstance(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of ChangeServerImageProductCode to be in the range (0 - 20)"))
		})

	})
})
