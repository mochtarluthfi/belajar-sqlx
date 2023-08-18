package controllers

import (
	"belajar-sqlx/db_client"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Post struct {
	Cbranch string `json:"cbranch"`
	Cdesc   string `json:"cdesc"`
	CAkses  string `json:"cakses"`
	Cnama   string `json:"cnama"`
}

func CreatePost(c *gin.Context) {
	var reqBody Post
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "Masukan Kode Cabang",
		})
		return
	}

	hasil, err := getBranchData(reqBody.Cbranch)
	if err != nil {
		fmt.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Masukan Kode Branch Yang Benar",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  hasil,
	})
}

func getBranchData(cbranch string) ([]map[string]interface{}, error) {
	query := "SELECT a.cbranch, a.cdesc, b.cakses, b.cnama FROM t_branch a LEFT JOIN t_AksesBranch b ON a.cBranch = b.cBranch WHERE a.cbranch in (@p1)"

	rows, err := db_client.DBClient.Queryx(query, cbranch)
	if err != nil {
		return nil, err
	}
	data := []map[string]interface{}{}
	for rows.Next() {
		rowData := make(map[string]interface{})
		if err := rows.MapScan(rowData); err != nil {
			return nil, err
		}
		data = append(data, rowData)
	}

	return data, nil
}

func GetAllUser(c *gin.Context) {
	query := `select a.cbranch,a.cdesc,b.cakses,b.cnama  from t_branch a left join t_AksesBranch b on a.cBranch=b.cBranch`
	rows, err := db_client.DBClient.Queryx(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan":  "Ada error :" + err.Error(),
			"status": "0",
		})
		return
	}

	hasil := make([]map[string]interface{}, 0)
	for rows.Next() {
		m := make(map[string]interface{})
		rows.MapScan(m)
		//m["nama"] = strings.TrimSpace(m["name"].(string))
		//delete(m, "name")
		hasil = append(hasil, m)
	}

	c.JSON(http.StatusOK, gin.H{
		"pesan": "ok",
		"data":  hasil,
	})
}
