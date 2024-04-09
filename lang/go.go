package lang

var GoLine = Lang{
	SingleLineComments: []string{`//`},
	MultiLineComments: []MultiLineComment{
		{
			Start: `/*`,
			End:   `*/`,
		},
	},
}
