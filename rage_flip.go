package main

// Exclude these two imports for a DigitalOcean function
import (
	"fmt"
	"os"
)

// Referral code https://m.do.co/c/db1d1305f10a

func flip(word string) string {
	forward := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890-=!@#$%^&*()_+ "
	backsidedown := " +‾()*⅋^%$#@¡=-068𝘓95ߤ↋↊⇂zʎxʍʌnʇsɹbdouɯʅʞɾᴉɥƃⅎǝpɔqɐZ⅄XϺɅՈꓕSꓤꝹԀONꟽ⅂ꓘᒋIH⅁ᖵƎᗡϽꓭ∀"
	upsidedown := ""
	upsidedownmap := make(map[string]string)
	flipped := ""

	for _, r := range backsidedown {
		upsidedown = string(r) + upsidedown
	}

	i := 0
	for _, r := range upsidedown {
		upsidedownmap[string(forward[i])] = string(r)
		i++
	}

	for _, r := range word {
		flipped = upsidedownmap[string(r)] + flipped
	}
	return flipped
}

func Main(args map[string]interface{}) map[string]interface{} {
	rageflip_front := "(ノಠ益ಠ)ノ彡┻"

	rageflip_back := "┻"
	name, ok := args["name"].(string)
	if !ok {
		name = "stranger"
	}
	msg := make(map[string]interface{})
	msg["body"] = rageflip_front + flip(name) + rageflip_back
	return msg
}

// Exclude func main() for the Digital Ocean function
func main() {

	args := make(map[string]interface{})
	args["name"] = os.Args[1]
	fmt.Println(Main(args)["body"].(string))
}
