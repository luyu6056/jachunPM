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
	return fmt.Sprintf(`<script>alert("%s")</script>`, message, message)
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
