package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/trotsdeveloper/web_project_backend/dao"
	"github.com/trotsdeveloper/web_project_backend/rest"
)

func main() {
	var err error
	dao.DBConf, err = dao.InitDB()
	if err != nil {
		fmt.Println("Error initing DB")
	} else {
		r := chi.NewRouter()
		// registered_users
		r.Get("/registered_users", rest.GetAllUsers)
		r.Get("/registered_users/{id}", rest.GetUser)
		r.Post("/registered_users/login", rest.LoginUser)
		r.Post("/registered_users/{id}", rest.CreateUser)
		r.Put("/registered_users/{id}", rest.UpdateUser)
		r.Delete("/registered_users/{id}", rest.DeleteUser)
        // payment_method_types		
		r.Get("/payment_method_types", rest.GetAllPaymentMethodTypes)
		r.Get("/payment_method_types/{id}", rest.GetPaymentMethodType)
		// payment_methods
        r.Get("/payment_methods", rest.GetAllPaymentMethods)
        r.Get("/payment_methods/{id}", rest.GetPaymentMethod)
        r.Post("/payment_methods/{id}", rest.CreatePaymentMethod)
        r.Put("/payment_methods/{id}", rest.UpdatePaymentMethod)
        r.Delete("/payment_methods/{id}", rest.DeletePaymentMethod)        
        // purchases
        r.Get("/purchases", rest.GetAllPurchases)
        r.Get("/purchases/{id}", rest.GetPurchase)
        r.Post("/purchases/{id}", rest.CreatePurchase)
        r.Put("/purchases/{id}", rest.UpdatePurchase)
        r.Delete("/purchases/{id}", rest.DeletePurchase)
        // product_categories
        r.Get("/product_categories", rest.GetAllProductCategories)
        r.Get("/product_categories/{id}", rest.GetProductCategory)
        // product_offers
        r.Get("/product_offers", rest.GetAllProductOffers)
        r.Get("/product_offers/{id}", rest.GetProductOffer)
        r.Post("/product_offers/{id}", rest.CreateProductOffer)
        r.Put("/product_offers/{id}", rest.UpdateProductOffer)
        r.Delete("/product_offers/{id}", rest.DeleteProductOffer)
        // product_offer_img_urls
        r.Get("/product_offer_img_urls", rest.GetAllProductOfferImgUrls)
        r.Post("/product_offer_img_urls/{id}", rest.CreateProductOfferImgUrl)
        r.Put("/product_offer_img_urls/{id}", rest.UpdateProductOfferImgUrl)
        r.Delete("/product_offer_img_urls/{id}", rest.DeleteProductOfferImgUrl)
        // purchase_details
        r.Get("/purchase_details", rest.GetAllPurchaseDetails)
        r.Get("/purchase_details/{id}", rest.GetPurchaseDetail)
        r.Post("/purchase_details/{id}", rest.CreatePurchaseDetail)
        r.Put("/purchase_details/{id}", rest.UpdatePurchaseDetail)
        r.Delete("/purchase_details/{id}", rest.DeletePurchaseDetail)
        http.ListenAndServe(":3000", r)
    }
}