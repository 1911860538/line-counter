package lang

var ScalaLine = Lang{
	SingleLineComments: []string{`//`},
	MultiLineComments: []MultiLineComment{
		{
			Start: `/*`,
			End:   `*/`,
		},
	},
}
