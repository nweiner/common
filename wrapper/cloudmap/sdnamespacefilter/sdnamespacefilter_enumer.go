// Code Generated By gen-enumer For "Enum Type: SdNamespaceFilter" - DO NOT EDIT;

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

package sdnamespacefilter

import (
	"fmt"
	"strconv"
)

// enum names constants
const (
	_SdNamespaceFilterName_0 = "UNKNOWN"
	_SdNamespaceFilterName_1 = "PublicDnsNamespace"
	_SdNamespaceFilterName_2 = "PrivateDnsNamespace"
	_SdNamespaceFilterName_3 = "Both"
)

// var declares of enum indexes
var (
	_SdNamespaceFilterIndex_0 = [...]uint8{0, 7}
	_SdNamespaceFilterIndex_1 = [...]uint8{0, 18}
	_SdNamespaceFilterIndex_2 = [...]uint8{0, 19}
	_SdNamespaceFilterIndex_3 = [...]uint8{0, 4}
)

func (i SdNamespaceFilter) String() string {
	switch {
	case i == UNKNOWN:
		return _SdNamespaceFilterName_0
	case i == PublicDnsNamespace:
		return _SdNamespaceFilterName_1
	case i == PrivateDnsNamespace:
		return _SdNamespaceFilterName_2
	case i == Both:
		return _SdNamespaceFilterName_3
	default:
		return ""
	}
}

var _SdNamespaceFilterValues = []SdNamespaceFilter{
	0, // UNKNOWN
	1, // PublicDnsNamespace
	2, // PrivateDnsNamespace
	3, // Both
}

var _SdNamespaceFilterNameToValueMap = map[string]SdNamespaceFilter{
	_SdNamespaceFilterName_0[0:7]:  0, // UNKNOWN
	_SdNamespaceFilterName_1[0:18]: 1, // PublicDnsNamespace
	_SdNamespaceFilterName_2[0:19]: 2, // PrivateDnsNamespace
	_SdNamespaceFilterName_3[0:4]:  3, // Both
}

var _SdNamespaceFilterValueToKeyMap = map[SdNamespaceFilter]string{
	0: _SdNamespaceFilterKey_0, // UNKNOWN
	1: _SdNamespaceFilterKey_1, // PublicDnsNamespace
	2: _SdNamespaceFilterKey_2, // PrivateDnsNamespace
	3: _SdNamespaceFilterKey_3, // Both
}

var _SdNamespaceFilterValueToCaptionMap = map[SdNamespaceFilter]string{
	0: _SdNamespaceFilterCaption_0, // UNKNOWN
	1: _SdNamespaceFilterCaption_1, // PublicDnsNamespace
	2: _SdNamespaceFilterCaption_2, // PrivateDnsNamespace
	3: _SdNamespaceFilterCaption_3, // Both
}

var _SdNamespaceFilterValueToDescriptionMap = map[SdNamespaceFilter]string{
	0: _SdNamespaceFilterDescription_0, // UNKNOWN
	1: _SdNamespaceFilterDescription_1, // PublicDnsNamespace
	2: _SdNamespaceFilterDescription_2, // PrivateDnsNamespace
	3: _SdNamespaceFilterDescription_3, // Both
}

// Valid returns 'true' if the value is listed in the SdNamespaceFilter enum map definition, 'false' otherwise
func (i SdNamespaceFilter) Valid() bool {
	for _, v := range _SdNamespaceFilterValues {
		if i == v {
			return true
		}
	}

	return false
}

// ParseByName retrieves a SdNamespaceFilter enum value from the enum string name,
// throws an error if the param is not part of the enum
func (i SdNamespaceFilter) ParseByName(s string) (SdNamespaceFilter, error) {
	if val, ok := _SdNamespaceFilterNameToValueMap[s]; ok {
		// parse ok
		return val, nil
	}

	// error
	return -1, fmt.Errorf("Enum Name of %s Not Expected In SdNamespaceFilter Values List", s)
}

// ParseByKey retrieves a SdNamespaceFilter enum value from the enum string key,
// throws an error if the param is not part of the enum
func (i SdNamespaceFilter) ParseByKey(s string) (SdNamespaceFilter, error) {
	for k, v := range _SdNamespaceFilterValueToKeyMap {
		if v == s {
			// parse ok
			return k, nil
		}
	}

	// error
	return -1, fmt.Errorf("Enum Key of %s Not Expected In SdNamespaceFilter Keys List", s)
}

// Key retrieves a SdNamespaceFilter enum string key
func (i SdNamespaceFilter) Key() string {
	if val, ok := _SdNamespaceFilterValueToKeyMap[i]; ok {
		// found
		return val
	} else {
		// not found
		return ""
	}
}

// Caption retrieves a SdNamespaceFilter enum string caption
func (i SdNamespaceFilter) Caption() string {
	if val, ok := _SdNamespaceFilterValueToCaptionMap[i]; ok {
		// found
		return val
	} else {
		// not found
		return ""
	}
}

// Description retrieves a SdNamespaceFilter enum string description
func (i SdNamespaceFilter) Description() string {
	if val, ok := _SdNamespaceFilterValueToDescriptionMap[i]; ok {
		// found
		return val
	} else {
		// not found
		return ""
	}
}

// IntValue gets the intrinsic enum integer value
func (i SdNamespaceFilter) IntValue() int {
	return int(i)
}

// IntString gets the intrinsic enum integer value represented in string format
func (i SdNamespaceFilter) IntString() string {
	return strconv.Itoa(int(i))
}

// ValueSlice returns all values of the enum SdNamespaceFilter in a slice
func (i SdNamespaceFilter) ValueSlice() []SdNamespaceFilter {
	return _SdNamespaceFilterValues
}

// NameMap returns all names of the enum SdNamespaceFilter in a K:name,V:SdNamespaceFilter map
func (i SdNamespaceFilter) NameMap() map[string]SdNamespaceFilter {
	return _SdNamespaceFilterNameToValueMap
}

// KeyMap returns all keys of the enum SdNamespaceFilter in a K:SdNamespaceFilter,V:key map
func (i SdNamespaceFilter) KeyMap() map[SdNamespaceFilter]string {
	return _SdNamespaceFilterValueToKeyMap
}

// CaptionMap returns all captions of the enum SdNamespaceFilter in a K:SdNamespaceFilter,V:caption map
func (i SdNamespaceFilter) CaptionMap() map[SdNamespaceFilter]string {
	return _SdNamespaceFilterValueToCaptionMap
}

// DescriptionMap returns all descriptions of the enum SdNamespaceFilter in a K:SdNamespaceFilter,V:description map
func (i SdNamespaceFilter) DescriptionMap() map[SdNamespaceFilter]string {
	return _SdNamespaceFilterValueToDescriptionMap
}