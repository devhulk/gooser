package main

type WhatsMyName struct {
	License    []string `json:"license"`
	Authors    []string `json:"authors"`
	Categories []string `json:"categories"`
	Sites      []struct {
		Name         string   `json:"name"`
		URICheck     string   `json:"uri_check"`
		URIPretty    string   `json:"uri_pretty"`
		PostBody     string   `json:"post_body"`
		StripBadChar string   `json:"strip_bad_char"`
		ECode        int      `json:"e_code"`
		EString      string   `json:"e_string"`
		MString      string   `json:"m_string"`
		MCode        int      `json:"m_code"`
		Known        []string `json:"known"`
		Cat          string   `json:"cat"`
		Valid        string   `json:"valid"`
		Headers      struct {
			Cookie string `json:"Cookie"`
			Accept string `json:"accept"`
		} `json:"headers"`
	} `json:"sites"`
}
