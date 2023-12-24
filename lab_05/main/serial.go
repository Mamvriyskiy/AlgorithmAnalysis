package main

func serialMtr(sliceMtr []matrixStruct, count, flag int) {
	if flag == 0 {
		el := sliceMtr[0]
		for i := 0; i < count; i++ {
			el.inverseMatrix = gaussMethodEx(el.matrixA)
			el.standartMul = mulStandart(el.inverseMatrix, el.matrixB, el.size)
			el.vinogradMul = mulVinogradov(el.inverseMatrix, el.matrixB, el.size)
		}
	} else {
		for i := range sliceMtr {
			sliceMtr[i].inverseMatrix = gaussMethodEx(sliceMtr[i].matrixA)
			sliceMtr[i].standartMul = mulStandart(sliceMtr[i].inverseMatrix, sliceMtr[i].matrixB, sliceMtr[i].size)
			sliceMtr[i].vinogradMul = mulVinogradov(sliceMtr[i].inverseMatrix, sliceMtr[i].matrixB, sliceMtr[i].size)
		}
	}
}
