package test

const (
	// 逻辑与
	LogicAnd = 1

	// 逻辑或
	LogicOr = 2

	// 逻辑非
	LogicNot = 3
)

type jumpTable struct {
	// 逻辑与跳转表
	LogicAndTable []int

	// 逻辑或跳转表
	LogicOrTable []int

	// 逻辑非跳转表
	LogicNotTable []int

	// 逻辑与跳转表
	LogicAndKeyWords map[string]int

	// 逻辑或跳转表
	LogicOrKeyWords map[string]int

	// 逻辑非跳转表
	LogicNotKeyWords map[string]int
}

		if found {
			indexedDoc := types.IndexedDocument{}
			indexedDoc.DocId = baseDocId
			if !countDocsOnly {
				docs = append(docs, indexedDoc)
			}
			numDocs++
		}