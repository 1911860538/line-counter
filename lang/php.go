package lang

var PhpLine = Lang{
	SingleLineComments: []string{`//`, `#`},
	MultiLineComments: []MultiLineComment{
		{
			Start: `/*`,
			End:   `*/`,
		},
	},
}
