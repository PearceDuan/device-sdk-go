// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2017-2018 Canonical Ltd
// Copyright (C) 2018-2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package service

import (
	contract "github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/pearceduan/device-sdk-go/internal/cache"
)

func (s *Service) GetDeviceCacheForName(name string) (contract.Device, bool) {
	return cache.Devices().ForName(name)
}

func (s *Service) GetDeviceCacheForId(id string) (contract.Device, bool) {
	return cache.Devices().ForId(id)
}

func (s *Service) GetDeviceCacheAll(name string) []contract.Device {
	return cache.Devices().All()
}

func (s *Service) AddDeviceCache(device contract.Device) (err error) {
	return cache.Devices().Add(device)
}

func (s *Service) RemoveDeviceCache(id string) error {
	return cache.Devices().Remove(id)
}

func (s *Service) RemoveDeviceCacheByName(name string) error {
	return cache.Devices().RemoveByName(name)
}

func (s *Service) UpdateDeviceCache(device contract.Device) error {
	return cache.Devices().Update(device)
}

func (s *Service) UpdateAdminState(id string, state contract.AdminState) error {
	return cache.Devices().UpdateAdminState(id, state)
}