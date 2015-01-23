package main

import (
	"fmt"
	"github.com/globalways/utils_go/location/baidumap"
)

func main() {

	var lat string = "30.552966442214"
	var lng string = "104.089695154970"

	// 从坐标到地址
	GEOToAddress, err := baidumap.GetAddressViaGEO(lat, lng)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("坐标到地址：", GEOToAddress)
		fmt.Println("坐标到地址 - 地址", GEOToAddress.Result.AddressComponent.String())
		fmt.Println("\n")
	}

	// 从地址到坐标
	address := "四川省成都市双流县中和镇中和大道2段红树湾小区"
	addressToGEO, err := baidumap.GetGeoViaAddress(address)
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println("从地址到坐标 - All", addressToGEO.Result.Location.String())
		fmt.Println("从地址到坐标 - Lat", addressToGEO.Result.Location.Lat)
		fmt.Println("从地址到坐标 - Lng", addressToGEO.Result.Location.Lng)
		fmt.Println("\n")
	}

	// 从IP到地址
	ipAddress := "171.217.46.110"
	IPToAddress, err := baidumap.GetAddressViaIP(ipAddress)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("从IP到地址：", IPToAddress)
		fmt.Println("从IP到地址 - 地址：", IPToAddress, IPToAddress.Content.Address_Detail)
		fmt.Println("\n")
	}
}
