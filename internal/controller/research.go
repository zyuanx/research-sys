package controller

import (
	"gin-research-sys/internal/form"
	"gin-research-sys/internal/model"
	"gin-research-sys/internal/response"
	"gin-research-sys/internal/service"
	"gin-research-sys/internal/util"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type IResearchController interface {
	List(c *gin.Context)
	Create(c *gin.Context)
	Retrieve(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)

	Square(c *gin.Context)

	DownloadExcel(ctx *gin.Context)
}
type ResearchController struct{}

func NewResearchController() IResearchController {
	return ResearchController{}
}

var researchServices = service.NewResearchService()

//var researchMgoServices = service.NewResearchMgoService()

func (r ResearchController) List(ctx *gin.Context) {
	pagination := form.Pagination{}
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		log.Println(err.Error())
		util.Success(ctx, nil, "query error")
		return
	}

	var researches []model.Research
	var total int64

	// not admin, retrieve created by yourself
	query := make(map[string]interface{})
	claims := jwt.ExtractClaims(ctx)
	id := int(claims["id"].(float64))
	username := claims["username"].(string)
	if username != "admin" {
		query["user_id"] = id
	}
	if err := researchServices.List(pagination.Page, pagination.Size, &researches, &total, query); err != nil {
		log.Println(err.Error())
		util.Success(ctx, nil, "list error")
	}
	util.Success(ctx, gin.H{
		"page":    pagination.Page,
		"size":    pagination.Size,
		"results": researches,
		"total":   total,
	}, "")
}

func (r ResearchController) Retrieve(ctx *gin.Context) {
	//idString := ctx.Param("id")
	//id, err := strconv.Atoi(idString)
	//if err != nil {
	//	log.Println(err.Error())
	//	util.Fail(ctx, gin.H{}, "param is error")
	//	return
	//}
	//research := model.Research{}
	//if err = researchServices.Retrieve(&research, id); err != nil {
	//	log.Println(err.Error())
	//	util.Fail(ctx, gin.H{}, "retrieve1 fail")
	//	return
	//}
	//researchMgo := model.ResearchMgo{}
	//if err = researchMgoServices.Retrieve(&researchMgo, research.ResearchID); err != nil {
	//	log.Println(err.Error())
	//	util.Fail(ctx, gin.H{}, "retrieve2 fail")
	//	return
	//}
	//
	//util.Success(ctx, gin.H{"research": gin.H{
	//	"id":          research.ID,
	//	"title":       research.Title,
	//	"desc":        research.Desc,
	//	"status":      research.Status,
	//	"once":        research.Once,
	//	"researchID":  research.ResearchID,
	//	"detail":      researchMgo.Detail,
	//	"rules":       researchMgo.Rules,
	//	"fieldsValue": researchMgo.FieldsValue,
	//	"creator":     research.UserID,
	//}}, "")
}

func (r ResearchController) Create(ctx *gin.Context) {
	//createForm := form.ResearchCreateForm{}
	//if err := ctx.ShouldBindJSON(&createForm); err != nil {
	//	util.Fail(ctx, gin.H{}, "payload error")
	//	return
	//}
	//// there needs mongo transaction
	//researchMgo := model.ResearchMgo{
	//	Detail:      createForm.Detail,
	//	Rules:       createForm.Rules,
	//	FieldsValue: createForm.FieldsValue,
	//}
	//result, err := researchMgoServices.Create(&researchMgo)
	//if err != nil {
	//	util.Fail(ctx, gin.H{}, "create error")
	//	return
	//}
	//claims := jwt.ExtractClaims(ctx)
	//id := int(claims["id"].(float64))
	//research := model.Research{
	//	Title:  createForm.Title,
	//	Desc:   createForm.Desc,
	//	Access: createForm.Access,
	//	Once:   createForm.Once,
	//	UserID: id,
	//}
	//if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
	//	research.ResearchID = oid.Hex()
	//} else {
	//	log.Println(ok)
	//	util.Fail(ctx, gin.H{}, "create fail")
	//	return
	//}
	//if err = researchServices.Create(&research); err != nil {
	//	log.Println(err.Error())
	//	util.Fail(ctx, gin.H{}, "create fail")
	//	return
	//}
	//util.Success(ctx, gin.H{}, "create success")
}

func (r ResearchController) Update(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "param is error")
		return
	}
	updateForm := form.ResearchUpdateForm{}
	if err = ctx.ShouldBindJSON(&updateForm); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "payload is error")
		return
	}
	research := model.Research{}
	if err = researchServices.Retrieve(&research, id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "retrieve fail")
		return
	}
	research.Title = updateForm.Title
	research.Desc = updateForm.Desc
	research.Once = updateForm.Once
	research.Status = updateForm.Status
	if err = researchServices.Update(&research); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "update fail")
		return
	}
	//res.Success(ctx, gin.H{}, "update success")
	//research := bson.M{}
	//if err := researchServices.Update(idString, research); err != nil {
	//	res.Fail(ctx, gin.H{}, "update fail")
	//	log.Println(err.Error())
	//	return
	//}
	util.Success(ctx, gin.H{}, "update success")
}

func (r ResearchController) Destroy(ctx *gin.Context) {
	panic("implement me")
}

func (r ResearchController) Square(ctx *gin.Context) {
	pagination := form.Pagination{}
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		log.Println(err.Error())
		util.Success(ctx, nil, "query error")
		return
	}

	claims := jwt.ExtractClaims(ctx)
	id := int(claims["id"].(float64))
	user := model.User{}
	if err := userServices.Retrieve(&user, id); err != nil {
		util.Fail(ctx, gin.H{}, "record not found")
		return
	}

	var researches []response.ResearchResponse
	var total int64

	query := make(map[string]interface{})
	query["status"] = 1
	query["access"] = user.College
	if err := researchServices.FindByAccess(pagination.Page, pagination.Size, &researches, &total, query); err != nil {
		log.Println(err.Error())
		util.Success(ctx, nil, "list error")
	}
	util.Success(ctx, gin.H{
		"page":    pagination.Page,
		"size":    pagination.Size,
		"results": researches,
		"total":   total,
	}, "")
}

func (r ResearchController) DownloadExcel(ctx *gin.Context) {
	//idString := ctx.Param("id")
	//// get research
	//researchMgo := model.ResearchMgo{}
	//if err := researchMgoServices.Retrieve(&researchMgo, idString); err != nil {
	//	log.Println(err.Error())
	//	util.Fail(ctx, gin.H{}, "retrieve2 fail")
	//	return
	//}
	//// set the excel title line
	//var fields []string
	//titleRow := make([]interface{}, 0)
	//titleRow = append(titleRow, "用户名", "IP地址", "填写时间")
	//for _, v := range researchMgo.Detail {
	//	titleRow = append(titleRow, v["label"].(string))
	//	fields = append(fields, v["fieldId"].(string))
	//}
	//
	//// get record list
	//var records []model.Record
	//var total int64
	//if err := recordServices.ListID(idString, &records, &total); err != nil {
	//	log.Println(err.Error())
	//	util.Success(ctx, nil, "list error")
	//	return
	//}
	//
	//// 1. start generate excel
	//xlsx := excelize.NewFile()
	//// 2. new StreamWriter
	//streamWriter, err := xlsx.NewStreamWriter("Sheet1")
	//if err != nil {
	//	println(err.Error())
	//}
	//if _, err = xlsx.NewStyle(`{"font":{"color":"#777777"}}`); err != nil {
	//	println(err.Error())
	//}
	//// 3. write title
	//if err = streamWriter.SetRow("A1", titleRow); err != nil {
	//	return
	//}
	//
	//// 4. write record data
	//for k, v := range records {
	//	row := make([]interface{}, 0)
	//	row = append(row, v.User.Username)
	//	row = append(row, v.IP)
	//	row = append(row, v.CreatedAt.Format("2006-01-02 15:04:05"))
	//	for colID := 0; colID < len(fields); colID++ {
	//		mgo := model.RecordMgo{}
	//		if err = recordMgoServices.Retrieve(&mgo, v.RecordID); err != nil {
	//			println(err.Error())
	//		}
	//		row = append(row, mgo.FieldsValue[fields[colID]])
	//	}
	//	cell, _ := excelize.CoordinatesToCellName(1, k+2)
	//	if err = streamWriter.SetRow(cell, row); err != nil {
	//		println(err.Error())
	//	}
	//}
	//// 5. flush streamWriter
	//if err = streamWriter.Flush(); err != nil {
	//	println(err.Error())
	//}
	//ctx.Header("response-type", "blob")
	//data, _ := xlsx.WriteToBuffer()
	//ctx.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", data.Bytes())
}

//func (r ResearchController) MgoRetrieve(ctx *gin.Context) {
//	idString := ctx.Param("id")
//
//	researchMgo := model.ResearchMgo{}
//	if err := researchMgoServices.Retrieve(&researchMgo, idString); err != nil {
//		log.Println(err.Error())
//		util.Fail(ctx, gin.H{}, "retrieve2 fail")
//		return
//	}
//
//	util.Success(ctx, gin.H{"research": gin.H{
//		"detail":      researchMgo.Detail,
//		"rules":       researchMgo.Rules,
//		"fieldsValue": researchMgo.FieldsValue,
//	}}, "")
//}
