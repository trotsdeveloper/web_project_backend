// Package for the declaration of the main structures
// for data handling in the project.
package dao

import (
	"database/sql"
	"fmt"
	"strconv"
	_ "github.com/lib/pq"
)

// Postgresql Settings
const (
	USER        = "admin"
	PASSWORD    = "admin"
	HOST        = "localhost"
	PORT        = 5432
	DATABASE    = "ecommerce"
)

// Declaration of global DB controller
var (
	DBConf *sql.DB
)

// Function for the inicialization of the global DB controller
func InitDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", USER, PASSWORD, HOST, strconv.Itoa(PORT), DATABASE)
	db, err := sql.Open("postgres", psqlInfo)
	return db, err
}

// Auxiliar function to add quotes to a string
func addQuotes(word string) string {
	return fmt.Sprintf(`'%v'`, word)
}

// registered_user
type RegisteredUser struct {
    Id          int         `json:"id"` // SERIAL PRIMARY KEY
    PhotoUrl    string      `json:"photo_url"`  // VARCHAR(100)
    Name        string      `json:"name"`       // VARCHAR(45)
    LastName    string      `json:"last_name"`  // VARCHAR(45)
    Email       string      `json:"e_mail"`      // VARCHAR(100)
    Country     string      `json:"country"`    // VARCHAR(45)
    Direction   string      `json:"direction"`  // VARCHAR(200)
    ZipCode     string      `json:"zip_code"`   // VARCHAR(30)
    Phone       string      `json:"phone"`      // VARCHAR(45)
    IsAdmin     bool        `json:"is_admin"`   // BOOLEAN
    Password    string      `json:"password"`   // VARCHAR(30) 
}

// RegisteredUser
// dbc *sql.DB

// payment_method_type
type PaymentMethodType struct {
    Id          int         `json:"id"`         // SERIAL PRIMARY KEY
    Name        string      `json:"name"`       // VARCHAR(45)
}
// payment_method
type PaymentMethod struct {
    Id          int         `json:"id"`         // SERIAL PRIMARY KEY
    AccountInfo string      `json:"account_info"`   // VARCHAR (100)
    Type        int         `json:"type"`       // INTEGER
    UserId      int         `json:"user_id"`    // INTEGER
}
// purchase
type Purchase struct {
    Id          int         `json:"id"`         // SERIAL PRIMARY KEY
    DirectionToSend string  `json:"direction_to_send"`  // VARCHAR(200)
    ZipCodeToSend   string  `json:"zip_code_to_send"`   // VARCHAR(30)
    Buyer       int         `json:"buyer"`      // INTEGER
    PaymentMethodId int     `json:"payment_method_id"`  // INTEGER
}
// product_category
type ProductCategory struct {
    Id          int         `json:"id"` // SERIAL PRIMARY KEY
    Name        string      `json:"name"`   // VARCHAR(45)
}
// product_offer
type ProductOffer struct {
    Id          int         `json:"id"` // SERIAL PRIMARY KEY
    Name        string      `json:"name"`   // VARCHAR(45)
    Description string      `json:"description"`  // VARCHAR(200)
    Price       float64     `json:"price"`  // REAL
    PrincipalImageUrl   string  `json:"principal_image_url"` // VARCHAR(100)    
}
// product_offer_img_urls
type ProductOfferImgUrl struct {
    Id          int         `json:"id"`         // SERIAL PRIMARY KEY
    ProductId   int         `json:"product_id"` // INTEGER
    ImgUrl      string      `json:"img_url"`    // VARCHAR(100)
}
// purchase_details
type PurchaseDetail struct {
    ProductOfferId  int     `json:"product_offer_id"`   // INTEGER
    PurchaseId      int     `json:"purchase_id"`        // INTEGER
    Quantity        int     `json:"quantity"`           // REAL
}

func (ru *RegisteredUser) SelectInDB() error {
    sqlStatement := `SELECT photo_url, name, last_name, e_mail, country,
    direction, zip_code, phone, is_admin, password FROM registered_user WHERE id=$1;`
    row := DBConf.QueryRow(sqlStatement, ru.Id)
    err := row.Scan(&ru.PhotoUrl, &ru.Name, &ru.LastName, &ru.Email, &ru.Country,
    &ru.Direction, &ru.ZipCode, &ru.Phone, &ru.IsAdmin, &ru.Password)
    return err
}

func (ru *RegisteredUser) CreateInDB() error {
    sqlStatement := `INSERT INTO registered_user (photo_url, name, last_name, e_mail,
    country, direction, zip_code, phone, is_admin, password) VALUES ($1, $2, $3, 
    $4, $5, $6, $7, $8, $9, $10) RETURNING id;`
    row := DBConf.QueryRow(sqlStatement, ru.PhotoUrl, ru.Name, ru.LastName,
    ru.Email, ru.Country, ru.Direction, ru.ZipCode, ru.Phone, ru.IsAdmin,
    ru.Password)
    err := row.Scan(&ru.Id)
    return err
}

func (ru *RegisteredUser) UpdateInDB() error {
    sqlStatement := `UPDATE registered_user SET photo_url = $2, name = $3, last_name = $4,
    e_mail = $5, country = $6, direction = $7, zip_code = $8, phone = $9, is_admin = $10,
    password = $11 WHERE id = $1;`
    _, err := DBConf.Exec(sqlStatement, ru.Id, ru.PhotoUrl, ru.Name, ru.LastName,
    ru.Email, ru.Country, ru.Direction, ru.ZipCode, ru.Phone, ru.IsAdmin, ru.Password)
    return err
}

func (ru *RegisteredUser) DeleteInDB() error {
    sqlStatement := `DELETE FROM registered_user WHERE id = $1;`
    _, err := DBConf.Exec(sqlStatement, ru.Id)
    return err
}

func (pmt *PaymentMethodType) SelectInDB() error {
    sqlStatement := `SELECT name FROM payment_method_type WHERE id=$1;`
    row := DBConf.QueryRow(sqlStatement, pmt.Id)
    err := row.Scan(&pmt.Name)
    return err
}

func (pmt *PaymentMethodType) CreateInDB() error {
    sqlStatement := `INSERT INTO payment_method_type (name) 
    VALUES ($1) RETURNING id;`
    row := DBConf.QueryRow(sqlStatement, pmt.Name)
    err := row.Scan(&pmt.Id)
    return err    
}

func (pmt *PaymentMethodType) UpdateInDB() error {
    sqlStatement := `UPDATE payment_method_type SET name = $2 WHERE id = $1;`
    _, err := DBConf.Exec(sqlStatement, pmt.Id, pmt.Name)
    return err
}

func (pmt *PaymentMethodType) DeleteInDB() error {
    sqlStatement := `DELETE FROM payment_method_type WHERE id = $1;`
    _, err := DBConf.Exec(sqlStatement, pmt.Id)
    return err
}

func (pm *PaymentMethod) SelectInDB() error {
    sqlStatement := `SELECT account_info, type, user_id FROM 
    payment_method WHERE id=$1;`
    row := DBConf.QueryRow(sqlStatement, pm.Id)
    err := row.Scan(&pm.AccountInfo, &pm.Type, &pm.UserId)
    return err
}

func (pm *PaymentMethod) CreateInDB() error {
    sqlStatement := `INSERT INTO payment_method (account_info, type, user_id)
    VALUES ($1, $2, $3) RETURNING id;`
    row := DBConf.QueryRow(sqlStatement, pm.AccountInfo, pm.Type, pm.UserId)
    err := row.Scan(&pm.Id)
    return err
}

func (pm *PaymentMethod) UpdateInDB() error {
    sqlStatement := `UPDATE payment_method SET account_info = $2, type = $3,
    user_id = $4 WHERE id = $1;`
    _, err := DBConf.Exec(sqlStatement, pm.Id, pm.AccountInfo, pm.Type, pm.UserId)
    return err
}

func (pm *PaymentMethod) DeleteInDB() error {
    sqlStatement := `DELETE FROM payment_method WHERE id = $1;`
    _, err := DBConf.Exec(sqlStatement, pm.Id)
    return err
}

func (p *Purchase) SelectInDB() error {
    sqlStatement := `SELECT direction_to_send, zip_code_to_send, buyer FROM
    purchase WHERE id=$1;`
    row := DBConf.QueryRow(sqlStatement, p.Id)
    err := row.Scan(&p.DirectionToSend, &p.ZipCodeToSend, &p.Buyer)
    return err
}
func (p *Purchase) CreateInDB() error {
    sqlStatement := `INSERT INTO purchase (direction_to_send, zip_code_to_send,
    buyer, payment_method_id) VALUES ($1, $2, $3, $4) RETURNING id;`
    row := DBConf.QueryRow(sqlStatement, p.DirectionToSend, p.ZipCodeToSend,
    p.Buyer, p.PaymentMethodId)
    err := row.Scan(&p.Id)
    return err
}

func (p *Purchase) UpdateInDB() error {
    sqlStatement := `UPDATE purchase SET direction_to_send = $2, zip_code_to_send = $3,
    buyer = $4, payment_method_id = $5 WHERE id = $1;`
    _, err := DBConf.Exec(sqlStatement, p.Id, p.DirectionToSend, p.ZipCodeToSend,
    p.Buyer, p.PaymentMethodId)
    return err
}

func (p *Purchase) DeleteInDB() error {
    sqlStatement := `DELETE FROM purchase WHERE id = $1;`
    _, err := DBConf.Exec(sqlStatement, p.Id)
    return err
}

func (pc *ProductCategory) SelectInDB() error {
    sqlStatement := `SELECT name FROM product_category WHERE id=$1;`
    row := DBConf.QueryRow(sqlStatement, pc.Id)
    err := row.Scan(&pc.Name)
    return err
}

func (pc *ProductCategory) CreateInDB() error {
    sqlStatement := `INSERT INTO product_category (name) VALUES ($1) RETURNING id;`
    row := DBConf.QueryRow(sqlStatement, pc.Name)
    err := row.Scan(&pc.Id)
    return err    
}

func (pc *ProductCategory) UpdateInDB() error {
    sqlStatement := `UPDATE product_category SET name = $2 WHERE id = $1;`
    _, err := DBConf.Exec(sqlStatement, pc.Id, pc.Name)
    return err
}

func (pc *ProductCategory) DeleteInDB() error {
    sqlStatement := `DELETE FROM product_category WHERE id = $1;`
    _, err := DBConf.Exec(sqlStatement, pc.Id)
    return err
}

func (po *ProductOffer) SelectInDB() error {
    sqlStatement := `SELECT name, description, price, principal_image_url FROM product_offer WHERE id=$1;`
    row := DBConf.QueryRow(sqlStatement, po.Id)
    err := row.Scan(&po.Name, &po.Description, &po.Price, &po.PrincipalImageUrl)
    return err
}

func (po *ProductOffer) CreateInDB() error {
    sqlStatement := `INSERT INTO product_offer (name, description, price, 
    principal_image_url) VALUES($1, $2, $3, $4) RETURNING id;`
    row := DBConf.QueryRow(sqlStatement, po.Name, po.Description, po.Price,
    po.PrincipalImageUrl)
    err := row.Scan(&po.Id)
    return err
}

func (po *ProductOffer) UpdateInDB() error {
    sqlStatement := `UPDATE product_offer SET name = $2, description = $3, price = $4,
    principal_image_url = $5 WHERE id = $1;`
    _, err := DBConf.Exec(sqlStatement, po.Name, po.Description, po.Price,
    po.PrincipalImageUrl)
    return err
}

func (po *ProductOffer) DeleteInDB() error {
    sqlStatement := `DELETE FROM product_offer WHERE id = $1;`
    _, err := DBConf.Exec(sqlStatement, po.Id)
    return err
}

func (poiu *ProductOfferImgUrl) SelectInDB() error {
    sqlStatement := `SELECT product_id, img_url FROM product_offer_img_urls WHERE
    id = $1;`
    row := DBConf.QueryRow(sqlStatement, poiu.Id)
    err := row.Scan(&poiu.ProductId, &poiu.ImgUrl)
    return err
}

func (poiu *ProductOfferImgUrl) CreateInDB() error {
    sqlStatement := `INSERT INTO product_offer_img_urls (product_id, img_url) VALUES
    ($1, $2) RETURNING id;`
    row := DBConf.QueryRow(sqlStatement, poiu.ProductId, poiu.ImgUrl)
    err := row.Scan(&poiu.Id)
    return err
}

func (poiu *ProductOfferImgUrl) UpdateInDB() error {
    sqlStatement := `UPDATE product_offer_img_urls (product_id, img_url) SET product_id = $2,
    img_url = $3 WHERE id = $1;`
    _, err := DBConf.Exec(sqlStatement, poiu.Id, poiu.ProductId, poiu.ImgUrl)
    return err
}

func (poiu *ProductOfferImgUrl) DeleteInDB() error {
    sqlStatement := `DELETE FROM product_offer_img_urls WHERE id = $1;`
    _, err := DBConf.Exec(sqlStatement, poiu.Id)
    return err
}

func (pd *PurchaseDetail) SelectInDB() error {
    sqlStatement := `SELECT quantity FROM purchase_details WHERE product_offer_id = $1 
    AND purchase_id = $2;`
    row := DBConf.QueryRow(sqlStatement, pd.Quantity)
    err := row.Scan(&pd.ProductOfferId, &pd.PurchaseId)
    return err   
}

func (pd *PurchaseDetail) CreateInDB() error {
    sqlStatement := `INSERT INTO purchase_details (product_offer_id, purchase_id, quantity)
    VALUES ($1, $2, $3);`
    DBConf.QueryRow(sqlStatement, pd.ProductOfferId, pd.PurchaseId, pd.Quantity)
    return nil
}

func (pd *PurchaseDetail) UpdateInDB() error {
    sqlStatement := `UPDATE purchase_details (quantity) SET quantity = $3,
    WHERE product_offer_id = $1 AND purchase_id = $2;`
    _, err := DBConf.Exec(sqlStatement, pd.ProductOfferId, pd.PurchaseId, pd.Quantity)
    return err
}

func (pd *PurchaseDetail) DeleteInDB() error {
    sqlStatement := `DELETE FROM purchase_details WHERE product_offer_id = $1 AND purchase_id = $2;`
    _, err := DBConf.Exec(sqlStatement, pd.ProductOfferId, pd.PurchaseId)
    return err
}

func (ru *RegisteredUser) SelectAllInDB() ([]RegisteredUser, error) {
    var users []RegisteredUser
    sqlStatement := `SELECT id, photo_url, name, last_name, e_mail, country,
    direction, zip_code, phone, is_admin, password FROM registered_user;`
    rows, err := DBConf.Query(sqlStatement)
    if err != nil {
        return users, err               
    }
    defer rows.Close()
    
    for rows.Next() {
        user := RegisteredUser{}
        err = rows.Scan(&user.Id, &user.PhotoUrl, &user.Name, &user.LastName,
        &user.Email, &user.Country, &user.Direction, &user.ZipCode, &user.Phone,
        &user.IsAdmin, &user.Password)
        if err != nil {
            return users, err
        }
        users = append(users, user)        
    }
    // get any error encountered during iteration
    err = rows.Err()
    if err != nil {
        return users, err
    }
    return users, err    
}

func (pmt *PaymentMethodType) SelectAllInDB() ([]PaymentMethodType, error) {
    var types []PaymentMethodType
    sqlStatement := `SELECT id, name FROM payment_method_type;`
    rows, err := DBConf.Query(sqlStatement)
    if err != nil {
        return types, err
    }
    defer rows.Close()
    
    for rows.Next() {
        paymentMethodType := PaymentMethodType{}
        err = rows.Scan(&paymentMethodType.Id, &paymentMethodType.Name)
        if err != nil {
            return types, err
        }
        types = append(types, paymentMethodType)
    }
    err = rows.Err()
    if err != nil {
        return types, err
    }
    
    return types, err    
}
func (pm *PaymentMethod) SelectAllInDB() ([]PaymentMethod, error) {
    var paymentMethods []PaymentMethod
    sqlStatement := `SELECT id, account_info, type, user_id FROM payment_method;`
    rows, err := DBConf.Query(sqlStatement)
    if err != nil {
        return paymentMethods, err
    }
    defer rows.Close()
    
    for rows.Next() {
        paymentMethod := PaymentMethod{}
        err = rows.Scan(&paymentMethod.Id, &paymentMethod.AccountInfo, 
        &paymentMethod.Type, &paymentMethod.UserId)
        if err != nil {
            return paymentMethods, err
        }
        paymentMethods = append(paymentMethods, paymentMethod)
    }
    err = rows.Err()
    if err != nil {
        return paymentMethods, err
    }
    return paymentMethods, err
}
func (p *Purchase) SelectAllInDB() ([]Purchase, error) {
    var purchases []Purchase
    sqlStatement := `SELECT id, direction_to_send, zip_code_to_send, 
    buyer, payment_method_id FROM purchase;`
    rows, err := DBConf.Query(sqlStatement)
    if err != nil {
        return purchases, err
    }
    defer rows.Close()
    
    for rows.Next() {
        purchase := Purchase{}
        err = rows.Scan(&purchase.Id, &purchase.DirectionToSend, 
        &purchase.ZipCodeToSend, &purchase.Buyer, &purchase.PaymentMethodId)
        if err != nil {
            return purchases, err
        }
        purchases = append(purchases, purchase)
    }
    err = rows.Err()
    if err != nil {
        return purchases, err
    }
    return purchases, err
}
func (pc *ProductCategory) SelectAllInDB() ([]ProductCategory, error) {
    var productCategories []ProductCategory
    sqlStatement := `SELECT id, name FROM product_category;`
    rows, err := DBConf.Query(sqlStatement)
    if err != nil {
        return productCategories, err
    }
    defer rows.Close()
    
    for rows.Next() {
        productCategory := ProductCategory{}
        err = rows.Scan(&productCategory.Id, &productCategory.Name)
        if err != nil {
            return productCategories, err
        }
        productCategories = append(productCategories, productCategory)
    }
    err = rows.Err()
    if err != nil {
        return productCategories, err
    }
    return productCategories, err
}
func (po *ProductOffer) SelectAllInDB() ([]ProductOffer, error) {
    var productOffers []ProductOffer
    sqlStatement := `SELECT id, name, description, price, 
    principal_image_url FROM product_offer;`
    rows, err := DBConf.Query(sqlStatement)
    if err != nil {
        return productOffers, err
    }
    defer rows.Close()
    
    for rows.Next() {
        productOffer := ProductOffer{}
        err = rows.Scan(&productOffer.Id, &productOffer.Name, &productOffer.Description, 
        &productOffer.Price, &productOffer.PrincipalImageUrl)
        if err != nil {
            return productOffers, err
        }
        productOffers = append(productOffers, productOffer)
    }
    err = rows.Err()
    if err != nil {
        return productOffers, err
    }
    return productOffers, err
}
func (poiu *ProductOfferImgUrl) SelectAllInDB() ([]ProductOfferImgUrl, error) {
    var imgUrls []ProductOfferImgUrl
    sqlStatement := `SELECT id, product_id, img_url FROM product_offer_img_urls;`
    rows, err := DBConf.Query(sqlStatement)
    if err != nil {
        return imgUrls, err
    }
    defer rows.Close()
    
    for rows.Next() {
        imgUrl := ProductOfferImgUrl{}
        err = rows.Scan(&imgUrl.Id, &imgUrl.ProductId, &imgUrl.ImgUrl)
        if err != nil {
            return imgUrls, err
        }
        imgUrls = append(imgUrls, imgUrl)
    }
    err = rows.Err()
    if err != nil {
        return imgUrls, err
    }
    return imgUrls, err
}
func (pd *PurchaseDetail) SelectAllInDB() ([]PurchaseDetail, error) {
    var purchaseDetails []PurchaseDetail
    sqlStatement := `SELECT product_offer_id, purchase_id, quantity FROM purchase_details;`
    rows, err := DBConf.Query(sqlStatement)
    if err != nil {
        return purchaseDetails, err
    }
    defer rows.Close()
    
    for rows.Next() {
        purchaseDetail := PurchaseDetail{}
        err = rows.Scan(&purchaseDetail.ProductOfferId, &purchaseDetail.PurchaseId, &purchaseDetail.Quantity)
        if err != nil {
            return purchaseDetails, err
        }
        purchaseDetails = append(purchaseDetails, purchaseDetail)
    }
    err = rows.Err()
    if err != nil {
        return purchaseDetails, err
    }
    return purchaseDetails, err 
}
