package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create Block Storage Snapshot Instance", func() {
	Describe("Create Block Storage Snapshot Instance", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusOK).BodyString(`<createBlockStorageSnapshotInstanceResponse>
   				<script />
   				<requestId>8ee82b50-bb6c-428e-bd1a-55af5864a1b2</requestId>
   				<returnCode>0</returnCode>
   				<returnMessage>success</returnMessage>
   				<totalRows>1</totalRows>
   				<blockStorageSnapshotInstanceList>
      				<blockStorageSnapshot>
     					<blockStorageSnapshotInstanceNo>694783</blockStorageSnapshotInstanceNo>
     					<blockStorageSnapshotName>snap1644516a9a5</blockStorageSnapshotName>
     					<blockStorageSnapshotVolumeSize>53687091200</blockStorageSnapshotVolumeSize>
     					<originalBlockStorageInstanceNo>694553</originalBlockStorageInstanceNo>
     					<originalBlockStorageName>s-5ds0yzequgdxx</originalBlockStorageName>
     					<blockStorageSnapshotInstanceStatus>
        					<code>INIT</code>
        					<codeName>Block storage INIT state</codeName>
     					</blockStorageSnapshotInstanceStatus>
     					<blockStorageSnapshotInstanceOperation>
        					<code>NULL</code>
        					<codeName>Block Storage NULLOP</codeName>
     					</blockStorageSnapshotInstanceOperation>
     					<blockStorageSnapshotInstanceStatusName>initialized</blockStorageSnapshotInstanceStatusName>
     					<createDate>2018-06-28T15:31:10+0900</createDate>
     					<blockStorageSnapshotInstanceDescription />
     					<serverImageProductCode>SPSW0LINUX000046</serverImageProductCode>
     					<osInformation>CentOS 7.3 (64-bit)</osInformation>
      				</blockStorageSnapshot>
   				</blockStorageSnapshotInstanceList>
			</createBlockStorageSnapshotInstanceResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should create Block Storage Snapshot Instance", func() {
			reqParams := new(RequestCreateBlockStorageSnapshotInstance)
			reqParams.BlockStorageInstanceNo = "694783"
			reqParams.BlockStorageSnapshotName = "snap1644516a9a5"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.CreateBlockStorageSnapshotInstance(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.BlockStorageSnapshotInstanceList)).To(Equal(1))

			blockStorageSnapshotInstance := result.BlockStorageSnapshotInstanceList[0]
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotInstanceNo).To(Equal("694783"))
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotName).To(Equal("snap1644516a9a5"))
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotVolumeSize).To(Equal(53687091200))
			Expect(blockStorageSnapshotInstance.OriginalBlockStorageInstanceNo).To(Equal("694553"))
			Expect(blockStorageSnapshotInstance.OriginalBlockStorageName).To(Equal("s-5ds0yzequgdxx"))
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotInstanceStatus.Code).To(Equal("INIT"))
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotInstanceStatus.CodeName).To(Equal("Block storage INIT state"))
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotInstanceOperation.Code).To(Equal("NULL"))
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotInstanceOperation.CodeName).To(Equal("Block Storage NULLOP"))
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotInstanceStatusName).To(Equal("initialized"))
			Expect(blockStorageSnapshotInstance.CreateDate).To(Equal("2018-06-28T15:31:10+0900"))
			Expect(blockStorageSnapshotInstance.ServerImageProductCode).To(Equal("SPSW0LINUX000046"))
			Expect(blockStorageSnapshotInstance.OsInformation).To(Equal("CentOS 7.3 (64-bit)"))

		})
	})

	Describe("Unable to create Block Storage Snapshot Instance", func() {
		Context("when invalid BlockStorageInstanceNo is invalid", func() {
			BeforeEach(func() {
				gock.New("https://ncloud.apigw.ntruss.com").
					Post("/server").
					Reply(http.StatusInternalServerError).BodyString(`<responseError>
						<returnCode>1300</returnCode>
						<returnMessage>Please try your call again later.
					Temporarily out of service.
					If error continue, Please contact our customer service center.</returnMessage>
					</responseError>`)
			})
			AfterEach(func() {
				gock.Off()
			})

			It("should fail", func() {
				reqParams := new(RequestCreateBlockStorageSnapshotInstance)
				reqParams.BlockStorageInstanceNo = "blockstorage"
				reqParams.BlockStorageSnapshotName = "BlockStorageSnapshotInstance"
				reqParams.BlockStorageSnapshotDescription = "Block Storage Snapshot Instance by SDK go"

				conn := NewConnection(accessKey, secretKey)
				result, err := conn.CreateBlockStorageSnapshotInstance(reqParams)

				Expect(err.Error()).To(ContainSubstring("500 Internal Server Error"))
				Expect(result.ReturnCode).To(Equal(1300))
				Expect(result.ReturnMessage).To(Equal(`Please try your call again later.
					Temporarily out of service.
					If error continue, Please contact our customer service center.`))
			})
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
		It("should be failed : 401 Unauthorized", func() {
			reqParams := new(RequestCreateBlockStorageSnapshotInstance)
			reqParams.BlockStorageInstanceNo = "694783"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.CreateBlockStorageSnapshotInstance(reqParams)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Invalid consumerKey"))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})

	Describe("Check Arguments", func() {
		It("should be error : BlockStorageInstanceNo field is required", func() {
			reqParams := new(RequestCreateBlockStorageSnapshotInstance)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateBlockStorageSnapshotInstance(reqParams)

			Expect(err.Error()).To(ContainSubstring("BlockStorageInstanceNo field is required"))
		})
	})

})
