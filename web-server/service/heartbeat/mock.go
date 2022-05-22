package heartbeat

import "common/models/endpoint"

/*
 这个是mock用的，后面这个要存mysql
*/

var openServiceMap = make(map[int64]*endpoint.EndpointModel)
var openNodeMap = make(map[int64]*endpoint.EndpointModel)

var ipServiceMap = make(map[string]int64)
var ipNodeMap = make(map[string]int64)

func SaveReport(endpoint *endpoint.EndpointModel) {
	switch endpoint.EndpointType {
	case "NODE":
		openNodeMap[endpoint.Id] = endpoint
		ipNodeMap[endpoint.Ip] = endpoint.Id
	case "SERVICE":
		openServiceMap[endpoint.Id] = endpoint
		ipServiceMap[endpoint.Ip] = endpoint.Id
	default:

	}
}

func GetIdByIp(ip string, endType string) int64 {
	switch endType {
	case "NODE":
		return ipNodeMap[ip]
	case "SERVICE":
		return ipServiceMap[ip]
	default:
		return -1
	}
}

func GetEndPointById(id int64, endType string) *endpoint.EndpointModel {
	switch endType {
	case "NODE":
		return openNodeMap[id]
	case "SERVICE":
		return openServiceMap[id]
	default:
		return nil
	}
}
