package cmd

type DataflowStruct struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Data struct {
	Dataflows      []DataflowStruct `json:"dataflows"`
	DataStructures []DataStructure  `json:"datastructures"`
	Codelists      []Codelist       `json:"codelists"`
}

type Root struct {
	Data Data `json:"data"`
}

type DataStructure struct {
	Id           string       `json:"id"`
	DsComponents DsComponents `json:"datastructurecomponents"`
}

type DsComponents struct {
	AttributeList AttributeList `json:"attributelist"`
	DimList       DimList       `json:"dimensionlist"`
	//MeasureList   string `json:"measurelist"`
}

type DimList struct {
	Dimensions []Dimension `json:"dimensions"`
}

type Dimension struct {
	Id       string `json:"id"`
	Position int    `json:"position"`
	Type     string `json:"type"`
}

type AttributeList struct {
	Attributes []Attribute `json:"attributes"`
}

type Attribute struct {
	Id string `json:"id"`
}

type Codelist struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Codes []Code `json:"codes"`
}

type Code struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Annotations []Annotation `json:"annotations"`
}

type Annotation struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
