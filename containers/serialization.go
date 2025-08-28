package containers

type JSONSerializer interface {
	ToJSON() ([]byte, error)

	MarshalJSON() ([]byte, error)
}

type JSONDeserializer interface {
	FromJSON([]byte) error
	UnmarshalJSON([]byte) error
}
