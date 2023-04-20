package spider

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	VERSION = "1.0.0"
)

func fetch(url string) (html string, err error) {
	var (
		doc *goquery.Document
	)

	log.Printf("Fetch url %s \n", url)

	doc, err = goquery.NewDocument(url)
	if err != nil {
		return
	}
	html, err = doc.Find(".markdown-body").Eq(0).Html()
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(strings.NewReader(html))
	for scanner.Scan() {
		html += scanner.Text()
		html += "0x0A"
	}
	if err = scanner.Err(); err != nil {
		return
	}

	return
}

func Parse(url string, out string) {
	html, err := fetch(url)
	if err != nil {
		log.Fatalln(err)
	}
	md, err := html2md(html)
	if err != nil {
		log.Fatalln(err)
	}
	runPath, err := GetRunPath()
	if err != nil {
		log.Fatalln(err)
	}
	fileName := path.Join(runPath, out)
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = file.WriteString(md)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("解析完成，输出文件地址: ", fileName)
}
