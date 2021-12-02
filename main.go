package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

////TestContent implements the Content interface provided by merkletree and represents the content stored in the tree.
//type TestContent struct {
//	x string
//}
//
type Transaction struct {
	id string
}

////CalculateHash hashes the values of a TestContent
func (t Transaction) CalculateHash() ([]byte, error) {
	//h := sha256.New()
	//if _, err := h.Write([]byte(t.id)); err != nil {
	//	return nil, err
	//}

	//return h.Sum(nil), nil

	return []byte(t.id), nil
	//hex, _ := hex.DecodeString(t.id)
	//var i int64
	//// エンディアンの変換
	//binary.Read(bytes.NewReader(hex), binary.BigEndian, &i)
	//result := make([]byte, binary.MaxVarintLen32)
	//binary.PutVarint(result, i)
	//return result, nil
}

//Equals tests for equality of two Contents
func (t Transaction) Equals(other Content) (bool, error) {
	return t.id == other.(Transaction).id, nil
}

func parseBE(string2 string) string {
	result := ""

	for i := 0; i < 32; i++ {
		result = string2[i*2:i*2+2] + result
		//fmt.Println(string2)
		//fmt.Println(result)
	}

	return result
}

func main() {
	var txId []Content
	fmt.Println("トランザクションIDを入力してください")
	fmt.Println("Yを入力すると入力されたトランザクションIDからマークルツリーを作成します。")
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

		be := parseBE(in)
		//fmt.Println(be)
		txId = append(txId, Transaction{id: be})
	}

	t, err := NewTree(txId)
	if err != nil {
		log.Fatal(err)
	}
	mr := t.MerkleRoot()
	log.Println(mr)
	log.Println(fmt.Sprintf("%s", mr))
	log.Println(hex.EncodeToString(mr))

}
