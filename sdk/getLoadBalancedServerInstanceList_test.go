package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Load Balanced Server Instance List", func() {
	Describe("Get all Load Balanced Server Instance List", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/loadbalancer").
				Reply(http.StatusOK).BodyString(`
					<getLoadBalancedServerInstanceListResponse>
					<requestId>dd803a78-615e-4a57-8b09-f573284b0632</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<serverInstanceList>
						<serverInstance>
							<serverInstanceNo>805847</serverInstanceNo>
							<serverName>aaa007</serverName>
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
							<privateIp>10.41.0.154</privateIp>
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
							<createDate>2018-06-04T15:16:07+0900</createDate>
							<uptime>2018-06-04T15:19:54+0900</uptime>
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
					</serverInstanceList>
				</getLoadBalancedServerInstanceListResponse>
					`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get Load Balanced Server Instance list", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetLoadBalancedServerInstanceList("811670")

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.ServerInstanceList)).To(Equal(1))
		})
	})

	Describe("There is no Load Balanced Server Instance list", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/loadbalancer").
				Reply(http.StatusOK).BodyString(`<GetLoadBalancedServerInstanceListResponse>
					<requestId>7cf7aa28-89da-4dee-8122-61ca61f92e17</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>0</totalRows>
					<serverInstanceList/>
					</GetLoadBalancedServerInstanceListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get No Load Balanced Server Instance list", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetLoadBalancedServerInstanceList("814712")

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
				Get("/loadbalancer").
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
			result, err := conn.GetLoadBalancedServerInstanceList("814712")

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Expired url."))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})

	Describe("Check Arguments", func() {
		It("should be error : 'LoadBalancerInstanceNo is required'", func() {
			conn := NewConnection(accessKey, secretKey)
			_, err := conn.GetLoadBalancedServerInstanceList("")

			Expect(err.Error()).To(Equal("LoadBalancerInstanceNo is required field"))
		})
	})
})
