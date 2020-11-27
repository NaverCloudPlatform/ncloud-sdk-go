package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create Login Key", func() {
	Describe("Create Login Key", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusOK).BodyString(`<createLoginKeyResponse>
					<requestId>ca6f393c-8a90-4023-ac40-93bcf77dd653</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<privateKey>
					-----BEGIN RSA PRIVATE KEY----- MIIEogIBAAKCAQEAnR2zOAQKvmp/wQ6+92OCkVmc4Jv/OHCdzYgHBrO/GbikWq2z 9tn4387eTq6oexourg/metEEGN3d1k57qC8P22xIYR5ncftccLR1WEZ20iQLVwKg l/lsKPx0ufaLtn9wPcTx9mewUwibflTCL4dvU8TacAmK/hanx6ny/i2D5VyB3bGc JV681mt6dtpg4U9MGOPLoAWAjm+FsJXirZqeCm2GJUltysPfFpqWT+Q6/3yFovhu 2X74FE56BgWKLFos40+a2Oq9Ez+OnDhqMVotadKJk5Wg3JllSmSCGaTglRRbszax +/6YLKahSgGtxMII+KuIc1OeRLyTxK66si7tewIDAQABAoIBADz1zReq2dBP2gm6 AvZRsdNbhbbQUBJB/7hDUxaSn08HiR110CGzKRENXswAkHZTDlGS8cYtWv2a9/uy r9/tgqWMOfkUpuZetNs16WG61VdrjMrRxiWxLPoqGdAdjM/eWlpdE9SbHprY+Bvo Hv/5+IP0GRznVvqHJ8occp+UYVyYEEjo+wk0mkcJSqSKG60Qs3t2/lajTdtsFhtf rtkRd9NcqukR8wpgrl9WP1USHdDQ88dBEx82aQARi5SA+5qD9zqQyihLHqmg2emX UzddhcKl+KoTB1Necxah65xyGc/Z4kpX0s1sJk+hspahDfoe1CQreRphm+QpL412 4fHFIYECgYEAzoc7rTGdCdTYpynjDf+88UlgtOonsRJuYnamKHFDbrAAp85usshP +emFZW7AdncCWYcGInxDffbEN1wHxj6KDdI3mh9p5zHdGW6NjkQmJuaMUj/9T8dh t2NWgbJ5odftziERZja9O6EGP0G1z8+SLsPmWDXwca0O2ZBPqakO37ECgYEAwsBm WWfR7vkAEAKefBYUqi6lgHwaFy3HiQkmmAnWjUF+KXd3BPu69xioy8RU6LTyxtOc qL3yYkARDK5LZ+pvQMm77+FwVcWu76godHHEORdFdTjRclJ9zk85eKCJiThCwti3 /XjVE54cUDrCS/Hqua13DPwAPy3/cHFGtXQsdusCgYBtG7akTiS3r/HxhkiE+cmY 83oaueXUP3d/n00y6nuEe4ktHfsPipHt5rfpmXQmYcMlBxdvmpSRdI++62jUC/Al Uxy62MO7N3WmXOh983TwcjZJsbKitiPinaFzDZYCsa/ZiVDG3j5kxTImsxFZQ0Do YehMS9rGQ3Yn7HyZlk3cUQKBgFamUiN/XNmNsihtceRV0rXxcTOcg+NCPCBNkP1e izZoKGD8xVjTQ9mDAD9BSGINeGNcbbxqXvi19qAEwIootX1HBLZzQI75GJ+K2w0M PTkFkKLaN948s7riZvizvN3vyUx2Lk65v1wuIfzUV/SAW2ZntqtxSA7UQW3fCI3d dZcnAoGACJWxbkMZPvHPfhes630LQ8JUBkI5vmhZ1waKMkbMEnb7TrqrIBrP+Enb 9cgxushdPpqI8KGOadpS8OFwqhYp3Oa7dYFA4+zIERHRqmBYk6KnTDELh/GzYt8G QG44e7ZqmB9ziny+gns9S+omP95EMHzgz+ZGaajGYQQvxd60BTE= -----END RSA PRIVATE KEY-----
					</privateKey>
					</createLoginKeyResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should create login key", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.CreateLoginKey("test01912")

			Expect(err).To(BeNil())
			Expect(result.RequestID).To(Equal("ca6f393c-8a90-4023-ac40-93bcf77dd653"))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(result.PrivateKey).To(Equal("-----BEGIN RSA PRIVATE KEY----- MIIEogIBAAKCAQEAnR2zOAQKvmp/wQ6+92OCkVmc4Jv/OHCdzYgHBrO/GbikWq2z 9tn4387eTq6oexourg/metEEGN3d1k57qC8P22xIYR5ncftccLR1WEZ20iQLVwKg l/lsKPx0ufaLtn9wPcTx9mewUwibflTCL4dvU8TacAmK/hanx6ny/i2D5VyB3bGc JV681mt6dtpg4U9MGOPLoAWAjm+FsJXirZqeCm2GJUltysPfFpqWT+Q6/3yFovhu 2X74FE56BgWKLFos40+a2Oq9Ez+OnDhqMVotadKJk5Wg3JllSmSCGaTglRRbszax +/6YLKahSgGtxMII+KuIc1OeRLyTxK66si7tewIDAQABAoIBADz1zReq2dBP2gm6 AvZRsdNbhbbQUBJB/7hDUxaSn08HiR110CGzKRENXswAkHZTDlGS8cYtWv2a9/uy r9/tgqWMOfkUpuZetNs16WG61VdrjMrRxiWxLPoqGdAdjM/eWlpdE9SbHprY+Bvo Hv/5+IP0GRznVvqHJ8occp+UYVyYEEjo+wk0mkcJSqSKG60Qs3t2/lajTdtsFhtf rtkRd9NcqukR8wpgrl9WP1USHdDQ88dBEx82aQARi5SA+5qD9zqQyihLHqmg2emX UzddhcKl+KoTB1Necxah65xyGc/Z4kpX0s1sJk+hspahDfoe1CQreRphm+QpL412 4fHFIYECgYEAzoc7rTGdCdTYpynjDf+88UlgtOonsRJuYnamKHFDbrAAp85usshP +emFZW7AdncCWYcGInxDffbEN1wHxj6KDdI3mh9p5zHdGW6NjkQmJuaMUj/9T8dh t2NWgbJ5odftziERZja9O6EGP0G1z8+SLsPmWDXwca0O2ZBPqakO37ECgYEAwsBm WWfR7vkAEAKefBYUqi6lgHwaFy3HiQkmmAnWjUF+KXd3BPu69xioy8RU6LTyxtOc qL3yYkARDK5LZ+pvQMm77+FwVcWu76godHHEORdFdTjRclJ9zk85eKCJiThCwti3 /XjVE54cUDrCS/Hqua13DPwAPy3/cHFGtXQsdusCgYBtG7akTiS3r/HxhkiE+cmY 83oaueXUP3d/n00y6nuEe4ktHfsPipHt5rfpmXQmYcMlBxdvmpSRdI++62jUC/Al Uxy62MO7N3WmXOh983TwcjZJsbKitiPinaFzDZYCsa/ZiVDG3j5kxTImsxFZQ0Do YehMS9rGQ3Yn7HyZlk3cUQKBgFamUiN/XNmNsihtceRV0rXxcTOcg+NCPCBNkP1e izZoKGD8xVjTQ9mDAD9BSGINeGNcbbxqXvi19qAEwIootX1HBLZzQI75GJ+K2w0M PTkFkKLaN948s7riZvizvN3vyUx2Lk65v1wuIfzUV/SAW2ZntqtxSA7UQW3fCI3d dZcnAoGACJWxbkMZPvHPfhes630LQ8JUBkI5vmhZ1waKMkbMEnb7TrqrIBrP+Enb 9cgxushdPpqI8KGOadpS8OFwqhYp3Oa7dYFA4+zIERHRqmBYk6KnTDELh/GzYt8G QG44e7ZqmB9ziny+gns9S+omP95EMHzgz+ZGaajGYQQvxd60BTE= -----END RSA PRIVATE KEY-----"))
		})
	})

	Describe("Exist login key name", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Post("/server").
				Reply(http.StatusBadRequest).BodyString(`<responseError>
					<returnCode>10405</returnCode>
					<returnMessage>The key name already exists.</returnMessage>
					</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should fail if key name is exist", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.CreateLoginKey("test01912")

			Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
			Expect(result.ReturnCode).To(Equal(10405))
			Expect(result.ReturnMessage).To(Equal("The key name already exists."))
		})
	})
})
