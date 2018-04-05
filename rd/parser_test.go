package rd_test

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	rd "github.com/shivammg/parsers/rd"
	t "github.com/shivammg/parsers/types"
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
	// Epsilon represents empty string.
	const Epsilon = "ε"
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
								"Children": [
								  {
									"Symbol": "ε",
									"Children": null
								  }
								]
							  }
							]
						  }
						]
					  },
					  {
						"Symbol": "E'",
						"Children": [
						  {
							"Symbol": "ε",
							"Children": null
						  }
						]
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
				"Children": [
				  {
					"Symbol": "ε",
					"Children": null
				  }
				]
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
					"Children": [
					  {
						"Symbol": "ε",
						"Children": null
					  }
					]
				  }
				]
			  },
			  {
				"Symbol": "E'",
				"Children": [
				  {
					"Symbol": "ε",
					"Children": null
				  }
				]
			  }
			]
		  }
		]
	  }`
	var want, got interface{}
	json.Unmarshal([]byte(wantJSON), &want)
	p := rd.NewParser([]string{"(", "id", "*", "id", ")", "+", "id"})

	p.Register("E", func() (*t.Tree, error) {
		T, err := p.Run("T")
		if err != nil {
			return nil, err
		}
		EP, err := p.Run("E'")
		if err != nil {
			return nil, err
		}
		return rd.T("E", T, EP), nil
	})

	p.Register("E'", func() (*t.Tree, error) {
		if p.Match("+") {
			TP, err := p.Run("T")
			if err != nil {
				return nil, err
			}
			EP, err := p.Run("E'")
			if err != nil {
				return nil, err
			}
			return rd.T("E'", rd.T("+"), TP, EP), nil
		}
		// epsilon exists for the rule
		return rd.T("E'", rd.T(Epsilon)), nil
	})

	p.Register("T", func() (*t.Tree, error) {
		F, err := p.Run("F")
		if err != nil {
			return nil, err
		}
		TP, err := p.Run("T'")
		if err != nil {
			return nil, err
		}
		return rd.T("T", F, TP), nil
	})

	p.Register("T'", func() (*t.Tree, error) {
		if p.Match("*") {
			F, err := p.Run("F")
			if err != nil {
				return nil, err
			}
			TP, err := p.Run("T'")
			if err != nil {
				return nil, err
			}
			return rd.T("T'", rd.T("*"), F, TP), nil
		}
		// epsilon exists for the rule
		return rd.T("T'", rd.T(Epsilon)), nil
	})

	p.Register("F", func() (*t.Tree, error) {
		if p.Match("id") {
			return rd.T("F", rd.T("id")), nil
		}
		if p.Match("(") {
			E, err := p.Run("E")
			if err != nil {
				return nil, err
			}
			if p.Match(")") {
				return rd.T("F", rd.T("("), E, rd.T(")")), nil
			}
		}
		return nil, errors.New(rd.ErrNoMatch)
	})

	tree, err := p.Run("E")
	if err != nil {
		test.Fatal(err)
	}
	gotJSON, _ := json.Marshal(tree)
	json.Unmarshal(gotJSON, &got)
	if !reflect.DeepEqual(want, got) {
		test.Errorf("Expected: %v\nGot: %v\n", want, got)
	}
}

func TestInvalidInput(test *testing.T) {
	p := rd.NewParser([]string{"a", "c"})

	p.Register("E", func() (*t.Tree, error) {
		if p.Match("a") {
			F, err := p.Run("F")
			if err == nil {
				return rd.T("E", rd.T("a"), F), nil
			}
			// explicitly backtrack since there was no incorrect Match,
			// and we need to run next production.
			p.Backtrack()
		}
		G, err := p.Run("G")
		if err == nil {
			return rd.T("E", G), nil
		}
		return nil, errors.New(rd.ErrNoMatch)
	})

	p.Register("F", func() (*t.Tree, error) {
		if p.Match("b") {
			return rd.T("F", rd.T("b")), nil
		}
		return nil, errors.New(rd.ErrNoMatch)
	})

	p.Register("G", func() (*t.Tree, error) {
		if p.Match("c") {
			return rd.T("G", rd.T("c")), nil
		}
		return nil, errors.New(rd.ErrNoMatch)
	})

	_, err := p.Run("E")
	if err != nil && err.Error() != rd.ErrNoMatch {
		test.Errorf("Run should've failed. Expected:%s Got:%s", rd.ErrNoMatch, err.Error())
	}
}
