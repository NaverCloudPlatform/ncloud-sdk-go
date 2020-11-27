package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Load Balancer SSL Certicate", func() {
	Describe("Get all Load Balancer SSL Certicate", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/loadbalancer").
				Reply(http.StatusOK).BodyString(`
					<getLoadBalancerSslCertificateListResponse>
					<requestId>601e4e6a-bdf5-4032-935e-728c6896a14c</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>2</totalRows>
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
						<sslCertificate>
							<certificateName>bbb</certificateName>
							<privateKey>-----BEGIN RSA PRIVATE KEY-----
				MIIEpAIBAAKCAQEAxXSlpetuPOKnDzvkPpkoAF3spAW8cfjp9YDHBFvDP+OtCKw9
				yinAlO61J3X18MJ0sKwvrb5L+gAkWsUrfYEzGAr8V5qZ6mP1JLflVlHJY7msqsjp
				rkeQQUx7HPP6qvmM3uZ7+mKGrEw+yKe0SrHojtqO9cBAY5Kv2qIMjUhPpcnb75q2
				BaqUOiTfsewzIh4wnHR9GSZGv3Pl/N+OCrGlsSmhe9ePBcbAtL9oYln+J7p8JGJS
				cMK6NtUJd48IYjgC1AeNetyZYrC2FnhQTCxkDK/yq2mMSO+l9WWmaR8uU4zxHqaj
				BycHcv1ywftpZOnBVdsp+Ik3PKzG54X7LyxXiQIDAQABAoIBAEOQY3H/uivZPmLH
				EpWc4IQnn2aMk+vHyX54/yBtqcS9yiKSlV4MpVoQyCnlgi9MypL9iB8CY4r663Wn
				y/bY87vBXpE3VH1QkLxstGux9qBKE1wo/VTmJeVCH0pL7bT9SQeohDmr5vsj58PP
				JrD8aWAgRxSuIRoxQj0kf/kECkTm0IVD0YqMe4b956aul5noet8wbQ+2xEjsnBCy
				Z4rI8JKMq+V2brgOn7XK4BLfL0YbIFnS3s4A5zu0xvi7aAK5gb7igz0zMVLp/Rhc
				u8HNnxsDkC7Fj8XeWEw4I4IGfS1y7Pz83LJfqM9kfdhjwn4h25KWOPxS6bOwpVqz
				wzZV1fUCgYEA6RTWBW8h15FZ7jlUG5JAzqQGxtKCFIFHZmkIiVfV4c/45XVfD3Kz
				5xg6RsQCUMF3sBTcLNZYuRcS7es2eL6aCXXTkfT4I5KUoMuSN5YoDvvRZLz+JZMy
				XmhYzzy/Uh7ZsOTc9Fdo/YXJIXD0Cl1P0nGiunOH2Qhz3eno4gJtJBcCgYEA2N8I
				sjRy9GW+Ha2Jt13RHPoXhQbFk39zZAn9B8Rx2z4FaAO//l1S8PDL417yivdLY1UW
				7Tu2MOGE3YgYTefAlN61Rtv/QjNjvYA2c4GjA2yNvVCKoPLV5sudhyKUZ6mQGLdu
				oiBa/jDFbtbWl7u9NiHyvEX6zbfmYnN+AKn0hV8CgYBXTCiEzITeWmBWaz5nPTXs
				r16iZQG3cFwvrTM3TaCb/Or59iXugUWETny1OICtgmizmHyGhpmgaVX7qlcyjiDf
				XjQpvJibqjDksJpJG4JRaluY4XhG1oTM+0QYCmaV+VwLdwySr5JxMgSM8+NTZnOZ
				HFqYfuDoltPez9cbn1EFbQKBgQCdu8olYsRhQUa/ayKI/XFEdBl7JWu6Va5limZA
				qf5tiXSBLIkNxm6200xXuQ0LScXJH3AnZ5ChiMUMIxoaP37wR/Ls8MF9MsdOYtw3
				sogPy3pjwRqy6SvuSxXt3Za2trsZXwDWZlYIHwzaCuPVRDTgFFzp1rQNv72OyZVR
				gktYXQKBgQCegh5CycXxo+l2AZiok8qEYythpSMoWgbYIXOl3ewDYPBXc7czwKm0
				tvFCSA63gNUikyzOqW0MT7cuvZvR1/7HfvqIB6ZKN2icWndTdHj7TOflaJr7TspN
				GTfhUTV7jTQ0dt9U1E+oxRkjqC2HFYlpewXP0rcQxhtK7p6kiaUDIw==
				-----END RSA PRIVATE KEY-----</privateKey>
							<publicKeyCertificate>-----BEGIN CERTIFICATE-----
				MIIDGDCCAgACCQDGDiYiQixnsTANBgkqhkiG9w0BAQsFADBOMQswCQYDVQQGEwJL
				UjEOMAwGA1UECAwFc2VvdWwxDjAMBgNVBAcMBXNlb3VsMQwwCgYDVQQKDANuYnAx
				ETAPBgNVBAsMCG5jbG91ZGV2MB4XDTE4MDYwODA5NTEyOVoXDTE4MDcwODA5NTEy
				OVowTjELMAkGA1UEBhMCS1IxDjAMBgNVBAgMBXNlb3VsMQ4wDAYDVQQHDAVzZW91
				bDEMMAoGA1UECgwDbmJwMREwDwYDVQQLDAhuY2xvdWRldjCCASIwDQYJKoZIhvcN
				AQEBBQADggEPADCCAQoCggEBAMV0paXrbjzipw875D6ZKABd7KQFvHH46fWAxwRb
				wz/jrQisPcopwJTutSd19fDCdLCsL62+S/oAJFrFK32BMxgK/Feamepj9SS35VZR
				yWO5rKrI6a5HkEFMexzz+qr5jN7me/pihqxMPsintEqx6I7ajvXAQGOSr9qiDI1I
				T6XJ2++atgWqlDok37HsMyIeMJx0fRkmRr9z5fzfjgqxpbEpoXvXjwXGwLS/aGJZ
				/ie6fCRiUnDCujbVCXePCGI4AtQHjXrcmWKwthZ4UEwsZAyv8qtpjEjvpfVlpmkf
				LlOM8R6mowcnB3L9csH7aWTpwVXbKfiJNzysxueF+y8sV4kCAwEAATANBgkqhkiG
				9w0BAQsFAAOCAQEABDFl73C8ta9zYfyQXIbtv2tXt6oIhphjD5sV5KO6lgVSw7db
				XoiDlpQb5/LIXVYEwf8GLlSPORLsU36DQk0EyE6veTDYq4Cexkp5U1ca+jAGSum3
				Rtl02Dj6w7pz44sFnZc1IwKWUI7nTc0rQoloKRVyWnb5EPoNEe5QI1R5HJh7vV2M
				OUpdTkOueLjy+4BfUAt+LqeNz5u9WhvVgkFul1V9e3UXOaf0KvMYaRUYRj+IThaK
				7/UsUgXKmbhVOZrgNPjQT1cTC1lSD+xfCeA5UT/6xqxq91LL6cSNBOqEoDwyWGt5
				Kp6GlcmqTBU7RXAwPE4MeQ20yMmdipoPMb8nug==
				-----END CERTIFICATE-----</publicKeyCertificate>
							<certificateChain></certificateChain>
						</sslCertificate>
					</sslCertificateList>
				</getLoadBalancerSslCertificateListResponse>
					`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get Load Balancer SSL Certicate", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetLoadBalancerSslCertificateList("")

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(2))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.SslCertificateList)).To(Equal(2))

			sslCertificate := result.SslCertificateList[0]
			Expect(sslCertificate.CertificateName).To(Equal("aaa"))
			Expect(sslCertificate.PrivateKey).To(ContainSubstring("BEGIN RSA PRIVATE KEY"))
			Expect(sslCertificate.PublicKeyCertificate).To(ContainSubstring("BEGIN CERTIFICATE"))
		})
	})

	Describe("Get One Load Balancer SSL Certicate which name is aaa", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/loadbalancer").
				Reply(http.StatusOK).BodyString(`
					<getLoadBalancerSslCertificateListResponse>
					<requestId>601e4e6a-bdf5-4032-935e-728c6896a14c</requestId>
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
				</getLoadBalancerSslCertificateListResponse>
					`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get One Load Balancer SSL Certificate", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetLoadBalancerSslCertificateList("aaa")

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.SslCertificateList)).To(Equal(1))

			sslCertificate := result.SslCertificateList[0]
			Expect(sslCertificate.CertificateName).To(Equal("aaa"))
			Expect(sslCertificate.PrivateKey).To(ContainSubstring("BEGIN RSA PRIVATE KEY"))
			Expect(sslCertificate.PublicKeyCertificate).To(ContainSubstring("BEGIN CERTIFICATE"))
		})
	})

	Describe("There is no Load Balancer SSL Certificate", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/loadbalancer").
				Reply(http.StatusOK).BodyString(`
							<getLoadBalancerSslCertificateListResponse>
		<requestId>08ef7f0f-76af-45d7-80d3-6d98f44cf8d3</requestId>
		<returnCode>0</returnCode>
		<returnMessage>success</returnMessage>
		<totalRows>0</totalRows>
		<sslCertificateList/>
		</getLoadBalancerSslCertificateListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get No Load Balancer SSL Certificate", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetLoadBalancerSslCertificateList("333333")

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(0))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.SslCertificateList)).To(Equal(0))
		})
	})

	Describe("Authorize fail", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/loadbalancer").
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
			result, err := conn.GetLoadBalancerSslCertificateList("aaa")

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Invalid consumerKey"))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})

	Describe("Expired URL", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/loadbalancer").
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
			result, err := conn.GetLoadBalancerSslCertificateList("aaa")

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Expired url."))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})
})
