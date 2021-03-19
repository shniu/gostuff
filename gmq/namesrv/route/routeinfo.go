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

type QueueData struct {
    BrokerName string
    ReadQueueNums int
    WriteQueueNums int
    Perm int
    TopicSyncFlag int
}

func NewRouteInfoManager() {
	fmt.Print("New route infos.")
}

type routeManager struct {

}

func (r *routeManager) DeleteTopic() {

}
