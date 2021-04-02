// Generated by tomltestgen for toml-test ref 39e37e6 on 2019-03-19T23:58:45-07:00
package toml_test

import (
	"testing"
)

func TestInvalidDatetimeMalformedNoLeads(t *testing.T) {
	input := `no-leads = 1987-7-05T17:45:00Z`
	testgenInvalid(t, input)
}

func TestInvalidDatetimeMalformedNoSecs(t *testing.T) {
	input := `no-secs = 1987-07-05T17:45Z`
	testgenInvalid(t, input)
}

func TestInvalidDatetimeMalformedNoT(t *testing.T) {
	input := `no-t = 1987-07-0517:45:00Z`
	testgenInvalid(t, input)
}

func TestInvalidDatetimeMalformedWithMilli(t *testing.T) {
	input := `with-milli = 1987-07-5T17:45:00.12Z`
	testgenInvalid(t, input)
}

func TestInvalidDuplicateKeyTable(t *testing.T) {
	input := `[fruit]
type = "apple"

[fruit.type]
apple = "yes"`
	testgenInvalid(t, input)
}

func TestInvalidDuplicateKeys(t *testing.T) {
	input := `dupe = false
dupe = true`
	testgenInvalid(t, input)
}

func TestInvalidDuplicateTables(t *testing.T) {
	input := `[a]
[a]`
	testgenInvalid(t, input)
}

func TestInvalidEmptyImplicitTable(t *testing.T) {
	input := `[naughty..naughty]`
	testgenInvalid(t, input)
}

func TestInvalidEmptyTable(t *testing.T) {
	input := `[]`
	testgenInvalid(t, input)
}

func TestInvalidFloatNoLeadingZero(t *testing.T) {
	input := `answer = .12345
neganswer = -.12345`
	testgenInvalid(t, input)
}

func TestInvalidFloatNoTrailingDigits(t *testing.T) {
	input := `answer = 1.
neganswer = -1.`
	testgenInvalid(t, input)
}

func TestInvalidKeyEmpty(t *testing.T) {
	input := ` = 1`
	testgenInvalid(t, input)
}

func TestInvalidKeyHash(t *testing.T) {
	input := `a# = 1`
	testgenInvalid(t, input)
}

func TestInvalidKeyNewline(t *testing.T) {
	input := `a
= 1`
	testgenInvalid(t, input)
}

func TestInvalidKeyOpenBracket(t *testing.T) {
	input := `[abc = 1`
	testgenInvalid(t, input)
}

func TestInvalidKeySingleOpenBracket(t *testing.T) {
	input := `[`
	testgenInvalid(t, input)
}

func TestInvalidKeySpace(t *testing.T) {
	input := `a b = 1`
	testgenInvalid(t, input)
}

func TestInvalidKeyStartBracket(t *testing.T) {
	input := `[a]
[xyz = 5
[b]`
	testgenInvalid(t, input)
}

func TestInvalidKeyTwoEquals(t *testing.T) {
	input := `key= = 1`
	testgenInvalid(t, input)
}

func TestInvalidStringBadByteEscape(t *testing.T) {
	input := `naughty = "\xAg"`
	testgenInvalid(t, input)
}

func TestInvalidStringBadEscape(t *testing.T) {
	input := `invalid-escape = "This string has a bad \a escape character."`
	testgenInvalid(t, input)
}

func TestInvalidStringByteEscapes(t *testing.T) {
	input := `answer = "\x33"`
	testgenInvalid(t, input)
}

func TestInvalidStringNoClose(t *testing.T) {
	input := `no-ending-quote = "One time, at band camp`
	testgenInvalid(t, input)
}

func TestInvalidTableArrayImplicit(t *testing.T) {
	input := "# This test is a bit tricky. It should fail because the first use of\n" +
		"# `[[albums.songs]]` without first declaring `albums` implies that `albums`\n" +
		"# must be a table. The alternative would be quite weird. Namely, it wouldn't\n" +
		"# comply with the TOML spec: \"Each double-bracketed sub-table will belong to \n" +
		"# the most *recently* defined table element *above* it.\"\n" +
		"#\n" +
		"# This is in contrast to the *valid* test, table-array-implicit where\n" +
		"# `[[albums.songs]]` works by itself, so long as `[[albums]]` isn't declared\n" +
		"# later. (Although, `[albums]` could be.)\n" +
		"[[albums.songs]]\n" +
		"name = \"Glory Days\"\n" +
		"\n" +
		"[[albums]]\n" +
		"name = \"Born in the USA\"\n"
	testgenInvalid(t, input)
}

func TestInvalidTableArrayMalformedBracket(t *testing.T) {
	input := `[[albums]
name = "Born to Run"`
	testgenInvalid(t, input)
}

func TestInvalidTableArrayMalformedEmpty(t *testing.T) {
	input := `[[]]
name = "Born to Run"`
	testgenInvalid(t, input)
}

func TestInvalidTableEmpty(t *testing.T) {
	input := `[]`
	testgenInvalid(t, input)
}

func TestInvalidTableNestedBracketsClose(t *testing.T) {
	input := `[a]b]
zyx = 42`
	testgenInvalid(t, input)
}

func TestInvalidTableNestedBracketsOpen(t *testing.T) {
	input := `[a[b]
zyx = 42`
	testgenInvalid(t, input)
}

func TestInvalidTableWhitespace(t *testing.T) {
	input := `[invalid key]`
	testgenInvalid(t, input)
}

func TestInvalidTableWithPound(t *testing.T) {
	input := `[key#group]
answer = 42`
	testgenInvalid(t, input)
}

func TestInvalidTextAfterArrayEntries(t *testing.T) {
	input := `array = [
  "Is there life after an array separator?", No
  "Entry"
]`
	testgenInvalid(t, input)
}

func TestInvalidTextAfterInteger(t *testing.T) {
	input := `answer = 42 the ultimate answer?`
	testgenInvalid(t, input)
}

func TestInvalidTextAfterString(t *testing.T) {
	input := `string = "Is there life after strings?" No.`
	testgenInvalid(t, input)
}

func TestInvalidTextAfterTable(t *testing.T) {
	input := `[error] this shouldn't be here`
	testgenInvalid(t, input)
}

func TestInvalidTextBeforeArraySeparator(t *testing.T) {
	input := `array = [
  "Is there life before an array separator?" No,
  "Entry"
]`
	testgenInvalid(t, input)
}

func TestInvalidTextInArray(t *testing.T) {
	input := `array = [
  "Entry 1",
  I don't belong,
  "Entry 2",
]`
	testgenInvalid(t, input)
}

func TestValidArrayEmpty(t *testing.T) {
	input := `thevoid = [[[[[]]]]]`
	jsonRef := `{
    "thevoid": { "type": "array", "value": [
        {"type": "array", "value": [
            {"type": "array", "value": [
                {"type": "array", "value": [
                    {"type": "array", "value": []}
                ]}
            ]}
        ]}
    ]}
}`
	testgenValid(t, input, jsonRef)
}

func TestValidArrayNospaces(t *testing.T) {
	input := `ints = [1,2,3]`
	jsonRef := `{
    "ints": {
        "type": "array",
        "value": [
            {"type": "integer", "value": "1"},
            {"type": "integer", "value": "2"},
            {"type": "integer", "value": "3"}
        ]
    }
}`
	testgenValid(t, input, jsonRef)
}

func TestValidArraysHetergeneous(t *testing.T) {
	input := `mixed = [[1, 2], ["a", "b"], [1.1, 2.1]]`
	jsonRef := `{
    "mixed": {
        "type": "array",
        "value": [
            {"type": "array", "value": [
                {"type": "integer", "value": "1"},
                {"type": "integer", "value": "2"}
            ]},
            {"type": "array", "value": [
                {"type": "string", "value": "a"},
                {"type": "string", "value": "b"}
            ]},
            {"type": "array", "value": [
                {"type": "float", "value": "1.1"},
                {"type": "float", "value": "2.1"}
            ]}
        ]
    }
}`
	testgenValid(t, input, jsonRef)
}

func TestValidArraysNested(t *testing.T) {
	input := `nest = [["a"], ["b"]]`
	jsonRef := `{
    "nest": {
        "type": "array",
        "value": [
            {"type": "array", "value": [
                {"type": "string", "value": "a"}
            ]},
            {"type": "array", "value": [
                {"type": "string", "value": "b"}
            ]}
        ]
    }
}`
	testgenValid(t, input, jsonRef)
}

func TestValidArrays(t *testing.T) {
	input := `ints = [1, 2, 3]
floats = [1.1, 2.1, 3.1]
strings = ["a", "b", "c"]
dates = [
  1987-07-05T17:45:00Z,
  1979-05-27T07:32:00Z,
  2006-06-01T11:00:00Z,
]`
	jsonRef := `{
    "ints": {
        "type": "array",
        "value": [
            {"type": "integer", "value": "1"},
            {"type": "integer", "value": "2"},
            {"type": "integer", "value": "3"}
        ]
    },
    "floats": {
        "type": "array",
        "value": [
            {"type": "float", "value": "1.1"},
            {"type": "float", "value": "2.1"},
            {"type": "float", "value": "3.1"}
        ]
    },
    "strings": {
        "type": "array",
        "value": [
            {"type": "string", "value": "a"},
            {"type": "string", "value": "b"},
            {"type": "string", "value": "c"}
        ]
    },
    "dates": {
        "type": "array",
        "value": [
            {"type": "datetime", "value": "1987-07-05T17:45:00Z"},
            {"type": "datetime", "value": "1979-05-27T07:32:00Z"},
            {"type": "datetime", "value": "2006-06-01T11:00:00Z"}
        ]
    }
}`
	testgenValid(t, input, jsonRef)
}

func TestValidBool(t *testing.T) {
	input := `t = true
f = false`
	jsonRef := `{
    "f": {"type": "bool", "value": "false"},
    "t": {"type": "bool", "value": "true"}
}`
	testgenValid(t, input, jsonRef)
}

func TestValidCommentsEverywhere(t *testing.T) {
	input := `# Top comment.
  # Top comment.
# Top comment.

# [no-extraneous-groups-please]

[group] # Comment
answer = 42 # Comment
# no-extraneous-keys-please = 999
# Between comment.
more = [ # Comment
  # What about multiple # comments?
  # Can you handle it?
  #
          # Evil.
# Evil.
  42, 42, # Comments within arrays are fun.
  # What about multiple # comments?
  # Can you handle it?
  #
          # Evil.
# Evil.
# ] Did I fool you?
] # Hopefully not.`
	jsonRef := `{
    "group": {
        "answer": {"type": "integer", "value": "42"},
        "more": {
            "type": "array",
            "value": [
                {"type": "integer", "value": "42"},
                {"type": "integer", "value": "42"}
            ]
        }
    }
}`
	testgenValid(t, input, jsonRef)
}

func TestValidDatetime(t *testing.T) {
	input := `bestdayever = 1987-07-05T17:45:00Z`
	jsonRef := `{
    "bestdayever": {"type": "datetime", "value": "1987-07-05T17:45:00Z"}
}`
	testgenValid(t, input, jsonRef)
}

func TestValidEmpty(t *testing.T) {
	input := ``
	jsonRef := `{}`
	testgenValid(t, input, jsonRef)
}

func TestValidExample(t *testing.T) {
	input := `best-day-ever = 1987-07-05T17:45:00Z

[numtheory]
boring = false
perfection = [6, 28, 496]`
	jsonRef := `{
  "best-day-ever": {"type": "datetime", "value": "1987-07-05T17:45:00Z"},
  "numtheory": {
    "boring": {"type": "bool", "value": "false"},
    "perfection": {
      "type": "array",
      "value": [
        {"type": "integer", "value": "6"},
        {"type": "integer", "value": "28"},
        {"type": "integer", "value": "496"}
      ]
    }
  }
}`
	testgenValid(t, input, jsonRef)
}

func TestValidFloat(t *testing.T) {
	input := `pi = 3.14
negpi = -3.14`
	jsonRef := `{
    "pi": {"type": "float", "value": "3.14"},
    "negpi": {"type": "float", "value": "-3.14"}
}`
	testgenValid(t, input, jsonRef)
}

func TestValidImplicitAndExplicitAfter(t *testing.T) {
	input := `[a.b.c]
answer = 42

[a]
better = 43`
	jsonRef := `{
    "a": {
        "better": {"type": "integer", "value": "43"},
        "b": {
            "c": {
                "answer": {"type": "integer", "value": "42"}
            }
        }
    }
}`
	testgenValid(t, input, jsonRef)
}

func TestValidImplicitAndExplicitBefore(t *testing.T) {
	input := `[a]
better = 43

[a.b.c]
answer = 42`
	jsonRef := `{
    "a": {
        "better": {"type": "integer", "value": "43"},
        "b": {
            "c": {
                "answer": {"type": "integer", "value": "42"}
            }
        }
    }
}`
	testgenValid(t, input, jsonRef)
}

func TestValidImplicitGroups(t *testing.T) {
	input := `[a.b.c]
answer = 42`
	jsonRef := `{
    "a": {
        "b": {
            "c": {
                "answer": {"type": "integer", "value": "42"}
            }
        }
    }
}`
	testgenValid(t, input, jsonRef)
}

func TestValidInteger(t *testing.T) {
	input := `answer = 42
neganswer = -42`
	jsonRef := `{
    "answer": {"type": "integer", "value": "42"},
    "neganswer": {"type": "integer", "value": "-42"}
}`
	testgenValid(t, input, jsonRef)
}

func TestValidKeyEqualsNospace(t *testing.T) {
	input := `answer=42`
	jsonRef := `{
    "answer": {"type": "integer", "value": "42"}
}`
	testgenValid(t, input, jsonRef)
}

func TestValidKeySpace(t *testing.T) {
	input := `"a b" = 1`
	jsonRef := `{
    "a b": {"type": "integer", "value": "1"}
}`
	testgenValid(t, input, jsonRef)
}

func TestValidKeySpecialChars(t *testing.T) {
	input := "\"~!@$^&*()_+-`1234567890[]|/?><.,;:'\" = 1\n"
	jsonRef := "{\n" +
		"    \"~!@$^&*()_+-`1234567890[]|/?><.,;:'\": {\n" +
		"        \"type\": \"integer\", \"value\": \"1\"\n" +
		"    }\n" +
		"}\n"
	testgenValid(t, input, jsonRef)
}

func TestValidLongFloat(t *testing.T) {
	input := `longpi = 3.141592653589793
neglongpi = -3.141592653589793`
	jsonRef := `{
    "longpi": {"type": "float", "value": "3.141592653589793"},
    "neglongpi": {"type": "float", "value": "-3.141592653589793"}
}`
	testgenValid(t, input, jsonRef)
}

func TestValidLongInteger(t *testing.T) {
	input := `answer = 9223372036854775807
neganswer = -9223372036854775808`
	jsonRef := `{
    "answer": {"type": "integer", "value": "9223372036854775807"},
    "neganswer": {"type": "integer", "value": "-9223372036854775808"}
}`
	testgenValid(t, input, jsonRef)
}

func TestValidMultilineString(t *testing.T) {
	input := `multiline_empty_one = """"""
multiline_empty_two = """
"""
multiline_empty_three = """\
    """
multiline_empty_four = """\
   \
   \
   """

equivalent_one = "The quick brown fox jumps over the lazy dog."
equivalent_two = """
The quick brown \


  fox jumps over \
    the lazy dog."""

equivalent_three = """\
       The quick brown \
       fox jumps over \
       the lazy dog.\
       """`
	jsonRef := `{
    "multiline_empty_one": {
        "type": "string",
        "value": ""
    },
    "multiline_empty_two": {
        "type": "string",
        "value": ""
    },
    "multiline_empty_three": {
        "type": "string",
        "value": ""
    },
    "multiline_empty_four": {
        "type": "string",
        "value": ""
    },
    "equivalent_one": {
        "type": "string",
        "value": "The quick brown fox jumps over the lazy dog."
    },
    "equivalent_two": {
        "type": "string",
        "value": "The quick brown fox jumps over the lazy dog."
    },
    "equivalent_three": {
        "type": "string",
        "value": "The quick brown fox jumps over the lazy dog."
    }
}`
	testgenValid(t, input, jsonRef)
}

func TestValidRawMultilineString(t *testing.T) {
	input := `oneline = '''This string has a ' quote character.'''
firstnl = '''
This string has a ' quote character.'''
multiline = '''
This string
has ' a quote character
and more than
one newline
in it.'''`
	jsonRef := `{
    "oneline": {
        "type": "string",
        "value": "This string has a ' quote character."
    },
    "firstnl": {
        "type": "string",
        "value": "This string has a ' quote character."
    },
    "multiline": {
        "type": "string",
        "value": "This string\nhas ' a quote character\nand more than\none newline\nin it."
    }
}`
	testgenValid(t, input, jsonRef)
}

func TestValidRawString(t *testing.T) {
	input := `backspace = 'This string has a \b backspace character.'
tab = 'This string has a \t tab character.'
newline = 'This string has a \n new line character.'
formfeed = 'This string has a \f form feed character.'
carriage = 'This string has a \r carriage return character.'
slash = 'This string has a \/ slash character.'
backslash = 'This string has a \\ backslash character.'`
	jsonRef := `{
    "backspace": {
        "type": "string",
        "value": "This string has a \\b backspace character."
    },
    "tab": {
        "type": "string",
        "value": "This string has a \\t tab character."
    },
    "newline": {
        "type": "string",
        "value": "This string has a \\n new line character."
    },
    "formfeed": {
        "type": "string",
        "value": "This string has a \\f form feed character."
    },
    "carriage": {
        "type": "string",
        "value": "This string has a \\r carriage return character."
    },
    "slash": {
        "type": "string",
        "value": "This string has a \\/ slash character."
    },
    "backslash": {
        "type": "string",
        "value": "This string has a \\\\ backslash character."
    }
}`
	testgenValid(t, input, jsonRef)
}

func TestValidStringEmpty(t *testing.T) {
	input := `answer = ""`
	jsonRef := `{
    "answer": {
        "type": "string",
        "value": ""
    }
}`
	testgenValid(t, input, jsonRef)
}

func TestValidStringEscapes(t *testing.T) {
	input := `backspace = "This string has a \b backspace character."
tab = "This string has a \t tab character."
newline = "This string has a \n new line character."
formfeed = "This string has a \f form feed character."
carriage = "This string has a \r carriage return character."
quote = "This string has a \" quote character."
backslash = "This string has a \\ backslash character."
notunicode1 = "This string does not have a unicode \\u escape."
notunicode2 = "This string does not have a unicode \u005Cu escape."
notunicode3 = "This string does not have a unicode \\u0075 escape."
notunicode4 = "This string does not have a unicode \\\u0075 escape."`
	jsonRef := `{
    "backspace": {
        "type": "string",
        "value": "This string has a \u0008 backspace character."
    },
    "tab": {
        "type": "string",
        "value": "This string has a \u0009 tab character."
    },
    "newline": {
        "type": "string",
        "value": "This string has a \u000A new line character."
    },
    "formfeed": {
        "type": "string",
        "value": "This string has a \u000C form feed character."
    },
    "carriage": {
        "type": "string",
        "value": "This string has a \u000D carriage return character."
    },
    "quote": {
        "type": "string",
        "value": "This string has a \u0022 quote character."
    },
    "backslash": {
        "type": "string",
        "value": "This string has a \u005C backslash character."
    },
    "notunicode1": {
        "type": "string",
        "value": "This string does not have a unicode \\u escape."
    },
    "notunicode2": {
        "type": "string",
        "value": "This string does not have a unicode \u005Cu escape."
    },
    "notunicode3": {
        "type": "string",
        "value": "This string does not have a unicode \\u0075 escape."
    },
    "notunicode4": {
        "type": "string",
        "value": "This string does not have a unicode \\\u0075 escape."
    }
}`
	testgenValid(t, input, jsonRef)
}

func TestValidStringSimple(t *testing.T) {
	input := `answer = "You are not drinking enough whisky."`
	jsonRef := `{
    "answer": {
        "type": "string",
        "value": "You are not drinking enough whisky."
    }
}`
	testgenValid(t, input, jsonRef)
}

func TestValidStringWithPound(t *testing.T) {
	input := `pound = "We see no # comments here."
poundcomment = "But there are # some comments here." # Did I # mess you up?`
	jsonRef := `{
    "pound": {"type": "string", "value": "We see no # comments here."},
    "poundcomment": {
        "type": "string",
        "value": "But there are # some comments here."
    }
}`
	testgenValid(t, input, jsonRef)
}

func TestValidTableArrayImplicit(t *testing.T) {
	input := `[[albums.songs]]
name = "Glory Days"`
	jsonRef := `{
    "albums": {
       "songs": [
           {"name": {"type": "string", "value": "Glory Days"}}
       ]
    }
}`
	testgenValid(t, input, jsonRef)
}

func TestValidTableArrayMany(t *testing.T) {
	input := `[[people]]
first_name = "Bruce"
last_name = "Springsteen"

[[people]]
first_name = "Eric"
last_name = "Clapton"

[[people]]
first_name = "Bob"
last_name = "Seger"`
	jsonRef := `{
    "people": [
        {
            "first_name": {"type": "string", "value": "Bruce"},
            "last_name": {"type": "string", "value": "Springsteen"}
        },
        {
            "first_name": {"type": "string", "value": "Eric"},
            "last_name": {"type": "string", "value": "Clapton"}
        },
        {
            "first_name": {"type": "string", "value": "Bob"},
            "last_name": {"type": "string", "value": "Seger"}
        }
    ]
}`
	testgenValid(t, input, jsonRef)
}

func TestValidTableArrayNest(t *testing.T) {
	input := `[[albums]]
name = "Born to Run"

  [[albums.songs]]
  name = "Jungleland"

  [[albums.songs]]
  name = "Meeting Across the River"

[[albums]]
name = "Born in the USA"
  
  [[albums.songs]]
  name = "Glory Days"

  [[albums.songs]]
  name = "Dancing in the Dark"`
	jsonRef := `{
    "albums": [
        {
            "name": {"type": "string", "value": "Born to Run"},
            "songs": [
                {"name": {"type": "string", "value": "Jungleland"}},
                {"name": {"type": "string", "value": "Meeting Across the River"}}
            ]
        },
        {
            "name": {"type": "string", "value": "Born in the USA"},
            "songs": [
                {"name": {"type": "string", "value": "Glory Days"}},
                {"name": {"type": "string", "value": "Dancing in the Dark"}}
            ]
        }
    ]
}`
	testgenValid(t, input, jsonRef)
}

func TestValidTableArrayOne(t *testing.T) {
	input := `[[people]]
first_name = "Bruce"
last_name = "Springsteen"`
	jsonRef := `{
    "people": [
        {
            "first_name": {"type": "string", "value": "Bruce"},
            "last_name": {"type": "string", "value": "Springsteen"}
        }
    ]
}`
	testgenValid(t, input, jsonRef)
}

func TestValidTableEmpty(t *testing.T) {
	input := `[a]`
	jsonRef := `{
    "a": {}
}`
	testgenValid(t, input, jsonRef)
}

func TestValidTableSubEmpty(t *testing.T) {
	input := `[a]
[a.b]`
	jsonRef := `{
    "a": { "b": {} }
}`
	testgenValid(t, input, jsonRef)
}

func TestValidTableWhitespace(t *testing.T) {
	input := `["valid key"]`
	jsonRef := `{
    "valid key": {}
}`
	testgenValid(t, input, jsonRef)
}

func TestValidTableWithPound(t *testing.T) {
	input := `["key#group"]
answer = 42`
	jsonRef := `{
    "key#group": {
        "answer": {"type": "integer", "value": "42"}
    }
}`
	testgenValid(t, input, jsonRef)
}

func TestValidUnicodeEscape(t *testing.T) {
	input := `answer4 = "\u03B4"
answer8 = "\U000003B4"`
	jsonRef := `{
    "answer4": {"type": "string", "value": "\u03B4"},
    "answer8": {"type": "string", "value": "\u03B4"}
}`
	testgenValid(t, input, jsonRef)
}

func TestValidUnicodeLiteral(t *testing.T) {
	input := `answer = "δ"`
	jsonRef := `{
    "answer": {"type": "string", "value": "δ"}
}`
	testgenValid(t, input, jsonRef)
}
