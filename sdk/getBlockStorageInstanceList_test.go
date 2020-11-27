package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Block Storage Instance List", func() {
	Describe("Get all Block Storage Instance List", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getBlockStorageInstanceListResponse>
					<requestId>79e97c32-5cd2-472d-8074-d92edd3a9c99</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>2</totalRows>
					<blockStorageInstanceList>
						<blockStorageInstance>
							<blockStorageInstanceNo>559512</blockStorageInstanceNo>
							<serverInstanceNo>559511</serverInstanceNo>
							<serverName>svr-9bb6f7aa8f7420d</serverName>
							<blockStorageType>
								<code>BASIC</code>
								<codeName>Basic BS</codeName>
							</blockStorageType>
							<blockStorageName>svr-9bb6f7aa8f7420d</blockStorageName>
							<blockStorageSize>53687091200</blockStorageSize>
							<deviceName>/dev/xvda</deviceName>
							<blockStorageProductCode>SPBSTBSTBS000001</blockStorageProductCode>
							<blockStorageInstanceStatus>
								<code>ATTAC</code>
								<codeName>Block storage ATTACHED state</codeName>
							</blockStorageInstanceStatus>
							<blockStorageInstanceOperation>
								<code>NULL</code>
								<codeName>Block Storage NULLOP</codeName>
							</blockStorageInstanceOperation>
							<blockStorageInstanceStatusName>attached</blockStorageInstanceStatusName>
							<createDate>2017-12-10T14:15:03+0900</createDate>
							<blockStorageInstanceDescription>svr-9bb6f7aa8f7420d&apos;s basic storage</blockStorageInstanceDescription>
							<diskType>
								<code>NET</code>
								<codeName>Network Storage</codeName>
							</diskType>
							<diskDetailType>
								<code>HDD</code>
								<codeName>HDD</codeName>
							</diskDetailType>
							<zone>
								<zoneNo>2</zoneNo>
								<zoneName>KR-1</zoneName>
								<zoneCode>KR-1</zoneCode>
								<zoneDescription>가산 zone</zoneDescription>
								<regionNo>1</regionNo>
							</zone>
						</blockStorageInstance>
						<blockStorageInstance>
							<blockStorageInstanceNo>559515</blockStorageInstanceNo>
							<serverInstanceNo>559514</serverInstanceNo>
							<serverName>svr-9bb6f7ae08d6966</serverName>
							<blockStorageType>
								<code>BASIC</code>
								<codeName>Basic BS</codeName>
							</blockStorageType>
							<blockStorageName>svr-9bb6f7ae08d6966</blockStorageName>
							<blockStorageSize>53687091200</blockStorageSize>
							<deviceName>/dev/xvda</deviceName>
							<blockStorageProductCode>SPBSTBSTBS000001</blockStorageProductCode>
							<blockStorageInstanceStatus>
								<code>ATTAC</code>
								<codeName>Block storage ATTACHED state</codeName>
							</blockStorageInstanceStatus>
							<blockStorageInstanceOperation>
								<code>NULL</code>
								<codeName>Block Storage NULLOP</codeName>
							</blockStorageInstanceOperation>
							<blockStorageInstanceStatusName>attached</blockStorageInstanceStatusName>
							<createDate>2017-12-10T14:17:32+0900</createDate>
							<blockStorageInstanceDescription>svr-9bb6f7ae08d6966&apos;s basic storage</blockStorageInstanceDescription>
							<diskType>
								<code>NET</code>
								<codeName>Network Storage</codeName>
							</diskType>
							<diskDetailType>
								<code>HDD</code>
								<codeName>HDD</codeName>
							</diskDetailType>
							<zone>
								<zoneNo>2</zoneNo>
								<zoneName>KR-1</zoneName>
								<zoneCode>KR-1</zoneCode>
								<zoneDescription>가산 zone</zoneDescription>
								<regionNo>1</regionNo>
							</zone>
						</blockStorageInstance>
					</blockStorageInstanceList>
				</getBlockStorageInstanceListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get Block Storage Instance list", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetBlockStorageInstance(nil)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(2))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.BlockStorageInstance)).To(Equal(2))

			blockStorageInstance := result.BlockStorageInstance[0]
			Expect(blockStorageInstance.BlockStorageInstanceNo).To(Equal("559512"))
			Expect(blockStorageInstance.ServerInstanceNo).To(Equal("559511"))
			Expect(blockStorageInstance.ServerName).To(Equal("svr-9bb6f7aa8f7420d"))
			Expect(blockStorageInstance.BlockStorageType.Code).To(Equal("BASIC"))
			Expect(blockStorageInstance.BlockStorageType.CodeName).To(Equal("Basic BS"))
			Expect(blockStorageInstance.BlockStorageName).To(Equal("svr-9bb6f7aa8f7420d"))
			Expect(blockStorageInstance.BlockStorageSize).To(Equal(53687091200))
			Expect(blockStorageInstance.DeviceName).To(Equal("/dev/xvda"))
			Expect(blockStorageInstance.BlockStorageProductCode).To(Equal("SPBSTBSTBS000001"))
			Expect(blockStorageInstance.BlockStorageInstanceStatus.Code).To(Equal("ATTAC"))
			Expect(blockStorageInstance.BlockStorageInstanceStatus.CodeName).To(Equal("Block storage ATTACHED state"))
			Expect(blockStorageInstance.BlockStorageInstanceOperation.Code).To(Equal("NULL"))
			Expect(blockStorageInstance.BlockStorageInstanceOperation.CodeName).To(Equal("Block Storage NULLOP"))
			Expect(blockStorageInstance.BlockStorageInstanceStatusName).To(Equal("attached"))
			Expect(blockStorageInstance.CreateDate).To(Equal("2017-12-10T14:15:03+0900"))
			Expect(blockStorageInstance.BlockStorageInstanceDescription).To(Equal("svr-9bb6f7aa8f7420d's basic storage"))
			Expect(blockStorageInstance.DiskType.Code).To(Equal("NET"))
			Expect(blockStorageInstance.DiskType.CodeName).To(Equal("Network Storage"))
			Expect(blockStorageInstance.DiskDetailType.Code).To(Equal("HDD"))
			Expect(blockStorageInstance.DiskDetailType.CodeName).To(Equal("HDD"))
			Expect(blockStorageInstance.Zone.ZoneNo).To(Equal("2"))
			Expect(blockStorageInstance.Zone.ZoneName).To(Equal("KR-1"))
			Expect(blockStorageInstance.Zone.ZoneCode).To(Equal("KR-1"))
			Expect(blockStorageInstance.Zone.ZoneDescription).To(Equal("가산 zone"))
			Expect(blockStorageInstance.Zone.RegionNo).To(Equal("1"))

			blockStorageInstance = result.BlockStorageInstance[1]
			Expect(blockStorageInstance.BlockStorageInstanceNo).To(Equal("559515"))
			Expect(blockStorageInstance.ServerInstanceNo).To(Equal("559514"))
			Expect(blockStorageInstance.ServerName).To(Equal("svr-9bb6f7ae08d6966"))
			Expect(blockStorageInstance.BlockStorageType.Code).To(Equal("BASIC"))
			Expect(blockStorageInstance.BlockStorageType.CodeName).To(Equal("Basic BS"))
			Expect(blockStorageInstance.BlockStorageName).To(Equal("svr-9bb6f7ae08d6966"))
			Expect(blockStorageInstance.BlockStorageSize).To(Equal(53687091200))
			Expect(blockStorageInstance.DeviceName).To(Equal("/dev/xvda"))
			Expect(blockStorageInstance.BlockStorageProductCode).To(Equal("SPBSTBSTBS000001"))
			Expect(blockStorageInstance.BlockStorageInstanceStatus.Code).To(Equal("ATTAC"))
			Expect(blockStorageInstance.BlockStorageInstanceStatus.CodeName).To(Equal("Block storage ATTACHED state"))
			Expect(blockStorageInstance.BlockStorageInstanceOperation.Code).To(Equal("NULL"))
			Expect(blockStorageInstance.BlockStorageInstanceOperation.CodeName).To(Equal("Block Storage NULLOP"))
			Expect(blockStorageInstance.BlockStorageInstanceStatusName).To(Equal("attached"))
			Expect(blockStorageInstance.CreateDate).To(Equal("2017-12-10T14:17:32+0900"))
			Expect(blockStorageInstance.BlockStorageInstanceDescription).To(Equal("svr-9bb6f7ae08d6966's basic storage"))
			Expect(blockStorageInstance.DiskType.Code).To(Equal("NET"))
			Expect(blockStorageInstance.DiskType.CodeName).To(Equal("Network Storage"))
			Expect(blockStorageInstance.DiskDetailType.Code).To(Equal("HDD"))
			Expect(blockStorageInstance.DiskDetailType.CodeName).To(Equal("HDD"))
			Expect(blockStorageInstance.Zone.ZoneNo).To(Equal("2"))
			Expect(blockStorageInstance.Zone.ZoneName).To(Equal("KR-1"))
			Expect(blockStorageInstance.Zone.ZoneCode).To(Equal("KR-1"))
			Expect(blockStorageInstance.Zone.ZoneDescription).To(Equal("가산 zone"))
			Expect(blockStorageInstance.Zone.RegionNo).To(Equal("1"))
		})
	})

	Describe("Get One Block Storage Instance List which ServerInstanceNoList.1 is 320897", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getBlockStorageInstanceListResponse>
					<requestId>f5a87a5d-4fbb-4f57-8ec7-3b1543ed3167</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<blockStorageInstanceList>
						<blockStorageInstance>
							<blockStorageInstanceNo>559512</blockStorageInstanceNo>
							<serverInstanceNo>559511</serverInstanceNo>
							<serverName>svr-9bb6f7aa8f7420d</serverName>
							<blockStorageType>
								<code>BASIC</code>
								<codeName>Basic BS</codeName>
							</blockStorageType>
							<blockStorageName>svr-9bb6f7aa8f7420d</blockStorageName>
							<blockStorageSize>53687091200</blockStorageSize>
							<deviceName>/dev/xvda</deviceName>
							<blockStorageProductCode>SPBSTBSTBS000001</blockStorageProductCode>
							<blockStorageInstanceStatus>
								<code>ATTAC</code>
								<codeName>Block storage ATTACHED state</codeName>
							</blockStorageInstanceStatus>
							<blockStorageInstanceOperation>
								<code>NULL</code>
								<codeName>Block Storage NULLOP</codeName>
							</blockStorageInstanceOperation>
							<blockStorageInstanceStatusName>attached</blockStorageInstanceStatusName>
							<createDate>2017-12-10T14:15:03+0900</createDate>
							<blockStorageInstanceDescription>svr-9bb6f7aa8f7420d&apos;s basic storage</blockStorageInstanceDescription>
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
				</getBlockStorageInstanceListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get One Block Storage Instance", func() {
			reqParams := new(RequestBlockStorageInstanceList)
			reqParams.BlockStorageInstanceNoList = []string{"559512"}
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetBlockStorageInstance(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.BlockStorageInstance)).To(Equal(1))

			blockStorageInstance := result.BlockStorageInstance[0]
			Expect(blockStorageInstance.BlockStorageInstanceNo).To(Equal("559512"))
			Expect(blockStorageInstance.ServerInstanceNo).To(Equal("559511"))
			Expect(blockStorageInstance.ServerName).To(Equal("svr-9bb6f7aa8f7420d"))
			Expect(blockStorageInstance.BlockStorageType.Code).To(Equal("BASIC"))
			Expect(blockStorageInstance.BlockStorageType.CodeName).To(Equal("Basic BS"))
			Expect(blockStorageInstance.BlockStorageName).To(Equal("svr-9bb6f7aa8f7420d"))
			Expect(blockStorageInstance.BlockStorageSize).To(Equal(53687091200))
			Expect(blockStorageInstance.DeviceName).To(Equal("/dev/xvda"))
			Expect(blockStorageInstance.BlockStorageProductCode).To(Equal("SPBSTBSTBS000001"))
			Expect(blockStorageInstance.BlockStorageInstanceStatus.Code).To(Equal("ATTAC"))
			Expect(blockStorageInstance.BlockStorageInstanceStatus.CodeName).To(Equal("Block storage ATTACHED state"))
			Expect(blockStorageInstance.BlockStorageInstanceOperation.Code).To(Equal("NULL"))
			Expect(blockStorageInstance.BlockStorageInstanceOperation.CodeName).To(Equal("Block Storage NULLOP"))
			Expect(blockStorageInstance.BlockStorageInstanceStatusName).To(Equal("attached"))
			Expect(blockStorageInstance.CreateDate).To(Equal("2017-12-10T14:15:03+0900"))
			Expect(blockStorageInstance.BlockStorageInstanceDescription).To(Equal("svr-9bb6f7aa8f7420d's basic storage"))
			Expect(blockStorageInstance.DiskType.Code).To(Equal("NET"))
			Expect(blockStorageInstance.DiskType.CodeName).To(Equal("Network Storage"))
			Expect(blockStorageInstance.DiskDetailType.Code).To(Equal("HDD"))
			Expect(blockStorageInstance.DiskDetailType.CodeName).To(Equal("HDD"))

		})
	})

	Describe("There is no Block Storage Instance list", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getBlockStorageInstanceListResponse>
					<requestId>4500fb43-d21e-4db8-a89f-2ca44cb20f91</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>0</totalRows>
					<blockStorageInstanceList/>
				</getBlockStorageInstanceListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get No Block Storage Instance list", func() {
			reqParams := new(RequestBlockStorageInstanceList)
			reqParams.BlockStorageInstanceNoList = []string{"559513"}
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetBlockStorageInstance(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(0))
			Expect(len(result.BlockStorageInstance)).To(Equal(0))
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
			result, err := conn.GetServerInstanceList(nil)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Expired url."))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})
})
