package libraries

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

const (
	STATICURL = "/static/"
)

type DISCUZCODE struct {
	num  int
	html map[int]string
}

var bbcodebufpool = sync.Pool{
	New: func() interface{} {
		return &MsgBuffer{}
	},
}

func Html2bbcode(str string) string {

	code := &DISCUZCODE{num: -1, html: make(map[int]string)}
	m, _ := Preg_match_result(`<div\sclass=[""]?blockcode[""]?>[\s\S]*?<blockquote>([\s\S]+?)<\/blockquote>[\s\S]*?<\/div>`, str, -1)
	for _, v := range m {
		str = strings.Replace(str, v[0], code.codetag(v[1]), -1)
	}

	var postbg = ""
	m, _ = Preg_match_result(`<style[^>]+name="editorpostbg"[^>]*>body{background-image:url\("([^\[\<\r\n"\"\?\(\)]+?)"\)}<\/style>`, str, -1)
	for _, v := range m {
		v[3] = strings.Replace(v[3], "/static/image/postbg/", "", -1)
		str = strings.Replace(str, v[0], "[postbg]"+v[3]+"[/postbg]", -1)
	}

	m, _ = Preg_match_result(`\[postbg\]\s*([^\[\<\r\n"\"\?\(\)]+?)\s*\[\/postbg\]`, str, -1)
	for _, v := range m {
		postbg = v[1]
		str = strings.Replace(str, v[0], "", -1)
	}
	if postbg != "" {
		str = "[postbg]" + postbg + "[/postbg]" + str
	}

	for _, s := range []string{`<style.*?>[\s\S]*?</style>`, `<script.*?>[\s\S]*?</script>`, `<noscript.*?>[\s\S]*?</noscript>`, `<select.*?>[\s\S]*?</select>`, `<object.*?>[\s\S]*?</object>`, `<!--[\s\S]*?-->`, ` on[a-zA-Z]{3,16}\s?=\s?"[\s\S]*?"`, `(\r\n|\n|\r)`, `&((#(32|127|160|173))|shy|nbsp)`, `<br\s+?style=([""]?)clear: both?(\1)[^\>]*>`} {
		str, _ = Preg_replace(s, "", str)
	}

	str, _ = Preg_replace(`([^>=\]""\/]|^)((((https?|ftp):\/\/)|www\.)([\w\-]+\.)*[\w\-\u4e00-\u9fa5]+\.([\.a-zA-Z0-9]+|\u4E2D\u56FD|\u7F51\u7EDC|\u516C\u53F8)((\?|\/|:)+[\w\.\/=\?%\-&~`+"`"+`@":+!]*)+\.(jpg|gif|png|bmp))`, "$1[img]$2[/img]", str)

	str = code.parseurl(str, "bbcode", false)

	str, _ = Preg_replace(`<br[^\>]*>`, "\n", str)

	m, _ = Preg_match_result(`<table[^>]*float:\s*(left|right)[^>]*><tbody><tr><td>\s*([\s\S]+?)\s*<\/td><\/tr></tbody><\/table>`, str, -1)
	for _, v := range m {
		str = strings.Replace(str, v[0], "[float="+v[1]+"]"+v[2]+"[/float]", -1)
	}

	m, _ = Preg_match_result(`<table([^>]*(width|background|background-color|backcolor)[^>]*)>`, str, -1)
	for _, v := range m {
		str = strings.Replace(str, v[0], code.tabletag(v[1]), -1)
	}

	str, _ = Preg_replace(`<table[^>]*>`, "[table]\n", str)
	m, _ = Preg_match_result(`<tr[^>]*(?:background|background-color|backcolor)[:=]\s*(["\"]?)([\(\)\s%,#\w]+)(\1)[^>]*>`, str, -1)
	for _, v := range m {
		str = strings.Replace(str, v[0], "[tr="+v[2]+"]", -1)
	}
	str, _ = Preg_replace(`<tr[^>]*>`, "[tr]", str)
	m, _ = Preg_match_result(`(<t[dh]([^>]*(left|center|right)[^>]*)>)\s*([\s\S]+?)\s*(<\/t[dh]>)`, str, -1)
	for _, v := range m {
		str = strings.Replace(str, v[0], v[1]+"[align="+v[3]+"]"+v[4]+"[/align]"+v[5], -1)
	}
	m, _ = Preg_match_result(`<t[dh]([^>]*(width|colspan|rowspan)[^>]*)>`, str, -1)
	for _, v := range m {
		str = strings.Replace(str, v[0], code.tdtag(v[1]), -1)
	}

	str, _ = Preg_replace(`<t[dh][^>]*>`, "[td]", str)
	str, _ = Preg_replace(`<\/t[dh]>`, "[/td]", str)
	str, _ = Preg_replace(`<\/tr>`, "[/tr]\n", str)
	str, _ = Preg_replace(`<\/table>`, "[/table]", str)
	str, _ = Preg_replace(`<\/table>`, "[/table]", str)
	str, _ = Preg_replace(`<h\d[^>]*>`, "[b]", str)
	str, _ = Preg_replace(`<\/h\d>`, "[/b]", str)
	m, _ = Preg_match_result(`<h([0-9]+)[^>]*>([\s\S]*?)<\/h\1>`, str, -1)
	for _, v := range m {
		s, _ := strconv.Atoi(v[1])
		size := 7 - s
		str = strings.Replace(str, v[0], "[size="+strconv.Itoa(size)+"]"+v[2]+"[/size]\n\n", -1)
	}
	str, _ = Preg_replace(`<hr[^>]*>`, "[hr]", str)
	m, _ = Preg_match_result(`<img[^>]+smilieid=([""]?)(\d+)(\1)[^>]*>`, str, -1)
	for _, v := range m {
		str = strings.Replace(str, v[0], code.smileycode(v[2]), -1)
	}
	m, _ = Preg_match_result(`<img([^>]*src[^>]*)>`, str, -1)
	for _, v := range m {
		str = strings.Replace(str, v[0], code.imgtag(v[1]), -1)
	}

	str, _ = Preg_replace(`<a\s+?name=([""]?)(.+?)(\1)[\s\S]*?>([\s\S]*?)<\/a>`, "$4", str)
	str, _ = Preg_replace(`<div[^>]*quote[^>]*><blockquote>([\s\S]*?)<\/blockquote><\/div>([\s\S]*?)(<br[^>]*>)?`, "[quote]$1[/quote]", str)
	str, _ = Preg_replace(`<div[^>]*blockcode[^>]*><blockquote>([\s\S]*?)<\/blockquote><\/div>([\s\S]*?)(<br[^>]*>)?`, "[code]$1[/code]", str)

	str = code.recursion("b", str, code.simpletag, "b")
	str = code.recursion("strong", str, code.simpletag, "b")

	str = code.recursion("i", str, code.simpletag, "i")

	str = code.recursion("em", str, code.simpletag, "i")
	str = code.recursion("u", str, code.simpletag, "u")
	str = code.recursion("strike", str, code.simpletag, "s")
	str = code.recursion("a", str, code.atag, "")
	str = code.recursion("font", str, code.fonttag, "")
	str = code.recursion("blockquote", str, code.simpletag, "indent")
	str = code.recursion("ol", str, code.listtag, "")
	str = code.recursion("ul", str, code.listtag, "")
	str = code.recursion("div", str, code.dstag, "")
	str = code.recursion("p", str, code.ptag, "")
	str = code.recursion("span", str, code.fonttag, "")

	str, _ = Preg_replace(`<[\/\!]*?[^<>]*?>`, "", str)

	str = code.clearcode(str)
	str = strings.Replace(str, "&nbsp;", " ", -1)
	str = strings.Replace(str, "&lt;", "<", -1)
	str = strings.Replace(str, "&gt;", ">", -1)
	str = strings.Replace(str, "&amp;", "&", -1)

	if code.num >= 0 {
		for i := 0; i <= code.num; i++ {
			str = strings.Replace(str, "[\tDISCUZ_CODE_"+strconv.Itoa(i)+"\t]", code.html[i], -1)
		}
	}

	return str
}

func (code *DISCUZCODE) codetag(text string) string {

	code.num++
	text, _ = Preg_replace(`\$`, "&&", text)
	code.html[code.num] = "[code]" + text + "[/code]"
	return "[\tDISCUZ_CODE_" + strconv.Itoa(code.num) + "\t]"
}
func (code *DISCUZCODE) parseurl(str string, mode string, parsecode bool) string {

	if parsecode {
		m, _ := Preg_match_result(`\[code\]([\s\S]+?)\[\/code\]`, str, -1)
		for _, v := range m {
			str = strings.Replace(str, v[0], code.codetag(v[1]), -1)
		}

	}
	str, _ = Preg_replace(`([^>=\]"'\/]|^)((((https?|ftp):\/\/)|www\.)([\w\-]+\.)*[\w\-\u4e00-\u9fa5]+\.([\.a-zA-Z0-9]+|\u4E2D\u56FD|\u7F51\u7EDC|\u516C\u53F8)((\?|\/|:)+[\w\.\/=\?%\-&~`+"`"+`@':+!]*)+\.(swf|flv))`, "$1[flash]$2[/flash]", str)
	str, _ = Preg_replace(`([^>=\]"'\/]|^)((((https?|ftp):\/\/)|www\.)([\w\-]+\.)*[\w\-\u4e00-\u9fa5]+\.([\.a-zA-Z0-9]+|\u4E2D\u56FD|\u7F51\u7EDC|\u516C\u53F8)((\?|\/|:)+[\w\.\/=\?%\-&~`+"`"+`@':+!]*)+\.(mp3|wma))`, "$1[audio]$2[/audio]", str)
	if mode == "html" {
		str, _ = Preg_replace(`([^>=\]"'\/@]|^)((((https?|ftp|gopher|news|telnet|rtsp|mms|callto|bctp|ed2k|thunder|qqdl|synacast):\/\/))([\w\-]+\.)*[:\.@\-\w\u4e00-\u9fa5]+\.([\.a-zA-Z0-9]+|\u4E2D\u56FD|\u7F51\u7EDC|\u516C\u53F8)((\?|\/|:)+[\w\.\/=\?%\-&;~`+"`"+`@':+!#]*)*)`, `$1<a href="$2" target="_blank">$2</a>`, str)
	} else {
		str, _ = Preg_replace(`([^>=\]"'\/@]|^)((((https?|ftp|gopher|news|telnet|rtsp|mms|callto|bctp|ed2k|thunder|qqdl|synacast):\/\/))([\w\-]+\.)*[:\.@\-\w\u4e00-\u9fa5]+\.([\.a-zA-Z0-9]+|\u4E2D\u56FD|\u7F51\u7EDC|\u516C\u53F8)((\?|\/|:)+[\w\.\/=\?%\-&;~`+"`"+`@':+!#]*)*)`, "$1[url]$2[/url]", str)
	}
	if mode == "html" {
		str, _ = Preg_replace(`([^\w>=\]"'\/@]|^)((www\.)([\w\-]+\.)*[:\.@\-\w\u4e00-\u9fa5]+\.([\.a-zA-Z0-9]+|\u4E2D\u56FD|\u7F51\u7EDC|\u516C\u53F8)((\?|\/|:)+[\w\.\/=\?%\-&;~`+"`"+`@':+!#]*)*)`, `$1<a href="$2" target="_blank">$2</a>`, str)
	} else {
		str, _ = Preg_replace(`([^\w>=\]"'\/@]|^)((www\.)([\w\-]+\.)*[:\.@\-\w\u4e00-\u9fa5]+\.([\.a-zA-Z0-9]+|\u4E2D\u56FD|\u7F51\u7EDC|\u516C\u53F8)((\?|\/|:)+[\w\.\/=\?%\-&;~`+"`"+`@':+!#]*)*)`, "$1[url]$2[/url]", str)
	}
	if mode == "html" {
		str, _ = Preg_replace(`([^\w->=\]:"'\.\/]|^)(([\-\.\w]+@[\.\-\w]+(\.\w+)+))`, `$1<a href="mailto:$2">$2</a>`, str)
	} else {
		str, _ = Preg_replace(`([^\w->=\]:"'\.\/]|^)(([\-\.\w]+@[\.\-\w]+(\.\w+)+))`, "$1[email]$2[/email]", str)
	}

	if parsecode && code.num >= 0 {
		for i := 0; i <= code.num; i++ {
			str = strings.Replace(str, "[\tDISCUZ_CODE_"+strconv.Itoa(i)+"\t]", code.html[i], -1)
		}
	}
	return str
}
func (code *DISCUZCODE) tabletag(attributes string) string {
	var width = ""
	matches, _ := Preg_match_result(`width=(["']?)(\d{1,4}%?)(\1)`, attributes, 1)

	if len(matches) == 1 {
		width = matches[0][2][len(matches[0][2])-1 : len(matches[0][2])]
		if width == `%` {
			width = matches[0][2][:len(matches[0][2])-1]
			if width < "98" {
				width = matches[0][2]
			} else {
				width = `98%`
			}
		} else {
			if matches[0][2] <= "560" {
				width = matches[0][2]
			} else {
				width = `98%`
			}
		}

	} else {
		matches, _ := Preg_match_result(`width\s?:\s?(\d{1,4})([px|%])`, attributes, 1)
		if len(matches) == 1 {
			if matches[0][2] == `%` {
				if matches[0][1] <= "98" {
					width = matches[0][1] + `%`
				} else {
					width = `98%`
				}
			} else {
				if matches[0][1] <= "560" {
					width = matches[0][1]
				} else {
					width = `98%`
				}
			}
		}
	}
	var bgcolor = ""
	matches, _ = Preg_match_result(`(?:background|background-color|bgcolor)[:=]\s*(["']?)((rgb\(\d{1,3}%?,\s*\d{1,3}%?,\s*\d{1,3}%?\))|(#[0-9a-fA-F]{3,6})|([a-zA-Z]{1,20}))(\1)`, attributes, 1)

	if len(matches) == 1 {
		bgcolor = matches[0][2]
		if width == "" {
			width = `98%`
		}

	}
	if bgcolor != "" {
		return "[table=" + width + "," + bgcolor + "]\n"
	} else {
		if width != "" {
			return "[table=" + width + "]\n"
		}
	}
	return "[table]\n"
}
func (code *DISCUZCODE) tdtag(attributes string) string {

	var colspan = "1"
	var rowspan = "1"
	var width = ""
	matches, _ := Preg_match_result(`colspan=(["']?)(\d{1,2})(\1)`, attributes, 1)

	if len(matches) == 1 {
		colspan = matches[0][2]
	}
	matches, _ = Preg_match_result(`rowspan=(["']?)(\d{1,2})(\1)`, attributes, 1)

	if len(matches) == 1 {
		rowspan = matches[0][2]
	}
	matches, _ = Preg_match_result(`width=(["']?)(\d{1,4}%?)(\1)`, attributes, 1)

	if len(matches) == 1 {
		width = matches[0][2]
	}
	if In_slice(width, []string{"", "0", `100%`}) {
		if colspan == "1" && rowspan == "1" {
			return "[td]"
		} else {
			return "[td=" + colspan + "," + rowspan + "]"
		}
	}
	if colspan == "1" && rowspan == "1" {
		return "[td=" + width + "]"
	}
	return "[td=" + colspan + "," + rowspan + "," + width + "]"

}
func (code *DISCUZCODE) smileycode(smileyid string) string {

	for _, v := range smilies_array {
		for _, vv := range v {
			for _, v3 := range vv {
				if v3[0] == smileyid {

					return v3[1]

				}
			}
		}
	}
	return ""
}
func (code *DISCUZCODE) imgtag(attributes string) string {
	var width, height, src string

	matches, _ := Preg_match_result(`src=(["']?)([\s\S]*?)(\1)`, attributes, 1)

	if len(matches) == 1 {
		src = matches[0][2]
	} else {
		return ""
	}
	matches, _ = Preg_match_result(`(max-)?width\s?:\s?(\d{1,4})(px)?`, attributes, 1)

	if len(matches) == 1 && matches[0][1] == "" {
		width = matches[0][2]
	}
	matches, _ = Preg_match_result(`height\s?:\s?(\d{1,4})(px)?`, attributes, 1)
	if len(matches) == 1 {
		height = matches[0][1]
	}

	if width == "" {
		matches, _ = Preg_match_result(`width=(["']?)(\d+)(\1)`, attributes, 1)

		if len(matches) == 1 {
			width = matches[0][2]
		}
	}

	if height == "" {
		matches, _ = Preg_match_result(`height=(["']?)(\d+)(\1)`, attributes, 1)
		if len(matches) == 1 {
			height = matches[0][2]
		}
	}
	aid := "0"
	matches, _ = Preg_match_result(`aid=(["']?)attachimg_(\d+)(\1)`, attributes, 1)

	if len(matches) == 1 {
		aid = matches[0][2]
	}
	if width < "0" {
		width = "0"
	}
	if height < "0" {
		height = "0"
	}
	if width > "0" || height > "0" {
		return "[img_" + aid + "=" + width + "," + height + "]" + src + "[/img]"
	}
	return "[img_" + aid + "]" + src + "[/img]"
}
func (code *DISCUZCODE) recursion(tagname, text string, dofunction func(string, string, string, string) string, extraargs string) string {

	tagname = strings.ToLower(tagname)

	var open_tag = "<" + tagname
	var open_tag_len = len(open_tag)
	var close_tag = "</" + tagname + ">"
	var close_tag_len = len(close_tag)
	var beginsearchpos = 0

	for {

		var textlower = strings.ToLower(text)
		var tagbegin = code.index(textlower, open_tag, beginsearchpos)
		var strlen = len(text)
		if tagbegin == -1 {
			break
		}

		var inquote = ""
		var found = false
		var tagnameend = 0
		var optionend = 0
		var t_char = ""

		for optionend = tagbegin; optionend < strlen; optionend++ {
			t_char = textlower[optionend : optionend+1]
			if (t_char == `"` || t_char == "'") && inquote == "" {
				inquote = t_char
			} else if (t_char == `"` || t_char == "'") && inquote == t_char {
				inquote = ``
			} else if t_char == `>` && inquote == "" {
				found = true
				break
			} else if (t_char == `=` || t_char == " ") && tagnameend == 0 {
				tagnameend = optionend
			}
		}

		if !found {
			break
		}
		if tagnameend == 0 {
			tagnameend = optionend
		}

		var tagoptions = textlower[tagbegin+open_tag_len : optionend]
		var acttagname = textlower[tagbegin*1+1 : tagnameend]
		if acttagname != tagname {
			beginsearchpos = optionend
			continue
		}

		var tagend = code.index(textlower, close_tag, optionend)
		if tagend == -1 {
			break
		}
		var nestedopenpos = code.index(textlower, open_tag, optionend)
		for nestedopenpos != -1 && tagend != -1 {
			if nestedopenpos > tagend {
				break
			}
			tagend = code.index(textlower, close_tag, tagend+close_tag_len)
			nestedopenpos = code.index(textlower, open_tag, nestedopenpos+open_tag_len)
		}

		if tagend == -1 {
			beginsearchpos = optionend

			continue
		}

		var localbegin = optionend + 1
		var localtext = dofunction(tagoptions, text[localbegin:tagend], tagname, extraargs)

		text = text[0:tagbegin] + localtext + text[tagend+close_tag_len:len(text)]
		beginsearchpos = tagbegin + len(localtext)

		if tagbegin == -1 {
			break
		}

	}

	return text
}
func (code *DISCUZCODE) simpletag(options, text, tagname, parseto string) string {
	if strings.Trim(text, " ") == "" {
		return ""
	}
	text = code.recursion(tagname, text, code.simpletag, parseto)
	return "[" + parseto + "]" + text + "[/" + parseto + "]"
}
func (code *DISCUZCODE) atag(aoptions, text, a, b string) string {
	if strings.Trim(text, " ") == "" {
		return ""
	}
	var pend = code.parsestyle(aoptions, "", "")
	href := code.getoptionvalue("href", aoptions)

	if len(href) >= 11 && href[0:11] == "javascript:" {
		return strings.Trim(code.recursion("a", text, code.atag, ""), " ")
	}

	return pend["prepend"] + "[url=" + href + "]" + strings.Trim(code.recursion("a", text, code.atag, ""), " ") + "[/url]" + pend["append"]
}
func (code *DISCUZCODE) fonttag(fontoptions, text, a, b string) string {
	var prepend = ""
	var ap = ""
	var tags = map[string]string{
		"font": "face=", "size": "size=", "color": "color=",
	}

	for bbcode, v := range tags {
		optionvalue := code.fetchoptionvalue(v, fontoptions)
		if optionvalue != "" {
			prepend += "[" + bbcode + "=" + optionvalue + "]"
			ap = "[/" + bbcode + "]" + ap
		}
	}

	var pend = code.parsestyle(fontoptions, prepend, ap)
	return pend["prepend"] + code.recursion("font", text, code.fonttag, "") + pend["append"]
}
func (code *DISCUZCODE) listtag(listoptions, text, tagname, b string) string {
	text, _ = Preg_replace(`<li>(([\s\S](?!<\/li))*?)(?=<\/?ol|<\/?ul|<li|\[list|\[\/list)`, "<li>$1</li>", text)
	text += "</li>"
	text = code.recursion("li", text, code.litag, "")
	var opentag = "[list]"
	var listtype = code.fetchoptionvalue("type=", listoptions)
	if listtype == "" && tagname == "ol" {
		listtype = "1"
	}

	if In_slice(listtype, []string{"1", "a", "A"}) {
		opentag = "[list=" + listtype + "]"
	}
	if text == "" {
		return ""
	}
	return opentag + "\n" + code.recursion(tagname, text, code.listtag, "") + "[/list]"
}

func (code *DISCUZCODE) dstag(options, text, tagname, b string) string {
	if strings.Trim(text, " ") == "" {
		return "\n"
	}
	var pend = code.parsestyle(options, "", "")
	var prepend = pend["prepend"]
	var ap = pend["append"]
	if In_slice(tagname, []string{"div", "p"}) {
		align := code.getoptionvalue("align", options)
		if In_slice(align, []string{"left", "center", "right"}) {
			prepend = "[align=" + align + "]" + prepend
			ap += "[/align]"
		} else {
			ap += "\n"
		}
	}
	return prepend + code.recursion(tagname, text, code.dstag, "") + ap
}
func (code *DISCUZCODE) ptag(options, text, tagname, b string) string {
	if strings.Trim(text, " ") == "" {
		return "\n"
	}
	if strings.Trim(options, " ") == "" {
		return text + "\n"
	}

	var lineHeight = ""
	var textIndent = ""
	var align = ""
	matches, _ := Preg_match_result(`line-height\s?:\s?(\d{1,3})px`, options, 1)
	if len(matches) == 1 {
		lineHeight = matches[0][1]
	}
	matches, _ = Preg_match_result(`text-indent\s?:\s?(\d{1,3})em`, options, 1)
	if len(matches) == 1 {
		textIndent = matches[0][1]
	}
	matches, _ = Preg_match_result(`text-align\s?:\s?(left|center|right)`, options, 1)
	if len(matches) == 1 {
		align = matches[0][1]
	} else {
		align = code.getoptionvalue("align", options)
	}
	if !In_slice(align, []string{"left", "center", "right"}) {
		align = "left"
	}

	style := code.getoptionvalue("style", options)
	for _, re := range []string{`line-height\s?:\s?(\d{1,3})px`, `text-indent\s?:\s?(\d{1,3})em`, `text-align\s?:\s?(left|center|right)`} {
		style, _ = Preg_replace(re, "", style)
	}
	if style != "" {
		text = `<span style="` + style + `">` + text + `</span>`
	}
	if lineHeight == "" && textIndent == "" {

		return "[align=" + align + "]" + text + `[/align]`
	} else {
		return `[p=` + lineHeight + `, ` + textIndent + `, ` + align + `]` + text + `[/p]`
	}
}

func (code *DISCUZCODE) clearcode(str string) string {
	for _, v := range []string{`\[url\]\[\/url\]`, `\[url=((https?|ftp|gopher|news|telnet|rtsp|mms|callto|bctp|thunder|qqdl|synacast){1}:\/\/|www\.|mailto:)?([^\s\[\"']+?)\]\[\/url\]`, `\[email\]\[\/email\]`, `\[email=(.[^\[]*)\]\[\/email\]`, `\[color=([^\[\<]+?)\]\[\/color\]`, `\[size=(\d+?)\]\[\/size\]`, `\[size=(\d+(\.\d+)?(px|pt)+?)\]\[\/size\]`, `\[font=([^\[\<]+?)\]\[\/font\]`, `\[align=([^\[\<]+?)\]\[\/align\]`, `\[p=(\d{1,2}), (\d{1,2}), (left|center|right)\]\[\/p\]`, `\[float=([^\[\<]+?)\]\[\/float\]`, `\[quote\]\[\/quote\]`, `\[code\]\[\/code\]`, `\[table\]\[\/table\]`, `\[free\]\[\/free\]`, `\[b\]\[\/b]`, `\[u\]\[\/u]`, `\[i\]\[\/i]`, `\[s\]\[\/s]`} {
		str, _ = Preg_replace(v, "", str)
	}

	return str
}
func (code *DISCUZCODE) parsestyle(tagoptions, prepend, ap string) map[string]string {
	var searchlist = [][]string{
		[]string{"align", `true`, `text-align:\s*(left|center|right);?`, "1", ""},
		[]string{"float", `true`, `float:\s*(left|right);?`, "1", ""},
		[]string{"color", `true`, `(^|[;\s])color:\s*([^;]+);?`, "2", ""},
		[]string{`backcolor`, `true`, `(^|[;\s])background-color:\s*([^;]+);?`, "2", ""},
		[]string{`font`, `true`, `font-family:\s*([^;]+);?`, "1", ""},
		[]string{`size`, `true`, `font-size:\s*(\d+(\.\d+)?(px|pt|in|cm|mm|pc|em|ex|%|));?`, "1", ""},
		[]string{`size`, `true`, `font-size:\s*(x\-small|small|medium|large|x\-large|xx\-large|\-webkit\-xxx\-large);?`, "1", `size`},
		[]string{`b`, `false`, `font-weight:\s*(bold);?`, "", ""},
		[]string{`i`, `false`, `font-style:\s*(italic);?`, "", ""},
		[]string{`u`, `false`, `text-decoration:\s*(underline);?`, "", ""},
		[]string{`s`, `false`, `text-decoration:\s*(line-through);?`, "", ""},
	}
	var sizealias = map[string]string{"x-small": `1`, "small": `2`, "medium": `3`, "large": `4`, "x-large": `5`, "xx-large": `6`, "-webkit-xxx-large": `7`}
	var style = code.getoptionvalue("style", tagoptions)
	m, _ := Preg_match_result(`^(?:\s|)color:\s*rgb\((\d+),\s*(\d+),\s*(\d+)\)(;?)`, style, -1)
	for _, v := range m {
		style = strings.Replace(style, v[0], "color:#"+code.tohex(v[1])+code.tohex(v[2])+code.tohex(v[3])+v[4], -1)
	}

	for _, v := range searchlist {
		matches, _ := Preg_match_result(v[2], style, 1)

		if len(matches) == 1 {

			var opnvalue string
			index, err := strconv.Atoi(v[3])
			if err == nil {
				opnvalue = matches[0][index]

			}
			if v[4] == `size` {
				opnvalue = sizealias[opnvalue]
			}
			prepend += `[` + v[0]
			if v[1] == "true" {
				prepend += "=" + opnvalue
			}
			prepend += `]`

			ap = `[/` + v[0] + `]` + ap
		}
	}
	return map[string]string{`prepend`: prepend, `append`: ap}
}
func (code *DISCUZCODE) tohex(ten string) string {
	i, _ := strconv.Atoi(ten)
	return fmt.Sprintf("%X", i)
}
func (code *DISCUZCODE) getoptionvalue(option, text string) string {
	matches, _ := Preg_match_result(option+`(\s+?)?\=(\s+?)?[\"']?(.+?)([\"']|$|>)`, text, 1)
	if len(matches) == 1 {
		return strings.Trim(matches[0][3], " ")
	}
	return ``
}
func (code *DISCUZCODE) fetchoptionvalue(option, text string) string {
	position := strings.Index(strings.ToLower(text), strings.ToLower(option))
	if position > -1 {
		delimiter := position + len(option)
		var delimchar = " "
		if text[delimiter:delimiter+1] == `"` {
			delimchar = `"`
		} else if text[delimiter:delimiter+1] == `'` {
			delimchar = `'`
		}
		delimloc := code.index(strings.ToLower(text), strings.ToLower(delimchar), delimiter+1)
		if delimloc == -1 {
			delimloc = len(text)
		} else if delimchar == `"` || delimchar == `'` {
			delimiter++
		}
		return strings.Trim(text[delimiter:delimloc], " ")
	} else {
		return ""
	}
}
func (code *DISCUZCODE) litag(listoptions, text, a, b string) string {
	text, _ = Preg_replace(`(\s+)$`, ``, text)
	return "[*]" + text + "\n"
}
func (code *DISCUZCODE) index(str string, find string, offset int) int {
	if offset > len(str) {
		return -1
	}
	var found = strings.Index(str[offset:], find)
	if found > -1 {
		found += offset
	}
	return found
}

var smilies_array [][][][]string
var Smiliesarray []string
var smilies_type = map[string][]string{
	"_1": []string{"默认", "default"},
	"_2": []string{"酷猴", "coolmonkey"},
	"_3": []string{"呆呆男", "grapeman"},
}
var postimg_type = struct {
	hrline []string
	postbg []string
}{
	hrline: []string{"line1.png", "line6.png", "2.gif", "3.gif", "line8.png", "line5.png", "line7.png", "line9.png", "0.gif", "line2.png", "1.gif", "4.gif", "5.gif", "line4.png", "line3.png"},
	postbg: []string{"bg1.png", "bg9.png", "1.jpg", "bg5.png", "3.jpg", "bg2.png", "0.gif", "bg3.png", "bg4.png", "bg8.png", "bg7.png", "2.jpg", "bg10.png", "bg6.png"},
}

func Bbcode2html(str string, allowbbcode, allowimgurl, allowhtml, allowsmilies, allowimgcode, parseurloff bool) string {

	if str == "" {
		return ""
	}
	code := &DISCUZCODE{num: -1, html: make(map[int]string)}

	parsetype := 0

	if allowbbcode && parsetype != 1 {
		m, _ := Preg_match_result(`\[code\]([\s\S]+?)\[\/code\]`, str, -1)
		for _, v := range m {

			str = strings.Replace(str, v[0], code.parsecode(v[1]), -1)
		}
	}

	if allowimgurl {
		str, _ = Preg_replace(`([^>=\]"'\/]|^)((((https?|ftp):\/\/)|www\.)([\w\-]+\.)*[\w\-\u4e00-\u9fa5]+\.([\.a-zA-Z0-9]+|\u4E2D\u56FD|\u7F51\u7EDC|\u516C\u53F8)((\?|\/|:)+[\w\.\/=\?%\-&~`+"`"+`@':+!]*)+\.(jpg|gif|png|bmp))`, "$1[img]$2[/img]", str)
	}

	if !allowhtml {
		str = strings.Replace(str, "<", "&lt;", -1)
		str = strings.Replace(str, ">", "&gt;", -1)

		if !parseurloff {
			str = code.parseurl(str, "html", false)
		}
	}

	if allowsmilies {
		for typeid, v := range smilies_array {
			for _, vv := range v {
				for _, v3 := range vv {
					str = strings.Replace(str, v3[1], `<img src="`+STATICURL+`image/smiley/`+smilies_type["_"+strconv.Itoa(typeid)][1]+`/`+v3[2]+`" border="0" smilieid="`+v3[0]+`" alt="`+v3[1]+`" />`, -1)
				}
			}
		}

	}

	if allowbbcode {
		str = code.clearcode(str)
		m, _ := Preg_match_result(`\[url\]\s*((https?|ftp|gopher|news|telnet|rtsp|mms|callto|bctp|thunder|qqdl|synacast){1}:\/\/|www\.)([^\[\"']+?)\s*\[\/url\]`, str, -1)
		for _, v := range m {
			str = strings.Replace(str, v[0], code.cuturl(v[1]+v[3]), -1)
		}

		str, _ = Preg_replace(`\[url=((https?|ftp|gopher|news|telnet|rtsp|mms|callto|bctp|thunder|qqdl|synacast){1}:\/\/|www\.|mailto:)?([^\r\n\[\"']+?)\]([\s\S]+?)\[\/url\]`, `<a href="$1$3" target="_blank">$4</a>`, str)
		str, _ = Preg_replace(`\[email\](.[^\\=[]*)\[\/email\]`, `<a href="mailto:$1">$1</a>`, str)
		str, _ = Preg_replace(`\[email=(.[^\\=[]*)\](.*?)\[\/email\]`, `<a href="mailto:$1" target="_blank">$2</a>`, str)
		m, _ = Preg_match_result(`\[postbg\]\s*([^\[\<\r\n;'\"\?\(\)]+?)\s*\[\/postbg\]`, str, -1)
		for _, v := range m {
			addCSS := ""
			if In_slice(v[1], postimg_type.postbg) {
				addCSS = `<style type="text/css" name="editorpostbg">body{background-image:url("` + STATICURL + `image/postbg/` + v[1] + `");}</style>`
			}
			str = strings.Replace(str, v[0], addCSS, -1)
		}
		str, _ = Preg_replace(`\[color=([\w#\(\),\s]+?)\]`, `<font color="$1">`, str)
		str, _ = Preg_replace(`\[backcolor=([\w#\(\),\s]+?)\]`, `<font style="background-color:$1">`, str)

		str, _ = Preg_replace(`\[size=(\d+?)\]`, `<font size="$1">`, str)
		str, _ = Preg_replace(`\[size=(\d+(\.\d+)?(px|pt)+?)\]`, `<font style="font-size: $1">`, str)
		str, _ = Preg_replace(`\[font=([^\[\<\=]+?)\]`, `<font face="$1">`, str)
		str, _ = Preg_replace(`\[align=([^\[\<\=]+?)\]`, `<div align="$1">`, str)
		str, _ = Preg_replace(`\[p=(\d{1,2}|null|), (\d{1,2}|null|), (left|center|right)\]`, `<p style="line-height: $1px; text-indent: $2em; text-align: $3;">`, str)
		str, _ = Preg_replace(`\[float=left\]`, `<br style="clear: both"><span style="float: left; margin-right: 5px;">`, str)
		str, _ = Preg_replace(`\[float=right\]`, `<br style="clear: both"><span style="float: right; margin-left: 5px;">`, str)

		if parsetype != 1 {
			str, _ = Preg_replace(`\[quote]([\s\S]*?)\[\/quote\]\s?\s?`, `<div class="quote"><blockquote>$1</blockquote></div>`, str)

		}
		for _, v := range [][]string{[]string{`\[\/color\]`, `</font>`},
			[]string{`\[\/backcolor\]`, `</font>`},
			[]string{`\[\/size\]`, `</font>`},
			[]string{`\[\/font\]`, `</font>`},
			[]string{`\[\/align\]`, `</div>`},
			[]string{`\[\/p\]`, `</p>`},
			[]string{`\[b\]`, `<b>`},
			[]string{`\[\/b\]`, `</b>`},
			[]string{`\[i\]`, `<i>`},
			[]string{`\[\/i\]`, `</i>`},
			[]string{`\[u\]`, `<u>`},
			[]string{`\[\/u\]`, `</u>`},
			[]string{`\[s\]`, `<strike>`},
			[]string{`\[\/s\]`, `</strike>`},
			[]string{`\[hr\]`, `<hr class="l" />`},
			[]string{`\[list\]`, `<ul>`},
			[]string{`\[list=1\]`, `<ul type=1 class="litype_1">`},
			[]string{`\[list=a\]`, `<ul type=a class="litype_2">`},
			[]string{`\[list=A\]`, `<ul type=A class="litype_3">`},
			[]string{`\\s?\[\*\]`, `<li>`},
			[]string{`\[\/list\]`, `</ul>`},
			[]string{`\[indent\]`, `<blockquote>`},
			[]string{`\[\/indent\]`, `</blockquote>`},
			[]string{`\[\/float\]`, `</span>`},
		} {
			str, _ = Preg_replace(v[0], v[1], str)
		}

		for m, _ = Preg_match_result(`\[table(?:=(\d{1,4}%?)(?:,([\(\)%,#\w ]+))?)?\]\s*([\s\S]+?)\s*\[\/table\]`, str, -1); len(m) > 0; m, _ = Preg_match_result(`\[table(?:=(\d{1,4}%?)(?:,([\(\)%,#\w ]+))?)?\]\s*([\s\S]+?)\s*\[\/table\]`, str, -1) {
			for k, v := range m {
				if k >= 4 {
					break
				}
				str = strings.Replace(str, v[0], code.parsetable(v[1], v[2], v[3]), -1)
			}
		}

		if allowimgcode {
			str, _ = Preg_replace(`\[img_(\d+)\]\s*([^\[\"\<\r\n]+?)\s*\[\/img\]`, `<img aid="attachimg_$1" src="$2" border="0" alt=""  />`, str)
			m, _ := Preg_match_result(`\[img_(\d+)=(\d{1,4})[x|\,](\d{1,4})\]\s*([^\[\"\<\r\n]+?)\s*\[\/img\]`, str, -1)
			for _, v := range m {
				tmp := bbcodebufpool.Get().(*MsgBuffer)
				tmp.Reset()
				tmp.WriteString(`<img aid="attachimg_`)
				tmp.WriteString(v[1])
				if v[2] > "0" {
					tmp.WriteString(` width="`)
					tmp.WriteString(v[2])
					tmp.WriteString(`"`)
				}
				if v[3] > "0" {
					tmp.WriteString(` _height="`)
					tmp.WriteString(v[3])
					tmp.WriteString(`"`)
				}
				tmp.WriteString(` src=`)
				tmp.WriteString(v[4])
				tmp.WriteString(`" border="0" alt="" />`)
				str = strings.Replace(str, v[0], tmp.String(), -1)
				bbcodebufpool.Put(tmp)
			}
			str, _ = Preg_replace(`\[img\]\s*([^\[\"\<\r\n]+?)\s*\[\/img\]`, `<img src="$1" border="0" alt=""  />`, str)

		} else {
			str, _ = Preg_replace(`\[img\]\s*([^\[\"\<\r\n]+?)\s*\[\/img\]`, `<a href="$1" target="_blank">$1</a>`, str)
			str, _ = Preg_replace(`\[img=(\d{1,4})[x|\,](\d{1,4})\]\s*([^\[\"\<\r\n]+?)\s*\[\/img\]`, `<a href="$3" target="_blank">$3</a>`, str)

		}

	}
	for i := 0; i <= code.num; i++ {
		str = strings.Replace(str, "[\tDISCUZ_CODE_"+strconv.Itoa(i)+"\t]", code.html[i], -1)
	}

	m, _ := Preg_match_result(`(^|>)([^<]+)(?=<|$)`, str, -1)
	for _, v := range m {
		for _, vv := range [][]string{
			[]string{`\t`, `&nbsp; &nbsp; &nbsp; &nbsp; `},
			[]string{`   `, `&nbsp; &nbsp;`},
			[]string{`  `, `&nbsp;&nbsp;`},
			[]string{`(\r\n|\n|\r)`, `<br />`},
		} {
			v[2], _ = Preg_replace(vv[0], vv[1], v[2])
		}
		str = strings.Replace(str, v[0], v[1]+v[2], -1)
	}

	return str
}
func (code *DISCUZCODE) parsecode(text string) string {
	code.num++
	code.html[code.num] = `<div class="blockcode"><blockquote>` + code.htmlspecialchars(text) + `</blockquote></div>`
	return "[\tDISCUZ_CODE_" + strconv.Itoa(code.num) + "\t]"
}
func (code *DISCUZCODE) htmlspecialchars(str string) string {
	str = strings.Replace(str, "&", "&amp;", -1)
	str = strings.Replace(str, "<", "&lt;", -1)
	str = strings.Replace(str, ">", "&gt;", -1)
	str = strings.Replace(str, `"`, "&quot;", -1)
	return str
}

func (code *DISCUZCODE) cuturl(url string) string {
	var length = 65
	var urllink = `<a href="`
	if strings.ToLower(url)[0:4] == "www." {
		urllink += `'http://`
	}
	urllink += url + `" target="_blank">`

	if len(url) > length {
		url = url[0:length/2] + ` ... ` + url[len(url)-length/3:]
	}
	urllink += url + `</a>`
	return urllink
}
func (code *DISCUZCODE) parsetable(width, bgcolor, str string) string {
	if str == "" {
		return ""
	}
	if width != "" {
		if width[len(width)-1:] == `%` {
			if width[:len(width)-1] > "98" {
				width = "98%"
			}
		} else if width > "560" {
			width = "98%"
		}
	}

	if !strings.Contains(str, `[/tr]`) && strings.Contains(str, `[/td]`) {
		str = `[tr]` + str + `[/tr]`
	}
	var simple string
	tmp := bbcodebufpool.Get().(*MsgBuffer)
	if !strings.Contains(str, `[/tr]`) && !strings.Contains(str, `[/td]`) {

		var rows = strings.Split(str, "\n")
		tmp.Reset()
		for _, r := range rows {
			for _, v := range [][]string{
				[]string{`\r`, ``},
				[]string{`\\\|`, `&#124;`},
				[]string{`\|`, `</td><td>`},
				[]string{`\\n`, "\n"},
			} {
				r, _ = Preg_replace(v[0], v[1], r)
			}
			tmp.WriteString(`<tr><td>`)
			tmp.WriteString(r)
			tmp.WriteString(`</td></tr>`)
		}
		str = tmp.String()
		simple = ` simpletable`
	} else {
		var m [][]string
		for m, _ = Preg_match_result(`\[tr(?:=([\(\)\s%,#\w]+))?\]\s*\[td(?:=(\d{1,4}%?))?\]`, str, -1); len(m) > 0; m, _ = Preg_match_result(`\[tr(?:=([\(\)\s%,#\w]+))?\]\s*\[td(?:=(\d{1,4}%?))?\]`, str, -1) {
			for _, v := range m {

				tmp.Reset()
				tmp.WriteString("<tr")
				if v[1] != "" {
					tmp.WriteString(` style="background-color: `)
					tmp.WriteString(v[1])
					tmp.WriteString(`"`)
				}
				tmp.WriteString("><td")
				if v[2] != "" {
					tmp.WriteString(` width="`)
					tmp.WriteString(v[2])
					tmp.WriteString(`"`)
				}
				tmp.WriteString(">")
				str = strings.Replace(str, v[0], tmp.String(), 1)

			}
		}

		for m, _ = Preg_match_result(`\[tr(?:=([\(\)\s%,#\w]+))?\]\s*\[td(?:=(\d{1,2}),(\d{1,2})(?:,(\d{1,4}%?))?)?\]`, str, -1); len(m) > 0; m, _ = Preg_match_result(`\[tr(?:=([\(\)\s%,#\w]+))?\]\s*\[td(?:=(\d{1,2}),(\d{1,2})(?:,(\d{1,4}%?))?)?\]`, str, -1) {
			for _, v := range m {

				tmp.Reset()
				tmp.WriteString("<tr")
				if v[1] != "" {
					tmp.WriteString(` style="background-color: `)
					tmp.WriteString(v[1])
					tmp.WriteString(`"`)
				}
				tmp.WriteString("><td")
				if v[2] != "" {
					tmp.WriteString(` colspan="`)
					tmp.WriteString(v[2])
					tmp.WriteString(`"`)
				}
				if v[3] != "" {
					tmp.WriteString(` rowspan="`)
					tmp.WriteString(v[3])
					tmp.WriteString(`"`)
				}
				if v[4] != "" {
					tmp.WriteString(` width="`)
					tmp.WriteString(v[4])
					tmp.WriteString(`"`)
				}
				tmp.WriteString(">")
				str = strings.Replace(str, v[0], tmp.String(), 1)

			}
		}

		for m, _ = Preg_match_result(`\[\/td\]\s*\[td(?:=(\d{1,4}%?))?\]`, str, -1); len(m) > 0; m, _ = Preg_match_result(`\[\/td\]\s*\[td(?:=(\d{1,4}%?))?\]`, str, -1) {
			for _, v := range m {

				tmp.Reset()
				tmp.WriteString("</td><td")

				if v[1] != "" {
					tmp.WriteString(` width="`)
					tmp.WriteString(v[1])
					tmp.WriteString(`"`)
				}
				tmp.WriteString(">")
				str = strings.Replace(str, v[0], tmp.String(), 1)

			}
		}

		for m, _ = Preg_match_result(`\[\/td\]\s*\[td(?:=(\d{1,2}),(\d{1,2})(?:,(\d{1,4}%?))?)?\]`, str, -1); len(m) > 0; m, _ = Preg_match_result(`\[\/td\]\s*\[td(?:=(\d{1,2}),(\d{1,2})(?:,(\d{1,4}%?))?)?\]`, str, -1) {
			for _, v := range m {

				tmp.Reset()
				tmp.WriteString("</td><td")

				if v[1] != "" {
					tmp.WriteString(` colspan="`)
					tmp.WriteString(v[1])
					tmp.WriteString(`"`)
				}
				if v[2] != "" {
					tmp.WriteString(` rowspan="`)
					tmp.WriteString(v[2])
					tmp.WriteString(`"`)
				}
				if v[3] != "" {
					tmp.WriteString(` width="`)
					tmp.WriteString(v[3])
					tmp.WriteString(`"`)
				}
				tmp.WriteString(">")
				str = strings.Replace(str, v[0], tmp.String(), 1)

			}
		}
		for m, _ = Preg_match_result(`\[tr(?:=([\(\)\s%,#\w]+))?\]\s*\[td(?:=(\d{1,4}%?))?\]`, str, -1); len(m) > 0; m, _ = Preg_match_result(`\[tr(?:=([\(\)\s%,#\w]+))?\]\s*\[td(?:=(\d{1,4}%?))?\]`, str, -1) {
			for _, v := range m {

				tmp.Reset()
				tmp.WriteString("<tr")
				if v[1] != "" {
					tmp.WriteString(` style="background-color: `)
					tmp.WriteString(v[1])
					tmp.WriteString(`"`)
				}
				tmp.WriteString("><td")
				if v[2] != "" {
					tmp.WriteString(` width="`)
					tmp.WriteString(v[2])
					tmp.WriteString(`"`)
				}
				tmp.WriteString(">")
				str = strings.Replace(str, v[0], tmp.String(), 1)

			}
		}

		str, _ = Preg_replace(`\[\/td\]\s*\[\/tr\]\s*`, "</td></tr>", str)
		str, _ = Preg_replace(`<td> <\/td>`, `<td>&nbsp;</td>`, str)
	}
	tmp.Reset()
	tmp.WriteString(`<table `)
	if width != `` {
		tmp.WriteString(`width="`)
		tmp.WriteString(width)
		tmp.WriteString(`"`)
	}
	tmp.WriteString(`class="t_table"`)
	if bgcolor != "" {
		tmp.WriteString(` style="background-color: `)
		tmp.WriteString(bgcolor)
		tmp.WriteString(`"`)
	}
	tmp.WriteString(simple)
	tmp.WriteString(`>`)
	tmp.WriteString(str)
	tmp.WriteString(`</table>`)
	s := tmp.String()
	bbcodebufpool.Put(tmp)
	return s
}
func init() {
	go func() {
		JsonUnmarshal([]byte(`[null,[null,[["1",":)","smile.gif","20","20","20"],["2",":(","sad.gif","20","20","20"],["3",":D","biggrin.gif","20","20","20"],["4",":'(","cry.gif","20","20","20"],["5",":@","huffy.gif","20","20","20"],["6",":o","shocked.gif","20","20","20"],["7",":P","tongue.gif","20","20","20"],["8",":$","shy.gif","20","20","20"],["9",";P","titter.gif","20","20","20"],["10",":L","sweat.gif","20","20","20"],["11",":Q","mad.gif","20","20","20"],["12",":lol","lol.gif","20","20","20"],["13",":loveliness:","loveliness.gif","20","20","20"],["14",":funk:","funk.gif","20","20","20"],["15",":curse:","curse.gif","20","20","20"],["16",":dizzy:","dizzy.gif","20","20","20"],["17",":shutup:","shutup.gif","20","20","20"],["18",":sleepy:","sleepy.gif","20","20","20"],["19",":hug:","hug.gif","20","20","20"],["20",":victory:","victory.gif","20","20","20"],["21",":time:","time.gif","20","20","20"],["22",":kiss:","kiss.gif","20","20","20"],["23",":handshake","handshake.gif","20","20","20"],["24",":call:","call.gif","20","20","20"]]],[null,[["25","{:2_25:}","01.gif","20","20","48"],["26","{:2_26:}","02.gif","20","20","48"],["27","{:2_27:}","03.gif","20","20","48"],["28","{:2_28:}","04.gif","20","20","48"],["29","{:2_29:}","05.gif","20","20","48"],["30","{:2_30:}","06.gif","20","20","48"],["31","{:2_31:}","07.gif","20","20","48"],["32","{:2_32:}","08.gif","20","20","48"],["33","{:2_33:}","09.gif","20","20","48"],["34","{:2_34:}","10.gif","20","20","48"],["35","{:2_35:}","11.gif","20","20","48"],["36","{:2_36:}","12.gif","20","20","48"],["37","{:2_37:}","13.gif","20","20","48"],["38","{:2_38:}","14.gif","20","20","48"],["39","{:2_39:}","15.gif","20","20","48"],["40","{:2_40:}","16.gif","20","20","48"]]],[null,[["41","{:3_41:}","01.gif","20","20","48"],["42","{:3_42:}","02.gif","20","20","48"],["43","{:3_43:}","03.gif","20","20","48"],["44","{:3_44:}","04.gif","20","20","48"],["45","{:3_45:}","05.gif","20","20","48"],["46","{:3_46:}","06.gif","20","20","48"],["47","{:3_47:}","07.gif","20","20","48"],["48","{:3_48:}","08.gif","20","20","48"],["49","{:3_49:}","09.gif","20","20","48"],["50","{:3_50:}","10.gif","20","20","48"],["51","{:3_51:}","11.gif","20","20","48"],["52","{:3_52:}","12.gif","20","20","48"],["53","{:3_53:}","13.gif","20","20","48"],["54","{:3_54:}","14.gif","20","20","48"],["55","{:3_55:}","15.gif","20","20","48"],["56","{:3_56:}","16.gif","20","20","48"],["57","{:3_57:}","17.gif","20","20","48"],["58","{:3_58:}","18.gif","20","20","48"],["59","{:3_59:}","19.gif","20","20","48"],["60","{:3_60:}","20.gif","20","20","48"],["61","{:3_61:}","21.gif","20","20","48"],["62","{:3_62:}","22.gif","20","20","48"],["63","{:3_63:}","23.gif","20","20","48"],["64","{:3_64:}","24.gif","20","20","48"]]]]`), &smilies_array)
		for typeid, v := range smilies_array {
			for _, vv := range v {
				for _, v3 := range vv {
					Smiliesarray = append(Smiliesarray, `<img src="`+STATICURL+`image/smiley/`+smilies_type["_"+strconv.Itoa(typeid)][1]+`/`+v3[2]+`" border="0" smilieid="`+v3[0]+`" alt="`+v3[1]+`" />`)
				}
			}
		}
	}()
}
