package course_maps

import (
	"fmt"
)

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	websites := map[string]string{
		"google": "www.google.com",
		"Amazon": "www.aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Amazon"])

	websites["Linkedin"] = "www.linkedin.com"
	fmt.Println(websites["Linkedin"])

	delete(websites, "Amazon")
	fmt.Println(websites)

}
