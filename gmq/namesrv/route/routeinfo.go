package route

import "fmt"

type RouteInfoManager interface {
    DeleteTopic()
    GetAllClusterInfo()
    GetAllTopicList()
    RegisterBroker()
    UnregisterBroker()

    ScanNotActiveBroker()
}

func NewRouteInfoManager() {
	fmt.Print("New route infos.")
}

type routeManager struct {

}

func (r *routeManager) DeleteTopic() {

}
