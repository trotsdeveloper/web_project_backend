package rest

import (
    "encoding/json"
    "net/http"
    "github.com/trotsdeveloper/web_project_backend/dao"
)

type RegisteredUsers []dao.RegisteredUser
type PaymentMethodType []dao.PaymentMethodType
type PaymentMethods []dao.PaymentMethod
type Purchases []dao.Purchase
type ProductCategories []dao.ProductCategory
type ProductOffers []dao.ProductOffer
type ProductOfferImgUrls []dao.ProductOfferImgUrl
type PurchaseDetails []dao.PurchaseDetail

type Response struct {
    Code string `json:"code"`
    Content interface{} `json:"content"`
}
// registered_users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
    user := dao.RegisteredUser{}
    var users RegisteredUsers
    var err error
    users, err = user.SelectAllInDB()
    response := Response{}
    if err != nil {
        response.Code = "200"
        response.Content = users       
    }
    respB, _ := json.Marshal(users)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])    
}
func GetUser(w http.ResponseWriter, r *http.Request) {}
func LoginUser(w http.ResponseWriter, r *http.Request) {}
func CreateUser(w http.ResponseWriter, r *http.Request) {}
func UpdateUser(w http.ResponseWriter, r *http.Request) {}
func DeleteUser(w http.ResponseWriter, r *http.Request) {}

// payment_method_types
func GetAllPaymentMethodTypes(w http.ResponseWriter, r *http.Request) {}
func GetPaymentMethodType(w http.ResponseWriter, r *http.Request) {}

// payment_methods
func GetAllPaymentMethods(w http.ResponseWriter, r *http.Request) {}
func GetPaymentMethod(w http.ResponseWriter, r *http.Request) {}
func CreatePaymentMethod(w http.ResponseWriter, r *http.Request) {}
func UpdatePaymentMethod(w http.ResponseWriter, r *http.Request) {}
func DeletePaymentMethod(w http.ResponseWriter, r *http.Request) {}

// purchases
func GetAllPurchases(w http.ResponseWriter, r *http.Request) {}
func GetPurchase(w http.ResponseWriter, r *http.Request) {}
func CreatePurchase(w http.ResponseWriter, r *http.Request) {}
func UpdatePurchase(w http.ResponseWriter, r *http.Request) {}
func DeletePurchase(w http.ResponseWriter, r *http.Request) {}

// product_categories 
func GetAllProductCategories(w http.ResponseWriter, r *http.Request) {}
func GetProductCategory(w http.ResponseWriter, r *http.Request) {}

// product_offers
func GetAllProductOffers(w http.ResponseWriter, r *http.Request) {}
func GetProductOffer(w http.ResponseWriter, r *http.Request) {}
func CreateProductOffer(w http.ResponseWriter, r *http.Request) {}
func UpdateProductOffer(w http.ResponseWriter, r *http.Request) {}
func DeleteProductOffer(w http.ResponseWriter, r *http.Request) {}

// product_offer_img_urls
func GetAllProductOfferImgUrls(w http.ResponseWriter, r *http.Request) {}
func CreateProductOfferImgUrl(w http.ResponseWriter, r *http.Request) {}
func UpdateProductOfferImgUrl(w http.ResponseWriter, r *http.Request) {}
func DeleteProductOfferImgUrl(w http.ResponseWriter, r *http.Request) {}

// purchase_details
func GetAllPurchaseDetails(w http.ResponseWriter, r *http.Request) {}
func GetPurchaseDetail(w http.ResponseWriter, r *http.Request) {}
func CreatePurchaseDetail(w http.ResponseWriter, r *http.Request) {}
func UpdatePurchaseDetail(w http.ResponseWriter, r *http.Request) {}
func DeletePurchaseDetail(w http.ResponseWriter, r *http.Request) {}
