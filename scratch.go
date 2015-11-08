package scratch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type ProjectSearch struct {
	Fields   Fields `json:"fields"`
	ImageURL string `json:"image_url"`
	Model    string `json:"model"`
	Pk       int    `json:"pk"`
}

type Fields struct {
	Creator              Creator `json:"creator"`
	FavoriteCount        int     `json:"favorite_count"`
	LoveCount            int     `json:"love_count"`
	RemixersCount        int     `json:"remixers_count"`
	Thumbnail            string  `json:"thumbnail"`
	Title                string  `json:"title"`
	UncachedThumbnailURL string  `json:"uncached_thumbnail_url"`
	ViewCount            int     `json:"view_count"`
}

type Project struct {
	Creator        Creator `json:"creator"`
	DatetimeShared string  `json:"datetime_shared"`
	Description    string  `json:"description"`
	FavoriteCount  string  `json:"favorite_count"`
	ID             int     `json:"id"`
	LoveCount      string  `json:"love_count"`
	ResourceURI    string  `json:"resource_uri"`
	Thumbnail      string  `json:"thumbnail"`
	Title          string  `json:"title"`
	ViewCount      string  `json:"view_count"`
}

type Creator struct {
	Username     string      `json:"username"`
	Userprofile  Userprofile `json:"userprofile"`
	Admin        bool        `json:"admin"`
	Pk           int         `json:"pk"`
	ThumbnailURL string      `json:"thumbnail_url"`
}

type Userprofile struct {
	Bio     string `json:"bio"`
	Country string `json:"country"`
	Status  string `json:"status"`
}

func GetProject(id string) (*Project, error) {
	url := "https://scratch.mit.edu/api/v1/project/" + id + "/?format=json"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	p := new(Project)
	err = json.Unmarshal(body, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func GetProjects(tag string) ([]*Project, error) {

	url := "https://scratch.mit.edu/site-api/explore/projects/" + tag + "/?date=this_month"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ps := make([]ProjectSearch, 0)
	err = json.Unmarshal(body, &ps)
	if err != nil {
		return nil, err
	}

	projects := []*Project{}

	for _, project := range ps {
		id := project.ImageURL
		id = strings.TrimPrefix(id, "//cdn2.scratch.mit.edu/get_image/project/")
		id = strings.Split(id, "_")[0]

		p, err := GetProject(id)
		if err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, nil

}

//<img src="//cdn2.scratch.mit.edu/get_image/project/86658772_216x163.png?v=1446739374.02" class="image">ck
