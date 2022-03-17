package controller

import (
	"fmt"
	"io"
	"mime/multipart"
)

func init() {
	m := map[string]interface{}{

		"/exam/error": func() (interface{}, error) {
			return srv.Error()
		},

		"/exam/ioreader": func(r io.ReadCloser) (interface{}, error) {
			defer r.Close()
			b, err := io.ReadAll(r)
			fmt.Printf("%s", b)
			return nil, err
		},

		"/exam/multipart": func(r *multipart.Form) (interface{}, error) {
			fmt.Printf("%v", r)
			return nil, nil
		},
	}

	for p, h := range m {
		controllers[p] = h
	}
}
