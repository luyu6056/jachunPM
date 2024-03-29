package handler

import (
	"jachunPM_test/db"
	"protocol"
	"strconv"
)

var HostConn *protocol.RpcClient

func Handler(in *protocol.Msg) bool {
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
			return true
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
	case *protocol.MSG_TEST_CASE_getTaskCasePairs:
		case_getTaskCasePairs(data, in)
	case *protocol.MSG_TEST_doRawSelect:
		out := protocol.GET_MSG_TEST_doRawSelect_result()
		var err error
		if out.List, err = in.DB.Raw(data.Sql).SelectMap(); err != nil {
			in.WriteErr(err)
		} else {
			in.SendResult(out)
		}
	default:
		return false
	}
	return true
}
