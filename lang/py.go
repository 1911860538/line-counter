package lang

var PyLine = Lang{
	SingleLineComments: []string{`#`},
	MultiLineComments: []MultiLineComment{
		{
			Start: `"""`,
			End:   `"""`,
		},
		{
			Start: `'''`,
			End:   `'''`,
		},
	},
}
