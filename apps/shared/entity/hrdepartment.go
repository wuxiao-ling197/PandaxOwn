package entity

// ResCompany  被引用的表是主表，外键存在的表是从表 所以应该定义在主表中
type ResCompanyB struct {
	ID       int64        `gorm:"primaryKey;column:id;default:nextval(res_company_id_seq::regclass);NOT NULL"`
	Name     string       `gorm:"column:name;NOT NULL" json:"name,omitempty"`
	ParentId int64        `gorm:"column:parent_id" json:"parent_id,omitempty"`
	Email    string       `gorm:"column:email" json:"email,omitempty"`
	Phone    string       `gorm:"column:phone" json:"phone,omitempty"`
	Active   bool         `gorm:"column:active" json:"active,omitempty"`
	Children []ResCompany `gorm:"-" json:"children,omitempty"`
}

type ResCompany struct {
	ResCompanyB
	// ID                     int64  `gorm:"primaryKey;column:id;default:nextval(res_company_id_seq::regclass);NOT NULL"`
	// Name                   string `gorm:"column:name;NOT NULL" json:"name,omitempty"`
	PartnerId  int64  `gorm:"column:partner_id;NOT NULL" json:"partner_id,omitempty"`
	CurrencyId int64  `gorm:"column:currency_id;NOT NULL" json:"currency_id,omitempty"`
	Sequence   int64  `gorm:"column:sequence" json:"sequence,omitempty"`
	CreateDate string `gorm:"column:create_date" json:"create_date,omitempty"`
	ParentPath string `gorm:"column:parent_path" json:"parent_path,omitempty"`
	// ParentId               int64  `gorm:"column:parent_id" json:"parent_id,omitempty"`
	PaperformatId          int64 `gorm:"column:paperformat_id" json:"paperformat_id,omitempty"`
	ExternalReportLayoutId int64 `gorm:"column:external_report_layout_id" json:"external_report_layout_id,omitempty"`
	CreateUid              int64 `gorm:"column:create_uid" json:"create_uid,omitempty"`
	WriteUid               int64 `gorm:"column:write_uid" json:"write_uid,omitempty"`
	// Email                  string `gorm:"column:email" json:"email,omitempty"`
	// Phone                  string `gorm:"column:phone" json:"phone,omitempty"`
	Mobile           string `gorm:"column:mobile" json:"mobile,omitempty"`
	Font             string `gorm:"column:font" json:"font,omitempty"`
	PrimaryColor     string `gorm:"column:primary_color;" json:"primary_color,omitempty"`
	SecondaryColor   string `gorm:"column:secondary_color" json:"secondary_color,omitempty"`
	LayoutBackground string `gorm:"column:layout_background;NOT NULL" json:"layout_background,omitempty"`
	ReportHeader     string `gorm:"column:report_header" json:"report_header,omitempty"`
	ReportFooter     string `gorm:"column:report_footer" json:"report_footer,omitempty"`
	CompanyDetails   string `gorm:"column:company_details" json:"company_details,omitempty"`
	// Active                                      bool           `gorm:"column:active" json:"active,omitempty"`
	UsesDefaultLogo                             bool           `gorm:"column:uses_default_logo" json:"uses_default_logo,omitempty"`
	WriteDate                                   string         `gorm:"column:write_date" json:"write_date,omitempty"`
	LogoWeb                                     string         `gorm:"column:logo_web" json:"logo_web,omitempty"`
	NomenclatureId                              int64          `gorm:"column:nomenclature_id" json:"nomenclature_id,omitempty"`
	ResourceCalendarId                          int64          `gorm:"column:resource_calendar_id" json:"resource_calendar_id,omitempty"`
	AliasDomainId                               int64          `gorm:"column:alias_domain_id" json:"alias_domain_id,omitempty"`
	AliasDomainName                             string         `gorm:"column:alias_domain_name" json:"alias_domain_name,omitempty"`
	EmailPrimaryColor                           string         `gorm:"column:email_primary_color;" json:"email_primary_color,omitempty"`
	EmailSecondaryColor                         string         `gorm:"column:email_secondary_color" json:"email_secondary_color,omitempty"`
	PartnerGid                                  int64          `gorm:"column:partner_gid" json:"partner_gid,omitempty"`
	IapEnrichAutoDone                           bool           `gorm:"column:iap_enrich_auto_done" json:"iap_enrich_auto_done,omitempty"`
	SnailmailColor                              bool           `gorm:"column:snailmail_color" json:"snailmail_color,omitempty"`
	SnailmailCover                              bool           `gorm:"column:snailmail_cover" json:"snailmail_cover,omitempty"`
	SnailmailDuplex                             bool           `gorm:"column:snailmail_duplex" json:"snailmail_duplex,omitempty"`
	PaymentOnboardingPaymentMethod              string         `gorm:"column:payment_onboarding_payment_method" json:"payment_onboarding_payment_method,omitempty"`
	FiscalyearLastDay                           string         `gorm:"column:fiscalyear_last_day;NOT NULL" json:"fiscalyear_last_day,omitempty"`
	TransferAccountId                           int64          `gorm:"column:transfer_account_id" json:"transfer_account_id,omitempty"`
	DefaultCashDifferenceIncomeAccountId        int64          `gorm:"column:default_cash_difference_income_account_id" json:"default_cash_difference_income_account_id,omitempty"`
	DefaultCashDifferenceExpenseAccountId       int64          `gorm:"column:default_cash_difference_expense_account_id" json:"default_cash_difference_expense_account_id,omitempty"`
	AccountJournalSuspenseAccountId             int64          `gorm:"column:account_journal_suspense_account_id" json:"account_journal_suspense_account_id,omitempty"`
	AccountJournalEarlyPayDiscountGainAccountId int64          `gorm:"column:account_journal_early_pay_discount_gain_account_id" json:"account_journal_early_pay_discount_gain_account_id,omitempty"`
	AccountJournalEarlyPayDiscountLossAccountId int64          `gorm:"column:account_journal_early_pay_discount_loss_account_id" json:"account_journal_early_pay_discount_loss_account_id,omitempty"`
	AccountSaleTaxId                            int64          `gorm:"column:account_sale_tax_id" json:"account_sale_tax_id,omitempty"`
	AccountPurchaseTaxId                        int64          `gorm:"column:account_purchase_tax_id" json:"account_purchase_tax_id,omitempty"`
	CurrencyExchangeJournalId                   int64          `gorm:"column:currency_exchange_journal_id" json:"currency_exchange_journal_id,omitempty"`
	IncomeCurrencyExchangeAccountId             int64          `gorm:"column:income_currency_exchange_account_id" json:"income_currency_exchange_account_id,omitempty"`
	ExpenseCurrencyExchangeAccountId            int64          `gorm:"column:expense_currency_exchange_account_id" json:"expense_currency_exchange_account_id,omitempty"`
	IncotermId                                  int64          `gorm:"column:incoterm_id" json:"incoterm_id,omitempty"`
	BatchPaymentSequenceId                      int64          `gorm:"column:batch_payment_sequence_id" json:"batch_payment_sequence_id,omitempty"`
	AccountOpeningMoveId                        int64          `gorm:"column:account_opening_move_id" json:"account_opening_move_id,omitempty"`
	AccountDefaultPosReceivableAccountId        int64          `gorm:"column:account_default_pos_receivable_account_id" json:"id,omitempty"`
	ExpenseAccrualAccountId                     int64          `gorm:"column:expense_accrual_account_id" json:"expense_accrual_account_id,omitempty"`
	RevenueAccrualAccountId                     int64          `gorm:"column:revenue_accrual_account_id" json:"revenue_accrual_account_id,omitempty"`
	AutomaticEntryDefaultJournalId              int64          `gorm:"column:automatic_entry_default_journal_id" json:"automatic_entry_default_journal_id,omitempty"`
	AccountFiscalCountryId                      int64          `gorm:"column:account_fiscal_country_id" json:"account_fiscal_country_id,omitempty"`
	TaxCashBasisJournalId                       int64          `gorm:"column:tax_cash_basis_journal_id" json:"tax_cash_basis_journal_id,omitempty"`
	AccountCashBasisBaseAccountId               int64          `gorm:"column:account_cash_basis_base_account_id" json:"account_cash_basis_base_account_id,omitempty"`
	AccountDiscountIncomeAllocationId           int64          `gorm:"column:account_discount_income_allocation_id" json:"account_discount_income_allocation_id,omitempty"`
	AccountDiscountExpenseAllocationId          int64          `gorm:"column:account_discount_expense_allocation_id" json:"account_discount_expense_allocation_id,omitempty"`
	FiscalyearLastMonth                         string         `gorm:"column:fiscalyear_last_month;NOT NULL" json:"fiscalyear_last_month,omitempty"`
	ChartTemplate                               string         `gorm:"column:chart_template" json:"chart_template,omitempty"`
	BankAccountCodePrefix                       string         `gorm:"column:bank_account_code_prefix" json:"bank_account_code_prefix,omitempty"`
	CashAccountCodePrefix                       string         `gorm:"column:cash_account_code_prefix" json:"cash_account_code_prefix,omitempty"`
	TransferAccountCodePrefix                   string         `gorm:"column:transfer_account_code_prefix" json:"transfer_account_code_prefix,omitempty"`
	TaxCalculationRoundingMethod                string         `gorm:"column:tax_calculation_rounding_method" json:"tax_calculation_rounding_method,omitempty"`
	TermsType                                   string         `gorm:"column:terms_type" json:"terms_type,omitempty"`
	QuickEditMode                               string         `gorm:"column:quick_edit_mode" json:"quick_edit_mode,omitempty"`
	AccountPriceInclude                         string         `gorm:"column:account_price_include;NOT NULL" json:"account_price_include,omitempty"`
	FiscalyearLockDate                          string         `gorm:"column:fiscalyear_lock_date" json:"fiscalyear_lock_date,omitempty"`
	TaxLockDate                                 string         `gorm:"column:tax_lock_date" json:"tax_lock_date,omitempty"`
	SaleLockDate                                string         `gorm:"column:sale_lock_date" json:"sale_lock_date,omitempty"`
	PurchaseLockDate                            string         `gorm:"column:purchase_lock_date" json:"purchase_lock_date,omitempty"`
	HardLockDate                                string         `gorm:"column:hard_lock_date" json:"hard_lock_date,omitempty"`
	AccountOpeningDate                          string         `gorm:"column:account_opening_date;NOT NULL" json:"account_opening_date,omitempty"`
	InvoiceTerms                                string         `gorm:"column:invoice_terms" json:"invoice_terms,omitempty"`
	InvoiceTermsHtml                            string         `gorm:"column:invoice_terms_html" json:"invoice_terms_html,omitempty"`
	ExpectsChartOfAccounts                      bool           `gorm:"column:expects_chart_of_accounts" json:"expects_chart_of_accounts,omitempty"`
	AngloSaxonAccounting                        bool           `gorm:"column:anglo_saxon_accounting" json:"anglo_saxon_accounting,omitempty"`
	QrCode                                      bool           `gorm:"column:qr_code" json:"qr_code,omitempty"`
	DisplayInvoiceAmountTotalWords              bool           `gorm:"column:display_invoice_amount_total_words" json:"display_invoice_amount_total_words,omitempty"`
	DisplayInvoiceTaxCompanyCurrency            bool           `gorm:"column:display_invoice_tax_company_currency" json:"display_invoice_tax_company_currency,omitempty"`
	AccountUseCreditLimit                       bool           `gorm:"column:account_use_credit_limit" json:"account_use_credit_limit,omitempty"`
	TaxExigibility                              bool           `gorm:"column:tax_exigibility" json:"tax_exigibility,omitempty"`
	AccountStorno                               bool           `gorm:"column:account_storno" json:"account_storno,omitempty"`
	CheckAccountAuditTrail                      bool           `gorm:"column:check_account_audit_trail" json:"check_account_audit_trail,omitempty"`
	AutopostBills                               bool           `gorm:"column:autopost_bills" json:"autopost_bills,omitempty"`
	HrPresenceControlEmailAmount                int64          `gorm:"column:hr_presence_control_email_amount" json:"hr_presence_control_email_amount,omitempty"`
	HrPresenceControlIpList                     string         `gorm:"column:hr_presence_control_ip_list" json:"hr_presence_control_ip_list,omitempty"`
	EmployeePropertiesDefinition                string         `gorm:"column:employee_properties_definition" json:"employee_properties_definition,omitempty"`
	HrPresenceControlLogin                      bool           `gorm:"column:hr_presence_control_login" json:"hr_presence_control_login,omitempty"`
	HrPresenceControlEmail                      bool           `gorm:"column:hr_presence_control_email" json:"hr_presence_control_email,omitempty"`
	HrPresenceControlIp                         bool           `gorm:"column:hr_presence_control_ip" json:"hr_presence_control_ip,omitempty"`
	HrPresenceControlAttendance                 bool           `gorm:"column:hr_presence_control_attendance" json:"hr_presence_control_attendance,omitempty"`
	OvertimeCompanyThreshold                    int64          `gorm:"column:overtime_company_threshold" json:"overtime_company_threshold,omitempty"`
	OvertimeEmployeeThreshold                   int64          `gorm:"column:overtime_employee_threshold" json:"iovertime_employee_thresholdd,omitempty"`
	AttendanceKioskDelay                        int64          `gorm:"column:attendance_kiosk_delay" json:"attendance_kiosk_delay,omitempty"`
	AttendanceKioskMode                         string         `gorm:"column:attendance_kiosk_mode" json:"attendance_kiosk_mode,omitempty"`
	AttendanceBarcodeSource                     string         `gorm:"column:attendance_barcode_source" json:"attendance_barcode_source,omitempty"`
	AttendanceKioskKey                          string         `gorm:"column:attendance_kiosk_key" json:"attendance_kiosk_key,omitempty"`
	AttendanceOvertimeValidation                string         `gorm:"column:attendance_overtime_validation" json:"attendance_overtime_validation,omitempty"`
	HrAttendanceDisplayOvertime                 bool           `gorm:"column:hr_attendance_display_overtime" json:"hr_attendance_display_overtime,omitempty"`
	AttendanceKioskUsePin                       bool           `gorm:"column:attendance_kiosk_use_pin" json:"attendance_kiosk_use_pin,omitempty"`
	AttendanceFromSystray                       bool           `gorm:"column:attendance_from_systray" json:"attendance_from_systray,omitempty"`
	AutoCheckOut                                bool           `gorm:"column:auto_check_out" json:"auto_check_out,omitempty"`
	AbsenceManagement                           bool           `gorm:"column:absence_management" json:"absence_management,omitempty"`
	AutoCheckOutTolerance                       float64        `gorm:"column:auto_check_out_tolerance" json:"auto_check_out_tolerance,omitempty"`
	ContractExpirationNoticePeriod              int64          `gorm:"column:contract_expiration_notice_period" json:"contract_expiration_notice_period,omitempty"`
	WorkPermitExpirationNoticePeriod            int64          `gorm:"column:work_permit_expiration_notice_period" json:"work_permit_expiration_notice_period,omitempty"`
	ExpenseJournalId                            int64          `gorm:"column:expense_journal_id" json:"expense_journal_id,omitempty"`
	ExpenseOutstandingAccountId                 int64          `gorm:"column:expense_outstanding_account_id" json:"expense_outstanding_account_id,omitempty"`
	CandidatePropertiesDefinition               string         `gorm:"column:candidate_properties_definition" json:"candidate_properties_definition,omitempty"`
	JobPropertiesDefinition                     string         `gorm:"column:job_properties_definition" json:"job_properties_definition,omitempty"`
	ProjectTimeModeId                           int64          `gorm:"column:project_time_mode_id" json:"project_time_mode_id,omitempty"`
	TimesheetEncodeUomId                        int64          `gorm:"column:timesheet_encode_uom_id" json:"timesheet_encode_uom_id,omitempty"`
	InternalProjectId                           int64          `gorm:"column:internal_project_id" json:"internal_project_id,omitempty"`
	LeaveTimesheetTaskId                        int64          `gorm:"column:leave_timesheet_task_id" json:"leave_timesheet_task_id,omitempty"`
	SocialTwitter                               string         `gorm:"column:social_twitter" json:"social_twitter,omitempty"`
	SocialFacebook                              string         `gorm:"column:social_facebook" json:"social_facebook,omitempty"`
	SocialGithub                                string         `gorm:"column:social_github" json:"social_github,omitempty"`
	SocialLinkedin                              string         `gorm:"column:social_linkedin" json:"social_linkedin,omitempty"`
	SocialYoutube                               string         `gorm:"column:social_youtube" json:"social_youtube,omitempty"`
	SocialInstagram                             string         `gorm:"column:social_instagram" json:"social_instagram,omitempty"`
	SocialTiktok                                string         `gorm:"column:social_tiktok" json:"social_tiktok,omitempty"`
	InternalTransitLocationId                   string         `gorm:"column:internal_transit_location_id" json:"internal_transit_location_id,omitempty"`
	StockMailConfirmationTemplateId             string         `gorm:"column:stock_mail_confirmation_template_id" json:"stock_mail_confirmation_template_id,omitempty"`
	AnnualInventoryDay                          int64          `gorm:"column:annual_inventory_day" json:"annual_inventory_day,omitempty"`
	AnnualInventoryMonth                        string         `gorm:"column:annual_inventory_month" json:"annual_inventory_month,omitempty"`
	StockMoveEmailValidation                    bool           `gorm:"column:stock_move_email_validation" json:"comstock_move_email_validationpany_id,omitempty"`
	AccountProductionWipAccountId               int64          `gorm:"column:account_production_wip_account_id" json:"account_production_wip_account_id,omitempty"`
	AccountProductionWipOverheadAccountId       int64          `gorm:"column:account_production_wip_overhead_account_id" json:"account_production_wip_overhead_account_id,omitempty"`
	StockSmsConfirmationTemplateId              int64          `gorm:"column:stock_sms_confirmation_template_id" json:"stock_sms_confirmation_template_id,omitempty"`
	StockMoveSmsValidation                      bool           `gorm:"column:stock_move_sms_validation" json:"stock_move_sms_validation,omitempty"`
	HasReceivedWarningStockSms                  bool           `gorm:"column:has_received_warning_stock_sms" json:"has_received_warning_stock_sms,omitempty"`
	SecurityLead                                float64        `gorm:"column:security_lead" json:"security_lead,omitempty"`
	PoLock                                      string         `gorm:"column:po_lock" json:"po_lock,omitempty"`
	PoDoubleValidation                          string         `gorm:"column:po_double_validation" json:"po_double_validation,omitempty"`
	PoDoubleValidationAmount                    float64        `gorm:"column:po_double_validation_amount" json:"po_double_validation_amount,omitempty"`
	PoLead                                      float64        `gorm:"column:po_lead" json:"po_lead,omitempty"`
	DaysToPurchase                              float64        `gorm:"column:days_to_purchase" json:"days_to_purchase,omitempty"`
	WebsiteId                                   int64          `gorm:"column:website_id" json:"website_id,omitempty"`
	QuotationValidityDays                       int64          `gorm:"column:quotation_validity_days" json:"quotation_validity_days,omitempty"`
	SaleDiscountProductId                       int64          `gorm:"column:sale_discount_product_id" json:"sale_discount_product_id,omitempty"`
	SaleOnboardingPaymentMethod                 string         `gorm:"column:sale_onboarding_payment_method" json:"sale_onboarding_payment_method,omitempty"`
	PortalConfirmationSign                      bool           `gorm:"column:portal_confirmation_sign" json:"portal_confirmation_sign,omitempty"`
	PortalConfirmationPay                       bool           `gorm:"column:portal_confirmation_pay" json:"portal_confirmation_pay,omitempty"`
	PrepaymentPercent                           float64        `gorm:"column:prepayment_percent" json:"prepayment_percent,omitempty"`
	SaleOrderTemplateId                         int64          `gorm:"column:sale_order_template_id" json:"sale_order_template_id,omitempty"`
	Departments                                 []HrDepartment `gorm:"foreignKey:CompanyId;" ` //references:ID json:"id,omitempty"
}

type CompanyWithDepatment struct {
	ResCompanyB
	Departments []HrDepartment `gorm:"foreignKey:CompanyId" ` //json:"department,omitempty"
}

// HrDepartment 是company的从表，employee的主表 对于外部关联如Employees创建时最好手打 因为复制过来的一直shit的报错
type HrDepartment struct {
	ID                 int64          `gorm:"primaryKey;column:id;default:nextval(hr_department_id_seq::regclass);NOT NULL" json:"id,omitempty"`
	CompanyId          int64          `gorm:"column:company_id;comment:'company'" json:"company_id,omitempty"`
	ParentId           int64          `gorm:"column:parent_id;comment:'parent department'" json:"parent_id,omitempty"`
	ManagerId          int64          `gorm:"column:manager_id;comment:'manager'" json:"idmanager_id,omitempty"`
	Color              int64          `gorm:"column:color;comment:'color index'" json:"color,omitempty"`
	MasterDepartmentId int64          `gorm:"column:master_department_id;comment:'master department'" json:"master_department_id,omitempty"`
	CreateUid          string         `gorm:"column:create_uid;comment:'created by'" json:"create_uid,omitempty"`
	WriteUid           string         `gorm:"column:write_uid;comment:'last updated by'" json:"write_uid,omitempty"`
	CompleteName       string         `gorm:"column:complete_name;comment:'complete name'" json:"complete_name,omitempty"`
	ParentPath         string         `gorm:"column:parent_path;comment:'parent path'" json:"parent_path,omitempty"`
	Name               string         `gorm:"column:name;NOT NULL;comment:'department name'" json:"name,omitempty"`
	Note               string         `gorm:"column:note;comment:'note'" json:"note,omitempty"`
	Active             bool           `gorm:"column:active;comment:'active'" json:"active,omitempty"`
	CreateDate         string         `gorm:"column:create_date;comment:'created on'" json:"create_date,omitempty"`
	WriteDate          string         `gorm:"column:write_date;comment:'last updated on'" json:"write_date,omitempty"`
	Employees          []HrEmployee   `gorm:"foreignKey:DepartmentId" json:"employees,omitempty"`
	Children           []HrDepartment `gorm:"-" json:"children,omitempty"`

	// 40添加字段
	// Property int64  `gorm:"column:property;comment:'科室属性'" json:"property,omitempty"`
	// DutyId   int64  `gorm:"column:duty_id;comment:'负责人'" json:"duty_id,omitempty"`
	// Virtual  bool   `gorm:"column:virtual;comment:'虚拟部门标志'" json:"virtual,omitempty"`
	// Code     string `gorm:"column:code;comment:'科室编码'" json:"code,omitempty"`
}

// 部门树
type DepartmentLable struct {
	DepartmentId   int64             `gorm:"-" json:"departmentId"`
	DepartmentName string            `gorm:"-" json:"departmentName"`
	CompanyId      int64             `gorm:"-" json:"companyId"`
	CompanyName    string            `gorm:"-" json:"companyName"`
	Children       []DepartmentLable `gorm:"-" json:"children"`
}

func (ResCompany) TableName() string {
	return "res_company" // 确保预加载preload时的表名正确
}
func (HrDepartment) TableName() string {
	return "hr_department"
}

func (CompanyWithDepatment) TableName() string {
	return "res_company" // 确保预加载preload时的表名正确
}

// type OdooName struct {
// 	ZHCN string `gorm:"-" json: "zh_CN"`
// 	ENUS string `gorm:"-" json: "en_US"`
// }
