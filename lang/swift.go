package lang

var SwiftLine = Lang{
	SingleLineComments: []string{`//`},
	MultiLineComments: []MultiLineComment{
		{
			Start: `/*`,
			End:   `*/`,
		},
	},
}
