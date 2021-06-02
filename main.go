package main

import (
	"fmt"
	"os"
)

func init(){
	checkKey()
}




func main(){
	fmt.Println(" #####         ######         #######                     ")
	fmt.Println("#     #  ####  #     # #####     #    #####  ###### ######")
	fmt.Println("#       #    # #     # #    #    #    #    # #      #     ")
	fmt.Println("#  #### #    # ######  #    #    #    #    # #####  ##### ")
	fmt.Println("#     # #    # #     # #####     #    #####  #      #     ")
	fmt.Println("#     # #    # #     # #         #    #   #  #      #     ")
	fmt.Println(" #####   ####  ######  #         #    #    # ###### ######")
	op := 0
	tree := NewTree()

	for{
		fmt.Println("input your choice:")
		fmt.Println("1.create a BpTree")
		fmt.Println("2.load BpTree from disk")
		fmt.Println("3.query value from BpTree")
		fmt.Println("4.add node in the tree")
		fmt.Println("5.delete node in BpTree")
		fmt.Println("6.show the BpTree")
		fmt.Println("7.save the BpTree to file")
		fmt.Println("8.exit")
		fmt.Scanf("%d",&op)
		switch op {
		case 1:
			//fmt.Println("please input the key and value")
			CreateTree(tree)
			break
		case 2:
			tree =LoadTreeFromFile()
			break
		case 3:
			QueryValue(tree)
			break
		case 4:
			AddNode(tree)
			break
		case 5:
			DeleteValue(tree)
			break
		case 6:
			tree.PrintTree()
			break
		case 7:
			SaveFile()
			break
		case 8:
			os.Exit(0)
		default:
			fmt.Println("u enter the wrong choice")

		}

	}

}
