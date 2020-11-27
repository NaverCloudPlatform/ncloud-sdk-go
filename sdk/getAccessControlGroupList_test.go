package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Access Control Group List", func() {
	Describe("Get all Access Control Group List", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getAccessControlGroupListResponse>
					<requestId>4c50ff5c-81d5-4073-a581-12a47d4ff714</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>2</totalRows>
					<accessControlGroupList>
					  <accessControlGroup>
						<accessControlGroupConfigurationNo>4964</accessControlGroupConfigurationNo>
						<accessControlGroupName>ncloud-default-acg</accessControlGroupName>
						<accessControlGroupDescription>Default AccessControlGroup</accessControlGroupDescription>
						<isDefault>true</isDefault>
						<createDate>2017-02-23T10:25:39+0900</createDate>
					  </accessControlGroup>
					  <accessControlGroup>
						<accessControlGroupConfigurationNo>4965</accessControlGroupConfigurationNo>
						<accessControlGroupName>ncloud-load-balancer</accessControlGroupName>
						<accessControlGroupDescription>Default Loadbalancer AccessControlGroup</accessControlGroupDescription>
						<isDefault>true</isDefault>
						<createDate>2017-02-23T10:25:39+0900</createDate>
					  </accessControlGroup>
					</accessControlGroupList>
				  </getAccessControlGroupListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get Acess Control Group list", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetAccessControlGroupList(nil)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(2))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.AccessControlGroup)).To(Equal(2))

			accessControlGroup := result.AccessControlGroup[0]
			Expect(accessControlGroup.AccessControlGroupConfigurationNo).To(Equal("4964"))
			Expect(accessControlGroup.AccessControlGroupName).To(Equal("ncloud-default-acg"))
			Expect(accessControlGroup.AccessControlGroupDescription).To(Equal("Default AccessControlGroup"))
			Expect(accessControlGroup.IsDefault).To(BeTrue())
			Expect(accessControlGroup.CreateDate).To(Equal("2017-02-23T10:25:39+0900"))

			accessControlGroup = result.AccessControlGroup[1]
			Expect(accessControlGroup.AccessControlGroupConfigurationNo).To(Equal("4965"))
			Expect(accessControlGroup.AccessControlGroupName).To(Equal("ncloud-load-balancer"))
			Expect(accessControlGroup.AccessControlGroupDescription).To(Equal("Default Loadbalancer AccessControlGroup"))
			Expect(accessControlGroup.IsDefault).To(BeTrue())
			Expect(accessControlGroup.CreateDate).To(Equal("2017-02-23T10:25:39+0900"))
		})
	})

	Describe("Get Access Control Group List which is defalut", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getAccessControlGroupListResponse>
					<requestId>509dac74-d836-4b4d-a420-01c0d99e3146</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<accessControlGroupList>
					  <accessControlGroup>
						<accessControlGroupConfigurationNo>4964</accessControlGroupConfigurationNo>
						<accessControlGroupName>ncloud-default-acg</accessControlGroupName>
						<accessControlGroupDescription>Default AccessControlGroup</accessControlGroupDescription>
						<isDefault>true</isDefault>
						<createDate>2017-02-23T10:25:39+0900</createDate>
					  </accessControlGroup>
					</accessControlGroupList>
				  </getAccessControlGroupListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get One Acess Control Group", func() {
			reqParams := new(RequestAccessControlGroupList)
			reqParams.IsDefault = "true"
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetAccessControlGroupList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.AccessControlGroup)).To(Equal(1))

			accessControlGroup := result.AccessControlGroup[0]
			Expect(accessControlGroup.AccessControlGroupConfigurationNo).To(Equal("4964"))
			Expect(accessControlGroup.AccessControlGroupName).To(Equal("ncloud-default-acg"))
			Expect(accessControlGroup.AccessControlGroupDescription).To(Equal("Default AccessControlGroup"))
			Expect(accessControlGroup.IsDefault).To(BeTrue())
			Expect(accessControlGroup.CreateDate).To(Equal("2017-02-23T10:25:39+0900"))
		})
	})

	Describe("There is no Access Control Group list", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getAccessControlGroupListResponse>
					<requestId>b14395d6-5e72-4e0e-9f54-94d6802f588b</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>0</totalRows>
					<accessControlGroupList/>
				  </getAccessControlGroupListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get No Acess Control Group list", func() {
			reqParams := new(RequestAccessControlGroupList)
			reqParams.AccessControlGroupConfigurationNoList = []string{"4966"}
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetAccessControlGroupList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(0))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.AccessControlGroup)).To(Equal(0))
		})
	})
})
