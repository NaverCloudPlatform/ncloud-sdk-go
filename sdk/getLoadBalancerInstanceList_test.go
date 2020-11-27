package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Load Balancer Instance List", func() {
	Describe("Get all Load Balancer Instance List", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/loadbalancer").
				Reply(http.StatusOK).BodyString(`
					<getLoadBalancerInstanceListResponse>
					<requestId>045d6705-d974-4912-9f3c-08f3c519257f</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>2</totalRows>
					<loadBalancerInstanceList>
						<loadBalancerInstance>
							<loadBalancerInstanceNo>810070</loadBalancerInstanceNo>
							<virtualIp>10.62.8.99</virtualIp>
							<loadBalancerName>aaa</loadBalancerName>
							<loadBalancerAlgorithmType>
								<code>RR</code>
								<codeName>Round Robin</codeName>
							</loadBalancerAlgorithmType>
							<loadBalancerDescription>aaaa</loadBalancerDescription>
							<createDate>2018-06-07T12:03:58+0900</createDate>
							<domainName></domainName>
							<internetLineType>
								<code>PUBLC</code>
								<codeName>PUBLC</codeName>
							</internetLineType>
							<loadBalancerInstanceStatusName>used</loadBalancerInstanceStatusName>
							<loadBalancerInstanceStatus>
								<code>USED</code>
								<codeName>NET USED state</codeName>
							</loadBalancerInstanceStatus>
							<loadBalancerInstanceOperation>
								<code>NULL</code>
								<codeName>NET NULL OP</codeName>
							</loadBalancerInstanceOperation>
							<networkUsageType>
								<code>PRVT</code>
								<codeName>Private </codeName>
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
									<l7HealthCheckPath>/monitor/l7check</l7HealthCheckPath>
									<certificateName></certificateName>
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
												<code>HTTP</code>
												<codeName>http</codeName>
											</protocolType>
											<loadBalancerPort>80</loadBalancerPort>
											<serverPort>80</serverPort>
											<l7HealthCheckPath>/monitor/l7check</l7HealthCheckPath>
											<proxyProtocolUseYn>N</proxyProtocolUseYn>
											<serverStatus>false</serverStatus>
										</serverHealthCheckStatus>
									</serverHealthCheckStatusList>
								</loadBalancedServerInstance>
								<loadBalancedServerInstance>
									<serverInstance>
										<serverInstanceNo>805832</serverInstanceNo>
										<serverName>aaa002</serverName>
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
										<privateIp>10.41.0.27</privateIp>
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
										<uptime>2018-06-04T15:19:05+0900</uptime>
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
											<l7HealthCheckPath>/monitor/l7check</l7HealthCheckPath>
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
												<code>HTTP</code>
												<codeName>http</codeName>
											</protocolType>
											<loadBalancerPort>80</loadBalancerPort>
											<serverPort>80</serverPort>
											<l7HealthCheckPath>/monitor/l7check</l7HealthCheckPath>
											<proxyProtocolUseYn>N</proxyProtocolUseYn>
											<serverStatus>false</serverStatus>
										</serverHealthCheckStatus>
									</serverHealthCheckStatusList>
								</loadBalancedServerInstance>
							</loadBalancedServerInstanceList>
						</loadBalancerInstance>
						<loadBalancerInstance>
							<loadBalancerInstanceNo>810075</loadBalancerInstanceNo>
							<virtualIp>220.230.126.154,210.89.187.132</virtualIp>
							<loadBalancerName>https</loadBalancerName>
							<loadBalancerAlgorithmType>
								<code>RR</code>
								<codeName>Round Robin</codeName>
							</loadBalancerAlgorithmType>
							<loadBalancerDescription>ggggg</loadBalancerDescription>
							<createDate>2018-06-07T12:05:45+0900</createDate>
							<domainName>slb-810075.ncloudslb.com</domainName>
							<internetLineType>
								<code>PUBLC</code>
								<codeName>PUBLC</codeName>
							</internetLineType>
							<loadBalancerInstanceStatusName>used</loadBalancerInstanceStatusName>
							<loadBalancerInstanceStatus>
								<code>USED</code>
								<codeName>NET USED state</codeName>
							</loadBalancerInstanceStatus>
							<loadBalancerInstanceOperation>
								<code>NULL</code>
								<codeName>NET NULL OP</codeName>
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
									<serverPort>80</serverPort>
									<l7HealthCheckPath>/monitor/l7check</l7HealthCheckPath>
									<certificateName>aaa</certificateName>
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
												<code>HTTPS</code>
												<codeName>https</codeName>
											</protocolType>
											<loadBalancerPort>443</loadBalancerPort>
											<serverPort>80</serverPort>
											<l7HealthCheckPath>/monitor/l7check</l7HealthCheckPath>
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
											<serverPort>80</serverPort>
											<l7HealthCheckPath>/monitor/l7check</l7HealthCheckPath>
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
												<code>HTTPS</code>
												<codeName>https</codeName>
											</protocolType>
											<loadBalancerPort>443</loadBalancerPort>
											<serverPort>80</serverPort>
											<l7HealthCheckPath>/monitor/l7check</l7HealthCheckPath>
											<proxyProtocolUseYn>N</proxyProtocolUseYn>
											<serverStatus>false</serverStatus>
										</serverHealthCheckStatus>
									</serverHealthCheckStatusList>
								</loadBalancedServerInstance>
							</loadBalancedServerInstanceList>
						</loadBalancerInstance>
					</loadBalancerInstanceList>
				</getLoadBalancerInstanceListResponse>
					`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get Load Balancer Instance list", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetLoadBalancerInstanceList(nil)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(2))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.LoadBalancerInstanceList)).To(Equal(2))

			loadBalancerInstance := result.LoadBalancerInstanceList[0]
			Expect(loadBalancerInstance.LoadBalancerInstanceNo).To(Equal("810070"))
			Expect(loadBalancerInstance.VirtualIP).To(Equal("10.62.8.99"))
			Expect(loadBalancerInstance.LoadBalancerName).To(Equal("aaa"))
			Expect(loadBalancerInstance.LoadBalancerAlgorithmType.Code).To(Equal("RR"))
			Expect(loadBalancerInstance.LoadBalancerAlgorithmType.CodeName).To(Equal("Round Robin"))
			Expect(loadBalancerInstance.LoadBalancerDescription).To(Equal("aaaa"))
			Expect(loadBalancerInstance.CreateDate).To(Equal("2018-06-07T12:03:58+0900"))
			Expect(loadBalancerInstance.DomainName).To(Equal(""))
			Expect(loadBalancerInstance.InternetLineType.Code).To(Equal("PUBLC"))
			Expect(loadBalancerInstance.InternetLineType.CodeName).To(Equal("PUBLC"))
			Expect(loadBalancerInstance.LoadBalancerInstanceStatusName).To(Equal("used"))
			Expect(loadBalancerInstance.LoadBalancerInstanceStatus.Code).To(Equal("USED"))
			Expect(loadBalancerInstance.LoadBalancerInstanceStatus.CodeName).To(Equal("NET USED state"))
			Expect(loadBalancerInstance.LoadBalancerInstanceOperation.Code).To(Equal("NULL"))
			Expect(loadBalancerInstance.LoadBalancerInstanceOperation.CodeName).To(Equal("NET NULL OP"))
			Expect(loadBalancerInstance.NetworkUsageType.Code).To(Equal("PRVT"))
			Expect(loadBalancerInstance.NetworkUsageType.CodeName).To(Equal("Private "))
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
			Expect(loadBalancerRule.L7HealthCheckPath).To(Equal("/monitor/l7check"))
			Expect(loadBalancerRule.CertificateName).To(Equal(""))
			Expect(loadBalancerRule.ProxyProtocolUseYn).To(Equal("N"))

			loadBalancedServerInstance := loadBalancerInstance.LoadBalancedServerInstanceList[0]
			Expect(len(loadBalancedServerInstance.ServerInstanceList)).To(Equal(1))
			Expect(len(loadBalancedServerInstance.ServerHealthCheckStatusList)).To(Equal(1))

			serverInstance := loadBalancedServerInstance.ServerInstanceList[0]
			serverHealthCheckStatus := loadBalancedServerInstance.ServerHealthCheckStatusList[0]

			Expect(serverInstance.ServerInstanceNo).To(Equal("805838"))
			Expect(serverInstance.ServerName).To(Equal("aaa004"))

			Expect(serverHealthCheckStatus.ProtocolType.Code).To(Equal("HTTP"))
			Expect(serverHealthCheckStatus.ProtocolType.CodeName).To(Equal("http"))
			Expect(serverHealthCheckStatus.LoadBalancerPort).To(Equal(80))
			Expect(serverHealthCheckStatus.ServerPort).To(Equal(80))
			Expect(serverHealthCheckStatus.L7HealthCheckPath).To(Equal("/monitor/l7check"))
			Expect(serverHealthCheckStatus.ProxyProtocolUseYn).To(Equal("N"))
			Expect(serverHealthCheckStatus.ServerStatus).To(Equal(false))
		})
	})

	Describe("Get One Load Balancer Instance List which LoadBalancerInstanceNoList.1 is 810070", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/loadbalancer").
				Reply(http.StatusOK).BodyString(`
					<getLoadBalancerInstanceListResponse>
					<requestId>045d6705-d974-4912-9f3c-08f3c519257f</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<loadBalancerInstanceList>
						<loadBalancerInstance>
							<loadBalancerInstanceNo>810070</loadBalancerInstanceNo>
							<virtualIp>10.62.8.99</virtualIp>
							<loadBalancerName>aaa</loadBalancerName>
							<loadBalancerAlgorithmType>
								<code>RR</code>
								<codeName>Round Robin</codeName>
							</loadBalancerAlgorithmType>
							<loadBalancerDescription>aaaa</loadBalancerDescription>
							<createDate>2018-06-07T12:03:58+0900</createDate>
							<domainName></domainName>
							<internetLineType>
								<code>PUBLC</code>
								<codeName>PUBLC</codeName>
							</internetLineType>
							<loadBalancerInstanceStatusName>used</loadBalancerInstanceStatusName>
							<loadBalancerInstanceStatus>
								<code>USED</code>
								<codeName>NET USED state</codeName>
							</loadBalancerInstanceStatus>
							<loadBalancerInstanceOperation>
								<code>NULL</code>
								<codeName>NET NULL OP</codeName>
							</loadBalancerInstanceOperation>
							<networkUsageType>
								<code>PRVT</code>
								<codeName>Private </codeName>
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
									<l7HealthCheckPath>/monitor/l7check</l7HealthCheckPath>
									<certificateName></certificateName>
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
												<code>HTTP</code>
												<codeName>http</codeName>
											</protocolType>
											<loadBalancerPort>80</loadBalancerPort>
											<serverPort>80</serverPort>
											<l7HealthCheckPath>/monitor/l7check</l7HealthCheckPath>
											<proxyProtocolUseYn>N</proxyProtocolUseYn>
											<serverStatus>false</serverStatus>
										</serverHealthCheckStatus>
									</serverHealthCheckStatusList>
								</loadBalancedServerInstance>
								<loadBalancedServerInstance>
									<serverInstance>
										<serverInstanceNo>805832</serverInstanceNo>
										<serverName>aaa002</serverName>
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
										<privateIp>10.41.0.27</privateIp>
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
										<uptime>2018-06-04T15:19:05+0900</uptime>
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
											<l7HealthCheckPath>/monitor/l7check</l7HealthCheckPath>
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
												<code>HTTP</code>
												<codeName>http</codeName>
											</protocolType>
											<loadBalancerPort>80</loadBalancerPort>
											<serverPort>80</serverPort>
											<l7HealthCheckPath>/monitor/l7check</l7HealthCheckPath>
											<proxyProtocolUseYn>N</proxyProtocolUseYn>
											<serverStatus>false</serverStatus>
										</serverHealthCheckStatus>
									</serverHealthCheckStatusList>
								</loadBalancedServerInstance>
							</loadBalancedServerInstanceList>
						</loadBalancerInstance>
					</loadBalancerInstanceList>
				</getLoadBalancerInstanceListResponse>
					`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get One Load Balancer Instance", func() {
			reqParams := new(RequestLoadBalancerInstanceList)
			reqParams.LoadBalancerInstanceNoList = []string{"810070"}
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetLoadBalancerInstanceList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.LoadBalancerInstanceList)).To(Equal(1))

			loadBalancerInstance := result.LoadBalancerInstanceList[0]
			Expect(loadBalancerInstance.LoadBalancerInstanceNo).To(Equal("810070"))
			Expect(loadBalancerInstance.VirtualIP).To(Equal("10.62.8.99"))
			Expect(loadBalancerInstance.LoadBalancerName).To(Equal("aaa"))
			Expect(loadBalancerInstance.LoadBalancerAlgorithmType.Code).To(Equal("RR"))
			Expect(loadBalancerInstance.LoadBalancerAlgorithmType.CodeName).To(Equal("Round Robin"))
			Expect(loadBalancerInstance.LoadBalancerDescription).To(Equal("aaaa"))
			Expect(loadBalancerInstance.CreateDate).To(Equal("2018-06-07T12:03:58+0900"))
			Expect(loadBalancerInstance.DomainName).To(Equal(""))
			Expect(loadBalancerInstance.InternetLineType.Code).To(Equal("PUBLC"))
			Expect(loadBalancerInstance.InternetLineType.CodeName).To(Equal("PUBLC"))
			Expect(loadBalancerInstance.LoadBalancerInstanceStatusName).To(Equal("used"))
			Expect(loadBalancerInstance.LoadBalancerInstanceStatus.Code).To(Equal("USED"))
			Expect(loadBalancerInstance.LoadBalancerInstanceStatus.CodeName).To(Equal("NET USED state"))
			Expect(loadBalancerInstance.LoadBalancerInstanceOperation.Code).To(Equal("NULL"))
			Expect(loadBalancerInstance.LoadBalancerInstanceOperation.CodeName).To(Equal("NET NULL OP"))
			Expect(loadBalancerInstance.NetworkUsageType.Code).To(Equal("PRVT"))
			Expect(loadBalancerInstance.NetworkUsageType.CodeName).To(Equal("Private "))
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
			Expect(loadBalancerRule.L7HealthCheckPath).To(Equal("/monitor/l7check"))
			Expect(loadBalancerRule.CertificateName).To(Equal(""))
			Expect(loadBalancerRule.ProxyProtocolUseYn).To(Equal("N"))

			loadBalancedServerInstance := loadBalancerInstance.LoadBalancedServerInstanceList[0]
			Expect(len(loadBalancedServerInstance.ServerInstanceList)).To(Equal(1))
			Expect(len(loadBalancedServerInstance.ServerHealthCheckStatusList)).To(Equal(1))

			serverInstance := loadBalancedServerInstance.ServerInstanceList[0]
			serverHealthCheckStatus := loadBalancedServerInstance.ServerHealthCheckStatusList[0]

			Expect(serverInstance.ServerInstanceNo).To(Equal("805838"))
			Expect(serverInstance.ServerName).To(Equal("aaa004"))

			Expect(serverHealthCheckStatus.ProtocolType.Code).To(Equal("HTTP"))
			Expect(serverHealthCheckStatus.ProtocolType.CodeName).To(Equal("http"))
			Expect(serverHealthCheckStatus.LoadBalancerPort).To(Equal(80))
			Expect(serverHealthCheckStatus.ServerPort).To(Equal(80))
			Expect(serverHealthCheckStatus.L7HealthCheckPath).To(Equal("/monitor/l7check"))
			Expect(serverHealthCheckStatus.ProxyProtocolUseYn).To(Equal("N"))
			Expect(serverHealthCheckStatus.ServerStatus).To(Equal(false))
		})
	})

	Describe("There is no Load Balancer Instance list", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/loadbalancer").
				Reply(http.StatusOK).BodyString(`<getLoadBalancerInstanceListResponse>
					<requestId>9b6e2bac-3ca2-4c97-8dec-b432f756b198</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>0</totalRows>
					<loadBalancerInstanceList/>
					</getLoadBalancerInstanceListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get No Load Balancer Instance list", func() {
			reqParams := new(RequestLoadBalancerInstanceList)
			reqParams.LoadBalancerInstanceNoList = []string{"111111"}
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetLoadBalancerInstanceList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(0))
			Expect(len(result.LoadBalancerInstanceList)).To(Equal(0))
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
					<returnMessage>Invalid consumerKey</returnMessage>
					</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should be failed by authorization fail", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetLoadBalancerInstanceList(nil)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Invalid consumerKey"))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})

	Describe("Expired URL", func() {
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
			result, err := conn.GetLoadBalancerInstanceList(nil)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Expired url."))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})

	Describe("Check Arguments", func() {
		It("should be error : 'InternetLineTypeCode should be PUBLC or GLBL'", func() {
			reqParams := new(RequestLoadBalancerInstanceList)
			reqParams.InternetLineTypeCode = String(5)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.GetLoadBalancerInstanceList(reqParams)

			Expect(err.Error()).To(Equal("InternetLineTypeCode should be PUBLC or GLBL"))
		})

		It("should be error : 'NetworkUsageTypeCode should be PBLIP or PRVT'", func() {
			reqParams := new(RequestLoadBalancerInstanceList)
			reqParams.NetworkUsageTypeCode = String(5)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.GetLoadBalancerInstanceList(reqParams)

			Expect(err.Error()).To(Equal("NetworkUsageTypeCode should be PBLIP or PRVT"))
		})

		It("should be error : 'Value of PageNo should be min 0 or max 2147483648'", func() {
			reqParams := new(RequestLoadBalancerInstanceList)
			reqParams.PageNo = -1

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.GetLoadBalancerInstanceList(reqParams)

			Expect(err.Error()).To(Equal("\"PageNo\" cannot be lower than 0: -1"))
		})

		It("should be error : 'Value of PageNo should be min 0 or max 2147483648'", func() {
			reqParams := new(RequestLoadBalancerInstanceList)
			reqParams.PageNo = 2147483648

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.GetLoadBalancerInstanceList(reqParams)

			Expect(err.Error()).To(Equal("\"PageNo\" cannot be higher than 2147483647: 2147483648"))
		})

		It("should be error : 'Value of PageSize should be min 0 or max 2147483648'", func() {
			reqParams := new(RequestLoadBalancerInstanceList)
			reqParams.PageSize = -1

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.GetLoadBalancerInstanceList(reqParams)

			Expect(err.Error()).To(Equal("\"PageSize\" cannot be lower than 0: -1"))
		})

		It("should be error : 'Value of PageSize should be min 0 or max 2147483648'", func() {
			reqParams := new(RequestLoadBalancerInstanceList)
			reqParams.PageSize = 2147483648

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.GetLoadBalancerInstanceList(reqParams)

			Expect(err.Error()).To(Equal("\"PageSize\" cannot be higher than 2147483647: 2147483648"))
		})

		It("should be error : 'SortedBy should be loadBalancerName or loadBalancerInstanceNo'", func() {
			reqParams := new(RequestLoadBalancerInstanceList)
			reqParams.SortedBy = "wrong sortedByValue"

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.GetLoadBalancerInstanceList(reqParams)

			Expect(err.Error()).To(Equal("SortedBy should be loadBalancerName or loadBalancerInstanceNo"))
		})

		It("should be error : 'SortingOrder should be ascending or descending'", func() {
			reqParams := new(RequestLoadBalancerInstanceList)
			reqParams.SortingOrder = "wrong SortingOrderValue"

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.GetLoadBalancerInstanceList(reqParams)

			Expect(err.Error()).To(Equal("SortingOrder should be ascending or descending"))
		})
	})
})
