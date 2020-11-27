package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Public IP Instance List", func() {
	Describe("Get all Public IP Instance List", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getPublicIpInstanceListResponse>
					<requestId>66fe1605-a8d8-4d96-a19b-357169205c93</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<publicIpInstanceList>
						<publicIpInstance>
							<publicIpInstanceNo>806088</publicIpInstanceNo>
							<publicIp>45.119.146.54</publicIp>
							<publicIpDescription></publicIpDescription>
							<createDate>2018-06-04T19:17:20+0900</createDate>
							<internetLineType>
								<code>PUBLC</code>
								<codeName>PUBLC</codeName>
							</internetLineType>
							<publicIpInstanceStatusName>created</publicIpInstanceStatusName>
							<publicIpInstanceStatus>
								<code>CREAT</code>
								<codeName>NET CREATE state</codeName>
							</publicIpInstanceStatus>
							<publicIpInstanceOperation>
								<code>NULL</code>
								<codeName>NET NULL OP</codeName>
							</publicIpInstanceOperation>
							<publicIpKindType>
								<code>GEN</code>
								<codeName>General</codeName>
							</publicIpKindType>
							<serverInstanceAssociatedWithPublicIp/>
							<zone>
								<zoneNo>3</zoneNo>
								<zoneName>KR-2</zoneName>
								<zoneCode>KR-2</zoneCode>
								<zoneDescription>평촌 zone</zoneDescription>
								<regionNo>1</regionNo>
							</zone>
						</publicIpInstance>
					</publicIpInstanceList>
				</getPublicIpInstanceListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get Public IP Instance List", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetPublicIPInstanceList(nil)

			Expect(err).To(BeNil())
			Expect(result.RequestID).To(Equal("66fe1605-a8d8-4d96-a19b-357169205c93"))
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.PublicIPInstanceList)).To(Equal(1))

			publicIPInstance := result.PublicIPInstanceList[0]
			Expect(publicIPInstance.PublicIPInstanceNo).To(Equal("806088"))
			Expect(publicIPInstance.PublicIP).To(Equal("45.119.146.54"))
			Expect(publicIPInstance.CreateDate).To(Equal("2018-06-04T19:17:20+0900"))
			Expect(publicIPInstance.PublicIPInstanceStatusName).To(Equal("created"))
			Expect(publicIPInstance.PublicIPInstanceStatus.Code).To(Equal("CREAT"))
			Expect(publicIPInstance.PublicIPInstanceStatus.CodeName).To(Equal("NET CREATE state"))
			Expect(publicIPInstance.PublicIPInstanceOperation.Code).To(Equal("NULL"))
			Expect(publicIPInstance.PublicIPInstanceOperation.CodeName).To(Equal("NET NULL OP"))
			Expect(publicIPInstance.PublicIPKindType.Code).To(Equal("GEN"))
			Expect(publicIPInstance.PublicIPKindType.CodeName).To(Equal("General"))
			Expect(publicIPInstance.Zone.ZoneNo).To(Equal("3"))
			Expect(publicIPInstance.Zone.ZoneName).To(Equal("KR-2"))
			Expect(publicIPInstance.Zone.ZoneCode).To(Equal("KR-2"))
			Expect(publicIPInstance.Zone.ZoneDescription).To(Equal("평촌 zone"))
			Expect(publicIPInstance.Zone.RegionNo).To(Equal("1"))
		})
	})

	Describe("Get Public IP Instance List which is parameters", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getPublicIpInstanceListResponse>
					<requestId>3c0668ae-6c6d-4b60-9bb5-874204f81802</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<publicIpInstanceList>
						<publicIpInstance>
							<publicIpInstanceNo>565346</publicIpInstanceNo>
							<publicIp>210.89.181.250</publicIp>
							<publicIpDescription></publicIpDescription>
							<createDate>2017-12-18T15:16:33+0900</createDate>
							<publicIpInstanceStatusName>used</publicIpInstanceStatusName>
							<publicIpInstanceStatus>
								<code>USED</code>
								<codeName>NET USED state</codeName>
							</publicIpInstanceStatus>
							<publicIpInstanceOperation>
								<code>NULL</code>
								<codeName>NET NULL OP</codeName>
							</publicIpInstanceOperation>
							<publicIpKindType>
								<code>GEN</code>
								<codeName>General</codeName>
							</publicIpKindType>
							<serverInstanceAssociatedWithPublicIp>
								<serverInstanceNo>565284</serverInstanceNo>
								<serverName>svr-9bb8d497d29584a</serverName>
								<serverDescription></serverDescription>
								<cpuCount>2</cpuCount>
								<memorySize>4294967296</memorySize>
								<baseBlockStorageSize>53687091200</baseBlockStorageSize>
								<platformType>
									<code>LNX64</code>
									<codeName>Linux 64 Bit</codeName>
								</platformType>
								<loginKeyName>packer-1513575006</loginKeyName>
								<isFeeChargingMonitoring>false</isFeeChargingMonitoring>
								<publicIp>210.89.181.250</publicIp>
								<privateIp>10.33.2.206</privateIp>
								<serverImageName>centos-5.11-64</serverImageName>
								<serverInstanceStatus>
									<code>RUN</code>
									<codeName>Server run state</codeName>
								</serverInstanceStatus>
								<serverInstanceOperation>
									<code>NULL</code>
									<codeName>Server NULL OP</codeName>
								</serverInstanceOperation>
								<serverInstanceStatusName>running</serverInstanceStatusName>
								<createDate>2017-12-18T14:30:08+0900</createDate>
								<uptime>2017-12-18T14:35:09+0900</uptime>
								<serverImageProductCode>SPSW0LINUX000043</serverImageProductCode>
								<serverProductCode>SPSVRSTAND000004</serverProductCode>
								<isProtectServerTermination>false</isProtectServerTermination>
								<portForwardingPublicIp>211.249.60.222</portForwardingPublicIp>
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
								<accessControlGroupList>
									<accessControlGroup>
										<accessControlGroupConfigurationNo>4964</accessControlGroupConfigurationNo>
										<accessControlGroupName>ncloud-default-acg</accessControlGroupName>
										<accessControlGroupDescription>Default AccessControlGroup</accessControlGroupDescription>
										<isDefault>true</isDefault>
										<createDate>2017-02-23T10:25:39+0900</createDate>
									</accessControlGroup>
								</accessControlGroupList>
							</serverInstanceAssociatedWithPublicIp>
						</publicIpInstance>
					</publicIpInstanceList>
				</getPublicIpInstanceListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get One Public IP Instance List", func() {
			reqParams := new(RequestPublicIPInstanceList)
			reqParams.PublicIPInstanceNoList = []string{"565346"}

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetPublicIPInstanceList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.RequestID).To(Equal("3c0668ae-6c6d-4b60-9bb5-874204f81802"))
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.PublicIPInstanceList)).To(Equal(1))

			publicIPInstance := result.PublicIPInstanceList[0]
			Expect(publicIPInstance.PublicIPInstanceNo).To(Equal("565346"))
			Expect(publicIPInstance.PublicIP).To(Equal("210.89.181.250"))
			Expect(publicIPInstance.CreateDate).To(Equal("2017-12-18T15:16:33+0900"))
			Expect(publicIPInstance.PublicIPInstanceStatusName).To(Equal("used"))
			Expect(publicIPInstance.PublicIPInstanceStatus.Code).To(Equal("USED"))
			Expect(publicIPInstance.PublicIPInstanceStatus.CodeName).To(Equal("NET USED state"))
			Expect(publicIPInstance.PublicIPInstanceOperation.Code).To(Equal("NULL"))
			Expect(publicIPInstance.PublicIPInstanceOperation.CodeName).To(Equal("NET NULL OP"))
			Expect(publicIPInstance.PublicIPKindType.Code).To(Equal("GEN"))
			Expect(publicIPInstance.PublicIPKindType.CodeName).To(Equal("General"))
			Expect(publicIPInstance.ServerInstance.ServerInstanceNo).To(Equal("565284"))
		})
	})

	Describe("There is no Public IP Instance list", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getPublicIpInstanceListResponse>
					<requestId>f7791c1d-941f-4968-bc24-a7224e16d39f</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>0</totalRows>
					<publicIpInstanceList/>
				</getPublicIpInstanceListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get No Public IP Instance list", func() {
			reqParams := new(RequestPublicIPInstanceList)
			reqParams.PublicIPInstanceNoList = []string{"5635346"}

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetPublicIPInstanceList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.RequestID).To(Equal("f7791c1d-941f-4968-bc24-a7224e16d39f"))
			Expect(result.TotalRows).To(Equal(0))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.PublicIPInstanceList)).To(Equal(0))
		})
	})
})
