package main

import (
	"database/sql"
	"encoding/gob"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"time"
	_ "github.com/go-sql-driver/mysql"
	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/Rhymen/go-whatsapp"
	"github.com/Rhymen/go-whatsapp/binary/proto"
)

type parameters struct {
	mysqlURL string
	whatsappSessionFile string
	whatsappGroup string
	piwigoImageFolder string
	piwigoBaseURL string
}

func loadParameters() (parameters) {
	param := new(parameters)
	flag.StringVar(&param.mysqlURL, "mysql-url", "", "The full url of the MySQL server to connect to")
	flag.StringVar(&param.whatsappSessionFile, "whatsapp-session-file", "", "The file to save the WhatsApp session to")
	flag.StringVar(&param.whatsappGroup, "whatsapp-group", "", "The ID of the WhatsApp group to send the message to")
	flag.StringVar(&param.piwigoImageFolder, "piwigo-image-folder", "", "The folder where the Piwigo images are stored")
	flag.StringVar(&param.piwigoBaseURL, "piwigo-base-url", "", "The base url of the Piwigo server")
	flag.Parse()
	return *param
}

func login(wac *whatsapp.Conn, sessionFile string) error {
	// Load saved session
	session, err := readSession(sessionFile)
	if err == nil {
		// Restore session
		session, err = wac.RestoreWithSession(session)
		if err != nil {
			return fmt.Errorf("restoring failed: %v", err)
		}
	} else {
		// No saved session -> regular login
		qr := make(chan string)
		go func() {
			terminal := qrcodeTerminal.New()
			terminal.Get(<-qr).Print()
		}()
		session, err = wac.Login(qr)
		if err != nil {
			return fmt.Errorf("error during login: %v", err)
		}
	}

	// Save session
	err = writeSession(session, sessionFile)
	if err != nil {
		return fmt.Errorf("error saving session: %v", err)
	}
	return nil
}

func readSession(sessionFile string) (whatsapp.Session, error) {
	session := whatsapp.Session{}
	file, err := os.Open(sessionFile)
	if err != nil {
		return session, err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&session)
	if err != nil {
		return session, err
	}
	return session, nil
}

func writeSession(session whatsapp.Session, sessionFile string) error {
	file, err := os.Create(sessionFile)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(session)
	if err != nil {
		return err
	}
	return nil
}

func sendMessage(wac *whatsapp.Conn, group string, message string, title string, thumbnail []byte) error {
	ts := uint64(time.Now().Unix())
	status := proto.WebMessageInfo_PENDING
	b := make([]byte, 10)
	rand.Read(b)
	id := strings.ToUpper(hex.EncodeToString(b))
	fromMe := true
	msg := &proto.WebMessageInfo{
		Key: &proto.MessageKey{
			FromMe: &fromMe,
			Id: &id,
			RemoteJid: &group,
		},
		MessageTimestamp: &ts,
		Message: &proto.Message{
			ExtendedTextMessage: &proto.ExtendedTextMessage{
				Title: &title,
				Text: &message,
				JpegThumbnail: thumbnail,
			},
		},
		Status: &status,
	}
	msgId, err := wac.Send(msg)
	if err != nil {
		return fmt.Errorf("Error sending message with title '%s': %v", title, err)
	}

	fmt.Fprintf(os.Stdout, "Message with title '%s' and id '%d' sent\n", title, msgId)
	return nil
}

func testConnexions(param parameters) error {
	// Create new WhatsApp connection and connect
	wac, err := whatsapp.NewConn(20 * time.Second)
	if err != nil {
		return fmt.Errorf("Error creating connection to WhatsApp: %v", err)
	}
	err = login(wac, param.whatsappSessionFile)
	if err != nil {
		return fmt.Errorf("Error logging in WhatsApp: %v", err)
	}
	<-time.After(3 * time.Second)
	defer wac.Disconnect()

	// Prints the available groups if none provided
	if param.whatsappGroup == "" {
		fmt.Fprintf(os.Stdout, "No WhatsApp group provided, showing all available groups\n")
		for _, chatNode := range wac.Store.Chats {
			fmt.Fprintf(os.Stdout, "%s | %s\n", chatNode.Jid, chatNode.Name)
		}

		return fmt.Errorf("No WhatsApp group provided")
	} else {
		_, exists := wac.Store.Chats[param.whatsappGroup]
		if (!exists) {
			return fmt.Errorf("Unknown WhatsApp group %s", param.whatsappGroup)
		}
	}

	// Connect to MySQL and execute a test query
	db, err := sql.Open("mysql", param.mysqlURL + "?parseTime=true")
	if err != nil {
		return fmt.Errorf("Error connecting to MySQL: %v", err)
	}
	defer db.Close()
	results, err := db.Query("SELECT Version();")
	if err != nil {
		return fmt.Errorf("Error executing MySQL query: %v", err)
	}
	defer results.Close()

	// Check the existence of the piwigo thumbnails directory
	_, err = os.Stat(fmt.Sprintf("%s/galleries", param.piwigoImageFolder))
	if err != nil {
		return fmt.Errorf("Could not find Piwigo thumbnail directory: %v", err)
	}

	return nil
}

func runLoop(param parameters) error {
	// Create new WhatsApp connection and connect
	wac, err := whatsapp.NewConn(20 * time.Second)
	if err != nil {
		return fmt.Errorf("Error creating connection to WhatsApp: %v", err)
	}
	err = login(wac, param.whatsappSessionFile)
	if err != nil {
		time.Sleep(30 * time.Second)
		err = login(wac, param.whatsappSessionFile)
		if err != nil {
			return fmt.Errorf("Error logging in WhatsApp: %v", err)
		}
	}
	<-time.After(3 * time.Second)
	defer wac.Disconnect()

	// Connect to MySQL and execute the query
	db, err := sql.Open("mysql", param.mysqlURL + "?parseTime=true")
	if err != nil {
		return fmt.Errorf("Error connecting to MySQL: %v", err)
	}
	defer db.Close()
	results, err := db.Query("SELECT piwigo_sharealbum.code, piwigo_categories.name, representatives.path, representatives.representative_ext, MIN(piwigo_images.date_creation) AS date_creation FROM piwigo_sharealbum JOIN piwigo_categories ON piwigo_sharealbum.cat = piwigo_categories.id JOIN piwigo_image_category ON piwigo_image_category.category_id = piwigo_categories.id JOIN piwigo_images ON piwigo_image_category.image_id = piwigo_images.id JOIN piwigo_images AS representatives ON piwigo_categories.representative_picture_id = representatives.id WHERE piwigo_images.date_creation IS NOT NULL GROUP BY piwigo_categories.id;")
	if err != nil {
		return fmt.Errorf("Error executing MySQL query: %v", err)
	}
	defer results.Close()

	var albumCode string
	var albumName string
	var representativePath string
	var representativeExt sql.NullString
	var albumDate time.Time
	for results.Next() {
		err = results.Scan(&albumCode, &albumName, &representativePath, &representativeExt, &albumDate)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error retrieving MySQL results for album '%s': %v\n", albumName, err)
			continue
		}

		if (albumDate.Month() == time.Now().Month()) && (albumDate.Day() == time.Now().Day()) {
			// Prepare the message
			url := fmt.Sprintf("%s/?xauth=%s", param.piwigoBaseURL, albumCode)
			imagePath := representativePath
			if (strings.HasPrefix(imagePath, "./")) {
				imagePath = imagePath[2:len(imagePath)]
			}
			imageExt := imagePath[strings.LastIndex(imagePath, ".")+1:len(imagePath)]
			if representativeExt.Valid {
				imageExt = representativeExt.String
				imagePath = imagePath[0:strings.LastIndex(imagePath, "/")] + "/pwg_representative" + imagePath[strings.LastIndex(imagePath, "/"):len(imagePath)]
			}
			imagePath = imagePath[0:strings.LastIndex(imagePath, ".")] + "-th." + imageExt
			thumbnail, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", param.piwigoImageFolder, imagePath))
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading thumbnail for album '%s': %v\n", albumName, err)
				continue
			}

			// Send the message
			sendMessage(wac, param.whatsappGroup, fmt.Sprintf("Il y a %d an(s) : %s", time.Now().Year()-albumDate.Year(), url), albumName, thumbnail)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error sending message to WhatsApp for album '%s': %v\n", albumName, err)
				continue
			}
		}
	}

	return nil
}

func main() {
	// Handle interrupts to clean properly
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		select {
			case sig := <-c:
				fmt.Printf("Got %s signal. Aborting...\n", sig)
				os.Exit(1)
		}
	}()

	// Load the parameters
	param := loadParameters()

	// Test the connexion on startup
	err := testConnexions(param)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect: %v\n", err)
		return
	}

	// Run the loop everyday at 7
	for {
		t := time.Now()
		n := time.Date(t.Year(), t.Month(), t.Day(), 7, 0, 0, 0, t.Location())
		d := n.Sub(t)
		if d < 0 {
			n = n.Add(24 * time.Hour)
			d = n.Sub(t)
		}
		fmt.Fprintf(os.Stderr, "Sleeping for: %s\n", d)
		time.Sleep(d)

		err := runLoop(param)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			continue
		}
	}
}
