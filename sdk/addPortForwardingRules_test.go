package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Add Port Forwarding Rules", func() {
	Describe("Add Port Forwarding Rules", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusOK).BodyString(`<addPortForwardingRulesResponse>
    <requestId>943e41d1-f2b9-43a9-b308-5f12152a2f7f</requestId>
    <returnCode>0</returnCode>
    <returnMessage>success</returnMessage>
    <portForwardingConfigurationNo>1676</portForwardingConfigurationNo>
    <portForwardingPublicIp>192.168.120.28</portForwardingPublicIp>
    <totalRows>2</totalRows>
    <portForwardingRuleList>
        <portForwardingRule>
            <portForwardingExternalPort>1025</portForwardingExternalPort>
            <portForwardingInternalPort>3389</portForwardingInternalPort>
            <serverInstance>
                <serverInstanceNo>274079</serverInstanceNo>
                <serverName>pjaser-2</serverName>
                <serverDescription />
                <cpuCount>1</cpuCount>
                <memorySize>2147483648</memorySize>
                <baseBlockStorageSize>53687091200</baseBlockStorageSize>
                <platformType>
                    <code>WND32</code>
                    <codeName>Windows 32 Bit</codeName>
                </platformType>
                <loginKeyName>pja-1126</loginKeyName>
                <isFeeChargingMonitoring>false</isFeeChargingMonitoring>
                <publicIp />
                <privateIp>10.101.6.101</privateIp>
                <serverImageName>win-2008-32-R1</serverImageName>
                <serverInstanceStatus>
                    <code>FSTOP</code>
                    <codeName>Server failure stopped state</codeName>
                </serverInstanceStatus>
                <serverInstanceOperation>
                    <code>NULL</code>
                    <codeName>Server NULL OP</codeName>
                </serverInstanceOperation>
                <serverInstanceStatusName>repairing</serverInstanceStatusName>
                <createDate>2015-03-04T15:12:47+0900</createDate>
                <uptime>2015-03-27T11:36:29+0900</uptime>
                <serverImageProductCode>SPSW0WINNT000013</serverImageProductCode>
                <serverProductCode>SPSVRSTAND000003</serverProductCode>
                <isProtectServerTermination>false</isProtectServerTermination>
                <portForwardingPublicIp>192.168.120.28</portForwardingPublicIp>
                <portForwardingExternalPort>1025</portForwardingExternalPort>
                <portForwardingInternalPort>3389</portForwardingInternalPort>
                <zone>
                    <zoneNo>2</zoneNo>
                    <zoneName>nang_zone</zoneName>
                    <zoneDescription>nang zone</zoneDescription>
                </zone>
                <baseBlockStorageDiskType>
                    <code>NET</code>
                    <codeName>Network Storage</codeName>
                </baseBlockStorageDiskType>
                <userData />
                <accessControlGroupList>
                    <accessControlGroup>
                        <accessControlGroupConfigurationNo>3321</accessControlGroupConfigurationNo>
                        <accessControlGroupName>pja-acg</accessControlGroupName>
                        <accessControlGroupDescription />
                        <isDefault>false</isDefault>
                        <createDate>2015-03-04T15:12:21+0900</createDate>
                    </accessControlGroup>
                </accessControlGroupList>
            </serverInstance>
        </portForwardingRule>
        <portForwardingRule>
            <portForwardingExternalPort>1026</portForwardingExternalPort>
            <portForwardingInternalPort>22</portForwardingInternalPort>
            <serverInstance>
                <serverInstanceNo>274317</serverInstanceNo>
                <serverName>x0312b</serverName>
                <serverDescription />
                <cpuCount>1</cpuCount>
                <memorySize>2147483648</memorySize>
                <baseBlockStorageSize>53687091200</baseBlockStorageSize>
                <platformType>
                    <code>LNX64</code>
                    <codeName>Linux 64 Bit</codeName>
                </platformType>
                <loginKeyName>pja-1126</loginKeyName>
                <isFeeChargingMonitoring>false</isFeeChargingMonitoring>
                <publicIp />
                <privateIp>10.101.6.115</privateIp>
                <serverImageName>centos-6.3-64</serverImageName>
                <serverInstanceStatus>
                    <code>FSTOP</code>
                    <codeName>Server failure stopped state</codeName>
                </serverInstanceStatus>
                <serverInstanceOperation>
                    <code>NULL</code>
                    <codeName>Server NULL OP</codeName>
                </serverInstanceOperation>
                <serverInstanceStatusName>repairing</serverInstanceStatusName>
                <createDate>2015-03-12T18:29:39+0900</createDate>
                <uptime>2015-03-13T10:08:45+0900</uptime>
                <serverImageProductCode>SPSW0LINUX000031</serverImageProductCode>
                <serverProductCode>SPSVRSTAND000003</serverProductCode>
                <isProtectServerTermination>false</isProtectServerTermination>
                <portForwardingPublicIp>192.168.120.28</portForwardingPublicIp>
                <portForwardingExternalPort>1026</portForwardingExternalPort>
                <portForwardingInternalPort>22</portForwardingInternalPort>
                <zone>
                    <zoneNo>2</zoneNo>
                    <zoneName>nang_zone</zoneName>
                    <zoneDescription>nang zone</zoneDescription>
                </zone>
                <baseBlockStorageDiskType>
                    <code>NET</code>
                    <codeName>Network Storage</codeName>
                </baseBlockStorageDiskType>
                <userData />
                <accessControlGroupList>
                    <accessControlGroup>
                        <accessControlGroupConfigurationNo>1038</accessControlGroupConfigurationNo>
                        <accessControlGroupName>ncloud-default-acg</accessControlGroupName>
                        <accessControlGroupDescription>Default AccessControlGroup</accessControlGroupDescription>
                        <isDefault>true</isDefault>
                        <createDate>2013-12-03T10:37:39+0900</createDate>
                    </accessControlGroup>
                </accessControlGroupList>
            </serverInstance>
        </portForwardingRule>
    </portForwardingRuleList>
				</addPortForwardingRulesResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should add Port Forwarding Rules", func() {
			reqParams := &RequestAddPortForwardingRules{
				PortForwardingConfigurationNo: "1676",
				PortForwardingRuleList: []PortForwardingRule{
					{
						ServerInstanceNo:           "274079",
						PortForwardingExternalPort: "1025",
						PortForwardingInternalPort: "3389",
					},
				},
			}

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.AddPortForwardingRules(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(2))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.PortForwardingRuleList)).To(Equal(2))

			portForwardingRule := result.PortForwardingRuleList[0]
			Expect(portForwardingRule.PortForwardingExternalPort).To(Equal("1025"))
			Expect(portForwardingRule.PortForwardingInternalPort).To(Equal("3389"))
		})
	})

	Describe("Invalid internal port has been configured", func() {
		Context("Invalid internal port has been configured", func() {
			BeforeEach(func() {
				gock.New("https://ncloud.apigw.ntruss.com").
					Post("/server").
					Reply(http.StatusBadRequest).BodyString(fmt.Sprintf(`<responseError>
						<returnCode>24069</returnCode>
						<returnMessage>Invalid internal port has been configured. Usable internal port(LINUX type : 22, WINDOWS type : 3389)</returnMessage>
					</responseError>`))
			})
			AfterEach(func() {
				gock.Off()
			})

			It("should fail", func() {
				reqParams := &RequestAddPortForwardingRules{
					PortForwardingConfigurationNo: "1676",
					PortForwardingRuleList: []PortForwardingRule{
						{
							ServerInstanceNo:           "274079",
							PortForwardingExternalPort: "20",
							PortForwardingInternalPort: "3389",
						},
					},
				}
				conn := NewConnection(accessKey, secretKey)
				result, err := conn.AddPortForwardingRules(reqParams)

				Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
				Expect(result.ReturnCode).To(Equal(24069))
				Expect(result.ReturnMessage).To(Equal("Invalid internal port has been configured. Usable internal port(LINUX type : 22, WINDOWS type : 3389)"))
			})
		})
	})

	Describe("A single external port number is used in multiple rules", func() {
		Context("A single external port number is used in multiple rules", func() {
			externalPort := 2022

			BeforeEach(func() {
				gock.New("https://ncloud.apigw.ntruss.com").
					Post("/server").
					Reply(http.StatusBadRequest).BodyString(fmt.Sprintf(`<responseError>
						<returnCode>24070</returnCode>
						<returnMessage>A single external port number is used in multiple rules. External port number : %d</returnMessage>
					</responseError>`, externalPort))
			})
			AfterEach(func() {
				gock.Off()
			})

			It("should fail", func() {
				reqParams := &RequestAddPortForwardingRules{
					PortForwardingConfigurationNo: "1676",
					PortForwardingRuleList: []PortForwardingRule{
						{
							ServerInstanceNo:           "274079",
							PortForwardingExternalPort: string(externalPort),
							PortForwardingInternalPort: "3389",
						},
					},
				}
				conn := NewConnection(accessKey, secretKey)
				result, err := conn.AddPortForwardingRules(reqParams)

				Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
				Expect(result.ReturnCode).To(Equal(24070))
				Expect(result.ReturnMessage).To(Equal(fmt.Sprintf(`A single external port number is used in multiple rules. External port number : %d`, externalPort)))
			})
		})
	})

	Describe("Single server is existing in multiple rules", func() {
		Context("Single server is existing in multiple rules", func() {
			BeforeEach(func() {
				gock.New("https://ncloud.apigw.ntruss.com").
					Post("/server").
					Reply(http.StatusBadRequest).BodyString(fmt.Sprintf(`<responseError>
						<returnCode>24071</returnCode>
						<returnMessage>Single server is existing in multiple rules. Server IP address : 192.168.120.28</returnMessage>
					</responseError>`))
			})
			AfterEach(func() {
				gock.Off()
			})

			It("should fail", func() {
				reqParams := &RequestAddPortForwardingRules{
					PortForwardingConfigurationNo: "1676",
					PortForwardingRuleList: []PortForwardingRule{
						{
							ServerInstanceNo:           "274079",
							PortForwardingExternalPort: "22",
							PortForwardingInternalPort: "3389",
						},
					},
				}
				conn := NewConnection(accessKey, secretKey)
				result, err := conn.AddPortForwardingRules(reqParams)

				Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
				Expect(result.ReturnCode).To(Equal(24071))
				Expect(result.ReturnMessage).To(ContainSubstring("Single server is existing in multiple rules"))
			})
		})
	})

	Describe("External internal port as well as server instance number are not designated to port forward rule", func() {
		Context("External internal port as well as server instance number are not designated to port forward rule", func() {
			BeforeEach(func() {
				gock.New("https://ncloud.apigw.ntruss.com").
					Post("/server").
					Reply(http.StatusBadRequest).BodyString(fmt.Sprintf(`<responseError>
						<returnCode>24074</returnCode>
						<returnMessage>External internal port as well as server instance number are not designated to port forward rule.</returnMessage>
					</responseError>`))
			})
			AfterEach(func() {
				gock.Off()
			})

			It("should fail", func() {
				reqParams := &RequestAddPortForwardingRules{
					PortForwardingConfigurationNo: "1676",
					PortForwardingRuleList: []PortForwardingRule{
						{
							ServerInstanceNo:           "274079",
							PortForwardingExternalPort: "",
							PortForwardingInternalPort: "",
						},
					},
				}
				conn := NewConnection(accessKey, secretKey)
				result, err := conn.AddPortForwardingRules(reqParams)

				Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
				Expect(result.ReturnCode).To(Equal(24074))
				Expect(result.ReturnMessage).To(ContainSubstring("External internal port as well as server instance number are not designated to port forward rule"))
			})
		})
	})

	Describe("Check Arguments", func() {
		It("should be error : 'PortForwardingConfigurationNo is required'", func() {
			reqParams := new(RequestAddPortForwardingRules)
			conn := NewConnection(accessKey, secretKey)
			_, err := conn.AddPortForwardingRules(reqParams)

			Expect(err.Error()).To(Equal("PortForwardingConfigurationNo field is required"))
		})

		It("should be error : 'PortForwardingRuleList is required'", func() {
			reqParams := new(RequestAddPortForwardingRules)
			reqParams.PortForwardingConfigurationNo = "1"
			conn := NewConnection(accessKey, secretKey)
			_, err := conn.AddPortForwardingRules(reqParams)

			Expect(err.Error()).To(Equal("PortForwardingRuleList is required"))
		})
	})
})
