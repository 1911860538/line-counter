package lang

var SqlLine = Lang{
	SingleLineComments: []string{`--`},
	MultiLineComments: []MultiLineComment{
		{
			Start: `/*`,
			End:   `*/`,
		},
	},
}
