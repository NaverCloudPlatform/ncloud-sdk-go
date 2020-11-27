package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetRegionList", func() {
	Describe("Get all region list", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getRegionListResponse>
					<requestId>8d548f30-daac-4599-a9c0-e5bb4a3061cc</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>6</totalRows>
					<regionList>
						<region>
							<regionNo>1</regionNo>
							<regionCode>KR</regionCode>
							<regionName>Korea</regionName>
						</region>
						<region>
							<regionNo>2</regionNo>
							<regionCode>USW</regionCode>
							<regionName>US-West</regionName>
						</region>
						<region>
							<regionNo>3</regionNo>
							<regionCode>HK</regionCode>
							<regionName>HongKong</regionName>
						</region>
						<region>
							<regionNo>4</regionNo>
							<regionCode>SG</regionCode>
							<regionName>Singapore</regionName>
						</region>
						<region>
							<regionNo>5</regionNo>
							<regionCode>JP</regionCode>
							<regionName>Japan</regionName>
						</region>
						<region>
							<regionNo>6</regionNo>
							<regionCode>DE</regionCode>
							<regionName>Germany</regionName>
						</region>
					</regionList>
				</getRegionListResponse>`)

		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get region list", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetRegionList()

			Expect(err).To(BeNil())
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(result.TotalRows).To(Equal(6))

			region := result.RegionList[0]
			Expect(region.RegionNo).To(Equal("1"))
			Expect(region.RegionCode).To(Equal("KR"))
			Expect(region.RegionName).To(Equal("Korea"))

		})
	})

})
