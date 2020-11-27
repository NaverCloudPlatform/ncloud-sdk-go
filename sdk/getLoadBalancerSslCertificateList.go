package sdk

import (
	"encoding/xml"
	"fmt"
	"net/http"

	common "github.com/NaverCloudPlatform/ncloud-sdk-go/common"
	request "github.com/NaverCloudPlatform/ncloud-sdk-go/request"
)

// GetLoadBalancerSslCertificateList get SSL Certificate
func (s *Conn) GetLoadBalancerSslCertificateList(certificateName string) (*SslCertificateList, error) {
	params := make(map[string]string)

	if certificateName != "" {
		params["certificateName"] = certificateName
	}

	bytes, resp, err := request.NewRequest(s.accessKey, s.secretKey, "GET", s.apiURL, "/loadbalancer/v2/getLoadBalancerSslCertificateList", params)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		responseError, err := common.ParseErrorResponse(bytes)
		if err != nil {
			return nil, err
		}

		respError := SslCertificateList{}
		respError.ReturnCode = responseError.ReturnCode
		respError.ReturnMessage = responseError.ReturnMessage

		return &respError, fmt.Errorf("%s %s - error code: %d , error message: %s", resp.Status, string(bytes), responseError.ReturnCode, responseError.ReturnMessage)
	}

	var SslCertificateList = SslCertificateList{}
	if err := xml.Unmarshal([]byte(bytes), &SslCertificateList); err != nil {
		return nil, err
	}

	return &SslCertificateList, nil
}
