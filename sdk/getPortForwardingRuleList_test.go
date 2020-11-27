package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Port Forwarding List", func() {
	Describe("Get Port Forwarding List", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getPortForwardingRuleListResponse>
    <requestId>08842664-ea53-49b5-9192-7e4d165b86d6</requestId>
    <returnCode>0</returnCode>
    <returnMessage>success</returnMessage>
    <portForwardingConfigurationNo>1676</portForwardingConfigurationNo>
    <portForwardingPublicIp>192.168.120.28</portForwardingPublicIp>
    <totalRows>2</totalRows>
	<zone>
		<zoneNo>2</zoneNo>
		<zoneName>KR-1</zoneName>
		<zoneCode>KR-1</zoneCode>
		<zoneDescription>가산 zone</zoneDescription>
		<regionNo>1</regionNo>
	</zone>
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
</getPortForwardingRuleListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get Port Forwarding List", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetPortForwardingRuleList(nil)

			Expect(err).To(BeNil())
			Expect(result.RequestID).To(Equal("08842664-ea53-49b5-9192-7e4d165b86d6"))
			Expect(result.TotalRows).To(Equal(2))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.PortForwardingRuleList)).To(Equal(2))

			Expect(result.PortForwardingConfigurationNo).To(Equal(1676))
			Expect(result.PortForwardingPublicIp).To(Equal("192.168.120.28"))
			Expect(result.Zone.ZoneCode).To(Equal("KR-1"))

			portForwardingRule := result.PortForwardingRuleList[0]
			Expect(portForwardingRule.PortForwardingExternalPort).To(Equal("1025"))
			Expect(portForwardingRule.PortForwardingInternalPort).To(Equal("3389"))
			Expect(portForwardingRule.ServerInstanceNo).To(Equal("274079"))
			Expect(portForwardingRule.PortForwardingPublicIp).To(Equal("192.168.120.28"))
		})
	})

	Describe("There is no Port Forwarding List", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getPortForwardingRuleListResponse>
				<requestId>08842664-ea53-49b5-9192-7e4d165b86d6</requestId>
				<returnCode>0</returnCode>
			    <returnMessage>success</returnMessage>
    			<portForwardingConfigurationNo>1676</portForwardingConfigurationNo>
    			<portForwardingPublicIp>192.168.120.28</portForwardingPublicIp>
    			<totalRows>0</totalRows>
				<portForwardingRuleList/>
				</getPortForwardingRuleListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get No Port Forwarding List", func() {
			reqParams := &RequestPortForwardingRuleList{
				InternetLineTypeCode: "PUBLC",
			}

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetPortForwardingRuleList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.RequestID).To(Equal("08842664-ea53-49b5-9192-7e4d165b86d6"))
			Expect(result.TotalRows).To(Equal(0))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.PortForwardingRuleList)).To(Equal(0))
		})
	})
})
