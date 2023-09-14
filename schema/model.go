package schema

import "time"

type Post struct {
	ID        uint      `gorm:"column:Id;primaryKey;autoIncrement" json:"id"`
	Title     string    `gorm:"column:Title;type:varchar(100)" json:"title"`
	Content   string    `gorm:"column:Content;type:text" json:"content"`
	Category  string    `gorm:"column:Category;type:varchar(100)" json:"category"`
	CreatedAt time.Time `gorm:"column:Created_date;type:timestamp" json:"-"`
	UpdatedAt time.Time `gorm:"column:Updated_date;type:timestamp" json:"-"`
	Status    string    `gorm:"column:Status;type:varchar(100);comment:Publish,Draft,Thrash" json:"status"`
}

func (p *Post) FromRequest(r PostRequest) {
	p.Title = r.Title
	p.Content = r.Content
	p.Category = r.Category
	p.Status = r.Status
}
