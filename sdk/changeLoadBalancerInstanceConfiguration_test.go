package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Change Load Balancer Instance", func() {
	Describe("Change http Load Balancer Instance", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/loadbalancer").
				Reply(http.StatusOK).BodyString(`
					<changeLoadBalancerInstanceConfigurationResponse>
					<requestId>fdf9dc33-7d6f-45c7-a86c-fb93e5e3a6eb</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<loadBalancerInstanceList>
						<loadBalancerInstance>
							<loadBalancerInstanceNo>811670</loadBalancerInstanceNo>
							<virtualIp>210.89.187.105,210.89.187.245</virtualIp>
							<loadBalancerName>a12</loadBalancerName>
							<loadBalancerAlgorithmType>
								<code>SIPHS</code>
								<codeName>Source IP Hash</codeName>
							</loadBalancerAlgorithmType>
							<loadBalancerDescription>LBTest Changed Description</loadBalancerDescription>
							<createDate>2018-06-08T11:25:28+0900</createDate>
							<domainName>slb-811670.ncloudslb.com</domainName>
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
				</changeLoadBalancerInstanceConfigurationResponse>
					`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should Change Load Balancer Instance", func() {
			conn := NewConnection(accessKey, secretKey)
			reqParams := new(RequestChangeLoadBalancerInstanceConfiguration)
			reqParams.LoadBalancerInstanceNo = "811670"
			reqParams.LoadBalancerAlgorithmTypeCode = "SIPHS"
			reqParams.LoadBalancerDescription = "LBTest Changed Description"
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode:  "HTTPS",
					LoadBalancerPort:  443,
					ServerPort:        443,
					L7HealthCheckPath: "/monitor",
					CertificateName:   "aaa",
				},
			}
			result, err := conn.ChangeLoadBalancerInstanceConfiguration(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.LoadBalancerInstanceList)).To(Equal(1))

			loadBalancerInstance := result.LoadBalancerInstanceList[0]
			Expect(loadBalancerInstance.LoadBalancerInstanceNo).To(Equal(reqParams.LoadBalancerInstanceNo))
			Expect(loadBalancerInstance.VirtualIP).To(Equal("210.89.187.105,210.89.187.245"))
			Expect(loadBalancerInstance.LoadBalancerAlgorithmType.Code).To(Equal(reqParams.LoadBalancerAlgorithmTypeCode))
			Expect(loadBalancerInstance.LoadBalancerAlgorithmType.CodeName).To(Equal("Source IP Hash"))
			Expect(loadBalancerInstance.LoadBalancerDescription).To(Equal(reqParams.LoadBalancerDescription))
			Expect(loadBalancerInstance.CreateDate).To(Equal("2018-06-08T11:25:28+0900"))
			Expect(loadBalancerInstance.DomainName).To(Equal("slb-811670.ncloudslb.com"))
			Expect(loadBalancerInstance.InternetLineType.Code).To(Equal("PUBLC"))
			Expect(loadBalancerInstance.InternetLineType.CodeName).To(Equal("PUBLC"))
			Expect(loadBalancerInstance.LoadBalancerInstanceStatusName).To(Equal("changing"))
			Expect(loadBalancerInstance.LoadBalancerInstanceStatus.Code).To(Equal("USED"))
			Expect(loadBalancerInstance.LoadBalancerInstanceStatus.CodeName).To(Equal("NET USED state"))
			Expect(loadBalancerInstance.LoadBalancerInstanceOperation.Code).To(Equal("CHANG"))
			Expect(loadBalancerInstance.LoadBalancerInstanceOperation.CodeName).To(Equal("NET CHANGE OP"))
			Expect(loadBalancerInstance.NetworkUsageType.Code).To(Equal("PBLIP"))
			Expect(loadBalancerInstance.NetworkUsageType.CodeName).To(Equal("Public"))
			Expect(loadBalancerInstance.IsHTTPKeepAlive).To(Equal(false))
			Expect(loadBalancerInstance.ConnectionTimeout).To(Equal(60))
			Expect(loadBalancerInstance.CertificateName).To(Equal("aaa"))

			Expect(len(loadBalancerInstance.LoadBalancerRuleList)).To(Equal(1))
			Expect(len(loadBalancerInstance.LoadBalancedServerInstanceList)).To(Equal(1))

			loadBalancerRule := loadBalancerInstance.LoadBalancerRuleList[0]
			Expect(loadBalancerRule.ProtocolType.Code).To(Equal(reqParams.LoadBalancerRuleList[0].ProtocolTypeCode))
			Expect(loadBalancerRule.ProtocolType.CodeName).To(Equal("https"))
			Expect(loadBalancerRule.LoadBalancerPort).To(Equal(reqParams.LoadBalancerRuleList[0].LoadBalancerPort))
			Expect(loadBalancerRule.ServerPort).To(Equal(reqParams.LoadBalancerRuleList[0].ServerPort))
			Expect(loadBalancerRule.L7HealthCheckPath).To(Equal(reqParams.LoadBalancerRuleList[0].L7HealthCheckPath))
			Expect(loadBalancerRule.CertificateName).To(Equal(reqParams.LoadBalancerRuleList[0].CertificateName))
			Expect(loadBalancerRule.ProxyProtocolUseYn).To(Equal("N"))
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
			reqParams := new(RequestChangeLoadBalancerInstanceConfiguration)
			reqParams.LoadBalancerInstanceNo = "811670"
			reqParams.LoadBalancerAlgorithmTypeCode = "SIPHS"
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode:  "HTTPS",
					LoadBalancerPort:  443,
					ServerPort:        443,
					L7HealthCheckPath: "/monitor",
					CertificateName:   "aaa",
				},
			}
			result, err := conn.ChangeLoadBalancerInstanceConfiguration(reqParams)

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
			reqParams := new(RequestChangeLoadBalancerInstanceConfiguration)
			reqParams.LoadBalancerInstanceNo = "811670"
			reqParams.LoadBalancerAlgorithmTypeCode = "SIPHS"
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode:  "HTTPS",
					LoadBalancerPort:  443,
					ServerPort:        443,
					L7HealthCheckPath: "/monitor",
					CertificateName:   "aaa",
				},
			}
			result, err := conn.ChangeLoadBalancerInstanceConfiguration(reqParams)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Expired url."))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})

	Describe("Check Arguments", func() {
		It("should be error : 'LoadBalancerInstanceNo is required'", func() {
			reqParams := new(RequestChangeLoadBalancerInstanceConfiguration)
			conn := NewConnection(accessKey, secretKey)
			_, err := conn.ChangeLoadBalancerInstanceConfiguration(reqParams)

			Expect(err.Error()).To(Equal("LoadBalancerInstanceNo field is required"))
		})

		It("should be error : 'LoadBalancerInstanceNo is required'", func() {
			conn := NewConnection(accessKey, secretKey)
			_, err := conn.ChangeLoadBalancerInstanceConfiguration(nil)

			Expect(err.Error()).To(Equal("LoadBalancerInstanceNo field is required"))
		})

		It("should be error : 'LoadBalancerAlgorithmTypeCode should be required'", func() {
			reqParams := new(RequestChangeLoadBalancerInstanceConfiguration)
			reqParams.LoadBalancerInstanceNo = "811670"
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode:  "HTTPS",
					LoadBalancerPort:  443,
					ServerPort:        443,
					L7HealthCheckPath: "/monitor",
					CertificateName:   "aaa",
				},
			}
			conn := NewConnection(accessKey, secretKey)
			_, err := conn.ChangeLoadBalancerInstanceConfiguration(reqParams)

			Expect(err.Error()).To(Equal("LoadBalancerAlgorithmTypeCode field is required"))
		})

		It("should be error : 'LoadBalancerAlgorithmTypeCode should be RR or LC or SIPHS'", func() {
			reqParams := new(RequestChangeLoadBalancerInstanceConfiguration)
			reqParams.LoadBalancerInstanceNo = "811670"
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode:  "HTTPS",
					LoadBalancerPort:  443,
					ServerPort:        443,
					L7HealthCheckPath: "/monitor",
					CertificateName:   "aaa",
				},
			}
			reqParams.LoadBalancerAlgorithmTypeCode = String(6)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.ChangeLoadBalancerInstanceConfiguration(reqParams)

			Expect(err.Error()).To(Equal("LoadBalancerAlgorithmTypeCode should be RR or LC or SIPHS"))
		})

		It("should be error : 'length of LoadBalancerDescription to be in the range (1 - 1000)'", func() {
			reqParams := new(RequestChangeLoadBalancerInstanceConfiguration)
			reqParams.LoadBalancerInstanceNo = "811670"
			reqParams.LoadBalancerAlgorithmTypeCode = "SIPHS"
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
			_, err := conn.ChangeLoadBalancerInstanceConfiguration(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of LoadBalancerDescription to be in the range (1 - 1000), got "))
		})

		It("should be error : 'LoadBalancerRuleList.1.ProtocolTypeCode should be HTTP or HTTPS or TCP or SSL'", func() {
			reqParams := new(RequestChangeLoadBalancerInstanceConfiguration)
			reqParams.LoadBalancerInstanceNo = "811670"
			reqParams.LoadBalancerAlgorithmTypeCode = "SIPHS"
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
			_, err := conn.ChangeLoadBalancerInstanceConfiguration(reqParams)

			Expect(err.Error()).To(Equal("loadBalancerRuleList.1.protocolTypeCode should be HTTP or HTTPS or TCP or SSL"))
		})

		It("should be error : 'LoadBalancerRuleList.1.loadBalancerPort cannot be lower than 65534: 44300'", func() {
			reqParams := new(RequestChangeLoadBalancerInstanceConfiguration)
			reqParams.LoadBalancerInstanceNo = "811670"
			reqParams.LoadBalancerAlgorithmTypeCode = "SIPHS"
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
			_, err := conn.ChangeLoadBalancerInstanceConfiguration(reqParams)

			Expect(err.Error()).To(Equal("\"loadBalancerRuleList.1.loadBalancerPort\" cannot be higher than 65534: 443000"))
		})

		It("should be error : 'LoadBalancerRuleList.1.serverPort cannot be lower than 65534: 44300'", func() {
			reqParams := new(RequestChangeLoadBalancerInstanceConfiguration)
			reqParams.LoadBalancerInstanceNo = "811670"
			reqParams.LoadBalancerAlgorithmTypeCode = "SIPHS"
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
			_, err := conn.ChangeLoadBalancerInstanceConfiguration(reqParams)

			Expect(err.Error()).To(Equal("\"loadBalancerRuleList.1.serverPort\" cannot be higher than 65534: 443000"))
		})

		It("should be error : 'expected length of L7HealthCheckPath to be in the range (1 - 600)'", func() {
			reqParams := new(RequestChangeLoadBalancerInstanceConfiguration)
			reqParams.LoadBalancerInstanceNo = "811670"
			reqParams.LoadBalancerAlgorithmTypeCode = "SIPHS"
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
			_, err := conn.ChangeLoadBalancerInstanceConfiguration(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of L7HealthCheckPath to be in the range (1 - 600), got "))
		})

		It("should be error : 'loadBalancerRuleList.1.l7HealthCheckPath field is required'", func() {
			reqParams := new(RequestChangeLoadBalancerInstanceConfiguration)
			reqParams.LoadBalancerInstanceNo = "811670"
			reqParams.LoadBalancerAlgorithmTypeCode = "SIPHS"
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode: "HTTP",
					LoadBalancerPort: 443,
					ServerPort:       443,
					CertificateName:  "aaa",
				},
			}

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.ChangeLoadBalancerInstanceConfiguration(reqParams)

			Expect(err.Error()).To(ContainSubstring("loadBalancerRuleList.1.l7HealthCheckPath field is required"))
		})

		It("should be error : 'expected length of CertificateName to be in the range (1 - 300)'", func() {
			reqParams := new(RequestChangeLoadBalancerInstanceConfiguration)
			reqParams.LoadBalancerInstanceNo = "811670"
			reqParams.LoadBalancerAlgorithmTypeCode = "SIPHS"
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
			_, err := conn.ChangeLoadBalancerInstanceConfiguration(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of loadBalancerRuleList.1.certificateName to be in the range (1 - 300), got "))
		})

		It("should be error : 'loadBalancerRuleList.1.certificateName field is required'", func() {
			reqParams := new(RequestChangeLoadBalancerInstanceConfiguration)
			reqParams.LoadBalancerInstanceNo = "811670"
			reqParams.LoadBalancerAlgorithmTypeCode = "SIPHS"
			reqParams.LoadBalancerRuleList = []RequestLoadBalancerRule{
				{
					ProtocolTypeCode:  "HTTPS",
					LoadBalancerPort:  443,
					ServerPort:        443,
					L7HealthCheckPath: "/monitor",
				},
			}

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.ChangeLoadBalancerInstanceConfiguration(reqParams)

			Expect(err.Error()).To(ContainSubstring("loadBalancerRuleList.1.certificateName field is required"))
		})

		It("should be error : 'loadBalancerRuleList.1.proxyProtocolUseYn should be Y or N'", func() {
			reqParams := new(RequestChangeLoadBalancerInstanceConfiguration)
			reqParams.LoadBalancerInstanceNo = "811670"
			reqParams.LoadBalancerAlgorithmTypeCode = "SIPHS"
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
			_, err := conn.ChangeLoadBalancerInstanceConfiguration(reqParams)

			Expect(err.Error()).To(ContainSubstring("loadBalancerRuleList.1.proxyProtocolUseYn should be Y or N"))
		})
	})
})
