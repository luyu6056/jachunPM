package handler

import (
	"github.com/luyu6056/cache"
	"jachunPM_oa/db"
	"time"
)

func getHoliday(date time.Time) (res *db.Holiday, err error) {
	key := date.Format("holiday_cache_2006")
	holidayCache := cache.Hget(key, "holiday")
	var list []*db.Holiday
	if ok := holidayCache.Get("cache", &list); !ok {
		if err = HostConn.DB.Table(db.TABLE_HOLIDAY).Where("`Year`=?", date.Format("2006")).Select(&list); err != nil {
			return
		}
		holidayCache.Store("cache", list)
	}
	for _, l := range list {
		if l.Date == date {
			res = l
		}

	}
	holidayCache.Expire(1)
	return
}
