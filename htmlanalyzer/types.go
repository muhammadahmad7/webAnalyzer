package htmlanalyzer

type AnalysisReport struct {
	InputURL     string
	HtmlVersion string
	PageTitle string
	H1Count int
	H2Count int
	H3Count int
	H4Count int
	H5Count int
	H6Count int
	InternalLinksCount int
	ExternalLinksCount int
	InaccessibleLinksCount int
	IsLogin bool
	Error error

}

type analysisParam struct {
	InputURL     string
	HtmlVersion string
	PageTitle string
	IsTitleFound bool
	H1Count int
	H2Count int
	H3Count int
	H4Count int
	H5Count int
	H6Count int
	Links  chan string
	InternalLinksCount int
	ExternalLinksCount int
	InaccessibleLinksCount int
	PasswordCount int
}

//DTD for diffrent html versions
var HtmlDTD =map[string]string{
	"-//W3C//DTD HTML 4.01//EN"              : "HTML 4.01 Strict",
	"-//W3C//DTD HTML 4.01 Transitional//EN" : "HTML 4.01 Transitional",
	"-//W3C//DTD HTML 4.01 Frameset//EN"     : "HTML 4.01 Frameset",
	"-//W3C//DTD XHTML 1.0 Strict//EN"       : "XHTML 1.0 Strict",
	"-//W3C//DTD XHTML 1.0 Transitional//EN" : "XHTML 1.0 Transitional",
	"-//W3C//DTD XHTML 1.0 Frameset//EN"	 : "XHTML 1.0 Frameset",
	"-//W3C//DTD XHTML 1.1//EN"				 : "XHTML 1.1",
}