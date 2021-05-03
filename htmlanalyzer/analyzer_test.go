package htmlanalyzer

import (
	"os"
	"testing"
)

func TestAnalyserHtmlDoc(t *testing.T) {
	testPage, _ := os.Open("../index.html")
	defer testPage.Close()
	 ap,err:= analyserHtmlDoc(testPage,"127.0.0.1:8080")
	 if err!=nil{
		 t.Fatalf("error: %v", err)
	 }else if ap.PageTitle!="URL info"{
	 	t.Fatalf("invalid page title expected: URL info got: "+ap.PageTitle )
	 }else if ap.HtmlVersion!="HTML 5"{
		 t.Fatalf("invalid HtmlVersion expected: HTML 5 got: "+ap.HtmlVersion )
	 }else if ap.PasswordCount!=0{
		 t.Fatalf("invalid password count expected: 0 got: %v", ap.PasswordCount )
	 }else if !(ap.ExternalLinksCount == 0 || ap.InternalLinksCount == 0 || ap.H1Count == 0 || ap.H2Count == 0 || ap.H3Count == 0 || ap.H4Count == 0 || ap.H5Count == 0 || ap.H6Count == 0){
		 t.Fatalf("Invalid result" )
	 }
	 t.Logf("passed")
}


func TestValidateUrl(t *testing.T) {

	err:=ValidateUrl("https://github.com/")
	if err!=nil{
		t.Fatalf("error: %v", err)
	}

	err=ValidateUrl("abfbfvf")
	if err==nil{
		t.Fatalf("abfbfvf is inavlid")
	}

	t.Logf("passed")
}