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

	"github.com/mendersoftware/deployments/resources/images"
)

// ImageByNameAndDeviceTyper is an autogenerated mock type for the ImageByNameAndDeviceTyper type
type ImageByNameAndDeviceTyper struct {
	mock.Mock
}

// ImageByNameAndDeviceType provides a mock function with given fields: name, deviceType
func (_m *ImageByNameAndDeviceTyper) ImageByNameAndDeviceType(name string, deviceType string) (*images.SoftwareImage, error) {
	ret := _m.Called(name, deviceType)

	var r0 *images.SoftwareImage
	if rf, ok := ret.Get(0).(func(string, string) *images.SoftwareImage); ok {
		r0 = rf(name, deviceType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*images.SoftwareImage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, deviceType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
