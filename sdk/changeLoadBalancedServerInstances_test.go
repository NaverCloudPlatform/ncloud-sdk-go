package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Change Load Balanced Instance", func() {
	Describe("Change http Load Balanced Instance", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/loadbalancer").
				Reply(http.StatusOK).BodyString(`
					<changeLoadBalancedServerInstancesResponse>
					<requestId>b056782c-6fac-4a37-8e18-8e92fd04e707</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<loadBalancerInstanceList>
						<loadBalancerInstance>
							<loadBalancerInstanceNo>814712</loadBalancerInstanceNo>
							<virtualIp>210.89.187.37,210.89.187.172</virtualIp>
							<loadBalancerName>bbg</loadBalancerName>
							<loadBalancerAlgorithmType>
								<code>RR</code>
								<codeName>Round Robin</codeName>
							</loadBalancerAlgorithmType>
							<loadBalancerDescription></loadBalancerDescription>
							<createDate>2018-06-10T16:58:06+0900</createDate>
							<domainName>slb-814712.ncloudslb.com</domainName>
							<internetLineType>
								<code>PUBLC</code>
								<codeName>PUBLC</codeName>
							</internetLineType>
							<loadBalancerInstanceStatusName>changing</loadBalancerInstanceStatusName>
							<loadBalancerInstanceStatus>
								<code>USED</code>
								<codeName>NET USED state</codeName>
							</loadBalancerInstanceStatus>
							<loadBalancerInstanceOperation>
								<code>CHANG</code>
								<codeName>NET CHANGE OP</codeName>
							</loadBalancerInstanceOperation>
							<networkUsageType>
								<code>PBLIP</code>
								<codeName>Public</codeName>
							</networkUsageType>
							<isHttpKeepAlive>false</isHttpKeepAlive>
							<connectionTimeout>60</connectionTimeout>
							<certificateName></certificateName>
							<loadBalancerRuleList>
								<loadBalancerRule>
									<protocolType>
										<code>HTTP</code>
										<codeName>http</codeName>
									</protocolType>
									<loadBalancerPort>80</loadBalancerPort>
									<serverPort>80</serverPort>
									<l7HealthCheckPath>/</l7HealthCheckPath>
									<certificateName></certificateName>
									<proxyProtocolUseYn>N</proxyProtocolUseYn>
								</loadBalancerRule>
							</loadBalancerRuleList>
							<loadBalancedServerInstanceList>
								<loadBalancedServerInstance>
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
									<serverHealthCheckStatusList>
										<serverHealthCheckStatus>
											<protocolType>
												<code>HTTP</code>
												<codeName>http</codeName>
											</protocolType>
											<loadBalancerPort>80</loadBalancerPort>
											<serverPort>80</serverPort>
											<l7HealthCheckPath>/</l7HealthCheckPath>
											<proxyProtocolUseYn>N</proxyProtocolUseYn>
											<serverStatus>false</serverStatus>
										</serverHealthCheckStatus>
									</serverHealthCheckStatusList>
								</loadBalancedServerInstance>
								<loadBalancedServerInstance>
									<serverInstance>
										<serverInstanceNo>805844</serverInstanceNo>
										<serverName>aaa006</serverName>
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
										<privateIp>10.41.3.33</privateIp>
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
										<createDate>2018-06-04T15:16:06+0900</createDate>
										<uptime>2018-06-04T15:19:23+0900</uptime>
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
									<serverHealthCheckStatusList>
										<serverHealthCheckStatus>
											<protocolType>
												<code>HTTP</code>
												<codeName>http</codeName>
											</protocolType>
											<loadBalancerPort>80</loadBalancerPort>
											<serverPort>80</serverPort>
											<l7HealthCheckPath>/</l7HealthCheckPath>
											<proxyProtocolUseYn>N</proxyProtocolUseYn>
											<serverStatus>false</serverStatus>
										</serverHealthCheckStatus>
									</serverHealthCheckStatusList>
								</loadBalancedServerInstance>
								<loadBalancedServerInstance>
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
									<serverHealthCheckStatusList>
										<serverHealthCheckStatus>
											<protocolType>
												<code>HTTP</code>
												<codeName>http</codeName>
											</protocolType>
											<loadBalancerPort>80</loadBalancerPort>
											<serverPort>80</serverPort>
											<l7HealthCheckPath>/</l7HealthCheckPath>
											<proxyProtocolUseYn>N</proxyProtocolUseYn>
											<serverStatus>false</serverStatus>
										</serverHealthCheckStatus>
									</serverHealthCheckStatusList>
								</loadBalancedServerInstance>
							</loadBalancedServerInstanceList>
						</loadBalancerInstance>
					</loadBalancerInstanceList>
				</changeLoadBalancedServerInstancesResponse>
					`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should Change Load Balanced Server Instances", func() {
			conn := NewConnection(accessKey, secretKey)
			reqParams := new(RequestChangeLoadBalancedServerInstances)
			reqParams.LoadBalancerInstanceNo = "814712"
			reqParams.ServerInstanceNoList = []string{"805841", "805844", "805847"}
			result, err := conn.ChangeLoadBalancedServerInstances(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.LoadBalancerInstanceList)).To(Equal(1))

			loadBalancerInstance := result.LoadBalancerInstanceList[0]
			Expect(len(loadBalancerInstance.LoadBalancedServerInstanceList)).To(Equal(3))

			loadBalancedServerInstance := loadBalancerInstance.LoadBalancedServerInstanceList[0]
			Expect(len(loadBalancedServerInstance.ServerInstanceList)).To(Equal(1))

			serverInstance := loadBalancedServerInstance.ServerInstanceList[0]
			Expect(serverInstance.ServerInstanceNo).To(Equal("805847"))

			loadBalancedServerInstance = loadBalancerInstance.LoadBalancedServerInstanceList[1]
			Expect(len(loadBalancedServerInstance.ServerInstanceList)).To(Equal(1))

			serverInstance = loadBalancedServerInstance.ServerInstanceList[0]
			Expect(serverInstance.ServerInstanceNo).To(Equal("805844"))

			loadBalancedServerInstance = loadBalancerInstance.LoadBalancedServerInstanceList[2]
			Expect(len(loadBalancedServerInstance.ServerInstanceList)).To(Equal(1))

			serverInstance = loadBalancedServerInstance.ServerInstanceList[0]
			Expect(serverInstance.ServerInstanceNo).To(Equal("805841"))
		})
	})

	Describe("Authorize fail", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/loadbalancer").
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
			reqParams := new(RequestChangeLoadBalancedServerInstances)
			reqParams.LoadBalancerInstanceNo = "814712"
			result, err := conn.ChangeLoadBalancedServerInstances(reqParams)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Invalid consumerKey"))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})

	Describe("Expired URL", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/loadbalancer").
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
			reqParams := new(RequestChangeLoadBalancedServerInstances)
			reqParams.LoadBalancerInstanceNo = "814712"
			result, err := conn.ChangeLoadBalancedServerInstances(reqParams)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Expired url."))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})

	Describe("Check Arguments", func() {
		It("should be error : 'LoadBalancerInstanceNo field is required'", func() {
			conn := NewConnection(accessKey, secretKey)
			_, err := conn.ChangeLoadBalancedServerInstances(nil)

			Expect(err.Error()).To(Equal("LoadBalancerInstanceNo field is required"))
		})

		It("should be error : 'LoadBalancerInstanceNo field is required'", func() {
			reqParams := new(RequestChangeLoadBalancedServerInstances)
			conn := NewConnection(accessKey, secretKey)
			_, err := conn.ChangeLoadBalancedServerInstances(reqParams)

			Expect(err.Error()).To(Equal("LoadBalancerInstanceNo field is required"))
		})
	})
})
