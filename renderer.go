package markdown

import(
    "io/ioutil"
    "github.com/russross/blackfriday"
)

type Markdown struct {
    cache map[string]string
}

var this *Markdown

// Render the given file as markdown.
func Render(f string) (string) {

    if this == nil {
        this = &Markdown{}
        Clean()
    }

    c, ok := this.cache[f]

    if ok == false {
        d, err := ioutil.ReadFile(f)
        if err != nil {
            return ""
        }
        s := blackfriday.MarkdownBasic(d)
        this.cache[f] = string(s)
        c = this.cache[f]
    }

    return c
}

// Clean the cache.
func Clean() {
    this.cache = map[string]string{}
}
