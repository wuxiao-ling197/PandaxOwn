package entity

/**结构体定义的字段类型请完全匹配数据库字段类型 否则在进行odoorpc调用时即便返回true但实际上并没有成功修改后端数据*/

type ResUsersId struct {
	ID int64 `gorm:"primaryKey;column:id;default:nextval(res_users_id_seq::regclass);NOT NULL" json:"id,omitempty"`
}
type LoginO struct {
	Login    string `gorm:"column:login;NOT NULL" json:"login,omitempty"`
	Password string `gorm:"column:password" json:"password,omitempty" `
	// CompanyId int64  `gorm:"column:company_id;NOT NULL" json:"companyId "`
}

type ResUsersB struct {
	// Id               string         `gorm:"column:id;default:nextval(res_users_id_seq::regclass);NOT NULL" json:"id"`
	CompanyId  int64  `gorm:"column:company_id;NOT NULL" json:"company_id,omitempty"`
	PartnerId  int64  `gorm:"column:partner_id;NOT NULL" json:"partner_id,omitempty"`
	Active     bool   `gorm:"column:active;default:true" json:"active,omitempty"`
	CreateDate string `gorm:"column:create_date" json:"create_date,omitempty"`
	// Login            string         `gorm:"column:login;NOT NULL" json:"login"`
	// Password         string `gorm:"column:password" json:"password"`
	ActionId         int64  `gorm:"column:action_id;comment:'home action'" json:"action_id,omitempty"`
	CreateUid        int64  `gorm:"column:create_uid;comment:'created by'" json:"create_uid,omitempty"`
	WriteUid         int64  `gorm:"column:write_uid;comment:'last updated by'" json:"write_uid,omitempty"`
	Signature        string `gorm:"column:signature;comment:'email signature'" json:"signature,omitempty"`
	Share            bool   `gorm:"column:share;comment:'share user'" json:"share,omitempty"`
	WriteDate        string `gorm:"column:write_date;comment:'last updated on'" json:"write_date,omitempty"`
	NestSecret       string `gorm:"column:nest_secret;comment:'nest totp secret'" json:"nest_secret,omitempty"`
	TotpSecret       string `gorm:"column:totp_secret" json:"totp_secret,omitempty"`
	TourEnabled      bool   `gorm:"column:tour_enabled;comment:'onboarding'" json:"tour_enabled,omitempty"`
	NotificationType string `gorm:"column:notification_type;NOT NULL;comment:'notification'" json:"notification_type,omitempty"`
	Karma            string `gorm:"column:karma;comment:'karma'" json:"karma,omitempty"`
	RankId           int64  `gorm:"column:rank_id;comment:'rank'" json:"rank_id,omitempty"`
	NextRankId       int64  `gorm:"column:next_rank_id;comment:'next rank'" json:"next_rank_id,omitempty"`
	OdoobotState     string `gorm:"column:odoobot_state;comment:'odoobot status'" json:"odoobot_state,omitempty"`
	OdoobotFailed    bool   `gorm:"column:odoobot_failed;comment:'odoobot failed'" json:"odoobot_failed,omitempty"`
	PandaxSecret     string `gorm:"column:pandax_secret;comment:'pandax secret'" json:"pandax_secret,omitempty"`
}

type ResUsers struct {
	ResUsersId
	ResUsersB
	LoginO
	// HrEmployee HrEmployee `gorm:"foreignKey:UserId"` // 指定外键
}

// 自定义序列化方法
// func (r ResUsers) MarshalJSON() ([]byte, error) {
// 	// 创建一个临时 map，用于存储字段值
// 	temp := make(map[string]interface{})
// 	// 强制保留 bool 类型的字段，即使值为 false
// 	temp["active"] = r.Active
// 	temp["share"] = r.Share
// 	temp["tour_enabled"] = r.TourEnabled
// 	// 其他字段按 omitempty 规则处理
// 	if r.Login != "" {
// 		temp["name"] = r.Login
// 	}
// 	if r.Password != "" {
// 		temp["password"] = r.Password
// 	}
// 	// 序列化 map
// 	return json.Marshal(temp)
// }

// 用户重置密码 odoo中是通过发送邮件根据内部提供的网址实现
type ResUsersPwd struct {
	OldPassword string `json:"oldPassword" form:"oldPassword"`
	NewPassword string `json:"newPassword" form:"newPassword"`
}

// TableName "res_users"
type ResUsersPage struct {
	ResUsersId
	ResUsersB
	LoginO
	Employee HrEmployee `gorm:"foreignKey:UserId" json:"employee,omitempty"` // 指定外键
	// HrDepartmentName string     `gorm:"-" json:"hrDepartmentName"`
}

func (ResUsersPage) TableName() string {
	return "res_users" // 或者 "HR_EMPLOYEES" 根据实际表名调整
}

type ResUsersView struct {
	ResUsersId
	ResUsersB
	LoginO
	Employee        HrEmployee `gorm:"foreignKey:UserId"` // 指定外键
	ResCompanyName  string     `gorm:"-" json:"resCompanyName,omitempty"`
	ResUserTypeName string     `gorm:"-" json:"resUserTypeName,omitempty"`
}

func (ResUsersView) TableName() string {
	return "res_users" // 或者 "HR_EMPLOYEES" 根据实际表名调整
}

type HrEmployee struct {
	ID                          int64  `gorm:"primaryKey;column:id;default:nextval(hr_employee_id_seq::regclass);NOT NULL" json:"id,omitempty"`
	ResourceId                  int64  `gorm:"column:resource_id;NOT NULL" json:"resource_id,omitempty"`
	CompanyId                   int64  `gorm:"column:company_id;NOT NULL" json:"company_id,omitempty"`
	ResourceCalendarId          int64  `gorm:"column:resource_calendar_id" json:"resource_calendar_id,omitempty"`
	MessageMainAttachmentId     int64  `gorm:"column:message_main_attachment_id" json:"message_main_attachment_id,omitempty"`
	Color                       int64  `gorm:"column:color" json:"color,omitempty"`
	DepartmentId                int64  `gorm:"column:department_id;" json:"department_id,omitempty"`
	JobId                       int64  `gorm:"column:job_id" json:"job_id,omitempty"`
	AddressId                   int64  `gorm:"column:address_id" json:"address_id,omitempty"`
	WorkContactId               int64  `gorm:"column:work_contact_id" json:"work_contact_id,omitempty"`
	WorkLocationId              int64  `gorm:"column:work_location_id" json:"work_location_id,omitempty"`
	UserId                      int64  `gorm:"column:user_id;" json:"user_id,omitempty"`
	ParentId                    int64  `gorm:"column:parent_id" json:"parent_id,omitempty"`
	CoachId                     int64  `gorm:"column:coach_id" json:"coach_id,omitempty"`
	PrivateStateId              int64  `gorm:"column:private_state_id" json:"private_state_id,omitempty"`
	PrivateCountryId            int64  `gorm:"column:private_country_id" json:"private_country_id,omitempty"`
	CountryId                   int64  `gorm:"column:country_id" json:"country_id,omitempty"`
	Children                    int64  `gorm:"column:children" json:"children,omitempty"`
	CountryOfBirth              int64  `gorm:"column:country_of_birth" json:"country_of_birth,omitempty"`
	BankAccountId               int64  `gorm:"column:bank_account_id" json:"bank_account_id,omitempty"`
	DistanceHomeWork            int64  `gorm:"column:distance_home_work" json:"distance_home_work,omitempty"`
	KmHomeWork                  int64  `gorm:"column:km_home_work" json:"km_home_work,omitempty"`
	DepartureReasonId           int64  `gorm:"column:departure_reason_id" json:"departure_reason_id,omitempty"`
	CreateUid                   string `gorm:"column:create_uid" json:"create_uid,omitempty"`
	WriteUid                    string `gorm:"column:write_uid" json:"write_uid,omitempty"`
	Name                        string `gorm:"column:name" json:"name,omitempty"`
	JobTitle                    string `gorm:"column:job_title" json:"job_title,omitempty"`
	WorkPhone                   string `gorm:"column:work_phone" json:"work_phone,omitempty"`
	MobilePhone                 string `gorm:"column:mobile_phone" json:"mobile_phone,omitempty"`
	WorkEmail                   string `gorm:"column:work_email" json:"work_email,omitempty"`
	PrivateStreet               string `gorm:"column:private_street" json:"private_street,omitempty"`
	PrivateStreet2              string `gorm:"column:private_street2" json:"private_street2,omitempty"`
	PrivateCity                 string `gorm:"column:private_city" json:"private_city,omitempty"`
	PrivateZip                  string `gorm:"column:private_zip" json:"private_zip,omitempty"`
	PrivatePhone                string `gorm:"column:private_phone" json:"resprivate_phoneource_id,omitempty"`
	PrivateEmail                string `gorm:"column:private_email" json:"private_email,omitempty"`
	Lang                        string `gorm:"column:lang" json:"lang,omitempty"`
	Gender                      string `gorm:"column:gender" json:"gender,omitempty"`
	Marital                     string `gorm:"column:marital;NOT NULL" json:"marital,omitempty"`
	SpouseCompleteName          string `gorm:"column:spouse_complete_name" json:"spouse_complete_name,omitempty"`
	PlaceOfBirth                string `gorm:"column:place_of_birth" json:"place_of_birth,omitempty"`
	Ssnid                       string `gorm:"column:ssnid" json:"ssnid,omitempty"`
	Sinid                       string `gorm:"column:sinid" json:"sinid,omitempty"`
	IdentificationId            string `gorm:"column:identification_id" json:"identification_id,omitempty"`
	PassportId                  string `gorm:"column:passport_id" json:"passport_id,omitempty"`
	PermitNo                    string `gorm:"column:permit_no" json:"permit_no,omitempty"`
	VisaNo                      string `gorm:"column:visa_no" json:"visa_no,omitempty"`
	Certificate                 string `gorm:"column:certificate" json:"certificate,omitempty"`
	StudyField                  string `gorm:"column:study_field" json:"study_field,omitempty"`
	StudySchool                 string `gorm:"column:study_school" json:"study_school,omitempty"`
	EmergencyContact            string `gorm:"column:emergency_contact" json:"emergency_contact,omitempty"`
	EmergencyPhone              string `gorm:"column:emergency_phone" json:"emergency_phone,omitempty"`
	DistanceHomeWorkUnit        string `gorm:"column:distance_home_work_unit;NOT NULL" json:"distance_home_work_unit,omitempty"`
	EmployeeType                string `gorm:"column:employee_type;NOT NULL" json:"employee_type,omitempty"`
	Barcode                     string `gorm:"column:barcode" json:"barcode,omitempty"`
	Pin                         string `gorm:"column:pin" json:"pin,omitempty"`
	PrivateCarPlate             string `gorm:"column:private_car_plate" json:"private_car_plate,omitempty"`
	SpouseBirthdate             string `gorm:"column:spouse_birthdate" json:"spouse_birthdate,omitempty"`
	Birthday                    string `gorm:"column:birthday" json:"birthday,omitempty"`
	VisaExpire                  string `gorm:"column:visa_expire" json:"visa_expire,omitempty"`
	WorkPermitExpirationDate    string `gorm:"column:work_permit_expiration_date" json:"work_permit_expiration_date,omitempty"`
	DepartureDate               string `gorm:"column:departure_date" json:"departure_date,omitempty"`
	EmployeeProperties          string `gorm:"column:employee_properties" json:"employee_properties,omitempty"`
	AdditionalNote              string `gorm:"column:additional_note" json:"additional_note,omitempty"`
	Notes                       string `gorm:"column:notes" json:"notes,omitempty"`
	DepartureDescription        string `gorm:"column:departure_description" json:"departure_description,omitempty"`
	Active                      bool   `gorm:"column:active" json:"active,omitempty"`
	IsFlexible                  bool   `gorm:"column:is_flexible" json:"is_flexible,omitempty"`
	IsFullyFlexible             bool   `gorm:"column:is_fully_flexible" json:"is_fully_flexible,omitempty"`
	WorkPermitScheduledActivity bool   `gorm:"column:work_permit_scheduled_activity" json:"work_permit_scheduled_activity,omitempty"`
	CreateDate                  string `gorm:"column:create_date" json:"create_date,omitempty"`
	WriteDate                   string `gorm:"column:write_date" json:"write_date,omitempty"`
	AttendanceManagerId         int64  `gorm:"column:attendance_manager_id" json:"attendance_manager_id,omitempty"`
	LastAttendanceId            int64  `gorm:"column:last_attendance_id" json:"last_attendance_id,omitempty"`
	LastCheckIn                 string `gorm:"column:last_check_in" json:"last_check_in,omitempty"`
	LastCheckOut                string `gorm:"column:last_check_out" json:"last_check_out,omitempty"`
	ContractId                  int64  `gorm:"column:contract_id" json:"contract_id,omitempty"`
	LegalName                   string `gorm:"column:legal_name" json:"legal_name,omitempty"`
	Vehicle                     string `gorm:"column:vehicle" json:"vehicle,omitempty"`
	FirstContractDate           string `gorm:"column:first_contract_date" json:"first_contract_date,omitempty"`
	ContractWarning             bool   `gorm:"column:contract_warning" json:"contract_warning,omitempty"`
	ExpenseManagerId            int64  `gorm:"column:expense_manager_id" json:"expense_manager_id,omitempty"`
	LeaveManagerId              int64  `gorm:"column:leave_manager_id" json:"leave_manager_id,omitempty"`
	HourlyCost                  string `gorm:"column:hourly_cost" json:"hourly_cost,omitempty"`
	PersonalMobile              string `gorm:"column:personal_mobile" json:"personal_mobile,omitempty"`
	JoiningDate                 string `gorm:"column:joining_date" json:"joining_date,omitempty"`
	IdExpiryDate                string `gorm:"column:id_expiry_date" json:"id_expiry_date,omitempty"`
	PassportExpiryDate          string `gorm:"column:passport_expiry_date" json:"passport_expiry_date,omitempty"`
	ResignDate                  string `gorm:"column:resign_date" json:"resign_date,omitempty"`
	Resigned                    bool   `gorm:"column:resigned" json:"resigned,omitempty"`
	Fired                       bool   `gorm:"column:fired" json:"fired,omitempty"`
}

func (HrEmployee) TableName() string {
	return "hr_employee" // 或者 "HR_EMPLOYEES" 根据实际表名调整
}

// 主员工数据 外加关联用户的登录名和状态
type EmployeeWithUser struct {
	HrEmployee
	UserName       string `gorm:"references:username" json:"username,omitempty"` //查询时指定别名为 username，因此需要定义column
	UserActive     bool   `gorm:"references:user_active" json:"user_active,omitempty"`
	DepartmentName string `gorm:"references:department_name" json:"department_name,omitempty"`
	CompanyName    string `gorm:"references:company_name" json:"company_name,omitempty"`
}

// 用户-新建 web_save
type CreateUserDto struct {
	Name      string `json:"name,omitempty"`
	Login     string `json:"login,omitempty"`
	Email     string `json:"email,omitempty"`
	CompanyId int64  `json:"company_id,omitempty"`
}

/**omitempty 序列化时忽略零值/空值
如果不定义非空字段
则对结构体序列化的结果为：
map[action_id: active:false company_id:0 create_date: create_uid: id:0 karma:100 login: nest_secret: next_rank_id:
notification_type: odoobot_failed:false odoobot_state: partner_id: password: rank_id: share:false signature: totp_secret:
tour_enabled:false write_date: write_uid:]

如果定义非空字段 那么属于默认零值的字段值(false、0、"" 等)在序列化时会完全被忽略，此时无法正确序列化，例如：
write := entity.ResUsers{
		ResUsersB: entity.ResUsersB{
			Active:      false,
			Share:       false,
			TourEnabled: false,
			Karma:       "100",
		},
	}
		最终里劣化结果为 map[karma:100]
**/
// 自定义序列化方法
// func (r HrEmployee) MarshalJSON() ([]byte, error) {
// 	// 创建一个临时 map，用于存储字段值
// 	temp := make(map[string]interface{})
// 	// 强制保留 bool 类型的字段，即使值为 false  保留后即便在传参的结构体中没有给这些字段赋值也会与被赋值的字段一起序列化返回
// 	//  而且若没有指定值会将默认值填充进去
// 	temp["active"] = r.Active
// 	temp["is_flexible"] = r.IsFlexible
// 	temp["is_fully_flexible"] = r.IsFullyFlexible
// 	temp["contract_warning"] = r.ContractWarning
// 	temp["resigned"] = r.Resigned
// 	temp["fired"] = r.Fired
// 	// 如果要设置int类型数据为空只需为其赋值为0
// 	// temp["department_id"] = r.DepartmentId
// 	// temp["user_id"] = r.UserId
// 	// temp["job_title"] = r.JobTitle
// 	// 其他字段按 omitempty 规则处理
// 	// if r.DepartmentId != 0 {
// 	// 	temp["department_id"] = r.DepartmentId
// 	// }
// 	// if r.UserId != 0 {
// 	// 	temp["user_id"] = r.UserId
// 	// }
// 	// 序列化 map
// 	return json.Marshal(temp)
// }

// TOTP双重验证
type AuthTotpWizard struct {
	Id         int64  `gorm:"column:id;default:nextval(auth_totp_wizard_id_seq::regclass);NOT NULL" json:"id,omitempty"`
	UserId     int64  `gorm:"column:user_id;NOT NULL" json:"user_id,omitempty"`
	CreateUid  int64  `gorm:"column:create_uid" json:"create_uid,omitempty"`
	WriteUid   int64  `gorm:"column:write_uid" json:"write_uid,omitempty"`
	Secret     string `gorm:"column:secret;NOT NULL" json:"secret,omitempty"`
	Url        string `gorm:"column:url" json:"url,omitempty"`
	Code       string `gorm:"column:code" json:"code,omitempty"`
	CreateDate string `gorm:"column:create_date" json:"create_date,omitempty"`
	WriteDate  string `gorm:"column:write_date" json:"write_date,omitempty"`
	Qrcode     string `gorm:"column:qrcode" json:"qrcode,omitempty"`
}

func (AuthTotpWizard) TableName() string {
	return "auth_totp_wizard"
}
