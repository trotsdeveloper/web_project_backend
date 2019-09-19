package rest

import (
    "encoding/json"
    "net/http"
    "github.com/go-chi/chi"
    "github.com/trotsdeveloper/web_project_backend/dao"
    "strconv"
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
func GetUser(w http.ResponseWriter, r *http.Request) {
    userId ,_ := strconv.Atoi(chi.URLParam(r, "id"))
    user := dao.RegisteredUser{Id: userId}
    var err error
    err = user.SelectInDB()
    response := Response{}
    respB, err2 := json.Marshal(user)
    
    if err != nil {
        response.Code = "404"
        response.Content = err              
    } else if err2 != nil {
        response.Code = "404"
        response.Content = err2
    } else {
        response.Code = "200"
        response.Content = user
    }
    
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])      
}
func LoginUser(w http.ResponseWriter, r *http.Request) {}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var err error
    var user dao.RegisteredUser
    json.NewDecoder(r.Body).Decode(&user)

    err = user.CreateInDB()
    response := Response{Code: "200"}
    respB, _ := json.Marshal(response)
    if err != nil {
        response.Code = "404"      
    } else {
        response.Code = "200"
    }
    
    w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])    
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
    var err error
    var user dao.RegisteredUser
    json.NewDecoder(r.Body).Decode(&user)
    userId ,_ := strconv.Atoi(chi.URLParam(r, "id"))
    user.Id = userId    
    err = user.UpdateInDB()
    response := Response{Code: "200"}
    respB, _ := json.Marshal(response)
    if err != nil {
        response.Code = "404"       
    } else {
        response.Code = "200"
    }
    
    w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])     
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    var err error
    id ,_ := strconv.Atoi(chi.URLParam(r, "id"))
    user := dao.RegisteredUser{Id: id}
    err = user.DeleteInDB()
    response := Response{Code: "200"}
    respB, _ := json.Marshal(response)
    
    if err != nil {
        response.Code = "404"       
    } else {
        response.Code = "200"
    }
    
    w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])
}

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
func GetPaymentMethodType(w http.ResponseWriter, r *http.Request) {
    paymentMethodTypeId ,_ := strconv.Atoi(chi.URLParam(r, "id"))
    paymentMethodType := dao.PaymentMethodType{Id: paymentMethodTypeId}
    var err error
    err = paymentMethodType.SelectInDB()
    response := Response{}
    respB, err2 := json.Marshal(paymentMethodType)
    
    if err != nil {
        response.Code = "404"
        response.Content = err              
    } else if err2 != nil {
        response.Code = "404"
        response.Content = err2
    } else {
        response.Code = "200"
        response.Content = paymentMethodType
    }
    
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])
    
}

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
func GetPaymentMethod(w http.ResponseWriter, r *http.Request) {
    paymentMethodId ,_ := strconv.Atoi(chi.URLParam(r, "id"))
    paymentMethod := dao.PaymentMethod{Id: paymentMethodId}
    var err error
    err = paymentMethod.SelectInDB()
    response := Response{}
    respB, err2 := json.Marshal(paymentMethod)
    
    if err != nil {
        response.Code = "404"
        response.Content = err              
    } else if err2 != nil {
        response.Code = "404"
        response.Content = err2
    } else {
        response.Code = "200"
        response.Content = paymentMethod
    }
    
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])
}
func CreatePaymentMethod(w http.ResponseWriter, r *http.Request) {
    var err error
    var paymentMethod dao.PaymentMethod
    json.NewDecoder(r.Body).Decode(&paymentMethod)

    err = paymentMethod.CreateInDB()
    response := Response{Code: "200"}
    respB, _ := json.Marshal(response)
    
    if err != nil {
        response.Code = "404"       
    } else {
        response.Code = "200"
    }
    
    w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:]) 
}
func UpdatePaymentMethod(w http.ResponseWriter, r *http.Request) {
    var err error
    var paymentMethod dao.PaymentMethod
    json.NewDecoder(r.Body).Decode(&paymentMethod)
    paymentMethodId ,_ := strconv.Atoi(chi.URLParam(r, "id"))
    paymentMethod.Id = paymentMethodId
    err = paymentMethod.UpdateInDB()
    response := Response{Code: "200"}
    respB, _ := json.Marshal(response)
    
    if err != nil {
        response.Code = "404"       
    } else {
        response.Code = "200"
    }
    
    w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:]) 
}
func DeletePaymentMethod(w http.ResponseWriter, r *http.Request) {
    var err error
    id ,_ := strconv.Atoi(chi.URLParam(r, "id"))
    paymentMethod := dao.PaymentMethod{Id: id}
    err = paymentMethod.DeleteInDB()
    response := Response{Code: "200"}
    respB, _ := json.Marshal(response)
    
    if err != nil {
        response.Code = "404"       
    } else {
        response.Code = "200"
    }
    
    w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])
}

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
func GetPurchase(w http.ResponseWriter, r *http.Request) {
    purchaseId ,_ := strconv.Atoi(chi.URLParam(r, "id"))
    purchase := dao.Purchase{Id: purchaseId}
    var err error
    err = purchase.SelectInDB()
    response := Response{}
    respB, err2 := json.Marshal(purchase)
    
    if err != nil {
        response.Code = "404"
        response.Content = err              
    } else if err2 != nil {
        response.Code = "404"
        response.Content = err2
    } else {
        response.Code = "200"
        response.Content = purchase
    }
    
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])    
    
}
func CreatePurchase(w http.ResponseWriter, r *http.Request) {
    var err error
    var purchase dao.Purchase
    json.NewDecoder(r.Body).Decode(&purchase)

    err = purchase.CreateInDB()
    response := Response{Code: "200"}
    respB, _ := json.Marshal(response)
    
    if err != nil {
        response.Code = "404"       
    } else {
        response.Code = "200"
    }
    
    w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])     
}
func UpdatePurchase(w http.ResponseWriter, r *http.Request) {
    var err error
    var purchase dao.Purchase
    json.NewDecoder(r.Body).Decode(&purchase)
    purchaseId ,_ := strconv.Atoi(chi.URLParam(r, "id"))
    purchase.Id = purchaseId

    err = purchase.UpdateInDB()
    response := Response{Code: "200"}
    respB, _ := json.Marshal(response)
    
    if err != nil {
        response.Code = "404"       
    } else {
        response.Code = "200"
    }
    
    w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:]) 
}
func DeletePurchase(w http.ResponseWriter, r *http.Request) {
    var err error
    id ,_ := strconv.Atoi(chi.URLParam(r, "id"))
    purchase := dao.Purchase{Id: id}
    err = purchase.DeleteInDB()
    response := Response{Code: "200"}
    respB, _ := json.Marshal(response)
    
    if err != nil {
        response.Code = "404"       
    } else {
        response.Code = "200"
    }
    
    w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])
}

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
func GetProductCategory(w http.ResponseWriter, r *http.Request) {
    productCategoryId ,_ := strconv.Atoi(chi.URLParam(r, "id"))
    productCategory := dao.ProductCategory{Id: productCategoryId}
    var err error
    err = productCategory.SelectInDB()
    response := Response{}
    respB, err2 := json.Marshal(productCategory)
    
    if err != nil {
        response.Code = "404"
        response.Content = err              
    } else if err2 != nil {
        response.Code = "404"
        response.Content = err2
    } else {
        response.Code = "200"
        response.Content = productCategory
    }
    
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:]) 
}

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
func GetProductOffer(w http.ResponseWriter, r *http.Request) {
    productOfferId ,_ := strconv.Atoi(chi.URLParam(r, "id"))
    productOffer := dao.ProductOffer{Id: productOfferId}
    var err error
    err = productOffer.SelectInDB()
    response := Response{}
    respB, err2 := json.Marshal(productOffer)
    
    if err != nil {
        response.Code = "404"
        response.Content = err              
    } else if err2 != nil {
        response.Code = "404"
        response.Content = err2
    } else {
        response.Code = "200"
        response.Content = productOffer
    }
    
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])    
    
}
func CreateProductOffer(w http.ResponseWriter, r *http.Request) {
    var err error
    var productOffer dao.ProductOffer
    json.NewDecoder(r.Body).Decode(&productOffer)

    err = productOffer.CreateInDB()
    response := Response{Code: "200"}
    respB, _ := json.Marshal(response)
    
    if err != nil {
        response.Code = "404"       
    } else {
        response.Code = "200"
    }
    
    w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])  
    
}
func UpdateProductOffer(w http.ResponseWriter, r *http.Request) {
    var err error
    var productOffer dao.ProductOffer
    json.NewDecoder(r.Body).Decode(&productOffer)

    productOfferId ,_ := strconv.Atoi(chi.URLParam(r, "id"))
    productOffer.Id = productOfferId
    
    err = productOffer.UpdateInDB()
    response := Response{Code: "200"}
    respB, _ := json.Marshal(response)
    
    if err != nil {
        response.Code = "404"       
    } else {
        response.Code = "200"
    }
    
    w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:]) 
}
func DeleteProductOffer(w http.ResponseWriter, r *http.Request) {
    var err error
    id ,_ := strconv.Atoi(chi.URLParam(r, "id"))
    productOffer := dao.ProductOffer{Id: id}
    err = productOffer.DeleteInDB()
    response := Response{Code: "200"}
    respB, _ := json.Marshal(response)
    
    if err != nil {
        response.Code = "404"       
    } else {
        response.Code = "200"
    }
    
    w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])
}

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
func CreateProductOfferImgUrl(w http.ResponseWriter, r *http.Request) {
    var err error
    var productOfferImgUrl dao.ProductOfferImgUrl
    json.NewDecoder(r.Body).Decode(&productOfferImgUrl)

    err = productOfferImgUrl.CreateInDB()
    response := Response{Code: "200"}
    respB, _ := json.Marshal(response)
    
    if err != nil {
        response.Code = "404"       
    } else {
        response.Code = "200"
    }
    
    w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])  
}
func UpdateProductOfferImgUrl(w http.ResponseWriter, r *http.Request) {
    var err error
    var productOfferImgUrl dao.ProductOfferImgUrl
    json.NewDecoder(r.Body).Decode(&productOfferImgUrl)
    
    productOfferImgUrlId ,_ := strconv.Atoi(chi.URLParam(r, "id"))
    productOfferImgUrl.Id = productOfferImgUrlId
    err = productOfferImgUrl.UpdateInDB()
    response := Response{Code: "200"}
    respB, _ := json.Marshal(response)
    
    if err != nil {
        response.Code = "404"       
    } else {
        response.Code = "200"
    }
    
    w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])
    
}
func DeleteProductOfferImgUrl(w http.ResponseWriter, r *http.Request) {
    var err error
    id ,_ := strconv.Atoi(chi.URLParam(r, "id"))
    productOfferImgUrl := dao.ProductOfferImgUrl{Id: id}
    err = productOfferImgUrl.DeleteInDB()
    response := Response{Code: "200"}
    respB, _ := json.Marshal(response)
    
    if err != nil {
        response.Code = "404"       
    } else {
        response.Code = "200"
    }
    
    w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])    
}

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

func CreatePurchaseDetail(w http.ResponseWriter, r *http.Request) {
    var err error
    var purchaseDetail dao.PurchaseDetail
    json.NewDecoder(r.Body).Decode(&purchaseDetail)

    err = purchaseDetail.CreateInDB()
    response := Response{Code: "200"}
    respB, _ := json.Marshal(response)
    
    if err != nil {
        response.Code = "404"       
    } else {
        response.Code = "200"
    }
    
    w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(respB[:])     
}
func UpdatePurchaseDetail(w http.ResponseWriter, r *http.Request) {

}
func DeletePurchaseDetail(w http.ResponseWriter, r *http.Request) {
  
}
