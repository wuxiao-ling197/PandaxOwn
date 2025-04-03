package tool

// 组织树处理
import (
	"pandax/apps/shared/entity"
)

// 创建完整的组织树结构 结构见ResCompanyB结构体 但实际不好操作
func BuildFullOrganizationTree(companies []entity.ResCompanyB, departments []entity.HrDepartment) ([]entity.ResCompanyB, error) {
	// 构建公司ID到公司的映射
	companyMap := make(map[int64]*entity.ResCompanyB)
	for i := range companies {
		company := &companies[i]
		companyMap[company.ID] = company
	}
	// 构建部门ID到部门的映射
	deptMap := make(map[int64]*entity.HrDepartment)
	// 按公司ID分组的部门列表
	companyDeptMap := make(map[int64][]entity.HrDepartment)
	for i := range departments {
		dept := &departments[i]
		deptMap[dept.ID] = dept
		companyDeptMap[dept.CompanyId] = append(companyDeptMap[dept.CompanyId], *dept)
	}
	// 为每个公司构建部门树
	for companyId, depts := range companyDeptMap {
		if company, exists := companyMap[companyId]; exists {
			company.Departments = BuildDepartmentTree(depts, deptMap)
		}
	}
	// 构建公司树
	var roots []entity.ResCompanyB
	for _, company := range companies {
		if company.ParentId == 0 { // 顶级公司
			roots = append(roots, *buildCompanyTree(companyMap, company))
		}
	}

	return roots, nil
}

// 创建公司树结构
func buildCompanyTree(companyMap map[int64]*entity.ResCompanyB, company entity.ResCompanyB) *entity.ResCompanyB {
	node := &entity.ResCompanyB{
		ID:          company.ID,
		Name:        company.Name,
		ParentId:    company.ParentId,
		ParentPath:  company.ParentPath,
		Email:       company.Email,
		Phone:       company.Phone,
		Active:      company.Active,
		Children:    make([]entity.ResCompanyB, 0),
		Departments: company.Departments, // 保留已构建的部门树
	}

	// node := new(entity.ResCompanyB)
	// node = &company

	// 查找所有子公司
	for _, child := range companyMap {
		if child.ParentId == company.ID {
			node.Children = append(node.Children, *buildCompanyTree(companyMap, *child))
		}
	}

	return node
}

// 创建部门树结构
func BuildDepartmentTree(departments []entity.HrDepartment, deptMap map[int64]*entity.HrDepartment) []entity.HrDepartment {
	// 记录顶级部门
	var roots []entity.HrDepartment
	for _, dept := range departments {
		if dept.ParentId == 0 { // 顶级部门
			roots = append(roots, *buildDepartmentSubtree(deptMap, dept))
		}
	}
	return roots
}

// 创建部门的子级部门树结构
func buildDepartmentSubtree(deptMap map[int64]*entity.HrDepartment, dept entity.HrDepartment) *entity.HrDepartment {
	node := &entity.HrDepartment{
		ID:                 dept.ID,
		Name:               dept.Name,
		CompanyId:          dept.CompanyId,
		ParentId:           dept.ParentId,
		ParentPath:         dept.ParentPath,
		ManagerId:          dept.ManagerId,
		MasterDepartmentId: dept.MasterDepartmentId,
		Active:             dept.Active,
		Color:              dept.Color,
		CreateUid:          dept.CreateUid,
		CreateDate:         dept.CreateDate,
		WriteUid:           dept.WriteUid,
		WriteDate:          dept.WriteDate,
		Children:           make([]entity.HrDepartment, 0),
	}

	// 查找所有子部门
	for _, child := range deptMap {
		if child.ParentId == dept.ID {
			node.Children = append(node.Children, *buildDepartmentSubtree(deptMap, *child))
		}
	}
	return node
}
