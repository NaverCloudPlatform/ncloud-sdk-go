package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	"gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Change NAS Volume Size", func() {
	Describe("Change NAS Volume Size", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusOK).BodyString(`
					<changeNasVolumeSizeResponse>
    <requestId>4f7f3c48-cc33-4013-9983-26f3c04bff70</requestId>
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
</changeNasVolumeSizeResponse>
					`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should Change NAS Volume Size", func() {
			conn := NewConnection(accessKey, secretKey)
			reqParams := new(RequestChangeNasVolumeSize)
			reqParams.NasVolumeInstanceNo = "397767"
			reqParams.VolumeSize = 500
			result, err := conn.ChangeNasVolumeSize(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.NasVolumeInstanceList)).To(Equal(1))

			nasVolumeInstance := result.NasVolumeInstanceList[0]
			Expect(nasVolumeInstance.NasVolumeInstanceNo).To(Equal("397767"))
			Expect(nasVolumeInstance.NasVolumeInstanceStatus.Code).To(Equal("CREAT"))
			Expect(nasVolumeInstance.NasVolumeInstanceOperation.Code).To(Equal("NULL"))
			Expect(nasVolumeInstance.NasVolumeInstanceStatusName).To(Equal("created"))
			Expect(nasVolumeInstance.CreateDate).To(Equal("2018-02-27T13:13:05+0900"))
			Expect(nasVolumeInstance.MountInformation).To(Equal("10.105.84.82:/n000212_penguin"))
			Expect(nasVolumeInstance.VolumeAllotmentProtocolType.Code).To(Equal("NFS"))
			Expect(nasVolumeInstance.VolumeName).To(Equal("n000212_penguin"))
			Expect(nasVolumeInstance.VolumeTotalSize).To(Equal(547608330240))
			Expect(nasVolumeInstance.VolumeSize).To(Equal(547608330240))
			Expect(nasVolumeInstance.VolumeUseSize).To(Equal(258048))
			Expect(nasVolumeInstance.VolumeUseRatio).To(Equal(float32(0.0)))
			Expect(nasVolumeInstance.SnapshotVolumeConfigurationRatio).To(Equal(float32(0.0)))
			Expect(nasVolumeInstance.SnapshotVolumeSize).To(Equal(0))
			Expect(nasVolumeInstance.SnapshotVolumeUseSize).To(Equal(0))
			Expect(nasVolumeInstance.SnapshotVolumeUseRatio).To(Equal(float32(0.0)))
			Expect(nasVolumeInstance.IsSnapshotConfiguration).To(Equal(false))
			Expect(nasVolumeInstance.IsEventConfiguration).To(Equal(false))
			Expect(nasVolumeInstance.Zone.ZoneCode).To(Equal("KR-1"))
			Expect(nasVolumeInstance.Region.RegionCode).To(Equal("KR"))
		})
	})

	Describe("Authorize fail", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
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
			reqParams := new(RequestChangeNasVolumeSize)
			reqParams.NasVolumeInstanceNo = "397767"
			reqParams.VolumeSize = 500
			result, err := conn.ChangeNasVolumeSize(reqParams)
			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Invalid consumerKey"))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})

	Describe("Expired URL", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
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
			reqParams := new(RequestChangeNasVolumeSize)
			reqParams.NasVolumeInstanceNo = "397767"
			reqParams.VolumeSize = 500
			result, err := conn.ChangeNasVolumeSize(reqParams)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Expired url."))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})

	Describe("Check Arguments", func() {
		It("should be error : 'NasVolumeInstanceNo field is required'", func() {
			conn := NewConnection(accessKey, secretKey)
			_, err := conn.ChangeNasVolumeSize(nil)

			Expect(err.Error()).To(Equal("NasVolumeInstanceNo field is required"))
		})

		It("should be error : 'NasVolumeInstanceNo field is required'", func() {
			reqParams := new(RequestChangeNasVolumeSize)
			conn := NewConnection(accessKey, secretKey)
			_, err := conn.ChangeNasVolumeSize(reqParams)

			Expect(err.Error()).To(Equal("NasVolumeInstanceNo field is required"))
		})

		It("should be error : 'VolumeSize cannot be lower than 500'", func() {
			reqParams := new(RequestChangeNasVolumeSize)
			reqParams.NasVolumeInstanceNo = "397767"
			conn := NewConnection(accessKey, secretKey)
			_, err := conn.ChangeNasVolumeSize(reqParams)

			Expect(err.Error()).To(Equal("\"VolumeSize\" cannot be lower than 500: 0"))
		})

		It("should be error : 'VolumeSize must be a multiple of 100'", func() {
			reqParams := new(RequestChangeNasVolumeSize)
			reqParams.NasVolumeInstanceNo = "397767"
			reqParams.VolumeSize = 1024
			conn := NewConnection(accessKey, secretKey)
			_, err := conn.ChangeNasVolumeSize(reqParams)

			Expect(err.Error()).To(Equal("VolumeSize must be a multiple of 100"))
		})
	})
})
