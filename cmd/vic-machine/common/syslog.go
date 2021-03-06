// Copyright 2019 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	"net/url"

	"gopkg.in/urfave/cli.v1"
)

// general syslog
type Syslog struct {
	SyslogAddr string
}

func (s *Syslog) SyslogFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "syslog-address",
			Value:       "",
			Usage:       "Address of the syslog server to send Virtual Container Host logs to. Must be in the format transport://host[:port], where transport is udp or tcp. port defaults to 514 if not specified",
			Destination: &s.SyslogAddr,
			Hidden:      true,
		},
	}
}

func (s *Syslog) ProcessSyslog() (*url.URL, error) {
	if len(s.SyslogAddr) == 0 {
		return nil, nil
	}

	u, err := url.Parse(s.SyslogAddr)
	if err != nil {
		return nil, err
	}

	return u, nil
}
