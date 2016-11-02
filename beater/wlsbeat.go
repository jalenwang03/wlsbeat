package beater

import (
	"fmt"
	"strconv"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jalenwang03/wlsbeat/config"
)

//var (
//	Host     string
//	Port     string
//	Username string
//	Password string
//)

type wlsinstances struct {
	host     string
	port     string
	username string
	password string
}
type Servers struct {
	Name   string `xml:"name"`
	State  string `xml:"state"`
	Health string `xml:"health"`
}
type ServersBody struct {
	Items []Servers `xml:"items`
}
type ServersList struct {
	Body     ServersBody `xml:"data"`
	Messages []string    `xml:"messages`
}

type Server struct {
	Name                    string `json:"name"`
	State                   string `json:"state"`
	Health                  string `json:"health"`
	Clustername             string `json:"clusterName"`
	CurrentMachine          string `json:"currentMachine"`
	WeblogicVersion         string `json:"weblogicVersion"`
	OpenSocketsCurrentCount uint64 `json:"openSocketsCurrentCount"`
	HeapSizeCurrent         uint64 `json:"heapSizeCurrent"`
	HeapFreeCurrent         uint64 `json:"heapFreeCurrent"`
	heapSizeMax             uint64 `json:"heapSizeMax"`
	JavaVersion             string `json:"javaVersion"`
	OsName                  string `json:"osName"`
	OsVersion               string `json:"osVersion"`
	JvmProcessorLoad        uint64 `json:"jvmProcessorLoad"`
}

////中间结构
type ServerBody struct {
	Items Server `json:"item"`
}

////最外层结构
type ServerList struct {
	Body     ServerBody `json:"body"`
	Messages []string   `json:"messages`
}

type Apps struct {
	Name    string `json:"name"`
	Apptype string `json:"type"`
	State   string `json:"state"`
	Health  string `json:"health"`
}
type AppsBody struct {
	Items []Apps `json:"items"`
}

type AppList struct {
	Body     AppsBody `json:"body"`
	Messages []string `json:"messages`
}

type AppTargetstateStruct struct {
	Target string `json:"target"`
	State  string `json:"state"`
}

type AppWorkManagers struct {
	Name              string `json:"name"`
	pendingRequests   uint64 `json:"pendingRequests"`
	CompletedRequests uint64 `json:"completedRequests"`
	Server            string `json:"server"`
}
type AppDataSource struct {
	Name   string `json:"name"`
	Server uint64 `json:"server"`
	State  uint64 `json:"state"`
}

type AppminThreadsConstraints struct {
	Name                     string `json:"name"`
	Server                   string `json:"server"`
	PendingRequests          uint64 `json:"pendingRequests"`
	CompletedRequests        string `json:"completedRequests"`
	ExecutingRequests        uint64 `json:"executingRequests"`
	OutOfOrderExecutionCount uint64 `json:"outOfOrderExecutionCount"`
	MustRunCount             uint64 `json:"mustRunCount"`
	MaxWaitTime              uint64 `json:"maxWaitTime"`
	CurrentWaitTime          uint64 `json:"currentWaitTime"`
}

type AppmaxThreadsConstraints struct {
	Name              string `json:"name"`
	Server            string `json:"server"`
	DeferredRequests  uint64 `json:"deferredRequests"`
	ExecutingRequests uint64 `json:"executingRequests"`
}
type ApprequestClasses struct {
	Name                 string `json:"name"`
	Server               string `json:"server"`
	RequestClassType     string `json:"requestClassType"`
	CompletedCount       uint64 `json:"completedCount"`
	TotalThreadUse       uint64 `json:"totalThreadUse"`
	PendingRequestCount  uint64 `json:"pendingRequestCount"`
	VirtualTimeIncrement uint64 `json:"virtualTimeIncrement"`
}
type ApplicationItem struct {
	Name                  string                     `json:"name"`
	Apptype               string                     `json:"type"`
	State                 string                     `json:"state"`
	Health                string                     `json:"health"`
	TargetStates          []AppTargetstateStruct     `json:"targetStates"`
	DataSources           []AppDataSource            `json:"dataSources"`
	WorkManagers          []AppWorkManagers          `json:"workManagers"`
	MinThreadsConstraints []AppminThreadsConstraints `json:"minThreadsConstraints"`
	MaxThreadsConstraints []AppmaxThreadsConstraints `json:"maxThreadsConstraints"`
	RequestClasses        []ApprequestClasses        `json:"requestClasses"`
	JvmProcessorLoad      uint64                     `json:"jvmProcessorLoad"`
}
type ApplicationBody struct {
	Items ApplicationItem `json:"item"`
}

type Application struct {
	Body     ApplicationBody `json:"body"`
	Messages []string        `json:"messages`
}
type DSServerState struct {
	Server string `json:"server"`
	State  string `json:"state"`
}
type DSS struct {
	Name      string          `json:"name"`
	DStype    string          `json:"type"`
	Instances []DSServerState `json:"instances"`
}
type DSSBody struct {
	Items []DSS `json:"items"`
}

type DSList struct {
	Body     DSSBody  `json:"body"`
	Messages []string `json:"messages`
}

type DSState struct {
	Server                        string `json:"server"`
	State                         string `json:"state"`
	Enabled                       bool   `json:"enabled"`
	VersionJDBCDriver             string `json:"versionJDBCDriver"`
	ActiveConnectionsAverageCount uint64 `json:"activeConnectionsAverageCount"`
	ActiveConnectionsCurrentCount uint64 `json:"activeConnectionsCurrentCount"`
	ActiveConnectionsHighCount    uint64 `json:"activeConnectionsHighCount"`
	ConnectionDelayTime           uint64 `json:"connectionDelayTime"`

	ConnectionsTotalCount     uint64 `json:"connectionsTotalCount"`
	CurrCapacity              uint64 `json:"currCapacity"`
	CurrCapacityHighCount     uint64 `json:"currCapacityHighCount"`
	FailedReserveRequestCount uint64 `json:"failedReserveRequestCount"`
	FailuresToReconnectCount  uint64 `json:"failuresToReconnectCount"`
	HighestNumAvailable       uint64 `json:"highestNumAvailable"`
	LeakedConnectionCount     uint64 `json:"leakedConnectionCount"`
	NumAvailable              uint64 `json:"numAvailable"`
	NumUnavailable            uint64 `json:"numUnavailable"`
	PrepStmtCacheAccessCount  uint64 `json:"prepStmtCacheAccessCount"`
	PrepStmtCacheAddCount     uint64 `json:"prepStmtCacheAddCount"`
	PrepStmtCacheCurrentSize  uint64 `json:"prepStmtCacheCurrentSize"`

	PrepStmtCacheDeleteCount         uint64 `json:"prepStmtCacheDeleteCount"`
	PrepStmtCacheHitCount            uint64 `json:"prepStmtCacheHitCount"`
	PrepStmtCacheMissCount           uint64 `json:"prepStmtCacheMissCount"`
	ReserveRequestCount              uint64 `json:"reserveRequestCount"`
	WaitSecondsHighCount             uint64 `json:"waitSecondsHighCount"`
	WaitingForConnectionCurrentCount uint64 `json:"waitingForConnectionCurrentCount"`
	WaitingForConnectionFailureTotal uint64 `json:"waitingForConnectionFailureTotal"`
	WaitingForConnectionHighCount    uint64 `json:"waitingForConnectionHighCount"`

	WaitingForConnectionSuccessTotal   uint64 `json:"waitingForConnectionSuccessTotal"`
	WaitingForConnectionTotal          uint64 `json:"waitingForConnectionTotal"`
	SuccessfulRclbBasedBorrowCount     uint64 `json:"successfulRclbBasedBorrowCount"`
	FailedRCLBBasedBorrowCount         uint64 `json:"failedRCLBBasedBorrowCount"`
	SuccessfulAffinityBasedBorrowCount uint64 `json:"successfulAffinityBasedBorrowCount"`
	FailedAffinityBasedBorrowCount     uint64 `json:"failedAffinityBasedBorrowCount"`
}
type DS struct {
	Name      string    `json:"name"`
	DStype    string    `json:"type"`
	Instances []DSState `json:"instances"`
}
type DSBody struct {
	Item DS `json:"item"`
}

type Datasource struct {
	Body     DSBody   `json:"body"`
	Messages []string `json:"messages`
}
type wlsAll struct {
	Server_Info      ServerList  `Server_Info`
	Application_Info Application `Application_Info`
	Datasource_Info  DS          `Datasource_Info`
}

type Wlsbeat struct {
	done   chan struct{}
	config config.Config
	client publisher.Client
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Wlsbeat{
		done:   make(chan struct{}),
		config: config,
	}
	return bt, nil
}

func (bt *Wlsbeat) Run(b *beat.Beat) error {
	logp.Info("wlsbeat is running! Hit CTRL-C to stop it.")
	//	Host := bt.config.Host
	//	Port := bt.config.Port
	//	Username := bt.config.Username
	//	Password := bt.config.Password
	//	logp.Info("Host:" + Host)
	//	logp.Info("Port:" + Port)
	//	logp.Info("Username:" + Username)
	//	logp.Info("Password:" + Password)
	//	fmt.Println(bt.config.Instances)
	//	fmt.Println("Getting instances")
	instances := getInstances(bt)
	//	fmt.Println(instances)
	//	var serversList ServersList
	//	err := json.Unmarshal([]byte(GetPerfData(Host, Port, "servers", Username, Password)), &serversList)
	//	if err != nil {
	//		logp.Err("Error:", err)
	//	}
	//	for i := 0; i < len(serversList.Body.Items); i++ {

	//			logp.Info("Name:" + serversList.Body.Items[i].Name)
	//			logp.Info("Health:" + serversList.Body.Items[i].Health)
	//			logp.Info("State:" + serversList.Body.Items[i].State)

	//	}
	bt.client = b.Publisher.Connect()
	ticker := time.NewTicker(bt.config.Period)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}
		for j := 0; j < len(instances); j++ {
			Host := instances[j].host
			Port := instances[j].port
			Username := instances[j].username
			Password := instances[j].password
			logp.Info("Getting metrics from [" + strconv.Itoa(j) + "]" + Host + ":" + Port)
			GetServerInfo(Host, Port, Username, Password, bt, b)
			GetAppInfo(Host, Port, Username, Password, bt, b)
			GetDatasourceInfo(Host, Port, Username, Password, bt, b)
		}
	}
}

func (bt *Wlsbeat) Stop() {
	logp.Info("Startting to shutdown!")
	bt.client.Close()
	close(bt.done)
}

func GetServerInfo(host string, port string, username string, password string, bt *Wlsbeat, b *beat.Beat) {
	logp.Info("Getting server info...")
	var serversList ServersList
	err := json.Unmarshal([]byte(GetPerfData(host, port, "servers", username, password)), &serversList)
	if err != nil {
		logp.Err("Error:" + err.Error())
	}
	//	fmt.Println(serversList)
	i := 0
	for i < len(serversList.Body.Items) {
		//		mapServer := make(map[string]string)
		var s ServerList
		err := json.Unmarshal([]byte(GetPerfData(host, port, "servers/"+serversList.Body.Items[i].Name, username, password)), &s)
		if err != nil {
			logp.Err("Error:" + err.Error())
		}
		event := common.MapStr{
			"@timestamp":              common.Time(time.Now()),
			"type":                    b.Name,
			"Name":                    s.Body.Items.Name,
			"State":                   s.Body.Items.State,
			"Clustername":             s.Body.Items.Clustername,
			"CurrentMachine":          s.Body.Items.CurrentMachine,
			"WeblogicVersion":         s.Body.Items.WeblogicVersion,
			"OpenSocketsCurrentCount": s.Body.Items.OpenSocketsCurrentCount,
			"HeapSizeCurrent":         s.Body.Items.HeapSizeCurrent,
			"HeapFreeCurrent":         s.Body.Items.HeapFreeCurrent,
			"heapSizeMax":             s.Body.Items.heapSizeMax,
			"JavaVersion":             s.Body.Items.JavaVersion,
			"OsName":                  s.Body.Items.OsName,
			"OsVersion":               s.Body.Items.OsVersion,
			"JvmProcessorLoad":        s.Body.Items.JvmProcessorLoad,
		}
		bt.client.PublishEvent(event)
		logp.Info("Event sent")
		i++
	}
}

func GetAppInfo(host string, port string, username string, password string, bt *Wlsbeat, b *beat.Beat) {
	logp.Info("Getting webapp ...")
	var apps AppList
	err := json.Unmarshal([]byte(GetPerfData(host, port, "applications", username, password)), &apps)
	if err != nil {
		logp.Err("Error:" + err.Error())
	}
	i := 0
	for i < len(apps.Body.Items) {
		var s Application
		//		var dat map[string]interface{}
		//		returnstr := GetPerfData("applications/" + apps.Body.Items[i].Name)
		err := json.Unmarshal([]byte(GetPerfData(host, port, "applications/"+apps.Body.Items[i].Name, username, password)), &s)
		if err != nil {
			logp.Err("Error:" + err.Error())
		}
		event := common.MapStr{
			"@timestamp":            common.Time(time.Now()),
			"type":                  b.Name,
			"Name":                  s.Body.Items.Name,
			"Apptype":               s.Body.Items.Apptype,
			"State":                 s.Body.Items.State,
			"Health":                s.Body.Items.Health,
			"TargetStates":          s.Body.Items.TargetStates,
			"DataSources":           s.Body.Items.DataSources,
			"WorkManagers":          s.Body.Items.WorkManagers,
			"MinThreadsConstraints": s.Body.Items.MinThreadsConstraints,
			"MaxThreadsConstraints": s.Body.Items.MaxThreadsConstraints,
			"RequestClasses":        s.Body.Items.RequestClasses,
			"JvmProcessorLoad":      s.Body.Items.JvmProcessorLoad,
		}
		bt.client.PublishEvent(event)
		logp.Info("Event sent")
		i++
	}
}

func GetDatasourceInfo(host string, port string, username string, password string, bt *Wlsbeat, b *beat.Beat) {
	logp.Info("Getting datasource ...")
	var dss DSList
	err := json.Unmarshal([]byte(GetPerfData(host, port, "datasources", username, password)), &dss)
	if err != nil {
		logp.Err("Error:" + err.Error())
	}
	i := 0
	for i < len(dss.Body.Items) {
		//		var dat map[string]interface{}
		var s Datasource
		err := json.Unmarshal([]byte(GetPerfData(host, port, "datasources/"+dss.Body.Items[i].Name, username, password)), &s)
		if err != nil {
			logp.Err("Error:" + err.Error())
		}
		for j := 0; j < len(s.Body.Item.Instances); j++ {
			event := common.MapStr{
				"@timestamp":                    common.Time(time.Now()),
				"type":                          b.Name,
				"Server":                        s.Body.Item.Instances[j].Server,
				"State":                         s.Body.Item.Instances[j].State,
				"Enabled":                       s.Body.Item.Instances[j].Enabled,
				"VersionJDBCDriver":             s.Body.Item.Instances[j].VersionJDBCDriver,
				"ActiveConnectionsAverageCount": s.Body.Item.Instances[j].ActiveConnectionsAverageCount,
				"ActiveConnectionsCurrentCount": s.Body.Item.Instances[j].ActiveConnectionsCurrentCount,
				"ActiveConnectionsHighCount":    s.Body.Item.Instances[j].ActiveConnectionsHighCount,
				"ConnectionDelayTime":           s.Body.Item.Instances[j].ConnectionDelayTime,

				"ConnectionsTotalCount":     s.Body.Item.Instances[j].ConnectionsTotalCount,
				"CurrCapacity":              s.Body.Item.Instances[j].CurrCapacity,
				"CurrCapacityHighCount":     s.Body.Item.Instances[j].CurrCapacityHighCount,
				"FailedReserveRequestCount": s.Body.Item.Instances[j].FailedReserveRequestCount,
				"FailuresToReconnectCount":  s.Body.Item.Instances[j].FailuresToReconnectCount,
				"HighestNumAvailable":       s.Body.Item.Instances[j].HighestNumAvailable,
				"LeakedConnectionCount":     s.Body.Item.Instances[j].LeakedConnectionCount,
				"NumAvailable":              s.Body.Item.Instances[j].NumAvailable,
				"NumUnavailable":            s.Body.Item.Instances[j].NumUnavailable,
				"PrepStmtCacheAccessCount":  s.Body.Item.Instances[j].PrepStmtCacheAccessCount,
				"PrepStmtCacheAddCount":     s.Body.Item.Instances[j].PrepStmtCacheAddCount,
				"PrepStmtCacheCurrentSize":  s.Body.Item.Instances[j].PrepStmtCacheCurrentSize,

				"PrepStmtCacheDeleteCount":         s.Body.Item.Instances[j].PrepStmtCacheDeleteCount,
				"PrepStmtCacheHitCount":            s.Body.Item.Instances[j].PrepStmtCacheHitCount,
				"PrepStmtCacheMissCount":           s.Body.Item.Instances[j].PrepStmtCacheMissCount,
				"ReserveRequestCount":              s.Body.Item.Instances[j].ReserveRequestCount,
				"WaitSecondsHighCount":             s.Body.Item.Instances[j].WaitSecondsHighCount,
				"WaitingForConnectionCurrentCount": s.Body.Item.Instances[j].WaitingForConnectionCurrentCount,
				"WaitingForConnectionFailureTotal": s.Body.Item.Instances[j].WaitingForConnectionFailureTotal,
				"WaitingForConnectionHighCount":    s.Body.Item.Instances[j].WaitingForConnectionHighCount,

				"WaitingForConnectionSuccessTotal":   s.Body.Item.Instances[j].WaitingForConnectionSuccessTotal,
				"WaitingForConnectionTotal":          s.Body.Item.Instances[j].WaitingForConnectionTotal,
				"SuccessfulRclbBasedBorrowCount":     s.Body.Item.Instances[j].SuccessfulRclbBasedBorrowCount,
				"FailedRCLBBasedBorrowCount":         s.Body.Item.Instances[j].FailedRCLBBasedBorrowCount,
				"SuccessfulAffinityBasedBorrowCount": s.Body.Item.Instances[j].SuccessfulAffinityBasedBorrowCount,
				"FailedAffinityBasedBorrowCount":     s.Body.Item.Instances[j].FailedAffinityBasedBorrowCount,
			}
			bt.client.PublishEvent(event)

			logp.Info("Event sent")
		}
		i++
	}
}

func GetPerfData(host string, port string, path string, username string, password string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://"+host+":"+port+"/management/tenant-monitoring/"+path, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		logp.Err("Error", err)
		return ""
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logp.Err("Error", err)
	}
	logp.Info(string(bodyText))
	return string(bodyText)
}
func getInstances(bt *Wlsbeat) []wlsinstances {
	//	var instanceCount int
	//	instanceCount = len(bt.config.Instances)
	//	var instances [len(bt.config.Instances)]wlsinstances
	//	s := make([]int, n, 2*n)
	//	fmt.Println(len(bt.config.Instances))
	instances := make([]wlsinstances, len(bt.config.Instances), 2*len(bt.config.Instances))
	for i := 0; i < len(bt.config.Instances); i++ {
		var instance wlsinstances
		instance.host = bt.config.Instances[i].Host
		instance.port = bt.config.Instances[i].Port
		instance.username = bt.config.Instances[i].Username
		instance.password = bt.config.Instances[i].Password
		instances[i] = instance
	}
	return instances
}
