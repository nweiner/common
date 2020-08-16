// Code Generated By gen-enumer For "Enum Type: SdOperationFilter" - DO NOT EDIT;

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

package sdoperationfilter

import (
	"fmt"
	"strconv"
)

// enum names constants
const (
	_SdOperationFilterName_0 = "UNKNOWN"
	_SdOperationFilterName_1 = "EQ_NameSpaceID"
	_SdOperationFilterName_2 = "EQ_ServiceID"
	_SdOperationFilterName_3 = "EQ_Status"
	_SdOperationFilterName_4 = "EQ_Type"
	_SdOperationFilterName_5 = "IN_Status"
	_SdOperationFilterName_6 = "IN_Type"
	_SdOperationFilterName_7 = "BETWEEN_UpdateDate"
)

// var declares of enum indexes
var (
	_SdOperationFilterIndex_0 = [...]uint8{0, 7}
	_SdOperationFilterIndex_1 = [...]uint8{0, 14}
	_SdOperationFilterIndex_2 = [...]uint8{0, 12}
	_SdOperationFilterIndex_3 = [...]uint8{0, 9}
	_SdOperationFilterIndex_4 = [...]uint8{0, 7}
	_SdOperationFilterIndex_5 = [...]uint8{0, 9}
	_SdOperationFilterIndex_6 = [...]uint8{0, 7}
	_SdOperationFilterIndex_7 = [...]uint8{0, 18}
)

func (i SdOperationFilter) String() string {
	switch {
	case i == UNKNOWN:
		return _SdOperationFilterName_0
	case i == EQ_NameSpaceID:
		return _SdOperationFilterName_1
	case i == EQ_ServiceID:
		return _SdOperationFilterName_2
	case i == EQ_Status:
		return _SdOperationFilterName_3
	case i == EQ_Type:
		return _SdOperationFilterName_4
	case i == IN_Status:
		return _SdOperationFilterName_5
	case i == IN_Type:
		return _SdOperationFilterName_6
	case i == BETWEEN_UpdateDate:
		return _SdOperationFilterName_7
	default:
		return ""
	}
}

var _SdOperationFilterValues = []SdOperationFilter{
	0, // UNKNOWN
	1, // EQ_NameSpaceID
	2, // EQ_ServiceID
	3, // EQ_Status
	4, // EQ_Type
	5, // IN_Status
	6, // IN_Type
	7, // BETWEEN_UpdateDate
}

var _SdOperationFilterNameToValueMap = map[string]SdOperationFilter{
	_SdOperationFilterName_0[0:7]:  0, // UNKNOWN
	_SdOperationFilterName_1[0:14]: 1, // EQ_NameSpaceID
	_SdOperationFilterName_2[0:12]: 2, // EQ_ServiceID
	_SdOperationFilterName_3[0:9]:  3, // EQ_Status
	_SdOperationFilterName_4[0:7]:  4, // EQ_Type
	_SdOperationFilterName_5[0:9]:  5, // IN_Status
	_SdOperationFilterName_6[0:7]:  6, // IN_Type
	_SdOperationFilterName_7[0:18]: 7, // BETWEEN_UpdateDate
}

var _SdOperationFilterValueToKeyMap = map[SdOperationFilter]string{
	0: _SdOperationFilterKey_0, // UNKNOWN
	1: _SdOperationFilterKey_1, // EQ_NameSpaceID
	2: _SdOperationFilterKey_2, // EQ_ServiceID
	3: _SdOperationFilterKey_3, // EQ_Status
	4: _SdOperationFilterKey_4, // EQ_Type
	5: _SdOperationFilterKey_5, // IN_Status
	6: _SdOperationFilterKey_6, // IN_Type
	7: _SdOperationFilterKey_7, // BETWEEN_UpdateDate
}

var _SdOperationFilterValueToCaptionMap = map[SdOperationFilter]string{
	0: _SdOperationFilterCaption_0, // UNKNOWN
	1: _SdOperationFilterCaption_1, // EQ_NameSpaceID
	2: _SdOperationFilterCaption_2, // EQ_ServiceID
	3: _SdOperationFilterCaption_3, // EQ_Status
	4: _SdOperationFilterCaption_4, // EQ_Type
	5: _SdOperationFilterCaption_5, // IN_Status
	6: _SdOperationFilterCaption_6, // IN_Type
	7: _SdOperationFilterCaption_7, // BETWEEN_UpdateDate
}

var _SdOperationFilterValueToDescriptionMap = map[SdOperationFilter]string{
	0: _SdOperationFilterDescription_0, // UNKNOWN
	1: _SdOperationFilterDescription_1, // EQ_NameSpaceID
	2: _SdOperationFilterDescription_2, // EQ_ServiceID
	3: _SdOperationFilterDescription_3, // EQ_Status
	4: _SdOperationFilterDescription_4, // EQ_Type
	5: _SdOperationFilterDescription_5, // IN_Status
	6: _SdOperationFilterDescription_6, // IN_Type
	7: _SdOperationFilterDescription_7, // BETWEEN_UpdateDate
}

// Valid returns 'true' if the value is listed in the SdOperationFilter enum map definition, 'false' otherwise
func (i SdOperationFilter) Valid() bool {
	for _, v := range _SdOperationFilterValues {
		if i == v {
			return true
		}
	}

	return false
}

// ParseByName retrieves a SdOperationFilter enum value from the enum string name,
// throws an error if the param is not part of the enum
func (i SdOperationFilter) ParseByName(s string) (SdOperationFilter, error) {
	if val, ok := _SdOperationFilterNameToValueMap[s]; ok {
		// parse ok
		return val, nil
	}

	// error
	return -1, fmt.Errorf("Enum Name of %s Not Expected In SdOperationFilter Values List", s)
}

// ParseByKey retrieves a SdOperationFilter enum value from the enum string key,
// throws an error if the param is not part of the enum
func (i SdOperationFilter) ParseByKey(s string) (SdOperationFilter, error) {
	for k, v := range _SdOperationFilterValueToKeyMap {
		if v == s {
			// parse ok
			return k, nil
		}
	}

	// error
	return -1, fmt.Errorf("Enum Key of %s Not Expected In SdOperationFilter Keys List", s)
}

// Key retrieves a SdOperationFilter enum string key
func (i SdOperationFilter) Key() string {
	if val, ok := _SdOperationFilterValueToKeyMap[i]; ok {
		// found
		return val
	} else {
		// not found
		return ""
	}
}

// Caption retrieves a SdOperationFilter enum string caption
func (i SdOperationFilter) Caption() string {
	if val, ok := _SdOperationFilterValueToCaptionMap[i]; ok {
		// found
		return val
	} else {
		// not found
		return ""
	}
}

// Description retrieves a SdOperationFilter enum string description
func (i SdOperationFilter) Description() string {
	if val, ok := _SdOperationFilterValueToDescriptionMap[i]; ok {
		// found
		return val
	} else {
		// not found
		return ""
	}
}

// IntValue gets the intrinsic enum integer value
func (i SdOperationFilter) IntValue() int {
	return int(i)
}

// IntString gets the intrinsic enum integer value represented in string format
func (i SdOperationFilter) IntString() string {
	return strconv.Itoa(int(i))
}

// ValueSlice returns all values of the enum SdOperationFilter in a slice
func (i SdOperationFilter) ValueSlice() []SdOperationFilter {
	return _SdOperationFilterValues
}

// NameMap returns all names of the enum SdOperationFilter in a K:name,V:SdOperationFilter map
func (i SdOperationFilter) NameMap() map[string]SdOperationFilter {
	return _SdOperationFilterNameToValueMap
}

// KeyMap returns all keys of the enum SdOperationFilter in a K:SdOperationFilter,V:key map
func (i SdOperationFilter) KeyMap() map[SdOperationFilter]string {
	return _SdOperationFilterValueToKeyMap
}

// CaptionMap returns all captions of the enum SdOperationFilter in a K:SdOperationFilter,V:caption map
func (i SdOperationFilter) CaptionMap() map[SdOperationFilter]string {
	return _SdOperationFilterValueToCaptionMap
}

// DescriptionMap returns all descriptions of the enum SdOperationFilter in a K:SdOperationFilter,V:description map
func (i SdOperationFilter) DescriptionMap() map[SdOperationFilter]string {
	return _SdOperationFilterValueToDescriptionMap
}
