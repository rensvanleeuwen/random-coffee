package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Pair struct {
	first  string
	second string
	last   bool
}

type PairCollection struct {
	pairs             []Pair
	firstColumnLength int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	users, err := readUniqueUsers("users.txt")
	if err != nil {
		os.Exit(1)
	}

	pairCollection := getPairCollection(users)

	for _, pair := range pairCollection.pairs {
		fmt.Printf("%"+strconv.Itoa(pairCollection.firstColumnLength)+"s ↔ %s\n", pair.first, pair.second)
	}
}

func getPairCollection(users []string) (pairCollection PairCollection) {
	for {
		pair := getPair(&users)
		pairCollection.pairs = append(pairCollection.pairs, pair)

		if len(pair.first) > pairCollection.firstColumnLength {
			pairCollection.firstColumnLength = len(pair.first)
		}

		if pair.last {
			break
		}
	}

	return
}

func getPair(users *[]string) (pair Pair) {
	if pair.first, pair.last = getUser(users); pair.last {
		// If this is the last user available and it is the first of a pair,
		// then there is no one to meet up with.
		pair.second = "¯\\_(ツ)_/¯"
		return
	}

	pair.second, pair.last = getUser(users)

	return
}

func getUser(users *[]string) (user string, last bool) {
	oldUsers := *users

	userIndex := rand.Intn(len(oldUsers))
	user = oldUsers[userIndex]

	// Swap the selected one with the one at the end and then
	// delete the last element (which is the selected user)
	oldUsers[userIndex] = oldUsers[len(oldUsers)-1]
	*users = oldUsers[:len(oldUsers)-1]

	last = len(*users) == 0

	return
}

func readUniqueUsers(filename string) (users []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("unable to open file %s\n", filename)
		return
	}

	defer file.Close()

	knownUsers := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		user := strings.TrimSpace(scanner.Text())

		if len(user) == 0 {
			continue
		}

		if _, known := knownUsers[user]; !known {
			knownUsers[user] = true
			users = append(users, user)
		}
	}

	return
}
