package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"{package_name}/{project_name}/models"
	"{package_name}/{project_name}/pkg/setting"
	"{package_name}/{project_name}/pkg/util"
)

func Get{entity_capital_plural}(c *gin.Context) {
	maps := make(map[string]interface{})
	//models.InitTable()
	data, err := models.Get{entity_capital_plural}(util.GetPage(c), setting.PageSize, maps)
	if err == nil {
		c.JSON(http.StatusOK, data)
	} else {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func Get{entity_capital}(c *gin.Context) {
	id := c.Param("id")

	{entity}, err := models.Get{entity_capital}(id)
	if err == nil {
		c.JSON(http.StatusOK, {entity})
	} else if gorm.IsRecordNotFoundError(err) {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func Add{entity_capital}(c *gin.Context) {
	{entity} := models.{entity_capital}{}

	if err := c.ShouldBind(&{entity}); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	rtnVal, err := models.Add{entity_capital}(&{entity})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, rtnVal)
}

func Edit{entity_capital}(c *gin.Context) {
	{entity} := models.{entity_capital}{}

	if err := c.ShouldBind(&{entity}); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	rtnVal, err := models.Edit{entity_capital}(&{entity})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, rtnVal)
}

func Delete{entity_capital}(c *gin.Context) {
	id := c.Param("id")

	err := models.Delete{entity_capital}(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.AbortWithStatus(http.StatusNoContent)
}
