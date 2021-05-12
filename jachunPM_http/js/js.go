package js

import (
	"fmt"
	"protocol"
	"strings"
)

func Alert(format string, value ...interface{}) string {
	message := format
	if len(value) > 0 {
		message = fmt.Sprintf(format, value...)
	}
	message = strings.ReplaceAll(message, `"`, `\"`)
	return fmt.Sprintf(`<script>alert("%s");if(window.parent) window.parent.$.enableForm();</script>`, message)
}
func Reload(value ...string) string {
	buf := protocol.BufPoolGet()
	buf.WriteString("<script>")
	if len(value) == 1 {
		buf.WriteString(value[0])
	} else {
		buf.WriteString("self")
	}
	buf.WriteString(".location.reload(true);</script>")
	res := buf.String()
	protocol.BufPoolPut(buf)
	return res
}
func Confirm(message, okURL, cancleURL string, okcancleTarget ...string) string { // $okTarget = "self", $cancleTarget = "self"
	buf := protocol.BufPoolGet()
	buf.WriteString("<script>")
	buf.WriteString("if(confirm(\"")
	buf.WriteString(message)
	buf.WriteString("\")){\r\n")

	if strings.ToLower(okURL) == "back" {
		buf.WriteString("history.back(-1);")
	} else if okURL != "" {
		if len(okcancleTarget) > 0 {
			buf.WriteString(okcancleTarget[0])
		} else {
			buf.WriteString("self")
		}
		buf.WriteString(".location = '")
		buf.WriteString(okURL)
		buf.WriteString("';")
	}
	buf.WriteString("\r\n}else{\r\n")
	if strings.ToLower(cancleURL) == "back" {
		buf.WriteString("history.back(-1);")
	} else if cancleURL != "" {
		if len(okcancleTarget) > 1 {
			buf.WriteString(okcancleTarget[1])
		} else {
			buf.WriteString("self")
		}
		buf.WriteString(".location = '")
		buf.WriteString(cancleURL)
		buf.WriteString("';")

	}
	buf.WriteString("\r\n}</script>")
	res := buf.String()
	protocol.BufPoolPut(buf)
	return res
}
func Location(str string, window string) string {
	if window == "" {
		window = "self"
	}
	if str == "back" {
		return `<script>setTimeout("history.back()",1000)</script>`
	}
	return `<script>setTimeout('` + window + ".location.href=\"" + strings.ReplaceAll(strings.ReplaceAll(str, `"`, `\"`), "'", `\'`) + `',1000)"</script>`
}
func CloseModal(window, location, callback string) string {
	if window == "" {
		window = "self"
	}
	if location == "" {
		location = "this"
	}
	if callback == "" {
		callback = "null"
	}
	buf := protocol.BufPoolGet()
	buf.WriteString("<script>if(")
	buf.WriteString(window)
	buf.WriteString(".location.href == self.location.href){ ")
	buf.WriteString(window)
	buf.WriteString(".window.close();}else{")
	buf.WriteString(window)
	buf.WriteString(".$.cookie('selfClose', 1);")
	buf.WriteString(window)
	buf.WriteString(".$.closeModal(")
	buf.WriteString(callback)
	buf.WriteString(", '")
	buf.WriteString(location)
	buf.WriteString("');}</script>")
	res := buf.String()
	protocol.BufPoolPut(buf)
	return res

}
func Error(message string, full ...bool) string {
	isfull := true
	if len(full) == 1 {
		isfull = full[0]
	}
	buf := protocol.BufPoolGet()
	if isfull {
		buf.WriteString("<html><meta charset='utf-8'/><style>body{background:white}</style><script>")
	} else {
		buf.WriteString("<script>")
	}
	buf.WriteString("alert('")
	buf.WriteString(strings.ReplaceAll(message, "'", `\'`))
	buf.WriteString("');if(window.parent) window.parent.$.enableForm();</script>")
	res := buf.String()
	protocol.BufPoolPut(buf)
	return res
}
