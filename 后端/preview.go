package main

import "fmt"

var img = [...]string{"webp","bmp","gif","jpg","jpeg","png","svg","tiff","psd","swf"}
var mus = [...]string{"wav","mp3","flac","wma","aiff","ogg","midi","acc","amr","cda"}
var vid = [...]string{"mp4","mpeg","avi","mov","wmv","rmvb","flv","3gp","f4v"}
var text = [...]string{"txt","text","go","c","cpp","java","js","css","html","rs","py","sh","yml","ini","e","hs"}

func getContentType(extName string)string{
	for _,v := range img{
		if v==extName{
			return "image/"+extName
		}
	}
	for _,v  := range vid{
		fmt.Println(v)
		if v==extName{
			return "video/"+extName
		}
	}
	for _,v  := range mus{
		if v==extName{
			return "audio/"+extName
		}
	}
	for _,v  := range text{
		if v==extName{
			return "text/"+extName
		}
	}
	return extName
}