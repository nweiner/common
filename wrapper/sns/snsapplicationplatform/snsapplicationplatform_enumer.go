// Code Generated By gen-enumer For "Enum Type: SNSApplicationPlatform" - DO NOT EDIT;

/*
 * Copyright 2020 Aldelo, LP
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package snsapplicationplatform

import (
	"fmt"
	"strconv"
)

// enum names constants
const (
	_SNSApplicationPlatformName_0 = "UNKNOWN"
	_SNSApplicationPlatformName_1 = "ADM_AmazonDeviceMessaging"
	_SNSApplicationPlatformName_2 = "APNS_ApplePushNotificationService"
	_SNSApplicationPlatformName_3 = "APNS_Sandbox"
	_SNSApplicationPlatformName_4 = "FCM_FirebaseCloudMessaging"
)

// var declares of enum indexes
var (
	_SNSApplicationPlatformIndex_0 = [...]uint8{0, 7}
	_SNSApplicationPlatformIndex_1 = [...]uint8{0, 25}
	_SNSApplicationPlatformIndex_2 = [...]uint8{0, 33}
	_SNSApplicationPlatformIndex_3 = [...]uint8{0, 12}
	_SNSApplicationPlatformIndex_4 = [...]uint8{0, 26}
)

func (i SNSApplicationPlatform) String() string {
	switch {
	case i == UNKNOWN:
		return _SNSApplicationPlatformName_0
	case i == ADM_AmazonDeviceMessaging:
		return _SNSApplicationPlatformName_1
	case i == APNS_ApplePushNotificationService:
		return _SNSApplicationPlatformName_2
	case i == APNS_Sandbox:
		return _SNSApplicationPlatformName_3
	case i == FCM_FirebaseCloudMessaging:
		return _SNSApplicationPlatformName_4
	default:
		return ""
	}
}

var _SNSApplicationPlatformValues = []SNSApplicationPlatform{
	0, // UNKNOWN
	1, // ADM_AmazonDeviceMessaging
	2, // APNS_ApplePushNotificationService
	3, // APNS_Sandbox
	4, // FCM_FirebaseCloudMessaging
}

var _SNSApplicationPlatformNameToValueMap = map[string]SNSApplicationPlatform{
	_SNSApplicationPlatformName_0[0:7]:  0, // UNKNOWN
	_SNSApplicationPlatformName_1[0:25]: 1, // ADM_AmazonDeviceMessaging
	_SNSApplicationPlatformName_2[0:33]: 2, // APNS_ApplePushNotificationService
	_SNSApplicationPlatformName_3[0:12]: 3, // APNS_Sandbox
	_SNSApplicationPlatformName_4[0:26]: 4, // FCM_FirebaseCloudMessaging
}

var _SNSApplicationPlatformValueToKeyMap = map[SNSApplicationPlatform]string{
	0: _SNSApplicationPlatformKey_0, // UNKNOWN
	1: _SNSApplicationPlatformKey_1, // ADM_AmazonDeviceMessaging
	2: _SNSApplicationPlatformKey_2, // APNS_ApplePushNotificationService
	3: _SNSApplicationPlatformKey_3, // APNS_Sandbox
	4: _SNSApplicationPlatformKey_4, // FCM_FirebaseCloudMessaging
}

var _SNSApplicationPlatformValueToCaptionMap = map[SNSApplicationPlatform]string{
	0: _SNSApplicationPlatformCaption_0, // UNKNOWN
	1: _SNSApplicationPlatformCaption_1, // ADM_AmazonDeviceMessaging
	2: _SNSApplicationPlatformCaption_2, // APNS_ApplePushNotificationService
	3: _SNSApplicationPlatformCaption_3, // APNS_Sandbox
	4: _SNSApplicationPlatformCaption_4, // FCM_FirebaseCloudMessaging
}

var _SNSApplicationPlatformValueToDescriptionMap = map[SNSApplicationPlatform]string{
	0: _SNSApplicationPlatformDescription_0, // UNKNOWN
	1: _SNSApplicationPlatformDescription_1, // ADM_AmazonDeviceMessaging
	2: _SNSApplicationPlatformDescription_2, // APNS_ApplePushNotificationService
	3: _SNSApplicationPlatformDescription_3, // APNS_Sandbox
	4: _SNSApplicationPlatformDescription_4, // FCM_FirebaseCloudMessaging
}

// Valid returns 'true' if the value is listed in the SNSApplicationPlatform enum map definition, 'false' otherwise
func (i SNSApplicationPlatform) Valid() bool {
	for _, v := range _SNSApplicationPlatformValues {
		if i == v {
			return true
		}
	}

	return false
}

// ParseByName retrieves a SNSApplicationPlatform enum value from the enum string name,
// throws an error if the param is not part of the enum
func (i SNSApplicationPlatform) ParseByName(s string) (SNSApplicationPlatform, error) {
	if val, ok := _SNSApplicationPlatformNameToValueMap[s]; ok {
		// parse ok
		return val, nil
	}

	// error
	return -1, fmt.Errorf("Enum Name of %s Not Expected In SNSApplicationPlatform Values List", s)
}

// ParseByKey retrieves a SNSApplicationPlatform enum value from the enum string key,
// throws an error if the param is not part of the enum
func (i SNSApplicationPlatform) ParseByKey(s string) (SNSApplicationPlatform, error) {
	for k, v := range _SNSApplicationPlatformValueToKeyMap {
		if v == s {
			// parse ok
			return k, nil
		}
	}

	// error
	return -1, fmt.Errorf("Enum Key of %s Not Expected In SNSApplicationPlatform Keys List", s)
}

// Key retrieves a SNSApplicationPlatform enum string key
func (i SNSApplicationPlatform) Key() string {
	if val, ok := _SNSApplicationPlatformValueToKeyMap[i]; ok {
		// found
		return val
	} else {
		// not found
		return ""
	}
}

// Caption retrieves a SNSApplicationPlatform enum string caption
func (i SNSApplicationPlatform) Caption() string {
	if val, ok := _SNSApplicationPlatformValueToCaptionMap[i]; ok {
		// found
		return val
	} else {
		// not found
		return ""
	}
}

// Description retrieves a SNSApplicationPlatform enum string description
func (i SNSApplicationPlatform) Description() string {
	if val, ok := _SNSApplicationPlatformValueToDescriptionMap[i]; ok {
		// found
		return val
	} else {
		// not found
		return ""
	}
}

// IntValue gets the intrinsic enum integer value
func (i SNSApplicationPlatform) IntValue() int {
	return int(i)
}

// IntString gets the intrinsic enum integer value represented in string format
func (i SNSApplicationPlatform) IntString() string {
	return strconv.Itoa(int(i))
}

// ValueSlice returns all values of the enum SNSApplicationPlatform in a slice
func (i SNSApplicationPlatform) ValueSlice() []SNSApplicationPlatform {
	return _SNSApplicationPlatformValues
}

// NameMap returns all names of the enum SNSApplicationPlatform in a K:name,V:SNSApplicationPlatform map
func (i SNSApplicationPlatform) NameMap() map[string]SNSApplicationPlatform {
	return _SNSApplicationPlatformNameToValueMap
}

// KeyMap returns all keys of the enum SNSApplicationPlatform in a K:SNSApplicationPlatform,V:key map
func (i SNSApplicationPlatform) KeyMap() map[SNSApplicationPlatform]string {
	return _SNSApplicationPlatformValueToKeyMap
}

// CaptionMap returns all captions of the enum SNSApplicationPlatform in a K:SNSApplicationPlatform,V:caption map
func (i SNSApplicationPlatform) CaptionMap() map[SNSApplicationPlatform]string {
	return _SNSApplicationPlatformValueToCaptionMap
}

// DescriptionMap returns all descriptions of the enum SNSApplicationPlatform in a K:SNSApplicationPlatform,V:description map
func (i SNSApplicationPlatform) DescriptionMap() map[SNSApplicationPlatform]string {
	return _SNSApplicationPlatformValueToDescriptionMap
}
