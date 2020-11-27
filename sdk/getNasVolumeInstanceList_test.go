package sdk_test

import (
	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get NAS Volume Instance List", func() {
	Describe("Get NAS Volume Instance List", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(200).BodyString(`<getNasVolumeInstanceListResponse>
    <requestId>8d8434da-f69f-4675-837c-b570e4567795</requestId>
    <returnCode>0</returnCode>
    <returnMessage>success</returnMessage>
    <totalRows>1</totalRows>
    <nasVolumeInstanceList>
        <nasVolumeInstance>
            <nasVolumeInstanceNo>327798</nasVolumeInstanceNo>
			<nasVolumeInstanceStatus>
				<code>CREAT</code>
				<codeName>NAS create</codeName>
			</nasVolumeInstanceStatus>​
            <createDate>Mon Jun 12 22:16:56 KST 2017</createDate>
            <nasVolumeInstanceDescription />
            <serviceIP>10.99.67.199</serviceIP>
			<volumeAllotmentProtocolType>
				<code>NFS</code>
				<codeName>NFS</codeName>
			</volumeAllotmentProtocolType>
            <volumeName>n000203_nangtest01</volumeName>
            <volumeTotalSize>53687091200</volumeTotalSize>
            <volumeSize>53687091200</volumeSize>
            <volumeUseSize>344064</volumeUseSize>
            <volumeUseRatio>0.0</volumeUseRatio>
            <snapshotVolumeSize>0</snapshotVolumeSize>
            <snapshotVolumeUseSize>0</snapshotVolumeUseSize>
            <snapshotVolumeUseRatio>0.0</snapshotVolumeUseRatio>
            <isSnapshot>false</isSnapshot>
            <isEvent>true</isEvent>
            <nasVolumeInstanceCustomIpList>
                <nasVolumeInstanceCustomIp>
                    <customIp>10.113.177.15</customIp>
                </nasVolumeInstanceCustomIp>
            </nasVolumeInstanceCustomIpList>
            <nasVolumeServerInstanceList>
                <nasVolumeServerInstance>
                    <serverInstanceNo>321332</serverInstanceNo>
                    <serverName>svr-9ae8ae0ae0f8ad2</serverName>
                    <ip>10.113.177.14</ip>
                    <instanceStatus>NSTOP</instanceStatus>
                    <instanceStatusName>terminating</instanceStatusName>
                </nasVolumeServerInstance>
            </nasVolumeServerInstanceList>
        </nasVolumeInstance>
    </nasVolumeInstanceList>
</getNasVolumeInstanceListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get NAS Volume Instance List", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetNasVolumeInstanceList(nil)

			Expect(err).To(BeNil())
			Expect(result.RequestID).To(Equal("8d8434da-f69f-4675-837c-b570e4567795"))
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.NasVolumeInstanceList)).To(Equal(1))

			nasVolumeInstance := result.NasVolumeInstanceList[0]
			Expect(nasVolumeInstance.NasVolumeInstanceNo).To(Equal("327798"))
			Expect(nasVolumeInstance.NasVolumeInstanceStatus.Code).To(Equal("CREAT"))
			Expect(nasVolumeInstance.VolumeAllotmentProtocolType.Code).To(Equal("NFS"))
			Expect(nasVolumeInstance.VolumeName).To(Equal("n000203_nangtest01"))

			customIPList := nasVolumeInstance.NasVolumeInstanceCustomIPList[0]
			Expect(customIPList.CustomIP).To(Equal("10.113.177.15"))
		})
	})

	Describe("There is no NAS Volume Instance List", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(200).BodyString(`<getNasVolumeInstanceListResponse>
    	<requestId>8d8434da-f69f-4675-837c-b570e4567795</requestId>
    	<returnCode>0</returnCode>
    	<returnMessage>success</returnMessage>
    	<totalRows>0</totalRows>
    	<nasVolumeInstanceList/>
	</getNasVolumeInstanceListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get No NAS Volume Instance List", func() {
			reqParams := &RequestGetNasVolumeInstanceList{
				VolumeAllotmentProtocolTypeCode: "NFS",
				IsEventConfiguration:            "false",
				IsSnapshotConfiguration:         "false",
			}

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetNasVolumeInstanceList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.RequestID).To(Equal("8d8434da-f69f-4675-837c-b570e4567795"))
			Expect(result.TotalRows).To(Equal(0))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.NasVolumeInstanceList)).To(Equal(0))
		})
	})
})
