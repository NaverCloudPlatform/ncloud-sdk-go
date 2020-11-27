package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetLoginKeyList", func() {
	Describe("Get all login key List", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getLoginKeyListResponse>
					<requestId>be885b84-4fa7-4f74-a688-b7b8dee72950</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>3</totalRows>
					<loginKeyList>
						<loginKey>
							<fingerprint>a7:4b:ac:40:2a:8a:96:d6:83:4e:0c:0d:56:34:fa:1a</fingerprint>
							<keyName>fff</keyName>
							<createDate>2017-12-01T13:11:14+0900</createDate>
						</loginKey>
						<loginKey>
							<fingerprint>a2:b9:be:1b:ec:92:6a:2a:5d:d3:44:9e:75:ee:8d:fc</fingerprint>
							<keyName>packer-1512019307</keyName>
							<createDate>2017-11-30T14:21:47+0900</createDate>
						</loginKey>
						<loginKey>
							<fingerprint>f4:16:72:b9:36:b4:1e:9b:74:2a:79:5e:a0:75:98:e8</fingerprint>
							<keyName>packer-1512019276</keyName>
							<createDate>2017-11-30T14:21:17+0900</createDate>
						</loginKey>
					</loginKeyList>
				</getLoginKeyListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get login key list", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetLoginKeyList(nil)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(3))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.LoginKeyList)).To(Equal(3))

			loginKey := result.LoginKeyList[0]
			Expect(loginKey.Fingerprint).To(Equal("a7:4b:ac:40:2a:8a:96:d6:83:4e:0c:0d:56:34:fa:1a"))
			Expect(loginKey.KeyName).To(Equal("fff"))
			Expect(loginKey.CreateDate).To(Equal("2017-12-01T13:11:14+0900"))

			loginKey = result.LoginKeyList[1]
			Expect(loginKey.Fingerprint).To(Equal("a2:b9:be:1b:ec:92:6a:2a:5d:d3:44:9e:75:ee:8d:fc"))
			Expect(loginKey.KeyName).To(Equal("packer-1512019307"))
			Expect(loginKey.CreateDate).To(Equal("2017-11-30T14:21:47+0900"))

			loginKey = result.LoginKeyList[2]
			Expect(loginKey.Fingerprint).To(Equal("f4:16:72:b9:36:b4:1e:9b:74:2a:79:5e:a0:75:98:e8"))
			Expect(loginKey.KeyName).To(Equal("packer-1512019276"))
			Expect(loginKey.CreateDate).To(Equal("2017-11-30T14:21:17+0900"))
		})
	})

	Describe("Get one loginKey list which KeyName is packer-1512019276", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getLoginKeyListResponse>
					<requestId>53629aee-b34b-4371-b834-94c82e0077fe</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<loginKeyList>
						<loginKey>
							<fingerprint>f4:16:72:b9:36:b4:1e:9b:74:2a:79:5e:a0:75:98:e8</fingerprint>
							<keyName>packer-1512019276</keyName>
							<createDate>2017-11-30T14:21:17+0900</createDate>
						</loginKey>
					</loginKeyList>
				</getLoginKeyListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get only one login key list", func() {
			reqParams := new(RequestGetLoginKeyList)
			reqParams.KeyName = "packer-1512019276"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetLoginKeyList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.LoginKeyList)).To(Equal(1))

			loginKey := result.LoginKeyList[0]
			Expect(loginKey.Fingerprint).To(Equal("f4:16:72:b9:36:b4:1e:9b:74:2a:79:5e:a0:75:98:e8"))
			Expect(loginKey.KeyName).To(Equal("packer-1512019276"))
			Expect(loginKey.CreateDate).To(Equal("2017-11-30T14:21:17+0900"))
		})
	})

	Describe("There is no login key list list", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getLoginKeyListResponse>
					<requestId>b275cef4-117f-428d-9bad-825a93f51227</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>0</totalRows>
					<loginKeyList/>
				</getLoginKeyListResponse>`)

		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get no login key list", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetLoginKeyList(nil)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(0))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.LoginKeyList)).To(Equal(0))
		})
	})
})
