package markdown

import (
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestRender(t *testing.T) {

	Describe("Render()", func() {

		It("should return <h1>title</h1>", func() {
			s := Render("./fixtures/test.md")
			AssertEqual(s, "<h1>title</h1>\n")
		})
	})

	Report(t)
}
