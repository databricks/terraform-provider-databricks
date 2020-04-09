package model

type PyPi struct {
	Package string `json:"package,omitempty"`
	Repo    string `json:"repo,omitempty"`
}

type Maven struct {
	Coordinates string   `json:"coordinates,omitempty"`
	Repo        string   `json:"repo,omitempty"`
	Exclusions  []string `json:"exclusions,omitempty"`
}

type Cran struct {
	Package string `json:"package,omitempty"`
	Repo    string `json:"repo,omitempty"`
}

type Library struct {
	Jar   string `json:"jar,omitempty"`
	Egg   string `json:"egg,omitempty"`
	Whl   string `json:"whl,omitempty"`
	Pypi  *PyPi  `json:"pypi,omitempty"`
	Maven *Maven `json:"maven,omitempty"`
	Cran  *Cran  `json:"cran,omitempty"`
}

type LibraryStatus struct {
	Library                         *Library `json:"library,omitempty"`
	Status                          string   `json:"status,omitempty"`
	IsLibraryInstalledOnAllClusters bool     `json:"is_library_for_all_clusters,omitempty"`
	Messages                        []string `json:"messages,omitempty"`
}
