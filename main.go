package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var (
	sent          = prometheus.NewDesc("mandrill_sent_total", "Total number of sent mails.", []string{"username"}, nil)
	hardBounces   = prometheus.NewDesc("mandrill_hard_bounces", "Number of mails bounced hard", []string{"username"}, nil)
	softBounces   = prometheus.NewDesc("mandrill_soft_bounces", "Number of mails bounced soft", []string{"username"}, nil)
	rejects       = prometheus.NewDesc("mandrill_rejects", "Number of mails rejected", []string{"username"}, nil)
	complaints    = prometheus.NewDesc("mandrill_complaints", "Number of complaints", []string{"username"}, nil)
	unsubs        = prometheus.NewDesc("mandrill_unsubs", "Number of unsubscribes", []string{"username"}, nil)
	opens         = prometheus.NewDesc("mandrill_opens", "Number of mails opened", []string{"username"}, nil)
	clicks        = prometheus.NewDesc("mandrill_clicks", "Number of clicks inside mails", []string{"username"}, nil)
	unique_opens  = prometheus.NewDesc("mandrill_unique_opens", "Unique number of mails opened", []string{"username"}, nil)
	unique_clicks = prometheus.NewDesc("mandrill_unique_clicks", "Unique number of clicks", []string{"username"}, nil)
)

type config struct {
	Addr           string `envconfig:"MANDRILL_EXPORTER_EXPORTER_LISTEN_ADDR" default:":9861"`
	MandrillApiKey string `envconfig:"MANDRILL_EXPORTER_API_KEY" required:"true"`
}

func main() {
	var conf config
	err := envconfig.Process("", &conf)
	if err != nil {
		log.Fatal(err)
	}

	mc := collector{
		apiKey: conf.MandrillApiKey,
	}

	prometheus.MustRegister(mc)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Mandrill stats Exporter</title></head>
             <body>
             <h1>Madrill statistics Exporter</h1>
             <p><a href='metrics'>Metrics</a></p>
             </body>
             </html>`))
	})
	http.Handle("/metrics", promhttp.Handler())
	err = http.ListenAndServe(conf.Addr, nil)
	if err != nil {
		log.Println(err)
	}
}
