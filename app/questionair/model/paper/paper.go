package paper

type Paper struct {
	Id			int		`json:"id"`
	Title 		string	`json:"title"`
	Subtitle	string	`json:"subtitle"`
	Blocks		[]*Block `json:"-"`
}

type Block struct {
	Id			int 	`json:"id"`
	Title		string 	`json:"title"`
	PaperId		int		`json:"-"`
	Questions	[]*Question `json:"-"`
}

type Question struct {
	Id			int		`json:"id"`
	Stem		string 	`json:"string"`
	Options		[]*QuestionOption	`json:"-"`
}

type QuestionOption struct {
	Id			int		`json:"id"`
	Value		string	`json:"value"`
	QuestionId	int		`json:"-"`
}