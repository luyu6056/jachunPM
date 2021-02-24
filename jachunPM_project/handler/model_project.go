package handler

import (
	"jachunPM_project/db"
	"time"
)

func project_getProducts(projectID int32) (products []*db.Product, err error) {
	err = HostConn.DB.Table(db.TABLE_PRODUCT).Where(db.TABLE_PROJECT+".Id = ?", projectID).LeftJoin(db.TABLE_PROJECT).On(db.TABLE_PRODUCT + ".Id = " + db.TABLE_PROJECT + ".Product").Limit(0).Select(&products)
	return
}
func init() {
	go func() {
		time.Sleep(time.Second * 5)
		project_getProducts(1)
	}()
}
