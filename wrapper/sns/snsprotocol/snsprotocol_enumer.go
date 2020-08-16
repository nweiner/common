// Code Generated By gen-enumer For "Enum Type: SNSProtocol" - DO NOT EDIT;

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

package snsprotocol

import (
	"fmt"
	"strconv"
)

// enum names constants
const (
	_SNSProtocolName_0 = "UNKNOWN"
	_SNSProtocolName_1 = "Http"
	_SNSProtocolName_2 = "Https"
	_SNSProtocolName_3 = "Email"
	_SNSProtocolName_4 = "EmailJson"
	_SNSProtocolName_5 = "Sms"
	_SNSProtocolName_6 = "Sqs"
	_SNSProtocolName_7 = "ApplicationEndpoint"
	_SNSProtocolName_8 = "Lambda"
)

// var declares of enum indexes
var (
	_SNSProtocolIndex_0 = [...]uint8{0, 7}
	_SNSProtocolIndex_1 = [...]uint8{0, 4}
	_SNSProtocolIndex_2 = [...]uint8{0, 5}
	_SNSProtocolIndex_3 = [...]uint8{0, 5}
	_SNSProtocolIndex_4 = [...]uint8{0, 9}
	_SNSProtocolIndex_5 = [...]uint8{0, 3}
	_SNSProtocolIndex_6 = [...]uint8{0, 3}
	_SNSProtocolIndex_7 = [...]uint8{0, 19}
	_SNSProtocolIndex_8 = [...]uint8{0, 6}
)

func (i SNSProtocol) String() string {
	switch {
	case i == UNKNOWN:
		return _SNSProtocolName_0
	case i == Http:
		return _SNSProtocolName_1
	case i == Https:
		return _SNSProtocolName_2
	case i == Email:
		return _SNSProtocolName_3
	case i == EmailJson:
		return _SNSProtocolName_4
	case i == Sms:
		return _SNSProtocolName_5
	case i == Sqs:
		return _SNSProtocolName_6
	case i == ApplicationEndpoint:
		return _SNSProtocolName_7
	case i == Lambda:
		return _SNSProtocolName_8
	default:
		return ""
	}
}

var _SNSProtocolValues = []SNSProtocol{
	0, // UNKNOWN
	1, // Http
	2, // Https
	3, // Email
	4, // EmailJson
	5, // Sms
	6, // Sqs
	7, // ApplicationEndpoint
	8, // Lambda
}

var _SNSProtocolNameToValueMap = map[string]SNSProtocol{
	_SNSProtocolName_0[0:7]:  0, // UNKNOWN
	_SNSProtocolName_1[0:4]:  1, // Http
	_SNSProtocolName_2[0:5]:  2, // Https
	_SNSProtocolName_3[0:5]:  3, // Email
	_SNSProtocolName_4[0:9]:  4, // EmailJson
	_SNSProtocolName_5[0:3]:  5, // Sms
	_SNSProtocolName_6[0:3]:  6, // Sqs
	_SNSProtocolName_7[0:19]: 7, // ApplicationEndpoint
	_SNSProtocolName_8[0:6]:  8, // Lambda
}

var _SNSProtocolValueToKeyMap = map[SNSProtocol]string{
	0: _SNSProtocolKey_0, // UNKNOWN
	1: _SNSProtocolKey_1, // Http
	2: _SNSProtocolKey_2, // Https
	3: _SNSProtocolKey_3, // Email
	4: _SNSProtocolKey_4, // EmailJson
	5: _SNSProtocolKey_5, // Sms
	6: _SNSProtocolKey_6, // Sqs
	7: _SNSProtocolKey_7, // ApplicationEndpoint
	8: _SNSProtocolKey_8, // Lambda
}

var _SNSProtocolValueToCaptionMap = map[SNSProtocol]string{
	0: _SNSProtocolCaption_0, // UNKNOWN
	1: _SNSProtocolCaption_1, // Http
	2: _SNSProtocolCaption_2, // Https
	3: _SNSProtocolCaption_3, // Email
	4: _SNSProtocolCaption_4, // EmailJson
	5: _SNSProtocolCaption_5, // Sms
	6: _SNSProtocolCaption_6, // Sqs
	7: _SNSProtocolCaption_7, // ApplicationEndpoint
	8: _SNSProtocolCaption_8, // Lambda
}

var _SNSProtocolValueToDescriptionMap = map[SNSProtocol]string{
	0: _SNSProtocolDescription_0, // UNKNOWN
	1: _SNSProtocolDescription_1, // Http
	2: _SNSProtocolDescription_2, // Https
	3: _SNSProtocolDescription_3, // Email
	4: _SNSProtocolDescription_4, // EmailJson
	5: _SNSProtocolDescription_5, // Sms
	6: _SNSProtocolDescription_6, // Sqs
	7: _SNSProtocolDescription_7, // ApplicationEndpoint
	8: _SNSProtocolDescription_8, // Lambda
}

// Valid returns 'true' if the value is listed in the SNSProtocol enum map definition, 'false' otherwise
func (i SNSProtocol) Valid() bool {
	for _, v := range _SNSProtocolValues {
		if i == v {
			return true
		}
	}

	return false
}

// ParseByName retrieves a SNSProtocol enum value from the enum string name,
// throws an error if the param is not part of the enum
func (i SNSProtocol) ParseByName(s string) (SNSProtocol, error) {
	if val, ok := _SNSProtocolNameToValueMap[s]; ok {
		// parse ok
		return val, nil
	}

	// error
	return -1, fmt.Errorf("Enum Name of %s Not Expected In SNSProtocol Values List", s)
}

// ParseByKey retrieves a SNSProtocol enum value from the enum string key,
// throws an error if the param is not part of the enum
func (i SNSProtocol) ParseByKey(s string) (SNSProtocol, error) {
	for k, v := range _SNSProtocolValueToKeyMap {
		if v == s {
			// parse ok
			return k, nil
		}
	}

	// error
	return -1, fmt.Errorf("Enum Key of %s Not Expected In SNSProtocol Keys List", s)
}

// Key retrieves a SNSProtocol enum string key
func (i SNSProtocol) Key() string {
	if val, ok := _SNSProtocolValueToKeyMap[i]; ok {
		// found
		return val
	} else {
		// not found
		return ""
	}
}

// Caption retrieves a SNSProtocol enum string caption
func (i SNSProtocol) Caption() string {
	if val, ok := _SNSProtocolValueToCaptionMap[i]; ok {
		// found
		return val
	} else {
		// not found
		return ""
	}
}

// Description retrieves a SNSProtocol enum string description
func (i SNSProtocol) Description() string {
	if val, ok := _SNSProtocolValueToDescriptionMap[i]; ok {
		// found
		return val
	} else {
		// not found
		return ""
	}
}

// IntValue gets the intrinsic enum integer value
func (i SNSProtocol) IntValue() int {
	return int(i)
}

// IntString gets the intrinsic enum integer value represented in string format
func (i SNSProtocol) IntString() string {
	return strconv.Itoa(int(i))
}

// ValueSlice returns all values of the enum SNSProtocol in a slice
func (i SNSProtocol) ValueSlice() []SNSProtocol {
	return _SNSProtocolValues
}

// NameMap returns all names of the enum SNSProtocol in a K:name,V:SNSProtocol map
func (i SNSProtocol) NameMap() map[string]SNSProtocol {
	return _SNSProtocolNameToValueMap
}

// KeyMap returns all keys of the enum SNSProtocol in a K:SNSProtocol,V:key map
func (i SNSProtocol) KeyMap() map[SNSProtocol]string {
	return _SNSProtocolValueToKeyMap
}

// CaptionMap returns all captions of the enum SNSProtocol in a K:SNSProtocol,V:caption map
func (i SNSProtocol) CaptionMap() map[SNSProtocol]string {
	return _SNSProtocolValueToCaptionMap
}

// DescriptionMap returns all descriptions of the enum SNSProtocol in a K:SNSProtocol,V:description map
func (i SNSProtocol) DescriptionMap() map[SNSProtocol]string {
	return _SNSProtocolValueToDescriptionMap
}
