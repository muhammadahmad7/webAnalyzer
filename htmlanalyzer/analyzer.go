package htmlanalyzer

import (
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
)


func Analyser(url string)(*AnalysisReport,error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil,err
	}
	ap,err:= analyserHtmlDoc(resp.Body,url)
	if err != nil {
		return nil,err
	}
	return generateAnalysisReport(ap), nil
}

func analyserHtmlDoc(r io.Reader,InputURL string)(*analysisParam,error) {

	node, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	if node==nil||node.FirstChild==nil{
		return nil, errors.New("invalid Html")
	}
	ap:= analysisParam{}
	ap.InputURL=InputURL
	// for rate limiting
	ap.Links=make(chan string, 20 )
	wg:= sync.WaitGroup{}
	wg.Add(1)
	go func() {
		ap.InaccessibleLinksCount=countInassibeleURl(ap.Links)
		wg.Done()
	}()
	traverse(node, &ap)
	wg.Wait()
	close(ap.Links)
	return &ap , nil
}

func generateAnalysisReport(ap *analysisParam)*AnalysisReport{
	ar:=new(AnalysisReport)
	ar.InputURL=ap.InputURL
	if ap.HtmlVersion==""{
		ar.HtmlVersion="HTML Version is not defined it will depend on browser"
	}else {
		ar.HtmlVersion=ap.HtmlVersion
	}
	ar.PageTitle=ap.PageTitle
	ar.H1Count=ap.H1Count
	ar.H2Count=ap.H2Count
	ar.H3Count=ap.H3Count
	ar.H4Count=ap.H4Count
	ar.H5Count=ap.H5Count
	ar.H6Count=ap.H6Count
	if ap.PasswordCount%2!=0{
		ar.IsLogin=true
	}
	ar.InternalLinksCount=ap.InternalLinksCount
	ar.ExternalLinksCount=ap.ExternalLinksCount
	ar.InaccessibleLinksCount=ap.InaccessibleLinksCount
	return ar
}


func traverse(n *html.Node, param *analysisParam) *html.Node {
	if n.Type==html.ElementNode{
		switch n.Data {
		case "h1":
				param.H1Count++
		case "h2":
			param.H2Count++
		case "h3":
			param.H3Count++
		case "h4":
			param.H4Count++
		case "h5":
			param.H5Count++
		case "h6":
			param.H6Count++

		case "input":
			if atr,found:=getNodeAttributeByKey(n.Attr, "type");found{
				if atr.Val=="password"{
					param.PasswordCount++
				}
			}
			// attribute key is href and value is url
		case "a":
			if atr,found:=getNodeAttributeByKey(n.Attr, "href");found{
				checkLinkType(atr.Val,param)
			}

		case "title":{
			if !param.IsTitleFound{
				param.PageTitle,param.IsTitleFound=getHTMLTitle(n)
			}

		}
		}

	}else if n.Type==html.DoctypeNode{
		param.HtmlVersion=getHTMLVersion(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result := traverse(c,param)
		if result != nil {
			return result
		}
	}

	return nil
}

func getHTMLVersion(DTNode *html.Node)string{
	if strings.EqualFold(DTNode.Data,"html") {
		if  DTNode.Attr==nil{
			return "HTML 5"
		}else{
			if attr,found:=getNodeAttributeByKey(DTNode.Attr, "public");found{
				if value,ok:=HtmlDTD[attr.Val];ok{
					return value
				}
			}
		}
	}
	return "unknown HTML version"
}

func getHTMLTitle(TNode *html.Node)(string,bool){
	//if TNode.Parent.Data=="head"&& TNode.Parent.Parent.Data=="html"{
		if TNode.FirstChild!=nil{
			return  TNode.FirstChild.Data ,true
		}
		return "",true
	//}
	//return "",false
}

func getNodeAttributeByKey(artList []html.Attribute, key string)(*html.Attribute,bool){
	for _,each:= range artList{
		if each.Key==key{
			return &each,true
		}
	}
	return nil,false
}

func checkLinkType(Url string,param *analysisParam){
//	if ok,found:=param.Links[Url];!(ok&&found) {
		link, err := url.Parse(Url)
		if err == nil {
			if link.Scheme == "http" || link.Scheme == "https" || link.Scheme == "" {
				input, _ := url.ParseRequestURI(param.InputURL)
				if input.Host == link.Host || link.Host == "" {

					//convert relative path to abs path
					if link.Host==""{
						abUrl, err := input.Parse(Url)
						if err != nil {
							fmt.Print(err)
						}else {
							Url=abUrl.String()
						}
					}
					param.InternalLinksCount++
				}else  {
					param.ExternalLinksCount++
				}
			}else{
				return
			}
		}
		param.Links<-Url
	//}
}