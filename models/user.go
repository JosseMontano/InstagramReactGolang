package models
import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type User struct{
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name" json:"name,omitempty"`
	LastName string `bson:"lastName" json:"lastName,omitempty"`
	DateBirth time.Time `bson:"dateBirth" json:"dateBirth,omitempty"`
	Gmail string `bson:"gmail" json:"gmail"`
	Password string `bson:"password" json:"password,omitempty"`
	Avatar string `bson:"avatar" json:"avatar,omitempty"`
	Banner string `bson:"banner" json:"banner,omitempty"`
	Bibliography string `bson:"bibliography" json:"bibliography,omitempty"`
	Ubication string `bson:"ubication" json:"ubication,omitempty"`
	SiteWeb string `bson:"siteWeb" json:"siteWeb,omitempty"`
}