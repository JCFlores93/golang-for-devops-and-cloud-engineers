package api

import "testing"

type MockClient struct {
	ResponseOutput *http.Response
}

func (m MockClient) Get(url string) (resp *http.Response, err error) {
	return m.ResponseOutput
}

func TestDoGetRequest(t *testing.T) {
	words := WordsPage{
		Page: Page{"words"},
		Words: Words{
			Input: "abc",
			Words: []string{"a", "b"},
		},
	}

	wordsBytes, err := json.Marshal(words)
	if err != nil {
		t.Error("marshal error: %s", err)
	}

	apiInstance := api{
		Options: Options{},
		Client: MockClient{
			ResponseOutput: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(wordsBytes)),
			},
		},
	}
	response, err := apiInstance.DoRequest("htpp://localhost/words")
	if err != nil {
		t.Error("DoRequest error: %s", err)
	}
	if response == nil {
		t.Fatalf("Error is empty")
	}
	if response.GetResponse() != strings.Join([]string{"a", "b"}, ", ") {
		t.Errorf("Unexpected response: %s", response.GetResponse())
	}

}
