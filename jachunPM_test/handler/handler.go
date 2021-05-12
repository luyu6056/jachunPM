package handler

import (
	"jachunPM_test/db"
	"libraries"
	"protocol"
	"reflect"
	"strconv"
)

var HostConn *protocol.RpcClient

func Handler(in *protocol.Msg) {
	switch data := in.Data.(type) {
	case *protocol.MSG_TEST_bug_getCount:
		c, err := in.DB.Table(db.TABLE_BUG).Where(data.Where).Count()
		if err != nil {
			in.WriteErr(err)
		} else {
			out := protocol.GET_MSG_TEST_buf_getCount_result()
			out.Count = c
			in.SendResult(out)
			out.Put()
		}
	case *protocol.MSG_TEST_product_deleteBranch_check:
		out := protocol.GET_MSG_TEST_product_deleteBranch_result()
		out.Result = protocol.Success
		c, err := in.DB.Table(db.TABLE_BUG).Where("Branch=" + strconv.Itoa(int(data.BranchID)) + " and Deleted = 0").Count()
		defer func() {
			if err != nil {
				in.WriteErr(err)
				return
			}

			in.SendResult(out)
			out.Put()
		}()
		if c > 0 || err != nil {
			out.Result = protocol.Err_ProjectBranchCanNotDelete_BUG
			return
		}
		//$bug     = $this->dao->select('id')->from(TABLE_BUG)->where('branch')->eq($branchID)->andWhere('deleted')->eq(0)->limit(1)->fetch();
		// $case    = $this->dao->select('id')->from(TABLE_CASE)->where('branch')->eq($branchID)->andWhere('deleted')->eq(0)->limit(1)->fetch();
		//$release = $this->dao->select('id')->from(TABLE_RELEASE)->where('branch')->eq($branchID)->andWhere('deleted')->eq(0)->limit(1)->fetch();
		//$build   = $this->dao->select('id')->from(TABLE_BUILD)->where('branch')->eq($branchID)->andWhere('deleted')->eq(0)->limit(1)->fetch();
	case *protocol.MSG_TEST_testtask_getById:
		testtask_getById(data, in)
	case *protocol.MSG_TEST_testsuite_getById:
		testsuite_getById(data, in)
	case *protocol.MSG_TEST_bug_getPairs:
		bug_getPairs(data, in)
	case *protocol.MSG_TEST_bug_getCountByWhere:
		bug_getCountByWhere(data, in)
	default:
		libraries.ReleaseLog("未设置消息%s处理", reflect.TypeOf(data).Elem().Name())
	}
}
