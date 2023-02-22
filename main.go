package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	origin := "./tutti/all.png"
	width := 512
	height := 2856
	descr := fmt.Sprintf("Tarocco Piemontese da 14 carte per segno - %s", time.Now().Format("2006-01-02 15:04:05"))

	if err := generatePac(origin, descr, width, height); err != nil {
		log.Fatal("Err: ", err)
	}
	log.Println("That's all folks!")
}

func generatePac(fnameImg, descr string, width, height int) error {
	finp, err := os.Open(fnameImg)
	if err != nil {
		return err
	}
	defer finp.Close()

	bbPng, err := io.ReadAll(finp)
	if err != nil {
		return err
	}

	pac := &bytes.Buffer{}
	// primi 100 bytes sono la descrizione

	if len(descr) > 100 {
		descr = descr[0:100]
	}
	pac.Write([]byte(descr))
	for next := true; next; next = pac.Len() < 100 {
		pac.WriteByte(0)
	}
	log.Printf("Size of descr: Len %d, Cap %d", pac.Len(), pac.Cap())
	timestamp := time.Now().Unix()
	var ble32 [4]byte
	binary.LittleEndian.PutUint32(ble32[:], uint32(timestamp))
	pac.Write(ble32[:]) // Timestamp LE32 (4 bytes)
	pac.WriteByte(0x01) // number of animations

	var ble16 [2]byte
	binary.LittleEndian.PutUint16(ble16[:], uint16(width))
	pac.Write(ble16[:]) // png width (2 bytes)
	binary.LittleEndian.PutUint16(ble16[:], uint16(height))
	pac.Write(ble16[:]) // png height (2 bytes)

	pac.WriteByte(0x01) // frames Len LE 16 (2 bytes)
	pac.WriteByte(0x00) // frames Len is one word
	pac.WriteByte(0xff) // frames body LE 16 use oxffff and it is ignored in game
	pac.WriteByte(0xff)
	// PGN Image
	pac.Write(bbPng)
	outfname := "tarock_piemonte.pac"
	if err := os.WriteFile(outfname, pac.Bytes(), 644); err != nil {
		return err
	}
	log.Println("File created ", outfname)
	return nil
}
