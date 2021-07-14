package nvd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type NVD struct{}

func New() *NVD {
	return new(NVD)
}

func (s *NVD) Start(cve string, verbose bool) {
	var (
		client      *http.Client
		resp        *http.Response
		err         error
		description string
	)
	client = new(http.Client)
	url := fmt.Sprintf("https://nvd.nist.gov/vuln/detail/%s", cve)
	if verbose {
		log.Printf("Requesting %s\n", url)
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "CVESearch")
	if resp, err = client.Do(req); err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		fmt.Printf("Cannot find %s", cve)
    return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".vuln-description").Each(func(i int, sel *goquery.Selection) {
		// <p data-testid="vuln-analysis-description"></p>
    description = sel.Text()
	})

	fmt.Printf("%s\n", description)
}
