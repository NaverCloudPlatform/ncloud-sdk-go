package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Delete Launch Configuration", func() {
	Describe("Delete Launch Configuration", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/autoscaling").
				Reply(http.StatusOK).BodyString(`<deleteAutoScalingLaunchConfigurationResponse>
				<requestId>4b8f08f4-eef7-42ba-9a2c-791ad979cf1d</requestId>
				<returnCode>0</returnCode>
				<returnMessage>success</returnMessage>
			</deleteAutoScalingLaunchConfigurationResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should delete Launch Configuration", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.DeleteLaunchConfiguration("ffwefwef11")

			Expect(err).To(BeNil())
			Expect(result.RequestID).To(Equal("4b8f08f4-eef7-42ba-9a2c-791ad979cf1d"))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
		})
	})

	Describe("Not exist Launch Configuration name", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/autoscaling").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
				<returnCode>50120</returnCode>
				<returnMessage>An invalid or out-of-range value was supplied for the input parameter.
							The named Launch Configuration is not found.</returnMessage>
			</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should fail if Launch Configuration Name is not exist", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.DeleteLaunchConfiguration("test01912")

			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
			Expect(result.ReturnCode).To(Equal(50120))
			Expect(result.ReturnMessage).To(ContainSubstring("An invalid or out-of-range value was supplied for the input parameter."))
			Expect(result.ReturnMessage).To(ContainSubstring("The named Launch Configuration is not found."))
		})
	})

	Describe("Launch Configuration length", func() {
		It("should be smaller than 255", func() {
			conn := NewConnection(accessKey, secretKey)
			_, err := conn.DeleteLaunchConfiguration(String(256))

			Expect(err.Error()).To(ContainSubstring("expected length of LaunchConfigurationName to be in the range (0 - 255)"))
		})
	})
})
