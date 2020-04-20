package model

// PyPi is a python library hosted on PYPI
type PyPi struct {
	Package string `json:"package,omitempty"`
	Repo    string `json:"repo,omitempty"`
}

// Maven is a jar library hosted on Maven
type Maven struct {
	Coordinates string   `json:"coordinates,omitempty"`
	Repo        string   `json:"repo,omitempty"`
	Exclusions  []string `json:"exclusions,omitempty"`
}

// Cran is a R library hosted on Maven
type Cran struct {
	Package string `json:"package,omitempty"`
	Repo    string `json:"repo,omitempty"`
}

// Library is a construct that contains information of the location of the library and how to download it
type Library struct {
	Jar   string `json:"jar,omitempty"`
	Egg   string `json:"egg,omitempty"`
	Whl   string `json:"whl,omitempty"`
	Pypi  *PyPi  `json:"pypi,omitempty"`
	Maven *Maven `json:"maven,omitempty"`
	Cran  *Cran  `json:"cran,omitempty"`
}

// LibraryStatus is the status on a given cluster when using the libraries status api
type LibraryStatus struct {
	Library                         *Library `json:"library,omitempty"`
	Status                          string   `json:"status,omitempty"`
	IsLibraryInstalledOnAllClusters bool     `json:"is_library_for_all_clusters,omitempty"`
	Messages                        []string `json:"messages,omitempty"`
}
