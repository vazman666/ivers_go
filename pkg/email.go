package pkg

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"strings"
)

func Email(file string) {

	var (
		//serverAddr = "smtp.mail.ru"
		serverAddr = "smtp.jino.ru"
		//password   = "rjV,brjhv1"
		password = "rbhgbxb1"
		//emailAddr  = "japautozap@mail.ru"
		emailAddr = "vazman@vazman.ru"
		//portNumber = 465
		portNumber = 465
		tos        = []string{}
		cc         = []string{
			"",
		}
		attachmentFilePath = file
		filename           = file
		delimeter          = "**=myohmy689407924327"
	)
	tos = append(tos, "japautozap@mail.ru")

	//fmt.Printf("File %v sender %v pass %v receiver %v\n", attachmentFilePath, emailAddr, password, tos)

	/*log.Println("======= Test Gmail client (with attachment) =========")
	log.Println("NOTE: user need to turn on 'less secure apps' options")
	log.Println("URL:  https://myaccount.google.com/lesssecureapps\n\r")*/

	tlsConfig := tls.Config{
		ServerName:         serverAddr,
		InsecureSkipVerify: true,
	}

	/*log.Println("Establish TLS connection")*/
	conn, connErr := tls.Dial("tcp", fmt.Sprintf("%s:%d", serverAddr, portNumber), &tlsConfig)
	if connErr != nil {
		log.Panic(connErr)
	}
	defer conn.Close()

	//log.Println("create new email client")
	client, clientErr := smtp.NewClient(conn, serverAddr)
	if clientErr != nil {
		log.Panic(clientErr)
	}
	defer client.Close()

	//log.Println("setup authenticate credential")
	auth := smtp.PlainAuth("", emailAddr, password, serverAddr)

	if err := client.Auth(auth); err != nil {
		log.Panic(err)
	}

	//log.Println("Start write mail content")
	//log.Println("Set 'FROM'")
	if err := client.Mail(emailAddr); err != nil {
		log.Panic(err)
	}
	//log.Println("Set 'TO(s)'")
	for _, to := range tos {
		if err := client.Rcpt(to); err != nil {
			log.Panic(err)
		}
	}

	writer, writerErr := client.Data()
	if writerErr != nil {
		log.Panic(writerErr)
	}

	//basic email headers
	sampleMsg := fmt.Sprintf("From: %s\r\n", emailAddr)
	sampleMsg += fmt.Sprintf("To: %s\r\n", strings.Join(tos, ";"))
	if len(cc) > 0 {
		sampleMsg += fmt.Sprintf("Cc: %s\r\n", strings.Join(cc, ";"))
	}
	sampleMsg += "Subject:Накладная IVERS\r\n"

	//log.Println("Mark content to accept multiple contents")
	sampleMsg += "MIME-Version: 1.0\r\n"
	sampleMsg += fmt.Sprintf("Content-Type: multipart/mixed; boundary=\"%s\"\r\n", delimeter)

	//place HTML message
	//log.Println("Put HTML message")
	sampleMsg += fmt.Sprintf("\r\n--%s\r\n", delimeter)
	sampleMsg += "Content-Type: text/html; charset=\"utf-8\"\r\n"
	sampleMsg += "Content-Transfer-Encoding: 7bit\r\n"
	/*sampleMsg += fmt.Sprintf("\r\n%s", "<html><body><h1>Hi There</h1>"+
	"<p>this is sample email (with attachment) sent via golang program</p></body></html>\r\n")*/

	//place file
	//log.Println("Put file attachment")
	sampleMsg += fmt.Sprintf("\r\n--%s\r\n", delimeter)
	sampleMsg += "Content-Type: text/plain; charset=\"utf-8\"\r\n"
	sampleMsg += "Content-Transfer-Encoding: base64\r\n"
	sampleMsg += "Content-Disposition: attachment;filename=\"" + filename + "\"\r\n"
	//read file
	rawFile, fileErr := ioutil.ReadFile(attachmentFilePath)
	if fileErr != nil {
		fmt.Printf("Нет файла %v   --письмо не отправлено!\n", attachmentFilePath)
		return //log.Panic(fileErr)
	}
	sampleMsg += "\r\n" + base64.StdEncoding.EncodeToString(rawFile)

	//write into email client stream writter
	//log.Println("Write content into client writter I/O")
	if _, err := writer.Write([]byte(sampleMsg)); err != nil {
		log.Panic(err)
	}

	if closeErr := writer.Close(); closeErr != nil {
		log.Panic(closeErr)
	}

	client.Quit()

	//log.Print("done.")
	//fmt.Printf("File send \n")
}
