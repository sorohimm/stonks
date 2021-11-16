package quotes_models

type IntradayTSMongo struct {
	MetaData MetaMongo      `json:"_meta" bson:"_meta"`
	Series   []SessionMongo `json:"series,omitempty" bson:"series,omitempty"`
}

