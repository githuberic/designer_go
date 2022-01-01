package e2

import "fmt"

type HTTPDownloader struct {
	*template
}

func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *HTTPDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}

func (*HTTPDownloader) save() {
	fmt.Printf("http save\n")
}
