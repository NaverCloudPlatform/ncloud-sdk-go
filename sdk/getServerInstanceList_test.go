package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Server Instance List", func() {
	Describe("Get all Server Instance List", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getServerInstanceListResponse>
					<requestId>f8342548-b510-436e-87c0-c0cd63c14fcd</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>4</totalRows>
					<serverInstanceList>
						<serverInstance>
							<serverInstanceNo>320897</serverInstanceNo>
							<serverName>name-change-test111</serverName>
							<serverDescription>testffff</serverDescription>
							<cpuCount>1</cpuCount>
							<memorySize>2147483648</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<loginKeyName>bk-test</loginKeyName>
							<isFeeChargingMonitoring>false</isFeeChargingMonitoring>
							<publicIp></publicIp>
							<privateIp>10.113.177.17</privateIp>
							<serverImageName>centos-6.3-64</serverImageName>
							<serverInstanceStatus>
								<code>NSTOP</code>
								<codeName>Server normal stopped state</codeName>
							</serverInstanceStatus>
							<serverInstanceOperation>
								<code>NULL</code>
								<codeName>Server NULL OP</codeName>
							</serverInstanceOperation>
							<serverInstanceStatusName>stopped</serverInstanceStatusName>
							<createDate>2017-03-21T18:53:58+0900</createDate>
							<uptime>2017-11-27T14:35:26+0900</uptime>
							<serverImageProductCode>SPSW0LINUX000031</serverImageProductCode>
							<serverProductCode>SPSVRSTAND000003</serverProductCode>
							<isProtectServerTermination>false</isProtectServerTermination>
							<portForwardingPublicIp>192.168.120.140</portForwardingPublicIp>
							<portForwardingExternalPort>1028</portForwardingExternalPort>
							<portForwardingInternalPort>22</portForwardingInternalPort>
							<zone>
								<zoneNo>2</zoneNo>
								<zoneName>KR-1</zoneName>
								<zoneCode>KR-1</zoneCode>
								<zoneDescription>가산 zone</zoneDescription>
								<regionNo>1</regionNo>
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
							<internetLineType>
								<code>PUBLC</code>
								<codeName>PUBLC</codeName>
							</internetLineType>
							<userData></userData>
							<accessControlGroupList>
								<accessControlGroup>
									<accessControlGroupConfigurationNo>3934</accessControlGroupConfigurationNo>
									<accessControlGroupName>acgtest</accessControlGroupName>
									<accessControlGroupDescription>ACGtest</accessControlGroupDescription>
									<isDefault>false</isDefault>
									<createDate>2017-03-21T18:53:23+0900</createDate>
								</accessControlGroup>
							</accessControlGroupList>
						</serverInstance>
						<serverInstance>
							<serverInstanceNo>343571</serverInstanceNo>
							<serverName>compt-jinyong</serverName>
							<serverDescription></serverDescription>
							<cpuCount>1</cpuCount>
							<memorySize>2147483648</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<loginKeyName>penguin</loginKeyName>
							<isFeeChargingMonitoring>false</isFeeChargingMonitoring>
							<publicIp></publicIp>
							<privateIp>10.113.178.233</privateIp>
							<serverImageName>centos-7.2-64</serverImageName>
							<serverInstanceStatus>
								<code>CREAT</code>
								<codeName>Server Created State</codeName>
							</serverInstanceStatus>
							<serverInstanceOperation>
								<code>START</code>
								<codeName>Server START OP</codeName>
							</serverInstanceOperation>
							<serverInstanceStatusName>booting</serverInstanceStatusName>
							<createDate>2017-08-20T19:30:13+0900</createDate>
							<uptime>2017-08-20T19:30:13+0900</uptime>
							<serverImageProductCode>SPSW0LINUX000045</serverImageProductCode>
							<serverProductCode>SPSVRSSD00000001</serverProductCode>
							<isProtectServerTermination>false</isProtectServerTermination>
							<portForwardingPublicIp>192.168.120.140</portForwardingPublicIp>
							<zone>
								<zoneNo>3</zoneNo>
								<zoneName>KR-2</zoneName>
								<zoneCode>KR-2</zoneCode>
								<zoneDescription>평촌 zone</zoneDescription>
								<regionNo>1</regionNo>
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
							<userData></userData>
							<accessControlGroupList/>
						</serverInstance>
						<serverInstance>
							<serverInstanceNo>343577</serverInstanceNo>
							<serverName>s-7s1ojtpou9gpha</serverName>
							<serverDescription>Auto scaling group &apos;iamtest&apos; 소속의 자동 생성된 서버</serverDescription>
							<cpuCount>2</cpuCount>
							<memorySize>4294967296</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<loginKeyName>penguin</loginKeyName>
							<isFeeChargingMonitoring>false</isFeeChargingMonitoring>
							<publicIp></publicIp>
							<privateIp>10.113.177.10</privateIp>
							<serverImageName>centos-6.6-64</serverImageName>
							<serverInstanceStatus>
								<code>NSTOP</code>
								<codeName>Server normal stopped state</codeName>
							</serverInstanceStatus>
							<serverInstanceOperation>
								<code>START</code>
								<codeName>Server START OP</codeName>
							</serverInstanceOperation>
							<serverInstanceStatusName>booting</serverInstanceStatusName>
							<createDate>2017-08-20T19:36:30+0900</createDate>
							<uptime>2017-08-20T19:38:56+0900</uptime>
							<serverImageProductCode>SPSW0LINUX000044</serverImageProductCode>
							<serverProductCode>SPSVRSSD00000003</serverProductCode>
							<isProtectServerTermination>false</isProtectServerTermination>
							<portForwardingPublicIp>192.168.120.140</portForwardingPublicIp>
							<zone>
								<zoneNo>2</zoneNo>
								<zoneName>KR-1</zoneName>
								<zoneCode>KR-1</zoneCode>
								<zoneDescription>가산 zone</zoneDescription>
								<regionNo>1</regionNo>
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
							<userData></userData>
							<accessControlGroupList>
								<accessControlGroup>
									<accessControlGroupConfigurationNo>4715</accessControlGroupConfigurationNo>
									<accessControlGroupName>iamtest</accessControlGroupName>
									<accessControlGroupDescription></accessControlGroupDescription>
									<isDefault>false</isDefault>
									<createDate>2017-08-14T15:37:57+0900</createDate>
								</accessControlGroup>
							</accessControlGroupList>
						</serverInstance>
						<serverInstance>
							<serverInstanceNo>347848</serverInstanceNo>
							<serverName>name-change-test2</serverName>
							<serverDescription></serverDescription>
							<cpuCount>2</cpuCount>
							<memorySize>4294967296</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<loginKeyName>test004</loginKeyName>
							<isFeeChargingMonitoring>false</isFeeChargingMonitoring>
							<publicIp></publicIp>
							<privateIp>10.113.177.13</privateIp>
							<serverImageName>centos-6.6-64</serverImageName>
							<serverInstanceStatus>
								<code>RUN</code>
								<codeName>Server RUN State</codeName>
							</serverInstanceStatus>
							<serverInstanceOperation>
								<code>SHTDN</code>
								<codeName>Server SHUTDOWN OP</codeName>
							</serverInstanceOperation>
							<serverInstanceStatusName>shutting down</serverInstanceStatusName>
							<createDate>2017-08-25T18:52:14+0900</createDate>
							<uptime>2017-08-27T14:54:16+0900</uptime>
							<serverImageProductCode>SPSW0LINUX000044</serverImageProductCode>
							<serverProductCode>SPSVRSSD00000003</serverProductCode>
							<isProtectServerTermination>false</isProtectServerTermination>
							<portForwardingPublicIp>192.168.120.140</portForwardingPublicIp>
							<zone>
								<zoneNo>2</zoneNo>
								<zoneName>KR-1</zoneName>
								<zoneCode>KR-1</zoneCode>
								<zoneDescription>가산 zone</zoneDescription>
								<regionNo>1</regionNo>
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
							<userData></userData>
							<accessControlGroupList>
								<accessControlGroup>
									<accessControlGroupConfigurationNo>3992</accessControlGroupConfigurationNo>
									<accessControlGroupName>bktest1</accessControlGroupName>
									<accessControlGroupDescription>test</accessControlGroupDescription>
									<isDefault>false</isDefault>
									<createDate>2017-03-28T11:04:39+0900</createDate>
								</accessControlGroup>
							</accessControlGroupList>
						</serverInstance>
					</serverInstanceList>
				</getServerInstanceListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get Server Instance list", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetServerInstanceList(nil)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(4))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.ServerInstanceList)).To(Equal(4))

			serverinstance := result.ServerInstanceList[0]
			Expect(serverinstance.ServerInstanceNo).To(Equal("320897"))
			Expect(serverinstance.ServerName).To(Equal("name-change-test111"))
			Expect(serverinstance.PlatformType.Code).To(Equal("LNX64"))
			Expect(serverinstance.PlatformType.CodeName).To(Equal("Linux 64 Bit"))
			Expect(serverinstance.LoginKeyName).To(Equal("bk-test"))
			Expect(serverinstance.Region.RegionCode).To(Equal("KR"))
			Expect(serverinstance.Region.RegionName).To(Equal("Korea"))
			Expect(serverinstance.Zone.ZoneNo).To(Equal("2"))
			Expect(serverinstance.Zone.ZoneName).To(Equal("KR-1"))
			Expect(serverinstance.Zone.ZoneCode).To(Equal("KR-1"))
			Expect(serverinstance.Zone.ZoneDescription).To(Equal("가산 zone"))
			Expect(serverinstance.Zone.RegionNo).To(Equal("1"))
			Expect(serverinstance.PlatformType.Code).To(Equal("LNX64"))
			Expect(serverinstance.ServerInstanceStatusName).To(Equal("stopped"))
			Expect(serverinstance.PortForwardingPublicIP).To(Equal("192.168.120.140"))
			Expect(serverinstance.PortForwardingExternalPort).To(Equal(1028))
			Expect(serverinstance.PortForwardingInternalPort).To(Equal(22))
			Expect(serverinstance.Uptime).To(Equal("2017-11-27T14:35:26+0900"))
			Expect(serverinstance.BaseBlockStorageDiskType.Code).To(Equal("NET"))

			serverinstance = result.ServerInstanceList[1]
			Expect(serverinstance.ServerInstanceNo).To(Equal("343571"))
			Expect(serverinstance.ServerName).To(Equal("compt-jinyong"))
			Expect(serverinstance.PlatformType.Code).To(Equal("LNX64"))
			Expect(serverinstance.PlatformType.CodeName).To(Equal("Linux 64 Bit"))
			Expect(serverinstance.LoginKeyName).To(Equal("penguin"))
			Expect(serverinstance.Region.RegionCode).To(Equal("KR"))
			Expect(serverinstance.Region.RegionName).To(Equal("Korea"))
			Expect(serverinstance.Zone.ZoneNo).To(Equal("3"))
			Expect(serverinstance.Zone.ZoneName).To(Equal("KR-2"))
			Expect(serverinstance.Zone.ZoneCode).To(Equal("KR-2"))
			Expect(serverinstance.Zone.ZoneDescription).To(Equal("평촌 zone"))
			Expect(serverinstance.Zone.RegionNo).To(Equal("1"))
			Expect(serverinstance.PlatformType.Code).To(Equal("LNX64"))
			Expect(serverinstance.ServerInstanceStatusName).To(Equal("booting"))
			Expect(serverinstance.PortForwardingPublicIP).To(Equal("192.168.120.140"))
			Expect(serverinstance.Uptime).To(Equal("2017-08-20T19:30:13+0900"))
			Expect(serverinstance.BaseBlockStorageDiskType.Code).To(Equal("NET"))
		})
	})

	Describe("Get One Server Instance List which ServerInstanceNoList.1 is 320897", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getServerInstanceListResponse>
					<requestId>8df5a02f-9d58-489d-9653-65dae4a2a84c</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<serverInstanceList>
						<serverInstance>
							<serverInstanceNo>320897</serverInstanceNo>
							<serverName>name-change-test111</serverName>
							<serverDescription>testffff</serverDescription>
							<cpuCount>1</cpuCount>
							<memorySize>2147483648</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<loginKeyName>bk-test</loginKeyName>
							<isFeeChargingMonitoring>false</isFeeChargingMonitoring>
							<publicIp></publicIp>
							<privateIp>10.113.177.17</privateIp>
							<serverImageName>centos-6.3-64</serverImageName>
							<serverInstanceStatus>
								<code>NSTOP</code>
								<codeName>Server normal stopped state</codeName>
							</serverInstanceStatus>
							<serverInstanceOperation>
								<code>NULL</code>
								<codeName>Server NULL OP</codeName>
							</serverInstanceOperation>
							<serverInstanceStatusName>stopped</serverInstanceStatusName>
							<createDate>2017-03-21T18:53:58+0900</createDate>
							<uptime>2017-11-27T14:35:26+0900</uptime>
							<serverImageProductCode>SPSW0LINUX000031</serverImageProductCode>
							<serverProductCode>SPSVRSTAND000003</serverProductCode>
							<isProtectServerTermination>false</isProtectServerTermination>
							<portForwardingPublicIp>192.168.120.140</portForwardingPublicIp>
							<portForwardingExternalPort>1028</portForwardingExternalPort>
							<portForwardingInternalPort>22</portForwardingInternalPort>
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
								<code>HDD</code>
								<codeName>HDD</codeName>
							</baseBlockStroageDiskDetailType>
							<internetLineType>
								<code>PUBLC</code>
								<codeName>PUBLC</codeName>
							</internetLineType>
							<userData></userData>
							<accessControlGroupList>
								<accessControlGroup>
									<accessControlGroupConfigurationNo>3934</accessControlGroupConfigurationNo>
									<accessControlGroupName>acgtest</accessControlGroupName>
									<accessControlGroupDescription>ACGtest</accessControlGroupDescription>
									<isDefault>false</isDefault>
									<createDate>2017-03-21T18:53:23+0900</createDate>
								</accessControlGroup>
							</accessControlGroupList>
						</serverInstance>
					</serverInstanceList>
				</getServerInstanceListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get One Server Instance", func() {
			reqParams := new(RequestGetServerInstanceList)
			reqParams.ServerInstanceNoList = []string{"320897"}

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetServerInstanceList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(len(result.ServerInstanceList)).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))

			serverinstance := result.ServerInstanceList[0]
			Expect(serverinstance.ServerInstanceNo).To(Equal("320897"))
			Expect(serverinstance.ServerName).To(Equal("name-change-test111"))
			Expect(serverinstance.PlatformType.Code).To(Equal("LNX64"))
			Expect(serverinstance.PlatformType.CodeName).To(Equal("Linux 64 Bit"))
			Expect(serverinstance.LoginKeyName).To(Equal("bk-test"))
			Expect(serverinstance.Region.RegionCode).To(Equal("KR"))
			Expect(serverinstance.Region.RegionName).To(Equal("Korea"))
			Expect(serverinstance.Zone.ZoneName).To(Equal("KR-1"))
			Expect(serverinstance.PlatformType.Code).To(Equal("LNX64"))
			Expect(serverinstance.ServerInstanceStatusName).To(Equal("stopped"))
			Expect(serverinstance.PortForwardingPublicIP).To(Equal("192.168.120.140"))
			Expect(serverinstance.PortForwardingExternalPort).To(Equal(1028))
			Expect(serverinstance.PortForwardingInternalPort).To(Equal(22))
			Expect(serverinstance.Uptime).To(Equal("2017-11-27T14:35:26+0900"))
			Expect(serverinstance.BaseBlockStorageDiskType.Code).To(Equal("NET"))

		})
	})

	Describe("There is no Server Instance list", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getServerInstanceListResponse>
					<requestId>7683adc5-e057-4b7c-aeec-ad20171893fd</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>0</totalRows>
					<serverInstanceList/>
					</getServerInstanceListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get No Server Instance list", func() {
			reqParams := new(RequestGetServerInstanceList)
			reqParams.ServerInstanceNoList = []string{"3208397"}

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetServerInstanceList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(0))
			Expect(len(result.ServerInstanceList)).To(Equal(0))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
		})
	})

	Describe("Authorize fail", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusUnauthorized).BodyString(`<responseError>
					<returnCode>800</returnCode>
					<returnMessage>Expired url.</returnMessage>
					</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should be failed by authorization fail", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetServerInstanceList(nil)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Expired url."))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})
})
