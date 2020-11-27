package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Block Storage Snapshot Instance List", func() {
	Describe("Get all Block Storage Snapshot Instance List", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getBlockStorageSnapshotInstanceListResponse>
   					<requestId>78ca2d57-d729-4077-b0ff-ce35248b71de</requestId>
   					<returnCode>0</returnCode>
   					<returnMessage>success</returnMessage>
   					<totalRows>1</totalRows>
   					<blockStorageSnapshotInstanceList>
      					<blockStorageSnapshot>
         					<blockStorageSnapshotInstanceNo>334118</blockStorageSnapshotInstanceNo>
         					<blockStorageSnapshotName>s0615bc</blockStorageSnapshotName>
         					<blockStorageSnapshotVolumeSize>53687091200</blockStorageSnapshotVolumeSize>
         					<originalBlockStorageInstanceNo>334105</originalBlockStorageInstanceNo>
         					<originalBlockStorageName>x0627</originalBlockStorageName>
         					<blockStorageSnapshotInstanceStatus>
            					<code>CREAT</code>
            					<codeName>Block storage CREATED state</codeName>
         					</blockStorageSnapshotInstanceStatus>
         					<blockStorageSnapshotInstanceOperation>
	            				<code>NULL</code>
    	        				<codeName>Block Storage NULLOP</codeName>
        	 				</blockStorageSnapshotInstanceOperation>
         					<createDate>2017-06-27T18:56:42+0900</createDate>
         					<blockStorageSnapshotInstanceDescription />
         					<serverImageProductCode>SPSW0WINNT000016</serverImageProductCode>
	         				<osInformation>Windows Server 2016 (64-bit)</osInformation>
	      				</blockStorageSnapshot>
   					</blockStorageSnapshotInstanceList>
				</getBlockStorageSnapshotInstanceListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get Block Storage Snapshot Instance list", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetBlockStorageSnapshotInstanceList(nil)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.BlockStorageSnapshotInstanceList)).To(Equal(1))

			blockStorageSnapshotInstance := result.BlockStorageSnapshotInstanceList[0]
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotInstanceNo).To(Equal("334118"))
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotName).To(Equal("s0615bc"))
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotVolumeSize).To(Equal(53687091200))
			Expect(blockStorageSnapshotInstance.OriginalBlockStorageInstanceNo).To(Equal("334105"))
			Expect(blockStorageSnapshotInstance.OriginalBlockStorageName).To(Equal("x0627"))
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotInstanceStatus.Code).To(Equal("CREAT"))
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotInstanceStatus.CodeName).To(Equal("Block storage CREATED state"))
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotInstanceOperation.Code).To(Equal("NULL"))
			Expect(blockStorageSnapshotInstance.BlockStorageSnapshotInstanceOperation.CodeName).To(Equal("Block Storage NULLOP"))
			Expect(blockStorageSnapshotInstance.CreateDate).To(Equal("2017-06-27T18:56:42+0900"))
			Expect(blockStorageSnapshotInstance.ServerImageProductCode).To(Equal("SPSW0WINNT000016"))
			Expect(blockStorageSnapshotInstance.OsInformation).To(Equal("Windows Server 2016 (64-bit)"))
		})
	})

	Describe("There is no Block Storage Snapshot Instance list", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getBlockStorageSnapshotInstanceListResponse>
					<requestId>4500fb43-d21e-4db8-a89f-2ca44cb20f91</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>0</totalRows>
					<blockStorageSnapshotInstanceList/>
				</getBlockStorageSnapshotInstanceListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get No Block Storage Snapshot Instance list", func() {
			reqParams := new(RequestGetBlockStorageSnapshotInstanceList)
			reqParams.BlockStorageSnapshotInstanceNoList = []string{"559513"}
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetBlockStorageSnapshotInstanceList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(0))
			Expect(len(result.BlockStorageSnapshotInstanceList)).To(Equal(0))
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
		It("should be failed by authorization fail", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetBlockStorageSnapshotInstanceList(nil)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Expired url."))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})
})
