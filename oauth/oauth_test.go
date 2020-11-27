package oauth_test

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strconv"
	"time"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/oauth"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func encryptHmacSha256(message, secretKey string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(message))
	sha := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(sha)
}

var _ = Describe("Oauth", func() {
	Describe("Get method", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get server instance list", func() {
			const requestMethod = "GET"
			const requestURL = "https://ncloud.apigw.ntruss.com"
			const requestURN = "/server/v2/createServerInstances"

			c := NewConsumer(accessKey, secretKey, requestMethod, requestURL)

			c.AdditionalParams["serverImageProductCode"] = "SPSW0LINUX000046"
			c.AdditionalParams["serverProductCode"] = "SPSVRSSD00000003"
			c.AdditionalParams["serverCreateCount"] = "2"
			c.AdditionalParams["serverCreateStartNo"] = "001"

			now := time.Now()
			timestamp := strconv.FormatInt(int64(time.Nanosecond)*now.UnixNano()/int64(time.Millisecond), 10)

			reqURL, getParams, body, err := c.GetRequest(requestURL + requestURN)
			if err != nil {
				return nil, nil, err
			}

			reqURL += "?" + getParams
			requestURN += "?" + getParams

			reqUrl, _, err := c.GetRequest()
			Expect(err).To(BeNil())

			req, err := http.NewRequest("GET", reqUrl, nil)
			Expect(err).To(BeNil())

			message := "GET " + requestURN + "\n" + timestamp + "\n" + accessKey

			apigwSignature := encryptHmacSha256(message, secretKey)

			req.Header.Set("x-ncp-iam-access-key", accessKey)
			req.Header.Set("x-ncp-apigw-signature-v2", apigwSignature)
			req.Header.Set("x-ncp-apigw-timestamp", timestamp)

			client := &http.Client{}
			resp, err := client.Do(req)
			Expect(err).To(BeNil())

			defer resp.Body.Close()

			Expect(resp.StatusCode).To(Equal(200))
		})
	})

	// 2020-11-27 APIGW 에서 POST 방식이 사라져 주석처리
	//Describe("Post method", func() {
	//	BeforeEach(func() {
	//		gock.New("https://ncloud.apigw.ntruss.com").
	//			Post("/server/v2").
	//			Reply(http.StatusOK)
	//	})
	//	AfterEach(func() {
	//		gock.Off()
	//	})
	//	It("should create server instance list", func() {
	//		const requestMethod = "POST"
	//		const requestURL = "https://ncloud.apigw.ntruss.com/server/v2"
	//
	//		c := NewConsumer(accessKey, secretKey, requestMethod, requestURL)
	//
	//		c.AdditionalParams["action"] = "createLoginKey"
	//		c.AdditionalParams["keyName"] = "keyName"
	//
	//		reqUrl, body, err := c.GetRequest()
	//		Expect(err).To(BeNil())
	//
	//		req, err := http.NewRequest("POST", reqUrl, body)
	//		Expect(err).To(BeNil())
	//
	//		client := &http.Client{}
	//		resp, err := client.Do(req)
	//		Expect(err).To(BeNil())
	//
	//		defer resp.Body.Close()
	//
	//		Expect(resp.StatusCode).To(Equal(200))
	//	})
	//})
})
