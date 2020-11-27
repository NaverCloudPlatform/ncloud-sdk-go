package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Delete Load Balancer SSL Certicate", func() {
	Describe("Delete Load Balancer SSL Certicate", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/loadbalancer").
				Reply(http.StatusOK).BodyString(`
					<deleteLoadBalancerSslCertificateResponse>
					<requestId>bdee692c-ac4d-4ad4-a055-98ce20310855</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<sslCertificateList>
						<sslCertificate>
							<certificateName>aaa</certificateName>
							<privateKey>-----BEGIN RSA PRIVATE KEY-----
				MIIEowIBAAKCAQEAxTN/wk1YoFhzrrfSFeThKtOfjgqi+DiF6YAbHNxnnnDzi+Ty
				PC+QcT524x9xKfCbI3kkxGnMSgX4fVKHqEcXT952xDniRbOFLLfFGKnQnYKA3kDN
				EZAq6VJiE+ho4aRf1QiPgN9VGHTHcrBx077UpgWKba/YWB1YZ6cJr8QvPh685I93
				BKOU+IKu39DBCO5gbxDDFviroNkXzz3tzyXrB6kZCrE0r5MurXthv9MktGPHhAME
				c8BLVb+MQurhCpAyDQ/gxrbwP20RInCqZKrIkBqL8S+x78nWiJ1ZNWzev89DJXWV
				bVas/GC8b32u4jlcX9St5iwhedrNdMRMig7NMwIDAQABAoIBACcZWOFreIECSJ2B
				sNPKd6KIJwCAt2NDwblUPwvv31OYzZEVKbopLBhn7CaIG2XXYMsdv955o5mhqW0f
				qeoBfmvFjgLF/0kWmPcO0LNdKBGyyF5ItuAel4N/ZBbSY8kpUB7q/Zjtru+UZ98j
				gM4c7gik3Jd24AxIdTjH87G30i4cjPbO44gJXF1vtkCBrNQohAgdZsRPLbW+YA0D
				Z0B3OuYBRWelsiVQ62+8/46xuGyYkpXkqULrMxju6IjCocop1vQQzIE9vFFvrdvh
				Xzc5o9N/IHkZhach0iHALPi8SXsTWiziB9jaefeFN3P716P/fGoIkAI3VTdaOe8D
				gF4KIeECgYEA5JKHAVL7xChYxEeHO/nJpnubrIrxluigEGpC441GprYWU+1N5J3u
				60nCsvD+M2Nub9Ft8HWa09CX5cf1XlO6RAW5lPdhf+xFfpZ/mjZaCXTIAoKU8kNP
				XfNZKAAieCLtpFgKY/Xpr/I7Quyif//Oo2pE5NI+8VIvmfljK4k4rZkCgYEA3N1J
				Etp/qAP6FmgPQZDAj4qJ9EWx2jAz3wTFDZHZryebIo9l4MaE2ClbeD25u5CI7BRQ
				7Ys9G7WAg52IHhpp8bHqrQQMyedZO96Bc+/ib/aothxAstDHXELKThyjmsFaFuXD
				/iWKU83+kJ3wlhVhxRrRRad2rPgzCWAVmx8gmKsCgYBdQA5GEwXVzCQSx5+7bze+
				bPVg2jqft3391Gw/i5aEUwse6FR77ZbYdPoyqLD8ZBddFRMGI+Srf5FI1GAdQIlg
				UPmadaZYWxsq/tlGH+BxtSwYIhAEGkPZ73qzIALwAkfzYVuz1lKlTmyw6/uLvuAd
				uTAsLj4xk7cJ2T2FOHFgGQKBgQCnks2qnFVW6Q2KipFosacoQXjB+U8juIbiov/P
				d+Wt22L9Kcemb3jDeT2JUWvmP5djZ1avpIFM0L0dHxzzmh7f3pmrg/MgDYPKxNt7
				V3p6cK843N7JzVY1TgeolFodRK5RySqXXWxCgXyvxOWROc/geKRnMNP+EKLIbjp2
				Y3ub2QKBgGHd17CYLJJ9v/7sPFic3EsUxh0cwvTNSm9iNFMFYWjrpN/QzT5Jn33J
				BA8TU6/SZjhf8AkHkmUj+MGxEUYZGm/H2kXgSMCfFzCJ/P4GtcTq79z4h/d0029R
				hvptV0BoPrj29W3S7XTRyWjo5rAMZEgpxdkLIyj968G0yQO9Iq13
				-----END RSA PRIVATE KEY-----</privateKey>
							<publicKeyCertificate>-----BEGIN CERTIFICATE-----
				MIIDgjCCAmoCCQCwNwjms1rPLjANBgkqhkiG9w0BAQsFADCBgjELMAkGA1UEBhMC
				a3IxDjAMBgNVBAgMBXNlb3VsMQ4wDAYDVQQHDAVzZW91bDEMMAoGA1UECgwDbmJw
				MQwwCgYDVQQLDANkZXYxEzARBgNVBAMMCm5jbG91ZCBkZXYxIjAgBgkqhkiG9w0B
				CQEWE2JlYm9wamF6ekBnbWFpbC5jb20wHhcNMTgwNjA3MDMwMjE4WhcNMTgwNzA3
				MDMwMjE4WjCBgjELMAkGA1UEBhMCa3IxDjAMBgNVBAgMBXNlb3VsMQ4wDAYDVQQH
				DAVzZW91bDEMMAoGA1UECgwDbmJwMQwwCgYDVQQLDANkZXYxEzARBgNVBAMMCm5j
				bG91ZCBkZXYxIjAgBgkqhkiG9w0BCQEWE2JlYm9wamF6ekBnbWFpbC5jb20wggEi
				MA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDFM3/CTVigWHOut9IV5OEq05+O
				CqL4OIXpgBsc3GeecPOL5PI8L5BxPnbjH3Ep8JsjeSTEacxKBfh9UoeoRxdP3nbE
				OeJFs4Ust8UYqdCdgoDeQM0RkCrpUmIT6GjhpF/VCI+A31UYdMdysHHTvtSmBYpt
				r9hYHVhnpwmvxC8+Hrzkj3cEo5T4gq7f0MEI7mBvEMMW+Kug2RfPPe3PJesHqRkK
				sTSvky6te2G/0yS0Y8eEAwRzwEtVv4xC6uEKkDIND+DGtvA/bREicKpkqsiQGovx
				L7HvydaInVk1bN6/z0MldZVtVqz8YLxvfa7iOVxf1K3mLCF52s10xEyKDs0zAgMB
				AAEwDQYJKoZIhvcNAQELBQADggEBAI8wOj5iFfzcLRve523MnF4oNYDDFxdp9Y1K
				hPGNgXZ+MHPSAJylgLXkB59q+478aMFDU4SkWoGy9yLNHiSQ9KnjSPrydOP5m1x6
				mYWxbDqd9IRROuoWnWPCK3769roo7GHCBZywNvzD0SCLK1pGaCW4fTuI2kThoB7v
				sGhk0NC8gzFZasCSYrNJe6wYVJXRydzIXtYrj3H+bAYNu6mXXkDCNvlscfdE3jj+
				bwjTkXlzirWYBQx7ERI1j7E/RPN6RXG2T1zpzUnO/OOuHeYLHNDq0gnjF3/rz4sc
				8fiyW81ft/Zs31g/AoTBoq2NyfIlSHfRhGDGr0bpDz1jq73NHy0=
				-----END CERTIFICATE-----</publicKeyCertificate>
							<certificateChain></certificateChain>
						</sslCertificate>
					</sslCertificateList>
				</deleteLoadBalancerSslCertificateResponse>
					`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should delete Load Balancer SSL Certicate", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.DeleteLoadBalancerSslCertificate("bbbb")

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.SslCertificateList)).To(Equal(1))

			sslCertificate := result.SslCertificateList[0]
			Expect(sslCertificate.CertificateName).To(ContainSubstring("aaa"))
			Expect(sslCertificate.PrivateKey).To(ContainSubstring("-----BEGIN RSA PRIVATE KEY-----"))
			Expect(sslCertificate.PublicKeyCertificate).To(ContainSubstring("-----BEGIN CERTIFICATE-----"))
		})
	})

	Describe("Invalid Certificate Name", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/loadbalancer").
				Reply(http.StatusInternalServerError).BodyString(`
<responseError>
<returnCode>1</returnCode>
<returnMessage>Please try your call again later.
Temporarily out of service.
If error continue, Please contact our customer service center.</returnMessage>
</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should not Delete Load Balancer SSL Certificate by invalid Certificate Name", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.DeleteLoadBalancerSslCertificate("invalidCertificateName")

			Expect(result.ReturnCode).To(Equal(1))
			Expect(result.ReturnMessage).To(Equal(`Please try your call again later.
Temporarily out of service.
If error continue, Please contact our customer service center.`))
			Expect(err.Error()).To(ContainSubstring("500 Internal Server Error"))
		})
	})

	Describe("Authorize fail", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/loadbalancer").
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
			result, err := conn.DeleteLoadBalancerSslCertificate("aaa")

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Invalid consumerKey"))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})

	Describe("Expired URL", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/loadbalancer").
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
			result, err := conn.DeleteLoadBalancerSslCertificate("bbbb")

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Expired url."))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})
})
