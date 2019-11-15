package database

import (
	"encoding/base64"
	"path/filepath"

	"github.com/jaeles-project/jaeles/database/models"
	"github.com/jaeles-project/jaeles/libs"
)

func CleanRecords() {
	var rec []models.Record
	DB.Find(&rec)
	DB.Unscoped().Delete(&rec)
}

func ImportRecord(rec libs.Record) {
	rawOutput, _ := filepath.Abs(rec.RawOutput)
	ReqRaw := base64.StdEncoding.EncodeToString([]byte(rec.Request.Beautify))
	ResRaw := base64.StdEncoding.EncodeToString([]byte(rec.Response.Beautify))
	extraOutput := base64.StdEncoding.EncodeToString([]byte(rec.ExtraOutput))

	if rec.Sign.Info.Name == "" {
		rec.Sign.Info.Name = rec.Sign.ID
	}
	if rec.Sign.Info.Risk == "" {
		rec.Sign.Info.Risk = "Potential"
	}

	recObj := models.Record{
		ReqMethod:   rec.Request.Method,
		ReqURL:      rec.Request.URL,
		ReqRaw:      ReqRaw,
		ReqBody:     rec.Request.Body,
		ResLength:   rec.Response.Length,
		StatusCode:  rec.Response.StatusCode,
		ResTime:     rec.Response.ResponseTime,
		ResRaw:      ResRaw,
		RawFile:     rawOutput,
		ExtraOutput: extraOutput,
		Issues:      rec.Sign.ID,
		Risk:        rec.Sign.Info.Risk,
		ScanID:      rec.ScanID,
	}
	DB.Create(&recObj)

	// if dbObj := DB.Create(&recObj); dbObj.Error != nil != true {
	// 	fmt.Printf("[Error] something woringcreate record\n")
	// }
}
