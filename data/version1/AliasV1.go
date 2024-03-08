package data1

type AliasV1 struct {
	Id    string `json:"id" bson:"_id"`
	Alias string `json:"alias" bson:"alias"`
	Url   string `json:"url" bson:"url"`
}

func (a AliasV1) Clone() AliasV1 {
	return AliasV1{
		Id:    a.Id,
		Alias: a.Alias,
		Url:   a.Url,
	}
}
