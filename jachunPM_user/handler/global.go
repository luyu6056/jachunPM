package handler

import (
	"jachunPM_user/db"
	"sync"
)

var userinfo_cache []*db.User
var userinfo_cache_lock sync.Mutex
