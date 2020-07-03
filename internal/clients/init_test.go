// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"net"
	"testing"

	bootstrapConfig "github.com/edgexfoundry/go-mod-bootstrap/config"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"github.com/pearceduan/device-sdk-go/internal/common"
)

func TestCheckServiceAvailableByPingWithTimeoutError(test *testing.T) {
	var clientConfig = map[string]bootstrapConfig.ClientInfo{
		common.ClientData: bootstrapConfig.ClientInfo{
			Host:     "www.google.com",
			Port:     81,
			Protocol: "http",
		},
	}
	var config = &common.ConfigurationStruct{Clients: clientConfig}
	common.CurrentConfig = config
	common.LoggingClient = logger.NewClient("test_service", false, "./device-simple.log", "DEBUG")

	err := checkServiceAvailableByPing(common.ClientData)

	if err, ok := err.(net.Error); ok && !err.Timeout() {
		test.Fatal("Should be timeout error")
	}
}
