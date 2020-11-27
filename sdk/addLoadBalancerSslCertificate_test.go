package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Add Load Balancer SSL Certicate", func() {
	Describe("Add Load Balancer SSL Certicate", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/loadbalancer").
				Reply(http.StatusOK).BodyString(`
<addLoadBalancerSslCertificateResponse>
<requestId>e45a3cc1-b732-41b6-83de-568e09e0a93c</requestId>
<returnCode>0</returnCode>
<returnMessage>success</returnMessage>
<totalRows>1</totalRows>
<sslCertificateList>
	<sslCertificate>
		<certificateName>test</certificateName>
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
</addLoadBalancerSslCertificateResponse>
					`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should add Load Balancer SSL Certicate", func() {
			conn := NewConnection(accessKey, secretKey)
			reqParams := &RequestAddSslCertificate{
				CertificateName: "test",
				PrivateKey: `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`,
				PublicKeyCertificate: `-----BEGIN CERTIFICATE-----
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
-----END CERTIFICATE-----`,
			}
			result, err := conn.AddLoadBalancerSslCertificate(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.SslCertificateList)).To(Equal(1))

			sslCertificate := result.SslCertificateList[0]
			Expect(sslCertificate.CertificateName).To(Equal(reqParams.CertificateName))
			Expect(sslCertificate.PrivateKey).To(Equal(reqParams.PrivateKey))
			Expect(sslCertificate.PublicKeyCertificate).To(Equal(reqParams.PublicKeyCertificate))
		})
	})

	Describe("Invalid SSL Certificate keys", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/loadbalancer").
				Reply(http.StatusBadRequest).BodyString(`
					<responseError>
					<returnCode>26040</returnCode>
					<returnMessage>The private key is not consistent.</returnMessage>
				</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should not add Load Balancer SSL Certificate by invalid keys", func() {
			conn := NewConnection(accessKey, secretKey)
			reqParams := &RequestAddSslCertificate{
				CertificateName: "test",
				PrivateKey: `-----BEGIN RSA PRIVATE KEY-----
GTfhUTV7jTQ0dt9U1E+oxRkjqC2HFYlpewXP0rcQxhtK7p6kiaUDIw==
-----END RSA PRIVATE KEY-----`,
				PublicKeyCertificate: `-----BEGIN CERTIFICATE-----
Kp6GlcmqTBU7RXAwPE4MeQ20yMmdipoPMb8nug==
-----END CERTIFICATE-----`,
			}
			result, err := conn.AddLoadBalancerSslCertificate(reqParams)

			Expect(result.ReturnCode).To(Equal(26040))
			Expect(result.ReturnMessage).To(Equal("The private key is not consistent."))
			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
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
			reqParams := &RequestAddSslCertificate{
				CertificateName: "test",
				PrivateKey: `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAxXSlpetuPOKnDzvkPpkoAF3spAW8cfjp9YDHBFvDP+OtCKw9
-----END RSA PRIVATE KEY-----`,
				PublicKeyCertificate: `-----BEGIN CERTIFICATE-----
MIIDGDCCAgACCQDGDiYiQixnsTANBgkqhkiG9w0BAQsFADBOMQswCQYDVQQGEwJL
-----END CERTIFICATE-----`,
			}
			result, err := conn.AddLoadBalancerSslCertificate(reqParams)

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
			reqParams := &RequestAddSslCertificate{
				CertificateName: "test",
				PrivateKey: `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`,
				PublicKeyCertificate: `-----BEGIN CERTIFICATE-----
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
-----END CERTIFICATE-----`,
			}
			result, err := conn.AddLoadBalancerSslCertificate(reqParams)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Expired url."))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})
})
