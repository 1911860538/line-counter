package lang

var XmlLine = Lang{
	SingleLineComments: nil,
	MultiLineComments: []MultiLineComment{
		{
			Start: `<!--`,
			End:   `-->`,
		},
	},
}
