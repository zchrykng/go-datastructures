package containers

type JSONSerializer interface {
	ToJSON() ([]byte, error)

	MarshelJSON() ([]byte, error)
}

type JSONDeserializer interface {
	FromJSON([]byte) error
	UnmarshalJSON([]byte) error
}
