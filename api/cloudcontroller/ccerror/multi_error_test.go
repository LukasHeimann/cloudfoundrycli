package ccerror_test

import (
	"net/http"

	. "github.com/LukasHeimann/cloudfoundrycli/v8/api/cloudcontroller/ccerror"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MultiError", func() {
	Describe("Error", func() {
		It("returns all errors", func() {
			err := MultiError{
				ResponseCode: http.StatusTeapot,
				Errors: []V3Error{
					{
						Code:   282010,
						Detail: "detail 1",
						Title:  "title-1",
					},
					{
						Code:   10242013,
						Detail: "detail 2",
						Title:  "title-2",
					},
				},
			}

			Expect(err).To(MatchError(`Multiple errors occurred:
Response Code: 418
Code: 282010, Title: title-1, Detail: detail 1
Code: 10242013, Title: title-2, Detail: detail 2`))
		})
	})
})
