package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Delete Load Balancer Instances", func() {
	Describe("Delete Load Balancer Instance", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/loadbalancer").
				Reply(http.StatusOK).BodyString(`
					<deleteLoadBalancerInstancesResponse>
					<requestId>ebd98972-db81-4432-b0bc-27f0fdbfe8bb</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<loadBalancerInstanceList>
						<loadBalancerInstance>
							<loadBalancerInstanceNo>811612</loadBalancerInstanceNo>
							<virtualIp>210.89.187.243,210.89.187.60</virtualIp>
							<loadBalancerName>aaa</loadBalancerName>
							<loadBalancerAlgorithmType>
								<code>RR</code>
								<codeName>Round Robin</codeName>
							</loadBalancerAlgorithmType>
							<loadBalancerDescription/>
							<createDate>2018-06-08T10:26:43+0900</createDate>
							<domainName>slb-811612.ncloudslb.com</domainName>
							<internetLineType>
								<code>PUBLC</code>
								<codeName>PUBLC</codeName>
							</internetLineType>
							<loadBalancerInstanceStatusName>terminating</loadBalancerInstanceStatusName>
							<loadBalancerInstanceStatus>
								<code>USED</code>
								<codeName>NET USED state</codeName>
							</loadBalancerInstanceStatus>
							<loadBalancerInstanceOperation>
								<code>TERMT</code>
								<codeName>NET TERMINATED OP</codeName>
							</loadBalancerInstanceOperation>
							<networkUsageType>
								<code>PBLIP</code>
								<codeName>Public</codeName>
							</networkUsageType>
							<isHttpKeepAlive>false</isHttpKeepAlive>
							<connectionTimeout>60</connectionTimeout>
							<certificateName/>
							<loadBalancerRuleList>
								<loadBalancerRule>
									<protocolType>
										<code>HTTP</code>
										<codeName>http</codeName>
									</protocolType>
									<loadBalancerPort>80</loadBalancerPort>
									<serverPort>80</serverPort>
									<l7HealthCheckPath>/</l7HealthCheckPath>
									<certificateName/>
									<proxyProtocolUseYn>N</proxyProtocolUseYn>
								</loadBalancerRule>
							</loadBalancerRuleList>
							<loadBalancedServerInstanceList>
								<loadBalancedServerInstance>
									<serverInstance>
										<serverInstanceNo>805847</serverInstanceNo>
										<serverName>aaa007</serverName>
										<serverDescription/>
										<cpuCount>2</cpuCount>
										<memorySize>4294967296</memorySize>
										<baseBlockStorageSize>53687091200</baseBlockStorageSize>
										<platformType>
											<code>LNX64</code>
											<codeName>Linux 64 Bit</codeName>
										</platformType>
										<loginKeyName>consolekey</loginKeyName>
										<isFeeChargingMonitoring>false</isFeeChargingMonitoring>
										<publicIp/>
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
										<userData/>
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
				</deleteLoadBalancerInstancesResponse>
					`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should delete Load Balancer Instance", func() {
			conn := NewConnection(accessKey, secretKey)
			reqParams := new(RequestDeleteLoadBalancerInstances)
			reqParams.LoadBalancerInstanceNoList = []string{"811649"}
			result, err := conn.DeleteLoadBalancerInstances(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.LoadBalancerInstanceList)).To(Equal(1))

			loadBalancerInstance := result.LoadBalancerInstanceList[0]
			Expect(loadBalancerInstance.LoadBalancerInstanceNo).To(Equal("811612"))
			Expect(loadBalancerInstance.VirtualIP).To(Equal("210.89.187.243,210.89.187.60"))
			Expect(loadBalancerInstance.LoadBalancerName).To(Equal("aaa"))
			Expect(loadBalancerInstance.LoadBalancerAlgorithmType.Code).To(Equal("RR"))
			Expect(loadBalancerInstance.LoadBalancerAlgorithmType.CodeName).To(Equal("Round Robin"))
			Expect(loadBalancerInstance.LoadBalancerDescription).To(Equal(""))
			Expect(loadBalancerInstance.CreateDate).To(Equal("2018-06-08T10:26:43+0900"))
			Expect(loadBalancerInstance.DomainName).To(Equal("slb-811612.ncloudslb.com"))
			Expect(loadBalancerInstance.InternetLineType.Code).To(Equal("PUBLC"))
			Expect(loadBalancerInstance.InternetLineType.CodeName).To(Equal("PUBLC"))
			Expect(loadBalancerInstance.LoadBalancerInstanceStatusName).To(Equal("terminating"))
			Expect(loadBalancerInstance.LoadBalancerInstanceStatus.Code).To(Equal("USED"))
			Expect(loadBalancerInstance.LoadBalancerInstanceStatus.CodeName).To(Equal("NET USED state"))
			Expect(loadBalancerInstance.LoadBalancerInstanceOperation.Code).To(Equal("TERMT"))
			Expect(loadBalancerInstance.LoadBalancerInstanceOperation.CodeName).To(Equal("NET TERMINATED OP"))
			Expect(loadBalancerInstance.NetworkUsageType.Code).To(Equal("PBLIP"))
			Expect(loadBalancerInstance.NetworkUsageType.CodeName).To(Equal("Public"))
			Expect(loadBalancerInstance.IsHTTPKeepAlive).To(Equal(false))
			Expect(loadBalancerInstance.ConnectionTimeout).To(Equal(60))
			Expect(loadBalancerInstance.CertificateName).To(Equal(""))

			Expect(len(loadBalancerInstance.LoadBalancerRuleList)).To(Equal(1))
			Expect(len(loadBalancerInstance.LoadBalancedServerInstanceList)).To(Equal(1))

			loadBalancerRule := loadBalancerInstance.LoadBalancerRuleList[0]
			Expect(loadBalancerRule.ProtocolType.Code).To(Equal("HTTP"))
			Expect(loadBalancerRule.ProtocolType.CodeName).To(Equal("http"))
			Expect(loadBalancerRule.LoadBalancerPort).To(Equal(80))
			Expect(loadBalancerRule.ServerPort).To(Equal(80))
			Expect(loadBalancerRule.L7HealthCheckPath).To(Equal("/"))
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
			Expect(serverHealthCheckStatus.L7HealthCheckPath).To(Equal("/"))
			Expect(serverHealthCheckStatus.ProxyProtocolUseYn).To(Equal("N"))
			Expect(serverHealthCheckStatus.ServerStatus).To(Equal(false))
		})
	})

	Describe("There is no Load Balancer Instances to be delete", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/loadbalancer").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
					<returnCode>10713</returnCode>
					<returnMessage>
				Not found contract information. Please check your input parameter.
					</returnMessage>
				</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should be failed by No Load Balancer Instances", func() {
			reqParams := new(RequestDeleteLoadBalancerInstances)
			reqParams.LoadBalancerInstanceNoList = []string{"810070"}
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.DeleteLoadBalancerInstances(reqParams)

			Expect(result.ReturnCode).To(Equal(10713))
			Expect(result.ReturnMessage).To(Equal("Not found contract information. Please check your input parameter."))
			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
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
			reqParams := new(RequestDeleteLoadBalancerInstances)
			reqParams.LoadBalancerInstanceNoList = []string{"810070"}
			result, err := conn.DeleteLoadBalancerInstances(reqParams)

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
			reqParams := new(RequestDeleteLoadBalancerInstances)
			reqParams.LoadBalancerInstanceNoList = []string{"810070"}
			result, err := conn.DeleteLoadBalancerInstances(reqParams)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Expired url."))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})

	Describe("Check Arguments", func() {
		It("should be error : 'LoadBalancerInstanceNoList is required'", func() {
			conn := NewConnection(accessKey, secretKey)
			_, err := conn.DeleteLoadBalancerInstances(nil)

			Expect(err.Error()).To(Equal("LoadBalancerInstanceNoList is required field"))
		})

		It("should be error : 'LoadBalancerInstanceNoList is required'", func() {
			reqParams := new(RequestDeleteLoadBalancerInstances)
			conn := NewConnection(accessKey, secretKey)
			_, err := conn.DeleteLoadBalancerInstances(reqParams)

			Expect(err.Error()).To(Equal("LoadBalancerInstanceNoList is required field"))
		})
	})
})
