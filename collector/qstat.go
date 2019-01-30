package collector

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/taylor840326/go_pbspro/qstat"
)

func init() {
	registerCollector(qstatCollectorSubSystem, defaultEnabled, NewQstatCollector)
}

type qstatCollector struct {
	server_state string
}

func (c *qstatCollector) Update(ch chan<- prometheus.Metric) error {
	c.updateQstatServer(ch)
	c.updateQstatQueue(ch)
	return nil
}

type qstatMetric struct {
	name            string
	desc            string
	value           float64
	metricType      prometheus.ValueType
	extraLabel      []string
	extraLabelValue string
}

func NewQstatCollector() (Collector, error) {
	qc := new(qstatCollector)
	return &qstatCollector{server_state: qc.server_state}, nil
}

func (c *qstatCollector) updateQstatServer(ch chan<- prometheus.Metric) {

	var allMetrics []qstatMetric
	//var metrics []qstatMetric
	var labelsValue []string

	qstat, err := qstat.NewQstat("172.18.7.10")
	if err != nil {
		fmt.Println(err.Error())
	}

	qstat.SetAttribs(nil)
	qstat.SetExtend("")

	err = qstat.ConnectPBS()
	if err != nil {
		fmt.Println("ConnectPBS Error")
	}
	defer qstat.DisconnectPBS()

	err = qstat.PbsServerState()
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, ss := range qstat.ServerState {
		allMetrics = []qstatMetric{
			{
				name:       "server_state",
				desc:       "pbspro_exporter: server state. 1 is Active",
				value:      float64(ss.ServerState),
				metricType: prometheus.GaugeValue,
			},
			{

				name:       "server_scheduling",
				desc:       "pbspro_exporter: Server Scheduling. 1 is True",
				value:      float64(ss.ServerScheduling),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_total_jobs",
				desc:       "pbspro_exporter: Server Total Jobs.",
				value:      float64(ss.TotalJobs),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_transit_state_count",
				desc:       "pbspro_exporter: Server Transit State Count.",
				value:      float64(ss.StateCountTransit),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_queued_state_count",
				desc:       "pbspro_exporter: Server Queued State Count.",
				value:      float64(ss.StateCountQueued),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_held_state_count",
				desc:       "pbspro_exporter: Server Held State Count.",
				value:      float64(ss.StateCountHeld),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_waiting_state_count",
				desc:       "pbspro_exporter: Server Waiting State Count.",
				value:      float64(ss.StateCountWaiting),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_running_state_count",
				desc:       "pbspro_exporter: Server Running State Count.",
				value:      float64(ss.StateCountRunning),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_exiting_state_count",
				desc:       "pbspro_exporter: Server Exiting State Count.",
				value:      float64(ss.StateCountExiting),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_begun_state_count",
				desc:       "pbspro_exporter: Server Begun State Count.",
				value:      float64(ss.StateCountBegun),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_log_events",
				desc:       "pbspro_exporter: Server Log Events.",
				value:      float64(ss.LogEvents),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_query_other_jobs",
				desc:       "pbspro_exporter: Server Query Other Jobs. 1 is True",
				value:      float64(ss.QueryOtherJobs),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_resources_default_ncpus",
				desc:       "pbspro_exporter: Server Resources Default Ncpus.",
				value:      float64(ss.ResourcesDefaultNcpus),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_default_chunk_ncpus",
				desc:       "pbspro_exporter: Server Default Chunk Ncpus.",
				value:      float64(ss.DefaultChunkNcpus),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_resources_assigned_ncpus",
				desc:       "pbspro_exporter: Server Resources Assigned Ncpus.",
				value:      float64(ss.ResourcesAssignedNcpus),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_resources_assigned_nodect",
				desc:       "pbspro_exporter: Server Resources Assigned Nodect.",
				value:      float64(ss.ResourcesAssignedNodect),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_scheduler_iteration",
				desc:       "pbspro_exporter: Server Scheudler Iteration.",
				value:      float64(ss.SchedulerIteration),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_flicenses",
				desc:       "pbspro_exporter: Server Flicense.",
				value:      float64(ss.Flicenses),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_resv_enable",
				desc:       "pbspro_exporter: Server Resv Enable. 1 is True",
				value:      float64(ss.ResvEnable),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_node_fail_requeue",
				desc:       "pbspro_exporter: Server Node Fail Requeue.",
				value:      float64(ss.NodeFailRequeue),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_max_array_size",
				desc:       "pbspro_exporter: Server Max Array Size.",
				value:      float64(ss.MaxArraySize),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_pbs_license_min",
				desc:       "pbspro_exporter: Server PBS License Min.",
				value:      float64(ss.PBSLicenseMin),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_pbs_license_max",
				desc:       "pbspro_exporter: Server PBS License Max.",
				value:      float64(ss.PBSLicenseMax),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_pbs_license_linger_time",
				desc:       "pbspro_exporter: Server PBS License Linger Time.",
				value:      float64(ss.PBSLicenseLingerTime),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_license_count_avail_global",
				desc:       "pbspro_exporter: Server License Count Avail Global.",
				value:      float64(ss.LicenseCountAvailGlobal),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_license_count_avail_local",
				desc:       "pbspro_exporter: Server License Count Avail Global.",
				value:      float64(ss.LicenseCountAvailLocal),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_license_count_used",
				desc:       "pbspro_exporter: Server License Used.",
				value:      float64(ss.LicenseCountUsed),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_license_count_high_use",
				desc:       "pbspro_exporter: Server License Count High Use.",
				value:      float64(ss.LicenseCountHighUse),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_eligible_time_enable",
				desc:       "pbspro_exporter: Server Eligible Time Enable.1 is True",
				value:      float64(ss.EligibleTimeEnable),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_job_history_enable",
				desc:       "pbspro_exporter: Server Job History Enable.1 is True",
				value:      float64(ss.JobHistoryEnable),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_job_history_duration",
				desc:       "pbspro_exporter: Server Job History Duration.",
				value:      float64(ss.JobHistoryDuration),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_max_concurrent_provision",
				desc:       "pbspro_exporter: Server Max Concurrent Provision.",
				value:      float64(ss.MaxConcurrentProvision),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_power_provisioning",
				desc:       "pbspro_exporter: Server Power Provisioning. 1 is True",
				value:      float64(ss.PowerProvisioning),
				metricType: prometheus.GaugeValue,
			},
		}
		labelsValue = []string{ss.ServerName, ss.ServerHost, ss.DefaultQueue, ss.MailFrom, ss.PBSVersion}
	}

	for _, m := range allMetrics {

		labelsName := []string{"ServerName", "ServerHost", "DefaultQueue", "MailFrom", "PBSVersion"}

		desc := prometheus.NewDesc(
			prometheus.BuildFQName(namespace, qstatCollectorSubSystem, m.name),
			m.desc,
			labelsName,
			nil,
		)

		ch <- prometheus.MustNewConstMetric(
			desc,
			m.metricType,
			m.value,
			labelsValue...,
		)
	}

}

func (c *qstatCollector) updateQstatQueue(ch chan<- prometheus.Metric) {

	var allMetrics []qstatMetric
	//var metrics []qstatMetric
	var labelsValue []string

	qstat, err := qstat.NewQstat("172.18.7.10")
	if err != nil {
		fmt.Println(err.Error())
	}

	qstat.SetAttribs(nil)
	qstat.SetExtend("")

	err = qstat.ConnectPBS()
	if err != nil {
		fmt.Println("ConnectPBS Error")
	}
	defer qstat.DisconnectPBS()

	err = qstat.PbsQueueState()
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, ss := range qstat.QueueState {
		allMetrics = []qstatMetric{
			{
				name:       "queue_total_jobs",
				desc:       "pbspro_exporter: Queue Total Jobs.",
				value:      float64(ss.TotalJobs),
				metricType: prometheus.GaugeValue,
			},
			{

				name:       "queue_transit_state_count",
				desc:       "pbspro_exporter: Queue Transit State Count.",
				value:      float64(ss.StateCountTransit),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "queue_queued_state_count",
				desc:       "pbspro_exporter: Queue Queued State Count.",
				value:      float64(ss.StateCountQueued),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "queue_held_state_count",
				desc:       "pbspro_exporter: Queue Held State Count.",
				value:      float64(ss.StateCountHeld),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "queue_waiting_state_count",
				desc:       "pbspro_exporter: Queue Waiting State Count.",
				value:      float64(ss.StateCountWaiting),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "queue_running_state_count",
				desc:       "pbspro_exporter: Queue Running State Count.",
				value:      float64(ss.StateCountRunning),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "queue_exiting_state_count",
				desc:       "pbspro_exporter: Queue Exiting State Count.",
				value:      float64(ss.StateCountExiting),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "queue_begun_state_count",
				desc:       "pbspro_exporter: Queue Begun State Count.",
				value:      float64(ss.StateCountBegun),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "queue_resources_assigned_ncpus",
				desc:       "pbspro_exporter: Queue Resources Assigned Ncpus.",
				value:      float64(ss.ResourcesAssignedNcpus),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "queue_resources_assigned_nodect",
				desc:       "pbspro_exporter: Queue Resources Assigned Nodect.",
				value:      float64(ss.ResourcesAssignedNodect),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "queue_enable",
				desc:       "pbspro_exporter: Queue Enable. 1 is True",
				value:      float64(ss.Enable),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "queue_started",
				desc:       "pbspro_exporter: Queue Started. 1 is True",
				value:      float64(ss.Started),
				metricType: prometheus.GaugeValue,
			},
		}
		labelsValue = []string{ss.QueueName, ss.QueueType}
	}

	for _, m := range allMetrics {

		labelsName := []string{"QueueName", "QueueType"}

		desc := prometheus.NewDesc(
			prometheus.BuildFQName(namespace, qstatCollectorSubSystem, m.name),
			m.desc,
			labelsName,
			nil,
		)

		ch <- prometheus.MustNewConstMetric(
			desc,
			m.metricType,
			m.value,
			labelsValue...,
		)
	}

}
