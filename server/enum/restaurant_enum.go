package enum

type RestaurantEnum int

const (
	// CheckRestaurant -
	CheckRestaurant = iota
)

func (action RestaurantEnum) String() string {
	return [...]string{
		"查看餐廳",
	}[action]
}
