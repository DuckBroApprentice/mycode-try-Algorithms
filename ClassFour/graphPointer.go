package classfour

type graphPointer struct {
	//可以指向任意個節點或者不指向
	//節點可能要用[]*graphPointer記錄
	//問題在於該怎麼找到其他節點
}

//存在多個資料，資料與資料間會有描述關係(若沒有任何描述則沒有鏈接)，比如二維陣列實現的graph即描述了資料間的關係
//晚點看看leetcode怎麼設計題目  先下課
