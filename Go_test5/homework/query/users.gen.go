// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"Go_test5/homework/model"
)

func newPeople(db *gorm.DB) people {
	_people := people{}

	_people.peopleDo.UseDB(db)
	_people.peopleDo.UseModel(&model.People{})

	tableName := _people.peopleDo.TableName()
	_people.ALL = field.NewField(tableName, "*")
	_people.UUID = field.NewString(tableName, "uuid")
	_people.Name = field.NewString(tableName, "name")
	_people.Age = field.NewInt64(tableName, "age")
	_people.Version = field.NewInt64(tableName, "version")

	_people.fillFieldMap()

	return _people
}

type people struct {
	peopleDo

	ALL     field.Field
	UUID    field.String
	Name    field.String
	Age     field.Int64
	Version field.Int64

	fieldMap map[string]field.Expr
}

func (p people) Table(newTableName string) *people {
	p.peopleDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p people) As(alias string) *people {
	p.peopleDo.DO = *(p.peopleDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *people) updateTableName(table string) *people {
	p.ALL = field.NewField(table, "*")
	p.UUID = field.NewString(table, "uuid")
	p.Name = field.NewString(table, "name")
	p.Age = field.NewInt64(table, "age")
	p.Version = field.NewInt64(table, "version")

	p.fillFieldMap()

	return p
}

func (p *people) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *people) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 4)
	p.fieldMap["uuid"] = p.UUID
	p.fieldMap["name"] = p.Name
	p.fieldMap["age"] = p.Age
	p.fieldMap["version"] = p.Version
}

func (p people) clone(db *gorm.DB) people {
	p.peopleDo.ReplaceDB(db)
	return p
}

type peopleDo struct{ gen.DO }

//sql(SELECT COUNT(version) FROM users GROUP BY version ORDER BY version DESC LIMIT 1)
func (p peopleDo) GetMaxVersionCount() (result int, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT version FROM users GROUP BY version ORDER BY version DESC LIMIT 1 ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String()).Take(&result)
	err = executeSQL.Error
	return
}

func (p peopleDo) Debug() *peopleDo {
	return p.withDO(p.DO.Debug())
}

func (p peopleDo) WithContext(ctx context.Context) *peopleDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p peopleDo) Clauses(conds ...clause.Expression) *peopleDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p peopleDo) Returning(value interface{}, columns ...string) *peopleDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p peopleDo) Not(conds ...gen.Condition) *peopleDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p peopleDo) Or(conds ...gen.Condition) *peopleDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p peopleDo) Select(conds ...field.Expr) *peopleDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p peopleDo) Where(conds ...gen.Condition) *peopleDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p peopleDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *peopleDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p peopleDo) Order(conds ...field.Expr) *peopleDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p peopleDo) Distinct(cols ...field.Expr) *peopleDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p peopleDo) Omit(cols ...field.Expr) *peopleDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p peopleDo) Join(table schema.Tabler, on ...field.Expr) *peopleDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p peopleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *peopleDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p peopleDo) RightJoin(table schema.Tabler, on ...field.Expr) *peopleDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p peopleDo) Group(cols ...field.Expr) *peopleDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p peopleDo) Having(conds ...gen.Condition) *peopleDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p peopleDo) Limit(limit int) *peopleDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p peopleDo) Offset(offset int) *peopleDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p peopleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *peopleDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p peopleDo) Unscoped() *peopleDo {
	return p.withDO(p.DO.Unscoped())
}

func (p peopleDo) Create(values ...*model.People) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p peopleDo) CreateInBatches(values []*model.People, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p peopleDo) Save(values ...*model.People) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p peopleDo) First() (*model.People, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.People), nil
	}
}

func (p peopleDo) Take() (*model.People, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.People), nil
	}
}

func (p peopleDo) Last() (*model.People, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.People), nil
	}
}

func (p peopleDo) Find() ([]*model.People, error) {
	result, err := p.DO.Find()
	return result.([]*model.People), err
}

func (p peopleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.People, err error) {
	buf := make([]*model.People, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p peopleDo) FindInBatches(result *[]*model.People, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p peopleDo) Attrs(attrs ...field.AssignExpr) *peopleDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p peopleDo) Assign(attrs ...field.AssignExpr) *peopleDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p peopleDo) Joins(fields ...field.RelationField) *peopleDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p peopleDo) Preload(fields ...field.RelationField) *peopleDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p peopleDo) FirstOrInit() (*model.People, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.People), nil
	}
}

func (p peopleDo) FirstOrCreate() (*model.People, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.People), nil
	}
}

func (p peopleDo) FindByPage(offset int, limit int) (result []*model.People, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p peopleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p *peopleDo) withDO(do gen.Dao) *peopleDo {
	p.DO = *do.(*gen.DO)
	return p
}
