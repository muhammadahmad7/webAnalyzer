package htmlanalyzer

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

func ValidateUrl(Url string)error{
	if _,err := url.ParseRequestURI(Url);err!=nil{
		return errors.New(" invalid URL : "+ Url)
	}
	return nil
}


func countInassibeleURl(chan1 chan string) int{
	count:=0
	var wg sync.WaitGroup

	for url := range chan1{
		wg.Add(1)
		u := url
		go func() {
			if !isURLAccessible(u){
				count++
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return count
}

func isURLAccessible(url string)bool{
	if ValidateUrl(url)==nil{
		if _, err := http.Get(url); err==nil{
			return true
		}else {
			fmt.Print(err)
		}
	}
	return false
}