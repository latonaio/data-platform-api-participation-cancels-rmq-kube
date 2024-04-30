package requests

type Header struct {
	Participation	int     `json:"Participation"`
	IsCancelled		*bool	`json:"IsCancelled"`
}
