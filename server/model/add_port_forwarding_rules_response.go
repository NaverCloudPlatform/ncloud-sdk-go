/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-07-02T09:06:17Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AddPortForwardingRulesResponse struct {

	// 포트포워딩설정번호
	PortForwardingConfigurationNo string `json:"portForwardingConfigurationNo,omitempty"`

	// 포트포워딩공인IP
	PortForwardingPublicIp string `json:"portForwardingPublicIp,omitempty"`

	// ZONE
	Zone *Zone `json:"zone,omitempty"`

	// 인터넷회선구분
	InternetLineType *CommonCode `json:"internetLineType,omitempty"`

	TotalRows int32 `json:"totalRows,omitempty"`

	PortForwardingRuleList []PortForwardingRule `json:"portForwardingRuleList,omitempty"`
}
