package entity

type HrJob struct {
	Id                            int64  `gorm:"column:id;default:nextval(hr_job_id_seq::regclass);NOT NULL" json:"id,omitempty"`
	Sequence                      int64  `gorm:"column:sequence" json:"sequence,omitempty"`
	ExpectedEmployees             int64  `gorm:"column:expected_employees" json:"expected_employees,omitempty"`
	NoOfEmployee                  int64  `gorm:"column:no_of_employee" json:"no_of_employee,omitempty"`
	NoOfRecruitment               int64  `gorm:"column:no_of_recruitment" json:"no_of_recruitment,omitempty"`
	DepartmentId                  int64  `gorm:"column:department_id" json:"department_id,omitempty"`
	CompanyId                     int64  `gorm:"column:company_id" json:"company_id,omitempty"`
	ContractTypeId                int64  `gorm:"column:contract_type_id" json:"contract_type_id,omitempty"`
	CreateUid                     int64  `gorm:"column:create_uid" json:"create_uid,omitempty"`
	WriteUid                      int64  `gorm:"column:write_uid" json:"write_uid,omitempty"`
	Name                          string `gorm:"column:name;NOT NULL" json:"name,omitempty"`
	Description                   string `gorm:"column:description" json:"description,omitempty"`
	Requirements                  string `gorm:"column:requirements" json:"requirements,omitempty"`
	Active                        bool   `gorm:"column:active" json:"active,omitempty"`
	CreateDate                    string `gorm:"column:create_date" json:"create_date,omitempty"`
	WriteDate                     string `gorm:"column:write_date" json:"write_date,omitempty"`
	AliasId                       int64  `gorm:"column:alias_id;NOT NULL" json:"alias_id,omitempty"`
	AddressId                     int64  `gorm:"column:address_id" json:"address_id,omitempty"`
	ManagerId                     int64  `gorm:"column:manager_id" json:"manager_id,omitempty"`
	UserID                        int64  `gorm:"column:user_id,comment:'recuiter'" json:"user_id,omitempty"`
	Color                         int64  `gorm:"column:color" json:"color,omitempty"`
	IndustryId                    int64  `gorm:"column:industry_id" json:"industry_id,omitempty"`
	NoOfHiredEmployee             int64  `gorm:"column:no_of_hired_employee" json:"no_of_hired_employee,omitempty"`
	DateFrom                      string `gorm:"column:date_from" json:"date_from,omitempty"`
	DateTo                        string `gorm:"column:date_to" json:"date_to,omitempty"`
	JobProperties                 string `gorm:"column:job_properties" json:"job_properties,omitempty"`
	ApplicantPropertiesDefinition string `gorm:"column:applicant_properties_definition" json:"applicant_properties_definition,omitempty"`
	SurveyId                      int64  `gorm:"column:survey_id" json:"survey_id,omitempty"`
	WebsiteId                     int64  `gorm:"column:website_id" json:"website_id,omitempty"`
	WebsiteMetaOgImg              string `gorm:"column:website_meta_og_img" json:"website_meta_og_img,omitempty"`
	PublishedDate                 string `gorm:"column:published_date" json:"published_date,omitempty"`
	WebsiteMetaTitle              string `gorm:"column:website_meta_title" json:"website_meta_title,omitempty"`
	WebsiteMetaDescription        string `gorm:"column:website_meta_description" json:"website_meta_description,omitempty"`
	WebsiteMetaKeywords           string `gorm:"column:website_meta_keywords" json:"website_meta_keywords,omitempty"`
	SeoName                       string `gorm:"column:seo_name" json:"seo_name,omitempty"`
	WebsiteDescription            string `gorm:"column:website_description" json:"website_description,omitempty"`
	JobDetails                    string `gorm:"column:job_details" json:"job_details,omitempty"`
	IsPublished                   bool   `gorm:"column:is_published" json:"is_published,omitempty"`
}

// TableName 表名
func (HrJob) TableName() string {
	return "hr_job"
}
