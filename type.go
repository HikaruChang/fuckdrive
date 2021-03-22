package main

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Cookie   string `json:"cookie"`
	Project  int    `json:"project"`
	Prefer   struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"prefer"`
	Size int `json:"size"`
}

type URLConst struct {
	base                     string
	mobileLogin              string
	appType                  string
	isLogin                  string
	getUserInfo              string
	getDaysOrder             string
	getInfoByTimeId          string
	getTimesInfoByCoachIDNew string
	orderCoachNew            string
}

type ClientConst struct {
	accept      string
	contentType string
	host        string
	origin      string
	referer     string
	userAgent   string
}

type AppTypeResp struct {
	Result    bool   `json:"Result"`
	IsSuccess bool   `json:"isSuccess"`
	ExType    int    `json:"ExType"`
	Message   string `json:"Message"`
	Data      int    `json:"Data"`
}

type MobileLoginParams struct {
	LoginType int    `json:"LoginType"`
	UserName  string `json:"UserName"`
	Password  string `json:"Password"`
}

type MobileLoginResp struct {
	D struct {
		Type      string `json:"__type"`
		ErrInfo   string `json:"ErrInfo"`
		SverTime  string `json:"sverTime"`
		Result    bool   `json:"Result"`
		State     int    `json:"State"`
		IsSuccess bool   `json:"isSuccess"`
		ExType    int    `json:"ExType"`
		Message   string `json:"Message"`
		Data      string `json:"Data"`
	} `json:"d"`
}

type GetUserInfoResp struct {
	D struct {
		Type      string `json:"__type"`
		ErrInfo   string `json:"ErrInfo"`
		SverTime  string `json:"sverTime"`
		Result    bool   `json:"Result"`
		State     int    `json:"State"`
		IsSuccess bool   `json:"isSuccess"`
		ExType    int    `json:"ExType"`
		Message   string `json:"Message"`
		Data      struct {
			HasOrerTimes      int    `json:"HasOrerTimes"`
			MaxOrderTimes     int    `json:"MaxOrderTimes"`
			NoMoreMoneyTimeID string `json:"NoMoreMoneyTimeID"`
			StuDepID          string `json:"StuDepID"`
			CardNumber        string `json:"cardNumber"`
			CardPwd           string `json:"cardPwd"`
			DtOrder           string `json:"dtOrder"`
			Introducername    string `json:"introducername"`
			Introducerphone   string `json:"introducerphone"`
			IsBindCard        int    `json:"isBindCard"`
			IsBlackList       bool   `json:"isBlackList"`
			IsEidtPaw         string `json:"isEidtPaw"`
			IsNeedPL          bool   `json:"isNeedPL"`
			IsOpenClass       bool   `json:"isOpenClass"`
			IsOver            bool   `json:"isOver"`
			LoginRecordNumber string `json:"loginRecordNumber"`
			NHasDuNote        int    `json:"nHasDuNote"`
			NNoteAllCount     int    `json:"nNoteAllCount"`
			NOrderCount       int    `json:"nOrderCount"`
			NPLCount          int    `json:"nPLCount"`
			NPjCount          int    `json:"nPjCount"`
			NQueryNoReadNote  int    `json:"nQueryNoReadNote"`
			NRRecordCount     int    `json:"nRRecordCount"`
			NRecordCount      int    `json:"nRecordCount"`
			NSessionType      int    `json:"nSessionType"`
			NWeiduNote        int    `json:"nWeiduNote"`
			Openid            string `json:"openid"`
			StuO2state        string `json:"stuO2state"`
			StuO3state        string `json:"stuO3state"`
			StuState          string `json:"stuState"`
			Stucarcode        string `json:"stucarcode"`
			Stunum            string `json:"stunum"`
			SvrTime           string `json:"svrTime"`
			TimeSpan          int    `json:"timeSpan"`
			一卡通单价             int    `json:"一卡通单价"`
			一卡通类型             int    `json:"一卡通类型"`
			分教练               bool   `json:"分教练"`
			分组                bool   `json:"分组"`
			分组信息排序            string `json:"分组信息排序"`
			分组名称              string `json:"分组名称"`
			分组排序              int    `json:"分组排序"`
			分组状态              string `json:"分组状态"`
			分组编号              string `json:"分组编号"`
			包含设备              bool   `json:"包含设备"`
			卡类型               string `json:"卡类型"`
			取消休眠去向            string `json:"取消休眠去向"`
			可预约天数             int    `json:"可预约天数"`
			可预约次数             int    `json:"可预约次数"`
			培训模式              string `json:"培训模式"`
			培训模式编号            string `json:"培训模式编号"`
			姓名                string `json:"姓名"`
			学时                string `json:"学时"`
			开始预约时间            string `json:"开始预约时间"`
			当前流程状态            int    `json:"当前流程状态"`
			录入时间              string `json:"录入时间"`
			手机                string `json:"手机"`
			教练排序              string `json:"教练排序"`
			显示单价              int    `json:"显示单价"`
			显示签到记录            bool   `json:"显示签到记录"`
			机构编号              string `json:"机构编号"`
			流程标志              int    `json:"流程标志"`
			流程类型              int    `json:"流程类型"`
			照片                string `json:"照片"`
			班别编号              string `json:"班别编号"`
			登录标志              string `json:"登录标志"`
			登陆方式              string `json:"登陆方式"`
			科三学时              string `json:"科三学时"`
			科二学时              string `json:"科二学时"`
			科目                string `json:"科目"`
			科目代号              int    `json:"科目代号"`
			结束预约时间            string `json:"结束预约时间"`
			编号                string `json:"编号"`
			考试标志              int    `json:"考试标志"`
			自主考试登记            bool   `json:"自主考试登记"`
			评价模式              int    `json:"评价模式"`
			身份证号码             string `json:"身份证号码"`
			车型                string `json:"车型"`
			违规规则              string `json:"违规规则"`
			违规规则编号            string `json:"违规规则编号"`
			选组方式              int    `json:"选组方式"`
			金额                string `json:"金额"`
			预约时间限制            bool   `json:"预约时间限制"`
			预约模式              string `json:"预约模式"`
			预约模式分类            int    `json:"预约模式分类"`
			预约模式周期            string `json:"预约模式周期"`
			预约模式编号            string `json:"预约模式编号"`
		} `json:"Data"`
	} `json:"d"`
}

type Dresp struct {
	D string `json:"d"`
}

type GetDaysOrderResp struct {
	Result    bool     `json:"Result"`
	IsSuccess bool     `json:"isSuccess"`
	ExType    int      `json:"ExType"`
	Message   string   `json:"Message"`
	Data      []string `json:"Data"`
}

type GetInfoByTimeIdParams struct {
	ArgDay     string `json:"argDay"`
	NPageIndex int    `json:"nPageIndex"`
	NUnitNO    int    `json:"nUnitNO"`
}

type GetInfoByTimeIdResp struct {
	BillNumber   string `json:"BillNumber"`
	BusinessData struct {
		CoachInfo struct {
			Data []struct {
				ROWID          int    `json:"ROWID"`
				CoachID        string `json:"教练__编号"`
				CoachPhoto     string `json:"教练__照片"`
				CoachName      string `json:"教练__姓名"`
				Guid_T         string `json:"guid_T"`
				Full_T         int    `json:"满度_T"`
				CoachFull      string `json:"教练__满度"`
				CoachCarType   string `json:"教练__准驾车型__名称"`
				CoachField     string `json:"教练__场地"`
				CoachPhone     string `json:"教练__手机"`
				CoachCarNo     string `json:"教练__车辆信息__车牌号"`
				CoachScoreTime int    `json:"教练__评分次数"`
				CoachDesc      string `json:"教练__自我介绍"`
			} `json:"Data"`
			RowCount           int    `json:"RowCount"`
			BID                string `json:"BID"`
			IntroductFieldName string `json:"IntroductFieldName"`
			IntroductTableName string `json:"IntroductTableName"`
			KeyField           string `json:"KeyField"`
			Name               string `json:"Name"`
			Page               struct {
				Creat           bool     `json:"Creat"`
				IsPage          bool     `json:"isPage"`
				PageCount       int      `json:"PageCount"`
				PageCurrent     int      `json:"PageCurrent"`
				PageTotal       int      `json:"PageTotal"`
				SystemWhereItem []string `json:"SystemWhereItem"`
				Total           int      `json:"Total"`
				WhereItem       []string `json:"WhereItem"`
			} `json:"Page"`
			PrimaryKeyFields []struct {
			} `json:"PrimaryKeyFields"`
			UIIndex string `json:"UIIndex"`
		} `json:"教练信息"`
	} `json:"BusinessData"`
	Common     string `json:"common"`
	CreatID    string `json:"CreatID"`
	Name       string `json:"Name"`
	PluginData string `json:"PluginData"`
	State      int    `json:"state"`
	Summary    string `json:"summary"`
	Type       int    `json:"Type"`
	UiIndex    string `json:"uiIndex"`
}

type GetTimesInfoByCoachIDNewParams struct {
	CoachID string `json:"coachID"`
	Date    string `json:"date"`
	Subid   int    `json:"subid"`
}

type Store struct {
	Cookie  string
	CoachID string
	StateID string
	OrderNo int
	Data    []struct {
		CoachName    string  `json:"姓名"`
		CoachPhone   string  `json:"教练__手机"`
		TrainingData string  `json:"培训日期"`
		Number       int     `json:"编号"`
		CoachID      string  `json:"教练__编号"`
		TBTimeNo     string  `json:"TB_时间段__编号"`
		Status       string  `json:"状态"`
		StartTime    string  `json:"开始时间"`
		EndTime      string  `json:"结束时间"`
		OrderStu     string  `json:"预约学员"`
		OrderNo      int     `json:"预约人数"`
		OrderedNo    int     `json:"已预约人数"`
		TrainDate    string  `json:"培训日期1"`
		UnitPrice    float64 `json:"单价"`
		Discount     int     `json:"折扣"`
		SubID        int     `json:"科目__编号"`
		TrainRes     int     `json:"培训资源"`
		Optional     bool    `json:"是否自选"`
	} `json:"Data"`
}

type GetTimesInfoByCoachIDNewResp struct {
	Result    bool   `json:"Result"`
	IsSuccess bool   `json:"isSuccess"`
	ExType    int    `json:"ExType"`
	Message   string `json:"Message"`
	Data      []struct {
		CoachName    string  `json:"姓名"`
		CoachPhone   string  `json:"教练__手机"`
		TrainingData string  `json:"培训日期"`
		Number       int     `json:"编号"`
		CoachID      string  `json:"教练__编号"`
		TBTimeNo     string  `json:"TB_时间段__编号"`
		Status       string  `json:"状态"`
		StartTime    string  `json:"开始时间"`
		EndTime      string  `json:"结束时间"`
		OrderStu     string  `json:"预约学员"`
		OrderNo      int     `json:"预约人数"`
		OrderedNo    int     `json:"已预约人数"`
		TrainDate    string  `json:"培训日期1"`
		UnitPrice    float64 `json:"单价"`
		Discount     int     `json:"折扣"`
		SubID        int     `json:"科目__编号"`
		TrainRes     int     `json:"培训资源"`
		Optional     bool    `json:"是否自选"`
	} `json:"Data"`
}

type OrderCoachNewParams struct {
	PXResID string `json:"PXResID"`
	SubID   string `json:"SubID"`
	StateID string `json:"stateID"`
}

type OrderCoachNewResp struct {
	D struct {
		Item1 bool     `json:"Item1"`
		Item2 []string `json:"Item2"`
	} `json:"d"`
}
