package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Girl struct {
	hot   bool
	match bool
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup            // Create waitGroup to determine when all goRoutines are finished.
	rand.Seed(time.Now().UnixNano()) // To be sure we have a random result every try.

	//Init Mockdata
	theGirls := make([]Girl, 0)
	theGirls = InitGirls(3000)

	//Channels
	likedChannel := make(chan Girl, 3) // Buffered channel
	matchedChannel := make(chan Girl)

	//Start GoRoutines
	wg.Add(3)
	go SwipeAndLike(theGirls, likedChannel, &wg)
	go FindMatches(likedChannel, matchedChannel, &wg)
	go StartConversations(matchedChannel, &wg)
	wg.Wait()

	elapsed := time.Since(start)
	log.Printf("Finished in %s", elapsed)
}

//Used to Swipe and like HotGirls if found.
func SwipeAndLike(girls []Girl, likedChannel chan Girl, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, girl := range girls {
		if girl.hot == true {
			fmt.Println("::..Liked One..:: **licklick")
			likedChannel <- girl
		}
	}
	close(likedChannel)
}

//Used to filter our matches from the girls we liked.
func FindMatches(likedChannel chan Girl, matchedChannel chan Girl, wg *sync.WaitGroup) {
	defer wg.Done()
	for girl := range likedChannel {
		if girl.match == true {
			fmt.Println("..::New Match!::..")
			matchedChannel <- girl
		}
	}
	close(matchedChannel)
}

//Used to start a conversation with our matches.
func StartConversations(matchedChannel chan Girl, wg *sync.WaitGroup) {
	defer wg.Done()
	for _ = range matchedChannel {
		fmt.Println("..::Started Conversation::..")
	}
}

func RandomBool() bool {
	return rand.Float32() < 0.5
}

//Used as Mockdata and creates 30 random girls.
func InitGirls(amount int) []Girl {
	girls := make([]Girl, 0)
	for i := 0; i <= amount; i++ {
		g := Girl{RandomBool(), RandomBool()}
		girls = append(girls, g)
	}
	return girls
}
