package  main 


func SelectionSort(ar []int){

	for i:=0;i<len(ar);i++{
		val:=i
		for j:=i+1;j<len(ar);j++{
			if ar[val]>ar[j]{
				val=j
			}
		}
		ar[i],ar[val]=ar[val],ar[i]
	}

}