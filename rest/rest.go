package rest

import (
    "encoding/json"
    "net/http"
    "github.com/trotsdeveloper/web_project_backend/dao"
)

type RegisteredUsers []dao.RegisteredUser
type PaymentMethodTypes []dao.PaymentMethodType
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
    respB, err2 := json.Marshal(users)
    
    if err != nil {
        response.Code = "404"
        response.Content = err              
    } else if err2 != nil {
        response.Code = "404"
        response.Content = err2
    } else {
        response.Code = "200"
        response.Content = users
    }
    
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
func GetAllPaymentMethodTypes(w http.ResponseWriter, r *http.Request) {
    paymentMethodType := dao.PaymentMethodType{}
    var paymentMethodTypes PaymentMethodTypes
    var err error
    paymentMethodTypes, err = paymentMethodType.SelectAllInDB()
    response := Response{}
    respB, err2 := json.Marshal(paymentMethodTypes)
    
    if err != nil {
        response.Code = "404"
        response.Content = err              
    } else if err2 != nil {
        response.Code = "404"
        response.Content = err2
    } else {
        response.Code = "200"
        response.Content = paymentMethodTypes
    }
    
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])   
    
}
func GetPaymentMethodType(w http.ResponseWriter, r *http.Request) {}

// payment_methods
func GetAllPaymentMethods(w http.ResponseWriter, r *http.Request) {
    paymentMethod := dao.PaymentMethod{}
    var paymentMethods PaymentMethods
    var err error
    paymentMethods, err = paymentMethod.SelectAllInDB()
    response := Response{}   
    respB, err2 := json.Marshal(paymentMethods)
    
    if err != nil {
        response.Code = "404"
        response.Content = err              
    } else if err2 != nil {
        response.Code = "404"
        response.Content = err2
    } else {
        response.Code = "200"
        response.Content = paymentMethods
    }
    
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])   
}
func GetPaymentMethod(w http.ResponseWriter, r *http.Request) {}
func CreatePaymentMethod(w http.ResponseWriter, r *http.Request) {}
func UpdatePaymentMethod(w http.ResponseWriter, r *http.Request) {}
func DeletePaymentMethod(w http.ResponseWriter, r *http.Request) {}

// purchases
func GetAllPurchases(w http.ResponseWriter, r *http.Request) {
    purchase := dao.Purchase{}
    var purchases Purchases
    var err error
    purchases, err = purchase.SelectAllInDB()
    response := Response{}
    respB, err2 := json.Marshal(purchases)
    
    if err != nil {
        response.Code = "404"
        response.Content = err              
    } else if err2 != nil {
        response.Code = "404"
        response.Content = err2
    } else {
        response.Code = "200"
        response.Content = purchases
    }
    
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])   
}
func GetPurchase(w http.ResponseWriter, r *http.Request) {}
func CreatePurchase(w http.ResponseWriter, r *http.Request) {}
func UpdatePurchase(w http.ResponseWriter, r *http.Request) {}
func DeletePurchase(w http.ResponseWriter, r *http.Request) {}

// product_categories 
func GetAllProductCategories(w http.ResponseWriter, r *http.Request) {
    productCategory := dao.ProductCategory{}
    var productCategories ProductCategories
    var err error
    productCategories, err = productCategory.SelectAllInDB()
    response := Response{}    
    respB, err2 := json.Marshal(productCategories)
    
    if err != nil {
        response.Code = "404"
        response.Content = err              
    } else if err2 != nil {
        response.Code = "404"
        response.Content = err2
    } else {
        response.Code = "200"
        response.Content = productCategories
    }
    
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])
}
func GetProductCategory(w http.ResponseWriter, r *http.Request) {}

// product_offers
func GetAllProductOffers(w http.ResponseWriter, r *http.Request) {
    productOffer := dao.ProductOffer{}
    var products ProductOffers
    var err error
    products, err = productOffer.SelectAllInDB()
    response := Response{}    
    respB, err2 := json.Marshal(products)
    
    if err != nil {
        response.Code = "404"
        response.Content = err              
    } else if err2 != nil {
        response.Code = "404"
        response.Content = err2
    } else {
        response.Code = "200"
        response.Content = products
    }
    
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])
}
func GetProductOffer(w http.ResponseWriter, r *http.Request) {}
func CreateProductOffer(w http.ResponseWriter, r *http.Request) {}
func UpdateProductOffer(w http.ResponseWriter, r *http.Request) {}
func DeleteProductOffer(w http.ResponseWriter, r *http.Request) {}

// product_offer_img_urls
func GetAllProductOfferImgUrls(w http.ResponseWriter, r *http.Request) {
    productOfferImgUrl := dao.ProductOfferImgUrl{}
    var imgUrls ProductOfferImgUrls
    var err error
    imgUrls, err = productOfferImgUrl.SelectAllInDB()
    response := Response{}    
    respB, err2 := json.Marshal(imgUrls)
    
    if err != nil {
        response.Code = "404"
        response.Content = err              
    } else if err2 != nil {
        response.Code = "404"
        response.Content = err2
    } else {
        response.Code = "200"
        response.Content = imgUrls
    }
    
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])   
    
}
func CreateProductOfferImgUrl(w http.ResponseWriter, r *http.Request) {}
func UpdateProductOfferImgUrl(w http.ResponseWriter, r *http.Request) {}
func DeleteProductOfferImgUrl(w http.ResponseWriter, r *http.Request) {}

// purchase_details
func GetAllPurchaseDetails(w http.ResponseWriter, r *http.Request) {
    purchaseDetail := dao.PurchaseDetail{}
    var purchaseDetails PurchaseDetails
    var err error
    purchaseDetails, err = purchaseDetail.SelectAllInDB()    
    response := Response{}    
    respB, err2 := json.Marshal(purchaseDetails)
    
    if err != nil {
        response.Code = "404"
        response.Content = err              
    } else if err2 != nil {
        response.Code = "404"
        response.Content = err2
    } else {
        response.Code = "200"
        response.Content = purchaseDetails
    }
    w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])
}
func GetPurchaseDetail(w http.ResponseWriter, r *http.Request) {}
func CreatePurchaseDetail(w http.ResponseWriter, r *http.Request) {}
func UpdatePurchaseDetail(w http.ResponseWriter, r *http.Request) {}
func DeletePurchaseDetail(w http.ResponseWriter, r *http.Request) {}
