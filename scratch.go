package scratch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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
	Username    string      `json:"username"`
	Userprofile Userprofile `json:"userprofile"`
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

//<img src="//cdn2.scratch.mit.edu/get_image/project/86658772_216x163.png?v=1446739374.02" class="image">ck
