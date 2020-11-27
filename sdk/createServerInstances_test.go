package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create Server Instance", func() {
	Describe("Create Two Linux Server Instance", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusOK).BodyString(`<createServerInstancesResponse>
					<requestId>eb22dc06-ff48-4d05-933d-22ca85102508</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>2</totalRows>
					<serverInstanceList>
						<serverInstance>
							<serverInstanceNo>550718</serverInstanceNo>
							<serverName>svr-9bb4df5a0e07841011</serverName>
							<serverDescription></serverDescription>
							<cpuCount>2</cpuCount>
							<memorySize>17179869184</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<loginKeyName>packer-1512025059</loginKeyName>
							<isFeeChargingMonitoring>false</isFeeChargingMonitoring>
							<publicIp></publicIp>
							<privateIp>10.33.3.29</privateIp>
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
							<createDate>2017-12-01T14:36:41+0900</createDate>
							<uptime>2017-12-01T14:36:41+0900</uptime>
							<serverImageProductCode>SPSW0LINUX000046</serverImageProductCode>
							<serverProductCode>SPSVRSSD00000011</serverProductCode>
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
							<accessControlGroupList/>
						</serverInstance>
						<serverInstance>
							<serverInstanceNo>550715</serverInstanceNo>
							<serverName>svr-9bb4df5a0e07841010</serverName>
							<serverDescription></serverDescription>
							<cpuCount>2</cpuCount>
							<memorySize>17179869184</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<loginKeyName>packer-1512025059</loginKeyName>
							<isFeeChargingMonitoring>false</isFeeChargingMonitoring>
							<publicIp></publicIp>
							<privateIp>10.34.18.113</privateIp>
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
							<createDate>2017-12-01T14:36:40+0900</createDate>
							<uptime>2017-12-01T14:36:40+0900</uptime>
							<serverImageProductCode>SPSW0LINUX000046</serverImageProductCode>
							<serverProductCode>SPSVRSSD00000011</serverProductCode>
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
							<accessControlGroupList/>
						</serverInstance>
					</serverInstanceList>
				</createServerInstancesResponse>`)

		})
		AfterEach(func() {
			gock.Off()
		})
		It("should create two linux server instances", func() {
			reqParams := new(RequestCreateServerInstance)
			reqParams.ServerImageProductCode = "SPSW0LINUX000046"
			reqParams.ServerProductCode = "SPSVRSSD00000011"
			reqParams.ServerCreateCount = 2
			reqParams.ServerCreateStartNo = 10

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.CreateServerInstances(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(2))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.ServerInstanceList)).To(Equal(2))

			serverinstance := result.ServerInstanceList[0]
			Expect(serverinstance.ServerInstanceNo).To(Equal("550718"))
			Expect(serverinstance.ServerName).To(Equal("svr-9bb4df5a0e07841011"))
			Expect(serverinstance.Region.RegionCode).To(Equal("KR"))
			Expect(serverinstance.Region.RegionName).To(Equal("Korea"))
			Expect(serverinstance.Zone.ZoneName).To(Equal("KR-1"))
			Expect(serverinstance.PlatformType.Code).To(Equal("LNX64"))
			Expect(serverinstance.ServerInstanceStatusName).To(Equal("init"))
			Expect(serverinstance.PortForwardingPublicIP).To(Equal("220.230.115.116"))
			Expect(serverinstance.Uptime).To(Equal("2017-12-01T14:36:41+0900"))
			Expect(serverinstance.BaseBlockStorageDiskType.Code).To(Equal("NET"))

			serverinstance = result.ServerInstanceList[1]
			Expect(serverinstance.ServerInstanceNo).To(Equal("550715"))
			Expect(serverinstance.ServerName).To(Equal("svr-9bb4df5a0e07841010"))
			Expect(serverinstance.Region.RegionCode).To(Equal("KR"))
			Expect(serverinstance.Region.RegionName).To(Equal("Korea"))
			Expect(serverinstance.Zone.ZoneName).To(Equal("KR-1"))
			Expect(serverinstance.PlatformType.Code).To(Equal("LNX64"))
			Expect(serverinstance.ServerInstanceStatusName).To(Equal("init"))
			Expect(serverinstance.PortForwardingPublicIP).To(Equal("220.230.115.116"))
			Expect(serverinstance.Uptime).To(Equal("2017-12-01T14:36:40+0900"))
			Expect(serverinstance.BaseBlockStorageDiskType.Code).To(Equal("NET"))
		})
	})

	Describe("Create one linux server instance", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusOK).BodyString(`<createServerInstancesResponse>
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
				</createServerInstancesResponse>`)

		})
		AfterEach(func() {
			gock.Off()
		})
		It("should create new linux server instance", func() {
			reqParams := new(RequestCreateServerInstance)
			reqParams.ServerImageProductCode = "SPSW0LINUX000046"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.CreateServerInstances(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.ServerInstanceList)).To(Equal(1))

			serverinstance := result.ServerInstanceList[0]
			Expect(serverinstance.ServerInstanceNo).To(Equal("550601"))
			Expect(serverinstance.ServerName).To(Equal("svr-9bb4ded6ecfc049"))
			Expect(serverinstance.Region.RegionCode).To(Equal("KR"))
			Expect(serverinstance.Region.RegionName).To(Equal("Korea"))
			Expect(serverinstance.Zone.ZoneName).To(Equal("KR-1"))
			Expect(serverinstance.PlatformType.Code).To(Equal("LNX64"))
			Expect(serverinstance.ServerInstanceStatusName).To(Equal("init"))
			Expect(serverinstance.PortForwardingPublicIP).To(Equal("220.230.115.116"))
			Expect(serverinstance.Uptime).To(Equal("2017-12-01T13:50:46+0900"))
			Expect(serverinstance.BaseBlockStorageDiskType.Code).To(Equal("NET"))
		})
	})

	Describe("Authorize fail", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusUnauthorized).BodyString(`<responseError>
					<returnCode>800</returnCode>
					<returnMessage>Invalid consumerKey</returnMessage>
					</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should be failed : 401 Unauthorized", func() {
			reqParams := new(RequestCreateServerInstance)
			reqParams.ServerImageProductCode = "SPSW0LINUX000046"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.CreateServerInstances(reqParams)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Invalid consumerKey"))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})

	Describe("No arguments", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
					<returnCode>900</returnCode>
					<returnMessage>Required field is not specified. location : serverImageProductCode or memberServerImageNo.</returnMessage>
				</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should be failed : 400 Bad Request", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.CreateServerInstances(nil)

			Expect(result.ReturnCode).To(Equal(900))
			Expect(result.ReturnMessage).To(Equal("Required field is not specified. location : serverImageProductCode or memberServerImageNo."))
			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
		})
	})

	Describe("creation limitation case", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
					<returnCode>10102</returnCode>
					<returnMessage>
					Unable to create server anymore since creation limitation setting has been exceeded. Creation limitation: 100 Created : 102
					</returnMessage>
					</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should be failed : 400 Bad Request", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.CreateServerInstances(nil)

			Expect(result.ReturnCode).To(Equal(10102))
			Expect(result.ReturnMessage).To(Equal("Unable to create server anymore since creation limitation setting has been exceeded. Creation limitation: 100 Created : 102"))
			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
		})
	})

	Describe("Check Arguments", func() {
		It("should be error : 'Length of ServerImageProductCode should be min 1 or max 20'", func() {
			reqParams := new(RequestCreateServerInstance)
			reqParams.ServerImageProductCode = String(21)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateServerInstances(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of ServerImageProductCode to be in the range (0 - 20)"))
		})

		It("should be error : 'Length of ServerProductCode should be min  or max 20'", func() {
			reqParams := new(RequestCreateServerInstance)
			reqParams.ServerProductCode = String(21)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateServerInstances(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of ServerProductCode to be in the range (0 - 20)"))
		})

		It("should be error : 'Length of ServerName should be min 1 or max 20'", func() {
			reqParams := new(RequestCreateServerInstance)
			reqParams.ServerName = String(31)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateServerInstances(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of ServerName to be in the range (3 - 30)"))

			reqParams.ServerName = "Se"

			_, err = conn.CreateServerInstances(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of ServerName to be in the range (3 - 30)"))
		})

		It("should be error : 'Length of ServerDescription should be min 1 or max 1000'", func() {
			reqParams := new(RequestCreateServerInstance)
			reqParams.ServerDescription = String(1001)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateServerInstances(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of ServerDescription to be in the range (0 - 1000)"))
		})

		It("should be error : 'Length of LoginKeyName should be min 3 or max 30'", func() {
			reqParams := new(RequestCreateServerInstance)
			reqParams.LoginKeyName = String(2)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateServerInstances(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of LoginKeyName to be in the range (3 - 30)"))

			reqParams.LoginKeyName = String(31)
			_, err = conn.CreateServerInstances(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of LoginKeyName to be in the range (3 - 30)"))
		})

		It("should be error : 'ServerCreateCount should be min 1 or max 20'", func() {
			reqParams := new(RequestCreateServerInstance)
			reqParams.ServerCreateCount = 21

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateServerInstances(reqParams)

			Expect(err.Error()).To(ContainSubstring("\"ServerCreateCount\" cannot be higher than 20: 21"))
		})

		It("should be error : 'Sum of ServerCreateCount and ServerCreateStartNo should be less than 1000'", func() {
			reqParams := new(RequestCreateServerInstance)
			reqParams.ServerCreateCount = 1
			reqParams.ServerCreateStartNo = 1000

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateServerInstances(reqParams)

			Expect(err.Error()).To(Equal("\"Sum of ServerCreateCount and ServerCreateStartNo\" cannot be higher than 1000: 1001"))
		})

		It("should be error : 'InternetLineTypeCode should be PUBLC or GLBL'", func() {
			reqParams := new(RequestCreateServerInstance)
			reqParams.InternetLineTypeCode = String(5)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateServerInstances(reqParams)

			Expect(err.Error()).To(Equal("InternetLineTypeCode should be PUBLC or GLBL"))
		})

		It("should be error : 'FeeSystemTypeCode should be FXSUM or MTRAT'", func() {
			reqParams := new(RequestCreateServerInstance)
			reqParams.FeeSystemTypeCode = String(5)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateServerInstances(reqParams)

			Expect(err.Error()).To(Equal("FeeSystemTypeCode should be FXSUM or MTRAT"))
		})

		It("should be error : 'expected length of UserData to be in the range (0 - 21847)'", func() {
			reqParams := new(RequestCreateServerInstance)
			reqParams.UserData = String(21848)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateServerInstances(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of UserData to be in the range (0 - 21847)"))
		})
	})
})
