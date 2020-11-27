package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create Load Balancer Instance", func() {
	Describe("Create http Load Balancer Instance", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/loadbalancer").
				Reply(http.StatusOK).BodyString(`
					<createLoadBalancerInstanceResponse>
					<requestId>f291ceb6-b06c-42db-85bf-a6286e2a7b50</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<loadBalancerInstanceList>
						<loadBalancerInstance>
							<loadBalancerInstanceNo>810982</loadBalancerInstanceNo>
							<virtualIp>210.89.187.214,210.89.187.208</virtualIp>
							<loadBalancerName>LBTest</loadBalancerName>
							<loadBalancerAlgorithmType>
								<code>RR</code>
								<codeName>Round Robin</codeName>
							</loadBalancerAlgorithmType>
							<loadBalancerDescription>LBTest Description</loadBalancerDescription>
							<createDate>2018-06-07T20:56:27+0900</createDate>
							<domainName>slb-810982.ncloudslb.com</domainName>
							<internetLineType>
								<code>PUBLC</code>
								<codeName>PUBLC</codeName>
							</internetLineType>
							<loadBalancerInstanceStatusName>creating</loadBalancerInstanceStatusName>
							<loadBalancerInstanceStatus>
								<code>INIT</code>
								<codeName>NET INIT state</codeName>
							</loadBalancerInstanceStatus>
							<loadBalancerInstanceOperation>
								<code>USE</code>
								<codeName>NET USE OP</codeName>
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
									<l7HealthCheckPath>/monitor</l7HealthCheckPath>
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
											<l7HealthCheckPath>/monitor</l7HealthCheckPath>
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
											<l7HealthCheckPath>/monitor</l7HealthCheckPath>
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
											<l7HealthCheckPath>/monitor</l7HealthCheckPath>
											<proxyProtocolUseYn>N</proxyProtocolUseYn>
											<serverStatus>false</serverStatus>
										</serverHealthCheckStatus>
									</serverHealthCheckStatusList>
								</loadBalancedServerInstance>
							</loadBalancedServerInstanceList>
						</loadBalancerInstance>
					</loadBalancerInstanceList>
				</createLoadBalancerInstanceResponse>
					`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should create Load Balancer Instance (http)", func() {
			conn := NewConnection(accessKey, secretKey)
			reqParams := new(RequestCreateLoadBalancerInstance)
			reqParams.LoadBalancerName = "LBTest"
			reqParams.LoadBalancerDescription = "LBTest Description"
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode:  "HTTP",
					LoadBalancerPort:  80,
					ServerPort:        80,
					L7HealthCheckPath: "/monitor",
				},
			}
			reqParams.ServerInstanceNoList = []string{"805841", "805844", "805847"}
			result, err := conn.CreateLoadBalancerInstance(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.LoadBalancerInstanceList)).To(Equal(1))

			loadBalancerInstance := result.LoadBalancerInstanceList[0]
			Expect(loadBalancerInstance.LoadBalancerInstanceNo).To(Equal("810982"))
			Expect(loadBalancerInstance.VirtualIP).To(Equal("210.89.187.214,210.89.187.208"))
			Expect(loadBalancerInstance.LoadBalancerName).To(Equal(reqParams.LoadBalancerName))
			Expect(loadBalancerInstance.LoadBalancerAlgorithmType.Code).To(Equal("RR"))
			Expect(loadBalancerInstance.LoadBalancerAlgorithmType.CodeName).To(Equal("Round Robin"))
			Expect(loadBalancerInstance.LoadBalancerDescription).To(Equal(reqParams.LoadBalancerDescription))
			Expect(loadBalancerInstance.CreateDate).To(Equal("2018-06-07T20:56:27+0900"))
			Expect(loadBalancerInstance.DomainName).To(Equal("slb-810982.ncloudslb.com"))
			Expect(loadBalancerInstance.InternetLineType.Code).To(Equal("PUBLC"))
			Expect(loadBalancerInstance.InternetLineType.CodeName).To(Equal("PUBLC"))
			Expect(loadBalancerInstance.LoadBalancerInstanceStatusName).To(Equal("creating"))
			Expect(loadBalancerInstance.LoadBalancerInstanceStatus.Code).To(Equal("INIT"))
			Expect(loadBalancerInstance.LoadBalancerInstanceStatus.CodeName).To(Equal("NET INIT state"))
			Expect(loadBalancerInstance.LoadBalancerInstanceOperation.Code).To(Equal("USE"))
			Expect(loadBalancerInstance.LoadBalancerInstanceOperation.CodeName).To(Equal("NET USE OP"))
			Expect(loadBalancerInstance.NetworkUsageType.Code).To(Equal("PBLIP"))
			Expect(loadBalancerInstance.NetworkUsageType.CodeName).To(Equal("Public"))
			Expect(loadBalancerInstance.IsHTTPKeepAlive).To(Equal(false))
			Expect(loadBalancerInstance.ConnectionTimeout).To(Equal(60))
			Expect(loadBalancerInstance.CertificateName).To(Equal(""))

			Expect(len(loadBalancerInstance.LoadBalancerRuleList)).To(Equal(1))
			Expect(len(loadBalancerInstance.LoadBalancedServerInstanceList)).To(Equal(3))

			loadBalancerRule := loadBalancerInstance.LoadBalancerRuleList[0]
			Expect(loadBalancerRule.ProtocolType.Code).To(Equal("HTTP"))
			Expect(loadBalancerRule.ProtocolType.CodeName).To(Equal("http"))
			Expect(loadBalancerRule.LoadBalancerPort).To(Equal(80))
			Expect(loadBalancerRule.ServerPort).To(Equal(80))
			Expect(loadBalancerRule.L7HealthCheckPath).To(Equal("/monitor"))
			Expect(loadBalancerRule.CertificateName).To(Equal(""))
			Expect(loadBalancerRule.ProxyProtocolUseYn).To(Equal("N"))

			loadBalancedServerInstance := loadBalancerInstance.LoadBalancedServerInstanceList[0]
			Expect(len(loadBalancedServerInstance.ServerInstanceList)).To(Equal(1))
			Expect(len(loadBalancedServerInstance.ServerHealthCheckStatusList)).To(Equal(1))

			serverInstance := loadBalancedServerInstance.ServerInstanceList[0]
			serverHealthCheckStatus := loadBalancedServerInstance.ServerHealthCheckStatusList[0]

			Expect(serverInstance.ServerInstanceNo).To(Equal("805847"))
			Expect(serverInstance.ServerName).To(Equal("aaa007"))

			Expect(serverHealthCheckStatus.ProtocolType.Code).To(Equal("HTTP"))
			Expect(serverHealthCheckStatus.ProtocolType.CodeName).To(Equal("http"))
			Expect(serverHealthCheckStatus.LoadBalancerPort).To(Equal(80))
			Expect(serverHealthCheckStatus.ServerPort).To(Equal(80))
			Expect(serverHealthCheckStatus.L7HealthCheckPath).To(Equal("/monitor"))
			Expect(serverHealthCheckStatus.ProxyProtocolUseYn).To(Equal("N"))
			Expect(serverHealthCheckStatus.ServerStatus).To(Equal(false))
		})
	})

	Describe("Create https Load Balancer Instance", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/loadbalancer").
				Reply(http.StatusOK).BodyString(`
					<createLoadBalancerInstanceResponse>
					<requestId>7caffad3-689a-4d79-b1a1-525132b2fdb7</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<loadBalancerInstanceList>
						<loadBalancerInstance>
							<loadBalancerInstanceNo>810984</loadBalancerInstanceNo>
							<virtualIp>49.236.149.111,49.236.149.189</virtualIp>
							<loadBalancerName>LBTest2</loadBalancerName>
							<loadBalancerAlgorithmType>
								<code>RR</code>
								<codeName>Round Robin</codeName>
							</loadBalancerAlgorithmType>
							<loadBalancerDescription>LBTest2 Description</loadBalancerDescription>
							<createDate>2018-06-07T21:05:52+0900</createDate>
							<domainName>slb-810984.ncloudslb.com</domainName>
							<internetLineType>
								<code>PUBLC</code>
								<codeName>PUBLC</codeName>
							</internetLineType>
							<loadBalancerInstanceStatusName>creating</loadBalancerInstanceStatusName>
							<loadBalancerInstanceStatus>
								<code>INIT</code>
								<codeName>NET INIT state</codeName>
							</loadBalancerInstanceStatus>
							<loadBalancerInstanceOperation>
								<code>USE</code>
								<codeName>NET USE OP</codeName>
							</loadBalancerInstanceOperation>
							<networkUsageType>
								<code>PBLIP</code>
								<codeName>Public</codeName>
							</networkUsageType>
							<isHttpKeepAlive>false</isHttpKeepAlive>
							<connectionTimeout>60</connectionTimeout>
							<certificateName>aaa</certificateName>
							<loadBalancerRuleList>
								<loadBalancerRule>
									<protocolType>
										<code>HTTPS</code>
										<codeName>https</codeName>
									</protocolType>
									<loadBalancerPort>443</loadBalancerPort>
									<serverPort>443</serverPort>
									<l7HealthCheckPath>/monitor</l7HealthCheckPath>
									<certificateName>aaa</certificateName>
									<proxyProtocolUseYn>N</proxyProtocolUseYn>
								</loadBalancerRule>
							</loadBalancerRuleList>
							<loadBalancedServerInstanceList>
								<loadBalancedServerInstance>
									<serverInstance>
										<serverInstanceNo>805838</serverInstanceNo>
										<serverName>aaa004</serverName>
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
										<privateIp>10.41.0.171</privateIp>
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
										<createDate>2018-06-04T15:16:04+0900</createDate>
										<uptime>2018-06-04T15:19:24+0900</uptime>
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
												<code>HTTPS</code>
												<codeName>https</codeName>
											</protocolType>
											<loadBalancerPort>443</loadBalancerPort>
											<serverPort>443</serverPort>
											<l7HealthCheckPath>/monitor</l7HealthCheckPath>
											<proxyProtocolUseYn>N</proxyProtocolUseYn>
											<serverStatus>false</serverStatus>
										</serverHealthCheckStatus>
									</serverHealthCheckStatusList>
								</loadBalancedServerInstance>
								<loadBalancedServerInstance>
									<serverInstance>
										<serverInstanceNo>805835</serverInstanceNo>
										<serverName>aaa003</serverName>
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
										<privateIp>10.41.5.185</privateIp>
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
										<createDate>2018-06-04T15:16:03+0900</createDate>
										<uptime>2018-06-04T15:18:56+0900</uptime>
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
												<code>HTTPS</code>
												<codeName>https</codeName>
											</protocolType>
											<loadBalancerPort>443</loadBalancerPort>
											<serverPort>443</serverPort>
											<l7HealthCheckPath>/monitor</l7HealthCheckPath>
											<proxyProtocolUseYn>N</proxyProtocolUseYn>
											<serverStatus>false</serverStatus>
										</serverHealthCheckStatus>
									</serverHealthCheckStatusList>
								</loadBalancedServerInstance>
							</loadBalancedServerInstanceList>
						</loadBalancerInstance>
					</loadBalancerInstanceList>
				</createLoadBalancerInstanceResponse>
					`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should create Load Balancer Instance (https)", func() {
			conn := NewConnection(accessKey, secretKey)
			reqParams := new(RequestCreateLoadBalancerInstance)
			reqParams.LoadBalancerName = "LBTest2"
			reqParams.LoadBalancerDescription = "LBTest2 Description"
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode:  "HTTPS",
					LoadBalancerPort:  443,
					ServerPort:        443,
					L7HealthCheckPath: "/monitor",
					CertificateName:   "aaa",
				},
			}
			reqParams.ServerInstanceNoList = []string{"805835", "805838"}
			result, err := conn.CreateLoadBalancerInstance(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.LoadBalancerInstanceList)).To(Equal(1))

			loadBalancerInstance := result.LoadBalancerInstanceList[0]
			Expect(loadBalancerInstance.LoadBalancerInstanceNo).To(Equal("810984"))
			Expect(loadBalancerInstance.VirtualIP).To(Equal("49.236.149.111,49.236.149.189"))
			Expect(loadBalancerInstance.LoadBalancerName).To(Equal(reqParams.LoadBalancerName))
			Expect(loadBalancerInstance.LoadBalancerAlgorithmType.Code).To(Equal("RR"))
			Expect(loadBalancerInstance.LoadBalancerAlgorithmType.CodeName).To(Equal("Round Robin"))
			Expect(loadBalancerInstance.LoadBalancerDescription).To(Equal(reqParams.LoadBalancerDescription))
			Expect(loadBalancerInstance.CreateDate).To(Equal("2018-06-07T21:05:52+0900"))
			Expect(loadBalancerInstance.DomainName).To(Equal("slb-810984.ncloudslb.com"))
			Expect(loadBalancerInstance.InternetLineType.Code).To(Equal("PUBLC"))
			Expect(loadBalancerInstance.InternetLineType.CodeName).To(Equal("PUBLC"))
			Expect(loadBalancerInstance.LoadBalancerInstanceStatusName).To(Equal("creating"))
			Expect(loadBalancerInstance.LoadBalancerInstanceStatus.Code).To(Equal("INIT"))
			Expect(loadBalancerInstance.LoadBalancerInstanceStatus.CodeName).To(Equal("NET INIT state"))
			Expect(loadBalancerInstance.LoadBalancerInstanceOperation.Code).To(Equal("USE"))
			Expect(loadBalancerInstance.LoadBalancerInstanceOperation.CodeName).To(Equal("NET USE OP"))
			Expect(loadBalancerInstance.NetworkUsageType.Code).To(Equal("PBLIP"))
			Expect(loadBalancerInstance.NetworkUsageType.CodeName).To(Equal("Public"))
			Expect(loadBalancerInstance.IsHTTPKeepAlive).To(Equal(false))
			Expect(loadBalancerInstance.ConnectionTimeout).To(Equal(60))
			Expect(loadBalancerInstance.CertificateName).To(Equal("aaa"))

			Expect(len(loadBalancerInstance.LoadBalancerRuleList)).To(Equal(1))
			Expect(len(loadBalancerInstance.LoadBalancedServerInstanceList)).To(Equal(2))

			loadBalancerRule := loadBalancerInstance.LoadBalancerRuleList[0]
			Expect(loadBalancerRule.ProtocolType.Code).To(Equal("HTTPS"))
			Expect(loadBalancerRule.ProtocolType.CodeName).To(Equal("https"))
			Expect(loadBalancerRule.LoadBalancerPort).To(Equal(443))
			Expect(loadBalancerRule.ServerPort).To(Equal(443))
			Expect(loadBalancerRule.L7HealthCheckPath).To(Equal("/monitor"))
			Expect(loadBalancerRule.CertificateName).To(Equal("aaa"))
			Expect(loadBalancerRule.ProxyProtocolUseYn).To(Equal("N"))

			loadBalancedServerInstance := loadBalancerInstance.LoadBalancedServerInstanceList[0]
			Expect(len(loadBalancedServerInstance.ServerInstanceList)).To(Equal(1))
			Expect(len(loadBalancedServerInstance.ServerHealthCheckStatusList)).To(Equal(1))

			serverInstance := loadBalancedServerInstance.ServerInstanceList[0]
			serverHealthCheckStatus := loadBalancedServerInstance.ServerHealthCheckStatusList[0]

			Expect(serverInstance.ServerInstanceNo).To(Equal("805838"))
			Expect(serverInstance.ServerName).To(Equal("aaa004"))

			Expect(serverHealthCheckStatus.ProtocolType.Code).To(Equal("HTTPS"))
			Expect(serverHealthCheckStatus.ProtocolType.CodeName).To(Equal("https"))
			Expect(serverHealthCheckStatus.LoadBalancerPort).To(Equal(443))
			Expect(serverHealthCheckStatus.ServerPort).To(Equal(443))
			Expect(serverHealthCheckStatus.L7HealthCheckPath).To(Equal("/monitor"))
			Expect(serverHealthCheckStatus.ProxyProtocolUseYn).To(Equal("N"))
			Expect(serverHealthCheckStatus.ServerStatus).To(Equal(false))
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
			reqParams := new(RequestCreateLoadBalancerInstance)
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode:  "HTTPS",
					LoadBalancerPort:  443,
					ServerPort:        443,
					L7HealthCheckPath: "/monitor",
					CertificateName:   "aaa",
				},
			}
			result, err := conn.CreateLoadBalancerInstance(reqParams)

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
			reqParams := new(RequestCreateLoadBalancerInstance)
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode:  "HTTPS",
					LoadBalancerPort:  443,
					ServerPort:        443,
					L7HealthCheckPath: "/monitor",
					CertificateName:   "aaa",
				},
			}
			result, err := conn.CreateLoadBalancerInstance(reqParams)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Expired url."))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})

	Describe("Check Arguments", func() {
		It("should be error : 'LoadBalancerRuleList is required'", func() {
			reqParams := new(RequestCreateLoadBalancerInstance)
			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateLoadBalancerInstance(reqParams)

			Expect(err.Error()).To(Equal("LoadBalancerRuleList is required"))
		})

		It("should be error : 'LoadBalancerRuleList is required'", func() {
			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateLoadBalancerInstance(nil)

			Expect(err.Error()).To(Equal("LoadBalancerRuleList is required"))
		})

		It("should be error : 'length of LoadBalancerName to be in the range (3 - 30)'", func() {
			reqParams := new(RequestCreateLoadBalancerInstance)
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode:  "HTTPS",
					LoadBalancerPort:  443,
					ServerPort:        443,
					L7HealthCheckPath: "/monitor",
					CertificateName:   "aaa",
				},
			}
			reqParams.LoadBalancerName = String(35)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateLoadBalancerInstance(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of LoadBalancerName to be in the range (3 - 30), got "))
		})

		It("should be error : 'LoadBalancerAlgorithmTypeCode should be RR or LC or SIPHS'", func() {
			reqParams := new(RequestCreateLoadBalancerInstance)
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode:  "HTTPS",
					LoadBalancerPort:  443,
					ServerPort:        443,
					L7HealthCheckPath: "/monitor",
					CertificateName:   "aaa",
				},
			}
			reqParams.LoadBalancerAlgorithmTypeCode = String(5)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateLoadBalancerInstance(reqParams)

			Expect(err.Error()).To(Equal("LoadBalancerAlgorithmTypeCode should be RR or LC or SIPHS"))
		})

		It("should be error : 'length of LoadBalancerDescription to be in the range (1 - 1000)'", func() {
			reqParams := new(RequestCreateLoadBalancerInstance)
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode:  "HTTPS",
					LoadBalancerPort:  443,
					ServerPort:        443,
					L7HealthCheckPath: "/monitor",
					CertificateName:   "aaa",
				},
			}
			reqParams.LoadBalancerDescription = String(1100)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateLoadBalancerInstance(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of LoadBalancerDescription to be in the range (1 - 1000), got "))
		})

		It("should be error : 'LoadBalancerRuleList.1.ProtocolTypeCode should be HTTP or HTTPS or TCP or SSL'", func() {
			reqParams := new(RequestCreateLoadBalancerInstance)
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode:  "H",
					LoadBalancerPort:  443,
					ServerPort:        443,
					L7HealthCheckPath: "/monitor",
					CertificateName:   "aaa",
				},
			}

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateLoadBalancerInstance(reqParams)

			Expect(err.Error()).To(Equal("loadBalancerRuleList.1.protocolTypeCode should be HTTP or HTTPS or TCP or SSL"))
		})

		It("should be error : 'LoadBalancerRuleList.1.loadBalancerPort cannot be lower than 65534: 44300'", func() {
			reqParams := new(RequestCreateLoadBalancerInstance)
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode:  "HTTP",
					LoadBalancerPort:  443000,
					ServerPort:        443,
					L7HealthCheckPath: "/monitor",
					CertificateName:   "aaa",
				},
			}

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateLoadBalancerInstance(reqParams)

			Expect(err.Error()).To(Equal("\"loadBalancerRuleList.1.loadBalancerPort\" cannot be higher than 65534: 443000"))
		})

		It("should be error : 'LoadBalancerRuleList.1.serverPort cannot be lower than 65534: 44300'", func() {
			reqParams := new(RequestCreateLoadBalancerInstance)
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode:  "HTTP",
					LoadBalancerPort:  443,
					ServerPort:        443000,
					L7HealthCheckPath: "/monitor",
					CertificateName:   "aaa",
				},
			}

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateLoadBalancerInstance(reqParams)

			Expect(err.Error()).To(Equal("\"loadBalancerRuleList.1.serverPort\" cannot be higher than 65534: 443000"))
		})

		It("should be error : 'expected length of L7HealthCheckPath to be in the range (1 - 600)'", func() {
			reqParams := new(RequestCreateLoadBalancerInstance)
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode:  "HTTP",
					LoadBalancerPort:  443,
					ServerPort:        443,
					L7HealthCheckPath: String(6450),
					CertificateName:   "aaa",
				},
			}

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateLoadBalancerInstance(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of L7HealthCheckPath to be in the range (1 - 600), got "))
		})

		It("should be error : 'loadBalancerRuleList.1.l7HealthCheckPath field is required'", func() {
			reqParams := new(RequestCreateLoadBalancerInstance)
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode: "HTTP",
					LoadBalancerPort: 443,
					ServerPort:       443,
					CertificateName:  "aaa",
				},
			}

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateLoadBalancerInstance(reqParams)

			Expect(err.Error()).To(ContainSubstring("loadBalancerRuleList.1.l7HealthCheckPath field is required"))
		})

		It("should be error : 'expected length of CertificateName to be in the range (1 - 300)'", func() {
			reqParams := new(RequestCreateLoadBalancerInstance)
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode:  "HTTPS",
					LoadBalancerPort:  443,
					ServerPort:        443,
					L7HealthCheckPath: "/monitor",
					CertificateName:   String(6450),
				},
			}

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateLoadBalancerInstance(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of loadBalancerRuleList.1.certificateName to be in the range (1 - 300), got "))
		})

		It("should be error : 'loadBalancerRuleList.1.certificateName field is required'", func() {
			reqParams := new(RequestCreateLoadBalancerInstance)
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode:  "HTTPS",
					LoadBalancerPort:  443,
					ServerPort:        443,
					L7HealthCheckPath: "/monitor",
				},
			}

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateLoadBalancerInstance(reqParams)

			Expect(err.Error()).To(ContainSubstring("loadBalancerRuleList.1.certificateName field is required"))
		})

		It("should be error : 'loadBalancerRuleList.1.proxyProtocolUseYn should be Y or N'", func() {
			reqParams := new(RequestCreateLoadBalancerInstance)
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode:   "HTTPS",
					LoadBalancerPort:   443,
					ServerPort:         443,
					L7HealthCheckPath:  "/monitor",
					CertificateName:    "aaa",
					ProxyProtocolUseYn: "K",
				},
			}

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateLoadBalancerInstance(reqParams)

			Expect(err.Error()).To(ContainSubstring("loadBalancerRuleList.1.proxyProtocolUseYn should be Y or N"))
		})

		It("should be error : 'InternetLineTypeCode should be PUBLC or GLBL'", func() {
			reqParams := new(RequestCreateLoadBalancerInstance)
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode:  "HTTPS",
					LoadBalancerPort:  443,
					ServerPort:        443,
					L7HealthCheckPath: "/monitor",
					CertificateName:   "aaa",
				},
			}
			reqParams.InternetLineTypeCode = String(5)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateLoadBalancerInstance(reqParams)

			Expect(err.Error()).To(Equal("InternetLineTypeCode should be PUBLC or GLBL"))
		})

		It("should be error : 'NetworkUsageTypeCode should be PBLIP or PRVT'", func() {
			reqParams := new(RequestCreateLoadBalancerInstance)
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode:  "HTTPS",
					LoadBalancerPort:  443,
					ServerPort:        443,
					L7HealthCheckPath: "/monitor",
					CertificateName:   "aaa",
				},
			}
			reqParams.NetworkUsageTypeCode = String(5)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateLoadBalancerInstance(reqParams)

			Expect(err.Error()).To(Equal("NetworkUsageTypeCode should be PBLIP or PRVT"))
		})
	})
})
