# httptester

`httptester` provides a format for running related tests in a loop. This encourages negative testing and can help facilitate better coverage.

## Add to your project
```bash
go get github.com/dfirebaugh/httptester
```

## Example

```go

func TestGetEntity(t *testing.T) {
	comicRepo := in_memory.Repo[entity.Comic]{}
	comicService := service.New[entity.Comic](&comicRepo)
	comicHandler := v1.EntityHandler[entity.Comic]{
		comicService,
		responder.Responder{},
	}

	tests := []httptester.HTTPTest{
		{
			Method: http.MethodGet,
			URL:    "/v1/user",
			Body:   nil,
			Tests: []httptester.TestCase{
				{
					Post: func(res *httptest.ResponseRecorder) {
						// should respond with latest
					},
				},
			},
		},
		{
			Method: http.MethodGet,
			URL:    "/v1/user?latest",
			Body:   nil,
			Tests: []httptester.TestCase{
				{
					Post: func(res *httptest.ResponseRecorder) {
						// should respond with latest
					},
				},
			},
		},
		{
			Method: http.MethodGet,
			URL:    "/v1/user?first",
			Body:   nil,
			Tests: []httptester.TestCase{
				{
					Post: func(res *httptest.ResponseRecorder) {
						// should respond with first
					},
				},
			},
		},
		{
			Method: http.MethodGet,
			URL:    "/v1/user?start=01-02-1988",
			Body:   nil,
			Tests: []httptester.TestCase{
				{
					Post: func(res *httptest.ResponseRecorder) {
						// should fail because end is require
					},
				},
			},
		},
		{
			Method: http.MethodGet,
			URL:    "/v1/user?end=01-02-1988",
			Body:   nil,
			Tests: []httptester.TestCase{
				{
					Post: func(res *httptest.ResponseRecorder) {
						// should fail because start is require
					},
				},
			},
		},
		{
			Method: http.MethodGet,
			URL:    "/v1/user?start=01-02-1988end=02-29-2022",
			Body:   nil,
			Tests: []httptester.TestCase{
				{
					Post: func(res *httptest.ResponseRecorder) {
						// should fail because it's an invalid date
					},
				},
			},
		},
		{
			Method: http.MethodGet,
			URL:    "/v1/user?start=01-02-1988end=02-29-2022",
			Body:   nil,
			Tests: []httptester.TestCase{
				{
					Post: func(res *httptest.ResponseRecorder) {
						// should respond with comics in the valid range
					},
				},
			},
		},
	}

	for _, el := range tests {
		el.Execute(t, comicHandler.Get)
	}
}

```