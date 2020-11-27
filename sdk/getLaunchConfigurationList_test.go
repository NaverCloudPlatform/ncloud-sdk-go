package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Launch Configuration List", func() {
	Describe("Get all Launch Configuration List", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/autoscaling").
				Reply(http.StatusOK).BodyString(`
				<getLaunchConfigurationListResponse>
				<requestId>d8c2690e-ae18-4a76-baa5-6c4c54901eaf</requestId>
				<returnCode>0</returnCode>
				<returnMessage>success</returnMessage>
				<totalRows>1</totalRows>
				<launchConfigurationList>
					<launchConfiguration>
						<launchConfigurationName>wefwfwf</launchConfigurationName>
						<serverImageProductCode>SPSW0LINUX000046</serverImageProductCode>
						<serverProductCode>SPSVRSSD00000003</serverProductCode>
						<memberServerImageNo></memberServerImageNo>
						<loginKeyName>ckey</loginKeyName>
						<createDate>2018-07-10T11:01:58+0900</createDate>
						<userData></userData>
						<accessControlGroupList>
							<accessControlGroup>
								<accessControlGroupConfigurationNo>4964</accessControlGroupConfigurationNo>
								<accessControlGroupName>ncloud-default-acg</accessControlGroupName>
								<accessControlGroupDescription>Default AccessControlGroup</accessControlGroupDescription>
								<isDefault>true</isDefault>
								<createDate>2017-02-23T10:25:39+0900</createDate>
							</accessControlGroup>
						</accessControlGroupList>
					</launchConfiguration>
				</launchConfigurationList>
			</getLaunchConfigurationListResponse>
					`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get Launch Configuration list", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetLaunchConfigurationList(nil)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.LaunchConfigurationList)).To(Equal(1))

			lc := result.LaunchConfigurationList[0]
			Expect(lc.LaunchConfigurationName).To(Equal("wefwfwf"))
			Expect(lc.ServerImageProductCode).To(Equal("SPSW0LINUX000046"))
			Expect(lc.ServerProductCode).To(Equal("SPSVRSSD00000003"))
			Expect(lc.MemberServerImageNo).To(Equal(""))
			Expect(lc.LoginKeyName).To(Equal("ckey"))
			Expect(lc.UserData).To(Equal(""))
			Expect(lc.CreateDate).To(Equal("2018-07-10T11:01:58+0900"))

			Expect(len(lc.AccessControlGroupList)).To(Equal(1))

			acg := lc.AccessControlGroupList[0]
			Expect(acg.AccessControlGroupConfigurationNo).To(Equal("4964"))
			Expect(acg.AccessControlGroupName).To(Equal("ncloud-default-acg"))
			Expect(acg.AccessControlGroupDescription).To(Equal("Default AccessControlGroup"))
			Expect(acg.IsDefault).To(Equal(true))
			Expect(acg.CreateDate).To(Equal("2017-02-23T10:25:39+0900"))
		})
	})

	Describe("Get One Launch Configuration List which LaunchConfigurationNameList.1 is wefwfwf", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/autoscaling").
				Reply(http.StatusOK).BodyString(`
				<getLaunchConfigurationListResponse>
				<requestId>936b76f0-b00e-4fbf-80a6-d4db3ff663a1</requestId>
				<returnCode>0</returnCode>
				<returnMessage>success</returnMessage>
				<totalRows>1</totalRows>
				<launchConfigurationList>
					<launchConfiguration>
						<launchConfigurationName>wefwfwf</launchConfigurationName>
						<serverImageProductCode>SPSW0LINUX000046</serverImageProductCode>
						<serverProductCode>SPSVRSSD00000003</serverProductCode>
						<memberServerImageNo></memberServerImageNo>
						<loginKeyName>ckey</loginKeyName>
						<createDate>2018-07-10T11:01:58+0900</createDate>
						<userData></userData>
						<accessControlGroupList>
							<accessControlGroup>
								<accessControlGroupConfigurationNo>4964</accessControlGroupConfigurationNo>
								<accessControlGroupName>ncloud-default-acg</accessControlGroupName>
								<accessControlGroupDescription>Default AccessControlGroup</accessControlGroupDescription>
								<isDefault>true</isDefault>
								<createDate>2017-02-23T10:25:39+0900</createDate>
							</accessControlGroup>
						</accessControlGroupList>
					</launchConfiguration>
				</launchConfigurationList>
			</getLaunchConfigurationListResponse>
					`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get One Launch Configuration", func() {
			reqParams := new(RequestGetLaunchConfigurationList)
			reqParams.LaunchConfigurationNameList = []string{"wefwfwf"}
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetLaunchConfigurationList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.LaunchConfigurationList)).To(Equal(1))

			lc := result.LaunchConfigurationList[0]
			Expect(lc.LaunchConfigurationName).To(Equal("wefwfwf"))
			Expect(lc.ServerImageProductCode).To(Equal("SPSW0LINUX000046"))
			Expect(lc.ServerProductCode).To(Equal("SPSVRSSD00000003"))
			Expect(lc.MemberServerImageNo).To(Equal(""))
			Expect(lc.LoginKeyName).To(Equal("ckey"))
			Expect(lc.UserData).To(Equal(""))
			Expect(lc.CreateDate).To(Equal("2018-07-10T11:01:58+0900"))

			Expect(len(lc.AccessControlGroupList)).To(Equal(1))

			acg := lc.AccessControlGroupList[0]
			Expect(acg.AccessControlGroupConfigurationNo).To(Equal("4964"))
			Expect(acg.AccessControlGroupName).To(Equal("ncloud-default-acg"))
			Expect(acg.AccessControlGroupDescription).To(Equal("Default AccessControlGroup"))
			Expect(acg.IsDefault).To(Equal(true))
			Expect(acg.CreateDate).To(Equal("2017-02-23T10:25:39+0900"))
		})
	})

	Describe("There is no Launch Configuration list", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/autoscaling").
				Reply(http.StatusOK).BodyString(`<getLaunchConfigurationListResponse>
					<requestId>d16f40b8-f2f1-4c2d-9d0c-ed214c6b9803</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>0</totalRows>
					<launchConfigurationList/>
				</getLaunchConfigurationListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get No Launch Configuration list", func() {
			reqParams := new(RequestGetLaunchConfigurationList)
			reqParams.LaunchConfigurationNameList = []string{"noname"}
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetLaunchConfigurationList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(0))
			Expect(len(result.LaunchConfigurationList)).To(Equal(0))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
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
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetLaunchConfigurationList(nil)

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
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetLaunchConfigurationList(nil)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Expired url."))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})

	Describe("Check Arguments", func() {
		It("should be error : 'Value of PageNo should be min 0 or max 2147483648'", func() {
			reqParams := new(RequestGetLaunchConfigurationList)
			reqParams.PageNo = -1

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.GetLaunchConfigurationList(reqParams)

			Expect(err.Error()).To(Equal("\"PageNo\" cannot be lower than 0: -1"))
		})

		It("should be error : 'Value of PageNo should be min 0 or max 2147483648'", func() {
			reqParams := new(RequestGetLaunchConfigurationList)
			reqParams.PageNo = 2147483648

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.GetLaunchConfigurationList(reqParams)

			Expect(err.Error()).To(Equal("\"PageNo\" cannot be higher than 2147483647: 2147483648"))
		})

		It("should be error : 'Value of PageSize should be min 0 or max 2147483648'", func() {
			reqParams := new(RequestGetLaunchConfigurationList)
			reqParams.PageSize = -1

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.GetLaunchConfigurationList(reqParams)

			Expect(err.Error()).To(Equal("\"PageSize\" cannot be lower than 0: -1"))
		})

		It("should be error : 'Value of PageSize should be min 0 or max 2147483648'", func() {
			reqParams := new(RequestGetLaunchConfigurationList)
			reqParams.PageSize = 2147483648

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.GetLaunchConfigurationList(reqParams)

			Expect(err.Error()).To(Equal("\"PageSize\" cannot be higher than 2147483647: 2147483648"))
		})

		It("should be error : 'SortedBy should be launchConfigurationName or createDate'", func() {
			reqParams := new(RequestGetLaunchConfigurationList)
			reqParams.SortedBy = "wrong sortedByValue"

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.GetLaunchConfigurationList(reqParams)

			Expect(err.Error()).To(Equal("SortedBy should be launchConfigurationName or createDate"))
		})

		It("should be error : 'SortingOrder should be ascending or descending'", func() {
			reqParams := new(RequestGetLaunchConfigurationList)
			reqParams.SortingOrder = "wrong SortingOrderValue"

			conn := NewConnection(accessKey, secretKey)
			_, err := conn.GetLaunchConfigurationList(reqParams)

			Expect(err.Error()).To(Equal("SortingOrder should be ascending or descending"))
		})
	})
})
