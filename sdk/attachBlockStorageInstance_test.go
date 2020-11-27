package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Attach Block Storage Instance", func() {
	Describe("Attach Block Storage Instance", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusOK).BodyString(`
					<attachBlockStorageInstanceResponse>
					<requestId>74fdc12b-9dd3-4855-a335-13dcf57a1bf4</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<blockStorageInstanceList>
						<blockStorageInstance>
							<blockStorageInstanceNo>825649</blockStorageInstanceNo>
							<serverInstanceNo></serverInstanceNo>
							<serverName></serverName>
							<blockStorageType>
								<code>SVRBS</code>
								<codeName>Server BS</codeName>
							</blockStorageType>
							<blockStorageName>bgss</blockStorageName>
							<blockStorageSize>10737418240</blockStorageSize>
							<deviceName></deviceName>
							<blockStorageProductCode>SPBSTBSTAD000006</blockStorageProductCode>
							<blockStorageInstanceStatus>
								<code>CREAT</code>
								<codeName>Block storage CREATED state</codeName>
							</blockStorageInstanceStatus>
							<blockStorageInstanceOperation>
								<code>ATTAC</code>
								<codeName>Block Storage  STOP OP</codeName>
							</blockStorageInstanceOperation>
							<blockStorageInstanceStatusName>attaching</blockStorageInstanceStatusName>
							<createDate>2018-06-18T19:52:57+0900</createDate>
							<blockStorageInstanceDescription></blockStorageInstanceDescription>
							<diskType>
								<code>NET</code>
								<codeName>Network Storage</codeName>
							</diskType>
							<diskDetailType>
								<code>SSD</code>
								<codeName>SSD</codeName>
							</diskDetailType>
							<maxIopsThroughput>4000</maxIopsThroughput>
							<zone>
								<zoneNo>3</zoneNo>
								<zoneName>KR-2</zoneName>
								<zoneCode>KR-2</zoneCode>
								<zoneDescription>평촌 zone</zoneDescription>
								<regionNo>1</regionNo>
							</zone>
						</blockStorageInstance>
					</blockStorageInstanceList>
				</attachBlockStorageInstanceResponse>
					`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should Attach Block Storage Instance", func() {
			conn := NewConnection(accessKey, secretKey)
			reqParams := &RequestAttachBlockStorageInstance{
				ServerInstanceNo:       "805847",
				BlockStorageInstanceNo: "825649",
			}
			result, err := conn.AttachBlockStorageInstance(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.BlockStorageInstance)).To(Equal(1))

			b := result.BlockStorageInstance[0]
			Expect(b.BlockStorageInstanceNo).To(Equal(reqParams.BlockStorageInstanceNo))
		})
	})

	Describe("Invalid ServerInstanceNo", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusInternalServerError).BodyString(`
					<responseError>
					<returnCode>1300</returnCode>
					<returnMessage>Please try your call again later.
				Temporarily out of service.
				If error continue, Please contact our customer service center.</returnMessage>
				</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should not  Attach Block Storage Instance by invalid ServerInstanceNo", func() {
			conn := NewConnection(accessKey, secretKey)
			reqParams := &RequestAttachBlockStorageInstance{
				ServerInstanceNo:       "805000",
				BlockStorageInstanceNo: "825649",
			}
			result, err := conn.AttachBlockStorageInstance(reqParams)

			Expect(result.ReturnCode).To(Equal(1300))
			Expect(result.ReturnMessage).To(ContainSubstring("Please try your call again later."))
			Expect(err.Error()).To(ContainSubstring("500 Internal Server Error"))
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
			reqParams := &RequestAttachBlockStorageInstance{
				ServerInstanceNo:       "805000",
				BlockStorageInstanceNo: "825649",
			}
			result, err := conn.AttachBlockStorageInstance(reqParams)

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
			reqParams := &RequestAttachBlockStorageInstance{
				ServerInstanceNo:       "805000",
				BlockStorageInstanceNo: "825649",
			}
			result, err := conn.AttachBlockStorageInstance(reqParams)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Expired url."))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})
})
