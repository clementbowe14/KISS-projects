package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var words = []string{
	"apple", "apricot", "avocado", "banana", "blackberry", "blueberry", "cantaloupe", "cherry", "clementine", "coconut",
	"cranberry", "date", "dragonfruit", "durian", "elderberry", "fig", "grape", "grapefruit", "guava", "honeydew",
	"jackfruit", "kiwi", "lemon", "lime", "longan", "lychee", "mango", "mangosteen", "melon", "mulberry",
	"nectarine", "orange", "papaya", "passionfruit", "peach", "pear", "persimmon", "pineapple", "plantain", "plum",
	"pomegranate", "pomelo", "quince", "raspberry", "rhubarb", "starfruit", "strawberry", "tangerine", "watermelon", "yuzu",
	"artichoke", "asparagus", "bamboo", "bean", "beet", "broccoli", "cabbage", "carrot", "cassava", "cauliflower",
	"celery", "chard", "chicory", "chive", "cilantro", "collard", "corn", "cucumber", "daikon", "edamame",
	"eggplant", "endive", "fennel", "garlic", "ginger", "kale", "kohlrabi", "leek", "lentil", "lettuce",
	"mushroom", "mustard", "okra", "onion", "parsnip", "pea", "pepper", "potato", "pumpkin", "radish",
	"rutabaga", "shallot", "spinach", "squash", "turnip", "yam", "zucchini", "almond", "cashew", "chestnut",
	"hazelnut", "macadamia", "peanut", "pecan", "pignoli", "pistachio", "walnut", "barley", "buckwheat", "bulgur",
	"chia", "flax", "millet", "oat", "quinoa", "rice", "rye", "sorghum", "spelt", "wheat",
	"beef", "bacon", "chicken", "duck", "ham", "lamb", "pork", "turkey", "veal", "venison",
	"anchovy", "bass", "catfish", "cod", "crab", "flounder", "haddock", "halibut", "herring", "lobster",
	"mackerel", "mahi", "mussel", "octopus", "oyster", "salmon", "sardine", "scallop", "shrimp", "snapper",
	"sole", "squid", "tilapia", "trout", "tuna", "walleye", "butter", "cheese", "cream", "milk",
	"yogurt", "bagel", "biscuit", "bread", "bun", "croissant", "muffin", "pancake", "pastry", "waffle",
	"alligator", "alpaca", "ant", "anteater", "antelope", "ape", "armadillo", "baboon", "badger", "barracuda",
	"bat", "bear", "beaver", "bee", "bison", "boar", "buffalo", "butterfly", "camel", "capybara",
	"caribou", "cat", "caterpillar", "cattle", "chamois", "cheetah", "chicken", "chimpanzee", "chinchilla", "clam",
	"cobra", "cockroach", "cod", "coyote", "crab", "crane", "crocodile", "crow", "deer", "dinosaur",
	"dog", "dolphin", "donkey", "dove", "dragonfly", "duck", "dugong", "eagle", "echidna", "eel",
	"eland", "elephant", "elk", "emu", "falcon", "ferret", "finch", "fish", "flamingo", "fly",
	"fox", "frog", "gaur", "gazelle", "gerbil", "giraffe", "gnat", "gnu", "goat", "goldfish",
	"goose", "gorilla", "goshawk", "grasshopper", "grouse", "guanaco", "gull", "hamster", "hare", "hawk",
	"hedgehog", "heron", "herring", "hippopotamus", "hornet", "horse", "human", "hummingbird", "hyena", "ibex",
	"ibis", "jackal", "jaguar", "jay", "jellyfish", "kangaroo", "kinkajou", "koala", "kouprey", "kudu",
	"lapwing", "lark", "lemur", "leopard", "lion", "llama", "lobster", "locust", "loris", "louse",
	"lyrebird", "magpie", "mallard", "manatee", "mandrill", "mantis", "marten", "meerkat", "mink", "mole",
	"mongoose", "monkey", "moose", "mosquito", "mouse", "mule", "narwhal", "newt", "nightingale", "octopus",
	"okapi", "opossum", "oryx", "ostrich", "otter", "owl", "oyster", "panther", "parrot", "partridge",
	"peafowl", "pelican", "penguin", "pheasant", "pig", "pigeon", "pony", "porcupine", "porpoise", "quail",
	"quelea", "quetzal", "rabbit", "raccoon", "rail", "ram", "rat", "raven", "reindeer", "rhinoceros",
	"rook", "salamander", "salmon", "sardine", "scorpion", "seahorse", "seal", "shark", "sheep", "shrew",
	"skunk", "snail", "snake", "snipe", "sparrow", "spider", "spoonbill", "squid", "squirrel", "starling",
	"stingray", "stinkbug", "stork", "swallow", "swan", "tapir", "tarsier", "termite", "tiger", "toad",
	"trout", "turkey", "turtle", "viper", "vulture", "wallaby", "walrus", "wasp", "weasel", "whale",
	"wolverine", "wombat", "woodcock", "woodpecker", "worm", "wren", "yak", "zebra", "amber", "amethyst",
	"aquamarine", "azure", "beige", "black", "blue", "bronze", "brown", "burgundy", "charcoal", "chartreuse",
	"chestnut", "chocolate", "cobalt", "copper", "coral", "crimson", "cyan", "emerald", "fuchsia", "gold",
	"gray", "green", "indigo", "ivory", "jade", "khaki", "lavender", "lemon", "lilac", "magenta",
	"mahogany", "maroon", "mint", "mustard", "navy", "ochre", "olive", "orange", "orchid", "peach",
	"pear", "periwinkle", "pink", "plum", "puce", "purple", "raspberry", "red", "rose", "ruby",
	"rust", "salmon", "sapphire", "scarlet", "sepia", "silver", "tan", "tangerine", "taupe", "teal",
	"turquoise", "ultramarine", "vermilion", "violet", "white", "yellow", "actor", "author", "baker", "barber",
	"builder", "butcher", "carpenter", "chef", "clerk", "cook", "dentist", "doctor", "driver", "editor",
	"engineer", "farmer", "firefighter", "fisherman", "gardener", "grocer", "hunter", "inspector", "judge", "lawyer",
	"mechanic", "musician", "nurse", "painter", "pilot", "plumber", "police", "reporter", "sailor", "scientist",
	"secretary", "singer", "soldier", "tailor", "teacher", "trainer", "writer", "anchor", "anvil", "arrow",
	"axe", "badge", "balloon", "basket", "bed", "bell", "blanket", "book", "bottle", "bowl",
	"box", "broom", "bucket", "button", "candle", "carpet", "chain", "chair", "clock", "cloth",
	"coin", "comb", "compass", "cup", "curtain", "desk", "door", "drawer", "envelope", "fan",
	"feather", "flag", "flask", "fork", "frame", "funnel", "glove", "hammer", "hook", "ink",
	"jar", "key", "knife", "ladder", "lamp", "lantern", "lock", "magnet", "map", "mask",
	"mirror", "needle", "net", "newspaper", "padlock", "pan", "paper", "pen", "pencil", "pillow",
	"pin", "pipe", "plate", "pot", "purse", "ribbon", "ring", "rope", "sack", "saddle",
	"scale", "scissors", "screw", "shield", "shoe", "shovel", "soap", "sock", "spoon", "stamp",
	"sword", "table", "towel", "toy", "tray", "umbrella", "vase", "wallet", "watch", "wheel",
}

func writeWordsToFile(filename string, words []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, word := range words {
		_, err := writer.WriteString(word + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func main() {
	writeWordsToFile("words.txt", words)
	trie := NewTrie()
	file, err := os.Open("/Users/clementbowe/Documents/KISS-projects/go-projects/wordfinder/words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		word := scanner.Text()
		trie.Insert(word)
	}

	// Example usage of the trie
	fmt.Println(trie.Search("apple"))  // Output: true or false
	fmt.Println(trie.Search("banana")) // Output: true or false
}
