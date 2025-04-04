package controllers

import (
	"encoding/json"
	"net/http"

	"company_db/backend/config"
	"company_db/backend/models"
	"company_db/backend/requests"
	"company_db/backend/responses"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// 企業データ登録エンドポイント
func CreateCompany(c *gin.Context) {
	var req requests.CreateCompanyRequest

	// リクエストの JSON をバインド. ここでbinding:"required"が指定されているので、リクエストのJSONに必須のフィールドが含まれていない場合はエラーを返す.
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.CreateCompanyErrorResponse{Error: err.Error()})
		return
	}

	// []string を JSON に変換
	urlJSON, err := json.Marshal(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.CreateCompanyErrorResponse{Error: "Failed to convert URL to JSON"})
		return
	}

	// postal_codeのIDを取得
	var postalCode models.PostalCode
	if err := config.DB.Where("postal_code = ?", req.PostalCode).First(&postalCode).Error; err != nil {
		c.JSON(http.StatusBadRequest, responses.CreateCompanyErrorResponse{Error: "Postal code not found"})
		return
	}

	// 企業データを作成
	company := models.Company{
		Name:         req.Name,
		URL:          datatypes.JSON(urlJSON), // ✅ 修正：[]byte に変換して代入
		PhoneNumber:  req.PhoneNumber,
		PostalCodeID: postalCode.ID,
		Address:      req.Address,
	}

	err = config.DB.Transaction(func(tx *gorm.DB) error {
		// トランザクション内は常に tx を使う！
		if err := tx.Create(&company).Error; err != nil {
			return err
		}

		for _, tagName := range req.Tags {
			var tag models.Tag
			if err := tx.Where("name = ?", tagName).First(&tag).Error; err != nil {
				tag = models.Tag{Name: tagName}
				if err := tx.Create(&tag).Error; err != nil {
					return err
				}
			}
			if err := tx.Create(&models.TagCompany{
				TagID:     tag.ID,
				CompanyID: company.ID,
			}).Error; err != nil {
				return err
			}
		}

		return nil // すべて成功 → commit
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company"})
		return
	}

	// 成功レスポンス
	c.JSON(http.StatusCreated, responses.CreateCompanyResponse{Message: "Company created successfully"})
}

func SearchCompany(c *gin.Context) {
	var req requests.SearchCompanyRequest

	// リクエストの JSON をバインド. ここでbinding:"required"が指定されているので、リクエストのJSONに必須のフィールドが含まれていない場合はエラーを返す.
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.CreateCompanyErrorResponse{Error: err.Error()})
		return
	}

	var companies []models.Company
	tx := config.DB.Model(&models.Company{}).
		Preload("PostalCode"). // 外部キーの郵便番号
		Preload("Tags")        // 多対多のタグ

	if req.Name != "" {
		tx = tx.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if len(req.URL) > 0 {
		tx = tx.Where("url LIKE ?", "%"+req.URL[0]+"%")
	}
	if req.PhoneNumber != "" {
		tx = tx.Where("phone_number LIKE ?", "%"+req.PhoneNumber+"%")
	}
	if req.PostalCode != "" {
		tx = tx.Where("postal_code_id IN (SELECT id FROM postal_codes WHERE postal_code LIKE ?)", "%"+req.PostalCode+"%")
	}
	if req.Address != "" {
		tx = tx.Where("address LIKE ?", "%"+req.Address+"%")
	}
	if len(req.Tags) > 0 {
		tx = tx.Joins("JOIN tag_companies ON tag_companies.company_id = companies.id").
			Joins("JOIN tags ON tags.id = tag_companies.tag_id").
			Where("tags.name IN (?)", req.Tags).
			Group("companies.id").
			Having("COUNT(tags.id) = ?", len(req.Tags))
	}

	if err := tx.Find(&companies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search company"})
		return
	}

	var companyData []responses.CompanyData
	for _, company := range companies {
		var urls []string
		json.Unmarshal(company.URL, &urls)
		data := responses.CompanyData{
			Name:        company.Name,
			URL:         urls,
			PhoneNumber: company.PhoneNumber,
			PostalCode:  company.PostalCode.PostalCode,
			Address:     company.Address,
			Tags:        make([]string, len(company.Tags)),
		}
		for i, tag := range company.Tags {
			data.Tags[i] = tag.Name
		}
		companyData = append(companyData, data)
	}
	// 成功レスポンス
	c.JSON(http.StatusOK, responses.CompanyResponse{
		Num:     len(companyData),
		Company: companyData,
	})
}
