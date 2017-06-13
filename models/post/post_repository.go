package post

import (
	"github.com/astaxie/beego/orm"
)

type PostRepository struct {
}

func init() {
	orm.RegisterModel(new(Post))
}

func (this *PostRepository) FindAll() ([]*Post, error) {
	o := orm.NewOrm()
	var posts []*Post
	_, err := o.QueryTable("post").All(&posts)
	return posts, err
}

func (this *PostRepository) FindById(id int64) (*Post, error) {
	o := orm.NewOrm()
	post := Post{Id: id}
	err := o.Read(&post)
	return &post, err
}

func (this *PostRepository) Save(p *Post) error {

	var err error
	o := orm.NewOrm()

	if p.Id != 0 {
		err = o.Read(p)
		if err != nil {
			_, err = o.Update(p)
		}
	} else {
		_, err = o.Insert(p)
	}

	return err
}

func (this *PostRepository) Delete(p *Post) error {

	o := orm.NewOrm()
	_, err := o.Delete(p)
	return err
}
