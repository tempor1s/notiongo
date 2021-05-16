package notion

import "testing"

// TestPage_marshal will test marshalling functionality with mock Notion Page data
func TestPage_marshal(t *testing.T) {
	// test empty body
	testJSONMarshal(t, &Page{}, "{}")

	// mock
	// TODO(imthaghost): create a function that will generate random member values
	p := &Page{
		Object:         "page",
		ID:             "0692b915742242bb988172b9705873d6",
		CreatedTime:    "2020-03-17T19:10:04.968Z",
		LastEditedTime: "2020-03-17T21:49:37.913Z",
	}

	// expected values
	want := `
		{
			"object": "page",
			"id": "0692b915742242bb988172b9705873d6",
			"created_time": "2020-03-17T19:10:04.968Z",
			"last_edited_time": "2020-03-17T21:49:37.913Z"
		}
	`

	testJSONMarshal(t, p, want)

}
