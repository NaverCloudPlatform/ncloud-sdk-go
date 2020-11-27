package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Delete Port Forwarding Rules", func() {
	Describe("Delete Port Forwarding Rules", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusOK).BodyString(`<deletePortForwardingRulesResponse>
    <requestId>54b32354-3acb-489c-8bdb-09a722555a75</requestId>
    <returnCode>0</returnCode>
    <returnMessage>success</returnMessage>
    <portForwardingConfigurationNo>1676</portForwardingConfigurationNo>
    <portForwardingPublicIp>192.168.120.28</portForwardingPublicIp>
    <totalRows>1</totalRows>
    <portForwardingRuleList>
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
</deletePortForwardingRulesResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should delete Port Forwarding Rules", func() {
			reqParams := &RequestDeletePortForwardingRules{
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
			result, err := conn.DeletePortForwardingRules(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.PortForwardingRuleList)).To(Equal(1))

			portForwardingRule := result.PortForwardingRuleList[0]
			Expect(portForwardingRule.PortForwardingExternalPort).To(Equal("1026"))
			Expect(portForwardingRule.PortForwardingInternalPort).To(Equal("22"))
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
