package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ServerProductList", func() {
	Describe("Get Server Product List by Server Image Product Code", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getServerProductListResponse>
					<requestId>0e790dc1-f636-4791-922b-a11c9ad79de2</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<productList>
					  <product>
						<productCode>SPSVRSTAND000056</productCode>
						<productName>vCPU 1EA, Memory 1GB, Disk 50GB</productName>
						<productType>
						  <code>MICRO</code>
						  <codeName>Micro Server</codeName>
						</productType>
						<productDescription>vCPU 1EA, Memory 1GB, Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>1</cpuCount>
						<memorySize>1073741824</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000003</productCode>
						<productName>vCPU 1EA, Memory 2GB, Disk 50GB</productName>
						<productType>
						  <code>COMPT</code>
						  <codeName>Compact Server</codeName>
						</productType>
						<productDescription>vCPU 1개, 메모리 2GB, 디스크 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>1</cpuCount>
						<memorySize>2147483648</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000049</productCode>
						<productName>vCPU 2EA, Memory 2GB, Disk 50GB</productName>
						<productType>
						  <code>COMPT</code>
						  <codeName>Compact Server</codeName>
						</productType>
						<productDescription>vCPU 2EA, Memory 2GB, Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>2</cpuCount>
						<memorySize>2147483648</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSSD00000001</productCode>
						<productName>vCPU 1EA, Memory 2GB, [SSD]Disk 50GB</productName>
						<productType>
						  <code>COMPT</code>
						  <codeName>Compact Server</codeName>
						</productType>
						<productDescription>vCPU 1개, 메모리 2GB, [SSD]디스크 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>1</cpuCount>
						<memorySize>2147483648</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSSD00000002</productCode>
						<productName>vCPU 2EA, Memory 2GB, [SSD]Disk 50GB</productName>
						<productType>
						  <code>COMPT</code>
						  <codeName>Compact Server</codeName>
						</productType>
						<productDescription>vCPU 2EA, Memory 2GB, [SSD]Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>2</cpuCount>
						<memorySize>2147483648</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000004</productCode>
						<productName>vCPU 2EA, Memory 4GB, Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 2개, 메모리 4GB, 디스크 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>2</cpuCount>
						<memorySize>4294967296</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000024</productCode>
						<productName>vCPU 2EA, Memory 8GB, Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 2개, 메모리 8GB, 디스크 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>2</cpuCount>
						<memorySize>8589934592</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000052</productCode>
						<productName>vCPU 2EA, Memory 16GB, Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 2EA, Memory 16GB, Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>2</cpuCount>
						<memorySize>17179869184</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000050</productCode>
						<productName>vCPU 4EA, Memory 4GB, Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 4EA, Memory 4GB, Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>4</cpuCount>
						<memorySize>4294967296</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000005</productCode>
						<productName>vCPU 4EA, Memory 8GB, Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 4개, 메모리 8GB, 디스크 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>4</cpuCount>
						<memorySize>8589934592</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000025</productCode>
						<productName>vCPU 4EA, Memory 16GB, Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 4개, 메모리 16GB, 디스크 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>4</cpuCount>
						<memorySize>17179869184</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000054</productCode>
						<productName>vCPU 4EA, Memory 32GB, Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 4EA, Memory 32GB, Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>4</cpuCount>
						<memorySize>34359738368</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000051</productCode>
						<productName>vCPU 8EA, Memory 8GB, Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 8EA, Memory 8GB, Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>8</cpuCount>
						<memorySize>8589934592</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000006</productCode>
						<productName>vCPU 8EA, Memory 16GB, Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 8개, 메모리 16GB, 디스크 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>8</cpuCount>
						<memorySize>17179869184</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000055</productCode>
						<productName>vCPU 8EA, Memory 32GB, Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 8EA, Memory 32GB, Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>8</cpuCount>
						<memorySize>34359738368</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000030</productCode>
						<productName>vCPU 12EA, Memory 16GB, Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 12개, 메모리 16GB, 디스크 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>12</cpuCount>
						<memorySize>17179869184</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000022</productCode>
						<productName>vCPU 12EA, Memory 32GB, Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 12개, 메모리 32GB, 디스크 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>12</cpuCount>
						<memorySize>34359738368</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000053</productCode>
						<productName>vCPU 16EA, Memory 16GB, Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 16EA, Memory 16GB, Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>16</cpuCount>
						<memorySize>17179869184</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000046</productCode>
						<productName>vCPU 16EA, Memory 32GB, Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 16EA, Memory 32GB, Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>16</cpuCount>
						<memorySize>34359738368</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSSD00000003</productCode>
						<productName>vCPU 2EA, Memory 4GB, [SSD]Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 2개, 메모리 4GB, [SSD]디스크 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>2</cpuCount>
						<memorySize>4294967296</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSSD00000010</productCode>
						<productName>vCPU 2EA, Memory 8GB, [SSD]Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 2개, 메모리 8GB, [SSD]디스크 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>2</cpuCount>
						<memorySize>8589934592</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSSD00000011</productCode>
						<productName>vCPU 2EA, Memory 16GB, [SSD]Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 2EA, Memory 16GB, [SSD]Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>2</cpuCount>
						<memorySize>17179869184</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSSD00000004</productCode>
						<productName>vCPU 4EA, Memory 4GB, [SSD]Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 4EA, Memory 4GB, [SSD]Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>4</cpuCount>
						<memorySize>4294967296</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSSD00000005</productCode>
						<productName>vCPU 4EA, Memory 8GB, [SSD]Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 4개, 메모리 8GB, [SSD]디스크 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>4</cpuCount>
						<memorySize>8589934592</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSSD00000012</productCode>
						<productName>vCPU 4EA, Memory 16GB, [SSD]Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 4개, 메모리 16GB, [SSD]디스크 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>4</cpuCount>
						<memorySize>17179869184</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSSD00000013</productCode>
						<productName>vCPU 4EA, Memory 32GB, [SSD]Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 4EA, Memory 32GB, [SSD]Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>4</cpuCount>
						<memorySize>34359738368</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSSD00000006</productCode>
						<productName>vCPU 8EA, Memory 8GB, [SSD]Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 8EA, Memory 8GB, [SSD]Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>8</cpuCount>
						<memorySize>8589934592</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSSD00000007</productCode>
						<productName>vCPU 8EA, Memory 16GB, [SSD]Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 8개, 메모리 16GB, [SSD]디스크 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>8</cpuCount>
						<memorySize>17179869184</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSSD00000014</productCode>
						<productName>vCPU 8EA, Memory 32GB, [SSD]Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 8EA, Memory 32GB, [SSD]Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>8</cpuCount>
						<memorySize>34359738368</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSSD00000008</productCode>
						<productName>vCPU 12EA, Memory 16GB, [SSD]Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 12개, 메모리 16GB, [SSD]디스크 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>12</cpuCount>
						<memorySize>17179869184</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSSD00000015</productCode>
						<productName>vCPU 12EA, Memory 32GB, [SSD]Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 12개, 메모리 32GB, [SSD]디스크 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>12</cpuCount>
						<memorySize>34359738368</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSSD00000009</productCode>
						<productName>vCPU 16EA, Memory 16GB, [SSD]Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 16EA, Memory 16GB, [SSD]Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>16</cpuCount>
						<memorySize>17179869184</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSSD00000016</productCode>
						<productName>vCPU 16EA, Memory 32GB, [SSD]Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 16EA, Memory 32GB, [SSD]Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>16</cpuCount>
						<memorySize>34359738368</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000059</productCode>
						<productName>vCPU 8EA, Memory 64GB, Disk 50GB</productName>
						<productType>
						  <code>HIMEM</code>
						  <codeName>High Memory</codeName>
						</productType>
						<productDescription>vCPU 8EA, Memory 64GB, Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>8</cpuCount>
						<memorySize>68719476736</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000060</productCode>
						<productName>vCPU 16EA, Memory 64GB, Disk 50GB</productName>
						<productType>
						  <code>HIMEM</code>
						  <codeName>High Memory</codeName>
						</productType>
						<productDescription>vCPU 16EA, Memory 64GB, Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>16</cpuCount>
						<memorySize>68719476736</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000061</productCode>
						<productName>vCPU 16EA, Memory 128GB, Disk 50GB</productName>
						<productType>
						  <code>HIMEM</code>
						  <codeName>High Memory</codeName>
						</productType>
						<productDescription>vCPU 16EA, Memory 128GB, Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>16</cpuCount>
						<memorySize>137438953472</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000063</productCode>
						<productName>vCPU 8EA, Memory 64GB, [SSD]Disk 50GB</productName>
						<productType>
						  <code>HIMEM</code>
						  <codeName>High Memory</codeName>
						</productType>
						<productDescription>vCPU 8EA, Memory 64GB, [SSD]Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>8</cpuCount>
						<memorySize>68719476736</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000064</productCode>
						<productName>vCPU 16EA, Memory 64GB, [SSD]Disk 50GB</productName>
						<productType>
						  <code>HIMEM</code>
						  <codeName>High Memory</codeName>
						</productType>
						<productDescription>vCPU 16EA, Memory 128GB, [SSD]Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>16</cpuCount>
						<memorySize>68719476736</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000065</productCode>
						<productName>vCPU 16EA, Memory 128GB, [SSD]Disk 50GB</productName>
						<productType>
						  <code>HIMEM</code>
						  <codeName>High Memory</codeName>
						</productType>
						<productDescription>vCPU 16EA, Memory 128GB, [SSD]Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>16</cpuCount>
						<memorySize>137438953472</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000043</productCode>
						<productName>vCPU 2EA, Memory 4GB, Disk 100GB</productName>
						<productType>
						  <code>LDISK</code>
						  <codeName>Local Disk Server</codeName>
						</productType>
						<productDescription>vCPU 2개, 메모리 4GB, 디스크 100GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>2</cpuCount>
						<memorySize>4294967296</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>LOCAL</code>
						  <codeName>Local storage</codeName>
						</diskType>
						<addBlockStroageSize>53687091200</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000028</productCode>
						<productName>vCPU 4EA, Memory 8GB, Disk 200GB</productName>
						<productType>
						  <code>LDISK</code>
						  <codeName>Local Disk Server</codeName>
						</productType>
						<productDescription>vCPU 4개, 메모리 8GB, 디스크 200GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>4</cpuCount>
						<memorySize>8589934592</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>LOCAL</code>
						  <codeName>Local storage</codeName>
						</diskType>
						<addBlockStroageSize>161061273600</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000029</productCode>
						<productName>vCPU 8EA, Memory 16GB, Disk 400GB</productName>
						<productType>
						  <code>LDISK</code>
						  <codeName>Local Disk Server</codeName>
						</productType>
						<productDescription>vCPU 8개, 메모리 16GB, 디스크 400GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>8</cpuCount>
						<memorySize>17179869184</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>LOCAL</code>
						  <codeName>Local storage</codeName>
						</diskType>
						<addBlockStroageSize>375809638400</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSSD00000021</productCode>
						<productName>vCPU 20EA, Memory 80GB, [SSD]Disk 1000GB (Booting 50GB/Data 950GB)</productName>
						<productType>
						  <code>VDS</code>
						  <codeName>VDS Server</codeName>
						</productType>
						<productDescription>vCPU 20EA, Memory 80GB, [SSD]Disk 1000GB (Booting 50GB/Data 950GB)</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>20</cpuCount>
						<memorySize>85899345920</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>LOCAL</code>
						  <codeName>Local storage</codeName>
						</diskType>
						<addBlockStroageSize>1020054732800</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSSD00000023</productCode>
						<productName>vCPU 20EA, Memory 80GB, [SSD]Disk 2000GB (Booting 50GB/Data 1950GB)</productName>
						<productType>
						  <code>VDS</code>
						  <codeName>VDS Server</codeName>
						</productType>
						<productDescription>vCPU 20EA, Memory 80GB, [SSD]Disk 2000GB (Booting 50GB/Data 1950GB)</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>20</cpuCount>
						<memorySize>85899345920</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>LOCAL</code>
						  <codeName>Local storage</codeName>
						</diskType>
						<addBlockStroageSize>2093796556800</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSSD00000025</productCode>
						<productName>vCPU 32EA, Memory 122GB, [SSD]Disk 1000GB (Booting 50GB/Data 950GB)</productName>
						<productType>
						  <code>VDS</code>
						  <codeName>VDS Server</codeName>
						</productType>
						<productDescription>vCPU 32EA, Memory 122GB, [SSD]Disk 1000GB (Booting 50GB/Data 950GB)</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>32</cpuCount>
						<memorySize>130996502528</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>LOCAL</code>
						  <codeName>Local storage</codeName>
						</diskType>
						<addBlockStroageSize>1020054732800</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSSD00000027</productCode>
						<productName>vCPU 32EA, Memory 122GB, [SSD]Disk 2000GB (Booting 50GB/Data 1950GB)</productName>
						<productType>
						  <code>VDS</code>
						  <codeName>VDS Server</codeName>
						</productType>
						<productDescription>vCPU 32EA, Memory 122GB, [SSD]Disk 2000GB (Booting 50GB/Data 1950GB)</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>32</cpuCount>
						<memorySize>130996502528</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>LOCAL</code>
						  <codeName>Local storage</codeName>
						</diskType>
						<addBlockStroageSize>2093796556800</addBlockStroageSize>
					  </product>
					</productList>
					<totalRows>46</totalRows>
				  </getServerProductListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get server product list by server image product code", func() {
			conn := NewConnection(accessKey, secretKey)
			reqParams := new(RequestGetServerProductList)
			reqParams.ServerImageProductCode = "SPSW0LINUX000043"
			result, err := conn.GetServerProductList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(46))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.Product)).To(Equal(46))

			product := result.Product[0]
			Expect(product.ProductCode).To(Equal("SPSVRSTAND000056"))
			Expect(product.ProductName).To(Equal("vCPU 1EA, Memory 1GB, Disk 50GB"))
			Expect(product.InfraResourceType.Code).To(Equal("SVR"))
			Expect(product.InfraResourceType.CodeName).To(Equal("Server"))

			product = result.Product[1]
			Expect(product.ProductCode).To(Equal("SPSVRSTAND000003"))
			Expect(product.ProductName).To(Equal("vCPU 1EA, Memory 2GB, Disk 50GB"))
			Expect(product.InfraResourceType.Code).To(Equal("SVR"))
			Expect(product.InfraResourceType.CodeName).To(Equal("Server"))

		})
	})

	Describe("Get Server Product List by Server Image Product Code with a specific region no", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getServerProductListResponse>
					<requestId>ed5b199a-3f02-42f4-9b0e-4e3ab202cb10</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<productList>
					  <product>
						<productCode>SPSVRSTAND000004</productCode>
						<productName>vCPU 2EA, Memory 4GB, Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 2개, 메모리 4GB, 디스크 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>2</cpuCount>
						<memorySize>4294967296</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000050</productCode>
						<productName>vCPU 4EA, Memory 4GB, Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 4EA, Memory 4GB, Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>4</cpuCount>
						<memorySize>4294967296</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000046</productCode>
						<productName>vCPU 16EA, Memory 32GB, Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 16EA, Memory 32GB, Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>16</cpuCount>
						<memorySize>34359738368</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					</productList>
					<totalRows>3</totalRows>
				  </getServerProductListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get server product list by a specific region", func() {
			conn := NewConnection(accessKey, secretKey)
			reqParams := new(RequestGetServerProductList)
			reqParams.ServerImageProductCode = "SPSW0LINUX000043"
			reqParams.RegionNo = "2"
			result, err := conn.GetServerProductList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(3))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.Product)).To(Equal(3))

			product := result.Product[0]
			Expect(product.ProductCode).To(Equal("SPSVRSTAND000004"))
			Expect(product.ProductName).To(Equal("vCPU 2EA, Memory 4GB, Disk 50GB"))
			Expect(product.InfraResourceType.Code).To(Equal("SVR"))
			Expect(product.InfraResourceType.CodeName).To(Equal("Server"))

			product = result.Product[1]
			Expect(product.ProductCode).To(Equal("SPSVRSTAND000050"))
			Expect(product.ProductName).To(Equal("vCPU 4EA, Memory 4GB, Disk 50GB"))
			Expect(product.InfraResourceType.Code).To(Equal("SVR"))
			Expect(product.InfraResourceType.CodeName).To(Equal("Server"))

			product = result.Product[2]
			Expect(product.ProductCode).To(Equal("SPSVRSTAND000046"))
			Expect(product.ProductName).To(Equal("vCPU 16EA, Memory 32GB, Disk 50GB"))
			Expect(product.InfraResourceType.Code).To(Equal("SVR"))
			Expect(product.InfraResourceType.CodeName).To(Equal("Server"))
		})
	})

	Describe("Get Server Product List by Server Image Product Code with a specific region no and exclusive product code", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getServerProductListResponse>
					<requestId>bade0f6c-0bc2-40d9-af49-5fdced57ddb4</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<productList>
					  <product>
						<productCode>SPSVRSTAND000050</productCode>
						<productName>vCPU 4EA, Memory 4GB, Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 4EA, Memory 4GB, Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>4</cpuCount>
						<memorySize>4294967296</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					  <product>
						<productCode>SPSVRSTAND000046</productCode>
						<productName>vCPU 16EA, Memory 32GB, Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 16EA, Memory 32GB, Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>16</cpuCount>
						<memorySize>34359738368</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					</productList>
					<totalRows>2</totalRows>
				  </getServerProductListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get server product list by a specific region with exclusive product code", func() {
			conn := NewConnection(accessKey, secretKey)
			reqParams := new(RequestGetServerProductList)
			reqParams.ServerImageProductCode = "SPSW0LINUX000043"
			reqParams.ExclusionProductCode = "SPSVRSTAND000004"
			reqParams.RegionNo = "2"
			result, err := conn.GetServerProductList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(2))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.Product)).To(Equal(2))

			product := result.Product[0]
			Expect(product.ProductCode).To(Equal("SPSVRSTAND000050"))
			Expect(product.ProductName).To(Equal("vCPU 4EA, Memory 4GB, Disk 50GB"))
			Expect(product.InfraResourceType.Code).To(Equal("SVR"))
			Expect(product.InfraResourceType.CodeName).To(Equal("Server"))

			product = result.Product[1]
			Expect(product.ProductCode).To(Equal("SPSVRSTAND000046"))
			Expect(product.ProductName).To(Equal("vCPU 16EA, Memory 32GB, Disk 50GB"))
			Expect(product.InfraResourceType.Code).To(Equal("SVR"))
			Expect(product.InfraResourceType.CodeName).To(Equal("Server"))

		})
	})

	Describe("Get Server Product List by Server Image Product Code with a specific region no and a specific product code", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getServerProductListResponse>
					<requestId>79cdf702-de73-45ed-9b87-b2052acb0947</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<productList>
					  <product>
						<productCode>SPSVRSTAND000050</productCode>
						<productName>vCPU 4EA, Memory 4GB, Disk 50GB</productName>
						<productType>
						  <code>STAND</code>
						  <codeName>Standard</codeName>
						</productType>
						<productDescription>vCPU 4EA, Memory 4GB, Disk 50GB</productDescription>
						<infraResourceType>
						  <code>SVR</code>
						  <codeName>Server</codeName>
						</infraResourceType>
						<cpuCount>4</cpuCount>
						<memorySize>4294967296</memorySize>
						<baseBlockStorageSize>53687091200</baseBlockStorageSize>
						<osInformation></osInformation>
						<diskType>
						  <code>NET</code>
						  <codeName>Network Storage</codeName>
						</diskType>
						<addBlockStroageSize>0</addBlockStroageSize>
					  </product>
					</productList>
					<totalRows>1</totalRows>
				  </getServerProductListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get server product list by a specific region and a specific product code", func() {
			conn := NewConnection(accessKey, secretKey)
			reqParams := new(RequestGetServerProductList)
			reqParams.ServerImageProductCode = "SPSW0LINUX000043"
			reqParams.ProductCode = "SPSVRSTAND000050"
			reqParams.RegionNo = "2"
			result, err := conn.GetServerProductList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.Product)).To(Equal(1))

			product := result.Product[0]
			Expect(product.ProductCode).To(Equal("SPSVRSTAND000050"))
			Expect(product.ProductName).To(Equal("vCPU 4EA, Memory 4GB, Disk 50GB"))
			Expect(product.InfraResourceType.Code).To(Equal("SVR"))
			Expect(product.InfraResourceType.CodeName).To(Equal("Server"))

		})
	})
})
