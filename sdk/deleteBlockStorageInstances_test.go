package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Delete Block Storage Instances", func() {
	Describe("Delete Block Storage Instances", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusOK).BodyString(`<deleteBlockStorageInstancesResponse>
					<requestId>4d8f4e3e-e8da-42f8-87ad-21986e96fdae</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<blockStorageInstanceList>
						<blockStorageInstance>
							<blockStorageInstanceNo>480691</blockStorageInstanceNo>
							<serverInstanceNo>449157</serverInstanceNo>
							<serverName>ms-ub-01</serverName>
							<blockStorageType>
								<code>SVRBS</code>
								<codeName>Server BS</codeName>
							</blockStorageType>
							<blockStorageName>dddd</blockStorageName>
							<blockStorageSize>10737418240</blockStorageSize>
							<deviceName>/dev/xvdb</deviceName>
							<blockStorageProductCode>SPBSTBSTAD000002</blockStorageProductCode>
							<blockStorageInstanceStatus>
								<code>ATTAC</code>
								<codeName>Block storage ATTACHED state</codeName>
							</blockStorageInstanceStatus>
							<blockStorageInstanceOperation>
								<code>DETAC</code>
								<codeName>Block Storage RESTART OP</codeName>
							</blockStorageInstanceOperation>
							<blockStorageInstanceStatusName>detaching</blockStorageInstanceStatusName>
							<createDate>2017-12-10T13:03:27+0900</createDate>
							<blockStorageInstanceDescription></blockStorageInstanceDescription>
							<diskType>
								<code>NET</code>
								<codeName>Network Storage</codeName>
							</diskType>
							<diskDetailType>
								<code>HDD</code>
								<codeName>HDD</codeName>
							</diskDetailType>
						</blockStorageInstance>
					</blockStorageInstanceList>
				</deleteBlockStorageInstancesResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should Delete Block Storage Instance", func() {
			reqParams := []string{"480691"}
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.DeleteBlockStorageInstances(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.BlockStorageInstance)).To(Equal(1))

			blockStorageInstance := result.BlockStorageInstance[0]
			Expect(blockStorageInstance.BlockStorageInstanceNo).To(Equal("480691"))
			Expect(blockStorageInstance.ServerInstanceNo).To(Equal("449157"))
			Expect(blockStorageInstance.ServerName).To(Equal("ms-ub-01"))
			Expect(blockStorageInstance.BlockStorageType.Code).To(Equal("SVRBS"))
			Expect(blockStorageInstance.BlockStorageType.CodeName).To(Equal("Server BS"))
			Expect(blockStorageInstance.BlockStorageName).To(Equal("dddd"))
			Expect(blockStorageInstance.BlockStorageSize).To(Equal(10737418240))
			Expect(blockStorageInstance.DeviceName).To(Equal("/dev/xvdb"))
			Expect(blockStorageInstance.BlockStorageProductCode).To(Equal("SPBSTBSTAD000002"))
			Expect(blockStorageInstance.BlockStorageInstanceStatus.Code).To(Equal("ATTAC"))
			Expect(blockStorageInstance.BlockStorageInstanceStatus.CodeName).To(Equal("Block storage ATTACHED state"))
			Expect(blockStorageInstance.BlockStorageInstanceOperation.Code).To(Equal("DETAC"))
			Expect(blockStorageInstance.BlockStorageInstanceOperation.CodeName).To(Equal("Block Storage RESTART OP"))
			Expect(blockStorageInstance.BlockStorageInstanceStatusName).To(Equal("detaching"))
			Expect(blockStorageInstance.BlockStorageInstanceDescription).To(Equal(""))
			Expect(blockStorageInstance.DiskType.Code).To(Equal("NET"))
			Expect(blockStorageInstance.DiskType.CodeName).To(Equal("Network Storage"))
			Expect(blockStorageInstance.DiskDetailType.Code).To(Equal("HDD"))
			Expect(blockStorageInstance.DiskDetailType.CodeName).To(Equal("HDD"))
		})
	})

	Describe("Unable to Delete Block Storage Instance", func() {
		Context("when invalid block storage instance no is invalid", func() {
			BeforeEach(func() {
				gock.New("https://ncloud.apigw.ntruss.com").
					Post("/server").
					Reply(http.StatusBadRequest).BodyString(`<responseError>
						<returnCode>24121</returnCode>
						<returnMessage>The input parameter storage instance number is invalid. </returnMessage>
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
				Expect(result.ReturnMessage).To(Equal(`The input parameter storage instance number is invalid.`))
			})
		})
	})
})
