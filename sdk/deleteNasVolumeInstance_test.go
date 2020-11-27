package sdk_test

import (
	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Delete NAS Volume Instance", func() {
	Describe("Delete NAS Volume Instance", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(200).BodyString(`<deleteNasVolumeInstanceResponse>
    <requestId>d131a03a-ecc4-407d-88d9-059911b34f74</requestId>
    <returnCode>0</returnCode>
    <returnMessage>success</returnMessage>
    <totalRows>1</totalRows>
    <nasVolumeInstanceList>
        <nasVolumeInstance>
            <nasVolumeInstanceNo>397767</nasVolumeInstanceNo>
            <nasVolumeInstanceStatus>
                <code>CREAT</code>
                <codeName>NAS create</codeName>
            </nasVolumeInstanceStatus>
            <nasVolumeInstanceOperation>
                <code>NULL</code>
                <codeName>NAS NULL OP</codeName>
            </nasVolumeInstanceOperation>
            <nasVolumeInstanceStatusName>created</nasVolumeInstanceStatusName>
            <createDate>2018-02-27T13:13:05+0900</createDate>
            <nasVolumeInstanceDescription />
            <mountInformation>10.105.84.82:/n000212_penguin</mountInformation>
            <volumeAllotmentProtocolType>
                <code>NFS</code>
                <codeName>NFS</codeName>
            </volumeAllotmentProtocolType>
            <volumeName>n000212_penguin</volumeName>
            <volumeTotalSize>547608330240</volumeTotalSize>
            <volumeSize>547608330240</volumeSize>
            <volumeUseSize>258048</volumeUseSize>
            <volumeUseRatio>0.0</volumeUseRatio>
            <snapshotVolumeConfigurationRatio>0.0</snapshotVolumeConfigurationRatio>
            <snapshotVolumeSize>0</snapshotVolumeSize>
            <snapshotVolumeUseSize>0</snapshotVolumeUseSize>
            <snapshotVolumeUseRatio>0.0</snapshotVolumeUseRatio>
            <isSnapshotConfiguration>false</isSnapshotConfiguration>
            <isEventConfiguration>false</isEventConfiguration>
            <region>
                <regionNo>1</regionNo>
                <regionCode>KR</regionCode>
                <regionName>Korea</regionName>
            </region>
            <zone>
                <zoneNo>2</zoneNo>
                <zoneName>KR-1</zoneName>
                <zoneCode>KR-1</zoneCode>
                <zoneDescription>KR-1 zone</zoneDescription>
                <regionNo>1</regionNo>
            </zone>
            <nasVolumeInstanceCustomIpList />
            <nasVolumeServerInstanceList />
        </nasVolumeInstance>
    </nasVolumeInstanceList>
</deleteNasVolumeInstanceResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should delete NAS Volume Instance", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.DeleteNasVolumeInstance("397766")

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.NasVolumeInstanceList)).To(Equal(1))

			nasVolumeInstance := result.NasVolumeInstanceList[0]
			Expect(nasVolumeInstance.VolumeName).To(Equal("n000212_penguin"))
			Expect(nasVolumeInstance.VolumeTotalSize).To(Equal(547608330240))
		})
	})

})
