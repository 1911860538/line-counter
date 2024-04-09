package lang

var CsLine = Lang{
	SingleLineComments: []string{`//`, `///`},
	MultiLineComments: []MultiLineComment{
		{
			Start: `/*`,
			End:   `*/`,
		},
	},
}
