package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create Launch Configuration", func() {
	Describe("Create Launch Configuration", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/autoscaling").
				Reply(http.StatusOK).BodyString(`<createLaunchConfigurationResponse>
				<requestId>f440cba9-60a0-4b8b-8a8a-466b8b99a23f</requestId>
				<returnCode>0</returnCode>
				<returnMessage>success</returnMessage>
				<totalRows>1</totalRows>
				<launchConfigurationList>
					<launchConfiguration>
						<launchConfigurationName>lctest</launchConfigurationName>
						<serverImageProductCode>SPSW0LINUX000045</serverImageProductCode>
						<serverProductCode>SPSVRSTAND000003</serverProductCode>
						<memberServerImageNo></memberServerImageNo>
						<loginKeyName>tf-key-test</loginKeyName>
						<createDate>2018-07-10T13:25:07+0900</createDate>
						<userData></userData>
						<accessControlGroupList>
							<accessControlGroup>
								<accessControlGroupConfigurationNo>4964</accessControlGroupConfigurationNo>
								<accessControlGroupName>ncloud-default-acg</accessControlGroupName>
								<accessControlGroupDescription>Default AccessControlGroup</accessControlGroupDescription>
								<isDefault>true</isDefault>
								<createDate>2017-02-23T10:25:39+0900</createDate>
							</accessControlGroup>
							<accessControlGroup>
								<accessControlGroupConfigurationNo>30067</accessControlGroupConfigurationNo>
								<accessControlGroupName>httpdport</accessControlGroupName>
								<accessControlGroupDescription></accessControlGroupDescription>
								<isDefault>false</isDefault>
								<createDate>2018-01-07T10:17:14+0900</createDate>
							</accessControlGroup>
						</accessControlGroupList>
					</launchConfiguration>
				</launchConfigurationList>
			</createLaunchConfigurationResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should Create Launch Configuration", func() {
			reqParams := new(RequestCreateLaunchConfiguration)
			reqParams.LaunchConfigurationName = "lctest"
			reqParams.ServerImageProductCode = "SPSW0LINUX000045"
			reqParams.AccessControlGroupConfigurationNoList = []string{"4964", "30067"}

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.CreateLaunchConfiguration(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.LaunchConfigurationList)).To(Equal(1))

			lc := result.LaunchConfigurationList[0]
			Expect(lc.LaunchConfigurationName).To(Equal(reqParams.LaunchConfigurationName))
			Expect(lc.ServerImageProductCode).To(Equal("SPSW0LINUX000045"))
			Expect(lc.ServerProductCode).To(Equal("SPSVRSTAND000003"))
			Expect(lc.MemberServerImageNo).To(Equal(""))
			Expect(lc.LoginKeyName).To(Equal("tf-key-test"))
			Expect(lc.UserData).To(Equal(""))
			Expect(lc.CreateDate).To(Equal("2018-07-10T13:25:07+0900"))

			Expect(len(lc.AccessControlGroupList)).To(Equal(2))

			acg := lc.AccessControlGroupList[0]
			Expect(acg.AccessControlGroupConfigurationNo).To(Equal("4964"))
			Expect(acg.AccessControlGroupName).To(Equal("ncloud-default-acg"))
			Expect(acg.AccessControlGroupDescription).To(Equal("Default AccessControlGroup"))
			Expect(acg.IsDefault).To(Equal(true))
			Expect(acg.CreateDate).To(Equal("2017-02-23T10:25:39+0900"))
		})
	})

	Describe("Authorize fail", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/autoscaling").
				Reply(http.StatusUnauthorized).BodyString(`<responseError>
					<returnCode>800</returnCode>
					<returnMessage>Invalid consumerKey</returnMessage>
					</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should be failed by authorization fail", func() {
			reqParams := new(RequestCreateLaunchConfiguration)
			reqParams.ServerImageProductCode = "SPSW0LINUX000045"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.CreateLaunchConfiguration(reqParams)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Invalid consumerKey"))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})

	Describe("Expired URL", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/autoscaling").
				Reply(http.StatusUnauthorized).BodyString(`<responseError>
					<returnCode>800</returnCode>
					<returnMessage>Expired url.</returnMessage>
					</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should be failed by authorization fail", func() {
			reqParams := new(RequestCreateLaunchConfiguration)
			reqParams.ServerImageProductCode = "SPSW0LINUX000045"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.CreateLaunchConfiguration(reqParams)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Expired url."))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})

	Describe("Check Arguments", func() {
		It("should be error : 'Required field is not specified. location : serverImageProductCode or memberServerImageNo'", func() {
			reqParams := new(RequestCreateLaunchConfiguration)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateLaunchConfiguration(reqParams)

			Expect(err.Error()).To(Equal("Required field is not specified. location : serverImageProductCode or memberServerImageNo"))
		})

		It("should be error : 'Only one field is required. location : serverImageProductCode or memberServerImageNo'", func() {
			reqParams := new(RequestCreateLaunchConfiguration)
			reqParams.ServerImageProductCode = "aaaa"
			reqParams.MemberServerImageNo = "aaaa"

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateLaunchConfiguration(reqParams)

			Expect(err.Error()).To(Equal("Only one field is required. location : serverImageProductCode or memberServerImageNo"))
		})

		It("should be error : 'Length of ServerImageProductCode should be min 1 or max 20'", func() {
			reqParams := new(RequestCreateLaunchConfiguration)
			reqParams.ServerImageProductCode = String(21)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateLaunchConfiguration(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of ServerImageProductCode to be in the range (0 - 20)"))
		})

		It("should be error : 'Length of ServerProductCode should be min or max 20'", func() {
			reqParams := new(RequestCreateLaunchConfiguration)
			reqParams.ServerImageProductCode = "aaaa"
			reqParams.ServerProductCode = String(21)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateLaunchConfiguration(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of ServerProductCode to be in the range (0 - 20)"))
		})

		It("should be error : 'Length of LoginKeyName should be min 3 or max 30'", func() {
			reqParams := new(RequestCreateLaunchConfiguration)
			reqParams.ServerImageProductCode = "aaaa"
			reqParams.LoginKeyName = String(2)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateLaunchConfiguration(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of LoginKeyName to be in the range (3 - 30)"))

			reqParams.LoginKeyName = String(31)
			_, err = conn.CreateLaunchConfiguration(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of LoginKeyName to be in the range (3 - 30)"))
		})

		It("should be error : 'expected length of UserData to be in the range (0 - 21847)'", func() {
			reqParams := new(RequestCreateLaunchConfiguration)
			reqParams.ServerImageProductCode = "aaaa"
			reqParams.UserData = String(21848)

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.CreateLaunchConfiguration(reqParams)

			Expect(err.Error()).To(ContainSubstring("expected length of UserData to be in the range (0 - 21847)"))
		})
	})
})
