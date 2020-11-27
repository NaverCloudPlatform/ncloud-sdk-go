package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create Server Image", func() {
	Describe("Create Server Image", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusOK).BodyString(`<createMemberServerImageResponse>
					<requestId>148e5d84-1e01-411b-812f-f092c959da3a</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<memberServerImageList>
						<memberServerImage>
							<memberServerImageNo>4328</memberServerImageNo>
							<memberServerImageName>imageBySDK</memberServerImageName>
							<memberServerImageDescription>server image by SDK go</memberServerImageDescription>
							<originalServerInstanceNo>551236</originalServerInstanceNo>
							<originalServerProductCode>SPSVRSSD00000003</originalServerProductCode>
							<originalServerName>fffffee004</originalServerName>
							<originalBaseBlockStorageDiskType>
								<code>NET</code>
								<codeName>Network Storage</codeName>
							</originalBaseBlockStorageDiskType>
							<originalServerImageProductCode>SPSW0LINUX000046</originalServerImageProductCode>
							<originalOsInformation>CentOS 7.3 (64-bit)</originalOsInformation>
							<originalServerImageName>centos-7.3-64</originalServerImageName>
							<memberServerImageStatusName>creating</memberServerImageStatusName>
							<memberServerImageStatus>
								<code>INIT</code>
								<codeName>NSI INIT state</codeName>
							</memberServerImageStatus>
							<memberServerImageOperation>
								<code>CREAT</code>
								<codeName>NSI CREAT OP</codeName>
							</memberServerImageOperation>
							<memberServerImagePlatformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</memberServerImagePlatformType>
							<createDate>2017-12-01T22:03:03+0900</createDate>
							<region>
								<regionNo>1</regionNo>
								<regionCode>KR</regionCode>
								<regionName>Korea</regionName>
							</region>
							<zone>
								<zoneNo>2</zoneNo>
								<zoneName>KR-1</zoneName>
								<zoneDescription>가산 NANG zone</zoneDescription>
							</zone>
							<memberServerImageBlockStorageTotalRows>0</memberServerImageBlockStorageTotalRows>
							<memberServerImageBlockStorageTotalSize>0</memberServerImageBlockStorageTotalSize>
						</memberServerImage>
					</memberServerImageList>
				</createMemberServerImageResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should create server image", func() {
			reqParams := new(RequestCreateServerImage)
			reqParams.MemberServerImageName = "imageBySDK"
			reqParams.MemberServerImageDescription = "server image by SDK go"
			reqParams.ServerInstanceNo = "551236"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.CreateMemberServerImage(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.MemberServerImageList)).To(Equal(1))

			memberServerImage := result.MemberServerImageList[0]
			Expect(memberServerImage.MemberServerImageNo).To(Equal("4328"))
			Expect(memberServerImage.MemberServerImageName).To(Equal("imageBySDK"))
			Expect(memberServerImage.MemberServerImageDescription).To(Equal("server image by SDK go"))
			Expect(memberServerImage.OriginalServerName).To(Equal("fffffee004"))
			Expect(memberServerImage.OriginalServerProductCode).To(Equal("SPSVRSSD00000003"))
			Expect(memberServerImage.Region.RegionCode).To(Equal("KR"))
			Expect(memberServerImage.MemberServerImageBlockStorageTotalRows).To(Equal(0))
			Expect(memberServerImage.MemberServerImageStatusName).To(Equal("creating"))
		})
	})

	Describe("Unable to create member server image", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
					<returnCode>25018</returnCode>
					<returnMessage>Operation is currently unavailable since (other) user is either operating server or creating server image of the server. </returnMessage>
				</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})

		It("should fail if (other) user is either operating server or creating server image of the server.", func() {
			reqParams := new(RequestCreateServerImage)
			reqParams.MemberServerImageName = "imageBySDK23"
			reqParams.MemberServerImageDescription = "server image by SDK go"
			reqParams.ServerInstanceNo = "551236"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.CreateMemberServerImage(reqParams)

			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
			Expect(result.ReturnCode).To(Equal(25018))
			Expect(result.ReturnMessage).To(Equal("Operation is currently unavailable since (other) user is either operating server or creating server image of the server."))
		})
	})

	Describe("Unable to create member server image", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
					<returnCode>25008</returnCode>
					<returnMessage>Unable to create server image since server is either not in suspended status or in operation.</returnMessage>
				</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})

		It("should fail if server is either not in suspended status or in operation.", func() {
			reqParams := new(RequestCreateServerImage)
			reqParams.MemberServerImageName = "imageBySDK23"
			reqParams.MemberServerImageDescription = "server image by SDK go"
			reqParams.ServerInstanceNo = "551236"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.CreateMemberServerImage(reqParams)

			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
			Expect(result.ReturnCode).To(Equal(25008))
			Expect(result.ReturnMessage).To(Equal("Unable to create server image since server is either not in suspended status or in operation."))
		})
	})

	Describe("Duplicate server image name", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
					<returnCode>10300</returnCode>
					<returnMessage>Instance name is already in use. please use other name.</returnMessage>
				</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})

		It("should fail if (other) user is either operating server or creating server image of the server.", func() {
			reqParams := new(RequestCreateServerImage)
			reqParams.MemberServerImageName = "imageBySDK"
			reqParams.MemberServerImageDescription = "server image by SDK go"
			reqParams.ServerInstanceNo = "551236"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.CreateMemberServerImage(reqParams)

			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
			Expect(result.ReturnCode).To(Equal(10300))
			Expect(result.ReturnMessage).To(Equal("Instance name is already in use. please use other name."))
		})
	})
})
