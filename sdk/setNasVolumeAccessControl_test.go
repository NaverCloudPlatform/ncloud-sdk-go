package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Set Nas Volume Access Control", func() {
	Describe("Set Nas Volume Access Control", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusOK).BodyString(`
					<setNasVolumeAccessControlResponse>
					<requestId>8979f7b6-d1eb-486f-960f-e2aa374eb860</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<nasVolumeInstanceList>
						<nasVolumeInstance>
							<nasVolumeInstanceNo>823525</nasVolumeInstanceNo>
							<nasVolumeInstanceStatus>
								<code>CREAT</code>
								<codeName>NAS create</codeName>
							</nasVolumeInstanceStatus>
							<nasVolumeInstanceOperation>
								<code>NULL</code>
								<codeName>NAS NULL OP</codeName>
							</nasVolumeInstanceOperation>
							<nasVolumeInstanceStatusName>created</nasVolumeInstanceStatusName>
							<createDate>2018-06-15T16:10:17+0900</createDate>
							<nasVolumeInstanceDescription></nasVolumeInstanceDescription>
							<mountInformation>10.250.53.74:/n003666_aaaa</mountInformation>
							<volumeAllotmentProtocolType>
								<code>NFS</code>
								<codeName>NFS</codeName>
							</volumeAllotmentProtocolType>
							<volumeName>n003666_aaaa</volumeName>
							<volumeTotalSize>536870912000</volumeTotalSize>
							<volumeSize>536870912000</volumeSize>
							<volumeUseSize>303104</volumeUseSize>
							<volumeUseRatio>0.0</volumeUseRatio>
							<snapshotVolumeConfigurationRatio>0.0</snapshotVolumeConfigurationRatio>
							<snapshotVolumeSize>0</snapshotVolumeSize>
							<snapshotVolumeUseSize>0</snapshotVolumeUseSize>
							<snapshotVolumeUseRatio>0.0</snapshotVolumeUseRatio>
							<isSnapshotConfiguration>false</isSnapshotConfiguration>
							<isEventConfiguration>false</isEventConfiguration>
							<region>
								<regionNo>1</regionNo>
								<regionCode>KR</regionCode>
								<regionName>Korea</regionName>
							</region>
							<zone>
								<zoneNo>3</zoneNo>
								<zoneName>KR-2</zoneName>
								<zoneCode>KR-2</zoneCode>
								<zoneDescription>평촌 zone</zoneDescription>
								<regionNo>1</regionNo>
							</zone>
							<nasVolumeInstanceCustomIpList>
								<nasVolumeInstanceCustomIp>
									<customIp>10.1.1.1</customIp>
								</nasVolumeInstanceCustomIp>
								<nasVolumeInstanceCustomIp>
									<customIp>10.1.2.1</customIp>
								</nasVolumeInstanceCustomIp>
							</nasVolumeInstanceCustomIpList>
							<nasVolumeServerInstanceList>
								<serverInstance>
									<serverInstanceNo>805829</serverInstanceNo>
									<serverName>aaa001</serverName>
									<serverDescription></serverDescription>
									<cpuCount>2</cpuCount>
									<memorySize>4294967296</memorySize>
									<baseBlockStorageSize>53687091200</baseBlockStorageSize>
									<platformType>
										<code>LNX64</code>
										<codeName>Linux 64 Bit</codeName>
									</platformType>
									<loginKeyName>consolekey</loginKeyName>
									<isFeeChargingMonitoring>false</isFeeChargingMonitoring>
									<publicIp></publicIp>
									<privateIp>10.41.2.72</privateIp>
									<serverImageName>centos-7.3-64</serverImageName>
									<serverInstanceStatus>
										<code>RUN</code>
										<codeName>Server run state</codeName>
									</serverInstanceStatus>
									<serverInstanceOperation>
										<code>NULL</code>
										<codeName>Server NULL OP</codeName>
									</serverInstanceOperation>
									<serverInstanceStatusName>running</serverInstanceStatusName>
									<createDate>2018-06-04T15:16:02+0900</createDate>
									<uptime>2018-06-04T15:18:53+0900</uptime>
									<serverImageProductCode>SPSW0LINUX000046</serverImageProductCode>
									<serverProductCode>SPSVRSSD00000003</serverProductCode>
									<isProtectServerTermination>false</isProtectServerTermination>
									<portForwardingPublicIp>106.10.41.173</portForwardingPublicIp>
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
									<baseBlockStorageDiskDetailType>
										<code>SSD</code>
										<codeName>SSD</codeName>
									</baseBlockStorageDiskDetailType>
									<internetLineType>
										<code>PUBLC</code>
										<codeName>PUBLC</codeName>
									</internetLineType>
									<serverInstanceType>
										<code>STAND</code>
										<codeName>Standard</codeName>
									</serverInstanceType>
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
								<serverInstance>
									<serverInstanceNo>805841</serverInstanceNo>
									<serverName>aaa005</serverName>
									<serverDescription></serverDescription>
									<cpuCount>2</cpuCount>
									<memorySize>4294967296</memorySize>
									<baseBlockStorageSize>53687091200</baseBlockStorageSize>
									<platformType>
										<code>LNX64</code>
										<codeName>Linux 64 Bit</codeName>
									</platformType>
									<loginKeyName>consolekey</loginKeyName>
									<isFeeChargingMonitoring>false</isFeeChargingMonitoring>
									<publicIp></publicIp>
									<privateIp>10.41.4.88</privateIp>
									<serverImageName>centos-7.3-64</serverImageName>
									<serverInstanceStatus>
										<code>RUN</code>
										<codeName>Server run state</codeName>
									</serverInstanceStatus>
									<serverInstanceOperation>
										<code>NULL</code>
										<codeName>Server NULL OP</codeName>
									</serverInstanceOperation>
									<serverInstanceStatusName>running</serverInstanceStatusName>
									<createDate>2018-06-04T15:16:05+0900</createDate>
									<uptime>2018-06-04T15:19:25+0900</uptime>
									<serverImageProductCode>SPSW0LINUX000046</serverImageProductCode>
									<serverProductCode>SPSVRSSD00000003</serverProductCode>
									<isProtectServerTermination>false</isProtectServerTermination>
									<portForwardingPublicIp>106.10.41.173</portForwardingPublicIp>
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
									<baseBlockStorageDiskDetailType>
										<code>SSD</code>
										<codeName>SSD</codeName>
									</baseBlockStorageDiskDetailType>
									<internetLineType>
										<code>PUBLC</code>
										<codeName>PUBLC</codeName>
									</internetLineType>
									<serverInstanceType>
										<code>STAND</code>
										<codeName>Standard</codeName>
									</serverInstanceType>
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
							</nasVolumeServerInstanceList>
						</nasVolumeInstance>
					</nasVolumeInstanceList>
				</setNasVolumeAccessControlResponse>
					`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should set Nas Volume Access Control", func() {
			conn := NewConnection(accessKey, secretKey)
			reqParams := &RequestNasVolumeAccessControl{
				NasVolumeInstanceNo:  "823525",
				ServerInstanceNoList: []string{"805829", "805841"},
				CustomIPList:         []string{"10.1.1.1", "10.1.2.1"},
			}
			result, err := conn.SetNasVolumeAccessControl(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.NasVolumeInstanceList)).To(Equal(1))

			nvi := result.NasVolumeInstanceList[0]
			Expect(nvi.NasVolumeInstanceNo).To(Equal(reqParams.NasVolumeInstanceNo))
			Expect(len(nvi.NasVolumeServerInstanceList)).To(Equal(len(reqParams.ServerInstanceNoList)))
			Expect(len(nvi.NasVolumeInstanceCustomIPList)).To(Equal(len(reqParams.CustomIPList)))
		})
	})

	Describe("Invalid nasVolumeInstanceNo", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusBadRequest).BodyString(`
					<responseError>
					<returnCode>24109</returnCode>
					<returnMessage>The input parameter instance number is invalid.</returnMessage>
				</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should not set Nas Volume Access Control by invalid nasVolumeInstanceNo", func() {
			conn := NewConnection(accessKey, secretKey)
			reqParams := &RequestNasVolumeAccessControl{
				NasVolumeInstanceNo:  "8235251",
				ServerInstanceNoList: []string{"805829", "805841"},
				CustomIPList:         []string{"10.1.1.1", "10.1.2.1"},
			}
			result, err := conn.SetNasVolumeAccessControl(reqParams)

			Expect(result.ReturnCode).To(Equal(24109))
			Expect(result.ReturnMessage).To(Equal("The input parameter instance number is invalid."))
			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
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
		It("should be failed by authorization fail", func() {
			conn := NewConnection(accessKey, secretKey)
			reqParams := &RequestNasVolumeAccessControl{
				NasVolumeInstanceNo:  "8235251",
				ServerInstanceNoList: []string{"805829", "805841"},
				CustomIPList:         []string{"10.1.1.1", "10.1.2.1"},
			}
			result, err := conn.SetNasVolumeAccessControl(reqParams)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Invalid consumerKey"))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})

	Describe("Expired URL", func() {
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
			conn := NewConnection(accessKey, secretKey)
			reqParams := &RequestNasVolumeAccessControl{
				NasVolumeInstanceNo:  "8235251",
				ServerInstanceNoList: []string{"805829", "805841"},
				CustomIPList:         []string{"10.1.1.1", "10.1.2.1"},
			}
			result, err := conn.SetNasVolumeAccessControl(reqParams)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Expired url."))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})
})
