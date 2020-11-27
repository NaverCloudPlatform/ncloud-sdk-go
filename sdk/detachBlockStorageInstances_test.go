package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Detach Block Storage Instance", func() {
	Describe("Detach Block Storage Instance", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusOK).BodyString(`
					<detachBlockStorageInstancesResponse>
					<requestId>b9e4a598-a47b-4635-924f-9be17fba4a35</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>2</totalRows>
					<blockStorageInstanceList>
						<blockStorageInstance>
							<blockStorageInstanceNo>825683</blockStorageInstanceNo>
							<serverInstanceNo>805847</serverInstanceNo>
							<serverName>aaa007</serverName>
							<blockStorageType>
								<code>SVRBS</code>
								<codeName>Server BS</codeName>
							</blockStorageType>
							<blockStorageName>bgss1</blockStorageName>
							<blockStorageSize>10737418240</blockStorageSize>
							<deviceName>/dev/xvdc</deviceName>
							<blockStorageProductCode>SPBSTBSTAD000006</blockStorageProductCode>
							<blockStorageInstanceStatus>
								<code>ATTAC</code>
								<codeName>Block storage ATTACHED state</codeName>
							</blockStorageInstanceStatus>
							<blockStorageInstanceOperation>
								<code>DETAC</code>
								<codeName>Block Storage RESTART OP</codeName>
							</blockStorageInstanceOperation>
							<blockStorageInstanceStatusName>detaching</blockStorageInstanceStatusName>
							<createDate>2018-06-18T20:22:46+0900</createDate>
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
						<blockStorageInstance>
							<blockStorageInstanceNo>825649</blockStorageInstanceNo>
							<serverInstanceNo>805844</serverInstanceNo>
							<serverName>aaa006</serverName>
							<blockStorageType>
								<code>SVRBS</code>
								<codeName>Server BS</codeName>
							</blockStorageType>
							<blockStorageName>bgss</blockStorageName>
							<blockStorageSize>10737418240</blockStorageSize>
							<deviceName>/dev/xvdb</deviceName>
							<blockStorageProductCode>SPBSTBSTAD000006</blockStorageProductCode>
							<blockStorageInstanceStatus>
								<code>ATTAC</code>
								<codeName>Block storage ATTACHED state</codeName>
							</blockStorageInstanceStatus>
							<blockStorageInstanceOperation>
								<code>DETAC</code>
								<codeName>Block Storage RESTART OP</codeName>
							</blockStorageInstanceOperation>
							<blockStorageInstanceStatusName>detaching</blockStorageInstanceStatusName>
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
				</detachBlockStorageInstancesResponse>
					`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should Detach Block Storage Instance", func() {
			conn := NewConnection(accessKey, secretKey)
			reqParams := &RequestDetachBlockStorageInstance{
				BlockStorageInstanceNoList: []string{"825649", "825683"},
			}
			result, err := conn.DetachBlockStorageInstance(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(2))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.BlockStorageInstance)).To(Equal(len(reqParams.BlockStorageInstanceNoList)))
		})
	})

	Describe("Invalid BlockStorageInstanceNoList", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusBadRequest).BodyString(`
		<responseError>
		  <returnCode>24109</returnCode>
		  <returnMessage>The input parameter instance number is invalid.</returnMessage>
		</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should not  Detach Block Storage Instance by invalid BlockStorageInstanceNoList", func() {
			conn := NewConnection(accessKey, secretKey)
			reqParams := &RequestDetachBlockStorageInstance{
				BlockStorageInstanceNoList: []string{"825600", "825601"},
			}
			result, err := conn.DetachBlockStorageInstance(reqParams)

			Expect(result.ReturnCode).To(Equal(24109))
			Expect(result.ReturnMessage).To(ContainSubstring("The input parameter instance number is invalid."))
			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
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
			reqParams := &RequestDetachBlockStorageInstance{
				BlockStorageInstanceNoList: []string{"825649"},
			}
			result, err := conn.DetachBlockStorageInstance(reqParams)

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
			reqParams := &RequestDetachBlockStorageInstance{
				BlockStorageInstanceNoList: []string{"825649"},
			}
			result, err := conn.DetachBlockStorageInstance(reqParams)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Expired url."))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})
})
