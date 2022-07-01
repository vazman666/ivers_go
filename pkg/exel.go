package pkg

import (
	"fmt"
	"ivers_api/models"
	"strconv"
	"strings"
	"time"

	"github.com/tealeg/xlsx/v3"
)

func Exel(data []models.Sql) {
	type xl struct {
		Firm       string
		PartNumber string
		Name       string
		Price      float32
		Quantity   int
		Summ       float32
		Remark     string
		//ACTT       string
	}
	var xls []xl
	var pul []int
	last := 0
	for _, i := range data {
		var tmp xl
		tmp.Firm = i.Firm
		tmp.PartNumber = i.Number
		tmp.Name = i.Name
		if s, err := strconv.ParseFloat(i.Price, 32); err == nil {
			tmp.Price = float32(s)
		}
		if s, err := strconv.Atoi(i.Quantity); err == nil {
			tmp.Quantity = s
		}

		tmp.Summ = tmp.Price * float32(tmp.Quantity)
		tmp.Remark = i.Remark
		//tmp.ACTT = i.ACTT
		xls = append(xls, tmp)
	}
	wb := xlsx.NewFile() //создаём новый экскиз экселя

	sheetTest, err := wb.AddSheet("Sheet") //добавляем страничку
	if err != nil {
		panic(err)
	}
	sheetTest.SetColWidth(1, 1, 16)
	sheetTest.SetColWidth(2, 2, 16)
	sheetTest.SetColWidth(3, 3, 55)
	sheetTest.SetColWidth(4, 6, 13)
	sheetTest.SetColWidth(7, 8, 20)
	//sheetTest.SetColWidth(11, 11, 12)
	//sheetTest.SetColWidth(12, 12, 1)

	row1 := sheetTest.AddRow()
	last++
	row1.SetHeight(15)
	cell := row1.AddCell()
	cell.Value = "Firm"
	cell = row1.AddCell()
	cell.Value = "PartNumber"
	cell = row1.AddCell()
	cell.Value = "Name"
	cell = row1.AddCell()
	cell.Value = "Цена"
	cell = row1.AddCell()
	cell.Value = "Количество"
	cell = row1.AddCell()
	cell.Value = "Сумма"
	cell = row1.AddCell()
	cell.Value = "Примечание"
	/*cell = row1.AddCell()
	cell.Value = "ACTT"*/
	row1 = sheetTest.AddRow()
	last++
	row1.SetHeight(15)

	/*for i := 0; i < 6; i++ {
		_ = sheetTest.SetColAutoWidth(i, xlsx.DefaultAutoWidth)
	}*/
	for _, value := range xls {

		row1 := sheetTest.AddRow()
		last++                           //добавляем строку
		_ = row1.WriteStruct(&value, -1) //и вставляе в эту строку строку из прайс
		if strings.Contains(strings.ToLower(value.Remark), "пул") ||
			strings.Contains(strings.ToLower(value.Remark), "pul") ||
			strings.ToLower(value.Remark)=="r"{
			pul = append(pul, last)
		}

		row1.SetHeight(15)

	}
	row1 = sheetTest.AddRow()
	row1 = sheetTest.AddRow()
	cell = row1.AddCell()
	cell = row1.AddCell()
	cell = row1.AddCell()
	cell = row1.AddCell()
	cell = row1.AddCell()
	cell.Value = "Итого:"
	cell = row1.AddCell()
	//cell.Value = "Итого:"
	cell.SetFormula("=SUM(F3:F" + strconv.Itoa(last) + ")")
	row1 = sheetTest.AddRow()
	row1 = sheetTest.AddRow()
	cell = row1.AddCell()
	cell = row1.AddCell()
	cell = row1.AddCell()
	cell = row1.AddCell()
	cell = row1.AddCell()
	//cell = row1.AddCell()
	cell.Value = "Пулмарт:"
	formulaPul := ""
	if len(pul) != 0 {
		formulaPul = "=SUM(F"
		for _, i := range pul {
			formulaPul = formulaPul + strconv.Itoa(i) + "+F"
		}
		formulaPul = strings.Trim(formulaPul, "+F")
		formulaPul = formulaPul + ")"
	} else {
		formulaPul = "=0"
	}
	//fmt.Printf("Formula=%s", formulaPul)
	cell = row1.AddCell()
	cell.SetFormula(formulaPul)

	t := time.Now()
	today := fmt.Sprintf("%02d_%02d_%d", t.Day(), t.Month(), t.Year())

	filename := "/home/vazman/Yandex.Disk/Документы/ivers/ivers" + today + ".xlsx"
	err = wb.Save(filename)
	if err != nil {
		panic(err)
	}
	sheetTest.Close()
	Email(filename)

}
