package main

import (
	"bufio"
	_ "bufio"
	"crypto/sha256"
	_ "crypto/sha256"
	"fmt"
	_ "fmt"
	"math/rand"
	_ "math/rand"
	"os"
	_ "os"
	"strconv"
	_ "strconv"
	"strings"
)

type block struct {
	transaction  string
	nonce        int
	currentHash string
	previousHash string
}

type blockchain struct {
    list []*block
}

func NewBlock(transaction string, nonce int, currentHash string, previousHash string) *block {
	b := new(block)
	b.transaction = transaction
	b.nonce = nonce
	b.currentHash = currentHash
	b.previousHash = previousHash
	return b
}

func ChangeBlock(block blockchain, blockno int, transaction string, nonce int) {
	if transaction == "none" {
		block.list[blockno].nonce = nonce
	} else {
		block.list[blockno].transaction = transaction
	}
}

func CalculateHash (stringToHash string) string{
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
}

func ListBlocks(block blockchain){
	fmt.Println("The blockchain is as follows")
	for i := 0; i < len(block.list); i++ {
		fmt.Printf("\n%s Block ID : %v %s", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Printf("\nTransaction : %vNonce : %v\nCurrent Hash : %v\nPrevious Hash : %v\n", block.list[i].transaction, block.list[i].nonce, block.list[i].currentHash, block.list[i].previousHash)
		fmt.Printf("%s\n", strings.Repeat("=", 64))
	}
}

func VerifyChain(blockchain blockchain, count int){
	fmt.Println("Verifying chain...")
	if count == 0 {
		fmt.Print("No hashes to verify")
	} else {
		hashVerify := true
		hash := "null"
		for i := 0; i < count; i++ {
			// if i == count -1 {
			// 	break;
			// }
			if hash != blockchain.list[i].previousHash {
				fmt.Printf("\n%s Block ID : %v %s", strings.Repeat("=", 25), i - 1, strings.Repeat("=", 25))
				fmt.Printf("\nTransaction : %vNonce : %v\nPrevious Hash : %v\n", blockchain.list[i - 1].transaction, blockchain.list[i - 1].nonce, blockchain.list[i - 1].previousHash)
				fmt.Printf("%s\n", strings.Repeat("=", 64))
				hashVerify = false
				break
			}
			hash = CalculateHash(blockchain.list[i].transaction + strconv.Itoa(blockchain.list[i].nonce) + blockchain.list[i].previousHash)
		}
		if hashVerify {
			fmt.Print("The chain is valid!\n")
		} else {
			fmt.Print("Changes have been made to the chain!\n")
		}
	}
}

func main() {
	choice := 0
	// var transaction string
	var nonce int
	previousHash := "null"
    blockchain := new(blockchain)
	count := 0
    in := bufio.NewReader(os.Stdin)
    for {
		fmt.Println("Menu\n1) Add a new block\n2) Change Block\n3) Verify Chain\n4) List Blocks\n0) Exit")
		
		fmt.Scanln(&choice)
        if choice == 0 {
            break;
        }
		if choice > 4 || choice < 0 {
			for {
				fmt.Println("Invalid input! Please enter again : ")
				fmt.Scanln(&choice)
				if choice >= 0 && choice <= 4 {
					break;
				}
			}
		}
		if choice == 0 {
			break;
		} else if choice == 1 {
			fmt.Println("Enter the transaction to be added : ")
			transaction, err := in.ReadString('\n')
			_ = err
			// fmt.Print(transaction)
			nonce = rand.Intn(10000)
			var fullblock string
			fullblock += transaction + strconv.Itoa(nonce) + previousHash
			currentHash := CalculateHash(fullblock)
			blockchain.list = append(blockchain.list, NewBlock(transaction, nonce, currentHash, previousHash))
			// fmt.Println(blockchain.list[count])
			previousHash = CalculateHash(fullblock)
			count += 1
		} else if choice == 2 {
			fmt.Printf("The current number of blocks is %v \n", count)
			if count != 0 {
				var blockno int
				fmt.Printf("Enter the block number you would like to change : (0 - %v) ", count-1)
				fmt.Scanln(&blockno)
				if blockno < 0 || blockno >= count {
					for {
						fmt.Print("Invalid input! Enter again : ")
						fmt.Scanln(&blockno)
						if blockno >= 0 && blockno < count {
							break;
						}
					}
				}
				var changeNo int
				fmt.Printf("\n%s Block ID : %v %s", strings.Repeat("=", 25), blockno, strings.Repeat("=", 25))
				fmt.Printf("\nTransaction : %vNonce : %v\nPrevious Hash : %v\n", blockchain.list[blockno].transaction, blockchain.list[blockno].nonce, blockchain.list[blockno].previousHash)
				fmt.Printf("%s\n", strings.Repeat("=", 64))
				fmt.Print("What would you like to change?\n1) Transaction\n2) Nonce\n")
				fmt.Scanln(&changeNo)
				if changeNo < 1 || changeNo > 2 {
					for {
						fmt.Println("Invalid Input! Enter again : ")
						fmt.Scanln(&changeNo)
						if changeNo == 1 || changeNo == 2 {
							break;
						}
					}
				}
				if changeNo == 1 {
					var newTrans string
					fmt.Println("Enter new transaction : ")
					newTrans, err := in.ReadString('\n')
					_ = err
					// fmt.Scanln(&newTrans)
					ChangeBlock(*blockchain, blockno, newTrans, 0)
				} else if changeNo == 2 {
					var newNonce int
					fmt.Println("Enter new nonce : ")
					fmt.Scanln(&newNonce)
					ChangeBlock(*blockchain, blockno, "none", newNonce)
				}
				fmt.Printf("\n%s Block ID : %v %s", strings.Repeat("=", 25), blockno, strings.Repeat("=", 25))
				fmt.Printf("\nTransaction : %vNonce : %v\nPrevious Hash : %v\n", blockchain.list[blockno].transaction, blockchain.list[blockno].nonce, blockchain.list[blockno].previousHash)
				fmt.Printf("%s\n", strings.Repeat("=", 64))
			}
		} else if choice == 3 {
			VerifyChain(*blockchain, count)
		} else if choice == 4 {
			ListBlocks(*blockchain)
		} else {
			fmt.Println("Invalid input!")
		}

	}
}