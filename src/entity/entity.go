package entity

const (
	LabelBean        string = "Bean"
	LabelLettuce     string = "Lettuce"
	LabelCPD         string = "CPD"
	LabelChickenCoop string = "Chicken Coop"
	LabelServerX     string = "Server X"
)

type Item struct {
	Label string  `json:"label"`
	Tags  Tags    `json:"tag"`
	Value float32 `json:"value"`
}

type Tag struct {
	Label string `json:"labe"`
	Value string `json:"value"`
}

type Tags []Tag

func (ib Item) ToInfluxDb() map[string]interface{} {
	return map[string]interface{}{
		"value": ib.Value,
	}
}

func (t Tags) ToInfluxDb() map[string]string {
	extractedTags := make(map[string]string)
	for _, tag := range t {
		extractedTags[tag.Label] = tag.Value
	}
	return extractedTags
}
