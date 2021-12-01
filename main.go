package main

import (
	"bufio"
	"log"

	//"crypto/sha256"
	"fmt"
	//"log"
	"os"
)

////TestContent implements the Content interface provided by merkletree and represents the content stored in the tree.
//type TestContent struct {
//	x string
//}
//
////CalculateHash hashes the values of a TestContent
//func (t TestContent) CalculateHash() ([]byte, error) {
//	h := sha256.New()
//	if _, err := h.Write([]byte(t.x)); err != nil {
//		return nil, err
//	}
//
//	return h.Sum(nil), nil
//}
//
////Equals tests for equality of two Contents
//func (t TestContent) Equals(other merkletree.Content) (bool, error) {
//	return t.x == other.(TestContent).x, nil
//}

func main() {
	var txId []string
	fmt.Println("トランザクションIDを入力してください")
	fmt.Println("Yを入力するとその時点でのマークルツリーを作成します。")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		in := scanner.Text()
		if in == "Y" {
			break
		}

		if len(in) != 64 {
			fmt.Println("不正な入力です。入力し直してください。")
		}

		txId = append(txId, in)
	}

	////Build list of Content to build tree
	//var list []merkletree.Content
	//list = append(list, TestContent{x: "Hello"})
	//list = append(list, TestContent{x: "Hi"})
	//list = append(list, TestContent{x: "Hey"})
	//list = append(list, TestContent{x: "Hola"})
	//
	////Create a new Merkle Tree from the list of Content
	t, err := merkletree.NewTree(txId)
	if err != nil {
		log.Fatal(err)
	}
	//
	////Get the Merkle Root of the tree
	//mr := t.MerkleRoot()
	//log.Println(mr)
	//
	////Verify the entire tree (hashes for each node) is valid
	//vt, err := t.VerifyTree()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("Verify Tree: ", vt)
	//
	////Verify a specific content in in the tree
	//vc, err := t.VerifyContent(list[0])
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("Verify Content: ", vc)
	//
	////String representation
	//log.Println(t)
}
