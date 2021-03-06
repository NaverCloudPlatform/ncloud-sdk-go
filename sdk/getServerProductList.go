package sdk

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"

	common "github.com/NaverCloudPlatform/ncloud-sdk-go/common"
	request "github.com/NaverCloudPlatform/ncloud-sdk-go/request"
)

func processGetServerProductListParams(reqParams *RequestGetServerProductList) (map[string]string, error) {
	params := make(map[string]string)

	if reqParams == nil || reqParams.ServerImageProductCode == "" {
		return params, errors.New("ServerImageProductCode is required field")
	}

	if len(reqParams.ServerImageProductCode) > 20 {
		return params, errors.New("Length of serverImageProductCode should be max 20")
	}

	params["serverImageProductCode"] = reqParams.ServerImageProductCode

	if reqParams.ExclusionProductCode != "" {
		if len(reqParams.ExclusionProductCode) > 20 {
			return params, errors.New("Length of exclusionProductCode should be max 20")
		}
		params["exclusionProductCode"] = reqParams.ExclusionProductCode
	}

	if reqParams.ProductCode != "" {
		if len(reqParams.ProductCode) > 20 {
			return params, errors.New("Length of productCode should be max 20")
		}
		params["productCode"] = reqParams.ProductCode
	}

	if reqParams.InternetLineTypeCode != "" {
		if err := validateIncludeValues("InternetLineTypeCode", reqParams.InternetLineTypeCode, []string{"PUBLC", "GLBL"}); err != nil {
			return nil, err
		}
		params["internetLineTypeCode"] = reqParams.InternetLineTypeCode
	}

	if reqParams.ZoneNo != "" {
		params["zoneNo"] = reqParams.ZoneNo
	}

	if reqParams.RegionNo != "" {
		params["regionNo"] = reqParams.RegionNo
	}

	return params, nil
}

// GetServerProductList : Get Server product list with server image product code by default.
func (s *Conn) GetServerProductList(reqParams *RequestGetServerProductList) (*ProductList, error) {
	params, err := processGetServerProductListParams(reqParams)
	if err != nil {
		return nil, err
	}

	bytes, resp, err := request.NewRequest(s.accessKey, s.secretKey, "GET", s.apiURL, "/server/v2/getServerProductList", params)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		responseError, err := common.ParseErrorResponse(bytes)
		if err != nil {
			return nil, err
		}

		respError := ProductList{}
		respError.ReturnCode = responseError.ReturnCode
		respError.ReturnMessage = responseError.ReturnMessage

		return &respError, fmt.Errorf("%s %s - error code: %d , error message: %s", resp.Status, string(bytes), responseError.ReturnCode, responseError.ReturnMessage)
	}

	var productListResp = ProductList{}
	if err := xml.Unmarshal([]byte(bytes), &productListResp); err != nil {
		return nil, err
	}

	return &productListResp, nil
}
