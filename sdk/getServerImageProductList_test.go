package sdk_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server Image Product List", func() {
	Describe("Get all Server Image Product List", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getServerImageProductListResponse>
					<requestId>03a18f33-dc4b-4bc4-be14-fd886cbf2a8d</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<productList>
						<product>
							<productCode>SPSW0LINUX000046</productCode>
							<productName>centos-7.3-64</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>CentOS 7.3 (64-bit)</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<osInformation>CentOS 7.3 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000045</productCode>
							<productName>centos-7.2-64</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>CentOS 7.2(64bit)</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<osInformation>CentOS 7.2 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000044</productCode>
							<productName>centos-6.6-64</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>CentOS 6.6(64bit)</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<osInformation>CentOS 6.6 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000031</productCode>
							<productName>centos-6.3-64</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>CentOS 6.3(64bit)</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<osInformation>CentOS 6.3 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000032</productCode>
							<productName>centos-6.3-32</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>CentOS 6.3(32bit)</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX32</code>
								<codeName>Linux 32 Bit</codeName>
							</platformType>
							<osInformation>CentOS 6.3 (32-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000043</productCode>
							<productName>centos-5.11-64</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>CentOS 5.11(64bit)</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<osInformation>CentOS 5.11 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000029</productCode>
							<productName>ubuntu-16.04-64-server</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>Ubuntu Server 16.04 (64-bit) </productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>UBS64</code>
								<codeName>Ubuntu Server 64 Bit</codeName>
							</platformType>
							<osInformation>Ubuntu Server 16.04 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000028</productCode>
							<productName>ubuntu-14.04-64-server</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>Ubuntu Server 14.04 </productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>UBS64</code>
								<codeName>Ubuntu Server 64 Bit</codeName>
							</platformType>
							<osInformation>Ubuntu Server 14.04 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000026</productCode>
							<productName>ubuntu-12.04-64-server</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>Ubuntu 12.04 LTS Server (64bit)</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>UBS64</code>
								<codeName>Ubuntu Server 64 Bit</codeName>
							</platformType>
							<osInformation>Ubuntu 12.04 Server (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000027</productCode>
							<productName>ubuntu-12.04-64-desktop</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>Ubuntu 12.04 LTS Desktop (64bit)</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>UBD64</code>
								<codeName>Ubuntu Desktop 64 Bit</codeName>
							</platformType>
							<osInformation>Ubuntu 12.04 Desktop (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000047</productCode>
							<productName>mysql(5.6)-centos-6.6-64</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>CentOS 6.6 with MySQL 5.6</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<osInformation>CentOS 6.6 with MySQL 5.6 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000050</productCode>
							<productName>mysql(5.7)-centos-7.2-64</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>CentOS 7.2 with MySQL 5.7</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<osInformation>CentOS 7.2 with MySQL 5.7 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000048</productCode>
							<productName>mysql(5.7)-centos-6.6-64</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>CentOS 6.6 with MySQL 5.7</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<osInformation>CentOS 6.6 with MySQL 5.7 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000058</productCode>
							<productName>cubrid(9.2)-ubuntu-14.04-64-server</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>Ubuntu Server 14.04 with Cubrid 9.2</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>UBS64</code>
								<codeName>Ubuntu Server 64 Bit</codeName>
							</platformType>
							<osInformation>Ubuntu Server 14.04 with Cubrid 9.2 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000062</productCode>
							<productName>redis(3.2.8)-centos-7.3-64</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>CentOS 7.3 with Redis 3.2.8</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<osInformation>CentOS 7.3 with Redis 3.2.8 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000065</productCode>
							<productName>tensorflow-ubuntu-16.04-64-server</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>Ubuntu Server 16.04 with Tensorflow</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>UBS64</code>
								<codeName>Ubuntu Server 64 Bit</codeName>
							</platformType>
							<osInformation>Ubuntu Server 16.04 with Tensorflow (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000067</productCode>
							<productName>Jenkins 2.73-ubuntu-16.04-64-server</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>Ubuntu Server 16.04 with Jenkins 2.73</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>UBS64</code>
								<codeName>Ubuntu Server 64 Bit</codeName>
							</platformType>
							<osInformation>Ubuntu Server 16.04 with Jenkins 2.73 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000071</productCode>
							<productName>Jenkins 2.73-ubuntu-14.04-64-server</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>Ubuntu Server 14.04 with Jenkins 2.73</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>UBS64</code>
								<codeName>Ubuntu Server 64 Bit</codeName>
							</platformType>
							<osInformation>Ubuntu Server 14.04 with Jenkins 2.73 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000073</productCode>
							<productName>mariadb(10.2)-ubuntu-16.04-64-server</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>Ubuntu Server 16.04 with MariaDB 10.2</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>UBS64</code>
								<codeName>Ubuntu Server 64 Bit</codeName>
							</platformType>
							<osInformation>Ubuntu Server 16.04 with MariaDB 10.2 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000066</productCode>
							<productName>tensorflow-centos-7.3-64</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>CentOS 7.3 with Tensorflow</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<osInformation>CentOS 7.3 with Tensorflow (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000068</productCode>
							<productName>Jenkins 2.73-centos-7.3-64</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>CentOS 7.3 with Jenkins 2.73</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<osInformation>CentOS 7.3 with Jenkins 2.73 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000069</productCode>
							<productName>PostgreSQL 9.4-centos-7.3-64</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>CentOS 7.3 with PostgreSQL 9.4</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<osInformation>CentOS 7.3 with PostgreSQL 9.4 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000070</productCode>
							<productName>PostgreSQL 9.4-centos-6.6-64</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>CentOS 6.6 with PostgreSQL 9.4</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<osInformation>CentOS 6.6 with PostgreSQL 9.4 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000072</productCode>
							<productName>Jenkins 2.73-centos-6.6-64</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>CentOS 6.6 with Jenkins 2.73</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<osInformation>CentOS 6.6 with Jenkins 2.73 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0LINUX000074</productCode>
							<productName>mariadb(10.2)-centos-7.3-64</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>CentOS 7.3 with MariaDB 10.2</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<osInformation>CentOS 7.3 with MariaDB 10.2 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0WINNT000016</productCode>
							<productName>win-2016-64</productName>
							<productType>
								<code>WINNT</code>
								<codeName>Windows NT</codeName>
							</productType>
							<productDescription>Windows Server 2016 (64-bit)</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>WND64</code>
								<codeName>Windows 64 Bit</codeName>
							</platformType>
							<osInformation>Windows Server 2016 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0WINNT000015</productCode>
							<productName>win-2012-64-R2</productName>
							<productType>
								<code>WINNT</code>
								<codeName>Windows NT</codeName>
							</productType>
							<productDescription>Windows Server 2012(64bit) R2</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>WND64</code>
								<codeName>Windows 64 Bit</codeName>
							</platformType>
							<osInformation>Windows Server 2012 R2 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0WINNT000014</productCode>
							<productName>win-2008-64-R2</productName>
							<productType>
								<code>WINNT</code>
								<codeName>Windows NT</codeName>
							</productType>
							<productDescription>Windows Server 2008(64bit) R2 with SP1</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>WND64</code>
								<codeName>Windows 64 Bit</codeName>
							</platformType>
							<osInformation>Windows Server 2008 R2 with SP1 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0WINNT000035</productCode>
							<productName>mssql(2012std)-win-2012-64-R2</productName>
							<productType>
								<code>WINNT</code>
								<codeName>Windows NT</codeName>
							</productType>
							<productDescription>Windows 2012 Server R2 with MSSQL 2012 Standard</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>WND64</code>
								<codeName>Windows 64 Bit</codeName>
							</platformType>
							<osInformation>Windows Server 2012 R2 with MSSQL 2012 Standard (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
						<product>
							<productCode>SPSW0WINNT000034</productCode>
							<productName>mssql(2008std)-win-2008-64-R2</productName>
							<productType>
								<code>WINNT</code>
								<codeName>Windows NT</codeName>
							</productType>
							<productDescription>Windows 2008 Server R2 with MSSQL Standard</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>WND64</code>
								<codeName>Windows 64 Bit</codeName>
							</platformType>
							<osInformation>Windows Server 2008 R2 with MSSQL 2008 Standard (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
					</productList>
					<totalRows>30</totalRows>
				</getServerImageProductListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get all server image product list", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetServerImageProductList(nil)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(30))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.Product)).To(Equal(30))

			product := result.Product[0]
			Expect(product.ProductCode).To(Equal("SPSW0LINUX000046"))
			Expect(product.ProductName).To(Equal("centos-7.3-64"))
			Expect(product.ProductType.Code).To(Equal("LINUX"))
			Expect(product.ProductType.CodeName).To(Equal("Linux"))
			Expect(product.InfraResourceType.Code).To(Equal("SW"))
			Expect(product.InfraResourceType.CodeName).To(Equal("Software"))
			Expect(product.CPUCount).To(Equal(0))
			Expect(product.MemorySize).To(Equal(0))
			Expect(product.BaseBlockStorageSize).To(Equal(53687091200))
			Expect(product.PlatformType.Code).To(Equal("LNX64"))
			Expect(product.PlatformType.CodeName).To(Equal("Linux 64 Bit"))
			Expect(product.OsInformation).To(Equal("CentOS 7.3 (64-bit)"))
			Expect(product.AddBlockStroageSize).To(Equal(0))
		})
	})

	Describe("Get One Server Image Product List which product Code : SPSW0LINUX000046", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getServerImageProductListResponse>
					<requestId>94f76273-6410-46ed-809b-352ce50f544d</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<productList>
						<product>
							<productCode>SPSW0LINUX000046</productCode>
							<productName>centos-7.3-64</productName>
							<productType>
								<code>LINUX</code>
								<codeName>Linux</codeName>
							</productType>
							<productDescription>CentOS 7.3 (64-bit)</productDescription>
							<infraResourceType>
								<code>SW</code>
								<codeName>Software</codeName>
							</infraResourceType>
							<cpuCount>0</cpuCount>
							<memorySize>0</memorySize>
							<baseBlockStorageSize>53687091200</baseBlockStorageSize>
							<platformType>
								<code>LNX64</code>
								<codeName>Linux 64 Bit</codeName>
							</platformType>
							<osInformation>CentOS 7.3 (64-bit)</osInformation>
							<addBlockStroageSize>0</addBlockStroageSize>
						</product>
					</productList>
					<totalRows>1</totalRows>
				</getServerImageProductListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get one server image product list which product Code : SPSW0LINUX000046", func() {
			reqParams := new(RequestGetServerImageProductList)
			reqParams.ProductCode = "SPSW0LINUX000046"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetServerImageProductList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.Product)).To(Equal(1))

			product := result.Product[0]
			Expect(product.ProductCode).To(Equal("SPSW0LINUX000046"))
			Expect(product.ProductName).To(Equal("centos-7.3-64"))
			Expect(product.ProductType.Code).To(Equal("LINUX"))
			Expect(product.ProductType.CodeName).To(Equal("Linux"))
			Expect(product.InfraResourceType.Code).To(Equal("SW"))
			Expect(product.InfraResourceType.CodeName).To(Equal("Software"))
			Expect(product.CPUCount).To(Equal(0))
			Expect(product.MemorySize).To(Equal(0))
			Expect(product.BaseBlockStorageSize).To(Equal(53687091200))
			Expect(product.PlatformType.Code).To(Equal("LNX64"))
			Expect(product.PlatformType.CodeName).To(Equal("Linux 64 Bit"))
			Expect(product.OsInformation).To(Equal("CentOS 7.3 (64-bit)"))
			Expect(product.AddBlockStroageSize).To(Equal(0))
		})
	})

	Describe("There is no server image product list", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusOK).BodyString(`<getServerImageProductListResponse>
					<requestId>33630663-e742-4283-b289-bdcafa04a768</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<productList/>
					<totalRows>0</totalRows>
				</getServerImageProductListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get no server image product list", func() {
			reqParams := new(RequestGetServerImageProductList)
			reqParams.ProductCode = "SPSW0LINUX000030"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetServerImageProductList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(0))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.Product)).To(Equal(0))
		})
	})

	Describe("Authorize fail", func() {
		BeforeEach(func() {
			gock.New("https://ncloud.apigw.ntruss.com").
				Get("/server").
				Reply(http.StatusUnauthorized).BodyString(`<responseError>
					<returnCode>800</returnCode>
					<returnMessage>Expired url.</returnMessage>
					</responseError>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get server image list", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetServerImageProductList(nil)

			Expect(result.ReturnCode).To(Equal(800))
			Expect(result.ReturnMessage).To(Equal("Expired url."))
			Expect(err.Error()).To(ContainSubstring("401 Unauthorized"))
		})
	})
})
