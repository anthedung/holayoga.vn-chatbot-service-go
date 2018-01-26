package vn_holayoga_dialogflow_service

type Config struct {
	Datastore *DatastoreConfig
}

type DatastoreConfig struct {
	// Kind (i.e. table) for yoga categories
	CategoryKind string
}

func NewDefaultConfig() *Config {
	return &Config{
		Datastore: &DatastoreConfig{
			CategoryKind: "Category",
		},
	}
}