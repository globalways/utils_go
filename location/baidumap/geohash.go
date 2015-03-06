// Copyright 2015 mint.zhao.chiu@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
package baidumap

import "github.com/gansidui/geohash"

func GEOHash(addr string) (hash string, lat, lng float64) {
	loc, err := GetGeoViaAddress(addr)
	if err != nil {
		return
	}

	lat = loc.Result.Location.Lat
	lng = loc.Result.Location.Lng
	hash, _ = geohash.Encode(lat, lng, 10)

	return
}
