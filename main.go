package main

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/json"
	"fmt"
	zmq "github.com/go-zeromq/zmq4"
	"io"
	"log"
	"reflect"
	"strconv"
	"strings"
)

var schemaMap = map[string]func() interface{}{
	"https://eddn.edcd.io/schemas/fcmaterials_capi/1":    func() interface{} { return &FCMaterialsMessage{} },
	"https://eddn.edcd.io/schemas/commodity/3":           func() interface{} { return &CommodityMessage{} },
	"https://eddn.edcd.io/schemas/journal/1":             func() interface{} { return &JournalMessage{} },
	"https://eddn.edcd.io/schemas/approachsettlement/1":  func() interface{} { return &ApproachSettlementMessage{} },
	"https://eddn.edcd.io/schemas/blackmarket/1":         func() interface{} { return &BlackMarketMessage{} },
	"https://eddn.edcd.io/schemas/fcmaterials_journal/1": func() interface{} { return &FCMaterialsJournalMessage{} },
	"https://eddn.edcd.io/schemas/dockinggranted/1":      func() interface{} { return &DockingGrantedMessage{} },
	"https://eddn.edcd.io/schemas/outfitting/2":          func() interface{} { return &OutfittingMessage{} },
	"https://eddn.edcd.io/schemas/navroute/1":            func() interface{} { return &NavRouteMessage{} },
	"https://eddn.edcd.io/schemas/fsssignaldiscovered/1": func() interface{} { return &FSSSignalDiscoveredMessage{} },
	"https://eddn.edcd.io/schemas/fssallbodiesfound/1":   func() interface{} { return &FSSAllBodiesFoundMessage{} },
	"https://eddn.edcd.io/schemas/scanbarycentre/1":      func() interface{} { return &ScanBaryCentreMessage{} },
	"https://eddn.edcd.io/schemas/dockingdenied/1":       func() interface{} { return &DockingDeniedMessage{} },
	"https://eddn.edcd.io/schemas/fssdiscoveryscan/1":    func() interface{} { return &FSSDiscoveryScanMessage{} },
	"https://eddn.edcd.io/schemas/codexentry/1":          func() interface{} { return &CodexEntryMessage{} },
	"https://eddn.edcd.io/schemas/shipyard/2":            func() interface{} { return &ShipyardMessage{} },
	"https://eddn.edcd.io/schemas/fssbodysignals/1":      func() interface{} { return &FSSBodySignalsMessage{} },
	"https://eddn.edcd.io/schemas/navbeaconscan/1":       func() interface{} { return &NavBeaconScanMessage{} },
	// Add more schemas as needed
}

type Mat struct {
	count int
	price int
}

// Sortable slice of Mat items
type MatList []struct {
	Name string
	Mat  Mat
}

// Implement sort.Interface for MatList based on the price field
func (ml MatList) Len() int           { return len(ml) }
func (ml MatList) Swap(i, j int)      { ml[i], ml[j] = ml[j], ml[i] }
func (ml MatList) Less(i, j int) bool { return ml[i].Mat.price < ml[j].Mat.price }

func main() {
	sub := zmq.NewSub(context.Background())
	defer sub.Close()

	err := sub.Dial("tcp://eddn.edcd.io:9500")
	if err != nil {
		log.Fatal(err)
	}

	sub.SetOption(zmq.OptionSubscribe, "")

	fmt.Println("Listening for EDDN messages...")

	msgChan := make(chan []byte)

	go func() {
		for {
			msg, err := sub.Recv()
			if err != nil {
				log.Fatal(err)
			}
			msgChan <- msg.Frames[0]
		}
	}()

	for msg := range msgChan {
		decompressedMsg, err := decompressZlib(msg)
		if err != nil {
			log.Printf("Error decompressing message: %v\n", err)
			continue
		}

		//fmt.Printf("Raw decompressed message: %s\n", string(decompressedMsg))

		if !json.Valid(decompressedMsg) {
			log.Printf("Invalid JSON received after decompression: %s\n", string(decompressedMsg))
			continue
		}

		var eddnMsg EDDN
		if err := json.Unmarshal(decompressedMsg, &eddnMsg); err != nil {
			log.Printf("Error parsing EDDN JSON: %v\n", err)
			//	log.Printf("Decompressed Message: %s\n", string(decompressedMsg))
			continue
		}
		//fmt.Printf("Parsed EDDN Message: %+v\n", eddnMsg)

		schemaFunc, exists := schemaMap[eddnMsg.SchemaRef]
		if !exists {
			log.Printf("Unknown schema: %s\n", eddnMsg.SchemaRef)
			continue
		}

		specificMsg := schemaFunc()

		if err := json.Unmarshal(eddnMsg.Message, specificMsg); err != nil {
			log.Printf("Error parsing specific message for schema %s: %v\n", eddnMsg.SchemaRef, err)
			//	log.Printf("Message content: %s\n", string(eddnMsg.Message))
			continue
		}

		// Process the specific message type
		switch v := specificMsg.(type) {
		case *NavBeaconScanMessage:
			//fmt.Printf("Processing FSSDiscoveryScanMessage: %+v\n", v)
		case *FSSBodySignalsMessage:
			/*var genericMap map[string]interface{}
			if err := json.Unmarshal(decompressedMsg, &genericMap); err != nil {
				log.Printf("Error unmarshalling into map: %v\n", err)
			} else {
				fmt.Printf("Unmarshalled map: %+v\n", genericMap)
			}*/
			//fmt.Printf("Processing FSSDiscoveryScanMessage: %+v\n", v)
		case *ShipyardMessage:
			//fmt.Printf("Processing FSSDiscoveryScanMessage: %+v\n", v)
		case *CodexEntryMessage:
			//fmt.Printf("Processing FSSDiscoveryScanMessage: %+v\n", v)
		case *FSSDiscoveryScanMessage:
			//fmt.Printf("Processing FSSDiscoveryScanMessage: %+v\n", v)
		case *DockingDeniedMessage:
			//fmt.Printf("Processing DockingDeniedMessage: %+v\n", v)
			// Handle DockingDeniedMessage
		case *ScanBaryCentreMessage:
			//fmt.Printf("Processing ScanBaryCentreMessage: %+v\n", v)
			// Handle ScanBaryCentreMessage
		case *FSSAllBodiesFoundMessage:
			//fmt.Printf("Processing FSSAllBodiesFoundMessage: %+v\n", v)
			// Handle FSSAllBodiesFoundMessage
		case *FSSSignalDiscoveredMessage:
			//fmt.Printf("Processing FSSSignalDiscoveredMessage: %+v\n", v)
			// Handle FSSSignalDiscoveredMessage
			for _, FSSSignal := range v.Signals {
				if strings.Contains(FSSSignal.SignalType, "High") || strings.Contains(FSSSignal.SignalType, "grade") {
					println("yay")
				}
			}

		case *NavRouteMessage:
			//fmt.Printf("Processing NavRouteMessage: %+v\n", v)
			// Handle NavRouteMessage
		case *OutfittingMessage:
			//fmt.Printf("Processing OutfittingMessage: %+v\n", v)
			// Handle OutfittingMessage
		case *FCMaterialsJournalMessage:
			//fmt.Printf("Processing FCMaterialsJournalMessage: %+v\n", v)
			// Handle FCMaterialsJournalMessage
		case *DockingGrantedMessage:
			//fmt.Printf("Processing DockingGrantedMessage: %+v\n", v)
			// Handle DockingGrantedMessage
		case *FCMaterialsMessage:
			//fmt.Printf("Processing FCMaterialsMessage: %+v\n", v)
			// Handle FCMaterialsMessage
		case *CommodityMessage:
			/*var genericMap map[string]interface{}
			if err := json.Unmarshal(decompressedMsg, &genericMap); err != nil {
				log.Printf("Error unmarshalling into map: %v\n", err)
			} else {
				fmt.Printf("Unmarshalled map: %+v\n", genericMap)
			}*/

		case *JournalMessage:
		/*	if v.Event != "Docked" {
			fmt.Printf("Processing JournalMessage: %+v\n", v)
			var genericMap map[string]interface{}
			if err := json.Unmarshal(decompressedMsg, &genericMap); err != nil {
				log.Printf("Error unmarshalling into map: %v\n", err)
			} else {
				fmt.Printf("Unmarshalled map: %+v\n", genericMap)
			}

		}*/

		// Handle JournalMessage
		case *ApproachSettlementMessage:
			//fmt.Printf("Processing ApproachSettlementMessage: %+v\n", v)
			// Handle ApproachSettlementMessage
		case *BlackMarketMessage:
			//fmt.Printf("Processing BlackMarketMessage: %+v\n", v)
			// Handle BlackMarketMessage
		// Add cases for other schemas
		default:
			log.Printf("Unhandled message type for schema: %s\n", eddnMsg.SchemaRef)
		}
	}
}

// decompressZlib decompresses the zlib-compressed message.
func decompressZlib(data []byte) ([]byte, error) {
	reader, err := zlib.NewReader(io.NopCloser(bytes.NewReader(data)))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	decompressedData, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return decompressedData, nil
}

func formatCurrency(amount int) string {
	// Convert the integer to a string with commas
	str := strconv.FormatInt(int64(amount), 10)
	n := len(str)
	if n <= 3 {
		return "$" + str
	}

	// Add commas every three digits
	result := ""
	for i, c := range str {
		if (n-i)%3 == 0 && i != 0 {
			result += ","
		}
		result += string(c)
	}

	return "$" + result
}

func printMapWithTypes(data map[string]interface{}) {
	for key, value := range data {
		v := reflect.ValueOf(value)
		fmt.Printf("Key: %s, Type: %s, Value: %v\n", key, v.Type(), value)

		// If the value is another map or slice, recurse into it
		switch v.Kind() {
		case reflect.Map:
			fmt.Printf("Entering nested map at key: %s\n", key)
			nestedMap, ok := value.(map[string]interface{})
			if ok {
				printMapWithTypes(nestedMap)
			}
		case reflect.Slice:
			fmt.Printf("Entering slice at key: %s\n", key)
			slice, ok := value.([]interface{})
			if ok {
				for i, item := range slice {
					fmt.Printf("Index: %d, Type: %s, Value: %v\n", i, reflect.TypeOf(item), item)
					if reflect.TypeOf(item).Kind() == reflect.Map {
						nestedMap, ok := item.(map[string]interface{})
						if ok {
							printMapWithTypes(nestedMap)
						}
					}
				}
			}
		}
	}
}
