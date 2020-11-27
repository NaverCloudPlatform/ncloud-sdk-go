package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gock "gopkg.in/h2non/gock.v1"
)

var _ = Describe("Get Zone List", func() {
	Describe("Get Zone List", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getZoneListResponse>
					<requestId>f650ca08-937c-49e5-9a33-aaa4e1d3412a</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<zoneList>
						<zone>
							<zoneNo>2</zoneNo>
							<zoneName>KR-1</zoneName>
							<zoneDescription>KR-1 zone</zoneDescription>
						</zone>
					</zoneList>
				</getZoneListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get Zone List", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetZoneList("")

			Expect(err).To(BeNil())
			Expect(result.RequestID).To(Equal("f650ca08-937c-49e5-9a33-aaa4e1d3412a"))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.Zone)).To(Equal(1))

			zone := result.Zone[0]
			Expect(zone.ZoneNo).To(Equal("2"))
			Expect(zone.ZoneName).To(Equal("KR-1"))
			Expect(zone.ZoneDescription).To(Equal("KR-1 zone"))
		})

	})

	Describe("Invalid RegionNo", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
					<returnCode>1300</returnCode>
					<returnMessage>Please try your call again later.
				Temporarily out of service.
				If error continue, Please contact our customer service center.</returnMessage>
				</responseError>`)
		})

		AfterEach(func() {
			gock.Off()
		})

		It("should fail if RegionNo is invalid", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetZoneList("0")

			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
			Expect(result.ReturnCode).To(Equal(1300))
			Expect(result.ReturnMessage).To(Equal(`Please try your call again later.
				Temporarily out of service.
				If error continue, Please contact our customer service center.`))
		})
	})
})
