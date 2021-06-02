package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)
var dist = make(map[string]string)
func CreateTree(tree *Tree){

	key:=0
	str :=""

	fmt.Println("enter the -1 to end")
	for{
		fmt.Println("please input the key")
		fmt.Scanf("%d",&key)
		if key==-1{
			break
		}
		fmt.Println("please input the value")
		fmt.Scanf("%s", &str)
		dist[strconv.Itoa(key)] =str

		value :=[]byte(str)
		err := tree.Insert(key, value)

		if err != nil {
			fmt.Errorf("%s", err)
		}

		r, err := tree.Find(key, false)
		if err != nil {
			fmt.Errorf("%s\n", err)
		}

		if r == nil {
			fmt.Errorf("returned nil \n")
		}

		if !reflect.DeepEqual(r.Value, value) {
			fmt.Errorf("expected %v and got %v \n", value, r.Value)
		}
	}
	tree.PrintTree()
	//fmt.Println(distValue)
	//fmt.Println(distKey)
}

func AddNode(tree *Tree){

	key:=0
	str :=""

	fmt.Println("enter the -1 to end")
	for{
		fmt.Println("please input the key")
		fmt.Scanf("%d",&key)
		if key==-1{
			break
		}
		fmt.Println("please input the value")
		fmt.Scanf("%s", &str)
		dist[strconv.Itoa(key)] =str

		value :=[]byte(str)
		err := tree.Insert(key, value)

		if err != nil {
			fmt.Errorf("%s", err)
		}

		r, err := tree.Find(key, false)
		if err != nil {
			fmt.Errorf("%s\n", err)
		}

		if r == nil {
			fmt.Errorf("returned nil \n")
		}

		if !reflect.DeepEqual(r.Value, value) {
			fmt.Errorf("expected %v and got %v \n", value, r.Value)
		}
	}
	tree.PrintTree()
	//fmt.Println(distValue)
	//fmt.Println(distKey)
}

func QueryValue(tree *Tree){
	key:=-2
	for{
		fmt.Println("please input the key to query, input -1 to end")
		fmt.Scanf("%d",&key)
		if key==-1{
			return
		}
		r, err := tree.Find(key, false)
		fmt.Printf("the value is %s \n",r.Value)
		if err != nil {
			fmt.Errorf("%s\n", err)
		}
		if r == nil {
			fmt.Errorf("returned nil, nothing find \n")
		}

	}

}
func DeleteValue(tree *Tree){
	key:=-2
	for{
		fmt.Println("please input the key to delete, input -1 to end")
		fmt.Scanf("%d",&key)
		if key==-1{
			return
		}
		r,err := tree.Find(key, false)
		fmt.Printf("the value is %s \n",r.Value)
		if err != nil {
			fmt.Errorf("%s\n", err)
		}
		if r == nil {
			fmt.Errorf("returned nil, nothing find \n")
		}
		err =tree.Delete(key)
		if err!=nil{
			fmt.Println("delete node failed")
		}
		delete(dist,strconv.Itoa(key))

	}

}

func SaveFile(){
	fmt.Println("input the file name you want to save")
	filename := ""
	fmt.Scanf("%s",&filename)
	Location:= "./tmp/"+filename+".db"
	f, err := os.Create(Location)
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
		return
	}

	w := bufio.NewWriter(f)
	for k, v := range dist {
		lineStr := fmt.Sprintf("%s:%s", k, v)
		fmt.Fprintln(w, lineStr)
	}

	w.Flush()
	f.Close()
	encryptFile(Location,Location+".enc")
	err1 := os.Remove(Location)

	if err1 != nil {
		// 删除失败
		fmt.Println("fail to delete the file")

	} else {
		// 删除成功
		fmt.Println("finish to save the BpTree")
	}

}

func LoadTreeFromFile() *Tree{
	tree := NewTree()

	fmt.Println("input the file name you want to load")
	filename := ""
	fmt.Scanf("%s",&filename)
	Location:= "./tmp/"+filename+".db.enc"
	decryptFile(Location,Location+".dec")
	dist = make(map[string]string)//清空map

	f, err := os.Open(Location+".dec")
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
		return nil
	}

	buf := bufio.NewReader(f)
	for {
		//遇到\n结束读取
		b, errR := buf.ReadBytes('\n')
		if errR != nil {
			if errR == io.EOF {
				break
			}
			fmt.Println(errR.Error())
		}
		//fmt.Println(string(b))
		line := string(b)
		result :=strings.Split(line,":")
		tmp :=result[0]
		value := strings.Replace(result[1],"\n","",-1)
		dist[tmp] =value
	}
	f.Close()
	for k, v := range dist {
		keyNum,_ := strconv.Atoi(k)
		tree.Insert(keyNum, []byte(v))
	}
	err1 := os.Remove(Location+".dec")
	if err1!=nil{
		fmt.Println("fail to delete the file")
	}

	return tree

}