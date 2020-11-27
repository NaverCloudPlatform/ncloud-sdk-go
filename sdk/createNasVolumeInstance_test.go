package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create NAS Volume Instance", func() {
	Describe("Create NAS Volume Instance", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusOK).BodyString(`<createNasVolumeInstanceResponse>
    <requestId>58574752-49cb-411f-8710-e0fc6c28b8a7</requestId>
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
            <volumeTotalSize>536870912000</volumeTotalSize>
            <volumeSize>536870912000</volumeSize>
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
</createNasVolumeInstanceResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should create NAS Volume", func() {
			reqParams := &RequestCreateNasVolumeInstance{
				VolumeName:                      "n000212_penguin",
				VolumeSize:                      600,
				VolumeAllotmentProtocolTypeCode: "NFS",
			}

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.CreateNasVolumeInstance(reqParams)

			Expect(err).To(BeNil())

			Expect(result.RequestID).To(Equal("58574752-49cb-411f-8710-e0fc6c28b8a7"))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.NasVolumeInstanceList)).To(Equal(1))

			nasVolume := result.NasVolumeInstanceList[0]
			Expect(nasVolume.NasVolumeInstanceNo).To(Equal("397767"))
			Expect(nasVolume.NasVolumeInstanceStatus.Code).To(Equal("CREAT"))
			Expect(nasVolume.VolumeAllotmentProtocolType.Code).To(Equal("NFS"))
			Expect(nasVolume.VolumeName).To(Equal("n000212_penguin"))
		})
	})

	Describe("The server information for setting up NFS is invalid", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
					<returnCode>24138</returnCode>
					<returnMessage>
					The server information for setting up NFS is invalid.
					</returnMessage>
					</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should fail", func() {
			reqParams := &RequestCreateNasVolumeInstance{
				VolumeName:                      "n000212_penguin",
				VolumeSize:                      600,
				VolumeAllotmentProtocolTypeCode: "NFS",
			}
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.CreateNasVolumeInstance(reqParams)

			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
			Expect(result.ReturnCode).To(Equal(24138))
			Expect(result.ReturnMessage).To(Equal("The server information for setting up NFS is invalid."))
		})
	})

})
