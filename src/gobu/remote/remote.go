package remote

import (
	"errors"
	"log"
	"strings"

	"github.com/anisus/query"
	req "github.com/levigross/grequests"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Remote helps instrument downloads
type Remote struct {
	data     string
	Versions []Version
}

// LatestVersion returns current stable version
func (r *Remote) LatestVersion() Version {
	return r.Versions[0]
}

// GetVersion returns selected version
func (r *Remote) GetVersion(release string) (Version, error) {
	for _, v := range r.Versions {
		if v.Release == release {
			return v, nil
		}
	}
	return Version{}, errors.New("release not found")
}

// Update fetches and updates version's for current GOOS and GOARCH
func (r *Remote) Update() {
	if r.data == "" {
		resp, err := req.Get(
			"https://golang.org/dl",
			&req.RequestOptions{
				UserAgent: "Mozilla/5.0 (compatible; Gobu; +https://github.com/dz0ny/gobu)",
			},
		)

		if err != nil {
			log.Fatalln("Unable to make request:", err)
		}
		r.data = resp.String()
	}
	for _, v := range r.parseDownloads(r.data) {
		if v.Compatible() {
			r.Versions = append(r.Versions, v)
		}
	}
}

func (r *Remote) parseDownloads(data string) []Version {

	versions := []Version{}
	body, err := html.Parse(strings.NewReader(data))
	if err != nil {
		log.Println(err)
		return versions
	}
	trs := query.Set{body}.Find(query.ByTag(atom.Tr))
	for _, trr := range trs {
		tr := query.Set{trr}
		td := tr.Find(query.ByTag(atom.Td))
		if t := td.Eq(1).Text(); t == "Archive" {

			a := tr.Find(query.ByTag(atom.A))
			hash := tr.Find(query.ByTag(atom.Tt)).Text()

			version := Version{
				remoteURL: a.Attr("href"),
				hash:      hash,
				name:      a.Text(),
				beta:      strings.Contains(a.Text(), "beta"),
			}

			if err = version.parseVersion(); err == nil {
				versions = append(versions, version)
			}

		}
	}
	return versions
}
