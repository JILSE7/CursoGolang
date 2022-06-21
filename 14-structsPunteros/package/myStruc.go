package myPhone

import "fmt"

// Phone can access public
type PhonePublic struct {
	Ram    int
	Memory int
	Brand  string
	Price  int
}

// IncrementMemoryPublic
func (phone *PhonePublic) IncrementMemoryPublic() {
	phone.Ram = phone.Ram * 2
	phone.Memory = phone.Memory * 2
}

// PrintFeaturesPublic
func (phone PhonePublic) PrintFeaturesPublic() {
	fmt.Println("ram", phone.Ram, "memory", phone.Memory, "price", phone.Price, "Brand", phone.Brand)
}
