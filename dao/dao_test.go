package dao

import (
    "fmt"
	"testing"
)

func TestAddQuotes(t *testing.T) {
    var err error
    DBConf, err = InitDB()
    if err != nil {
        fmt.Println(err)
    }
    
    user := RegisteredUser{Name: "Manuel", LastName: "Mena"}
    err = user.CreateInDB()
    if err != nil {
        fmt.Println(err)
    }
    
    if user.Id == 0 {
        t.Errorf("Failed to get User Id.")
    }
    
    paymentMethodType := PaymentMethodType{}
    err = paymentMethodType.CreateInDB()
    if err != nil {
        fmt.Println("Exception in payment method type.")
        fmt.Println(err)
    }
    
    if paymentMethodType.Id == 0 {
        t.Errorf("Failed to get Payment Method Type Id.")
    }
    
    paymentMethod := PaymentMethod{}
    paymentMethod.Type = paymentMethodType.Id
    paymentMethod.UserId = user.Id
    err = paymentMethod.CreateInDB()
    if err != nil {
        fmt.Println("Exception in payment method.")
        fmt.Println(err)
    }
    
    if paymentMethod.Id == 0 {
        t.Errorf("Failed to get Payment Method Id.")
    }       
    
    purchase := Purchase{}
    purchase.Buyer = user.Id
    purchase.PaymentMethodId = paymentMethod.Id
    err = purchase.CreateInDB()
    if err != nil {
        fmt.Println(err)
    }
    
    if purchase.Id == 0 {
        t.Errorf("Failed to get Purchase Id.")
    }    
    
    productCategory := ProductCategory{}
    err = productCategory.CreateInDB()
    if err != nil {
        fmt.Println(err)
    }
    
    if productCategory.Id == 0 {
        t.Errorf("Failed to get Product Category Id.")
    }
    
    productOffer := ProductOffer{}
    err = productOffer.CreateInDB()
    if err != nil {
        fmt.Println(err)
    }
    
    if productOffer.Id == 0 {
        t.Errorf("Failed to get Product Offer Id.")
    }
    
    productOfferImgUrl := ProductOfferImgUrl{}    
    productOfferImgUrl.ProductId = productOffer.Id
    err = productOfferImgUrl.CreateInDB()
    if err != nil {
        fmt.Println(err)
    }
    
    if productOfferImgUrl.Id == 0 {
        t.Errorf("Failed to get Product Offer Img Url Id.")
    }
    
    purchaseDetail := PurchaseDetail{}
    purchaseDetail.ProductOfferId = productOffer.Id
    purchaseDetail.PurchaseId = purchase.Id
    err = purchaseDetail.CreateInDB()
    if err != nil {
        fmt.Println(err)
    }
    users, err := user.SelectAllInDB()
    if err != nil {
        fmt.Println("Error listing users.")
        fmt.Println(err)
    }
    if len(users) == 0 {    
        t.Errorf("Failed to get users.")
    }
    paymentMethodTypes, err := paymentMethodType.SelectAllInDB()
    if err != nil {
        fmt.Println("Error listing payment method types.")
        fmt.Println(err)
    }
    if len(paymentMethodTypes) == 0 {
        t.Errorf("Failed to get payment method types.")
    }
    
    paymentMethods, err := paymentMethod.SelectAllInDB()
    if err != nil {
        fmt.Println("Error listing payment methods.")
        fmt.Println(err)
    }
    
    if len(paymentMethods) == 0 {
        t.Errorf("Failed to get payment methods.")
    }
    
    purchases, err := purchase.SelectAllInDB()
    if err != nil {
        fmt.Println("Error listing purchases.")
        fmt.Println(err)
    }
    
    if len(purchases) == 0 {
        t.Errorf("Failed to get purchases")
    }
    
    productCategories, err := productCategory.SelectAllInDB()
    if err != nil {
        fmt.Println("Error listing product categories.")
        fmt.Println(err)  
    }
    
    if len(productCategories) == 0 {
        t.Errorf("Failed to get product categories.")
    }
    
    productOffers, err := productOffer.SelectAllInDB()
    if err != nil {
        fmt.Println("Error listing product offers.")
        fmt.Println(err)
    }
    
    if len(productOffers) == 0 {
        t.Errorf("Failed to get product offers.")
    }
    
    imgUrls, err := productOfferImgUrl.SelectAllInDB()
    if err != nil {
        fmt.Println("Error listing image urls.")
        fmt.Println(err)
    }
    
    if len(imgUrls) == 0 {
        t.Errorf("Failed to get product offer image urls.")
    }
    
    purchaseDetails, err := purchaseDetail.SelectAllInDB()
    if err != nil {
        fmt.Println("Error listing purchase details.")
        fmt.Println(err)
    }
    if len(purchaseDetails) == 0 {
        t.Errorf("Failed to get purchase details.")
    }   
    
}