package lang

var CppLine = Lang{
	SingleLineComments: []string{`//`},
	MultiLineComments: []MultiLineComment{
		{
			Start: `/*`,
			End:   `*/`,
		},
	},
}
