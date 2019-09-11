package models

import (
	//"fmt"
	//"os"

	"github.com/astaxie/beego/orm"
)

type Name600 struct {
	StockName
}

func (a *Name600) TableName() string {
	return TableName("name_600")
}

func Name600Add(a *Name600) (int64, error) {
	o1 := orm.NewOrm()
	o1.Using("hcs")
	return o1.Insert(a)
}

func Name600GetList() ([]*Name600, int64) {
	list := make([]*Name600, 0)
	o1 := orm.NewOrm()
	o1.Using("hcs")
	query := o1.QueryTable(TableName("name_600"))
	total, _ := query.Count()
	//fmt.Fprintf(os.Stdout, "total[%d]\n", total)

	query.OrderBy("-id").All(&list)
	//for id := 0; id < len(list); id++ {
	//	fmt.Fprintf(os.Stdout, "ID[%d], [%s]\n", id, list[id])
	//}

	return list, total
}

func Name600GetById(id int) (*Name600, error) {
	name := new(Name600)
	o1 := orm.NewOrm()
	o1.Using("hcs")
	err := o1.QueryTable(TableName("name_600")).Filter("id", id).One(name)
	if err != nil {
		return nil, err
	}
	return name, nil
}

func Name600GetByCode(code string) (*Name600, error) {
	name := new(Name600)
	o1 := orm.NewOrm()
	o1.Using("hcs")
	err := o1.QueryTable(TableName("name_600")).Filter("code", code).One(name)
	if err != nil {
		return nil, err
	}
	return name, nil
}

func (a *Name600) Update(fields ...string) error {
	o1 := orm.NewOrm()
	o1.Using("hcs")
	if _, err := o1.Update(a, fields...); err != nil {
		return err
	}
	return nil
}
