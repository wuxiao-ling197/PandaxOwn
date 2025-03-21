package entity

import (
	"github.com/PandaXGO/PandaKit/model"
)

type LoginM struct {
	Username string `gorm:"type:varchar(64)" json:"username"`
	Password string `gorm:"type:varchar(128)" json:"password"`
}

type SysUserId struct {
	UserId int64 `gorm:"primary_key;AUTO_INCREMENT"  json:"userId"` // 编码
}

type SysUserB struct {
	NickName       string `gorm:"type:varchar(128)" json:"nickName"` // 昵称
	Phone          string `gorm:"type:varchar(11)" json:"phone"`     // 手机号
	RoleId         int64  `gorm:"type:int" json:"roleId"`            // 角色编码
	Salt           string `gorm:"type:varchar(255)" json:"salt"`     //盐
	Avatar         string `gorm:"type:varchar(255)" json:"avatar"`   //头像
	Sex            string `gorm:"type:varchar(255)" json:"sex"`      //性别
	Email          string `gorm:"type:varchar(128)" json:"email"`    //邮箱
	OrganizationId int64  `gorm:"type:int" json:"organizationId"`    //组织编码
	PostId         int64  `gorm:"type:int" json:"postId"`            //职位编码
	RoleIds        string `gorm:"type:varchar(255)" json:"roleIds"`  //多角色
	PostIds        string `gorm:"type:varchar(255)" json:"postIds"`  // 多岗位
	CreateBy       string `gorm:"type:varchar(128)" json:"createBy"` //
	UpdateBy       string `gorm:"type:varchar(128)" json:"updateBy"` //
	Remark         string `gorm:"type:varchar(255)" json:"remark"`   //备注
	Status         string `gorm:"type:varchar(1);" json:"status"`
	model.BaseModel
}

type SysUser struct {
	SysUserId
	SysUserB
	LoginM
}

type SysUserPwd struct {
	OldPassword string `json:"oldPassword" form:"oldPassword"`
	NewPassword string `json:"newPassword" form:"newPassword"`
}

type SysUserPage struct {
	SysUserId
	SysUserB
	LoginM
	OrganizationName string `gorm:"-" json:"organizationName"`
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
	UUID     string `form:"UUID" json:"uuid" binding:"required"`
}

type SysUserView struct {
	SysUserId
	SysUserB
	LoginM
	RoleName string `gorm:"column:role_name"  json:"role_name"`
}

// Odoo migrate
// type ResUsersId struct {
// 	ID int64 `gorm:"column:id;default:nextval(res_users_id_seq::regclass);NOT NULL" json:"id"`
// }
// type LoginO struct {
// 	Login    string `gorm:"column:login;NOT NULL" json:"login"`
// 	Password string `gorm:"column:password" json:"password"`
// }

// type ResUsersB struct {
// 	// Id               string         `gorm:"column:id;default:nextval(res_users_id_seq::regclass);NOT NULL" json:"id"`
// 	CompanyId  int64          `gorm:"column:company_id;NOT NULL" json:"companyId "`
// 	PartnerId  string         `gorm:"column:partner_id;NOT NULL" json:"partnerId"`
// 	Active     bool           `gorm:"column:active;default:true" json:"active"`
// 	CreateDate sql.NullString `gorm:"column:create_date"`
// 	// Login            string         `gorm:"column:login;NOT NULL" json:"login"`
// 	// Password         sql.NullString `gorm:"column:password" json:"password"`
// 	ActionId         sql.NullString `gorm:"column:action_id;comment:'home action'"`
// 	CreateUid        sql.NullString `gorm:"column:create_uid;comment:'created by'"`
// 	WriteUid         sql.NullString `gorm:"column:write_uid;comment:'last updated by'"`
// 	Signature        sql.NullString `gorm:"column:signature;comment:'email signature'"`
// 	Share            sql.NullBool   `gorm:"column:share;comment:'share user'"`
// 	WriteDate        sql.NullString `gorm:"column:write_date;comment:'last updated on'"`
// 	NestSecret       sql.NullString `gorm:"column:nest_secret;comment:'nest totp secret'"`
// 	TotpSecret       sql.NullString `gorm:"column:totp_secret"`
// 	TourEnabled      sql.NullBool   `gorm:"column:tour_enabled;comment:'onboarding'"`
// 	NotificationType string         `gorm:"column:notification_type;NOT NULL;comment:'notification'"`
// 	Karma            sql.NullString `gorm:"column:karma;comment:'karma'"`
// 	RankId           sql.NullString `gorm:"column:rank_id;comment:'rank'"`
// 	NextRankId       sql.NullString `gorm:"column:next_rank_id;comment:'next rank'"`
// 	OdoobotState     sql.NullString `gorm:"column:odoobot_state;comment:'odoobot status'"`
// 	OdoobotFailed    sql.NullBool   `gorm:"column:odoobot_failed;comment:'odoobot failed'"`
// }

// type ResUsers struct {
// 	ResUsersId
// 	ResUsersB
// 	LoginO
// }

// type ResUsersPwd struct {
// 	OldPassword string `json:"oldPassword" form:"oldPassword"`
// 	NewPassword string `json:"newPassword" form:"newPassword"`
// }

// type ResUsersPage struct {
// 	ResUsersId
// 	ResUsersB
// 	LoginO
// 	HrEmployee
// 	HrDepartmentName string `gorm:"-" json:"hrDepartmentName"`
// }

// type ResUsersView struct {
// 	ResUsersId
// 	ResUsersB
// 	LoginO
// 	HrEmployee
// 	ResGroupsName   string `gorm:"-" json:"resGroupsName"`
// 	ResUserTypeName string `gorm:"-" json:"resUserTypeName"`
// }

// type HrEmployee struct {
// 	Id                          string         `gorm:"column:id;default:nextval(hr_employee_id_seq::regclass);NOT NULL"`
// 	ResourceId                  string         `gorm:"column:resource_id;NOT NULL;comment:'resource'"`
// 	CompanyId                   string         `gorm:"column:company_id;NOT NULL;comment:'company'"`
// 	ResourceCalendarId          sql.NullString `gorm:"column:resource_calendar_id;comment:'working hours'"`
// 	MessageMainAttachmentId     sql.NullString `gorm:"column:message_main_attachment_id;comment:'main attachment'"`
// 	Color                       sql.NullString `gorm:"column:color;comment:'color index'"`
// 	DepartmentId                sql.NullString `gorm:"column:department_id;comment:'department'"`
// 	JobId                       sql.NullString `gorm:"column:job_id;comment:'job position'"`
// 	AddressId                   sql.NullString `gorm:"column:address_id;comment:'work address'"`
// 	WorkContactId               sql.NullString `gorm:"column:work_contact_id;comment:'work contact'"`
// 	WorkLocationId              sql.NullString `gorm:"column:work_location_id;comment:'work location'"`
// 	UserId                      sql.NullString `gorm:"column:user_id;comment:'user'"`
// 	ParentId                    sql.NullString `gorm:"column:parent_id;comment:'manager'"`
// 	CoachId                     sql.NullString `gorm:"column:coach_id;comment:'coach'"`
// 	PrivateStateId              sql.NullString `gorm:"column:private_state_id;comment:'private state'"`
// 	PrivateCountryId            sql.NullString `gorm:"column:private_country_id;comment:'private country'"`
// 	CountryId                   sql.NullString `gorm:"column:country_id;comment:'nationality (country)'"`
// 	Children                    sql.NullString `gorm:"column:children;comment:'number of dependent children'"`
// 	CountryOfBirth              sql.NullString `gorm:"column:country_of_birth;comment:'country of birth'"`
// 	BankAccountId               sql.NullString `gorm:"column:bank_account_id;comment:'bank account'"`
// 	DistanceHomeWork            sql.NullString `gorm:"column:distance_home_work;comment:'home-work distance'"`
// 	KmHomeWork                  sql.NullString `gorm:"column:km_home_work;comment:'home-work distance in km'"`
// 	DepartureReasonId           sql.NullString `gorm:"column:departure_reason_id;comment:'departure reason'"`
// 	CreateUid                   sql.NullString `gorm:"column:create_uid;comment:'created by'"`
// 	WriteUid                    sql.NullString `gorm:"column:write_uid;comment:'last updated by'"`
// 	Name                        sql.NullString `gorm:"column:name;comment:'employee name'"`
// 	JobTitle                    sql.NullString `gorm:"column:job_title;comment:'job title'"`
// 	WorkPhone                   sql.NullString `gorm:"column:work_phone;comment:'work phone'"`
// 	MobilePhone                 sql.NullString `gorm:"column:mobile_phone;comment:'work mobile'"`
// 	WorkEmail                   sql.NullString `gorm:"column:work_email;comment:'work email'"`
// 	PrivateStreet               sql.NullString `gorm:"column:private_street;comment:'private street'"`
// 	PrivateStreet2              sql.NullString `gorm:"column:private_street2;comment:'private street2'"`
// 	PrivateCity                 sql.NullString `gorm:"column:private_city;comment:'private city'"`
// 	PrivateZip                  sql.NullString `gorm:"column:private_zip;comment:'private zip'"`
// 	PrivatePhone                sql.NullString `gorm:"column:private_phone;comment:'private phone'"`
// 	PrivateEmail                sql.NullString `gorm:"column:private_email;comment:'private email'"`
// 	Lang                        sql.NullString `gorm:"column:lang;comment:'lang'"`
// 	Gender                      sql.NullString `gorm:"column:gender;comment:'gender'"`
// 	Marital                     string         `gorm:"column:marital;NOT NULL;comment:'marital status'"`
// 	SpouseCompleteName          sql.NullString `gorm:"column:spouse_complete_name;comment:'spouse complete name'"`
// 	PlaceOfBirth                sql.NullString `gorm:"column:place_of_birth;comment:'place of birth'"`
// 	Ssnid                       sql.NullString `gorm:"column:ssnid;comment:'ssn no'"`
// 	Sinid                       sql.NullString `gorm:"column:sinid;comment:'sin no'"`
// 	IdentificationId            sql.NullString `gorm:"column:identification_id;comment:'identification no'"`
// 	PassportId                  sql.NullString `gorm:"column:passport_id;comment:'passport no'"`
// 	PermitNo                    sql.NullString `gorm:"column:permit_no;comment:'work permit no'"`
// 	VisaNo                      sql.NullString `gorm:"column:visa_no;comment:'visa no'"`
// 	Certificate                 sql.NullString `gorm:"column:certificate;comment:'certificate level'"`
// 	StudyField                  sql.NullString `gorm:"column:study_field;comment:'field of study'"`
// 	StudySchool                 sql.NullString `gorm:"column:study_school;comment:'school'"`
// 	EmergencyContact            sql.NullString `gorm:"column:emergency_contact;comment:'contact name'"`
// 	EmergencyPhone              sql.NullString `gorm:"column:emergency_phone;comment:'contact phone'"`
// 	DistanceHomeWorkUnit        string         `gorm:"column:distance_home_work_unit;NOT NULL;comment:'home-work distance unit'"`
// 	EmployeeType                string         `gorm:"column:employee_type;NOT NULL;comment:'employee type'"`
// 	Barcode                     sql.NullString `gorm:"column:barcode;comment:'badge id'"`
// 	Pin                         sql.NullString `gorm:"column:pin;comment:'pin'"`
// 	PrivateCarPlate             sql.NullString `gorm:"column:private_car_plate;comment:'private car plate'"`
// 	SpouseBirthdate             sql.NullString `gorm:"column:spouse_birthdate;comment:'spouse birthdate'"`
// 	Birthday                    sql.NullString `gorm:"column:birthday;comment:'date of birth'"`
// 	VisaExpire                  sql.NullString `gorm:"column:visa_expire;comment:'visa expiration date'"`
// 	WorkPermitExpirationDate    sql.NullString `gorm:"column:work_permit_expiration_date;comment:'work permit expiration date'"`
// 	DepartureDate               sql.NullString `gorm:"column:departure_date;comment:'departure date'"`
// 	EmployeeProperties          sql.NullString `gorm:"column:employee_properties;comment:'properties'"`
// 	AdditionalNote              sql.NullString `gorm:"column:additional_note;comment:'additional note'"`
// 	Notes                       sql.NullString `gorm:"column:notes;comment:'notes'"`
// 	DepartureDescription        sql.NullString `gorm:"column:departure_description;comment:'additional information'"`
// 	Active                      sql.NullBool   `gorm:"column:active;comment:'active'"`
// 	IsFlexible                  sql.NullBool   `gorm:"column:is_flexible;comment:'is flexible'"`
// 	IsFullyFlexible             sql.NullBool   `gorm:"column:is_fully_flexible;comment:'is fully flexible'"`
// 	WorkPermitScheduledActivity sql.NullBool   `gorm:"column:work_permit_scheduled_activity;comment:'work permit scheduled activity'"`
// 	CreateDate                  sql.NullString `gorm:"column:create_date;comment:'created on'"`
// 	WriteDate                   sql.NullString `gorm:"column:write_date;comment:'last updated on'"`
// 	AttendanceManagerId         sql.NullString `gorm:"column:attendance_manager_id;comment:'attendance manager'"`
// 	LastAttendanceId            sql.NullString `gorm:"column:last_attendance_id;comment:'last attendance'"`
// 	LastCheckIn                 sql.NullString `gorm:"column:last_check_in;comment:'check in'"`
// 	LastCheckOut                sql.NullString `gorm:"column:last_check_out;comment:'check out'"`
// 	ContractId                  sql.NullString `gorm:"column:contract_id;comment:'current contract'"`
// 	LegalName                   sql.NullString `gorm:"column:legal_name;comment:'legal name'"`
// 	Vehicle                     sql.NullString `gorm:"column:vehicle;comment:'company vehicle'"`
// 	FirstContractDate           sql.NullString `gorm:"column:first_contract_date;comment:'first contract date'"`
// 	ContractWarning             sql.NullBool   `gorm:"column:contract_warning;comment:'contract warning'"`
// 	ExpenseManagerId            sql.NullString `gorm:"column:expense_manager_id;comment:'expense'"`
// 	LeaveManagerId              sql.NullString `gorm:"column:leave_manager_id;comment:'time off'"`
// 	HourlyCost                  sql.NullString `gorm:"column:hourly_cost;comment:'hourly cost'"`
// 	PersonalMobile              sql.NullString `gorm:"column:personal_mobile;comment:'mobile'"`
// 	JoiningDate                 sql.NullString `gorm:"column:joining_date;comment:'joining date'"`
// 	IdExpiryDate                sql.NullString `gorm:"column:id_expiry_date;comment:'expiry date'"`
// 	PassportExpiryDate          sql.NullString `gorm:"column:passport_expiry_date;comment:'expiry date'"`
// 	ResignDate                  sql.NullString `gorm:"column:resign_date;comment:'resign date'"`
// 	Resigned                    sql.NullBool   `gorm:"column:resigned;comment:'resigned'"`
// 	Fired                       sql.NullBool   `gorm:"column:fired;comment:'fired'"`
// }
