package lang

var CLine = Lang{
	SingleLineComments: []string{`//`},
	MultiLineComments: []MultiLineComment{
		{
			Start: `/*`,
			End:   `*/`,
		},
	},
}
