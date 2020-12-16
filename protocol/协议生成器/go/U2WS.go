package protocol

import (
	"sync"
	"bbs/libraries"
)

const (
	CMD_MSG_U2WS_Login_Gethash = -814838689
	CMD_MSG_WS2U_Login_Gethash = 1595670514
	CMD_MSG_U2WS_UserLogin = 866031831
	CMD_MSG_WS2U_UserLogin = 1506905850
	CMD_MSG_U2WS_forum_index = 1000231794
	CMD_MSG_WS2U_forum_index = -278866301
	CMD_MSG_U2WS_GetHead = -446167730
	CMD_MSG_diy_info = -740027387
	CMD_MSG_diy_tab = -1451137317
	CMD_MSG_diy_frame = 814047536
	CMD_MSG_diy_attr = 471008838
	CMD_MSG_diy_title = -1370806890
	CMD_MSG_diy_first = 386708394
	CMD_MSG_diy_column = -1092979582
	CMD_MSG_diy_block = 107399903
	CMD_MSG_U2WS_diy_save = -1378201131
	CMD_MSG_U2WS_Checkseccode = -704884081
	CMD_MSG_U2WS_GetRegister = 1849790434
	CMD_MSG_WS2U_GetRegister = -1161885165
	CMD_MSG_U2WS_Register = -961222056
	CMD_MSG_WSU2_Register = 1935059451
	CMD_MSG_U2WS_CheckRegister = -204476842
	CMD_MSG_WS2U_CheckRegister = 1673174523
	CMD_MSG_diy_block_info = -1627801678
	CMD_MSG_block_style = -1418299074
	CMD_MSG_block_template = 578970688
	CMD_MSG_block_template_s = 256949028
	CMD_MSG_block_template_order = 1617053915
	CMD_MSG_block_item = 117099868
	CMD_MSG_block_item_field = 1448265555
	CMD_MSG_block_item_showstyle = 405274830
	CMD_MSG_block_item_showstyle_info = 1959565559
	CMD_MSG_forum_index_cart = 1761115921
	CMD_MSG_forum_extra = 177318912
	CMD_MSG_forum = 755023070
	CMD_MSG_forum_idname = -1334594352
	CMD_MSG_forum_lastpost = -35410442
	CMD_MSG_U2WS_forum = -1863283440
	CMD_MSG_WS2U_forum = 616977199
	CMD_MSG_forum_modrecommend = 1189303805
	CMD_MSG_U2WS_forum_modcp = 774094179
	CMD_MSG_U2WS_forum_recyclebin = -866489977
	CMD_MSG_forum_threadtype = 104018229
	CMD_MSG_forum_type = -1867493945
	CMD_MSG_forum_thread = 213874677
	CMD_MSG_forum_threadrush = 2145302771
	CMD_MSG_U2WS_forum_newthread = -1375386023
	CMD_MSG_WS2U_forum_newthread = 2090873870
	CMD_MSG_forum_savethread = 1315240586
	CMD_MSG_forum_replycredit = -1811127086
	CMD_MSG_forum_post_rush = 806931670
	CMD_MSG_forum_group = 711910816
	CMD_MSG_forum_attach = -480884745
	CMD_MSG_forum_imgattach = 1883458624
	CMD_MSG_forum_album = 2117521190
	CMD_MSG_U2WS_Getlogin = 56019999
	CMD_MSG_WS2U_Getlogin = -60812530
	CMD_MSG_U2WS_Forum_newthread_submit = 1149592144
	CMD_MSG_WS2U_Forum_newthread_submit = -796428201
	CMD_MSG_add_extcredits = -1451728518
	CMD_MSG_U2WS_forum_viewthread = -1106524571
	CMD_MSG_WS2U_forum_viewthread = 1272358720
	CMD_MSG_forum_thread_forum = 1303724428
	CMD_MSG_forum_post = 1189377635
	CMD_MSG_forum_post_medal = -2088438621
	CMD_MSG_postreview = -1141941364
	CMD_MSG_post_member_profile = -493836804
	CMD_MSG_post_ratelog = 1296010750
	CMD_MSG_post_ratelog_score = 239411683
	CMD_MSG_post_relateitem = 1127661264
	CMD_MSG_threadmod = 2087726285
	CMD_MSG_post_comment = -1934823944
	CMD_MSG_U2WS_threadfastpost = 1969915046
	CMD_MSG_WS2U_threadfastpost = -576901504
	CMD_MSG_U2WS_nextset = -1590634053
	CMD_MSG_WS2U_nextset = -1736066974
	CMD_MSG_U2WS_upload_image = 1031321379
	CMD_MSG_WS2U_upload_image = -149517524
	CMD_MSG_U2WS_upload_tmp_image = -895601579
	CMD_MSG_WS2U_upload_tmp_image = 1061377392
	CMD_MSG_U2WS_delete_attach = 1199011583
	CMD_MSG_WS2U_delete_attach = -687733422
	CMD_MSG_U2WS_threadmod = 1848054995
	CMD_MSG_U2WS_viewthreadmod = 885980426
	CMD_MSG_WS2U_viewthreadmod = -1531212121
	CMD_MSG_U2WS_forum_refresh = 1160234390
	CMD_MSG_U2WS_forum_carlist = 664970303
	CMD_MSG_WS2U_forum_carlist = -1210591342
	CMD_MSG_forum_cart = 399348313
	CMD_MSG_forum_cart_child = -1249765676
	CMD_MSG_U2WS_GetPostWarnList = -1077254663
	CMD_MSG_U2WS_space = 1017770727
	CMD_MSG_WS2U_space = -2002834216
	CMD_MSG_userprofiles = -2109377824
	CMD_MSG_usergroup = -1523481925
	CMD_MSG_U2WS_SpaceThread = -1799904226
	CMD_MSG_WS2U_SpaceThread = 1078534639
	CMD_MSG_SpaceThread = 1760833800
	CMD_MSG_SpacePost = -950215658
	CMD_MSG_U2WS_searchThread = 1503238307
	CMD_MSG_WS2U_searchThread = -1812486996
	CMD_MSG_searchThread = -1380880793
	CMD_MSG_WS2U_threadmod = 73925374
	CMD_MSG_U2WS_spacecp = 494965013
	CMD_MSG_WS2U_spacecp = 607514828
	CMD_MSG_U2WS_tpl_success = -1265560393
	CMD_MSG_WS2U_tpl_success = 1617924422
	CMD_MSG_U2WS_upload_avatar = 1329334923
	CMD_MSG_WS2U_upload_avatar = -548832986
	CMD_MSG_U2WS_Edit_Profile = 1027411378
	CMD_MSG_U2WS_RecommendThread = 1475637593
	CMD_MSG_WS2U_RecommendThread = -2058288370
	CMD_MSG_U2WS_SpacecpGroup = 127228310
	CMD_MSG_GroupIdName = 689990192
	CMD_MSG_WS2U_SpacecpGroup = -839162471
	CMD_MSG_U2WS_SpacecpForum = -276913250
	CMD_MSG_WS2U_SpacecpForum = 621846417
	CMD_MSG_SpacecpGroupPermission = 831351040
	CMD_MSG_U2WS_ChangePasswd_Gethash = 899199166
	CMD_MSG_WS2U_ChangePasswd_Gethash = 304149348
	CMD_MSG_U2WS_ChangePasswd = 54726888
	CMD_MSG_U2WS_Email_Verify = 305846558
	CMD_MSG_WS2U_Email_Verify = -665524975
	CMD_MSG_U2WS_LostPW = 1334552719
	CMD_MSG_WS2U_LostPW = -111561876
	CMD_MSG_U2WS_ResetPW = 1419751991
	CMD_MSG_WS2U_ResetPW = 1831523310
	CMD_MSG_U2WS_QQLoginUrl = -1619289269
	CMD_MSG_WS2U_QQLoginUrl = -623998936
	CMD_MSG_U2WS_QQLogin = -105825607
	CMD_MSG_WS2U_QQLogin = -1073456288
	CMD_MSG_U2WS_BindAccount = -2077234963
	CMD_MSG_WS2U_BindAccount = 1355902236
	CMD_MSG_U2WS_GetThreadBind = -302662362
	CMD_MSG_ThreadBind = 874163309
	CMD_MSG_WS2U_GetThreadBind = 2105593483
	CMD_MSG_U2WS_GetChangeBindUrl = -713593028
	CMD_MSG_WS2U_GetChangeBindUrl = 548015641
	CMD_MSG_U2WS_ChangeBind = 276821048
	CMD_MSG_WS2U_ChangeBind = 1439360859
	CMD_MSG_Poll_info = 292242050
	CMD_MSG_poll_option = -380025767
	CMD_MSG_U2WS_PollThread = -1328936102
	CMD_MSG_WS2U_PollThread_sucess = 1184955545
)

type MSG_U2WS_Login_Gethash struct {
	Name string
}

var pool_MSG_U2WS_Login_Gethash = sync.Pool{New: func() interface{} { return &MSG_U2WS_Login_Gethash{} }}

func GET_MSG_U2WS_Login_Gethash() *MSG_U2WS_Login_Gethash {
	return pool_MSG_U2WS_Login_Gethash.Get().(*MSG_U2WS_Login_Gethash)
}

func (data *MSG_U2WS_Login_Gethash) Put() {
	data.Name = ``
	pool_MSG_U2WS_Login_Gethash.Put(data)
}
func (data *MSG_U2WS_Login_Gethash) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_Login_Gethash)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_Login_Gethash(data, buf)
}

func WRITE_MSG_U2WS_Login_Gethash(data *MSG_U2WS_Login_Gethash, buf *libraries.MsgBuffer) {
	WRITE_string(data.Name, buf)
}

func READ_MSG_U2WS_Login_Gethash(buf *libraries.MsgBuffer) (data *MSG_U2WS_Login_Gethash) {
	data = pool_MSG_U2WS_Login_Gethash.Get().(*MSG_U2WS_Login_Gethash)
	data.Name = READ_string(buf)
	return
}

type MSG_WS2U_Login_Gethash struct {
	Hash string
	Hash2 string
}

var pool_MSG_WS2U_Login_Gethash = sync.Pool{New: func() interface{} { return &MSG_WS2U_Login_Gethash{} }}

func GET_MSG_WS2U_Login_Gethash() *MSG_WS2U_Login_Gethash {
	return pool_MSG_WS2U_Login_Gethash.Get().(*MSG_WS2U_Login_Gethash)
}

func (data *MSG_WS2U_Login_Gethash) Put() {
	data.Hash = ``
	data.Hash2 = ``
	pool_MSG_WS2U_Login_Gethash.Put(data)
}
func (data *MSG_WS2U_Login_Gethash) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_Login_Gethash)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_Login_Gethash(data, buf)
}

func WRITE_MSG_WS2U_Login_Gethash(data *MSG_WS2U_Login_Gethash, buf *libraries.MsgBuffer) {
	WRITE_string(data.Hash, buf)
	WRITE_string(data.Hash2, buf)
}

func READ_MSG_WS2U_Login_Gethash(buf *libraries.MsgBuffer) (data *MSG_WS2U_Login_Gethash) {
	data = pool_MSG_WS2U_Login_Gethash.Get().(*MSG_WS2U_Login_Gethash)
	data.Hash = READ_string(buf)
	data.Hash2 = READ_string(buf)
	return
}

type MSG_U2WS_UserLogin struct {
	Username string
	Passwd string
	Seccode string
	Type string
	State string
	Openid string
	Access_token string
}

var pool_MSG_U2WS_UserLogin = sync.Pool{New: func() interface{} { return &MSG_U2WS_UserLogin{} }}

func GET_MSG_U2WS_UserLogin() *MSG_U2WS_UserLogin {
	return pool_MSG_U2WS_UserLogin.Get().(*MSG_U2WS_UserLogin)
}

func (data *MSG_U2WS_UserLogin) Put() {
	data.Username = ``
	data.Passwd = ``
	data.Seccode = ``
	data.Type = ``
	data.State = ``
	data.Openid = ``
	data.Access_token = ``
	pool_MSG_U2WS_UserLogin.Put(data)
}
func (data *MSG_U2WS_UserLogin) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_UserLogin)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_UserLogin(data, buf)
}

func WRITE_MSG_U2WS_UserLogin(data *MSG_U2WS_UserLogin, buf *libraries.MsgBuffer) {
	WRITE_string(data.Username, buf)
	WRITE_string(data.Passwd, buf)
	WRITE_string(data.Seccode, buf)
	WRITE_string(data.Type, buf)
	WRITE_string(data.State, buf)
	WRITE_string(data.Openid, buf)
	WRITE_string(data.Access_token, buf)
}

func READ_MSG_U2WS_UserLogin(buf *libraries.MsgBuffer) (data *MSG_U2WS_UserLogin) {
	data = pool_MSG_U2WS_UserLogin.Get().(*MSG_U2WS_UserLogin)
	data.Username = READ_string(buf)
	data.Passwd = READ_string(buf)
	data.Seccode = READ_string(buf)
	data.Type = READ_string(buf)
	data.State = READ_string(buf)
	data.Openid = READ_string(buf)
	data.Access_token = READ_string(buf)
	return
}

type MSG_WS2U_UserLogin struct {
	Result int16
	Head *MSG_WS2U_Common_Head
	Token []byte
}

var pool_MSG_WS2U_UserLogin = sync.Pool{New: func() interface{} { return &MSG_WS2U_UserLogin{} }}

func GET_MSG_WS2U_UserLogin() *MSG_WS2U_UserLogin {
	return pool_MSG_WS2U_UserLogin.Get().(*MSG_WS2U_UserLogin)
}

func (data *MSG_WS2U_UserLogin) Put() {
	data.Result = 0
	if data.Head != nil {
		data.Head.Put()
		data.Head = nil
	}
	data.Token = data.Token[:0]
	pool_MSG_WS2U_UserLogin.Put(data)
}
func (data *MSG_WS2U_UserLogin) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_UserLogin)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_UserLogin(data, buf)
}

func WRITE_MSG_WS2U_UserLogin(data *MSG_WS2U_UserLogin, buf *libraries.MsgBuffer) {
	WRITE_int16(data.Result, buf)
	if data.Head == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_WS2U_Common_Head(data.Head, buf)
	}
	WRITE_int32(int32(len(data.Token)), buf)
	buf.Write(data.Token)
}

func READ_MSG_WS2U_UserLogin(buf *libraries.MsgBuffer) (data *MSG_WS2U_UserLogin) {
	data = pool_MSG_WS2U_UserLogin.Get().(*MSG_WS2U_UserLogin)
	data.Result = READ_int16(buf)
	Head_len := int(READ_int8(buf))
	if Head_len == 1 {
		data.Head = READ_MSG_WS2U_Common_Head(buf)
	}else{
		data.Head = nil
	}
	Token_len := int(READ_int32(buf))
	data.Token = make([]byte, Token_len)
	copy(data.Token,buf.Next(Token_len))
	return
}

type MSG_U2WS_forum_index struct {
	Gid int32
}

var pool_MSG_U2WS_forum_index = sync.Pool{New: func() interface{} { return &MSG_U2WS_forum_index{} }}

func GET_MSG_U2WS_forum_index() *MSG_U2WS_forum_index {
	return pool_MSG_U2WS_forum_index.Get().(*MSG_U2WS_forum_index)
}

func (data *MSG_U2WS_forum_index) Put() {
	data.Gid = 0
	pool_MSG_U2WS_forum_index.Put(data)
}
func (data *MSG_U2WS_forum_index) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_forum_index)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_forum_index(data, buf)
}

func WRITE_MSG_U2WS_forum_index(data *MSG_U2WS_forum_index, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Gid, buf)
}

func READ_MSG_U2WS_forum_index(buf *libraries.MsgBuffer) (data *MSG_U2WS_forum_index) {
	data = pool_MSG_U2WS_forum_index.Get().(*MSG_U2WS_forum_index)
	data.Gid = READ_int32(buf)
	return
}

type MSG_WS2U_forum_index struct {
	Threads int64
	Posts int64
	Onlinenum int32
	Usernum int32
	Lastname string
	Catlist []*MSG_forum_index_cart
	Diy_list []*MSG_diy_info
	History []*MSG_forum
	Recommend []*MSG_forum
}

var pool_MSG_WS2U_forum_index = sync.Pool{New: func() interface{} { return &MSG_WS2U_forum_index{} }}

func GET_MSG_WS2U_forum_index() *MSG_WS2U_forum_index {
	return pool_MSG_WS2U_forum_index.Get().(*MSG_WS2U_forum_index)
}

func (data *MSG_WS2U_forum_index) Put() {
	data.Threads = 0
	data.Posts = 0
	data.Onlinenum = 0
	data.Usernum = 0
	data.Lastname = ``
	for _,v := range data.Catlist {
		v.Put()
	}
	data.Catlist = data.Catlist[:0]
	for _,v := range data.Diy_list {
		v.Put()
	}
	data.Diy_list = data.Diy_list[:0]
	for _,v := range data.History {
		v.Put()
	}
	data.History = data.History[:0]
	for _,v := range data.Recommend {
		v.Put()
	}
	data.Recommend = data.Recommend[:0]
	pool_MSG_WS2U_forum_index.Put(data)
}
func (data *MSG_WS2U_forum_index) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_forum_index)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_forum_index(data, buf)
}

func WRITE_MSG_WS2U_forum_index(data *MSG_WS2U_forum_index, buf *libraries.MsgBuffer) {
	WRITE_int64(data.Threads, buf)
	WRITE_int64(data.Posts, buf)
	WRITE_int32(data.Onlinenum, buf)
	WRITE_int32(data.Usernum, buf)
	WRITE_string(data.Lastname, buf)
	WRITE_int32(int32(len(data.Catlist)), buf)
	for _, v := range data.Catlist{
		WRITE_MSG_forum_index_cart(v, buf)
	}
	WRITE_int32(int32(len(data.Diy_list)), buf)
	for _, v := range data.Diy_list{
		WRITE_MSG_diy_info(v, buf)
	}
	WRITE_int32(int32(len(data.History)), buf)
	for _, v := range data.History{
		WRITE_MSG_forum(v, buf)
	}
	WRITE_int32(int32(len(data.Recommend)), buf)
	for _, v := range data.Recommend{
		WRITE_MSG_forum(v, buf)
	}
}

func READ_MSG_WS2U_forum_index(buf *libraries.MsgBuffer) (data *MSG_WS2U_forum_index) {
	data = pool_MSG_WS2U_forum_index.Get().(*MSG_WS2U_forum_index)
	data.Threads = READ_int64(buf)
	data.Posts = READ_int64(buf)
	data.Onlinenum = READ_int32(buf)
	data.Usernum = READ_int32(buf)
	data.Lastname = READ_string(buf)
	Catlist_len := int(READ_int32(buf))
	for i := 0; i < Catlist_len; i++ {
		data.Catlist = append(data.Catlist, READ_MSG_forum_index_cart(buf))
	}
	Diy_list_len := int(READ_int32(buf))
	for i := 0; i < Diy_list_len; i++ {
		data.Diy_list = append(data.Diy_list, READ_MSG_diy_info(buf))
	}
	History_len := int(READ_int32(buf))
	for i := 0; i < History_len; i++ {
		data.History = append(data.History, READ_MSG_forum(buf))
	}
	Recommend_len := int(READ_int32(buf))
	for i := 0; i < Recommend_len; i++ {
		data.Recommend = append(data.Recommend, READ_MSG_forum(buf))
	}
	return
}

type MSG_U2WS_GetHead struct {
}

var pool_MSG_U2WS_GetHead = sync.Pool{New: func() interface{} { return &MSG_U2WS_GetHead{} }}

func GET_MSG_U2WS_GetHead() *MSG_U2WS_GetHead {
	return pool_MSG_U2WS_GetHead.Get().(*MSG_U2WS_GetHead)
}

func (data *MSG_U2WS_GetHead) Put() {
	pool_MSG_U2WS_GetHead.Put(data)
}
func (data *MSG_U2WS_GetHead) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_GetHead)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_GetHead(data, buf)
}

func WRITE_MSG_U2WS_GetHead(data *MSG_U2WS_GetHead, buf *libraries.MsgBuffer) {
}

func READ_MSG_U2WS_GetHead(buf *libraries.MsgBuffer) (data *MSG_U2WS_GetHead) {
	data = pool_MSG_U2WS_GetHead.Get().(*MSG_U2WS_GetHead)
	return
}

type MSG_diy_info struct {
	Id string
	Tab *MSG_diy_tab
	Frame *MSG_diy_frame
}

var pool_MSG_diy_info = sync.Pool{New: func() interface{} { return &MSG_diy_info{} }}

func GET_MSG_diy_info() *MSG_diy_info {
	return pool_MSG_diy_info.Get().(*MSG_diy_info)
}

func (data *MSG_diy_info) Put() {
	data.Id = ``
	if data.Tab != nil {
		data.Tab.Put()
		data.Tab = nil
	}
	if data.Frame != nil {
		data.Frame.Put()
		data.Frame = nil
	}
	pool_MSG_diy_info.Put(data)
}
func (data *MSG_diy_info) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_diy_info)
	WRITE_int32(cmd, buf)
	WRITE_MSG_diy_info(data, buf)
}

func WRITE_MSG_diy_info(data *MSG_diy_info, buf *libraries.MsgBuffer) {
	WRITE_string(data.Id, buf)
	if data.Tab == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_diy_tab(data.Tab, buf)
	}
	if data.Frame == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_diy_frame(data.Frame, buf)
	}
}

func READ_MSG_diy_info(buf *libraries.MsgBuffer) (data *MSG_diy_info) {
	data = pool_MSG_diy_info.Get().(*MSG_diy_info)
	data.Id = READ_string(buf)
	Tab_len := int(READ_int8(buf))
	if Tab_len == 1 {
		data.Tab = READ_MSG_diy_tab(buf)
	}else{
		data.Tab = nil
	}
	Frame_len := int(READ_int8(buf))
	if Frame_len == 1 {
		data.Frame = READ_MSG_diy_frame(buf)
	}else{
		data.Frame = nil
	}
	return
}

type MSG_diy_tab struct {
	Attr *MSG_diy_attr
	Column *MSG_diy_column
}

var pool_MSG_diy_tab = sync.Pool{New: func() interface{} { return &MSG_diy_tab{} }}

func GET_MSG_diy_tab() *MSG_diy_tab {
	return pool_MSG_diy_tab.Get().(*MSG_diy_tab)
}

func (data *MSG_diy_tab) Put() {
	if data.Attr != nil {
		data.Attr.Put()
		data.Attr = nil
	}
	if data.Column != nil {
		data.Column.Put()
		data.Column = nil
	}
	pool_MSG_diy_tab.Put(data)
}
func (data *MSG_diy_tab) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_diy_tab)
	WRITE_int32(cmd, buf)
	WRITE_MSG_diy_tab(data, buf)
}

func WRITE_MSG_diy_tab(data *MSG_diy_tab, buf *libraries.MsgBuffer) {
	if data.Attr == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_diy_attr(data.Attr, buf)
	}
	if data.Column == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_diy_column(data.Column, buf)
	}
}

func READ_MSG_diy_tab(buf *libraries.MsgBuffer) (data *MSG_diy_tab) {
	data = pool_MSG_diy_tab.Get().(*MSG_diy_tab)
	Attr_len := int(READ_int8(buf))
	if Attr_len == 1 {
		data.Attr = READ_MSG_diy_attr(buf)
	}else{
		data.Attr = nil
	}
	Column_len := int(READ_int8(buf))
	if Column_len == 1 {
		data.Column = READ_MSG_diy_column(buf)
	}else{
		data.Column = nil
	}
	return
}

type MSG_diy_frame struct {
	Attr *MSG_diy_attr
	Column *MSG_diy_column
}

var pool_MSG_diy_frame = sync.Pool{New: func() interface{} { return &MSG_diy_frame{} }}

func GET_MSG_diy_frame() *MSG_diy_frame {
	return pool_MSG_diy_frame.Get().(*MSG_diy_frame)
}

func (data *MSG_diy_frame) Put() {
	if data.Attr != nil {
		data.Attr.Put()
		data.Attr = nil
	}
	if data.Column != nil {
		data.Column.Put()
		data.Column = nil
	}
	pool_MSG_diy_frame.Put(data)
}
func (data *MSG_diy_frame) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_diy_frame)
	WRITE_int32(cmd, buf)
	WRITE_MSG_diy_frame(data, buf)
}

func WRITE_MSG_diy_frame(data *MSG_diy_frame, buf *libraries.MsgBuffer) {
	if data.Attr == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_diy_attr(data.Attr, buf)
	}
	if data.Column == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_diy_column(data.Column, buf)
	}
}

func READ_MSG_diy_frame(buf *libraries.MsgBuffer) (data *MSG_diy_frame) {
	data = pool_MSG_diy_frame.Get().(*MSG_diy_frame)
	Attr_len := int(READ_int8(buf))
	if Attr_len == 1 {
		data.Attr = READ_MSG_diy_attr(buf)
	}else{
		data.Attr = nil
	}
	Column_len := int(READ_int8(buf))
	if Column_len == 1 {
		data.Column = READ_MSG_diy_column(buf)
	}else{
		data.Column = nil
	}
	return
}

type MSG_diy_attr struct {
	Name string
	Moveable string
	ClassName string
	Titles *MSG_diy_title
}

var pool_MSG_diy_attr = sync.Pool{New: func() interface{} { return &MSG_diy_attr{} }}

func GET_MSG_diy_attr() *MSG_diy_attr {
	return pool_MSG_diy_attr.Get().(*MSG_diy_attr)
}

func (data *MSG_diy_attr) Put() {
	data.Name = ``
	data.Moveable = ``
	data.ClassName = ``
	if data.Titles != nil {
		data.Titles.Put()
		data.Titles = nil
	}
	pool_MSG_diy_attr.Put(data)
}
func (data *MSG_diy_attr) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_diy_attr)
	WRITE_int32(cmd, buf)
	WRITE_MSG_diy_attr(data, buf)
}

func WRITE_MSG_diy_attr(data *MSG_diy_attr, buf *libraries.MsgBuffer) {
	WRITE_string(data.Name, buf)
	WRITE_string(data.Moveable, buf)
	WRITE_string(data.ClassName, buf)
	if data.Titles == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_diy_title(data.Titles, buf)
	}
}

func READ_MSG_diy_attr(buf *libraries.MsgBuffer) (data *MSG_diy_attr) {
	data = pool_MSG_diy_attr.Get().(*MSG_diy_attr)
	data.Name = READ_string(buf)
	data.Moveable = READ_string(buf)
	data.ClassName = READ_string(buf)
	Titles_len := int(READ_int8(buf))
	if Titles_len == 1 {
		data.Titles = READ_MSG_diy_title(buf)
	}else{
		data.Titles = nil
	}
	return
}

type MSG_diy_title struct {
	Style string
	ClassName []string
	SwitchType []string
	First *MSG_diy_first
}

var pool_MSG_diy_title = sync.Pool{New: func() interface{} { return &MSG_diy_title{} }}

func GET_MSG_diy_title() *MSG_diy_title {
	return pool_MSG_diy_title.Get().(*MSG_diy_title)
}

func (data *MSG_diy_title) Put() {
	data.Style = ``
	data.ClassName = data.ClassName[:0]
	data.SwitchType = data.SwitchType[:0]
	if data.First != nil {
		data.First.Put()
		data.First = nil
	}
	pool_MSG_diy_title.Put(data)
}
func (data *MSG_diy_title) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_diy_title)
	WRITE_int32(cmd, buf)
	WRITE_MSG_diy_title(data, buf)
}

func WRITE_MSG_diy_title(data *MSG_diy_title, buf *libraries.MsgBuffer) {
	WRITE_string(data.Style, buf)
	WRITE_int32(int32(len(data.ClassName)), buf)
	for _, v := range data.ClassName{
		WRITE_string(v, buf)
	}
	WRITE_int32(int32(len(data.SwitchType)), buf)
	for _, v := range data.SwitchType{
		WRITE_string(v, buf)
	}
	if data.First == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_diy_first(data.First, buf)
	}
}

func READ_MSG_diy_title(buf *libraries.MsgBuffer) (data *MSG_diy_title) {
	data = pool_MSG_diy_title.Get().(*MSG_diy_title)
	data.Style = READ_string(buf)
	ClassName_len := int(READ_int32(buf))
	for i := 0; i < ClassName_len; i++ {
		data.ClassName = append(data.ClassName, READ_string(buf))
	}
	SwitchType_len := int(READ_int32(buf))
	for i := 0; i < SwitchType_len; i++ {
		data.SwitchType = append(data.SwitchType, READ_string(buf))
	}
	First_len := int(READ_int8(buf))
	if First_len == 1 {
		data.First = READ_MSG_diy_first(buf)
	}else{
		data.First = nil
	}
	return
}

type MSG_diy_first struct {
	Text string
	Href string
	Color string
	Float string
	Margin string
	Font_size string
	ClassName string
	Src string
}

var pool_MSG_diy_first = sync.Pool{New: func() interface{} { return &MSG_diy_first{} }}

func GET_MSG_diy_first() *MSG_diy_first {
	return pool_MSG_diy_first.Get().(*MSG_diy_first)
}

func (data *MSG_diy_first) Put() {
	data.Text = ``
	data.Href = ``
	data.Color = ``
	data.Float = ``
	data.Margin = ``
	data.Font_size = ``
	data.ClassName = ``
	data.Src = ``
	pool_MSG_diy_first.Put(data)
}
func (data *MSG_diy_first) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_diy_first)
	WRITE_int32(cmd, buf)
	WRITE_MSG_diy_first(data, buf)
}

func WRITE_MSG_diy_first(data *MSG_diy_first, buf *libraries.MsgBuffer) {
	WRITE_string(data.Text, buf)
	WRITE_string(data.Href, buf)
	WRITE_string(data.Color, buf)
	WRITE_string(data.Float, buf)
	WRITE_string(data.Margin, buf)
	WRITE_string(data.Font_size, buf)
	WRITE_string(data.ClassName, buf)
	WRITE_string(data.Src, buf)
}

func READ_MSG_diy_first(buf *libraries.MsgBuffer) (data *MSG_diy_first) {
	data = pool_MSG_diy_first.Get().(*MSG_diy_first)
	data.Text = READ_string(buf)
	data.Href = READ_string(buf)
	data.Color = READ_string(buf)
	data.Float = READ_string(buf)
	data.Margin = READ_string(buf)
	data.Font_size = READ_string(buf)
	data.ClassName = READ_string(buf)
	data.Src = READ_string(buf)
	return
}

type MSG_diy_column struct {
	Attr *MSG_diy_attr
	Block []*MSG_diy_block
}

var pool_MSG_diy_column = sync.Pool{New: func() interface{} { return &MSG_diy_column{} }}

func GET_MSG_diy_column() *MSG_diy_column {
	return pool_MSG_diy_column.Get().(*MSG_diy_column)
}

func (data *MSG_diy_column) Put() {
	if data.Attr != nil {
		data.Attr.Put()
		data.Attr = nil
	}
	for _,v := range data.Block {
		v.Put()
	}
	data.Block = data.Block[:0]
	pool_MSG_diy_column.Put(data)
}
func (data *MSG_diy_column) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_diy_column)
	WRITE_int32(cmd, buf)
	WRITE_MSG_diy_column(data, buf)
}

func WRITE_MSG_diy_column(data *MSG_diy_column, buf *libraries.MsgBuffer) {
	if data.Attr == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_diy_attr(data.Attr, buf)
	}
	WRITE_int32(int32(len(data.Block)), buf)
	for _, v := range data.Block{
		WRITE_MSG_diy_block(v, buf)
	}
}

func READ_MSG_diy_column(buf *libraries.MsgBuffer) (data *MSG_diy_column) {
	data = pool_MSG_diy_column.Get().(*MSG_diy_column)
	Attr_len := int(READ_int8(buf))
	if Attr_len == 1 {
		data.Attr = READ_MSG_diy_attr(buf)
	}else{
		data.Attr = nil
	}
	Block_len := int(READ_int32(buf))
	for i := 0; i < Block_len; i++ {
		data.Block = append(data.Block, READ_MSG_diy_block(buf))
	}
	return
}

type MSG_diy_block struct {
	Attr *MSG_diy_attr
	Info *MSG_diy_block_info
}

var pool_MSG_diy_block = sync.Pool{New: func() interface{} { return &MSG_diy_block{} }}

func GET_MSG_diy_block() *MSG_diy_block {
	return pool_MSG_diy_block.Get().(*MSG_diy_block)
}

func (data *MSG_diy_block) Put() {
	if data.Attr != nil {
		data.Attr.Put()
		data.Attr = nil
	}
	if data.Info != nil {
		data.Info.Put()
		data.Info = nil
	}
	pool_MSG_diy_block.Put(data)
}
func (data *MSG_diy_block) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_diy_block)
	WRITE_int32(cmd, buf)
	WRITE_MSG_diy_block(data, buf)
}

func WRITE_MSG_diy_block(data *MSG_diy_block, buf *libraries.MsgBuffer) {
	if data.Attr == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_diy_attr(data.Attr, buf)
	}
	if data.Info == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_diy_block_info(data.Info, buf)
	}
}

func READ_MSG_diy_block(buf *libraries.MsgBuffer) (data *MSG_diy_block) {
	data = pool_MSG_diy_block.Get().(*MSG_diy_block)
	Attr_len := int(READ_int8(buf))
	if Attr_len == 1 {
		data.Attr = READ_MSG_diy_attr(buf)
	}else{
		data.Attr = nil
	}
	Info_len := int(READ_int8(buf))
	if Info_len == 1 {
		data.Info = READ_MSG_diy_block_info(buf)
	}else{
		data.Info = nil
	}
	return
}

type MSG_U2WS_diy_save struct {
	List string
	TemplateName string
}

var pool_MSG_U2WS_diy_save = sync.Pool{New: func() interface{} { return &MSG_U2WS_diy_save{} }}

func GET_MSG_U2WS_diy_save() *MSG_U2WS_diy_save {
	return pool_MSG_U2WS_diy_save.Get().(*MSG_U2WS_diy_save)
}

func (data *MSG_U2WS_diy_save) Put() {
	data.List = ``
	data.TemplateName = ``
	pool_MSG_U2WS_diy_save.Put(data)
}
func (data *MSG_U2WS_diy_save) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_diy_save)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_diy_save(data, buf)
}

func WRITE_MSG_U2WS_diy_save(data *MSG_U2WS_diy_save, buf *libraries.MsgBuffer) {
	WRITE_string(data.List, buf)
	WRITE_string(data.TemplateName, buf)
}

func READ_MSG_U2WS_diy_save(buf *libraries.MsgBuffer) (data *MSG_U2WS_diy_save) {
	data = pool_MSG_U2WS_diy_save.Get().(*MSG_U2WS_diy_save)
	data.List = READ_string(buf)
	data.TemplateName = READ_string(buf)
	return
}

type MSG_U2WS_Checkseccode struct {
	Code string
}

var pool_MSG_U2WS_Checkseccode = sync.Pool{New: func() interface{} { return &MSG_U2WS_Checkseccode{} }}

func GET_MSG_U2WS_Checkseccode() *MSG_U2WS_Checkseccode {
	return pool_MSG_U2WS_Checkseccode.Get().(*MSG_U2WS_Checkseccode)
}

func (data *MSG_U2WS_Checkseccode) Put() {
	data.Code = ``
	pool_MSG_U2WS_Checkseccode.Put(data)
}
func (data *MSG_U2WS_Checkseccode) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_Checkseccode)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_Checkseccode(data, buf)
}

func WRITE_MSG_U2WS_Checkseccode(data *MSG_U2WS_Checkseccode, buf *libraries.MsgBuffer) {
	WRITE_string(data.Code, buf)
}

func READ_MSG_U2WS_Checkseccode(buf *libraries.MsgBuffer) (data *MSG_U2WS_Checkseccode) {
	data = pool_MSG_U2WS_Checkseccode.Get().(*MSG_U2WS_Checkseccode)
	data.Code = READ_string(buf)
	return
}

type MSG_U2WS_GetRegister struct {
}

var pool_MSG_U2WS_GetRegister = sync.Pool{New: func() interface{} { return &MSG_U2WS_GetRegister{} }}

func GET_MSG_U2WS_GetRegister() *MSG_U2WS_GetRegister {
	return pool_MSG_U2WS_GetRegister.Get().(*MSG_U2WS_GetRegister)
}

func (data *MSG_U2WS_GetRegister) Put() {
	pool_MSG_U2WS_GetRegister.Put(data)
}
func (data *MSG_U2WS_GetRegister) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_GetRegister)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_GetRegister(data, buf)
}

func WRITE_MSG_U2WS_GetRegister(data *MSG_U2WS_GetRegister, buf *libraries.MsgBuffer) {
}

func READ_MSG_U2WS_GetRegister(buf *libraries.MsgBuffer) (data *MSG_U2WS_GetRegister) {
	data = pool_MSG_U2WS_GetRegister.Get().(*MSG_U2WS_GetRegister)
	return
}

type MSG_WS2U_GetRegister struct {
}

var pool_MSG_WS2U_GetRegister = sync.Pool{New: func() interface{} { return &MSG_WS2U_GetRegister{} }}

func GET_MSG_WS2U_GetRegister() *MSG_WS2U_GetRegister {
	return pool_MSG_WS2U_GetRegister.Get().(*MSG_WS2U_GetRegister)
}

func (data *MSG_WS2U_GetRegister) Put() {
	pool_MSG_WS2U_GetRegister.Put(data)
}
func (data *MSG_WS2U_GetRegister) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_GetRegister)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_GetRegister(data, buf)
}

func WRITE_MSG_WS2U_GetRegister(data *MSG_WS2U_GetRegister, buf *libraries.MsgBuffer) {
}

func READ_MSG_WS2U_GetRegister(buf *libraries.MsgBuffer) (data *MSG_WS2U_GetRegister) {
	data = pool_MSG_WS2U_GetRegister.Get().(*MSG_WS2U_GetRegister)
	return
}

type MSG_U2WS_Register struct {
	Username string
	Passwd string
	Email string
	Seccode string
	Type string
	State string
	Openid string
	Access_token string
}

var pool_MSG_U2WS_Register = sync.Pool{New: func() interface{} { return &MSG_U2WS_Register{} }}

func GET_MSG_U2WS_Register() *MSG_U2WS_Register {
	return pool_MSG_U2WS_Register.Get().(*MSG_U2WS_Register)
}

func (data *MSG_U2WS_Register) Put() {
	data.Username = ``
	data.Passwd = ``
	data.Email = ``
	data.Seccode = ``
	data.Type = ``
	data.State = ``
	data.Openid = ``
	data.Access_token = ``
	pool_MSG_U2WS_Register.Put(data)
}
func (data *MSG_U2WS_Register) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_Register)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_Register(data, buf)
}

func WRITE_MSG_U2WS_Register(data *MSG_U2WS_Register, buf *libraries.MsgBuffer) {
	WRITE_string(data.Username, buf)
	WRITE_string(data.Passwd, buf)
	WRITE_string(data.Email, buf)
	WRITE_string(data.Seccode, buf)
	WRITE_string(data.Type, buf)
	WRITE_string(data.State, buf)
	WRITE_string(data.Openid, buf)
	WRITE_string(data.Access_token, buf)
}

func READ_MSG_U2WS_Register(buf *libraries.MsgBuffer) (data *MSG_U2WS_Register) {
	data = pool_MSG_U2WS_Register.Get().(*MSG_U2WS_Register)
	data.Username = READ_string(buf)
	data.Passwd = READ_string(buf)
	data.Email = READ_string(buf)
	data.Seccode = READ_string(buf)
	data.Type = READ_string(buf)
	data.State = READ_string(buf)
	data.Openid = READ_string(buf)
	data.Access_token = READ_string(buf)
	return
}

type MSG_WSU2_Register struct {
	Head *MSG_WS2U_Common_Head
}

var pool_MSG_WSU2_Register = sync.Pool{New: func() interface{} { return &MSG_WSU2_Register{} }}

func GET_MSG_WSU2_Register() *MSG_WSU2_Register {
	return pool_MSG_WSU2_Register.Get().(*MSG_WSU2_Register)
}

func (data *MSG_WSU2_Register) Put() {
	if data.Head != nil {
		data.Head.Put()
		data.Head = nil
	}
	pool_MSG_WSU2_Register.Put(data)
}
func (data *MSG_WSU2_Register) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WSU2_Register)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WSU2_Register(data, buf)
}

func WRITE_MSG_WSU2_Register(data *MSG_WSU2_Register, buf *libraries.MsgBuffer) {
	if data.Head == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_WS2U_Common_Head(data.Head, buf)
	}
}

func READ_MSG_WSU2_Register(buf *libraries.MsgBuffer) (data *MSG_WSU2_Register) {
	data = pool_MSG_WSU2_Register.Get().(*MSG_WSU2_Register)
	Head_len := int(READ_int8(buf))
	if Head_len == 1 {
		data.Head = READ_MSG_WS2U_Common_Head(buf)
	}else{
		data.Head = nil
	}
	return
}

type MSG_U2WS_CheckRegister struct {
	Type int8
	Name string
}

var pool_MSG_U2WS_CheckRegister = sync.Pool{New: func() interface{} { return &MSG_U2WS_CheckRegister{} }}

func GET_MSG_U2WS_CheckRegister() *MSG_U2WS_CheckRegister {
	return pool_MSG_U2WS_CheckRegister.Get().(*MSG_U2WS_CheckRegister)
}

func (data *MSG_U2WS_CheckRegister) Put() {
	data.Type = 0
	data.Name = ``
	pool_MSG_U2WS_CheckRegister.Put(data)
}
func (data *MSG_U2WS_CheckRegister) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_CheckRegister)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_CheckRegister(data, buf)
}

func WRITE_MSG_U2WS_CheckRegister(data *MSG_U2WS_CheckRegister, buf *libraries.MsgBuffer) {
	WRITE_int8(data.Type, buf)
	WRITE_string(data.Name, buf)
}

func READ_MSG_U2WS_CheckRegister(buf *libraries.MsgBuffer) (data *MSG_U2WS_CheckRegister) {
	data = pool_MSG_U2WS_CheckRegister.Get().(*MSG_U2WS_CheckRegister)
	data.Type = READ_int8(buf)
	data.Name = READ_string(buf)
	return
}

type MSG_WS2U_CheckRegister struct {
	Result int16
}

var pool_MSG_WS2U_CheckRegister = sync.Pool{New: func() interface{} { return &MSG_WS2U_CheckRegister{} }}

func GET_MSG_WS2U_CheckRegister() *MSG_WS2U_CheckRegister {
	return pool_MSG_WS2U_CheckRegister.Get().(*MSG_WS2U_CheckRegister)
}

func (data *MSG_WS2U_CheckRegister) Put() {
	data.Result = 0
	pool_MSG_WS2U_CheckRegister.Put(data)
}
func (data *MSG_WS2U_CheckRegister) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_CheckRegister)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_CheckRegister(data, buf)
}

func WRITE_MSG_WS2U_CheckRegister(data *MSG_WS2U_CheckRegister, buf *libraries.MsgBuffer) {
	WRITE_int16(data.Result, buf)
}

func READ_MSG_WS2U_CheckRegister(buf *libraries.MsgBuffer) (data *MSG_WS2U_CheckRegister) {
	data = pool_MSG_WS2U_CheckRegister.Get().(*MSG_WS2U_CheckRegister)
	data.Result = READ_int16(buf)
	return
}

type MSG_diy_block_info struct {
	Name string
	Summary string
	Blockclass string
	Style *MSG_block_style
	Title string
	Bid int32
	Itemlist []*MSG_block_item
	Hidedisplay int8
}

var pool_MSG_diy_block_info = sync.Pool{New: func() interface{} { return &MSG_diy_block_info{} }}

func GET_MSG_diy_block_info() *MSG_diy_block_info {
	return pool_MSG_diy_block_info.Get().(*MSG_diy_block_info)
}

func (data *MSG_diy_block_info) Put() {
	data.Name = ``
	data.Summary = ``
	data.Blockclass = ``
	if data.Style != nil {
		data.Style.Put()
		data.Style = nil
	}
	data.Title = ``
	data.Bid = 0
	for _,v := range data.Itemlist {
		v.Put()
	}
	data.Itemlist = data.Itemlist[:0]
	data.Hidedisplay = 0
	pool_MSG_diy_block_info.Put(data)
}
func (data *MSG_diy_block_info) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_diy_block_info)
	WRITE_int32(cmd, buf)
	WRITE_MSG_diy_block_info(data, buf)
}

func WRITE_MSG_diy_block_info(data *MSG_diy_block_info, buf *libraries.MsgBuffer) {
	WRITE_string(data.Name, buf)
	WRITE_string(data.Summary, buf)
	WRITE_string(data.Blockclass, buf)
	if data.Style == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_block_style(data.Style, buf)
	}
	WRITE_string(data.Title, buf)
	WRITE_int32(data.Bid, buf)
	WRITE_int32(int32(len(data.Itemlist)), buf)
	for _, v := range data.Itemlist{
		WRITE_MSG_block_item(v, buf)
	}
	WRITE_int8(data.Hidedisplay, buf)
}

func READ_MSG_diy_block_info(buf *libraries.MsgBuffer) (data *MSG_diy_block_info) {
	data = pool_MSG_diy_block_info.Get().(*MSG_diy_block_info)
	data.Name = READ_string(buf)
	data.Summary = READ_string(buf)
	data.Blockclass = READ_string(buf)
	Style_len := int(READ_int8(buf))
	if Style_len == 1 {
		data.Style = READ_MSG_block_style(buf)
	}else{
		data.Style = nil
	}
	data.Title = READ_string(buf)
	data.Bid = READ_int32(buf)
	Itemlist_len := int(READ_int32(buf))
	for i := 0; i < Itemlist_len; i++ {
		data.Itemlist = append(data.Itemlist, READ_MSG_block_item(buf))
	}
	data.Hidedisplay = READ_int8(buf)
	return
}

type MSG_block_style struct {
	Template *MSG_block_template
	Moreurl int8
	Fields []string
}

var pool_MSG_block_style = sync.Pool{New: func() interface{} { return &MSG_block_style{} }}

func GET_MSG_block_style() *MSG_block_style {
	return pool_MSG_block_style.Get().(*MSG_block_style)
}

func (data *MSG_block_style) Put() {
	if data.Template != nil {
		data.Template.Put()
		data.Template = nil
	}
	data.Moreurl = 0
	data.Fields = data.Fields[:0]
	pool_MSG_block_style.Put(data)
}
func (data *MSG_block_style) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_block_style)
	WRITE_int32(cmd, buf)
	WRITE_MSG_block_style(data, buf)
}

func WRITE_MSG_block_style(data *MSG_block_style, buf *libraries.MsgBuffer) {
	if data.Template == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_block_template(data.Template, buf)
	}
	WRITE_int8(data.Moreurl, buf)
	WRITE_int32(int32(len(data.Fields)), buf)
	for _, v := range data.Fields{
		WRITE_string(v, buf)
	}
}

func READ_MSG_block_style(buf *libraries.MsgBuffer) (data *MSG_block_style) {
	data = pool_MSG_block_style.Get().(*MSG_block_style)
	Template_len := int(READ_int8(buf))
	if Template_len == 1 {
		data.Template = READ_MSG_block_template(buf)
	}else{
		data.Template = nil
	}
	data.Moreurl = READ_int8(buf)
	Fields_len := int(READ_int32(buf))
	for i := 0; i < Fields_len; i++ {
		data.Fields = append(data.Fields, READ_string(buf))
	}
	return
}

type MSG_block_template struct {
	Raw string
	Footer string
	Header string
	Indexplus *MSG_block_template_s
	Index *MSG_block_template_s
	Orderplus []*MSG_block_template_s
	Order *MSG_block_template_s
	Loopplus []string
	Loop string
}

var pool_MSG_block_template = sync.Pool{New: func() interface{} { return &MSG_block_template{} }}

func GET_MSG_block_template() *MSG_block_template {
	return pool_MSG_block_template.Get().(*MSG_block_template)
}

func (data *MSG_block_template) Put() {
	data.Raw = ``
	data.Footer = ``
	data.Header = ``
	if data.Indexplus != nil {
		data.Indexplus.Put()
		data.Indexplus = nil
	}
	if data.Index != nil {
		data.Index.Put()
		data.Index = nil
	}
	for _,v := range data.Orderplus {
		v.Put()
	}
	data.Orderplus = data.Orderplus[:0]
	if data.Order != nil {
		data.Order.Put()
		data.Order = nil
	}
	data.Loopplus = data.Loopplus[:0]
	data.Loop = ``
	pool_MSG_block_template.Put(data)
}
func (data *MSG_block_template) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_block_template)
	WRITE_int32(cmd, buf)
	WRITE_MSG_block_template(data, buf)
}

func WRITE_MSG_block_template(data *MSG_block_template, buf *libraries.MsgBuffer) {
	WRITE_string(data.Raw, buf)
	WRITE_string(data.Footer, buf)
	WRITE_string(data.Header, buf)
	if data.Indexplus == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_block_template_s(data.Indexplus, buf)
	}
	if data.Index == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_block_template_s(data.Index, buf)
	}
	WRITE_int32(int32(len(data.Orderplus)), buf)
	for _, v := range data.Orderplus{
		WRITE_MSG_block_template_s(v, buf)
	}
	if data.Order == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_block_template_s(data.Order, buf)
	}
	WRITE_int32(int32(len(data.Loopplus)), buf)
	for _, v := range data.Loopplus{
		WRITE_string(v, buf)
	}
	WRITE_string(data.Loop, buf)
}

func READ_MSG_block_template(buf *libraries.MsgBuffer) (data *MSG_block_template) {
	data = pool_MSG_block_template.Get().(*MSG_block_template)
	data.Raw = READ_string(buf)
	data.Footer = READ_string(buf)
	data.Header = READ_string(buf)
	Indexplus_len := int(READ_int8(buf))
	if Indexplus_len == 1 {
		data.Indexplus = READ_MSG_block_template_s(buf)
	}else{
		data.Indexplus = nil
	}
	Index_len := int(READ_int8(buf))
	if Index_len == 1 {
		data.Index = READ_MSG_block_template_s(buf)
	}else{
		data.Index = nil
	}
	Orderplus_len := int(READ_int32(buf))
	for i := 0; i < Orderplus_len; i++ {
		data.Orderplus = append(data.Orderplus, READ_MSG_block_template_s(buf))
	}
	Order_len := int(READ_int8(buf))
	if Order_len == 1 {
		data.Order = READ_MSG_block_template_s(buf)
	}else{
		data.Order = nil
	}
	Loopplus_len := int(READ_int32(buf))
	for i := 0; i < Loopplus_len; i++ {
		data.Loopplus = append(data.Loopplus, READ_string(buf))
	}
	data.Loop = READ_string(buf)
	return
}

type MSG_block_template_s struct {
	Key string
	Order []*MSG_block_template_order
	Odd string
	Even string
}

var pool_MSG_block_template_s = sync.Pool{New: func() interface{} { return &MSG_block_template_s{} }}

func GET_MSG_block_template_s() *MSG_block_template_s {
	return pool_MSG_block_template_s.Get().(*MSG_block_template_s)
}

func (data *MSG_block_template_s) Put() {
	data.Key = ``
	for _,v := range data.Order {
		v.Put()
	}
	data.Order = data.Order[:0]
	data.Odd = ``
	data.Even = ``
	pool_MSG_block_template_s.Put(data)
}
func (data *MSG_block_template_s) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_block_template_s)
	WRITE_int32(cmd, buf)
	WRITE_MSG_block_template_s(data, buf)
}

func WRITE_MSG_block_template_s(data *MSG_block_template_s, buf *libraries.MsgBuffer) {
	WRITE_string(data.Key, buf)
	WRITE_int32(int32(len(data.Order)), buf)
	for _, v := range data.Order{
		WRITE_MSG_block_template_order(v, buf)
	}
	WRITE_string(data.Odd, buf)
	WRITE_string(data.Even, buf)
}

func READ_MSG_block_template_s(buf *libraries.MsgBuffer) (data *MSG_block_template_s) {
	data = pool_MSG_block_template_s.Get().(*MSG_block_template_s)
	data.Key = READ_string(buf)
	Order_len := int(READ_int32(buf))
	for i := 0; i < Order_len; i++ {
		data.Order = append(data.Order, READ_MSG_block_template_order(buf))
	}
	data.Odd = READ_string(buf)
	data.Even = READ_string(buf)
	return
}

type MSG_block_template_order struct {
	Key string
	Value string
}

var pool_MSG_block_template_order = sync.Pool{New: func() interface{} { return &MSG_block_template_order{} }}

func GET_MSG_block_template_order() *MSG_block_template_order {
	return pool_MSG_block_template_order.Get().(*MSG_block_template_order)
}

func (data *MSG_block_template_order) Put() {
	data.Key = ``
	data.Value = ``
	pool_MSG_block_template_order.Put(data)
}
func (data *MSG_block_template_order) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_block_template_order)
	WRITE_int32(cmd, buf)
	WRITE_MSG_block_template_order(data, buf)
}

func WRITE_MSG_block_template_order(data *MSG_block_template_order, buf *libraries.MsgBuffer) {
	WRITE_string(data.Key, buf)
	WRITE_string(data.Value, buf)
}

func READ_MSG_block_template_order(buf *libraries.MsgBuffer) (data *MSG_block_template_order) {
	data = pool_MSG_block_template_order.Get().(*MSG_block_template_order)
	data.Key = READ_string(buf)
	data.Value = READ_string(buf)
	return
}

type MSG_block_item struct {
	Position int16
	Itemid int32
	Fields *MSG_block_item_field
	Showstyle []*MSG_block_item_showstyle
	Picsize int16
	Picflag int8
}

var pool_MSG_block_item = sync.Pool{New: func() interface{} { return &MSG_block_item{} }}

func GET_MSG_block_item() *MSG_block_item {
	return pool_MSG_block_item.Get().(*MSG_block_item)
}

func (data *MSG_block_item) Put() {
	data.Position = 0
	data.Itemid = 0
	if data.Fields != nil {
		data.Fields.Put()
		data.Fields = nil
	}
	for _,v := range data.Showstyle {
		v.Put()
	}
	data.Showstyle = data.Showstyle[:0]
	data.Picsize = 0
	data.Picflag = 0
	pool_MSG_block_item.Put(data)
}
func (data *MSG_block_item) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_block_item)
	WRITE_int32(cmd, buf)
	WRITE_MSG_block_item(data, buf)
}

func WRITE_MSG_block_item(data *MSG_block_item, buf *libraries.MsgBuffer) {
	WRITE_int16(data.Position, buf)
	WRITE_int32(data.Itemid, buf)
	if data.Fields == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_block_item_field(data.Fields, buf)
	}
	WRITE_int32(int32(len(data.Showstyle)), buf)
	for _, v := range data.Showstyle{
		WRITE_MSG_block_item_showstyle(v, buf)
	}
	WRITE_int16(data.Picsize, buf)
	WRITE_int8(data.Picflag, buf)
}

func READ_MSG_block_item(buf *libraries.MsgBuffer) (data *MSG_block_item) {
	data = pool_MSG_block_item.Get().(*MSG_block_item)
	data.Position = READ_int16(buf)
	data.Itemid = READ_int32(buf)
	Fields_len := int(READ_int8(buf))
	if Fields_len == 1 {
		data.Fields = READ_MSG_block_item_field(buf)
	}else{
		data.Fields = nil
	}
	Showstyle_len := int(READ_int32(buf))
	for i := 0; i < Showstyle_len; i++ {
		data.Showstyle = append(data.Showstyle, READ_MSG_block_item_showstyle(buf))
	}
	data.Picsize = READ_int16(buf)
	data.Picflag = READ_int8(buf)
	return
}

type MSG_block_item_field struct {
	Fulltitle string
	Threads string
	Author string
	Authorid string
	Avatar string
	Avatar_middle string
	Avatar_big string
	Posts string
	Todayposts string
	Lastposter string
	Lastpost string
	Dateline string
	Replies string
	Forumurl string
	Forumname string
	Typename string
	Typeicon string
	Typeurl string
	Sortname string
	Sorturl string
	Views string
	Heats string
	Recommends string
	Hourviews string
	Todayviews string
	Weekviews string
	Monthviews string
}

var pool_MSG_block_item_field = sync.Pool{New: func() interface{} { return &MSG_block_item_field{} }}

func GET_MSG_block_item_field() *MSG_block_item_field {
	return pool_MSG_block_item_field.Get().(*MSG_block_item_field)
}

func (data *MSG_block_item_field) Put() {
	data.Fulltitle = ``
	data.Threads = ``
	data.Author = ``
	data.Authorid = ``
	data.Avatar = ``
	data.Avatar_middle = ``
	data.Avatar_big = ``
	data.Posts = ``
	data.Todayposts = ``
	data.Lastposter = ``
	data.Lastpost = ``
	data.Dateline = ``
	data.Replies = ``
	data.Forumurl = ``
	data.Forumname = ``
	data.Typename = ``
	data.Typeicon = ``
	data.Typeurl = ``
	data.Sortname = ``
	data.Sorturl = ``
	data.Views = ``
	data.Heats = ``
	data.Recommends = ``
	data.Hourviews = ``
	data.Todayviews = ``
	data.Weekviews = ``
	data.Monthviews = ``
	pool_MSG_block_item_field.Put(data)
}
func (data *MSG_block_item_field) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_block_item_field)
	WRITE_int32(cmd, buf)
	WRITE_MSG_block_item_field(data, buf)
}

func WRITE_MSG_block_item_field(data *MSG_block_item_field, buf *libraries.MsgBuffer) {
	WRITE_string(data.Fulltitle, buf)
	WRITE_string(data.Threads, buf)
	WRITE_string(data.Author, buf)
	WRITE_string(data.Authorid, buf)
	WRITE_string(data.Avatar, buf)
	WRITE_string(data.Avatar_middle, buf)
	WRITE_string(data.Avatar_big, buf)
	WRITE_string(data.Posts, buf)
	WRITE_string(data.Todayposts, buf)
	WRITE_string(data.Lastposter, buf)
	WRITE_string(data.Lastpost, buf)
	WRITE_string(data.Dateline, buf)
	WRITE_string(data.Replies, buf)
	WRITE_string(data.Forumurl, buf)
	WRITE_string(data.Forumname, buf)
	WRITE_string(data.Typename, buf)
	WRITE_string(data.Typeicon, buf)
	WRITE_string(data.Typeurl, buf)
	WRITE_string(data.Sortname, buf)
	WRITE_string(data.Sorturl, buf)
	WRITE_string(data.Views, buf)
	WRITE_string(data.Heats, buf)
	WRITE_string(data.Recommends, buf)
	WRITE_string(data.Hourviews, buf)
	WRITE_string(data.Todayviews, buf)
	WRITE_string(data.Weekviews, buf)
	WRITE_string(data.Monthviews, buf)
}

func READ_MSG_block_item_field(buf *libraries.MsgBuffer) (data *MSG_block_item_field) {
	data = pool_MSG_block_item_field.Get().(*MSG_block_item_field)
	data.Fulltitle = READ_string(buf)
	data.Threads = READ_string(buf)
	data.Author = READ_string(buf)
	data.Authorid = READ_string(buf)
	data.Avatar = READ_string(buf)
	data.Avatar_middle = READ_string(buf)
	data.Avatar_big = READ_string(buf)
	data.Posts = READ_string(buf)
	data.Todayposts = READ_string(buf)
	data.Lastposter = READ_string(buf)
	data.Lastpost = READ_string(buf)
	data.Dateline = READ_string(buf)
	data.Replies = READ_string(buf)
	data.Forumurl = READ_string(buf)
	data.Forumname = READ_string(buf)
	data.Typename = READ_string(buf)
	data.Typeicon = READ_string(buf)
	data.Typeurl = READ_string(buf)
	data.Sortname = READ_string(buf)
	data.Sorturl = READ_string(buf)
	data.Views = READ_string(buf)
	data.Heats = READ_string(buf)
	data.Recommends = READ_string(buf)
	data.Hourviews = READ_string(buf)
	data.Todayviews = READ_string(buf)
	data.Weekviews = READ_string(buf)
	data.Monthviews = READ_string(buf)
	return
}

type MSG_block_item_showstyle struct {
	Key string
	Info *MSG_block_item_showstyle_info
}

var pool_MSG_block_item_showstyle = sync.Pool{New: func() interface{} { return &MSG_block_item_showstyle{} }}

func GET_MSG_block_item_showstyle() *MSG_block_item_showstyle {
	return pool_MSG_block_item_showstyle.Get().(*MSG_block_item_showstyle)
}

func (data *MSG_block_item_showstyle) Put() {
	data.Key = ``
	if data.Info != nil {
		data.Info.Put()
		data.Info = nil
	}
	pool_MSG_block_item_showstyle.Put(data)
}
func (data *MSG_block_item_showstyle) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_block_item_showstyle)
	WRITE_int32(cmd, buf)
	WRITE_MSG_block_item_showstyle(data, buf)
}

func WRITE_MSG_block_item_showstyle(data *MSG_block_item_showstyle, buf *libraries.MsgBuffer) {
	WRITE_string(data.Key, buf)
	if data.Info == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_block_item_showstyle_info(data.Info, buf)
	}
}

func READ_MSG_block_item_showstyle(buf *libraries.MsgBuffer) (data *MSG_block_item_showstyle) {
	data = pool_MSG_block_item_showstyle.Get().(*MSG_block_item_showstyle)
	data.Key = READ_string(buf)
	Info_len := int(READ_int8(buf))
	if Info_len == 1 {
		data.Info = READ_MSG_block_item_showstyle_info(buf)
	}else{
		data.Info = nil
	}
	return
}

type MSG_block_item_showstyle_info struct {
	B int8
	I int8
	U int8
	C string
}

var pool_MSG_block_item_showstyle_info = sync.Pool{New: func() interface{} { return &MSG_block_item_showstyle_info{} }}

func GET_MSG_block_item_showstyle_info() *MSG_block_item_showstyle_info {
	return pool_MSG_block_item_showstyle_info.Get().(*MSG_block_item_showstyle_info)
}

func (data *MSG_block_item_showstyle_info) Put() {
	data.B = 0
	data.I = 0
	data.U = 0
	data.C = ``
	pool_MSG_block_item_showstyle_info.Put(data)
}
func (data *MSG_block_item_showstyle_info) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_block_item_showstyle_info)
	WRITE_int32(cmd, buf)
	WRITE_MSG_block_item_showstyle_info(data, buf)
}

func WRITE_MSG_block_item_showstyle_info(data *MSG_block_item_showstyle_info, buf *libraries.MsgBuffer) {
	WRITE_int8(data.B, buf)
	WRITE_int8(data.I, buf)
	WRITE_int8(data.U, buf)
	WRITE_string(data.C, buf)
}

func READ_MSG_block_item_showstyle_info(buf *libraries.MsgBuffer) (data *MSG_block_item_showstyle_info) {
	data = pool_MSG_block_item_showstyle_info.Get().(*MSG_block_item_showstyle_info)
	data.B = READ_int8(buf)
	data.I = READ_int8(buf)
	data.U = READ_int8(buf)
	data.C = READ_string(buf)
	return
}

type MSG_forum_index_cart struct {
	Endrows string
	Extra *MSG_forum_extra
	Fid int32
	Forumcolumns int8
	Forums []*MSG_forum
	Forumscount int32
	Moderators string
	Name string
}

var pool_MSG_forum_index_cart = sync.Pool{New: func() interface{} { return &MSG_forum_index_cart{} }}

func GET_MSG_forum_index_cart() *MSG_forum_index_cart {
	return pool_MSG_forum_index_cart.Get().(*MSG_forum_index_cart)
}

func (data *MSG_forum_index_cart) Put() {
	data.Endrows = ``
	if data.Extra != nil {
		data.Extra.Put()
		data.Extra = nil
	}
	data.Fid = 0
	data.Forumcolumns = 0
	for _,v := range data.Forums {
		v.Put()
	}
	data.Forums = data.Forums[:0]
	data.Forumscount = 0
	data.Moderators = ``
	data.Name = ``
	pool_MSG_forum_index_cart.Put(data)
}
func (data *MSG_forum_index_cart) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_forum_index_cart)
	WRITE_int32(cmd, buf)
	WRITE_MSG_forum_index_cart(data, buf)
}

func WRITE_MSG_forum_index_cart(data *MSG_forum_index_cart, buf *libraries.MsgBuffer) {
	WRITE_string(data.Endrows, buf)
	if data.Extra == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_forum_extra(data.Extra, buf)
	}
	WRITE_int32(data.Fid, buf)
	WRITE_int8(data.Forumcolumns, buf)
	WRITE_int32(int32(len(data.Forums)), buf)
	for _, v := range data.Forums{
		WRITE_MSG_forum(v, buf)
	}
	WRITE_int32(data.Forumscount, buf)
	WRITE_string(data.Moderators, buf)
	WRITE_string(data.Name, buf)
}

func READ_MSG_forum_index_cart(buf *libraries.MsgBuffer) (data *MSG_forum_index_cart) {
	data = pool_MSG_forum_index_cart.Get().(*MSG_forum_index_cart)
	data.Endrows = READ_string(buf)
	Extra_len := int(READ_int8(buf))
	if Extra_len == 1 {
		data.Extra = READ_MSG_forum_extra(buf)
	}else{
		data.Extra = nil
	}
	data.Fid = READ_int32(buf)
	data.Forumcolumns = READ_int8(buf)
	Forums_len := int(READ_int32(buf))
	for i := 0; i < Forums_len; i++ {
		data.Forums = append(data.Forums, READ_MSG_forum(buf))
	}
	data.Forumscount = READ_int32(buf)
	data.Moderators = READ_string(buf)
	data.Name = READ_string(buf)
	return
}

type MSG_forum_extra struct {
	Namecolor string
	Iconwidth int16
}

var pool_MSG_forum_extra = sync.Pool{New: func() interface{} { return &MSG_forum_extra{} }}

func GET_MSG_forum_extra() *MSG_forum_extra {
	return pool_MSG_forum_extra.Get().(*MSG_forum_extra)
}

func (data *MSG_forum_extra) Put() {
	data.Namecolor = ``
	data.Iconwidth = 0
	pool_MSG_forum_extra.Put(data)
}
func (data *MSG_forum_extra) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_forum_extra)
	WRITE_int32(cmd, buf)
	WRITE_MSG_forum_extra(data, buf)
}

func WRITE_MSG_forum_extra(data *MSG_forum_extra, buf *libraries.MsgBuffer) {
	WRITE_string(data.Namecolor, buf)
	WRITE_int16(data.Iconwidth, buf)
}

func READ_MSG_forum_extra(buf *libraries.MsgBuffer) (data *MSG_forum_extra) {
	data = pool_MSG_forum_extra.Get().(*MSG_forum_extra)
	data.Namecolor = READ_string(buf)
	data.Iconwidth = READ_int16(buf)
	return
}

type MSG_forum struct {
	Fup int32
	Extra *MSG_forum_extra
	Fid int32
	Icon string
	Lastpost *MSG_forum_lastpost
	Moderators string
	Name string
	Description string
	Orderid int32
	Permission int8
	Posts int32
	Subforums int8
	Threads int32
	Todayposts int32
	Simple int8
	Level_three []*MSG_forum_idname
}

var pool_MSG_forum = sync.Pool{New: func() interface{} { return &MSG_forum{} }}

func GET_MSG_forum() *MSG_forum {
	return pool_MSG_forum.Get().(*MSG_forum)
}

func (data *MSG_forum) Put() {
	data.Fup = 0
	if data.Extra != nil {
		data.Extra.Put()
		data.Extra = nil
	}
	data.Fid = 0
	data.Icon = ``
	if data.Lastpost != nil {
		data.Lastpost.Put()
		data.Lastpost = nil
	}
	data.Moderators = ``
	data.Name = ``
	data.Description = ``
	data.Orderid = 0
	data.Permission = 0
	data.Posts = 0
	data.Subforums = 0
	data.Threads = 0
	data.Todayposts = 0
	data.Simple = 0
	for _,v := range data.Level_three {
		v.Put()
	}
	data.Level_three = data.Level_three[:0]
	pool_MSG_forum.Put(data)
}
func (data *MSG_forum) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_forum)
	WRITE_int32(cmd, buf)
	WRITE_MSG_forum(data, buf)
}

func WRITE_MSG_forum(data *MSG_forum, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Fup, buf)
	if data.Extra == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_forum_extra(data.Extra, buf)
	}
	WRITE_int32(data.Fid, buf)
	WRITE_string(data.Icon, buf)
	if data.Lastpost == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_forum_lastpost(data.Lastpost, buf)
	}
	WRITE_string(data.Moderators, buf)
	WRITE_string(data.Name, buf)
	WRITE_string(data.Description, buf)
	WRITE_int32(data.Orderid, buf)
	WRITE_int8(data.Permission, buf)
	WRITE_int32(data.Posts, buf)
	WRITE_int8(data.Subforums, buf)
	WRITE_int32(data.Threads, buf)
	WRITE_int32(data.Todayposts, buf)
	WRITE_int8(data.Simple, buf)
	WRITE_int32(int32(len(data.Level_three)), buf)
	for _, v := range data.Level_three{
		WRITE_MSG_forum_idname(v, buf)
	}
}

func READ_MSG_forum(buf *libraries.MsgBuffer) (data *MSG_forum) {
	data = pool_MSG_forum.Get().(*MSG_forum)
	data.Fup = READ_int32(buf)
	Extra_len := int(READ_int8(buf))
	if Extra_len == 1 {
		data.Extra = READ_MSG_forum_extra(buf)
	}else{
		data.Extra = nil
	}
	data.Fid = READ_int32(buf)
	data.Icon = READ_string(buf)
	Lastpost_len := int(READ_int8(buf))
	if Lastpost_len == 1 {
		data.Lastpost = READ_MSG_forum_lastpost(buf)
	}else{
		data.Lastpost = nil
	}
	data.Moderators = READ_string(buf)
	data.Name = READ_string(buf)
	data.Description = READ_string(buf)
	data.Orderid = READ_int32(buf)
	data.Permission = READ_int8(buf)
	data.Posts = READ_int32(buf)
	data.Subforums = READ_int8(buf)
	data.Threads = READ_int32(buf)
	data.Todayposts = READ_int32(buf)
	data.Simple = READ_int8(buf)
	Level_three_len := int(READ_int32(buf))
	for i := 0; i < Level_three_len; i++ {
		data.Level_three = append(data.Level_three, READ_MSG_forum_idname(buf))
	}
	return
}

type MSG_forum_idname struct {
	Fid int32
	Name string
}

var pool_MSG_forum_idname = sync.Pool{New: func() interface{} { return &MSG_forum_idname{} }}

func GET_MSG_forum_idname() *MSG_forum_idname {
	return pool_MSG_forum_idname.Get().(*MSG_forum_idname)
}

func (data *MSG_forum_idname) Put() {
	data.Fid = 0
	data.Name = ``
	pool_MSG_forum_idname.Put(data)
}
func (data *MSG_forum_idname) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_forum_idname)
	WRITE_int32(cmd, buf)
	WRITE_MSG_forum_idname(data, buf)
}

func WRITE_MSG_forum_idname(data *MSG_forum_idname, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Fid, buf)
	WRITE_string(data.Name, buf)
}

func READ_MSG_forum_idname(buf *libraries.MsgBuffer) (data *MSG_forum_idname) {
	data = pool_MSG_forum_idname.Get().(*MSG_forum_idname)
	data.Fid = READ_int32(buf)
	data.Name = READ_string(buf)
	return
}

type MSG_forum_lastpost struct {
	Tid int32
	Subject string
	Dateline int32
	Author string
}

var pool_MSG_forum_lastpost = sync.Pool{New: func() interface{} { return &MSG_forum_lastpost{} }}

func GET_MSG_forum_lastpost() *MSG_forum_lastpost {
	return pool_MSG_forum_lastpost.Get().(*MSG_forum_lastpost)
}

func (data *MSG_forum_lastpost) Put() {
	data.Tid = 0
	data.Subject = ``
	data.Dateline = 0
	data.Author = ``
	pool_MSG_forum_lastpost.Put(data)
}
func (data *MSG_forum_lastpost) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_forum_lastpost)
	WRITE_int32(cmd, buf)
	WRITE_MSG_forum_lastpost(data, buf)
}

func WRITE_MSG_forum_lastpost(data *MSG_forum_lastpost, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Tid, buf)
	WRITE_string(data.Subject, buf)
	WRITE_int32(data.Dateline, buf)
	WRITE_string(data.Author, buf)
}

func READ_MSG_forum_lastpost(buf *libraries.MsgBuffer) (data *MSG_forum_lastpost) {
	data = pool_MSG_forum_lastpost.Get().(*MSG_forum_lastpost)
	data.Tid = READ_int32(buf)
	data.Subject = READ_string(buf)
	data.Dateline = READ_int32(buf)
	data.Author = READ_string(buf)
	return
}

type MSG_U2WS_forum struct {
	Fid int32
	Typeid string
	Dateline int32
	Orderby string
	Page int16
	Specialtype string
	Rewardtype int8
	Filter string
}

var pool_MSG_U2WS_forum = sync.Pool{New: func() interface{} { return &MSG_U2WS_forum{} }}

func GET_MSG_U2WS_forum() *MSG_U2WS_forum {
	return pool_MSG_U2WS_forum.Get().(*MSG_U2WS_forum)
}

func (data *MSG_U2WS_forum) Put() {
	data.Fid = 0
	data.Typeid = ``
	data.Dateline = 0
	data.Orderby = ``
	data.Page = 0
	data.Specialtype = ``
	data.Rewardtype = 0
	data.Filter = ``
	pool_MSG_U2WS_forum.Put(data)
}
func (data *MSG_U2WS_forum) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_forum)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_forum(data, buf)
}

func WRITE_MSG_U2WS_forum(data *MSG_U2WS_forum, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Fid, buf)
	WRITE_string(data.Typeid, buf)
	WRITE_int32(data.Dateline, buf)
	WRITE_string(data.Orderby, buf)
	WRITE_int16(data.Page, buf)
	WRITE_string(data.Specialtype, buf)
	WRITE_int8(data.Rewardtype, buf)
	WRITE_string(data.Filter, buf)
}

func READ_MSG_U2WS_forum(buf *libraries.MsgBuffer) (data *MSG_U2WS_forum) {
	data = pool_MSG_U2WS_forum.Get().(*MSG_U2WS_forum)
	data.Fid = READ_int32(buf)
	data.Typeid = READ_string(buf)
	data.Dateline = READ_int32(buf)
	data.Orderby = READ_string(buf)
	data.Page = READ_int16(buf)
	data.Specialtype = READ_string(buf)
	data.Rewardtype = READ_int8(buf)
	data.Filter = READ_string(buf)
	return
}

type MSG_WS2U_forum struct {
	Parent *MSG_forum_idname
	Allow int16
	Group_status int32
	Allowstickthread int8
	Fid int32
	Name string
	Modrecommend *MSG_forum_modrecommend
	Todayposts int32
	Threads int32
	Favtimes int32
	Threadmodcount int32
	Threadtypes *MSG_forum_threadtype
	Rules string
	Threadlist []*MSG_forum_thread
	Livethread *MSG_forum_thread
	Status int8
	Livemessage string
	Separatepos int32
	Threadscount int32
	Oldrank int16
	Rank int16
	Yesterdayposts int32
	Moderators string
	Lastpost int32
	Forum_history []*MSG_forum_idname
}

var pool_MSG_WS2U_forum = sync.Pool{New: func() interface{} { return &MSG_WS2U_forum{} }}

func GET_MSG_WS2U_forum() *MSG_WS2U_forum {
	return pool_MSG_WS2U_forum.Get().(*MSG_WS2U_forum)
}

func (data *MSG_WS2U_forum) Put() {
	if data.Parent != nil {
		data.Parent.Put()
		data.Parent = nil
	}
	data.Allow = 0
	data.Group_status = 0
	data.Allowstickthread = 0
	data.Fid = 0
	data.Name = ``
	if data.Modrecommend != nil {
		data.Modrecommend.Put()
		data.Modrecommend = nil
	}
	data.Todayposts = 0
	data.Threads = 0
	data.Favtimes = 0
	data.Threadmodcount = 0
	if data.Threadtypes != nil {
		data.Threadtypes.Put()
		data.Threadtypes = nil
	}
	data.Rules = ``
	for _,v := range data.Threadlist {
		v.Put()
	}
	data.Threadlist = data.Threadlist[:0]
	if data.Livethread != nil {
		data.Livethread.Put()
		data.Livethread = nil
	}
	data.Status = 0
	data.Livemessage = ``
	data.Separatepos = 0
	data.Threadscount = 0
	data.Oldrank = 0
	data.Rank = 0
	data.Yesterdayposts = 0
	data.Moderators = ``
	data.Lastpost = 0
	for _,v := range data.Forum_history {
		v.Put()
	}
	data.Forum_history = data.Forum_history[:0]
	pool_MSG_WS2U_forum.Put(data)
}
func (data *MSG_WS2U_forum) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_forum)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_forum(data, buf)
}

func WRITE_MSG_WS2U_forum(data *MSG_WS2U_forum, buf *libraries.MsgBuffer) {
	if data.Parent == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_forum_idname(data.Parent, buf)
	}
	WRITE_int16(data.Allow, buf)
	WRITE_int32(data.Group_status, buf)
	WRITE_int8(data.Allowstickthread, buf)
	WRITE_int32(data.Fid, buf)
	WRITE_string(data.Name, buf)
	if data.Modrecommend == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_forum_modrecommend(data.Modrecommend, buf)
	}
	WRITE_int32(data.Todayposts, buf)
	WRITE_int32(data.Threads, buf)
	WRITE_int32(data.Favtimes, buf)
	WRITE_int32(data.Threadmodcount, buf)
	if data.Threadtypes == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_forum_threadtype(data.Threadtypes, buf)
	}
	WRITE_string(data.Rules, buf)
	WRITE_int32(int32(len(data.Threadlist)), buf)
	for _, v := range data.Threadlist{
		WRITE_MSG_forum_thread(v, buf)
	}
	if data.Livethread == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_forum_thread(data.Livethread, buf)
	}
	WRITE_int8(data.Status, buf)
	WRITE_string(data.Livemessage, buf)
	WRITE_int32(data.Separatepos, buf)
	WRITE_int32(data.Threadscount, buf)
	WRITE_int16(data.Oldrank, buf)
	WRITE_int16(data.Rank, buf)
	WRITE_int32(data.Yesterdayposts, buf)
	WRITE_string(data.Moderators, buf)
	WRITE_int32(data.Lastpost, buf)
	WRITE_int32(int32(len(data.Forum_history)), buf)
	for _, v := range data.Forum_history{
		WRITE_MSG_forum_idname(v, buf)
	}
}

func READ_MSG_WS2U_forum(buf *libraries.MsgBuffer) (data *MSG_WS2U_forum) {
	data = pool_MSG_WS2U_forum.Get().(*MSG_WS2U_forum)
	Parent_len := int(READ_int8(buf))
	if Parent_len == 1 {
		data.Parent = READ_MSG_forum_idname(buf)
	}else{
		data.Parent = nil
	}
	data.Allow = READ_int16(buf)
	data.Group_status = READ_int32(buf)
	data.Allowstickthread = READ_int8(buf)
	data.Fid = READ_int32(buf)
	data.Name = READ_string(buf)
	Modrecommend_len := int(READ_int8(buf))
	if Modrecommend_len == 1 {
		data.Modrecommend = READ_MSG_forum_modrecommend(buf)
	}else{
		data.Modrecommend = nil
	}
	data.Todayposts = READ_int32(buf)
	data.Threads = READ_int32(buf)
	data.Favtimes = READ_int32(buf)
	data.Threadmodcount = READ_int32(buf)
	Threadtypes_len := int(READ_int8(buf))
	if Threadtypes_len == 1 {
		data.Threadtypes = READ_MSG_forum_threadtype(buf)
	}else{
		data.Threadtypes = nil
	}
	data.Rules = READ_string(buf)
	Threadlist_len := int(READ_int32(buf))
	for i := 0; i < Threadlist_len; i++ {
		data.Threadlist = append(data.Threadlist, READ_MSG_forum_thread(buf))
	}
	Livethread_len := int(READ_int8(buf))
	if Livethread_len == 1 {
		data.Livethread = READ_MSG_forum_thread(buf)
	}else{
		data.Livethread = nil
	}
	data.Status = READ_int8(buf)
	data.Livemessage = READ_string(buf)
	data.Separatepos = READ_int32(buf)
	data.Threadscount = READ_int32(buf)
	data.Oldrank = READ_int16(buf)
	data.Rank = READ_int16(buf)
	data.Yesterdayposts = READ_int32(buf)
	data.Moderators = READ_string(buf)
	data.Lastpost = READ_int32(buf)
	Forum_history_len := int(READ_int32(buf))
	for i := 0; i < Forum_history_len; i++ {
		data.Forum_history = append(data.Forum_history, READ_MSG_forum_idname(buf))
	}
	return
}

type MSG_forum_modrecommend struct {
	Sort int8
	Imagewidth int16
	Imageheight int16
	Imagenum int16
}

var pool_MSG_forum_modrecommend = sync.Pool{New: func() interface{} { return &MSG_forum_modrecommend{} }}

func GET_MSG_forum_modrecommend() *MSG_forum_modrecommend {
	return pool_MSG_forum_modrecommend.Get().(*MSG_forum_modrecommend)
}

func (data *MSG_forum_modrecommend) Put() {
	data.Sort = 0
	data.Imagewidth = 0
	data.Imageheight = 0
	data.Imagenum = 0
	pool_MSG_forum_modrecommend.Put(data)
}
func (data *MSG_forum_modrecommend) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_forum_modrecommend)
	WRITE_int32(cmd, buf)
	WRITE_MSG_forum_modrecommend(data, buf)
}

func WRITE_MSG_forum_modrecommend(data *MSG_forum_modrecommend, buf *libraries.MsgBuffer) {
	WRITE_int8(data.Sort, buf)
	WRITE_int16(data.Imagewidth, buf)
	WRITE_int16(data.Imageheight, buf)
	WRITE_int16(data.Imagenum, buf)
}

func READ_MSG_forum_modrecommend(buf *libraries.MsgBuffer) (data *MSG_forum_modrecommend) {
	data = pool_MSG_forum_modrecommend.Get().(*MSG_forum_modrecommend)
	data.Sort = READ_int8(buf)
	data.Imagewidth = READ_int16(buf)
	data.Imageheight = READ_int16(buf)
	data.Imagenum = READ_int16(buf)
	return
}

type MSG_U2WS_forum_modcp struct {
	Fid int32
}

var pool_MSG_U2WS_forum_modcp = sync.Pool{New: func() interface{} { return &MSG_U2WS_forum_modcp{} }}

func GET_MSG_U2WS_forum_modcp() *MSG_U2WS_forum_modcp {
	return pool_MSG_U2WS_forum_modcp.Get().(*MSG_U2WS_forum_modcp)
}

func (data *MSG_U2WS_forum_modcp) Put() {
	data.Fid = 0
	pool_MSG_U2WS_forum_modcp.Put(data)
}
func (data *MSG_U2WS_forum_modcp) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_forum_modcp)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_forum_modcp(data, buf)
}

func WRITE_MSG_U2WS_forum_modcp(data *MSG_U2WS_forum_modcp, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Fid, buf)
}

func READ_MSG_U2WS_forum_modcp(buf *libraries.MsgBuffer) (data *MSG_U2WS_forum_modcp) {
	data = pool_MSG_U2WS_forum_modcp.Get().(*MSG_U2WS_forum_modcp)
	data.Fid = READ_int32(buf)
	return
}

type MSG_U2WS_forum_recyclebin struct {
	Fid int32
}

var pool_MSG_U2WS_forum_recyclebin = sync.Pool{New: func() interface{} { return &MSG_U2WS_forum_recyclebin{} }}

func GET_MSG_U2WS_forum_recyclebin() *MSG_U2WS_forum_recyclebin {
	return pool_MSG_U2WS_forum_recyclebin.Get().(*MSG_U2WS_forum_recyclebin)
}

func (data *MSG_U2WS_forum_recyclebin) Put() {
	data.Fid = 0
	pool_MSG_U2WS_forum_recyclebin.Put(data)
}
func (data *MSG_U2WS_forum_recyclebin) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_forum_recyclebin)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_forum_recyclebin(data, buf)
}

func WRITE_MSG_U2WS_forum_recyclebin(data *MSG_U2WS_forum_recyclebin, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Fid, buf)
}

func READ_MSG_U2WS_forum_recyclebin(buf *libraries.MsgBuffer) (data *MSG_U2WS_forum_recyclebin) {
	data = pool_MSG_U2WS_forum_recyclebin.Get().(*MSG_U2WS_forum_recyclebin)
	data.Fid = READ_int32(buf)
	return
}

type MSG_forum_threadtype struct {
	Types []*MSG_forum_type
	Status int8
	Prefix int8
}

var pool_MSG_forum_threadtype = sync.Pool{New: func() interface{} { return &MSG_forum_threadtype{} }}

func GET_MSG_forum_threadtype() *MSG_forum_threadtype {
	return pool_MSG_forum_threadtype.Get().(*MSG_forum_threadtype)
}

func (data *MSG_forum_threadtype) Put() {
	for _,v := range data.Types {
		v.Put()
	}
	data.Types = data.Types[:0]
	data.Status = 0
	data.Prefix = 0
	pool_MSG_forum_threadtype.Put(data)
}
func (data *MSG_forum_threadtype) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_forum_threadtype)
	WRITE_int32(cmd, buf)
	WRITE_MSG_forum_threadtype(data, buf)
}

func WRITE_MSG_forum_threadtype(data *MSG_forum_threadtype, buf *libraries.MsgBuffer) {
	WRITE_int32(int32(len(data.Types)), buf)
	for _, v := range data.Types{
		WRITE_MSG_forum_type(v, buf)
	}
	WRITE_int8(data.Status, buf)
	WRITE_int8(data.Prefix, buf)
}

func READ_MSG_forum_threadtype(buf *libraries.MsgBuffer) (data *MSG_forum_threadtype) {
	data = pool_MSG_forum_threadtype.Get().(*MSG_forum_threadtype)
	Types_len := int(READ_int32(buf))
	for i := 0; i < Types_len; i++ {
		data.Types = append(data.Types, READ_MSG_forum_type(buf))
	}
	data.Status = READ_int8(buf)
	data.Prefix = READ_int8(buf)
	return
}

type MSG_forum_type struct {
	Id int16
	Name string
	Icon string
	Count int32
	Ismoderator int8
}

var pool_MSG_forum_type = sync.Pool{New: func() interface{} { return &MSG_forum_type{} }}

func GET_MSG_forum_type() *MSG_forum_type {
	return pool_MSG_forum_type.Get().(*MSG_forum_type)
}

func (data *MSG_forum_type) Put() {
	data.Id = 0
	data.Name = ``
	data.Icon = ``
	data.Count = 0
	data.Ismoderator = 0
	pool_MSG_forum_type.Put(data)
}
func (data *MSG_forum_type) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_forum_type)
	WRITE_int32(cmd, buf)
	WRITE_MSG_forum_type(data, buf)
}

func WRITE_MSG_forum_type(data *MSG_forum_type, buf *libraries.MsgBuffer) {
	WRITE_int16(data.Id, buf)
	WRITE_string(data.Name, buf)
	WRITE_string(data.Icon, buf)
	WRITE_int32(data.Count, buf)
	WRITE_int8(data.Ismoderator, buf)
}

func READ_MSG_forum_type(buf *libraries.MsgBuffer) (data *MSG_forum_type) {
	data = pool_MSG_forum_type.Get().(*MSG_forum_type)
	data.Id = READ_int16(buf)
	data.Name = READ_string(buf)
	data.Icon = READ_string(buf)
	data.Count = READ_int32(buf)
	data.Ismoderator = READ_int8(buf)
	return
}

type MSG_forum_thread struct {
	Displayorder int8
	Tid int32
	Allreplies int32
	Fid int32
	Closed int8
	Isgroup int8
	Typeid int16
	Icon int8
	Status int16
	Rushinfo *MSG_forum_threadrush
	Readperm int16
	Price int16
	Special int8
	Attachment int8
	Digest int8
	Heats int32
	Replycredit int16
	Dateline int32
	Lastpost int32
	Lastposter string
	Authorid int32
	Author string
	Subject string
	Folder string
	Views int32
	Recommend_add int32
	Recommend_sub int32
	Recommends int32
	Relay int32
	Replies int32
	Replycredit_rule *MSG_forum_replycredit
	Rewardfloor int8
	Groupname string
	Forumname string
	Groupcolor string
	Verify string
	Groupviews int32
	Highlight int8
	Stamp int8
	Cutmessage string
}

var pool_MSG_forum_thread = sync.Pool{New: func() interface{} { return &MSG_forum_thread{} }}

func GET_MSG_forum_thread() *MSG_forum_thread {
	return pool_MSG_forum_thread.Get().(*MSG_forum_thread)
}

func (data *MSG_forum_thread) Put() {
	data.Displayorder = 0
	data.Tid = 0
	data.Allreplies = 0
	data.Fid = 0
	data.Closed = 0
	data.Isgroup = 0
	data.Typeid = 0
	data.Icon = 0
	data.Status = 0
	if data.Rushinfo != nil {
		data.Rushinfo.Put()
		data.Rushinfo = nil
	}
	data.Readperm = 0
	data.Price = 0
	data.Special = 0
	data.Attachment = 0
	data.Digest = 0
	data.Heats = 0
	data.Replycredit = 0
	data.Dateline = 0
	data.Lastpost = 0
	data.Lastposter = ``
	data.Authorid = 0
	data.Author = ``
	data.Subject = ``
	data.Folder = ``
	data.Views = 0
	data.Recommend_add = 0
	data.Recommend_sub = 0
	data.Recommends = 0
	data.Relay = 0
	data.Replies = 0
	if data.Replycredit_rule != nil {
		data.Replycredit_rule.Put()
		data.Replycredit_rule = nil
	}
	data.Rewardfloor = 0
	data.Groupname = ``
	data.Forumname = ``
	data.Groupcolor = ``
	data.Verify = ``
	data.Groupviews = 0
	data.Highlight = 0
	data.Stamp = 0
	data.Cutmessage = ``
	pool_MSG_forum_thread.Put(data)
}
func (data *MSG_forum_thread) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_forum_thread)
	WRITE_int32(cmd, buf)
	WRITE_MSG_forum_thread(data, buf)
}

func WRITE_MSG_forum_thread(data *MSG_forum_thread, buf *libraries.MsgBuffer) {
	WRITE_int8(data.Displayorder, buf)
	WRITE_int32(data.Tid, buf)
	WRITE_int32(data.Allreplies, buf)
	WRITE_int32(data.Fid, buf)
	WRITE_int8(data.Closed, buf)
	WRITE_int8(data.Isgroup, buf)
	WRITE_int16(data.Typeid, buf)
	WRITE_int8(data.Icon, buf)
	WRITE_int16(data.Status, buf)
	if data.Rushinfo == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_forum_threadrush(data.Rushinfo, buf)
	}
	WRITE_int16(data.Readperm, buf)
	WRITE_int16(data.Price, buf)
	WRITE_int8(data.Special, buf)
	WRITE_int8(data.Attachment, buf)
	WRITE_int8(data.Digest, buf)
	WRITE_int32(data.Heats, buf)
	WRITE_int16(data.Replycredit, buf)
	WRITE_int32(data.Dateline, buf)
	WRITE_int32(data.Lastpost, buf)
	WRITE_string(data.Lastposter, buf)
	WRITE_int32(data.Authorid, buf)
	WRITE_string(data.Author, buf)
	WRITE_string(data.Subject, buf)
	WRITE_string(data.Folder, buf)
	WRITE_int32(data.Views, buf)
	WRITE_int32(data.Recommend_add, buf)
	WRITE_int32(data.Recommend_sub, buf)
	WRITE_int32(data.Recommends, buf)
	WRITE_int32(data.Relay, buf)
	WRITE_int32(data.Replies, buf)
	if data.Replycredit_rule == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_forum_replycredit(data.Replycredit_rule, buf)
	}
	WRITE_int8(data.Rewardfloor, buf)
	WRITE_string(data.Groupname, buf)
	WRITE_string(data.Forumname, buf)
	WRITE_string(data.Groupcolor, buf)
	WRITE_string(data.Verify, buf)
	WRITE_int32(data.Groupviews, buf)
	WRITE_int8(data.Highlight, buf)
	WRITE_int8(data.Stamp, buf)
	WRITE_string(data.Cutmessage, buf)
}

func READ_MSG_forum_thread(buf *libraries.MsgBuffer) (data *MSG_forum_thread) {
	data = pool_MSG_forum_thread.Get().(*MSG_forum_thread)
	data.Displayorder = READ_int8(buf)
	data.Tid = READ_int32(buf)
	data.Allreplies = READ_int32(buf)
	data.Fid = READ_int32(buf)
	data.Closed = READ_int8(buf)
	data.Isgroup = READ_int8(buf)
	data.Typeid = READ_int16(buf)
	data.Icon = READ_int8(buf)
	data.Status = READ_int16(buf)
	Rushinfo_len := int(READ_int8(buf))
	if Rushinfo_len == 1 {
		data.Rushinfo = READ_MSG_forum_threadrush(buf)
	}else{
		data.Rushinfo = nil
	}
	data.Readperm = READ_int16(buf)
	data.Price = READ_int16(buf)
	data.Special = READ_int8(buf)
	data.Attachment = READ_int8(buf)
	data.Digest = READ_int8(buf)
	data.Heats = READ_int32(buf)
	data.Replycredit = READ_int16(buf)
	data.Dateline = READ_int32(buf)
	data.Lastpost = READ_int32(buf)
	data.Lastposter = READ_string(buf)
	data.Authorid = READ_int32(buf)
	data.Author = READ_string(buf)
	data.Subject = READ_string(buf)
	data.Folder = READ_string(buf)
	data.Views = READ_int32(buf)
	data.Recommend_add = READ_int32(buf)
	data.Recommend_sub = READ_int32(buf)
	data.Recommends = READ_int32(buf)
	data.Relay = READ_int32(buf)
	data.Replies = READ_int32(buf)
	Replycredit_rule_len := int(READ_int8(buf))
	if Replycredit_rule_len == 1 {
		data.Replycredit_rule = READ_MSG_forum_replycredit(buf)
	}else{
		data.Replycredit_rule = nil
	}
	data.Rewardfloor = READ_int8(buf)
	data.Groupname = READ_string(buf)
	data.Forumname = READ_string(buf)
	data.Groupcolor = READ_string(buf)
	data.Verify = READ_string(buf)
	data.Groupviews = READ_int32(buf)
	data.Highlight = READ_int8(buf)
	data.Stamp = READ_int8(buf)
	data.Cutmessage = READ_string(buf)
	return
}

type MSG_forum_threadrush struct {
	Tid int32
	Starttimefrom int32
}

var pool_MSG_forum_threadrush = sync.Pool{New: func() interface{} { return &MSG_forum_threadrush{} }}

func GET_MSG_forum_threadrush() *MSG_forum_threadrush {
	return pool_MSG_forum_threadrush.Get().(*MSG_forum_threadrush)
}

func (data *MSG_forum_threadrush) Put() {
	data.Tid = 0
	data.Starttimefrom = 0
	pool_MSG_forum_threadrush.Put(data)
}
func (data *MSG_forum_threadrush) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_forum_threadrush)
	WRITE_int32(cmd, buf)
	WRITE_MSG_forum_threadrush(data, buf)
}

func WRITE_MSG_forum_threadrush(data *MSG_forum_threadrush, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Tid, buf)
	WRITE_int32(data.Starttimefrom, buf)
}

func READ_MSG_forum_threadrush(buf *libraries.MsgBuffer) (data *MSG_forum_threadrush) {
	data = pool_MSG_forum_threadrush.Get().(*MSG_forum_threadrush)
	data.Tid = READ_int32(buf)
	data.Starttimefrom = READ_int32(buf)
	return
}

type MSG_U2WS_forum_newthread struct {
	Fid int32
	Special int8
	Type int8
	Tid int32
	Position int32
}

var pool_MSG_U2WS_forum_newthread = sync.Pool{New: func() interface{} { return &MSG_U2WS_forum_newthread{} }}

func GET_MSG_U2WS_forum_newthread() *MSG_U2WS_forum_newthread {
	return pool_MSG_U2WS_forum_newthread.Get().(*MSG_U2WS_forum_newthread)
}

func (data *MSG_U2WS_forum_newthread) Put() {
	data.Fid = 0
	data.Special = 0
	data.Type = 0
	data.Tid = 0
	data.Position = 0
	pool_MSG_U2WS_forum_newthread.Put(data)
}
func (data *MSG_U2WS_forum_newthread) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_forum_newthread)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_forum_newthread(data, buf)
}

func WRITE_MSG_U2WS_forum_newthread(data *MSG_U2WS_forum_newthread, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Fid, buf)
	WRITE_int8(data.Special, buf)
	WRITE_int8(data.Type, buf)
	WRITE_int32(data.Tid, buf)
	WRITE_int32(data.Position, buf)
}

func READ_MSG_U2WS_forum_newthread(buf *libraries.MsgBuffer) (data *MSG_U2WS_forum_newthread) {
	data = pool_MSG_U2WS_forum_newthread.Get().(*MSG_U2WS_forum_newthread)
	data.Fid = READ_int32(buf)
	data.Special = READ_int8(buf)
	data.Type = READ_int8(buf)
	data.Tid = READ_int32(buf)
	data.Position = READ_int32(buf)
	return
}

type MSG_WS2U_forum_newthread struct {
	Parent *MSG_forum_idname
	Subject string
	Tid int32
	Fid int32
	Name string
	Savethreads []*MSG_forum_savethread
	Ismoderator int8
	Allow int32
	Group_status int32
	Threadtypes *MSG_forum_threadtype
	Typeid int16
	Displayorder int8
	Special int8
	Stand int8
	Message string
	Attachextensions string
	Allowat int8
	Extcreditstype int8
	Userextcredit int32
	Status int16
	Maxprice int16
	Replycredit int16
	Replycredit_rule *MSG_forum_replycredit
	Readperm int16
	Price int16
	Dateline int32
	Tag string
	Recent_use_tag []string
	Rush *MSG_forum_post_rush
	Anonymous int8
	Htmlon int8
	Replies int32
	Mygroups []*MSG_forum_group
	Attachs *MSG_forum_attach
	Imgattachs *MSG_forum_attach
	Maxattachsize int32
	Allowuploadnum int16
	Allowuploadsize int32
	Albumlist []*MSG_forum_album
	Poll *MSG_Poll_info
}

var pool_MSG_WS2U_forum_newthread = sync.Pool{New: func() interface{} { return &MSG_WS2U_forum_newthread{} }}

func GET_MSG_WS2U_forum_newthread() *MSG_WS2U_forum_newthread {
	return pool_MSG_WS2U_forum_newthread.Get().(*MSG_WS2U_forum_newthread)
}

func (data *MSG_WS2U_forum_newthread) Put() {
	if data.Parent != nil {
		data.Parent.Put()
		data.Parent = nil
	}
	data.Subject = ``
	data.Tid = 0
	data.Fid = 0
	data.Name = ``
	for _,v := range data.Savethreads {
		v.Put()
	}
	data.Savethreads = data.Savethreads[:0]
	data.Ismoderator = 0
	data.Allow = 0
	data.Group_status = 0
	if data.Threadtypes != nil {
		data.Threadtypes.Put()
		data.Threadtypes = nil
	}
	data.Typeid = 0
	data.Displayorder = 0
	data.Special = 0
	data.Stand = 0
	data.Message = ``
	data.Attachextensions = ``
	data.Allowat = 0
	data.Extcreditstype = 0
	data.Userextcredit = 0
	data.Status = 0
	data.Maxprice = 0
	data.Replycredit = 0
	if data.Replycredit_rule != nil {
		data.Replycredit_rule.Put()
		data.Replycredit_rule = nil
	}
	data.Readperm = 0
	data.Price = 0
	data.Dateline = 0
	data.Tag = ``
	data.Recent_use_tag = data.Recent_use_tag[:0]
	if data.Rush != nil {
		data.Rush.Put()
		data.Rush = nil
	}
	data.Anonymous = 0
	data.Htmlon = 0
	data.Replies = 0
	for _,v := range data.Mygroups {
		v.Put()
	}
	data.Mygroups = data.Mygroups[:0]
	if data.Attachs != nil {
		data.Attachs.Put()
		data.Attachs = nil
	}
	if data.Imgattachs != nil {
		data.Imgattachs.Put()
		data.Imgattachs = nil
	}
	data.Maxattachsize = 0
	data.Allowuploadnum = 0
	data.Allowuploadsize = 0
	for _,v := range data.Albumlist {
		v.Put()
	}
	data.Albumlist = data.Albumlist[:0]
	if data.Poll != nil {
		data.Poll.Put()
		data.Poll = nil
	}
	pool_MSG_WS2U_forum_newthread.Put(data)
}
func (data *MSG_WS2U_forum_newthread) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_forum_newthread)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_forum_newthread(data, buf)
}

func WRITE_MSG_WS2U_forum_newthread(data *MSG_WS2U_forum_newthread, buf *libraries.MsgBuffer) {
	if data.Parent == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_forum_idname(data.Parent, buf)
	}
	WRITE_string(data.Subject, buf)
	WRITE_int32(data.Tid, buf)
	WRITE_int32(data.Fid, buf)
	WRITE_string(data.Name, buf)
	WRITE_int32(int32(len(data.Savethreads)), buf)
	for _, v := range data.Savethreads{
		WRITE_MSG_forum_savethread(v, buf)
	}
	WRITE_int8(data.Ismoderator, buf)
	WRITE_int32(data.Allow, buf)
	WRITE_int32(data.Group_status, buf)
	if data.Threadtypes == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_forum_threadtype(data.Threadtypes, buf)
	}
	WRITE_int16(data.Typeid, buf)
	WRITE_int8(data.Displayorder, buf)
	WRITE_int8(data.Special, buf)
	WRITE_int8(data.Stand, buf)
	WRITE_string(data.Message, buf)
	WRITE_string(data.Attachextensions, buf)
	WRITE_int8(data.Allowat, buf)
	WRITE_int8(data.Extcreditstype, buf)
	WRITE_int32(data.Userextcredit, buf)
	WRITE_int16(data.Status, buf)
	WRITE_int16(data.Maxprice, buf)
	WRITE_int16(data.Replycredit, buf)
	if data.Replycredit_rule == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_forum_replycredit(data.Replycredit_rule, buf)
	}
	WRITE_int16(data.Readperm, buf)
	WRITE_int16(data.Price, buf)
	WRITE_int32(data.Dateline, buf)
	WRITE_string(data.Tag, buf)
	WRITE_int32(int32(len(data.Recent_use_tag)), buf)
	for _, v := range data.Recent_use_tag{
		WRITE_string(v, buf)
	}
	if data.Rush == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_forum_post_rush(data.Rush, buf)
	}
	WRITE_int8(data.Anonymous, buf)
	WRITE_int8(data.Htmlon, buf)
	WRITE_int32(data.Replies, buf)
	WRITE_int32(int32(len(data.Mygroups)), buf)
	for _, v := range data.Mygroups{
		WRITE_MSG_forum_group(v, buf)
	}
	if data.Attachs == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_forum_attach(data.Attachs, buf)
	}
	if data.Imgattachs == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_forum_attach(data.Imgattachs, buf)
	}
	WRITE_int32(data.Maxattachsize, buf)
	WRITE_int16(data.Allowuploadnum, buf)
	WRITE_int32(data.Allowuploadsize, buf)
	WRITE_int32(int32(len(data.Albumlist)), buf)
	for _, v := range data.Albumlist{
		WRITE_MSG_forum_album(v, buf)
	}
	if data.Poll == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_Poll_info(data.Poll, buf)
	}
}

func READ_MSG_WS2U_forum_newthread(buf *libraries.MsgBuffer) (data *MSG_WS2U_forum_newthread) {
	data = pool_MSG_WS2U_forum_newthread.Get().(*MSG_WS2U_forum_newthread)
	Parent_len := int(READ_int8(buf))
	if Parent_len == 1 {
		data.Parent = READ_MSG_forum_idname(buf)
	}else{
		data.Parent = nil
	}
	data.Subject = READ_string(buf)
	data.Tid = READ_int32(buf)
	data.Fid = READ_int32(buf)
	data.Name = READ_string(buf)
	Savethreads_len := int(READ_int32(buf))
	for i := 0; i < Savethreads_len; i++ {
		data.Savethreads = append(data.Savethreads, READ_MSG_forum_savethread(buf))
	}
	data.Ismoderator = READ_int8(buf)
	data.Allow = READ_int32(buf)
	data.Group_status = READ_int32(buf)
	Threadtypes_len := int(READ_int8(buf))
	if Threadtypes_len == 1 {
		data.Threadtypes = READ_MSG_forum_threadtype(buf)
	}else{
		data.Threadtypes = nil
	}
	data.Typeid = READ_int16(buf)
	data.Displayorder = READ_int8(buf)
	data.Special = READ_int8(buf)
	data.Stand = READ_int8(buf)
	data.Message = READ_string(buf)
	data.Attachextensions = READ_string(buf)
	data.Allowat = READ_int8(buf)
	data.Extcreditstype = READ_int8(buf)
	data.Userextcredit = READ_int32(buf)
	data.Status = READ_int16(buf)
	data.Maxprice = READ_int16(buf)
	data.Replycredit = READ_int16(buf)
	Replycredit_rule_len := int(READ_int8(buf))
	if Replycredit_rule_len == 1 {
		data.Replycredit_rule = READ_MSG_forum_replycredit(buf)
	}else{
		data.Replycredit_rule = nil
	}
	data.Readperm = READ_int16(buf)
	data.Price = READ_int16(buf)
	data.Dateline = READ_int32(buf)
	data.Tag = READ_string(buf)
	Recent_use_tag_len := int(READ_int32(buf))
	for i := 0; i < Recent_use_tag_len; i++ {
		data.Recent_use_tag = append(data.Recent_use_tag, READ_string(buf))
	}
	Rush_len := int(READ_int8(buf))
	if Rush_len == 1 {
		data.Rush = READ_MSG_forum_post_rush(buf)
	}else{
		data.Rush = nil
	}
	data.Anonymous = READ_int8(buf)
	data.Htmlon = READ_int8(buf)
	data.Replies = READ_int32(buf)
	Mygroups_len := int(READ_int32(buf))
	for i := 0; i < Mygroups_len; i++ {
		data.Mygroups = append(data.Mygroups, READ_MSG_forum_group(buf))
	}
	Attachs_len := int(READ_int8(buf))
	if Attachs_len == 1 {
		data.Attachs = READ_MSG_forum_attach(buf)
	}else{
		data.Attachs = nil
	}
	Imgattachs_len := int(READ_int8(buf))
	if Imgattachs_len == 1 {
		data.Imgattachs = READ_MSG_forum_attach(buf)
	}else{
		data.Imgattachs = nil
	}
	data.Maxattachsize = READ_int32(buf)
	data.Allowuploadnum = READ_int16(buf)
	data.Allowuploadsize = READ_int32(buf)
	Albumlist_len := int(READ_int32(buf))
	for i := 0; i < Albumlist_len; i++ {
		data.Albumlist = append(data.Albumlist, READ_MSG_forum_album(buf))
	}
	Poll_len := int(READ_int8(buf))
	if Poll_len == 1 {
		data.Poll = READ_MSG_Poll_info(buf)
	}else{
		data.Poll = nil
	}
	return
}

type MSG_forum_savethread struct {
	Tid int32
	Position int32
	Fid int32
	Subject string
	Dateline int32
}

var pool_MSG_forum_savethread = sync.Pool{New: func() interface{} { return &MSG_forum_savethread{} }}

func GET_MSG_forum_savethread() *MSG_forum_savethread {
	return pool_MSG_forum_savethread.Get().(*MSG_forum_savethread)
}

func (data *MSG_forum_savethread) Put() {
	data.Tid = 0
	data.Position = 0
	data.Fid = 0
	data.Subject = ``
	data.Dateline = 0
	pool_MSG_forum_savethread.Put(data)
}
func (data *MSG_forum_savethread) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_forum_savethread)
	WRITE_int32(cmd, buf)
	WRITE_MSG_forum_savethread(data, buf)
}

func WRITE_MSG_forum_savethread(data *MSG_forum_savethread, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Tid, buf)
	WRITE_int32(data.Position, buf)
	WRITE_int32(data.Fid, buf)
	WRITE_string(data.Subject, buf)
	WRITE_int32(data.Dateline, buf)
}

func READ_MSG_forum_savethread(buf *libraries.MsgBuffer) (data *MSG_forum_savethread) {
	data = pool_MSG_forum_savethread.Get().(*MSG_forum_savethread)
	data.Tid = READ_int32(buf)
	data.Position = READ_int32(buf)
	data.Fid = READ_int32(buf)
	data.Subject = READ_string(buf)
	data.Dateline = READ_int32(buf)
	return
}

type MSG_forum_replycredit struct {
	Tid int32
	Extcredits int32
	Extcreditstype int8
	Times int16
	Membertimes int16
	Random int32
}

var pool_MSG_forum_replycredit = sync.Pool{New: func() interface{} { return &MSG_forum_replycredit{} }}

func GET_MSG_forum_replycredit() *MSG_forum_replycredit {
	return pool_MSG_forum_replycredit.Get().(*MSG_forum_replycredit)
}

func (data *MSG_forum_replycredit) Put() {
	data.Tid = 0
	data.Extcredits = 0
	data.Extcreditstype = 0
	data.Times = 0
	data.Membertimes = 0
	data.Random = 0
	pool_MSG_forum_replycredit.Put(data)
}
func (data *MSG_forum_replycredit) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_forum_replycredit)
	WRITE_int32(cmd, buf)
	WRITE_MSG_forum_replycredit(data, buf)
}

func WRITE_MSG_forum_replycredit(data *MSG_forum_replycredit, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Tid, buf)
	WRITE_int32(data.Extcredits, buf)
	WRITE_int8(data.Extcreditstype, buf)
	WRITE_int16(data.Times, buf)
	WRITE_int16(data.Membertimes, buf)
	WRITE_int32(data.Random, buf)
}

func READ_MSG_forum_replycredit(buf *libraries.MsgBuffer) (data *MSG_forum_replycredit) {
	data = pool_MSG_forum_replycredit.Get().(*MSG_forum_replycredit)
	data.Tid = READ_int32(buf)
	data.Extcredits = READ_int32(buf)
	data.Extcreditstype = READ_int8(buf)
	data.Times = READ_int16(buf)
	data.Membertimes = READ_int16(buf)
	data.Random = READ_int32(buf)
	return
}

type MSG_forum_post_rush struct {
	Tid int32
	Starttimefrom int32
	Starttimeto int32
	Rewardfloor string
	Replylimit int16
	Stopfloor int32
	Creditlimit int32
}

var pool_MSG_forum_post_rush = sync.Pool{New: func() interface{} { return &MSG_forum_post_rush{} }}

func GET_MSG_forum_post_rush() *MSG_forum_post_rush {
	return pool_MSG_forum_post_rush.Get().(*MSG_forum_post_rush)
}

func (data *MSG_forum_post_rush) Put() {
	data.Tid = 0
	data.Starttimefrom = 0
	data.Starttimeto = 0
	data.Rewardfloor = ``
	data.Replylimit = 0
	data.Stopfloor = 0
	data.Creditlimit = 0
	pool_MSG_forum_post_rush.Put(data)
}
func (data *MSG_forum_post_rush) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_forum_post_rush)
	WRITE_int32(cmd, buf)
	WRITE_MSG_forum_post_rush(data, buf)
}

func WRITE_MSG_forum_post_rush(data *MSG_forum_post_rush, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Tid, buf)
	WRITE_int32(data.Starttimefrom, buf)
	WRITE_int32(data.Starttimeto, buf)
	WRITE_string(data.Rewardfloor, buf)
	WRITE_int16(data.Replylimit, buf)
	WRITE_int32(data.Stopfloor, buf)
	WRITE_int32(data.Creditlimit, buf)
}

func READ_MSG_forum_post_rush(buf *libraries.MsgBuffer) (data *MSG_forum_post_rush) {
	data = pool_MSG_forum_post_rush.Get().(*MSG_forum_post_rush)
	data.Tid = READ_int32(buf)
	data.Starttimefrom = READ_int32(buf)
	data.Starttimeto = READ_int32(buf)
	data.Rewardfloor = READ_string(buf)
	data.Replylimit = READ_int16(buf)
	data.Stopfloor = READ_int32(buf)
	data.Creditlimit = READ_int32(buf)
	return
}

type MSG_forum_group struct {
	Fid int32
	Name string
}

var pool_MSG_forum_group = sync.Pool{New: func() interface{} { return &MSG_forum_group{} }}

func GET_MSG_forum_group() *MSG_forum_group {
	return pool_MSG_forum_group.Get().(*MSG_forum_group)
}

func (data *MSG_forum_group) Put() {
	data.Fid = 0
	data.Name = ``
	pool_MSG_forum_group.Put(data)
}
func (data *MSG_forum_group) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_forum_group)
	WRITE_int32(cmd, buf)
	WRITE_MSG_forum_group(data, buf)
}

func WRITE_MSG_forum_group(data *MSG_forum_group, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Fid, buf)
	WRITE_string(data.Name, buf)
}

func READ_MSG_forum_group(buf *libraries.MsgBuffer) (data *MSG_forum_group) {
	data = pool_MSG_forum_group.Get().(*MSG_forum_group)
	data.Fid = READ_int32(buf)
	data.Name = READ_string(buf)
	return
}

type MSG_forum_attach struct {
	Unused []*MSG_forum_imgattach
	Used []*MSG_forum_imgattach
}

var pool_MSG_forum_attach = sync.Pool{New: func() interface{} { return &MSG_forum_attach{} }}

func GET_MSG_forum_attach() *MSG_forum_attach {
	return pool_MSG_forum_attach.Get().(*MSG_forum_attach)
}

func (data *MSG_forum_attach) Put() {
	for _,v := range data.Unused {
		v.Put()
	}
	data.Unused = data.Unused[:0]
	for _,v := range data.Used {
		v.Put()
	}
	data.Used = data.Used[:0]
	pool_MSG_forum_attach.Put(data)
}
func (data *MSG_forum_attach) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_forum_attach)
	WRITE_int32(cmd, buf)
	WRITE_MSG_forum_attach(data, buf)
}

func WRITE_MSG_forum_attach(data *MSG_forum_attach, buf *libraries.MsgBuffer) {
	WRITE_int32(int32(len(data.Unused)), buf)
	for _, v := range data.Unused{
		WRITE_MSG_forum_imgattach(v, buf)
	}
	WRITE_int32(int32(len(data.Used)), buf)
	for _, v := range data.Used{
		WRITE_MSG_forum_imgattach(v, buf)
	}
}

func READ_MSG_forum_attach(buf *libraries.MsgBuffer) (data *MSG_forum_attach) {
	data = pool_MSG_forum_attach.Get().(*MSG_forum_attach)
	Unused_len := int(READ_int32(buf))
	for i := 0; i < Unused_len; i++ {
		data.Unused = append(data.Unused, READ_MSG_forum_imgattach(buf))
	}
	Used_len := int(READ_int32(buf))
	for i := 0; i < Used_len; i++ {
		data.Used = append(data.Used, READ_MSG_forum_imgattach(buf))
	}
	return
}

type MSG_forum_imgattach struct {
	Aid int64
	Filenametitle string
	Src string
	Dateline int32
	Filename string
}

var pool_MSG_forum_imgattach = sync.Pool{New: func() interface{} { return &MSG_forum_imgattach{} }}

func GET_MSG_forum_imgattach() *MSG_forum_imgattach {
	return pool_MSG_forum_imgattach.Get().(*MSG_forum_imgattach)
}

func (data *MSG_forum_imgattach) Put() {
	data.Aid = 0
	data.Filenametitle = ``
	data.Src = ``
	data.Dateline = 0
	data.Filename = ``
	pool_MSG_forum_imgattach.Put(data)
}
func (data *MSG_forum_imgattach) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_forum_imgattach)
	WRITE_int32(cmd, buf)
	WRITE_MSG_forum_imgattach(data, buf)
}

func WRITE_MSG_forum_imgattach(data *MSG_forum_imgattach, buf *libraries.MsgBuffer) {
	WRITE_int64(data.Aid, buf)
	WRITE_string(data.Filenametitle, buf)
	WRITE_string(data.Src, buf)
	WRITE_int32(data.Dateline, buf)
	WRITE_string(data.Filename, buf)
}

func READ_MSG_forum_imgattach(buf *libraries.MsgBuffer) (data *MSG_forum_imgattach) {
	data = pool_MSG_forum_imgattach.Get().(*MSG_forum_imgattach)
	data.Aid = READ_int64(buf)
	data.Filenametitle = READ_string(buf)
	data.Src = READ_string(buf)
	data.Dateline = READ_int32(buf)
	data.Filename = READ_string(buf)
	return
}

type MSG_forum_album struct {
	Albumid int32
	Albumname string
}

var pool_MSG_forum_album = sync.Pool{New: func() interface{} { return &MSG_forum_album{} }}

func GET_MSG_forum_album() *MSG_forum_album {
	return pool_MSG_forum_album.Get().(*MSG_forum_album)
}

func (data *MSG_forum_album) Put() {
	data.Albumid = 0
	data.Albumname = ``
	pool_MSG_forum_album.Put(data)
}
func (data *MSG_forum_album) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_forum_album)
	WRITE_int32(cmd, buf)
	WRITE_MSG_forum_album(data, buf)
}

func WRITE_MSG_forum_album(data *MSG_forum_album, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Albumid, buf)
	WRITE_string(data.Albumname, buf)
}

func READ_MSG_forum_album(buf *libraries.MsgBuffer) (data *MSG_forum_album) {
	data = pool_MSG_forum_album.Get().(*MSG_forum_album)
	data.Albumid = READ_int32(buf)
	data.Albumname = READ_string(buf)
	return
}

type MSG_U2WS_Getlogin struct {
}

var pool_MSG_U2WS_Getlogin = sync.Pool{New: func() interface{} { return &MSG_U2WS_Getlogin{} }}

func GET_MSG_U2WS_Getlogin() *MSG_U2WS_Getlogin {
	return pool_MSG_U2WS_Getlogin.Get().(*MSG_U2WS_Getlogin)
}

func (data *MSG_U2WS_Getlogin) Put() {
	pool_MSG_U2WS_Getlogin.Put(data)
}
func (data *MSG_U2WS_Getlogin) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_Getlogin)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_Getlogin(data, buf)
}

func WRITE_MSG_U2WS_Getlogin(data *MSG_U2WS_Getlogin, buf *libraries.MsgBuffer) {
}

func READ_MSG_U2WS_Getlogin(buf *libraries.MsgBuffer) (data *MSG_U2WS_Getlogin) {
	data = pool_MSG_U2WS_Getlogin.Get().(*MSG_U2WS_Getlogin)
	return
}

type MSG_WS2U_Getlogin struct {
	Islogin int8
	Img []byte
}

var pool_MSG_WS2U_Getlogin = sync.Pool{New: func() interface{} { return &MSG_WS2U_Getlogin{} }}

func GET_MSG_WS2U_Getlogin() *MSG_WS2U_Getlogin {
	return pool_MSG_WS2U_Getlogin.Get().(*MSG_WS2U_Getlogin)
}

func (data *MSG_WS2U_Getlogin) Put() {
	data.Islogin = 0
	data.Img = data.Img[:0]
	pool_MSG_WS2U_Getlogin.Put(data)
}
func (data *MSG_WS2U_Getlogin) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_Getlogin)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_Getlogin(data, buf)
}

func WRITE_MSG_WS2U_Getlogin(data *MSG_WS2U_Getlogin, buf *libraries.MsgBuffer) {
	WRITE_int8(data.Islogin, buf)
	WRITE_int32(int32(len(data.Img)), buf)
	buf.Write(data.Img)
}

func READ_MSG_WS2U_Getlogin(buf *libraries.MsgBuffer) (data *MSG_WS2U_Getlogin) {
	data = pool_MSG_WS2U_Getlogin.Get().(*MSG_WS2U_Getlogin)
	data.Islogin = READ_int8(buf)
	Img_len := int(READ_int32(buf))
	data.Img = make([]byte, Img_len)
	copy(data.Img,buf.Next(Img_len))
	return
}

type MSG_U2WS_Forum_newthread_submit struct {
	Fid int32
	Tid int32
	Position int32
	Type int8
	Typeid int16
	Special int8
	Subject string
	Message string
	Seccode string
	Other int16
	Readperm int16
	Tags string
	Aids []int64
	Specialext []byte
}

var pool_MSG_U2WS_Forum_newthread_submit = sync.Pool{New: func() interface{} { return &MSG_U2WS_Forum_newthread_submit{} }}

func GET_MSG_U2WS_Forum_newthread_submit() *MSG_U2WS_Forum_newthread_submit {
	return pool_MSG_U2WS_Forum_newthread_submit.Get().(*MSG_U2WS_Forum_newthread_submit)
}

func (data *MSG_U2WS_Forum_newthread_submit) Put() {
	data.Fid = 0
	data.Tid = 0
	data.Position = 0
	data.Type = 0
	data.Typeid = 0
	data.Special = 0
	data.Subject = ``
	data.Message = ``
	data.Seccode = ``
	data.Other = 0
	data.Readperm = 0
	data.Tags = ``
	data.Aids = data.Aids[:0]
	data.Specialext = data.Specialext[:0]
	pool_MSG_U2WS_Forum_newthread_submit.Put(data)
}
func (data *MSG_U2WS_Forum_newthread_submit) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_Forum_newthread_submit)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_Forum_newthread_submit(data, buf)
}

func WRITE_MSG_U2WS_Forum_newthread_submit(data *MSG_U2WS_Forum_newthread_submit, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Fid, buf)
	WRITE_int32(data.Tid, buf)
	WRITE_int32(data.Position, buf)
	WRITE_int8(data.Type, buf)
	WRITE_int16(data.Typeid, buf)
	WRITE_int8(data.Special, buf)
	WRITE_string(data.Subject, buf)
	WRITE_string(data.Message, buf)
	WRITE_string(data.Seccode, buf)
	WRITE_int16(data.Other, buf)
	WRITE_int16(data.Readperm, buf)
	WRITE_string(data.Tags, buf)
	WRITE_int32(int32(len(data.Aids)), buf)
	for _, v := range data.Aids{
		WRITE_int64(v, buf)
	}
	WRITE_int32(int32(len(data.Specialext)), buf)
	buf.Write(data.Specialext)
}

func READ_MSG_U2WS_Forum_newthread_submit(buf *libraries.MsgBuffer) (data *MSG_U2WS_Forum_newthread_submit) {
	data = pool_MSG_U2WS_Forum_newthread_submit.Get().(*MSG_U2WS_Forum_newthread_submit)
	data.Fid = READ_int32(buf)
	data.Tid = READ_int32(buf)
	data.Position = READ_int32(buf)
	data.Type = READ_int8(buf)
	data.Typeid = READ_int16(buf)
	data.Special = READ_int8(buf)
	data.Subject = READ_string(buf)
	data.Message = READ_string(buf)
	data.Seccode = READ_string(buf)
	data.Other = READ_int16(buf)
	data.Readperm = READ_int16(buf)
	data.Tags = READ_string(buf)
	Aids_len := int(READ_int32(buf))
	for i := 0; i < Aids_len; i++ {
		data.Aids = append(data.Aids, READ_int64(buf))
	}
	Specialext_len := int(READ_int32(buf))
	data.Specialext = make([]byte, Specialext_len)
	copy(data.Specialext,buf.Next(Specialext_len))
	return
}

type MSG_WS2U_Forum_newthread_submit struct {
	Result int16
	Tid int32
	Extcredits []*MSG_add_extcredits
}

var pool_MSG_WS2U_Forum_newthread_submit = sync.Pool{New: func() interface{} { return &MSG_WS2U_Forum_newthread_submit{} }}

func GET_MSG_WS2U_Forum_newthread_submit() *MSG_WS2U_Forum_newthread_submit {
	return pool_MSG_WS2U_Forum_newthread_submit.Get().(*MSG_WS2U_Forum_newthread_submit)
}

func (data *MSG_WS2U_Forum_newthread_submit) Put() {
	data.Result = 0
	data.Tid = 0
	for _,v := range data.Extcredits {
		v.Put()
	}
	data.Extcredits = data.Extcredits[:0]
	pool_MSG_WS2U_Forum_newthread_submit.Put(data)
}
func (data *MSG_WS2U_Forum_newthread_submit) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_Forum_newthread_submit)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_Forum_newthread_submit(data, buf)
}

func WRITE_MSG_WS2U_Forum_newthread_submit(data *MSG_WS2U_Forum_newthread_submit, buf *libraries.MsgBuffer) {
	WRITE_int16(data.Result, buf)
	WRITE_int32(data.Tid, buf)
	WRITE_int32(int32(len(data.Extcredits)), buf)
	for _, v := range data.Extcredits{
		WRITE_MSG_add_extcredits(v, buf)
	}
}

func READ_MSG_WS2U_Forum_newthread_submit(buf *libraries.MsgBuffer) (data *MSG_WS2U_Forum_newthread_submit) {
	data = pool_MSG_WS2U_Forum_newthread_submit.Get().(*MSG_WS2U_Forum_newthread_submit)
	data.Result = READ_int16(buf)
	data.Tid = READ_int32(buf)
	Extcredits_len := int(READ_int32(buf))
	for i := 0; i < Extcredits_len; i++ {
		data.Extcredits = append(data.Extcredits, READ_MSG_add_extcredits(buf))
	}
	return
}

type MSG_add_extcredits struct {
	Id int8
	Value int32
}

var pool_MSG_add_extcredits = sync.Pool{New: func() interface{} { return &MSG_add_extcredits{} }}

func GET_MSG_add_extcredits() *MSG_add_extcredits {
	return pool_MSG_add_extcredits.Get().(*MSG_add_extcredits)
}

func (data *MSG_add_extcredits) Put() {
	data.Id = 0
	data.Value = 0
	pool_MSG_add_extcredits.Put(data)
}
func (data *MSG_add_extcredits) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_add_extcredits)
	WRITE_int32(cmd, buf)
	WRITE_MSG_add_extcredits(data, buf)
}

func WRITE_MSG_add_extcredits(data *MSG_add_extcredits, buf *libraries.MsgBuffer) {
	WRITE_int8(data.Id, buf)
	WRITE_int32(data.Value, buf)
}

func READ_MSG_add_extcredits(buf *libraries.MsgBuffer) (data *MSG_add_extcredits) {
	data = pool_MSG_add_extcredits.Get().(*MSG_add_extcredits)
	data.Id = READ_int8(buf)
	data.Value = READ_int32(buf)
	return
}

type MSG_U2WS_forum_viewthread struct {
	Tid int32
	Page int16
	Ordertype int8
	Stand int8
	Authorid int32
	Position int32
	Fromuid int32
}

var pool_MSG_U2WS_forum_viewthread = sync.Pool{New: func() interface{} { return &MSG_U2WS_forum_viewthread{} }}

func GET_MSG_U2WS_forum_viewthread() *MSG_U2WS_forum_viewthread {
	return pool_MSG_U2WS_forum_viewthread.Get().(*MSG_U2WS_forum_viewthread)
}

func (data *MSG_U2WS_forum_viewthread) Put() {
	data.Tid = 0
	data.Page = 0
	data.Ordertype = 0
	data.Stand = 0
	data.Authorid = 0
	data.Position = 0
	data.Fromuid = 0
	pool_MSG_U2WS_forum_viewthread.Put(data)
}
func (data *MSG_U2WS_forum_viewthread) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_forum_viewthread)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_forum_viewthread(data, buf)
}

func WRITE_MSG_U2WS_forum_viewthread(data *MSG_U2WS_forum_viewthread, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Tid, buf)
	WRITE_int16(data.Page, buf)
	WRITE_int8(data.Ordertype, buf)
	WRITE_int8(data.Stand, buf)
	WRITE_int32(data.Authorid, buf)
	WRITE_int32(data.Position, buf)
	WRITE_int32(data.Fromuid, buf)
}

func READ_MSG_U2WS_forum_viewthread(buf *libraries.MsgBuffer) (data *MSG_U2WS_forum_viewthread) {
	data = pool_MSG_U2WS_forum_viewthread.Get().(*MSG_U2WS_forum_viewthread)
	data.Tid = READ_int32(buf)
	data.Page = READ_int16(buf)
	data.Ordertype = READ_int8(buf)
	data.Stand = READ_int8(buf)
	data.Authorid = READ_int32(buf)
	data.Position = READ_int32(buf)
	data.Fromuid = READ_int32(buf)
	return
}

type MSG_WS2U_forum_viewthread struct {
	Avatar string
	Fid int32
	Name string
	Parent *MSG_forum_idname
	Forum *MSG_forum_thread_forum
	Group_status int32
	Admin_status int32
	Allowstickthread int8
	Modmenu int8
	Medal_list []*MSG_forum_post_medal
	Thread *MSG_forum_thread
	Page int16
	Postlist []*MSG_forum_post
	Credits int32
	Posts int32
	Digestposts int32
	Blockedpids []int32
	Firststand int8
	Maxattachsize int32
	Allowuploadnum int16
	Allowuploadsize int32
	Allowrecommend int8
	Edittimelimit int32
	Recent_use_tag []string
	Reason string
	Poll *MSG_Poll_info
	Forum_history []*MSG_forum_idname
}

var pool_MSG_WS2U_forum_viewthread = sync.Pool{New: func() interface{} { return &MSG_WS2U_forum_viewthread{} }}

func GET_MSG_WS2U_forum_viewthread() *MSG_WS2U_forum_viewthread {
	return pool_MSG_WS2U_forum_viewthread.Get().(*MSG_WS2U_forum_viewthread)
}

func (data *MSG_WS2U_forum_viewthread) Put() {
	data.Avatar = ``
	data.Fid = 0
	data.Name = ``
	if data.Parent != nil {
		data.Parent.Put()
		data.Parent = nil
	}
	if data.Forum != nil {
		data.Forum.Put()
		data.Forum = nil
	}
	data.Group_status = 0
	data.Admin_status = 0
	data.Allowstickthread = 0
	data.Modmenu = 0
	for _,v := range data.Medal_list {
		v.Put()
	}
	data.Medal_list = data.Medal_list[:0]
	if data.Thread != nil {
		data.Thread.Put()
		data.Thread = nil
	}
	data.Page = 0
	for _,v := range data.Postlist {
		v.Put()
	}
	data.Postlist = data.Postlist[:0]
	data.Credits = 0
	data.Posts = 0
	data.Digestposts = 0
	data.Blockedpids = data.Blockedpids[:0]
	data.Firststand = 0
	data.Maxattachsize = 0
	data.Allowuploadnum = 0
	data.Allowuploadsize = 0
	data.Allowrecommend = 0
	data.Edittimelimit = 0
	data.Recent_use_tag = data.Recent_use_tag[:0]
	data.Reason = ``
	if data.Poll != nil {
		data.Poll.Put()
		data.Poll = nil
	}
	for _,v := range data.Forum_history {
		v.Put()
	}
	data.Forum_history = data.Forum_history[:0]
	pool_MSG_WS2U_forum_viewthread.Put(data)
}
func (data *MSG_WS2U_forum_viewthread) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_forum_viewthread)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_forum_viewthread(data, buf)
}

func WRITE_MSG_WS2U_forum_viewthread(data *MSG_WS2U_forum_viewthread, buf *libraries.MsgBuffer) {
	WRITE_string(data.Avatar, buf)
	WRITE_int32(data.Fid, buf)
	WRITE_string(data.Name, buf)
	if data.Parent == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_forum_idname(data.Parent, buf)
	}
	if data.Forum == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_forum_thread_forum(data.Forum, buf)
	}
	WRITE_int32(data.Group_status, buf)
	WRITE_int32(data.Admin_status, buf)
	WRITE_int8(data.Allowstickthread, buf)
	WRITE_int8(data.Modmenu, buf)
	WRITE_int32(int32(len(data.Medal_list)), buf)
	for _, v := range data.Medal_list{
		WRITE_MSG_forum_post_medal(v, buf)
	}
	if data.Thread == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_forum_thread(data.Thread, buf)
	}
	WRITE_int16(data.Page, buf)
	WRITE_int32(int32(len(data.Postlist)), buf)
	for _, v := range data.Postlist{
		WRITE_MSG_forum_post(v, buf)
	}
	WRITE_int32(data.Credits, buf)
	WRITE_int32(data.Posts, buf)
	WRITE_int32(data.Digestposts, buf)
	WRITE_int32(int32(len(data.Blockedpids)), buf)
	for _, v := range data.Blockedpids{
		WRITE_int32(v, buf)
	}
	WRITE_int8(data.Firststand, buf)
	WRITE_int32(data.Maxattachsize, buf)
	WRITE_int16(data.Allowuploadnum, buf)
	WRITE_int32(data.Allowuploadsize, buf)
	WRITE_int8(data.Allowrecommend, buf)
	WRITE_int32(data.Edittimelimit, buf)
	WRITE_int32(int32(len(data.Recent_use_tag)), buf)
	for _, v := range data.Recent_use_tag{
		WRITE_string(v, buf)
	}
	WRITE_string(data.Reason, buf)
	if data.Poll == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_Poll_info(data.Poll, buf)
	}
	WRITE_int32(int32(len(data.Forum_history)), buf)
	for _, v := range data.Forum_history{
		WRITE_MSG_forum_idname(v, buf)
	}
}

func READ_MSG_WS2U_forum_viewthread(buf *libraries.MsgBuffer) (data *MSG_WS2U_forum_viewthread) {
	data = pool_MSG_WS2U_forum_viewthread.Get().(*MSG_WS2U_forum_viewthread)
	data.Avatar = READ_string(buf)
	data.Fid = READ_int32(buf)
	data.Name = READ_string(buf)
	Parent_len := int(READ_int8(buf))
	if Parent_len == 1 {
		data.Parent = READ_MSG_forum_idname(buf)
	}else{
		data.Parent = nil
	}
	Forum_len := int(READ_int8(buf))
	if Forum_len == 1 {
		data.Forum = READ_MSG_forum_thread_forum(buf)
	}else{
		data.Forum = nil
	}
	data.Group_status = READ_int32(buf)
	data.Admin_status = READ_int32(buf)
	data.Allowstickthread = READ_int8(buf)
	data.Modmenu = READ_int8(buf)
	Medal_list_len := int(READ_int32(buf))
	for i := 0; i < Medal_list_len; i++ {
		data.Medal_list = append(data.Medal_list, READ_MSG_forum_post_medal(buf))
	}
	Thread_len := int(READ_int8(buf))
	if Thread_len == 1 {
		data.Thread = READ_MSG_forum_thread(buf)
	}else{
		data.Thread = nil
	}
	data.Page = READ_int16(buf)
	Postlist_len := int(READ_int32(buf))
	for i := 0; i < Postlist_len; i++ {
		data.Postlist = append(data.Postlist, READ_MSG_forum_post(buf))
	}
	data.Credits = READ_int32(buf)
	data.Posts = READ_int32(buf)
	data.Digestposts = READ_int32(buf)
	Blockedpids_len := int(READ_int32(buf))
	for i := 0; i < Blockedpids_len; i++ {
		data.Blockedpids = append(data.Blockedpids, READ_int32(buf))
	}
	data.Firststand = READ_int8(buf)
	data.Maxattachsize = READ_int32(buf)
	data.Allowuploadnum = READ_int16(buf)
	data.Allowuploadsize = READ_int32(buf)
	data.Allowrecommend = READ_int8(buf)
	data.Edittimelimit = READ_int32(buf)
	Recent_use_tag_len := int(READ_int32(buf))
	for i := 0; i < Recent_use_tag_len; i++ {
		data.Recent_use_tag = append(data.Recent_use_tag, READ_string(buf))
	}
	data.Reason = READ_string(buf)
	Poll_len := int(READ_int8(buf))
	if Poll_len == 1 {
		data.Poll = READ_MSG_Poll_info(buf)
	}else{
		data.Poll = nil
	}
	Forum_history_len := int(READ_int32(buf))
	for i := 0; i < Forum_history_len; i++ {
		data.Forum_history = append(data.Forum_history, READ_MSG_forum_idname(buf))
	}
	return
}

type MSG_forum_thread_forum struct {
	Picstyle int8
	Threadtypes *MSG_forum_threadtype
	Status int8
	Ismoderator int8
	Allowspecialonly int8
	Alloweditpost int8
	Disablecollect int8
}

var pool_MSG_forum_thread_forum = sync.Pool{New: func() interface{} { return &MSG_forum_thread_forum{} }}

func GET_MSG_forum_thread_forum() *MSG_forum_thread_forum {
	return pool_MSG_forum_thread_forum.Get().(*MSG_forum_thread_forum)
}

func (data *MSG_forum_thread_forum) Put() {
	data.Picstyle = 0
	if data.Threadtypes != nil {
		data.Threadtypes.Put()
		data.Threadtypes = nil
	}
	data.Status = 0
	data.Ismoderator = 0
	data.Allowspecialonly = 0
	data.Alloweditpost = 0
	data.Disablecollect = 0
	pool_MSG_forum_thread_forum.Put(data)
}
func (data *MSG_forum_thread_forum) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_forum_thread_forum)
	WRITE_int32(cmd, buf)
	WRITE_MSG_forum_thread_forum(data, buf)
}

func WRITE_MSG_forum_thread_forum(data *MSG_forum_thread_forum, buf *libraries.MsgBuffer) {
	WRITE_int8(data.Picstyle, buf)
	if data.Threadtypes == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_forum_threadtype(data.Threadtypes, buf)
	}
	WRITE_int8(data.Status, buf)
	WRITE_int8(data.Ismoderator, buf)
	WRITE_int8(data.Allowspecialonly, buf)
	WRITE_int8(data.Alloweditpost, buf)
	WRITE_int8(data.Disablecollect, buf)
}

func READ_MSG_forum_thread_forum(buf *libraries.MsgBuffer) (data *MSG_forum_thread_forum) {
	data = pool_MSG_forum_thread_forum.Get().(*MSG_forum_thread_forum)
	data.Picstyle = READ_int8(buf)
	Threadtypes_len := int(READ_int8(buf))
	if Threadtypes_len == 1 {
		data.Threadtypes = READ_MSG_forum_threadtype(buf)
	}else{
		data.Threadtypes = nil
	}
	data.Status = READ_int8(buf)
	data.Ismoderator = READ_int8(buf)
	data.Allowspecialonly = READ_int8(buf)
	data.Alloweditpost = READ_int8(buf)
	data.Disablecollect = READ_int8(buf)
	return
}

type MSG_forum_post struct {
	StatusBool int32
	Adminid int16
	Grouptitle string
	Header string
	Footer string
	Groupcolor string
	Memberstatus int32
	Invisible int8
	Author string
	Avatar string
	Authorid int32
	Dateline int32
	Gender int32
	Groupid int16
	Imagelistcount int8
	Message string
	Mobiletype int8
	Number int32
	Position int32
	Profile *MSG_post_member_profile
	Ratelog []*MSG_post_ratelog
	Ratelogextcredits []*MSG_post_ratelog_score
	Relateitem *MSG_post_relateitem
	Replycredit int32
	Signature string
	Stand int8
	Subject string
	Tags []string
	Useip string
	Username string
	Voters int32
	Lastmod *MSG_threadmod
	Comments *MSG_post_comment
	Totalcomment int16
	Commentcount int16
	Totalrate int16
	Location string
	Attachlist []string
	Imagelist []*MSG_forum_imgattach
	Releatcollectionnum int32
	Threadnum int32
	Digestnum int32
	Extcredits1 int32
	Extcredits2 int32
	Extcredits3 int32
}

var pool_MSG_forum_post = sync.Pool{New: func() interface{} { return &MSG_forum_post{} }}

func GET_MSG_forum_post() *MSG_forum_post {
	return pool_MSG_forum_post.Get().(*MSG_forum_post)
}

func (data *MSG_forum_post) Put() {
	data.StatusBool = 0
	data.Adminid = 0
	data.Grouptitle = ``
	data.Header = ``
	data.Footer = ``
	data.Groupcolor = ``
	data.Memberstatus = 0
	data.Invisible = 0
	data.Author = ``
	data.Avatar = ``
	data.Authorid = 0
	data.Dateline = 0
	data.Gender = 0
	data.Groupid = 0
	data.Imagelistcount = 0
	data.Message = ``
	data.Mobiletype = 0
	data.Number = 0
	data.Position = 0
	if data.Profile != nil {
		data.Profile.Put()
		data.Profile = nil
	}
	for _,v := range data.Ratelog {
		v.Put()
	}
	data.Ratelog = data.Ratelog[:0]
	for _,v := range data.Ratelogextcredits {
		v.Put()
	}
	data.Ratelogextcredits = data.Ratelogextcredits[:0]
	if data.Relateitem != nil {
		data.Relateitem.Put()
		data.Relateitem = nil
	}
	data.Replycredit = 0
	data.Signature = ``
	data.Stand = 0
	data.Subject = ``
	data.Tags = data.Tags[:0]
	data.Useip = ``
	data.Username = ``
	data.Voters = 0
	if data.Lastmod != nil {
		data.Lastmod.Put()
		data.Lastmod = nil
	}
	if data.Comments != nil {
		data.Comments.Put()
		data.Comments = nil
	}
	data.Totalcomment = 0
	data.Commentcount = 0
	data.Totalrate = 0
	data.Location = ``
	data.Attachlist = data.Attachlist[:0]
	for _,v := range data.Imagelist {
		v.Put()
	}
	data.Imagelist = data.Imagelist[:0]
	data.Releatcollectionnum = 0
	data.Threadnum = 0
	data.Digestnum = 0
	data.Extcredits1 = 0
	data.Extcredits2 = 0
	data.Extcredits3 = 0
	pool_MSG_forum_post.Put(data)
}
func (data *MSG_forum_post) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_forum_post)
	WRITE_int32(cmd, buf)
	WRITE_MSG_forum_post(data, buf)
}

func WRITE_MSG_forum_post(data *MSG_forum_post, buf *libraries.MsgBuffer) {
	WRITE_int32(data.StatusBool, buf)
	WRITE_int16(data.Adminid, buf)
	WRITE_string(data.Grouptitle, buf)
	WRITE_string(data.Header, buf)
	WRITE_string(data.Footer, buf)
	WRITE_string(data.Groupcolor, buf)
	WRITE_int32(data.Memberstatus, buf)
	WRITE_int8(data.Invisible, buf)
	WRITE_string(data.Author, buf)
	WRITE_string(data.Avatar, buf)
	WRITE_int32(data.Authorid, buf)
	WRITE_int32(data.Dateline, buf)
	WRITE_int32(data.Gender, buf)
	WRITE_int16(data.Groupid, buf)
	WRITE_int8(data.Imagelistcount, buf)
	WRITE_string(data.Message, buf)
	WRITE_int8(data.Mobiletype, buf)
	WRITE_int32(data.Number, buf)
	WRITE_int32(data.Position, buf)
	if data.Profile == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_post_member_profile(data.Profile, buf)
	}
	WRITE_int32(int32(len(data.Ratelog)), buf)
	for _, v := range data.Ratelog{
		WRITE_MSG_post_ratelog(v, buf)
	}
	WRITE_int32(int32(len(data.Ratelogextcredits)), buf)
	for _, v := range data.Ratelogextcredits{
		WRITE_MSG_post_ratelog_score(v, buf)
	}
	if data.Relateitem == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_post_relateitem(data.Relateitem, buf)
	}
	WRITE_int32(data.Replycredit, buf)
	WRITE_string(data.Signature, buf)
	WRITE_int8(data.Stand, buf)
	WRITE_string(data.Subject, buf)
	WRITE_int32(int32(len(data.Tags)), buf)
	for _, v := range data.Tags{
		WRITE_string(v, buf)
	}
	WRITE_string(data.Useip, buf)
	WRITE_string(data.Username, buf)
	WRITE_int32(data.Voters, buf)
	if data.Lastmod == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_threadmod(data.Lastmod, buf)
	}
	if data.Comments == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_post_comment(data.Comments, buf)
	}
	WRITE_int16(data.Totalcomment, buf)
	WRITE_int16(data.Commentcount, buf)
	WRITE_int16(data.Totalrate, buf)
	WRITE_string(data.Location, buf)
	WRITE_int32(int32(len(data.Attachlist)), buf)
	for _, v := range data.Attachlist{
		WRITE_string(v, buf)
	}
	WRITE_int32(int32(len(data.Imagelist)), buf)
	for _, v := range data.Imagelist{
		WRITE_MSG_forum_imgattach(v, buf)
	}
	WRITE_int32(data.Releatcollectionnum, buf)
	WRITE_int32(data.Threadnum, buf)
	WRITE_int32(data.Digestnum, buf)
	WRITE_int32(data.Extcredits1, buf)
	WRITE_int32(data.Extcredits2, buf)
	WRITE_int32(data.Extcredits3, buf)
}

func READ_MSG_forum_post(buf *libraries.MsgBuffer) (data *MSG_forum_post) {
	data = pool_MSG_forum_post.Get().(*MSG_forum_post)
	data.StatusBool = READ_int32(buf)
	data.Adminid = READ_int16(buf)
	data.Grouptitle = READ_string(buf)
	data.Header = READ_string(buf)
	data.Footer = READ_string(buf)
	data.Groupcolor = READ_string(buf)
	data.Memberstatus = READ_int32(buf)
	data.Invisible = READ_int8(buf)
	data.Author = READ_string(buf)
	data.Avatar = READ_string(buf)
	data.Authorid = READ_int32(buf)
	data.Dateline = READ_int32(buf)
	data.Gender = READ_int32(buf)
	data.Groupid = READ_int16(buf)
	data.Imagelistcount = READ_int8(buf)
	data.Message = READ_string(buf)
	data.Mobiletype = READ_int8(buf)
	data.Number = READ_int32(buf)
	data.Position = READ_int32(buf)
	Profile_len := int(READ_int8(buf))
	if Profile_len == 1 {
		data.Profile = READ_MSG_post_member_profile(buf)
	}else{
		data.Profile = nil
	}
	Ratelog_len := int(READ_int32(buf))
	for i := 0; i < Ratelog_len; i++ {
		data.Ratelog = append(data.Ratelog, READ_MSG_post_ratelog(buf))
	}
	Ratelogextcredits_len := int(READ_int32(buf))
	for i := 0; i < Ratelogextcredits_len; i++ {
		data.Ratelogextcredits = append(data.Ratelogextcredits, READ_MSG_post_ratelog_score(buf))
	}
	Relateitem_len := int(READ_int8(buf))
	if Relateitem_len == 1 {
		data.Relateitem = READ_MSG_post_relateitem(buf)
	}else{
		data.Relateitem = nil
	}
	data.Replycredit = READ_int32(buf)
	data.Signature = READ_string(buf)
	data.Stand = READ_int8(buf)
	data.Subject = READ_string(buf)
	Tags_len := int(READ_int32(buf))
	for i := 0; i < Tags_len; i++ {
		data.Tags = append(data.Tags, READ_string(buf))
	}
	data.Useip = READ_string(buf)
	data.Username = READ_string(buf)
	data.Voters = READ_int32(buf)
	Lastmod_len := int(READ_int8(buf))
	if Lastmod_len == 1 {
		data.Lastmod = READ_MSG_threadmod(buf)
	}else{
		data.Lastmod = nil
	}
	Comments_len := int(READ_int8(buf))
	if Comments_len == 1 {
		data.Comments = READ_MSG_post_comment(buf)
	}else{
		data.Comments = nil
	}
	data.Totalcomment = READ_int16(buf)
	data.Commentcount = READ_int16(buf)
	data.Totalrate = READ_int16(buf)
	data.Location = READ_string(buf)
	Attachlist_len := int(READ_int32(buf))
	for i := 0; i < Attachlist_len; i++ {
		data.Attachlist = append(data.Attachlist, READ_string(buf))
	}
	Imagelist_len := int(READ_int32(buf))
	for i := 0; i < Imagelist_len; i++ {
		data.Imagelist = append(data.Imagelist, READ_MSG_forum_imgattach(buf))
	}
	data.Releatcollectionnum = READ_int32(buf)
	data.Threadnum = READ_int32(buf)
	data.Digestnum = READ_int32(buf)
	data.Extcredits1 = READ_int32(buf)
	data.Extcredits2 = READ_int32(buf)
	data.Extcredits3 = READ_int32(buf)
	return
}

type MSG_forum_post_medal struct {
	Id int16
	Name string
	Description string
}

var pool_MSG_forum_post_medal = sync.Pool{New: func() interface{} { return &MSG_forum_post_medal{} }}

func GET_MSG_forum_post_medal() *MSG_forum_post_medal {
	return pool_MSG_forum_post_medal.Get().(*MSG_forum_post_medal)
}

func (data *MSG_forum_post_medal) Put() {
	data.Id = 0
	data.Name = ``
	data.Description = ``
	pool_MSG_forum_post_medal.Put(data)
}
func (data *MSG_forum_post_medal) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_forum_post_medal)
	WRITE_int32(cmd, buf)
	WRITE_MSG_forum_post_medal(data, buf)
}

func WRITE_MSG_forum_post_medal(data *MSG_forum_post_medal, buf *libraries.MsgBuffer) {
	WRITE_int16(data.Id, buf)
	WRITE_string(data.Name, buf)
	WRITE_string(data.Description, buf)
}

func READ_MSG_forum_post_medal(buf *libraries.MsgBuffer) (data *MSG_forum_post_medal) {
	data = pool_MSG_forum_post_medal.Get().(*MSG_forum_post_medal)
	data.Id = READ_int16(buf)
	data.Name = READ_string(buf)
	data.Description = READ_string(buf)
	return
}

type MSG_postreview struct {
	Support int16
	Against int16
}

var pool_MSG_postreview = sync.Pool{New: func() interface{} { return &MSG_postreview{} }}

func GET_MSG_postreview() *MSG_postreview {
	return pool_MSG_postreview.Get().(*MSG_postreview)
}

func (data *MSG_postreview) Put() {
	data.Support = 0
	data.Against = 0
	pool_MSG_postreview.Put(data)
}
func (data *MSG_postreview) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_postreview)
	WRITE_int32(cmd, buf)
	WRITE_MSG_postreview(data, buf)
}

func WRITE_MSG_postreview(data *MSG_postreview, buf *libraries.MsgBuffer) {
	WRITE_int16(data.Support, buf)
	WRITE_int16(data.Against, buf)
}

func READ_MSG_postreview(buf *libraries.MsgBuffer) (data *MSG_postreview) {
	data = pool_MSG_postreview.Get().(*MSG_postreview)
	data.Support = READ_int16(buf)
	data.Against = READ_int16(buf)
	return
}

type MSG_post_member_profile struct {
	Mobile string
	Wx string
	Qq int64
	Site string
}

var pool_MSG_post_member_profile = sync.Pool{New: func() interface{} { return &MSG_post_member_profile{} }}

func GET_MSG_post_member_profile() *MSG_post_member_profile {
	return pool_MSG_post_member_profile.Get().(*MSG_post_member_profile)
}

func (data *MSG_post_member_profile) Put() {
	data.Mobile = ``
	data.Wx = ``
	data.Qq = 0
	data.Site = ``
	pool_MSG_post_member_profile.Put(data)
}
func (data *MSG_post_member_profile) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_post_member_profile)
	WRITE_int32(cmd, buf)
	WRITE_MSG_post_member_profile(data, buf)
}

func WRITE_MSG_post_member_profile(data *MSG_post_member_profile, buf *libraries.MsgBuffer) {
	WRITE_string(data.Mobile, buf)
	WRITE_string(data.Wx, buf)
	WRITE_int64(data.Qq, buf)
	WRITE_string(data.Site, buf)
}

func READ_MSG_post_member_profile(buf *libraries.MsgBuffer) (data *MSG_post_member_profile) {
	data = pool_MSG_post_member_profile.Get().(*MSG_post_member_profile)
	data.Mobile = READ_string(buf)
	data.Wx = READ_string(buf)
	data.Qq = READ_int64(buf)
	data.Site = READ_string(buf)
	return
}

type MSG_post_ratelog struct {
	Username string
	Avatar string
	Score []*MSG_post_ratelog_score
	Reason string
	Uid int32
}

var pool_MSG_post_ratelog = sync.Pool{New: func() interface{} { return &MSG_post_ratelog{} }}

func GET_MSG_post_ratelog() *MSG_post_ratelog {
	return pool_MSG_post_ratelog.Get().(*MSG_post_ratelog)
}

func (data *MSG_post_ratelog) Put() {
	data.Username = ``
	data.Avatar = ``
	for _,v := range data.Score {
		v.Put()
	}
	data.Score = data.Score[:0]
	data.Reason = ``
	data.Uid = 0
	pool_MSG_post_ratelog.Put(data)
}
func (data *MSG_post_ratelog) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_post_ratelog)
	WRITE_int32(cmd, buf)
	WRITE_MSG_post_ratelog(data, buf)
}

func WRITE_MSG_post_ratelog(data *MSG_post_ratelog, buf *libraries.MsgBuffer) {
	WRITE_string(data.Username, buf)
	WRITE_string(data.Avatar, buf)
	WRITE_int32(int32(len(data.Score)), buf)
	for _, v := range data.Score{
		WRITE_MSG_post_ratelog_score(v, buf)
	}
	WRITE_string(data.Reason, buf)
	WRITE_int32(data.Uid, buf)
}

func READ_MSG_post_ratelog(buf *libraries.MsgBuffer) (data *MSG_post_ratelog) {
	data = pool_MSG_post_ratelog.Get().(*MSG_post_ratelog)
	data.Username = READ_string(buf)
	data.Avatar = READ_string(buf)
	Score_len := int(READ_int32(buf))
	for i := 0; i < Score_len; i++ {
		data.Score = append(data.Score, READ_MSG_post_ratelog_score(buf))
	}
	data.Reason = READ_string(buf)
	data.Uid = READ_int32(buf)
	return
}

type MSG_post_ratelog_score struct {
	Id int8
	Score int32
}

var pool_MSG_post_ratelog_score = sync.Pool{New: func() interface{} { return &MSG_post_ratelog_score{} }}

func GET_MSG_post_ratelog_score() *MSG_post_ratelog_score {
	return pool_MSG_post_ratelog_score.Get().(*MSG_post_ratelog_score)
}

func (data *MSG_post_ratelog_score) Put() {
	data.Id = 0
	data.Score = 0
	pool_MSG_post_ratelog_score.Put(data)
}
func (data *MSG_post_ratelog_score) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_post_ratelog_score)
	WRITE_int32(cmd, buf)
	WRITE_MSG_post_ratelog_score(data, buf)
}

func WRITE_MSG_post_ratelog_score(data *MSG_post_ratelog_score, buf *libraries.MsgBuffer) {
	WRITE_int8(data.Id, buf)
	WRITE_int32(data.Score, buf)
}

func READ_MSG_post_ratelog_score(buf *libraries.MsgBuffer) (data *MSG_post_ratelog_score) {
	data = pool_MSG_post_ratelog_score.Get().(*MSG_post_ratelog_score)
	data.Id = READ_int8(buf)
	data.Score = READ_int32(buf)
	return
}

type MSG_post_relateitem struct {
	Tid int32
	Subject string
}

var pool_MSG_post_relateitem = sync.Pool{New: func() interface{} { return &MSG_post_relateitem{} }}

func GET_MSG_post_relateitem() *MSG_post_relateitem {
	return pool_MSG_post_relateitem.Get().(*MSG_post_relateitem)
}

func (data *MSG_post_relateitem) Put() {
	data.Tid = 0
	data.Subject = ``
	pool_MSG_post_relateitem.Put(data)
}
func (data *MSG_post_relateitem) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_post_relateitem)
	WRITE_int32(cmd, buf)
	WRITE_MSG_post_relateitem(data, buf)
}

func WRITE_MSG_post_relateitem(data *MSG_post_relateitem, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Tid, buf)
	WRITE_string(data.Subject, buf)
}

func READ_MSG_post_relateitem(buf *libraries.MsgBuffer) (data *MSG_post_relateitem) {
	data = pool_MSG_post_relateitem.Get().(*MSG_post_relateitem)
	data.Tid = READ_int32(buf)
	data.Subject = READ_string(buf)
	return
}

type MSG_threadmod struct {
	Uid int32
	Modactiontype string
	Modusername string
	Moddateline int32
	Expiration int32
	Reason string
	Stamp int8
	Status int8
}

var pool_MSG_threadmod = sync.Pool{New: func() interface{} { return &MSG_threadmod{} }}

func GET_MSG_threadmod() *MSG_threadmod {
	return pool_MSG_threadmod.Get().(*MSG_threadmod)
}

func (data *MSG_threadmod) Put() {
	data.Uid = 0
	data.Modactiontype = ``
	data.Modusername = ``
	data.Moddateline = 0
	data.Expiration = 0
	data.Reason = ``
	data.Stamp = 0
	data.Status = 0
	pool_MSG_threadmod.Put(data)
}
func (data *MSG_threadmod) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_threadmod)
	WRITE_int32(cmd, buf)
	WRITE_MSG_threadmod(data, buf)
}

func WRITE_MSG_threadmod(data *MSG_threadmod, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Uid, buf)
	WRITE_string(data.Modactiontype, buf)
	WRITE_string(data.Modusername, buf)
	WRITE_int32(data.Moddateline, buf)
	WRITE_int32(data.Expiration, buf)
	WRITE_string(data.Reason, buf)
	WRITE_int8(data.Stamp, buf)
	WRITE_int8(data.Status, buf)
}

func READ_MSG_threadmod(buf *libraries.MsgBuffer) (data *MSG_threadmod) {
	data = pool_MSG_threadmod.Get().(*MSG_threadmod)
	data.Uid = READ_int32(buf)
	data.Modactiontype = READ_string(buf)
	data.Modusername = READ_string(buf)
	data.Moddateline = READ_int32(buf)
	data.Expiration = READ_int32(buf)
	data.Reason = READ_string(buf)
	data.Stamp = READ_int8(buf)
	data.Status = READ_int8(buf)
	return
}

type MSG_post_comment struct {
	Authorid int32
	Avatar string
	Author string
	Comment string
	Rpid string
	Useip string
	Id int32
	Dateline int32
}

var pool_MSG_post_comment = sync.Pool{New: func() interface{} { return &MSG_post_comment{} }}

func GET_MSG_post_comment() *MSG_post_comment {
	return pool_MSG_post_comment.Get().(*MSG_post_comment)
}

func (data *MSG_post_comment) Put() {
	data.Authorid = 0
	data.Avatar = ``
	data.Author = ``
	data.Comment = ``
	data.Rpid = ``
	data.Useip = ``
	data.Id = 0
	data.Dateline = 0
	pool_MSG_post_comment.Put(data)
}
func (data *MSG_post_comment) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_post_comment)
	WRITE_int32(cmd, buf)
	WRITE_MSG_post_comment(data, buf)
}

func WRITE_MSG_post_comment(data *MSG_post_comment, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Authorid, buf)
	WRITE_string(data.Avatar, buf)
	WRITE_string(data.Author, buf)
	WRITE_string(data.Comment, buf)
	WRITE_string(data.Rpid, buf)
	WRITE_string(data.Useip, buf)
	WRITE_int32(data.Id, buf)
	WRITE_int32(data.Dateline, buf)
}

func READ_MSG_post_comment(buf *libraries.MsgBuffer) (data *MSG_post_comment) {
	data = pool_MSG_post_comment.Get().(*MSG_post_comment)
	data.Authorid = READ_int32(buf)
	data.Avatar = READ_string(buf)
	data.Author = READ_string(buf)
	data.Comment = READ_string(buf)
	data.Rpid = READ_string(buf)
	data.Useip = READ_string(buf)
	data.Id = READ_int32(buf)
	data.Dateline = READ_int32(buf)
	return
}

type MSG_U2WS_threadfastpost struct {
	Tid int32
	Position int32
	Subject string
	Message string
	Seccode string
	Other int8
}

var pool_MSG_U2WS_threadfastpost = sync.Pool{New: func() interface{} { return &MSG_U2WS_threadfastpost{} }}

func GET_MSG_U2WS_threadfastpost() *MSG_U2WS_threadfastpost {
	return pool_MSG_U2WS_threadfastpost.Get().(*MSG_U2WS_threadfastpost)
}

func (data *MSG_U2WS_threadfastpost) Put() {
	data.Tid = 0
	data.Position = 0
	data.Subject = ``
	data.Message = ``
	data.Seccode = ``
	data.Other = 0
	pool_MSG_U2WS_threadfastpost.Put(data)
}
func (data *MSG_U2WS_threadfastpost) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_threadfastpost)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_threadfastpost(data, buf)
}

func WRITE_MSG_U2WS_threadfastpost(data *MSG_U2WS_threadfastpost, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Tid, buf)
	WRITE_int32(data.Position, buf)
	WRITE_string(data.Subject, buf)
	WRITE_string(data.Message, buf)
	WRITE_string(data.Seccode, buf)
	WRITE_int8(data.Other, buf)
}

func READ_MSG_U2WS_threadfastpost(buf *libraries.MsgBuffer) (data *MSG_U2WS_threadfastpost) {
	data = pool_MSG_U2WS_threadfastpost.Get().(*MSG_U2WS_threadfastpost)
	data.Tid = READ_int32(buf)
	data.Position = READ_int32(buf)
	data.Subject = READ_string(buf)
	data.Message = READ_string(buf)
	data.Seccode = READ_string(buf)
	data.Other = READ_int8(buf)
	return
}

type MSG_WS2U_threadfastpost struct {
	Result int16
	Page int16
	Add_info *MSG_forum_post
}

var pool_MSG_WS2U_threadfastpost = sync.Pool{New: func() interface{} { return &MSG_WS2U_threadfastpost{} }}

func GET_MSG_WS2U_threadfastpost() *MSG_WS2U_threadfastpost {
	return pool_MSG_WS2U_threadfastpost.Get().(*MSG_WS2U_threadfastpost)
}

func (data *MSG_WS2U_threadfastpost) Put() {
	data.Result = 0
	data.Page = 0
	if data.Add_info != nil {
		data.Add_info.Put()
		data.Add_info = nil
	}
	pool_MSG_WS2U_threadfastpost.Put(data)
}
func (data *MSG_WS2U_threadfastpost) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_threadfastpost)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_threadfastpost(data, buf)
}

func WRITE_MSG_WS2U_threadfastpost(data *MSG_WS2U_threadfastpost, buf *libraries.MsgBuffer) {
	WRITE_int16(data.Result, buf)
	WRITE_int16(data.Page, buf)
	if data.Add_info == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_forum_post(data.Add_info, buf)
	}
}

func READ_MSG_WS2U_threadfastpost(buf *libraries.MsgBuffer) (data *MSG_WS2U_threadfastpost) {
	data = pool_MSG_WS2U_threadfastpost.Get().(*MSG_WS2U_threadfastpost)
	data.Result = READ_int16(buf)
	data.Page = READ_int16(buf)
	Add_info_len := int(READ_int8(buf))
	if Add_info_len == 1 {
		data.Add_info = READ_MSG_forum_post(buf)
	}else{
		data.Add_info = nil
	}
	return
}

type MSG_U2WS_nextset struct {
	Next int8
	Tid int32
}

var pool_MSG_U2WS_nextset = sync.Pool{New: func() interface{} { return &MSG_U2WS_nextset{} }}

func GET_MSG_U2WS_nextset() *MSG_U2WS_nextset {
	return pool_MSG_U2WS_nextset.Get().(*MSG_U2WS_nextset)
}

func (data *MSG_U2WS_nextset) Put() {
	data.Next = 0
	data.Tid = 0
	pool_MSG_U2WS_nextset.Put(data)
}
func (data *MSG_U2WS_nextset) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_nextset)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_nextset(data, buf)
}

func WRITE_MSG_U2WS_nextset(data *MSG_U2WS_nextset, buf *libraries.MsgBuffer) {
	WRITE_int8(data.Next, buf)
	WRITE_int32(data.Tid, buf)
}

func READ_MSG_U2WS_nextset(buf *libraries.MsgBuffer) (data *MSG_U2WS_nextset) {
	data = pool_MSG_U2WS_nextset.Get().(*MSG_U2WS_nextset)
	data.Next = READ_int8(buf)
	data.Tid = READ_int32(buf)
	return
}

type MSG_WS2U_nextset struct {
	Tid int32
}

var pool_MSG_WS2U_nextset = sync.Pool{New: func() interface{} { return &MSG_WS2U_nextset{} }}

func GET_MSG_WS2U_nextset() *MSG_WS2U_nextset {
	return pool_MSG_WS2U_nextset.Get().(*MSG_WS2U_nextset)
}

func (data *MSG_WS2U_nextset) Put() {
	data.Tid = 0
	pool_MSG_WS2U_nextset.Put(data)
}
func (data *MSG_WS2U_nextset) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_nextset)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_nextset(data, buf)
}

func WRITE_MSG_WS2U_nextset(data *MSG_WS2U_nextset, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Tid, buf)
}

func READ_MSG_WS2U_nextset(buf *libraries.MsgBuffer) (data *MSG_WS2U_nextset) {
	data = pool_MSG_WS2U_nextset.Get().(*MSG_WS2U_nextset)
	data.Tid = READ_int32(buf)
	return
}

type MSG_U2WS_upload_image struct {
	Filename string
	Data []byte
}

var pool_MSG_U2WS_upload_image = sync.Pool{New: func() interface{} { return &MSG_U2WS_upload_image{} }}

func GET_MSG_U2WS_upload_image() *MSG_U2WS_upload_image {
	return pool_MSG_U2WS_upload_image.Get().(*MSG_U2WS_upload_image)
}

func (data *MSG_U2WS_upload_image) Put() {
	data.Filename = ``
	data.Data = data.Data[:0]
	pool_MSG_U2WS_upload_image.Put(data)
}
func (data *MSG_U2WS_upload_image) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_upload_image)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_upload_image(data, buf)
}

func WRITE_MSG_U2WS_upload_image(data *MSG_U2WS_upload_image, buf *libraries.MsgBuffer) {
	WRITE_string(data.Filename, buf)
	WRITE_int32(int32(len(data.Data)), buf)
	buf.Write(data.Data)
}

func READ_MSG_U2WS_upload_image(buf *libraries.MsgBuffer) (data *MSG_U2WS_upload_image) {
	data = pool_MSG_U2WS_upload_image.Get().(*MSG_U2WS_upload_image)
	data.Filename = READ_string(buf)
	Data_len := int(READ_int32(buf))
	data.Data = make([]byte, Data_len)
	copy(data.Data,buf.Next(Data_len))
	return
}

type MSG_WS2U_upload_image struct {
	Img string
	Aid int64
	Title string
}

var pool_MSG_WS2U_upload_image = sync.Pool{New: func() interface{} { return &MSG_WS2U_upload_image{} }}

func GET_MSG_WS2U_upload_image() *MSG_WS2U_upload_image {
	return pool_MSG_WS2U_upload_image.Get().(*MSG_WS2U_upload_image)
}

func (data *MSG_WS2U_upload_image) Put() {
	data.Img = ``
	data.Aid = 0
	data.Title = ``
	pool_MSG_WS2U_upload_image.Put(data)
}
func (data *MSG_WS2U_upload_image) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_upload_image)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_upload_image(data, buf)
}

func WRITE_MSG_WS2U_upload_image(data *MSG_WS2U_upload_image, buf *libraries.MsgBuffer) {
	WRITE_string(data.Img, buf)
	WRITE_int64(data.Aid, buf)
	WRITE_string(data.Title, buf)
}

func READ_MSG_WS2U_upload_image(buf *libraries.MsgBuffer) (data *MSG_WS2U_upload_image) {
	data = pool_MSG_WS2U_upload_image.Get().(*MSG_WS2U_upload_image)
	data.Img = READ_string(buf)
	data.Aid = READ_int64(buf)
	data.Title = READ_string(buf)
	return
}

type MSG_U2WS_upload_tmp_image struct {
	Filename string
	Data []byte
}

var pool_MSG_U2WS_upload_tmp_image = sync.Pool{New: func() interface{} { return &MSG_U2WS_upload_tmp_image{} }}

func GET_MSG_U2WS_upload_tmp_image() *MSG_U2WS_upload_tmp_image {
	return pool_MSG_U2WS_upload_tmp_image.Get().(*MSG_U2WS_upload_tmp_image)
}

func (data *MSG_U2WS_upload_tmp_image) Put() {
	data.Filename = ``
	data.Data = data.Data[:0]
	pool_MSG_U2WS_upload_tmp_image.Put(data)
}
func (data *MSG_U2WS_upload_tmp_image) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_upload_tmp_image)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_upload_tmp_image(data, buf)
}

func WRITE_MSG_U2WS_upload_tmp_image(data *MSG_U2WS_upload_tmp_image, buf *libraries.MsgBuffer) {
	WRITE_string(data.Filename, buf)
	WRITE_int32(int32(len(data.Data)), buf)
	buf.Write(data.Data)
}

func READ_MSG_U2WS_upload_tmp_image(buf *libraries.MsgBuffer) (data *MSG_U2WS_upload_tmp_image) {
	data = pool_MSG_U2WS_upload_tmp_image.Get().(*MSG_U2WS_upload_tmp_image)
	data.Filename = READ_string(buf)
	Data_len := int(READ_int32(buf))
	data.Data = make([]byte, Data_len)
	copy(data.Data,buf.Next(Data_len))
	return
}

type MSG_WS2U_upload_tmp_image struct {
	Aid int64
}

var pool_MSG_WS2U_upload_tmp_image = sync.Pool{New: func() interface{} { return &MSG_WS2U_upload_tmp_image{} }}

func GET_MSG_WS2U_upload_tmp_image() *MSG_WS2U_upload_tmp_image {
	return pool_MSG_WS2U_upload_tmp_image.Get().(*MSG_WS2U_upload_tmp_image)
}

func (data *MSG_WS2U_upload_tmp_image) Put() {
	data.Aid = 0
	pool_MSG_WS2U_upload_tmp_image.Put(data)
}
func (data *MSG_WS2U_upload_tmp_image) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_upload_tmp_image)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_upload_tmp_image(data, buf)
}

func WRITE_MSG_WS2U_upload_tmp_image(data *MSG_WS2U_upload_tmp_image, buf *libraries.MsgBuffer) {
	WRITE_int64(data.Aid, buf)
}

func READ_MSG_WS2U_upload_tmp_image(buf *libraries.MsgBuffer) (data *MSG_WS2U_upload_tmp_image) {
	data = pool_MSG_WS2U_upload_tmp_image.Get().(*MSG_WS2U_upload_tmp_image)
	data.Aid = READ_int64(buf)
	return
}

type MSG_U2WS_delete_attach struct {
	Ids []int64
}

var pool_MSG_U2WS_delete_attach = sync.Pool{New: func() interface{} { return &MSG_U2WS_delete_attach{} }}

func GET_MSG_U2WS_delete_attach() *MSG_U2WS_delete_attach {
	return pool_MSG_U2WS_delete_attach.Get().(*MSG_U2WS_delete_attach)
}

func (data *MSG_U2WS_delete_attach) Put() {
	data.Ids = data.Ids[:0]
	pool_MSG_U2WS_delete_attach.Put(data)
}
func (data *MSG_U2WS_delete_attach) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_delete_attach)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_delete_attach(data, buf)
}

func WRITE_MSG_U2WS_delete_attach(data *MSG_U2WS_delete_attach, buf *libraries.MsgBuffer) {
	WRITE_int32(int32(len(data.Ids)), buf)
	for _, v := range data.Ids{
		WRITE_int64(v, buf)
	}
}

func READ_MSG_U2WS_delete_attach(buf *libraries.MsgBuffer) (data *MSG_U2WS_delete_attach) {
	data = pool_MSG_U2WS_delete_attach.Get().(*MSG_U2WS_delete_attach)
	Ids_len := int(READ_int32(buf))
	for i := 0; i < Ids_len; i++ {
		data.Ids = append(data.Ids, READ_int64(buf))
	}
	return
}

type MSG_WS2U_delete_attach struct {
	Ids []int64
}

var pool_MSG_WS2U_delete_attach = sync.Pool{New: func() interface{} { return &MSG_WS2U_delete_attach{} }}

func GET_MSG_WS2U_delete_attach() *MSG_WS2U_delete_attach {
	return pool_MSG_WS2U_delete_attach.Get().(*MSG_WS2U_delete_attach)
}

func (data *MSG_WS2U_delete_attach) Put() {
	data.Ids = data.Ids[:0]
	pool_MSG_WS2U_delete_attach.Put(data)
}
func (data *MSG_WS2U_delete_attach) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_delete_attach)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_delete_attach(data, buf)
}

func WRITE_MSG_WS2U_delete_attach(data *MSG_WS2U_delete_attach, buf *libraries.MsgBuffer) {
	WRITE_int32(int32(len(data.Ids)), buf)
	for _, v := range data.Ids{
		WRITE_int64(v, buf)
	}
}

func READ_MSG_WS2U_delete_attach(buf *libraries.MsgBuffer) (data *MSG_WS2U_delete_attach) {
	data = pool_MSG_WS2U_delete_attach.Get().(*MSG_WS2U_delete_attach)
	Ids_len := int(READ_int32(buf))
	for i := 0; i < Ids_len; i++ {
		data.Ids = append(data.Ids, READ_int64(buf))
	}
	return
}

type MSG_U2WS_threadmod struct {
	Tids []int32
	Expiration int32
	Action int8
	Param int8
	Param1 int32
	Reason string
	Sendreasonpm int8
}

var pool_MSG_U2WS_threadmod = sync.Pool{New: func() interface{} { return &MSG_U2WS_threadmod{} }}

func GET_MSG_U2WS_threadmod() *MSG_U2WS_threadmod {
	return pool_MSG_U2WS_threadmod.Get().(*MSG_U2WS_threadmod)
}

func (data *MSG_U2WS_threadmod) Put() {
	data.Tids = data.Tids[:0]
	data.Expiration = 0
	data.Action = 0
	data.Param = 0
	data.Param1 = 0
	data.Reason = ``
	data.Sendreasonpm = 0
	pool_MSG_U2WS_threadmod.Put(data)
}
func (data *MSG_U2WS_threadmod) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_threadmod)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_threadmod(data, buf)
}

func WRITE_MSG_U2WS_threadmod(data *MSG_U2WS_threadmod, buf *libraries.MsgBuffer) {
	WRITE_int32(int32(len(data.Tids)), buf)
	for _, v := range data.Tids{
		WRITE_int32(v, buf)
	}
	WRITE_int32(data.Expiration, buf)
	WRITE_int8(data.Action, buf)
	WRITE_int8(data.Param, buf)
	WRITE_int32(data.Param1, buf)
	WRITE_string(data.Reason, buf)
	WRITE_int8(data.Sendreasonpm, buf)
}

func READ_MSG_U2WS_threadmod(buf *libraries.MsgBuffer) (data *MSG_U2WS_threadmod) {
	data = pool_MSG_U2WS_threadmod.Get().(*MSG_U2WS_threadmod)
	Tids_len := int(READ_int32(buf))
	for i := 0; i < Tids_len; i++ {
		data.Tids = append(data.Tids, READ_int32(buf))
	}
	data.Expiration = READ_int32(buf)
	data.Action = READ_int8(buf)
	data.Param = READ_int8(buf)
	data.Param1 = READ_int32(buf)
	data.Reason = READ_string(buf)
	data.Sendreasonpm = READ_int8(buf)
	return
}

type MSG_U2WS_viewthreadmod struct {
	Tid int32
}

var pool_MSG_U2WS_viewthreadmod = sync.Pool{New: func() interface{} { return &MSG_U2WS_viewthreadmod{} }}

func GET_MSG_U2WS_viewthreadmod() *MSG_U2WS_viewthreadmod {
	return pool_MSG_U2WS_viewthreadmod.Get().(*MSG_U2WS_viewthreadmod)
}

func (data *MSG_U2WS_viewthreadmod) Put() {
	data.Tid = 0
	pool_MSG_U2WS_viewthreadmod.Put(data)
}
func (data *MSG_U2WS_viewthreadmod) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_viewthreadmod)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_viewthreadmod(data, buf)
}

func WRITE_MSG_U2WS_viewthreadmod(data *MSG_U2WS_viewthreadmod, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Tid, buf)
}

func READ_MSG_U2WS_viewthreadmod(buf *libraries.MsgBuffer) (data *MSG_U2WS_viewthreadmod) {
	data = pool_MSG_U2WS_viewthreadmod.Get().(*MSG_U2WS_viewthreadmod)
	data.Tid = READ_int32(buf)
	return
}

type MSG_WS2U_viewthreadmod struct {
	Param int32
	List []*MSG_threadmod
}

var pool_MSG_WS2U_viewthreadmod = sync.Pool{New: func() interface{} { return &MSG_WS2U_viewthreadmod{} }}

func GET_MSG_WS2U_viewthreadmod() *MSG_WS2U_viewthreadmod {
	return pool_MSG_WS2U_viewthreadmod.Get().(*MSG_WS2U_viewthreadmod)
}

func (data *MSG_WS2U_viewthreadmod) Put() {
	data.Param = 0
	for _,v := range data.List {
		v.Put()
	}
	data.List = data.List[:0]
	pool_MSG_WS2U_viewthreadmod.Put(data)
}
func (data *MSG_WS2U_viewthreadmod) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_viewthreadmod)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_viewthreadmod(data, buf)
}

func WRITE_MSG_WS2U_viewthreadmod(data *MSG_WS2U_viewthreadmod, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Param, buf)
	WRITE_int32(int32(len(data.List)), buf)
	for _, v := range data.List{
		WRITE_MSG_threadmod(v, buf)
	}
}

func READ_MSG_WS2U_viewthreadmod(buf *libraries.MsgBuffer) (data *MSG_WS2U_viewthreadmod) {
	data = pool_MSG_WS2U_viewthreadmod.Get().(*MSG_WS2U_viewthreadmod)
	data.Param = READ_int32(buf)
	List_len := int(READ_int32(buf))
	for i := 0; i < List_len; i++ {
		data.List = append(data.List, READ_MSG_threadmod(buf))
	}
	return
}

type MSG_U2WS_forum_refresh struct {
	Fid int32
	Lastpost int32
}

var pool_MSG_U2WS_forum_refresh = sync.Pool{New: func() interface{} { return &MSG_U2WS_forum_refresh{} }}

func GET_MSG_U2WS_forum_refresh() *MSG_U2WS_forum_refresh {
	return pool_MSG_U2WS_forum_refresh.Get().(*MSG_U2WS_forum_refresh)
}

func (data *MSG_U2WS_forum_refresh) Put() {
	data.Fid = 0
	data.Lastpost = 0
	pool_MSG_U2WS_forum_refresh.Put(data)
}
func (data *MSG_U2WS_forum_refresh) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_forum_refresh)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_forum_refresh(data, buf)
}

func WRITE_MSG_U2WS_forum_refresh(data *MSG_U2WS_forum_refresh, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Fid, buf)
	WRITE_int32(data.Lastpost, buf)
}

func READ_MSG_U2WS_forum_refresh(buf *libraries.MsgBuffer) (data *MSG_U2WS_forum_refresh) {
	data = pool_MSG_U2WS_forum_refresh.Get().(*MSG_U2WS_forum_refresh)
	data.Fid = READ_int32(buf)
	data.Lastpost = READ_int32(buf)
	return
}

type MSG_U2WS_forum_carlist struct {
}

var pool_MSG_U2WS_forum_carlist = sync.Pool{New: func() interface{} { return &MSG_U2WS_forum_carlist{} }}

func GET_MSG_U2WS_forum_carlist() *MSG_U2WS_forum_carlist {
	return pool_MSG_U2WS_forum_carlist.Get().(*MSG_U2WS_forum_carlist)
}

func (data *MSG_U2WS_forum_carlist) Put() {
	pool_MSG_U2WS_forum_carlist.Put(data)
}
func (data *MSG_U2WS_forum_carlist) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_forum_carlist)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_forum_carlist(data, buf)
}

func WRITE_MSG_U2WS_forum_carlist(data *MSG_U2WS_forum_carlist, buf *libraries.MsgBuffer) {
}

func READ_MSG_U2WS_forum_carlist(buf *libraries.MsgBuffer) (data *MSG_U2WS_forum_carlist) {
	data = pool_MSG_U2WS_forum_carlist.Get().(*MSG_U2WS_forum_carlist)
	return
}

type MSG_WS2U_forum_carlist struct {
	Catlist []*MSG_forum_cart
}

var pool_MSG_WS2U_forum_carlist = sync.Pool{New: func() interface{} { return &MSG_WS2U_forum_carlist{} }}

func GET_MSG_WS2U_forum_carlist() *MSG_WS2U_forum_carlist {
	return pool_MSG_WS2U_forum_carlist.Get().(*MSG_WS2U_forum_carlist)
}

func (data *MSG_WS2U_forum_carlist) Put() {
	for _,v := range data.Catlist {
		v.Put()
	}
	data.Catlist = data.Catlist[:0]
	pool_MSG_WS2U_forum_carlist.Put(data)
}
func (data *MSG_WS2U_forum_carlist) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_forum_carlist)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_forum_carlist(data, buf)
}

func WRITE_MSG_WS2U_forum_carlist(data *MSG_WS2U_forum_carlist, buf *libraries.MsgBuffer) {
	WRITE_int32(int32(len(data.Catlist)), buf)
	for _, v := range data.Catlist{
		WRITE_MSG_forum_cart(v, buf)
	}
}

func READ_MSG_WS2U_forum_carlist(buf *libraries.MsgBuffer) (data *MSG_WS2U_forum_carlist) {
	data = pool_MSG_WS2U_forum_carlist.Get().(*MSG_WS2U_forum_carlist)
	Catlist_len := int(READ_int32(buf))
	for i := 0; i < Catlist_len; i++ {
		data.Catlist = append(data.Catlist, READ_MSG_forum_cart(buf))
	}
	return
}

type MSG_forum_cart struct {
	Name string
	Catid int32
	Forums []*MSG_forum_cart_child
}

var pool_MSG_forum_cart = sync.Pool{New: func() interface{} { return &MSG_forum_cart{} }}

func GET_MSG_forum_cart() *MSG_forum_cart {
	return pool_MSG_forum_cart.Get().(*MSG_forum_cart)
}

func (data *MSG_forum_cart) Put() {
	data.Name = ``
	data.Catid = 0
	for _,v := range data.Forums {
		v.Put()
	}
	data.Forums = data.Forums[:0]
	pool_MSG_forum_cart.Put(data)
}
func (data *MSG_forum_cart) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_forum_cart)
	WRITE_int32(cmd, buf)
	WRITE_MSG_forum_cart(data, buf)
}

func WRITE_MSG_forum_cart(data *MSG_forum_cart, buf *libraries.MsgBuffer) {
	WRITE_string(data.Name, buf)
	WRITE_int32(data.Catid, buf)
	WRITE_int32(int32(len(data.Forums)), buf)
	for _, v := range data.Forums{
		WRITE_MSG_forum_cart_child(v, buf)
	}
}

func READ_MSG_forum_cart(buf *libraries.MsgBuffer) (data *MSG_forum_cart) {
	data = pool_MSG_forum_cart.Get().(*MSG_forum_cart)
	data.Name = READ_string(buf)
	data.Catid = READ_int32(buf)
	Forums_len := int(READ_int32(buf))
	for i := 0; i < Forums_len; i++ {
		data.Forums = append(data.Forums, READ_MSG_forum_cart_child(buf))
	}
	return
}

type MSG_forum_cart_child struct {
	Fid int32
	Fup int32
	Name string
	Threadtypes *MSG_forum_threadtype
}

var pool_MSG_forum_cart_child = sync.Pool{New: func() interface{} { return &MSG_forum_cart_child{} }}

func GET_MSG_forum_cart_child() *MSG_forum_cart_child {
	return pool_MSG_forum_cart_child.Get().(*MSG_forum_cart_child)
}

func (data *MSG_forum_cart_child) Put() {
	data.Fid = 0
	data.Fup = 0
	data.Name = ``
	if data.Threadtypes != nil {
		data.Threadtypes.Put()
		data.Threadtypes = nil
	}
	pool_MSG_forum_cart_child.Put(data)
}
func (data *MSG_forum_cart_child) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_forum_cart_child)
	WRITE_int32(cmd, buf)
	WRITE_MSG_forum_cart_child(data, buf)
}

func WRITE_MSG_forum_cart_child(data *MSG_forum_cart_child, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Fid, buf)
	WRITE_int32(data.Fup, buf)
	WRITE_string(data.Name, buf)
	if data.Threadtypes == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_forum_threadtype(data.Threadtypes, buf)
	}
}

func READ_MSG_forum_cart_child(buf *libraries.MsgBuffer) (data *MSG_forum_cart_child) {
	data = pool_MSG_forum_cart_child.Get().(*MSG_forum_cart_child)
	data.Fid = READ_int32(buf)
	data.Fup = READ_int32(buf)
	data.Name = READ_string(buf)
	Threadtypes_len := int(READ_int8(buf))
	if Threadtypes_len == 1 {
		data.Threadtypes = READ_MSG_forum_threadtype(buf)
	}else{
		data.Threadtypes = nil
	}
	return
}

type MSG_U2WS_GetPostWarnList struct {
	Tid int32
	Position int32
}

var pool_MSG_U2WS_GetPostWarnList = sync.Pool{New: func() interface{} { return &MSG_U2WS_GetPostWarnList{} }}

func GET_MSG_U2WS_GetPostWarnList() *MSG_U2WS_GetPostWarnList {
	return pool_MSG_U2WS_GetPostWarnList.Get().(*MSG_U2WS_GetPostWarnList)
}

func (data *MSG_U2WS_GetPostWarnList) Put() {
	data.Tid = 0
	data.Position = 0
	pool_MSG_U2WS_GetPostWarnList.Put(data)
}
func (data *MSG_U2WS_GetPostWarnList) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_GetPostWarnList)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_GetPostWarnList(data, buf)
}

func WRITE_MSG_U2WS_GetPostWarnList(data *MSG_U2WS_GetPostWarnList, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Tid, buf)
	WRITE_int32(data.Position, buf)
}

func READ_MSG_U2WS_GetPostWarnList(buf *libraries.MsgBuffer) (data *MSG_U2WS_GetPostWarnList) {
	data = pool_MSG_U2WS_GetPostWarnList.Get().(*MSG_U2WS_GetPostWarnList)
	data.Tid = READ_int32(buf)
	data.Position = READ_int32(buf)
	return
}

type MSG_U2WS_space struct {
	Uid int32
	Name string
}

var pool_MSG_U2WS_space = sync.Pool{New: func() interface{} { return &MSG_U2WS_space{} }}

func GET_MSG_U2WS_space() *MSG_U2WS_space {
	return pool_MSG_U2WS_space.Get().(*MSG_U2WS_space)
}

func (data *MSG_U2WS_space) Put() {
	data.Uid = 0
	data.Name = ``
	pool_MSG_U2WS_space.Put(data)
}
func (data *MSG_U2WS_space) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_space)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_space(data, buf)
}

func WRITE_MSG_U2WS_space(data *MSG_U2WS_space, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Uid, buf)
	WRITE_string(data.Name, buf)
}

func READ_MSG_U2WS_space(buf *libraries.MsgBuffer) (data *MSG_U2WS_space) {
	data = pool_MSG_U2WS_space.Get().(*MSG_U2WS_space)
	data.Uid = READ_int32(buf)
	data.Name = READ_string(buf)
	return
}

type MSG_WS2U_space struct {
	Status int32
	Username string
	Avatar string
	Uid int32
	Views int32
	Email string
	Customstatus string
	Sightml string
	Posts int32
	Threads int32
	Profiles []*MSG_userprofiles
	Admingroup *MSG_usergroup
	Group *MSG_usergroup
	Upgradecredit int32
	Credits int32
	Groupexpiry int32
	Oltime int32
	Regdate int32
	Lastvisit int32
	Regip string
	Lastip string
	Lastpost int32
	Lastsendmail int32
	Attachsize int32
}

var pool_MSG_WS2U_space = sync.Pool{New: func() interface{} { return &MSG_WS2U_space{} }}

func GET_MSG_WS2U_space() *MSG_WS2U_space {
	return pool_MSG_WS2U_space.Get().(*MSG_WS2U_space)
}

func (data *MSG_WS2U_space) Put() {
	data.Status = 0
	data.Username = ``
	data.Avatar = ``
	data.Uid = 0
	data.Views = 0
	data.Email = ``
	data.Customstatus = ``
	data.Sightml = ``
	data.Posts = 0
	data.Threads = 0
	for _,v := range data.Profiles {
		v.Put()
	}
	data.Profiles = data.Profiles[:0]
	if data.Admingroup != nil {
		data.Admingroup.Put()
		data.Admingroup = nil
	}
	if data.Group != nil {
		data.Group.Put()
		data.Group = nil
	}
	data.Upgradecredit = 0
	data.Credits = 0
	data.Groupexpiry = 0
	data.Oltime = 0
	data.Regdate = 0
	data.Lastvisit = 0
	data.Regip = ``
	data.Lastip = ``
	data.Lastpost = 0
	data.Lastsendmail = 0
	data.Attachsize = 0
	pool_MSG_WS2U_space.Put(data)
}
func (data *MSG_WS2U_space) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_space)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_space(data, buf)
}

func WRITE_MSG_WS2U_space(data *MSG_WS2U_space, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Status, buf)
	WRITE_string(data.Username, buf)
	WRITE_string(data.Avatar, buf)
	WRITE_int32(data.Uid, buf)
	WRITE_int32(data.Views, buf)
	WRITE_string(data.Email, buf)
	WRITE_string(data.Customstatus, buf)
	WRITE_string(data.Sightml, buf)
	WRITE_int32(data.Posts, buf)
	WRITE_int32(data.Threads, buf)
	WRITE_int32(int32(len(data.Profiles)), buf)
	for _, v := range data.Profiles{
		WRITE_MSG_userprofiles(v, buf)
	}
	if data.Admingroup == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_usergroup(data.Admingroup, buf)
	}
	if data.Group == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_usergroup(data.Group, buf)
	}
	WRITE_int32(data.Upgradecredit, buf)
	WRITE_int32(data.Credits, buf)
	WRITE_int32(data.Groupexpiry, buf)
	WRITE_int32(data.Oltime, buf)
	WRITE_int32(data.Regdate, buf)
	WRITE_int32(data.Lastvisit, buf)
	WRITE_string(data.Regip, buf)
	WRITE_string(data.Lastip, buf)
	WRITE_int32(data.Lastpost, buf)
	WRITE_int32(data.Lastsendmail, buf)
	WRITE_int32(data.Attachsize, buf)
}

func READ_MSG_WS2U_space(buf *libraries.MsgBuffer) (data *MSG_WS2U_space) {
	data = pool_MSG_WS2U_space.Get().(*MSG_WS2U_space)
	data.Status = READ_int32(buf)
	data.Username = READ_string(buf)
	data.Avatar = READ_string(buf)
	data.Uid = READ_int32(buf)
	data.Views = READ_int32(buf)
	data.Email = READ_string(buf)
	data.Customstatus = READ_string(buf)
	data.Sightml = READ_string(buf)
	data.Posts = READ_int32(buf)
	data.Threads = READ_int32(buf)
	Profiles_len := int(READ_int32(buf))
	for i := 0; i < Profiles_len; i++ {
		data.Profiles = append(data.Profiles, READ_MSG_userprofiles(buf))
	}
	Admingroup_len := int(READ_int8(buf))
	if Admingroup_len == 1 {
		data.Admingroup = READ_MSG_usergroup(buf)
	}else{
		data.Admingroup = nil
	}
	Group_len := int(READ_int8(buf))
	if Group_len == 1 {
		data.Group = READ_MSG_usergroup(buf)
	}else{
		data.Group = nil
	}
	data.Upgradecredit = READ_int32(buf)
	data.Credits = READ_int32(buf)
	data.Groupexpiry = READ_int32(buf)
	data.Oltime = READ_int32(buf)
	data.Regdate = READ_int32(buf)
	data.Lastvisit = READ_int32(buf)
	data.Regip = READ_string(buf)
	data.Lastip = READ_string(buf)
	data.Lastpost = READ_int32(buf)
	data.Lastsendmail = READ_int32(buf)
	data.Attachsize = READ_int32(buf)
	return
}

type MSG_userprofiles struct {
	Title string
	Value string
}

var pool_MSG_userprofiles = sync.Pool{New: func() interface{} { return &MSG_userprofiles{} }}

func GET_MSG_userprofiles() *MSG_userprofiles {
	return pool_MSG_userprofiles.Get().(*MSG_userprofiles)
}

func (data *MSG_userprofiles) Put() {
	data.Title = ``
	data.Value = ``
	pool_MSG_userprofiles.Put(data)
}
func (data *MSG_userprofiles) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_userprofiles)
	WRITE_int32(cmd, buf)
	WRITE_MSG_userprofiles(data, buf)
}

func WRITE_MSG_userprofiles(data *MSG_userprofiles, buf *libraries.MsgBuffer) {
	WRITE_string(data.Title, buf)
	WRITE_string(data.Value, buf)
}

func READ_MSG_userprofiles(buf *libraries.MsgBuffer) (data *MSG_userprofiles) {
	data = pool_MSG_userprofiles.Get().(*MSG_userprofiles)
	data.Title = READ_string(buf)
	data.Value = READ_string(buf)
	return
}

type MSG_usergroup struct {
	Id int16
	Color string
	Grouptitle string
	Icon string
}

var pool_MSG_usergroup = sync.Pool{New: func() interface{} { return &MSG_usergroup{} }}

func GET_MSG_usergroup() *MSG_usergroup {
	return pool_MSG_usergroup.Get().(*MSG_usergroup)
}

func (data *MSG_usergroup) Put() {
	data.Id = 0
	data.Color = ``
	data.Grouptitle = ``
	data.Icon = ``
	pool_MSG_usergroup.Put(data)
}
func (data *MSG_usergroup) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_usergroup)
	WRITE_int32(cmd, buf)
	WRITE_MSG_usergroup(data, buf)
}

func WRITE_MSG_usergroup(data *MSG_usergroup, buf *libraries.MsgBuffer) {
	WRITE_int16(data.Id, buf)
	WRITE_string(data.Color, buf)
	WRITE_string(data.Grouptitle, buf)
	WRITE_string(data.Icon, buf)
}

func READ_MSG_usergroup(buf *libraries.MsgBuffer) (data *MSG_usergroup) {
	data = pool_MSG_usergroup.Get().(*MSG_usergroup)
	data.Id = READ_int16(buf)
	data.Color = READ_string(buf)
	data.Grouptitle = READ_string(buf)
	data.Icon = READ_string(buf)
	return
}

type MSG_U2WS_SpaceThread struct {
	Uid int32
	Type int8
	Page int16
}

var pool_MSG_U2WS_SpaceThread = sync.Pool{New: func() interface{} { return &MSG_U2WS_SpaceThread{} }}

func GET_MSG_U2WS_SpaceThread() *MSG_U2WS_SpaceThread {
	return pool_MSG_U2WS_SpaceThread.Get().(*MSG_U2WS_SpaceThread)
}

func (data *MSG_U2WS_SpaceThread) Put() {
	data.Uid = 0
	data.Type = 0
	data.Page = 0
	pool_MSG_U2WS_SpaceThread.Put(data)
}
func (data *MSG_U2WS_SpaceThread) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_SpaceThread)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_SpaceThread(data, buf)
}

func WRITE_MSG_U2WS_SpaceThread(data *MSG_U2WS_SpaceThread, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Uid, buf)
	WRITE_int8(data.Type, buf)
	WRITE_int16(data.Page, buf)
}

func READ_MSG_U2WS_SpaceThread(buf *libraries.MsgBuffer) (data *MSG_U2WS_SpaceThread) {
	data = pool_MSG_U2WS_SpaceThread.Get().(*MSG_U2WS_SpaceThread)
	data.Uid = READ_int32(buf)
	data.Type = READ_int8(buf)
	data.Page = READ_int16(buf)
	return
}

type MSG_WS2U_SpaceThread struct {
	Uid int32
	Threadlist []*MSG_SpaceThread
	Threadcount int32
}

var pool_MSG_WS2U_SpaceThread = sync.Pool{New: func() interface{} { return &MSG_WS2U_SpaceThread{} }}

func GET_MSG_WS2U_SpaceThread() *MSG_WS2U_SpaceThread {
	return pool_MSG_WS2U_SpaceThread.Get().(*MSG_WS2U_SpaceThread)
}

func (data *MSG_WS2U_SpaceThread) Put() {
	data.Uid = 0
	for _,v := range data.Threadlist {
		v.Put()
	}
	data.Threadlist = data.Threadlist[:0]
	data.Threadcount = 0
	pool_MSG_WS2U_SpaceThread.Put(data)
}
func (data *MSG_WS2U_SpaceThread) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_SpaceThread)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_SpaceThread(data, buf)
}

func WRITE_MSG_WS2U_SpaceThread(data *MSG_WS2U_SpaceThread, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Uid, buf)
	WRITE_int32(int32(len(data.Threadlist)), buf)
	for _, v := range data.Threadlist{
		WRITE_MSG_SpaceThread(v, buf)
	}
	WRITE_int32(data.Threadcount, buf)
}

func READ_MSG_WS2U_SpaceThread(buf *libraries.MsgBuffer) (data *MSG_WS2U_SpaceThread) {
	data = pool_MSG_WS2U_SpaceThread.Get().(*MSG_WS2U_SpaceThread)
	data.Uid = READ_int32(buf)
	Threadlist_len := int(READ_int32(buf))
	for i := 0; i < Threadlist_len; i++ {
		data.Threadlist = append(data.Threadlist, READ_MSG_SpaceThread(buf))
	}
	data.Threadcount = READ_int32(buf)
	return
}

type MSG_SpaceThread struct {
	Tid int32
	Folder string
	Special int8
	Displayorder int8
	Subject string
	Digest int8
	Attachment int8
	Multipage int8
	Closed int8
	Fid int32
	ForumName string
	Authorid int32
	Author string
	Dateline int32
	Replies int32
	Views int32
	Lastposter string
	Lastpost int32
	Postlist []*MSG_SpacePost
}

var pool_MSG_SpaceThread = sync.Pool{New: func() interface{} { return &MSG_SpaceThread{} }}

func GET_MSG_SpaceThread() *MSG_SpaceThread {
	return pool_MSG_SpaceThread.Get().(*MSG_SpaceThread)
}

func (data *MSG_SpaceThread) Put() {
	data.Tid = 0
	data.Folder = ``
	data.Special = 0
	data.Displayorder = 0
	data.Subject = ``
	data.Digest = 0
	data.Attachment = 0
	data.Multipage = 0
	data.Closed = 0
	data.Fid = 0
	data.ForumName = ``
	data.Authorid = 0
	data.Author = ``
	data.Dateline = 0
	data.Replies = 0
	data.Views = 0
	data.Lastposter = ``
	data.Lastpost = 0
	for _,v := range data.Postlist {
		v.Put()
	}
	data.Postlist = data.Postlist[:0]
	pool_MSG_SpaceThread.Put(data)
}
func (data *MSG_SpaceThread) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_SpaceThread)
	WRITE_int32(cmd, buf)
	WRITE_MSG_SpaceThread(data, buf)
}

func WRITE_MSG_SpaceThread(data *MSG_SpaceThread, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Tid, buf)
	WRITE_string(data.Folder, buf)
	WRITE_int8(data.Special, buf)
	WRITE_int8(data.Displayorder, buf)
	WRITE_string(data.Subject, buf)
	WRITE_int8(data.Digest, buf)
	WRITE_int8(data.Attachment, buf)
	WRITE_int8(data.Multipage, buf)
	WRITE_int8(data.Closed, buf)
	WRITE_int32(data.Fid, buf)
	WRITE_string(data.ForumName, buf)
	WRITE_int32(data.Authorid, buf)
	WRITE_string(data.Author, buf)
	WRITE_int32(data.Dateline, buf)
	WRITE_int32(data.Replies, buf)
	WRITE_int32(data.Views, buf)
	WRITE_string(data.Lastposter, buf)
	WRITE_int32(data.Lastpost, buf)
	WRITE_int32(int32(len(data.Postlist)), buf)
	for _, v := range data.Postlist{
		WRITE_MSG_SpacePost(v, buf)
	}
}

func READ_MSG_SpaceThread(buf *libraries.MsgBuffer) (data *MSG_SpaceThread) {
	data = pool_MSG_SpaceThread.Get().(*MSG_SpaceThread)
	data.Tid = READ_int32(buf)
	data.Folder = READ_string(buf)
	data.Special = READ_int8(buf)
	data.Displayorder = READ_int8(buf)
	data.Subject = READ_string(buf)
	data.Digest = READ_int8(buf)
	data.Attachment = READ_int8(buf)
	data.Multipage = READ_int8(buf)
	data.Closed = READ_int8(buf)
	data.Fid = READ_int32(buf)
	data.ForumName = READ_string(buf)
	data.Authorid = READ_int32(buf)
	data.Author = READ_string(buf)
	data.Dateline = READ_int32(buf)
	data.Replies = READ_int32(buf)
	data.Views = READ_int32(buf)
	data.Lastposter = READ_string(buf)
	data.Lastpost = READ_int32(buf)
	Postlist_len := int(READ_int32(buf))
	for i := 0; i < Postlist_len; i++ {
		data.Postlist = append(data.Postlist, READ_MSG_SpacePost(buf))
	}
	return
}

type MSG_SpacePost struct {
	Position int32
	Message string
}

var pool_MSG_SpacePost = sync.Pool{New: func() interface{} { return &MSG_SpacePost{} }}

func GET_MSG_SpacePost() *MSG_SpacePost {
	return pool_MSG_SpacePost.Get().(*MSG_SpacePost)
}

func (data *MSG_SpacePost) Put() {
	data.Position = 0
	data.Message = ``
	pool_MSG_SpacePost.Put(data)
}
func (data *MSG_SpacePost) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_SpacePost)
	WRITE_int32(cmd, buf)
	WRITE_MSG_SpacePost(data, buf)
}

func WRITE_MSG_SpacePost(data *MSG_SpacePost, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Position, buf)
	WRITE_string(data.Message, buf)
}

func READ_MSG_SpacePost(buf *libraries.MsgBuffer) (data *MSG_SpacePost) {
	data = pool_MSG_SpacePost.Get().(*MSG_SpacePost)
	data.Position = READ_int32(buf)
	data.Message = READ_string(buf)
	return
}

type MSG_U2WS_searchThread struct {
	Word string
	Page int16
}

var pool_MSG_U2WS_searchThread = sync.Pool{New: func() interface{} { return &MSG_U2WS_searchThread{} }}

func GET_MSG_U2WS_searchThread() *MSG_U2WS_searchThread {
	return pool_MSG_U2WS_searchThread.Get().(*MSG_U2WS_searchThread)
}

func (data *MSG_U2WS_searchThread) Put() {
	data.Word = ``
	data.Page = 0
	pool_MSG_U2WS_searchThread.Put(data)
}
func (data *MSG_U2WS_searchThread) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_searchThread)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_searchThread(data, buf)
}

func WRITE_MSG_U2WS_searchThread(data *MSG_U2WS_searchThread, buf *libraries.MsgBuffer) {
	WRITE_string(data.Word, buf)
	WRITE_int16(data.Page, buf)
}

func READ_MSG_U2WS_searchThread(buf *libraries.MsgBuffer) (data *MSG_U2WS_searchThread) {
	data = pool_MSG_U2WS_searchThread.Get().(*MSG_U2WS_searchThread)
	data.Word = READ_string(buf)
	data.Page = READ_int16(buf)
	return
}

type MSG_WS2U_searchThread struct {
	Threadlist []*MSG_searchThread
	Threadcount int32
	Time int64
}

var pool_MSG_WS2U_searchThread = sync.Pool{New: func() interface{} { return &MSG_WS2U_searchThread{} }}

func GET_MSG_WS2U_searchThread() *MSG_WS2U_searchThread {
	return pool_MSG_WS2U_searchThread.Get().(*MSG_WS2U_searchThread)
}

func (data *MSG_WS2U_searchThread) Put() {
	for _,v := range data.Threadlist {
		v.Put()
	}
	data.Threadlist = data.Threadlist[:0]
	data.Threadcount = 0
	data.Time = 0
	pool_MSG_WS2U_searchThread.Put(data)
}
func (data *MSG_WS2U_searchThread) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_searchThread)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_searchThread(data, buf)
}

func WRITE_MSG_WS2U_searchThread(data *MSG_WS2U_searchThread, buf *libraries.MsgBuffer) {
	WRITE_int32(int32(len(data.Threadlist)), buf)
	for _, v := range data.Threadlist{
		WRITE_MSG_searchThread(v, buf)
	}
	WRITE_int32(data.Threadcount, buf)
	WRITE_int64(data.Time, buf)
}

func READ_MSG_WS2U_searchThread(buf *libraries.MsgBuffer) (data *MSG_WS2U_searchThread) {
	data = pool_MSG_WS2U_searchThread.Get().(*MSG_WS2U_searchThread)
	Threadlist_len := int(READ_int32(buf))
	for i := 0; i < Threadlist_len; i++ {
		data.Threadlist = append(data.Threadlist, READ_MSG_searchThread(buf))
	}
	data.Threadcount = READ_int32(buf)
	data.Time = READ_int64(buf)
	return
}

type MSG_searchThread struct {
	Tid int32
	Fid int32
	Subject string
	Replies int32
	Views int32
	Authorid int32
	Author string
	Post string
	ForumName string
	Cutmessage string
	Totalpage int16
	Dateline int32
}

var pool_MSG_searchThread = sync.Pool{New: func() interface{} { return &MSG_searchThread{} }}

func GET_MSG_searchThread() *MSG_searchThread {
	return pool_MSG_searchThread.Get().(*MSG_searchThread)
}

func (data *MSG_searchThread) Put() {
	data.Tid = 0
	data.Fid = 0
	data.Subject = ``
	data.Replies = 0
	data.Views = 0
	data.Authorid = 0
	data.Author = ``
	data.Post = ``
	data.ForumName = ``
	data.Cutmessage = ``
	data.Totalpage = 0
	data.Dateline = 0
	pool_MSG_searchThread.Put(data)
}
func (data *MSG_searchThread) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_searchThread)
	WRITE_int32(cmd, buf)
	WRITE_MSG_searchThread(data, buf)
}

func WRITE_MSG_searchThread(data *MSG_searchThread, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Tid, buf)
	WRITE_int32(data.Fid, buf)
	WRITE_string(data.Subject, buf)
	WRITE_int32(data.Replies, buf)
	WRITE_int32(data.Views, buf)
	WRITE_int32(data.Authorid, buf)
	WRITE_string(data.Author, buf)
	WRITE_string(data.Post, buf)
	WRITE_string(data.ForumName, buf)
	WRITE_string(data.Cutmessage, buf)
	WRITE_int16(data.Totalpage, buf)
	WRITE_int32(data.Dateline, buf)
}

func READ_MSG_searchThread(buf *libraries.MsgBuffer) (data *MSG_searchThread) {
	data = pool_MSG_searchThread.Get().(*MSG_searchThread)
	data.Tid = READ_int32(buf)
	data.Fid = READ_int32(buf)
	data.Subject = READ_string(buf)
	data.Replies = READ_int32(buf)
	data.Views = READ_int32(buf)
	data.Authorid = READ_int32(buf)
	data.Author = READ_string(buf)
	data.Post = READ_string(buf)
	data.ForumName = READ_string(buf)
	data.Cutmessage = READ_string(buf)
	data.Totalpage = READ_int16(buf)
	data.Dateline = READ_int32(buf)
	return
}

type MSG_WS2U_threadmod struct {
	Result int16
}

var pool_MSG_WS2U_threadmod = sync.Pool{New: func() interface{} { return &MSG_WS2U_threadmod{} }}

func GET_MSG_WS2U_threadmod() *MSG_WS2U_threadmod {
	return pool_MSG_WS2U_threadmod.Get().(*MSG_WS2U_threadmod)
}

func (data *MSG_WS2U_threadmod) Put() {
	data.Result = 0
	pool_MSG_WS2U_threadmod.Put(data)
}
func (data *MSG_WS2U_threadmod) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_threadmod)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_threadmod(data, buf)
}

func WRITE_MSG_WS2U_threadmod(data *MSG_WS2U_threadmod, buf *libraries.MsgBuffer) {
	WRITE_int16(data.Result, buf)
}

func READ_MSG_WS2U_threadmod(buf *libraries.MsgBuffer) (data *MSG_WS2U_threadmod) {
	data = pool_MSG_WS2U_threadmod.Get().(*MSG_WS2U_threadmod)
	data.Result = READ_int16(buf)
	return
}

type MSG_U2WS_spacecp struct {
	Token string
}

var pool_MSG_U2WS_spacecp = sync.Pool{New: func() interface{} { return &MSG_U2WS_spacecp{} }}

func GET_MSG_U2WS_spacecp() *MSG_U2WS_spacecp {
	return pool_MSG_U2WS_spacecp.Get().(*MSG_U2WS_spacecp)
}

func (data *MSG_U2WS_spacecp) Put() {
	data.Token = ``
	pool_MSG_U2WS_spacecp.Put(data)
}
func (data *MSG_U2WS_spacecp) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_spacecp)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_spacecp(data, buf)
}

func WRITE_MSG_U2WS_spacecp(data *MSG_U2WS_spacecp, buf *libraries.MsgBuffer) {
	WRITE_string(data.Token, buf)
}

func READ_MSG_U2WS_spacecp(buf *libraries.MsgBuffer) (data *MSG_U2WS_spacecp) {
	data = pool_MSG_U2WS_spacecp.Get().(*MSG_U2WS_spacecp)
	data.Token = READ_string(buf)
	return
}

type MSG_WS2U_spacecp struct {
	Uid int32
	Name string
	GroupId int16
	Allow int8
	Customstatus string
	Mobile string
	Qq int64
	Wx string
	WebSite string
	Sightml string
	GroupTitle string
	SiteGroups []*MSG_GroupIdName
	CommonGroups []*MSG_GroupIdName
	Email string
	EmailNew string
}

var pool_MSG_WS2U_spacecp = sync.Pool{New: func() interface{} { return &MSG_WS2U_spacecp{} }}

func GET_MSG_WS2U_spacecp() *MSG_WS2U_spacecp {
	return pool_MSG_WS2U_spacecp.Get().(*MSG_WS2U_spacecp)
}

func (data *MSG_WS2U_spacecp) Put() {
	data.Uid = 0
	data.Name = ``
	data.GroupId = 0
	data.Allow = 0
	data.Customstatus = ``
	data.Mobile = ``
	data.Qq = 0
	data.Wx = ``
	data.WebSite = ``
	data.Sightml = ``
	data.GroupTitle = ``
	for _,v := range data.SiteGroups {
		v.Put()
	}
	data.SiteGroups = data.SiteGroups[:0]
	for _,v := range data.CommonGroups {
		v.Put()
	}
	data.CommonGroups = data.CommonGroups[:0]
	data.Email = ``
	data.EmailNew = ``
	pool_MSG_WS2U_spacecp.Put(data)
}
func (data *MSG_WS2U_spacecp) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_spacecp)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_spacecp(data, buf)
}

func WRITE_MSG_WS2U_spacecp(data *MSG_WS2U_spacecp, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Uid, buf)
	WRITE_string(data.Name, buf)
	WRITE_int16(data.GroupId, buf)
	WRITE_int8(data.Allow, buf)
	WRITE_string(data.Customstatus, buf)
	WRITE_string(data.Mobile, buf)
	WRITE_int64(data.Qq, buf)
	WRITE_string(data.Wx, buf)
	WRITE_string(data.WebSite, buf)
	WRITE_string(data.Sightml, buf)
	WRITE_string(data.GroupTitle, buf)
	WRITE_int32(int32(len(data.SiteGroups)), buf)
	for _, v := range data.SiteGroups{
		WRITE_MSG_GroupIdName(v, buf)
	}
	WRITE_int32(int32(len(data.CommonGroups)), buf)
	for _, v := range data.CommonGroups{
		WRITE_MSG_GroupIdName(v, buf)
	}
	WRITE_string(data.Email, buf)
	WRITE_string(data.EmailNew, buf)
}

func READ_MSG_WS2U_spacecp(buf *libraries.MsgBuffer) (data *MSG_WS2U_spacecp) {
	data = pool_MSG_WS2U_spacecp.Get().(*MSG_WS2U_spacecp)
	data.Uid = READ_int32(buf)
	data.Name = READ_string(buf)
	data.GroupId = READ_int16(buf)
	data.Allow = READ_int8(buf)
	data.Customstatus = READ_string(buf)
	data.Mobile = READ_string(buf)
	data.Qq = READ_int64(buf)
	data.Wx = READ_string(buf)
	data.WebSite = READ_string(buf)
	data.Sightml = READ_string(buf)
	data.GroupTitle = READ_string(buf)
	SiteGroups_len := int(READ_int32(buf))
	for i := 0; i < SiteGroups_len; i++ {
		data.SiteGroups = append(data.SiteGroups, READ_MSG_GroupIdName(buf))
	}
	CommonGroups_len := int(READ_int32(buf))
	for i := 0; i < CommonGroups_len; i++ {
		data.CommonGroups = append(data.CommonGroups, READ_MSG_GroupIdName(buf))
	}
	data.Email = READ_string(buf)
	data.EmailNew = READ_string(buf)
	return
}

type MSG_U2WS_tpl_success struct {
}

var pool_MSG_U2WS_tpl_success = sync.Pool{New: func() interface{} { return &MSG_U2WS_tpl_success{} }}

func GET_MSG_U2WS_tpl_success() *MSG_U2WS_tpl_success {
	return pool_MSG_U2WS_tpl_success.Get().(*MSG_U2WS_tpl_success)
}

func (data *MSG_U2WS_tpl_success) Put() {
	pool_MSG_U2WS_tpl_success.Put(data)
}
func (data *MSG_U2WS_tpl_success) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_tpl_success)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_tpl_success(data, buf)
}

func WRITE_MSG_U2WS_tpl_success(data *MSG_U2WS_tpl_success, buf *libraries.MsgBuffer) {
}

func READ_MSG_U2WS_tpl_success(buf *libraries.MsgBuffer) (data *MSG_U2WS_tpl_success) {
	data = pool_MSG_U2WS_tpl_success.Get().(*MSG_U2WS_tpl_success)
	return
}

type MSG_WS2U_tpl_success struct {
}

var pool_MSG_WS2U_tpl_success = sync.Pool{New: func() interface{} { return &MSG_WS2U_tpl_success{} }}

func GET_MSG_WS2U_tpl_success() *MSG_WS2U_tpl_success {
	return pool_MSG_WS2U_tpl_success.Get().(*MSG_WS2U_tpl_success)
}

func (data *MSG_WS2U_tpl_success) Put() {
	pool_MSG_WS2U_tpl_success.Put(data)
}
func (data *MSG_WS2U_tpl_success) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_tpl_success)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_tpl_success(data, buf)
}

func WRITE_MSG_WS2U_tpl_success(data *MSG_WS2U_tpl_success, buf *libraries.MsgBuffer) {
}

func READ_MSG_WS2U_tpl_success(buf *libraries.MsgBuffer) (data *MSG_WS2U_tpl_success) {
	data = pool_MSG_WS2U_tpl_success.Get().(*MSG_WS2U_tpl_success)
	return
}

type MSG_U2WS_upload_avatar struct {
	Imgdata []byte
}

var pool_MSG_U2WS_upload_avatar = sync.Pool{New: func() interface{} { return &MSG_U2WS_upload_avatar{} }}

func GET_MSG_U2WS_upload_avatar() *MSG_U2WS_upload_avatar {
	return pool_MSG_U2WS_upload_avatar.Get().(*MSG_U2WS_upload_avatar)
}

func (data *MSG_U2WS_upload_avatar) Put() {
	data.Imgdata = data.Imgdata[:0]
	pool_MSG_U2WS_upload_avatar.Put(data)
}
func (data *MSG_U2WS_upload_avatar) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_upload_avatar)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_upload_avatar(data, buf)
}

func WRITE_MSG_U2WS_upload_avatar(data *MSG_U2WS_upload_avatar, buf *libraries.MsgBuffer) {
	WRITE_int32(int32(len(data.Imgdata)), buf)
	buf.Write(data.Imgdata)
}

func READ_MSG_U2WS_upload_avatar(buf *libraries.MsgBuffer) (data *MSG_U2WS_upload_avatar) {
	data = pool_MSG_U2WS_upload_avatar.Get().(*MSG_U2WS_upload_avatar)
	Imgdata_len := int(READ_int32(buf))
	data.Imgdata = make([]byte, Imgdata_len)
	copy(data.Imgdata,buf.Next(Imgdata_len))
	return
}

type MSG_WS2U_upload_avatar struct {
	Result int16
	Avatar string
}

var pool_MSG_WS2U_upload_avatar = sync.Pool{New: func() interface{} { return &MSG_WS2U_upload_avatar{} }}

func GET_MSG_WS2U_upload_avatar() *MSG_WS2U_upload_avatar {
	return pool_MSG_WS2U_upload_avatar.Get().(*MSG_WS2U_upload_avatar)
}

func (data *MSG_WS2U_upload_avatar) Put() {
	data.Result = 0
	data.Avatar = ``
	pool_MSG_WS2U_upload_avatar.Put(data)
}
func (data *MSG_WS2U_upload_avatar) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_upload_avatar)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_upload_avatar(data, buf)
}

func WRITE_MSG_WS2U_upload_avatar(data *MSG_WS2U_upload_avatar, buf *libraries.MsgBuffer) {
	WRITE_int16(data.Result, buf)
	WRITE_string(data.Avatar, buf)
}

func READ_MSG_WS2U_upload_avatar(buf *libraries.MsgBuffer) (data *MSG_WS2U_upload_avatar) {
	data = pool_MSG_WS2U_upload_avatar.Get().(*MSG_WS2U_upload_avatar)
	data.Result = READ_int16(buf)
	data.Avatar = READ_string(buf)
	return
}

type MSG_U2WS_Edit_Profile struct {
	Qq int64
	Wx string
	Mobile string
	WebSite string
	Sightml string
	Customstatus string
}

var pool_MSG_U2WS_Edit_Profile = sync.Pool{New: func() interface{} { return &MSG_U2WS_Edit_Profile{} }}

func GET_MSG_U2WS_Edit_Profile() *MSG_U2WS_Edit_Profile {
	return pool_MSG_U2WS_Edit_Profile.Get().(*MSG_U2WS_Edit_Profile)
}

func (data *MSG_U2WS_Edit_Profile) Put() {
	data.Qq = 0
	data.Wx = ``
	data.Mobile = ``
	data.WebSite = ``
	data.Sightml = ``
	data.Customstatus = ``
	pool_MSG_U2WS_Edit_Profile.Put(data)
}
func (data *MSG_U2WS_Edit_Profile) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_Edit_Profile)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_Edit_Profile(data, buf)
}

func WRITE_MSG_U2WS_Edit_Profile(data *MSG_U2WS_Edit_Profile, buf *libraries.MsgBuffer) {
	WRITE_int64(data.Qq, buf)
	WRITE_string(data.Wx, buf)
	WRITE_string(data.Mobile, buf)
	WRITE_string(data.WebSite, buf)
	WRITE_string(data.Sightml, buf)
	WRITE_string(data.Customstatus, buf)
}

func READ_MSG_U2WS_Edit_Profile(buf *libraries.MsgBuffer) (data *MSG_U2WS_Edit_Profile) {
	data = pool_MSG_U2WS_Edit_Profile.Get().(*MSG_U2WS_Edit_Profile)
	data.Qq = READ_int64(buf)
	data.Wx = READ_string(buf)
	data.Mobile = READ_string(buf)
	data.WebSite = READ_string(buf)
	data.Sightml = READ_string(buf)
	data.Customstatus = READ_string(buf)
	return
}

type MSG_U2WS_RecommendThread struct {
	Tid int32
	Status int8
}

var pool_MSG_U2WS_RecommendThread = sync.Pool{New: func() interface{} { return &MSG_U2WS_RecommendThread{} }}

func GET_MSG_U2WS_RecommendThread() *MSG_U2WS_RecommendThread {
	return pool_MSG_U2WS_RecommendThread.Get().(*MSG_U2WS_RecommendThread)
}

func (data *MSG_U2WS_RecommendThread) Put() {
	data.Tid = 0
	data.Status = 0
	pool_MSG_U2WS_RecommendThread.Put(data)
}
func (data *MSG_U2WS_RecommendThread) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_RecommendThread)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_RecommendThread(data, buf)
}

func WRITE_MSG_U2WS_RecommendThread(data *MSG_U2WS_RecommendThread, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Tid, buf)
	WRITE_int8(data.Status, buf)
}

func READ_MSG_U2WS_RecommendThread(buf *libraries.MsgBuffer) (data *MSG_U2WS_RecommendThread) {
	data = pool_MSG_U2WS_RecommendThread.Get().(*MSG_U2WS_RecommendThread)
	data.Tid = READ_int32(buf)
	data.Status = READ_int8(buf)
	return
}

type MSG_WS2U_RecommendThread struct {
	Status int8
	Result int16
	Recommend int32
	Recommend_add int32
	Recommend_sub int32
}

var pool_MSG_WS2U_RecommendThread = sync.Pool{New: func() interface{} { return &MSG_WS2U_RecommendThread{} }}

func GET_MSG_WS2U_RecommendThread() *MSG_WS2U_RecommendThread {
	return pool_MSG_WS2U_RecommendThread.Get().(*MSG_WS2U_RecommendThread)
}

func (data *MSG_WS2U_RecommendThread) Put() {
	data.Status = 0
	data.Result = 0
	data.Recommend = 0
	data.Recommend_add = 0
	data.Recommend_sub = 0
	pool_MSG_WS2U_RecommendThread.Put(data)
}
func (data *MSG_WS2U_RecommendThread) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_RecommendThread)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_RecommendThread(data, buf)
}

func WRITE_MSG_WS2U_RecommendThread(data *MSG_WS2U_RecommendThread, buf *libraries.MsgBuffer) {
	WRITE_int8(data.Status, buf)
	WRITE_int16(data.Result, buf)
	WRITE_int32(data.Recommend, buf)
	WRITE_int32(data.Recommend_add, buf)
	WRITE_int32(data.Recommend_sub, buf)
}

func READ_MSG_WS2U_RecommendThread(buf *libraries.MsgBuffer) (data *MSG_WS2U_RecommendThread) {
	data = pool_MSG_WS2U_RecommendThread.Get().(*MSG_WS2U_RecommendThread)
	data.Status = READ_int8(buf)
	data.Result = READ_int16(buf)
	data.Recommend = READ_int32(buf)
	data.Recommend_add = READ_int32(buf)
	data.Recommend_sub = READ_int32(buf)
	return
}

type MSG_U2WS_SpacecpGroup struct {
	Groupid int16
}

var pool_MSG_U2WS_SpacecpGroup = sync.Pool{New: func() interface{} { return &MSG_U2WS_SpacecpGroup{} }}

func GET_MSG_U2WS_SpacecpGroup() *MSG_U2WS_SpacecpGroup {
	return pool_MSG_U2WS_SpacecpGroup.Get().(*MSG_U2WS_SpacecpGroup)
}

func (data *MSG_U2WS_SpacecpGroup) Put() {
	data.Groupid = 0
	pool_MSG_U2WS_SpacecpGroup.Put(data)
}
func (data *MSG_U2WS_SpacecpGroup) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_SpacecpGroup)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_SpacecpGroup(data, buf)
}

func WRITE_MSG_U2WS_SpacecpGroup(data *MSG_U2WS_SpacecpGroup, buf *libraries.MsgBuffer) {
	WRITE_int16(data.Groupid, buf)
}

func READ_MSG_U2WS_SpacecpGroup(buf *libraries.MsgBuffer) (data *MSG_U2WS_SpacecpGroup) {
	data = pool_MSG_U2WS_SpacecpGroup.Get().(*MSG_U2WS_SpacecpGroup)
	data.Groupid = READ_int16(buf)
	return
}

type MSG_GroupIdName struct {
	Name string
	Id int16
}

var pool_MSG_GroupIdName = sync.Pool{New: func() interface{} { return &MSG_GroupIdName{} }}

func GET_MSG_GroupIdName() *MSG_GroupIdName {
	return pool_MSG_GroupIdName.Get().(*MSG_GroupIdName)
}

func (data *MSG_GroupIdName) Put() {
	data.Name = ``
	data.Id = 0
	pool_MSG_GroupIdName.Put(data)
}
func (data *MSG_GroupIdName) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_GroupIdName)
	WRITE_int32(cmd, buf)
	WRITE_MSG_GroupIdName(data, buf)
}

func WRITE_MSG_GroupIdName(data *MSG_GroupIdName, buf *libraries.MsgBuffer) {
	WRITE_string(data.Name, buf)
	WRITE_int16(data.Id, buf)
}

func READ_MSG_GroupIdName(buf *libraries.MsgBuffer) (data *MSG_GroupIdName) {
	data = pool_MSG_GroupIdName.Get().(*MSG_GroupIdName)
	data.Name = READ_string(buf)
	data.Id = READ_int16(buf)
	return
}

type MSG_WS2U_SpacecpGroup struct {
	Groupid int16
	Grouptitle string
	Readaccess int8
	Allow int16
	Maxsigsize int16
	Allowrecommend int8
	Maxattachsize int32
	Maxsizeperday int32
	Maxattachnum int16
	Attachextensions string
}

var pool_MSG_WS2U_SpacecpGroup = sync.Pool{New: func() interface{} { return &MSG_WS2U_SpacecpGroup{} }}

func GET_MSG_WS2U_SpacecpGroup() *MSG_WS2U_SpacecpGroup {
	return pool_MSG_WS2U_SpacecpGroup.Get().(*MSG_WS2U_SpacecpGroup)
}

func (data *MSG_WS2U_SpacecpGroup) Put() {
	data.Groupid = 0
	data.Grouptitle = ``
	data.Readaccess = 0
	data.Allow = 0
	data.Maxsigsize = 0
	data.Allowrecommend = 0
	data.Maxattachsize = 0
	data.Maxsizeperday = 0
	data.Maxattachnum = 0
	data.Attachextensions = ``
	pool_MSG_WS2U_SpacecpGroup.Put(data)
}
func (data *MSG_WS2U_SpacecpGroup) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_SpacecpGroup)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_SpacecpGroup(data, buf)
}

func WRITE_MSG_WS2U_SpacecpGroup(data *MSG_WS2U_SpacecpGroup, buf *libraries.MsgBuffer) {
	WRITE_int16(data.Groupid, buf)
	WRITE_string(data.Grouptitle, buf)
	WRITE_int8(data.Readaccess, buf)
	WRITE_int16(data.Allow, buf)
	WRITE_int16(data.Maxsigsize, buf)
	WRITE_int8(data.Allowrecommend, buf)
	WRITE_int32(data.Maxattachsize, buf)
	WRITE_int32(data.Maxsizeperday, buf)
	WRITE_int16(data.Maxattachnum, buf)
	WRITE_string(data.Attachextensions, buf)
}

func READ_MSG_WS2U_SpacecpGroup(buf *libraries.MsgBuffer) (data *MSG_WS2U_SpacecpGroup) {
	data = pool_MSG_WS2U_SpacecpGroup.Get().(*MSG_WS2U_SpacecpGroup)
	data.Groupid = READ_int16(buf)
	data.Grouptitle = READ_string(buf)
	data.Readaccess = READ_int8(buf)
	data.Allow = READ_int16(buf)
	data.Maxsigsize = READ_int16(buf)
	data.Allowrecommend = READ_int8(buf)
	data.Maxattachsize = READ_int32(buf)
	data.Maxsizeperday = READ_int32(buf)
	data.Maxattachnum = READ_int16(buf)
	data.Attachextensions = READ_string(buf)
	return
}

type MSG_U2WS_SpacecpForum struct {
}

var pool_MSG_U2WS_SpacecpForum = sync.Pool{New: func() interface{} { return &MSG_U2WS_SpacecpForum{} }}

func GET_MSG_U2WS_SpacecpForum() *MSG_U2WS_SpacecpForum {
	return pool_MSG_U2WS_SpacecpForum.Get().(*MSG_U2WS_SpacecpForum)
}

func (data *MSG_U2WS_SpacecpForum) Put() {
	pool_MSG_U2WS_SpacecpForum.Put(data)
}
func (data *MSG_U2WS_SpacecpForum) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_SpacecpForum)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_SpacecpForum(data, buf)
}

func WRITE_MSG_U2WS_SpacecpForum(data *MSG_U2WS_SpacecpForum, buf *libraries.MsgBuffer) {
}

func READ_MSG_U2WS_SpacecpForum(buf *libraries.MsgBuffer) (data *MSG_U2WS_SpacecpForum) {
	data = pool_MSG_U2WS_SpacecpForum.Get().(*MSG_U2WS_SpacecpForum)
	return
}

type MSG_WS2U_SpacecpForum struct {
	Catlist []*MSG_SpacecpGroupPermission
}

var pool_MSG_WS2U_SpacecpForum = sync.Pool{New: func() interface{} { return &MSG_WS2U_SpacecpForum{} }}

func GET_MSG_WS2U_SpacecpForum() *MSG_WS2U_SpacecpForum {
	return pool_MSG_WS2U_SpacecpForum.Get().(*MSG_WS2U_SpacecpForum)
}

func (data *MSG_WS2U_SpacecpForum) Put() {
	for _,v := range data.Catlist {
		v.Put()
	}
	data.Catlist = data.Catlist[:0]
	pool_MSG_WS2U_SpacecpForum.Put(data)
}
func (data *MSG_WS2U_SpacecpForum) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_SpacecpForum)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_SpacecpForum(data, buf)
}

func WRITE_MSG_WS2U_SpacecpForum(data *MSG_WS2U_SpacecpForum, buf *libraries.MsgBuffer) {
	WRITE_int32(int32(len(data.Catlist)), buf)
	for _, v := range data.Catlist{
		WRITE_MSG_SpacecpGroupPermission(v, buf)
	}
}

func READ_MSG_WS2U_SpacecpForum(buf *libraries.MsgBuffer) (data *MSG_WS2U_SpacecpForum) {
	data = pool_MSG_WS2U_SpacecpForum.Get().(*MSG_WS2U_SpacecpForum)
	Catlist_len := int(READ_int32(buf))
	for i := 0; i < Catlist_len; i++ {
		data.Catlist = append(data.Catlist, READ_MSG_SpacecpGroupPermission(buf))
	}
	return
}

type MSG_SpacecpGroupPermission struct {
	Fid int32
	Name string
	Child []*MSG_SpacecpGroupPermission
	Allow int8
}

var pool_MSG_SpacecpGroupPermission = sync.Pool{New: func() interface{} { return &MSG_SpacecpGroupPermission{} }}

func GET_MSG_SpacecpGroupPermission() *MSG_SpacecpGroupPermission {
	return pool_MSG_SpacecpGroupPermission.Get().(*MSG_SpacecpGroupPermission)
}

func (data *MSG_SpacecpGroupPermission) Put() {
	data.Fid = 0
	data.Name = ``
	for _,v := range data.Child {
		v.Put()
	}
	data.Child = data.Child[:0]
	data.Allow = 0
	pool_MSG_SpacecpGroupPermission.Put(data)
}
func (data *MSG_SpacecpGroupPermission) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_SpacecpGroupPermission)
	WRITE_int32(cmd, buf)
	WRITE_MSG_SpacecpGroupPermission(data, buf)
}

func WRITE_MSG_SpacecpGroupPermission(data *MSG_SpacecpGroupPermission, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Fid, buf)
	WRITE_string(data.Name, buf)
	WRITE_int32(int32(len(data.Child)), buf)
	for _, v := range data.Child{
		WRITE_MSG_SpacecpGroupPermission(v, buf)
	}
	WRITE_int8(data.Allow, buf)
}

func READ_MSG_SpacecpGroupPermission(buf *libraries.MsgBuffer) (data *MSG_SpacecpGroupPermission) {
	data = pool_MSG_SpacecpGroupPermission.Get().(*MSG_SpacecpGroupPermission)
	data.Fid = READ_int32(buf)
	data.Name = READ_string(buf)
	Child_len := int(READ_int32(buf))
	for i := 0; i < Child_len; i++ {
		data.Child = append(data.Child, READ_MSG_SpacecpGroupPermission(buf))
	}
	data.Allow = READ_int8(buf)
	return
}

type MSG_U2WS_ChangePasswd_Gethash struct {
	Seccode string
}

var pool_MSG_U2WS_ChangePasswd_Gethash = sync.Pool{New: func() interface{} { return &MSG_U2WS_ChangePasswd_Gethash{} }}

func GET_MSG_U2WS_ChangePasswd_Gethash() *MSG_U2WS_ChangePasswd_Gethash {
	return pool_MSG_U2WS_ChangePasswd_Gethash.Get().(*MSG_U2WS_ChangePasswd_Gethash)
}

func (data *MSG_U2WS_ChangePasswd_Gethash) Put() {
	data.Seccode = ``
	pool_MSG_U2WS_ChangePasswd_Gethash.Put(data)
}
func (data *MSG_U2WS_ChangePasswd_Gethash) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_ChangePasswd_Gethash)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_ChangePasswd_Gethash(data, buf)
}

func WRITE_MSG_U2WS_ChangePasswd_Gethash(data *MSG_U2WS_ChangePasswd_Gethash, buf *libraries.MsgBuffer) {
	WRITE_string(data.Seccode, buf)
}

func READ_MSG_U2WS_ChangePasswd_Gethash(buf *libraries.MsgBuffer) (data *MSG_U2WS_ChangePasswd_Gethash) {
	data = pool_MSG_U2WS_ChangePasswd_Gethash.Get().(*MSG_U2WS_ChangePasswd_Gethash)
	data.Seccode = READ_string(buf)
	return
}

type MSG_WS2U_ChangePasswd_Gethash struct {
	Hash string
	Hash2 string
	Hash3 string
}

var pool_MSG_WS2U_ChangePasswd_Gethash = sync.Pool{New: func() interface{} { return &MSG_WS2U_ChangePasswd_Gethash{} }}

func GET_MSG_WS2U_ChangePasswd_Gethash() *MSG_WS2U_ChangePasswd_Gethash {
	return pool_MSG_WS2U_ChangePasswd_Gethash.Get().(*MSG_WS2U_ChangePasswd_Gethash)
}

func (data *MSG_WS2U_ChangePasswd_Gethash) Put() {
	data.Hash = ``
	data.Hash2 = ``
	data.Hash3 = ``
	pool_MSG_WS2U_ChangePasswd_Gethash.Put(data)
}
func (data *MSG_WS2U_ChangePasswd_Gethash) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_ChangePasswd_Gethash)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_ChangePasswd_Gethash(data, buf)
}

func WRITE_MSG_WS2U_ChangePasswd_Gethash(data *MSG_WS2U_ChangePasswd_Gethash, buf *libraries.MsgBuffer) {
	WRITE_string(data.Hash, buf)
	WRITE_string(data.Hash2, buf)
	WRITE_string(data.Hash3, buf)
}

func READ_MSG_WS2U_ChangePasswd_Gethash(buf *libraries.MsgBuffer) (data *MSG_WS2U_ChangePasswd_Gethash) {
	data = pool_MSG_WS2U_ChangePasswd_Gethash.Get().(*MSG_WS2U_ChangePasswd_Gethash)
	data.Hash = READ_string(buf)
	data.Hash2 = READ_string(buf)
	data.Hash3 = READ_string(buf)
	return
}

type MSG_U2WS_ChangePasswd struct {
	Passwd string
	Newpwd []byte
	Email string
}

var pool_MSG_U2WS_ChangePasswd = sync.Pool{New: func() interface{} { return &MSG_U2WS_ChangePasswd{} }}

func GET_MSG_U2WS_ChangePasswd() *MSG_U2WS_ChangePasswd {
	return pool_MSG_U2WS_ChangePasswd.Get().(*MSG_U2WS_ChangePasswd)
}

func (data *MSG_U2WS_ChangePasswd) Put() {
	data.Passwd = ``
	data.Newpwd = data.Newpwd[:0]
	data.Email = ``
	pool_MSG_U2WS_ChangePasswd.Put(data)
}
func (data *MSG_U2WS_ChangePasswd) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_ChangePasswd)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_ChangePasswd(data, buf)
}

func WRITE_MSG_U2WS_ChangePasswd(data *MSG_U2WS_ChangePasswd, buf *libraries.MsgBuffer) {
	WRITE_string(data.Passwd, buf)
	WRITE_int32(int32(len(data.Newpwd)), buf)
	buf.Write(data.Newpwd)
	WRITE_string(data.Email, buf)
}

func READ_MSG_U2WS_ChangePasswd(buf *libraries.MsgBuffer) (data *MSG_U2WS_ChangePasswd) {
	data = pool_MSG_U2WS_ChangePasswd.Get().(*MSG_U2WS_ChangePasswd)
	data.Passwd = READ_string(buf)
	Newpwd_len := int(READ_int32(buf))
	data.Newpwd = make([]byte, Newpwd_len)
	copy(data.Newpwd,buf.Next(Newpwd_len))
	data.Email = READ_string(buf)
	return
}

type MSG_U2WS_Email_Verify struct {
	Uid int32
	Code string
}

var pool_MSG_U2WS_Email_Verify = sync.Pool{New: func() interface{} { return &MSG_U2WS_Email_Verify{} }}

func GET_MSG_U2WS_Email_Verify() *MSG_U2WS_Email_Verify {
	return pool_MSG_U2WS_Email_Verify.Get().(*MSG_U2WS_Email_Verify)
}

func (data *MSG_U2WS_Email_Verify) Put() {
	data.Uid = 0
	data.Code = ``
	pool_MSG_U2WS_Email_Verify.Put(data)
}
func (data *MSG_U2WS_Email_Verify) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_Email_Verify)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_Email_Verify(data, buf)
}

func WRITE_MSG_U2WS_Email_Verify(data *MSG_U2WS_Email_Verify, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Uid, buf)
	WRITE_string(data.Code, buf)
}

func READ_MSG_U2WS_Email_Verify(buf *libraries.MsgBuffer) (data *MSG_U2WS_Email_Verify) {
	data = pool_MSG_U2WS_Email_Verify.Get().(*MSG_U2WS_Email_Verify)
	data.Uid = READ_int32(buf)
	data.Code = READ_string(buf)
	return
}

type MSG_WS2U_Email_Verify struct {
	Result int16
}

var pool_MSG_WS2U_Email_Verify = sync.Pool{New: func() interface{} { return &MSG_WS2U_Email_Verify{} }}

func GET_MSG_WS2U_Email_Verify() *MSG_WS2U_Email_Verify {
	return pool_MSG_WS2U_Email_Verify.Get().(*MSG_WS2U_Email_Verify)
}

func (data *MSG_WS2U_Email_Verify) Put() {
	data.Result = 0
	pool_MSG_WS2U_Email_Verify.Put(data)
}
func (data *MSG_WS2U_Email_Verify) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_Email_Verify)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_Email_Verify(data, buf)
}

func WRITE_MSG_WS2U_Email_Verify(data *MSG_WS2U_Email_Verify, buf *libraries.MsgBuffer) {
	WRITE_int16(data.Result, buf)
}

func READ_MSG_WS2U_Email_Verify(buf *libraries.MsgBuffer) (data *MSG_WS2U_Email_Verify) {
	data = pool_MSG_WS2U_Email_Verify.Get().(*MSG_WS2U_Email_Verify)
	data.Result = READ_int16(buf)
	return
}

type MSG_U2WS_LostPW struct {
	Name string
	Seccode string
}

var pool_MSG_U2WS_LostPW = sync.Pool{New: func() interface{} { return &MSG_U2WS_LostPW{} }}

func GET_MSG_U2WS_LostPW() *MSG_U2WS_LostPW {
	return pool_MSG_U2WS_LostPW.Get().(*MSG_U2WS_LostPW)
}

func (data *MSG_U2WS_LostPW) Put() {
	data.Name = ``
	data.Seccode = ``
	pool_MSG_U2WS_LostPW.Put(data)
}
func (data *MSG_U2WS_LostPW) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_LostPW)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_LostPW(data, buf)
}

func WRITE_MSG_U2WS_LostPW(data *MSG_U2WS_LostPW, buf *libraries.MsgBuffer) {
	WRITE_string(data.Name, buf)
	WRITE_string(data.Seccode, buf)
}

func READ_MSG_U2WS_LostPW(buf *libraries.MsgBuffer) (data *MSG_U2WS_LostPW) {
	data = pool_MSG_U2WS_LostPW.Get().(*MSG_U2WS_LostPW)
	data.Name = READ_string(buf)
	data.Seccode = READ_string(buf)
	return
}

type MSG_WS2U_LostPW struct {
	Result int16
	Email string
}

var pool_MSG_WS2U_LostPW = sync.Pool{New: func() interface{} { return &MSG_WS2U_LostPW{} }}

func GET_MSG_WS2U_LostPW() *MSG_WS2U_LostPW {
	return pool_MSG_WS2U_LostPW.Get().(*MSG_WS2U_LostPW)
}

func (data *MSG_WS2U_LostPW) Put() {
	data.Result = 0
	data.Email = ``
	pool_MSG_WS2U_LostPW.Put(data)
}
func (data *MSG_WS2U_LostPW) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_LostPW)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_LostPW(data, buf)
}

func WRITE_MSG_WS2U_LostPW(data *MSG_WS2U_LostPW, buf *libraries.MsgBuffer) {
	WRITE_int16(data.Result, buf)
	WRITE_string(data.Email, buf)
}

func READ_MSG_WS2U_LostPW(buf *libraries.MsgBuffer) (data *MSG_WS2U_LostPW) {
	data = pool_MSG_WS2U_LostPW.Get().(*MSG_WS2U_LostPW)
	data.Result = READ_int16(buf)
	data.Email = READ_string(buf)
	return
}

type MSG_U2WS_ResetPW struct {
	Name string
	Passwd []byte
}

var pool_MSG_U2WS_ResetPW = sync.Pool{New: func() interface{} { return &MSG_U2WS_ResetPW{} }}

func GET_MSG_U2WS_ResetPW() *MSG_U2WS_ResetPW {
	return pool_MSG_U2WS_ResetPW.Get().(*MSG_U2WS_ResetPW)
}

func (data *MSG_U2WS_ResetPW) Put() {
	data.Name = ``
	data.Passwd = data.Passwd[:0]
	pool_MSG_U2WS_ResetPW.Put(data)
}
func (data *MSG_U2WS_ResetPW) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_ResetPW)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_ResetPW(data, buf)
}

func WRITE_MSG_U2WS_ResetPW(data *MSG_U2WS_ResetPW, buf *libraries.MsgBuffer) {
	WRITE_string(data.Name, buf)
	WRITE_int32(int32(len(data.Passwd)), buf)
	buf.Write(data.Passwd)
}

func READ_MSG_U2WS_ResetPW(buf *libraries.MsgBuffer) (data *MSG_U2WS_ResetPW) {
	data = pool_MSG_U2WS_ResetPW.Get().(*MSG_U2WS_ResetPW)
	data.Name = READ_string(buf)
	Passwd_len := int(READ_int32(buf))
	data.Passwd = make([]byte, Passwd_len)
	copy(data.Passwd,buf.Next(Passwd_len))
	return
}

type MSG_WS2U_ResetPW struct {
}

var pool_MSG_WS2U_ResetPW = sync.Pool{New: func() interface{} { return &MSG_WS2U_ResetPW{} }}

func GET_MSG_WS2U_ResetPW() *MSG_WS2U_ResetPW {
	return pool_MSG_WS2U_ResetPW.Get().(*MSG_WS2U_ResetPW)
}

func (data *MSG_WS2U_ResetPW) Put() {
	pool_MSG_WS2U_ResetPW.Put(data)
}
func (data *MSG_WS2U_ResetPW) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_ResetPW)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_ResetPW(data, buf)
}

func WRITE_MSG_WS2U_ResetPW(data *MSG_WS2U_ResetPW, buf *libraries.MsgBuffer) {
}

func READ_MSG_WS2U_ResetPW(buf *libraries.MsgBuffer) (data *MSG_WS2U_ResetPW) {
	data = pool_MSG_WS2U_ResetPW.Get().(*MSG_WS2U_ResetPW)
	return
}

type MSG_U2WS_QQLoginUrl struct {
}

var pool_MSG_U2WS_QQLoginUrl = sync.Pool{New: func() interface{} { return &MSG_U2WS_QQLoginUrl{} }}

func GET_MSG_U2WS_QQLoginUrl() *MSG_U2WS_QQLoginUrl {
	return pool_MSG_U2WS_QQLoginUrl.Get().(*MSG_U2WS_QQLoginUrl)
}

func (data *MSG_U2WS_QQLoginUrl) Put() {
	pool_MSG_U2WS_QQLoginUrl.Put(data)
}
func (data *MSG_U2WS_QQLoginUrl) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_QQLoginUrl)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_QQLoginUrl(data, buf)
}

func WRITE_MSG_U2WS_QQLoginUrl(data *MSG_U2WS_QQLoginUrl, buf *libraries.MsgBuffer) {
}

func READ_MSG_U2WS_QQLoginUrl(buf *libraries.MsgBuffer) (data *MSG_U2WS_QQLoginUrl) {
	data = pool_MSG_U2WS_QQLoginUrl.Get().(*MSG_U2WS_QQLoginUrl)
	return
}

type MSG_WS2U_QQLoginUrl struct {
	Url string
}

var pool_MSG_WS2U_QQLoginUrl = sync.Pool{New: func() interface{} { return &MSG_WS2U_QQLoginUrl{} }}

func GET_MSG_WS2U_QQLoginUrl() *MSG_WS2U_QQLoginUrl {
	return pool_MSG_WS2U_QQLoginUrl.Get().(*MSG_WS2U_QQLoginUrl)
}

func (data *MSG_WS2U_QQLoginUrl) Put() {
	data.Url = ``
	pool_MSG_WS2U_QQLoginUrl.Put(data)
}
func (data *MSG_WS2U_QQLoginUrl) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_QQLoginUrl)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_QQLoginUrl(data, buf)
}

func WRITE_MSG_WS2U_QQLoginUrl(data *MSG_WS2U_QQLoginUrl, buf *libraries.MsgBuffer) {
	WRITE_string(data.Url, buf)
}

func READ_MSG_WS2U_QQLoginUrl(buf *libraries.MsgBuffer) (data *MSG_WS2U_QQLoginUrl) {
	data = pool_MSG_WS2U_QQLoginUrl.Get().(*MSG_WS2U_QQLoginUrl)
	data.Url = READ_string(buf)
	return
}

type MSG_U2WS_QQLogin struct {
	Openid string
	State string
}

var pool_MSG_U2WS_QQLogin = sync.Pool{New: func() interface{} { return &MSG_U2WS_QQLogin{} }}

func GET_MSG_U2WS_QQLogin() *MSG_U2WS_QQLogin {
	return pool_MSG_U2WS_QQLogin.Get().(*MSG_U2WS_QQLogin)
}

func (data *MSG_U2WS_QQLogin) Put() {
	data.Openid = ``
	data.State = ``
	pool_MSG_U2WS_QQLogin.Put(data)
}
func (data *MSG_U2WS_QQLogin) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_QQLogin)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_QQLogin(data, buf)
}

func WRITE_MSG_U2WS_QQLogin(data *MSG_U2WS_QQLogin, buf *libraries.MsgBuffer) {
	WRITE_string(data.Openid, buf)
	WRITE_string(data.State, buf)
}

func READ_MSG_U2WS_QQLogin(buf *libraries.MsgBuffer) (data *MSG_U2WS_QQLogin) {
	data = pool_MSG_U2WS_QQLogin.Get().(*MSG_U2WS_QQLogin)
	data.Openid = READ_string(buf)
	data.State = READ_string(buf)
	return
}

type MSG_WS2U_QQLogin struct {
	Result int16
	Head *MSG_WS2U_Common_Head
	Openid string
}

var pool_MSG_WS2U_QQLogin = sync.Pool{New: func() interface{} { return &MSG_WS2U_QQLogin{} }}

func GET_MSG_WS2U_QQLogin() *MSG_WS2U_QQLogin {
	return pool_MSG_WS2U_QQLogin.Get().(*MSG_WS2U_QQLogin)
}

func (data *MSG_WS2U_QQLogin) Put() {
	data.Result = 0
	if data.Head != nil {
		data.Head.Put()
		data.Head = nil
	}
	data.Openid = ``
	pool_MSG_WS2U_QQLogin.Put(data)
}
func (data *MSG_WS2U_QQLogin) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_QQLogin)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_QQLogin(data, buf)
}

func WRITE_MSG_WS2U_QQLogin(data *MSG_WS2U_QQLogin, buf *libraries.MsgBuffer) {
	WRITE_int16(data.Result, buf)
	if data.Head == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_WS2U_Common_Head(data.Head, buf)
	}
	WRITE_string(data.Openid, buf)
}

func READ_MSG_WS2U_QQLogin(buf *libraries.MsgBuffer) (data *MSG_WS2U_QQLogin) {
	data = pool_MSG_WS2U_QQLogin.Get().(*MSG_WS2U_QQLogin)
	data.Result = READ_int16(buf)
	Head_len := int(READ_int8(buf))
	if Head_len == 1 {
		data.Head = READ_MSG_WS2U_Common_Head(buf)
	}else{
		data.Head = nil
	}
	data.Openid = READ_string(buf)
	return
}

type MSG_U2WS_BindAccount struct {
	Type string
	Openid string
	Access_token string
	State string
}

var pool_MSG_U2WS_BindAccount = sync.Pool{New: func() interface{} { return &MSG_U2WS_BindAccount{} }}

func GET_MSG_U2WS_BindAccount() *MSG_U2WS_BindAccount {
	return pool_MSG_U2WS_BindAccount.Get().(*MSG_U2WS_BindAccount)
}

func (data *MSG_U2WS_BindAccount) Put() {
	data.Type = ``
	data.Openid = ``
	data.Access_token = ``
	data.State = ``
	pool_MSG_U2WS_BindAccount.Put(data)
}
func (data *MSG_U2WS_BindAccount) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_BindAccount)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_BindAccount(data, buf)
}

func WRITE_MSG_U2WS_BindAccount(data *MSG_U2WS_BindAccount, buf *libraries.MsgBuffer) {
	WRITE_string(data.Type, buf)
	WRITE_string(data.Openid, buf)
	WRITE_string(data.Access_token, buf)
	WRITE_string(data.State, buf)
}

func READ_MSG_U2WS_BindAccount(buf *libraries.MsgBuffer) (data *MSG_U2WS_BindAccount) {
	data = pool_MSG_U2WS_BindAccount.Get().(*MSG_U2WS_BindAccount)
	data.Type = READ_string(buf)
	data.Openid = READ_string(buf)
	data.Access_token = READ_string(buf)
	data.State = READ_string(buf)
	return
}

type MSG_WS2U_BindAccount struct {
	Result int16
	Nickname string
	Img string
}

var pool_MSG_WS2U_BindAccount = sync.Pool{New: func() interface{} { return &MSG_WS2U_BindAccount{} }}

func GET_MSG_WS2U_BindAccount() *MSG_WS2U_BindAccount {
	return pool_MSG_WS2U_BindAccount.Get().(*MSG_WS2U_BindAccount)
}

func (data *MSG_WS2U_BindAccount) Put() {
	data.Result = 0
	data.Nickname = ``
	data.Img = ``
	pool_MSG_WS2U_BindAccount.Put(data)
}
func (data *MSG_WS2U_BindAccount) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_BindAccount)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_BindAccount(data, buf)
}

func WRITE_MSG_WS2U_BindAccount(data *MSG_WS2U_BindAccount, buf *libraries.MsgBuffer) {
	WRITE_int16(data.Result, buf)
	WRITE_string(data.Nickname, buf)
	WRITE_string(data.Img, buf)
}

func READ_MSG_WS2U_BindAccount(buf *libraries.MsgBuffer) (data *MSG_WS2U_BindAccount) {
	data = pool_MSG_WS2U_BindAccount.Get().(*MSG_WS2U_BindAccount)
	data.Result = READ_int16(buf)
	data.Nickname = READ_string(buf)
	data.Img = READ_string(buf)
	return
}

type MSG_U2WS_GetThreadBind struct {
}

var pool_MSG_U2WS_GetThreadBind = sync.Pool{New: func() interface{} { return &MSG_U2WS_GetThreadBind{} }}

func GET_MSG_U2WS_GetThreadBind() *MSG_U2WS_GetThreadBind {
	return pool_MSG_U2WS_GetThreadBind.Get().(*MSG_U2WS_GetThreadBind)
}

func (data *MSG_U2WS_GetThreadBind) Put() {
	pool_MSG_U2WS_GetThreadBind.Put(data)
}
func (data *MSG_U2WS_GetThreadBind) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_GetThreadBind)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_GetThreadBind(data, buf)
}

func WRITE_MSG_U2WS_GetThreadBind(data *MSG_U2WS_GetThreadBind, buf *libraries.MsgBuffer) {
}

func READ_MSG_U2WS_GetThreadBind(buf *libraries.MsgBuffer) (data *MSG_U2WS_GetThreadBind) {
	data = pool_MSG_U2WS_GetThreadBind.Get().(*MSG_U2WS_GetThreadBind)
	return
}

type MSG_ThreadBind struct {
	Name string
	Nickname string
	Img string
}

var pool_MSG_ThreadBind = sync.Pool{New: func() interface{} { return &MSG_ThreadBind{} }}

func GET_MSG_ThreadBind() *MSG_ThreadBind {
	return pool_MSG_ThreadBind.Get().(*MSG_ThreadBind)
}

func (data *MSG_ThreadBind) Put() {
	data.Name = ``
	data.Nickname = ``
	data.Img = ``
	pool_MSG_ThreadBind.Put(data)
}
func (data *MSG_ThreadBind) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_ThreadBind)
	WRITE_int32(cmd, buf)
	WRITE_MSG_ThreadBind(data, buf)
}

func WRITE_MSG_ThreadBind(data *MSG_ThreadBind, buf *libraries.MsgBuffer) {
	WRITE_string(data.Name, buf)
	WRITE_string(data.Nickname, buf)
	WRITE_string(data.Img, buf)
}

func READ_MSG_ThreadBind(buf *libraries.MsgBuffer) (data *MSG_ThreadBind) {
	data = pool_MSG_ThreadBind.Get().(*MSG_ThreadBind)
	data.Name = READ_string(buf)
	data.Nickname = READ_string(buf)
	data.Img = READ_string(buf)
	return
}

type MSG_WS2U_GetThreadBind struct {
	Thread []*MSG_ThreadBind
}

var pool_MSG_WS2U_GetThreadBind = sync.Pool{New: func() interface{} { return &MSG_WS2U_GetThreadBind{} }}

func GET_MSG_WS2U_GetThreadBind() *MSG_WS2U_GetThreadBind {
	return pool_MSG_WS2U_GetThreadBind.Get().(*MSG_WS2U_GetThreadBind)
}

func (data *MSG_WS2U_GetThreadBind) Put() {
	for _,v := range data.Thread {
		v.Put()
	}
	data.Thread = data.Thread[:0]
	pool_MSG_WS2U_GetThreadBind.Put(data)
}
func (data *MSG_WS2U_GetThreadBind) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_GetThreadBind)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_GetThreadBind(data, buf)
}

func WRITE_MSG_WS2U_GetThreadBind(data *MSG_WS2U_GetThreadBind, buf *libraries.MsgBuffer) {
	WRITE_int32(int32(len(data.Thread)), buf)
	for _, v := range data.Thread{
		WRITE_MSG_ThreadBind(v, buf)
	}
}

func READ_MSG_WS2U_GetThreadBind(buf *libraries.MsgBuffer) (data *MSG_WS2U_GetThreadBind) {
	data = pool_MSG_WS2U_GetThreadBind.Get().(*MSG_WS2U_GetThreadBind)
	Thread_len := int(READ_int32(buf))
	for i := 0; i < Thread_len; i++ {
		data.Thread = append(data.Thread, READ_MSG_ThreadBind(buf))
	}
	return
}

type MSG_U2WS_GetChangeBindUrl struct {
	Type string
	Passwd string
}

var pool_MSG_U2WS_GetChangeBindUrl = sync.Pool{New: func() interface{} { return &MSG_U2WS_GetChangeBindUrl{} }}

func GET_MSG_U2WS_GetChangeBindUrl() *MSG_U2WS_GetChangeBindUrl {
	return pool_MSG_U2WS_GetChangeBindUrl.Get().(*MSG_U2WS_GetChangeBindUrl)
}

func (data *MSG_U2WS_GetChangeBindUrl) Put() {
	data.Type = ``
	data.Passwd = ``
	pool_MSG_U2WS_GetChangeBindUrl.Put(data)
}
func (data *MSG_U2WS_GetChangeBindUrl) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_GetChangeBindUrl)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_GetChangeBindUrl(data, buf)
}

func WRITE_MSG_U2WS_GetChangeBindUrl(data *MSG_U2WS_GetChangeBindUrl, buf *libraries.MsgBuffer) {
	WRITE_string(data.Type, buf)
	WRITE_string(data.Passwd, buf)
}

func READ_MSG_U2WS_GetChangeBindUrl(buf *libraries.MsgBuffer) (data *MSG_U2WS_GetChangeBindUrl) {
	data = pool_MSG_U2WS_GetChangeBindUrl.Get().(*MSG_U2WS_GetChangeBindUrl)
	data.Type = READ_string(buf)
	data.Passwd = READ_string(buf)
	return
}

type MSG_WS2U_GetChangeBindUrl struct {
	Type string
	Url string
}

var pool_MSG_WS2U_GetChangeBindUrl = sync.Pool{New: func() interface{} { return &MSG_WS2U_GetChangeBindUrl{} }}

func GET_MSG_WS2U_GetChangeBindUrl() *MSG_WS2U_GetChangeBindUrl {
	return pool_MSG_WS2U_GetChangeBindUrl.Get().(*MSG_WS2U_GetChangeBindUrl)
}

func (data *MSG_WS2U_GetChangeBindUrl) Put() {
	data.Type = ``
	data.Url = ``
	pool_MSG_WS2U_GetChangeBindUrl.Put(data)
}
func (data *MSG_WS2U_GetChangeBindUrl) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_GetChangeBindUrl)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_GetChangeBindUrl(data, buf)
}

func WRITE_MSG_WS2U_GetChangeBindUrl(data *MSG_WS2U_GetChangeBindUrl, buf *libraries.MsgBuffer) {
	WRITE_string(data.Type, buf)
	WRITE_string(data.Url, buf)
}

func READ_MSG_WS2U_GetChangeBindUrl(buf *libraries.MsgBuffer) (data *MSG_WS2U_GetChangeBindUrl) {
	data = pool_MSG_WS2U_GetChangeBindUrl.Get().(*MSG_WS2U_GetChangeBindUrl)
	data.Type = READ_string(buf)
	data.Url = READ_string(buf)
	return
}

type MSG_U2WS_ChangeBind struct {
	Type string
	State string
	Openid string
	Access_token string
}

var pool_MSG_U2WS_ChangeBind = sync.Pool{New: func() interface{} { return &MSG_U2WS_ChangeBind{} }}

func GET_MSG_U2WS_ChangeBind() *MSG_U2WS_ChangeBind {
	return pool_MSG_U2WS_ChangeBind.Get().(*MSG_U2WS_ChangeBind)
}

func (data *MSG_U2WS_ChangeBind) Put() {
	data.Type = ``
	data.State = ``
	data.Openid = ``
	data.Access_token = ``
	pool_MSG_U2WS_ChangeBind.Put(data)
}
func (data *MSG_U2WS_ChangeBind) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_ChangeBind)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_ChangeBind(data, buf)
}

func WRITE_MSG_U2WS_ChangeBind(data *MSG_U2WS_ChangeBind, buf *libraries.MsgBuffer) {
	WRITE_string(data.Type, buf)
	WRITE_string(data.State, buf)
	WRITE_string(data.Openid, buf)
	WRITE_string(data.Access_token, buf)
}

func READ_MSG_U2WS_ChangeBind(buf *libraries.MsgBuffer) (data *MSG_U2WS_ChangeBind) {
	data = pool_MSG_U2WS_ChangeBind.Get().(*MSG_U2WS_ChangeBind)
	data.Type = READ_string(buf)
	data.State = READ_string(buf)
	data.Openid = READ_string(buf)
	data.Access_token = READ_string(buf)
	return
}

type MSG_WS2U_ChangeBind struct {
	Result int16
	Msg string
}

var pool_MSG_WS2U_ChangeBind = sync.Pool{New: func() interface{} { return &MSG_WS2U_ChangeBind{} }}

func GET_MSG_WS2U_ChangeBind() *MSG_WS2U_ChangeBind {
	return pool_MSG_WS2U_ChangeBind.Get().(*MSG_WS2U_ChangeBind)
}

func (data *MSG_WS2U_ChangeBind) Put() {
	data.Result = 0
	data.Msg = ``
	pool_MSG_WS2U_ChangeBind.Put(data)
}
func (data *MSG_WS2U_ChangeBind) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_ChangeBind)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_ChangeBind(data, buf)
}

func WRITE_MSG_WS2U_ChangeBind(data *MSG_WS2U_ChangeBind, buf *libraries.MsgBuffer) {
	WRITE_int16(data.Result, buf)
	WRITE_string(data.Msg, buf)
}

func READ_MSG_WS2U_ChangeBind(buf *libraries.MsgBuffer) (data *MSG_WS2U_ChangeBind) {
	data = pool_MSG_WS2U_ChangeBind.Get().(*MSG_WS2U_ChangeBind)
	data.Result = READ_int16(buf)
	data.Msg = READ_string(buf)
	return
}

type MSG_Poll_info struct {
	Polloption []*MSG_poll_option
	Maxchoices int8
	Expiration int32
	Visible int8
	Overt int8
	Voterscount int32
	Isimagepoll int8
	AllreadyPoll int8
}

var pool_MSG_Poll_info = sync.Pool{New: func() interface{} { return &MSG_Poll_info{} }}

func GET_MSG_Poll_info() *MSG_Poll_info {
	return pool_MSG_Poll_info.Get().(*MSG_Poll_info)
}

func (data *MSG_Poll_info) Put() {
	for _,v := range data.Polloption {
		v.Put()
	}
	data.Polloption = data.Polloption[:0]
	data.Maxchoices = 0
	data.Expiration = 0
	data.Visible = 0
	data.Overt = 0
	data.Voterscount = 0
	data.Isimagepoll = 0
	data.AllreadyPoll = 0
	pool_MSG_Poll_info.Put(data)
}
func (data *MSG_Poll_info) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_Poll_info)
	WRITE_int32(cmd, buf)
	WRITE_MSG_Poll_info(data, buf)
}

func WRITE_MSG_Poll_info(data *MSG_Poll_info, buf *libraries.MsgBuffer) {
	WRITE_int32(int32(len(data.Polloption)), buf)
	for _, v := range data.Polloption{
		WRITE_MSG_poll_option(v, buf)
	}
	WRITE_int8(data.Maxchoices, buf)
	WRITE_int32(data.Expiration, buf)
	WRITE_int8(data.Visible, buf)
	WRITE_int8(data.Overt, buf)
	WRITE_int32(data.Voterscount, buf)
	WRITE_int8(data.Isimagepoll, buf)
	WRITE_int8(data.AllreadyPoll, buf)
}

func READ_MSG_Poll_info(buf *libraries.MsgBuffer) (data *MSG_Poll_info) {
	data = pool_MSG_Poll_info.Get().(*MSG_Poll_info)
	Polloption_len := int(READ_int32(buf))
	for i := 0; i < Polloption_len; i++ {
		data.Polloption = append(data.Polloption, READ_MSG_poll_option(buf))
	}
	data.Maxchoices = READ_int8(buf)
	data.Expiration = READ_int32(buf)
	data.Visible = READ_int8(buf)
	data.Overt = READ_int8(buf)
	data.Voterscount = READ_int32(buf)
	data.Isimagepoll = READ_int8(buf)
	data.AllreadyPoll = READ_int8(buf)
	return
}

type MSG_poll_option struct {
	Id int32
	Name string
	Displayorder int8
	Aid int64
	Imginfo *MSG_forum_imgattach
	Votes int32
}

var pool_MSG_poll_option = sync.Pool{New: func() interface{} { return &MSG_poll_option{} }}

func GET_MSG_poll_option() *MSG_poll_option {
	return pool_MSG_poll_option.Get().(*MSG_poll_option)
}

func (data *MSG_poll_option) Put() {
	data.Id = 0
	data.Name = ``
	data.Displayorder = 0
	data.Aid = 0
	if data.Imginfo != nil {
		data.Imginfo.Put()
		data.Imginfo = nil
	}
	data.Votes = 0
	pool_MSG_poll_option.Put(data)
}
func (data *MSG_poll_option) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_poll_option)
	WRITE_int32(cmd, buf)
	WRITE_MSG_poll_option(data, buf)
}

func WRITE_MSG_poll_option(data *MSG_poll_option, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Id, buf)
	WRITE_string(data.Name, buf)
	WRITE_int8(data.Displayorder, buf)
	WRITE_int64(data.Aid, buf)
	if data.Imginfo == nil {
		WRITE_int8(0, buf)
	} else {
		WRITE_int8(1, buf)
		WRITE_MSG_forum_imgattach(data.Imginfo, buf)
	}
	WRITE_int32(data.Votes, buf)
}

func READ_MSG_poll_option(buf *libraries.MsgBuffer) (data *MSG_poll_option) {
	data = pool_MSG_poll_option.Get().(*MSG_poll_option)
	data.Id = READ_int32(buf)
	data.Name = READ_string(buf)
	data.Displayorder = READ_int8(buf)
	data.Aid = READ_int64(buf)
	Imginfo_len := int(READ_int8(buf))
	if Imginfo_len == 1 {
		data.Imginfo = READ_MSG_forum_imgattach(buf)
	}else{
		data.Imginfo = nil
	}
	data.Votes = READ_int32(buf)
	return
}

type MSG_U2WS_PollThread struct {
	Tid int32
	Ids []int32
}

var pool_MSG_U2WS_PollThread = sync.Pool{New: func() interface{} { return &MSG_U2WS_PollThread{} }}

func GET_MSG_U2WS_PollThread() *MSG_U2WS_PollThread {
	return pool_MSG_U2WS_PollThread.Get().(*MSG_U2WS_PollThread)
}

func (data *MSG_U2WS_PollThread) Put() {
	data.Tid = 0
	data.Ids = data.Ids[:0]
	pool_MSG_U2WS_PollThread.Put(data)
}
func (data *MSG_U2WS_PollThread) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_U2WS_PollThread)
	WRITE_int32(cmd, buf)
	WRITE_MSG_U2WS_PollThread(data, buf)
}

func WRITE_MSG_U2WS_PollThread(data *MSG_U2WS_PollThread, buf *libraries.MsgBuffer) {
	WRITE_int32(data.Tid, buf)
	WRITE_int32(int32(len(data.Ids)), buf)
	for _, v := range data.Ids{
		WRITE_int32(v, buf)
	}
}

func READ_MSG_U2WS_PollThread(buf *libraries.MsgBuffer) (data *MSG_U2WS_PollThread) {
	data = pool_MSG_U2WS_PollThread.Get().(*MSG_U2WS_PollThread)
	data.Tid = READ_int32(buf)
	Ids_len := int(READ_int32(buf))
	for i := 0; i < Ids_len; i++ {
		data.Ids = append(data.Ids, READ_int32(buf))
	}
	return
}

type MSG_WS2U_PollThread_sucess struct {
}

var pool_MSG_WS2U_PollThread_sucess = sync.Pool{New: func() interface{} { return &MSG_WS2U_PollThread_sucess{} }}

func GET_MSG_WS2U_PollThread_sucess() *MSG_WS2U_PollThread_sucess {
	return pool_MSG_WS2U_PollThread_sucess.Get().(*MSG_WS2U_PollThread_sucess)
}

func (data *MSG_WS2U_PollThread_sucess) Put() {
	pool_MSG_WS2U_PollThread_sucess.Put(data)
}
func (data *MSG_WS2U_PollThread_sucess) WRITE(buf *libraries.MsgBuffer) {
	buf.Reset()
	cmd := int32(CMD_MSG_WS2U_PollThread_sucess)
	WRITE_int32(cmd, buf)
	WRITE_MSG_WS2U_PollThread_sucess(data, buf)
}

func WRITE_MSG_WS2U_PollThread_sucess(data *MSG_WS2U_PollThread_sucess, buf *libraries.MsgBuffer) {
}

func READ_MSG_WS2U_PollThread_sucess(buf *libraries.MsgBuffer) (data *MSG_WS2U_PollThread_sucess) {
	data = pool_MSG_WS2U_PollThread_sucess.Get().(*MSG_WS2U_PollThread_sucess)
	return
}

