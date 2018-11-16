// Common tools and helper functions
package common

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/go-playground/validator.v8"

	"github.com/gin-gonic/gin/binding"
	"github.com/gin-gonic/gin"
	//"github.com/appleboy/gorush/rpc/proto"
	//"context"
	"encoding/json"
	"log"
	"math"
)

const (
	TEMPLATE_DB_CONSTRING = `%s:%s@tcp(%s:%s)/%s`
	FacebookProvider = "FACEBOOK"
	NormalProvider = "NORMAL"
)

const (
	ORDER_CREATED_STATUS = 1
	ORDER_ACCEPTED_BY_STORE = 2
	ORDER_ACCEPTED_BY_DELIVERY = 3
	ORDER_CONFIRM = 4
	ORDER_IN_STORE = 5
	ORDER_LAUNDRYING = 6
	ORDER_FINISH_LAUNDRYING = 7
	ORDER_DELIVERY_BACK_TO_CUSTOMER = 8
	ORDER_COMPLETE = 9
	ORDER_CANCEL = 10
	DELIVERY_CANNOT_RECEIVE_CLOTHES = 11
	DELIVERY_CANNOT_GIVE_BACK_CLOTHES = 12
	DELIVERY_REFUSE_TO_DELIVER = 13
)

const (
	HCM_LOWER_LATITUDE  =  10.6925
	HCM_UPPER_LATITUDE  =  10.8610
	HCM_LOWER_LONGITUDE = 106.5583
	HCM_UPPER_LONGITUDE = 106.7753
)
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// A helper function to generate random string
func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Keep this two config private, it should not expose to open source
const NBSecretPassword = "A String Very Very Very Strong!!@##$!@#$"
const NBRandomPassword = "A String Very Very Very Niubilty!!@##$!@#4"

const Page = "page";
const Limit = "limit";
const PageDefault = "1";
const LimitDefault = "10";

// A Util function to generate jwt_token which can be used in the request header
func GenToken(id uint) string {
	jwt_token := jwt.New(jwt.GetSigningMethod("HS256"))
	// Set some claims
	jwt_token.Claims = jwt.MapClaims{
		"id":  id,
		//"role_id": roleId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	// Sign and get the complete encoded token as a string
	token, _ := jwt_token.SignedString([]byte(NBSecretPassword))
	return token
}

// My own Error type that will help return my customized Error info
//  {"database": {"hello":"no such table", error: "not_exists"}}
type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

// To handle the error returned by c.Bind in gin framework
// https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
func NewValidatorError(err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		// can translate each error one at a time.
		//fmt.Println("gg",v.NameNamespace)
		if v.Param != "" {
			res.Errors[v.Field] = fmt.Sprintf("{%v: %v}", v.Tag, v.Param)
		} else {
			res.Errors[v.Field] = fmt.Sprintf("{key: %v}", v.Tag)
		}

	}
	return res
}

// Warp the error info in a object
func NewError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}

// Changed the c.MustBindWith() ->  c.ShouldBindWith().
// I don't want to auto return 400 when error happened.
// origin function is here: https://github.com/gin-gonic/gin/blob/master/context.go
func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}

func ProduceMessage(queue string,message interface{}){
	messageBytes, err := json.Marshal(message)
	if err != nil {
		// handle error
		log.Print(err)
	}

	switch(queue){
		case FIREBASE_QUEUE:
			GetFirebaseMQ().PublishBytes(messageBytes)
		case NOTIFICATION_QUEUE:
			GetNotificationMQ().PublishBytes(messageBytes)
	}
}


// haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// Distance function returns the distance (in meters) between two points of
//     a given longitude and latitude relatively accurately (using a spherical
//     approximation of the Earth) through the Haversin Distance Formula for
//     great arc distance on a sphere with accuracy for small distances
//
// point coordinates are supplied in degrees and converted into rad. in the func
//
// distance returned is METERS!!!!!!
// http://en.wikipedia.org/wiki/Haversine_formula
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

//Check if location is in ho chi minh
func IsLocationaInHoChiMinhCity(lat float32,long float32) bool{
	if lat >= HCM_LOWER_LATITUDE && lat <= HCM_UPPER_LATITUDE{
		if long >= HCM_LOWER_LONGITUDE && long <= HCM_UPPER_LONGITUDE{
			return true
		}
	}
	return false;
}