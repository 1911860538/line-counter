package lang

var HsLine = Lang{
	SingleLineComments: []string{`--`},
	MultiLineComments: []MultiLineComment{
		{
			Start: `{-`,
			End:   `-}`,
		},
	},
}
