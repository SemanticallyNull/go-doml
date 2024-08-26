package doml

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestDOML(t *testing.T) {
	c, err := Parse("🏁✌️key✌️▶️value◀️✌️array_key✌️⏩▶️Obnoxious◀️▶️language◀️⏪🛑")
	assert.NilError(t, err)

	assert.DeepEqual(t, c, map[string]any{
		"array_key": []string{"Obnoxious", "language"},
		"key":       "value",
	})
}

func TestDOMLWhitespace(t *testing.T) {
	c, err := Parse(` 🏁
	✌️key✌️
		▶️value◀️
    ✌️array_key✌️	⏩
		▶️Obnoxious◀️
		▶️language◀️
     ⏪
	✌️key✌️
		▶️value◀️
🛑

`)
	assert.NilError(t, err)

	assert.DeepEqual(t, c, map[string]any{
		"array_key": []string{"Obnoxious", "language"},
		"key":       "value",
	})
}
