package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type input struct {
	Stuno string `json:"stu_no"`
	Name  string `json:"name"`
}

type output struct {
	CreatedAt string `json:"created_at"`
	Name      string `json:"name"`
	Stuno     string `json:"stu_no"`
	ID        uint   `json:"id"`
}

func (s *Service) getAllLost(c *gin.Context) {
	data := make([]*lost, 0, 100)
	out := make([]*output, 0, 100)
	s.DB.Model(data).Find(&data)

	for _, v := range data {
		out = append(out, &output{
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			Name:      v.Name,
			Stuno:     v.Stuno,
			ID:        v.ID,
		})
	}
	c.JSON(makeSuccessReturn(200, out))
	return
}

func (s *Service) getLost(c *gin.Context) {
	ID, err := analysisID(c)
	if err != nil {
		return
	}
	//fmt.Println(ID)

	temp := new(lost)
	s.DB.Where("id = ?",ID).Find(temp)
	if temp.Stuno == "" {
		c.JSON(makeErrorReturn(300, 30000, "doesn't exist"))
		return
	}

	out := new(output)
	out.ID = temp.ID
	out.Stuno = temp.Stuno
	out.Name =temp.Name
	out.CreatedAt = temp.CreatedAt.Format("2006-01-02 15:04:05")
	c.JSON(makeSuccessReturn(200, out))
	return
}

func (s *Service) modify(c *gin.Context) {
	//fmt.Println(1)
	ID, err := analysisID(c)
	//fmt.Println(1)
	if err != nil {
		//fmt.Println("1")
		return
	}

	//fmt.Println(1)
	temp := new(lost)
	s.DB.Model(&lost{}).Where("id = ?", ID).Find(temp)
	if temp.Stuno == "" {
		c.JSON(makeErrorReturn(400, 40000, "doesn't exist"))
		return
	}

	//fmt.Println(1)
	err = c.BindJSON(temp)
	if err != nil {
		DealError(err)
		c.JSON(makeErrorReturn(300, 30000, "json wrong format"))
		return
	}

	//fmt.Println(temp)
	//fmt.Println("1")
	tx := s.DB.Begin()
	if tx.Model(&lost{}).Where("id = ?", ID).Updates(lost{
		Stuno: temp.Stuno,
		Name:  temp.Name,
	}).RowsAffected != 1 {
		tx.Rollback()
		c.JSON(makeErrorReturn(500, 50000, "can't add it"))
		return
	}
	tx.Commit()
	c.JSON(makeSuccessReturn(200, ""))
	return
}

func (s *Service) delete(c *gin.Context) {
	ID, err := analysisID(c)
	if err != nil {
		return
	}

	temp := new(lost)
	s.DB.Model(&lost{}).Where("id = ?", ID).Find(temp)
	if temp.Stuno == "" {
		c.JSON(makeErrorReturn(400, 40000, "doesn't exist"))
		return
	}

	tx := s.DB.Begin()
	if tx.Model(&lost{}).Where("id = ?",ID).Delete(&lost{}).RowsAffected != 1 {
		tx.Rollback()
		c.JSON(makeErrorReturn(400, 40000, "can't delete it"))
		return
	}
	tx.Commit()
	c.JSON(makeSuccessReturn(200, ""))
	return
}

func (s *Service) add(c *gin.Context) {
	temp := new(input)
	err := c.BindJSON(temp)
	DealError(err)
	if err != nil {
		c.JSON(makeErrorReturn(400, 40000, "json wrong format"))
		return
	}

	tx := s.DB.Begin()
	if tx.Create(&lost{
		Stuno: temp.Stuno,
		Name:  temp.Name,
	}).RowsAffected != 1 {
		tx.Rollback()
		c.JSON(makeErrorReturn(500, 50000, "can't add it"))
		return
	}
	tx.Commit()
	c.JSON(makeSuccessReturn(200, ""))
	return
}

func makeSuccessReturn(status int, data interface{}) (int, interface{}) {
	return status, gin.H{
		"error": 0,
		"msg":   "success",
		"data":  data,
	}
}

func makeErrorReturn(status int, error int, msg string) (int, interface{}) {
	return status, gin.H{
		"error": error,
		"msg":   msg,
	}
}

func analysisID(c *gin.Context) (uint, error) {
	var id int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(makeErrorReturn(300, 30000, "url wrong format"))
		return uint(id), err
	} else {
		return uint(id), err
	}
}
