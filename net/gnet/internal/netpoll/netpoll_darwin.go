// Copyright (c) 2019 Andy Pan
// Copyright (c) 2017 Joshua J Baker
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package netpoll

import (
	"os"

	"golang.org/x/sys/unix"
)

// SetKeepAlive sets the keepalive for the connection.
func SetKeepAlive(fd, secs int) error {
	if err := os.NewSyscallError("setsockopt", unix.SetsockoptInt(fd, unix.SOL_SOCKET, unix.SO_KEEPALIVE, 1)); err != nil {
		return err
	}
	switch err := os.NewSyscallError("setsockopt", unix.SetsockoptInt(fd, unix.IPPROTO_TCP, unix.TCP_KEEPINTVL, secs)); err {
	case nil, unix.ENOPROTOOPT: // OS X 10.7 and earlier don't support this option
	default:
		return err
	}
	return os.NewSyscallError("setsockopt", unix.SetsockoptInt(fd, unix.IPPROTO_TCP, unix.TCP_KEEPALIVE, secs))
}
