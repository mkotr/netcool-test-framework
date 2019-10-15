module git02.ae.sda.corp.telstra.com/scm/wian/netcool-test-automation/nc-test-webapp

go 1.12

require (
	github.com/astaxie/beego v1.12.0
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337
	golang.org/x/crypto v0.0.0-20190927123631-a832865fa7ad // indirect
	golang.org/x/net v0.0.0-20190926025831-c00fd9afed17 // indirect
	golang.org/x/text v0.3.2 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
)

replace git02.ae.sda.corp.telstra.com/scm/wian/netcool-test-automation/nc-test-webapp/service => ./service

replace git02.ae.sda.corp.telstra.com/scm/wian/netcool-test-automation/nc-test-webapp/domain => ./domain
