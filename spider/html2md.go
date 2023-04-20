package spider

import (
	"bytes"
	"regexp"
	"strings"
	"unicode/utf8"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func getCodeWithoutTags(startNode *html.Node) []byte {
	var buf bytes.Buffer

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && (n.Data == "style" || n.Data == "script" || n.Data == "textarea") {
			return
		}
		if n.Type == html.ElementNode && (n.Data == "br" || n.Data == "div") {
			buf.WriteString("\n")
		}

		if n.Type == html.TextNode {
			buf.WriteString(n.Data)
			return
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(startNode)

	return buf.Bytes()
}

func getCodeContent(selec *goquery.Selection) string {
	if len(selec.Nodes) == 0 {
		return ""
	}

	code := getCodeWithoutTags(selec.Nodes[0])

	return string(code)
}

func html2md(html string) (res string, err error) {
	converter := md.NewConverter("", true, nil)
	converter.AddRules(md.Rule{
		Filter: []string{"pre"},
		Replacement: func(content string, selec *goquery.Selection, options *md.Options) *string {
			codeElement := selec.Find("code")
			language := codeElement.AttrOr("class", "")
			re := regexp.MustCompile("language-(\\w*)")
			language = string(re.FindSubmatch([]byte(language))[1])

			code := getCodeContent(selec)

			fenceChar, _ := utf8.DecodeRuneInString(options.Fence)
			fence := md.CalculateCodeFence(fenceChar, code)

			text := "\n\n" + fence + language + "\n" +
				code +
				"\n" + fence + "\n\n"
			return &text
		},
	})
	converter.After(func(markdown string) string {
		return strings.ReplaceAll(markdown, "0x0A", "\n")
	})
	res, err = converter.ConvertString(html)
	if err != nil {
		return
	}
	return
}
