package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"log"
)

type collector struct {
	apiKey string
}

func (m collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- sent
	ch <- hardBounces
	ch <- softBounces
	ch <- rejects
	ch <- complaints
	ch <- unsubs
	ch <- opens
	ch <- clicks
	ch <- unique_opens
	ch <- unique_clicks
}

type mandrillUserInfo struct {
	Username string `json:"username"`
	Stats    struct {
		AllTime struct {
			Sent          int `json:"sent"`
			SoftBounces   int `json:"soft_bounces"`
			HardBounces   int `json:"hard_bounces"`
			Rejects       int `json:"rejects"`
			Complaints    int `json:"complaints"`
			Unsubs        int `json:"unsubs"`
			Opens         int `json:"opens"`
			Clicks        int `json:"clicks"`
			Unique_opens  int `json:"unique_opens"`
			Unique_clicks int `json:"unique_clicks"`
		} `json:"all_time"`
	}
}

func (m collector) Collect(ch chan<- prometheus.Metric) {
	info, err := getMandrillUserInfo(m.apiKey)
	if err != nil {
		log.Println(err)
		return
	}

	stats := info.Stats.AllTime
	ch <- prometheus.MustNewConstMetric(sent, prometheus.CounterValue, float64(stats.Sent), info.Username)
	ch <- prometheus.MustNewConstMetric(hardBounces, prometheus.CounterValue, float64(stats.HardBounces), info.Username)
	ch <- prometheus.MustNewConstMetric(softBounces, prometheus.CounterValue, float64(stats.SoftBounces), info.Username)
	ch <- prometheus.MustNewConstMetric(rejects, prometheus.CounterValue, float64(stats.Rejects), info.Username)
	ch <- prometheus.MustNewConstMetric(complaints, prometheus.CounterValue, float64(stats.Complaints), info.Username)
	ch <- prometheus.MustNewConstMetric(unsubs, prometheus.CounterValue, float64(stats.Unsubs), info.Username)
	ch <- prometheus.MustNewConstMetric(opens, prometheus.CounterValue, float64(stats.Opens), info.Username)
	ch <- prometheus.MustNewConstMetric(clicks, prometheus.CounterValue, float64(stats.Clicks), info.Username)
	ch <- prometheus.MustNewConstMetric(unique_opens, prometheus.CounterValue, float64(stats.Unique_opens), info.Username)
	ch <- prometheus.MustNewConstMetric(unique_clicks, prometheus.CounterValue, float64(stats.Unique_clicks), info.Username)
}
