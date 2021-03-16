# Mandrill statistic exporter for Prometheus
```
docker run -p 9861:9861 \
    -e MANDRILL_EXPORTER_API_KEY=your-twilio-account-id \
    rentberry/mandrill-prometheus-exporter:latest
```

## What's exported?
- ``mandrill_sent_total`` - Total number of sent mails
- ``mandrill_hard_bounces`` - Number of mails bounced hard
- ``mandrill_soft_bounces`` - Number of mails bounced soft
- ``mandrill_rejects`` - Number of mails rejected
- ``mandrill_complaints`` - Number of complaints
- ``mandrill_unsubs`` - Number of unsubscribes
- ``mandrill_opens`` - Number of mails opened
- ``mandrill_clicks`` - Number of clicks inside mails
- ``mandrill_unique_opens`` - Unique number of mails opened
- ``mandrill_unique_clicks`` - Unique number of clicks
