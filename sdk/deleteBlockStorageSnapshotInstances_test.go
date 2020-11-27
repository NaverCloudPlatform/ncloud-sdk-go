package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Delete Block Storage Snapshot Instances", func() {
	Describe("Delete Block Storage Snapshot Instances", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusOK).BodyString(`<deleteBlockStorageSnapshotInstancesResponse>
				<script />
   				<requestId>15104c45-243f-4e98-8841-4cff02e08bb6</requestId>
   				<returnCode>0</returnCode>
				<returnMessage>success</returnMessage>
   				<totalRows>1</totalRows>
   				<blockStorageSnapshotInstanceList>
      				<blockStorageSnapshot>
     					<blockStorageSnapshotInstanceNo>693535</blockStorageSnapshotInstanceNo>
     					<blockStorageSnapshotName>snap164355fe5f4</blockStorageSnapshotName>
     					<blockStorageSnapshotVolumeSize>53687091200</blockStorageSnapshotVolumeSize>
     					<originalBlockStorageInstanceNo>691658</originalBlockStorageInstanceNo>
     					<originalBlockStorageName>ms-ad-test-01</originalBlockStorageName>
     					<blockStorageSnapshotInstanceStatus>
        					<code>CREAT</code>
        					<codeName>Block storage CREATED state</codeName>
     					</blockStorageSnapshotInstanceStatus>
     					<blockStorageSnapshotInstanceOperation>
        					<code>TERMT</code>
        					<codeName>Block Storage TERMINATE OP</codeName>
						</blockStorageSnapshotInstanceOperation>
     					<blockStorageSnapshotInstanceStatusName>terminating</blockStorageSnapshotInstanceStatusName>
     					<createDate>2018-06-25T14:17:14+0900</createDate>
     					<blockStorageSnapshotInstanceDescription />
     					<serverImageProductCode>SPSW0WINNTEN0016</serverImageProductCode>
     					<osInformation>Windows Server 2016 (64-bit) English Edition</osInformation>
      				</blockStorageSnapshot>
   				</blockStorageSnapshotInstanceList>
			</deleteBlockStorageSnapshotInstancesResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should Delete Block Storage Snapthot Instance", func() {
			reqParams := []string{"693535"}
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.DeleteBlockStorageSnapshotInstances(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.BlockStorageSnapshotInstanceList)).To(Equal(1))

			blockStorageSnapshotInstance := result.BlockStorageSnapshotInstanceList[0]
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotInstanceNo).To(Equal("693535"))
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotName).To(Equal("snap164355fe5f4"))
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotVolumeSize).To(Equal(53687091200))
			Expect(blockStorageSnapshotInstance.OriginalBlockStorageInstanceNo).To(Equal("691658"))
			Expect(blockStorageSnapshotInstance.OriginalBlockStorageName).To(Equal("ms-ad-test-01"))
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotInstanceStatus.Code).To(Equal("CREAT"))
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotInstanceStatus.CodeName).To(Equal("Block storage CREATED state"))
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotInstanceOperation.Code).To(Equal("TERMT"))
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotInstanceOperation.CodeName).To(Equal("Block Storage TERMINATE OP"))
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotInstanceStatusName).To(Equal("terminating"))
			Expect(blockStorageSnapshotInstance.CreateDate).To(Equal("2018-06-25T14:17:14+0900"))
			Expect(blockStorageSnapshotInstance.ServerImageProductCode).To(Equal("SPSW0WINNTEN0016"))
			Expect(blockStorageSnapshotInstance.OsInformation).To(Equal("Windows Server 2016 (64-bit) English Edition"))
		})
	})

	Describe("Unable to Delete Block Storage Snapshot Instance", func() {
		Context("when invalid Block Storage Snapshot Instance no is invalid", func() {
			BeforeEach(func() {
				gock.New("https://ncloud.apigw.ntruss.com").
					Post("/server").
					Reply(http.StatusBadRequest).BodyString(`<responseError>
						<returnCode>24121</returnCode>
						<returnMessage>The input parameter storage snapshot instance number is invalid. </returnMessage>
					</responseError>`)
			})
			AfterEach(func() {
				gock.Off()
			})

			It("should fail", func() {
				reqParams := []string{"480691"}
				conn := NewConnection(accessKey, secretKey)
				result, err := conn.DeleteBlockStorageInstances(reqParams)

				Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
				Expect(result.ReturnCode).To(Equal(24121))
				Expect(result.ReturnMessage).To(Equal(`The input parameter storage snapshot instance number is invalid.`))
			})
		})
	})
})
