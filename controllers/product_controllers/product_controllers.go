package product_controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ihksanghazi/restful-api-gorillamux/helpers"
	"github.com/ihksanghazi/restful-api-gorillamux/models"
	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request){
	var products []models.Products

	if err := models.DB.Find(&products).Error; err != nil {
		helpers.ResponseError(w,http.StatusInternalServerError,err.Error())
		return
	}
	
	helpers.ResponseJSON(w,http.StatusOK,products)
	
}

func Show(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err:= strconv.ParseInt(vars["id"],10,64)
	if err != nil{
		helpers.ResponseError(w,http.StatusBadRequest,err.Error())
		return
	}

	var product models.Products
	if err:= models.DB.First(&product,id).Error; err != nil{
		switch err {
		case gorm.ErrRecordNotFound:
			helpers.ResponseError(w,http.StatusNotFound,"Product Tidak Ditemukan")
			return
		default:
			helpers.ResponseError(w,http.StatusInternalServerError,err.Error())
			return
		}
	}
	helpers.ResponseJSON(w,http.StatusOK,product)
}

func Create(w http.ResponseWriter, r *http.Request){
	var product models.Products

	decode:= json.NewDecoder(r.Body)
	if err:= decode.Decode(&product); err != nil {
		helpers.ResponseError(w,http.StatusInternalServerError,err.Error())
		return
	}

	defer r.Body.Close()

	if err:= models.DB.Create(&product).Error; err != nil {
		helpers.ResponseError(w,http.StatusInternalServerError,err.Error())
		return
	}

	helpers.ResponseJSON(w,http.StatusCreated,product)

}

func Update(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err:= strconv.ParseInt(vars["id"],10,64)
	if err != nil{
		helpers.ResponseError(w,http.StatusBadRequest,err.Error())
		return
	}

	var product models.Products
	decode:= json.NewDecoder(r.Body)
	if err:= decode.Decode(&product); err != nil {
		helpers.ResponseError(w,http.StatusInternalServerError,err.Error())
		return
	}

	defer r.Body.Close()

	if models.DB.Where("id = ?",id).Updates(&product).RowsAffected == 0 {
		helpers.ResponseError(w,http.StatusBadRequest,"Tidak Dapat Mengupdate Product")
		return
	}

	product.Id = id

	helpers.ResponseJSON(w,http.StatusOK,product)
}

func Delete(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err:= strconv.ParseInt(vars["id"],10,64)
	if err != nil{
		helpers.ResponseError(w,http.StatusBadRequest,err.Error())
		return
	}

	var product models.Products

	if models.DB.Delete(&product,id).RowsAffected == 0 {
		helpers.ResponseError(w,http.StatusBadRequest,"Tidak Dapat Menghapus Product")
		return
	}
	
	helpers.ResponseJSON(w,http.StatusOK, map[string]string{"msg":"Product Berhasil Dihapus"})
}