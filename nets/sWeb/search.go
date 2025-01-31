package sWeb

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sStrings"

	"golang.org/x/net/html"
)

type linksFinder struct {
	url     string
	dir     string
	fileExt []string
}

func NewLinksFinder(url, dir string, fileExt ...string) *linksFinder {
	return &linksFinder{
		url:     url,
		dir:     dir,
		fileExt: fileExt,
	}
}

// SearchLinks, use limit = -1 for nonstop
func (dp *linksFinder) SearchLinks(path_prefix string, patterns []string, download bool, limit int) int {

	url_pp := dp.url
	path_pp := dp.dir

	if len(path_prefix) > 0 {
		url_pp = fmt.Sprintf("%s/%s", url_pp, path_prefix)
		path_pp = fmt.Sprintf("%s/%s", path_pp, path_prefix)
	}

	page, err := parse(url_pp)
	if err != nil {
		sLog.Error("SearchLinks: Error getting page: %q  ..  err: %s \n", dp.url, err)
		return limit
	}

	links := pageLinks(nil, page)
	for _, link := range links {

		clearLink := strings.TrimRight(link, "/")

		file_url := fmt.Sprintf("%s/%s", url_pp, clearLink)

		if sStrings.FindSuffix(link, dp.fileExt...) {
			if sStrings.FindPatterns(link, patterns...) {
				if download {
					// log.Printf("Download: url: %s ", file_url)
					// log.Printf("Download: path: %s/%s ", path_pp, clearLink)

					err := DownloadFile(path_pp, clearLink, file_url)
					if err == nil {
						limit--
						sLog.Info("SearchLinks: Download success !!  file_url: %s ", file_url)
					} else {
						sLog.Error("SearchLinks: Error file_url: %s  ..  err: %s  ", file_url, err)
					}
				}
			}
		} else {
			// log.Printf("dir_url: %s \n", file_url)
			// log.Printf("dir_path: %s/%s \n\n", path_pp, clearLink)

			if len(path_prefix) > 0 {
				d := false
				if sStrings.FindPatterns(link, patterns...) {
					d = download
				}
				limit = dp.SearchLinks(fmt.Sprintf("%s/%s", path_prefix, clearLink),
					patterns, d, limit)
			} else {
				limit = dp.SearchLinks(clearLink, patterns, true, limit)
			}
		}

		if limit == 0 {
			break
		}
	}
	return limit
}

func findFile(url string, patterns ...string) bool {
	if sStrings.FindSuffix(url, ".gz", ".csv", ".zip") {
		return sStrings.FindPatterns(url, patterns...)
	}
	return false
}

func Search(base_url, path_prefix string, needle func(string) bool, limit int) ([]string, error) {
	// Formateo de URL base y path_prefix utilizando la librería `net/url` para mayor seguridad
	base, err := url.Parse(base_url)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %s", err)
	}

	if len(path_prefix) > 0 {
		prefix, err := url.Parse(path_prefix)
		if err != nil {
			return nil, fmt.Errorf("invalid path prefix: %s", err)
		}
		// Unir los segmentos de la ruta de forma segura
		base.Path = path.Join(append([]string{base.Path}, prefix.Path)...)
	}

	// Parsear la página
	page, err := parse(base.String())
	if err != nil {
		return nil, fmt.Errorf("parse(%s): %s", base.String(), err)
	}

	var urls []string
	links := pageLinks(nil, page)

	// Filtrar links
	for _, link := range links {
		if needle(link) {
			clear_link := strings.TrimRight(link, "/")

			file_url := base.ResolveReference(&url.URL{Path: clear_link}).String()

			sLog.Info("file_url: %s", file_url)

			urls = append(urls, file_url)

			if len(urls) >= limit {
				break
			}
		}
	}

	return urls, nil
}

func parse(page_url string) (*html.Node, error) {
	sLog.Info("Parse: %q", page_url)

	resp, err := http.Get(page_url)
	if err != nil {
		return nil, fmt.Errorf("cannot get page: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
	}

	node, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot parse page: %w", err)
	}

	return node, nil
}

func pageLinks(links []string, n *html.Node) []string {
	// Si el nodo es un elemento <a> y tiene un atributo href
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	// Recursión para procesar todos los hijos
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = pageLinks(links, c)
	}

	return links
}
