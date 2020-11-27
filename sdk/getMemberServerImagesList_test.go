package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MemberServerImagesList", func() {
	Describe("Get all Member Server Image List", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getMemberServerImageListResponse>
					<requestId>ceff19a9-dd17-4c95-89c7-21b617b20628</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>2</totalRows>
					<memberServerImageList>
						<memberServerImage>
							<memberServerImageNo>2440</memberServerImageNo>
							<memberServerImageName>test</memberServerImageName>
							<memberServerImageDescription>test</memberServerImageDescription>
							<originalServerInstanceNo>326664</originalServerInstanceNo>
							<originalServerProductCode>SPSVRSTAND000004</originalServerProductCode>
							<originalServerName>ysw-test</originalServerName>
							<originalBaseBlockStorageDiskType>
								<code>NET</code>
								<codeName>Network Storage</codeName>
							</originalBaseBlockStorageDiskType>
							<originalServerImageProductCode>SPSW0LINUX000044</originalServerImageProductCode>
							<originalOsInformation>CentOS 6.6 (64-bit)</originalOsInformation>
							<originalServerImageName>centos-6.6-64</originalServerImageName>
							<memberServerImageStatusName>terminating</memberServerImageStatusName>
							<memberServerImageStatus>
								<code>CREAT</code>
								<codeName>NSI CREATED State</codeName>
							</memberServerImageStatus>
							<memberServerImageOperation>
								<code>TERMT</code>
								<codeName>NSI TERMINATE OP</codeName>
							</memberServerImageOperation>
							<memberServerImagePlatformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</memberServerImagePlatformType>
							<createDate>2017-06-05T11:42:56+0900</createDate>
							<region>
								<regionNo>1</regionNo>
								<regionCode>KR</regionCode>
								<regionName>Korea</regionName>
							</region>
							<zone>
								<zoneNo>2</zoneNo>
								<zoneName>KR-1</zoneName>
								<zoneDescription>KR-1 zone</zoneDescription>
							</zone>
							<memberServerImageBlockStorageTotalRows>2</memberServerImageBlockStorageTotalRows>
							<memberServerImageBlockStorageTotalSize>107374182400</memberServerImageBlockStorageTotalSize>
						</memberServerImage>
						<memberServerImage>
						<memberServerImageNo>2441</memberServerImageNo>
						<memberServerImageName>test111</memberServerImageName>
						<memberServerImageDescription>test</memberServerImageDescription>
						<originalServerInstanceNo>326664</originalServerInstanceNo>
						<originalServerProductCode>SPSVRSTAND000004</originalServerProductCode>
						<originalServerName>ysw-test</originalServerName>
						<originalBaseBlockStorageDiskType>
							<code>NET</code>
							<codeName>Network Storage</codeName>
						</originalBaseBlockStorageDiskType>
						<originalServerImageProductCode>SPSW0LINUX000044</originalServerImageProductCode>
						<originalOsInformation>CentOS 6.6 (64-bit)</originalOsInformation>
						<originalServerImageName>centos-6.6-64</originalServerImageName>
						<memberServerImageStatusName>terminating</memberServerImageStatusName>
						<memberServerImageStatus>
							<code>CREAT</code>
							<codeName>NSI CREATED State</codeName>
						</memberServerImageStatus>
						<memberServerImageOperation>
							<code>TERMT</code>
							<codeName>NSI TERMINATE OP</codeName>
						</memberServerImageOperation>
						<memberServerImagePlatformType>
							<code>LNX64</code>
							<codeName>Linux 64 Bit</codeName>
						</memberServerImagePlatformType>
						<createDate>2017-06-05T11:42:56+0900</createDate>
						<region>
							<regionNo>1</regionNo>
							<regionCode>KR</regionCode>
							<regionName>Korea</regionName>
						</region>
						<zone>
							<zoneNo>2</zoneNo>
							<zoneName>KR-1</zoneName>
							<zoneDescription>KR-1 zone</zoneDescription>
						</zone>
						<memberServerImageBlockStorageTotalRows>2</memberServerImageBlockStorageTotalRows>
						<memberServerImageBlockStorageTotalSize>107374182400</memberServerImageBlockStorageTotalSize>
					</memberServerImage>
					</memberServerImageList>
				</getMemberServerImageListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get member server image list", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetMemberServerImageList(nil)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(2))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.MemberServerImageList)).To(Equal(2))

			memberServerImage := result.MemberServerImageList[0]
			Expect(memberServerImage.MemberServerImageNo).To(Equal("2440"))
			Expect(memberServerImage.OriginalServerName).To(Equal("ysw-test"))
			Expect(memberServerImage.OriginalServerProductCode).To(Equal("SPSVRSTAND000004"))
			Expect(memberServerImage.Region.RegionCode).To(Equal("KR"))
			Expect(memberServerImage.MemberServerImageBlockStorageTotalRows).To(Equal(2))

			memberServerImage = result.MemberServerImageList[1]
			Expect(memberServerImage.MemberServerImageNo).To(Equal("2441"))
			Expect(memberServerImage.MemberServerImageName).To(Equal("test111"))
			Expect(memberServerImage.Zone.ZoneDescription).To(Equal("KR-1 zone"))
			Expect(memberServerImage.MemberServerImageBlockStorageTotalRows).To(Equal(2))
		})
	})

	Describe("Get One Member Server Image List which memberServerImageNoList.1 is 2440", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getMemberServerImageListResponse>
					<requestId>b65188cf-4cfb-46c4-ab39-2e41107c5e2b</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<memberServerImageList>
						<memberServerImage>
							<memberServerImageNo>2440</memberServerImageNo>
							<memberServerImageName>test</memberServerImageName>
							<memberServerImageDescription>test</memberServerImageDescription>
							<originalServerInstanceNo>326664</originalServerInstanceNo>
							<originalServerProductCode>SPSVRSTAND000004</originalServerProductCode>
							<originalServerName>ysw-test</originalServerName>
							<originalBaseBlockStorageDiskType>
								<code>NET</code>
								<codeName>Network Storage</codeName>
							</originalBaseBlockStorageDiskType>
							<originalServerImageProductCode>SPSW0LINUX000044</originalServerImageProductCode>
							<originalOsInformation>CentOS 6.6 (64-bit)</originalOsInformation>
							<originalServerImageName>centos-6.6-64</originalServerImageName>
							<memberServerImageStatusName>terminating</memberServerImageStatusName>
							<memberServerImageStatus>
								<code>CREAT</code>
								<codeName>NSI CREATED State</codeName>
							</memberServerImageStatus>
							<memberServerImageOperation>
								<code>TERMT</code>
								<codeName>NSI TERMINATE OP</codeName>
							</memberServerImageOperation>
							<memberServerImagePlatformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</memberServerImagePlatformType>
							<createDate>2017-06-05T11:42:56+0900</createDate>
							<region>
								<regionNo>1</regionNo>
								<regionCode>KR</regionCode>
								<regionName>Korea</regionName>
							</region>
							<zone>
								<zoneNo>2</zoneNo>
								<zoneName>KR-1</zoneName>
								<zoneDescription>KR-1 zone</zoneDescription>
							</zone>
							<memberServerImageBlockStorageTotalRows>2</memberServerImageBlockStorageTotalRows>
							<memberServerImageBlockStorageTotalSize>107374182400</memberServerImageBlockStorageTotalSize>
						</memberServerImage>
					</memberServerImageList>
				</getMemberServerImageListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get server image list", func() {
			reqParams := new(RequestServerImageList)
			reqParams.MemberServerImageNoList = []string{"4426"}

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetMemberServerImageList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(len(result.MemberServerImageList)).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))

			memberServerImage := result.MemberServerImageList[0]
			Expect(memberServerImage.OriginalServerProductCode).To(Equal("SPSVRSTAND000004"))
			Expect(memberServerImage.Region.RegionCode).To(Equal("KR"))
			Expect(memberServerImage.MemberServerImageBlockStorageTotalRows).To(Equal(2))
			Expect(memberServerImage.Region.RegionNo).To(Equal("1"))
			Expect(memberServerImage.Region.RegionCode).To(Equal("KR"))
			Expect(memberServerImage.Region.RegionName).To(Equal("Korea"))
			Expect(memberServerImage.Zone.ZoneNo).To(Equal("2"))
			Expect(memberServerImage.Zone.ZoneName).To(Equal("KR-1"))
			Expect(memberServerImage.Zone.ZoneDescription).To(Equal("KR-1 zone"))

		})
	})

	Describe("There is no member server image list", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getMemberServerImageListResponse>
					<requestId>97484034-36a6-4ec6-abd4-cfc987890149</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>0</totalRows>
					<memberServerImageList/>
				</getMemberServerImageListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get member server image list", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetMemberServerImageList(nil)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(0))
			Expect(len(result.MemberServerImageList)).To(Equal(0))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
		})
	})

	Describe("Authorize fail", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusUnauthorized).BodyString(`<responseError>
					<returnCode>800</returnCode>
					<returnMessage>Expired url.</returnMessage>
					</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get member server image list", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetMemberServerImageList(nil)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Expired url."))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})
})
