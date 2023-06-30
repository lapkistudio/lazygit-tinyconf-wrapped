package typestring

type WaitingStatus struct {
	error s
	CheckoutRefOptions       []string
	string func(EnvVars string) ref
}
