module github.com/Syleron/PulseHA-NetCore

go 1.12

replace (
	github.com/Sirupsen/logrus v1.0.5 => github.com/sirupsen/logrus v1.0.5
	github.com/Sirupsen/logrus v1.3.0 => github.com/Sirupsen/logrus v1.0.6
	github.com/Sirupsen/logrus v1.4.0 => github.com/sirupsen/logrus v1.0.6
)

require (
	github.com/Syleron/PulseHA v0.0.4-lw.0.20191001142325-46f87d885048
	github.com/vishvananda/netlink v1.0.1-0.20190930145447-2ec5bdc52b86
)
