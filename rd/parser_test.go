package rd_test

import (
	"encoding/json"
	"reflect"
	"testing"

	rd "github.com/shivammg/parsers/rd"
)

func TestArithGrammar(test *testing.T) {
	/*
	   	"(", "id", "*", "id", ")", "+", "id"
	   Recursive Descent parser for the following grammar:
	   	E  -> TE'
	   	E' -> +TE'|ε
	   	T  -> FT'
	   	T' -> *FT'|ε
	   	F  -> id|(E)

	   - Used in parsing addition and multiplication arithmetic expressions.
	   - ε represents empty string.
	*/
	wantJSON := `{
		"Symbol": "E",
		"Children": [
			{
				"Symbol": "T",
				"Children": [
					{
						"Symbol": "F",
						"Children": [
							{
								"Symbol": "(",
								"Children": null
							},
							{
								"Symbol": "E",
								"Children": [
									{
										"Symbol": "T",
										"Children": [
											{
												"Symbol": "F",
												"Children": [
													{
														"Symbol": "id",
														"Children": null
													}
												]
											},
											{
												"Symbol": "T'",
												"Children": [
													{
														"Symbol": "*",
														"Children": null
													},
													{
														"Symbol": "F",
														"Children": [
															{
																"Symbol": "id",
																"Children": null
															}
														]
													},
													{
														"Symbol": "T'",
														"Children": null
													}
												]
											}
										]
									},
									{
										"Symbol": "E'",
										"Children": null
									}
								]
							},
							{
								"Symbol": ")",
								"Children": null
							}
						]
					},
					{
						"Symbol": "T'",
						"Children": null
					}
				]
			},
			{
				"Symbol": "E'",
				"Children": [
					{
						"Symbol": "+",
						"Children": null
					},
					{
						"Symbol": "T",
						"Children": [
							{
								"Symbol": "F",
								"Children": [
									{
										"Symbol": "id",
										"Children": null
									}
								]
							},
							{
								"Symbol": "T'",
								"Children": null
							}
						]
					},
					{
						"Symbol": "E'",
						"Children": null
					}
				]
			}
		]
	}`
	var want, got interface{}
	json.Unmarshal([]byte(wantJSON), &want)
	p := rd.NewParser([]string{"(", "id", "*", "id", ")", "+", "id"})

	p.Register("E", func() bool {
		return p.Match("T") && p.Match("E'")
	})

	p.Register("E'", func() bool {
		if p.Match("+") &&
			p.Match("T") &&
			p.Match("E'") {
			return true
		}
		// epsilon exists for the rule
		return true
	})

	p.Register("T", func() bool {
		if p.Match("F") &&
			p.Match("T'") {
			return true
		}
		return false
	})

	p.Register("T'", func() bool {
		if p.Match("*") &&
			p.Match("F") &&
			p.Match("T'") {
			return true
		}
		// epsilon exists for the rule
		return true
	})

	p.Register("F", func() bool {
		if p.Match("id") {
			return true
		}
		return p.Match("(") && p.Match("E") && p.Match(")")
	})

	gotOK := p.Match("E")
	if gotOK != true {
		test.Fatal("Parsing failed")
	}
	gotJSON, _ := json.Marshal(p.Tree())
	json.Unmarshal(gotJSON, &got)
	if !reflect.DeepEqual(want, got) {
		test.Errorf("Expected: %v\nGot: %v\n", want, got)
	}
}

func TestInvalidInput(test *testing.T) {
	p := rd.NewParser([]string{"a", "c"})

	p.Register("E", func() bool {
		if p.Match("a") &&
			p.Match("F") {
			return true
		}
		return p.Match("G")
	})

	p.Register("F", func() bool {
		return p.Match("b")
	})

	p.Register("G", func() bool {
		return p.Match("c")
	})

	ok := p.Match("E")
	if ok {
		test.Error("Match should've failed")
	}
}
