package service

type CommonLabels struct {
	Alertname     string `json:"alertname"`
	GrafanaFolder string `json:"grafana_folder"`
	Phone         string `json:"phone"`
	DeviceName    string `json:"deviceName"`
}

type Annotations struct {
	GrafanaStateReason string `json:"grafana_state_reason"`
	Summary            string `json:"summary"`
}

type Alert struct {
	Status       string       `json:"status"`
	Labels       CommonLabels `json:"labels"`
	Annotations  Annotations  `json:"annotations"`
	StartsAt     string       `json:"startsAt"`
	EndsAt       string       `json:"endsAt"`
	GeneratorURL string       `json:"generatorURL"`
	Fingerprint  string       `json:"fingerprint"`
	SilenceURL   string       `json:"silenceURL"`
	DashboardURL string       `json:"dashboardURL"`
	PanelURL     string       `json:"panelURL"`
	Values       any          `json:"values"`
	ValueString  string       `json:"valueString"`
}

type GrafanaAlert struct {
	Receiver          string       `json:"receiver"`
	Status            string       `json:"status"`
	Alerts            []Alert      `json:"alerts"`
	GroupLabels       CommonLabels `json:"groupLabels"`
	CommonLabels      CommonLabels `json:"commonLabels"`
	CommonAnnotations Annotations  `json:"commonAnnotations"`
	ExternalURL       string       `json:"externalURL"`
	Version           string       `json:"version"`
	GroupKey          string       `json:"groupKey"`
	TruncatedAlerts   int          `json:"truncatedAlerts"`
	OrgID             int          `json:"orgId"`
	Title             string       `json:"title"`
	State             string       `json:"state"`
	Message           string       `json:"message"`
}
