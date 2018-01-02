/*
   Pulse-NetCore - PulseHA Networking Plugin
   Copyright (C) 2017  Andrew Zak <andrew@pulseha.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
	"github.com/Syleron/PulseHA/src/netUtils"
	"errors"
	"github.com/Syleron/PulseHA/src/utils"
)

type PulseNetCore bool

const PluginName = "PulseHA-NetCore"
const PluginVersion = 1.0

/**
 * Return Plugin name
 **/
func (e PulseNetCore) Name() string {
	return PluginName
}

/**
 * Return Plugin version
 **/
func (e PulseNetCore) Version() float64 {
	return PluginVersion
}

func (e PulseNetCore) BringUpIPs(iface string, ips []string) error {
	for _, ip := range ips {
		success, err := netUtils.BringIPup(iface, ip)
		if !success && err != nil {
			return err
		} else if success && err != nil {
			return err
		}
		if utils.IsIPv6(ip) {
			go netUtils.IPv6NDP(iface)
		} else {
			go netUtils.SendGARP(iface, ip)
		}
	}
	return nil
}

func (e PulseNetCore) BringDownIPs(iface string, ips []string) error {
	for _, ip := range ips {
		success, err := netUtils.BringIPdown(iface, ip)
		if err != nil {
			return err
		}
		if !success {
			return errors.New("Failed to take down " + ip + " on interface " + iface)
		}
	}
	return nil
}

var PluginNet PulseNetCore

