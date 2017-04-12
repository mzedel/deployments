// Copyright 2016 Mender Software AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/mendersoftware/deployments/resources/deployments"
)

// DeviceDeploymentLogStorage is an autogenerated mock type for the DeviceDeploymentStorage type
type DeviceDeploymentLogStorage struct {
	mock.Mock
}

func (_m *DeviceDeploymentLogStorage) SaveDeviceDeploymentLog(log deployments.DeploymentLog) error {
	ret := _m.Called(log)

	return ret.Error(0)
}

func (_m *DeviceDeploymentLogStorage) GetDeviceDeploymentLog(deviceID, deploymentID string) (*deployments.DeploymentLog, error) {
	ret := _m.Called(deviceID, deploymentID)

	return ret.Get(0).(*deployments.DeploymentLog), ret.Error(1)
}
